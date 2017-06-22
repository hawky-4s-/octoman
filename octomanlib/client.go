package octoman

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// orgas -> repos
// team: repos: [repo1, repo2, repo3...]
// team: members: [member1, member2, member3...]

// collaborators

func NewGitHubClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

//func GetAllOrganizationRepositories(ctx context.Context, organization string) {
//    client.Organizations.Get(ctx, organization)
//
//    client.Organizations.AddTeamMembership()
//    client.Organizations.AddTeamRepo()
//
//    team, _, error := client.Organizations.CreateTeam()
//
//    client.Organizations.ConcealMembership()
//    client.Organizations.PublicizeMembership()
//
//    client.Organizations.DeleteTeam()
//    client.Organizations.EditOrgMembership()
//    client.Organizations.EditTeam()
//    client.Organizations.GetOrgMembership()
//    client.Organizations.GetTeam()
//    client.Organizations.GetTeamMembership()
//
//    client.Organizations.ListMembers()
//    client.Organizations.ListOrgMemberships()
//    client.Organizations.ListOutsideCollaborators()
//    client.Organizations.ListTeamMembers()
//    client.Organizations.ListTeams()
//    client.Organizations.ListTeamRepos()
//    client.Organizations.ListPendingOrgInvitations()
//    client.Organizations.ListPendingTeamInvitations()
//
//
//    client.Organizations.RemoveMember()
//    client.Organizations.RemoveOrgMembership()
//    client.Organizations.RemoveTeamMembership()
//    client.Organizations.RemoveTeamRepo()
//
//    orga, _, error := client.Organizations.Get(ctx, organization)
//}
//
//func ListAllRepositoriesForUser(ctx context.Context, user string) {
//    client.Organizations.
//
//	// list all repositories for the authenticated user
//	repos, _, err := client.Repositories.List(ctx, "", nil)
//
//    for repo := range repos {
//
//    }
//}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func RetrieveStateForAllGitHubOrganizations(ctx context.Context) {
	orgas, _, err := client.Organizations.List(ctx, "", &github.ListOptions{PerPage: 100})
	handleError(err)
	for _, orga := range orgas {
		fmt.Println("Name:" + orga.GetLogin())
	}
	for _, orga := range orgas {
		RetrieveRemoteOrganizationState(ctx, orga.GetLogin())
	}
}

func RetrieveRemoteOrganizationState(ctx context.Context, organization string) {
	orga := retrieveOrganization(ctx, organization) // org-members
	fmt.Println("%v+\n", orga)
	//repos := retrieveRepositories(ctx, organization) // repo members and teams
	//fmt.Printf("%v\n", repos)
	//teams := retrieveTeams(ctx, organization) // all teams and team-members
	//fmt.Println("%v+\n", teams)
	//orgaMembers := retrieveOrganizationMembers(ctx, organization) // all teams and team-members
	//fmt.Println("%v+\n", orgaMembers)
	collaborators := retrieveOrganizationCollaborators(ctx, organization)
	for _, collaborator := range collaborators {
		fmt.Printf("%s - %d \n", collaborator.GetLogin(), collaborator.GetID())
	}
}

func retrieveOrganization(ctx context.Context, organization string) *github.Organization {
	orga, _, err := client.Organizations.Get(ctx, organization)
	handleError(err)
	return orga
}

func retrieveTeams(ctx context.Context, organization string) []*github.Team {
	teams, _, err := client.Organizations.ListTeams(ctx, organization, &github.ListOptions{PerPage: 100})
	handleError(err)
	return teams
}

func retrieveRepositories(ctx context.Context, organization string) []*github.Repository {
	opts := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	repos, _, err := client.Repositories.ListByOrg(ctx, organization, opts)
	handleError(err)
	return repos
}

func retrieveOrganizationMembers(ctx context.Context, organization string) []*github.User {
	opts := &github.ListMembersOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	members, _, err := client.Organizations.ListMembers(ctx, organization, opts)
	handleError(err)
	return members
}

func retrieveOrganizationCollaborators(ctx context.Context, organization string) []*github.User {
	opts := &github.ListOutsideCollaboratorsOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	collaborators, _, err := client.Organizations.ListOutsideCollaborators(ctx, organization, opts)
	handleError(err)
	return collaborators
}

/**
 * Orga -> Repos, Teams, Collaborators
 * Repo -> Teams, Collaborators
 * Team -> Teammembers
 */
