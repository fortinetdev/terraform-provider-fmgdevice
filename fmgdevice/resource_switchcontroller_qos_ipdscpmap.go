// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch QoS IP precedence/DSCP.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerQosIpDscpMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerQosIpDscpMapCreate,
		Read:   resourceSwitchControllerQosIpDscpMapRead,
		Update: resourceSwitchControllerQosIpDscpMapUpdate,
		Delete: resourceSwitchControllerQosIpDscpMapDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"map": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cos_queue": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"diffserv": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ip_precedence": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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

func resourceSwitchControllerQosIpDscpMapCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerQosIpDscpMap(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosIpDscpMap resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerQosIpDscpMap(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosIpDscpMap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerQosIpDscpMapRead(d, m)
}

func resourceSwitchControllerQosIpDscpMapUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerQosIpDscpMap(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosIpDscpMap resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerQosIpDscpMap(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosIpDscpMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerQosIpDscpMapRead(d, m)
}

func resourceSwitchControllerQosIpDscpMapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerQosIpDscpMap(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerQosIpDscpMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerQosIpDscpMapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerQosIpDscpMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosIpDscpMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerQosIpDscpMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosIpDscpMap resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerQosIpDscpMapDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosIpDscpMapMap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos_queue"
		if _, ok := i["cos-queue"]; ok {
			v := flattenSwitchControllerQosIpDscpMapMapCosQueue(i["cos-queue"], d, pre_append)
			tmp["cos_queue"] = fortiAPISubPartPatch(v, "SwitchControllerQosIpDscpMap-Map-CosQueue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffserv"
		if _, ok := i["diffserv"]; ok {
			v := flattenSwitchControllerQosIpDscpMapMapDiffserv(i["diffserv"], d, pre_append)
			tmp["diffserv"] = fortiAPISubPartPatch(v, "SwitchControllerQosIpDscpMap-Map-Diffserv")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_precedence"
		if _, ok := i["ip-precedence"]; ok {
			v := flattenSwitchControllerQosIpDscpMapMapIpPrecedence(i["ip-precedence"], d, pre_append)
			tmp["ip_precedence"] = fortiAPISubPartPatch(v, "SwitchControllerQosIpDscpMap-Map-IpPrecedence")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSwitchControllerQosIpDscpMapMapName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SwitchControllerQosIpDscpMap-Map-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenSwitchControllerQosIpDscpMapMapValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "SwitchControllerQosIpDscpMap-Map-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerQosIpDscpMapMapCosQueue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosIpDscpMapMapDiffserv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerQosIpDscpMapMapIpPrecedence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerQosIpDscpMapMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosIpDscpMapMapValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosIpDscpMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerQosIpDscpMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("description", flattenSwitchControllerQosIpDscpMapDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerQosIpDscpMap-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("map", flattenSwitchControllerQosIpDscpMapMap(o["map"], d, "map")); err != nil {
			if vv, ok := fortiAPIPatch(o["map"], "SwitchControllerQosIpDscpMap-Map"); ok {
				if err = d.Set("map", vv); err != nil {
					return fmt.Errorf("Error reading map: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading map: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("map"); ok {
			if err = d.Set("map", flattenSwitchControllerQosIpDscpMapMap(o["map"], d, "map")); err != nil {
				if vv, ok := fortiAPIPatch(o["map"], "SwitchControllerQosIpDscpMap-Map"); ok {
					if err = d.Set("map", vv); err != nil {
						return fmt.Errorf("Error reading map: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading map: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSwitchControllerQosIpDscpMapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerQosIpDscpMap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerQosIpDscpMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerQosIpDscpMapDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosIpDscpMapMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos_queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos-queue"], _ = expandSwitchControllerQosIpDscpMapMapCosQueue(d, i["cos_queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffserv"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["diffserv"], _ = expandSwitchControllerQosIpDscpMapMapDiffserv(d, i["diffserv"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_precedence"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-precedence"], _ = expandSwitchControllerQosIpDscpMapMapIpPrecedence(d, i["ip_precedence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSwitchControllerQosIpDscpMapMapName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandSwitchControllerQosIpDscpMapMapValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerQosIpDscpMapMapCosQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosIpDscpMapMapDiffserv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerQosIpDscpMapMapIpPrecedence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerQosIpDscpMapMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosIpDscpMapMapValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosIpDscpMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerQosIpDscpMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerQosIpDscpMapDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("map"); ok || d.HasChange("map") {
		t, err := expandSwitchControllerQosIpDscpMapMap(d, v, "map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["map"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerQosIpDscpMapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
