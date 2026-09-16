---
subcategory: "Switch Controller"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_securitypolicy_admin"
description: |-
  Configure fortiswitch's admin security-policy.
---

# fmgdevice_switchcontroller_securitypolicy_admin
Configure fortiswitch's admin security-policy.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auto` - Automatically set based on the host ip connected via the Fortilink interface. Valid values: `disable`, `enable`.

* `ip6_trusthost1` - Trusted IPv6 host.
* `ip6_trusthost10` - Trusted IPv6 host.
* `ip6_trusthost2` - Trusted IPv6 host.
* `ip6_trusthost3` - Trusted IPv6 host.
* `ip6_trusthost4` - Trusted IPv6 host.
* `ip6_trusthost5` - Trusted IPv6 host.
* `ip6_trusthost6` - Trusted IPv6 host.
* `ip6_trusthost7` - Trusted IPv6 host.
* `ip6_trusthost8` - Trusted IPv6 host.
* `ip6_trusthost9` - Trusted IPv6 host.
* `name` - Policy name.
* `trusthost1` - Trusted IPv4 host.
* `trusthost10` - Trusted IPv4 host.
* `trusthost2` - Trusted IPv4 host.
* `trusthost3` - Trusted IPv4 host.
* `trusthost4` - Trusted IPv4 host.
* `trusthost5` - Trusted IPv4 host.
* `trusthost6` - Trusted IPv4 host.
* `trusthost7` - Trusted IPv4 host.
* `trusthost8` - Trusted IPv4 host.
* `trusthost9` - Trusted IPv4 host.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController SecurityPolicyAdmin can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_securitypolicy_admin.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

