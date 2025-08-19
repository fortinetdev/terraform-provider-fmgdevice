// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure location GPS coordinates.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerLocationCoordinates() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerLocationCoordinatesUpdate,
		Read:   resourceSwitchControllerLocationCoordinatesRead,
		Update: resourceSwitchControllerLocationCoordinatesUpdate,
		Delete: resourceSwitchControllerLocationCoordinatesDelete,

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
			"altitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"altitude_unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"latitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"longitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerLocationCoordinatesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerLocationCoordinates(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLocationCoordinates resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerLocationCoordinates(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLocationCoordinates resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerLocationCoordinates")

	return resourceSwitchControllerLocationCoordinatesRead(d, m)
}

func resourceSwitchControllerLocationCoordinatesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerLocationCoordinates(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerLocationCoordinates resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLocationCoordinatesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerLocationCoordinates(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLocationCoordinates resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerLocationCoordinates(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLocationCoordinates resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerLocationCoordinatesAltitude2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationCoordinatesAltitudeUnit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationCoordinatesDatum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationCoordinatesLatitude2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationCoordinatesLongitude2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLocationCoordinatesParentKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerLocationCoordinates(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("altitude", flattenSwitchControllerLocationCoordinatesAltitude2edl(o["altitude"], d, "altitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["altitude"], "SwitchControllerLocationCoordinates-Altitude"); ok {
			if err = d.Set("altitude", vv); err != nil {
				return fmt.Errorf("Error reading altitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading altitude: %v", err)
		}
	}

	if err = d.Set("altitude_unit", flattenSwitchControllerLocationCoordinatesAltitudeUnit2edl(o["altitude-unit"], d, "altitude_unit")); err != nil {
		if vv, ok := fortiAPIPatch(o["altitude-unit"], "SwitchControllerLocationCoordinates-AltitudeUnit"); ok {
			if err = d.Set("altitude_unit", vv); err != nil {
				return fmt.Errorf("Error reading altitude_unit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading altitude_unit: %v", err)
		}
	}

	if err = d.Set("datum", flattenSwitchControllerLocationCoordinatesDatum2edl(o["datum"], d, "datum")); err != nil {
		if vv, ok := fortiAPIPatch(o["datum"], "SwitchControllerLocationCoordinates-Datum"); ok {
			if err = d.Set("datum", vv); err != nil {
				return fmt.Errorf("Error reading datum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading datum: %v", err)
		}
	}

	if err = d.Set("latitude", flattenSwitchControllerLocationCoordinatesLatitude2edl(o["latitude"], d, "latitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["latitude"], "SwitchControllerLocationCoordinates-Latitude"); ok {
			if err = d.Set("latitude", vv); err != nil {
				return fmt.Errorf("Error reading latitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latitude: %v", err)
		}
	}

	if err = d.Set("longitude", flattenSwitchControllerLocationCoordinatesLongitude2edl(o["longitude"], d, "longitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["longitude"], "SwitchControllerLocationCoordinates-Longitude"); ok {
			if err = d.Set("longitude", vv); err != nil {
				return fmt.Errorf("Error reading longitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading longitude: %v", err)
		}
	}

	if err = d.Set("parent_key", flattenSwitchControllerLocationCoordinatesParentKey2edl(o["parent-key"], d, "parent_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["parent-key"], "SwitchControllerLocationCoordinates-ParentKey"); ok {
			if err = d.Set("parent_key", vv); err != nil {
				return fmt.Errorf("Error reading parent_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading parent_key: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerLocationCoordinatesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerLocationCoordinatesAltitude2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationCoordinatesAltitudeUnit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationCoordinatesDatum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationCoordinatesLatitude2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationCoordinatesLongitude2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLocationCoordinatesParentKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerLocationCoordinates(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("altitude"); ok || d.HasChange("altitude") {
		t, err := expandSwitchControllerLocationCoordinatesAltitude2edl(d, v, "altitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["altitude"] = t
		}
	}

	if v, ok := d.GetOk("altitude_unit"); ok || d.HasChange("altitude_unit") {
		t, err := expandSwitchControllerLocationCoordinatesAltitudeUnit2edl(d, v, "altitude_unit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["altitude-unit"] = t
		}
	}

	if v, ok := d.GetOk("datum"); ok || d.HasChange("datum") {
		t, err := expandSwitchControllerLocationCoordinatesDatum2edl(d, v, "datum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["datum"] = t
		}
	}

	if v, ok := d.GetOk("latitude"); ok || d.HasChange("latitude") {
		t, err := expandSwitchControllerLocationCoordinatesLatitude2edl(d, v, "latitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latitude"] = t
		}
	}

	if v, ok := d.GetOk("longitude"); ok || d.HasChange("longitude") {
		t, err := expandSwitchControllerLocationCoordinatesLongitude2edl(d, v, "longitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["longitude"] = t
		}
	}

	if v, ok := d.GetOk("parent_key"); ok || d.HasChange("parent_key") {
		t, err := expandSwitchControllerLocationCoordinatesParentKey2edl(d, v, "parent_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parent-key"] = t
		}
	}

	return &obj, nil
}
