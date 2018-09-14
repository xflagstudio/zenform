---
title: Credentials
permalink: /docs/configuration/
---

## Setup

zenform loads YAML file from current directory which holds your credentials of Zendesk instance.  
First of all, create `zenform.yml` according to the format as bellow.

```yaml
zendesk:
  # e.g. zenform.zendesk.com
  subdomain: zenform
  email: john.doe@example.com
  token: blablablabla
```
