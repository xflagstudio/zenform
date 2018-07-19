package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/spf13/cobra"
	"github.com/xflagstudio/zenform/command/apply"
	"github.com/xflagstudio/zenform/command/common"
	"github.com/xflagstudio/zenform/command/step"
	"github.com/xflagstudio/zenform/config"
)

var configFormat string

const stateFileName = "zfstate.json"

func init() {
	applyCommand.Flags().StringVarP(&configFormat, "format", "f", "csv", "Configuration file format. Only \"csv\" is supported currently.")
}

var applyCommand = &cobra.Command{
	Use:   "apply",
	Short: "Apply configuration to Zendesk",
	Long:  "Apply configuration to Zendesk",
	Run:   applyFunc,
}

func applyFunc(cmd *cobra.Command, args []string) {
	exe := step.NewExecutor()
	zd, _ := zendesk.NewClient(nil)
	zfconfig := &config.ZenformConfig{}
	currentState := config.NewZenformState()
	var conf config.Config

	// If specified zenform project directory path,
	// `apply` is executed in it.
	exe.Step("Running apply", func() error {
		if len(args) > 0 {
			os.Chdir(args[0])
		}
		cwd, _ := os.Getwd()
		exe.Info("Check into directory: " + cwd)
		return nil
	})

	exe.Step("Checking zenform config", common.StepLoadZenformConfig(exe, zfconfig, zd))

	// If zfstate.json exists,
	// recover last config data as current Zendesk state
	exe.Step("Checking state file", func() error {
		if _, err := os.Stat(stateFileName); os.IsNotExist(err) {
			exe.Info("Not Found")
			return nil
		}
		exe.Success("Found " + stateFileName)

		exe.Step("Loading state", func() error {
			stateJsonStr, _ := ioutil.ReadFile(stateFileName)
			err := json.Unmarshal(stateJsonStr, &currentState)
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			return nil
		})
		return nil
	})

	// Create parser according to config extension
	configExtension := "." + configFormat
	parser, err := config.NewParserFromExtension(configExtension)
	if err != nil {
		fmt.Println(err)
		return
	}

	//TODO: check if all configs exists
	exe.Step("Detecting config files", func() error {
		exe.Success("Found (number of files) files")
		return nil
	})

	exe.Step("Loading configuration files", func() error {
		exe.Step("Loading ticket_fields"+configExtension, func() error {
			confText, err := ioutil.ReadFile("./ticket_fields" + configExtension)
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			conf.TicketFields, err = parser.ParseTicketFields(string(confText))
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			exe.Success("OK")
			return nil
		})

		exe.Step("Loading ticket_forms"+configExtension, func() error {
			confText, err := ioutil.ReadFile("./ticket_forms" + configExtension)
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			conf.TicketForms, err = parser.ParseTicketForms(string(confText))
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			exe.Success("OK")
			return nil
		})

		exe.Step("Loading triggers"+configExtension, func() error {
			confText, err := ioutil.ReadFile("./triggers" + configExtension)
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			conf.Triggers, err = parser.ParseTriggers(string(confText))
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			exe.Success("OK")
			return nil
		})
		return nil
	})

	exe.Step("Applying patch to Zendesk", func() error {
		if err := exe.StepIf(len(conf.TicketFields) > 0, "Creating ticket fields...", apply.StepCreateTicketFields(exe, zd, conf, currentState)); err != nil {
			return err
		}
		if err := exe.StepIf(len(conf.TicketForms) > 0, "Creating ticket forms...", apply.StepCreateTicketForms(exe, zd, conf, currentState)); err != nil {
			return err
		}
		if err := exe.StepIf(len(conf.Triggers) > 0, "Creating triggers...", apply.StepCreateTriggers(exe, zd, conf, currentState)); err != nil {
			return err
		}
		return nil
	})

	// Write new state file
	exe.Step("Updating zfstate.json", func() error {
		jsonStr, err := json.MarshalIndent(currentState, "", "    ") // indent with 4 spaces
		if err != nil {
			exe.Error(err.Error())
			return err
		}
		ioutil.WriteFile(stateFileName, jsonStr, 0644)
		exe.Success("Done")
		return nil
	})
}
