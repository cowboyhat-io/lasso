package ado

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

type CreateWorkItemArg struct {
	C    workitemtracking.Client
	CTX  context.Context
	PROJ string
	Type string
	Item *webapi.JsonPatchOperation
}
