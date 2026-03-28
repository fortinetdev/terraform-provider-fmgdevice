---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_krbkeytab"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure Kerberos keytab entries.
---

# fmgdevice_user_krbkeytab
<i>This object will be purged after policy copy and install.</i> Configure Kerberos keytab entries.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `keytab` - Base64 coded keytab file containing a pre-shared key.
* `ldap_server` - LDAP server name(s).
* `name` - Kerberos keytab entry name.
* `pac_data` - Enable/disable parsing PAC data in the ticket. Valid values: `disable`, `enable`.

* `password` - Password for keytab.
* `principal` - Kerberos service principal. For example, HTTP/myfgt.example.com@example.com.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User KrbKeytab can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_krbkeytab.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

