---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_lldp_networkpolicy"
description: |-
  Configure LLDP network policy.
---

# fmgdevice_system_lldp_networkpolicy
Configure LLDP network policy.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `guest`: `fmgdevice_system_lldp_networkpolicy_guest`
>- `guest_voice_signaling`: `fmgdevice_system_lldp_networkpolicy_guestvoicesignaling`
>- `softphone`: `fmgdevice_system_lldp_networkpolicy_softphone`
>- `streaming_video`: `fmgdevice_system_lldp_networkpolicy_streamingvideo`
>- `video_conferencing`: `fmgdevice_system_lldp_networkpolicy_videoconferencing`
>- `video_signaling`: `fmgdevice_system_lldp_networkpolicy_videosignaling`
>- `voice`: `fmgdevice_system_lldp_networkpolicy_voice`
>- `voice_signaling`: `fmgdevice_system_lldp_networkpolicy_voicesignaling`



## Example Usage

```hcl
resource "fmgdevice_system_lldp_networkpolicy" "trname" {
  comment = "your own value"
  guest {
    dscp     = 10
    priority = 10
    status   = "disable"
    tag      = "dot1p"
    vlan     = 10
  }

  guest_voice_signaling {
    dscp     = 10
    priority = 10
    status   = "disable"
    tag      = "none"
    vlan     = 10
  }

  name = "your own value"
  softphone {
    dscp     = 10
    priority = 10
    status   = "enable"
    tag      = "none"
    vlan     = 10
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `guest` - Guest. The structure of `guest` block is documented below.
* `guest_voice_signaling` - Guest-Voice-Signaling. The structure of `guest_voice_signaling` block is documented below.
* `name` - LLDP network policy name.
* `softphone` - Softphone. The structure of `softphone` block is documented below.
* `streaming_video` - Streaming-Video. The structure of `streaming_video` block is documented below.
* `video_conferencing` - Video-Conferencing. The structure of `video_conferencing` block is documented below.
* `video_signaling` - Video-Signaling. The structure of `video_signaling` block is documented below.
* `voice` - Voice. The structure of `voice` block is documented below.
* `voice_signaling` - Voice-Signaling. The structure of `voice_signaling` block is documented below.

The `guest` block supports:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy. Valid values: `disable`, `enable`.

* `tag` - Advertise tagged or untagged traffic. Valid values: `none`, `dot1q`, `dot1p`.

* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).

The `guest_voice_signaling` block supports:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy. Valid values: `disable`, `enable`.

* `tag` - Advertise tagged or untagged traffic. Valid values: `none`, `dot1q`, `dot1p`.

* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).

The `softphone` block supports:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy. Valid values: `disable`, `enable`.

* `tag` - Advertise tagged or untagged traffic. Valid values: `none`, `dot1q`, `dot1p`.

* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).

The `streaming_video` block supports:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy. Valid values: `disable`, `enable`.

* `tag` - Advertise tagged or untagged traffic. Valid values: `none`, `dot1q`, `dot1p`.

* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).

The `video_conferencing` block supports:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy. Valid values: `disable`, `enable`.

* `tag` - Advertise tagged or untagged traffic. Valid values: `none`, `dot1q`, `dot1p`.

* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).

The `video_signaling` block supports:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy. Valid values: `disable`, `enable`.

* `tag` - Advertise tagged or untagged traffic. Valid values: `none`, `dot1q`, `dot1p`.

* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).

The `voice` block supports:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy. Valid values: `disable`, `enable`.

* `tag` - Advertise tagged or untagged traffic. Valid values: `none`, `dot1q`, `dot1p`.

* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).

The `voice_signaling` block supports:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy. Valid values: `disable`, `enable`.

* `tag` - Advertise tagged or untagged traffic. Valid values: `none`, `dot1q`, `dot1p`.

* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System LldpNetworkPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_lldp_networkpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

