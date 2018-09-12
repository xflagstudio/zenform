---
title: Installation
permalink: /docs/
---

Use of released version is recommended.

## Released Binary

Download the binary from [release page](https://github.com/xflagstudio/zenform/releases/latest) according to your operating system.

| OS      | Binary Name                                   |
|:--------|:----------------------------------------------|
| Linux   | zenform-linux-amd64<br/>zenform-linux-386     |
| macOS   | zenform-darwin-amd64<br/>zenform-darwin-386   |
| Windows | zenform-windows-amd64<br/>zenform-windows-386 |

## From Source

You can build master branch, if you want to use pre-release version.  

It requires **Go >= 1.9** and [dep](https://github.com/golang/dep).

```sh
$ git clone git@github.com:xflagstudio/zenform.git
$ cd zenform
$ dep ensure
$ go install
```
