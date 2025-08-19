---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ztna_webproxy_apigateway_quic"
description: |-
  QUIC setting.
---

# fmgdevice_ztna_webproxy_apigateway_quic
QUIC setting.

~> This resource is a sub resource for variable `quic` of resource `fmgdevice_ztna_webproxy_apigateway`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `web_proxy` - Web Proxy.
* `api_gateway` - Api Gateway.

* `ack_delay_exponent` - ACK delay exponent (1 - 20, default = 3).
* `active_connection_id_limit` - Active connection ID limit (1 - 8, default = 2).
* `active_migration` - Enable/disable active migration (default = disable). Valid values: `disable`, `enable`.

* `grease_quic_bit` - Enable/disable grease QUIC bit (default = enable). Valid values: `disable`, `enable`.

* `max_ack_delay` - Maximum ACK delay in milliseconds (1 - 16383, default = 25).
* `max_datagram_frame_size` - Maximum datagram frame size in bytes (1 - 1500, default = 1500).
* `max_idle_timeout` - Maximum idle timeout milliseconds (1 - 60000, default = 30000).
* `max_udp_payload_size` - Maximum UDP payload size in bytes (1200 - 1500, default = 1500).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ztna WebProxyApiGatewayQuic can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "web_proxy=YOUR_VALUE", "api_gateway=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ztna_webproxy_apigateway_quic.labelname ZtnaWebProxyApiGatewayQuic
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

