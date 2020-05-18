package goenv

import (
	"os"
	"testing"
)

type EnvVars struct {
	VarWithDefault    string `json:"var_with_default"`
	VarWithoutDefault string `json:"var_without_default"`
}

func TestSyncEnvVar(t *testing.T) {
	envs := EnvVars{
		VarWithDefault: "defaultval",
	}

	err := os.Setenv("VAR_WITHOUT_DEFAULT", "envval")
	if err != nil {
		t.Fatal("failed to set env var")
	}

	type args struct {
		env interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "should pass with valid values",
			args: args{
				env: &envs,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SyncEnvVar(tt.args.env)
			if envs.VarWithDefault != "defaultval" || envs.VarWithoutDefault != "envval" {
				t.Log(envs)
				t.Error("env values are not synced currectly in the struct")
			}
		})
	}
}
