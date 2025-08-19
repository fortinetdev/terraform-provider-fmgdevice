// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Policies which can override the 'default' for specific ISL/ICL/FortiLink interface.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerAutoConfigCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerAutoConfigCustomCreate,
		Read:   resourceSwitchControllerAutoConfigCustomRead,
		Update: resourceSwitchControllerAutoConfigCustomUpdate,
		Delete: resourceSwitchControllerAutoConfigCustomDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"switch_binding": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"switch_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchControllerAutoConfigCustomCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerAutoConfigCustom(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAutoConfigCustom resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerAutoConfigCustom(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAutoConfigCustom resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerAutoConfigCustomRead(d, m)
}

func resourceSwitchControllerAutoConfigCustomUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerAutoConfigCustom(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigCustom resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerAutoConfigCustom(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerAutoConfigCustomRead(d, m)
}

func resourceSwitchControllerAutoConfigCustomDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerAutoConfigCustom(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerAutoConfigCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAutoConfigCustomRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerAutoConfigCustom(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAutoConfigCustom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigCustom resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAutoConfigCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigCustomSwitchBinding(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy"
		if _, ok := i["policy"]; ok {
			v := flattenSwitchControllerAutoConfigCustomSwitchBindingPolicy(i["policy"], d, pre_append)
			tmp["policy"] = fortiAPISubPartPatch(v, "SwitchControllerAutoConfigCustom-SwitchBinding-Policy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := i["switch-id"]; ok {
			v := flattenSwitchControllerAutoConfigCustomSwitchBindingSwitchId(i["switch-id"], d, pre_append)
			tmp["switch_id"] = fortiAPISubPartPatch(v, "SwitchControllerAutoConfigCustom-SwitchBinding-SwitchId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerAutoConfigCustomSwitchBindingPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerAutoConfigCustomSwitchBindingSwitchId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerAutoConfigCustom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenSwitchControllerAutoConfigCustomName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerAutoConfigCustom-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("switch_binding", flattenSwitchControllerAutoConfigCustomSwitchBinding(o["switch-binding"], d, "switch_binding")); err != nil {
			if vv, ok := fortiAPIPatch(o["switch-binding"], "SwitchControllerAutoConfigCustom-SwitchBinding"); ok {
				if err = d.Set("switch_binding", vv); err != nil {
					return fmt.Errorf("Error reading switch_binding: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading switch_binding: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("switch_binding"); ok {
			if err = d.Set("switch_binding", flattenSwitchControllerAutoConfigCustomSwitchBinding(o["switch-binding"], d, "switch_binding")); err != nil {
				if vv, ok := fortiAPIPatch(o["switch-binding"], "SwitchControllerAutoConfigCustom-SwitchBinding"); ok {
					if err = d.Set("switch_binding", vv); err != nil {
						return fmt.Errorf("Error reading switch_binding: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading switch_binding: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerAutoConfigCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerAutoConfigCustomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigCustomSwitchBinding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["policy"], _ = expandSwitchControllerAutoConfigCustomSwitchBindingPolicy(d, i["policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["switch-id"], _ = expandSwitchControllerAutoConfigCustomSwitchBindingSwitchId(d, i["switch_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerAutoConfigCustomSwitchBindingPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerAutoConfigCustomSwitchBindingSwitchId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerAutoConfigCustom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerAutoConfigCustomName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("switch_binding"); ok || d.HasChange("switch_binding") {
		t, err := expandSwitchControllerAutoConfigCustomSwitchBinding(d, v, "switch_binding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-binding"] = t
		}
	}

	return &obj, nil
}
