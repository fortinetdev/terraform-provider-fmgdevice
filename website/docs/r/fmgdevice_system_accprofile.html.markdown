---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_accprofile"
description: |-
  Configure access profiles for system administrators.
---

# fmgdevice_system_accprofile
Configure access profiles for system administrators.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `fwgrp_permission`: `fmgdevice_system_accprofile_fwgrppermission`
>- `loggrp_permission`: `fmgdevice_system_accprofile_loggrppermission`
>- `netgrp_permission`: `fmgdevice_system_accprofile_netgrppermission`
>- `sysgrp_permission`: `fmgdevice_system_accprofile_sysgrppermission`
>- `utmgrp_permission`: `fmgdevice_system_accprofile_utmgrppermission`



## Example Usage

```hcl
resource "fmgdevice_system_accprofile" "trname" {
  admintimeout          = 10
  admintimeout_override = "enable"
  authgrp               = "none"
  cli_config            = "enable"
  cli_diagnose          = "disable"
  name                  = "your own value"
  device_name           = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `admintimeout` - Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
* `admintimeout_override` - Enable/disable overriding the global administrator idle timeout. Valid values: `disable`, `enable`.

* `authgrp` - Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.

* `cli_config` - Enable/disable permission to run config commands. Valid values: `disable`, `enable`.

* `cli_diagnose` - Enable/disable permission to run diagnostic commands. Valid values: `disable`, `enable`.

* `cli_exec` - Enable/disable permission to run execute commands. Valid values: `disable`, `enable`.

* `cli_get` - Enable/disable permission to run get commands. Valid values: `disable`, `enable`.

* `cli_show` - Enable/disable permission to run show commands. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `ftviewgrp` - FortiView. Valid values: `none`, `read`, `read-write`.

* `fwgrp` - Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.

* `fwgrp_permission` - Fwgrp-Permission. The structure of `fwgrp_permission` block is documented below.
* `loggrp` - Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.

* `loggrp_permission` - Loggrp-Permission. The structure of `loggrp_permission` block is documented below.
* `name` - Profile name.
* `netgrp` - Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.

* `netgrp_permission` - Netgrp-Permission. The structure of `netgrp_permission` block is documented below.
* `scope` - Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.

* `secfabgrp` - Security Fabric. Valid values: `none`, `read`, `read-write`.

* `sysgrp` - System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.

* `sysgrp_permission` - Sysgrp-Permission. The structure of `sysgrp_permission` block is documented below.
* `system_diagnostics` - System-Diagnostics. Valid values: `disable`, `enable`.

* `system_execute_ssh` - Enable/disable permission to execute SSH commands. Valid values: `disable`, `enable`.

* `system_execute_telnet` - Enable/disable permission to execute TELNET commands. Valid values: `disable`, `enable`.

* `utmgrp` - Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.

* `utmgrp_permission` - Utmgrp-Permission. The structure of `utmgrp_permission` block is documented below.
* `vpngrp` - Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.

* `wanoptgrp` - Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.

* `wifi` - Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.


The `fwgrp_permission` block supports:

* `address` - Address Configuration. Valid values: `none`, `read`, `read-write`.

* `others` - Other Firewall Configuration. Valid values: `none`, `read`, `read-write`.

* `policy` - Policy Configuration. Valid values: `none`, `read`, `read-write`.

* `schedule` - Schedule Configuration. Valid values: `none`, `read`, `read-write`.

* `service` - Service Configuration. Valid values: `none`, `read`, `read-write`.


The `loggrp_permission` block supports:

* `config` - Log & Report configuration. Valid values: `none`, `read`, `read-write`.

* `data_access` - Log & Report Data Access. Valid values: `none`, `read`, `read-write`.

* `report_access` - Log & Report Report Access. Valid values: `none`, `read`, `read-write`.

* `threat_weight` - Log & Report Threat Weight. Valid values: `none`, `read`, `read-write`.


The `netgrp_permission` block supports:

* `cfg` - Network Configuration. Valid values: `none`, `read`, `read-write`.

* `packet_capture` - Packet Capture Configuration. Valid values: `none`, `read`, `read-write`.

* `route_cfg` - Router Configuration. Valid values: `none`, `read`, `read-write`.


The `sysgrp_permission` block supports:

* `admin` - Administrator Users. Valid values: `none`, `read`, `read-write`.

* `cfg` - System Configuration. Valid values: `none`, `read`, `read-write`.

* `mnt` - Maintenance. Valid values: `none`, `read`, `read-write`.

* `upd` - FortiGuard Updates. Valid values: `none`, `read`, `read-write`.


The `utmgrp_permission` block supports:

* `antivirus` - Antivirus profiles and settings. Valid values: `none`, `read`, `read-write`.

* `application_control` - Application Control profiles and settings. Valid values: `none`, `read`, `read-write`.

* `casb` - Inline CASB filter profile and settings Valid values: `none`, `read`, `read-write`.

* `data_leak_prevention` - DLP profiles and settings. Valid values: `none`, `read`, `read-write`.

* `data_loss_prevention` - DLP profiles and settings. Valid values: `none`, `read`, `read-write`.

* `dlp` - DLP profiles and settings. Valid values: `none`, `read`, `read-write`.

* `dnsfilter` - DNS Filter profiles and settings. Valid values: `none`, `read`, `read-write`.

* `emailfilter` - Email Filter and settings. Valid values: `none`, `read`, `read-write`.

* `endpoint_control` - FortiClient Profiles. Valid values: `none`, `read`, `read-write`.

* `file_filter` - File-filter profiles and settings. Valid values: `none`, `read`, `read-write`.

* `icap` - ICAP profiles and settings. Valid values: `none`, `read`, `read-write`.

* `ips` - IPS profiles and settings. Valid values: `none`, `read`, `read-write`.

* `mmsgtp` - UTM permission. Valid values: `none`, `read`, `read-write`.

* `videofilter` - Video filter profiles and settings. Valid values: `none`, `read`, `read-write`.

* `virtual_patch` - Virtual patch profiles and settings. Valid values: `none`, `read`, `read-write`.

* `voip` - VoIP profiles and settings. Valid values: `none`, `read`, `read-write`.

* `waf` - Web Application Firewall profiles and settings. Valid values: `none`, `read`, `read-write`.

* `webfilter` - Web Filter profiles and settings. Valid values: `none`, `read`, `read-write`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Accprofile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_accprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

