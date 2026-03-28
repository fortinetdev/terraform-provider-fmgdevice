---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sdnconnector_compartmentlist"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure OCI compartment list.
---

# fmgdevice_system_sdnconnector_compartmentlist
<i>This object will be purged after policy copy and install.</i> Configure OCI compartment list.

~> This resource is a sub resource for variable `compartment_list` of resource `fmgdevice_system_sdnconnector`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `sdn_connector` - Sdn Connector.

* `compartment_id` - OCI compartment ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{compartment_id}}.

## Import

System SdnConnectorCompartmentList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "sdn_connector=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sdnconnector_compartmentlist.labelname {{compartment_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

