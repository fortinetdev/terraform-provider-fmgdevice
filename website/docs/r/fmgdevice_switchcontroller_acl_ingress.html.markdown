---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_acl_ingress"
description: |-
  Configure ingress ACL policies to be applied on managed FortiSwitch ports.
---

# fmgdevice_switchcontroller_acl_ingress
Configure ingress ACL policies to be applied on managed FortiSwitch ports.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `action`: `fmgdevice_switchcontroller_acl_ingress_action`
>- `classifier`: `fmgdevice_switchcontroller_acl_ingress_classifier`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_acl_ingress" "trname" {
  action {
    count = "disable"
    drop  = "disable"
  }

  classifier {
    dst_ip_prefix = ["your own value"]
    dst_mac       = "your own value"
    src_ip_prefix = ["your own value"]
    src_mac       = "your own value"
    vlan          = 10
  }

  description = "your own value"
  fosid       = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `action` - Action. The structure of `action` block is documented below.
* `classifier` - Classifier. The structure of `classifier` block is documented below.
* `description` - Description for the ACL policy.
* `fosid` - ACL ID.

The `action` block supports:

* `count` - Enable/disable count. Valid values: `disable`, `enable`.

* `drop` - Enable/disable drop. Valid values: `disable`, `enable`.


The `classifier` block supports:

* `dst_ip_prefix` - Destination IP address to be matched.
* `dst_mac` - Destination MAC address to be matched.
* `src_ip_prefix` - Source IP address to be matched.
* `src_mac` - Source MAC address to be matched.
* `vlan` - VLAN ID to be matched.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController AclIngress can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_acl_ingress.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

