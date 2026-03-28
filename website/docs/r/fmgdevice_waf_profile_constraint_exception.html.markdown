---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_waf_profile_constraint_exception"
description: |-
  <i>This object will be purged after policy copy and install.</i> HTTP constraint exception.
---

# fmgdevice_waf_profile_constraint_exception
<i>This object will be purged after policy copy and install.</i> HTTP constraint exception.

~> This resource is a sub resource for variable `exception` of resource `fmgdevice_waf_profile_constraint`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `address` - Host address.
* `content_length` - HTTP content length in request. Valid values: `disable`, `enable`.

* `header_length` - HTTP header length in request. Valid values: `disable`, `enable`.

* `hostname` - Enable/disable hostname check. Valid values: `disable`, `enable`.

* `fosid` - Exception ID.
* `line_length` - HTTP line length in request. Valid values: `disable`, `enable`.

* `malformed` - Enable/disable malformed HTTP request check. Valid values: `disable`, `enable`.

* `max_cookie` - Maximum number of cookies in HTTP request. Valid values: `disable`, `enable`.

* `max_header_line` - Maximum number of HTTP header line. Valid values: `disable`, `enable`.

* `max_range_segment` - Maximum number of range segments in HTTP range line. Valid values: `disable`, `enable`.

* `max_url_param` - Maximum number of parameters in URL. Valid values: `disable`, `enable`.

* `method` - Enable/disable HTTP method check. Valid values: `disable`, `enable`.

* `param_length` - Maximum length of parameter in URL, HTTP POST request or HTTP body. Valid values: `disable`, `enable`.

* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `disable`, `enable`.

* `url_param_length` - Maximum length of parameter in URL. Valid values: `disable`, `enable`.

* `version` - Enable/disable HTTP version check. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Waf ProfileConstraintException can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_waf_profile_constraint_exception.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

