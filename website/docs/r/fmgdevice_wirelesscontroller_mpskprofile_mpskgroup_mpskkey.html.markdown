---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_mpskprofile_mpskgroup_mpskkey"
description: |-
  List of multiple PSK entries.
---

# fmgdevice_wirelesscontroller_mpskprofile_mpskgroup_mpskkey
List of multiple PSK entries.

~> This resource is a sub resource for variable `mpsk_key` of resource `fmgdevice_wirelesscontroller_mpskprofile_mpskgroup`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_mpskprofile_mpskgroup_mpskkey" "trname" {
  comment                      = "your own value"
  concurrent_client_limit_type = "specified"
  concurrent_clients           = 10
  key_type                     = "wpa3-sae"
  mac                          = "your own value"
  name                         = "your own value"
  device_name                  = var.device_name # not required if setting is at provider
  device_vdom                  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `mpsk_profile` - Mpsk Profile.
* `mpsk_group` - Mpsk Group.

* `comment` - Comment.
* `concurrent_client_limit_type` - MPSK client limit type options. Valid values: `default`, `unlimited`, `specified`.

* `concurrent_clients` - Number of clients that can connect using this pre-shared key (1 - 65535, default is 256).
* `key_type` - Select the type of the key. Valid values: `wpa2-personal`, `wpa3-sae`.

* `mac` - MAC address.
* `mpsk_schedules` - Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid.
* `name` - Pre-shared key name.
* `passphrase` - WPA Pre-shared key.
* `pmk` - WPA PMK.
* `sae_password` - WPA3 SAE password.
* `sae_pk` - Enable/disable WPA3 SAE-PK (default = disable). Valid values: `disable`, `enable`.

* `sae_private_key` - Private key used for WPA3 SAE-PK authentication.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController MpskProfileMpskGroupMpskKey can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "mpsk_profile=YOUR_VALUE", "mpsk_group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_mpskprofile_mpskgroup_mpskkey.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

