---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ssl_web_realm"
description: |-
  <i>This object will be purged after policy copy and install.</i> Realm.
---

# fmgdevice_vpn_ssl_web_realm
<i>This object will be purged after policy copy and install.</i> Realm.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `login_page` - Replacement HTML for Agentless VPN login page.
* `max_concurrent_user` - Maximum concurrent users (0 - 65535, 0 means unlimited).
* `nas_ip` - IP address used as a NAS-IP to communicate with the RADIUS server.
* `radius_port` - RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
* `radius_server` - RADIUS server associated with realm.
* `url_path` - URL path to access Agentless VPN login page.
* `virtual_host` - Virtual host name for realm.
* `virtual_host_only` - Enable/disable enforcement of virtual host method for Agentless VPN client access. Valid values: `disable`, `enable`.

* `virtual_host_server_cert` - Name of the server certificate to used for this realm.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{url_path}}.

## Import

Vpn SslWebRealm can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ssl_web_realm.labelname {{url_path}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

