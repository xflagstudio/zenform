package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	version = "0.0.0"
)

var RootCommand = &cobra.Command{
	Use:   "zenform",
	Short: "Zendesk provisioning tool",
	Long:  "zenform is provisioning tool for Zendesk. It enables to build up support environment according to config files.",
}

func init() {
	cobra.OnInitialize()

	RootCommand.AddCommand(versionCommand)
	RootCommand.AddCommand(applyCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Output version",
	Long:  "Output version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`
zenform - v%s

(C) 2018 CRE Team, XFLAG Studio
    https://xflag.com

Do you have any question or trouble?
Please refer to https://github.com/xflagstudio/zenform/issues

`, version)
	},
}
