// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch SNMP v1/v2c communities globally.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSnmpCommunity() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSnmpCommunityCreate,
		Read:   resourceSwitchControllerSnmpCommunityRead,
		Update: resourceSwitchControllerSnmpCommunityUpdate,
		Delete: resourceSwitchControllerSnmpCommunityDelete,

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
			"events": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
			"query_v1_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"query_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_v2c_port": &schema.Schema{
				Type:     schema.TypeInt,
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
			"trap_v1_lport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_v1_rport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_v2c_lport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_v2c_rport": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSwitchControllerSnmpCommunityCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerSnmpCommunity(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSnmpCommunity resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerSnmpCommunity(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSnmpCommunity resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerSnmpCommunityRead(d, m)
}

func resourceSwitchControllerSnmpCommunityUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerSnmpCommunity(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpCommunity resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSnmpCommunity(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpCommunity resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerSnmpCommunityRead(d, m)
}

func resourceSwitchControllerSnmpCommunityDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerSnmpCommunity(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSnmpCommunity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSnmpCommunityRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerSnmpCommunity(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpCommunity resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSnmpCommunity(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpCommunity resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSnmpCommunityEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSnmpCommunityHosts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSwitchControllerSnmpCommunityHostsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SwitchControllerSnmpCommunity-Hosts-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSwitchControllerSnmpCommunityHostsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SwitchControllerSnmpCommunity-Hosts-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerSnmpCommunityHostsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityHostsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSnmpCommunityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityQueryV1Port(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityQueryV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityQueryV2CPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityQueryV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityTrapV1Lport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityTrapV1Rport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityTrapV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityTrapV2CLport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityTrapV2CRport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityTrapV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerSnmpCommunity(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("events", flattenSwitchControllerSnmpCommunityEvents(o["events"], d, "events")); err != nil {
		if vv, ok := fortiAPIPatch(o["events"], "SwitchControllerSnmpCommunity-Events"); ok {
			if err = d.Set("events", vv); err != nil {
				return fmt.Errorf("Error reading events: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("hosts", flattenSwitchControllerSnmpCommunityHosts(o["hosts"], d, "hosts")); err != nil {
			if vv, ok := fortiAPIPatch(o["hosts"], "SwitchControllerSnmpCommunity-Hosts"); ok {
				if err = d.Set("hosts", vv); err != nil {
					return fmt.Errorf("Error reading hosts: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading hosts: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hosts"); ok {
			if err = d.Set("hosts", flattenSwitchControllerSnmpCommunityHosts(o["hosts"], d, "hosts")); err != nil {
				if vv, ok := fortiAPIPatch(o["hosts"], "SwitchControllerSnmpCommunity-Hosts"); ok {
					if err = d.Set("hosts", vv); err != nil {
						return fmt.Errorf("Error reading hosts: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading hosts: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenSwitchControllerSnmpCommunityId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SwitchControllerSnmpCommunity-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerSnmpCommunityName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerSnmpCommunity-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("query_v1_port", flattenSwitchControllerSnmpCommunityQueryV1Port(o["query-v1-port"], d, "query_v1_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v1-port"], "SwitchControllerSnmpCommunity-QueryV1Port"); ok {
			if err = d.Set("query_v1_port", vv); err != nil {
				return fmt.Errorf("Error reading query_v1_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v1_port: %v", err)
		}
	}

	if err = d.Set("query_v1_status", flattenSwitchControllerSnmpCommunityQueryV1Status(o["query-v1-status"], d, "query_v1_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v1-status"], "SwitchControllerSnmpCommunity-QueryV1Status"); ok {
			if err = d.Set("query_v1_status", vv); err != nil {
				return fmt.Errorf("Error reading query_v1_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v1_status: %v", err)
		}
	}

	if err = d.Set("query_v2c_port", flattenSwitchControllerSnmpCommunityQueryV2CPort(o["query-v2c-port"], d, "query_v2c_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v2c-port"], "SwitchControllerSnmpCommunity-QueryV2CPort"); ok {
			if err = d.Set("query_v2c_port", vv); err != nil {
				return fmt.Errorf("Error reading query_v2c_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v2c_port: %v", err)
		}
	}

	if err = d.Set("query_v2c_status", flattenSwitchControllerSnmpCommunityQueryV2CStatus(o["query-v2c-status"], d, "query_v2c_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v2c-status"], "SwitchControllerSnmpCommunity-QueryV2CStatus"); ok {
			if err = d.Set("query_v2c_status", vv); err != nil {
				return fmt.Errorf("Error reading query_v2c_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v2c_status: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchControllerSnmpCommunityStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SwitchControllerSnmpCommunity-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_v1_lport", flattenSwitchControllerSnmpCommunityTrapV1Lport(o["trap-v1-lport"], d, "trap_v1_lport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v1-lport"], "SwitchControllerSnmpCommunity-TrapV1Lport"); ok {
			if err = d.Set("trap_v1_lport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v1_lport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v1_lport: %v", err)
		}
	}

	if err = d.Set("trap_v1_rport", flattenSwitchControllerSnmpCommunityTrapV1Rport(o["trap-v1-rport"], d, "trap_v1_rport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v1-rport"], "SwitchControllerSnmpCommunity-TrapV1Rport"); ok {
			if err = d.Set("trap_v1_rport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v1_rport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v1_rport: %v", err)
		}
	}

	if err = d.Set("trap_v1_status", flattenSwitchControllerSnmpCommunityTrapV1Status(o["trap-v1-status"], d, "trap_v1_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v1-status"], "SwitchControllerSnmpCommunity-TrapV1Status"); ok {
			if err = d.Set("trap_v1_status", vv); err != nil {
				return fmt.Errorf("Error reading trap_v1_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v1_status: %v", err)
		}
	}

	if err = d.Set("trap_v2c_lport", flattenSwitchControllerSnmpCommunityTrapV2CLport(o["trap-v2c-lport"], d, "trap_v2c_lport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v2c-lport"], "SwitchControllerSnmpCommunity-TrapV2CLport"); ok {
			if err = d.Set("trap_v2c_lport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v2c_lport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v2c_lport: %v", err)
		}
	}

	if err = d.Set("trap_v2c_rport", flattenSwitchControllerSnmpCommunityTrapV2CRport(o["trap-v2c-rport"], d, "trap_v2c_rport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v2c-rport"], "SwitchControllerSnmpCommunity-TrapV2CRport"); ok {
			if err = d.Set("trap_v2c_rport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v2c_rport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v2c_rport: %v", err)
		}
	}

	if err = d.Set("trap_v2c_status", flattenSwitchControllerSnmpCommunityTrapV2CStatus(o["trap-v2c-status"], d, "trap_v2c_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v2c-status"], "SwitchControllerSnmpCommunity-TrapV2CStatus"); ok {
			if err = d.Set("trap_v2c_status", vv); err != nil {
				return fmt.Errorf("Error reading trap_v2c_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v2c_status: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSnmpCommunityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSnmpCommunityEvents(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerSnmpCommunityHosts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSwitchControllerSnmpCommunityHostsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSwitchControllerSnmpCommunityHostsIp(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerSnmpCommunityHostsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityHostsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerSnmpCommunityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityQueryV1Port(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityQueryV1Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityQueryV2CPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityQueryV2CStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityTrapV1Lport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityTrapV1Rport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityTrapV1Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityTrapV2CLport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityTrapV2CRport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityTrapV2CStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSnmpCommunity(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("events"); ok || d.HasChange("events") {
		t, err := expandSwitchControllerSnmpCommunityEvents(d, v, "events")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["events"] = t
		}
	}

	if v, ok := d.GetOk("hosts"); ok || d.HasChange("hosts") {
		t, err := expandSwitchControllerSnmpCommunityHosts(d, v, "hosts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosts"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSwitchControllerSnmpCommunityId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerSnmpCommunityName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("query_v1_port"); ok || d.HasChange("query_v1_port") {
		t, err := expandSwitchControllerSnmpCommunityQueryV1Port(d, v, "query_v1_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v1-port"] = t
		}
	}

	if v, ok := d.GetOk("query_v1_status"); ok || d.HasChange("query_v1_status") {
		t, err := expandSwitchControllerSnmpCommunityQueryV1Status(d, v, "query_v1_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v1-status"] = t
		}
	}

	if v, ok := d.GetOk("query_v2c_port"); ok || d.HasChange("query_v2c_port") {
		t, err := expandSwitchControllerSnmpCommunityQueryV2CPort(d, v, "query_v2c_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v2c-port"] = t
		}
	}

	if v, ok := d.GetOk("query_v2c_status"); ok || d.HasChange("query_v2c_status") {
		t, err := expandSwitchControllerSnmpCommunityQueryV2CStatus(d, v, "query_v2c_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v2c-status"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSwitchControllerSnmpCommunityStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_lport"); ok || d.HasChange("trap_v1_lport") {
		t, err := expandSwitchControllerSnmpCommunityTrapV1Lport(d, v, "trap_v1_lport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-lport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_rport"); ok || d.HasChange("trap_v1_rport") {
		t, err := expandSwitchControllerSnmpCommunityTrapV1Rport(d, v, "trap_v1_rport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-rport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_status"); ok || d.HasChange("trap_v1_status") {
		t, err := expandSwitchControllerSnmpCommunityTrapV1Status(d, v, "trap_v1_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-status"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_lport"); ok || d.HasChange("trap_v2c_lport") {
		t, err := expandSwitchControllerSnmpCommunityTrapV2CLport(d, v, "trap_v2c_lport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-lport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_rport"); ok || d.HasChange("trap_v2c_rport") {
		t, err := expandSwitchControllerSnmpCommunityTrapV2CRport(d, v, "trap_v2c_rport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-rport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_status"); ok || d.HasChange("trap_v2c_status") {
		t, err := expandSwitchControllerSnmpCommunityTrapV2CStatus(d, v, "trap_v2c_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-status"] = t
		}
	}

	return &obj, nil
}
