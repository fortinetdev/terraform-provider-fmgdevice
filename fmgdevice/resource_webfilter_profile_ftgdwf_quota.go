// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> FortiGuard traffic quota settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterProfileFtgdWfQuota() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterProfileFtgdWfQuotaCreate,
		Read:   resourceWebfilterProfileFtgdWfQuotaRead,
		Update: resourceWebfilterProfileFtgdWfQuotaUpdate,
		Delete: resourceWebfilterProfileFtgdWfQuotaDelete,

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
			"category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"duration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"override_replacemsg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceWebfilterProfileFtgdWfQuotaCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterProfileFtgdWfQuota(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterProfileFtgdWfQuota resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebfilterProfileFtgdWfQuota(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebfilterProfileFtgdWfQuota(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebfilterProfileFtgdWfQuota resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebfilterProfileFtgdWfQuota(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebfilterProfileFtgdWfQuota resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebfilterProfileFtgdWfQuotaRead(d, m)
}

func resourceWebfilterProfileFtgdWfQuotaUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterProfileFtgdWfQuota(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfileFtgdWfQuota resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterProfileFtgdWfQuota(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfileFtgdWfQuota resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebfilterProfileFtgdWfQuotaRead(d, m)
}

func resourceWebfilterProfileFtgdWfQuotaDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebfilterProfileFtgdWfQuota(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterProfileFtgdWfQuota resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterProfileFtgdWfQuotaRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterProfileFtgdWfQuota(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterProfileFtgdWfQuota resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterProfileFtgdWfQuota(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterProfileFtgdWfQuota resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterProfileFtgdWfQuotaCategory3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfQuotaDuration3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaOverrideReplacemsg3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaUnit3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterProfileFtgdWfQuota(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("category", flattenWebfilterProfileFtgdWfQuotaCategory3rdl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "WebfilterProfileFtgdWfQuota-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("duration", flattenWebfilterProfileFtgdWfQuotaDuration3rdl(o["duration"], d, "duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["duration"], "WebfilterProfileFtgdWfQuota-Duration"); ok {
			if err = d.Set("duration", vv); err != nil {
				return fmt.Errorf("Error reading duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duration: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWebfilterProfileFtgdWfQuotaId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WebfilterProfileFtgdWfQuota-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("override_replacemsg", flattenWebfilterProfileFtgdWfQuotaOverrideReplacemsg3rdl(o["override-replacemsg"], d, "override_replacemsg")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-replacemsg"], "WebfilterProfileFtgdWfQuota-OverrideReplacemsg"); ok {
			if err = d.Set("override_replacemsg", vv); err != nil {
				return fmt.Errorf("Error reading override_replacemsg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_replacemsg: %v", err)
		}
	}

	if err = d.Set("type", flattenWebfilterProfileFtgdWfQuotaType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "WebfilterProfileFtgdWfQuota-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("unit", flattenWebfilterProfileFtgdWfQuotaUnit3rdl(o["unit"], d, "unit")); err != nil {
		if vv, ok := fortiAPIPatch(o["unit"], "WebfilterProfileFtgdWfQuota-Unit"); ok {
			if err = d.Set("unit", vv); err != nil {
				return fmt.Errorf("Error reading unit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unit: %v", err)
		}
	}

	if err = d.Set("value", flattenWebfilterProfileFtgdWfQuotaValue3rdl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "WebfilterProfileFtgdWfQuota-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenWebfilterProfileFtgdWfQuotaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterProfileFtgdWfQuotaCategory3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfQuotaDuration3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaOverrideReplacemsg3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaUnit3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterProfileFtgdWfQuota(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandWebfilterProfileFtgdWfQuotaCategory3rdl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("duration"); ok || d.HasChange("duration") {
		t, err := expandWebfilterProfileFtgdWfQuotaDuration3rdl(d, v, "duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duration"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWebfilterProfileFtgdWfQuotaId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("override_replacemsg"); ok || d.HasChange("override_replacemsg") {
		t, err := expandWebfilterProfileFtgdWfQuotaOverrideReplacemsg3rdl(d, v, "override_replacemsg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-replacemsg"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandWebfilterProfileFtgdWfQuotaType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("unit"); ok || d.HasChange("unit") {
		t, err := expandWebfilterProfileFtgdWfQuotaUnit3rdl(d, v, "unit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unit"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandWebfilterProfileFtgdWfQuotaValue3rdl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
