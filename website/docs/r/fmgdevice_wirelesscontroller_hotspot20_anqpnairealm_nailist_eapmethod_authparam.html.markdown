---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod_authparam"
description: |-
  EAP auth param.
---

# fmgdevice_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod_authparam
EAP auth param.

~> This resource is a sub resource for variable `auth_param` of resource `fmgdevice_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod_authparam" "trname" {
  fosid       = "tunneled-credential"
  index       = 10
  val         = "eap-md5"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `anqp_nai_realm` - Anqp Nai Realm.
* `nai_list` - Nai List.
* `eap_method` - Eap Method.

* `fosid` - ID of authentication parameter. Valid values: `non-eap-inner-auth`, `inner-auth-eap`, `credential`, `tunneled-credential`.

* `index` - Param index.
* `val` - Value of authentication parameter. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`, `non-eap-pap`, `non-eap-chap`, `non-eap-mschap`, `non-eap-mschapv2`, `cred-sim`, `cred-usim`, `cred-nfc`, `cred-hardware-token`, `cred-softoken`, `cred-certificate`, `cred-user-pwd`, `cred-none`, `cred-vendor-specific`, `tun-cred-sim`, `tun-cred-usim`, `tun-cred-nfc`, `tun-cred-hardware-token`, `tun-cred-softoken`, `tun-cred-certificate`, `tun-cred-user-pwd`, `tun-cred-anonymous`, `tun-cred-vendor-specific`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

WirelessController Hotspot20AnqpNaiRealmNaiListEapMethodAuthParam can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "anqp_nai_realm=YOUR_VALUE", "nai_list=YOUR_VALUE", "eap_method=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod_authparam.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

