---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_smsserver"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure SMS server for sending SMS messages to support user authentication.
---

# fmgdevice_system_smsserver
<i>This object will be purged after policy copy and install.</i> Configure SMS server for sending SMS messages to support user authentication.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `mail_server` - Email-to-SMS server domain name.
* `name` - Name of SMS server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SmsServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_smsserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

