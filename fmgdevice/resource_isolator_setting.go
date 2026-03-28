// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device IsolatorSetting

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIsolatorSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceIsolatorSettingUpdate,
		Read:   resourceIsolatorSettingRead,
		Update: resourceIsolatorSettingUpdate,
		Delete: resourceIsolatorSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"default_isolator_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"defective_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unmatched_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIsolatorSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIsolatorSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating IsolatorSetting resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIsolatorSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IsolatorSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("IsolatorSetting")

	return resourceIsolatorSettingRead(d, m)
}

func resourceIsolatorSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteIsolatorSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IsolatorSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIsolatorSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadIsolatorSetting(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IsolatorSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIsolatorSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IsolatorSetting resource from API: %v", err)
	}
	return nil
}

func flattenIsolatorSettingDefaultIsolatorProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIsolatorSettingDefectiveSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIsolatorSettingUnmatchedSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIsolatorSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("default_isolator_profile", flattenIsolatorSettingDefaultIsolatorProfile(o["default-isolator-profile"], d, "default_isolator_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-isolator-profile"], "IsolatorSetting-DefaultIsolatorProfile"); ok {
			if err = d.Set("default_isolator_profile", vv); err != nil {
				return fmt.Errorf("Error reading default_isolator_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_isolator_profile: %v", err)
		}
	}

	if err = d.Set("defective_session", flattenIsolatorSettingDefectiveSession(o["defective-session"], d, "defective_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["defective-session"], "IsolatorSetting-DefectiveSession"); ok {
			if err = d.Set("defective_session", vv); err != nil {
				return fmt.Errorf("Error reading defective_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading defective_session: %v", err)
		}
	}

	if err = d.Set("unmatched_session", flattenIsolatorSettingUnmatchedSession(o["unmatched-session"], d, "unmatched_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["unmatched-session"], "IsolatorSetting-UnmatchedSession"); ok {
			if err = d.Set("unmatched_session", vv); err != nil {
				return fmt.Errorf("Error reading unmatched_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unmatched_session: %v", err)
		}
	}

	return nil
}

func flattenIsolatorSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIsolatorSettingDefaultIsolatorProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIsolatorSettingDefectiveSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIsolatorSettingUnmatchedSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIsolatorSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_isolator_profile"); ok || d.HasChange("default_isolator_profile") {
		t, err := expandIsolatorSettingDefaultIsolatorProfile(d, v, "default_isolator_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-isolator-profile"] = t
		}
	}

	if v, ok := d.GetOk("defective_session"); ok || d.HasChange("defective_session") {
		t, err := expandIsolatorSettingDefectiveSession(d, v, "defective_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["defective-session"] = t
		}
	}

	if v, ok := d.GetOk("unmatched_session"); ok || d.HasChange("unmatched_session") {
		t, err := expandIsolatorSettingUnmatchedSession(d, v, "unmatched_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unmatched-session"] = t
		}
	}

	return &obj, nil
}
