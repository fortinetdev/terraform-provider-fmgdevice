// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure location civic address.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerLocationAddressCivic() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerLocationAddressCivicUpdate,
		Read:   resourceSwitchControllerLocationAddressCivicRead,
		Update: resourceSwitchControllerLocationAddressCivicUpdate,
		Delete: resourceSwitchControllerLocationAddressCivicDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"additional": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"additional_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"block": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"branch_road": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"building": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"city": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"city_division": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"country_subdivision": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"county": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"floor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"landmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"number_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"place_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"post_office_box": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"postal_community": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary_road": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"road_section": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"room": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"seat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"street": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"street_name_post_mod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"street_name_pre_mod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"street_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sub_branch_road": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trailing_str_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"zip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerLocationAddressCivicUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	location := d.Get("location").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["location"] = location

	obj, err := getObjectSwitchControllerLocationAddressCivic(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLocationAddressCivic resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerLocationAddressCivic(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLocationAddressCivic resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerLocationAddressCivic")

	return resourceSwitchControllerLocationAddressCivicRead(d, m)
}

func resourceSwitchControllerLocationAddressCivicDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	location := d.Get("location").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["location"] = location

	err = c.DeleteSwitchControllerLocationAddressCivic(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerLocationAddressCivic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLocationAddressCivicRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	location := d.Get("location").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if device_vdom == "" {
		device_vdom = importOptionChecking(m.(*FortiClient).Cfg, "device_vdom")
		if device_vdom == "" {
			return fmt.Errorf("Parameter device_vdom is missing")
		}
		if err = d.Set("device_vdom", device_vdom); err != nil {
			return fmt.Errorf("Error set params device_vdom: %v", err)
		}
	}
	if location == "" {
		location = importOptionChecking(m.(*FortiClient).Cfg, "location")
		if location == "" {
			return fmt.Errorf("Parameter location is missing")
		}
		if err = d.Set("location", location); err != nil {
			return fmt.Errorf("Error set params location: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["location"] = location

	o, err := c.ReadSwitchControllerLocationAddressCivic(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLocationAddressCivic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerLocationAddressCivic(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLocationAddressCivic resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerLocationAddressCivicAdditional2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicAdditionalCode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicBlock2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicBranchRoad2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicBuilding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicCity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicCityDivision2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicCountry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicCountrySubdivision2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicCounty2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicFloor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicLandmark2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicLanguage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicNumber2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicNumberSuffix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicParentKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicPlaceType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicPostOfficeBox2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicPostalCommunity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicPrimaryRoad2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicRoadSection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicRoom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicScript2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicSeat2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicStreet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicStreetNamePostMod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicStreetNamePreMod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicStreetSuffix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicSubBranchRoad2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicTrailingStrSuffix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicUnit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationAddressCivicZip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerLocationAddressCivic(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("additional", flattenSwitchControllerLocationAddressCivicAdditional2edl(o["additional"], d, "additional")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional"], "SwitchControllerLocationAddressCivic-Additional"); ok {
			if err = d.Set("additional", vv); err != nil {
				return fmt.Errorf("Error reading additional: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional: %v", err)
		}
	}

	if err = d.Set("additional_code", flattenSwitchControllerLocationAddressCivicAdditionalCode2edl(o["additional-code"], d, "additional_code")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-code"], "SwitchControllerLocationAddressCivic-AdditionalCode"); ok {
			if err = d.Set("additional_code", vv); err != nil {
				return fmt.Errorf("Error reading additional_code: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_code: %v", err)
		}
	}

	if err = d.Set("block", flattenSwitchControllerLocationAddressCivicBlock2edl(o["block"], d, "block")); err != nil {
		if vv, ok := fortiAPIPatch(o["block"], "SwitchControllerLocationAddressCivic-Block"); ok {
			if err = d.Set("block", vv); err != nil {
				return fmt.Errorf("Error reading block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block: %v", err)
		}
	}

	if err = d.Set("branch_road", flattenSwitchControllerLocationAddressCivicBranchRoad2edl(o["branch-road"], d, "branch_road")); err != nil {
		if vv, ok := fortiAPIPatch(o["branch-road"], "SwitchControllerLocationAddressCivic-BranchRoad"); ok {
			if err = d.Set("branch_road", vv); err != nil {
				return fmt.Errorf("Error reading branch_road: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading branch_road: %v", err)
		}
	}

	if err = d.Set("building", flattenSwitchControllerLocationAddressCivicBuilding2edl(o["building"], d, "building")); err != nil {
		if vv, ok := fortiAPIPatch(o["building"], "SwitchControllerLocationAddressCivic-Building"); ok {
			if err = d.Set("building", vv); err != nil {
				return fmt.Errorf("Error reading building: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading building: %v", err)
		}
	}

	if err = d.Set("city", flattenSwitchControllerLocationAddressCivicCity2edl(o["city"], d, "city")); err != nil {
		if vv, ok := fortiAPIPatch(o["city"], "SwitchControllerLocationAddressCivic-City"); ok {
			if err = d.Set("city", vv); err != nil {
				return fmt.Errorf("Error reading city: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading city: %v", err)
		}
	}

	if err = d.Set("city_division", flattenSwitchControllerLocationAddressCivicCityDivision2edl(o["city-division"], d, "city_division")); err != nil {
		if vv, ok := fortiAPIPatch(o["city-division"], "SwitchControllerLocationAddressCivic-CityDivision"); ok {
			if err = d.Set("city_division", vv); err != nil {
				return fmt.Errorf("Error reading city_division: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading city_division: %v", err)
		}
	}

	if err = d.Set("country", flattenSwitchControllerLocationAddressCivicCountry2edl(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "SwitchControllerLocationAddressCivic-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("country_subdivision", flattenSwitchControllerLocationAddressCivicCountrySubdivision2edl(o["country-subdivision"], d, "country_subdivision")); err != nil {
		if vv, ok := fortiAPIPatch(o["country-subdivision"], "SwitchControllerLocationAddressCivic-CountrySubdivision"); ok {
			if err = d.Set("country_subdivision", vv); err != nil {
				return fmt.Errorf("Error reading country_subdivision: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country_subdivision: %v", err)
		}
	}

	if err = d.Set("county", flattenSwitchControllerLocationAddressCivicCounty2edl(o["county"], d, "county")); err != nil {
		if vv, ok := fortiAPIPatch(o["county"], "SwitchControllerLocationAddressCivic-County"); ok {
			if err = d.Set("county", vv); err != nil {
				return fmt.Errorf("Error reading county: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading county: %v", err)
		}
	}

	if err = d.Set("direction", flattenSwitchControllerLocationAddressCivicDirection2edl(o["direction"], d, "direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["direction"], "SwitchControllerLocationAddressCivic-Direction"); ok {
			if err = d.Set("direction", vv); err != nil {
				return fmt.Errorf("Error reading direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("floor", flattenSwitchControllerLocationAddressCivicFloor2edl(o["floor"], d, "floor")); err != nil {
		if vv, ok := fortiAPIPatch(o["floor"], "SwitchControllerLocationAddressCivic-Floor"); ok {
			if err = d.Set("floor", vv); err != nil {
				return fmt.Errorf("Error reading floor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading floor: %v", err)
		}
	}

	if err = d.Set("landmark", flattenSwitchControllerLocationAddressCivicLandmark2edl(o["landmark"], d, "landmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["landmark"], "SwitchControllerLocationAddressCivic-Landmark"); ok {
			if err = d.Set("landmark", vv); err != nil {
				return fmt.Errorf("Error reading landmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading landmark: %v", err)
		}
	}

	if err = d.Set("language", flattenSwitchControllerLocationAddressCivicLanguage2edl(o["language"], d, "language")); err != nil {
		if vv, ok := fortiAPIPatch(o["language"], "SwitchControllerLocationAddressCivic-Language"); ok {
			if err = d.Set("language", vv); err != nil {
				return fmt.Errorf("Error reading language: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading language: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerLocationAddressCivicName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerLocationAddressCivic-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("number", flattenSwitchControllerLocationAddressCivicNumber2edl(o["number"], d, "number")); err != nil {
		if vv, ok := fortiAPIPatch(o["number"], "SwitchControllerLocationAddressCivic-Number"); ok {
			if err = d.Set("number", vv); err != nil {
				return fmt.Errorf("Error reading number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading number: %v", err)
		}
	}

	if err = d.Set("number_suffix", flattenSwitchControllerLocationAddressCivicNumberSuffix2edl(o["number-suffix"], d, "number_suffix")); err != nil {
		if vv, ok := fortiAPIPatch(o["number-suffix"], "SwitchControllerLocationAddressCivic-NumberSuffix"); ok {
			if err = d.Set("number_suffix", vv); err != nil {
				return fmt.Errorf("Error reading number_suffix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading number_suffix: %v", err)
		}
	}

	if err = d.Set("parent_key", flattenSwitchControllerLocationAddressCivicParentKey2edl(o["parent-key"], d, "parent_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["parent-key"], "SwitchControllerLocationAddressCivic-ParentKey"); ok {
			if err = d.Set("parent_key", vv); err != nil {
				return fmt.Errorf("Error reading parent_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading parent_key: %v", err)
		}
	}

	if err = d.Set("place_type", flattenSwitchControllerLocationAddressCivicPlaceType2edl(o["place-type"], d, "place_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["place-type"], "SwitchControllerLocationAddressCivic-PlaceType"); ok {
			if err = d.Set("place_type", vv); err != nil {
				return fmt.Errorf("Error reading place_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading place_type: %v", err)
		}
	}

	if err = d.Set("post_office_box", flattenSwitchControllerLocationAddressCivicPostOfficeBox2edl(o["post-office-box"], d, "post_office_box")); err != nil {
		if vv, ok := fortiAPIPatch(o["post-office-box"], "SwitchControllerLocationAddressCivic-PostOfficeBox"); ok {
			if err = d.Set("post_office_box", vv); err != nil {
				return fmt.Errorf("Error reading post_office_box: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading post_office_box: %v", err)
		}
	}

	if err = d.Set("postal_community", flattenSwitchControllerLocationAddressCivicPostalCommunity2edl(o["postal-community"], d, "postal_community")); err != nil {
		if vv, ok := fortiAPIPatch(o["postal-community"], "SwitchControllerLocationAddressCivic-PostalCommunity"); ok {
			if err = d.Set("postal_community", vv); err != nil {
				return fmt.Errorf("Error reading postal_community: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading postal_community: %v", err)
		}
	}

	if err = d.Set("primary_road", flattenSwitchControllerLocationAddressCivicPrimaryRoad2edl(o["primary-road"], d, "primary_road")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary-road"], "SwitchControllerLocationAddressCivic-PrimaryRoad"); ok {
			if err = d.Set("primary_road", vv); err != nil {
				return fmt.Errorf("Error reading primary_road: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_road: %v", err)
		}
	}

	if err = d.Set("road_section", flattenSwitchControllerLocationAddressCivicRoadSection2edl(o["road-section"], d, "road_section")); err != nil {
		if vv, ok := fortiAPIPatch(o["road-section"], "SwitchControllerLocationAddressCivic-RoadSection"); ok {
			if err = d.Set("road_section", vv); err != nil {
				return fmt.Errorf("Error reading road_section: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading road_section: %v", err)
		}
	}

	if err = d.Set("room", flattenSwitchControllerLocationAddressCivicRoom2edl(o["room"], d, "room")); err != nil {
		if vv, ok := fortiAPIPatch(o["room"], "SwitchControllerLocationAddressCivic-Room"); ok {
			if err = d.Set("room", vv); err != nil {
				return fmt.Errorf("Error reading room: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading room: %v", err)
		}
	}

	if err = d.Set("script", flattenSwitchControllerLocationAddressCivicScript2edl(o["script"], d, "script")); err != nil {
		if vv, ok := fortiAPIPatch(o["script"], "SwitchControllerLocationAddressCivic-Script"); ok {
			if err = d.Set("script", vv); err != nil {
				return fmt.Errorf("Error reading script: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	if err = d.Set("seat", flattenSwitchControllerLocationAddressCivicSeat2edl(o["seat"], d, "seat")); err != nil {
		if vv, ok := fortiAPIPatch(o["seat"], "SwitchControllerLocationAddressCivic-Seat"); ok {
			if err = d.Set("seat", vv); err != nil {
				return fmt.Errorf("Error reading seat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seat: %v", err)
		}
	}

	if err = d.Set("street", flattenSwitchControllerLocationAddressCivicStreet2edl(o["street"], d, "street")); err != nil {
		if vv, ok := fortiAPIPatch(o["street"], "SwitchControllerLocationAddressCivic-Street"); ok {
			if err = d.Set("street", vv); err != nil {
				return fmt.Errorf("Error reading street: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading street: %v", err)
		}
	}

	if err = d.Set("street_name_post_mod", flattenSwitchControllerLocationAddressCivicStreetNamePostMod2edl(o["street-name-post-mod"], d, "street_name_post_mod")); err != nil {
		if vv, ok := fortiAPIPatch(o["street-name-post-mod"], "SwitchControllerLocationAddressCivic-StreetNamePostMod"); ok {
			if err = d.Set("street_name_post_mod", vv); err != nil {
				return fmt.Errorf("Error reading street_name_post_mod: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading street_name_post_mod: %v", err)
		}
	}

	if err = d.Set("street_name_pre_mod", flattenSwitchControllerLocationAddressCivicStreetNamePreMod2edl(o["street-name-pre-mod"], d, "street_name_pre_mod")); err != nil {
		if vv, ok := fortiAPIPatch(o["street-name-pre-mod"], "SwitchControllerLocationAddressCivic-StreetNamePreMod"); ok {
			if err = d.Set("street_name_pre_mod", vv); err != nil {
				return fmt.Errorf("Error reading street_name_pre_mod: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading street_name_pre_mod: %v", err)
		}
	}

	if err = d.Set("street_suffix", flattenSwitchControllerLocationAddressCivicStreetSuffix2edl(o["street-suffix"], d, "street_suffix")); err != nil {
		if vv, ok := fortiAPIPatch(o["street-suffix"], "SwitchControllerLocationAddressCivic-StreetSuffix"); ok {
			if err = d.Set("street_suffix", vv); err != nil {
				return fmt.Errorf("Error reading street_suffix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading street_suffix: %v", err)
		}
	}

	if err = d.Set("sub_branch_road", flattenSwitchControllerLocationAddressCivicSubBranchRoad2edl(o["sub-branch-road"], d, "sub_branch_road")); err != nil {
		if vv, ok := fortiAPIPatch(o["sub-branch-road"], "SwitchControllerLocationAddressCivic-SubBranchRoad"); ok {
			if err = d.Set("sub_branch_road", vv); err != nil {
				return fmt.Errorf("Error reading sub_branch_road: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sub_branch_road: %v", err)
		}
	}

	if err = d.Set("trailing_str_suffix", flattenSwitchControllerLocationAddressCivicTrailingStrSuffix2edl(o["trailing-str-suffix"], d, "trailing_str_suffix")); err != nil {
		if vv, ok := fortiAPIPatch(o["trailing-str-suffix"], "SwitchControllerLocationAddressCivic-TrailingStrSuffix"); ok {
			if err = d.Set("trailing_str_suffix", vv); err != nil {
				return fmt.Errorf("Error reading trailing_str_suffix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trailing_str_suffix: %v", err)
		}
	}

	if err = d.Set("unit", flattenSwitchControllerLocationAddressCivicUnit2edl(o["unit"], d, "unit")); err != nil {
		if vv, ok := fortiAPIPatch(o["unit"], "SwitchControllerLocationAddressCivic-Unit"); ok {
			if err = d.Set("unit", vv); err != nil {
				return fmt.Errorf("Error reading unit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unit: %v", err)
		}
	}

	if err = d.Set("zip", flattenSwitchControllerLocationAddressCivicZip2edl(o["zip"], d, "zip")); err != nil {
		if vv, ok := fortiAPIPatch(o["zip"], "SwitchControllerLocationAddressCivic-Zip"); ok {
			if err = d.Set("zip", vv); err != nil {
				return fmt.Errorf("Error reading zip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading zip: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerLocationAddressCivicFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerLocationAddressCivicAdditional2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicAdditionalCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicBlock2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicBranchRoad2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicBuilding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicCity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicCityDivision2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicCountry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicCountrySubdivision2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicCounty2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicFloor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicLandmark2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicLanguage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicNumber2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicNumberSuffix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicParentKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicPlaceType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicPostOfficeBox2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicPostalCommunity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicPrimaryRoad2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicRoadSection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicRoom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicScript2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicSeat2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicStreet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicStreetNamePostMod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicStreetNamePreMod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicStreetSuffix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicSubBranchRoad2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicTrailingStrSuffix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicUnit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationAddressCivicZip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerLocationAddressCivic(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("additional"); ok || d.HasChange("additional") {
		t, err := expandSwitchControllerLocationAddressCivicAdditional2edl(d, v, "additional")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional"] = t
		}
	}

	if v, ok := d.GetOk("additional_code"); ok || d.HasChange("additional_code") {
		t, err := expandSwitchControllerLocationAddressCivicAdditionalCode2edl(d, v, "additional_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-code"] = t
		}
	}

	if v, ok := d.GetOk("block"); ok || d.HasChange("block") {
		t, err := expandSwitchControllerLocationAddressCivicBlock2edl(d, v, "block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block"] = t
		}
	}

	if v, ok := d.GetOk("branch_road"); ok || d.HasChange("branch_road") {
		t, err := expandSwitchControllerLocationAddressCivicBranchRoad2edl(d, v, "branch_road")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["branch-road"] = t
		}
	}

	if v, ok := d.GetOk("building"); ok || d.HasChange("building") {
		t, err := expandSwitchControllerLocationAddressCivicBuilding2edl(d, v, "building")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["building"] = t
		}
	}

	if v, ok := d.GetOk("city"); ok || d.HasChange("city") {
		t, err := expandSwitchControllerLocationAddressCivicCity2edl(d, v, "city")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["city"] = t
		}
	}

	if v, ok := d.GetOk("city_division"); ok || d.HasChange("city_division") {
		t, err := expandSwitchControllerLocationAddressCivicCityDivision2edl(d, v, "city_division")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["city-division"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok || d.HasChange("country") {
		t, err := expandSwitchControllerLocationAddressCivicCountry2edl(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("country_subdivision"); ok || d.HasChange("country_subdivision") {
		t, err := expandSwitchControllerLocationAddressCivicCountrySubdivision2edl(d, v, "country_subdivision")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country-subdivision"] = t
		}
	}

	if v, ok := d.GetOk("county"); ok || d.HasChange("county") {
		t, err := expandSwitchControllerLocationAddressCivicCounty2edl(d, v, "county")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["county"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok || d.HasChange("direction") {
		t, err := expandSwitchControllerLocationAddressCivicDirection2edl(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("floor"); ok || d.HasChange("floor") {
		t, err := expandSwitchControllerLocationAddressCivicFloor2edl(d, v, "floor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["floor"] = t
		}
	}

	if v, ok := d.GetOk("landmark"); ok || d.HasChange("landmark") {
		t, err := expandSwitchControllerLocationAddressCivicLandmark2edl(d, v, "landmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["landmark"] = t
		}
	}

	if v, ok := d.GetOk("language"); ok || d.HasChange("language") {
		t, err := expandSwitchControllerLocationAddressCivicLanguage2edl(d, v, "language")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["language"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerLocationAddressCivicName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("number"); ok || d.HasChange("number") {
		t, err := expandSwitchControllerLocationAddressCivicNumber2edl(d, v, "number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["number"] = t
		}
	}

	if v, ok := d.GetOk("number_suffix"); ok || d.HasChange("number_suffix") {
		t, err := expandSwitchControllerLocationAddressCivicNumberSuffix2edl(d, v, "number_suffix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["number-suffix"] = t
		}
	}

	if v, ok := d.GetOk("parent_key"); ok || d.HasChange("parent_key") {
		t, err := expandSwitchControllerLocationAddressCivicParentKey2edl(d, v, "parent_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parent-key"] = t
		}
	}

	if v, ok := d.GetOk("place_type"); ok || d.HasChange("place_type") {
		t, err := expandSwitchControllerLocationAddressCivicPlaceType2edl(d, v, "place_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["place-type"] = t
		}
	}

	if v, ok := d.GetOk("post_office_box"); ok || d.HasChange("post_office_box") {
		t, err := expandSwitchControllerLocationAddressCivicPostOfficeBox2edl(d, v, "post_office_box")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["post-office-box"] = t
		}
	}

	if v, ok := d.GetOk("postal_community"); ok || d.HasChange("postal_community") {
		t, err := expandSwitchControllerLocationAddressCivicPostalCommunity2edl(d, v, "postal_community")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["postal-community"] = t
		}
	}

	if v, ok := d.GetOk("primary_road"); ok || d.HasChange("primary_road") {
		t, err := expandSwitchControllerLocationAddressCivicPrimaryRoad2edl(d, v, "primary_road")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-road"] = t
		}
	}

	if v, ok := d.GetOk("road_section"); ok || d.HasChange("road_section") {
		t, err := expandSwitchControllerLocationAddressCivicRoadSection2edl(d, v, "road_section")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["road-section"] = t
		}
	}

	if v, ok := d.GetOk("room"); ok || d.HasChange("room") {
		t, err := expandSwitchControllerLocationAddressCivicRoom2edl(d, v, "room")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["room"] = t
		}
	}

	if v, ok := d.GetOk("script"); ok || d.HasChange("script") {
		t, err := expandSwitchControllerLocationAddressCivicScript2edl(d, v, "script")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script"] = t
		}
	}

	if v, ok := d.GetOk("seat"); ok || d.HasChange("seat") {
		t, err := expandSwitchControllerLocationAddressCivicSeat2edl(d, v, "seat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seat"] = t
		}
	}

	if v, ok := d.GetOk("street"); ok || d.HasChange("street") {
		t, err := expandSwitchControllerLocationAddressCivicStreet2edl(d, v, "street")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["street"] = t
		}
	}

	if v, ok := d.GetOk("street_name_post_mod"); ok || d.HasChange("street_name_post_mod") {
		t, err := expandSwitchControllerLocationAddressCivicStreetNamePostMod2edl(d, v, "street_name_post_mod")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["street-name-post-mod"] = t
		}
	}

	if v, ok := d.GetOk("street_name_pre_mod"); ok || d.HasChange("street_name_pre_mod") {
		t, err := expandSwitchControllerLocationAddressCivicStreetNamePreMod2edl(d, v, "street_name_pre_mod")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["street-name-pre-mod"] = t
		}
	}

	if v, ok := d.GetOk("street_suffix"); ok || d.HasChange("street_suffix") {
		t, err := expandSwitchControllerLocationAddressCivicStreetSuffix2edl(d, v, "street_suffix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["street-suffix"] = t
		}
	}

	if v, ok := d.GetOk("sub_branch_road"); ok || d.HasChange("sub_branch_road") {
		t, err := expandSwitchControllerLocationAddressCivicSubBranchRoad2edl(d, v, "sub_branch_road")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub-branch-road"] = t
		}
	}

	if v, ok := d.GetOk("trailing_str_suffix"); ok || d.HasChange("trailing_str_suffix") {
		t, err := expandSwitchControllerLocationAddressCivicTrailingStrSuffix2edl(d, v, "trailing_str_suffix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trailing-str-suffix"] = t
		}
	}

	if v, ok := d.GetOk("unit"); ok || d.HasChange("unit") {
		t, err := expandSwitchControllerLocationAddressCivicUnit2edl(d, v, "unit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unit"] = t
		}
	}

	if v, ok := d.GetOk("zip"); ok || d.HasChange("zip") {
		t, err := expandSwitchControllerLocationAddressCivicZip2edl(d, v, "zip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zip"] = t
		}
	}

	return &obj, nil
}
