---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_diameterfilter_profile"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure Diameter filter profiles.
---

# fmgdevice_diameterfilter_profile
<i>This object will be purged after policy copy and install.</i> Configure Diameter filter profiles.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `cmd_flags_reserve_set` - Action to be taken for messages with cmd flag reserve bits set. Valid values: `block`, `reset`, `monitor`, `allow`.

* `command_code_invalid` - Action to be taken for messages with invalid command code. Valid values: `block`, `reset`, `monitor`, `allow`.

* `command_code_range` - Valid range for command codes (0-16777215).
* `comment` - Comment.
* `log_packet` - Enable/disable packet log for triggered diameter settings. Valid values: `disable`, `enable`.

* `message_length_invalid` - Action to be taken for invalid message length. Valid values: `block`, `reset`, `monitor`, `allow`.

* `missing_request_action` - Action to be taken for answers without corresponding request. Valid values: `block`, `reset`, `monitor`, `allow`.

* `monitor_all_messages` - Enable/disable logging for all User Name and Result Code AVP messages. Valid values: `disable`, `enable`.

* `name` - Profile name.
* `protocol_version_invalid` - Action to be taken for invalid protocol version. Valid values: `block`, `reset`, `monitor`, `allow`.

* `request_error_flag_set` - Action to be taken for request messages with error flag set. Valid values: `block`, `reset`, `monitor`, `allow`.

* `track_requests_answers` - Enable/disable validation that each answer has a corresponding request. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

DiameterFilter Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_diameterfilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

