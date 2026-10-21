package pike

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// awsManagedPolicyDenylist is the set of AWS-managed policy names that grant
// admin-equivalent or near-admin access and should not be attached blindly.
var awsManagedPolicyDenylist = map[string]bool{
	"AdministratorAccess": true,
	"IAMFullAccess":       true,
	"PowerUserAccess":     true,
}

// awsPolicyAttachmentTypes are the resource types whose policy_arn may
// reference a customer-managed aws_iam_policy defined in the same scan —
// collectAuditFindings' attachment-graph pass looks for exactly these.
var awsPolicyAttachmentTypes = map[string]bool{
	"aws_iam_role_policy_attachment":  true,
	"aws_iam_user_policy_attachment":  true,
	"aws_iam_group_policy_attachment": true,
	"aws_iam_policy_attachment":       true,
}

func init() {
	for _, r := range []string{
		"aws_iam_policy", "aws_iam_role_policy",
		"aws_iam_user_policy", "aws_iam_group_policy",
	} {
		auditHandlers["resource."+r] = auditAWSInlinePolicy
	}
	auditHandlers["resource.aws_iam_role"] = auditAWSRoleInlinePolicies
	auditHandlers["data.aws_iam_policy_document"] = auditAWSPolicyDocument
	for r := range awsPolicyAttachmentTypes {
		auditHandlers["resource."+r] = auditAWSPolicyAttachment
	}
}

// ---- policy document model ----------------------------------------------

type stringOrSlice []string

func (s *stringOrSlice) UnmarshalJSON(data []byte) error {
	var one string
	if err := json.Unmarshal(data, &one); err == nil {
		*s = []string{one}
		return nil
	}
	var many []string
	if err := json.Unmarshal(data, &many); err != nil {
		return err
	}
	*s = many
	return nil
}

type awsStmt struct {
	Sid         string          `json:"Sid"`
	Effect      string          `json:"Effect"`
	Action      stringOrSlice   `json:"Action"`
	NotAction   stringOrSlice   `json:"NotAction"`
	Resource    stringOrSlice   `json:"Resource"`
	NotResource stringOrSlice   `json:"NotResource"`
	Condition   json.RawMessage `json:"Condition"`
	// Principal is only meaningful on a trust policy (assume_role_policy)
	// statement; it's absent/zero on an identity/resource policy statement.
	Principal awsPrincipal `json:"Principal"`
	// ConditionKeys holds the condition variable names (e.g.
	// "iam:PassedToService", "aws:PrincipalTag/team") extracted from
	// Condition, when its structure could be statically walked.
	ConditionKeys []string `json:"-"`
}

// awsPrincipal models the shapes a trust policy's Principal can take: the
// bare string "*", or an object with AWS/Service/Federated keys, each either
// a single string or a list.
type awsPrincipal struct {
	Wildcard  bool // Principal: "*", or Principal: {"AWS": "*"}
	AWS       []string
	Service   []string
	Federated []string
}

func (p *awsPrincipal) UnmarshalJSON(data []byte) error {
	var asString string
	if err := json.Unmarshal(data, &asString); err == nil {
		p.Wildcard = asString == "*"
		return nil
	}

	var asObject struct {
		AWS       stringOrSlice `json:"AWS"`
		Service   stringOrSlice `json:"Service"`
		Federated stringOrSlice `json:"Federated"`
	}
	if err := json.Unmarshal(data, &asObject); err != nil {
		return err
	}

	p.AWS = asObject.AWS
	p.Service = asObject.Service
	p.Federated = asObject.Federated
	p.Wildcard = slices.Contains(p.AWS, "*")

	return nil
}

// escalationScopingConditionKeys are AWS condition keys that genuinely
// narrow the blast radius of an escalation-class action (e.g. restricting
// iam:PassRole to a specific service, or an action to specific
// tagged/sourced principals/resources). A condition that only tests
// something unrelated (region, time, TLS, ...) does not make the grant
// safe, so it must not suppress AWS004.
var escalationScopingConditionKeys = []string{
	"iam:passedtoservice",
	"iam:awsservicename",
	"iam:permissionsboundary",
	"iam:policyarn",
	"iam:resourcetag",
	"aws:resourcetag",
	"aws:requesttag",
	"aws:principaltag",
	"aws:sourcearn",
	"aws:sourceaccount",
	"aws:resourceaccount",
}

func isScopingConditionKey(key string) bool {
	lk := strings.ToLower(key)
	for _, scoping := range escalationScopingConditionKeys {
		if lk == scoping || strings.HasPrefix(lk, scoping+"/") {
			return true
		}
	}
	return false
}

// hasScopingCondition reports whether the statement's condition actually
// restricts an escalation-class grant, as opposed to merely existing.
func (s awsStmt) hasScopingCondition() bool {
	return slices.ContainsFunc(s.ConditionKeys, isScopingConditionKey)
}

// hasAnyCondition reports whether the statement has a Condition block at
// all, regardless of which keys it tests.
func (s awsStmt) hasAnyCondition() bool {
	return len(s.ConditionKeys) > 0
}

// hasConditionKey reports whether any of the statement's condition keys
// case-insensitively match one of want (exact match, ignoring a trailing
// "/Tag"-style qualifier).
func (s awsStmt) hasConditionKey(want ...string) bool {
	return slices.ContainsFunc(s.ConditionKeys, func(key string) bool {
		lk := strings.ToLower(key)
		for _, w := range want {
			if lk == w {
				return true
			}
		}
		return false
	})
}

// hasSubOrAudCondition reports whether the statement scopes a federated/OIDC
// principal by a claim condition, e.g. "token.actions.githubusercontent.com:sub"
// or "accounts.google.com:aud" — any condition key ending in ":sub" or
// ":aud", regardless of which OIDC provider it names.
func (s awsStmt) hasSubOrAudCondition() bool {
	return slices.ContainsFunc(s.ConditionKeys, func(key string) bool {
		lk := strings.ToLower(key)
		return strings.HasSuffix(lk, ":sub") || strings.HasSuffix(lk, ":aud")
	})
}

// hasWildcardRootAccount reports whether principals contains an
// account-root ARN with a wildcard account segment, e.g.
// "arn:aws:iam::*:root" — which matches every account, not one specific one.
func hasWildcardRootAccount(principals []string) bool {
	return slices.ContainsFunc(principals, func(arn string) bool {
		return strings.Contains(arn, ":iam::*:root")
	})
}

// containsActionFold reports whether actions contains want, ignoring case —
// AWS action names are case-insensitive.
func containsActionFold(actions []string, want string) bool {
	return slices.ContainsFunc(actions, func(a string) bool {
		return strings.EqualFold(a, want)
	})
}

// conditionKeysFromRaw extracts condition variable names from a JSON
// Condition object of the form {"Operator": {"key": value, ...}, ...}.
func conditionKeysFromRaw(raw json.RawMessage) []string {
	if len(raw) == 0 {
		return nil
	}
	var ops map[string]json.RawMessage
	if err := json.Unmarshal(raw, &ops); err != nil {
		return nil
	}
	var keys []string
	for _, v := range ops {
		var kv map[string]json.RawMessage
		if err := json.Unmarshal(v, &kv); err != nil {
			continue
		}
		for k := range kv {
			keys = append(keys, k)
		}
	}
	return keys
}

type awsPolicyDoc struct {
	Statement []awsStmt
}

func (d *awsPolicyDoc) UnmarshalJSON(data []byte) error {
	var raw struct {
		Statement json.RawMessage `json:"Statement"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if len(raw.Statement) == 0 {
		return nil
	}
	if raw.Statement[0] == '[' {
		return json.Unmarshal(raw.Statement, &d.Statement)
	}
	var one awsStmt
	if err := json.Unmarshal(raw.Statement, &one); err != nil {
		return err
	}
	d.Statement = []awsStmt{one}
	return nil
}

// ---- extraction ----------------------------------------------------------

// extractAWSPolicy turns a `policy` attribute expression into an awsPolicyDoc.
// Returns (nil, reason) when the value cannot be statically analysed.
func extractAWSPolicy(expr hclsyntax.Expression, baseDir string) (*awsPolicyDoc, string) {
	switch e := expr.(type) {
	case *hclsyntax.FunctionCallExpr:
		switch e.Name {
		case "jsonencode":
			if len(e.Args) != 1 {
				return nil, "jsonencode() with unexpected arity"
			}
			obj, ok := e.Args[0].(*hclsyntax.ObjectConsExpr)
			if !ok {
				return nil, "jsonencode() argument is not an object literal"
			}
			return policyFromObjectExpr(obj), ""
		case "file", "templatefile":
			if len(e.Args) == 0 {
				return nil, e.Name + "() with no path argument"
			}
			path := extractStringValue(e.Args[0])
			if path == "" || strings.Contains(path, "${") {
				return nil, e.Name + "() path is not a literal"
			}
			full := filepath.Join(baseDir, path)
			data, err := os.ReadFile(full) // #nosec G304 -- path comes from user's own Terraform source
			if err != nil {
				return nil, fmt.Sprintf("cannot read %s: %v", path, err)
			}
			var doc awsPolicyDoc
			if err := json.Unmarshal(data, &doc); err != nil {
				return nil, fmt.Sprintf("%s is not valid policy JSON: %v", path, err)
			}
			populateConditionKeys(&doc)
			return &doc, ""
		default:
			return nil, fmt.Sprintf("policy is computed by %s()", e.Name)
		}
	case *hclsyntax.TemplateExpr:
		s := extractStringValue(e)
		if s == "" || strings.Contains(s, "${") {
			return nil, "policy heredoc contains interpolation"
		}
		var doc awsPolicyDoc
		if err := json.Unmarshal([]byte(s), &doc); err != nil {
			return nil, fmt.Sprintf("policy heredoc is not valid JSON: %v", err)
		}
		populateConditionKeys(&doc)
		return &doc, ""
	case *hclsyntax.ScopeTraversalExpr:
		return nil, fmt.Sprintf("policy references %s", extractStringValue(e))
	default:
		return nil, "policy expression form not supported"
	}
}

// populateConditionKeys fills in ConditionKeys for statements parsed
// straight out of JSON (file/templatefile/heredoc), where Condition already
// holds the real Condition object.
func populateConditionKeys(doc *awsPolicyDoc) {
	for i := range doc.Statement {
		doc.Statement[i].ConditionKeys = conditionKeysFromRaw(doc.Statement[i].Condition)
	}
}

// policyFromObjectExpr walks a jsonencode({...}) object literal and extracts
// the fields the rules care about. Non-literal leaves are dropped, so a
// statement with an interpolated Resource still has its literal Action checked.
func policyFromObjectExpr(obj *hclsyntax.ObjectConsExpr) *awsPolicyDoc {
	var doc awsPolicyDoc
	for _, item := range obj.Items {
		if !keyEquals(item.KeyExpr, "Statement") {
			continue
		}
		switch v := unwrapExpr(item.ValueExpr).(type) {
		case *hclsyntax.TupleConsExpr:
			for _, se := range v.Exprs {
				if so, ok := unwrapExpr(se).(*hclsyntax.ObjectConsExpr); ok {
					doc.Statement = append(doc.Statement, stmtFromObjectExpr(so))
				}
			}
		case *hclsyntax.ObjectConsExpr:
			doc.Statement = append(doc.Statement, stmtFromObjectExpr(v))
		}
	}
	return &doc
}

func stmtFromObjectExpr(obj *hclsyntax.ObjectConsExpr) awsStmt {
	var s awsStmt
	for _, item := range obj.Items {
		val := unwrapExpr(item.ValueExpr)
		switch {
		case keyEquals(item.KeyExpr, "Effect"):
			s.Effect = extractStringValue(val)
		case keyEquals(item.KeyExpr, "Sid"):
			s.Sid = extractStringValue(val)
		case keyEquals(item.KeyExpr, "Action"):
			s.Action = stringOrListExpr(val)
		case keyEquals(item.KeyExpr, "NotAction"):
			s.NotAction = stringOrListExpr(val)
		case keyEquals(item.KeyExpr, "Resource"):
			s.Resource = stringOrListExpr(val)
		case keyEquals(item.KeyExpr, "NotResource"):
			s.NotResource = stringOrListExpr(val)
		case keyEquals(item.KeyExpr, "Condition"):
			s.Condition = json.RawMessage(`{}`)
			s.ConditionKeys = conditionKeysFromObjectExpr(val)
		case keyEquals(item.KeyExpr, "Principal"):
			s.Principal = principalFromExpr(val)
		}
	}
	return s
}

// principalFromExpr walks a Principal = "*" or
// Principal = { AWS = ..., Service = ..., Federated = ... } HCL literal.
func principalFromExpr(expr hclsyntax.Expression) awsPrincipal {
	var p awsPrincipal

	if s := extractStringValue(expr); s != "" {
		p.Wildcard = s == "*"
		return p
	}

	obj, ok := unwrapExpr(expr).(*hclsyntax.ObjectConsExpr)
	if !ok {
		return p
	}

	for _, item := range obj.Items {
		vals := stringOrListExpr(unwrapExpr(item.ValueExpr))
		switch {
		case keyEquals(item.KeyExpr, "AWS"):
			p.AWS = vals
		case keyEquals(item.KeyExpr, "Service"):
			p.Service = vals
		case keyEquals(item.KeyExpr, "Federated"):
			p.Federated = vals
		}
	}

	p.Wildcard = slices.Contains(p.AWS, "*")

	return p
}

// conditionKeysFromObjectExpr walks a Condition = { Operator = { key = val } }
// HCL object literal and returns the condition variable names (the
// second-level keys), e.g. "iam:PassedToService".
func conditionKeysFromObjectExpr(expr hclsyntax.Expression) []string {
	obj, ok := unwrapExpr(expr).(*hclsyntax.ObjectConsExpr)
	if !ok {
		return nil
	}
	var keys []string
	for _, opItem := range obj.Items {
		inner, ok := unwrapExpr(opItem.ValueExpr).(*hclsyntax.ObjectConsExpr)
		if !ok {
			continue
		}
		for _, kv := range inner.Items {
			if k := hclKeyName(kv.KeyExpr); k != "" {
				keys = append(keys, k)
			}
		}
	}
	return keys
}

func stringOrListExpr(expr hclsyntax.Expression) []string {
	if t, ok := expr.(*hclsyntax.TupleConsExpr); ok {
		var out []string
		for _, e := range t.Exprs {
			if s := extractStringValue(e); s != "" {
				out = append(out, s)
			}
		}
		return out
	}
	if s := extractStringValue(expr); s != "" {
		return []string{s}
	}
	return nil
}

func keyEquals(key hclsyntax.Expression, want string) bool {
	return hclKeyName(key) == want
}

// hclKeyName returns the string name of an object-literal key, whether it's
// a bare identifier (Statement = ...) or a quoted string ("Statement" = ...).
func hclKeyName(key hclsyntax.Expression) string {
	k := unwrapExpr(key)
	// Bare identifier keys (Statement = ...) parse as a single-step traversal.
	if t, ok := k.(*hclsyntax.ScopeTraversalExpr); ok && len(t.Traversal) == 1 {
		return t.Traversal.RootName()
	}
	return extractStringValue(k)
}

func unwrapExpr(e hclsyntax.Expression) hclsyntax.Expression {
	for {
		switch w := e.(type) {
		case *hclsyntax.ObjectConsKeyExpr:
			e = w.Wrapped
		case *hclsyntax.ParenthesesExpr:
			e = w.Expression
		default:
			return e
		}
	}
}

// ---- handlers ------------------------------------------------------------

func auditAWSInlinePolicy(block *hclsyntax.Block, baseDir string) []Finding {
	attr, ok := block.Body.Attributes["policy"]
	if !ok {
		return nil
	}
	doc, reason := extractAWSPolicy(attr.Expr, baseDir)
	if doc == nil {
		return []Finding{finding(block, "AWS000", SevInfo,
			"policy document not statically analysable", reason)}
	}
	return runAWSPolicyRules(block, doc)
}

func auditAWSRoleInlinePolicies(block *hclsyntax.Block, baseDir string) []Finding {
	var out []Finding

	if attr, ok := block.Body.Attributes["assume_role_policy"]; ok {
		doc, reason := extractAWSPolicy(attr.Expr, baseDir)
		if doc == nil {
			out = append(out, finding(block, "AWS000", SevInfo,
				"assume_role_policy not statically analysable", reason))
		} else {
			// A trust policy's statements carry a Principal, not a Resource, and
			// its actions (sts:AssumeRole*) aren't meaningfully scoped by the
			// same Conditions an identity policy's escalation-class actions are
			// — so run the Principal-aware trust rules only, not the generic
			// Action/Resource rules runAWSPolicyRules applies to identity/
			// resource policies.
			out = append(out, runAWSTrustPolicyRules(block, doc)...)
		}
	}

	for _, inner := range block.Body.Blocks {
		if inner.Type != "inline_policy" {
			continue
		}
		if attr, ok := inner.Body.Attributes["policy"]; ok {
			doc, reason := extractAWSPolicy(attr.Expr, baseDir)
			if doc == nil {
				out = append(out, finding(block, "AWS000", SevInfo,
					"inline_policy not statically analysable", reason))
				continue
			}
			out = append(out, runAWSPolicyRules(block, doc)...)
		}
	}
	return out
}

func auditAWSPolicyDocument(block *hclsyntax.Block, _ string) []Finding {
	var doc awsPolicyDoc
	for _, stmt := range block.Body.Blocks {
		if stmt.Type != "statement" {
			continue
		}
		s := awsStmt{
			Sid:       attrString(stmt, "sid"),
			Effect:    attrString(stmt, "effect"),
			Action:    attrStringList(stmt, "actions"),
			NotAction: attrStringList(stmt, "not_actions"),
			Resource:  attrStringList(stmt, "resources"),
		}
		if s.Effect == "" {
			s.Effect = "Allow"
		}
		for _, c := range stmt.Body.Blocks {
			if c.Type == "condition" {
				s.Condition = json.RawMessage(`{}`)
				if v := attrString(c, "variable"); v != "" {
					s.ConditionKeys = append(s.ConditionKeys, v)
				}
			}
		}
		doc.Statement = append(doc.Statement, s)
	}
	return runAWSPolicyRules(block, &doc)
}

func auditAWSPolicyAttachment(block *hclsyntax.Block, _ string) []Finding {
	arn := attrString(block, "policy_arn")
	if arn == "" {
		return nil
	}
	name := arn[strings.LastIndex(arn, "/")+1:]
	if awsManagedPolicyDenylist[name] || (strings.HasSuffix(name, "FullAccess") && strings.HasPrefix(arn, "arn:aws:iam::aws:policy/")) {
		return []Finding{finding(block, "AWS006", SevHigh,
			fmt.Sprintf("attaches admin-class managed policy %s", name),
			"prefer a scoped customer-managed policy")}
	}
	return nil
}

// awsAttachmentLabels returns the principal label(s) (e.g. "role
// aws_iam_role.ci") an attachment block names, for resolveReferenceGraph:
// a single role/user/group attribute (aws_iam_{role,user,group}_policy_
// attachment), or the plural roles/users/groups list form
// (aws_iam_policy_attachment). A reference attribute (role = aws_iam_role.
// ci.name) is identified by its resource address; a literal value (role =
// "ci") is used verbatim.
func awsAttachmentLabels(block *hclsyntax.Block) []string {
	var labels []string

	add := func(kind string, raw string) {
		if raw == "" {
			return
		}
		labels = append(labels, kind+" "+principalIdentifier(raw))
	}

	add("role", attrString(block, "role"))
	add("user", attrString(block, "user"))
	add("group", attrString(block, "group"))

	for _, v := range attrStringList(block, "roles") {
		add("role", v)
	}
	for _, v := range attrStringList(block, "users") {
		add("user", v)
	}
	for _, v := range attrStringList(block, "groups") {
		add("group", v)
	}

	return labels
}

// ---- rules ---------------------------------------------------------------

func runAWSPolicyRules(block *hclsyntax.Block, doc *awsPolicyDoc) []Finding {
	var out []Finding
	for i, s := range doc.Statement {
		if !strings.EqualFold(s.Effect, "Allow") {
			continue
		}
		ref := stmtRef(s.Sid, i)

		for _, a := range s.Action {
			escalates := escalationAWS[a]
			switch {
			case a == "*":
				out = append(out, finding(block, "AWS001", SevHigh,
					fmt.Sprintf("%s allows Action \"*\"", ref),
					"full-wildcard action grants every API call"))
			case strings.HasSuffix(a, ":*"):
				out = append(out, finding(block, "AWS002", SevMedium,
					fmt.Sprintf("%s allows service-wide action %s", ref, a),
					"prefer enumerating the specific actions required"))
				escalates = escalates || escalationAWSServiceWildcards[a]
			}
			if escalates && !s.hasScopingCondition() {
				out = append(out, finding(block, "AWS004", SevHigh,
					fmt.Sprintf("%s allows escalation-class action %s without a scoping Condition", ref, a),
					"add a service-scoping condition (e.g. iam:PassedToService) or restrict Resource"))
			}
		}

		if len(s.NotAction) > 0 {
			out = append(out, finding(block, "AWS005", SevMedium,
				fmt.Sprintf("%s uses Allow with NotAction", ref),
				"allow-except statements are easy to over-grant as new actions are added"))
		}

		if isWildcardResource(s.Resource) && hasResourceScopableAction(s.Action) {
			out = append(out, finding(block, "AWS003", SevMedium,
				fmt.Sprintf("%s uses Resource \"*\" with resource-scopable actions", ref),
				"restrict Resource to specific ARNs where the service supports it"))
		}
	}
	return out
}

// runAWSTrustPolicyRules inspects an assume_role_policy's Principal/
// Condition pairs for overly permissive trust relationships — who can
// assume the role, as opposed to runAWSPolicyRules' what the role can then
// do.
func runAWSTrustPolicyRules(block *hclsyntax.Block, doc *awsPolicyDoc) []Finding {
	var out []Finding
	for i, s := range doc.Statement {
		if !strings.EqualFold(s.Effect, "Allow") {
			continue
		}
		ref := stmtRef(s.Sid, i)
		p := s.Principal

		if p.Wildcard && !s.hasAnyCondition() {
			out = append(out, finding(block, "AWS007", SevHigh,
				fmt.Sprintf("%s allows any principal (Principal \"*\" or AWS \"*\") to assume this role with no Condition", ref),
				"anyone, in any account, can assume this role — restrict Principal or add a scoping Condition"))
		}

		if hasWildcardRootAccount(p.AWS) && !s.hasConditionKey("aws:principalorgid") {
			out = append(out, finding(block, "AWS008", SevHigh,
				fmt.Sprintf("%s trusts arn:aws:iam::*:root (any account's root user)", ref),
				"restrict to specific account IDs, or scope with a Condition on aws:PrincipalOrgID"))
		}

		if len(p.Service) > 0 && containsActionFold(s.Action, "sts:AssumeRole") &&
			!s.hasConditionKey("aws:sourcearn", "aws:sourceaccount") {
			out = append(out, finding(block, "AWS009", SevMedium,
				fmt.Sprintf("%s lets service principal %s call sts:AssumeRole with no aws:SourceArn/aws:SourceAccount Condition",
					ref, strings.Join(p.Service, ",")),
				"confused-deputy risk — scope with aws:SourceArn or aws:SourceAccount"))
		}

		if len(p.Federated) > 0 && containsActionFold(s.Action, "sts:AssumeRoleWithWebIdentity") &&
			!s.hasSubOrAudCondition() {
			out = append(out, finding(block, "AWS010", SevHigh,
				fmt.Sprintf("%s lets federated/OIDC principal %s call sts:AssumeRoleWithWebIdentity with no claim-scoping Condition (e.g. a :sub or :aud key)",
					ref, strings.Join(p.Federated, ",")),
				"without a condition on the provider's sub/aud claim, any subject at that OIDC provider (e.g. any GitHub repo) can assume this role"))
		}
	}
	return out
}

func stmtRef(sid string, i int) string {
	if sid != "" {
		return "statement " + sid
	}
	return fmt.Sprintf("statement[%d]", i)
}

func isWildcardResource(r []string) bool {
	return len(r) == 1 && r[0] == "*"
}

// hasResourceScopableAction returns true if at least one action in the list
// supports resource-level ARN scoping. We approximate by excluding read-only
// actions (which are commonly global) and pure wildcards.
func hasResourceScopableAction(actions []string) bool {
	for _, a := range actions {
		if a == "*" || strings.HasSuffix(a, ":*") {
			continue
		}
		if !isAWSReadPerm(a) {
			return true
		}
	}
	return false
}
