package pike

import (
	_ "embed"
	"testing"
)

func TestAZUREPolicy(t *testing.T) {
	type args struct {
		permissions []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"pass",
			args{[]string{"woof"}},
			"resource \"azurerm_role_definition\" \"terraform_pike\" {\n  role_definition_id = local.uuid\n  name               = \"terraform_pike\"\n  scope              = data.azurerm_subscription.primary.id\n\n  permissions {\n    actions = [\n    \"woof\"]\n    not_actions = []\n  }\n\n  assignable_scopes = [\n    data.azurerm_subscription.primary.id,\n  ]\n}\n\nlocals {\n  uuid = uuid()\n}\n\ndata \"azurerm_subscription\" \"primary\" {\n}\n",
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AZUREPolicy(tt.args.permissions)
			if (err != nil) != tt.wantErr {
				t.Errorf("AZUREPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			mingot := minify(got)
			minwant := minify(tt.want)
			if mingot != minwant {
				t.Errorf("AZUREPolicy() = %v, want %v", mingot, minwant)
			}
		})
	}
}
