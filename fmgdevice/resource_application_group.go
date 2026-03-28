// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure firewall application groups.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceApplicationGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationGroupCreate,
		Read:   resourceApplicationGroupRead,
		Update: resourceApplicationGroupUpdate,
		Delete: resourceApplicationGroupDelete,

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
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"behavior": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"popularity": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"protocols": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"risk": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"technology": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceApplicationGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectApplicationGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationGroup resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadApplicationGroup(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateApplicationGroup(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ApplicationGroup resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateApplicationGroup(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ApplicationGroup resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceApplicationGroupRead(d, m)
}

func resourceApplicationGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectApplicationGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationGroup resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateApplicationGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceApplicationGroupRead(d, m)
}

func resourceApplicationGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteApplicationGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ApplicationGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadApplicationGroup(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ApplicationGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationGroup resource from API: %v", err)
	}
	return nil
}

func flattenApplicationGroupApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenApplicationGroupBehavior(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationGroupCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenApplicationGroupComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationGroupPopularity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationGroupProtocols(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationGroupRisk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenApplicationGroupTechnology(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationGroupType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationGroupVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectApplicationGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("application", flattenApplicationGroupApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "ApplicationGroup-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("behavior", flattenApplicationGroupBehavior(o["behavior"], d, "behavior")); err != nil {
		if vv, ok := fortiAPIPatch(o["behavior"], "ApplicationGroup-Behavior"); ok {
			if err = d.Set("behavior", vv); err != nil {
				return fmt.Errorf("Error reading behavior: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading behavior: %v", err)
		}
	}

	if err = d.Set("category", flattenApplicationGroupCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "ApplicationGroup-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("comment", flattenApplicationGroupComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ApplicationGroup-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("name", flattenApplicationGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ApplicationGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("popularity", flattenApplicationGroupPopularity(o["popularity"], d, "popularity")); err != nil {
		if vv, ok := fortiAPIPatch(o["popularity"], "ApplicationGroup-Popularity"); ok {
			if err = d.Set("popularity", vv); err != nil {
				return fmt.Errorf("Error reading popularity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading popularity: %v", err)
		}
	}

	if err = d.Set("protocols", flattenApplicationGroupProtocols(o["protocols"], d, "protocols")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocols"], "ApplicationGroup-Protocols"); ok {
			if err = d.Set("protocols", vv); err != nil {
				return fmt.Errorf("Error reading protocols: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocols: %v", err)
		}
	}

	if err = d.Set("risk", flattenApplicationGroupRisk(o["risk"], d, "risk")); err != nil {
		if vv, ok := fortiAPIPatch(o["risk"], "ApplicationGroup-Risk"); ok {
			if err = d.Set("risk", vv); err != nil {
				return fmt.Errorf("Error reading risk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading risk: %v", err)
		}
	}

	if err = d.Set("technology", flattenApplicationGroupTechnology(o["technology"], d, "technology")); err != nil {
		if vv, ok := fortiAPIPatch(o["technology"], "ApplicationGroup-Technology"); ok {
			if err = d.Set("technology", vv); err != nil {
				return fmt.Errorf("Error reading technology: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading technology: %v", err)
		}
	}

	if err = d.Set("type", flattenApplicationGroupType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ApplicationGroup-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("vendor", flattenApplicationGroupVendor(o["vendor"], d, "vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["vendor"], "ApplicationGroup-Vendor"); ok {
			if err = d.Set("vendor", vv); err != nil {
				return fmt.Errorf("Error reading vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	return nil
}

func flattenApplicationGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandApplicationGroupApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandApplicationGroupBehavior(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationGroupCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandApplicationGroupComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupPopularity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationGroupProtocols(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationGroupRisk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandApplicationGroupTechnology(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationGroupType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectApplicationGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandApplicationGroupApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("behavior"); ok || d.HasChange("behavior") {
		t, err := expandApplicationGroupBehavior(d, v, "behavior")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["behavior"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandApplicationGroupCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandApplicationGroupComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandApplicationGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("popularity"); ok || d.HasChange("popularity") {
		t, err := expandApplicationGroupPopularity(d, v, "popularity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["popularity"] = t
		}
	}

	if v, ok := d.GetOk("protocols"); ok || d.HasChange("protocols") {
		t, err := expandApplicationGroupProtocols(d, v, "protocols")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocols"] = t
		}
	}

	if v, ok := d.GetOk("risk"); ok || d.HasChange("risk") {
		t, err := expandApplicationGroupRisk(d, v, "risk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["risk"] = t
		}
	}

	if v, ok := d.GetOk("technology"); ok || d.HasChange("technology") {
		t, err := expandApplicationGroupTechnology(d, v, "technology")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["technology"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandApplicationGroupType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok || d.HasChange("vendor") {
		t, err := expandApplicationGroupVendor(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	return &obj, nil
}
