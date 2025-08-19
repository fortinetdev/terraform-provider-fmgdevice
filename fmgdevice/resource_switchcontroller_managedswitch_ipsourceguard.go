// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IP source guard.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchIpSourceGuard() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchIpSourceGuardCreate,
		Read:   resourceSwitchControllerManagedSwitchIpSourceGuardRead,
		Update: resourceSwitchControllerManagedSwitchIpSourceGuardUpdate,
		Delete: resourceSwitchControllerManagedSwitchIpSourceGuardDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"binding_entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchIpSourceGuardCreate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchIpSourceGuard(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchIpSourceGuard resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerManagedSwitchIpSourceGuard(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchIpSourceGuard resource: %v", err)
	}

	d.SetId(getStringKey(d, "port"))

	return resourceSwitchControllerManagedSwitchIpSourceGuardRead(d, m)
}

func resourceSwitchControllerManagedSwitchIpSourceGuardUpdate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchIpSourceGuard(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchIpSourceGuard resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchIpSourceGuard(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchIpSourceGuard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "port"))

	return resourceSwitchControllerManagedSwitchIpSourceGuardRead(d, m)
}

func resourceSwitchControllerManagedSwitchIpSourceGuardDelete(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerManagedSwitchIpSourceGuard(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchIpSourceGuard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchIpSourceGuardRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	managed_switch := d.Get("managed_switch").(string)
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
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	o, err := c.ReadSwitchControllerManagedSwitchIpSourceGuard(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchIpSourceGuard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchIpSourceGuard(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchIpSourceGuard resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntry2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_name"
		if _, ok := i["entry-name"]; ok {
			v := flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName2edl(i["entry-name"], d, pre_append)
			tmp["entry_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIpSourceGuard-BindingEntry-EntryName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIpSourceGuard-BindingEntry-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac2edl(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIpSourceGuard-BindingEntry-Mac")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchIpSourceGuard(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("binding_entry", flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntry2edl(o["binding-entry"], d, "binding_entry")); err != nil {
			if vv, ok := fortiAPIPatch(o["binding-entry"], "SwitchControllerManagedSwitchIpSourceGuard-BindingEntry"); ok {
				if err = d.Set("binding_entry", vv); err != nil {
					return fmt.Errorf("Error reading binding_entry: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading binding_entry: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("binding_entry"); ok {
			if err = d.Set("binding_entry", flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntry2edl(o["binding-entry"], d, "binding_entry")); err != nil {
				if vv, ok := fortiAPIPatch(o["binding-entry"], "SwitchControllerManagedSwitchIpSourceGuard-BindingEntry"); ok {
					if err = d.Set("binding_entry", vv); err != nil {
						return fmt.Errorf("Error reading binding_entry: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading binding_entry: %v", err)
				}
			}
		}
	}

	if err = d.Set("description", flattenSwitchControllerManagedSwitchIpSourceGuardDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerManagedSwitchIpSourceGuard-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("port", flattenSwitchControllerManagedSwitchIpSourceGuardPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SwitchControllerManagedSwitchIpSourceGuard-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchIpSourceGuardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["entry-name"], _ = expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName2edl(d, i["entry_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp2edl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac2edl(d, i["mac"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchIpSourceGuard(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("binding_entry"); ok || d.HasChange("binding_entry") {
		t, err := expandSwitchControllerManagedSwitchIpSourceGuardBindingEntry2edl(d, v, "binding_entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["binding-entry"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerManagedSwitchIpSourceGuardDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSwitchControllerManagedSwitchIpSourceGuardPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	return &obj, nil
}
