---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_proberesponse"
description: |-
  Configure system probe response.
---

# fmgdevice_system_proberesponse
Configure system probe response.

## Example Usage

```hcl
resource "fmgdevice_system_proberesponse" "trname" {
  http_probe_value = "your own value"
  mode             = "twamp"
  password         = ["your own value"]
  port             = 10
  security_mode    = "none"
  device_name      = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `http_probe_value` - Value to respond to the monitoring server.
* `mode` - SLA response mode. Valid values: `none`, `http-probe`, `twamp`.

* `password` - TWAMP responder password in authentication mode.
* `port` - Port number to response.
* `security_mode` - TWAMP responder security mode. Valid values: `none`, `authentication`.

* `timeout` - An inactivity timer for a twamp test session.
* `ttl_mode` - Mode for TWAMP packet TTL modification. Valid values: `reinit`, `decrease`, `retain`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System ProbeResponse can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_proberesponse.labelname SystemProbeResponse
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

