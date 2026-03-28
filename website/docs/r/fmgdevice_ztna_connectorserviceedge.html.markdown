---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ztna_connectorserviceedge"
description: |-
  Device ZtnaConnectorServiceEdge
---

# fmgdevice_ztna_connectorserviceedge
Device ZtnaConnectorServiceEdge

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `interface` - Interface.
* `port` - Port.
* `server_cert` - Server-Cert.
* `ssl_max_version` - Ssl-Max-Version. Valid values: `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `status` - Status. Valid values: `disable`, `enable`.

* `trusted_client_ca` - Trusted-Client-Ca.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ztna ConnectorServiceEdge can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ztna_connectorserviceedge.labelname ZtnaConnectorServiceEdge
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

