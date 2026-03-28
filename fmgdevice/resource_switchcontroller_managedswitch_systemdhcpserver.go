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

func resourceSwitchControllerManagedSwitchSystemDhcpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchSystemDhcpServerCreate,
		Read:   resourceSwitchControllerManagedSwitchSystemDhcpServerRead,
		Update: resourceSwitchControllerManagedSwitchSystemDhcpServerUpdate,
		Delete: resourceSwitchControllerManagedSwitchSystemDhcpServerDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_gateway": &schema.Schema{
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
			"dns_service": &schema.Schema{
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
			"ip_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"lease_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"netmask": &schema.Schema{
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
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_id": &schema.Schema{
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

func resourceSwitchControllerManagedSwitchSystemDhcpServerCreate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectSwitchControllerManagedSwitchSystemDhcpServer(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchSystemDhcpServer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSwitchControllerManagedSwitchSystemDhcpServer(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSwitchControllerManagedSwitchSystemDhcpServer(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SwitchControllerManagedSwitchSystemDhcpServer resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSwitchControllerManagedSwitchSystemDhcpServer(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SwitchControllerManagedSwitchSystemDhcpServer resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerManagedSwitchSystemDhcpServerRead(d, m)
}

func resourceSwitchControllerManagedSwitchSystemDhcpServerUpdate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectSwitchControllerManagedSwitchSystemDhcpServer(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSystemDhcpServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSwitchControllerManagedSwitchSystemDhcpServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSystemDhcpServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerManagedSwitchSystemDhcpServerRead(d, m)
}

func resourceSwitchControllerManagedSwitchSystemDhcpServerDelete(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	wsParams["adom"] = adomv

	err = c.DeleteSwitchControllerManagedSwitchSystemDhcpServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchSystemDhcpServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchSystemDhcpServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerManagedSwitchSystemDhcpServer(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSystemDhcpServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchSystemDhcpServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSystemDhcpServer resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerDefaultGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerDnsServer12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerDnsServer22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerDnsServer32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerDnsService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerIpRange2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeEndIp2edl(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchSystemDhcpServer-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchSystemDhcpServer-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeStartIp2edl(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchSystemDhcpServer-IpRange-StartIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerLeaseTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerNetmask2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerNtpServer12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerNtpServer22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerNtpServer32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerNtpService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerOptions2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsCode2edl(i["code"], d, pre_append)
			tmp["code"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchSystemDhcpServer-Options-Code")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchSystemDhcpServer-Options-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchSystemDhcpServer-Options-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchSystemDhcpServer-Options-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchSystemDhcpServer-Options-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsCode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerSwitchId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerManagedSwitchSystemDhcpServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("default_gateway", flattenSwitchControllerManagedSwitchSystemDhcpServerDefaultGateway2edl(o["default-gateway"], d, "default_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-gateway"], "SwitchControllerManagedSwitchSystemDhcpServer-DefaultGateway"); ok {
			if err = d.Set("default_gateway", vv); err != nil {
				return fmt.Errorf("Error reading default_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_gateway: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenSwitchControllerManagedSwitchSystemDhcpServerDnsServer12edl(o["dns-server1"], d, "dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server1"], "SwitchControllerManagedSwitchSystemDhcpServer-DnsServer1"); ok {
			if err = d.Set("dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenSwitchControllerManagedSwitchSystemDhcpServerDnsServer22edl(o["dns-server2"], d, "dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server2"], "SwitchControllerManagedSwitchSystemDhcpServer-DnsServer2"); ok {
			if err = d.Set("dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_server3", flattenSwitchControllerManagedSwitchSystemDhcpServerDnsServer32edl(o["dns-server3"], d, "dns_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server3"], "SwitchControllerManagedSwitchSystemDhcpServer-DnsServer3"); ok {
			if err = d.Set("dns_server3", vv); err != nil {
				return fmt.Errorf("Error reading dns_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server3: %v", err)
		}
	}

	if err = d.Set("dns_service", flattenSwitchControllerManagedSwitchSystemDhcpServerDnsService2edl(o["dns-service"], d, "dns_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-service"], "SwitchControllerManagedSwitchSystemDhcpServer-DnsService"); ok {
			if err = d.Set("dns_service", vv); err != nil {
				return fmt.Errorf("Error reading dns_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_service: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSwitchControllerManagedSwitchSystemDhcpServerId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SwitchControllerManagedSwitchSystemDhcpServer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSwitchControllerManagedSwitchSystemDhcpServerInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SwitchControllerManagedSwitchSystemDhcpServer-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_range", flattenSwitchControllerManagedSwitchSystemDhcpServerIpRange2edl(o["ip-range"], d, "ip_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-range"], "SwitchControllerManagedSwitchSystemDhcpServer-IpRange"); ok {
				if err = d.Set("ip_range", vv); err != nil {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenSwitchControllerManagedSwitchSystemDhcpServerIpRange2edl(o["ip-range"], d, "ip_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-range"], "SwitchControllerManagedSwitchSystemDhcpServer-IpRange"); ok {
					if err = d.Set("ip_range", vv); err != nil {
						return fmt.Errorf("Error reading ip_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("lease_time", flattenSwitchControllerManagedSwitchSystemDhcpServerLeaseTime2edl(o["lease-time"], d, "lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["lease-time"], "SwitchControllerManagedSwitchSystemDhcpServer-LeaseTime"); ok {
			if err = d.Set("lease_time", vv); err != nil {
				return fmt.Errorf("Error reading lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("netmask", flattenSwitchControllerManagedSwitchSystemDhcpServerNetmask2edl(o["netmask"], d, "netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["netmask"], "SwitchControllerManagedSwitchSystemDhcpServer-Netmask"); ok {
			if err = d.Set("netmask", vv); err != nil {
				return fmt.Errorf("Error reading netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netmask: %v", err)
		}
	}

	if err = d.Set("ntp_server1", flattenSwitchControllerManagedSwitchSystemDhcpServerNtpServer12edl(o["ntp-server1"], d, "ntp_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server1"], "SwitchControllerManagedSwitchSystemDhcpServer-NtpServer1"); ok {
			if err = d.Set("ntp_server1", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server1: %v", err)
		}
	}

	if err = d.Set("ntp_server2", flattenSwitchControllerManagedSwitchSystemDhcpServerNtpServer22edl(o["ntp-server2"], d, "ntp_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server2"], "SwitchControllerManagedSwitchSystemDhcpServer-NtpServer2"); ok {
			if err = d.Set("ntp_server2", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server2: %v", err)
		}
	}

	if err = d.Set("ntp_server3", flattenSwitchControllerManagedSwitchSystemDhcpServerNtpServer32edl(o["ntp-server3"], d, "ntp_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server3"], "SwitchControllerManagedSwitchSystemDhcpServer-NtpServer3"); ok {
			if err = d.Set("ntp_server3", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server3: %v", err)
		}
	}

	if err = d.Set("ntp_service", flattenSwitchControllerManagedSwitchSystemDhcpServerNtpService2edl(o["ntp-service"], d, "ntp_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-service"], "SwitchControllerManagedSwitchSystemDhcpServer-NtpService"); ok {
			if err = d.Set("ntp_service", vv); err != nil {
				return fmt.Errorf("Error reading ntp_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_service: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("options", flattenSwitchControllerManagedSwitchSystemDhcpServerOptions2edl(o["options"], d, "options")); err != nil {
			if vv, ok := fortiAPIPatch(o["options"], "SwitchControllerManagedSwitchSystemDhcpServer-Options"); ok {
				if err = d.Set("options", vv); err != nil {
					return fmt.Errorf("Error reading options: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("options"); ok {
			if err = d.Set("options", flattenSwitchControllerManagedSwitchSystemDhcpServerOptions2edl(o["options"], d, "options")); err != nil {
				if vv, ok := fortiAPIPatch(o["options"], "SwitchControllerManagedSwitchSystemDhcpServer-Options"); ok {
					if err = d.Set("options", vv); err != nil {
						return fmt.Errorf("Error reading options: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading options: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenSwitchControllerManagedSwitchSystemDhcpServerStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SwitchControllerManagedSwitchSystemDhcpServer-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("switch_id", flattenSwitchControllerManagedSwitchSystemDhcpServerSwitchId2edl(o["switch-id"], d, "switch_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-id"], "SwitchControllerManagedSwitchSystemDhcpServer-SwitchId"); ok {
			if err = d.Set("switch_id", vv); err != nil {
				return fmt.Errorf("Error reading switch_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_id: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchSystemDhcpServerDefaultGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerDnsServer12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerDnsServer22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerDnsServer32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerDnsService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerIpRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["end-ip"], _ = expandSwitchControllerManagedSwitchSystemDhcpServerIpRangeEndIp2edl(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSwitchControllerManagedSwitchSystemDhcpServerIpRangeId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandSwitchControllerManagedSwitchSystemDhcpServerIpRangeStartIp2edl(d, i["start_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerIpRangeEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerIpRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerIpRangeStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerLeaseTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerNetmask2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerNtpServer12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerNtpServer22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerNtpServer32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerNtpService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["code"], _ = expandSwitchControllerManagedSwitchSystemDhcpServerOptionsCode2edl(d, i["code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSwitchControllerManagedSwitchSystemDhcpServerOptionsId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSwitchControllerManagedSwitchSystemDhcpServerOptionsIp2edl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSwitchControllerManagedSwitchSystemDhcpServerOptionsType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandSwitchControllerManagedSwitchSystemDhcpServerOptionsValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerOptionsCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerOptionsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerOptionsIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerOptionsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerOptionsValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerSwitchId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerManagedSwitchSystemDhcpServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_gateway"); ok || d.HasChange("default_gateway") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerDefaultGateway2edl(d, v, "default_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gateway"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok || d.HasChange("dns_server1") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerDnsServer12edl(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok || d.HasChange("dns_server2") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerDnsServer22edl(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_server3"); ok || d.HasChange("dns_server3") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerDnsServer32edl(d, v, "dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("dns_service"); ok || d.HasChange("dns_service") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerDnsService2edl(d, v, "dns_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-service"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerIpRange2edl(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok || d.HasChange("lease_time") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerLeaseTime2edl(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("netmask"); ok || d.HasChange("netmask") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerNetmask2edl(d, v, "netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netmask"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server1"); ok || d.HasChange("ntp_server1") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerNtpServer12edl(d, v, "ntp_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server1"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server2"); ok || d.HasChange("ntp_server2") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerNtpServer22edl(d, v, "ntp_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server2"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server3"); ok || d.HasChange("ntp_server3") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerNtpServer32edl(d, v, "ntp_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server3"] = t
		}
	}

	if v, ok := d.GetOk("ntp_service"); ok || d.HasChange("ntp_service") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerNtpService2edl(d, v, "ntp_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-service"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("switch_id"); ok || d.HasChange("switch_id") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerSwitchId2edl(d, v, "switch_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-id"] = t
		}
	}

	return &obj, nil
}
