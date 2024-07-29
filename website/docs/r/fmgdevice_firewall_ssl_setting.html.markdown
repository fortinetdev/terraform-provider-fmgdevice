---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ssl_setting"
description: |-
  SSL proxy settings.
---

# fmgdevice_firewall_ssl_setting
SSL proxy settings.

## Example Usage

```hcl
resource "fmgdevice_firewall_ssl_setting" "trname" {
  abbreviate_handshake      = "disable"
  cert_cache_capacity       = 10
  cert_cache_timeout        = 10
  kxp_queue_threshold       = 10
  no_matching_cipher_action = "bypass"
  device_name               = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `abbreviate_handshake` - Enable/disable use of SSL abbreviated handshake. Valid values: `disable`, `enable`.

* `cert_cache_capacity` - Maximum capacity of the host certificate cache (0 - 500, default = 200).
* `cert_cache_timeout` - Time limit to keep certificate cache (1 - 120 min, default = 10).
* `kxp_queue_threshold` - Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
* `no_matching_cipher_action` - Bypass or drop the connection when no matching cipher is found. Valid values: `drop`, `bypass`.

* `proxy_connect_timeout` - Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
* `session_cache_capacity` - Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
* `session_cache_timeout` - Time limit to keep SSL session state (1 - 60 min, default = 20).
* `ssl_dh_bits` - Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.

* `ssl_queue_threshold` - Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall SslSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ssl_setting.labelname FirewallSslSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

