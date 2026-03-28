---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ztna_serviceconnector"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_ztna_serviceconnector
<i>This object will be purged after policy copy and install.</i>

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `certificate` - Certificate.
* `connection_mode` - Connection-Mode. Valid values: `forward`, `reverse`.

* `encryption` - Encryption. Valid values: `disable`, `enable`.

* `forward_address` - Forward-Address.
* `forward_destination_cn` - Forward-Destination-Cn.
* `forward_port` - Forward-Port.
* `health_check_interval` - Health-Check-Interval.
* `log` - Log. Valid values: `disable`, `enable`.

* `name` - Name.
* `relay_dev_info` - Relay-Dev-Info. Valid values: `disable`, `enable`.

* `relay_user_info` - Relay-User-Info. Valid values: `disable`, `enable`.

* `ssl_max_version` - Ssl-Max-Version. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_min_version` - Ssl-Min-Version. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `status` - Status. Valid values: `disable`, `enable`.

* `trusted_ca` - Trusted-Ca.
* `url_map` - Url-Map.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ztna ServiceConnector can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ztna_serviceconnector.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

