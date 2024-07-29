---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_location_addresscivic"
description: |-
  Configure location civic address.
---

# fmgdevice_switchcontroller_location_addresscivic
Configure location civic address.

~> This resource is a sub resource for variable `address_civic` of resource `fmgdevice_switchcontroller_location`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_location_addresscivic" "trname" {
  additional      = "your own value"
  additional_code = "your own value"
  block           = "your own value"
  branch_road     = "your own value"
  building        = "your own value"
  device_name     = var.device_name # not required if setting is at provider
  device_vdom     = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `location` - Location.

* `additional` - Location additional details.
* `additional_code` - Location additional code details.
* `block` - Location block details.
* `branch_road` - Location branch road details.
* `building` - Location building details.
* `city` - Location city details.
* `city_division` - Location city division details.
* `country` - The two-letter ISO 3166 country code in capital ASCII letters eg. US, CA, DK, DE.
* `country_subdivision` - National subdivisions (state, canton, region, province, or prefecture).
* `county` - County, parish, gun (JP), or district (IN).
* `direction` - Leading street direction.
* `floor` - Floor.
* `landmark` - Landmark or vanity address.
* `language` - Language.
* `name` - Name (residence and office occupant).
* `number` - House number.
* `number_suffix` - House number suffix.
* `parent_key` - Parent-Key.
* `place_type` - Place type.
* `post_office_box` - Post office box.
* `postal_community` - Postal community name.
* `primary_road` - Primary road name.
* `road_section` - Road section.
* `room` - Room number.
* `script` - Script used to present the address information.
* `seat` - Seat number.
* `street` - Street.
* `street_name_post_mod` - Street name post modifier.
* `street_name_pre_mod` - Street name pre modifier.
* `street_suffix` - Street suffix.
* `sub_branch_road` - Sub branch road name.
* `trailing_str_suffix` - Trailing street suffix.
* `unit` - Unit (apartment, suite).
* `zip` - Postal/zip code.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController LocationAddressCivic can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "location=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_location_addresscivic.labelname SwitchControllerLocationAddressCivic
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

