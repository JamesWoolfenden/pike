package pike

import (
	_ "embed"
	"testing"
)

//nolint:funlen
func TestAZUREPolicy(t *testing.T) {
	t.Parallel()

	type args struct {
		permissions []string
		policyName  string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "pass",
			args: args{[]string{"woof"}, ""},
			want: "resource \"azurerm_role_definition\" \"terraform_pike\" {\n  role_definition_id = local.uuid\n" +
				"  name               = \"terraform_pike\"\n  scope              = data.azurerm_subscription.primary.id\n\n" +
				"  permissions {\n    actions = [\n    \"woof\"]\n    not_actions = []\n  }\n\n  assignable_scopes = [\n" +
				"    data.azurerm_subscription.primary.id,\n  ]\n}\n\nlocals {\n  uuid = uuid()\n}\n\ndata" +
				" \"azurerm_subscription\" \"primary\" {\n}\n",
			wantErr: false,
		},
		{
			name:    "fail",
			args:    args{[]string{}, ""},
			want:    "",
			wantErr: true,
		},
		{
			name: "Policy named",
			args: args{[]string{"woof"}, "pike"},
			want: "resource \"azurerm_role_definition\" \"pike\" {\n  role_definition_id = local.uuid\n" +
				"  name               = \"pike\"\n  scope              = data.azurerm_subscription.primary.id\n\n" +
				"  permissions {\n    actions = [\n    \"woof\"]\n    not_actions = []\n  }\n\n  assignable_scopes = [\n" +
				"    data.azurerm_subscription.primary.id,\n  ]\n}\n\nlocals {\n  uuid = uuid()\n}\n\ndata" +
				" \"azurerm_subscription\" \"primary\" {\n}\n",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := AZUREPolicy(tt.args.permissions, tt.args.policyName)

			if (err != nil) != tt.wantErr {
				t.Errorf("AZUREPolicy() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			minGot := Minify(got)
			minWant := Minify(tt.want)

			if minGot != minWant {
				t.Errorf("AZUREPolicy() = %v, want %v", minGot, minWant)
			}
		})
	}
}
