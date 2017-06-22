package octoman

import (
	"github.com/google/go-github/github"
)

// orgas -> repos
// team: repos: [repo1, repo2, repo3...]
// team: members: [member1, member2, member3...]

// collaborators
var client *github.Client

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

//func DiffOrganization(ctx context.Context, organization string) {
//    diffOrganization(ctx)
//    diffRepositories(ctx)
//    diffTeams(ctx)
//    diffTeamMemberships(ctx)
//}

func main() {
	token = "foo"
	// read token etc
	client = NewGitHubClient(token)
}

func ImportAllOrganizationsOfUser() []*github.Organization {
	// get all orgs of user and call ImportOrganization
}

func ImportOrganization(organization string) (bool, error) {

}

func ImportTeams(organization string) {

}

func ImportMembers(organization string) {

}

/**
 * Orga -> Repos, Teams, Collaborators
 * Repo -> Teams, Collaborators
 * Team -> Teammembers
 */
