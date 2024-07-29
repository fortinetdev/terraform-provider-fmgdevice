---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_vdomexception"
description: |-
  Global configuration objects that can be configured independently across different ha peers for all VDOMs or for the defined VDOM scope.
---

# fmgdevice_system_vdomexception
Global configuration objects that can be configured independently across different ha peers for all VDOMs or for the defined VDOM scope.

## Example Usage

```hcl
resource "fmgdevice_system_vdomexception" "trname" {
  fosid       = 10
  object      = "log.syslogd4.override-setting"
  oid         = 10
  scope       = "all"
  vdom        = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fosid` - Index (1 - 4096).
* `object` - Name of the configuration object that can be configured independently for all VDOMs. Valid values: `log.fortianalyzer.setting`, `log.fortianalyzer.override-setting`, `vpn.ipsec.phase1-interface`, `vpn.ipsec.phase2-interface`, `router.bgp`, `router.route-map`, `router.prefix-list`, `firewall.ippool`, `log.fortianalyzer2.setting`, `log.fortianalyzer2.override-setting`, `log.fortianalyzer3.setting`, `log.fortianalyzer3.override-setting`, `log.fortianalyzer-cloud.setting`, `log.fortianalyzer-cloud.override-setting`, `system.central-management`, `system.csf`, `user.radius`, `system.virtual-wan-link`, `log.syslogd.setting`, `log.syslogd.override-setting`, `log.syslogd2.setting`, `log.syslogd2.override-setting`, `log.syslogd3.setting`, `log.syslogd3.override-setting`, `log.syslogd4.setting`, `log.syslogd4.override-setting`, `firewall.vip`, `firewall.vip6`, `firewall.vip46`, `firewall.vip64`, `firewall.ippool6`, `router.static`, `router.static6`, `system.interface`, `system.sdwan`, `system.saml`, `router.policy`, `router.policy6`, `system.gre-tunnel`, `system.cluster-sync`, `system.standalone-cluster`, `system.snmp.sysinfo`, `system.snmp.community`, `system.snmp.user`, `firewall.address`.

* `oid` - Object ID.
* `scope` - Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.

* `vdom` - Names of the VDOMs.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System VdomException can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_vdomexception.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

