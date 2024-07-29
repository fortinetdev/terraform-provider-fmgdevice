---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_emailserver"
description: |-
  Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.
---

# fmgdevice_system_emailserver
Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.

## Example Usage

```hcl
resource "fmgdevice_system_emailserver" "trname" {
  authenticate            = "enable"
  interface               = ["port2"]
  interface_select_method = "specify"
  password                = ["your own value"]
  port                    = 10
  device_name             = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `authenticate` - Enable/disable authentication. Valid values: `disable`, `enable`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `password` - SMTP server user password for authentication.
* `port` - SMTP server port.
* `reply_to` - Reply-To email address.
* `security` - Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.

* `server` - SMTP server IP address or hostname.
* `source_ip` - SMTP server IPv4 source IP.
* `source_ip6` - SMTP server IPv6 source IP.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1-3`.

* `type` - Use FortiGuard Message service or custom email server. Valid values: `custom`.

* `username` - SMTP server user name for authentication.
* `validate_server` - Enable/disable validation of server certificate. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System EmailServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_emailserver.labelname SystemEmailServer
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

