---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_endpointcontrol_fctems"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure FortiClient Enterprise Management Server (EMS) entries.
---

# fmgdevice_endpointcontrol_fctems
<i>This object will be purged after policy copy and install.</i> Configure FortiClient Enterprise Management Server (EMS) entries.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `call_timeout` - FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
* `certificate` - FortiClient EMS certificate.
* `capabilities` - List of EMS capabilities. Valid values: `fabric-auth`, `silent-approval`, `websocket`, `websocket-malware`, `push-ca-certs`, `common-tags-api`, `tenant-id`, `single-vdom-connector`, `client-avatars`, `fgt-sysinfo-api`, `ztna-server-info`, `used-tags`.

* `certificate_fingerprint` - EMS certificate fingerprint.
* `cloud_authentication_access_key` - FortiClient EMS Cloud multitenancy access key
* `dirty_reason` - Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.

* `ems_id` - EMS ID in order (1 - 7).
* `fortinetone_cloud_authentication` - Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `disable`, `enable`.

* `https_port` - FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `name` - FortiClient Enterprise Management Server (EMS) name.
* `out_of_sync_threshold` - Outdated resource threshold in seconds (10 - 3600, default = 180).
* `pull_avatars` - Pull-Avatars. Valid values: `disable`, `enable`.

* `pull_malware_hash` - Enable/disable pulling FortiClient malware hash from EMS. Valid values: `disable`, `enable`.

* `pull_sysinfo` - Enable/disable pulling SysInfo from EMS. Valid values: `disable`, `enable`.

* `pull_tags` - Enable/disable pulling FortiClient user tags from EMS. Valid values: `disable`, `enable`.

* `pull_vulnerabilities` - Enable/disable pulling vulnerabilities from EMS. Valid values: `disable`, `enable`.

* `send_tags_to_all_vdoms` - Relax restrictions on tags to send all EMS tags to all VDOMs Valid values: `disable`, `enable`.

* `serial_number` - EMS Serial Number.
* `server` - FortiClient EMS FQDN or IPv4 address.
* `source_ip` - REST API call source IP.
* `status_check_interval` - FortiClient EMS call timeout in seconds (1 - 120 seconds, default = 5).
* `status` - Enable or disable this EMS configuration. Valid values: `disable`, `enable`.

* `tenant_id` - EMS Tenant ID.
* `trust_ca_cn` - Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal. Valid values: `disable`, `enable`.

* `verified_cn` - EMS certificate CN.
* `verifying_ca` - Lowest CA cert on Fortigate in verified EMS cert chain.
* `websocket_override` - Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.

* `ca_cn_info` - Ca-Cn-Info.
* `cloud_server_type` - Cloud-Server-Type. Valid values: `production`, `alpha`, `beta`.

* `preserve_ssl_session` - Preserve-Ssl-Session. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ems_id}}.

## Import

EndpointControl Fctems can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_endpointcontrol_fctems.labelname {{ems_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

