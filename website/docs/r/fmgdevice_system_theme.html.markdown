---
subcategory: "System"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_theme"
description: |-
  Configure custom gui themes.
---

# fmgdevice_system_theme
Configure custom gui themes.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `accent_color` - Hexidecimal color code for the accent color.
* `banner_msg` - Optional message to appear on the fortigate header.
* `banner_msg_severity` - Severity color of the banner message. Valid values: `info`, `warning`, `critical`.

* `base_theme` - Base theme that will be used as a template. Valid values: `light`, `dark`.

* `border_radius` - Enable/disable border radius. Valid values: `disable`, `enable`.

* `call_to_action_color` - Hexidecimal color code for the call to action color.
* `font` - Font family. Valid values: `lato`, `inter`.

* `font_weight` - Font weight. Valid values: `light`, `standard`.

* `header_color` - Hexidecimal color code for the header color.
* `name` - Custom theme name.
* `nav_color` - Color of the navigation bar. Valid values: `light`, `dark`, `gray`.

* `nav_style` - Navigation bar style. Valid values: `standard`, `full-height`.

* `selected_color` - Hexidecimal color code for the selected color.
* `table_style` - Table style. Valid values: `fill`, `float`.

* `theme_template` - Theme template that will be used in place of the base theme. Valid values: `melongene`, `mariner`, `neutrino`, `jade`, `graphite`, `dark-matter`, `onyx`, `eclipse`, `retro`, `jet-stream`, `security-fabric`, `none`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Theme can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_theme.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

