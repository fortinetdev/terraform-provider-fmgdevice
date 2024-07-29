// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPS URL filter settings for IPv6.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterIpsUrlfilterSetting6() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterIpsUrlfilterSetting6Update,
		Read:   resourceWebfilterIpsUrlfilterSetting6Read,
		Update: resourceWebfilterIpsUrlfilterSetting6Update,
		Delete: resourceWebfilterIpsUrlfilterSetting6Delete,

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
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gateway6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"geo_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWebfilterIpsUrlfilterSetting6Update(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectWebfilterIpsUrlfilterSetting6(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterSetting6 resource while getting object: %v", err)
	}

	_, err = c.UpdateWebfilterIpsUrlfilterSetting6(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterSetting6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WebfilterIpsUrlfilterSetting6")

	return resourceWebfilterIpsUrlfilterSetting6Read(d, m)
}

func resourceWebfilterIpsUrlfilterSetting6Delete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteWebfilterIpsUrlfilterSetting6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterIpsUrlfilterSetting6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterIpsUrlfilterSetting6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadWebfilterIpsUrlfilterSetting6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterIpsUrlfilterSetting6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterIpsUrlfilterSetting6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterIpsUrlfilterSetting6 resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterIpsUrlfilterSetting6Device(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterIpsUrlfilterSetting6Distance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterIpsUrlfilterSetting6Gateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterIpsUrlfilterSetting6GeoFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterIpsUrlfilterSetting6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("device", flattenWebfilterIpsUrlfilterSetting6Device(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "WebfilterIpsUrlfilterSetting6-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("distance", flattenWebfilterIpsUrlfilterSetting6Distance(o["distance"], d, "distance")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance"], "WebfilterIpsUrlfilterSetting6-Distance"); ok {
			if err = d.Set("distance", vv); err != nil {
				return fmt.Errorf("Error reading distance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("gateway6", flattenWebfilterIpsUrlfilterSetting6Gateway6(o["gateway6"], d, "gateway6")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway6"], "WebfilterIpsUrlfilterSetting6-Gateway6"); ok {
			if err = d.Set("gateway6", vv); err != nil {
				return fmt.Errorf("Error reading gateway6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway6: %v", err)
		}
	}

	if err = d.Set("geo_filter", flattenWebfilterIpsUrlfilterSetting6GeoFilter(o["geo-filter"], d, "geo_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["geo-filter"], "WebfilterIpsUrlfilterSetting6-GeoFilter"); ok {
			if err = d.Set("geo_filter", vv); err != nil {
				return fmt.Errorf("Error reading geo_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading geo_filter: %v", err)
		}
	}

	return nil
}

func flattenWebfilterIpsUrlfilterSetting6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterIpsUrlfilterSetting6Device(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterSetting6Distance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterSetting6Gateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterSetting6GeoFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterIpsUrlfilterSetting6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandWebfilterIpsUrlfilterSetting6Device(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok || d.HasChange("distance") {
		t, err := expandWebfilterIpsUrlfilterSetting6Distance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("gateway6"); ok || d.HasChange("gateway6") {
		t, err := expandWebfilterIpsUrlfilterSetting6Gateway6(d, v, "gateway6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway6"] = t
		}
	}

	if v, ok := d.GetOk("geo_filter"); ok || d.HasChange("geo_filter") {
		t, err := expandWebfilterIpsUrlfilterSetting6GeoFilter(d, v, "geo_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geo-filter"] = t
		}
	}

	return &obj, nil
}
