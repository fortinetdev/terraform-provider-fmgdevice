---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_acme_accounts"
description: |-
  ACME accounts list.
---

# fmgdevice_system_acme_accounts
ACME accounts list.

~> This resource is a sub resource for variable `accounts` of resource `fmgdevice_system_acme`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_acme_accounts" "trname" {
  ca_url      = "your own value"
  email       = "your own value"
  fosid       = "your own value"
  privatekey  = "your own value"
  status      = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `ca_url` - Account ca_url.
* `email` - Account email.
* `fosid` - Account id.
* `privatekey` - Account Private Key.
* `status` - Account status.
* `url` - Account url.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System AcmeAccounts can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_acme_accounts.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

