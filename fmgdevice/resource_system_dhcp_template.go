// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DHCP server templates.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDhcpTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcpTemplateCreate,
		Read:   resourceSystemDhcpTemplateRead,
		Update: resourceSystemDhcpTemplateUpdate,
		Delete: resourceSystemDhcpTemplateDelete,

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
			"auto_configuration": &schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lease_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"oui_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oui_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"start_ip_index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						"vendor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"fabric_force_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"ip_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lease_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"oui_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oui_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"reserve": &schema.Schema{
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
						"vendor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ipsec_lease_hold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
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
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
			"relay_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reserve_extra_addresses": &schema.Schema{
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
							Computed: true,
						},
						"circuit_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"circuit_id_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip_index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remote_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"remote_id_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
			"uuid": &schema.Schema{
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

func resourceSystemDhcpTemplateCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemDhcpTemplate(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcpTemplate resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemDhcpTemplate(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemDhcpTemplate(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemDhcpTemplate resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemDhcpTemplate(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemDhcpTemplate resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemDhcpTemplateRead(d, m)
}

func resourceSystemDhcpTemplateUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemDhcpTemplate(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpTemplate resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemDhcpTemplate(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpTemplate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemDhcpTemplateRead(d, m)
}

func resourceSystemDhcpTemplateDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteSystemDhcpTemplate(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcpTemplate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcpTemplateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDhcpTemplate(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemDhcpTemplate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcpTemplate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcpTemplate resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcpTemplateAutoConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateConflictedIpTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsKeyname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsUpdateOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDnsServer4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDnsService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemDhcpTemplateExcludeRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_count"
		if _, ok := i["ip-count"]; ok {
			v := flattenSystemDhcpTemplateExcludeRangeIpCount(i["ip-count"], d, pre_append)
			tmp["ip_count"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ExcludeRange-IpCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := i["lease-time"]; ok {
			v := flattenSystemDhcpTemplateExcludeRangeLeaseTime(i["lease-time"], d, pre_append)
			tmp["lease_time"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ExcludeRange-LeaseTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_match"
		if _, ok := i["oui-match"]; ok {
			v := flattenSystemDhcpTemplateExcludeRangeOuiMatch(i["oui-match"], d, pre_append)
			tmp["oui_match"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ExcludeRange-OuiMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_string"
		if _, ok := i["oui-string"]; ok {
			v := flattenSystemDhcpTemplateExcludeRangeOuiString(i["oui-string"], d, pre_append)
			tmp["oui_string"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ExcludeRange-OuiString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip_index"
		if _, ok := i["start-ip-index"]; ok {
			v := flattenSystemDhcpTemplateExcludeRangeStartIpIndex(i["start-ip-index"], d, pre_append)
			tmp["start_ip_index"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ExcludeRange-StartIpIndex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenSystemDhcpTemplateExcludeRangeUciMatch(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ExcludeRange-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenSystemDhcpTemplateExcludeRangeUciString(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ExcludeRange-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenSystemDhcpTemplateExcludeRangeVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ExcludeRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenSystemDhcpTemplateExcludeRangeVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ExcludeRange-VciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := i["vendor"]; ok {
			v := flattenSystemDhcpTemplateExcludeRangeVendor(i["vendor"], d, pre_append)
			tmp["vendor"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ExcludeRange-Vendor")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemDhcpTemplateExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeIpCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeOuiMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeOuiString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateExcludeRangeStartIpIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeUciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeUciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateExcludeRangeVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateExcludeRangeVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateFabricForceSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateFabricObjectSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateFilename(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateForticlientOnNetStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemDhcpTemplateIpRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_count"
		if _, ok := i["ip-count"]; ok {
			v := flattenSystemDhcpTemplateIpRangeIpCount(i["ip-count"], d, pre_append)
			tmp["ip_count"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-IpRange-IpCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := i["lease-time"]; ok {
			v := flattenSystemDhcpTemplateIpRangeLeaseTime(i["lease-time"], d, pre_append)
			tmp["lease_time"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-IpRange-LeaseTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_match"
		if _, ok := i["oui-match"]; ok {
			v := flattenSystemDhcpTemplateIpRangeOuiMatch(i["oui-match"], d, pre_append)
			tmp["oui_match"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-IpRange-OuiMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_string"
		if _, ok := i["oui-string"]; ok {
			v := flattenSystemDhcpTemplateIpRangeOuiString(i["oui-string"], d, pre_append)
			tmp["oui_string"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-IpRange-OuiString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reserve"
		if _, ok := i["reserve"]; ok {
			v := flattenSystemDhcpTemplateIpRangeReserve(i["reserve"], d, pre_append)
			tmp["reserve"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-IpRange-Reserve")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenSystemDhcpTemplateIpRangeUciMatch(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-IpRange-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenSystemDhcpTemplateIpRangeUciString(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-IpRange-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenSystemDhcpTemplateIpRangeVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-IpRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenSystemDhcpTemplateIpRangeVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-IpRange-VciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := i["vendor"]; ok {
			v := flattenSystemDhcpTemplateIpRangeVendor(i["vendor"], d, pre_append)
			tmp["vendor"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-IpRange-Vendor")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemDhcpTemplateIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeIpCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeOuiMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeOuiString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateIpRangeReserve(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeUciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeUciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateIpRangeVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateIpRangeVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpsecLeaseHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateMacAclDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateNextServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateNtpServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateNtpServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateNtpServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateNtpService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemDhcpTemplateOptionsCode(i["code"], d, pre_append)
			tmp["code"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-Options-Code")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemDhcpTemplateOptionsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-Options-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemDhcpTemplateOptionsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-Options-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSystemDhcpTemplateOptionsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-Options-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenSystemDhcpTemplateOptionsUciMatch(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-Options-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenSystemDhcpTemplateOptionsUciString(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-Options-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenSystemDhcpTemplateOptionsValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-Options-Value")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenSystemDhcpTemplateOptionsVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-Options-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenSystemDhcpTemplateOptionsVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-Options-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemDhcpTemplateOptionsCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsUciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsUciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateOptionsValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateRelayAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReserveExtraAddresses(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemDhcpTemplateReservedAddressAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ReservedAddress-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {
			v := flattenSystemDhcpTemplateReservedAddressCircuitId(i["circuit-id"], d, pre_append)
			tmp["circuit_id"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ReservedAddress-CircuitId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := i["circuit-id-type"]; ok {
			v := flattenSystemDhcpTemplateReservedAddressCircuitIdType(i["circuit-id-type"], d, pre_append)
			tmp["circuit_id_type"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ReservedAddress-CircuitIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenSystemDhcpTemplateReservedAddressDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ReservedAddress-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemDhcpTemplateReservedAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ReservedAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_index"
		if _, ok := i["ip-index"]; ok {
			v := flattenSystemDhcpTemplateReservedAddressIpIndex(i["ip-index"], d, pre_append)
			tmp["ip_index"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ReservedAddress-IpIndex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenSystemDhcpTemplateReservedAddressMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ReservedAddress-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			v := flattenSystemDhcpTemplateReservedAddressRemoteId(i["remote-id"], d, pre_append)
			tmp["remote_id"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ReservedAddress-RemoteId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := i["remote-id-type"]; ok {
			v := flattenSystemDhcpTemplateReservedAddressRemoteIdType(i["remote-id-type"], d, pre_append)
			tmp["remote_id_type"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ReservedAddress-RemoteIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSystemDhcpTemplateReservedAddressType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SystemDhcpTemplate-ReservedAddress-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemDhcpTemplateReservedAddressAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressCircuitId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressCircuitIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressIpIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressRemoteId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressRemoteIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateSharedSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateTftpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateTimezone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateTimezoneOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateWifiAcService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateWifiAc1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateWifiAc2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateWifiAc3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateWinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateWinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDhcpTemplate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auto_configuration", flattenSystemDhcpTemplateAutoConfiguration(o["auto-configuration"], d, "auto_configuration")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-configuration"], "SystemDhcpTemplate-AutoConfiguration"); ok {
			if err = d.Set("auto_configuration", vv); err != nil {
				return fmt.Errorf("Error reading auto_configuration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_configuration: %v", err)
		}
	}

	if err = d.Set("conflicted_ip_timeout", flattenSystemDhcpTemplateConflictedIpTimeout(o["conflicted-ip-timeout"], d, "conflicted_ip_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["conflicted-ip-timeout"], "SystemDhcpTemplate-ConflictedIpTimeout"); ok {
			if err = d.Set("conflicted_ip_timeout", vv); err != nil {
				return fmt.Errorf("Error reading conflicted_ip_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conflicted_ip_timeout: %v", err)
		}
	}

	if err = d.Set("ddns_auth", flattenSystemDhcpTemplateDdnsAuth(o["ddns-auth"], d, "ddns_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-auth"], "SystemDhcpTemplate-DdnsAuth"); ok {
			if err = d.Set("ddns_auth", vv); err != nil {
				return fmt.Errorf("Error reading ddns_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_auth: %v", err)
		}
	}

	if err = d.Set("ddns_keyname", flattenSystemDhcpTemplateDdnsKeyname(o["ddns-keyname"], d, "ddns_keyname")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-keyname"], "SystemDhcpTemplate-DdnsKeyname"); ok {
			if err = d.Set("ddns_keyname", vv); err != nil {
				return fmt.Errorf("Error reading ddns_keyname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_keyname: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenSystemDhcpTemplateDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-ip"], "SystemDhcpTemplate-DdnsServerIp"); ok {
			if err = d.Set("ddns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_ttl", flattenSystemDhcpTemplateDdnsTtl(o["ddns-ttl"], d, "ddns_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-ttl"], "SystemDhcpTemplate-DdnsTtl"); ok {
			if err = d.Set("ddns_ttl", vv); err != nil {
				return fmt.Errorf("Error reading ddns_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_ttl: %v", err)
		}
	}

	if err = d.Set("ddns_update", flattenSystemDhcpTemplateDdnsUpdate(o["ddns-update"], d, "ddns_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-update"], "SystemDhcpTemplate-DdnsUpdate"); ok {
			if err = d.Set("ddns_update", vv); err != nil {
				return fmt.Errorf("Error reading ddns_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_update: %v", err)
		}
	}

	if err = d.Set("ddns_update_override", flattenSystemDhcpTemplateDdnsUpdateOverride(o["ddns-update-override"], d, "ddns_update_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-update-override"], "SystemDhcpTemplate-DdnsUpdateOverride"); ok {
			if err = d.Set("ddns_update_override", vv); err != nil {
				return fmt.Errorf("Error reading ddns_update_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_update_override: %v", err)
		}
	}

	if err = d.Set("ddns_zone", flattenSystemDhcpTemplateDdnsZone(o["ddns-zone"], d, "ddns_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-zone"], "SystemDhcpTemplate-DdnsZone"); ok {
			if err = d.Set("ddns_zone", vv); err != nil {
				return fmt.Errorf("Error reading ddns_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_zone: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenSystemDhcpTemplateDnsServer1(o["dns-server1"], d, "dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server1"], "SystemDhcpTemplate-DnsServer1"); ok {
			if err = d.Set("dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenSystemDhcpTemplateDnsServer2(o["dns-server2"], d, "dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server2"], "SystemDhcpTemplate-DnsServer2"); ok {
			if err = d.Set("dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_server3", flattenSystemDhcpTemplateDnsServer3(o["dns-server3"], d, "dns_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server3"], "SystemDhcpTemplate-DnsServer3"); ok {
			if err = d.Set("dns_server3", vv); err != nil {
				return fmt.Errorf("Error reading dns_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server3: %v", err)
		}
	}

	if err = d.Set("dns_server4", flattenSystemDhcpTemplateDnsServer4(o["dns-server4"], d, "dns_server4")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server4"], "SystemDhcpTemplate-DnsServer4"); ok {
			if err = d.Set("dns_server4", vv); err != nil {
				return fmt.Errorf("Error reading dns_server4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server4: %v", err)
		}
	}

	if err = d.Set("dns_service", flattenSystemDhcpTemplateDnsService(o["dns-service"], d, "dns_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-service"], "SystemDhcpTemplate-DnsService"); ok {
			if err = d.Set("dns_service", vv); err != nil {
				return fmt.Errorf("Error reading dns_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_service: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemDhcpTemplateDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "SystemDhcpTemplate-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("exclude_range", flattenSystemDhcpTemplateExcludeRange(o["exclude-range"], d, "exclude_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["exclude-range"], "SystemDhcpTemplate-ExcludeRange"); ok {
				if err = d.Set("exclude_range", vv); err != nil {
					return fmt.Errorf("Error reading exclude_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("exclude_range"); ok {
			if err = d.Set("exclude_range", flattenSystemDhcpTemplateExcludeRange(o["exclude-range"], d, "exclude_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["exclude-range"], "SystemDhcpTemplate-ExcludeRange"); ok {
					if err = d.Set("exclude_range", vv); err != nil {
						return fmt.Errorf("Error reading exclude_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("fabric_force_sync", flattenSystemDhcpTemplateFabricForceSync(o["fabric-force-sync"], d, "fabric_force_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-force-sync"], "SystemDhcpTemplate-FabricForceSync"); ok {
			if err = d.Set("fabric_force_sync", vv); err != nil {
				return fmt.Errorf("Error reading fabric_force_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_force_sync: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenSystemDhcpTemplateFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "SystemDhcpTemplate-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("fabric_object_source", flattenSystemDhcpTemplateFabricObjectSource(o["fabric-object-source"], d, "fabric_object_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object-source"], "SystemDhcpTemplate-FabricObjectSource"); ok {
			if err = d.Set("fabric_object_source", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object_source: %v", err)
		}
	}

	if err = d.Set("filename", flattenSystemDhcpTemplateFilename(o["filename"], d, "filename")); err != nil {
		if vv, ok := fortiAPIPatch(o["filename"], "SystemDhcpTemplate-Filename"); ok {
			if err = d.Set("filename", vv); err != nil {
				return fmt.Errorf("Error reading filename: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filename: %v", err)
		}
	}

	if err = d.Set("forticlient_on_net_status", flattenSystemDhcpTemplateForticlientOnNetStatus(o["forticlient-on-net-status"], d, "forticlient_on_net_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-on-net-status"], "SystemDhcpTemplate-ForticlientOnNetStatus"); ok {
			if err = d.Set("forticlient_on_net_status", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_on_net_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_on_net_status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_range", flattenSystemDhcpTemplateIpRange(o["ip-range"], d, "ip_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-range"], "SystemDhcpTemplate-IpRange"); ok {
				if err = d.Set("ip_range", vv); err != nil {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenSystemDhcpTemplateIpRange(o["ip-range"], d, "ip_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-range"], "SystemDhcpTemplate-IpRange"); ok {
					if err = d.Set("ip_range", vv); err != nil {
						return fmt.Errorf("Error reading ip_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipsec_lease_hold", flattenSystemDhcpTemplateIpsecLeaseHold(o["ipsec-lease-hold"], d, "ipsec_lease_hold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-lease-hold"], "SystemDhcpTemplate-IpsecLeaseHold"); ok {
			if err = d.Set("ipsec_lease_hold", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenSystemDhcpTemplateLeaseTime(o["lease-time"], d, "lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["lease-time"], "SystemDhcpTemplate-LeaseTime"); ok {
			if err = d.Set("lease_time", vv); err != nil {
				return fmt.Errorf("Error reading lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("mac_acl_default_action", flattenSystemDhcpTemplateMacAclDefaultAction(o["mac-acl-default-action"], d, "mac_acl_default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-acl-default-action"], "SystemDhcpTemplate-MacAclDefaultAction"); ok {
			if err = d.Set("mac_acl_default_action", vv); err != nil {
				return fmt.Errorf("Error reading mac_acl_default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_acl_default_action: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemDhcpTemplateName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemDhcpTemplate-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("next_server", flattenSystemDhcpTemplateNextServer(o["next-server"], d, "next_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-server"], "SystemDhcpTemplate-NextServer"); ok {
			if err = d.Set("next_server", vv); err != nil {
				return fmt.Errorf("Error reading next_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_server: %v", err)
		}
	}

	if err = d.Set("ntp_server1", flattenSystemDhcpTemplateNtpServer1(o["ntp-server1"], d, "ntp_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server1"], "SystemDhcpTemplate-NtpServer1"); ok {
			if err = d.Set("ntp_server1", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server1: %v", err)
		}
	}

	if err = d.Set("ntp_server2", flattenSystemDhcpTemplateNtpServer2(o["ntp-server2"], d, "ntp_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server2"], "SystemDhcpTemplate-NtpServer2"); ok {
			if err = d.Set("ntp_server2", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server2: %v", err)
		}
	}

	if err = d.Set("ntp_server3", flattenSystemDhcpTemplateNtpServer3(o["ntp-server3"], d, "ntp_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server3"], "SystemDhcpTemplate-NtpServer3"); ok {
			if err = d.Set("ntp_server3", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server3: %v", err)
		}
	}

	if err = d.Set("ntp_service", flattenSystemDhcpTemplateNtpService(o["ntp-service"], d, "ntp_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-service"], "SystemDhcpTemplate-NtpService"); ok {
			if err = d.Set("ntp_service", vv); err != nil {
				return fmt.Errorf("Error reading ntp_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_service: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("options", flattenSystemDhcpTemplateOptions(o["options"], d, "options")); err != nil {
			if vv, ok := fortiAPIPatch(o["options"], "SystemDhcpTemplate-Options"); ok {
				if err = d.Set("options", vv); err != nil {
					return fmt.Errorf("Error reading options: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("options"); ok {
			if err = d.Set("options", flattenSystemDhcpTemplateOptions(o["options"], d, "options")); err != nil {
				if vv, ok := fortiAPIPatch(o["options"], "SystemDhcpTemplate-Options"); ok {
					if err = d.Set("options", vv); err != nil {
						return fmt.Errorf("Error reading options: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading options: %v", err)
				}
			}
		}
	}

	if err = d.Set("relay_agent", flattenSystemDhcpTemplateRelayAgent(o["relay-agent"], d, "relay_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["relay-agent"], "SystemDhcpTemplate-RelayAgent"); ok {
			if err = d.Set("relay_agent", vv); err != nil {
				return fmt.Errorf("Error reading relay_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading relay_agent: %v", err)
		}
	}

	if err = d.Set("reserve_extra_addresses", flattenSystemDhcpTemplateReserveExtraAddresses(o["reserve-extra-addresses"], d, "reserve_extra_addresses")); err != nil {
		if vv, ok := fortiAPIPatch(o["reserve-extra-addresses"], "SystemDhcpTemplate-ReserveExtraAddresses"); ok {
			if err = d.Set("reserve_extra_addresses", vv); err != nil {
				return fmt.Errorf("Error reading reserve_extra_addresses: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reserve_extra_addresses: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("reserved_address", flattenSystemDhcpTemplateReservedAddress(o["reserved-address"], d, "reserved_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["reserved-address"], "SystemDhcpTemplate-ReservedAddress"); ok {
				if err = d.Set("reserved_address", vv); err != nil {
					return fmt.Errorf("Error reading reserved_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading reserved_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("reserved_address"); ok {
			if err = d.Set("reserved_address", flattenSystemDhcpTemplateReservedAddress(o["reserved-address"], d, "reserved_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["reserved-address"], "SystemDhcpTemplate-ReservedAddress"); ok {
					if err = d.Set("reserved_address", vv); err != nil {
						return fmt.Errorf("Error reading reserved_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading reserved_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenSystemDhcpTemplateServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "SystemDhcpTemplate-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("shared_subnet", flattenSystemDhcpTemplateSharedSubnet(o["shared-subnet"], d, "shared_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["shared-subnet"], "SystemDhcpTemplate-SharedSubnet"); ok {
			if err = d.Set("shared_subnet", vv); err != nil {
				return fmt.Errorf("Error reading shared_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shared_subnet: %v", err)
		}
	}

	if err = d.Set("tftp_server", flattenSystemDhcpTemplateTftpServer(o["tftp-server"], d, "tftp_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tftp-server"], "SystemDhcpTemplate-TftpServer"); ok {
			if err = d.Set("tftp_server", vv); err != nil {
				return fmt.Errorf("Error reading tftp_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tftp_server: %v", err)
		}
	}

	if err = d.Set("timezone", flattenSystemDhcpTemplateTimezone(o["timezone"], d, "timezone")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone"], "SystemDhcpTemplate-Timezone"); ok {
			if err = d.Set("timezone", vv); err != nil {
				return fmt.Errorf("Error reading timezone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("timezone_option", flattenSystemDhcpTemplateTimezoneOption(o["timezone-option"], d, "timezone_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone-option"], "SystemDhcpTemplate-TimezoneOption"); ok {
			if err = d.Set("timezone_option", vv); err != nil {
				return fmt.Errorf("Error reading timezone_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone_option: %v", err)
		}
	}

	if err = d.Set("uuid", flattenSystemDhcpTemplateUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "SystemDhcpTemplate-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenSystemDhcpTemplateVciMatch(o["vci-match"], d, "vci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-match"], "SystemDhcpTemplate-VciMatch"); ok {
			if err = d.Set("vci_match", vv); err != nil {
				return fmt.Errorf("Error reading vci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

	if err = d.Set("vci_string", flattenSystemDhcpTemplateVciString(o["vci-string"], d, "vci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-string"], "SystemDhcpTemplate-VciString"); ok {
			if err = d.Set("vci_string", vv); err != nil {
				return fmt.Errorf("Error reading vci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_string: %v", err)
		}
	}

	if err = d.Set("wifi_ac_service", flattenSystemDhcpTemplateWifiAcService(o["wifi-ac-service"], d, "wifi_ac_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac-service"], "SystemDhcpTemplate-WifiAcService"); ok {
			if err = d.Set("wifi_ac_service", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac_service: %v", err)
		}
	}

	if err = d.Set("wifi_ac1", flattenSystemDhcpTemplateWifiAc1(o["wifi-ac1"], d, "wifi_ac1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac1"], "SystemDhcpTemplate-WifiAc1"); ok {
			if err = d.Set("wifi_ac1", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac1: %v", err)
		}
	}

	if err = d.Set("wifi_ac2", flattenSystemDhcpTemplateWifiAc2(o["wifi-ac2"], d, "wifi_ac2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac2"], "SystemDhcpTemplate-WifiAc2"); ok {
			if err = d.Set("wifi_ac2", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac2: %v", err)
		}
	}

	if err = d.Set("wifi_ac3", flattenSystemDhcpTemplateWifiAc3(o["wifi-ac3"], d, "wifi_ac3")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac3"], "SystemDhcpTemplate-WifiAc3"); ok {
			if err = d.Set("wifi_ac3", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac3: %v", err)
		}
	}

	if err = d.Set("wins_server1", flattenSystemDhcpTemplateWinsServer1(o["wins-server1"], d, "wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server1"], "SystemDhcpTemplate-WinsServer1"); ok {
			if err = d.Set("wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", flattenSystemDhcpTemplateWinsServer2(o["wins-server2"], d, "wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server2"], "SystemDhcpTemplate-WinsServer2"); ok {
			if err = d.Set("wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	return nil
}

func flattenSystemDhcpTemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDhcpTemplateAutoConfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateConflictedIpTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateDdnsKeyname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsUpdateOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDnsServer4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDnsService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemDhcpTemplateExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-count"], _ = expandSystemDhcpTemplateExcludeRangeIpCount(d, i["ip_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lease-time"], _ = expandSystemDhcpTemplateExcludeRangeLeaseTime(d, i["lease_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["oui-match"], _ = expandSystemDhcpTemplateExcludeRangeOuiMatch(d, i["oui_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["oui-string"], _ = expandSystemDhcpTemplateExcludeRangeOuiString(d, i["oui_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip_index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip-index"], _ = expandSystemDhcpTemplateExcludeRangeStartIpIndex(d, i["start_ip_index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandSystemDhcpTemplateExcludeRangeUciMatch(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandSystemDhcpTemplateExcludeRangeUciString(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandSystemDhcpTemplateExcludeRangeVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandSystemDhcpTemplateExcludeRangeVciString(d, i["vci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vendor"], _ = expandSystemDhcpTemplateExcludeRangeVendor(d, i["vendor"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeIpCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeOuiMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeOuiString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateExcludeRangeStartIpIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeUciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeUciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateExcludeRangeVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateExcludeRangeVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateFabricForceSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateFabricObjectSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateFilename(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateForticlientOnNetStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemDhcpTemplateIpRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-count"], _ = expandSystemDhcpTemplateIpRangeIpCount(d, i["ip_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lease-time"], _ = expandSystemDhcpTemplateIpRangeLeaseTime(d, i["lease_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["oui-match"], _ = expandSystemDhcpTemplateIpRangeOuiMatch(d, i["oui_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["oui-string"], _ = expandSystemDhcpTemplateIpRangeOuiString(d, i["oui_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reserve"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["reserve"], _ = expandSystemDhcpTemplateIpRangeReserve(d, i["reserve"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandSystemDhcpTemplateIpRangeUciMatch(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandSystemDhcpTemplateIpRangeUciString(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandSystemDhcpTemplateIpRangeVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandSystemDhcpTemplateIpRangeVciString(d, i["vci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vendor"], _ = expandSystemDhcpTemplateIpRangeVendor(d, i["vendor"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateIpRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeIpCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeOuiMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeOuiString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateIpRangeReserve(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeUciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeUciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateIpRangeVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateIpRangeVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpsecLeaseHold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateMacAclDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateNextServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateNtpServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateNtpServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateNtpServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateNtpService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["code"], _ = expandSystemDhcpTemplateOptionsCode(d, i["code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemDhcpTemplateOptionsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemDhcpTemplateOptionsIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSystemDhcpTemplateOptionsType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandSystemDhcpTemplateOptionsUciMatch(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandSystemDhcpTemplateOptionsUciString(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandSystemDhcpTemplateOptionsValue(d, i["value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandSystemDhcpTemplateOptionsVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandSystemDhcpTemplateOptionsVciString(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateOptionsCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsUciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsUciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateOptionsValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateRelayAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReserveExtraAddresses(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandSystemDhcpTemplateReservedAddressAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id"], _ = expandSystemDhcpTemplateReservedAddressCircuitId(d, i["circuit_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id-type"], _ = expandSystemDhcpTemplateReservedAddressCircuitIdType(d, i["circuit_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandSystemDhcpTemplateReservedAddressDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemDhcpTemplateReservedAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-index"], _ = expandSystemDhcpTemplateReservedAddressIpIndex(d, i["ip_index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandSystemDhcpTemplateReservedAddressMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id"], _ = expandSystemDhcpTemplateReservedAddressRemoteId(d, i["remote_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id-type"], _ = expandSystemDhcpTemplateReservedAddressRemoteIdType(d, i["remote_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSystemDhcpTemplateReservedAddressType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateReservedAddressAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressCircuitId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressCircuitIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressIpIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressRemoteId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressRemoteIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateSharedSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateTftpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateTimezone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateTimezoneOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateWifiAcService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateWifiAc1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateWifiAc2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateWifiAc3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateWinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateWinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDhcpTemplate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_configuration"); ok || d.HasChange("auto_configuration") {
		t, err := expandSystemDhcpTemplateAutoConfiguration(d, v, "auto_configuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-configuration"] = t
		}
	}

	if v, ok := d.GetOk("conflicted_ip_timeout"); ok || d.HasChange("conflicted_ip_timeout") {
		t, err := expandSystemDhcpTemplateConflictedIpTimeout(d, v, "conflicted_ip_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conflicted-ip-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ddns_auth"); ok || d.HasChange("ddns_auth") {
		t, err := expandSystemDhcpTemplateDdnsAuth(d, v, "ddns_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-auth"] = t
		}
	}

	if v, ok := d.GetOk("ddns_key"); ok || d.HasChange("ddns_key") {
		t, err := expandSystemDhcpTemplateDdnsKey(d, v, "ddns_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-key"] = t
		}
	}

	if v, ok := d.GetOk("ddns_keyname"); ok || d.HasChange("ddns_keyname") {
		t, err := expandSystemDhcpTemplateDdnsKeyname(d, v, "ddns_keyname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-keyname"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok || d.HasChange("ddns_server_ip") {
		t, err := expandSystemDhcpTemplateDdnsServerIp(d, v, "ddns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_ttl"); ok || d.HasChange("ddns_ttl") {
		t, err := expandSystemDhcpTemplateDdnsTtl(d, v, "ddns_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-ttl"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update"); ok || d.HasChange("ddns_update") {
		t, err := expandSystemDhcpTemplateDdnsUpdate(d, v, "ddns_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update_override"); ok || d.HasChange("ddns_update_override") {
		t, err := expandSystemDhcpTemplateDdnsUpdateOverride(d, v, "ddns_update_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update-override"] = t
		}
	}

	if v, ok := d.GetOk("ddns_zone"); ok || d.HasChange("ddns_zone") {
		t, err := expandSystemDhcpTemplateDdnsZone(d, v, "ddns_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-zone"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok || d.HasChange("dns_server1") {
		t, err := expandSystemDhcpTemplateDnsServer1(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok || d.HasChange("dns_server2") {
		t, err := expandSystemDhcpTemplateDnsServer2(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_server3"); ok || d.HasChange("dns_server3") {
		t, err := expandSystemDhcpTemplateDnsServer3(d, v, "dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("dns_server4"); ok || d.HasChange("dns_server4") {
		t, err := expandSystemDhcpTemplateDnsServer4(d, v, "dns_server4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server4"] = t
		}
	}

	if v, ok := d.GetOk("dns_service"); ok || d.HasChange("dns_service") {
		t, err := expandSystemDhcpTemplateDnsService(d, v, "dns_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-service"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandSystemDhcpTemplateDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("exclude_range"); ok || d.HasChange("exclude_range") {
		t, err := expandSystemDhcpTemplateExcludeRange(d, v, "exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("fabric_force_sync"); ok || d.HasChange("fabric_force_sync") {
		t, err := expandSystemDhcpTemplateFabricForceSync(d, v, "fabric_force_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-force-sync"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandSystemDhcpTemplateFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object_source"); ok || d.HasChange("fabric_object_source") {
		t, err := expandSystemDhcpTemplateFabricObjectSource(d, v, "fabric_object_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object-source"] = t
		}
	}

	if v, ok := d.GetOk("filename"); ok || d.HasChange("filename") {
		t, err := expandSystemDhcpTemplateFilename(d, v, "filename")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filename"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_on_net_status"); ok || d.HasChange("forticlient_on_net_status") {
		t, err := expandSystemDhcpTemplateForticlientOnNetStatus(d, v, "forticlient_on_net_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-on-net-status"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandSystemDhcpTemplateIpRange(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_lease_hold"); ok || d.HasChange("ipsec_lease_hold") {
		t, err := expandSystemDhcpTemplateIpsecLeaseHold(d, v, "ipsec_lease_hold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-lease-hold"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok || d.HasChange("lease_time") {
		t, err := expandSystemDhcpTemplateLeaseTime(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("mac_acl_default_action"); ok || d.HasChange("mac_acl_default_action") {
		t, err := expandSystemDhcpTemplateMacAclDefaultAction(d, v, "mac_acl_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-acl-default-action"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemDhcpTemplateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("next_server"); ok || d.HasChange("next_server") {
		t, err := expandSystemDhcpTemplateNextServer(d, v, "next_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-server"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server1"); ok || d.HasChange("ntp_server1") {
		t, err := expandSystemDhcpTemplateNtpServer1(d, v, "ntp_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server1"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server2"); ok || d.HasChange("ntp_server2") {
		t, err := expandSystemDhcpTemplateNtpServer2(d, v, "ntp_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server2"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server3"); ok || d.HasChange("ntp_server3") {
		t, err := expandSystemDhcpTemplateNtpServer3(d, v, "ntp_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server3"] = t
		}
	}

	if v, ok := d.GetOk("ntp_service"); ok || d.HasChange("ntp_service") {
		t, err := expandSystemDhcpTemplateNtpService(d, v, "ntp_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-service"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandSystemDhcpTemplateOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("relay_agent"); ok || d.HasChange("relay_agent") {
		t, err := expandSystemDhcpTemplateRelayAgent(d, v, "relay_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relay-agent"] = t
		}
	}

	if v, ok := d.GetOk("reserve_extra_addresses"); ok || d.HasChange("reserve_extra_addresses") {
		t, err := expandSystemDhcpTemplateReserveExtraAddresses(d, v, "reserve_extra_addresses")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reserve-extra-addresses"] = t
		}
	}

	if v, ok := d.GetOk("reserved_address"); ok || d.HasChange("reserved_address") {
		t, err := expandSystemDhcpTemplateReservedAddress(d, v, "reserved_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reserved-address"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandSystemDhcpTemplateServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("shared_subnet"); ok || d.HasChange("shared_subnet") {
		t, err := expandSystemDhcpTemplateSharedSubnet(d, v, "shared_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shared-subnet"] = t
		}
	}

	if v, ok := d.GetOk("tftp_server"); ok || d.HasChange("tftp_server") {
		t, err := expandSystemDhcpTemplateTftpServer(d, v, "tftp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tftp-server"] = t
		}
	}

	if v, ok := d.GetOk("timezone"); ok || d.HasChange("timezone") {
		t, err := expandSystemDhcpTemplateTimezone(d, v, "timezone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone"] = t
		}
	}

	if v, ok := d.GetOk("timezone_option"); ok || d.HasChange("timezone_option") {
		t, err := expandSystemDhcpTemplateTimezoneOption(d, v, "timezone_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone-option"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandSystemDhcpTemplateUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok || d.HasChange("vci_match") {
		t, err := expandSystemDhcpTemplateVciMatch(d, v, "vci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok || d.HasChange("vci_string") {
		t, err := expandSystemDhcpTemplateVciString(d, v, "vci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac_service"); ok || d.HasChange("wifi_ac_service") {
		t, err := expandSystemDhcpTemplateWifiAcService(d, v, "wifi_ac_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac-service"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac1"); ok || d.HasChange("wifi_ac1") {
		t, err := expandSystemDhcpTemplateWifiAc1(d, v, "wifi_ac1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac1"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac2"); ok || d.HasChange("wifi_ac2") {
		t, err := expandSystemDhcpTemplateWifiAc2(d, v, "wifi_ac2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac2"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac3"); ok || d.HasChange("wifi_ac3") {
		t, err := expandSystemDhcpTemplateWifiAc3(d, v, "wifi_ac3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac3"] = t
		}
	}

	if v, ok := d.GetOk("wins_server1"); ok || d.HasChange("wins_server1") {
		t, err := expandSystemDhcpTemplateWinsServer1(d, v, "wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("wins_server2"); ok || d.HasChange("wins_server2") {
		t, err := expandSystemDhcpTemplateWinsServer2(d, v, "wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server2"] = t
		}
	}

	return &obj, nil
}
