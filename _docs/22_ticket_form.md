---
title: Ticket Form
permalink: /docs/configuration/ticket-form/
---

Configuration format for Ticket Form.

{: .notice--warning}
:rotating_light: Currently, zenform supports only CSV format for Ticket Form.

## CSV Format

| Attribute Name     | Type                                  | Description                                                                          |
|:-------------------|:--------------------------------------|:-------------------------------------------------------------------------------------|
| `slug`             | string                                | Unique identifier of the ticket form                                                 |
| `name`             | string                                | Name of the ticket form                                                              |
| `position`         | integer                               | Sequential number which indicates the order of ticket forms                          |
| `end_user_visible` | boolean<br/>( `"TRUE"` \| `"FALSE"` ) | Whether end user can see the form on [Zendesk Guide](https://www.zendesk.com/guide/) |
| `display_name`     | string                                | Display name of ticket field                                                         |
| `ticket_field_ids` | []string                              | slugs of the ticket field to be added                                                |

{: .notice--info}
:pencil: This format corresponds to [Zendesk API Ticket Forms](https://developer.zendesk.com/rest_api/docs/core/ticket_forms)
