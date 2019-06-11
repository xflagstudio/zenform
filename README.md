# :building_construction: zenform
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![CircleCI](https://circleci.com/gh/xflagstudio/zenform/tree/master.svg?style=svg)](https://circleci.com/gh/xflagstudio/zenform/tree/master)

  zenform is a CLI tool to provision Zendesk instance.
  It converts your configuration files to HCL so that you can set up your Zendesk instance with [terraform-provider-zendesk](https://github.com/nukosuke/terraform-provider-zendesk).

## Usage

```sh
$ zenform conv input.csv > resources.tf

# Terraform with terraform-provider-zendesk
$ terraform init
$ terraform plan
$ terraform apply
```

## License

  (C) 2018 - CRE team, [XFLAG Studio](https://career.xflag.com/), mixi inc.

  This software is released under the MIT License. See LICENSE for details.
