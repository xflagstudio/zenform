---
title: apply
permalink: /docs/commands/apply/
---

Apply configurations to Zendesk.

## Usage

```
Usage:
  zenform apply [config directory] [flags]

Flags:
  -f, --format string   Configuration file format.
                        Only "csv" is supported currently.
                        default: "csv"

  -h, --help            help for apply
```

If target directory is not specified, zenform use current directory by default.

## What’s `zfstate.json` ?

`apply` will create `zfstate.json` into the same directory which holds configuration files.  
It has state of resources that have been created on Zendesk. This file prevents zenform to duplicate existing resource.
