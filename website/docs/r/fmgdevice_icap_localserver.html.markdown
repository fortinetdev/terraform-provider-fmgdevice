---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_icap_localserver"
description: |-
  Device IcapLocalServer
---

# fmgdevice_icap_localserver
Device IcapLocalServer

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `icap_service`: `fmgdevice_icap_localserver_icapservice`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `icap_incoming_port` - Icap-Incoming-Port.
* `icap_incoming_ssl_port` - Icap-Incoming-Ssl-Port.
* `icap_server_id` - Icap-Server-Id.
* `icap_service` - Icap-Service. The structure of `icap_service` block is documented below.
* `incoming_ip` - Incoming-Ip.
* `incoming_ipv6` - Incoming-Ipv6.
* `interface` - Interface.
* `message_preview` - Message-Preview. Valid values: `disable`, `enable`.

* `secure_connection` - Secure-Connection. Valid values: `disable`, `enable`.

* `srcaddr` - Srcaddr.
* `ssl_cert` - Ssl-Cert.
* `status` - Status. Valid values: `disable`, `enable`.

* `status_ipv6` - Status-Ipv6. Valid values: `disable`, `enable`.

* `strict_scheme_check` - Strict-Scheme-Check. Valid values: `disable`, `enable`.

* `timeout` - Timeout.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `icap_service` block supports:

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
* `id` - an identifier for the resource with format {{icap_server_id}}.

## Import

Icap LocalServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_icap_localserver.labelname {{icap_server_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

