---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_profileprotocoloptions_mailsignature"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure Mail signature.
---

# fmgdevice_firewall_profileprotocoloptions_mailsignature
<i>This object will be purged after policy copy and install.</i> Configure Mail signature.

~> This resource is a sub resource for variable `mail_signature` of resource `fmgdevice_firewall_profileprotocoloptions`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile_protocol_options` - Profile Protocol Options.

* `signature` - Email signature to be added to outgoing email (if the signature contains spaces, enclose with quotation marks).
* `status` - Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall ProfileProtocolOptionsMailSignature can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile_protocol_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_profileprotocoloptions_mailsignature.labelname FirewallProfileProtocolOptionsMailSignature
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

