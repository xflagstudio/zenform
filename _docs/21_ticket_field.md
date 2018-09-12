---
title: Ticket Field
permalink: /docs/configuration/ticket-field/
---

Configuration format for Ticket Field.

{: .notice--warning}
:rotating_light: Currently, zenform supports only CSV format for Ticket Field.

## CSV Format

| Attribute Name                | Type                                  | Description                                               |
|:------------------------------|:--------------------------------------|:----------------------------------------------------------|
| `slug`                        | string                                | Unique identifier of the ticket field                     |
| `title`                       | string                                | Title of the ticket field                                 |
| `visible_in_portal`           | boolean<br/>( `"TRUE"` \| `"FALSE"` ) | Whether the ticket field is visible on portal site        |
| `editable_in_portal`          | boolean<br/>( `"TRUE"` \| `"FALSE"` ) | Whether the ticket field is editable on portal site       |
| `required_in_portal`          | boolean<br/>( `"TRUE"` \| `"FALSE"` ) | Whether the ticket field is mandatory item on portal site |
| `description`                 | string                                | Description of the ticket field                           |
| `custom_field_options.names`  | []string                              | option name to custom                                     |
| `custom_field_options.values` | []string                              | Option value to custom                                    |

{: .notice--info}
:pencil: This format corresponds to [Zendesk API Ticket Fields](https://developer.zendesk.com/rest_api/docs/core/ticket_fields)

## Want `”hold”` status?

By default, Zendesk API returns error if you try to create a trigger using `"hold"` status. It's necessary to activate `"hold"` status by yourself.

See [this article by Zendesk](https://support.zendesk.com/hc/ja/articles/203661576-Zendesk-Support) for details.
