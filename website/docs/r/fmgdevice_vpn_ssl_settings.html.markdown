---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ssl_settings"
description: |-
  Configure SSL-VPN.
---

# fmgdevice_vpn_ssl_settings
Configure SSL-VPN.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `authentication_rule`: `fmgdevice_vpn_ssl_settings_authenticationrule`



## Example Usage

```hcl
resource "fmgdevice_vpn_ssl_settings" "trname" {
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
  authentication_rule {
    id               = 2
    source_interface = ["any"]
    source_address   = ["all"]
    users            = ["guest"]
    portal           = ["web-access"]
  }
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `algorithm` - Force the SSL-VPN security level. High allows only high. Medium allows medium and high. Low allows any. Valid values: `default`, `high`, `low`, `medium`.

* `auth_session_check_source_ip` - Enable/disable checking of source IP for authentication session. Valid values: `disable`, `enable`.

* `auth_timeout` - SSL-VPN authentication timeout (1 - 259200 sec (3 days), 0 for no timeout).
* `authentication_rule` - Authentication-Rule. The structure of `authentication_rule` block is documented below.
* `auto_tunnel_static_route` - Enable/disable to auto-create static routes for the SSL-VPN tunnel IP addresses. Valid values: `disable`, `enable`.

* `banned_cipher` - Select one or more cipher technologies that cannot be used in SSL-VPN negotiations. Only applies to TLS 1.2 and below. Valid values: `RSA`, `DH`, `DHE`, `ECDH`, `ECDHE`, `DSS`, `ECDSA`, `AES`, `AESGCM`, `CAMELLIA`, `3DES`, `SHA1`, `SHA256`, `SHA384`, `STATIC`, `CHACHA20`, `ARIA`, `AESCCM`.

* `browser_language_detection` - Enable/disable overriding the configured system language based on the preferred language of the browser. Valid values: `disable`, `enable`.

* `check_referer` - Enable/disable verification of referer field in HTTP request header. Valid values: `disable`, `enable`.

* `ciphersuite` - Select one or more TLS 1.3 ciphersuites to enable. Does not affect ciphers in TLS 1.2 and below. At least one must be enabled. To disable all, set ssl-max-proto-ver to tls1-2 or below. Valid values: `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-AES-128-CCM-SHA256`, `TLS-AES-128-CCM-8-SHA256`.

* `client_sigalgs` - Set signature algorithms related to client authentication. Affects TLS version &lt;= 1.2 only. Valid values: `no-rsa-pss`, `all`.

* `default_portal` - Default SSL-VPN portal.
* `deflate_compression_level` - Compression level (0~9).
* `deflate_min_data_size` - Minimum amount of data that triggers compression (200 - 65535 bytes).
* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_suffix` - DNS suffix used for SSL-VPN clients.
* `dtls_heartbeat_fail_count` - Number of missing heartbeats before the connection is considered dropped.
* `dtls_heartbeat_idle_timeout` - Idle timeout before DTLS heartbeat is sent.
* `dtls_heartbeat_interval` - Interval between DTLS heartbeat.
* `dtls_hello_timeout` - SSLVPN maximum DTLS hello timeout (10 - 60 sec, default = 10).
* `dtls_max_proto_ver` - DTLS maximum protocol version. Valid values: `dtls1-0`, `dtls1-2`.

* `dtls_min_proto_ver` - DTLS minimum protocol version. Valid values: `dtls1-0`, `dtls1-2`.

* `dtls_tunnel` - Enable/disable DTLS to prevent eavesdropping, tampering, or message forgery. Valid values: `disable`, `enable`.

* `dual_stack_mode` - Tunnel mode: enable parallel IPv4 and IPv6 tunnel. Web mode: support IPv4 and IPv6 bookmarks in the portal. Valid values: `disable`, `enable`.

* `encode_2f_sequence` - Encode 2F sequence to forward slash in URLs. Valid values: `disable`, `enable`.

* `encrypt_and_store_password` - Encrypt and store user passwords for SSL-VPN web sessions. Valid values: `disable`, `enable`.

* `force_two_factor_auth` - Enable/disable only PKI users with two-factor authentication for SSL-VPNs. Valid values: `disable`, `enable`.

* `header_x_forwarded_for` - Forward the same, add, or remove HTTP header. Valid values: `pass`, `add`, `remove`.

* `hsts_include_subdomains` - Add HSTS includeSubDomains response header. Valid values: `disable`, `enable`.

* `http_compression` - Enable/disable to allow HTTP compression over SSL-VPN tunnels. Valid values: `disable`, `enable`.

* `http_only_cookie` - Enable/disable SSL-VPN support for HttpOnly cookies. Valid values: `disable`, `enable`.

* `http_request_body_timeout` - SSL-VPN session is disconnected if an HTTP request body is not received within this time (1 - 60 sec, default = 20).
* `http_request_header_timeout` - SSL-VPN session is disconnected if an HTTP request header is not received within this time (1 - 60 sec, default = 20).
* `https_redirect` - Enable/disable redirect of port 80 to SSL-VPN port. Valid values: `disable`, `enable`.

* `idle_timeout` - SSL-VPN disconnects if idle for specified time in seconds.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `ipv6_wins_server1` - IPv6 WINS server 1.
* `ipv6_wins_server2` - IPv6 WINS server 2.
* `login_attempt_limit` - SSL-VPN maximum login attempt times before block (0 - 10, default = 2, 0 = no limit).
* `login_block_time` - Time for which a user is blocked from logging in after too many failed login attempts (0 - 86400 sec, default = 60).
* `login_timeout` - SSLVPN maximum login timeout (10 - 180 sec, default = 30).
* `port` - SSL-VPN access port (1 - 65535).
* `port_precedence` - Enable/disable, Enable means that if SSL-VPN connections are allowed on an interface admin GUI connections are blocked on that interface. Valid values: `disable`, `enable`.

* `reqclientcert` - Enable/disable to require client certificates for all SSL-VPN users. Valid values: `disable`, `enable`.

* `saml_redirect_port` - SAML local redirect port in the machine running FortiClient (0 - 65535). 0 is to disable redirection on FGT side.
* `server_hostname` - Server hostname for HTTPS. When set, will be used for SSL VPN web proxy host header for any redirection.
* `servercert` - Name of the server certificate to be used for SSL-VPNs.
* `source_address` - Source address of incoming traffic.
* `source_address_negate` - Enable/disable negated source address match. Valid values: `disable`, `enable`.

* `source_address6` - IPv6 source address of incoming traffic.
* `source_address6_negate` - Enable/disable negated source IPv6 address match. Valid values: `disable`, `enable`.

* `source_interface` - SSL-VPN source interface of incoming traffic.
* `ssl_client_renegotiation` - Enable/disable to allow client renegotiation by the server if the tunnel goes down. Valid values: `disable`, `enable`.

* `ssl_insert_empty_fragment` - Enable/disable insertion of empty fragment. Valid values: `disable`, `enable`.

* `ssl_max_proto_ver` - SSL maximum protocol version. Valid values: `tls1-0`, `tls1-1`, `tls1-2`, `tls1-3`.

* `ssl_min_proto_ver` - SSL minimum protocol version. Valid values: `tls1-0`, `tls1-1`, `tls1-2`, `tls1-3`.

* `status` - Enable/disable SSL-VPN. Valid values: `disable`, `enable`.

* `transform_backward_slashes` - Transform backward slashes to forward slashes in URLs. Valid values: `disable`, `enable`.

* `tunnel_addr_assigned_method` - Method used for assigning address for tunnel. Valid values: `first-available`, `round-robin`.

* `tunnel_connect_without_reauth` - Enable/disable tunnel connection without re-authorization if previous connection dropped. Valid values: `disable`, `enable`.

* `tunnel_ip_pools` - Names of the IPv4 IP Pool firewall objects that define the IP addresses reserved for remote clients.
* `tunnel_ipv6_pools` - Names of the IPv6 IP Pool firewall objects that define the IP addresses reserved for remote clients.
* `tunnel_user_session_timeout` - Number of seconds after which user sessions are cleaned up after tunnel connection is dropped (1 - 86400, default = 30).
* `unsafe_legacy_renegotiation` - Enable/disable unsafe legacy re-negotiation. Valid values: `disable`, `enable`.

* `url_obscuration` - Enable/disable to obscure the host name of the URL of the web browser display. Valid values: `disable`, `enable`.

* `user_peer` - Name of user peer.
* `web_mode_snat` - Enable/disable use of IP pools defined in firewall policy while using web-mode. Valid values: `disable`, `enable`.

* `wins_server1` - WINS server 1.
* `wins_server2` - WINS server 2.
* `x_content_type_options` - Add HTTP X-Content-Type-Options header. Valid values: `disable`, `enable`.

* `ztna_trusted_client` - Enable/disable verification of device certificate for SSLVPN ZTNA session. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `authentication_rule` block supports:

* `auth` - SSL-VPN authentication method restriction. Valid values: `any`, `local`, `radius`, `ldap`, `tacacs+`, `peer`.

* `cipher` - SSL-VPN cipher strength. Valid values: `any`, `high`, `medium`.

* `client_cert` - Enable/disable SSL-VPN client certificate restrictive. Valid values: `disable`, `enable`.

* `groups` - User groups.
* `id` - ID (0 - 4294967295).
* `portal` - SSL-VPN portal.
* `realm` - SSL-VPN realm.
* `source_address` - Source address of incoming traffic.
* `source_address_negate` - Enable/disable negated source address match. Valid values: `disable`, `enable`.

* `source_address6` - IPv6 source address of incoming traffic.
* `source_address6_negate` - Enable/disable negated source IPv6 address match. Valid values: `disable`, `enable`.

* `source_interface` - SSL-VPN source interface of incoming traffic.
* `user_peer` - Name of user peer.
* `users` - User name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn SslSettings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ssl_settings.labelname VpnSslSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

