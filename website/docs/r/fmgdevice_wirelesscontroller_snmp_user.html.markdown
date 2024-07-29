---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_snmp_user"
description: |-
  SNMP User Configuration.
---

# fmgdevice_wirelesscontroller_snmp_user
SNMP User Configuration.

~> This resource is a sub resource for variable `user` of resource `fmgdevice_wirelesscontroller_snmp`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_snmp_user" "trname" {
  auth_proto   = "md5"
  auth_pwd     = ["your own value"]
  name         = "your own value"
  notify_hosts = ["your own value"]
  priv_proto   = "aes256cisco"
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_proto` - Authentication protocol. Valid values: `md5`, `sha`.

* `auth_pwd` - Password for authentication protocol.
* `name` - SNMP user name.
* `notify_hosts` - Configure SNMP User Notify Hosts.
* `priv_proto` - Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.

* `priv_pwd` - Password for privacy (encryption) protocol.
* `queries` - Enable/disable SNMP queries for this user. Valid values: `disable`, `enable`.

* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.

* `status` - SNMP user enable. Valid values: `disable`, `enable`.

* `trap_status` - Enable/disable traps for this SNMP user. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController SnmpUser can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_snmp_user.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

