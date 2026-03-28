---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_fssopolling"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure FSSO active directory servers for polling mode.
---

# fmgdevice_user_fssopolling
<i>This object will be purged after policy copy and install.</i> Configure FSSO active directory servers for polling mode.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `adgrp`: `fmgdevice_user_fssopolling_adgrp`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `_gui_meta` - _Gui_Meta.
* `adgrp` - Adgrp. The structure of `adgrp` block is documented below.
* `default_domain` - Default domain managed by this Active Directory server.
* `fosid` - Active Directory server ID.
* `ldap_server` - LDAP server name used in LDAP connection strings.
* `logon_history` - Number of hours of logon history to keep, 0 means keep all history.
* `password` - Password required to log into this Active Directory server.
* `polling_frequency` - Polling frequency (every 1 to 30 seconds).
* `port` - Port to communicate with this Active Directory server.
* `server` - Host name or IP address of the Active Directory server.
* `smb_ntlmv1_auth` - Enable/disable support of NTLMv1 for Samba authentication. Valid values: `disable`, `enable`.

* `smbv1` - Enable/disable support of SMBv1 for Samba. Valid values: `disable`, `enable`.

* `status` - Enable/disable polling for the status of this Active Directory server. Valid values: `disable`, `enable`.

* `user` - User name required to log into this Active Directory server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `adgrp` block supports:

* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

User FssoPolling can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_fssopolling.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

