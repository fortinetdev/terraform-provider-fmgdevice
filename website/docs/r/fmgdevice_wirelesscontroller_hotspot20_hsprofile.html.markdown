---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_hsprofile"
description: |-
  Configure hotspot profile.
---

# fmgdevice_wirelesscontroller_hotspot20_hsprofile
Configure hotspot profile.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_hsprofile" "trname" {
  n3gpp_plmn              = ["your own value"]
  access_network_asra     = "disable"
  access_network_esr      = "enable"
  access_network_internet = "disable"
  access_network_type     = "chargeable-public-network"
  name                    = "your own value"
  device_name             = var.device_name # not required if setting is at provider
  device_vdom             = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `n3gpp_plmn` - 3GPP PLMN name.
* `access_network_asra` - Enable/disable additional step required for access (ASRA). Valid values: `disable`, `enable`.

* `access_network_esr` - Enable/disable emergency services reachable (ESR). Valid values: `disable`, `enable`.

* `access_network_internet` - Enable/disable connectivity to the Internet. Valid values: `disable`, `enable`.

* `access_network_type` - Access network type. Valid values: `private-network`, `private-network-with-guest-access`, `chargeable-public-network`, `free-public-network`, `personal-device-network`, `emergency-services-only-network`, `test-or-experimental`, `wildcard`.

* `access_network_uesa` - Enable/disable unauthenticated emergency service accessible (UESA). Valid values: `disable`, `enable`.

* `advice_of_charge` - Advice of charge.
* `anqp_domain_id` - ANQP Domain ID (0-65535).
* `bss_transition` - Enable/disable basic service set (BSS) transition Support. Valid values: `disable`, `enable`.

* `conn_cap` - Connection capability name.
* `deauth_request_timeout` - Deauthentication request timeout (in seconds).
* `dgaf` - Enable/disable downstream group-addressed forwarding (DGAF). Valid values: `disable`, `enable`.

* `domain_name` - Domain name.
* `gas_comeback_delay` - GAS comeback delay (0 or 100 - 10000 milliseconds, default = 500).
* `gas_fragmentation_limit` - GAS fragmentation limit (512 - 4096, default = 1024).
* `hessid` - Homogeneous extended service set identifier (HESSID).
* `ip_addr_type` - IP address type name.
* `l2tif` - Enable/disable Layer 2 traffic inspection and filtering. Valid values: `disable`, `enable`.

* `nai_realm` - NAI realm list name.
* `name` - Hotspot profile name.
* `network_auth` - Network authentication name.
* `oper_friendly_name` - Operator friendly name.
* `oper_icon` - Operator icon.
* `osu_provider` - Manually selected list of OSU provider(s).
* `osu_provider_nai` - OSU Provider NAI.
* `osu_ssid` - Online sign up (OSU) SSID.
* `pame_bi` - Enable/disable Pre-Association Message Exchange BSSID Independent (PAME-BI). Valid values: `disable`, `enable`.

* `proxy_arp` - Enable/disable Proxy ARP. Valid values: `disable`, `enable`.

* `qos_map` - QoS MAP set ID.
* `release` - Hotspot 2.0 Release number (1, 2, 3, default = 2).
* `roaming_consortium` - Roaming consortium list name.
* `terms_and_conditions` - Terms and conditions.
* `venue_group` - Venue group. Valid values: `unspecified`, `assembly`, `business`, `educational`, `factory`, `institutional`, `mercantile`, `residential`, `storage`, `utility`, `vehicular`, `outdoor`.

* `venue_name` - Venue name.
* `venue_type` - Venue type. Valid values: `unspecified`, `arena`, `stadium`, `passenger-terminal`, `amphitheater`, `amusement-park`, `place-of-worship`, `convention-center`, `library`, `museum`, `restaurant`, `theater`, `bar`, `coffee-shop`, `zoo-or-aquarium`, `emergency-center`, `doctor-office`, `bank`, `fire-station`, `police-station`, `post-office`, `professional-office`, `research-facility`, `attorney-office`, `primary-school`, `secondary-school`, `university-or-college`, `factory`, `hospital`, `long-term-care-facility`, `rehab-center`, `group-home`, `prison-or-jail`, `retail-store`, `grocery-market`, `auto-service-station`, `shopping-mall`, `gas-station`, `private`, `hotel-or-motel`, `dormitory`, `boarding-house`, `automobile`, `airplane`, `bus`, `ferry`, `ship-or-boat`, `train`, `motor-bike`, `muni-mesh-network`, `city-park`, `rest-area`, `traffic-control`, `bus-stop`, `kiosk`.

* `venue_url` - Venue name.
* `wan_metrics` - WAN metric name.
* `wnm_sleep_mode` - Enable/disable wireless network management (WNM) sleep mode. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20HsProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_hsprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

