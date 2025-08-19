// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Custom field name for CEF format logging.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogAzureSecurityCenterSettingCustomFieldName() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogAzureSecurityCenterSettingCustomFieldNameCreate,
		Read:   resourceLogAzureSecurityCenterSettingCustomFieldNameRead,
		Update: resourceLogAzureSecurityCenterSettingCustomFieldNameUpdate,
		Delete: resourceLogAzureSecurityCenterSettingCustomFieldNameDelete,

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
			"custom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceLogAzureSecurityCenterSettingCustomFieldNameCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectLogAzureSecurityCenterSettingCustomFieldName(d)
	if err != nil {
		return fmt.Errorf("Error creating LogAzureSecurityCenterSettingCustomFieldName resource while getting object: %v", err)
	}

	_, err = c.CreateLogAzureSecurityCenterSettingCustomFieldName(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating LogAzureSecurityCenterSettingCustomFieldName resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogAzureSecurityCenterSettingCustomFieldNameRead(d, m)
}

func resourceLogAzureSecurityCenterSettingCustomFieldNameUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectLogAzureSecurityCenterSettingCustomFieldName(d)
	if err != nil {
		return fmt.Errorf("Error updating LogAzureSecurityCenterSettingCustomFieldName resource while getting object: %v", err)
	}

	_, err = c.UpdateLogAzureSecurityCenterSettingCustomFieldName(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogAzureSecurityCenterSettingCustomFieldName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogAzureSecurityCenterSettingCustomFieldNameRead(d, m)
}

func resourceLogAzureSecurityCenterSettingCustomFieldNameDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteLogAzureSecurityCenterSettingCustomFieldName(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogAzureSecurityCenterSettingCustomFieldName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogAzureSecurityCenterSettingCustomFieldNameRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogAzureSecurityCenterSettingCustomFieldName(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogAzureSecurityCenterSettingCustomFieldName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogAzureSecurityCenterSettingCustomFieldName(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogAzureSecurityCenterSettingCustomFieldName resource from API: %v", err)
	}
	return nil
}

func flattenLogAzureSecurityCenterSettingCustomFieldNameCustom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogAzureSecurityCenterSettingCustomFieldNameId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogAzureSecurityCenterSettingCustomFieldNameName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogAzureSecurityCenterSettingCustomFieldName(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("custom", flattenLogAzureSecurityCenterSettingCustomFieldNameCustom2edl(o["custom"], d, "custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom"], "LogAzureSecurityCenterSettingCustomFieldName-Custom"); ok {
			if err = d.Set("custom", vv); err != nil {
				return fmt.Errorf("Error reading custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom: %v", err)
		}
	}

	if err = d.Set("fosid", flattenLogAzureSecurityCenterSettingCustomFieldNameId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "LogAzureSecurityCenterSettingCustomFieldName-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenLogAzureSecurityCenterSettingCustomFieldNameName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "LogAzureSecurityCenterSettingCustomFieldName-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenLogAzureSecurityCenterSettingCustomFieldNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogAzureSecurityCenterSettingCustomFieldNameCustom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogAzureSecurityCenterSettingCustomFieldNameId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogAzureSecurityCenterSettingCustomFieldNameName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogAzureSecurityCenterSettingCustomFieldName(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("custom"); ok || d.HasChange("custom") {
		t, err := expandLogAzureSecurityCenterSettingCustomFieldNameCustom2edl(d, v, "custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandLogAzureSecurityCenterSettingCustomFieldNameId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandLogAzureSecurityCenterSettingCustomFieldNameName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
