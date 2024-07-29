---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_authportal"
description: |-
  Configure firewall authentication portals.
---

# fmgdevice_firewall_authportal
Configure firewall authentication portals.

## Example Usage

```hcl
resource "fmgdevice_firewall_authportal" "trname" {
  groups               = ["your own value"]
  identity_based_route = ["your own value"]
  portal_addr          = "your own value"
  portal_addr6         = "your own value"
  proxy_auth           = "enable"
  device_name          = var.device_name # not required if setting is at provider
  device_vdom          = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `groups` - Firewall user groups permitted to authenticate through this portal. Separate group names with spaces.
* `identity_based_route` - Name of the identity-based route that applies to this portal.
* `portal_addr` - Address (or FQDN) of the authentication portal.
* `portal_addr6` - IPv6 address (or FQDN) of authentication portal.
* `proxy_auth` - Enable/disable authentication by proxy daemon (default = disable). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall AuthPortal can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_authportal.labelname FirewallAuthPortal
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

