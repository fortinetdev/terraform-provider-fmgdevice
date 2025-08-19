---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_automationtrigger"
description: |-
  Trigger for automation stitches.
---

# fmgdevice_system_automationtrigger
Trigger for automation stitches.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `fields`: `fmgdevice_system_automationtrigger_fields`



## Example Usage

```hcl
resource "fmgdevice_system_automationtrigger" "trname" {
  description           = "your own value"
  event_type            = "local-cert-near-expiry"
  fabric_event_name     = "your own value"
  fabric_event_severity = "your own value"
  faz_event_name        = "your own value"
  name                  = "your own value"
  device_name           = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `description` - Description.
* `event_type` - Event type. Valid values: `ioc`, `event-log`, `reboot`, `low-memory`, `high-cpu`, `license-near-expiry`, `ha-failover`, `config-change`, `security-rating-summary`, `virus-ips-db-updated`, `faz-event`, `incoming-webhook`, `fabric-event`, `ips-logs`, `anomaly-logs`, `virus-logs`, `ssh-logs`, `webfilter-violation`, `traffic-violation`, `local-cert-near-expiry`.

* `fabric_event_name` - Fabric connector event handler name.
* `fabric_event_severity` - Fabric connector event severity.
* `faz_event_name` - FortiAnalyzer event handler name.
* `faz_event_severity` - FortiAnalyzer event severity.
* `faz_event_tags` - FortiAnalyzer event tags.
* `fields` - Fields. The structure of `fields` block is documented below.
* `ioc_level` - IOC threat level. Valid values: `high`, `medium`.

* `license_type` - License type. Valid values: `forticare-support`, `fortiguard-webfilter`, `fortiguard-antispam`, `fortiguard-antivirus`, `fortiguard-ips`, `fortiguard-management`, `forticloud`, `any`.

* `logid` - Log IDs to trigger event.
* `name` - Name.
* `report_type` - Security Rating report. Valid values: `posture`, `coverage`, `optimization`, `any`.

* `serial` - Fabric connector serial number.
* `stitch_name` - Triggering stitch name.
* `trigger_datetime` - Trigger date and time (YYYY-MM-DD HH:MM:SS).
* `trigger_day` - Day within a month to trigger.
* `trigger_frequency` - Scheduled trigger frequency (default = daily). Valid values: `daily`, `weekly`, `once`, `monthly`, `hourly`.

* `trigger_hour` - Hour of the day on which to trigger (0 - 23, default = 1).
* `trigger_minute` - Minute of the hour on which to trigger (0 - 59, default = 0).
* `trigger_type` - Trigger type. Valid values: `event-based`, `scheduled`.

* `trigger_weekday` - Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.

* `vdom` - Virtual domain(s) that this trigger is valid for.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `fields` block supports:

* `id` - Entry ID.
* `name` - Name.
* `value` - Value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationTrigger can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_automationtrigger.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

