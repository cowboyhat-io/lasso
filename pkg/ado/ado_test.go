package ado

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
	"testing"
)

func TestCommented(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		displayName string
		want        bool
	}{
		"DidComment": {
			displayName: "Toul DeGuia-Cranmer",
			want:        true,
		},
		"DidNotComment": {
			displayName: "Lord Frieza",
			want:        false,
		},
	}
	for displayName, tc := range tests {
		comments := make([]workitemtracking.Comment, 1)
		name := "Toul DeGuia-Cranmer"
		comments[0].CreatedBy = &webapi.IdentityRef{DisplayName: &name}
		t.Run(displayName, func(t *testing.T) {
			got := Commented(tc.displayName, comments)
			if got != tc.want {
				t.Errorf("Got: %v, want: %v", got, tc.want)
			}
		})
	}
}
