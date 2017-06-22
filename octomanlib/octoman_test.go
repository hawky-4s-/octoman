package octomanlib

import (
	"context"
	"testing"
)

var token = "foo"

func TestRetrieveStateForAllGitHubOrganizations(t *testing.T) {
	client = NewGitHubClient(token)
	RetrieveStateForAllGitHubOrganizations(context.Background())
}

func TestRetrieveRemoteOrganizationState(t *testing.T) {
	client = NewGitHubClient(token)
	orga := "camunda"

	RetrieveRemoteOrganizationState(context.Background(), orga)
}
