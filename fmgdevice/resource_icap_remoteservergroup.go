// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device IcapRemoteServerGroup

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIcapRemoteServerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceIcapRemoteServerGroupCreate,
		Read:   resourceIcapRemoteServerGroupRead,
		Update: resourceIcapRemoteServerGroupUpdate,
		Delete: resourceIcapRemoteServerGroupDelete,

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
			"ldb_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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

func resourceIcapRemoteServerGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIcapRemoteServerGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating IcapRemoteServerGroup resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIcapRemoteServerGroup(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIcapRemoteServerGroup(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IcapRemoteServerGroup resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateIcapRemoteServerGroup(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IcapRemoteServerGroup resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceIcapRemoteServerGroupRead(d, m)
}

func resourceIcapRemoteServerGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIcapRemoteServerGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating IcapRemoteServerGroup resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIcapRemoteServerGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IcapRemoteServerGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceIcapRemoteServerGroupRead(d, m)
}

func resourceIcapRemoteServerGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteIcapRemoteServerGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IcapRemoteServerGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIcapRemoteServerGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadIcapRemoteServerGroup(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IcapRemoteServerGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIcapRemoteServerGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IcapRemoteServerGroup resource from API: %v", err)
	}
	return nil
}

func flattenIcapRemoteServerGroupLdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapRemoteServerGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapRemoteServerGroupServerList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenIcapRemoteServerGroupServerListName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "IcapRemoteServerGroup-ServerList-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenIcapRemoteServerGroupServerListWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "IcapRemoteServerGroup-ServerList-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenIcapRemoteServerGroupServerListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenIcapRemoteServerGroupServerListWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIcapRemoteServerGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("ldb_method", flattenIcapRemoteServerGroupLdbMethod(o["ldb-method"], d, "ldb_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldb-method"], "IcapRemoteServerGroup-LdbMethod"); ok {
			if err = d.Set("ldb_method", vv); err != nil {
				return fmt.Errorf("Error reading ldb_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("name", flattenIcapRemoteServerGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "IcapRemoteServerGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_list", flattenIcapRemoteServerGroupServerList(o["server-list"], d, "server_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-list"], "IcapRemoteServerGroup-ServerList"); ok {
				if err = d.Set("server_list", vv); err != nil {
					return fmt.Errorf("Error reading server_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_list"); ok {
			if err = d.Set("server_list", flattenIcapRemoteServerGroupServerList(o["server-list"], d, "server_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-list"], "IcapRemoteServerGroup-ServerList"); ok {
					if err = d.Set("server_list", vv); err != nil {
						return fmt.Errorf("Error reading server_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenIcapRemoteServerGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIcapRemoteServerGroupLdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapRemoteServerGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapRemoteServerGroupServerList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandIcapRemoteServerGroupServerListName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandIcapRemoteServerGroupServerListWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandIcapRemoteServerGroupServerListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapRemoteServerGroupServerListWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIcapRemoteServerGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ldb_method"); ok || d.HasChange("ldb_method") {
		t, err := expandIcapRemoteServerGroupLdbMethod(d, v, "ldb_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandIcapRemoteServerGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server_list"); ok || d.HasChange("server_list") {
		t, err := expandIcapRemoteServerGroupServerList(d, v, "server_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-list"] = t
		}
	}

	return &obj, nil
}
