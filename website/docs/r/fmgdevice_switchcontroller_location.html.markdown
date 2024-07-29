---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_location"
description: |-
  Configure FortiSwitch location services.
---

# fmgdevice_switchcontroller_location
Configure FortiSwitch location services.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `address_civic`: `fmgdevice_switchcontroller_location_addresscivic`
>- `coordinates`: `fmgdevice_switchcontroller_location_coordinates`
>- `elin_number`: `fmgdevice_switchcontroller_location_elinnumber`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_location" "trname" {
  address_civic {
    additional           = "your own value"
    additional_code      = "your own value"
    block                = "your own value"
    branch_road          = "your own value"
    building             = "your own value"
    city                 = "your own value"
    city_division        = "your own value"
    country              = "your own value"
    country_subdivision  = "your own value"
    county               = "your own value"
    direction            = "your own value"
    floor                = "your own value"
    landmark             = "your own value"
    language             = "your own value"
    name                 = "your own value"
    number               = "your own value"
    number_suffix        = "your own value"
    parent_key           = "your own value"
    place_type           = "your own value"
    post_office_box      = "your own value"
    postal_community     = "your own value"
    primary_road         = "your own value"
    road_section         = "your own value"
    room                 = "your own value"
    script               = "your own value"
    seat                 = "your own value"
    street               = "your own value"
    street_name_post_mod = "your own value"
    street_name_pre_mod  = "your own value"
    street_suffix        = "your own value"
    sub_branch_road      = "your own value"
    trailing_str_suffix  = "your own value"
    unit                 = "your own value"
    zip                  = "your own value"
  }

  coordinates {
    altitude      = "your own value"
    altitude_unit = "m"
    datum         = "WGS84"
    latitude      = "your own value"
    longitude     = "your own value"
    parent_key    = "your own value"
  }

  elin_number {
    elin_num   = "your own value"
    parent_key = "your own value"
  }

  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `address_civic` - Address-Civic. The structure of `address_civic` block is documented below.
* `coordinates` - Coordinates. The structure of `coordinates` block is documented below.
* `elin_number` - Elin-Number. The structure of `elin_number` block is documented below.
* `name` - Unique location item name.

The `address_civic` block supports:

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

The `coordinates` block supports:

* `altitude` - Plus or minus floating point number. For example, 117.47.
* `altitude_unit` - Configure the unit for which the altitude is to (m = meters, f = floors of a building). Valid values: `m`, `f`.

* `datum` - WGS84, NAD83, NAD83/MLLW. Valid values: `WGS84`, `NAD83`, `NAD83/MLLW`.

* `latitude` - Floating point starting with +/- or ending with (N or S). For example, +/-16.67 or 16.67N.
* `longitude` - Floating point starting with +/- or ending with (N or S). For example, +/-26.789 or 26.789E.
* `parent_key` - Parent-Key.

The `elin_number` block supports:

* `elin_num` - Configure ELIN callback number.
* `parent_key` - Parent-Key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController Location can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_location.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

