---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_ptp_profile"
description: |-
  Global PTP profile.
---

# fmgdevice_switchcontroller_ptp_profile
Global PTP profile.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_ptp_profile" "trname" {
  description         = "your own value"
  domain              = 10
  mode                = "transparent-e2e"
  name                = "your own value"
  pdelay_req_interval = "32sec"
  device_name         = var.device_name # not required if setting is at provider
  device_vdom         = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `description` - Description.
* `domain` - Configure PTP domain value (0 - 255, default = 254).
* `mode` - Select PTP mode. Valid values: `transparent-e2e`, `transparent-p2p`.

* `name` - Profile name.
* `pdelay_req_interval` - Configure PTP peer delay request interval. Valid values: `1sec`, `2sec`, `4sec`, `8sec`, `16sec`, `32sec`.

* `ptp_profile` - Configure PTP power profile. Valid values: `C37.238-2017`.

* `transport` - Configure PTP transport mode. Valid values: `l2-mcast`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController PtpProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_ptp_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

