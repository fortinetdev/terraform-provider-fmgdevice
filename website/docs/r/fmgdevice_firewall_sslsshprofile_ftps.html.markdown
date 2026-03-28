---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_sslsshprofile_ftps"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure FTPS options.
---

# fmgdevice_firewall_sslsshprofile_ftps
<i>This object will be purged after policy copy and install.</i> Configure FTPS options.

~> This resource is a sub resource for variable `ftps` of resource `fmgdevice_firewall_sslsshprofile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `ssl_ssh_profile` - Ssl Ssh Profile.

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.

* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.

* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.

* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.

* `min_allowed_ssl_version` - Minimum SSL version to be allowed. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.

* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `disable`, `enable`, `strict`.

* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.

* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.

* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.

* `unsupported_ssl_version` - Action based on the SSL version used being unsupported. Valid values: `block`, `allow`.

* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow`, `block`, `ignore`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall SslSshProfileFtps can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "ssl_ssh_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_sslsshprofile_ftps.labelname FirewallSslSshProfileFtps
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

