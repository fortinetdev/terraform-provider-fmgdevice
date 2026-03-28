---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dnsfilter_profile_dnstranslation"
description: |-
  <i>This object will be purged after policy copy and install.</i> DNS translation settings.
---

# fmgdevice_dnsfilter_profile_dnstranslation
<i>This object will be purged after policy copy and install.</i> DNS translation settings.

~> This resource is a sub resource for variable `dns_translation` of resource `fmgdevice_dnsfilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `addr_type` - DNS translation type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `dst` - IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
* `dst6` - IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
* `fosid` - ID.
* `netmask` - If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
* `prefix` - If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
* `src` - IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
* `src6` - IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
* `status` - Enable/disable this DNS translation entry. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Dnsfilter ProfileDnsTranslation can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dnsfilter_profile_dnstranslation.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

