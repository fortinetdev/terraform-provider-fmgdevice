---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ldbmonitor"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure server load balancing health monitors.
---

# fmgdevice_firewall_ldbmonitor
<i>This object will be purged after policy copy and install.</i> Configure server load balancing health monitors.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `dns_match_ip` - Response IP expected from DNS server.
* `dns_protocol` - Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.

* `dns_request_domain` - Fully qualified domain name to resolve for the DNS probe.
* `http_get` - Request URI used to send a GET request to check the health of an HTTP server. Optionally provide a hostname before the first '/' and it will be used as the HTTP Host Header.
* `http_match` - String to match the value expected in response to an HTTP-GET request.
* `http_max_redirects` - The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
* `interval` - Time between health checks (5 - 65535 sec, default = 10).
* `name` - Monitor name.
* `port` - Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65535, default = 0).
* `retry` - Number health check attempts before the server is considered down (1 - 255, default = 3).
* `src_ip` - Source IP for ldb-monitor.
* `timeout` - Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
* `type` - Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP | HTTPS | DNS). Valid values: `ping`, `tcp`, `http`, `passive-sip`, `https`, `dns`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall LdbMonitor can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ldbmonitor.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

