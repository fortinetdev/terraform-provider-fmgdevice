---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_sniffer"
description: |-
  Configure sniffer.
---

# fmgdevice_firewall_sniffer
Configure sniffer.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `anomaly`: `fmgdevice_firewall_sniffer_anomaly`



## Example Usage

```hcl
resource "fmgdevice_firewall_sniffer" "trname" {
  anomaly {
    action                   = "proxy"
    log                      = "enable"
    name                     = "your own value"
    quarantine               = "interface"
    quarantine_expiry        = "your own value"
    quarantine_log           = "disable"
    status                   = "enable"
    synproxy_tcp_mss         = "512"
    synproxy_tcp_sack        = "enable"
    synproxy_tcp_timestamp   = "enable"
    synproxy_tcp_window      = "16384"
    synproxy_tcp_windowscale = "6"
    synproxy_tos             = "18"
    synproxy_ttl             = "32"
    threshold                = 10
    thresholddefault         = 10
  }

  application_list        = ["your own value"]
  application_list_status = "enable"
  av_profile              = ["your own value"]
  av_profile_status       = "enable"
  fosid                   = 10
  device_name             = var.device_name # not required if setting is at provider
  device_vdom             = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `anomaly` - Anomaly. The structure of `anomaly` block is documented below.
* `application_list` - Name of an existing application list.
* `application_list_status` - Enable/disable application control profile. Valid values: `disable`, `enable`.

* `av_profile` - Name of an existing antivirus profile.
* `av_profile_status` - Enable/disable antivirus profile. Valid values: `disable`, `enable`.

* `dlp_sensor` - Name of an existing DLP sensor.
* `dlp_sensor_status` - Enable/disable DLP sensor. Valid values: `disable`, `enable`.

* `dlp_profile` - Name of an existing DLP profile.
* `dlp_profile_status` - Enable/disable DLP profile. Valid values: `disable`, `enable`.

* `dsri` - Enable/disable DSRI. Valid values: `disable`, `enable`.

* `emailfilter_profile` - Name of an existing email filter profile.
* `emailfilter_profile_status` - Enable/disable emailfilter. Valid values: `disable`, `enable`.

* `file_filter_profile` - Name of an existing file-filter profile.
* `file_filter_profile_status` - Enable/disable file filter. Valid values: `disable`, `enable`.

* `host` - Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
* `fosid` - Sniffer ID (0 - 9999).
* `interface` - Interface name that traffic sniffing will take place on.
* `ip_threatfeed` - Name of an existing IP threat feed.
* `ip_threatfeed_status` - Enable/disable IP threat feed. Valid values: `disable`, `enable`.

* `ips_dos_status` - Enable/disable IPS DoS anomaly detection. Valid values: `disable`, `enable`.

* `ips_sensor` - Name of an existing IPS sensor.
* `ips_sensor_status` - Enable/disable IPS sensor. Valid values: `disable`, `enable`.

* `ipv6` - Enable/disable sniffing IPv6 packets. Valid values: `disable`, `enable`.

* `logtraffic` - Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `disable`, `all`, `utm`.

* `max_packet_count` - Maximum packet count (1 - 1000000, default = 4000).
* `non_ip` - Enable/disable sniffing non-IP packets. Valid values: `disable`, `enable`.

* `port` - Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `status` - Enable/disable the active status of the sniffer. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `vlan` - List of VLANs to sniff.
* `webfilter_profile` - Name of an existing web filter profile.
* `webfilter_profile_status` - Enable/disable web filter profile. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `anomaly` block supports:

* `action` - Action taken when the threshold is reached. Valid values: `pass`, `block`, `proxy`.

* `log` - Enable/disable anomaly logging. Valid values: `disable`, `enable`.

* `name` - Anomaly name.
* `quarantine` - Quarantine method. Valid values: `none`, `attacker`, `both`, `interface`.

* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging. Valid values: `disable`, `enable`.

* `status` - Enable/disable this anomaly. Valid values: `disable`, `enable`.

* `synproxy_tcp_mss` - Determine TCP maximum segment size (MSS) value for packets replied by syn proxy module. Valid values: `0`, `256`, `512`, `1024`, `1300`, `1360`, `1460`, `1500`.

* `synproxy_tcp_sack` - enable/disable TCP selective acknowledage (SACK) for packets replied by syn proxy module. Valid values: `disable`, `enable`.

* `synproxy_tcp_timestamp` - enable/disable TCP timestamp option for packets replied by syn proxy module. Valid values: `disable`, `enable`.

* `synproxy_tcp_window` - Determine TCP Window size for packets replied by syn proxy module. Valid values: `4096`, `8192`, `16384`, `32768`.

* `synproxy_tcp_windowscale` - Determine TCP window scale option value for packets replied by syn proxy module. Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `13`, `14`.

* `synproxy_tos` - Determine TCP differentiated services code point value (type of service). Valid values: `0`, `10`, `12`, `14`, `18`, `20`, `22`, `26`, `28`, `30`, `34`, `36`, `38`, `40`, `46`, `255`.

* `synproxy_ttl` - Determine Time to live (TTL) value for packets replied by syn proxy module. Valid values: `32`, `64`, `128`, `255`.

* `threshold` - Anomaly threshold. Number of detected instances (packets per second or concurrent session number) that triggers the anomaly action.
* `thresholddefault` - Threshold(Default).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall Sniffer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_sniffer.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

