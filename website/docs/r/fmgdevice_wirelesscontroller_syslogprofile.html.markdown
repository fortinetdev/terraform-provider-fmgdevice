---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_syslogprofile"
description: |-
  Configure Wireless Termination Points (WTP) system log server profile.
---

# fmgdevice_wirelesscontroller_syslogprofile
Configure Wireless Termination Points (WTP) system log server profile.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_syslogprofile" "trname" {
  comment          = "your own value"
  log_level        = "information"
  name             = "your own value"
  server_addr_type = "fqdn"
  server_fqdn      = "your own value"
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `log_level` - Lowest level of log messages that FortiAP units send to this server (default = information). Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.

* `name` - WTP system log server profile name.
* `server_addr_type` - Syslog server address type (default = ip). Valid values: `fqdn`, `ip`.

* `server_fqdn` - FQDN of syslog server that FortiAP units send log messages to.
* `server_ip` - IP address of syslog server that FortiAP units send log messages to.
* `server_port` - Port number of syslog server that FortiAP units send log messages to (default = 514).
* `server_status` - Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `disable`, `enable`.

* `server_type` - Configure syslog server type (default = standard). Valid values: `standard`, `fortianalyzer`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController SyslogProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_syslogprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

