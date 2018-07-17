// zenform - Provisioning tool for Zendesk
//
//   (C) 2018- nukosuke <nukosuke@lavabit.com>
//
// This software is released under MIT License.
// See LICENSE.

package main

import (
	"os"

	"github.com/xflagstudio/zenform/command"
)

func main() {
	if err := command.RootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
