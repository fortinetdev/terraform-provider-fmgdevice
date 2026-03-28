---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_externalidentityprovider"
description: |-
  Configure external identity provider.
---

# fmgdevice_user_externalidentityprovider
Configure external identity provider.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `group_attr_name` - Group attribute name in authentication query.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `name` - External identity provider name.
* `port` - External identity provider service port number (0 to use default).
* `server_identity_check` - Enable/disable server's identity check against its certificate and subject alternative name(s). Valid values: `disable`, `enable`.

* `source_ip` - Use this IPv4/v6 address to connect to the external identity provider.
* `timeout` - Connection timeout value in seconds (default=5).
* `type` - External identity provider type. Valid values: `ms-graph`.

* `url` - Url.
* `user_attr_name` - User attribute name in authentication query.
* `version` - External identity API version. Valid values: `beta`, `v1.0`.

* `vrf_select` - VRF ID used for connection to server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User ExternalIdentityProvider can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_externalidentityprovider.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

