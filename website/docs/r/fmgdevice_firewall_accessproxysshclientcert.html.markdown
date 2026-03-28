---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_accessproxysshclientcert"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure Access Proxy SSH client certificate.
---

# fmgdevice_firewall_accessproxysshclientcert
<i>This object will be purged after policy copy and install.</i> Configure Access Proxy SSH client certificate.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `cert_extension`: `fmgdevice_firewall_accessproxysshclientcert_certextension`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_ca` - Name of the SSH server public key authentication CA.
* `cert_extension` - Cert-Extension. The structure of `cert_extension` block is documented below.
* `name` - SSH client certificate name.
* `permit_agent_forwarding` - Enable/disable appending permit-agent-forwarding certificate extension. Valid values: `disable`, `enable`.

* `permit_port_forwarding` - Enable/disable appending permit-port-forwarding certificate extension. Valid values: `disable`, `enable`.

* `permit_pty` - Enable/disable appending permit-pty certificate extension. Valid values: `disable`, `enable`.

* `permit_user_rc` - Enable/disable appending permit-user-rc certificate extension. Valid values: `disable`, `enable`.

* `permit_x11_forwarding` - Enable/disable appending permit-x11-forwarding certificate extension. Valid values: `disable`, `enable`.

* `source_address` - Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `cert_extension` block supports:

* `critical` - Critical option. Valid values: `no`, `yes`.

* `data` - Data of certificate extension.
* `name` - Name of certificate extension.
* `type` - Type of certificate extension. Valid values: `fixed`, `user`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall AccessProxySshClientCert can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_accessproxysshclientcert.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

