---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_csf"
description: |-
  Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.
---

# fmgdevice_system_csf
Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `fabric_connector`: `fmgdevice_system_csf_fabricconnector`
>- `fabric_device`: `fmgdevice_system_csf_fabricdevice`
>- `trusted_list`: `fmgdevice_system_csf_trustedlist`



## Example Usage

```hcl
resource "fmgdevice_system_csf" "trname" {
  accept_auth_by_cert        = "enable"
  authorization_request_type = "certificate"
  certificate                = ["your own value"]
  configuration_sync         = "default"
  downstream_access          = "disable"
  device_name                = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `accept_auth_by_cert` - Accept connections with unknown certificates and ask admin for approval. Valid values: `disable`, `enable`.

* `authorization_request_type` - Authorization request type. Valid values: `certificate`, `serial`.

* `certificate` - Certificate.
* `configuration_sync` - Configuration sync mode. Valid values: `default`, `local`.

* `downstream_access` - Enable/disable downstream device access to this device's configuration and data. Valid values: `disable`, `enable`.

* `downstream_accprofile` - Default access profile for requests from downstream devices.
* `fabric_connector` - Fabric-Connector. The structure of `fabric_connector` block is documented below.
* `fabric_device` - Fabric-Device. The structure of `fabric_device` block is documented below.
* `fabric_object_unification` - Fabric CMDB Object Unification. Valid values: `default`, `local`.

* `fabric_workers` - Number of worker processes for Security Fabric daemon.
* `file_mgmt` - Enable/disable Security Fabric daemon file management. Valid values: `disable`, `enable`.

* `file_quota` - Maximum amount of memory that can be used by the daemon files (in bytes).
* `file_quota_warning` - Warn when the set percentage of quota has been used.
* `fixed_key` - Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
* `forticloud_account_enforcement` - Fabric FortiCloud account unification. Valid values: `disable`, `enable`.

* `group_name` - Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
* `group_password` - Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
* `legacy_authentication` - Enable/disable legacy authentication. Valid values: `disable`, `enable`.

* `log_unification` - Enable/disable broadcast of discovery messages for log unification. Valid values: `disable`, `enable`.

* `management_ip` - Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
* `management_port` - Overriding port for management connection (Overrides admin port).
* `saml_configuration_sync` - SAML setting configuration synchronization. Valid values: `default`, `local`.

* `source_ip` - Source IP address for communication with the upstream FortiGate.
* `status` - Enable/disable Security Fabric. Valid values: `disable`, `enable`.

* `trusted_list` - Trusted-List. The structure of `trusted_list` block is documented below.
* `uid` - Unique ID of the current CSF node
* `upstream` - IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_ip` - IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_interface` - Specify outgoing interface to reach server.
* `upstream_interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `upstream_port` - The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `fabric_connector` block supports:

* `accprofile` - Override access profile.
* `configuration_write_access` - Enable/disable downstream device write access to configuration. Valid values: `disable`, `enable`.

* `serial` - Serial.
* `vdom` - Virtual domains that the connector has access to. If none are set, the connector will only have access to the VDOM that it joins the Security Fabric through.

The `fabric_device` block supports:

* `access_token` - Device access token.
* `device_ip` - Device IP.
* `https_port` - HTTPS port for fabric device.
* `name` - Device name.

The `trusted_list` block supports:

* `action` - Security fabric authorization action. Valid values: `deny`, `accept`.

* `authorization_type` - Authorization type. Valid values: `certificate`, `serial`.

* `certificate` - Certificate.
* `downstream_authorization` - Trust authorizations by this node's administrator. Valid values: `disable`, `enable`.

* `ha_members` - HA members.
* `index` - Index of the downstream in tree.
* `name` - Name.
* `serial` - Serial.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Csf can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_csf.labelname SystemCsf
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

