// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPS URL filter settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterIpsUrlfilterSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterIpsUrlfilterSettingUpdate,
		Read:   resourceWebfilterIpsUrlfilterSettingRead,
		Update: resourceWebfilterIpsUrlfilterSettingUpdate,
		Delete: resourceWebfilterIpsUrlfilterSettingDelete,

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
			"gateway": &schema.Schema{
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

func resourceWebfilterIpsUrlfilterSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterIpsUrlfilterSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateWebfilterIpsUrlfilterSetting(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WebfilterIpsUrlfilterSetting")

	return resourceWebfilterIpsUrlfilterSettingRead(d, m)
}

func resourceWebfilterIpsUrlfilterSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebfilterIpsUrlfilterSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterIpsUrlfilterSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterIpsUrlfilterSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterIpsUrlfilterSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterIpsUrlfilterSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterIpsUrlfilterSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterIpsUrlfilterSetting resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterIpsUrlfilterSettingDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWebfilterIpsUrlfilterSettingDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterIpsUrlfilterSettingGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterIpsUrlfilterSettingGeoFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterIpsUrlfilterSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("device", flattenWebfilterIpsUrlfilterSettingDevice(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "WebfilterIpsUrlfilterSetting-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("distance", flattenWebfilterIpsUrlfilterSettingDistance(o["distance"], d, "distance")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance"], "WebfilterIpsUrlfilterSetting-Distance"); ok {
			if err = d.Set("distance", vv); err != nil {
				return fmt.Errorf("Error reading distance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("gateway", flattenWebfilterIpsUrlfilterSettingGateway(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "WebfilterIpsUrlfilterSetting-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("geo_filter", flattenWebfilterIpsUrlfilterSettingGeoFilter(o["geo-filter"], d, "geo_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["geo-filter"], "WebfilterIpsUrlfilterSetting-GeoFilter"); ok {
			if err = d.Set("geo_filter", vv); err != nil {
				return fmt.Errorf("Error reading geo_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading geo_filter: %v", err)
		}
	}

	return nil
}

func flattenWebfilterIpsUrlfilterSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterIpsUrlfilterSettingDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWebfilterIpsUrlfilterSettingDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterSettingGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterSettingGeoFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterIpsUrlfilterSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandWebfilterIpsUrlfilterSettingDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok || d.HasChange("distance") {
		t, err := expandWebfilterIpsUrlfilterSettingDistance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandWebfilterIpsUrlfilterSettingGateway(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("geo_filter"); ok || d.HasChange("geo_filter") {
		t, err := expandWebfilterIpsUrlfilterSettingGeoFilter(d, v, "geo_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geo-filter"] = t
		}
	}

	return &obj, nil
}
