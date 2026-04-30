---
subcategory: "LLM"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_llm_proxy"
description: |-
  Device LlmProxy
---

# fmgdevice_llm_proxy
Device LlmProxy

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `hostname` - Hostname.
* `incoming_http_port` - Incoming-Http-Port.
* `incoming_https_port` - Incoming-Https-Port.
* `incoming_ip` - Incoming-Ip.
* `incoming_ip6` - Incoming-Ip6.
* `interface` - Interface.
* `ipv6_status` - Ipv6-Status. Valid values: `disable`, `enable`.

* `name` - Name.
* `protocol` - Protocol. Valid values: `http`, `https`.

* `srv_pool_max_concurrent_request` - Srv-Pool-Max-Concurrent-Request.
* `srv_pool_max_request` - Srv-Pool-Max-Request.
* `srv_pool_ttl` - Srv-Pool-Ttl.
* `ssl_certificate` - Ssl-Certificate.
* `ssl_max_version` - Ssl-Max-Version. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_min_version` - Ssl-Min-Version. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Llm Proxy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_llm_proxy.labelname LlmProxy
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

