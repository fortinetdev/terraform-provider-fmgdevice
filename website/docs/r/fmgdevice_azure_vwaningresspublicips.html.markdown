---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_azure_vwaningresspublicips"
description: |-
  Display Azure vWAN SLB ingress public IPs.
---

# fmgdevice_azure_vwaningresspublicips
Display Azure vWAN SLB ingress public IPs.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `ip` - Public IP address.
* `name` - Name of public IP.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Azure VwanIngressPublicIps can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_azure_vwaningresspublicips.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

