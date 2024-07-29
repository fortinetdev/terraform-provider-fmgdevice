---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_h2qpconncapability"
description: |-
  Configure connection capability.
---

# fmgdevice_wirelesscontroller_hotspot20_h2qpconncapability
Configure connection capability.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_h2qpconncapability" "trname" {
  esp_port    = "open"
  ftp_port    = "open"
  http_port   = "closed"
  icmp_port   = "closed"
  ikev2_port  = "open"
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `esp_port` - Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.

* `ftp_port` - Set FTP port service status. Valid values: `closed`, `open`, `unknown`.

* `http_port` - Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.

* `icmp_port` - Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.

* `ikev2_port` - Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.

* `ikev2_xx_port` - Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.

* `name` - Connection capability name.
* `pptp_vpn_port` - Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.

* `ssh_port` - Set SSH port service status. Valid values: `closed`, `open`, `unknown`.

* `tls_port` - Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.

* `voip_tcp_port` - Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.

* `voip_udp_port` - Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20H2QpConnCapability can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_h2qpconncapability.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

