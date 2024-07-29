---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_vap_portalmessageoverrides"
description: |-
  Individual message overrides.
---

# fmgdevice_wirelesscontroller_vap_portalmessageoverrides
Individual message overrides.

~> This resource is a sub resource for variable `portal_message_overrides` of resource `fmgdevice_wirelesscontroller_vap`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_vap_portalmessageoverrides" "trname" {
  auth_disclaimer_page   = "your own value"
  auth_login_failed_page = "your own value"
  auth_login_page        = "your own value"
  auth_reject_page       = "your own value"
  device_name            = var.device_name # not required if setting is at provider
  device_vdom            = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `vap` - Vap.

* `auth_disclaimer_page` - Override auth-disclaimer-page message with message from portal-message-overrides group.
* `auth_login_failed_page` - Override auth-login-failed-page message with message from portal-message-overrides group.
* `auth_login_page` - Override auth-login-page message with message from portal-message-overrides group.
* `auth_reject_page` - Override auth-reject-page message with message from portal-message-overrides group.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController VapPortalMessageOverrides can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "vap=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_vap_portalmessageoverrides.labelname WirelessControllerVapPortalMessageOverrides
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

