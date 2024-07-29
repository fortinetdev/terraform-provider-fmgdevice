// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AP local configuration profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerApcfgProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerApcfgProfileCreate,
		Read:   resourceWirelessControllerApcfgProfileRead,
		Update: resourceWirelessControllerApcfgProfileUpdate,
		Delete: resourceWirelessControllerApcfgProfileDelete,

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
			"ac_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ac_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ac_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ac_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"command_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"passwd_value": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerApcfgProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectWirelessControllerApcfgProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerApcfgProfile resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerApcfgProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerApcfgProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerApcfgProfileRead(d, m)
}

func resourceWirelessControllerApcfgProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectWirelessControllerApcfgProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerApcfgProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerApcfgProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerApcfgProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerApcfgProfileRead(d, m)
}

func resourceWirelessControllerApcfgProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteWirelessControllerApcfgProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerApcfgProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerApcfgProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerApcfgProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerApcfgProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerApcfgProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerApcfgProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerApcfgProfileAcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileAcPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileAcTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileAcType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileApFamily(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileCommandList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWirelessControllerApcfgProfileCommandListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerApcfgProfile-CommandList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWirelessControllerApcfgProfileCommandListName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WirelessControllerApcfgProfile-CommandList-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenWirelessControllerApcfgProfileCommandListType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "WirelessControllerApcfgProfile-CommandList-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenWirelessControllerApcfgProfileCommandListValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "WirelessControllerApcfgProfile-CommandList-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerApcfgProfileCommandListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileCommandListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileCommandListType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileCommandListValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerApcfgProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("ac_ip", flattenWirelessControllerApcfgProfileAcIp(o["ac-ip"], d, "ac_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ac-ip"], "WirelessControllerApcfgProfile-AcIp"); ok {
			if err = d.Set("ac_ip", vv); err != nil {
				return fmt.Errorf("Error reading ac_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ac_ip: %v", err)
		}
	}

	if err = d.Set("ac_port", flattenWirelessControllerApcfgProfileAcPort(o["ac-port"], d, "ac_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ac-port"], "WirelessControllerApcfgProfile-AcPort"); ok {
			if err = d.Set("ac_port", vv); err != nil {
				return fmt.Errorf("Error reading ac_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ac_port: %v", err)
		}
	}

	if err = d.Set("ac_timer", flattenWirelessControllerApcfgProfileAcTimer(o["ac-timer"], d, "ac_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["ac-timer"], "WirelessControllerApcfgProfile-AcTimer"); ok {
			if err = d.Set("ac_timer", vv); err != nil {
				return fmt.Errorf("Error reading ac_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ac_timer: %v", err)
		}
	}

	if err = d.Set("ac_type", flattenWirelessControllerApcfgProfileAcType(o["ac-type"], d, "ac_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ac-type"], "WirelessControllerApcfgProfile-AcType"); ok {
			if err = d.Set("ac_type", vv); err != nil {
				return fmt.Errorf("Error reading ac_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ac_type: %v", err)
		}
	}

	if err = d.Set("ap_family", flattenWirelessControllerApcfgProfileApFamily(o["ap-family"], d, "ap_family")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-family"], "WirelessControllerApcfgProfile-ApFamily"); ok {
			if err = d.Set("ap_family", vv); err != nil {
				return fmt.Errorf("Error reading ap_family: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_family: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("command_list", flattenWirelessControllerApcfgProfileCommandList(o["command-list"], d, "command_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["command-list"], "WirelessControllerApcfgProfile-CommandList"); ok {
				if err = d.Set("command_list", vv); err != nil {
					return fmt.Errorf("Error reading command_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading command_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("command_list"); ok {
			if err = d.Set("command_list", flattenWirelessControllerApcfgProfileCommandList(o["command-list"], d, "command_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["command-list"], "WirelessControllerApcfgProfile-CommandList"); ok {
					if err = d.Set("command_list", vv); err != nil {
						return fmt.Errorf("Error reading command_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading command_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenWirelessControllerApcfgProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerApcfgProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerApcfgProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerApcfgProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerApcfgProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerApcfgProfileAcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileAcPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileAcTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileAcType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileApFamily(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileCommandList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWirelessControllerApcfgProfileCommandListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWirelessControllerApcfgProfileCommandListName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passwd_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passwd-value"], _ = expandWirelessControllerApcfgProfileCommandListPasswdValue(d, i["passwd_value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandWirelessControllerApcfgProfileCommandListType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandWirelessControllerApcfgProfileCommandListValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerApcfgProfileCommandListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileCommandListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileCommandListPasswdValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerApcfgProfileCommandListType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileCommandListValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerApcfgProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ac_ip"); ok || d.HasChange("ac_ip") {
		t, err := expandWirelessControllerApcfgProfileAcIp(d, v, "ac_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-ip"] = t
		}
	}

	if v, ok := d.GetOk("ac_port"); ok || d.HasChange("ac_port") {
		t, err := expandWirelessControllerApcfgProfileAcPort(d, v, "ac_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-port"] = t
		}
	}

	if v, ok := d.GetOk("ac_timer"); ok || d.HasChange("ac_timer") {
		t, err := expandWirelessControllerApcfgProfileAcTimer(d, v, "ac_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-timer"] = t
		}
	}

	if v, ok := d.GetOk("ac_type"); ok || d.HasChange("ac_type") {
		t, err := expandWirelessControllerApcfgProfileAcType(d, v, "ac_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-type"] = t
		}
	}

	if v, ok := d.GetOk("ap_family"); ok || d.HasChange("ap_family") {
		t, err := expandWirelessControllerApcfgProfileApFamily(d, v, "ap_family")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-family"] = t
		}
	}

	if v, ok := d.GetOk("command_list"); ok || d.HasChange("command_list") {
		t, err := expandWirelessControllerApcfgProfileCommandList(d, v, "command_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-list"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerApcfgProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerApcfgProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
