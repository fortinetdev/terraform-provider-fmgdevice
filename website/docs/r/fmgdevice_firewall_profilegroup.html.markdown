---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_profilegroup"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure profile groups.
---

# fmgdevice_firewall_profilegroup
<i>This object will be purged after policy copy and install.</i> Configure profile groups.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `application_list` - Name of an existing Application list.
* `av_profile` - Name of an existing Antivirus profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `casb_profile` - Name of an existing CASB profile.
* `diameter_filter_profile` - Name of an existing Diameter filter profile.
* `dlp_profile` - Name of an existing DLP profile.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `emailfilter_profile` - Name of an existing email filter profile.
* `file_filter_profile` - Name of an existing file-filter profile.
* `icap_profile` - Name of an existing ICAP profile.
* `ips_sensor` - Name of an existing IPS sensor.
* `ips_voip_filter` - Name of an existing VoIP (ips) profile.
* `name` - Profile group name.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `telemetry_profile` - Name of an existing telemetry profile.
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `virtual_patch_profile` - Name of an existing virtual-patch profile.
* `voip_profile` - Name of an existing VoIP (voipd) profile.
* `waf_profile` - Name of an existing Web application firewall profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `dlp_sensor` - Dlp-Sensor.
* `ia_profile` - Ia-Profile.
* `isolator_profile` - Isolator-Profile.
* `redirect_profile` - Redirect-Profile.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ProfileGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_profilegroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

