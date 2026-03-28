---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_videofilter_profile"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure VideoFilter profile.
---

# fmgdevice_videofilter_profile
<i>This object will be purged after policy copy and install.</i> Configure VideoFilter profile.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `filters`: `fmgdevice_videofilter_profile_filters`
>- `fortiguard_category`: `fmgdevice_videofilter_profile_fortiguardcategory`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `dailymotion` - Enable/disable Dailymotion video source. Valid values: `disable`, `enable`.

* `default_action` - Default-Action. Valid values: `block`, `monitor`, `allow`.

* `filters` - Filters. The structure of `filters` block is documented below.
* `fortiguard_category` - Fortiguard-Category. The structure of `fortiguard_category` block is documented below.
* `log` - Log. Valid values: `disable`, `enable`.

* `name` - Name.
* `replacemsg_group` - Replacement message group.
* `vimeo` - Enable/disable Vimeo video source. Valid values: `disable`, `enable`.

* `vimeo_restrict` - Vimeo-Restrict.
* `youtube` - Enable/disable YouTube video source. Valid values: `disable`, `enable`.

* `youtube_channel_filter` - Youtube-Channel-Filter.
* `youtube_restrict` - Youtube-Restrict. Valid values: `strict`, `none`, `moderate`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `filters` block supports:

* `action` - Video filter action. Valid values: `block`, `monitor`, `allow`.

* `category` - FortiGuard category ID.
* `channel` - Channel ID.
* `comment` - Comment.
* `id` - ID.
* `keyword` - Video filter keyword ID.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `type` - Filter type. Valid values: `category`, `channel`, `title`, `description`.


The `fortiguard_category` block supports:

* `filters` - Filters. The structure of `filters` block is documented below.

The `filters` block supports:

* `action` - Action. Valid values: `block`, `monitor`, `allow`.

* `category_id` - Category-Id.
* `id` - Id.
* `log` - Log. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Videofilter Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_videofilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

