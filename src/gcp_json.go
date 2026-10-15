package pike

import (
	"encoding/json"
	"strings"
)

func gcpPermissionJSON(permissions []string) string {
	if len(permissions) == 0 {
		return ""
	}
	role := struct {
		IncludedPermissions []string `json:"includedPermissions"`
	}{IncludedPermissions: Unique(permissions)}
	data, _ := json.MarshalIndent(role, "", "    ")
	return string(data)
}

// gcpJSONOutput preserves the existing single-provider AWS output shape and
// uses provider keys only for mixed AWS/GCP scans. A GCP permission document
// deliberately omits project/name: permissions can require different scopes,
// including the provider's quota project rather than the resource project.
func (out OutputPolicy) gcpJSONOutput(legacy bool) string {
	gcp := json.RawMessage(out.gcpJSON)
	aws := json.RawMessage(strings.TrimSpace(out.AWS.JSONOut))
	if !legacy {
		type roles struct {
			Apply json.RawMessage `json:"apply,omitempty"`
			Plan  json.RawMessage `json:"plan,omitempty"`
		}
		gcp, _ = json.Marshal(roles{Apply: gcp, Plan: json.RawMessage(out.planGCPJSON)})
		if len(aws) > 0 {
			aws, _ = json.Marshal(roles{Apply: aws, Plan: json.RawMessage(strings.TrimSpace(out.PlanAWS.JSONOut))})
		}
	}
	var value any = gcp
	if len(aws) > 0 {
		value = map[string]json.RawMessage{"aws": aws, "gcp": gcp}
	}
	data, _ := json.MarshalIndent(value, "", "    ")
	return string(data) + "\n"
}
