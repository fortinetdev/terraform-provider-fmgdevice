// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard Web Filter local risk score.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterFtgdLocalRisk() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterFtgdLocalRiskCreate,
		Read:   resourceWebfilterFtgdLocalRiskRead,
		Update: resourceWebfilterFtgdLocalRiskUpdate,
		Delete: resourceWebfilterFtgdLocalRiskDelete,

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"risk_score": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceWebfilterFtgdLocalRiskCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWebfilterFtgdLocalRisk(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdLocalRisk resource while getting object: %v", err)
	}

	_, err = c.CreateWebfilterFtgdLocalRisk(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdLocalRisk resource: %v", err)
	}

	d.SetId(getStringKey(d, "url"))

	return resourceWebfilterFtgdLocalRiskRead(d, m)
}

func resourceWebfilterFtgdLocalRiskUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWebfilterFtgdLocalRisk(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdLocalRisk resource while getting object: %v", err)
	}

	_, err = c.UpdateWebfilterFtgdLocalRisk(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdLocalRisk resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "url"))

	return resourceWebfilterFtgdLocalRiskRead(d, m)
}

func resourceWebfilterFtgdLocalRiskDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWebfilterFtgdLocalRisk(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterFtgdLocalRisk resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterFtgdLocalRiskRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterFtgdLocalRisk(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdLocalRisk resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterFtgdLocalRisk(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdLocalRisk resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterFtgdLocalRiskComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFtgdLocalRiskRiskScore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFtgdLocalRiskStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFtgdLocalRiskUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterFtgdLocalRisk(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenWebfilterFtgdLocalRiskComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WebfilterFtgdLocalRisk-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("risk_score", flattenWebfilterFtgdLocalRiskRiskScore(o["risk-score"], d, "risk_score")); err != nil {
		if vv, ok := fortiAPIPatch(o["risk-score"], "WebfilterFtgdLocalRisk-RiskScore"); ok {
			if err = d.Set("risk_score", vv); err != nil {
				return fmt.Errorf("Error reading risk_score: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading risk_score: %v", err)
		}
	}

	if err = d.Set("status", flattenWebfilterFtgdLocalRiskStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WebfilterFtgdLocalRisk-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("url", flattenWebfilterFtgdLocalRiskUrl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "WebfilterFtgdLocalRisk-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func flattenWebfilterFtgdLocalRiskFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterFtgdLocalRiskComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdLocalRiskRiskScore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdLocalRiskStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdLocalRiskUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterFtgdLocalRisk(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWebfilterFtgdLocalRiskComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("risk_score"); ok || d.HasChange("risk_score") {
		t, err := expandWebfilterFtgdLocalRiskRiskScore(d, v, "risk_score")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["risk-score"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWebfilterFtgdLocalRiskStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandWebfilterFtgdLocalRiskUrl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
