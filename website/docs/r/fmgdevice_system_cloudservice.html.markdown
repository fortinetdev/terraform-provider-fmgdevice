---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_cloudservice"
description: |-
  Configure system cloud service.
---

# fmgdevice_system_cloudservice
Configure system cloud service.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `gck_access_token_lifetime` - Lifetime of automatically generated access tokens in minutes (default is 60 minutes).
* `gck_keyid` - Key id, also referred as "kid".
* `gck_private_key` - Service account private key in PEM format (e.g. "-----BEGIN PRIVATE KEY-----\ ...").
* `gck_service_account` - Service account (e.g. "account-id@sampledomain.com").
* `name` - Name.
* `traffic_vdom` - Vdom used to communicate with cloud service.
* `vendor` - Cloud service vendor. Valid values: `unknown`, `google-cloud-kms`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CloudService can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_cloudservice.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

