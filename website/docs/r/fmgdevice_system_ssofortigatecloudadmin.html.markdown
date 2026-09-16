---
subcategory: "System"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ssofortigatecloudadmin"
description: |-
  Configure FortiCloud SSO admin users.
---

# fmgdevice_system_ssofortigatecloudadmin
Configure FortiCloud SSO admin users.

## Example Usage

```hcl
resource "fmgdevice_system_ssofortigatecloudadmin" "trname" {
  accprofile                           = ["your own value"]
  gui_default_dashboard_template       = "your own value"
  gui_ignore_invalid_signature_version = "your own value"
  gui_ignore_release_overview_version  = "your own value"
  name                                 = "your own value"
  device_name                          = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `accprofile` - FortiCloud SSO admin user access profile.
* `gui_custom_theme` - Custom theme that overrides the default FortiGate theme.
* `gui_dashboard_id` - GUI Dashboard ID.
* `gui_default_dashboard_template` - The default dashboard template.
* `gui_ignore_invalid_signature_version` - FortiOS image build version to ignore invalid signature warning for.
* `gui_ignore_release_overview_version` - FortiOS version to ignore release overview prompt for.
* `gui_llm_provider` - Select the LLM provider. Valid values: `fortiai`, `openai`.

* `gui_theme` - Predefined theme that overrides the default FortiGate theme. Valid values: `melongene`, `mariner`, `neutrino`, `jade`, `graphite`, `dark-matter`, `onyx`, `eclipse`, `retro`, `jet-stream`, `security-fabric`, `none`.

* `gui_theme_type` - Use predefined themes or custom themes. Valid values: `predefined`, `custom`.

* `name` - FortiCloud SSO admin name.
* `openai_api_key` - OpenAI API key.
* `openai_api_key_part2` - OpenAI API key part 2 for excess length.
* `openai_model` - OpenAI model.
* `openai_org_id` - OpenAI organization ID.
* `openai_project_id` - OpenAI project ID.
* `vdom` - Virtual domain(s) that the administrator can access.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SsoFortigateCloudAdmin can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ssofortigatecloudadmin.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

