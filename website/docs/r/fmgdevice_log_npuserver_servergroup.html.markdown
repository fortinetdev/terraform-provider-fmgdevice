---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_npuserver_servergroup"
description: |-
  <i>This object will be purged after policy copy and install.</i> create server group.
---

# fmgdevice_log_npuserver_servergroup
<i>This object will be purged after policy copy and install.</i> create server group.

~> This resource is a sub resource for variable `server_group` of resource `fmgdevice_log_npuserver`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `group_name` - server group name.
* `log_format` - Set the log format Valid values: `syslog`, `netflow`.

* `log_gen_event` - Enable/disbale generating event for Per-Mapping log Valid values: `disable`, `enable`.

* `log_mode` - Set the log mode Valid values: `per-session`, `per-nat-mapping`, `per-session-ending`.

* `log_tx_mode` - Configure log transmit mode. Valid values: `multicast`, `roundrobin`.

* `log_user_info` - Enable/disbale logging user information. Valid values: `disable`, `enable`.

* `server_number` - server number in this group.
* `server_start_id` - the start id of the continuous server series in this group,[1,16].
* `sw_log_flags` - Set flags for software logging via driver. Valid values: `tcp-udp-only`, `enable-all-log`, `disable-all-log`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{group_name}}.

## Import

Log NpuServerServerGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_npuserver_servergroup.labelname {{group_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

