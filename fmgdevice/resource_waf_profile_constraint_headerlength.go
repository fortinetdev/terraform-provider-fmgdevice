// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> HTTP header length in request.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfileConstraintHeaderLength() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileConstraintHeaderLengthUpdate,
		Read:   resourceWafProfileConstraintHeaderLengthRead,
		Update: resourceWafProfileConstraintHeaderLengthUpdate,
		Delete: resourceWafProfileConstraintHeaderLengthDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWafProfileConstraintHeaderLengthUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectWafProfileConstraintHeaderLength(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileConstraintHeaderLength resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfileConstraintHeaderLength(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileConstraintHeaderLength resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WafProfileConstraintHeaderLength")

	return resourceWafProfileConstraintHeaderLengthRead(d, m)
}

func resourceWafProfileConstraintHeaderLengthDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteWafProfileConstraintHeaderLength(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfileConstraintHeaderLength resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileConstraintHeaderLengthRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadWafProfileConstraintHeaderLength(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfileConstraintHeaderLength resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfileConstraintHeaderLength(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfileConstraintHeaderLength resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileConstraintHeaderLengthAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthLength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthLog3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthSeverity3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWafProfileConstraintHeaderLength(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenWafProfileConstraintHeaderLengthAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "WafProfileConstraintHeaderLength-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("length", flattenWafProfileConstraintHeaderLengthLength3rdl(o["length"], d, "length")); err != nil {
		if vv, ok := fortiAPIPatch(o["length"], "WafProfileConstraintHeaderLength-Length"); ok {
			if err = d.Set("length", vv); err != nil {
				return fmt.Errorf("Error reading length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading length: %v", err)
		}
	}

	if err = d.Set("log", flattenWafProfileConstraintHeaderLengthLog3rdl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "WafProfileConstraintHeaderLength-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("severity", flattenWafProfileConstraintHeaderLengthSeverity3rdl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "WafProfileConstraintHeaderLength-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenWafProfileConstraintHeaderLengthStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WafProfileConstraintHeaderLength-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenWafProfileConstraintHeaderLengthFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileConstraintHeaderLengthAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthLength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthLog3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthSeverity3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWafProfileConstraintHeaderLength(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandWafProfileConstraintHeaderLengthAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("length"); ok || d.HasChange("length") {
		t, err := expandWafProfileConstraintHeaderLengthLength3rdl(d, v, "length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["length"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandWafProfileConstraintHeaderLengthLog3rdl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandWafProfileConstraintHeaderLengthSeverity3rdl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWafProfileConstraintHeaderLengthStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
