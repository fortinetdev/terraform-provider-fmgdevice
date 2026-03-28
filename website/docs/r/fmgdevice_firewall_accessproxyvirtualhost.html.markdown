---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_accessproxyvirtualhost"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure Access Proxy virtual hosts.
---

# fmgdevice_firewall_accessproxyvirtualhost
<i>This object will be purged after policy copy and install.</i> Configure Access Proxy virtual hosts.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `client_cert` - Enable/disable requesting client certificate. Valid values: `disable`, `enable`.

* `empty_cert_action` - Action for an empty client certificate. Valid values: `block`, `accept`, `accept-unmanageable`.

* `host` - The host name.
* `host_type` - Type of host pattern. Valid values: `sub-string`, `wildcard`.

* `name` - Virtual host name.
* `replacemsg_group` - Access-proxy-virtual-host replacement message override group.
* `ssl_certificate` - SSL certificates for this host.
* `user_agent_detect` - Enable/disable detecting device type by HTTP user-agent if no client certificate is provided. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall AccessProxyVirtualHost can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_accessproxyvirtualhost.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

