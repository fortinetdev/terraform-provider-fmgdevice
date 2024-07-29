---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ips_rule"
description: |-
  Configure IPS rules.
---

# fmgdevice_ips_rule
Configure IPS rules.

## Example Usage

```hcl
resource "fmgdevice_ips_rule" "trname" {
  action      = "pass"
  application = "your own value"
  date        = 10
  group       = "your own value"
  location    = "your own value"
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `action` - Action. Valid values: `pass`, `block`.

* `application` - Application.
* `date` - Date.
* `group` - Group.
* `location` - Location.
* `log` - Log. Valid values: `disable`, `enable`.

* `log_packet` - Log-Packet. Valid values: `disable`, `enable`.

* `name` - Rule name.
* `os` - Os.
* `rev` - Rev.
* `rule_id` - Rule-Id.
* `service` - Service.
* `severity` - Severity.
* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ips Rule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ips_rule.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

