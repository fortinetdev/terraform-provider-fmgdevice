---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_interfacepolicy6"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure IPv6 interface policies.
---

# fmgdevice_firewall_interfacepolicy6
<i>This object will be purged after policy copy and install.</i> Configure IPv6 interface policies.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `address_type` - Address-Type. Valid values: `ipv4`, `ipv6`.

* `application_list` - Application list name.
* `application_list_status` - Enable/disable application control. Valid values: `disable`, `enable`.

* `av_profile` - Antivirus profile.
* `av_profile_status` - Enable/disable antivirus. Valid values: `disable`, `enable`.

* `casb_profile` - CASB profile.
* `casb_profile_status` - Enable/disable CASB. Valid values: `disable`, `enable`.

* `comments` - Comments.
* `dlp_sensor` - DLP sensor name.
* `dlp_sensor_status` - Enable/disable DLP. Valid values: `disable`, `enable`.

* `dlp_profile` - DLP profile name.
* `dlp_profile_status` - Enable/disable DLP. Valid values: `disable`, `enable`.

* `dsri` - Enable/disable DSRI. Valid values: `disable`, `enable`.

* `dstaddr6` - IPv6 address object to limit traffic monitoring to network traffic sent to the specified address or range.
* `emailfilter_profile` - Email filter profile.
* `emailfilter_profile_status` - Enable/disable email filter. Valid values: `disable`, `enable`.

* `interface` - Monitored interface name from available interfaces.
* `ips_sensor` - IPS sensor name.
* `ips_sensor_status` - Enable/disable IPS. Valid values: `disable`, `enable`.

* `label` - Label.
* `logtraffic` - Logging type to be used in this policy (Options: all | utm | disable, Default: utm). Valid values: `disable`, `all`, `utm`.

* `policyid` - Policy ID (0 - 4294967295).
* `service6` - Service name.
* `srcaddr6` - IPv6 address object to limit traffic monitoring to network traffic sent from the specified address or range.
* `status` - Enable/disable this policy. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `webfilter_profile` - Web filter profile.
* `webfilter_profile_status` - Enable/disable web filtering. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall InterfacePolicy6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_interfacepolicy6.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

