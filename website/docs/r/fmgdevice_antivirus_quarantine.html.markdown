---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_antivirus_quarantine"
description: |-
  Configure quarantine options.
---

# fmgdevice_antivirus_quarantine
Configure quarantine options.

## Example Usage

```hcl
resource "fmgdevice_antivirus_quarantine" "trname" {
  agelimit       = 10
  destination    = "NULL"
  drop_blocked   = ["smtps"]
  drop_heuristic = ["nntp"]
  drop_infected  = ["mapi"]
  device_name    = var.device_name # not required if setting is at provider
  device_vdom    = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `agelimit` - Age limit for quarantined files (0 - 479 hours, 0 means forever).
* `destination` - Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.

* `drop_blocked` - Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `mm1`, `mm3`, `mm4`, `mm7`, `ftps`, `mapi`, `cifs`, `ssh`.

* `drop_heuristic` - Drop-Heuristic. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `im`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `mm1`, `mm3`, `mm4`, `mm7`, `ftps`, `mapi`, `cifs`, `ssh`.

* `drop_infected` - Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `im`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `mm1`, `mm3`, `mm4`, `mm7`, `ftps`, `mapi`, `cifs`, `ssh`.

* `drop_intercepted` - drop intercepted from a protocol Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `imaps`, `smtps`, `pop3s`, `https`, `mm1`, `mm3`, `mm4`, `mm7`, `ftps`, `mapi`.

* `drop_machine_learning` - Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.

* `lowspace` - Select the method for handling additional files when running low on disk space. Valid values: `ovrw-old`, `drop-new`.

* `maxfilesize` - Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
* `quarantine_quota` - The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
* `store_blocked` - Quarantine blocked files found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `mm1`, `mm3`, `mm4`, `mm7`, `ftps`, `mapi`, `cifs`, `ssh`.

* `store_heuristic` - Store-Heuristic. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `im`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `mm1`, `mm3`, `mm4`, `mm7`, `ftps`, `mapi`, `cifs`, `ssh`.

* `store_infected` - Quarantine infected files found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `im`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `mm1`, `mm3`, `mm4`, `mm7`, `ftps`, `mapi`, `cifs`, `ssh`.

* `store_intercepted` - quarantine intercepted from a protocol Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `imaps`, `smtps`, `pop3s`, `https`, `mm1`, `mm3`, `mm4`, `mm7`, `ftps`, `mapi`.

* `store_machine_learning` - Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Antivirus Quarantine can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_antivirus_quarantine.labelname AntivirusQuarantine
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

