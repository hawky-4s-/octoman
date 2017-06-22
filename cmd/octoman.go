package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	username string
	token    string

	verbose    bool
	logFile    string
	logging    bool
	verboseLog bool
)

// OctomanCmd is Octoman's root command.
// Every other command attached to OctomanCmd is a child command to it.
var OctomanCmd = &cobra.Command{
	Use:   "octoman",
	Short: "Ocotoman manages you GitHub organization(s)",
	Long: `Octoman allows you to manage your GitHub organization(s).

Octoman does this by fetching the remote state of your GitHub organization(s)
and comparing it with your local state.

Built with passion by hawky-4s- in Go.
`,
}

var octomanCmdV *cobra.Command

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	AddCommands()

	if c, err := OctomanCmd.ExecuteC(); err != nil {
		c.Println("")
		c.Println(c.UsageString())

		os.Exit(-1)
	}
}

func AddCommands() {
	OctomanCmd.AddCommand(organizationCmd)
	OctomanCmd.AddCommand(versionCmd)
}

// init initializes flags.
func init() {
	OctomanCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	OctomanCmd.PersistentFlags().BoolVar(&logging, "log", false, "enable Logging")
	OctomanCmd.PersistentFlags().StringVar(&logFile, "logFile", "", "log File path (if set, logging enabled automatically)")
	OctomanCmd.PersistentFlags().BoolVar(&verboseLog, "verboseLog", false, "verbose logging")

	initRootPersistentFlags()
	initHugoBuilderFlags(OctomanCmd)

	OctomanCmd.Flags().BoolVarP(&buildWatch, "watch", "w", false, "watch filesystem for changes and recreate as needed")
	octomanCmdV = OctomanCmd

	// Set bash-completion
	_ = OctomanCmd.PersistentFlags().SetAnnotation("logFile", cobra.BashCompFilenameExt, []string{})
}

func initRootPersistentFlags() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	OctomanCmd.PersistentFlags().StringVar(&username, "u", "", "username, defaults to be empty")
	OctomanCmd.PersistentFlags().StringVar(&token, "p", "", "password, defaults to be empty")

	// Set bash-completion
	validConfigFilenames := []string{"json", "js", "yaml", "yml", "toml", "tml"}
	_ = OctomanCmd.PersistentFlags().SetAnnotation("config", cobra.BashCompFilenameExt, validConfigFilenames)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// default values
	viper.SetDefault("username", "")
	viper.SetDefault("token", "")

	viper.SetEnvPrefix("octoman")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
