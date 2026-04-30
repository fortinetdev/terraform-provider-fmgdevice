---
subcategory: "ZTNA"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ztna_webportalbookmark_llmsecureproxy"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_ztna_webportalbookmark_llmsecureproxy
<i>This object will be purged after policy copy and install.</i>

~> This resource is a sub resource for variable `llm_secure_proxy` of resource `fmgdevice_ztna_webportalbookmark`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `web_portal_bookmark` - Web Portal Bookmark.

* `all_llm_servers` - All-Llm-Servers. Valid values: `disable`, `enable`.

* `llm_servers` - Llm-Servers.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ztna WebPortalBookmarkLlmSecureProxy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "web_portal_bookmark=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ztna_webportalbookmark_llmsecureproxy.labelname ZtnaWebPortalBookmarkLlmSecureProxy
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

