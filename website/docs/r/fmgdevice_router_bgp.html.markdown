---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_bgp"
description: |-
  Configure BGP.
---

# fmgdevice_router_bgp
Configure BGP.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `admin_distance`: `fmgdevice_router_bgp_admindistance`
>- `aggregate_address`: `fmgdevice_router_bgp_aggregateaddress`
>- `aggregate_address6`: `fmgdevice_router_bgp_aggregateaddress6`
>- `neighbor`: `fmgdevice_router_bgp_neighbor`
>- `neighbor_group`: `fmgdevice_router_bgp_neighborgroup`
>- `neighbor_range`: `fmgdevice_router_bgp_neighborrange`
>- `neighbor_range6`: `fmgdevice_router_bgp_neighborrange6`
>- `network`: `fmgdevice_router_bgp_network`
>- `network6`: `fmgdevice_router_bgp_network6`
>- `redistribute`: `fmgdevice_router_bgp_redistribute`
>- `redistribute6`: `fmgdevice_router_bgp_redistribute6`
>- `vrf`: `fmgdevice_router_bgp_vrf`
>- `vrf6`: `fmgdevice_router_bgp_vrf6`
>- `vrf_leak`: `fmgdevice_router_bgp_vrfleak`
>- `vrf_leak6`: `fmgdevice_router_bgp_vrfleak6`



## Example Usage

```hcl
resource "fmgdevice_router_bgp" "trname" {
  additional_path              = "disable"
  additional_path_select       = 10
  additional_path_select_vpnv4 = 10
  additional_path_select_vpnv6 = 10
  additional_path_select6      = 10
  device_name                  = var.device_name # not required if setting is at provider
  device_vdom                  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `additional_path` - Enable/disable selection of BGP IPv4 additional paths. Valid values: `disable`, `enable`.

* `additional_path_select` - Number of additional paths to be selected for each IPv4 NLRI.
* `additional_path_select_vpnv4` - Number of additional paths to be selected for each VPNv4 NLRI.
* `additional_path_select_vpnv6` - Number of additional paths to be selected for each VPNv6 NLRI.
* `additional_path_select6` - Number of additional paths to be selected for each IPv6 NLRI.
* `additional_path_vpnv4` - Enable/disable selection of BGP VPNv4 additional paths. Valid values: `disable`, `enable`.

* `additional_path_vpnv6` - Enable/disable selection of BGP VPNv6 additional paths. Valid values: `disable`, `enable`.

* `additional_path6` - Enable/disable selection of BGP IPv6 additional paths. Valid values: `disable`, `enable`.

* `admin_distance` - Admin-Distance. The structure of `admin_distance` block is documented below.
* `aggregate_address` - Aggregate-Address. The structure of `aggregate_address` block is documented below.
* `aggregate_address6` - Aggregate-Address6. The structure of `aggregate_address6` block is documented below.
* `always_compare_med` - Enable/disable always compare MED. Valid values: `disable`, `enable`.

* `as` - Router AS number, asplain/asdot/asdot+ format, 0 to disable BGP.
* `bestpath_as_path_ignore` - Enable/disable ignore AS path. Valid values: `disable`, `enable`.

* `bestpath_cmp_confed_aspath` - Enable/disable compare federation AS path length. Valid values: `disable`, `enable`.

* `bestpath_cmp_routerid` - Enable/disable compare router ID for identical EBGP paths. Valid values: `disable`, `enable`.

* `bestpath_med_confed` - Enable/disable compare MED among confederation paths. Valid values: `disable`, `enable`.

* `bestpath_med_missing_as_worst` - Enable/disable treat missing MED as least preferred. Valid values: `disable`, `enable`.

* `client_to_client_reflection` - Enable/disable client-to-client route reflection. Valid values: `disable`, `enable`.

* `cluster_id` - Route reflector cluster ID.
* `confederation_identifier` - Confederation identifier.
* `confederation_peers` - Confederation peers.
* `cross_family_conditional_adv` - Enable/disable cross address family conditional advertisement. Valid values: `disable`, `enable`.

* `dampening` - Enable/disable route-flap dampening. Valid values: `disable`, `enable`.

* `dampening_max_suppress_time` - Maximum minutes a route can be suppressed.
* `dampening_reachability_half_life` - Reachability half-life time for penalty (min).
* `dampening_reuse` - Threshold to reuse routes.
* `dampening_route_map` - Criteria for dampening.
* `dampening_suppress` - Threshold to suppress routes.
* `dampening_unreachability_half_life` - Unreachability half-life time for penalty (min).
* `default_local_preference` - Default local preference.
* `deterministic_med` - Enable/disable enforce deterministic comparison of MED. Valid values: `disable`, `enable`.

* `distance_external` - Distance for routes external to the AS.
* `distance_internal` - Distance for routes internal to the AS.
* `distance_local` - Distance for routes local to the AS.
* `ebgp_multipath` - Enable/disable EBGP multi-path. Valid values: `disable`, `enable`.

* `enforce_first_as` - Enable/disable enforce first AS for EBGP routes. Valid values: `disable`, `enable`.

* `fast_external_failover` - Enable/disable reset peer BGP session if link goes down. Valid values: `disable`, `enable`.

* `graceful_end_on_timer` - Enable/disable to exit graceful restart on timer only. Valid values: `disable`, `enable`.

* `graceful_restart` - Enable/disable BGP graceful restart capabilities. Valid values: `disable`, `enable`.

* `graceful_restart_time` - Time needed for neighbors to restart (sec).
* `graceful_stalepath_time` - Time to hold stale paths of restarting neighbor (sec).
* `graceful_update_delay` - Route advertisement/selection delay after restart (sec).
* `holdtime_timer` - Number of seconds to mark peer as dead.
* `ibgp_multipath` - Enable/disable IBGP multi-path. Valid values: `disable`, `enable`.

* `ignore_optional_capability` - Do not send unknown optional capability notification message. Valid values: `disable`, `enable`.

* `keepalive_timer` - Frequency to send keep alive requests.
* `log_neighbour_changes` - Log BGP neighbor changes. Valid values: `disable`, `enable`.

* `multipath_recursive_distance` - Enable/disable use of recursive distance to select multipath. Valid values: `disable`, `enable`.

* `neighbor` - Neighbor. The structure of `neighbor` block is documented below.
* `neighbor_group` - Neighbor-Group. The structure of `neighbor_group` block is documented below.
* `neighbor_range` - Neighbor-Range. The structure of `neighbor_range` block is documented below.
* `neighbor_range6` - Neighbor-Range6. The structure of `neighbor_range6` block is documented below.
* `network` - Network. The structure of `network` block is documented below.
* `network_import_check` - Enable/disable ensure BGP network route exists in IGP. Valid values: `disable`, `enable`.

* `network6` - Network6. The structure of `network6` block is documented below.
* `recursive_inherit_priority` - Enable/disable priority inheritance for recursive resolution. Valid values: `disable`, `enable`.

* `recursive_next_hop` - Enable/disable recursive resolution of next-hop using BGP route. Valid values: `disable`, `enable`.

* `redistribute` - Redistribute. The structure of `redistribute` block is documented below.
* `redistribute6` - Redistribute6. The structure of `redistribute6` block is documented below.
* `router_id` - Router ID.
* `scan_time` - Background scanner interval (sec), 0 to disable it.
* `synchronization` - Enable/disable only advertise routes from iBGP if routes present in an IGP. Valid values: `disable`, `enable`.

* `tag_resolve_mode` - Configure tag-match mode. Resolves BGP routes with other routes containing the same tag. Valid values: `disable`, `preferred`, `merge`.

* `vrf_leak` - Vrf-Leak. The structure of `vrf_leak` block is documented below.
* `vrf_leak6` - Vrf-Leak6. The structure of `vrf_leak6` block is documented below.
* `vrf` - Vrf. The structure of `vrf` block is documented below.
* `vrf6` - Vrf6. The structure of `vrf6` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `admin_distance` block supports:

* `distance` - Administrative distance to apply (1 - 255).
* `id` - ID.
* `neighbour_prefix` - Neighbor address prefix.
* `route_list` - Access list of routes to apply new distance to.

The `aggregate_address` block supports:

* `as_set` - Enable/disable generate AS set path information. Valid values: `disable`, `enable`.

* `id` - ID.
* `prefix` - Aggregate prefix.
* `summary_only` - Enable/disable filter more specific routes from updates. Valid values: `disable`, `enable`.


The `aggregate_address6` block supports:

* `as_set` - Enable/disable generate AS set path information. Valid values: `disable`, `enable`.

* `id` - ID.
* `prefix6` - Aggregate IPv6 prefix.
* `summary_only` - Enable/disable filter more specific routes from updates. Valid values: `disable`, `enable`.


The `neighbor` block supports:

* `activate` - Enable/disable address family IPv4 for this neighbor. Valid values: `disable`, `enable`.

* `activate_evpn` - Enable/disable address family L2VPN EVPN for this neighbor. Valid values: `disable`, `enable`.

* `activate_vpnv4` - Enable/disable address family VPNv4 for this neighbor. Valid values: `disable`, `enable`.

* `activate_vpnv6` - Enable/disable address family VPNv6 for this neighbor. Valid values: `disable`, `enable`.

* `activate6` - Enable/disable address family IPv6 for this neighbor. Valid values: `disable`, `enable`.

* `additional_path` - Enable/disable IPv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.

* `additional_path_vpnv4` - Enable/disable VPNv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.

* `additional_path_vpnv6` - Enable/disable VPNv6 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.

* `additional_path6` - Enable/disable IPv6 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.

* `adv_additional_path` - Number of IPv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path_vpnv4` - Number of VPNv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path_vpnv6` - Number of VPNv6 additional paths that can be advertised to this neighbor.
* `adv_additional_path6` - Number of IPv6 additional paths that can be advertised to this neighbor.
* `advertisement_interval` - Minimum interval (sec) between sending updates.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path. Valid values: `disable`, `enable`.

* `allowas_in_enable_evpn` - Enable/disable to allow my AS in AS path for L2VPN EVPN route. Valid values: `disable`, `enable`.

* `allowas_in_enable_vpnv4` - Enable/disable to allow my AS in AS path for VPNv4 route. Valid values: `disable`, `enable`.

* `allowas_in_enable_vpnv6` - Enable/disable use of my AS in AS path for VPNv6 route. Valid values: `disable`, `enable`.

* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path. Valid values: `disable`, `enable`.

* `allowas_in_evpn` - The maximum number of occurrence of my AS number allowed for L2VPN EVPN route.
* `allowas_in_vpnv4` - The maximum number of occurrence of my AS number allowed for VPNv4 route.
* `allowas_in_vpnv6` - The maximum number of occurrence of my AS number allowed for VPNv6 route.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4. Valid values: `disable`, `enable`.

* `as_override6` - Enable/disable replace peer AS with own AS for IPv6. Valid values: `disable`, `enable`.

* `attribute_unchanged` - IPv4 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.

* `attribute_unchanged_vpnv4` - List of attributes that should be unchanged for VPNv4 route. Valid values: `as-path`, `med`, `next-hop`.

* `attribute_unchanged_vpnv6` - List of attributes that should not be changed for VPNv6 route. Valid values: `as-path`, `med`, `next-hop`.

* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.

* `auth_options` - Key-chain name for TCP authentication options.
* `bfd` - Enable/disable BFD for this neighbor. Valid values: `disable`, `enable`.

* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor. Valid values: `disable`, `enable`.

* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor. Valid values: `disable`, `enable`.

* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor. Valid values: `disable`, `enable`.

* `capability_graceful_restart` - Enable/disable advertise IPv4 graceful restart capability to this neighbor. Valid values: `disable`, `enable`.

* `capability_graceful_restart_evpn` - Enable/disable advertisement of L2VPN EVPN graceful restart capability to this neighbor. Valid values: `disable`, `enable`.

* `capability_graceful_restart_vpnv4` - Enable/disable advertise VPNv4 graceful restart capability to this neighbor. Valid values: `disable`, `enable`.

* `capability_graceful_restart_vpnv6` - Enable/disable advertisement of VPNv6 graceful restart capability to this neighbor. Valid values: `disable`, `enable`.

* `capability_graceful_restart6` - Enable/disable advertise IPv6 graceful restart capability to this neighbor. Valid values: `disable`, `enable`.

* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor. Valid values: `none`, `send`, `receive`, `both`.

* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor. Valid values: `none`, `send`, `receive`, `both`.

* `capability_route_refresh` - Enable/disable advertise route refresh capability to this neighbor. Valid values: `disable`, `enable`.

* `conditional_advertise` - Conditional-Advertise. The structure of `conditional_advertise` block is documented below.
* `conditional_advertise6` - Conditional-Advertise6. The structure of `conditional_advertise6` block is documented below.
* `connect_timer` - Interval (sec) for connect timer.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `description` - Description.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `distribute_list_in_vpnv4` - Filter for VPNv4 updates from this neighbor.
* `distribute_list_in_vpnv6` - Filter for VPNv6 updates from this neighbor.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `distribute_list_out_vpnv4` - Filter for VPNv4 updates to this neighbor.
* `distribute_list_out_vpnv6` - Filter for VPNv6 updates to this neighbor.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `dont_capability_negotiate` - Do not negotiate capabilities with this neighbor. Valid values: `disable`, `enable`.

* `ebgp_enforce_multihop` - Enable/disable allow multi-hop EBGP neighbors. Valid values: `disable`, `enable`.

* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `filter_list_in` - BGP filter for IPv4 inbound routes.
* `filter_list_in_vpnv4` - BGP filter for VPNv4 inbound routes.
* `filter_list_in_vpnv6` - BGP filter for VPNv6 inbound routes.
* `filter_list_in6` - BGP filter for IPv6 inbound routes.
* `filter_list_out` - BGP filter for IPv4 outbound routes.
* `filter_list_out_vpnv4` - BGP filter for VPNv4 outbound routes.
* `filter_list_out_vpnv6` - BGP filter for VPNv6 outbound routes.
* `filter_list_out6` - BGP filter for IPv6 outbound routes.
* `holdtime_timer` - Interval (sec) before peer considered dead.
* `interface` - Specify outgoing interface for peer connection. For IPv6 peer, the interface should have link-local address.
* `ip` - IP/IPv6 address of neighbor.
* `keep_alive_timer` - Keep alive timer interval (sec).
* `link_down_failover` - Enable/disable failover upon link down. Valid values: `disable`, `enable`.

* `local_as` - Local AS number of neighbor.
* `local_as_no_prepend` - Do not prepend local-as to incoming updates. Valid values: `disable`, `enable`.

* `local_as_replace_as` - Replace real AS with local-as in outgoing updates. Valid values: `disable`, `enable`.

* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `maximum_prefix_evpn` - Maximum number of L2VPN EVPN prefixes to accept from this peer.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_evpn` - Maximum L2VPN EVPN prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_vpnv4` - Maximum VPNv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_vpnv6` - Maximum VPNv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_vpnv4` - Maximum number of VPNv4 prefixes to accept from this peer.
* `maximum_prefix_vpnv6` - Maximum number of VPNv6 prefixes to accept from this peer.
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when limit is exceeded. Valid values: `disable`, `enable`.

* `maximum_prefix_warning_only_evpn` - Enable/disable only sending warning message when exceeding limit of L2VPN EVPN routes. Valid values: `disable`, `enable`.

* `maximum_prefix_warning_only_vpnv4` - Enable/disable only giving warning message when limit is exceeded for VPNv4 routes. Valid values: `disable`, `enable`.

* `maximum_prefix_warning_only_vpnv6` - Enable/disable warning message when limit is exceeded for VPNv6 routes. Valid values: `disable`, `enable`.

* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when limit is exceeded. Valid values: `disable`, `enable`.

* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor. Valid values: `disable`, `enable`.

* `next_hop_self_rr` - Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes. Valid values: `disable`, `enable`.

* `next_hop_self_rr6` - Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes. Valid values: `disable`, `enable`.

* `next_hop_self_vpnv4` - Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor. Valid values: `disable`, `enable`.

* `next_hop_self_vpnv6` - Enable/disable use of outgoing interface's IP address as VPNv6 next-hop for this neighbor. Valid values: `disable`, `enable`.

* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor. Valid values: `disable`, `enable`.

* `override_capability` - Enable/disable override result of capability negotiation. Valid values: `disable`, `enable`.

* `passive` - Enable/disable sending of open messages to this neighbor. Valid values: `disable`, `enable`.

* `password` - Password used in MD5 authentication.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `prefix_list_in_vpnv4` - Inbound filter for VPNv4 updates from this neighbor.
* `prefix_list_in_vpnv6` - Inbound filter for VPNv6 updates from this neighbor.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `prefix_list_out_vpnv4` - Outbound filter for VPNv4 updates to this neighbor.
* `prefix_list_out_vpnv6` - Outbound filter for VPNv6 updates to this neighbor.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `remote_as` - AS number of neighbor.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates. Valid values: `disable`, `enable`.

* `remove_private_as_evpn` - Enable/disable removing private AS number from L2VPN EVPN outbound updates. Valid values: `disable`, `enable`.

* `remove_private_as_vpnv4` - Enable/disable remove private AS number from VPNv4 outbound updates. Valid values: `disable`, `enable`.

* `remove_private_as_vpnv6` - Enable/disable to remove private AS number from VPNv6 outbound updates. Valid values: `disable`, `enable`.

* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates. Valid values: `disable`, `enable`.

* `restart_time` - Graceful restart delay time (sec, 0 = global default).
* `retain_stale_time` - Time to retain stale routes.
* `route_map_in` - IPv4 Inbound route map filter.
* `route_map_in_evpn` - L2VPN EVPN inbound route map filter.
* `route_map_in_vpnv4` - VPNv4 inbound route map filter.
* `route_map_in_vpnv6` - VPNv6 inbound route map filter.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_out` - IPv4 outbound route map filter.
* `route_map_out_evpn` - L2VPN EVPN outbound route map filter.
* `route_map_out_preferable` - IPv4 outbound route map filter if the peer is preferred.
* `route_map_out_vpnv4` - VPNv4 outbound route map filter.
* `route_map_out_vpnv4_preferable` - VPNv4 outbound route map filter if the peer is preferred.
* `route_map_out_vpnv6` - VPNv6 outbound route map filter.
* `route_map_out_vpnv6_preferable` - VPNv6 outbound route map filter if this neighbor is preferred.
* `route_map_out6` - IPv6 Outbound route map filter.
* `route_map_out6_preferable` - IPv6 outbound route map filter if the peer is preferred.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client. Valid values: `disable`, `enable`.

* `route_reflector_client_evpn` - Enable/disable L2VPN EVPN AS route reflector client for this neighbor. Valid values: `disable`, `enable`.

* `route_reflector_client_vpnv4` - Enable/disable VPNv4 AS route reflector client for this neighbor. Valid values: `disable`, `enable`.

* `route_reflector_client_vpnv6` - Enable/disable VPNv6 AS route reflector client for this neighbor. Valid values: `disable`, `enable`.

* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client. Valid values: `disable`, `enable`.

* `route_server_client` - Enable/disable IPv4 AS route server client. Valid values: `disable`, `enable`.

* `route_server_client_evpn` - Enable/disable L2VPN EVPN AS route server client for this neighbor. Valid values: `disable`, `enable`.

* `route_server_client_vpnv4` - Enable/disable VPNv4 AS route server client for this neighbor. Valid values: `disable`, `enable`.

* `route_server_client_vpnv6` - Enable/disable VPNv6 AS route server client for this neighbor. Valid values: `disable`, `enable`.

* `route_server_client6` - Enable/disable IPv6 AS route server client. Valid values: `disable`, `enable`.

* `send_community` - IPv4 Send community attribute to neighbor. Valid values: `disable`, `standard`, `extended`, `both`.

* `send_community_evpn` - Enable/disable sending community attribute to neighbor for L2VPN EVPN address family. Valid values: `disable`, `standard`, `extended`, `both`.

* `send_community_vpnv4` - Send community attribute to neighbor for VPNv4 address family. Valid values: `disable`, `standard`, `extended`, `both`.

* `send_community_vpnv6` - Enable/disable sending community attribute to this neighbor for VPNv6 address family. Valid values: `disable`, `standard`, `extended`, `both`.

* `send_community6` - IPv6 Send community attribute to neighbor. Valid values: `disable`, `standard`, `extended`, `both`.

* `shutdown` - Enable/disable shutdown this neighbor. Valid values: `disable`, `enable`.

* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration. Valid values: `disable`, `enable`.

* `soft_reconfiguration_evpn` - Enable/disable L2VPN EVPN inbound soft reconfiguration. Valid values: `disable`, `enable`.

* `soft_reconfiguration_vpnv4` - Enable/disable allow VPNv4 inbound soft reconfiguration. Valid values: `disable`, `enable`.

* `soft_reconfiguration_vpnv6` - Enable/disable VPNv6 inbound soft reconfiguration. Valid values: `disable`, `enable`.

* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration. Valid values: `disable`, `enable`.

* `stale_route` - Enable/disable stale route after neighbor down. Valid values: `disable`, `enable`.

* `strict_capability_match` - Enable/disable strict capability matching. Valid values: `disable`, `enable`.

* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections.
* `weight` - Neighbor weight.

The `conditional_advertise` block supports:

* `advertise_routemap` - Name of advertising route map.
* `condition_routemap` - List of conditional route maps.
* `condition_type` - Type of condition. Valid values: `exist`, `non-exist`.


The `conditional_advertise6` block supports:

* `advertise_routemap` - Name of advertising route map.
* `condition_routemap` - List of conditional route maps.
* `condition_type` - Type of condition. Valid values: `exist`, `non-exist`.


The `neighbor_group` block supports:

* `activate` - Enable/disable address family IPv4 for this neighbor. Valid values: `disable`, `enable`.

* `activate_evpn` - Enable/disable address family L2VPN EVPN for this neighbor. Valid values: `disable`, `enable`.

* `activate_vpnv4` - Enable/disable address family VPNv4 for this neighbor. Valid values: `disable`, `enable`.

* `activate_vpnv6` - Enable/disable address family VPNv6 for this neighbor. Valid values: `disable`, `enable`.

* `activate6` - Enable/disable address family IPv6 for this neighbor. Valid values: `disable`, `enable`.

* `additional_path` - Enable/disable IPv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.

* `additional_path_vpnv4` - Enable/disable VPNv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.

* `additional_path_vpnv6` - Enable/disable VPNv6 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.

* `additional_path6` - Enable/disable IPv6 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.

* `adv_additional_path` - Number of IPv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path_vpnv4` - Number of VPNv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path_vpnv6` - Number of VPNv6 additional paths that can be advertised to this neighbor.
* `adv_additional_path6` - Number of IPv6 additional paths that can be advertised to this neighbor.
* `advertisement_interval` - Minimum interval (sec) between sending updates.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path. Valid values: `disable`, `enable`.

* `allowas_in_enable_evpn` - Enable/disable to allow my AS in AS path for L2VPN EVPN route. Valid values: `disable`, `enable`.

* `allowas_in_enable_vpnv4` - Enable/disable to allow my AS in AS path for VPNv4 route. Valid values: `disable`, `enable`.

* `allowas_in_enable_vpnv6` - Enable/disable use of my AS in AS path for VPNv6 route. Valid values: `disable`, `enable`.

* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path. Valid values: `disable`, `enable`.

* `allowas_in_evpn` - The maximum number of occurrence of my AS number allowed for L2VPN EVPN route.
* `allowas_in_vpnv4` - The maximum number of occurrence of my AS number allowed for VPNv4 route.
* `allowas_in_vpnv6` - The maximum number of occurrence of my AS number allowed for VPNv6 route.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4. Valid values: `disable`, `enable`.

* `as_override6` - Enable/disable replace peer AS with own AS for IPv6. Valid values: `disable`, `enable`.

* `attribute_unchanged` - IPv4 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.

* `attribute_unchanged_vpnv4` - List of attributes that should be unchanged for VPNv4 route. Valid values: `as-path`, `med`, `next-hop`.

* `attribute_unchanged_vpnv6` - List of attributes that should not be changed for VPNv6 route. Valid values: `as-path`, `med`, `next-hop`.

* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.

* `auth_options` - Key-chain name for TCP authentication options.
* `bfd` - Enable/disable BFD for this neighbor. Valid values: `disable`, `enable`.

* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor. Valid values: `disable`, `enable`.

* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor. Valid values: `disable`, `enable`.

* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor. Valid values: `disable`, `enable`.

* `capability_graceful_restart` - Enable/disable advertise IPv4 graceful restart capability to this neighbor. Valid values: `disable`, `enable`.

* `capability_graceful_restart_evpn` - Enable/disable advertisement of L2VPN EVPN graceful restart capability to this neighbor. Valid values: `disable`, `enable`.

* `capability_graceful_restart_vpnv4` - Enable/disable advertise VPNv4 graceful restart capability to this neighbor. Valid values: `disable`, `enable`.

* `capability_graceful_restart_vpnv6` - Enable/disable advertisement of VPNv6 graceful restart capability to this neighbor. Valid values: `disable`, `enable`.

* `capability_graceful_restart6` - Enable/disable advertise IPv6 graceful restart capability to this neighbor. Valid values: `disable`, `enable`.

* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor. Valid values: `none`, `send`, `receive`, `both`.

* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor. Valid values: `none`, `send`, `receive`, `both`.

* `capability_route_refresh` - Enable/disable advertise route refresh capability to this neighbor. Valid values: `disable`, `enable`.

* `connect_timer` - Interval (sec) for connect timer.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `description` - Description.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `distribute_list_in_vpnv4` - Filter for VPNv4 updates from this neighbor.
* `distribute_list_in_vpnv6` - Filter for VPNv6 updates from this neighbor.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `distribute_list_out_vpnv4` - Filter for VPNv4 updates to this neighbor.
* `distribute_list_out_vpnv6` - Filter for VPNv6 updates to this neighbor.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `dont_capability_negotiate` - Do not negotiate capabilities with this neighbor. Valid values: `disable`, `enable`.

* `ebgp_enforce_multihop` - Enable/disable allow multi-hop EBGP neighbors. Valid values: `disable`, `enable`.

* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `filter_list_in` - BGP filter for IPv4 inbound routes.
* `filter_list_in_vpnv4` - BGP filter for VPNv4 inbound routes.
* `filter_list_in_vpnv6` - BGP filter for VPNv6 inbound routes.
* `filter_list_in6` - BGP filter for IPv6 inbound routes.
* `filter_list_out` - BGP filter for IPv4 outbound routes.
* `filter_list_out_vpnv4` - BGP filter for VPNv4 outbound routes.
* `filter_list_out_vpnv6` - BGP filter for VPNv6 outbound routes.
* `filter_list_out6` - BGP filter for IPv6 outbound routes.
* `holdtime_timer` - Interval (sec) before peer considered dead.
* `interface` - Specify outgoing interface for peer connection. For IPv6 peer, the interface should have link-local address.
* `keep_alive_timer` - Keep alive timer interval (sec).
* `link_down_failover` - Enable/disable failover upon link down. Valid values: `disable`, `enable`.

* `local_as` - Local AS number of neighbor.
* `local_as_no_prepend` - Do not prepend local-as to incoming updates. Valid values: `disable`, `enable`.

* `local_as_replace_as` - Replace real AS with local-as in outgoing updates. Valid values: `disable`, `enable`.

* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `maximum_prefix_evpn` - Maximum number of L2VPN EVPN prefixes to accept from this peer.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_evpn` - Maximum L2VPN EVPN prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_vpnv4` - Maximum VPNv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_vpnv6` - Maximum VPNv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_vpnv4` - Maximum number of VPNv4 prefixes to accept from this peer.
* `maximum_prefix_vpnv6` - Maximum number of VPNv6 prefixes to accept from this peer.
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when limit is exceeded. Valid values: `disable`, `enable`.

* `maximum_prefix_warning_only_evpn` - Enable/disable only sending warning message when exceeding limit of L2VPN EVPN routes. Valid values: `disable`, `enable`.

* `maximum_prefix_warning_only_vpnv4` - Enable/disable only giving warning message when limit is exceeded for VPNv4 routes. Valid values: `disable`, `enable`.

* `maximum_prefix_warning_only_vpnv6` - Enable/disable warning message when limit is exceeded for VPNv6 routes. Valid values: `disable`, `enable`.

* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when limit is exceeded. Valid values: `disable`, `enable`.

* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `name` - Neighbor group name.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor. Valid values: `disable`, `enable`.

* `next_hop_self_rr` - Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes. Valid values: `disable`, `enable`.

* `next_hop_self_rr6` - Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes. Valid values: `disable`, `enable`.

* `next_hop_self_vpnv4` - Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor. Valid values: `disable`, `enable`.

* `next_hop_self_vpnv6` - Enable/disable use of outgoing interface's IP address as VPNv6 next-hop for this neighbor. Valid values: `disable`, `enable`.

* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor. Valid values: `disable`, `enable`.

* `override_capability` - Enable/disable override result of capability negotiation. Valid values: `disable`, `enable`.

* `passive` - Enable/disable sending of open messages to this neighbor. Valid values: `disable`, `enable`.

* `password` - Password used in MD5 authentication.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `prefix_list_in_vpnv4` - Inbound filter for VPNv4 updates from this neighbor.
* `prefix_list_in_vpnv6` - Inbound filter for VPNv6 updates from this neighbor.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `prefix_list_out_vpnv4` - Outbound filter for VPNv4 updates to this neighbor.
* `prefix_list_out_vpnv6` - Outbound filter for VPNv6 updates to this neighbor.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `remote_as` - AS number of neighbor.
* `remote_as_filter` - BGP filter for remote AS.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates. Valid values: `disable`, `enable`.

* `remove_private_as_evpn` - Enable/disable removing private AS number from L2VPN EVPN outbound updates. Valid values: `disable`, `enable`.

* `remove_private_as_vpnv4` - Enable/disable remove private AS number from VPNv4 outbound updates. Valid values: `disable`, `enable`.

* `remove_private_as_vpnv6` - Enable/disable to remove private AS number from VPNv6 outbound updates. Valid values: `disable`, `enable`.

* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates. Valid values: `disable`, `enable`.

* `restart_time` - Graceful restart delay time (sec, 0 = global default).
* `retain_stale_time` - Time to retain stale routes.
* `route_map_in` - IPv4 Inbound route map filter.
* `route_map_in_evpn` - L2VPN EVPN inbound route map filter.
* `route_map_in_vpnv4` - VPNv4 inbound route map filter.
* `route_map_in_vpnv6` - VPNv6 inbound route map filter.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_out` - IPv4 outbound route map filter.
* `route_map_out_evpn` - L2VPN EVPN outbound route map filter.
* `route_map_out_preferable` - IPv4 outbound route map filter if the peer is preferred.
* `route_map_out_vpnv4` - VPNv4 outbound route map filter.
* `route_map_out_vpnv4_preferable` - VPNv4 outbound route map filter if the peer is preferred.
* `route_map_out_vpnv6` - VPNv6 outbound route map filter.
* `route_map_out_vpnv6_preferable` - VPNv6 outbound route map filter if this neighbor is preferred.
* `route_map_out6` - IPv6 Outbound route map filter.
* `route_map_out6_preferable` - IPv6 outbound route map filter if the peer is preferred.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client. Valid values: `disable`, `enable`.

* `route_reflector_client_evpn` - Enable/disable L2VPN EVPN AS route reflector client for this neighbor. Valid values: `disable`, `enable`.

* `route_reflector_client_vpnv4` - Enable/disable VPNv4 AS route reflector client for this neighbor. Valid values: `disable`, `enable`.

* `route_reflector_client_vpnv6` - Enable/disable VPNv6 AS route reflector client for this neighbor. Valid values: `disable`, `enable`.

* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client. Valid values: `disable`, `enable`.

* `route_server_client` - Enable/disable IPv4 AS route server client. Valid values: `disable`, `enable`.

* `route_server_client_evpn` - Enable/disable L2VPN EVPN AS route server client for this neighbor. Valid values: `disable`, `enable`.

* `route_server_client_vpnv4` - Enable/disable VPNv4 AS route server client for this neighbor. Valid values: `disable`, `enable`.

* `route_server_client_vpnv6` - Enable/disable VPNv6 AS route server client for this neighbor. Valid values: `disable`, `enable`.

* `route_server_client6` - Enable/disable IPv6 AS route server client. Valid values: `disable`, `enable`.

* `send_community` - IPv4 Send community attribute to neighbor. Valid values: `disable`, `standard`, `extended`, `both`.

* `send_community_evpn` - Enable/disable sending community attribute to neighbor for L2VPN EVPN address family. Valid values: `disable`, `standard`, `extended`, `both`.

* `send_community_vpnv4` - Send community attribute to neighbor for VPNv4 address family. Valid values: `disable`, `standard`, `extended`, `both`.

* `send_community_vpnv6` - Enable/disable sending community attribute to this neighbor for VPNv6 address family. Valid values: `disable`, `standard`, `extended`, `both`.

* `send_community6` - IPv6 Send community attribute to neighbor. Valid values: `disable`, `standard`, `extended`, `both`.

* `shutdown` - Enable/disable shutdown this neighbor. Valid values: `disable`, `enable`.

* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration. Valid values: `disable`, `enable`.

* `soft_reconfiguration_evpn` - Enable/disable L2VPN EVPN inbound soft reconfiguration. Valid values: `disable`, `enable`.

* `soft_reconfiguration_vpnv4` - Enable/disable allow VPNv4 inbound soft reconfiguration. Valid values: `disable`, `enable`.

* `soft_reconfiguration_vpnv6` - Enable/disable VPNv6 inbound soft reconfiguration. Valid values: `disable`, `enable`.

* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration. Valid values: `disable`, `enable`.

* `stale_route` - Enable/disable stale route after neighbor down. Valid values: `disable`, `enable`.

* `strict_capability_match` - Enable/disable strict capability matching. Valid values: `disable`, `enable`.

* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections.
* `weight` - Neighbor weight.

The `neighbor_range` block supports:

* `id` - Neighbor range ID.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name.
* `prefix` - Neighbor range prefix.

The `neighbor_range6` block supports:

* `id` - IPv6 neighbor range ID.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name.
* `prefix6` - IPv6 prefix.

The `network` block supports:

* `backdoor` - Enable/disable route as backdoor. Valid values: `disable`, `enable`.

* `id` - ID.
* `network_import_check` - Configure insurance of BGP network route existence in IGP. Valid values: `disable`, `enable`, `global`.

* `prefix` - Network prefix.
* `route_map` - Route map to modify generated route.

The `network6` block supports:

* `backdoor` - Enable/disable route as backdoor. Valid values: `disable`, `enable`.

* `id` - ID.
* `network_import_check` - Configure insurance of BGP network route existence in IGP. Valid values: `disable`, `enable`, `global`.

* `prefix6` - Network IPv6 prefix.
* `route_map` - Route map to modify generated route.

The `redistribute` block supports:

* `name` - Distribute list entry name.
* `route_map` - Route map name.
* `status` - Status. Valid values: `disable`, `enable`.


The `redistribute6` block supports:

* `name` - Distribute list entry name.
* `route_map` - Route map name.
* `status` - Status. Valid values: `disable`, `enable`.


The `vrf_leak` block supports:

* `target` - Target. The structure of `target` block is documented below.
* `vrf` - Origin VRF ID (0 - 31).

The `target` block supports:

* `interface` - Interface which is used to leak routes to target VRF.
* `route_map` - Route map of VRF leaking.
* `vrf` - Target VRF ID (0 - 31).

The `vrf_leak6` block supports:

* `target` - Target. The structure of `target` block is documented below.
* `vrf` - Origin VRF ID (0 - 31).

The `target` block supports:

* `interface` - Interface which is used to leak routes to target VRF.
* `route_map` - Route map of VRF leaking.
* `vrf` - Target VRF ID (0 - 31).

The `vrf` block supports:

* `export_rt` - List of export route target.
* `import_route_map` - Import route map.
* `import_rt` - List of import route target.
* `leak_target` - Leak-Target. The structure of `leak_target` block is documented below.
* `rd` - Route Distinguisher: AA:NN|A.B.C.D:NN.
* `role` - VRF role. Valid values: `standalone`, `ce`, `pe`.

* `vrf` - Origin VRF ID (0 - 251).

The `leak_target` block supports:

* `interface` - Interface which is used to leak routes to target VRF.
* `route_map` - Route map of VRF leaking.
* `vrf` - Target VRF ID (0 - 251).

The `vrf6` block supports:

* `export_rt` - List of export route target.
* `import_route_map` - Import route map.
* `import_rt` - List of import route target.
* `leak_target` - Leak-Target. The structure of `leak_target` block is documented below.
* `rd` - Route Distinguisher: AA:NN|A.B.C.D:NN.
* `role` - VRF role. Valid values: `standalone`, `ce`, `pe`.

* `vrf` - Origin VRF ID (0 - 251).

The `leak_target` block supports:

* `interface` - Interface which is used to leak routes to target VRF.
* `route_map` - Route map of VRF leaking.
* `vrf` - Target VRF ID (0 - 251).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Bgp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_bgp.labelname RouterBgp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

