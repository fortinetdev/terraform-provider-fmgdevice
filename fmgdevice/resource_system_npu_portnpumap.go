// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure port to NPU group mapping.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNpuPortNpuMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNpuPortNpuMapCreate,
		Read:   resourceSystemNpuPortNpuMapRead,
		Update: resourceSystemNpuPortNpuMapUpdate,
		Delete: resourceSystemNpuPortNpuMapDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"npu_group_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemNpuPortNpuMapCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemNpuPortNpuMap(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemNpuPortNpuMap resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("interface")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemNpuPortNpuMap(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemNpuPortNpuMap(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemNpuPortNpuMap resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemNpuPortNpuMap(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemNpuPortNpuMap resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "interface"))

	return resourceSystemNpuPortNpuMapRead(d, m)
}

func resourceSystemNpuPortNpuMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemNpuPortNpuMap(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNpuPortNpuMap resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemNpuPortNpuMap(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemNpuPortNpuMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "interface"))

	return resourceSystemNpuPortNpuMapRead(d, m)
}

func resourceSystemNpuPortNpuMapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemNpuPortNpuMap(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNpuPortNpuMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNpuPortNpuMapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemNpuPortNpuMap(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemNpuPortNpuMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNpuPortNpuMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNpuPortNpuMap resource from API: %v", err)
	}
	return nil
}

func flattenSystemNpuPortNpuMapInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuPortNpuMapNpuGroupIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNpuPortNpuMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("interface", flattenSystemNpuPortNpuMapInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemNpuPortNpuMap-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("npu_group_index", flattenSystemNpuPortNpuMapNpuGroupIndex2edl(o["npu-group-index"], d, "npu_group_index")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-group-index"], "SystemNpuPortNpuMap-NpuGroupIndex"); ok {
			if err = d.Set("npu_group_index", vv); err != nil {
				return fmt.Errorf("Error reading npu_group_index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_group_index: %v", err)
		}
	}

	return nil
}

func flattenSystemNpuPortNpuMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNpuPortNpuMapInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuPortNpuMapNpuGroupIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNpuPortNpuMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemNpuPortNpuMapInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("npu_group_index"); ok || d.HasChange("npu_group_index") {
		t, err := expandSystemNpuPortNpuMapNpuGroupIndex2edl(d, v, "npu_group_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-group-index"] = t
		}
	}

	return &obj, nil
}
