---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_setting"
description: |-
  VDOM wireless controller configuration.
---

# fmgdevice_wirelesscontroller_setting
VDOM wireless controller configuration.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `offending_ssid`: `fmgdevice_wirelesscontroller_setting_offendingssid`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_setting" "trname" {
  account_id               = "your own value"
  country                  = "PH"
  darrp_optimize           = 10
  darrp_optimize_schedules = ["your own value"]
  device_holdoff           = 10
  device_name              = var.device_name # not required if setting is at provider
  device_vdom              = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `account_id` - FortiCloud customer account ID.
* `country` - Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available. Valid values: `AL`, `DZ`, `AR`, `AM`, `AU`, `AT`, `AZ`, `BH`, `BD`, `BY`, `BE`, `BZ`, `BO`, `BA`, `BR`, `BN`, `BG`, `CA`, `CL`, `CN`, `CO`, `CR`, `HR`, `CY`, `CZ`, `DK`, `DO`, `EC`, `EG`, `SV`, `EE`, `FI`, `FR`, `GE`, `DE`, `GR`, `GT`, `HN`, `HK`, `HU`, `IS`, `IN`, `ID`, `IR`, `IE`, `IL`, `IT`, `JM`, `JP`, `JO`, `KZ`, `KE`, `KP`, `KR`, `KW`, `LV`, `LB`, `LI`, `LT`, `LU`, `MO`, `MK`, `MY`, `MT`, `MX`, `MC`, `MA`, `NP`, `NL`, `AN`, `NZ`, `NO`, `OM`, `PK`, `PA`, `PG`, `PE`, `PH`, `PL`, `PT`, `PR`, `QA`, `RO`, `RU`, `SA`, `CS`, `SG`, `SK`, `SI`, `ZA`, `ES`, `LK`, `SE`, `CH`, `SY`, `TW`, `TH`, `TT`, `TN`, `TR`, `AE`, `UA`, `GB`, `US`, `PS`, `UY`, `UZ`, `VE`, `VN`, `YE`, `ZW`, `NA`, `BS`, `VC`, `KH`, `MV`, `AF`, `NG`, `TZ`, `ZM`, `SN`, `CI`, `GH`, `SD`, `CM`, `MW`, `AO`, `GA`, `ML`, `BJ`, `MG`, `TD`, `BW`, `LY`, `RW`, `MZ`, `GM`, `LS`, `MU`, `CG`, `UG`, `BF`, `SL`, `SO`, `CD`, `NE`, `CF`, `SZ`, `TG`, `LR`, `MR`, `DJ`, `RE`, `RS`, `ME`, `IQ`, `MD`, `KY`, `BB`, `BM`, `TC`, `VI`, `PM`, `MF`, `GD`, `IM`, `FO`, `GI`, `GL`, `TM`, `MN`, `VU`, `FJ`, `LA`, `GU`, `WF`, `MH`, `BT`, `FM`, `PF`, `NI`, `PY`, `HT`, `GY`, `AW`, `KN`, `GF`, `AS`, `MP`, `PW`, `MM`, `LC`, `GP`, `ET`, `SR`, `ZB`, `CX`, `DM`, `MQ`, `YT`, `BL`, `--`.

* `darrp_optimize` - Time for running Distributed Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
* `darrp_optimize_schedules` - Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space.
* `device_holdoff` - Lower limit of creation time of device for identification in minutes (0 - 60, default = 5).
* `device_idle` - Upper limit of idle time of device for identification in minutes (0 - 14400, default = 1440).
* `device_weight` - Upper limit of confidence of device for identification (0 - 255, default = 1, 0 = disable).
* `duplicate_ssid` - Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM. Valid values: `disable`, `enable`.

* `fake_ssid_action` - Actions taken for detected fake SSID. Valid values: `log`, `suppress`.

* `fapc_compatibility` - Enable/disable FAP-C series compatibility. Valid values: `disable`, `enable`.

* `firmware_provision_on_authorization` - Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `disable`, `enable`.

* `offending_ssid` - Offending-Ssid. The structure of `offending_ssid` block is documented below.
* `phishing_ssid_detect` - Enable/disable phishing SSID detection. Valid values: `disable`, `enable`.

* `rolling_wtp_upgrade` - Enable/disable rolling WTP upgrade (default = disable). Valid values: `disable`, `enable`.

* `wfa_compatibility` - Enable/disable WFA compatibility. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `offending_ssid` block supports:

* `action` - Actions taken for detected offending SSID. Valid values: `log`, `suppress`.

* `id` - ID.
* `ssid_pattern` - Define offending SSID pattern (case insensitive). For example, word, word*, *word, wo*rd.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Setting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_setting.labelname WirelessControllerSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

