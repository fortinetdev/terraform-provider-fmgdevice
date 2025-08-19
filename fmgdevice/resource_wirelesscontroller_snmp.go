// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SNMP.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerSnmp() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerSnmpUpdate,
		Read:   resourceWirelessControllerSnmpRead,
		Update: resourceWirelessControllerSnmpUpdate,
		Delete: resourceWirelessControllerSnmpDelete,

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
			"community": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
					},
				},
			},
			"contact_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"engine_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trap_high_cpu_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_high_mem_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_pwd": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"notify_hosts": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"priv_proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priv_pwd": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"queries": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"security_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trap_status": &schema.Schema{
							Type:     schema.TypeString,
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

func resourceWirelessControllerSnmpUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerSnmp(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSnmp resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerSnmp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSnmp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerSnmp")

	return resourceWirelessControllerSnmpRead(d, m)
}

func resourceWirelessControllerSnmpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerSnmp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerSnmp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSnmpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerSnmp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSnmp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerSnmp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSnmp resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerSnmpCommunity(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hosts"
		if _, ok := i["hosts"]; ok {
			v := flattenWirelessControllerSnmpCommunityHosts(i["hosts"], d, pre_append)
			tmp["hosts"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-Community-Hosts")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWirelessControllerSnmpCommunityId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-Community-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWirelessControllerSnmpCommunityName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-Community-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_status"
		if _, ok := i["query-v1-status"]; ok {
			v := flattenWirelessControllerSnmpCommunityQueryV1Status(i["query-v1-status"], d, pre_append)
			tmp["query_v1_status"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-Community-QueryV1Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_status"
		if _, ok := i["query-v2c-status"]; ok {
			v := flattenWirelessControllerSnmpCommunityQueryV2CStatus(i["query-v2c-status"], d, pre_append)
			tmp["query_v2c_status"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-Community-QueryV2CStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWirelessControllerSnmpCommunityStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-Community-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_status"
		if _, ok := i["trap-v1-status"]; ok {
			v := flattenWirelessControllerSnmpCommunityTrapV1Status(i["trap-v1-status"], d, pre_append)
			tmp["trap_v1_status"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-Community-TrapV1Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_status"
		if _, ok := i["trap-v2c-status"]; ok {
			v := flattenWirelessControllerSnmpCommunityTrapV2CStatus(i["trap-v2c-status"], d, pre_append)
			tmp["trap_v2c_status"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-Community-TrapV2CStatus")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerSnmpCommunityHosts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWirelessControllerSnmpCommunityHostsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerSnmpCommunity-Hosts-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenWirelessControllerSnmpCommunityHostsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "WirelessControllerSnmpCommunity-Hosts-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerSnmpCommunityHostsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityHostsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityQueryV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityQueryV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityTrapV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityTrapV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpContactInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpEngineId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpTrapHighCpuThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpTrapHighMemThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUser(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_proto"
		if _, ok := i["auth-proto"]; ok {
			v := flattenWirelessControllerSnmpUserAuthProto(i["auth-proto"], d, pre_append)
			tmp["auth_proto"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-User-AuthProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWirelessControllerSnmpUserName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-User-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "notify_hosts"
		if _, ok := i["notify-hosts"]; ok {
			v := flattenWirelessControllerSnmpUserNotifyHosts(i["notify-hosts"], d, pre_append)
			tmp["notify_hosts"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-User-NotifyHosts")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_proto"
		if _, ok := i["priv-proto"]; ok {
			v := flattenWirelessControllerSnmpUserPrivProto(i["priv-proto"], d, pre_append)
			tmp["priv_proto"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-User-PrivProto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queries"
		if _, ok := i["queries"]; ok {
			v := flattenWirelessControllerSnmpUserQueries(i["queries"], d, pre_append)
			tmp["queries"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-User-Queries")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_level"
		if _, ok := i["security-level"]; ok {
			v := flattenWirelessControllerSnmpUserSecurityLevel(i["security-level"], d, pre_append)
			tmp["security_level"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-User-SecurityLevel")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWirelessControllerSnmpUserStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-User-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_status"
		if _, ok := i["trap-status"]; ok {
			v := flattenWirelessControllerSnmpUserTrapStatus(i["trap-status"], d, pre_append)
			tmp["trap_status"] = fortiAPISubPartPatch(v, "WirelessControllerSnmp-User-TrapStatus")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerSnmpUserAuthProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserNotifyHosts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerSnmpUserPrivProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserQueries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserSecurityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserTrapStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerSnmp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("community", flattenWirelessControllerSnmpCommunity(o["community"], d, "community")); err != nil {
			if vv, ok := fortiAPIPatch(o["community"], "WirelessControllerSnmp-Community"); ok {
				if err = d.Set("community", vv); err != nil {
					return fmt.Errorf("Error reading community: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading community: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("community"); ok {
			if err = d.Set("community", flattenWirelessControllerSnmpCommunity(o["community"], d, "community")); err != nil {
				if vv, ok := fortiAPIPatch(o["community"], "WirelessControllerSnmp-Community"); ok {
					if err = d.Set("community", vv); err != nil {
						return fmt.Errorf("Error reading community: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading community: %v", err)
				}
			}
		}
	}

	if err = d.Set("contact_info", flattenWirelessControllerSnmpContactInfo(o["contact-info"], d, "contact_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["contact-info"], "WirelessControllerSnmp-ContactInfo"); ok {
			if err = d.Set("contact_info", vv); err != nil {
				return fmt.Errorf("Error reading contact_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading contact_info: %v", err)
		}
	}

	if err = d.Set("engine_id", flattenWirelessControllerSnmpEngineId(o["engine-id"], d, "engine_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["engine-id"], "WirelessControllerSnmp-EngineId"); ok {
			if err = d.Set("engine_id", vv); err != nil {
				return fmt.Errorf("Error reading engine_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading engine_id: %v", err)
		}
	}

	if err = d.Set("trap_high_cpu_threshold", flattenWirelessControllerSnmpTrapHighCpuThreshold(o["trap-high-cpu-threshold"], d, "trap_high_cpu_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-high-cpu-threshold"], "WirelessControllerSnmp-TrapHighCpuThreshold"); ok {
			if err = d.Set("trap_high_cpu_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
		}
	}

	if err = d.Set("trap_high_mem_threshold", flattenWirelessControllerSnmpTrapHighMemThreshold(o["trap-high-mem-threshold"], d, "trap_high_mem_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-high-mem-threshold"], "WirelessControllerSnmp-TrapHighMemThreshold"); ok {
			if err = d.Set("trap_high_mem_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_high_mem_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_high_mem_threshold: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("user", flattenWirelessControllerSnmpUser(o["user"], d, "user")); err != nil {
			if vv, ok := fortiAPIPatch(o["user"], "WirelessControllerSnmp-User"); ok {
				if err = d.Set("user", vv); err != nil {
					return fmt.Errorf("Error reading user: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading user: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("user"); ok {
			if err = d.Set("user", flattenWirelessControllerSnmpUser(o["user"], d, "user")); err != nil {
				if vv, ok := fortiAPIPatch(o["user"], "WirelessControllerSnmp-User"); ok {
					if err = d.Set("user", vv); err != nil {
						return fmt.Errorf("Error reading user: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading user: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerSnmpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerSnmpCommunity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hosts"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWirelessControllerSnmpCommunityHosts(d, i["hosts"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["hosts"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWirelessControllerSnmpCommunityId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWirelessControllerSnmpCommunityName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["query-v1-status"], _ = expandWirelessControllerSnmpCommunityQueryV1Status(d, i["query_v1_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["query-v2c-status"], _ = expandWirelessControllerSnmpCommunityQueryV2CStatus(d, i["query_v2c_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWirelessControllerSnmpCommunityStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["trap-v1-status"], _ = expandWirelessControllerSnmpCommunityTrapV1Status(d, i["trap_v1_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["trap-v2c-status"], _ = expandWirelessControllerSnmpCommunityTrapV2CStatus(d, i["trap_v2c_status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerSnmpCommunityHosts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerSnmpCommunityHostsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandWirelessControllerSnmpCommunityHostsIp(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerSnmpCommunityHostsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityHostsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityQueryV1Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityQueryV2CStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityTrapV1Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityTrapV2CStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpContactInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpEngineId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpTrapHighCpuThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpTrapHighMemThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-proto"], _ = expandWirelessControllerSnmpUserAuthProto(d, i["auth_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_pwd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-pwd"], _ = expandWirelessControllerSnmpUserAuthPwd(d, i["auth_pwd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWirelessControllerSnmpUserName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "notify_hosts"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["notify-hosts"], _ = expandWirelessControllerSnmpUserNotifyHosts(d, i["notify_hosts"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priv-proto"], _ = expandWirelessControllerSnmpUserPrivProto(d, i["priv_proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_pwd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priv-pwd"], _ = expandWirelessControllerSnmpUserPrivPwd(d, i["priv_pwd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queries"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queries"], _ = expandWirelessControllerSnmpUserQueries(d, i["queries"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_level"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-level"], _ = expandWirelessControllerSnmpUserSecurityLevel(d, i["security_level"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWirelessControllerSnmpUserStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["trap-status"], _ = expandWirelessControllerSnmpUserTrapStatus(d, i["trap_status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerSnmpUserAuthProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserAuthPwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerSnmpUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserNotifyHosts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerSnmpUserPrivProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserPrivPwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerSnmpUserQueries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserSecurityLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserTrapStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerSnmp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("community"); ok || d.HasChange("community") {
		t, err := expandWirelessControllerSnmpCommunity(d, v, "community")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["community"] = t
		}
	}

	if v, ok := d.GetOk("contact_info"); ok || d.HasChange("contact_info") {
		t, err := expandWirelessControllerSnmpContactInfo(d, v, "contact_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contact-info"] = t
		}
	}

	if v, ok := d.GetOk("engine_id"); ok || d.HasChange("engine_id") {
		t, err := expandWirelessControllerSnmpEngineId(d, v, "engine_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["engine-id"] = t
		}
	}

	if v, ok := d.GetOk("trap_high_cpu_threshold"); ok || d.HasChange("trap_high_cpu_threshold") {
		t, err := expandWirelessControllerSnmpTrapHighCpuThreshold(d, v, "trap_high_cpu_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-high-cpu-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_high_mem_threshold"); ok || d.HasChange("trap_high_mem_threshold") {
		t, err := expandWirelessControllerSnmpTrapHighMemThreshold(d, v, "trap_high_mem_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-high-mem-threshold"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandWirelessControllerSnmpUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	return &obj, nil
}
