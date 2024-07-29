---
layout: "fmgdevice"
page_title: "Provider: FmgDevice"
sidebar_current: "docs-fmgdevice-index"
description: |-
  The FmgDevice provider interacts with FortiManager managed devices.
---

# FmgDevice Provider

The FmgDevice provider is used to interact with the device related resources supported by FortiManager. We need to configure the provider with the proper credentials before it can be used. Please use the navigation on the left to read more details about the available resources.

## Example Usage

```hcl
# Configure the Provider for FortiManager
provider "fmgdevice" {
  hostname     = "192.168.52.178"
  username     = "admin"
  password     = "admin"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"

  device_name  = "terraform-test"
  device_vdom  = "root"
}

# Create a firewall vip object
resource "fmgdevice_router_static" "labelname" {
  device = "terraform-test"
  vdom = "root"
  dev = ["port6"]
  gateway = "1.1.1.2"
  dst = ["11.12.11.0", "255.255.255.0"]
}

```

Before using this provider, the permission level for rpc-permit need to be set. See `Guides->To Set the Permission Level for RPC-Permit` for details.

If it is used for testing, you can set `insecure` to "true" and unset `cabundlefile` to quickly set the provider up, for example:

```hcl
provider "fmgdevice" {
  hostname     = "192.168.52.178"
  username     = "admin"
  password     = "admin"
  insecure     = "true"

  device_name  = "terraform-test"
  device_vdom  = "root"
}
```

Please refer to the Argument Reference below for more help on `insecure` and `cabundlefile`.


## Authentication

The FmgDevice provider offers a means of providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

### Static credentials

Static credentials can be provided by adding credential keys in-line in the FmgDevice provider block. 

There are two kinds of credentials supported for on-prem FortiManager.
- `token` based authentication (Recommanded). User needs to generate an API token from FortiManager. *Note: Only FortiManager version >= v7.2.2 supports Token based authentication.*
- `username/password` authentication. User provide the username and password of the administrator. 

There are two kinds of credentials supported for FortiManager Cloud.
- Provide `fmg_cloud_token` directly. User needs to generate an FortiCloud token. *Note: The Token could be expired. Make sure the Token provided is valid.*
- `username/password` authentication. User provide the username and password of the FortiCloud API user. The provider will generate the FortiCloud token based on username/password. 

Usage:

```hcl
provider "fmgdevice" {
  hostname     = "192.168.52.178"
  token        = "4ktitbsdiuye6ja43aoxafuqcb15kzse"
  insecure     = "true"

  device_name  = "terraform-test"
  device_vdom  = "root"
}
```

#### Generate an API token for FortiManager

See the left navigation: `Guides` -> `Generate an API token for FortiManager`.

#### Create FortiCloud API user and generate FortiCloud token for FortiManager Cloud

See the left navigation: `Guides` -> `Generate an FortiCloud token for FortiManager Cloud`.

### Environment variables

You can provide your credentials via the `FORTIMANAGER_ACCESS_HOSTNAME`, `FORTIMANAGER_ACCESS_TOKEN`, `FORTIMANAGER_CLOUD_ACCESS_TOKEN`,`FORTIMANAGER_ACCESS_USERNAME`, `FORTIMANAGER_ACCESS_PASSWORD`, `FORTIMANAGER_INSECURE` and `FORTIMANAGER_CA_CABUNDLE` environment variables. Note that setting your FortiManager credentials using static credentials variables will override the environment variables.

Usage:

```shell
$ export "FORTIMANAGER_ACCESS_HOSTNAME"="192.168.52.178"
$ export "FORTIMANAGER_ACCESS_USERNAME"="admin"
$ export "FORTIMANAGER_ACCESS_PASSWORD"="admin"
$ export "FORTIMANAGER_ACCESS_TOKEN"="4ktitbsdiuye6ja43aoxafuqcb15kzse"
$ export "FORTIMANAGER_INSECURE"="false"
$ export "FORTIMANAGER_CA_CABUNDLE"="/path/yourCA.crt"
```

Then configure the FmgDevice Provider as following:

```hcl
provider "fmgdevice" {
  device_name = "terraform-test"
  device_vdom = "root"
}
```

## Argument Reference

The following arguments are supported:

* `hostname` - (String | Optional) The hostname or IP address of FortiManager unit. It must be provided, but it can also be sourced from the `FORTIMANAGER_ACCESS_HOSTNAME` environment variable.

* `token` - (String | Optional) The token of FortiManager unit. If omitted, the `FORTIMANAGER_ACCESS_TOKEN` environment variable will be used. If neither is set, username/password will be used.

* `fmg_cloud_token` - (String | Optional) The access token of FortiManager Cloud. If omitted, the `FORTIMANAGER_CLOUD_ACCESS_TOKEN` environment variable will be used. If neither is set, username/password will be used. Available only when `fmg_type` set to `forticloud`. 

* `fmg_type` - (String | Optional) FortiManager type. Valid values: `on-prem`, `forticloud`. Default is `on-prem`. Set to `forticloud` if using FortiManager Cloud under FortiCloud.

* `username` - (String | Optional) FortiManager username if `fmg_type` is `on-prem`, and it is required. FortiCloud API username if `fmg_type` is `forticloud`, and it is optional. It can also be sourced from the `FORTIMANAGER_ACCESS_USERNAME` environment variable.

* `password` - (String | Optional) FortiManager password if `fmg_type` is `on-prem`, and it is required. FortiCloud API password if `fmg_type` is `forticloud`, and it is optional. It can also be sourced from the `FORTIMANAGER_ACCESS_PASSWORD` environment variable.

* `insecure` - (Bool | Optional) Control whether the Provider to perform insecure SSL requests. If omitted, the `FORTIMANAGER_INSECURE` environment variable is used. If neither is set, default value is `false`.

* `cabundlefile` - (String | Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by the `FORTIMANAGER_CA_CABUNDLE` environment variable.

* `device_name` - (String | Optional) FortiManager managed device name. This variable is used in the request URL. Each resource can also set its own device name as needed. Resource level configuration has higher priority.

* `device_vdom` - (String | Optional) FortiManager managed device vdom. This variable is used in the request URL. Each resource can also set its own device vdom as needed. Resource level configuration has higher priority.

* `import_options` - (List | Optional) This parameter is only used for import in some special cases. When the resource to be imported includes pkg parameter, you need to assign a value to the parameter here, for example:

    ```hcl
    provider "fmgdevice" {
      hostname       = "192.168.52.178"
      username       = "admin"
      password       = "admin"
      insecure       = "true"

      device_name    = "terraform-test"
      device_vdom    = "root"

      import_options = ["pkg=default"]
    }
    ```

* `logsession` - (Bool | Optional) Save the session to a local file. Used to assist fmgdevice_exec_workspace_action resource. Valid values: `true`: log to file, `false`: do not log to file. Default is `false`. See `Guides -> To Lock for Restricting Configuration Changes` for details.

* `presession` - (String | Optional) The session saved earlier and within the validity period, used to reuse the previous session and assist fmgdevice_exec_workspace_action resource. See `Guides -> To Lock for Restricting Configuration Changes` for details. Default is empty.

* `clean_session` - (Bool | Optional) Whether clean sessions. **Only works on workspace mode set to `disabled`.** If set to `true`, the provider will generate and logout the session for each HTTPS request. If set to `false`, the provider will generate a session for each Terraform operation. But the session will stay exist until it expires. Default is `false`.


## Release
Check out the FmgDevice provider release notes and additional information from: [the FmgDevice provider releases](https://github.com/fortinetdev/terraform-provider-fmgdevice/releases).


## FmgDevice best practices

FortiManager is an integrated platform for the centralized management of products in a Fortinet security infrastructure, including FortiGates.

Once FortiGates are managed by a FortiManager that is operating in Normal Mode, whenever possible, configuration changes should be made on the FortiManager and not the FortiGate. This recommendation also applies when the configuration of FortiGates and FortiManager is executed through the [FortiOS](https://registry.terraform.io/providers/fortinetdev/fortios/latest) and [FortiManager](https://registry.terraform.io/providers/fortinetdev/fortimanager/latest) providers.

To help you get the most out of your FortiManager products, maximize performance, and avoid potential problems, please refer to the [FortiManager documentation](https://docs.fortinet.com/product/fortimanager), including the Administration Guide and Best Practices documents.

Fortinet also provides a developer community to help administrators and advanced users enhance and increase the effectiveness of Fortinet products. The [Fortinet Developer Network (FNDN)](https://fndn.fortinet.net/) provides the official documentation and advanced tools for developing custom solutions using Fortinet products.


## Versioning

The provider can cover FortiManager >=7.4.3 versions, the configuration of all parameters should be based on the relevant FortiManager version manual and FortiManager API guide.
