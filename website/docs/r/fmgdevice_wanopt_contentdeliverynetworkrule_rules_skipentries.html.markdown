---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wanopt_contentdeliverynetworkrule_rules_skipentries"
description: |-
  List of entries to skip.
---

# fmgdevice_wanopt_contentdeliverynetworkrule_rules_skipentries
List of entries to skip.

~> This resource is a sub resource for variable `skip_entries` of resource `fmgdevice_wanopt_contentdeliverynetworkrule_rules`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wanopt_contentdeliverynetworkrule_rules_skipentries" "trname" {
  fosid       = 10
  pattern     = ["your own value"]
  target      = "youtube-map"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `content_delivery_network_rule` - Content Delivery Network Rule.
* `rules` - Rules.

* `fosid` - Rule ID.
* `pattern` - Pattern string for matching target (Referrer or URL pattern). For example, a, a*c, *a*, a*c*e, and *.
* `target` - Option in HTTP header or URL parameter to match. Valid values: `path`, `parameter`, `referrer`, `youtube-map`, `youtube-id`, `youku-id`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Wanopt ContentDeliveryNetworkRuleRulesSkipEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "content_delivery_network_rule=YOUR_VALUE", "rules=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wanopt_contentdeliverynetworkrule_rules_skipentries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

