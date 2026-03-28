---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_profile_rule"
description: |-
  <i>This object will be purged after policy copy and install.</i> Set up DLP rules for this profile.
---

# fmgdevice_dlp_profile_rule
<i>This object will be purged after policy copy and install.</i> Set up DLP rules for this profile.

~> This resource is a sub resource for variable `rule` of resource `fmgdevice_dlp_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Action to take with content that this DLP profile matches. Valid values: `log-only`, `block`, `quarantine-ip`, `allow`.

* `archive` - Enable/disable DLP archiving. Valid values: `disable`, `enable`.

* `expiry` - Quarantine duration in days, hours, minutes (format = dddhhmm).
* `file_size` - Match files greater than or equal to this size (KB).
* `file_type` - Select the number of a DLP file pattern table to match.
* `filter_by` - Select the type of content to match. Valid values: `fingerprint`, `sensor`, `encrypted`, `none`, `mip`, `label`.

* `fosid` - ID.
* `label` - Select DLP label.
* `match_percentage` - Percentage of fingerprints in the fingerprint databases designated with the selected sensitivity to match.
* `name` - Filter name.
* `proto` - Check messages or files over one or more of these protocols. Valid values: `smtp`, `pop3`, `imap`, `http-post`, `http-get`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.

* `sensitivity` - Select a DLP file pattern sensitivity to match.
* `sensor` - Select DLP sensors.
* `severity` - Select the severity or threat level that matches this filter. Valid values: `info`, `low`, `medium`, `high`, `critical`.

* `type` - Select whether to check the content of messages (an email message) or files (downloaded files or email attachments). Valid values: `file`, `message`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Dlp ProfileRule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_profile_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

