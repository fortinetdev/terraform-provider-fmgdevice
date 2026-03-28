// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Main signature class.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfileSignatureMainClass() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileSignatureMainClassUpdate,
		Read:   resourceWafProfileSignatureMainClassRead,
		Update: resourceWafProfileSignatureMainClassUpdate,
		Delete: resourceWafProfileSignatureMainClassDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceWafProfileSignatureMainClassUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWafProfileSignatureMainClass(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileSignatureMainClass resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfileSignatureMainClass(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileSignatureMainClass resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WafProfileSignatureMainClass")

	return resourceWafProfileSignatureMainClassRead(d, m)
}

func resourceWafProfileSignatureMainClassDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWafProfileSignatureMainClass(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfileSignatureMainClass resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileSignatureMainClassRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWafProfileSignatureMainClass(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfileSignatureMainClass resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfileSignatureMainClass(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfileSignatureMainClass resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileSignatureMainClassAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenWafProfileSignatureMainClassLog3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassSeverity3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWafProfileSignatureMainClass(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenWafProfileSignatureMainClassAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "WafProfileSignatureMainClass-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWafProfileSignatureMainClassId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WafProfileSignatureMainClass-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("log", flattenWafProfileSignatureMainClassLog3rdl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "WafProfileSignatureMainClass-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("severity", flattenWafProfileSignatureMainClassSeverity3rdl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "WafProfileSignatureMainClass-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenWafProfileSignatureMainClassStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WafProfileSignatureMainClass-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenWafProfileSignatureMainClassFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileSignatureMainClassAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassLog3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassSeverity3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWafProfileSignatureMainClass(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandWafProfileSignatureMainClassAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWafProfileSignatureMainClassId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandWafProfileSignatureMainClassLog3rdl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandWafProfileSignatureMainClassSeverity3rdl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWafProfileSignatureMainClassStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
