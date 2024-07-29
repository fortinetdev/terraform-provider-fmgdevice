---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_acme"
description: |-
  Configure ACME client.
---

# fmgdevice_system_acme
Configure ACME client.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `accounts`: `fmgdevice_system_acme_accounts`



## Example Usage

```hcl
resource "fmgdevice_system_acme" "trname" {
  accounts {
    ca_url     = "your own value"
    email      = "your own value"
    fosid      = "your own value"
    privatekey = "your own value"
    status     = "your own value"
    url        = "your own value"
  }

  interface        = ["port2"]
  source_ip        = "your own value"
  source_ip6       = "your own value"
  store_passphrase = ["your own value"]
  device_name      = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `accounts` - Accounts. The structure of `accounts` block is documented below.
* `interface` - Interface(s) on which the ACME client will listen for challenges.
* `source_ip` - Source IPv4 address used to connect to the ACME server.
* `source_ip6` - Source IPv6 address used to connect to the ACME server.
* `store_passphrase` - mod_md store passphrase (Read-Only)
* `use_ha_direct` - Enable the use of 'ha-mgmt' interface to connect to the ACME server when 'ha-direct' is enabled in HA configuration Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `accounts` block supports:

* `ca_url` - Account ca_url.
* `email` - Account email.
* `id` - Account id.
* `privatekey` - Account Private Key.
* `status` - Account status.
* `url` - Account url.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Acme can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_acme.labelname SystemAcme
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

