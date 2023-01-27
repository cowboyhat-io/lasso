package ado

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

// Commented takes the person responsible for doing AD Check's name to see if they've commented
// on the NFR story or not
func Commented(a string, list []workitemtracking.Comment) bool {
	for _, b := range list {
		if *b.CreatedBy.DisplayName == a {
			return true
		}
	}
	return false
}
