---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sdnconnector_forwardingrule"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure GCP forwarding rule.
---

# fmgdevice_system_sdnconnector_forwardingrule
<i>This object will be purged after policy copy and install.</i> Configure GCP forwarding rule.

~> This resource is a sub resource for variable `forwarding_rule` of resource `fmgdevice_system_sdnconnector`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `sdn_connector` - Sdn Connector.

* `rule_name` - Forwarding rule name.
* `target` - Target instance name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{rule_name}}.

## Import

System SdnConnectorForwardingRule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "sdn_connector=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sdnconnector_forwardingrule.labelname {{rule_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

