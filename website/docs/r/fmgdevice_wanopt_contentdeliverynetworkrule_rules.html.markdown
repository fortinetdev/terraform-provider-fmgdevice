---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wanopt_contentdeliverynetworkrule_rules"
description: |-
  WAN optimization content delivery network rule entries.
---

# fmgdevice_wanopt_contentdeliverynetworkrule_rules
WAN optimization content delivery network rule entries.

~> This resource is a sub resource for variable `rules` of resource `fmgdevice_wanopt_contentdeliverynetworkrule`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `content_id`: `fmgdevice_wanopt_contentdeliverynetworkrule_rules_contentid`
>- `match_entries`: `fmgdevice_wanopt_contentdeliverynetworkrule_rules_matchentries`
>- `skip_entries`: `fmgdevice_wanopt_contentdeliverynetworkrule_rules_skipentries`



## Example Usage

```hcl
resource "fmgdevice_wanopt_contentdeliverynetworkrule_rules" "trname" {
  content_id {
    end_direction   = "forward"
    end_skip        = 10
    end_str         = "your own value"
    range_str       = "your own value"
    start_direction = "backward"
    start_skip      = 10
    start_str       = "your own value"
    target          = "youtube-map"
  }

  match_entries {
    fosid   = 10
    pattern = ["your own value"]
    target  = "path"
  }

  match_mode = "any"
  name       = "your own value"
  skip_entries {
    fosid   = 10
    pattern = ["your own value"]
    target  = "referrer"
  }

  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `content_delivery_network_rule` - Content Delivery Network Rule.

* `content_id` - Content-Id. The structure of `content_id` block is documented below.
* `match_entries` - Match-Entries. The structure of `match_entries` block is documented below.
* `match_mode` - Match criteria for collecting content ID. Valid values: `any`, `all`.

* `name` - WAN optimization content delivery network rule name.
* `skip_entries` - Skip-Entries. The structure of `skip_entries` block is documented below.
* `skip_rule_mode` - Skip mode when evaluating skip-rules. Valid values: `any`, `all`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `content_id` block supports:

* `end_direction` - Search direction from end-str match. Valid values: `forward`, `backward`.

* `end_skip` - Number of characters in URL to skip after end-str has been matched.
* `end_str` - String from which to end search.
* `range_str` - Name of content ID within the start string and end string.
* `start_direction` - Search direction from start-str match. Valid values: `forward`, `backward`.

* `start_skip` - Number of characters in URL to skip after start-str has been matched.
* `start_str` - String from which to start search.
* `target` - Option in HTTP header or URL parameter to match. Valid values: `path`, `parameter`, `referrer`, `youtube-map`, `youtube-id`, `youku-id`, `hls-manifest`, `dash-manifest`, `hls-fragment`, `dash-fragment`.


The `match_entries` block supports:

* `id` - Rule ID.
* `pattern` - Pattern string for matching target (Referrer or URL pattern). For example, a, a*c, *a*, a*c*e, and *.
* `target` - Option in HTTP header or URL parameter to match. Valid values: `path`, `parameter`, `referrer`, `youtube-map`, `youtube-id`, `youku-id`.


The `skip_entries` block supports:

* `id` - Rule ID.
* `pattern` - Pattern string for matching target (Referrer or URL pattern). For example, a, a*c, *a*, a*c*e, and *.
* `target` - Option in HTTP header or URL parameter to match. Valid values: `path`, `parameter`, `referrer`, `youtube-map`, `youtube-id`, `youku-id`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Wanopt ContentDeliveryNetworkRuleRules can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "content_delivery_network_rule=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wanopt_contentdeliverynetworkrule_rules.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

