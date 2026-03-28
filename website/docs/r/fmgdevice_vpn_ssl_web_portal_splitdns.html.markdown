---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ssl_web_portal_splitdns"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_vpn_ssl_web_portal_splitdns
<i>This object will be purged after policy copy and install.</i>

~> This resource is a sub resource for variable `split_dns` of resource `fmgdevice_vpn_ssl_web_portal`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `portal` - Portal.

* `dns_server1` - Dns-Server1.
* `dns_server2` - Dns-Server2.
* `domains` - Domains.
* `fosid` - Id.
* `ipv6_dns_server1` - Ipv6-Dns-Server1.
* `ipv6_dns_server2` - Ipv6-Dns-Server2.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Vpn SslWebPortalSplitDns can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "portal=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ssl_web_portal_splitdns.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

