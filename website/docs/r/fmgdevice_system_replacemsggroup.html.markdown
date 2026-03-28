---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_replacemsggroup"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure replacement message groups.
---

# fmgdevice_system_replacemsggroup
<i>This object will be purged after policy copy and install.</i> Configure replacement message groups.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `admin`: `fmgdevice_system_replacemsggroup_admin`
>- `alertmail`: `fmgdevice_system_replacemsggroup_alertmail`
>- `auth`: `fmgdevice_system_replacemsggroup_auth`
>- `automation`: `fmgdevice_system_replacemsggroup_automation`
>- `custom_message`: `fmgdevice_system_replacemsggroup_custommessage`
>- `device_detection_portal`: `fmgdevice_system_replacemsggroup_devicedetectionportal`
>- `fortiguard_wf`: `fmgdevice_system_replacemsggroup_fortiguardwf`
>- `ftp`: `fmgdevice_system_replacemsggroup_ftp`
>- `http`: `fmgdevice_system_replacemsggroup_http`
>- `icap`: `fmgdevice_system_replacemsggroup_icap`
>- `mail`: `fmgdevice_system_replacemsggroup_mail`
>- `nac_quar`: `fmgdevice_system_replacemsggroup_nacquar`
>- `nntp`: `fmgdevice_system_replacemsggroup_nntp`
>- `spam`: `fmgdevice_system_replacemsggroup_spam`
>- `sslvpn`: `fmgdevice_system_replacemsggroup_sslvpn`
>- `traffic_quota`: `fmgdevice_system_replacemsggroup_trafficquota`
>- `utm`: `fmgdevice_system_replacemsggroup_utm`
>- `webproxy`: `fmgdevice_system_replacemsggroup_webproxy`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `admin` - Admin. The structure of `admin` block is documented below.
* `alertmail` - Alertmail. The structure of `alertmail` block is documented below.
* `auth` - Auth. The structure of `auth` block is documented below.
* `automation` - Automation. The structure of `automation` block is documented below.
* `comment` - Comment.
* `custom_message` - Custom-Message. The structure of `custom_message` block is documented below.
* `device_detection_portal` - Device-Detection-Portal. The structure of `device_detection_portal` block is documented below.
* `fortiguard_wf` - Fortiguard-Wf. The structure of `fortiguard_wf` block is documented below.
* `ftp` - Ftp. The structure of `ftp` block is documented below.
* `group_type` - Group type. Valid values: `default`, `utm`, `auth`.

* `http` - Http. The structure of `http` block is documented below.
* `icap` - Icap. The structure of `icap` block is documented below.
* `mail` - Mail. The structure of `mail` block is documented below.
* `nac_quar` - Nac-Quar. The structure of `nac_quar` block is documented below.
* `name` - Group name.
* `nntp` - Nntp. The structure of `nntp` block is documented below.
* `spam` - Spam. The structure of `spam` block is documented below.
* `sslvpn` - Sslvpn. The structure of `sslvpn` block is documented below.
* `traffic_quota` - Traffic-Quota. The structure of `traffic_quota` block is documented below.
* `utm` - Utm. The structure of `utm` block is documented below.
* `webproxy` - Webproxy. The structure of `webproxy` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `admin` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `alertmail` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `id` - Id.
* `msg_type` - Message type.

The `auth` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `automation` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `custom_message` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `device_detection_portal` block supports:

* `buffer` - Buffer.
* `format` - Format. Valid values: `none`, `text`, `html`.

* `header` - Header. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Msg-Type.

The `fortiguard_wf` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `ftp` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `http` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `icap` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `mail` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `nac_quar` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `id` - Id.
* `msg_type` - Message type.

The `nntp` block supports:

* `buffer` - Buffer.
* `format` - Format. Valid values: `none`, `text`, `html`.

* `header` - Header. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Msg-Type.

The `spam` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `sslvpn` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `traffic_quota` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `utm` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `webproxy` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ReplacemsgGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_replacemsggroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

