---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_waf_profile_constraint"
description: |-
  <i>This object will be purged after policy copy and install.</i> WAF HTTP protocol restrictions.
---

# fmgdevice_waf_profile_constraint
<i>This object will be purged after policy copy and install.</i> WAF HTTP protocol restrictions.

~> This resource is a sub resource for variable `constraint` of resource `fmgdevice_waf_profile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `content_length`: `fmgdevice_waf_profile_constraint_contentlength`
>- `exception`: `fmgdevice_waf_profile_constraint_exception`
>- `header_length`: `fmgdevice_waf_profile_constraint_headerlength`
>- `hostname`: `fmgdevice_waf_profile_constraint_hostname`
>- `line_length`: `fmgdevice_waf_profile_constraint_linelength`
>- `malformed`: `fmgdevice_waf_profile_constraint_malformed`
>- `max_cookie`: `fmgdevice_waf_profile_constraint_maxcookie`
>- `max_header_line`: `fmgdevice_waf_profile_constraint_maxheaderline`
>- `max_range_segment`: `fmgdevice_waf_profile_constraint_maxrangesegment`
>- `max_url_param`: `fmgdevice_waf_profile_constraint_maxurlparam`
>- `method`: `fmgdevice_waf_profile_constraint_method`
>- `param_length`: `fmgdevice_waf_profile_constraint_paramlength`
>- `url_param_length`: `fmgdevice_waf_profile_constraint_urlparamlength`
>- `version`: `fmgdevice_waf_profile_constraint_version`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `content_length` - Content-Length. The structure of `content_length` block is documented below.
* `exception` - Exception. The structure of `exception` block is documented below.
* `header_length` - Header-Length. The structure of `header_length` block is documented below.
* `hostname` - Hostname. The structure of `hostname` block is documented below.
* `line_length` - Line-Length. The structure of `line_length` block is documented below.
* `malformed` - Malformed. The structure of `malformed` block is documented below.
* `max_cookie` - Max-Cookie. The structure of `max_cookie` block is documented below.
* `max_header_line` - Max-Header-Line. The structure of `max_header_line` block is documented below.
* `max_range_segment` - Max-Range-Segment. The structure of `max_range_segment` block is documented below.
* `max_url_param` - Max-Url-Param. The structure of `max_url_param` block is documented below.
* `method` - Method. The structure of `method` block is documented below.
* `param_length` - Param-Length. The structure of `param_length` block is documented below.
* `url_param_length` - Url-Param-Length. The structure of `url_param_length` block is documented below.
* `version` - Version. The structure of `version` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `content_length` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `length` - Length of HTTP content in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `exception` block supports:

* `address` - Host address.
* `content_length` - HTTP content length in request. Valid values: `disable`, `enable`.

* `header_length` - HTTP header length in request. Valid values: `disable`, `enable`.

* `hostname` - Enable/disable hostname check. Valid values: `disable`, `enable`.

* `id` - Exception ID.
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


The `header_length` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `length` - Length of HTTP header in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `hostname` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `line_length` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `length` - Length of HTTP line in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `malformed` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `max_cookie` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `max_cookie` - Maximum number of cookies in HTTP request (0 to 2147483647).
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `max_header_line` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `max_header_line` - Maximum number HTTP header lines (0 to 2147483647).
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `max_range_segment` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `max_range_segment` - Maximum number of range segments in HTTP range line (0 to 2147483647).
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `max_url_param` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `max_url_param` - Maximum number of parameters in URL (0 to 2147483647).
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `method` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `param_length` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `length` - Maximum length of parameter in URL, HTTP POST request or HTTP body in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `url_param_length` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `length` - Maximum length of URL parameter in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.


The `version` block supports:

* `action` - Action. Valid values: `allow`, `block`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Enable/disable the constraint. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Waf ProfileConstraint can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_waf_profile_constraint.labelname WafProfileConstraint
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

