// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Request headers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutomationActionHttpHeaders() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationActionHttpHeadersCreate,
		Read:   resourceSystemAutomationActionHttpHeadersRead,
		Update: resourceSystemAutomationActionHttpHeadersUpdate,
		Delete: resourceSystemAutomationActionHttpHeadersDelete,

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
			"automation_action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAutomationActionHttpHeadersCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	automation_action := d.Get("automation_action").(string)
	paradict["device"] = device_name
	paradict["automation_action"] = automation_action

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemAutomationActionHttpHeaders(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationActionHttpHeaders resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAutomationActionHttpHeaders(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationActionHttpHeaders resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAutomationActionHttpHeadersRead(d, m)
}

func resourceSystemAutomationActionHttpHeadersUpdate(d *schema.ResourceData, m interface{}) error {
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
	automation_action := d.Get("automation_action").(string)
	paradict["device"] = device_name
	paradict["automation_action"] = automation_action

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemAutomationActionHttpHeaders(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationActionHttpHeaders resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutomationActionHttpHeaders(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationActionHttpHeaders resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAutomationActionHttpHeadersRead(d, m)
}

func resourceSystemAutomationActionHttpHeadersDelete(d *schema.ResourceData, m interface{}) error {
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
	automation_action := d.Get("automation_action").(string)
	paradict["device"] = device_name
	paradict["automation_action"] = automation_action

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemAutomationActionHttpHeaders(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationActionHttpHeaders resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationActionHttpHeadersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	automation_action := d.Get("automation_action").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if automation_action == "" {
		automation_action = importOptionChecking(m.(*FortiClient).Cfg, "automation_action")
		if automation_action == "" {
			return fmt.Errorf("Parameter automation_action is missing")
		}
		if err = d.Set("automation_action", automation_action); err != nil {
			return fmt.Errorf("Error set params automation_action: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["automation_action"] = automation_action

	o, err := c.ReadSystemAutomationActionHttpHeaders(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationActionHttpHeaders resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationActionHttpHeaders(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationActionHttpHeaders resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationActionHttpHeadersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionHttpHeadersKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionHttpHeadersValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutomationActionHttpHeaders(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemAutomationActionHttpHeadersId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemAutomationActionHttpHeaders-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("key", flattenSystemAutomationActionHttpHeadersKey2edl(o["key"], d, "key")); err != nil {
		if vv, ok := fortiAPIPatch(o["key"], "SystemAutomationActionHttpHeaders-Key"); ok {
			if err = d.Set("key", vv); err != nil {
				return fmt.Errorf("Error reading key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}

	if err = d.Set("value", flattenSystemAutomationActionHttpHeadersValue2edl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "SystemAutomationActionHttpHeaders-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationActionHttpHeadersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutomationActionHttpHeadersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHttpHeadersKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHttpHeadersValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationActionHttpHeaders(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemAutomationActionHttpHeadersId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandSystemAutomationActionHttpHeadersKey2edl(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandSystemAutomationActionHttpHeadersValue2edl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
