---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_acl_ingress_classifier"
description: |-
  ACL classifiers.
---

# fmgdevice_switchcontroller_acl_ingress_classifier
ACL classifiers.

~> This resource is a sub resource for variable `classifier` of resource `fmgdevice_switchcontroller_acl_ingress`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_acl_ingress_classifier" "trname" {
  dst_ip_prefix = ["your own value"]
  dst_mac       = "your own value"
  src_ip_prefix = ["your own value"]
  src_mac       = "your own value"
  vlan          = 10
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `ingress` - Ingress.

* `dst_ip_prefix` - Destination IP address to be matched.
* `dst_mac` - Destination MAC address to be matched.
* `src_ip_prefix` - Source IP address to be matched.
* `src_mac` - Source MAC address to be matched.
* `vlan` - VLAN ID to be matched.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController AclIngressClassifier can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "ingress=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_acl_ingress_classifier.labelname SwitchControllerAclIngressClassifier
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

