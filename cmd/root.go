package cmd

import (
	listCmd "github.com/katiem0/gh-environments/cmd/list"
	secretsCmd "github.com/katiem0/gh-environments/cmd/secrets"
	variablesCmd "github.com/katiem0/gh-environments/cmd/variables"
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {

	cmdRoot := &cobra.Command{
		Use:   "environments <command> <subcommand> [flags]",
		Short: "List repo environments and metadata.",
		Long:  "List repo environments and metadata, including listing and creating environment secrets and variables.",
	}
	cmdRoot.PersistentFlags().Bool("help", false, "Show help for command")

	cmdRoot.AddCommand(listCmd.NewCmdList())
	cmdRoot.AddCommand(secretsCmd.NewCmdSecrets())
	cmdRoot.AddCommand(variablesCmd.NewCmdVariables())
	cmdRoot.CompletionOptions.DisableDefaultCmd = true
	cmdRoot.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	return cmdRoot
}
