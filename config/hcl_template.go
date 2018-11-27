package config

// TicketFieldHCL is template for ticket field
const TicketFieldHCL = `# {{.FileName}}.tf
#   This file is generated from {{.FileName}}.csv
{{range .TicketFields}}
resource "zendesk_ticket_field" "{{.Slug}}" {
  name               = "{{.Title}}"
  type               = "{{.Type}}"
  description        = "{{.Description}}"
  visible_in_portal  = {{if eq .VisibleInPortal "TRUE"}}true{{else}}false{{end}}
  editable_in_portal = {{if eq .EditableInPortal "TRUE"}}true{{else}}false{{end}}
  required_in_portal = {{if eq .RequiredInPortal "TRUE"}}true{{else}}false{{end}}{{if eq .Type "tagger"}}
{{range .CustomFieldOptions}}
  custom_field_option {
    name  = "{{.Name}}"
    value = "{{.Value}}"
  }
{{end}}{{end}}
}
{{end}}`

// TicketFormHCL is template for ticket form
const TicketFormHCL = `# {{.FileName}}.tf
#   This file is generated from {{.FileName}}.csv
{{range .TicketForms}}
resource "zendesk_ticket_form" "{{.Slug}}" {
  name             = "{{.Name}}"
  display_name     = "{{.DisplayName}}"
  position         = {{.Position}}
  end_user_visible = {{if eq .EndUserVisible "TRUE"}}true{{else}}false{{end}}

  ticket_field_ids = [{{range .TicketFieldIDs}}
    "${zendesk_ticket_field.{{.}}.id}",{{end}}
  ]
}
{{end}}`

// TriggerHCL is template for trigger
const TriggerHCL = `# {{.FileName}}.tf
#   This file is generated from {{.FileName}}.csv
{{range .Triggers}}
resource "zendesk_trigger" "{{.Slug}}" {
  title    = "{{.Title}}"
  position = {{.Position}}
{{range .All}}
  all {
    field    = "{{.Field}}"
    operator = "{{.Operator}}"
    value    = "{{.Value}}"
  }
{{end}}{{range .Any}}
  any {
    field    = "{{.Field}}"
    operator = "{{.Operator}}"
    value    = "{{.Value}}"
  }
{{end}}{{range .Actions}}
  action {
    field = "{{.Field}}"
    value = {{if eq .Field "notification_user"}}<<ENVELOPE
[
  "{{index .Value 0}}",
  "{{index .Value 1}}",
  "{{index .Value 2}}"
]
ENVELOPE{{else}}"{{index .Value 0}}"{{end}}
  }
{{end}}
}
{{end}}`
