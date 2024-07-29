---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sessionhelper"
description: |-
  Configure session helper.
---

# fmgdevice_system_sessionhelper
Configure session helper.

## Example Usage

```hcl
resource "fmgdevice_system_sessionhelper" "trname" {
  fosid       = 10
  name        = "dcerpc"
  port        = 10
  protocol    = 10
  seq         = 10
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fosid` - Session helper ID.
* `name` - Helper name. Valid values: `ftp`, `tftp`, `ras`, `h323`, `h245O`, `h245I`, `tns`, `mms`, `sip`, `pptp`, `rtsp`, `dns-udp`, `dns-tcp`, `pmap`, `rsh`, `dcerpc`, `mgcp`, `gtp-c`, `gtp-u`, `gtp-b`, `pfcp`.

* `port` - Protocol port.
* `protocol` - Protocol number.
* `seq` - Seq.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SessionHelper can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sessionhelper.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

