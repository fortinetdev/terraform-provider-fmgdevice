---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_pop3"
description: |-
  <i>This object will be purged after policy copy and install.</i> POP3 server entry configuration.
---

# fmgdevice_user_pop3
<i>This object will be purged after policy copy and install.</i> POP3 server entry configuration.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `name` - POP3 server entry name.
* `port` - POP3 service port number.
* `secure` - SSL connection. Valid values: `none`, `starttls`, `pop3s`.

* `server` - Server domain name or IP address.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1-3`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Pop3 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_pop3.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

