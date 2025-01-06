// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SNMP community configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSnmpCommunity() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpCommunityCreate,
		Read:   resourceSystemSnmpCommunityRead,
		Update: resourceSystemSnmpCommunityUpdate,
		Delete: resourceSystemSnmpCommunityDelete,

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
						"ha_direct": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"host_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"interface_select_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"hosts6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ha_direct": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"host_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"interface_select_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_ipv6": &schema.Schema{
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
			"mib_view": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
			"vdoms": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceSystemSnmpCommunityCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemSnmpCommunity(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunity resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSnmpCommunity(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunity resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSnmpCommunityRead(d, m)
}

func resourceSystemSnmpCommunityUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemSnmpCommunity(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunity resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSnmpCommunity(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunity resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSnmpCommunityRead(d, m)
}

func resourceSystemSnmpCommunityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemSnmpCommunity(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpCommunity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpCommunityRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemSnmpCommunity(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunity resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpCommunity(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunity resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpCommunityEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSnmpCommunityHosts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_direct"
		if _, ok := i["ha-direct"]; ok {
			v := flattenSystemSnmpCommunityHostsHaDirect(i["ha-direct"], d, pre_append)
			tmp["ha_direct"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts-HaDirect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if _, ok := i["host-type"]; ok {
			v := flattenSystemSnmpCommunityHostsHostType(i["host-type"], d, pre_append)
			tmp["host_type"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts-HostType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSnmpCommunityHostsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystemSnmpCommunityHostsInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenSystemSnmpCommunityHostsInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemSnmpCommunityHostsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenSystemSnmpCommunityHostsSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts-SourceIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSnmpCommunityHostsHaDirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHostsHostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHostsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHostsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSnmpCommunityHostsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHostsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSnmpCommunityHostsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_direct"
		if _, ok := i["ha-direct"]; ok {
			v := flattenSystemSnmpCommunityHosts6HaDirect(i["ha-direct"], d, pre_append)
			tmp["ha_direct"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts6-HaDirect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if _, ok := i["host-type"]; ok {
			v := flattenSystemSnmpCommunityHosts6HostType(i["host-type"], d, pre_append)
			tmp["host_type"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts6-HostType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSnmpCommunityHosts6Id(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts6-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystemSnmpCommunityHosts6Interface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts6-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenSystemSnmpCommunityHosts6InterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts6-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if _, ok := i["ipv6"]; ok {
			v := flattenSystemSnmpCommunityHosts6Ipv6(i["ipv6"], d, pre_append)
			tmp["ipv6"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts6-Ipv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ipv6"
		if _, ok := i["source-ipv6"]; ok {
			v := flattenSystemSnmpCommunityHosts6SourceIpv6(i["source-ipv6"], d, pre_append)
			tmp["source_ipv6"] = fortiAPISubPartPatch(v, "SystemSnmpCommunity-Hosts6-SourceIpv6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSnmpCommunityHosts6HaDirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6HostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSnmpCommunityHosts6InterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6Ipv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6SourceIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityMibView(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSnmpCommunityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityQueryV1Port(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityQueryV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityQueryV2CPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityQueryV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityTrapV1Lport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityTrapV1Rport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityTrapV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityTrapV2CLport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityTrapV2CRport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityTrapV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityVdoms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemSnmpCommunity(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("events", flattenSystemSnmpCommunityEvents(o["events"], d, "events")); err != nil {
		if vv, ok := fortiAPIPatch(o["events"], "SystemSnmpCommunity-Events"); ok {
			if err = d.Set("events", vv); err != nil {
				return fmt.Errorf("Error reading events: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("hosts", flattenSystemSnmpCommunityHosts(o["hosts"], d, "hosts")); err != nil {
			if vv, ok := fortiAPIPatch(o["hosts"], "SystemSnmpCommunity-Hosts"); ok {
				if err = d.Set("hosts", vv); err != nil {
					return fmt.Errorf("Error reading hosts: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading hosts: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hosts"); ok {
			if err = d.Set("hosts", flattenSystemSnmpCommunityHosts(o["hosts"], d, "hosts")); err != nil {
				if vv, ok := fortiAPIPatch(o["hosts"], "SystemSnmpCommunity-Hosts"); ok {
					if err = d.Set("hosts", vv); err != nil {
						return fmt.Errorf("Error reading hosts: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading hosts: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("hosts6", flattenSystemSnmpCommunityHosts6(o["hosts6"], d, "hosts6")); err != nil {
			if vv, ok := fortiAPIPatch(o["hosts6"], "SystemSnmpCommunity-Hosts6"); ok {
				if err = d.Set("hosts6", vv); err != nil {
					return fmt.Errorf("Error reading hosts6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading hosts6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hosts6"); ok {
			if err = d.Set("hosts6", flattenSystemSnmpCommunityHosts6(o["hosts6"], d, "hosts6")); err != nil {
				if vv, ok := fortiAPIPatch(o["hosts6"], "SystemSnmpCommunity-Hosts6"); ok {
					if err = d.Set("hosts6", vv); err != nil {
						return fmt.Errorf("Error reading hosts6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading hosts6: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenSystemSnmpCommunityId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSnmpCommunity-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("mib_view", flattenSystemSnmpCommunityMibView(o["mib-view"], d, "mib_view")); err != nil {
		if vv, ok := fortiAPIPatch(o["mib-view"], "SystemSnmpCommunity-MibView"); ok {
			if err = d.Set("mib_view", vv); err != nil {
				return fmt.Errorf("Error reading mib_view: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mib_view: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSnmpCommunityName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSnmpCommunity-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("query_v1_port", flattenSystemSnmpCommunityQueryV1Port(o["query-v1-port"], d, "query_v1_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v1-port"], "SystemSnmpCommunity-QueryV1Port"); ok {
			if err = d.Set("query_v1_port", vv); err != nil {
				return fmt.Errorf("Error reading query_v1_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v1_port: %v", err)
		}
	}

	if err = d.Set("query_v1_status", flattenSystemSnmpCommunityQueryV1Status(o["query-v1-status"], d, "query_v1_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v1-status"], "SystemSnmpCommunity-QueryV1Status"); ok {
			if err = d.Set("query_v1_status", vv); err != nil {
				return fmt.Errorf("Error reading query_v1_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v1_status: %v", err)
		}
	}

	if err = d.Set("query_v2c_port", flattenSystemSnmpCommunityQueryV2CPort(o["query-v2c-port"], d, "query_v2c_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v2c-port"], "SystemSnmpCommunity-QueryV2CPort"); ok {
			if err = d.Set("query_v2c_port", vv); err != nil {
				return fmt.Errorf("Error reading query_v2c_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v2c_port: %v", err)
		}
	}

	if err = d.Set("query_v2c_status", flattenSystemSnmpCommunityQueryV2CStatus(o["query-v2c-status"], d, "query_v2c_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-v2c-status"], "SystemSnmpCommunity-QueryV2CStatus"); ok {
			if err = d.Set("query_v2c_status", vv); err != nil {
				return fmt.Errorf("Error reading query_v2c_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_v2c_status: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSnmpCommunityStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSnmpCommunity-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_v1_lport", flattenSystemSnmpCommunityTrapV1Lport(o["trap-v1-lport"], d, "trap_v1_lport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v1-lport"], "SystemSnmpCommunity-TrapV1Lport"); ok {
			if err = d.Set("trap_v1_lport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v1_lport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v1_lport: %v", err)
		}
	}

	if err = d.Set("trap_v1_rport", flattenSystemSnmpCommunityTrapV1Rport(o["trap-v1-rport"], d, "trap_v1_rport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v1-rport"], "SystemSnmpCommunity-TrapV1Rport"); ok {
			if err = d.Set("trap_v1_rport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v1_rport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v1_rport: %v", err)
		}
	}

	if err = d.Set("trap_v1_status", flattenSystemSnmpCommunityTrapV1Status(o["trap-v1-status"], d, "trap_v1_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v1-status"], "SystemSnmpCommunity-TrapV1Status"); ok {
			if err = d.Set("trap_v1_status", vv); err != nil {
				return fmt.Errorf("Error reading trap_v1_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v1_status: %v", err)
		}
	}

	if err = d.Set("trap_v2c_lport", flattenSystemSnmpCommunityTrapV2CLport(o["trap-v2c-lport"], d, "trap_v2c_lport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v2c-lport"], "SystemSnmpCommunity-TrapV2CLport"); ok {
			if err = d.Set("trap_v2c_lport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v2c_lport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v2c_lport: %v", err)
		}
	}

	if err = d.Set("trap_v2c_rport", flattenSystemSnmpCommunityTrapV2CRport(o["trap-v2c-rport"], d, "trap_v2c_rport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v2c-rport"], "SystemSnmpCommunity-TrapV2CRport"); ok {
			if err = d.Set("trap_v2c_rport", vv); err != nil {
				return fmt.Errorf("Error reading trap_v2c_rport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v2c_rport: %v", err)
		}
	}

	if err = d.Set("trap_v2c_status", flattenSystemSnmpCommunityTrapV2CStatus(o["trap-v2c-status"], d, "trap_v2c_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-v2c-status"], "SystemSnmpCommunity-TrapV2CStatus"); ok {
			if err = d.Set("trap_v2c_status", vv); err != nil {
				return fmt.Errorf("Error reading trap_v2c_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_v2c_status: %v", err)
		}
	}

	if err = d.Set("vdoms", flattenSystemSnmpCommunityVdoms(o["vdoms"], d, "vdoms")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdoms"], "SystemSnmpCommunity-Vdoms"); ok {
			if err = d.Set("vdoms", vv); err != nil {
				return fmt.Errorf("Error reading vdoms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdoms: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpCommunityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSnmpCommunityEvents(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSnmpCommunityHosts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_direct"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-direct"], _ = expandSystemSnmpCommunityHostsHaDirect(d, i["ha_direct"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host-type"], _ = expandSystemSnmpCommunityHostsHostType(d, i["host_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSnmpCommunityHostsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystemSnmpCommunityHostsInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-select-method"], _ = expandSystemSnmpCommunityHostsInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemSnmpCommunityHostsIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip"], _ = expandSystemSnmpCommunityHostsSourceIp(d, i["source_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSnmpCommunityHostsHaDirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsHostType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSnmpCommunityHostsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemSnmpCommunityHostsSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_direct"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-direct"], _ = expandSystemSnmpCommunityHosts6HaDirect(d, i["ha_direct"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host-type"], _ = expandSystemSnmpCommunityHosts6HostType(d, i["host_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSnmpCommunityHosts6Id(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystemSnmpCommunityHosts6Interface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-select-method"], _ = expandSystemSnmpCommunityHosts6InterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6"], _ = expandSystemSnmpCommunityHosts6Ipv6(d, i["ipv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ipv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ipv6"], _ = expandSystemSnmpCommunityHosts6SourceIpv6(d, i["source_ipv6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSnmpCommunityHosts6HaDirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6HostType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSnmpCommunityHosts6InterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6Ipv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6SourceIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityMibView(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSnmpCommunityName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityQueryV1Port(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityQueryV1Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityQueryV2CPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityQueryV2CStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV1Lport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV1Rport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV1Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV2CLport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV2CRport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV2CStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityVdoms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemSnmpCommunity(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("events"); ok || d.HasChange("events") {
		t, err := expandSystemSnmpCommunityEvents(d, v, "events")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["events"] = t
		}
	}

	if v, ok := d.GetOk("hosts"); ok || d.HasChange("hosts") {
		t, err := expandSystemSnmpCommunityHosts(d, v, "hosts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosts"] = t
		}
	}

	if v, ok := d.GetOk("hosts6"); ok || d.HasChange("hosts6") {
		t, err := expandSystemSnmpCommunityHosts6(d, v, "hosts6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosts6"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSnmpCommunityId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("mib_view"); ok || d.HasChange("mib_view") {
		t, err := expandSystemSnmpCommunityMibView(d, v, "mib_view")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mib-view"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSnmpCommunityName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("query_v1_port"); ok || d.HasChange("query_v1_port") {
		t, err := expandSystemSnmpCommunityQueryV1Port(d, v, "query_v1_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v1-port"] = t
		}
	}

	if v, ok := d.GetOk("query_v1_status"); ok || d.HasChange("query_v1_status") {
		t, err := expandSystemSnmpCommunityQueryV1Status(d, v, "query_v1_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v1-status"] = t
		}
	}

	if v, ok := d.GetOk("query_v2c_port"); ok || d.HasChange("query_v2c_port") {
		t, err := expandSystemSnmpCommunityQueryV2CPort(d, v, "query_v2c_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v2c-port"] = t
		}
	}

	if v, ok := d.GetOk("query_v2c_status"); ok || d.HasChange("query_v2c_status") {
		t, err := expandSystemSnmpCommunityQueryV2CStatus(d, v, "query_v2c_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v2c-status"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSnmpCommunityStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_lport"); ok || d.HasChange("trap_v1_lport") {
		t, err := expandSystemSnmpCommunityTrapV1Lport(d, v, "trap_v1_lport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-lport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_rport"); ok || d.HasChange("trap_v1_rport") {
		t, err := expandSystemSnmpCommunityTrapV1Rport(d, v, "trap_v1_rport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-rport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_status"); ok || d.HasChange("trap_v1_status") {
		t, err := expandSystemSnmpCommunityTrapV1Status(d, v, "trap_v1_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-status"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_lport"); ok || d.HasChange("trap_v2c_lport") {
		t, err := expandSystemSnmpCommunityTrapV2CLport(d, v, "trap_v2c_lport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-lport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_rport"); ok || d.HasChange("trap_v2c_rport") {
		t, err := expandSystemSnmpCommunityTrapV2CRport(d, v, "trap_v2c_rport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-rport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_status"); ok || d.HasChange("trap_v2c_status") {
		t, err := expandSystemSnmpCommunityTrapV2CStatus(d, v, "trap_v2c_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-status"] = t
		}
	}

	if v, ok := d.GetOk("vdoms"); ok || d.HasChange("vdoms") {
		t, err := expandSystemSnmpCommunityVdoms(d, v, "vdoms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdoms"] = t
		}
	}

	return &obj, nil
}
