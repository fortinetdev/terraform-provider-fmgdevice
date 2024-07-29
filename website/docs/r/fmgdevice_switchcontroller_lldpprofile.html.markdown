---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_lldpprofile"
description: |-
  Configure FortiSwitch LLDP profiles.
---

# fmgdevice_switchcontroller_lldpprofile
Configure FortiSwitch LLDP profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `custom_tlvs`: `fmgdevice_switchcontroller_lldpprofile_customtlvs`
>- `med_location_service`: `fmgdevice_switchcontroller_lldpprofile_medlocationservice`
>- `med_network_policy`: `fmgdevice_switchcontroller_lldpprofile_mednetworkpolicy`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_lldpprofile" "trname" {
  n8021_tlvs            = ["port-vlan-id"]
  n8023_tlvs            = ["power-negotiation"]
  auto_isl              = "enable"
  auto_isl_auth         = "legacy"
  auto_isl_auth_encrypt = "must"
  name                  = "your own value"
  device_name           = var.device_name # not required if setting is at provider
  device_vdom           = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `n8021_tlvs` - Transmitted IEEE 802.1 TLVs. Valid values: `port-vlan-id`.

* `n8023_tlvs` - Transmitted IEEE 802.3 TLVs. Valid values: `max-frame-size`, `power-negotiation`.

* `auto_isl` - Enable/disable auto inter-switch LAG. Valid values: `disable`, `enable`.

* `auto_isl_auth` - Auto inter-switch LAG authentication mode. Valid values: `legacy`, `strict`, `relax`.

* `auto_isl_auth_encrypt` - Auto inter-switch LAG encryption mode. Valid values: `none`, `mixed`, `must`.

* `auto_isl_auth_identity` - Auto inter-switch LAG authentication identity.
* `auto_isl_auth_macsec_profile` - Auto inter-switch LAG macsec profile for encryption.
* `auto_isl_auth_reauth` - Auto inter-switch LAG authentication reauth period in seconds(10 - 3600, default = 3600).
* `auto_isl_auth_user` - Auto inter-switch LAG authentication user certificate.
* `auto_isl_hello_timer` - Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
* `auto_isl_port_group` - Auto inter-switch LAG port group ID (0 - 9).
* `auto_isl_receive_timeout` - Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
* `auto_mclag_icl` - Enable/disable MCLAG inter chassis link. Valid values: `disable`, `enable`.

* `custom_tlvs` - Custom-Tlvs. The structure of `custom_tlvs` block is documented below.
* `med_location_service` - Med-Location-Service. The structure of `med_location_service` block is documented below.
* `med_network_policy` - Med-Network-Policy. The structure of `med_network_policy` block is documented below.
* `med_tlvs` - Transmitted LLDP-MED TLVs (type-length-value descriptions). Valid values: `inventory-management`, `network-policy`, `power-management`, `location-identification`.

* `name` - Profile name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `custom_tlvs` block supports:

* `information_string` - Organizationally defined information string (0 - 507 hexadecimal bytes).
* `name` - TLV name (not sent).
* `oui` - Organizationally unique identifier (OUI), a 3-byte hexadecimal number, for this TLV.
* `subtype` - Organizationally defined subtype (0 - 255).

The `med_location_service` block supports:

* `name` - Location service type name.
* `status` - Enable or disable this TLV. Valid values: `disable`, `enable`.

* `sys_location_id` - Location service ID.

The `med_network_policy` block supports:

* `assign_vlan` - Enable/disable VLAN assignment when this profile is applied on managed FortiSwitch port. Valid values: `disable`, `enable`.

* `dscp` - Advertised Differentiated Services Code Point (DSCP) value, a packet header value indicating the level of service requested for traffic, such as high priority or best effort delivery.
* `name` - Policy type name.
* `priority` - Advertised Layer 2 priority (0 - 7; from lowest to highest priority).
* `status` - Enable or disable this TLV. Valid values: `disable`, `enable`.

* `vlan` - ID of VLAN to advertise, if configured on port (0 - 4094, 0 = priority tag).
* `vlan_intf` - VLAN interface to advertise; if configured on port.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController LldpProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_lldpprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

