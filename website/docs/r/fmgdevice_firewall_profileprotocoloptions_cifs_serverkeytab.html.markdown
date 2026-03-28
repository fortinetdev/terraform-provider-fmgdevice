---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_profileprotocoloptions_cifs_serverkeytab"
description: |-
  <i>This object will be purged after policy copy and install.</i> Server keytab.
---

# fmgdevice_firewall_profileprotocoloptions_cifs_serverkeytab
<i>This object will be purged after policy copy and install.</i> Server keytab.

~> This resource is a sub resource for variable `server_keytab` of resource `fmgdevice_firewall_profileprotocoloptions_cifs`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile_protocol_options` - Profile Protocol Options.

* `keytab` - Base64 encoded keytab file containing credential of the server.
* `password` - Password for keytab.
* `principal` - Service principal. For example, host/cifsserver.example.com@example.com.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{principal}}.

## Import

Firewall ProfileProtocolOptionsCifsServerKeytab can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile_protocol_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_profileprotocoloptions_cifs_serverkeytab.labelname {{principal}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

