---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_threatweight"
description: |-
  Configure threat weight settings.
---

# fmgdevice_log_threatweight
Configure threat weight settings.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `application`: `fmgdevice_log_threatweight_application`
>- `geolocation`: `fmgdevice_log_threatweight_geolocation`
>- `ips`: `fmgdevice_log_threatweight_ips`
>- `level`: `fmgdevice_log_threatweight_level`
>- `malware`: `fmgdevice_log_threatweight_malware`
>- `web`: `fmgdevice_log_threatweight_web`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `application` - Application. The structure of `application` block is documented below.
* `blocked_connection` - Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `botnet_connection_detected` - Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `failed_connection` - Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `geolocation` - Geolocation. The structure of `geolocation` block is documented below.
* `ips` - Ips. The structure of `ips` block is documented below.
* `level` - Level. The structure of `level` block is documented below.
* `malware` - Malware. The structure of `malware` block is documented below.
* `status` - Enable/disable the threat weight feature. Valid values: `disable`, `enable`.

* `url_block_detected` - Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `web` - Web. The structure of `web` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `application` block supports:

* `category` - Application category.
* `id` - Entry ID.
* `level` - Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.


The `geolocation` block supports:

* `country` - Country code.
* `id` - Entry ID.
* `level` - Threat weight score for Geolocation-based events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.


The `ips` block supports:

* `critical_severity` - Threat weight score for IPS critical severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `high_severity` - Threat weight score for IPS high severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `info_severity` - Threat weight score for IPS info severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `low_severity` - Threat weight score for IPS low severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `medium_severity` - Threat weight score for IPS medium severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.


The `level` block supports:

* `critical` - Critical level score value (1 - 100).
* `high` - High level score value (1 - 100).
* `low` - Low level score value (1 - 100).
* `medium` - Medium level score value (1 - 100).

The `malware` block supports:

* `command_blocked` - Threat weight score for blocked command detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `content_disarm` - Threat weight score for virus (content disarm) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `ems_threat_feed` - Threat weight score for virus (EMS threat feed) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `file_blocked` - Threat weight score for blocked file detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `fortiai` - Threat weight score for FortiAI-detected virus. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `fsa_high_risk` - Threat weight score for FortiSandbox high risk malware detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `fsa_malicious` - Threat weight score for FortiSandbox malicious malware detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `fsa_medium_risk` - Threat weight score for FortiSandbox medium risk malware detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `inline_block` - Threat weight score for malware detected by inline block. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `malware_list` - Threat weight score for virus (malware list) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `mimefragmented` - Threat weight score for mimefragmented detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `oversized` - Threat weight score for oversized file detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `switch_proto` - Threat weight score for switch proto detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `virus_file_type_executable` - Threat weight score for virus (file type executable) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `virus_infected` - Threat weight score for virus (infected) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `virus_outbreak_prevention` - Threat weight score for virus (outbreak prevention) event. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `virus_scan_error` - Threat weight score for virus (scan error) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `fortindr` - Fortindr. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

* `fortisandbox` - Fortisandbox. Valid values: `disable`, `low`, `medium`, `high`, `critical`.


The `web` block supports:

* `category` - Threat weight score for web category filtering matches.
* `id` - Entry ID.
* `level` - Threat weight score for web category filtering matches. Valid values: `disable`, `low`, `medium`, `high`, `critical`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log ThreatWeight can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_threatweight.labelname LogThreatWeight
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

