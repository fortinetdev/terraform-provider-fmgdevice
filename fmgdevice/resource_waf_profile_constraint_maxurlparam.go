// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Maximum number of parameters in URL.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfileConstraintMaxUrlParam() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileConstraintMaxUrlParamUpdate,
		Read:   resourceWafProfileConstraintMaxUrlParamRead,
		Update: resourceWafProfileConstraintMaxUrlParamUpdate,
		Delete: resourceWafProfileConstraintMaxUrlParamDelete,

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
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_url_param": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceWafProfileConstraintMaxUrlParamUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWafProfileConstraintMaxUrlParam(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileConstraintMaxUrlParam resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfileConstraintMaxUrlParam(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileConstraintMaxUrlParam resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WafProfileConstraintMaxUrlParam")

	return resourceWafProfileConstraintMaxUrlParamRead(d, m)
}

func resourceWafProfileConstraintMaxUrlParamDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWafProfileConstraintMaxUrlParam(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfileConstraintMaxUrlParam resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileConstraintMaxUrlParamRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWafProfileConstraintMaxUrlParam(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfileConstraintMaxUrlParam resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfileConstraintMaxUrlParam(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfileConstraintMaxUrlParam resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileConstraintMaxUrlParamAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamLog3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamMaxUrlParam3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamSeverity3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWafProfileConstraintMaxUrlParam(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenWafProfileConstraintMaxUrlParamAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "WafProfileConstraintMaxUrlParam-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("log", flattenWafProfileConstraintMaxUrlParamLog3rdl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "WafProfileConstraintMaxUrlParam-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("max_url_param", flattenWafProfileConstraintMaxUrlParamMaxUrlParam3rdl(o["max-url-param"], d, "max_url_param")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-url-param"], "WafProfileConstraintMaxUrlParam-MaxUrlParam"); ok {
			if err = d.Set("max_url_param", vv); err != nil {
				return fmt.Errorf("Error reading max_url_param: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_url_param: %v", err)
		}
	}

	if err = d.Set("severity", flattenWafProfileConstraintMaxUrlParamSeverity3rdl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "WafProfileConstraintMaxUrlParam-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenWafProfileConstraintMaxUrlParamStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WafProfileConstraintMaxUrlParam-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenWafProfileConstraintMaxUrlParamFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileConstraintMaxUrlParamAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamLog3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamMaxUrlParam3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamSeverity3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWafProfileConstraintMaxUrlParam(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandWafProfileConstraintMaxUrlParamAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandWafProfileConstraintMaxUrlParamLog3rdl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("max_url_param"); ok || d.HasChange("max_url_param") {
		t, err := expandWafProfileConstraintMaxUrlParamMaxUrlParam3rdl(d, v, "max_url_param")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandWafProfileConstraintMaxUrlParamSeverity3rdl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWafProfileConstraintMaxUrlParamStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
