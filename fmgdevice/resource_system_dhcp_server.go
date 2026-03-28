// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DHCP servers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDhcpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcpServerCreate,
		Read:   resourceSystemDhcpServerRead,
		Update: resourceSystemDhcpServerUpdate,
		Delete: resourceSystemDhcpServerDelete,

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
			"auto_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_managed_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conflicted_ip_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ddns_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ddns_keyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ddns_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_update_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_settings_from_fortiipam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"exclude_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"lease_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forticlient_on_net_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"lease_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ipsec_lease_hold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"lease_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mac_acl_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"next_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"uci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"uci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"relay_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reserved_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"circuit_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"circuit_id_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"remote_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"remote_id_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shared_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tftp_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"timezone_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vci_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vci_string": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wifi_ac_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wins_server2": &schema.Schema{
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

func resourceSystemDhcpServerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcpServer(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcpServer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemDhcpServer(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemDhcpServer(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemDhcpServer resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemDhcpServer(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemDhcpServer resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpServerRead(d, m)
}

func resourceSystemDhcpServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcpServer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemDhcpServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpServerRead(d, m)
}

func resourceSystemDhcpServerDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemDhcpServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcpServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcpServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDhcpServer(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemDhcpServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcpServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcpServer resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcpServerAutoConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerAutoManagedStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerConflictedIpTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsKeyname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsUpdateOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDefaultGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDhcpSettingsFromFortiipam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDnsServer4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDnsService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenSystemDhcpServerExcludeRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ExcludeRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemDhcpServerExcludeRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := i["lease-time"]; ok {
			v := flattenSystemDhcpServerExcludeRangeLeaseTime(i["lease-time"], d, pre_append)
			tmp["lease_time"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ExcludeRange-LeaseTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenSystemDhcpServerExcludeRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ExcludeRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenSystemDhcpServerExcludeRangeUciMatch(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ExcludeRange-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenSystemDhcpServerExcludeRangeUciString(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ExcludeRange-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenSystemDhcpServerExcludeRangeVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ExcludeRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenSystemDhcpServerExcludeRangeVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ExcludeRange-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemDhcpServerExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeUciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeUciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerExcludeRangeVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerFilename(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerForticlientOnNetStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenSystemDhcpServerIpRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "SystemDhcpServer-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemDhcpServerIpRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemDhcpServer-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := i["lease-time"]; ok {
			v := flattenSystemDhcpServerIpRangeLeaseTime(i["lease-time"], d, pre_append)
			tmp["lease_time"] = fortiAPISubPartPatch(v, "SystemDhcpServer-IpRange-LeaseTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenSystemDhcpServerIpRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "SystemDhcpServer-IpRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenSystemDhcpServerIpRangeUciMatch(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "SystemDhcpServer-IpRange-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenSystemDhcpServerIpRangeUciString(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "SystemDhcpServer-IpRange-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenSystemDhcpServerIpRangeVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "SystemDhcpServer-IpRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenSystemDhcpServerIpRangeVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "SystemDhcpServer-IpRange-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemDhcpServerIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerIpRangeLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerIpRangeUciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerIpRangeUciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerIpRangeVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerIpRangeVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerIpsecLeaseHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerMacAclDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerNextServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerNtpServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerNtpServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerNtpServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerNtpService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := i["code"]; ok {
			v := flattenSystemDhcpServerOptionsCode(i["code"], d, pre_append)
			tmp["code"] = fortiAPISubPartPatch(v, "SystemDhcpServer-Options-Code")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemDhcpServerOptionsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemDhcpServer-Options-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemDhcpServerOptionsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemDhcpServer-Options-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSystemDhcpServerOptionsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SystemDhcpServer-Options-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenSystemDhcpServerOptionsUciMatch(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "SystemDhcpServer-Options-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenSystemDhcpServerOptionsUciString(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "SystemDhcpServer-Options-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenSystemDhcpServerOptionsValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "SystemDhcpServer-Options-Value")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenSystemDhcpServerOptionsVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "SystemDhcpServer-Options-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenSystemDhcpServerOptionsVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "SystemDhcpServer-Options-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemDhcpServerOptionsCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerOptionsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsUciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsUciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerOptionsValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerRelayAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenSystemDhcpServerReservedAddressAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ReservedAddress-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {
			v := flattenSystemDhcpServerReservedAddressCircuitId(i["circuit-id"], d, pre_append)
			tmp["circuit_id"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ReservedAddress-CircuitId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := i["circuit-id-type"]; ok {
			v := flattenSystemDhcpServerReservedAddressCircuitIdType(i["circuit-id-type"], d, pre_append)
			tmp["circuit_id_type"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ReservedAddress-CircuitIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenSystemDhcpServerReservedAddressDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ReservedAddress-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemDhcpServerReservedAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ReservedAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemDhcpServerReservedAddressIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ReservedAddress-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenSystemDhcpServerReservedAddressMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ReservedAddress-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			v := flattenSystemDhcpServerReservedAddressRemoteId(i["remote-id"], d, pre_append)
			tmp["remote_id"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ReservedAddress-RemoteId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := i["remote-id-type"]; ok {
			v := flattenSystemDhcpServerReservedAddressRemoteIdType(i["remote-id-type"], d, pre_append)
			tmp["remote_id_type"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ReservedAddress-RemoteIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSystemDhcpServerReservedAddressType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SystemDhcpServer-ReservedAddress-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemDhcpServerReservedAddressAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressCircuitId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressCircuitIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressRemoteId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressRemoteIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerSharedSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerTftpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerTimezone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenSystemDhcpServerTimezoneOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerWifiAcService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerWifiAc1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerWifiAc2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerWifiAc3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerWinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerWinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDhcpServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auto_configuration", flattenSystemDhcpServerAutoConfiguration(o["auto-configuration"], d, "auto_configuration")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-configuration"], "SystemDhcpServer-AutoConfiguration"); ok {
			if err = d.Set("auto_configuration", vv); err != nil {
				return fmt.Errorf("Error reading auto_configuration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_configuration: %v", err)
		}
	}

	if err = d.Set("auto_managed_status", flattenSystemDhcpServerAutoManagedStatus(o["auto-managed-status"], d, "auto_managed_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-managed-status"], "SystemDhcpServer-AutoManagedStatus"); ok {
			if err = d.Set("auto_managed_status", vv); err != nil {
				return fmt.Errorf("Error reading auto_managed_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_managed_status: %v", err)
		}
	}

	if err = d.Set("conflicted_ip_timeout", flattenSystemDhcpServerConflictedIpTimeout(o["conflicted-ip-timeout"], d, "conflicted_ip_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["conflicted-ip-timeout"], "SystemDhcpServer-ConflictedIpTimeout"); ok {
			if err = d.Set("conflicted_ip_timeout", vv); err != nil {
				return fmt.Errorf("Error reading conflicted_ip_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conflicted_ip_timeout: %v", err)
		}
	}

	if err = d.Set("ddns_auth", flattenSystemDhcpServerDdnsAuth(o["ddns-auth"], d, "ddns_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-auth"], "SystemDhcpServer-DdnsAuth"); ok {
			if err = d.Set("ddns_auth", vv); err != nil {
				return fmt.Errorf("Error reading ddns_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_auth: %v", err)
		}
	}

	if err = d.Set("ddns_keyname", flattenSystemDhcpServerDdnsKeyname(o["ddns-keyname"], d, "ddns_keyname")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-keyname"], "SystemDhcpServer-DdnsKeyname"); ok {
			if err = d.Set("ddns_keyname", vv); err != nil {
				return fmt.Errorf("Error reading ddns_keyname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_keyname: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenSystemDhcpServerDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-ip"], "SystemDhcpServer-DdnsServerIp"); ok {
			if err = d.Set("ddns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_ttl", flattenSystemDhcpServerDdnsTtl(o["ddns-ttl"], d, "ddns_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-ttl"], "SystemDhcpServer-DdnsTtl"); ok {
			if err = d.Set("ddns_ttl", vv); err != nil {
				return fmt.Errorf("Error reading ddns_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_ttl: %v", err)
		}
	}

	if err = d.Set("ddns_update", flattenSystemDhcpServerDdnsUpdate(o["ddns-update"], d, "ddns_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-update"], "SystemDhcpServer-DdnsUpdate"); ok {
			if err = d.Set("ddns_update", vv); err != nil {
				return fmt.Errorf("Error reading ddns_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_update: %v", err)
		}
	}

	if err = d.Set("ddns_update_override", flattenSystemDhcpServerDdnsUpdateOverride(o["ddns-update-override"], d, "ddns_update_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-update-override"], "SystemDhcpServer-DdnsUpdateOverride"); ok {
			if err = d.Set("ddns_update_override", vv); err != nil {
				return fmt.Errorf("Error reading ddns_update_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_update_override: %v", err)
		}
	}

	if err = d.Set("ddns_zone", flattenSystemDhcpServerDdnsZone(o["ddns-zone"], d, "ddns_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-zone"], "SystemDhcpServer-DdnsZone"); ok {
			if err = d.Set("ddns_zone", vv); err != nil {
				return fmt.Errorf("Error reading ddns_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_zone: %v", err)
		}
	}

	if err = d.Set("default_gateway", flattenSystemDhcpServerDefaultGateway(o["default-gateway"], d, "default_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-gateway"], "SystemDhcpServer-DefaultGateway"); ok {
			if err = d.Set("default_gateway", vv); err != nil {
				return fmt.Errorf("Error reading default_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_gateway: %v", err)
		}
	}

	if err = d.Set("dhcp_settings_from_fortiipam", flattenSystemDhcpServerDhcpSettingsFromFortiipam(o["dhcp-settings-from-fortiipam"], d, "dhcp_settings_from_fortiipam")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-settings-from-fortiipam"], "SystemDhcpServer-DhcpSettingsFromFortiipam"); ok {
			if err = d.Set("dhcp_settings_from_fortiipam", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_settings_from_fortiipam: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_settings_from_fortiipam: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenSystemDhcpServerDnsServer1(o["dns-server1"], d, "dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server1"], "SystemDhcpServer-DnsServer1"); ok {
			if err = d.Set("dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenSystemDhcpServerDnsServer2(o["dns-server2"], d, "dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server2"], "SystemDhcpServer-DnsServer2"); ok {
			if err = d.Set("dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_server3", flattenSystemDhcpServerDnsServer3(o["dns-server3"], d, "dns_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server3"], "SystemDhcpServer-DnsServer3"); ok {
			if err = d.Set("dns_server3", vv); err != nil {
				return fmt.Errorf("Error reading dns_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server3: %v", err)
		}
	}

	if err = d.Set("dns_server4", flattenSystemDhcpServerDnsServer4(o["dns-server4"], d, "dns_server4")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server4"], "SystemDhcpServer-DnsServer4"); ok {
			if err = d.Set("dns_server4", vv); err != nil {
				return fmt.Errorf("Error reading dns_server4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server4: %v", err)
		}
	}

	if err = d.Set("dns_service", flattenSystemDhcpServerDnsService(o["dns-service"], d, "dns_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-service"], "SystemDhcpServer-DnsService"); ok {
			if err = d.Set("dns_service", vv); err != nil {
				return fmt.Errorf("Error reading dns_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_service: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemDhcpServerDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "SystemDhcpServer-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("exclude_range", flattenSystemDhcpServerExcludeRange(o["exclude-range"], d, "exclude_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["exclude-range"], "SystemDhcpServer-ExcludeRange"); ok {
				if err = d.Set("exclude_range", vv); err != nil {
					return fmt.Errorf("Error reading exclude_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("exclude_range"); ok {
			if err = d.Set("exclude_range", flattenSystemDhcpServerExcludeRange(o["exclude-range"], d, "exclude_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["exclude-range"], "SystemDhcpServer-ExcludeRange"); ok {
					if err = d.Set("exclude_range", vv); err != nil {
						return fmt.Errorf("Error reading exclude_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("filename", flattenSystemDhcpServerFilename(o["filename"], d, "filename")); err != nil {
		if vv, ok := fortiAPIPatch(o["filename"], "SystemDhcpServer-Filename"); ok {
			if err = d.Set("filename", vv); err != nil {
				return fmt.Errorf("Error reading filename: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filename: %v", err)
		}
	}

	if err = d.Set("forticlient_on_net_status", flattenSystemDhcpServerForticlientOnNetStatus(o["forticlient-on-net-status"], d, "forticlient_on_net_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-on-net-status"], "SystemDhcpServer-ForticlientOnNetStatus"); ok {
			if err = d.Set("forticlient_on_net_status", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_on_net_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_on_net_status: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemDhcpServerId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemDhcpServer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemDhcpServerInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemDhcpServer-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenSystemDhcpServerIpMode(o["ip-mode"], d, "ip_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-mode"], "SystemDhcpServer-IpMode"); ok {
			if err = d.Set("ip_mode", vv); err != nil {
				return fmt.Errorf("Error reading ip_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_range", flattenSystemDhcpServerIpRange(o["ip-range"], d, "ip_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-range"], "SystemDhcpServer-IpRange"); ok {
				if err = d.Set("ip_range", vv); err != nil {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenSystemDhcpServerIpRange(o["ip-range"], d, "ip_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-range"], "SystemDhcpServer-IpRange"); ok {
					if err = d.Set("ip_range", vv); err != nil {
						return fmt.Errorf("Error reading ip_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipsec_lease_hold", flattenSystemDhcpServerIpsecLeaseHold(o["ipsec-lease-hold"], d, "ipsec_lease_hold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-lease-hold"], "SystemDhcpServer-IpsecLeaseHold"); ok {
			if err = d.Set("ipsec_lease_hold", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenSystemDhcpServerLeaseTime(o["lease-time"], d, "lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["lease-time"], "SystemDhcpServer-LeaseTime"); ok {
			if err = d.Set("lease_time", vv); err != nil {
				return fmt.Errorf("Error reading lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("mac_acl_default_action", flattenSystemDhcpServerMacAclDefaultAction(o["mac-acl-default-action"], d, "mac_acl_default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-acl-default-action"], "SystemDhcpServer-MacAclDefaultAction"); ok {
			if err = d.Set("mac_acl_default_action", vv); err != nil {
				return fmt.Errorf("Error reading mac_acl_default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_acl_default_action: %v", err)
		}
	}

	if err = d.Set("netmask", flattenSystemDhcpServerNetmask(o["netmask"], d, "netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["netmask"], "SystemDhcpServer-Netmask"); ok {
			if err = d.Set("netmask", vv); err != nil {
				return fmt.Errorf("Error reading netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netmask: %v", err)
		}
	}

	if err = d.Set("next_server", flattenSystemDhcpServerNextServer(o["next-server"], d, "next_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-server"], "SystemDhcpServer-NextServer"); ok {
			if err = d.Set("next_server", vv); err != nil {
				return fmt.Errorf("Error reading next_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_server: %v", err)
		}
	}

	if err = d.Set("ntp_server1", flattenSystemDhcpServerNtpServer1(o["ntp-server1"], d, "ntp_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server1"], "SystemDhcpServer-NtpServer1"); ok {
			if err = d.Set("ntp_server1", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server1: %v", err)
		}
	}

	if err = d.Set("ntp_server2", flattenSystemDhcpServerNtpServer2(o["ntp-server2"], d, "ntp_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server2"], "SystemDhcpServer-NtpServer2"); ok {
			if err = d.Set("ntp_server2", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server2: %v", err)
		}
	}

	if err = d.Set("ntp_server3", flattenSystemDhcpServerNtpServer3(o["ntp-server3"], d, "ntp_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server3"], "SystemDhcpServer-NtpServer3"); ok {
			if err = d.Set("ntp_server3", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server3: %v", err)
		}
	}

	if err = d.Set("ntp_service", flattenSystemDhcpServerNtpService(o["ntp-service"], d, "ntp_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-service"], "SystemDhcpServer-NtpService"); ok {
			if err = d.Set("ntp_service", vv); err != nil {
				return fmt.Errorf("Error reading ntp_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_service: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("options", flattenSystemDhcpServerOptions(o["options"], d, "options")); err != nil {
			if vv, ok := fortiAPIPatch(o["options"], "SystemDhcpServer-Options"); ok {
				if err = d.Set("options", vv); err != nil {
					return fmt.Errorf("Error reading options: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("options"); ok {
			if err = d.Set("options", flattenSystemDhcpServerOptions(o["options"], d, "options")); err != nil {
				if vv, ok := fortiAPIPatch(o["options"], "SystemDhcpServer-Options"); ok {
					if err = d.Set("options", vv); err != nil {
						return fmt.Errorf("Error reading options: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading options: %v", err)
				}
			}
		}
	}

	if err = d.Set("relay_agent", flattenSystemDhcpServerRelayAgent(o["relay-agent"], d, "relay_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["relay-agent"], "SystemDhcpServer-RelayAgent"); ok {
			if err = d.Set("relay_agent", vv); err != nil {
				return fmt.Errorf("Error reading relay_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading relay_agent: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("reserved_address", flattenSystemDhcpServerReservedAddress(o["reserved-address"], d, "reserved_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["reserved-address"], "SystemDhcpServer-ReservedAddress"); ok {
				if err = d.Set("reserved_address", vv); err != nil {
					return fmt.Errorf("Error reading reserved_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading reserved_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("reserved_address"); ok {
			if err = d.Set("reserved_address", flattenSystemDhcpServerReservedAddress(o["reserved-address"], d, "reserved_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["reserved-address"], "SystemDhcpServer-ReservedAddress"); ok {
					if err = d.Set("reserved_address", vv); err != nil {
						return fmt.Errorf("Error reading reserved_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading reserved_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenSystemDhcpServerServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "SystemDhcpServer-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("shared_subnet", flattenSystemDhcpServerSharedSubnet(o["shared-subnet"], d, "shared_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["shared-subnet"], "SystemDhcpServer-SharedSubnet"); ok {
			if err = d.Set("shared_subnet", vv); err != nil {
				return fmt.Errorf("Error reading shared_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shared_subnet: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemDhcpServerStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemDhcpServer-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tftp_server", flattenSystemDhcpServerTftpServer(o["tftp-server"], d, "tftp_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tftp-server"], "SystemDhcpServer-TftpServer"); ok {
			if err = d.Set("tftp_server", vv); err != nil {
				return fmt.Errorf("Error reading tftp_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tftp_server: %v", err)
		}
	}

	if err = d.Set("timezone", flattenSystemDhcpServerTimezone(o["timezone"], d, "timezone")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone"], "SystemDhcpServer-Timezone"); ok {
			if err = d.Set("timezone", vv); err != nil {
				return fmt.Errorf("Error reading timezone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("timezone_option", flattenSystemDhcpServerTimezoneOption(o["timezone-option"], d, "timezone_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone-option"], "SystemDhcpServer-TimezoneOption"); ok {
			if err = d.Set("timezone_option", vv); err != nil {
				return fmt.Errorf("Error reading timezone_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone_option: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenSystemDhcpServerVciMatch(o["vci-match"], d, "vci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-match"], "SystemDhcpServer-VciMatch"); ok {
			if err = d.Set("vci_match", vv); err != nil {
				return fmt.Errorf("Error reading vci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

	if err = d.Set("vci_string", flattenSystemDhcpServerVciString(o["vci-string"], d, "vci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-string"], "SystemDhcpServer-VciString"); ok {
			if err = d.Set("vci_string", vv); err != nil {
				return fmt.Errorf("Error reading vci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_string: %v", err)
		}
	}

	if err = d.Set("wifi_ac_service", flattenSystemDhcpServerWifiAcService(o["wifi-ac-service"], d, "wifi_ac_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac-service"], "SystemDhcpServer-WifiAcService"); ok {
			if err = d.Set("wifi_ac_service", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac_service: %v", err)
		}
	}

	if err = d.Set("wifi_ac1", flattenSystemDhcpServerWifiAc1(o["wifi-ac1"], d, "wifi_ac1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac1"], "SystemDhcpServer-WifiAc1"); ok {
			if err = d.Set("wifi_ac1", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac1: %v", err)
		}
	}

	if err = d.Set("wifi_ac2", flattenSystemDhcpServerWifiAc2(o["wifi-ac2"], d, "wifi_ac2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac2"], "SystemDhcpServer-WifiAc2"); ok {
			if err = d.Set("wifi_ac2", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac2: %v", err)
		}
	}

	if err = d.Set("wifi_ac3", flattenSystemDhcpServerWifiAc3(o["wifi-ac3"], d, "wifi_ac3")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac3"], "SystemDhcpServer-WifiAc3"); ok {
			if err = d.Set("wifi_ac3", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac3: %v", err)
		}
	}

	if err = d.Set("wins_server1", flattenSystemDhcpServerWinsServer1(o["wins-server1"], d, "wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server1"], "SystemDhcpServer-WinsServer1"); ok {
			if err = d.Set("wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", flattenSystemDhcpServerWinsServer2(o["wins-server2"], d, "wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server2"], "SystemDhcpServer-WinsServer2"); ok {
			if err = d.Set("wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	return nil
}

func flattenSystemDhcpServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDhcpServerAutoConfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerAutoManagedStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerConflictedIpTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerDdnsKeyname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsUpdateOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDefaultGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDhcpSettingsFromFortiipam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDnsServer4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDnsService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandSystemDhcpServerExcludeRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemDhcpServerExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lease-time"], _ = expandSystemDhcpServerExcludeRangeLeaseTime(d, i["lease_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandSystemDhcpServerExcludeRangeStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandSystemDhcpServerExcludeRangeUciMatch(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandSystemDhcpServerExcludeRangeUciString(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandSystemDhcpServerExcludeRangeVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandSystemDhcpServerExcludeRangeVciString(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemDhcpServerExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeUciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeUciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerExcludeRangeVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerFilename(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerForticlientOnNetStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerIpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerIpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandSystemDhcpServerIpRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemDhcpServerIpRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lease-time"], _ = expandSystemDhcpServerIpRangeLeaseTime(d, i["lease_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandSystemDhcpServerIpRangeStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandSystemDhcpServerIpRangeUciMatch(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandSystemDhcpServerIpRangeUciString(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandSystemDhcpServerIpRangeVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandSystemDhcpServerIpRangeVciString(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemDhcpServerIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerIpRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerIpRangeLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerIpRangeUciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerIpRangeUciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerIpRangeVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerIpRangeVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerIpsecLeaseHold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerMacAclDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerNextServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerNtpServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerNtpServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerNtpServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerNtpService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["code"], _ = expandSystemDhcpServerOptionsCode(d, i["code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemDhcpServerOptionsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemDhcpServerOptionsIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSystemDhcpServerOptionsType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandSystemDhcpServerOptionsUciMatch(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandSystemDhcpServerOptionsUciString(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandSystemDhcpServerOptionsValue(d, i["value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandSystemDhcpServerOptionsVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandSystemDhcpServerOptionsVciString(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemDhcpServerOptionsCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerOptionsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsUciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsUciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerOptionsValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerRelayAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandSystemDhcpServerReservedAddressAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id"], _ = expandSystemDhcpServerReservedAddressCircuitId(d, i["circuit_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id-type"], _ = expandSystemDhcpServerReservedAddressCircuitIdType(d, i["circuit_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandSystemDhcpServerReservedAddressDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemDhcpServerReservedAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemDhcpServerReservedAddressIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandSystemDhcpServerReservedAddressMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id"], _ = expandSystemDhcpServerReservedAddressRemoteId(d, i["remote_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id-type"], _ = expandSystemDhcpServerReservedAddressRemoteIdType(d, i["remote_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSystemDhcpServerReservedAddressType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemDhcpServerReservedAddressAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressCircuitId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressCircuitIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressRemoteId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressRemoteIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerSharedSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerTftpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerTimezone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerTimezoneOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerWifiAcService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerWifiAc1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerWifiAc2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerWifiAc3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerWinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerWinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDhcpServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_configuration"); ok || d.HasChange("auto_configuration") {
		t, err := expandSystemDhcpServerAutoConfiguration(d, v, "auto_configuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-configuration"] = t
		}
	}

	if v, ok := d.GetOk("auto_managed_status"); ok || d.HasChange("auto_managed_status") {
		t, err := expandSystemDhcpServerAutoManagedStatus(d, v, "auto_managed_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-managed-status"] = t
		}
	}

	if v, ok := d.GetOk("conflicted_ip_timeout"); ok || d.HasChange("conflicted_ip_timeout") {
		t, err := expandSystemDhcpServerConflictedIpTimeout(d, v, "conflicted_ip_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conflicted-ip-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ddns_auth"); ok || d.HasChange("ddns_auth") {
		t, err := expandSystemDhcpServerDdnsAuth(d, v, "ddns_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-auth"] = t
		}
	}

	if v, ok := d.GetOk("ddns_key"); ok || d.HasChange("ddns_key") {
		t, err := expandSystemDhcpServerDdnsKey(d, v, "ddns_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-key"] = t
		}
	}

	if v, ok := d.GetOk("ddns_keyname"); ok || d.HasChange("ddns_keyname") {
		t, err := expandSystemDhcpServerDdnsKeyname(d, v, "ddns_keyname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-keyname"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok || d.HasChange("ddns_server_ip") {
		t, err := expandSystemDhcpServerDdnsServerIp(d, v, "ddns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_ttl"); ok || d.HasChange("ddns_ttl") {
		t, err := expandSystemDhcpServerDdnsTtl(d, v, "ddns_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-ttl"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update"); ok || d.HasChange("ddns_update") {
		t, err := expandSystemDhcpServerDdnsUpdate(d, v, "ddns_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update_override"); ok || d.HasChange("ddns_update_override") {
		t, err := expandSystemDhcpServerDdnsUpdateOverride(d, v, "ddns_update_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update-override"] = t
		}
	}

	if v, ok := d.GetOk("ddns_zone"); ok || d.HasChange("ddns_zone") {
		t, err := expandSystemDhcpServerDdnsZone(d, v, "ddns_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-zone"] = t
		}
	}

	if v, ok := d.GetOk("default_gateway"); ok || d.HasChange("default_gateway") {
		t, err := expandSystemDhcpServerDefaultGateway(d, v, "default_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gateway"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_settings_from_fortiipam"); ok || d.HasChange("dhcp_settings_from_fortiipam") {
		t, err := expandSystemDhcpServerDhcpSettingsFromFortiipam(d, v, "dhcp_settings_from_fortiipam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-settings-from-fortiipam"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok || d.HasChange("dns_server1") {
		t, err := expandSystemDhcpServerDnsServer1(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok || d.HasChange("dns_server2") {
		t, err := expandSystemDhcpServerDnsServer2(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_server3"); ok || d.HasChange("dns_server3") {
		t, err := expandSystemDhcpServerDnsServer3(d, v, "dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("dns_server4"); ok || d.HasChange("dns_server4") {
		t, err := expandSystemDhcpServerDnsServer4(d, v, "dns_server4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server4"] = t
		}
	}

	if v, ok := d.GetOk("dns_service"); ok || d.HasChange("dns_service") {
		t, err := expandSystemDhcpServerDnsService(d, v, "dns_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-service"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandSystemDhcpServerDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("exclude_range"); ok || d.HasChange("exclude_range") {
		t, err := expandSystemDhcpServerExcludeRange(d, v, "exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("filename"); ok || d.HasChange("filename") {
		t, err := expandSystemDhcpServerFilename(d, v, "filename")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filename"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_on_net_status"); ok || d.HasChange("forticlient_on_net_status") {
		t, err := expandSystemDhcpServerForticlientOnNetStatus(d, v, "forticlient_on_net_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-on-net-status"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemDhcpServerId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemDhcpServerInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip_mode"); ok || d.HasChange("ip_mode") {
		t, err := expandSystemDhcpServerIpMode(d, v, "ip_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandSystemDhcpServerIpRange(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_lease_hold"); ok || d.HasChange("ipsec_lease_hold") {
		t, err := expandSystemDhcpServerIpsecLeaseHold(d, v, "ipsec_lease_hold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-lease-hold"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok || d.HasChange("lease_time") {
		t, err := expandSystemDhcpServerLeaseTime(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("mac_acl_default_action"); ok || d.HasChange("mac_acl_default_action") {
		t, err := expandSystemDhcpServerMacAclDefaultAction(d, v, "mac_acl_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-acl-default-action"] = t
		}
	}

	if v, ok := d.GetOk("netmask"); ok || d.HasChange("netmask") {
		t, err := expandSystemDhcpServerNetmask(d, v, "netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netmask"] = t
		}
	}

	if v, ok := d.GetOk("next_server"); ok || d.HasChange("next_server") {
		t, err := expandSystemDhcpServerNextServer(d, v, "next_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-server"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server1"); ok || d.HasChange("ntp_server1") {
		t, err := expandSystemDhcpServerNtpServer1(d, v, "ntp_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server1"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server2"); ok || d.HasChange("ntp_server2") {
		t, err := expandSystemDhcpServerNtpServer2(d, v, "ntp_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server2"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server3"); ok || d.HasChange("ntp_server3") {
		t, err := expandSystemDhcpServerNtpServer3(d, v, "ntp_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server3"] = t
		}
	}

	if v, ok := d.GetOk("ntp_service"); ok || d.HasChange("ntp_service") {
		t, err := expandSystemDhcpServerNtpService(d, v, "ntp_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-service"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandSystemDhcpServerOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("relay_agent"); ok || d.HasChange("relay_agent") {
		t, err := expandSystemDhcpServerRelayAgent(d, v, "relay_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relay-agent"] = t
		}
	}

	if v, ok := d.GetOk("reserved_address"); ok || d.HasChange("reserved_address") {
		t, err := expandSystemDhcpServerReservedAddress(d, v, "reserved_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reserved-address"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandSystemDhcpServerServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("shared_subnet"); ok || d.HasChange("shared_subnet") {
		t, err := expandSystemDhcpServerSharedSubnet(d, v, "shared_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shared-subnet"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemDhcpServerStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tftp_server"); ok || d.HasChange("tftp_server") {
		t, err := expandSystemDhcpServerTftpServer(d, v, "tftp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tftp-server"] = t
		}
	}

	if v, ok := d.GetOk("timezone"); ok || d.HasChange("timezone") {
		t, err := expandSystemDhcpServerTimezone(d, v, "timezone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone"] = t
		}
	}

	if v, ok := d.GetOk("timezone_option"); ok || d.HasChange("timezone_option") {
		t, err := expandSystemDhcpServerTimezoneOption(d, v, "timezone_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone-option"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok || d.HasChange("vci_match") {
		t, err := expandSystemDhcpServerVciMatch(d, v, "vci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok || d.HasChange("vci_string") {
		t, err := expandSystemDhcpServerVciString(d, v, "vci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac_service"); ok || d.HasChange("wifi_ac_service") {
		t, err := expandSystemDhcpServerWifiAcService(d, v, "wifi_ac_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac-service"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac1"); ok || d.HasChange("wifi_ac1") {
		t, err := expandSystemDhcpServerWifiAc1(d, v, "wifi_ac1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac1"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac2"); ok || d.HasChange("wifi_ac2") {
		t, err := expandSystemDhcpServerWifiAc2(d, v, "wifi_ac2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac2"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac3"); ok || d.HasChange("wifi_ac3") {
		t, err := expandSystemDhcpServerWifiAc3(d, v, "wifi_ac3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac3"] = t
		}
	}

	if v, ok := d.GetOk("wins_server1"); ok || d.HasChange("wins_server1") {
		t, err := expandSystemDhcpServerWinsServer1(d, v, "wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("wins_server2"); ok || d.HasChange("wins_server2") {
		t, err := expandSystemDhcpServerWinsServer2(d, v, "wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server2"] = t
		}
	}

	return &obj, nil
}
