---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_qosprofile"
description: |-
  Configure WiFi quality of service (QoS) profiles.
---

# fmgdevice_wirelesscontroller_qosprofile
Configure WiFi quality of service (QoS) profiles.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_qosprofile" "trname" {
  bandwidth_admission_control = "disable"
  bandwidth_capacity          = 10
  burst                       = "disable"
  call_admission_control      = "enable"
  call_capacity               = 10
  name                        = "your own value"
  device_name                 = var.device_name # not required if setting is at provider
  device_vdom                 = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `bandwidth_admission_control` - Enable/disable WMM bandwidth admission control. Valid values: `disable`, `enable`.

* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `burst` - Enable/disable client rate burst. Valid values: `disable`, `enable`.

* `call_admission_control` - Enable/disable WMM call admission control. Valid values: `disable`, `enable`.

* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones allowed (0 - 60, default = 10).
* `comment` - Comment.
* `downlink` - Maximum downlink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).
* `downlink_sta` - Maximum downlink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).
* `dscp_wmm_be` - DSCP mapping for best effort access (default = 0 24).
* `dscp_wmm_bk` - DSCP mapping for background access (default = 8 16).
* `dscp_wmm_mapping` - Enable/disable Differentiated Services Code Point (DSCP) mapping. Valid values: `disable`, `enable`.

* `dscp_wmm_vi` - DSCP mapping for video access (default = 32 40).
* `dscp_wmm_vo` - DSCP mapping for voice access (default = 48 56).
* `name` - WiFi QoS profile name.
* `uplink` - Maximum uplink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).
* `uplink_sta` - Maximum uplink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).
* `wmm` - Enable/disable WiFi multi-media (WMM) control. Valid values: `disable`, `enable`.

* `wmm_be_dscp` - DSCP marking for best effort access (default = 0).
* `wmm_bk_dscp` - DSCP marking for background access (default = 8).
* `wmm_dscp_marking` - Enable/disable WMM Differentiated Services Code Point (DSCP) marking. Valid values: `disable`, `enable`.

* `wmm_uapsd` - Enable/disable WMM Unscheduled Automatic Power Save Delivery (U-APSD) power save mode. Valid values: `disable`, `enable`.

* `wmm_vi_dscp` - DSCP marking for video access (default = 32).
* `wmm_vo_dscp` - DSCP marking for voice access (default = 48).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController QosProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_qosprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

