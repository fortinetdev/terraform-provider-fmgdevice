---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_vip"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure virtual IP for IPv4.
---

# fmgdevice_firewall_vip
<i>This object will be purged after policy copy and install.</i> Configure virtual IP for IPv4.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fmgdevice_firewall_vip_dynamic_mapping`
>- `gslb_public_ips`: `fmgdevice_firewall_vip_gslbpublicips`
>- `quic`: `fmgdevice_firewall_vip_quic`
>- `realservers`: `fmgdevice_firewall_vip_realservers`
>- `ssl_cipher_suites`: `fmgdevice_firewall_vip_sslciphersuites`
>- `ssl_server_cipher_suites`: `fmgdevice_firewall_vip_sslserverciphersuites`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `add_nat46_route` - Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.

* `arp_reply` - Enable to respond to ARP requests for this virtual IP address. Enabled by default. Valid values: `disable`, `enable`.

* `client_cert` - Enable/disable requesting client certificate. Valid values: `disable`, `enable`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `dns_mapping_ttl` - DNS mapping TTL (Set to zero to use TTL in DNS response, default = 0).
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `empty_cert_action` - Action for an empty client certificate. Valid values: `accept`, `block`, `accept-unmanageable`.

* `extaddr` - External FQDN address name.
* `extintf` - Interface connected to the source network that receives the packets that will be forwarded to the destination network.
* `extip` - IP address or address range on the external interface that you want to map to an address or address range on the destination network.
* `extport` - Incoming port number range that you want to map to a port number range on the destination network.
* `gratuitous_arp_interval` - Enable to have the VIP send gratuitous ARPs. 0=disabled. Set from 5 up to 8640000 seconds to enable.
* `gslb_domain_name` - Domain to use when integrating with FortiGSLB.
* `gslb_hostname` - Hostname to use within the configured FortiGSLB domain.
* `gslb_public_ips` - Gslb-Public-Ips. The structure of `gslb_public_ips` block is documented below.
* `h2_support` - Enable/disable HTTP2 support (default = enable). Valid values: `disable`, `enable`.

* `h3_support` - Enable/disable HTTP3/QUIC support (default = disable). Valid values: `disable`, `enable`.

* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP. Valid values: `disable`, `enable`.

* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_share` - Control sharing of cookies across virtual servers. Use of same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing. Valid values: `disable`, `same-ip`.

* `http_ip_header` - For HTTP multiplexing, enable to add the original client IP address in the X-Forwarded-For HTTP header. Valid values: `disable`, `enable`.

* `http_ip_header_name` - For HTTP multiplexing, enter a custom HTTPS header name. The original client IP address is added to this header. If empty, X-Forwarded-For is used.
* `http_multiplex` - Enable/disable HTTP multiplexing. Valid values: `disable`, `enable`.

* `http_multiplex_max_concurrent_request` - Maximum number of concurrent requests that a multiplex server can handle (default = unlimited).
* `http_multiplex_max_request` - Maximum number of requests that a multiplex server can handle before disconnecting sessions (default = unlimited).
* `http_multiplex_ttl` - Time-to-live for idle connections to servers.
* `http_redirect` - Enable/disable redirection of HTTP to HTTPS. Valid values: `disable`, `enable`.

* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure. Valid values: `disable`, `enable`.

* `fosid` - Custom defined ID.
* `ipv6_mappedip` - Range of mapped IPv6 addresses. Specify the start IPv6 address followed by a space and the end IPv6 address.
* `ipv6_mappedport` - IPv6 port number range on the destination network to which the external port number range is mapped.
* `ldb_method` - Method used to distribute sessions to real servers. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`, `http-host`.

* `mapped_addr` - Mapped FQDN address name.
* `mappedip` - IP address or address range on the destination network to which the external IP address is mapped.
* `mappedport` - Port number range on the destination network to which the external port number range is mapped.
* `max_embryonic_connections` - Maximum number of incomplete connections.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
* `name` - Virtual IP name.
* `nat_source_vip` - Enable/disable forcing the source NAT mapped IP to the external IP for all traffic. Valid values: `disable`, `enable`.

* `nat44` - Enable/disable NAT44. Valid values: `disable`, `enable`.

* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.

* `one_click_gslb_server` - Enable/disable one click GSLB server integration with FortiGSLB. Valid values: `disable`, `enable`.

* `outlook_web_access` - Enable to add the Front-End-Https header for Microsoft Outlook Web Access. Valid values: `disable`, `enable`.

* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session. Valid values: `none`, `http-cookie`, `ssl-session-id`.

* `portforward` - Enable/disable port forwarding. Valid values: `disable`, `enable`.

* `portmapping_type` - Port mapping type. Valid values: `1-to-1`, `m-to-n`.

* `protocol` - Protocol to use when forwarding packets. Valid values: `tcp`, `udp`, `sctp`, `icmp`.

* `quic` - Quic. The structure of `quic` block is documented below.
* `realservers` - Realservers. The structure of `realservers` block is documented below.
* `server_type` - Protocol to be load balanced by the virtual server (also called the server load balance virtual IP). Valid values: `http`, `https`, `ssl`, `tcp`, `udp`, `ip`, `imaps`, `pop3s`, `smtps`.

* `service` - Service name.
* `src_filter` - Source address filter. Each address must be either an IP/subnet (x.x.x.x/n) or a range (x.x.x.x-y.y.y.y). Separate addresses with spaces.
* `src_vip_filter` - Enable/disable use of 'src-filter' to match destinations for the reverse SNAT rule. Valid values: `disable`, `enable`.

* `srcintf_filter` - Interfaces to which the VIP applies. Separate the names with spaces.
* `ssl_accept_ffdhe_groups` - Enable/disable FFDHE cipher suite for SSL key exchange. Valid values: `disable`, `enable`.

* `ssl_algorithm` - Permitted encryption algorithms for SSL sessions according to encryption strength. Valid values: `high`, `medium`, `low`, `custom`.

* `ssl_certificate` - Name of the certificate to use for SSL handshake.
* `ssl_cipher_suites` - Ssl-Cipher-Suites. The structure of `ssl_cipher_suites` block is documented below.
* `ssl_client_fallback` - Enable/disable support for preventing Downgrade Attacks on client connections (RFC 7507). Valid values: `disable`, `enable`.

* `ssl_client_rekey_count` - Maximum length of data in MB before triggering a client rekey (0 = disable).
* `ssl_client_renegotiation` - Allow, deny, or require secure renegotiation of client sessions to comply with RFC 5746. Valid values: `deny`, `allow`, `secure`.

* `ssl_client_session_state_max` - Maximum number of client to FortiGate SSL session states to keep.
* `ssl_client_session_state_timeout` - Number of minutes to keep client to FortiGate SSL session state.
* `ssl_client_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the client and the FortiGate. Valid values: `disable`, `time`, `count`, `both`.

* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions. Valid values: `768`, `1024`, `1536`, `2048`, `3072`, `4096`.

* `ssl_hpkp` - Enable/disable including HPKP header in response. Valid values: `disable`, `enable`, `report-only`.

* `ssl_hpkp_age` - Number of seconds the client should honor the HPKP setting.
* `ssl_hpkp_backup` - Certificate to generate backup HPKP pin from.
* `ssl_hpkp_include_subdomains` - Indicate that HPKP header applies to all subdomains. Valid values: `disable`, `enable`.

* `ssl_hpkp_primary` - Certificate to generate primary HPKP pin from.
* `ssl_hpkp_report_uri` - URL to report HPKP violations to.
* `ssl_hsts` - Enable/disable including HSTS header in response. Valid values: `disable`, `enable`.

* `ssl_hsts_age` - Number of seconds the client should honor the HSTS setting.
* `ssl_hsts_include_subdomains` - Indicate that HSTS header applies to all subdomains. Valid values: `disable`, `enable`.

* `ssl_http_location_conversion` - Enable to replace HTTP with HTTPS in the reply's Location HTTP header field. Valid values: `disable`, `enable`.

* `ssl_http_match_host` - Enable/disable HTTP host matching for location conversion. Valid values: `disable`, `enable`.

* `ssl_max_version` - Highest SSL/TLS version acceptable from a client. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_min_version` - Lowest SSL/TLS version acceptable from a client. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_mode` - Apply SSL offloading between the client and the FortiGate (half) or from the client to the FortiGate and from the FortiGate to the server (full). Valid values: `half`, `full`.

* `ssl_pfs` - Select the cipher suites that can be used for SSL perfect forward secrecy (PFS). Applies to both client and server sessions. Valid values: `require`, `deny`, `allow`.

* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid CBC IV attacks (SSL 3.0 & TLS 1.0 only). May need to be disabled for compatibility with older systems. Valid values: `disable`, `enable`.

* `ssl_server_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength. Valid values: `high`, `low`, `medium`, `custom`, `client`.

* `ssl_server_cipher_suites` - Ssl-Server-Cipher-Suites. The structure of `ssl_server_cipher_suites` block is documented below.
* `ssl_server_max_version` - Highest SSL/TLS version acceptable from a server. Use the client setting by default. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `client`, `tls-1.3`.

* `ssl_server_min_version` - Lowest SSL/TLS version acceptable from a server. Use the client setting by default. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `client`, `tls-1.3`.

* `ssl_server_renegotiation` - Enable/disable secure renegotiation to comply with RFC 5746. Valid values: `disable`, `enable`.

* `ssl_server_session_state_max` - Maximum number of FortiGate to Server SSL session states to keep.
* `ssl_server_session_state_timeout` - Number of minutes to keep FortiGate to Server SSL session state.
* `ssl_server_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the server and the FortiGate. Valid values: `disable`, `time`, `count`, `both`.

* `status` - Enable/disable VIP. Valid values: `disable`, `enable`.

* `type` - Configure a static NAT, load balance, server load balance, access proxy, DNS translation, or FQDN VIP. Valid values: `static-nat`, `load-balance`, `server-load-balance`, `dns-translation`, `fqdn`, `access-proxy`.

* `user_agent_detect` - Enable/disable detecting device type by HTTP user-agent if no client certificate is provided. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `weblogic_server` - Enable to add an HTTP header to indicate SSL offloading for a WebLogic server. Valid values: `disable`, `enable`.

* `websphere_server` - Enable to add an HTTP header to indicate SSL offloading for a WebSphere server. Valid values: `disable`, `enable`.

* `http_supported_max_version` - Http-Supported-Max-Version. Valid values: `http1`, `http2`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `add_nat46_route` - Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.

* `arp_reply` - Enable to respond to ARP requests for this virtual IP address. Enabled by default. Valid values: `disable`, `enable`.

* `client_cert` - Enable/disable requesting client certificate. Valid values: `disable`, `enable`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `dns_mapping_ttl` - DNS mapping TTL (Set to zero to use TTL in DNS response, default = 0).
* `empty_cert_action` - Action for an empty client certificate. Valid values: `accept`, `block`, `accept-unmanageable`.

* `extaddr` - External FQDN address name.
* `extintf` - Interface connected to the source network that receives the packets that will be forwarded to the destination network.
* `extip` - IP address or address range on the external interface that you want to map to an address or address range on the destination network.
* `extport` - Incoming port number range that you want to map to a port number range on the destination network.
* `gratuitous_arp_interval` - Enable to have the VIP send gratuitous ARPs. 0=disabled. Set from 5 up to 8640000 seconds to enable.
* `gslb_domain_name` - Domain to use when integrating with FortiGSLB.
* `gslb_hostname` - Hostname to use within the configured FortiGSLB domain.
* `h2_support` - Enable/disable HTTP2 support (default = enable). Valid values: `disable`, `enable`.

* `h3_support` - Enable/disable HTTP3/QUIC support (default = disable). Valid values: `disable`, `enable`.

* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP. Valid values: `disable`, `enable`.

* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_share` - Control sharing of cookies across virtual servers. Use of same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing. Valid values: `disable`, `same-ip`.

* `http_ip_header` - For HTTP multiplexing, enable to add the original client IP address in the X-Forwarded-For HTTP header. Valid values: `disable`, `enable`.

* `http_ip_header_name` - For HTTP multiplexing, enter a custom HTTPS header name. The original client IP address is added to this header. If empty, X-Forwarded-For is used.
* `http_multiplex` - Enable/disable HTTP multiplexing. Valid values: `disable`, `enable`.

* `http_multiplex_max_concurrent_request` - Maximum number of concurrent requests that a multiplex server can handle (default = unlimited).
* `http_multiplex_max_request` - Maximum number of requests that a multiplex server can handle before disconnecting sessions (default = unlimited).
* `http_multiplex_ttl` - Time-to-live for idle connections to servers.
* `http_redirect` - Enable/disable redirection of HTTP to HTTPS. Valid values: `disable`, `enable`.

* `http_supported_max_version` - Http-Supported-Max-Version. Valid values: `http1`, `http2`.

* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure. Valid values: `disable`, `enable`.

* `id` - Custom defined ID.
* `ipv6_mappedip` - Range of mapped IPv6 addresses. Specify the start IPv6 address followed by a space and the end IPv6 address.
* `ipv6_mappedport` - IPv6 port number range on the destination network to which the external port number range is mapped.
* `ldb_method` - Method used to distribute sessions to real servers. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`, `http-host`.

* `mapped_addr` - Mapped FQDN address name.
* `mappedip` - IP address or address range on the destination network to which the external IP address is mapped.
* `mappedport` - Port number range on the destination network to which the external port number range is mapped.
* `max_embryonic_connections` - Maximum number of incomplete connections.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
* `nat_source_vip` - Enable/disable forcing the source NAT mapped IP to the external IP for all traffic. Valid values: `disable`, `enable`.

* `nat44` - Enable/disable NAT44. Valid values: `disable`, `enable`.

* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.

* `one_click_gslb_server` - Enable/disable one click GSLB server integration with FortiGSLB. Valid values: `disable`, `enable`.

* `outlook_web_access` - Enable to add the Front-End-Https header for Microsoft Outlook Web Access. Valid values: `disable`, `enable`.

* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session. Valid values: `none`, `http-cookie`, `ssl-session-id`.

* `portforward` - Enable/disable port forwarding. Valid values: `disable`, `enable`.

* `portmapping_type` - Port mapping type. Valid values: `1-to-1`, `m-to-n`.

* `protocol` - Protocol to use when forwarding packets. Valid values: `tcp`, `udp`, `sctp`, `icmp`.

* `realservers` - Realservers. The structure of `realservers` block is documented below.
* `server_type` - Protocol to be load balanced by the virtual server (also called the server load balance virtual IP). Valid values: `http`, `https`, `ssl`, `tcp`, `udp`, `ip`, `imaps`, `pop3s`, `smtps`.

* `service` - Service name.
* `src_filter` - Source address filter. Each address must be either an IP/subnet (x.x.x.x/n) or a range (x.x.x.x-y.y.y.y). Separate addresses with spaces.
* `src_vip_filter` - Enable/disable use of 'src-filter' to match destinations for the reverse SNAT rule. Valid values: `disable`, `enable`.

* `srcintf_filter` - Interfaces to which the VIP applies. Separate the names with spaces.
* `ssl_accept_ffdhe_groups` - Enable/disable FFDHE cipher suite for SSL key exchange. Valid values: `disable`, `enable`.

* `ssl_algorithm` - Permitted encryption algorithms for SSL sessions according to encryption strength. Valid values: `high`, `medium`, `low`, `custom`.

* `ssl_certificate` - Name of the certificate to use for SSL handshake.
* `ssl_cipher_suites` - Ssl-Cipher-Suites. The structure of `ssl_cipher_suites` block is documented below.
* `ssl_client_fallback` - Enable/disable support for preventing Downgrade Attacks on client connections (RFC 7507). Valid values: `disable`, `enable`.

* `ssl_client_rekey_count` - Maximum length of data in MB before triggering a client rekey (0 = disable).
* `ssl_client_renegotiation` - Allow, deny, or require secure renegotiation of client sessions to comply with RFC 5746. Valid values: `deny`, `allow`, `secure`.

* `ssl_client_session_state_max` - Maximum number of client to FortiGate SSL session states to keep.
* `ssl_client_session_state_timeout` - Number of minutes to keep client to FortiGate SSL session state.
* `ssl_client_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the client and the FortiGate. Valid values: `disable`, `time`, `count`, `both`.

* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions. Valid values: `768`, `1024`, `1536`, `2048`, `3072`, `4096`.

* `ssl_hpkp` - Enable/disable including HPKP header in response. Valid values: `disable`, `enable`, `report-only`.

* `ssl_hpkp_age` - Number of seconds the client should honor the HPKP setting.
* `ssl_hpkp_backup` - Certificate to generate backup HPKP pin from.
* `ssl_hpkp_include_subdomains` - Indicate that HPKP header applies to all subdomains. Valid values: `disable`, `enable`.

* `ssl_hpkp_primary` - Certificate to generate primary HPKP pin from.
* `ssl_hpkp_report_uri` - URL to report HPKP violations to.
* `ssl_hsts` - Enable/disable including HSTS header in response. Valid values: `disable`, `enable`.

* `ssl_hsts_age` - Number of seconds the client should honor the HSTS setting.
* `ssl_hsts_include_subdomains` - Indicate that HSTS header applies to all subdomains. Valid values: `disable`, `enable`.

* `ssl_http_location_conversion` - Enable to replace HTTP with HTTPS in the reply's Location HTTP header field. Valid values: `disable`, `enable`.

* `ssl_http_match_host` - Enable/disable HTTP host matching for location conversion. Valid values: `disable`, `enable`.

* `ssl_max_version` - Highest SSL/TLS version acceptable from a client. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_min_version` - Lowest SSL/TLS version acceptable from a client. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_mode` - Apply SSL offloading between the client and the FortiGate (half) or from the client to the FortiGate and from the FortiGate to the server (full). Valid values: `half`, `full`.

* `ssl_pfs` - Select the cipher suites that can be used for SSL perfect forward secrecy (PFS). Applies to both client and server sessions. Valid values: `require`, `deny`, `allow`.

* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid CBC IV attacks (SSL 3.0 & TLS 1.0 only). May need to be disabled for compatibility with older systems. Valid values: `disable`, `enable`.

* `ssl_server_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength. Valid values: `high`, `low`, `medium`, `custom`, `client`.

* `ssl_server_max_version` - Highest SSL/TLS version acceptable from a server. Use the client setting by default. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `client`, `tls-1.3`.

* `ssl_server_min_version` - Lowest SSL/TLS version acceptable from a server. Use the client setting by default. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `client`, `tls-1.3`.

* `ssl_server_renegotiation` - Enable/disable secure renegotiation to comply with RFC 5746. Valid values: `disable`, `enable`.

* `ssl_server_session_state_max` - Maximum number of FortiGate to Server SSL session states to keep.
* `ssl_server_session_state_timeout` - Number of minutes to keep FortiGate to Server SSL session state.
* `ssl_server_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the server and the FortiGate. Valid values: `disable`, `time`, `count`, `both`.

* `status` - Enable/disable VIP. Valid values: `disable`, `enable`.

* `type` - Configure a static NAT, load balance, server load balance, access proxy, DNS translation, or FQDN VIP. Valid values: `static-nat`, `load-balance`, `server-load-balance`, `dns-translation`, `fqdn`, `access-proxy`.

* `user_agent_detect` - Enable/disable detecting device type by HTTP user-agent if no client certificate is provided. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `weblogic_server` - Enable to add an HTTP header to indicate SSL offloading for a WebLogic server. Valid values: `disable`, `enable`.

* `websphere_server` - Enable to add an HTTP header to indicate SSL offloading for a WebSphere server. Valid values: `disable`, `enable`.


The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `realservers` block supports:

* `address` - Dynamic address of the real server.
* `client_ip` - Only clients in this IP range can connect to this real server.
* `health_check_proto` - Health-Check-Proto. Valid values: `ping`, `http`.

* `healthcheck` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`, `vip`.

* `holddown_interval` - Time in seconds that the system waits before re-activating a previously down active server in the active-standby mode. This is to prevent any flapping issues.
* `http_host` - HTTP server domain name in HTTP header.
* `id` - Real server ID.
* `ip` - IP address of the real server.
* `max_connections` - Max number of active connections that can be directed to the real server. When reached, sessions are sent to other real servers.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
* `port` - Port for communicating with the real server. Required if port forwarding is enabled.
* `seq` - Seq.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.

* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server. Valid values: `disable`, `enable`.

* `type` - Type of address. Valid values: `ip`, `address`.

* `verify_cert` - Enable/disable certificate verification of the real server. Valid values: `disable`, `enable`.

* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.

The `ssl_cipher_suites` block supports:

* `cipher` - Cipher suite name. Valid values: `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA`.

* `id` - Id.
* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.


The `gslb_public_ips` block supports:

* `index` - Index of this public IP setting.
* `ip` - The publicly accessible IP address.

The `quic` block supports:

* `ack_delay_exponent` - ACK delay exponent (1 - 20, default = 3).
* `active_connection_id_limit` - Active connection ID limit (1 - 8, default = 2).
* `active_migration` - Enable/disable active migration (default = disable). Valid values: `disable`, `enable`.

* `grease_quic_bit` - Enable/disable grease QUIC bit (default = enable). Valid values: `disable`, `enable`.

* `max_ack_delay` - Maximum ACK delay in milliseconds (1 - 16383, default = 25).
* `max_datagram_frame_size` - Maximum datagram frame size in bytes (1 - 1500, default = 1500).
* `max_idle_timeout` - Maximum idle timeout milliseconds (1 - 60000, default = 30000).
* `max_udp_payload_size` - Maximum UDP payload size in bytes (1200 - 1500, default = 1500).

The `realservers` block supports:

* `address` - Dynamic address of the real server.
* `client_ip` - Only clients in this IP range can connect to this real server.
* `healthcheck` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`, `vip`.

* `holddown_interval` - Time in seconds that the system waits before re-activating a previously down active server in the active-standby mode. This is to prevent any flapping issues.
* `http_host` - HTTP server domain name in HTTP header.
* `id` - Real server ID.
* `ip` - IP address of the real server.
* `max_connections` - Max number of active connections that can be directed to the real server. When reached, sessions are sent to other real servers.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
* `port` - Port for communicating with the real server. Required if port forwarding is enabled.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.

* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server. Valid values: `disable`, `enable`.

* `type` - Type of address. Valid values: `ip`, `address`.

* `verify_cert` - Enable/disable certificate verification of the real server. Valid values: `disable`, `enable`.

* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `health_check_proto` - Health-Check-Proto. Valid values: `ping`, `http`.

* `seq` - Seq.

The `ssl_cipher_suites` block supports:

* `cipher` - Cipher suite name. Valid values: `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA`.

* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `id` - Id.

The `ssl_server_cipher_suites` block supports:

* `cipher` - Cipher suite name. Valid values: `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA`.

* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Vip can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_vip.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

