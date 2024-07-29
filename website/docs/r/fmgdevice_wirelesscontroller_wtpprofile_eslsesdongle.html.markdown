---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_wtpprofile_eslsesdongle"
description: |-
  ESL SES-imagotag dongle configuration.
---

# fmgdevice_wirelesscontroller_wtpprofile_eslsesdongle
ESL SES-imagotag dongle configuration.

~> This resource is a sub resource for variable `esl_ses_dongle` of resource `fmgdevice_wirelesscontroller_wtpprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_wtpprofile_eslsesdongle" "trname" {
  apc_addr_type = "fqdn"
  apc_fqdn      = "your own value"
  apc_ip        = "your own value"
  apc_port      = 10
  coex_level    = "none"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `wtp_profile` - Wtp Profile.

* `apc_addr_type` - ESL SES-imagotag APC address type (default = fqdn). Valid values: `fqdn`, `ip`.

* `apc_fqdn` - FQDN of ESL SES-imagotag Access Point Controller (APC).
* `apc_ip` - IP address of ESL SES-imagotag Access Point Controller (APC).
* `apc_port` - Port of ESL SES-imagotag Access Point Controller (APC).
* `coex_level` - ESL SES-imagotag dongle coexistence level (default = none). Valid values: `none`.

* `compliance_level` - Compliance levels for the ESL solution integration (default = compliance-level-2). Valid values: `compliance-level-2`.

* `esl_channel` - ESL SES-imagotag dongle channel (default = 127). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `127`, `-1`.

* `output_power` - ESL SES-imagotag dongle output power (default = A). Valid values: `a`, `b`, `c`, `d`, `e`, `f`, `g`, `h`.

* `scd_enable` - Enable/disable ESL SES-imagotag Serial Communication Daemon (SCD) (default = disable). Valid values: `disable`, `enable`.

* `tls_cert_verification` - Enable/disable TLS certificate verification (default = enable). Valid values: `disable`, `enable`.

* `tls_fqdn_verification` - Enable/disable TLS certificate verification (default = disable). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController WtpProfileEslSesDongle can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "wtp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_wtpprofile_eslsesdongle.labelname WirelessControllerWtpProfileEslSesDongle
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

