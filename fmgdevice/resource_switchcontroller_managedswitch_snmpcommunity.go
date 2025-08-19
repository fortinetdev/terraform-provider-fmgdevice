// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration method to edit Simple Network Management Protocol (SNMP) communities.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchSnmpCommunity() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchSnmpCommunityCreate,
		Read:   resourceSwitchControllerManagedSwitchSnmpCommunityRead,
		Update: resourceSwitchControllerManagedSwitchSnmpCommunityUpdate,
		Delete: resourceSwitchControllerManagedSwitchSnmpCommunityDelete,

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

func resourceSwitchControllerManagedSwitchSnmpCommunityCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerManagedSwitchSnmpCommunity(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchSnmpCommunity resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerManagedSwitchSnmpCommunity(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchSnmpCommunity resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerManagedSwitchSnmpCommunityRead(d, m)
}

func resourceSwitchControllerManagedSwitchSnmpCommunityUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerManagedSwitchSnmpCommunity(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSnmpCommunity resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchSnmpCommunity(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSnmpCommunity resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerManagedSwitchSnmpCommunityRead(d, m)
}

func resourceSwitchControllerManagedSwitchSnmpCommunityDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerManagedSwitchSnmpCommunity(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchSnmpCommunity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchSnmpCommunityRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerManagedSwitchSnmpCommunity(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSnmpCommunity resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchSnmpCommunity(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSnmpCommunity resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchSnmpCommunityEvents2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchSnmpCommunityHosts2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSwitchControllerManagedSwitchSnmpCommunityHostsId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchSnmpCommunity-Hosts-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSwitchControllerManagedSwitchSnmpCommunityHostsIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchSnmpCommunity-Hosts-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchSnmpCommunityHostsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityHostsIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchSnmpCommunityId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Port2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchSnmpCommunity(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("events", flattenSwitchControllerManagedSwitchSnmpCommunityEvents2edl(o["events"], d, "events")); err != nil {
		if vv, ok := fortiAPIPatch(o["events"], "SwitchControllerManagedSwitchSnmpCommunity-Events"); ok {
			if err = d.Set("events", vv); err != nil {
				return fmt.Errorf("Error reading events: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("hosts", flattenSwitchControllerManagedSwitchSnmpCommunityHosts2edl(o["hosts"], d, "hosts")); err != nil {
			if vv, ok := fortiAPIPatch(o["hosts"], "SwitchControllerManagedSwitchSnmpCommunity-Hosts"); ok {
				if err = d.Set("hosts", vv); err != nil {
					return fmt.Errorf("Error reading hosts: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading hosts: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hosts"); ok {
			if err = d.Set("hosts", flattenSwitchControllerManagedSwitchSnmpCommunityHosts2edl(o["hosts"], d, "hosts")); err != nil {
				if vv, ok := fortiAPIPatch(o["hosts"], "SwitchControllerManagedSwitchSnmpCommunity-Hosts"); ok {
					if err = d.Set("hosts", vv); err != nil {
						return fmt.Errorf("Error reading hosts: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading hosts: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenSwitchControllerManagedSwitchSnmpCommunityId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SwitchControllerManagedSwitchSnmpCommunity-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerManagedSwitchSnmpCommunityName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerManagedSwitchSnmpCommunity-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("query_v1_port", flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Port2edl(o["query-v1-port"], d, "query_v1_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v1-port"], "SwitchControllerManagedSwitchSnmpCommunity-QueryV1Port"); ok {
			if err = d.Set("query_v1_port", vv); err != nil {
				return fmt.Errorf("Error reading query_v1_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v1_port: %v", err)
		}
	}

	if err = d.Set("query_v1_status", flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Status2edl(o["query-v1-status"], d, "query_v1_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v1-status"], "SwitchControllerManagedSwitchSnmpCommunity-QueryV1Status"); ok {
			if err = d.Set("query_v1_status", vv); err != nil {
				return fmt.Errorf("Error reading query_v1_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v1_status: %v", err)
		}
	}

	if err = d.Set("query_v2c_port", flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort2edl(o["query-v2c-port"], d, "query_v2c_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v2c-port"], "SwitchControllerManagedSwitchSnmpCommunity-QueryV2CPort"); ok {
			if err = d.Set("query_v2c_port", vv); err != nil {
				return fmt.Errorf("Error reading query_v2c_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v2c_port: %v", err)
		}
	}

	if err = d.Set("query_v2c_status", flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus2edl(o["query-v2c-status"], d, "query_v2c_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v2c-status"], "SwitchControllerManagedSwitchSnmpCommunity-QueryV2CStatus"); ok {
			if err = d.Set("query_v2c_status", vv); err != nil {
				return fmt.Errorf("Error reading query_v2c_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v2c_status: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchControllerManagedSwitchSnmpCommunityStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SwitchControllerManagedSwitchSnmpCommunity-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_v1_lport", flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport2edl(o["trap-v1-lport"], d, "trap_v1_lport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v1-lport"], "SwitchControllerManagedSwitchSnmpCommunity-TrapV1Lport"); ok {
			if err = d.Set("trap_v1_lport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v1_lport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v1_lport: %v", err)
		}
	}

	if err = d.Set("trap_v1_rport", flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport2edl(o["trap-v1-rport"], d, "trap_v1_rport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v1-rport"], "SwitchControllerManagedSwitchSnmpCommunity-TrapV1Rport"); ok {
			if err = d.Set("trap_v1_rport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v1_rport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v1_rport: %v", err)
		}
	}

	if err = d.Set("trap_v1_status", flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Status2edl(o["trap-v1-status"], d, "trap_v1_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v1-status"], "SwitchControllerManagedSwitchSnmpCommunity-TrapV1Status"); ok {
			if err = d.Set("trap_v1_status", vv); err != nil {
				return fmt.Errorf("Error reading trap_v1_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v1_status: %v", err)
		}
	}

	if err = d.Set("trap_v2c_lport", flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport2edl(o["trap-v2c-lport"], d, "trap_v2c_lport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v2c-lport"], "SwitchControllerManagedSwitchSnmpCommunity-TrapV2CLport"); ok {
			if err = d.Set("trap_v2c_lport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v2c_lport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v2c_lport: %v", err)
		}
	}

	if err = d.Set("trap_v2c_rport", flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport2edl(o["trap-v2c-rport"], d, "trap_v2c_rport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v2c-rport"], "SwitchControllerManagedSwitchSnmpCommunity-TrapV2CRport"); ok {
			if err = d.Set("trap_v2c_rport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v2c_rport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v2c_rport: %v", err)
		}
	}

	if err = d.Set("trap_v2c_status", flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus2edl(o["trap-v2c-status"], d, "trap_v2c_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v2c-status"], "SwitchControllerManagedSwitchSnmpCommunity-TrapV2CStatus"); ok {
			if err = d.Set("trap_v2c_status", vv); err != nil {
				return fmt.Errorf("Error reading trap_v2c_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v2c_status: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchSnmpCommunityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchSnmpCommunityEvents2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityHosts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSwitchControllerManagedSwitchSnmpCommunityHostsId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSwitchControllerManagedSwitchSnmpCommunityHostsIp2edl(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityHostsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityHostsIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityQueryV1Port2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityQueryV1Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchSnmpCommunity(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("events"); ok || d.HasChange("events") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityEvents2edl(d, v, "events")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["events"] = t
		}
	}

	if v, ok := d.GetOk("hosts"); ok || d.HasChange("hosts") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityHosts2edl(d, v, "hosts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosts"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("query_v1_port"); ok || d.HasChange("query_v1_port") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityQueryV1Port2edl(d, v, "query_v1_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v1-port"] = t
		}
	}

	if v, ok := d.GetOk("query_v1_status"); ok || d.HasChange("query_v1_status") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityQueryV1Status2edl(d, v, "query_v1_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v1-status"] = t
		}
	}

	if v, ok := d.GetOk("query_v2c_port"); ok || d.HasChange("query_v2c_port") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort2edl(d, v, "query_v2c_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v2c-port"] = t
		}
	}

	if v, ok := d.GetOk("query_v2c_status"); ok || d.HasChange("query_v2c_status") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus2edl(d, v, "query_v2c_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v2c-status"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_lport"); ok || d.HasChange("trap_v1_lport") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport2edl(d, v, "trap_v1_lport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-lport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_rport"); ok || d.HasChange("trap_v1_rport") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport2edl(d, v, "trap_v1_rport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-rport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_status"); ok || d.HasChange("trap_v1_status") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Status2edl(d, v, "trap_v1_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-status"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_lport"); ok || d.HasChange("trap_v2c_lport") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport2edl(d, v, "trap_v2c_lport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-lport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_rport"); ok || d.HasChange("trap_v2c_rport") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport2edl(d, v, "trap_v2c_rport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-rport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_status"); ok || d.HasChange("trap_v2c_status") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus2edl(d, v, "trap_v2c_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-status"] = t
		}
	}

	return &obj, nil
}
