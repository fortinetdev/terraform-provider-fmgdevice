---
subcategory: "ZTNA"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ztna_connectoredge"
description: |-
  Configure ZTNA connector edge.
---

# fmgdevice_ztna_connectoredge
Configure ZTNA connector edge.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `interface` - Connector edge interface.
* `port` - Connector service edge port(1-65535, default 8443).
* `server_cert` - Server certificate for mTLS.
* `ssl_max_version` - Highest TLS version accepted by server. Valid values: `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_min_version` - Lowest TLS version accepted by server. Valid values: `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `status` - Connector service edge status. Valid values: `disable`, `enable`.

* `trusted_client_ca` - CA certificate used for client certificate verification.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ztna ConnectorEdge can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ztna_connectoredge.labelname ZtnaConnectorEdge
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

