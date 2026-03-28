---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ssl_keyringlist"
description: |-
  Device FirewallSslKeyringList
---

# fmgdevice_firewall_ssl_keyringlist
Device FirewallSslKeyringList

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `name` - Name.
* `uuid` - Uuid.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall SslKeyringList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ssl_keyringlist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

