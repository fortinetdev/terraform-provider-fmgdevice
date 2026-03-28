---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_sshfilter_profile"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure SSH filter profile.
---

# fmgdevice_sshfilter_profile
<i>This object will be purged after policy copy and install.</i> Configure SSH filter profile.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `shell_commands`: `fmgdevice_sshfilter_profile_shellcommands`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `block` - SSH blocking options. Valid values: `x11`, `shell`, `exec`, `port-forward`, `tun-forward`, `sftp`, `unknown`, `scp`.

* `default_command_log` - Enable/disable logging unmatched shell commands. Valid values: `disable`, `enable`.

* `log` - SSH logging options. Valid values: `x11`, `shell`, `exec`, `port-forward`, `tun-forward`, `sftp`, `unknown`, `scp`.

* `name` - SSH filter profile name.
* `shell_commands` - Shell-Commands. The structure of `shell_commands` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `shell_commands` block supports:

* `action` - Action to take for SSH shell command matches. Valid values: `block`, `allow`.

* `alert` - Enable/disable alert. Valid values: `disable`, `enable`.

* `id` - Id.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `pattern` - SSH shell command pattern.
* `severity` - Log severity. Valid values: `low`, `medium`, `high`, `critical`.

* `type` - Matching type. Valid values: `regex`, `simple`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SshFilter Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_sshfilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

