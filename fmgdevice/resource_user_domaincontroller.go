// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure domain controller entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserDomainController() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserDomainControllerCreate,
		Read:   resourceUserDomainControllerRead,
		Update: resourceUserDomainControllerUpdate,
		Delete: resourceUserDomainControllerDelete,

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
			"ad_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adlds_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"adlds_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adlds_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adlds_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"change_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"change_detection_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dns_srv_lookup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extra_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"source_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
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
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"replication_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_name_src": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceUserDomainControllerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserDomainController(d)
	if err != nil {
		return fmt.Errorf("Error creating UserDomainController resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserDomainController(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserDomainController(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserDomainController resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserDomainController(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserDomainController resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserDomainControllerRead(d, m)
}

func resourceUserDomainControllerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserDomainController(d)
	if err != nil {
		return fmt.Errorf("Error updating UserDomainController resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserDomainController(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserDomainController resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserDomainControllerRead(d, m)
}

func resourceUserDomainControllerDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserDomainController(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserDomainController resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserDomainControllerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserDomainController(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserDomainController resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserDomainController(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserDomainController resource from API: %v", err)
	}
	return nil
}

func flattenUserDomainControllerAdMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerAdldsDn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerAdldsIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerAdldsIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerAdldsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerChangeDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerChangeDetectionPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerDnsSrvLookup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserDomainControllerExtraServerId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "UserDomainController-ExtraServer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := i["ip-address"]; ok {
			v := flattenUserDomainControllerExtraServerIpAddress(i["ip-address"], d, pre_append)
			tmp["ip_address"] = fortiAPISubPartPatch(v, "UserDomainController-ExtraServer-IpAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenUserDomainControllerExtraServerPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "UserDomainController-ExtraServer-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip_address"
		if _, ok := i["source-ip-address"]; ok {
			v := flattenUserDomainControllerExtraServerSourceIpAddress(i["source-ip-address"], d, pre_append)
			tmp["source_ip_address"] = fortiAPISubPartPatch(v, "UserDomainController-ExtraServer-SourceIpAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := i["source-port"]; ok {
			v := flattenUserDomainControllerExtraServerSourcePort(i["source-port"], d, pre_append)
			tmp["source_port"] = fortiAPISubPartPatch(v, "UserDomainController-ExtraServer-SourcePort")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserDomainControllerExtraServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServerIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServerSourceIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServerSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserDomainControllerInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserDomainControllerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerReplicationPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerSourceIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerDomainNameSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserDomainController(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("ad_mode", flattenUserDomainControllerAdMode(o["ad-mode"], d, "ad_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ad-mode"], "UserDomainController-AdMode"); ok {
			if err = d.Set("ad_mode", vv); err != nil {
				return fmt.Errorf("Error reading ad_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ad_mode: %v", err)
		}
	}

	if err = d.Set("adlds_dn", flattenUserDomainControllerAdldsDn(o["adlds-dn"], d, "adlds_dn")); err != nil {
		if vv, ok := fortiAPIPatch(o["adlds-dn"], "UserDomainController-AdldsDn"); ok {
			if err = d.Set("adlds_dn", vv); err != nil {
				return fmt.Errorf("Error reading adlds_dn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adlds_dn: %v", err)
		}
	}

	if err = d.Set("adlds_ip_address", flattenUserDomainControllerAdldsIpAddress(o["adlds-ip-address"], d, "adlds_ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["adlds-ip-address"], "UserDomainController-AdldsIpAddress"); ok {
			if err = d.Set("adlds_ip_address", vv); err != nil {
				return fmt.Errorf("Error reading adlds_ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adlds_ip_address: %v", err)
		}
	}

	if err = d.Set("adlds_ip6", flattenUserDomainControllerAdldsIp6(o["adlds-ip6"], d, "adlds_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["adlds-ip6"], "UserDomainController-AdldsIp6"); ok {
			if err = d.Set("adlds_ip6", vv); err != nil {
				return fmt.Errorf("Error reading adlds_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adlds_ip6: %v", err)
		}
	}

	if err = d.Set("adlds_port", flattenUserDomainControllerAdldsPort(o["adlds-port"], d, "adlds_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["adlds-port"], "UserDomainController-AdldsPort"); ok {
			if err = d.Set("adlds_port", vv); err != nil {
				return fmt.Errorf("Error reading adlds_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adlds_port: %v", err)
		}
	}

	if err = d.Set("change_detection", flattenUserDomainControllerChangeDetection(o["change-detection"], d, "change_detection")); err != nil {
		if vv, ok := fortiAPIPatch(o["change-detection"], "UserDomainController-ChangeDetection"); ok {
			if err = d.Set("change_detection", vv); err != nil {
				return fmt.Errorf("Error reading change_detection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading change_detection: %v", err)
		}
	}

	if err = d.Set("change_detection_period", flattenUserDomainControllerChangeDetectionPeriod(o["change-detection-period"], d, "change_detection_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["change-detection-period"], "UserDomainController-ChangeDetectionPeriod"); ok {
			if err = d.Set("change_detection_period", vv); err != nil {
				return fmt.Errorf("Error reading change_detection_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading change_detection_period: %v", err)
		}
	}

	if err = d.Set("dns_srv_lookup", flattenUserDomainControllerDnsSrvLookup(o["dns-srv-lookup"], d, "dns_srv_lookup")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-srv-lookup"], "UserDomainController-DnsSrvLookup"); ok {
			if err = d.Set("dns_srv_lookup", vv); err != nil {
				return fmt.Errorf("Error reading dns_srv_lookup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_srv_lookup: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenUserDomainControllerDomainName(o["domain-name"], d, "domain_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-name"], "UserDomainController-DomainName"); ok {
			if err = d.Set("domain_name", vv); err != nil {
				return fmt.Errorf("Error reading domain_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("extra_server", flattenUserDomainControllerExtraServer(o["extra-server"], d, "extra_server")); err != nil {
			if vv, ok := fortiAPIPatch(o["extra-server"], "UserDomainController-ExtraServer"); ok {
				if err = d.Set("extra_server", vv); err != nil {
					return fmt.Errorf("Error reading extra_server: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading extra_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("extra_server"); ok {
			if err = d.Set("extra_server", flattenUserDomainControllerExtraServer(o["extra-server"], d, "extra_server")); err != nil {
				if vv, ok := fortiAPIPatch(o["extra-server"], "UserDomainController-ExtraServer"); ok {
					if err = d.Set("extra_server", vv); err != nil {
						return fmt.Errorf("Error reading extra_server: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading extra_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("hostname", flattenUserDomainControllerHostname(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "UserDomainController-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("interface", flattenUserDomainControllerInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "UserDomainController-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenUserDomainControllerInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "UserDomainController-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ip_address", flattenUserDomainControllerIpAddress(o["ip-address"], d, "ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-address"], "UserDomainController-IpAddress"); ok {
			if err = d.Set("ip_address", vv); err != nil {
				return fmt.Errorf("Error reading ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	if err = d.Set("ip6", flattenUserDomainControllerIp6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "UserDomainController-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserDomainControllerLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "UserDomainController-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("name", flattenUserDomainControllerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserDomainController-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenUserDomainControllerPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "UserDomainController-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("replication_port", flattenUserDomainControllerReplicationPort(o["replication-port"], d, "replication_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["replication-port"], "UserDomainController-ReplicationPort"); ok {
			if err = d.Set("replication_port", vv); err != nil {
				return fmt.Errorf("Error reading replication_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replication_port: %v", err)
		}
	}

	if err = d.Set("source_ip_address", flattenUserDomainControllerSourceIpAddress(o["source-ip-address"], d, "source_ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-address"], "UserDomainController-SourceIpAddress"); ok {
			if err = d.Set("source_ip_address", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_address: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenUserDomainControllerSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip6"], "UserDomainController-SourceIp6"); ok {
			if err = d.Set("source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("source_port", flattenUserDomainControllerSourcePort(o["source-port"], d, "source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-port"], "UserDomainController-SourcePort"); ok {
			if err = d.Set("source_port", vv); err != nil {
				return fmt.Errorf("Error reading source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	if err = d.Set("username", flattenUserDomainControllerUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "UserDomainController-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("domain_name_src", flattenUserDomainControllerDomainNameSrc(o["domain-name-src"], d, "domain_name_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-name-src"], "UserDomainController-DomainNameSrc"); ok {
			if err = d.Set("domain_name_src", vv); err != nil {
				return fmt.Errorf("Error reading domain_name_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_name_src: %v", err)
		}
	}

	return nil
}

func flattenUserDomainControllerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserDomainControllerAdMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerAdldsDn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerAdldsIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerAdldsIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerAdldsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerChangeDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerChangeDetectionPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerDnsSrvLookup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerDomainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandUserDomainControllerExtraServerId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-address"], _ = expandUserDomainControllerExtraServerIpAddress(d, i["ip_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandUserDomainControllerExtraServerPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip-address"], _ = expandUserDomainControllerExtraServerSourceIpAddress(d, i["source_ip_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-port"], _ = expandUserDomainControllerExtraServerSourcePort(d, i["source_port"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserDomainControllerExtraServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerSourceIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserDomainControllerInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserDomainControllerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserDomainControllerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerReplicationPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerSourceIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerDomainNameSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserDomainController(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ad_mode"); ok || d.HasChange("ad_mode") {
		t, err := expandUserDomainControllerAdMode(d, v, "ad_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ad-mode"] = t
		}
	}

	if v, ok := d.GetOk("adlds_dn"); ok || d.HasChange("adlds_dn") {
		t, err := expandUserDomainControllerAdldsDn(d, v, "adlds_dn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adlds-dn"] = t
		}
	}

	if v, ok := d.GetOk("adlds_ip_address"); ok || d.HasChange("adlds_ip_address") {
		t, err := expandUserDomainControllerAdldsIpAddress(d, v, "adlds_ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adlds-ip-address"] = t
		}
	}

	if v, ok := d.GetOk("adlds_ip6"); ok || d.HasChange("adlds_ip6") {
		t, err := expandUserDomainControllerAdldsIp6(d, v, "adlds_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adlds-ip6"] = t
		}
	}

	if v, ok := d.GetOk("adlds_port"); ok || d.HasChange("adlds_port") {
		t, err := expandUserDomainControllerAdldsPort(d, v, "adlds_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adlds-port"] = t
		}
	}

	if v, ok := d.GetOk("change_detection"); ok || d.HasChange("change_detection") {
		t, err := expandUserDomainControllerChangeDetection(d, v, "change_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["change-detection"] = t
		}
	}

	if v, ok := d.GetOk("change_detection_period"); ok || d.HasChange("change_detection_period") {
		t, err := expandUserDomainControllerChangeDetectionPeriod(d, v, "change_detection_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["change-detection-period"] = t
		}
	}

	if v, ok := d.GetOk("dns_srv_lookup"); ok || d.HasChange("dns_srv_lookup") {
		t, err := expandUserDomainControllerDnsSrvLookup(d, v, "dns_srv_lookup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-srv-lookup"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok || d.HasChange("domain_name") {
		t, err := expandUserDomainControllerDomainName(d, v, "domain_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("extra_server"); ok || d.HasChange("extra_server") {
		t, err := expandUserDomainControllerExtraServer(d, v, "extra_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra-server"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandUserDomainControllerHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandUserDomainControllerInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandUserDomainControllerInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ip_address"); ok || d.HasChange("ip_address") {
		t, err := expandUserDomainControllerIpAddress(d, v, "ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok || d.HasChange("ip6") {
		t, err := expandUserDomainControllerIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandUserDomainControllerLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserDomainControllerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandUserDomainControllerPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandUserDomainControllerPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("replication_port"); ok || d.HasChange("replication_port") {
		t, err := expandUserDomainControllerReplicationPort(d, v, "replication_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replication-port"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_address"); ok || d.HasChange("source_ip_address") {
		t, err := expandUserDomainControllerSourceIpAddress(d, v, "source_ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-address"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok || d.HasChange("source_ip6") {
		t, err := expandUserDomainControllerSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("source_port"); ok || d.HasChange("source_port") {
		t, err := expandUserDomainControllerSourcePort(d, v, "source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandUserDomainControllerUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("domain_name_src"); ok || d.HasChange("domain_name_src") {
		t, err := expandUserDomainControllerDomainNameSrc(d, v, "domain_name_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name-src"] = t
		}
	}

	return &obj, nil
}
