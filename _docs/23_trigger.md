---
title: Trigger
permalink: /docs/configuration/trigger/
---

Configuration format for Trigger.

{: .notice--warning}
:rotating_light: Currently, zenform supports only CSV format for Trigger.

## CSV Format

| Attribute Name             | Type       | Description                                                            |
|:---------------------------|:-----------|:-----------------------------------------------------------------------|
| `slug`                     | string     | Unique identifier of the trigger                                       |
| `title`                    | string     | Title of the trigger                                                   |
| `conditions.all.fields`    | []string   | Fields for AND condition                                               |
| `conditions.all.operators` | []string   | Operators for AND condition                                            |
| `conditions.all.values`    | []string   | Values for AND condition                                               |
| `conditions.any.fields`    | []string   | Fields for OR condition                                                |
| `conditions.any.operators` | []string   | Operators for OR condition                                             |
| `conditions.any.values`    | []string   | Values for OR condition                                                |
| `actions.fields`           | []string   | Fields to be updated by trigger                                        |
| `actions.values`           | [][]string | Values to be set to each field which was specified by `actions.fields` |

{: .notice--info}
:pencil: This format corresponds to [Zendesk API Triggers](https://developer.zendesk.com/rest_api/docs/core/triggers)
