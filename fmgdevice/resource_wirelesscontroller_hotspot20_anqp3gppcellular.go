// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure 3GPP public land mobile network (PLMN).

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20Anqp3GppCellular() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20Anqp3GppCellularCreate,
		Read:   resourceWirelessControllerHotspot20Anqp3GppCellularRead,
		Update: resourceWirelessControllerHotspot20Anqp3GppCellularUpdate,
		Delete: resourceWirelessControllerHotspot20Anqp3GppCellularDelete,

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
			"mcc_mnc_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mcc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mnc": &schema.Schema{
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

func resourceWirelessControllerHotspot20Anqp3GppCellularCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerHotspot20Anqp3GppCellular(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20Anqp3GppCellular resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20Anqp3GppCellular(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20Anqp3GppCellular resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20Anqp3GppCellularRead(d, m)
}

func resourceWirelessControllerHotspot20Anqp3GppCellularUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerHotspot20Anqp3GppCellular(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20Anqp3GppCellular resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20Anqp3GppCellular(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20Anqp3GppCellular resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20Anqp3GppCellularRead(d, m)
}

func resourceWirelessControllerHotspot20Anqp3GppCellularDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerHotspot20Anqp3GppCellular(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20Anqp3GppCellular resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20Anqp3GppCellularRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerHotspot20Anqp3GppCellular(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20Anqp3GppCellular resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20Anqp3GppCellular(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20Anqp3GppCellular resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20Anqp3GppCellularMccMncList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20Anqp3GppCellular-MccMncList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcc"
		if _, ok := i["mcc"]; ok {
			v := flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc(i["mcc"], d, pre_append)
			tmp["mcc"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20Anqp3GppCellular-MccMncList-Mcc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mnc"
		if _, ok := i["mnc"]; ok {
			v := flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc(i["mnc"], d, pre_append)
			tmp["mnc"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20Anqp3GppCellular-MccMncList-Mnc")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20Anqp3GppCellularName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20Anqp3GppCellular(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("mcc_mnc_list", flattenWirelessControllerHotspot20Anqp3GppCellularMccMncList(o["mcc-mnc-list"], d, "mcc_mnc_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["mcc-mnc-list"], "WirelessControllerHotspot20Anqp3GppCellular-MccMncList"); ok {
				if err = d.Set("mcc_mnc_list", vv); err != nil {
					return fmt.Errorf("Error reading mcc_mnc_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mcc_mnc_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mcc_mnc_list"); ok {
			if err = d.Set("mcc_mnc_list", flattenWirelessControllerHotspot20Anqp3GppCellularMccMncList(o["mcc-mnc-list"], d, "mcc_mnc_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["mcc-mnc-list"], "WirelessControllerHotspot20Anqp3GppCellular-MccMncList"); ok {
					if err = d.Set("mcc_mnc_list", vv); err != nil {
						return fmt.Errorf("Error reading mcc_mnc_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mcc_mnc_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenWirelessControllerHotspot20Anqp3GppCellularName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerHotspot20Anqp3GppCellular-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20Anqp3GppCellularFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20Anqp3GppCellularMccMncList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerHotspot20Anqp3GppCellularMccMncListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mcc"], _ = expandWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc(d, i["mcc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mnc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mnc"], _ = expandWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc(d, i["mnc"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20Anqp3GppCellularMccMncListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20Anqp3GppCellularName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20Anqp3GppCellular(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mcc_mnc_list"); ok || d.HasChange("mcc_mnc_list") {
		t, err := expandWirelessControllerHotspot20Anqp3GppCellularMccMncList(d, v, "mcc_mnc_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcc-mnc-list"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerHotspot20Anqp3GppCellularName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
