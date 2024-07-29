---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_snmpuser"
description: |-
  Configure FortiSwitch SNMP v3 users globally.
---

# fmgdevice_switchcontroller_snmpuser
Configure FortiSwitch SNMP v3 users globally.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_snmpuser" "trname" {
  auth_proto  = "sha512"
  auth_pwd    = ["your own value"]
  name        = "your own value"
  priv_proto  = "aes128"
  priv_pwd    = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_proto` - Authentication protocol. Valid values: `md5`, `sha`, `sha1`, `sha256`, `sha384`, `sha512`, `sha224`.

* `auth_pwd` - Password for authentication protocol.
* `name` - SNMP user name.
* `priv_proto` - Privacy (encryption) protocol. Valid values: `des`, `aes`, `aes128`, `aes192`, `aes256`, `aes192c`, `aes256c`.

* `priv_pwd` - Password for privacy (encryption) protocol.
* `queries` - Enable/disable SNMP queries for this user. Valid values: `disable`, `enable`.

* `query_port` - SNMPv3 query port (default = 161).
* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController SnmpUser can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_snmpuser.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

