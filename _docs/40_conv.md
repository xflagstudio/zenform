---
title: conv
permalink: /docs/commands/conv/
---

Convert configurations to HCL(.tf)

## Usage

```
Usage:
  zenform conv [config directory] [flags]

Flags:
  -f, --format string   Configuration file format.
                        Only "csv" is supported currently.
                        default: "csv"

  -h, --help            help for conv
```

If target directory is not specified, zenform use current directory by default.
