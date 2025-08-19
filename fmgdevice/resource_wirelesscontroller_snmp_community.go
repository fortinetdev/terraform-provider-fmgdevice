// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SNMP Community Configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerSnmpCommunity() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerSnmpCommunityCreate,
		Read:   resourceWirelessControllerSnmpCommunityRead,
		Update: resourceWirelessControllerSnmpCommunityUpdate,
		Delete: resourceWirelessControllerSnmpCommunityDelete,

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
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"query_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_v2c_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_v2c_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerSnmpCommunityCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerSnmpCommunity(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerSnmpCommunity resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerSnmpCommunity(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerSnmpCommunity resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWirelessControllerSnmpCommunityRead(d, m)
}

func resourceWirelessControllerSnmpCommunityUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerSnmpCommunity(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSnmpCommunity resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerSnmpCommunity(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSnmpCommunity resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWirelessControllerSnmpCommunityRead(d, m)
}

func resourceWirelessControllerSnmpCommunityDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerSnmpCommunity(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerSnmpCommunity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSnmpCommunityRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerSnmpCommunity(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSnmpCommunity resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerSnmpCommunity(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSnmpCommunity resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerSnmpCommunityHosts2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWirelessControllerSnmpCommunityHostsId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerSnmpCommunity-Hosts-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenWirelessControllerSnmpCommunityHostsIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "WirelessControllerSnmpCommunity-Hosts-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerSnmpCommunityHostsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityHostsIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityQueryV1Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityQueryV2CStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityTrapV1Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityTrapV2CStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerSnmpCommunity(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("hosts", flattenWirelessControllerSnmpCommunityHosts2edl(o["hosts"], d, "hosts")); err != nil {
			if vv, ok := fortiAPIPatch(o["hosts"], "WirelessControllerSnmpCommunity-Hosts"); ok {
				if err = d.Set("hosts", vv); err != nil {
					return fmt.Errorf("Error reading hosts: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading hosts: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hosts"); ok {
			if err = d.Set("hosts", flattenWirelessControllerSnmpCommunityHosts2edl(o["hosts"], d, "hosts")); err != nil {
				if vv, ok := fortiAPIPatch(o["hosts"], "WirelessControllerSnmpCommunity-Hosts"); ok {
					if err = d.Set("hosts", vv); err != nil {
						return fmt.Errorf("Error reading hosts: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading hosts: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenWirelessControllerSnmpCommunityId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WirelessControllerSnmpCommunity-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerSnmpCommunityName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerSnmpCommunity-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("query_v1_status", flattenWirelessControllerSnmpCommunityQueryV1Status2edl(o["query-v1-status"], d, "query_v1_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v1-status"], "WirelessControllerSnmpCommunity-QueryV1Status"); ok {
			if err = d.Set("query_v1_status", vv); err != nil {
				return fmt.Errorf("Error reading query_v1_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v1_status: %v", err)
		}
	}

	if err = d.Set("query_v2c_status", flattenWirelessControllerSnmpCommunityQueryV2CStatus2edl(o["query-v2c-status"], d, "query_v2c_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v2c-status"], "WirelessControllerSnmpCommunity-QueryV2CStatus"); ok {
			if err = d.Set("query_v2c_status", vv); err != nil {
				return fmt.Errorf("Error reading query_v2c_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v2c_status: %v", err)
		}
	}

	if err = d.Set("status", flattenWirelessControllerSnmpCommunityStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WirelessControllerSnmpCommunity-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_v1_status", flattenWirelessControllerSnmpCommunityTrapV1Status2edl(o["trap-v1-status"], d, "trap_v1_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v1-status"], "WirelessControllerSnmpCommunity-TrapV1Status"); ok {
			if err = d.Set("trap_v1_status", vv); err != nil {
				return fmt.Errorf("Error reading trap_v1_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v1_status: %v", err)
		}
	}

	if err = d.Set("trap_v2c_status", flattenWirelessControllerSnmpCommunityTrapV2CStatus2edl(o["trap-v2c-status"], d, "trap_v2c_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v2c-status"], "WirelessControllerSnmpCommunity-TrapV2CStatus"); ok {
			if err = d.Set("trap_v2c_status", vv); err != nil {
				return fmt.Errorf("Error reading trap_v2c_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v2c_status: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerSnmpCommunityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerSnmpCommunityHosts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerSnmpCommunityHostsId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandWirelessControllerSnmpCommunityHostsIp2edl(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerSnmpCommunityHostsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityHostsIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityQueryV1Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityQueryV2CStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityTrapV1Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityTrapV2CStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerSnmpCommunity(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("hosts"); ok || d.HasChange("hosts") {
		t, err := expandWirelessControllerSnmpCommunityHosts2edl(d, v, "hosts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosts"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWirelessControllerSnmpCommunityId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerSnmpCommunityName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("query_v1_status"); ok || d.HasChange("query_v1_status") {
		t, err := expandWirelessControllerSnmpCommunityQueryV1Status2edl(d, v, "query_v1_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v1-status"] = t
		}
	}

	if v, ok := d.GetOk("query_v2c_status"); ok || d.HasChange("query_v2c_status") {
		t, err := expandWirelessControllerSnmpCommunityQueryV2CStatus2edl(d, v, "query_v2c_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v2c-status"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWirelessControllerSnmpCommunityStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_status"); ok || d.HasChange("trap_v1_status") {
		t, err := expandWirelessControllerSnmpCommunityTrapV1Status2edl(d, v, "trap_v1_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-status"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_status"); ok || d.HasChange("trap_v2c_status") {
		t, err := expandWirelessControllerSnmpCommunityTrapV2CStatus2edl(d, v, "trap_v2c_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-status"] = t
		}
	}

	return &obj, nil
}
