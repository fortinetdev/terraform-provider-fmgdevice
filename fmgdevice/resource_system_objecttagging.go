// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure object tagging.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemObjectTagging() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemObjectTaggingCreate,
		Read:   resourceSystemObjectTaggingRead,
		Update: resourceSystemObjectTaggingUpdate,
		Delete: resourceSystemObjectTaggingDelete,

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
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multiple": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemObjectTaggingCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemObjectTagging(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemObjectTagging resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("category")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemObjectTagging(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemObjectTagging(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemObjectTagging resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemObjectTagging(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemObjectTagging resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "category"))

	return resourceSystemObjectTaggingRead(d, m)
}

func resourceSystemObjectTaggingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemObjectTagging(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemObjectTagging resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemObjectTagging(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemObjectTagging resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "category"))

	return resourceSystemObjectTaggingRead(d, m)
}

func resourceSystemObjectTaggingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemObjectTagging(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemObjectTagging resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemObjectTaggingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemObjectTagging(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemObjectTagging resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemObjectTagging(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemObjectTagging resource from API: %v", err)
	}
	return nil
}

func flattenSystemObjectTaggingAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemObjectTaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemObjectTaggingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemObjectTaggingDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemObjectTaggingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemObjectTaggingMultiple(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemObjectTaggingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemObjectTagging(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenSystemObjectTaggingAddress(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "SystemObjectTagging-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("category", flattenSystemObjectTaggingCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "SystemObjectTagging-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("color", flattenSystemObjectTaggingColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "SystemObjectTagging-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("device", flattenSystemObjectTaggingDevice(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "SystemObjectTagging-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemObjectTaggingInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemObjectTagging-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("multiple", flattenSystemObjectTaggingMultiple(o["multiple"], d, "multiple")); err != nil {
		if vv, ok := fortiAPIPatch(o["multiple"], "SystemObjectTagging-Multiple"); ok {
			if err = d.Set("multiple", vv); err != nil {
				return fmt.Errorf("Error reading multiple: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multiple: %v", err)
		}
	}

	if err = d.Set("tags", flattenSystemObjectTaggingTags(o["tags"], d, "tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["tags"], "SystemObjectTagging-Tags"); ok {
			if err = d.Set("tags", vv); err != nil {
				return fmt.Errorf("Error reading tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tags: %v", err)
		}
	}

	return nil
}

func flattenSystemObjectTaggingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemObjectTaggingAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingMultiple(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemObjectTagging(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandSystemObjectTaggingAddress(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandSystemObjectTaggingCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandSystemObjectTaggingColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandSystemObjectTaggingDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemObjectTaggingInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("multiple"); ok || d.HasChange("multiple") {
		t, err := expandSystemObjectTaggingMultiple(d, v, "multiple")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple"] = t
		}
	}

	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		t, err := expandSystemObjectTaggingTags(d, v, "tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tags"] = t
		}
	}

	return &obj, nil
}
