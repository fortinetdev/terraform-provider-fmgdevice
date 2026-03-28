// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NPU DSW DTS profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNpuDswDtsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNpuDswDtsProfileCreate,
		Read:   resourceSystemNpuDswDtsProfileRead,
		Update: resourceSystemNpuDswDtsProfileUpdate,
		Delete: resourceSystemNpuDswDtsProfileDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"min_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"profile_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"step": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemNpuDswDtsProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemNpuDswDtsProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemNpuDswDtsProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("profile_id")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemNpuDswDtsProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemNpuDswDtsProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemNpuDswDtsProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemNpuDswDtsProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemNpuDswDtsProfile resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "profile_id")))

	return resourceSystemNpuDswDtsProfileRead(d, m)
}

func resourceSystemNpuDswDtsProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemNpuDswDtsProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNpuDswDtsProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemNpuDswDtsProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemNpuDswDtsProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "profile_id")))

	return resourceSystemNpuDswDtsProfileRead(d, m)
}

func resourceSystemNpuDswDtsProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemNpuDswDtsProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNpuDswDtsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNpuDswDtsProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemNpuDswDtsProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemNpuDswDtsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNpuDswDtsProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNpuDswDtsProfile resource from API: %v", err)
	}
	return nil
}

func flattenSystemNpuDswDtsProfileAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuDswDtsProfileMinLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuDswDtsProfileProfileId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuDswDtsProfileStep2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNpuDswDtsProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenSystemNpuDswDtsProfileAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SystemNpuDswDtsProfile-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("min_limit", flattenSystemNpuDswDtsProfileMinLimit2edl(o["min-limit"], d, "min_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-limit"], "SystemNpuDswDtsProfile-MinLimit"); ok {
			if err = d.Set("min_limit", vv); err != nil {
				return fmt.Errorf("Error reading min_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_limit: %v", err)
		}
	}

	if err = d.Set("profile_id", flattenSystemNpuDswDtsProfileProfileId2edl(o["profile-id"], d, "profile_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-id"], "SystemNpuDswDtsProfile-ProfileId"); ok {
			if err = d.Set("profile_id", vv); err != nil {
				return fmt.Errorf("Error reading profile_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_id: %v", err)
		}
	}

	if err = d.Set("step", flattenSystemNpuDswDtsProfileStep2edl(o["step"], d, "step")); err != nil {
		if vv, ok := fortiAPIPatch(o["step"], "SystemNpuDswDtsProfile-Step"); ok {
			if err = d.Set("step", vv); err != nil {
				return fmt.Errorf("Error reading step: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading step: %v", err)
		}
	}

	return nil
}

func flattenSystemNpuDswDtsProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNpuDswDtsProfileAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDswDtsProfileMinLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDswDtsProfileProfileId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDswDtsProfileStep2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNpuDswDtsProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSystemNpuDswDtsProfileAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("min_limit"); ok || d.HasChange("min_limit") {
		t, err := expandSystemNpuDswDtsProfileMinLimit2edl(d, v, "min_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-limit"] = t
		}
	}

	if v, ok := d.GetOk("profile_id"); ok || d.HasChange("profile_id") {
		t, err := expandSystemNpuDswDtsProfileProfileId2edl(d, v, "profile_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-id"] = t
		}
	}

	if v, ok := d.GetOk("step"); ok || d.HasChange("step") {
		t, err := expandSystemNpuDswDtsProfileStep2edl(d, v, "step")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["step"] = t
		}
	}

	return &obj, nil
}
