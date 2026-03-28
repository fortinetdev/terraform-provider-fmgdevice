// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard - AntiSpam.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterFortishield() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterFortishieldUpdate,
		Read:   resourceEmailfilterFortishieldRead,
		Update: resourceEmailfilterFortishieldUpdate,
		Delete: resourceEmailfilterFortishieldDelete,

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
			"spam_submit_force": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_submit_srv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_submit_txt2htm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceEmailfilterFortishieldUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEmailfilterFortishield(d)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterFortishield resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEmailfilterFortishield(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterFortishield resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("EmailfilterFortishield")

	return resourceEmailfilterFortishieldRead(d, m)
}

func resourceEmailfilterFortishieldDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteEmailfilterFortishield(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterFortishield resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterFortishieldRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadEmailfilterFortishield(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EmailfilterFortishield resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterFortishield(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterFortishield resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterFortishieldSpamSubmitForce(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterFortishieldSpamSubmitSrv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterFortishieldSpamSubmitTxt2Htm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEmailfilterFortishield(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("spam_submit_force", flattenEmailfilterFortishieldSpamSubmitForce(o["spam-submit-force"], d, "spam_submit_force")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-submit-force"], "EmailfilterFortishield-SpamSubmitForce"); ok {
			if err = d.Set("spam_submit_force", vv); err != nil {
				return fmt.Errorf("Error reading spam_submit_force: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_submit_force: %v", err)
		}
	}

	if err = d.Set("spam_submit_srv", flattenEmailfilterFortishieldSpamSubmitSrv(o["spam-submit-srv"], d, "spam_submit_srv")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-submit-srv"], "EmailfilterFortishield-SpamSubmitSrv"); ok {
			if err = d.Set("spam_submit_srv", vv); err != nil {
				return fmt.Errorf("Error reading spam_submit_srv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_submit_srv: %v", err)
		}
	}

	if err = d.Set("spam_submit_txt2htm", flattenEmailfilterFortishieldSpamSubmitTxt2Htm(o["spam-submit-txt2htm"], d, "spam_submit_txt2htm")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-submit-txt2htm"], "EmailfilterFortishield-SpamSubmitTxt2Htm"); ok {
			if err = d.Set("spam_submit_txt2htm", vv); err != nil {
				return fmt.Errorf("Error reading spam_submit_txt2htm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_submit_txt2htm: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterFortishieldFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEmailfilterFortishieldSpamSubmitForce(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterFortishieldSpamSubmitSrv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterFortishieldSpamSubmitTxt2Htm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterFortishield(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("spam_submit_force"); ok || d.HasChange("spam_submit_force") {
		t, err := expandEmailfilterFortishieldSpamSubmitForce(d, v, "spam_submit_force")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-submit-force"] = t
		}
	}

	if v, ok := d.GetOk("spam_submit_srv"); ok || d.HasChange("spam_submit_srv") {
		t, err := expandEmailfilterFortishieldSpamSubmitSrv(d, v, "spam_submit_srv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-submit-srv"] = t
		}
	}

	if v, ok := d.GetOk("spam_submit_txt2htm"); ok || d.HasChange("spam_submit_txt2htm") {
		t, err := expandEmailfilterFortishieldSpamSubmitTxt2Htm(d, v, "spam_submit_txt2htm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-submit-txt2htm"] = t
		}
	}

	return &obj, nil
}
