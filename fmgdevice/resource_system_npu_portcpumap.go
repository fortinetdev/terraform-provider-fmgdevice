// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NPU interface to CPU core mapping.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNpuPortCpuMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNpuPortCpuMapCreate,
		Read:   resourceSystemNpuPortCpuMapRead,
		Update: resourceSystemNpuPortCpuMapUpdate,
		Delete: resourceSystemNpuPortCpuMapDelete,

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
			"cpu_core": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemNpuPortCpuMapCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemNpuPortCpuMap(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemNpuPortCpuMap resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("interface")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemNpuPortCpuMap(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemNpuPortCpuMap(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemNpuPortCpuMap resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemNpuPortCpuMap(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemNpuPortCpuMap resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "interface"))

	return resourceSystemNpuPortCpuMapRead(d, m)
}

func resourceSystemNpuPortCpuMapUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemNpuPortCpuMap(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNpuPortCpuMap resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemNpuPortCpuMap(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemNpuPortCpuMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "interface"))

	return resourceSystemNpuPortCpuMapRead(d, m)
}

func resourceSystemNpuPortCpuMapDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteSystemNpuPortCpuMap(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNpuPortCpuMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNpuPortCpuMapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemNpuPortCpuMap(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemNpuPortCpuMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNpuPortCpuMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNpuPortCpuMap resource from API: %v", err)
	}
	return nil
}

func flattenSystemNpuPortCpuMapCpuCore2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuPortCpuMapInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNpuPortCpuMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cpu_core", flattenSystemNpuPortCpuMapCpuCore2edl(o["cpu-core"], d, "cpu_core")); err != nil {
		if vv, ok := fortiAPIPatch(o["cpu-core"], "SystemNpuPortCpuMap-CpuCore"); ok {
			if err = d.Set("cpu_core", vv); err != nil {
				return fmt.Errorf("Error reading cpu_core: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cpu_core: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemNpuPortCpuMapInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemNpuPortCpuMap-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func flattenSystemNpuPortCpuMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNpuPortCpuMapCpuCore2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuPortCpuMapInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNpuPortCpuMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cpu_core"); ok || d.HasChange("cpu_core") {
		t, err := expandSystemNpuPortCpuMapCpuCore2edl(d, v, "cpu_core")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu-core"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemNpuPortCpuMapInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	return &obj, nil
}
