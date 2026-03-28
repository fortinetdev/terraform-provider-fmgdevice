---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_proxypolicy"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure proxy policies.
---

# fmgdevice_firewall_proxypolicy
<i>This object will be purged after policy copy and install.</i> Configure proxy policies.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `_policy_block` - _Policy_Block.
* `access_proxy` - IPv4 access proxy.
* `access_proxy6` - IPv6 access proxy.
* `action` - Accept or deny traffic matching the policy parameters. Valid values: `accept`, `deny`, `redirect`, `isolate`.

* `application_list` - Name of an existing Application list.
* `av_profile` - Name of an existing Antivirus profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `block_notification` - Enable/disable block notification. Valid values: `disable`, `enable`.

* `casb_profile` - Name of an existing CASB profile.
* `comments` - Optional comments.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `detect_https_in_http_request` - Enable/disable detection of HTTPS in HTTP request. Valid values: `disable`, `enable`.

* `device_ownership` - When enabled, the ownership enforcement will be done at policy level. Valid values: `disable`, `enable`.

* `diameter_filter_profile` - Name of an existing Diameter filter profile.
* `disclaimer` - Web proxy disclaimer setting: by domain, policy, or user. Valid values: `disable`, `domain`, `policy`, `user`.

* `dlp_sensor` - Name of an existing DLP sensor.
* `dlp_profile` - Name of an existing DLP profile.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `dstaddr` - Destination address objects.
* `dstaddr_negate` - When enabled, destination addresses match against any address EXCEPT the specified destination addresses. Valid values: `disable`, `enable`.

* `dstaddr6` - IPv6 destination address objects.
* `dstintf` - Destination interface names.
* `emailfilter_profile` - Name of an existing email filter profile.
* `file_filter_profile` - Name of an existing file-filter profile.
* `global_label` - Global web-based manager visible label.
* `groups` - Names of group objects.
* `http_tunnel_auth` - Enable/disable HTTP tunnel authentication. Valid values: `disable`, `enable`.

* `https_sub_category` - Enable/disable HTTPS sub-category policy matching. Valid values: `disable`, `enable`.

* `icap_profile` - Name of an existing ICAP profile.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `disable`, `enable`.

* `internet_service_custom` - Custom Internet Service name.
* `internet_service_custom_group` - Custom Internet Service group name.
* `internet_service_fortiguard` - FortiGuard Internet Service name.
* `internet_service_group` - Internet Service group name.
* `internet_service_name` - Internet Service name.
* `internet_service_negate` - When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service. Valid values: `disable`, `enable`.

* `internet_service6` - Enable/disable use of Internet Services IPv6 for this policy. If enabled, destination IPv6 address and service are not used. Valid values: `disable`, `enable`.

* `internet_service6_custom` - Custom Internet Service IPv6 name.
* `internet_service6_custom_group` - Custom Internet Service IPv6 group name.
* `internet_service6_fortiguard` - FortiGuard Internet Service IPv6 name.
* `internet_service6_group` - Internet Service IPv6 group name.
* `internet_service6_name` - Internet Service IPv6 name.
* `internet_service6_negate` - When enabled, Internet Services match against any internet service IPv6 EXCEPT the selected Internet Service IPv6. Valid values: `disable`, `enable`.

* `ips_sensor` - Name of an existing IPS sensor.
* `ips_voip_filter` - Name of an existing VoIP (ips) profile.
* `isolator_server` - Isolator server name.
* `label` - VDOM-specific GUI visible label.
* `log_http_transaction` - Enable/disable HTTP transaction log. Valid values: `disable`, `enable`.

* `logtraffic` - Enable/disable logging traffic through the policy. Valid values: `disable`, `all`, `utm`.

* `logtraffic_start` - Enable/disable policy log traffic start. Valid values: `disable`, `enable`.

* `name` - Policy name.
* `policyid` - Policy ID.
* `poolname` - Name of IP pool object.
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single`, `group`.

* `proxy` - Type of explicit proxy. Valid values: `explicit-web`, `transparent-web`, `ftp`, `wanopt`, `ssh`, `ssh-tunnel`, `access-proxy`, `ztna-proxy`.

* `redirect_url` - Redirect URL for further explicit web proxy processing.
* `replacemsg_override_group` - Authentication replacement message override group.
* `schedule` - Name of schedule object.
* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `service` - Name of service objects.
* `service_negate` - When enabled, services match against any service EXCEPT the specified destination services. Valid values: `disable`, `enable`.

* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `srcaddr` - Source address objects.
* `srcaddr_negate` - When enabled, source addresses match against any address EXCEPT the specified source addresses. Valid values: `disable`, `enable`.

* `srcaddr6` - IPv6 source address objects.
* `srcintf` - Source interface names.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy. Valid values: `disable`, `enable`.

* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `status` - Enable/disable the active status of the policy. Valid values: `disable`, `enable`.

* `telemetry_profile` - Name of an existing telemetry profile.
* `transparent` - Enable to use the IP address of the client to connect to the server. Valid values: `disable`, `enable`.

* `url_risk` - URL risk level name.
* `users` - Names of user objects.
* `utm_status` - Enable the use of UTM profiles/sensors/lists. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `voip_profile` - Name of an existing VoIP profile.
* `virtual_patch_profile` - Virtual-Patch-Profile.
* `waf_profile` - Name of an existing Web application firewall profile.
* `webcache` - Enable/disable web caching. Valid values: `disable`, `enable`.

* `webcache_https` - Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile). Valid values: `disable`, `enable`.

* `webfilter_profile` - Name of an existing Web filter profile.
* `webproxy_forward_server` - Web proxy forward server name.
* `webproxy_profile` - Name of web proxy profile.
* `ztna_ems_tag` - ZTNA EMS Tag names.
* `ztna_ems_tag_negate` - When enabled, ZTNA EMS tags match against any tag EXCEPT the specified ZTNA EMS tags. Valid values: `disable`, `enable`.

* `ztna_proxy` - ZTNA proxies.
* `ztna_tags_match_logic` - ZTNA tag matching logic. Valid values: `or`, `and`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall ProxyPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_proxypolicy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

