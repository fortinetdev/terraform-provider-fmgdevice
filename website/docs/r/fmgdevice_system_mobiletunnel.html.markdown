---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_mobiletunnel"
description: |-
  Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.
---

# fmgdevice_system_mobiletunnel
Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `network`: `fmgdevice_system_mobiletunnel_network`



## Example Usage

```hcl
resource "fmgdevice_system_mobiletunnel" "trname" {
  hash_algorithm = "hmac-md5"
  home_address   = "your own value"
  home_agent     = "your own value"
  lifetime       = 10
  n_mhae_key     = ["your own value"]
  name           = "your own value"
  device_name    = var.device_name # not required if setting is at provider
  device_vdom    = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `hash_algorithm` - Hash Algorithm (Keyed MD5). Valid values: `hmac-md5`.

* `home_address` - Home IP address (Format: xxx.xxx.xxx.xxx).
* `home_agent` - IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
* `lifetime` - NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
* `n_mhae_key` - NEMO authentication key.
* `n_mhae_key_type` - NEMO authentication key type (ASCII or base64). Valid values: `ascii`, `base64`.

* `n_mhae_spi` - NEMO authentication SPI (default: 256).
* `name` - Tunnel name.
* `network` - Network. The structure of `network` block is documented below.
* `reg_interval` - NMMO HA registration interval (5 - 300, default = 5).
* `reg_retry` - Maximum number of NMMO HA registration retries (1 to 30, default = 3).
* `renew_interval` - Time before lifetime expiration to send NMMO HA re-registration (5 - 60, default = 60).
* `roaming_interface` - Select the associated interface name from available options.
* `status` - Enable/disable this mobile tunnel. Valid values: `disable`, `enable`.

* `tunnel_mode` - NEMO tunnel mode (GRE tunnel). Valid values: `gre`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `network` block supports:

* `id` - Network entry ID.
* `interface` - Select the associated interface name from available options.
* `prefix` - Class IP and Netmask with correction (Format:xxx.xxx.xxx.xxx xxx.xxx.xxx.xxx or xxx.xxx.xxx.xxx/x).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System MobileTunnel can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_mobiletunnel.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

