---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_mpskprofile"
description: |-
  Configure MPSK profile.
---

# fmgdevice_wirelesscontroller_mpskprofile
Configure MPSK profile.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `mpsk_group`: `fmgdevice_wirelesscontroller_mpskprofile_mpskgroup`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_mpskprofile" "trname" {
  mpsk_concurrent_clients   = 10
  mpsk_external_server      = ["your own value"]
  mpsk_external_server_auth = "enable"
  mpsk_group {
    mpsk_key {
      comment                      = "your own value"
      concurrent_client_limit_type = "default"
      concurrent_clients           = 10
      key_type                     = "wpa3-sae"
      mac                          = "your own value"
      mpsk_schedules               = ["your own value"]
      name                         = "your own value"
      passphrase                   = ["your own value"]
      pmk                          = ["your own value"]
      sae_password                 = ["your own value"]
      sae_pk                       = "enable"
      sae_private_key              = "your own value"
    }

    name      = "your own value"
    vlan_id   = 10
    vlan_type = "fixed-vlan"
  }

  mpsk_type   = "wpa3-sae"
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `mpsk_concurrent_clients` - Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
* `mpsk_external_server` - RADIUS server to be used to authenticate MPSK users.
* `mpsk_external_server_auth` - Enable/Disable MPSK external server authentication (default = disable). Valid values: `disable`, `enable`.

* `mpsk_group` - Mpsk-Group. The structure of `mpsk_group` block is documented below.
* `mpsk_type` - Select the security type of keys for this profile. Valid values: `wpa2-personal`, `wpa3-sae`, `wpa3-sae-transition`.

* `name` - MPSK profile name.
* `ssid` - SSID of the VAP in which the MPSK profile is configured.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `mpsk_group` block supports:

* `mpsk_key` - Mpsk-Key. The structure of `mpsk_key` block is documented below.
* `name` - MPSK group name.
* `vlan_id` - Optional VLAN ID.
* `vlan_type` - MPSK group VLAN options. Valid values: `no-vlan`, `fixed-vlan`.


The `mpsk_key` block supports:

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

WirelessController MpskProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_mpskprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

