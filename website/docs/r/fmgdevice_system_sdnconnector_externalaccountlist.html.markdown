---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sdnconnector_externalaccountlist"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure AWS external account list.
---

# fmgdevice_system_sdnconnector_externalaccountlist
<i>This object will be purged after policy copy and install.</i> Configure AWS external account list.

~> This resource is a sub resource for variable `external_account_list` of resource `fmgdevice_system_sdnconnector`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `sdn_connector` - Sdn Connector.

* `external_id` - AWS external ID.
* `region_list` - AWS region name list.
* `role_arn` - AWS role ARN to assume.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{role_arn}}.

## Import

System SdnConnectorExternalAccountList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "sdn_connector=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sdnconnector_externalaccountlist.labelname {{role_arn}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

