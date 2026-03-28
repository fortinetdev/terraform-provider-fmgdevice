---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_ldap"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure LDAP server entries.
---

# fmgdevice_user_ldap
<i>This object will be purged after policy copy and install.</i> Configure LDAP server entries.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fmgdevice_user_ldap_dynamic_mapping`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `account_key_cert_field` - Define subject identity field in certificate for user access right checking. Valid values: `othername`, `rfc822name`, `dnsname`, `cn`.

* `account_key_filter` - Account key filter, using the UPN as the search filter.
* `account_key_processing` - Account key processing operation. The FortiGate will keep either the whole domain or strip the domain from the subject identity. Valid values: `same`, `strip`.

* `antiphish` - Enable/disable AntiPhishing credential backend. Valid values: `disable`, `enable`.

* `ca_cert` - CA certificate name.
* `client_cert` - Client certificate name.
* `client_cert_auth` - Enable/disable using client certificate for TLS authentication. Valid values: `disable`, `enable`.

* `cnid` - Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
* `dn` - Distinguished name used to look up entries on the LDAP server.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `group_filter` - Filter used for group matching.
* `group_member_check` - Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.

* `group_object_filter` - Filter used for group searching.
* `group_search_base` - Search base used for group searching.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `member_attr` - Name of attribute from which to get group membership.
* `name` - LDAP server entry name.
* `obtain_user_info` - Enable/disable obtaining of user information. Valid values: `disable`, `enable`.

* `password` - Password for initial binding.
* `password_attr` - Name of attribute to get password hash.
* `password_expiry_warning` - Enable/disable password expiry warnings. Valid values: `disable`, `enable`.

* `password_renewal` - Enable/disable online password renewal. Valid values: `disable`, `enable`.

* `port` - Port to be used for communication with the LDAP server (default = 389).
* `search_type` - Search type. Valid values: `recursive`.

* `secondary_server` - Secondary LDAP server CN domain name or IP.
* `secure` - Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.

* `server` - LDAP server CN domain name or IP.
* `server_identity_check` - Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `disable`, `enable`.

* `source_ip` - FortiGate IP address to be used for communication with the LDAP server.
* `source_ip_interface` - Source interface for communication with the LDAP server.
* `source_port` - Source port to be used for communication with the LDAP server.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1-3`.

* `status_ttl` - Time for which server reachability is cached so that when a server is unreachable, it will not be retried for at least this period of time (0 = cache disabled, default = 300).
* `tertiary_server` - Tertiary LDAP server CN domain name or IP.
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.

* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.

* `two_factor_filter` - Filter used to synchronize users to FortiToken Cloud.
* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.

* `type` - Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.

* `user_info_exchange_server` - MS Exchange server from which to fetch user information.
* `username` - Username (full DN) for initial binding.
* `account_key_upn_san` - Account-Key-Upn-San. Valid values: `othername`, `rfc822name`, `dnsname`.

* `vrf_select` - VRF ID used for connection to server.
* `max_connections` - Max-Connections.
* `ssl_max_proto_version` - Ssl-Max-Proto-Version. Valid values: `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1`, `TLSv1-3`.

* `validate_server_certificate` - Validate-Server-Certificate. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `account_key_cert_field` - Define subject identity field in certificate for user access right checking. Valid values: `othername`, `rfc822name`, `dnsname`, `cn`.

* `account_key_filter` - Account key filter, using the UPN as the search filter.
* `account_key_name` - Account-Key-Name.
* `account_key_processing` - Account key processing operation. The FortiGate will keep either the whole domain or strip the domain from the subject identity. Valid values: `same`, `strip`.

* `account_key_upn_san` - Account-Key-Upn-San. Valid values: `othername`, `rfc822name`, `dnsname`.

* `antiphish` - Enable/disable AntiPhishing credential backend. Valid values: `disable`, `enable`.

* `ca_cert` - CA certificate name.
* `client_cert` - Client certificate name.
* `client_cert_auth` - Enable/disable using client certificate for TLS authentication. Valid values: `disable`, `enable`.

* `cnid` - Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
* `dn` - Distinguished name used to look up entries on the LDAP server.
* `filter` - Filter.
* `group` - Group.
* `group_filter` - Filter used for group matching.
* `group_member_check` - Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.

* `group_object_filter` - Filter used for group searching.
* `group_object_search_base` - Group-Object-Search-Base.
* `group_search_base` - Search base used for group searching.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `max_connections` - Max-Connections.
* `member_attr` - Name of attribute from which to get group membership.
* `obtain_user_info` - Enable/disable obtaining of user information. Valid values: `disable`, `enable`.

* `password` - Password for initial binding.
* `password_attr` - Name of attribute to get password hash.
* `password_expiry_warning` - Enable/disable password expiry warnings. Valid values: `disable`, `enable`.

* `password_renewal` - Enable/disable online password renewal. Valid values: `disable`, `enable`.

* `port` - Port to be used for communication with the LDAP server (default = 389).
* `retrieve_protection_profile` - Retrieve-Protection-Profile.
* `search_type` - Search type. Valid values: `recursive`.

* `secondary_server` - Secondary LDAP server CN domain name or IP.
* `secure` - Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.

* `server` - LDAP server CN domain name or IP.
* `server_identity_check` - Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `disable`, `enable`.

* `source_ip` - FortiGate IP address to be used for communication with the LDAP server.
* `source_ip_interface` - Source interface for communication with the LDAP server.
* `source_port` - Source port to be used for communication with the LDAP server.
* `ssl_max_proto_version` - Ssl-Max-Proto-Version. Valid values: `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1`, `TLSv1-3`.

* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1-3`.

* `status_ttl` - Time for which server reachability is cached so that when a server is unreachable, it will not be retried for at least this period of time (0 = cache disabled, default = 300).
* `tertiary_server` - Tertiary LDAP server CN domain name or IP.
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.

* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.

* `two_factor_filter` - Filter used to synchronize users to FortiToken Cloud.
* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.

* `type` - Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.

* `user_info_exchange_server` - MS Exchange server from which to fetch user information.
* `username` - Username (full DN) for initial binding.
* `validate_server_certificate` - Validate-Server-Certificate. Valid values: `disable`, `enable`.

* `vrf_select` - VRF ID used for connection to server.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Ldap can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_ldap.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

