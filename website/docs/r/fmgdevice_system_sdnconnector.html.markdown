---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sdnconnector"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure connection to SDN Connector.
---

# fmgdevice_system_sdnconnector
<i>This object will be purged after policy copy and install.</i> Configure connection to SDN Connector.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `compartment_list`: `fmgdevice_system_sdnconnector_compartmentlist`
>- `external_account_list`: `fmgdevice_system_sdnconnector_externalaccountlist`
>- `external_ip`: `fmgdevice_system_sdnconnector_externalip`
>- `forwarding_rule`: `fmgdevice_system_sdnconnector_forwardingrule`
>- `gcp_project_list`: `fmgdevice_system_sdnconnector_gcpprojectlist`
>- `nic`: `fmgdevice_system_sdnconnector_nic`
>- `oci_region_list`: `fmgdevice_system_sdnconnector_ociregionlist`
>- `route`: `fmgdevice_system_sdnconnector_route`
>- `route_table`: `fmgdevice_system_sdnconnector_routetable`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `_local_cert` - _Local_Cert.
* `access_key` - AWS / ACS access key ID.
* `alt_resource_ip` - Enable/disable AWS alternative resource IP. Valid values: `disable`, `enable`.

* `api_key` - IBM cloud API key or service ID API key.
* `azure_region` - Azure server region. Valid values: `global`, `china`, `germany`, `usgov`, `local`.

* `client_id` - Azure client ID (application ID).
* `client_secret` - Azure client secret (application key).
* `compartment_list` - Compartment-List. The structure of `compartment_list` block is documented below.
* `compute_generation` - Compute-Generation.
* `domain` - Domain name.
* `external_account_list` - External-Account-List. The structure of `external_account_list` block is documented below.
* `external_ip` - External-Ip. The structure of `external_ip` block is documented below.
* `forwarding_rule` - Forwarding-Rule. The structure of `forwarding_rule` block is documented below.
* `gcp_project` - GCP project name.
* `gcp_project_list` - Gcp-Project-List. The structure of `gcp_project_list` block is documented below.
* `group_name` - Full path group name of computers.
* `ha_status` - Enable/disable use for FortiGate HA service. Valid values: `disable`, `enable`.

* `ibm_region` - IBM cloud region name. Valid values: `dallas`, `washington-dc`, `london`, `frankfurt`, `sydney`, `tokyo`, `osaka`, `toronto`, `sao-paulo`, `dallas-private`, `washington-dc-private`, `london-private`, `frankfurt-private`, `sydney-private`, `tokyo-private`, `osaka-private`, `toronto-private`, `sao-paulo-private`, `madrid`, `madrid-private`.

* `ibm_region_gen1` - Ibm-Region-Gen1. Valid values: `us-south`, `us-east`, `germany`, `great-britain`, `japan`, `australia`.

* `ibm_region_gen2` - Ibm-Region-Gen2. Valid values: `us-south`, `us-east`, `great-britain`.

* `key_passwd` - Private key password.
* `login_endpoint` - Azure Stack login endpoint.
* `message_server_port` - HTTP port number of the SAP message server.
* `microsoft_365` - Enable to use as Microsoft 365 connector. Valid values: `disable`, `enable`.

* `name` - SDN connector name.
* `nic` - Nic. The structure of `nic` block is documented below.
* `nsx_cert_fingerprint` - NSX certificate fingerprint.
* `oci_cert` - OCI certificate.
* `oci_fingerprint` - Oci-Fingerprint.
* `oci_region_list` - Oci-Region-List. The structure of `oci_region_list` block is documented below.
* `oci_region_type` - OCI region type. Valid values: `commercial`, `government`.

* `password` - Password of the remote SDN connector as login credentials.
* `private_key` - Private key of GCP service account.
* `proxy` - SDN proxy.
* `region` - AWS / ACS region name.
* `resource_group` - Azure resource group.
* `resource_url` - Azure Stack resource URL.
* `rest_interface` - Interface name for REST service to listen on. Valid values: `mgmt`, `sync`.

* `rest_password` - Password for REST service.
* `rest_sport` - REST service access port (1 - 65535).
* `rest_ssl` - Rest-Ssl. Valid values: `disable`, `enable`.

* `route` - Route. The structure of `route` block is documented below.
* `route_table` - Route-Table. The structure of `route_table` block is documented below.
* `secret_key` - AWS / ACS secret access key.
* `secret_token` - Secret token of Kubernetes service account.
* `server` - Server address of the remote SDN connector.
* `server_ca_cert` - Trust only those servers whose certificate is directly/indirectly signed by this certificate.
* `server_cert` - Trust servers that contain this certificate only.
* `server_list` - Server address list of the remote SDN connector.
* `server_port` - Port number of the remote SDN connector.
* `service_account` - GCP service account email.
* `status` - Enable/disable connection to the remote SDN connector. Valid values: `disable`, `enable`.

* `subscription_id` - Azure subscription ID.
* `tenant_id` - Tenant ID (directory ID).
* `type` - Type of SDN connector. Valid values: `aci`, `aws`, `nsx`, `nuage`, `azure`, `gcp`, `oci`, `openstack`, `kubernetes`, `vmware`, `alicloud`, `sepm`, `aci-direct`, `ibm`, `nutanix`, `sap`.

* `update_interval` - Dynamic object update interval (30 - 3600 sec, default = 60, 0 = disabled).
* `use_metadata_iam` - Enable/disable use of IAM role from metadata to call API. Valid values: `disable`, `enable`.

* `user_id` - User ID.
* `username` - Username of the remote SDN connector as login credentials.
* `vcenter_password` - vCenter server password for NSX quarantine.
* `vcenter_server` - vCenter server address for NSX quarantine.
* `vcenter_username` - vCenter server username for NSX quarantine.
* `vdom` - Virtual domain name of the remote SDN connector.
* `verify_certificate` - Enable/disable server certificate verification. Valid values: `disable`, `enable`.

* `vmx_image_url` - URL of web-hosted VMX image.
* `vmx_service_name` - VMX Service name.
* `vpc_id` - AWS VPC ID.
* `compartment_id` - Compartment-Id.
* `oci_region` - Oci-Region.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `compartment_list` block supports:

* `compartment_id` - OCI compartment ID.

The `external_account_list` block supports:

* `external_id` - AWS external ID.
* `region_list` - AWS region name list.
* `role_arn` - AWS role ARN to assume.

The `external_ip` block supports:

* `name` - External IP name.

The `forwarding_rule` block supports:

* `rule_name` - Forwarding rule name.
* `target` - Target instance name.

The `gcp_project_list` block supports:

* `gcp_zone_list` - Configure GCP zone list.
* `id` - GCP project ID.

The `nic` block supports:

* `ip` - Ip. The structure of `ip` block is documented below.
* `name` - Network interface name.
* `peer_nic` - Peer network interface name.

The `ip` block supports:

* `name` - IP configuration name.
* `private_ip` - Private IP address.
* `public_ip` - Public IP name.
* `resource_group` - Resource group of Azure public IP.

The `oci_region_list` block supports:

* `region` - OCI region.

The `route` block supports:

* `name` - Route name.

The `route_table` block supports:

* `name` - Route table name.
* `resource_group` - Resource group of Azure route table.
* `route` - Route. The structure of `route` block is documented below.
* `subscription_id` - Subscription ID of Azure route table.

The `route` block supports:

* `name` - Route name.
* `next_hop` - Next hop address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SdnConnector can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sdnconnector.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

