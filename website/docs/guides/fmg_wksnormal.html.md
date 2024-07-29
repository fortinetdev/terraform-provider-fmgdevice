---
subcategory: ""
layout: "fmgdevice"
page_title: "To Lock for Restricting Configuration Changes"
description: |-
  To Lock for Restricting Configuration Changes
---


# To Lock for Restricting Configuration Changes

Workspace enables locking ADOMs, devices, or policy packages so that an administrator can prevent other administrators from making changes to the elements that they are working in.

The provider currently supports `normal mode`. In Normal mode, ADOMs, or individual devices or policy packages must be locked before policy, object, or device changes can be made. Multiple administrators can lock devices and policy packages within a single, unlocked ADOM at the same time. When an individual device or policy package is locked, other administrators can only lock the ADOM that contains the locked device or policy package by disconnecting the administrator that locked it.

## Example

### Step 1 Set workspace_mode to normal

```hcl
provider "fmgdevice" {
  hostname     = "myfirewall"
  username     = "admin"
  password     = "admin"
  insecure     = "false"
  cabundlefile = "server.crt"

  device_name    = "terraform-test"
  device_vdom    = "root"
}

resource "fmgdevice_router_static" "labelname" {
  vdom = "root"
  dev = ["port6"]
  gateway = "1.1.1.2"
  dst = ["11.12.11.0", "255.255.255.0"]
}
```

### Step 2 Wrap the resources that need to be locked with two fmgdevice_exec_workspace_action resources

#### Example 

```hcl
resource "fmgdevice_exec_workspace_action" "lockres" { # lock root ADOM
  adom           = "global"
  action         = "lockbegin"
  target         = ""
  param          = ""
  force_recreate = uuid()
  comment        = ""
}

resource "fmgdevice_router_static" "labelname" {
  vdom = "root"
  dev = ["port6"]
  gateway = "1.1.1.2"
  dst = ["11.12.11.0", "255.255.255.0"]
  depends_on = [fmgdevice_exec_workspace_action.lockres]
}

resource "fmgdevice_exec_workspace_action" "unlockres" { # save change and unlock root ADOM
  adom           = "global"
  action         = "lockend"
  target         = ""
  param          = ""
  force_recreate = uuid()
  comment        = ""
  depends_on     = [fmgdevice_router_static.labelname]
}

```
