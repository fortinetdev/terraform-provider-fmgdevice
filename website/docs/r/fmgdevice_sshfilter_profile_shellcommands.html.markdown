---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_sshfilter_profile_shellcommands"
description: |-
  <i>This object will be purged after policy copy and install.</i> SSH command filter.
---

# fmgdevice_sshfilter_profile_shellcommands
<i>This object will be purged after policy copy and install.</i> SSH command filter.

~> This resource is a sub resource for variable `shell_commands` of resource `fmgdevice_sshfilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Action to take for SSH shell command matches. Valid values: `block`, `allow`.

* `alert` - Enable/disable alert. Valid values: `disable`, `enable`.

* `fosid` - Id.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `pattern` - SSH shell command pattern.
* `severity` - Log severity. Valid values: `low`, `medium`, `high`, `critical`.

* `type` - Matching type. Valid values: `regex`, `simple`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SshFilter ProfileShellCommands can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_sshfilter_profile_shellcommands.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

