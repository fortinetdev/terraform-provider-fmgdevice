---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_nsxt_servicechain_serviceindex_move"
description: |-
  Move service index.
---

# fmgdevice_nsxt_servicechain_serviceindex_move
Move service index.

## Example Usage

```hcl
resource "fmgdevice_nsxt_servicechain_serviceindex_move" "trname" {
  service_chain = fmgdevice_nsxt_servicechain.trname1.fosid
  service_index = 11
  target        = 10
  option        = "before"
  depends_on    = [fmgdevice_nsxt_servicechain_serviceindex.trname1, fmgdevice_nsxt_servicechain_serviceindex.trname2]
  device_name   = var.device_name # not required if setting is at provider
}

resource "fmgdevice_nsxt_servicechain_serviceindex" "trname1" {
  fosid         = 10
  device_name   = var.device_name # not required if setting is at provider
  service_chain = fmgdevice_nsxt_servicechain.trname1.fosid
  depends_on    = [fmgdevice_nsxt_servicechain.trname1]
}

resource "fmgdevice_nsxt_servicechain_serviceindex" "trname2" {
  fosid         = 11
  device_name   = var.device_name # not required if setting is at provider
  service_chain = fmgdevice_nsxt_servicechain.trname1.fosid
  depends_on    = [fmgdevice_nsxt_servicechain.trname1]
}

resource "fmgdevice_nsxt_servicechain" "trname1" {
  fosid       = 12
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `service_chain` - Service Chain.
* `service_index` - Service Index.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the service index changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of service indexes.
