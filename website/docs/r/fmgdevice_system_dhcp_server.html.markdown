---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_dhcp_server"
description: |-
  Configure DHCP servers.
---

# fmgdevice_system_dhcp_server
Configure DHCP servers.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `exclude_range`: `fmgdevice_system_dhcp_server_excluderange`
>- `ip_range`: `fmgdevice_system_dhcp_server_iprange`
>- `options`: `fmgdevice_system_dhcp_server_options`
>- `reserved_address`: `fmgdevice_system_dhcp_server_reservedaddress`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auto_configuration` - Enable/disable auto configuration. Valid values: `disable`, `enable`.

* `auto_managed_status` - Enable/disable use of this DHCP server once this interface has been assigned an IP address from FortiIPAM. Valid values: `disable`, `enable`.

* `conflicted_ip_timeout` - Time in seconds to wait after a conflicted IP address is removed from the DHCP range before it can be reused.
* `ddns_auth` - DDNS authentication mode. Valid values: `disable`, `tsig`.

* `ddns_key` - DDNS update key (base 64 encoding).
* `ddns_keyname` - DDNS update key name.
* `ddns_server_ip` - DDNS server IP.
* `ddns_ttl` - TTL.
* `ddns_update` - Enable/disable DDNS update for DHCP. Valid values: `disable`, `enable`.

* `ddns_update_override` - Enable/disable DDNS update override for DHCP. Valid values: `disable`, `enable`.

* `ddns_zone` - Zone of your domain name (ex. DDNS.com).
* `default_gateway` - Default gateway IP address assigned by the DHCP server.
* `dhcp_settings_from_fortiipam` - Enable/disable populating of DHCP server settings from FortiIPAM. Valid values: `disable`, `enable`.

* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `dns_server4` - DNS server 4.
* `dns_service` - Options for assigning DNS servers to DHCP clients. Valid values: `default`, `specify`, `local`.

* `domain` - Domain name suffix for the IP addresses that the DHCP server assigns to clients.
* `exclude_range` - Exclude-Range. The structure of `exclude_range` block is documented below.
* `filename` - Name of the boot file on the TFTP server.
* `forticlient_on_net_status` - Enable/disable FortiClient-On-Net service for this DHCP server. Valid values: `disable`, `enable`.

* `fosid` - ID.
* `interface` - DHCP server can assign IP configurations to clients connected to this interface.
* `ip_mode` - Method used to assign client IP. Valid values: `range`, `usrgrp`.

* `ip_range` - Ip-Range. The structure of `ip_range` block is documented below.
* `ipsec_lease_hold` - DHCP over IPsec leases expire this many seconds after tunnel down (0 to disable forced-expiry).
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `mac_acl_default_action` - MAC access control default action (allow or block assigning IP settings). Valid values: `assign`, `block`.

* `netmask` - Netmask assigned by the DHCP server.
* `next_server` - IP address of a server (for example, a TFTP sever) that DHCP clients can download a boot file from.
* `ntp_server1` - NTP server 1.
* `ntp_server2` - NTP server 2.
* `ntp_server3` - NTP server 3.
* `ntp_service` - Options for assigning Network Time Protocol (NTP) servers to DHCP clients. Valid values: `default`, `specify`, `local`.

* `options` - Options. The structure of `options` block is documented below.
* `relay_agent` - Relay agent IP.
* `reserved_address` - Reserved-Address. The structure of `reserved_address` block is documented below.
* `server_type` - DHCP server can be a normal DHCP server or an IPsec DHCP server. Valid values: `regular`, `ipsec`.

* `shared_subnet` - Enable/disable shared subnet. Valid values: `disable`, `enable`.

* `status` - Enable/disable this DHCP configuration. Valid values: `disable`, `enable`.

* `tftp_server` - One or more hostnames or IP addresses of the TFTP servers in quotes separated by spaces.
* `timezone` - Select the time zone to be assigned to DHCP clients.
* `timezone_option` - Options for the DHCP server to set the client's time zone. Valid values: `disable`, `default`, `specify`.

* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served. Valid values: `disable`, `enable`.

* `vci_string` - One or more VCI strings in quotes separated by spaces.
* `wifi_ac_service` - Options for assigning WiFi access controllers to DHCP clients. Valid values: `specify`, `local`.

* `wifi_ac1` - WiFi Access Controller 1 IP address (DHCP option 138, RFC 5417).
* `wifi_ac2` - WiFi Access Controller 2 IP address (DHCP option 138, RFC 5417).
* `wifi_ac3` - WiFi Access Controller 3 IP address (DHCP option 138, RFC 5417).
* `wins_server1` - WINS server 1.
* `wins_server2` - WINS server 2.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `exclude_range` block supports:

* `end_ip` - End of IP range.
* `id` - ID.
* `lease_time` - Lease time in seconds, 0 means default lease time.
* `start_ip` - Start of IP range.
* `uci_match` - Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this range. Valid values: `disable`, `enable`.

* `uci_string` - One or more UCI strings in quotes separated by spaces.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range. Valid values: `disable`, `enable`.

* `vci_string` - One or more VCI strings in quotes separated by spaces.

The `ip_range` block supports:

* `end_ip` - End of IP range.
* `id` - ID.
* `lease_time` - Lease time in seconds, 0 means default lease time.
* `start_ip` - Start of IP range.
* `uci_match` - Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this range. Valid values: `disable`, `enable`.

* `uci_string` - One or more UCI strings in quotes separated by spaces.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range. Valid values: `disable`, `enable`.

* `vci_string` - One or more VCI strings in quotes separated by spaces.

The `options` block supports:

* `code` - DHCP option code.
* `id` - ID.
* `ip` - DHCP option IPs.
* `type` - DHCP option type. Valid values: `hex`, `string`, `ip`, `fqdn`.

* `uci_match` - Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this option. Valid values: `disable`, `enable`.

* `uci_string` - One or more UCI strings in quotes separated by spaces.
* `value` - DHCP option value.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this option. Valid values: `disable`, `enable`.

* `vci_string` - One or more VCI strings in quotes separated by spaces.

The `reserved_address` block supports:

* `action` - Options for the DHCP server to configure the client with the reserved MAC address. Valid values: `assign`, `block`, `reserved`.

* `circuit_id` - Option 82 circuit-ID of the client that will get the reserved IP address.
* `circuit_id_type` - DHCP option type. Valid values: `hex`, `string`.

* `description` - Description.
* `id` - ID.
* `ip` - IP address to be reserved for the MAC address.
* `mac` - MAC address of the client that will get the reserved IP address.
* `remote_id` - Option 82 remote-ID of the client that will get the reserved IP address.
* `remote_id_type` - DHCP option type. Valid values: `hex`, `string`.

* `type` - DHCP reserved-address type. Valid values: `mac`, `option82`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System DhcpServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_dhcp_server.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

