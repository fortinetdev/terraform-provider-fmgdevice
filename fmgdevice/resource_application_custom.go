// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure custom application signatures.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceApplicationCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationCustomCreate,
		Read:   resourceApplicationCustomRead,
		Update: resourceApplicationCustomUpdate,
		Delete: resourceApplicationCustomDelete,

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
			"behavior": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"technology": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceApplicationCustomCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectApplicationCustom(d)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationCustom resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("tag")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadApplicationCustom(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateApplicationCustom(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ApplicationCustom resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateApplicationCustom(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ApplicationCustom resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "tag"))

	return resourceApplicationCustomRead(d, m)
}

func resourceApplicationCustomUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectApplicationCustom(d)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationCustom resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateApplicationCustom(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "tag"))

	return resourceApplicationCustomRead(d, m)
}

func resourceApplicationCustomDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteApplicationCustom(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ApplicationCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationCustomRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadApplicationCustom(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ApplicationCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationCustom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationCustom resource from API: %v", err)
	}
	return nil
}

func flattenApplicationCustomBehavior(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationCustomComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomSignature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomTechnology(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectApplicationCustom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("behavior", flattenApplicationCustomBehavior(o["behavior"], d, "behavior")); err != nil {
		if vv, ok := fortiAPIPatch(o["behavior"], "ApplicationCustom-Behavior"); ok {
			if err = d.Set("behavior", vv); err != nil {
				return fmt.Errorf("Error reading behavior: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading behavior: %v", err)
		}
	}

	if err = d.Set("category", flattenApplicationCustomCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "ApplicationCustom-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("comment", flattenApplicationCustomComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ApplicationCustom-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("fosid", flattenApplicationCustomId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ApplicationCustom-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenApplicationCustomName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ApplicationCustom-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenApplicationCustomProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ApplicationCustom-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("signature", flattenApplicationCustomSignature(o["signature"], d, "signature")); err != nil {
		if vv, ok := fortiAPIPatch(o["signature"], "ApplicationCustom-Signature"); ok {
			if err = d.Set("signature", vv); err != nil {
				return fmt.Errorf("Error reading signature: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signature: %v", err)
		}
	}

	if err = d.Set("tag", flattenApplicationCustomTag(o["tag"], d, "tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag"], "ApplicationCustom-Tag"); ok {
			if err = d.Set("tag", vv); err != nil {
				return fmt.Errorf("Error reading tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag: %v", err)
		}
	}

	if err = d.Set("technology", flattenApplicationCustomTechnology(o["technology"], d, "technology")); err != nil {
		if vv, ok := fortiAPIPatch(o["technology"], "ApplicationCustom-Technology"); ok {
			if err = d.Set("technology", vv); err != nil {
				return fmt.Errorf("Error reading technology: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading technology: %v", err)
		}
	}

	if err = d.Set("vendor", flattenApplicationCustomVendor(o["vendor"], d, "vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["vendor"], "ApplicationCustom-Vendor"); ok {
			if err = d.Set("vendor", vv); err != nil {
				return fmt.Errorf("Error reading vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	return nil
}

func flattenApplicationCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandApplicationCustomBehavior(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationCustomComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomTechnology(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectApplicationCustom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("behavior"); ok || d.HasChange("behavior") {
		t, err := expandApplicationCustomBehavior(d, v, "behavior")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["behavior"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandApplicationCustomCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandApplicationCustomComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandApplicationCustomId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandApplicationCustomName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandApplicationCustomProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("signature"); ok || d.HasChange("signature") {
		t, err := expandApplicationCustomSignature(d, v, "signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature"] = t
		}
	}

	if v, ok := d.GetOk("tag"); ok || d.HasChange("tag") {
		t, err := expandApplicationCustomTag(d, v, "tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag"] = t
		}
	}

	if v, ok := d.GetOk("technology"); ok || d.HasChange("technology") {
		t, err := expandApplicationCustomTechnology(d, v, "technology")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["technology"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok || d.HasChange("vendor") {
		t, err := expandApplicationCustomVendor(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	return &obj, nil
}
