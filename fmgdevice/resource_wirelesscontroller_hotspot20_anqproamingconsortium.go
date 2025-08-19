// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure roaming consortium.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpRoamingConsortium() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpRoamingConsortiumCreate,
		Read:   resourceWirelessControllerHotspot20AnqpRoamingConsortiumRead,
		Update: resourceWirelessControllerHotspot20AnqpRoamingConsortiumUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpRoamingConsortiumDelete,

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
			"oi_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"oi": &schema.Schema{
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

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerHotspot20AnqpRoamingConsortium(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpRoamingConsortium resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20AnqpRoamingConsortium(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpRoamingConsortium resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20AnqpRoamingConsortiumRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerHotspot20AnqpRoamingConsortium(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpRoamingConsortium resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20AnqpRoamingConsortium(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpRoamingConsortium resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20AnqpRoamingConsortiumRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerHotspot20AnqpRoamingConsortium(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpRoamingConsortium resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerHotspot20AnqpRoamingConsortium(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpRoamingConsortium resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpRoamingConsortium(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpRoamingConsortium resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpRoamingConsortium-OiList-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpRoamingConsortium-OiList-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oi"
		if _, ok := i["oi"]; ok {
			v := flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi(i["oi"], d, pre_append)
			tmp["oi"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpRoamingConsortium-OiList-Oi")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpRoamingConsortium(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenWirelessControllerHotspot20AnqpRoamingConsortiumName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerHotspot20AnqpRoamingConsortium-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("oi_list", flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiList(o["oi-list"], d, "oi_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["oi-list"], "WirelessControllerHotspot20AnqpRoamingConsortium-OiList"); ok {
				if err = d.Set("oi_list", vv); err != nil {
					return fmt.Errorf("Error reading oi_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading oi_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("oi_list"); ok {
			if err = d.Set("oi_list", flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiList(o["oi-list"], d, "oi_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["oi-list"], "WirelessControllerHotspot20AnqpRoamingConsortium-OiList"); ok {
					if err = d.Set("oi_list", vv); err != nil {
						return fmt.Errorf("Error reading oi_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading oi_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oi"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["oi"], _ = expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi(d, i["oi"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpRoamingConsortium(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerHotspot20AnqpRoamingConsortiumName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("oi_list"); ok || d.HasChange("oi_list") {
		t, err := expandWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d, v, "oi_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oi-list"] = t
		}
	}

	return &obj, nil
}
