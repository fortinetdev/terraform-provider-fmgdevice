---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ipsec_forticlient"
description: |-
  Configure FortiClient policy realm.
---

# fmgdevice_vpn_ipsec_forticlient
Configure FortiClient policy realm.

## Example Usage

```hcl
resource "fmgdevice_vpn_ipsec_forticlient" "trname" {
  phase2name    = ["your own value"]
  realm         = "your own value"
  status        = "enable"
  usergroupname = ["your own value"]
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `phase2name` - Phase 2 tunnel name that you defined in the FortiClient dialup configuration.
* `realm` - FortiClient realm name.
* `status` - Enable/disable this FortiClient configuration. Valid values: `disable`, `enable`.

* `usergroupname` - User group name for FortiClient users.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{realm}}.

## Import

Vpn IpsecForticlient can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ipsec_forticlient.labelname {{realm}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

