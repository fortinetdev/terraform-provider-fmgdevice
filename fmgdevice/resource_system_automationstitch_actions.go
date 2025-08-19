// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure stitch actions.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutomationStitchActions() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationStitchActionsCreate,
		Read:   resourceSystemAutomationStitchActionsRead,
		Update: resourceSystemAutomationStitchActionsUpdate,
		Delete: resourceSystemAutomationStitchActionsDelete,

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
			"automation_stitch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"required": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAutomationStitchActionsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	automation_stitch := d.Get("automation_stitch").(string)
	paradict["device"] = device_name
	paradict["automation_stitch"] = automation_stitch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemAutomationStitchActions(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationStitchActions resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAutomationStitchActions(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationStitchActions resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAutomationStitchActionsRead(d, m)
}

func resourceSystemAutomationStitchActionsUpdate(d *schema.ResourceData, m interface{}) error {
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
	automation_stitch := d.Get("automation_stitch").(string)
	paradict["device"] = device_name
	paradict["automation_stitch"] = automation_stitch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemAutomationStitchActions(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationStitchActions resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutomationStitchActions(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationStitchActions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAutomationStitchActionsRead(d, m)
}

func resourceSystemAutomationStitchActionsDelete(d *schema.ResourceData, m interface{}) error {
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
	automation_stitch := d.Get("automation_stitch").(string)
	paradict["device"] = device_name
	paradict["automation_stitch"] = automation_stitch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemAutomationStitchActions(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationStitchActions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationStitchActionsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	automation_stitch := d.Get("automation_stitch").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if automation_stitch == "" {
		automation_stitch = importOptionChecking(m.(*FortiClient).Cfg, "automation_stitch")
		if automation_stitch == "" {
			return fmt.Errorf("Parameter automation_stitch is missing")
		}
		if err = d.Set("automation_stitch", automation_stitch); err != nil {
			return fmt.Errorf("Error set params automation_stitch: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["automation_stitch"] = automation_stitch

	o, err := c.ReadSystemAutomationStitchActions(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationStitchActions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationStitchActions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationStitchActions resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationStitchActionsAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationStitchActionsDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationStitchActionsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationStitchActionsRequired2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutomationStitchActions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenSystemAutomationStitchActionsAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SystemAutomationStitchActions-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("delay", flattenSystemAutomationStitchActionsDelay2edl(o["delay"], d, "delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["delay"], "SystemAutomationStitchActions-Delay"); ok {
			if err = d.Set("delay", vv); err != nil {
				return fmt.Errorf("Error reading delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delay: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemAutomationStitchActionsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemAutomationStitchActions-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("required", flattenSystemAutomationStitchActionsRequired2edl(o["required"], d, "required")); err != nil {
		if vv, ok := fortiAPIPatch(o["required"], "SystemAutomationStitchActions-Required"); ok {
			if err = d.Set("required", vv); err != nil {
				return fmt.Errorf("Error reading required: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading required: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationStitchActionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutomationStitchActionsAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationStitchActionsDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchActionsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchActionsRequired2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationStitchActions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSystemAutomationStitchActionsAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("delay"); ok || d.HasChange("delay") {
		t, err := expandSystemAutomationStitchActionsDelay2edl(d, v, "delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delay"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemAutomationStitchActionsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("required"); ok || d.HasChange("required") {
		t, err := expandSystemAutomationStitchActionsRequired2edl(d, v, "required")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["required"] = t
		}
	}

	return &obj, nil
}
