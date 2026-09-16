---
subcategory: "ZTNA"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ztna_destination"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure ZTNA destination.
---

# fmgdevice_ztna_destination
<i>This object will be purged after policy copy and install.</i> Configure ZTNA destination.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `address` - Address or address group of the ZTNA destination.
* `conn_type` - Connection type. Valid values: `traffic-forwarding`, `ssh`.

* `domain` - Wildcard domain name of the ZTNA destination.
* `external_auth` - Enable/disable use of external browser as user-agent for SAML user authentication. Valid values: `disable`, `enable`.

* `mappedport` - Port for communicating with the real server.
* `name` - Destination name.
* `protocol` - Protocol type based on IANA numbers. Valid values: `TCP`, `UDP`, `ALL`.

* `saas_application` - SaaS application controlled by this ZTNA destination.
* `ssh_client_cert` - Configure access-proxy SSH client certificate profile.
* `ssh_host_key` - Configure host keys (one or more may be configured).
* `ssh_host_key_validation` - Enable/disable SSH host key validation. Valid values: `disable`, `enable`.

* `tunnel_encryption` - Tunnel encryption. Valid values: `disable`, `enable`.

* `type` - ZTNA destination type. Valid values: `on-premise`, `saas`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ztna Destination can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ztna_destination.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

