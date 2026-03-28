---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_icap_localserver_icapservice"
description: |-
  Device IcapLocalServerIcapService
---

# fmgdevice_icap_localserver_icapservice
Device IcapLocalServerIcapService

~> This resource is a sub resource for variable `icap_service` of resource `fmgdevice_icap_localserver`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `local_server` - Local Server.

* `av_profile` - Av-Profile.
* `dlp_profile` - Dlp-Profile.
* `dlp_sensor` - Dlp-Sensor.
* `extension_headers` - Extension-Headers. Valid values: `X-Virus-id`, `X-Infection-Found`, `X-Violation-Found`.

* `image_analyzer_profile` - Image-Analyzer-Profile.
* `name` - Name.
* `profile_protocol_options` - Profile-Protocol-Options.
* `service_id` - Service-Id.
* `webfilter_profile` - Webfilter-Profile.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{service_id}}.

## Import

Icap LocalServerIcapService can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "local_server=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_icap_localserver_icapservice.labelname {{service_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

