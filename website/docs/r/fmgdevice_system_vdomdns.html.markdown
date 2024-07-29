---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_vdomdns"
description: |-
  Configure DNS servers for a non-management VDOM.
---

# fmgdevice_system_vdomdns
Configure DNS servers for a non-management VDOM.

## Example Usage

```hcl
resource "fmgdevice_system_vdomdns" "trname" {
  dns_over_tls            = "enforce"
  alt_primary             = "your own value"
  alt_secondary           = "your own value"
  interface               = ["port2"]
  interface_select_method = "auto"
  device_name             = var.device_name # not required if setting is at provider
  device_vdom             = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `dns_over_tls` - Enable/disable/enforce DNS over TLS. Valid values: `disable`, `enable`, `enforce`.

* `alt_primary` - Alternate primary DNS server. This is not used as a failover DNS server.
* `alt_secondary` - Alternate secondary DNS server. This is not used as a failover DNS server.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ip6_primary` - Primary IPv6 DNS server IP address for the VDOM.
* `ip6_secondary` - Secondary IPv6 DNS server IP address for the VDOM.
* `primary` - Primary DNS server IP address for the VDOM.
* `protocol` - DNS transport protocols. Valid values: `cleartext`, `dot`, `doh`.

* `secondary` - Secondary DNS server IP address for the VDOM.
* `server_hostname` - DNS server host name list.
* `server_select_method` - Specify how configured servers are prioritized. Valid values: `least-rtt`, `failover`.

* `source_ip` - Source IP for communications with the DNS server.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `vdom_dns` - Enable/disable configuring DNS servers for the current VDOM. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System VdomDns can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_vdomdns.labelname SystemVdomDns
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

