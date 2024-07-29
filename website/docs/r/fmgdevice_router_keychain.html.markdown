---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_keychain"
description: |-
  Configure key-chain.
---

# fmgdevice_router_keychain
Configure key-chain.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `key`: `fmgdevice_router_keychain_key`



## Example Usage

```hcl
resource "fmgdevice_router_keychain" "trname" {
  key {
    accept_lifetime = "your own value"
    algorithm       = "cmac-aes128"
    fosid           = "your own value"
    key_string      = ["your own value"]
    send_lifetime   = "your own value"
  }

  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `key` - Key. The structure of `key` block is documented below.
* `name` - Key-chain name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `key` block supports:

* `accept_lifetime` - Lifetime of received authentication key (format: hh:mm:ss day month year).
* `algorithm` - Cryptographic algorithm. Valid values: `md5`, `hmac-sha1`, `hmac-sha256`, `hmac-sha384`, `hmac-sha512`, `cmac-aes128`.

* `id` - Key ID (0 - 2147483647).
* `key_string` - Password for the key (maximum = 64 characters).
* `send_lifetime` - Lifetime of sent authentication key (format: hh:mm:ss day month year).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router KeyChain can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_keychain.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

