---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_global"
description: |-
  Configure Web proxy global settings.
---

# fmgdevice_webproxy_global
Configure Web proxy global settings.

## Example Usage

```hcl
resource "fmgdevice_webproxy_global" "trname" {
  always_learn_client_ip          = "enable"
  fast_policy_match               = "disable"
  forward_proxy_auth              = "disable"
  forward_server_affinity_timeout = 10
  ldap_user_cache                 = "disable"
  device_name                     = var.device_name # not required if setting is at provider
  device_vdom                     = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `always_learn_client_ip` - Enable/disable learning the client's IP address from headers for every request. Valid values: `disable`, `enable`.

* `fast_policy_match` - Enable/disable fast matching algorithm for explicit and transparent proxy policy. Valid values: `disable`, `enable`.

* `forward_proxy_auth` - Enable/disable forwarding proxy authentication headers. Valid values: `disable`, `enable`.

* `forward_server_affinity_timeout` - Period of time before the source IP's traffic is no longer assigned to the forwarding server (6 - 60 min, default = 30).
* `ldap_user_cache` - Enable/disable LDAP user cache for explicit and transparent proxy user. Valid values: `disable`, `enable`.

* `learn_client_ip` - Enable/disable learning the client's IP address from headers. Valid values: `disable`, `enable`.

* `learn_client_ip_from_header` - Learn client IP address from the specified headers. Valid values: `true-client-ip`, `x-real-ip`, `x-forwarded-for`.

* `learn_client_ip_srcaddr` - Source address name (srcaddr or srcaddr6 must be set).
* `learn_client_ip_srcaddr6` - IPv6 Source address name (srcaddr or srcaddr6 must be set).
* `log_app_id` - Enable/disable always log application type in traffic log. Valid values: `disable`, `enable`.

* `log_forward_server` - Enable/disable forward server name logging in forward traffic log. Valid values: `disable`, `enable`.

* `log_policy_pending` - Enable/disable logging sessions that are pending on policy matching. Valid values: `disable`, `enable`.

* `max_message_length` - Maximum length of HTTP message, not including body (16 - 256 Kbytes, default = 32).
* `max_request_length` - Maximum length of HTTP request line (2 - 64 Kbytes, default = 8).
* `max_waf_body_cache_length` - Maximum length of HTTP messages processed by Web Application Firewall (WAF) (10 - 1024 Kbytes, default = 32).
* `policy_category_deep_inspect` - Enable/disable deep inspection for application level category policy matching. Valid values: `disable`, `enable`.

* `proxy_fqdn` - Fully Qualified Domain Name (FQDN) that clients connect to (default = default.fqdn) to connect to the explicit web proxy.
* `proxy_transparent_cert_inspection` - Enable/disable transparent proxy certificate inspection. Valid values: `disable`, `enable`.

* `request_obs_fold` - Action when HTTP/1.x request header contains obs-fold. Valid values: `block`, `replace-with-sp`, `keep`.

* `src_affinity_exempt_addr` - IPv4 source addresses to exempt proxy affinity.
* `src_affinity_exempt_addr6` - IPv6 source addresses to exempt proxy affinity.
* `ssl_ca_cert` - SSL CA certificate for SSL interception.
* `ssl_cert` - SSL certificate for SSL interception.
* `strict_web_check` - Enable/disable strict web checking to block web sites that send incorrect headers that don't conform to HTTP 1.1. Valid values: `disable`, `enable`.

* `tunnel_non_http` - Enable/disable allowing non-HTTP traffic. Allowed non-HTTP traffic is tunneled. Valid values: `disable`, `enable`.

* `unknown_http_version` - Action to take when an unknown version of HTTP is encountered: reject, allow (tunnel), or proceed with best-effort. Valid values: `best-effort`, `reject`, `tunnel`.

* `webproxy_profile` - Name of the web proxy profile to apply when explicit proxy traffic is allowed by default and traffic is accepted that does not match an explicit proxy policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WebProxy Global can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_global.labelname WebProxyGlobal
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

