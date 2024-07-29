---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wanopt_contentdeliverynetworkrule_rules_contentid"
description: |-
  Content ID settings.
---

# fmgdevice_wanopt_contentdeliverynetworkrule_rules_contentid
Content ID settings.

~> This resource is a sub resource for variable `content_id` of resource `fmgdevice_wanopt_contentdeliverynetworkrule_rules`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wanopt_contentdeliverynetworkrule_rules_contentid" "trname" {
  end_direction   = "backward"
  end_skip        = 10
  end_str         = "your own value"
  range_str       = "your own value"
  start_direction = "backward"
  device_name     = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `content_delivery_network_rule` - Content Delivery Network Rule.
* `rules` - Rules.

* `end_direction` - Search direction from end-str match. Valid values: `forward`, `backward`.

* `end_skip` - Number of characters in URL to skip after end-str has been matched.
* `end_str` - String from which to end search.
* `range_str` - Name of content ID within the start string and end string.
* `start_direction` - Search direction from start-str match. Valid values: `forward`, `backward`.

* `start_skip` - Number of characters in URL to skip after start-str has been matched.
* `start_str` - String from which to start search.
* `target` - Option in HTTP header or URL parameter to match. Valid values: `path`, `parameter`, `referrer`, `youtube-map`, `youtube-id`, `youku-id`, `hls-manifest`, `dash-manifest`, `hls-fragment`, `dash-fragment`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Wanopt ContentDeliveryNetworkRuleRulesContentId can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "content_delivery_network_rule=YOUR_VALUE", "rules=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wanopt_contentdeliverynetworkrule_rules_contentid.labelname WanoptContentDeliveryNetworkRuleRulesContentId
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

