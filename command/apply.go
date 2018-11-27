package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/xflagstudio/zenform/config"
)

var configFormat string

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
	ticketFieldTmpl, err := template.New("ticket_field").Parse(config.TicketFieldHCL)
	if err != nil {
		fmt.Errorf("Failed to load HCL template: ticket_field")
		os.Exit(1)
	}

	ticketFormTmpl, err := template.New("ticket_form").Parse(config.TicketFormHCL)
	if err != nil {
		fmt.Errorf("Failed to load HCL template: ticket_form")
		os.Exit(1)
	}

	triggerTmpl, err := template.New("trigger").Parse(config.TriggerHCL)
	if err != nil {
		fmt.Errorf("Failed to load HCL template: trigger")
		os.Exit(1)
	}

	var conf config.Config

	// If specified zenform project directory path,
	// `apply` is executed in it.
	if len(args) > 0 {
		os.Chdir(args[0])
	}

	// Create parser according to config extension
	configExtension := "." + configFormat
	parser, err := config.NewParserFromExtension(configExtension)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ticket field
	confText, err := ioutil.ReadFile("./ticket_fields" + configExtension)
	if err != nil {
		fmt.Printf("[E]: %s", err)
		os.Exit(1)
	}
	conf.TicketFields, err = parser.ParseTicketFields(string(confText))
	if err != nil {
		fmt.Printf("[E]: %s", err)
		os.Exit(1)
	}
	if err = ticketFieldTmpl.Execute(os.Stdout, map[string]interface{}{
		"FileName":     "ticket_fields",
		"TicketFields": conf.TicketFields,
	}); err != nil {
		fmt.Printf("[E]: %s", err)
		os.Exit(1)
	}

	// ticket form
	confText, err = ioutil.ReadFile("./ticket_forms" + configExtension)
	if err != nil {
		fmt.Printf("[E]: %s", err)
		os.Exit(1)
	}
	conf.TicketForms, err = parser.ParseTicketForms(string(confText))
	if err != nil {
		fmt.Printf("[E]: %s", err)
		os.Exit(1)
	}
	if err = ticketFormTmpl.Execute(os.Stdout, map[string]interface{}{
		"FileName":    "ticket_forms",
		"TicketForms": conf.TicketForms,
	}); err != nil {
		fmt.Printf("[E]: %s", err)
		os.Exit(1)
	}

	// trigger
	confText, err = ioutil.ReadFile("./triggers" + configExtension)
	if err != nil {
		fmt.Printf("[E]: %s", err)
		os.Exit(1)
	}
	conf.Triggers, err = parser.ParseTriggers(string(confText))
	if err != nil {
		fmt.Printf("[E]: %s", err)
		os.Exit(1)
	}
	if err = triggerTmpl.Execute(os.Stdout, map[string]interface{}{
		"FileName": "triggers",
		"Triggers": conf.Triggers,
	}); err != nil {
		fmt.Printf("[E]: %s", err)
		os.Exit(1)
	}
}
