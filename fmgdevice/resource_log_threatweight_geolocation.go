// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Geolocation-based threat weight settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogThreatWeightGeolocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogThreatWeightGeolocationCreate,
		Read:   resourceLogThreatWeightGeolocationRead,
		Update: resourceLogThreatWeightGeolocationUpdate,
		Delete: resourceLogThreatWeightGeolocationDelete,

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogThreatWeightGeolocationCreate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectLogThreatWeightGeolocation(d)
	if err != nil {
		return fmt.Errorf("Error creating LogThreatWeightGeolocation resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadLogThreatWeightGeolocation(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateLogThreatWeightGeolocation(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating LogThreatWeightGeolocation resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateLogThreatWeightGeolocation(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating LogThreatWeightGeolocation resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogThreatWeightGeolocationRead(d, m)
}

func resourceLogThreatWeightGeolocationUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectLogThreatWeightGeolocation(d)
	if err != nil {
		return fmt.Errorf("Error updating LogThreatWeightGeolocation resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLogThreatWeightGeolocation(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogThreatWeightGeolocation resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogThreatWeightGeolocationRead(d, m)
}

func resourceLogThreatWeightGeolocationDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteLogThreatWeightGeolocation(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogThreatWeightGeolocation resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogThreatWeightGeolocationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogThreatWeightGeolocation(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LogThreatWeightGeolocation resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogThreatWeightGeolocation(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogThreatWeightGeolocation resource from API: %v", err)
	}
	return nil
}

func flattenLogThreatWeightGeolocationCountry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogThreatWeightGeolocationId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightGeolocationLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogThreatWeightGeolocation(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("country", flattenLogThreatWeightGeolocationCountry2edl(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "LogThreatWeightGeolocation-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("fosid", flattenLogThreatWeightGeolocationId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "LogThreatWeightGeolocation-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("level", flattenLogThreatWeightGeolocationLevel2edl(o["level"], d, "level")); err != nil {
		if vv, ok := fortiAPIPatch(o["level"], "LogThreatWeightGeolocation-Level"); ok {
			if err = d.Set("level", vv); err != nil {
				return fmt.Errorf("Error reading level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading level: %v", err)
		}
	}

	return nil
}

func flattenLogThreatWeightGeolocationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogThreatWeightGeolocationCountry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogThreatWeightGeolocationId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightGeolocationLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogThreatWeightGeolocation(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("country"); ok || d.HasChange("country") {
		t, err := expandLogThreatWeightGeolocationCountry2edl(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandLogThreatWeightGeolocationId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("level"); ok || d.HasChange("level") {
		t, err := expandLogThreatWeightGeolocationLevel2edl(d, v, "level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["level"] = t
		}
	}

	return &obj, nil
}
