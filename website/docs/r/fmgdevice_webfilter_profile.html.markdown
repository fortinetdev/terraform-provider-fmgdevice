---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_profile"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure Web filter profiles.
---

# fmgdevice_webfilter_profile
<i>This object will be purged after policy copy and install.</i> Configure Web filter profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `antiphish`: `fmgdevice_webfilter_profile_antiphish`
>- `ftgd_wf`: `fmgdevice_webfilter_profile_ftgdwf`
>- `override`: `fmgdevice_webfilter_profile_override`
>- `url_extraction`: `fmgdevice_webfilter_profile_urlextraction`
>- `web`: `fmgdevice_webfilter_profile_web`
>- `youtube_channel_filter`: `fmgdevice_webfilter_profile_youtubechannelfilter`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `antiphish` - Antiphish. The structure of `antiphish` block is documented below.
* `comment` - Optional comments.
* `extended_log` - Enable/disable extended logging for web filtering. Valid values: `disable`, `enable`.

* `feature_set` - Flow/proxy feature set. Valid values: `proxy`, `flow`.

* `ftgd_wf` - Ftgd-Wf. The structure of `ftgd_wf` block is documented below.
* `https_replacemsg` - Enable replacement messages for HTTPS. Valid values: `disable`, `enable`.

* `log_all_url` - Enable/disable logging all URLs visited. Valid values: `disable`, `enable`.

* `name` - Profile name.
* `options` - Options. Valid values: `block-invalid-url`, `jscript`, `js`, `vbs`, `unknown`, `wf-referer`, `intrinsic`, `wf-cookie`, `activexfilter`, `cookiefilter`, `javafilter`, `per-user-bal`.

* `override` - Override. The structure of `override` block is documented below.
* `ovrd_perm` - Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.

* `post_action` - Action taken for HTTP POST traffic. Valid values: `normal`, `block`.

* `replacemsg_group` - Replacement message group.
* `url_extraction` - Url-Extraction. The structure of `url_extraction` block is documented below.
* `web` - Web. The structure of `web` block is documented below.
* `web_antiphishing_log` - Enable/disable logging of AntiPhishing checks. Valid values: `disable`, `enable`.

* `web_content_log` - Enable/disable logging logging blocked web content. Valid values: `disable`, `enable`.

* `web_extended_all_action_log` - Enable/disable extended any filter action logging for web filtering. Valid values: `disable`, `enable`.

* `web_filter_activex_log` - Enable/disable logging ActiveX. Valid values: `disable`, `enable`.

* `web_filter_applet_log` - Enable/disable logging Java applets. Valid values: `disable`, `enable`.

* `web_filter_command_block_log` - Enable/disable logging blocked commands. Valid values: `disable`, `enable`.

* `web_filter_cookie_log` - Enable/disable logging cookie filtering. Valid values: `disable`, `enable`.

* `web_filter_cookie_removal_log` - Enable/disable logging blocked cookies. Valid values: `disable`, `enable`.

* `web_filter_js_log` - Enable/disable logging Java scripts. Valid values: `disable`, `enable`.

* `web_filter_jscript_log` - Enable/disable logging JScripts. Valid values: `disable`, `enable`.

* `web_filter_referer_log` - Enable/disable logging referrers. Valid values: `disable`, `enable`.

* `web_filter_unknown_log` - Enable/disable logging unknown scripts. Valid values: `disable`, `enable`.

* `web_filter_vbs_log` - Enable/disable logging VBS scripts. Valid values: `disable`, `enable`.

* `web_flow_log_encoding` - Log encoding in flow mode. Valid values: `utf-8`, `punycode`.

* `web_ftgd_err_log` - Enable/disable logging rating errors. Valid values: `disable`, `enable`.

* `web_ftgd_quota_usage` - Enable/disable logging daily quota usage. Valid values: `disable`, `enable`.

* `web_invalid_domain_log` - Enable/disable logging invalid domain names. Valid values: `disable`, `enable`.

* `web_url_log` - Enable/disable logging URL filtering. Valid values: `disable`, `enable`.

* `wisp` - Enable/disable web proxy WISP. Valid values: `disable`, `enable`.

* `wisp_algorithm` - WISP server selection algorithm. Valid values: `auto-learning`, `primary-secondary`, `round-robin`.

* `wisp_servers` - WISP servers.
* `youtube_channel_filter` - Youtube-Channel-Filter. The structure of `youtube_channel_filter` block is documented below.
* `youtube_channel_status` - YouTube channel filter status. Valid values: `disable`, `blacklist`, `whitelist`.

* `ia_categorization` - Ia-Categorization. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `antiphish` block supports:

* `authentication` - Authentication methods. Valid values: `domain-controller`, `ldap`.

* `check_basic_auth` - Enable/disable checking of HTTP Basic Auth field for known credentials. Valid values: `disable`, `enable`.

* `check_uri` - Enable/disable checking of GET URI parameters for known credentials. Valid values: `disable`, `enable`.

* `check_username_only` - Enable/disable username only matching of credentials. Action will be taken for valid usernames regardless of password validity. Valid values: `disable`, `enable`.

* `custom_patterns` - Custom-Patterns. The structure of `custom_patterns` block is documented below.
* `default_action` - Action to be taken when there is no matching rule. Valid values: `log`, `block`, `exempt`.

* `domain_controller` - Domain for which to verify received credentials against.
* `inspection_entries` - Inspection-Entries. The structure of `inspection_entries` block is documented below.
* `ldap` - LDAP server for which to verify received credentials against.
* `max_body_len` - Maximum size of a POST body to check for credentials.
* `status` - Toggle AntiPhishing functionality. Valid values: `disable`, `enable`.


The `custom_patterns` block supports:

* `category` - Category that the pattern matches. Valid values: `username`, `password`.

* `pattern` - Target pattern.
* `type` - Pattern will be treated either as a regex pattern or literal string. Valid values: `regex`, `literal`.


The `inspection_entries` block supports:

* `action` - Action to be taken upon an AntiPhishing match. Valid values: `log`, `block`, `exempt`.

* `fortiguard_category` - FortiGuard category to match.
* `name` - Inspection target name.

The `ftgd_wf` block supports:

* `exempt_quota` - Do not stop quota for these categories.
* `filters` - Filters. The structure of `filters` block is documented below.
* `max_quota_timeout` - Maximum FortiGuard quota used by single page view in seconds (excludes streams).
* `options` - Options for FortiGuard Web Filter. Valid values: `error-allow`, `rate-server-ip`, `connect-request-bypass`, `ftgd-disable`.

* `ovrd` - Allow web filter profile overrides.
* `quota` - Quota. The structure of `quota` block is documented below.
* `rate_crl_urls` - Enable/disable rating CRL by URL. Valid values: `disable`, `enable`.

* `rate_css_urls` - Enable/disable rating CSS by URL. Valid values: `disable`, `enable`.

* `rate_image_urls` - Rate-Image-Urls. Valid values: `disable`, `enable`.

* `rate_javascript_urls` - Enable/disable rating JavaScript by URL. Valid values: `disable`, `enable`.

* `risk` - Risk. The structure of `risk` block is documented below.

The `filters` block supports:

* `action` - Action to take for matches. Valid values: `block`, `monitor`, `warning`, `authenticate`.

* `auth_usr_grp` - Groups with permission to authenticate.
* `category` - Categories and groups the filter examines.
* `id` - ID number.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `override_replacemsg` - Override replacement message.
* `warn_duration` - Duration of warnings.
* `warning_duration_type` - Re-display warning after closing browser or after a timeout. Valid values: `session`, `timeout`.

* `warning_prompt` - Warning prompts in each category or each domain. Valid values: `per-domain`, `per-category`.


The `quota` block supports:

* `category` - FortiGuard categories to apply quota to (category action must be set to monitor).
* `duration` - Duration of quota.
* `id` - ID number.
* `override_replacemsg` - Override replacement message.
* `type` - Quota type. Valid values: `time`, `traffic`.

* `unit` - Traffic quota unit of measurement. Valid values: `B`, `KB`, `MB`, `GB`.

* `value` - Traffic quota value.

The `risk` block supports:

* `action` - Action to take for matches. Valid values: `block`, `monitor`.

* `id` - ID number.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `risk_level` - Risk level to be examined.

The `override` block supports:

* `ovrd_cookie` - Allow/deny browser-based (cookie) overrides. Valid values: `deny`, `allow`.

* `ovrd_dur` - Override duration.
* `ovrd_dur_mode` - Override duration mode. Valid values: `constant`, `ask`.

* `ovrd_scope` - Override scope. Valid values: `user`, `user-group`, `ip`, `ask`, `browser`.

* `ovrd_user_group` - User groups with permission to use the override.
* `profile` - Web filter profile with permission to create overrides.
* `profile_attribute` - Profile attribute to retrieve from the RADIUS server. Valid values: `User-Name`, `NAS-IP-Address`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Filter-Id`, `Login-IP-Host`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `Class`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Zone`, `Acct-Session-Id`, `Acct-Multi-Session-Id`.

* `profile_type` - Override profile type. Valid values: `list`, `radius`.


The `url_extraction` block supports:

* `redirect_header` - HTTP header name to use for client redirect on blocked requests
* `redirect_no_content` - Enable / Disable empty message-body entity in HTTP response Valid values: `disable`, `enable`.

* `redirect_url` - HTTP header value to use for client redirect on blocked requests
* `server_fqdn` - URL extraction server FQDN (fully qualified domain name)
* `status` - Enable URL Extraction Valid values: `disable`, `enable`.


The `web` block supports:

* `blacklist` - Enable/disable automatic addition of URLs detected by FortiSandbox to blacklist. Valid values: `disable`, `enable`.

* `allowlist` - FortiGuard allowlist settings. Valid values: `exempt-av`, `exempt-webcontent`, `exempt-activex-java-cookie`, `exempt-dlp`, `exempt-rangeblock`, `extended-log-others`.

* `blocklist` - Enable/disable automatic addition of URLs detected by FortiSandbox to blocklist. Valid values: `disable`, `enable`.

* `bword_table` - Banned word table ID.
* `bword_threshold` - Banned word score threshold.
* `content_header_list` - Content header list.
* `keyword_match` - Search keywords to log when match is found.
* `log_search` - Enable/disable logging all search phrases. Valid values: `disable`, `enable`.

* `safe_search` - Safe search type. Valid values: `url`, `header`.

* `urlfilter_table` - URL filter table ID.
* `whitelist` - FortiGuard whitelist settings. Valid values: `exempt-av`, `exempt-webcontent`, `exempt-activex-java-cookie`, `exempt-dlp`, `exempt-rangeblock`, `extended-log-others`.

* `vimeo_restrict` - Set Vimeo-restrict ("7" = don't show mature content, "134" = don't show unrated and mature content). A value of cookie "content_rating".
* `youtube_restrict` - YouTube EDU filter level. Valid values: `strict`, `none`, `moderate`.

* `qwant_restrict` - Qwant-Restrict. Valid values: `strict`, `none`, `moderate`.


The `youtube_channel_filter` block supports:

* `channel_id` - YouTube channel ID to be filtered.
* `comment` - Comment.
* `id` - ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Webfilter Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

