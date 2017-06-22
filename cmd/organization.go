package cmd

import "github.com/spf13/cobra"

var organizationCmd = &cobra.Command{
	Use:   "octoman",
	Short: "Ocotoman manages you GitHub organization(s)",
	Long: `Octoman allows you to manage your GitHub organization(s).

Octoman does this by fetching the remote state of your GitHub organization(s)
and comparing it with your local state.

Built with passion by hawky-4s- in Go.
`,
}
