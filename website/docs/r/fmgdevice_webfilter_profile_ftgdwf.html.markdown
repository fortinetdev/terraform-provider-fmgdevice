---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_profile_ftgdwf"
description: |-
  <i>This object will be purged after policy copy and install.</i> FortiGuard Web Filter settings.
---

# fmgdevice_webfilter_profile_ftgdwf
<i>This object will be purged after policy copy and install.</i> FortiGuard Web Filter settings.

~> This resource is a sub resource for variable `ftgd_wf` of resource `fmgdevice_webfilter_profile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `filters`: `fmgdevice_webfilter_profile_ftgdwf_filters`
>- `quota`: `fmgdevice_webfilter_profile_ftgdwf_quota`
>- `risk`: `fmgdevice_webfilter_profile_ftgdwf_risk`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `exempt_quota` - Do not stop quota for these categories.
* `filters` - Filters. The structure of `filters` block is documented below.
* `max_quota_timeout` - Maximum FortiGuard quota used by single page view in seconds (excludes streams).
* `options` - Options for FortiGuard Web Filter. Valid values: `error-allow`, `rate-server-ip`, `connect-request-bypass`, `ftgd-disable`.

* `ovrd` - Allow web filter profile overrides.
* `quota` - Quota. The structure of `quota` block is documented below.
* `rate_crl_urls` - Enable/disable rating CRL by URL. Valid values: `disable`, `enable`.

* `rate_css_urls` - Enable/disable rating CSS by URL. Valid values: `disable`, `enable`.

* `rate_image_urls` - Rate-Image-Urls. Valid values: `disable`, `enable`.

* `rate_javascript_urls` - Enable/disable rating JavaScript by URL. Valid values: `disable`, `enable`.

* `risk` - Risk. The structure of `risk` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `filters` block supports:

* `action` - Action to take for matches. Valid values: `block`, `monitor`, `warning`, `authenticate`.

* `auth_usr_grp` - Groups with permission to authenticate.
* `category` - Categories and groups the filter examines.
* `id` - ID number.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `override_replacemsg` - Override replacement message.
* `warn_duration` - Duration of warnings.
* `warning_duration_type` - Re-display warning after closing browser or after a timeout. Valid values: `session`, `timeout`.

* `warning_prompt` - Warning prompts in each category or each domain. Valid values: `per-domain`, `per-category`.


The `quota` block supports:

* `category` - FortiGuard categories to apply quota to (category action must be set to monitor).
* `duration` - Duration of quota.
* `id` - ID number.
* `override_replacemsg` - Override replacement message.
* `type` - Quota type. Valid values: `time`, `traffic`.

* `unit` - Traffic quota unit of measurement. Valid values: `B`, `KB`, `MB`, `GB`.

* `value` - Traffic quota value.

The `risk` block supports:

* `action` - Action to take for matches. Valid values: `block`, `monitor`.

* `id` - ID number.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `risk_level` - Risk level to be examined.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Webfilter ProfileFtgdWf can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_profile_ftgdwf.labelname WebfilterProfileFtgdWf
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

