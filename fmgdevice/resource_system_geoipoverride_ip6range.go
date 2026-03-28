// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Table of IPv6 ranges assigned to country.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemGeoipOverrideIp6Range() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGeoipOverrideIp6RangeCreate,
		Read:   resourceSystemGeoipOverrideIp6RangeRead,
		Update: resourceSystemGeoipOverrideIp6RangeUpdate,
		Delete: resourceSystemGeoipOverrideIp6RangeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"geoip_override": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemGeoipOverrideIp6RangeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	geoip_override := d.Get("geoip_override").(string)
	paradict["device"] = device_name
	paradict["geoip_override"] = geoip_override

	obj, err := getObjectSystemGeoipOverrideIp6Range(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemGeoipOverrideIp6Range resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemGeoipOverrideIp6Range(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemGeoipOverrideIp6Range(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemGeoipOverrideIp6Range resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateSystemGeoipOverrideIp6Range(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemGeoipOverrideIp6Range resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceSystemGeoipOverrideIp6RangeRead(d, m)
			} else {
				return fmt.Errorf("Error creating SystemGeoipOverrideIp6Range resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemGeoipOverrideIp6RangeRead(d, m)
}

func resourceSystemGeoipOverrideIp6RangeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	geoip_override := d.Get("geoip_override").(string)
	paradict["device"] = device_name
	paradict["geoip_override"] = geoip_override

	obj, err := getObjectSystemGeoipOverrideIp6Range(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemGeoipOverrideIp6Range resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemGeoipOverrideIp6Range(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemGeoipOverrideIp6Range resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemGeoipOverrideIp6RangeRead(d, m)
}

func resourceSystemGeoipOverrideIp6RangeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	geoip_override := d.Get("geoip_override").(string)
	paradict["device"] = device_name
	paradict["geoip_override"] = geoip_override

	wsParams["adom"] = adomv

	err = c.DeleteSystemGeoipOverrideIp6Range(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemGeoipOverrideIp6Range resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGeoipOverrideIp6RangeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	geoip_override := d.Get("geoip_override").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if geoip_override == "" {
		geoip_override = importOptionChecking(m.(*FortiClient).Cfg, "geoip_override")
		if geoip_override == "" {
			return fmt.Errorf("Parameter geoip_override is missing")
		}
		if err = d.Set("geoip_override", geoip_override); err != nil {
			return fmt.Errorf("Error set params geoip_override: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["geoip_override"] = geoip_override

	o, err := c.ReadSystemGeoipOverrideIp6Range(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemGeoipOverrideIp6Range resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGeoipOverrideIp6Range(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemGeoipOverrideIp6Range resource from API: %v", err)
	}
	return nil
}

func flattenSystemGeoipOverrideIp6RangeEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeoipOverrideIp6RangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeoipOverrideIp6RangeStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemGeoipOverrideIp6Range(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("end_ip", flattenSystemGeoipOverrideIp6RangeEndIp2edl(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "SystemGeoipOverrideIp6Range-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemGeoipOverrideIp6RangeId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemGeoipOverrideIp6Range-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenSystemGeoipOverrideIp6RangeStartIp2edl(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "SystemGeoipOverrideIp6Range-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	return nil
}

func flattenSystemGeoipOverrideIp6RangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemGeoipOverrideIp6RangeEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIp6RangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIp6RangeStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGeoipOverrideIp6Range(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandSystemGeoipOverrideIp6RangeEndIp2edl(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemGeoipOverrideIp6RangeId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandSystemGeoipOverrideIp6RangeStartIp2edl(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	return &obj, nil
}
