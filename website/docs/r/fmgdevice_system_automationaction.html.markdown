---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_automationaction"
description: |-
  Action for automation stitches.
---

# fmgdevice_system_automationaction
Action for automation stitches.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `http_headers`: `fmgdevice_system_automationaction_httpheaders`



## Example Usage

```hcl
resource "fmgdevice_system_automationaction" "trname" {
  accprofile                 = ["your own value"]
  action_type                = "microsoft-teams-notification"
  alicloud_access_key_id     = "your own value"
  alicloud_access_key_secret = ["your own value"]
  alicloud_account_id        = "your own value"
  name                       = "your own value"
  device_name                = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `accprofile` - Access profile for CLI script action to access FortiGate features.
* `action_type` - Action type. Valid values: `email`, `ios-notification`, `alert`, `disable-ssid`, `quarantine`, `ban-ip`, `quarantine-forticlient`, `aws-lambda`, `webhook`, `quarantine-nsx`, `azure-function`, `cli-script`, `google-cloud-function`, `alicloud-function`, `slack-notification`, `quarantine-fortinac`, `microsoft-teams-notification`, `fortiexplorer-notification`, `system-actions`.

* `alicloud_access_key_id` - AliCloud AccessKey ID.
* `alicloud_access_key_secret` - AliCloud AccessKey secret.
* `alicloud_account_id` - AliCloud account ID.
* `alicloud_function` - AliCloud function name.
* `alicloud_function_authorization` - AliCloud function authorization type. Valid values: `anonymous`, `function`.

* `alicloud_function_domain` - AliCloud function domain.
* `alicloud_region` - AliCloud region.
* `alicloud_service` - AliCloud service name.
* `alicloud_version` - AliCloud version.
* `aws_api_id` - AWS API Gateway ID.
* `aws_api_key` - AWS API Gateway API key.
* `aws_api_path` - AWS API Gateway path.
* `aws_api_stage` - AWS API Gateway deployment stage name.
* `aws_domain` - AWS domain.
* `aws_region` - AWS region.
* `azure_api_key` - Azure function API key.
* `azure_app` - Azure function application name.
* `azure_domain` - Azure function domain.
* `azure_function` - Azure function name.
* `azure_function_authorization` - Azure function authorization level. Valid values: `anonymous`, `function`, `admin`.

* `delay` - Delay.
* `email_body` - Email body.
* `description` - Description.
* `email_from` - Email sender name.
* `email_subject` - Email subject.
* `email_to` - Email addresses.
* `gcp_function` - Google Cloud function name.
* `gcp_function_domain` - Google Cloud function domain.
* `gcp_function_region` - Google Cloud function region.
* `gcp_project` - Google Cloud Platform project name.
* `execute_security_fabric` - Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric. Valid values: `disable`, `enable`.

* `headers` - Headers.
* `forticare_email` - Enable/disable use of your FortiCare email address as the email-to address. Valid values: `disable`, `enable`.

* `http_body` - Request body (if necessary). Should be serialized json string.
* `http_headers` - Http-Headers. The structure of `http_headers` block is documented below.
* `message` - Message content.
* `message_type` - Message type. Valid values: `text`, `json`.

* `method` - Request method (POST, PUT, GET, PATCH or DELETE). Valid values: `delete`, `get`, `post`, `put`, `patch`.

* `minimum_interval` - Limit execution to no more than once in this interval (in seconds).
* `name` - Name.
* `output_size` - Number of megabytes to limit script output to (1 - 1024, default = 10).
* `port` - Protocol port.
* `protocol` - Request protocol. Valid values: `http`, `https`.

* `replacement_message` - Enable/disable replacement message. Valid values: `disable`, `enable`.

* `replacemsg_group` - Replacement message group.
* `required` - Required. Valid values: `disable`, `enable`.

* `script` - CLI script.
* `sdn_connector` - NSX SDN connector names.
* `security_tag` - NSX security tag.
* `system_action` - System action type. Valid values: `reboot`, `shutdown`, `backup-config`.

* `timeout` - Maximum running time for this script in seconds (0 = no timeout).
* `tls_certificate` - Custom TLS certificate for API request.
* `uri` - Request API URI.
* `verify_host_cert` - Enable/disable verification of the remote host certificate. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `http_headers` block supports:

* `id` - Entry ID.
* `key` - Request header key.
* `value` - Request header value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationAction can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_automationaction.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

