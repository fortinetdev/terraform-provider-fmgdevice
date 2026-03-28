---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ssl_web_portal"
description: |-
  <i>This object will be purged after policy copy and install.</i> Portal.
---

# fmgdevice_vpn_ssl_web_portal
<i>This object will be purged after policy copy and install.</i> Portal.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `bookmark_group`: `fmgdevice_vpn_ssl_web_portal_bookmarkgroup`
>- `landing_page`: `fmgdevice_vpn_ssl_web_portal_landingpage`
>- `mac_addr_check_rule`: `fmgdevice_vpn_ssl_web_portal_macaddrcheckrule`
>- `os_check_list`: `fmgdevice_vpn_ssl_web_portal_oschecklist`
>- `split_dns`: `fmgdevice_vpn_ssl_web_portal_splitdns`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `allow_user_access` - Allow user access to Agentless VPN applications. Valid values: `web`, `ftp`, `telnet`, `smb`, `vnc`, `rdp`, `ssh`, `ping`, `sftp`.

* `auto_connect` - Auto-Connect. Valid values: `disable`, `enable`.

* `bookmark_group` - Bookmark-Group. The structure of `bookmark_group` block is documented below.
* `client_src_range` - Client-Src-Range. Valid values: `disable`, `enable`.

* `clipboard` - Enable to support RDP/VPC clipboard functionality. Valid values: `disable`, `enable`.

* `custom_lang` - Change the web portal display language. Overrides config system global set language. You can use config system custom-language and execute system custom-language to add custom language files.
* `customize_forticlient_download_url` - Customize-Forticlient-Download-Url. Valid values: `disable`, `enable`.

* `default_protocol` - Application type that is set by default. Valid values: `web`, `ftp`, `telnet`, `smb`, `vnc`, `rdp`, `ssh`, `sftp`.

* `default_window_height` - Screen height (range from 0 - 65535, default = 768).
* `default_window_width` - Screen width (range from 0 - 65535, default = 1024).
* `dhcp_ip_overlap` - Dhcp-Ip-Overlap. Valid values: `use-old`, `use-new`.

* `dhcp_ra_giaddr` - Dhcp-Ra-Giaddr.
* `dhcp_reservation` - Dhcp-Reservation. Valid values: `disable`, `enable`.

* `dhcp6_ra_linkaddr` - Dhcp6-Ra-Linkaddr.
* `display_bookmark` - Enable to display the web portal bookmark widget. Valid values: `disable`, `enable`.

* `display_connection_tools` - Enable to display the web portal connection tools widget. Valid values: `disable`, `enable`.

* `display_history` - Enable to display the web portal user login history widget. Valid values: `disable`, `enable`.

* `display_status` - Enable to display the web portal status widget. Valid values: `disable`, `enable`.

* `dns_server1` - Dns-Server1.
* `dns_server2` - Dns-Server2.
* `dns_suffix` - DNS suffix.
* `exclusive_routing` - Exclusive-Routing. Valid values: `disable`, `enable`.

* `focus_bookmark` - Enable to prioritize the placement of the bookmark section over the quick-connection section in the Agentless VPN application. Valid values: `disable`, `enable`.

* `forticlient_download` - Forticlient-Download. Valid values: `disable`, `enable`.

* `forticlient_download_method` - Forticlient-Download-Method. Valid values: `direct`, `ssl-vpn`.

* `heading` - Web portal heading message.
* `hide_sso_credential` - Enable to prevent SSO credential being sent to client. Valid values: `disable`, `enable`.

* `host_check` - Host-Check. Valid values: `none`, `av`, `fw`, `av-fw`, `custom`.

* `host_check_interval` - Host-Check-Interval.
* `host_check_policy` - Host-Check-Policy.
* `ip_mode` - Ip-Mode. Valid values: `range`, `user-group`, `dhcp`, `no-ip`.

* `ip_pools` - Ip-Pools.
* `ipv6_dns_server1` - Ipv6-Dns-Server1.
* `ipv6_dns_server2` - Ipv6-Dns-Server2.
* `ipv6_exclusive_routing` - Ipv6-Exclusive-Routing. Valid values: `disable`, `enable`.

* `ipv6_pools` - Ipv6-Pools.
* `ipv6_service_restriction` - Ipv6-Service-Restriction. Valid values: `disable`, `enable`.

* `ipv6_split_tunneling` - Ipv6-Split-Tunneling. Valid values: `disable`, `enable`.

* `ipv6_split_tunneling_routing_address` - Ipv6-Split-Tunneling-Routing-Address.
* `ipv6_split_tunneling_routing_negate` - Ipv6-Split-Tunneling-Routing-Negate. Valid values: `disable`, `enable`.

* `ipv6_tunnel_mode` - Ipv6-Tunnel-Mode. Valid values: `disable`, `enable`.

* `ipv6_wins_server1` - Ipv6-Wins-Server1.
* `ipv6_wins_server2` - Ipv6-Wins-Server2.
* `keep_alive` - Keep-Alive. Valid values: `disable`, `enable`.

* `landing_page` - Landing-Page. The structure of `landing_page` block is documented below.
* `landing_page_mode` - Enable/disable Agentless VPN landing page mode. Valid values: `disable`, `enable`.

* `limit_user_logins` - Enable to limit each user to one Agentless VPN session at a time. Valid values: `disable`, `enable`.

* `mac_addr_action` - Mac-Addr-Action. Valid values: `deny`, `allow`.

* `mac_addr_check` - Mac-Addr-Check. Valid values: `disable`, `enable`.

* `mac_addr_check_rule` - Mac-Addr-Check-Rule. The structure of `mac_addr_check_rule` block is documented below.
* `macos_forticlient_download_url` - Macos-Forticlient-Download-Url.
* `name` - Portal name.
* `os_check` - Os-Check. Valid values: `disable`, `enable`.

* `os_check_list` - Os-Check-List. The structure of `os_check_list` block is documented below.
* `prefer_ipv6_dns` - Prefer to query IPv6 DNS server first if enabled. Valid values: `disable`, `enable`.

* `redir_url` - Client login redirect URL.
* `rewrite_ip_uri_ui` - Rewrite contents for URI contains IP and /ui/ (default = disable). Valid values: `disable`, `enable`.

* `save_password` - Save-Password. Valid values: `disable`, `enable`.

* `service_restriction` - Service-Restriction. Valid values: `disable`, `enable`.

* `skip_check_for_browser` - Skip-Check-For-Browser. Valid values: `disable`, `enable`.

* `skip_check_for_unsupported_os` - Skip-Check-For-Unsupported-Os. Valid values: `disable`, `enable`.

* `smb_max_version` - SMB maximum client protocol version. Valid values: `smbv1`, `smbv2`, `smbv3`.

* `smb_min_version` - SMB minimum client protocol version. Valid values: `smbv1`, `smbv2`, `smbv3`.

* `smb_ntlmv1_auth` - Enable support of NTLMv1 for Samba authentication. Valid values: `disable`, `enable`.

* `split_dns` - Split-Dns. The structure of `split_dns` block is documented below.
* `split_tunneling` - Split-Tunneling. Valid values: `disable`, `enable`.

* `split_tunneling_routing_address` - Split-Tunneling-Routing-Address.
* `split_tunneling_routing_negate` - Split-Tunneling-Routing-Negate. Valid values: `disable`, `enable`.

* `theme` - Web portal color scheme. Valid values: `melongene`, `mariner`, `neutrino`, `jade`, `graphite`, `dark-matter`, `onyx`, `eclipse`, `jet-stream`, `security-fabric`.

* `tunnel_mode` - Tunnel-Mode. Valid values: `disable`, `enable`.

* `use_sdwan` - Use SD-WAN rules to get output interface. Valid values: `disable`, `enable`.

* `user_bookmark` - Enable to allow web portal users to create their own bookmarks. Valid values: `disable`, `enable`.

* `user_group_bookmark` - Enable to allow web portal users to create bookmarks for all users in the same user group. Valid values: `disable`, `enable`.

* `web_mode` - Enable/disable Agentless VPN web mode. Valid values: `disable`, `enable`.

* `windows_forticlient_download_url` - Windows-Forticlient-Download-Url.
* `wins_server1` - Wins-Server1.
* `wins_server2` - Wins-Server2.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `bookmark_group` block supports:

* `bookmarks` - Bookmarks. The structure of `bookmarks` block is documented below.
* `name` - Bookmark group name.

The `bookmarks` block supports:

* `additional_params` - Additional parameters.
* `apptype` - Application type. Valid values: `web`, `telnet`, `ssh`, `ftp`, `smb`, `vnc`, `rdp`, `sftp`.

* `color_depth` - Color depth per pixel. Valid values: `8`, `16`, `32`.

* `description` - Description.
* `domain` - Login domain.
* `folder` - Network shared file folder parameter.
* `form_data` - Form-Data. The structure of `form_data` block is documented below.
* `height` - Screen height (range from 0 - 65535, default = 0).
* `host` - Host name/IP parameter.
* `keyboard_layout` - Keyboard layout. Valid values: `da`, `de`, `de-ch`, `en-uk`, `en-us`, `es`, `fi`, `fr`, `fr-be`, `fr-ca`, `fr-ch`, `hr`, `hu`, `it`, `ja`, `lt`, `mk`, `no`, `pt`, `pt-br`, `ru`, `sl`, `sv`, `ar-101`, `ar-102`, `ar-102-azerty`, `can-mul`, `cz`, `cz-qwerty`, `cz-pr`, `nl`, `de-ibm`, `en-uk-ext`, `en-us-dvorak`, `es-var`, `fi-sami`, `hu-101`, `it-142`, `ko`, `lt-ibm`, `lt-std`, `lav-std`, `lav-leg`, `mk-std`, `no-sami`, `pol-214`, `pol-pr`, `pt-br-abnt2`, `ru-mne`, `ru-t`, `sv-sami`, `tuk`, `tur-f`, `tur-q`, `zh-sym-sg-us`, `zh-sym-us`, `zh-tr-hk`, `zh-tr-mo`, `zh-tr-us`, `fr-apple`, `la-am`, `ja-106`.

* `listening_port` - Listening-Port.
* `load_balancing_info` - The load balancing information or cookie which should be provided to the connection broker.
* `logon_password` - Logon password.
* `logon_user` - Logon user.
* `name` - Bookmark name.
* `port` - Remote port.
* `preconnection_blob` - An arbitrary string which identifies the RDP source.
* `preconnection_id` - The numeric ID of the RDP source (0-4294967295).
* `remote_port` - Remote-Port.
* `restricted_admin` - Enable/disable restricted admin mode for RDP. Valid values: `disable`, `enable`.

* `security` - Security mode for RDP connection (default = any). Valid values: `rdp`, `nla`, `tls`, `any`.

* `send_preconnection_id` - Enable/disable sending of preconnection ID. Valid values: `disable`, `enable`.

* `server_layout` - Server-Layout. Valid values: `en-us-qwerty`, `de-de-qwertz`, `fr-fr-azerty`, `it-it-qwerty`, `sv-se-qwerty`, `failsafe`, `en-gb-qwerty`, `es-es-qwerty`, `fr-ch-qwertz`, `ja-jp-qwerty`, `pt-br-qwerty`, `tr-tr-qwerty`, `fr-ca-qwerty`.

* `show_status_window` - Show-Status-Window. Valid values: `disable`, `enable`.

* `sso` - Single sign-on. Valid values: `disable`, `static`, `auto`.

* `sso_credential` - Single sign-on credentials. Valid values: `sslvpn-login`, `alternative`.

* `sso_credential_sent_once` - Single sign-on credentials are only sent once to remote server. Valid values: `disable`, `enable`.

* `sso_password` - SSO password.
* `sso_username` - SSO user name.
* `url` - URL parameter.
* `vnc_keyboard_layout` - Keyboard layout. Valid values: `da`, `de`, `de-ch`, `en-uk`, `es`, `fi`, `fr`, `fr-be`, `it`, `no`, `pt`, `sv`, `nl`, `en-uk-ext`, `it-142`, `pt-br-abnt2`, `default`, `fr-ca-mul`, `gd`, `us-intl`.

* `width` - Screen width (range from 0 - 65535, default = 0).

The `form_data` block supports:

* `name` - Name.
* `value` - Value.

The `landing_page` block supports:

* `form_data` - Form-Data. The structure of `form_data` block is documented below.
* `sso` - Single sign-on. Valid values: `disable`, `static`, `auto`.

* `sso_credential` - Single sign-on credentials. Valid values: `sslvpn-login`, `alternative`.

* `sso_password` - SSO password.
* `sso_username` - SSO user name.
* `url` - Landing page URL.

The `form_data` block supports:

* `name` - Name.
* `value` - Value.

The `mac_addr_check_rule` block supports:

* `mac_addr_list` - Mac-Addr-List.
* `mac_addr_mask` - Mac-Addr-Mask.
* `name` - Name.

The `os_check_list` block supports:

* `action` - Action. Valid values: `allow`, `check-up-to-date`, `deny`.

* `latest_patch_level` - Latest-Patch-Level.
* `minor_version` - Minor-Version.
* `name` - Name.
* `tolerance` - Tolerance.

The `split_dns` block supports:

* `dns_server1` - Dns-Server1.
* `dns_server2` - Dns-Server2.
* `domains` - Domains.
* `id` - Id.
* `ipv6_dns_server1` - Ipv6-Dns-Server1.
* `ipv6_dns_server2` - Ipv6-Dns-Server2.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn SslWebPortal can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ssl_web_portal.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

