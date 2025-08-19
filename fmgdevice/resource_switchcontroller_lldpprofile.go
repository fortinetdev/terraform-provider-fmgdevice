// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch LLDP profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerLldpProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerLldpProfileCreate,
		Read:   resourceSwitchControllerLldpProfileRead,
		Update: resourceSwitchControllerLldpProfileUpdate,
		Delete: resourceSwitchControllerLldpProfileDelete,

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
			"n8021_tlvs": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"n8023_tlvs": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auto_isl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl_auth_encrypt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl_auth_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_isl_auth_macsec_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_isl_auth_reauth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_isl_auth_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_isl_hello_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_isl_port_group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auto_isl_receive_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_mclag_icl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_tlvs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"information_string": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"oui": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subtype": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"med_location_service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sys_location_id": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"med_network_policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assign_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vlan_intf": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"med_tlvs": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchControllerLldpProfileCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerLldpProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerLldpProfile resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerLldpProfile(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerLldpProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerLldpProfileRead(d, m)
}

func resourceSwitchControllerLldpProfileUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerLldpProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerLldpProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerLldpProfileRead(d, m)
}

func resourceSwitchControllerLldpProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerLldpProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerLldpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLldpProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerLldpProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerLldpProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpProfile resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerLldpProfile8021Tlvs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerLldpProfile8023Tlvs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerLldpProfileAutoIsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslAuthEncrypt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslAuthIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslAuthMacsecProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslAuthReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslAuthUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslHelloTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslPortGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslReceiveTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoMclagIcl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "information_string"
		if _, ok := i["information-string"]; ok {
			v := flattenSwitchControllerLldpProfileCustomTlvsInformationString(i["information-string"], d, pre_append)
			tmp["information_string"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-CustomTlvs-InformationString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSwitchControllerLldpProfileCustomTlvsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-CustomTlvs-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui"
		if _, ok := i["oui"]; ok {
			v := flattenSwitchControllerLldpProfileCustomTlvsOui(i["oui"], d, pre_append)
			tmp["oui"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-CustomTlvs-Oui")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtype"
		if _, ok := i["subtype"]; ok {
			v := flattenSwitchControllerLldpProfileCustomTlvsSubtype(i["subtype"], d, pre_append)
			tmp["subtype"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-CustomTlvs-Subtype")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerLldpProfileCustomTlvsInformationString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvsOui(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvsSubtype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedLocationService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSwitchControllerLldpProfileMedLocationServiceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-MedLocationService-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSwitchControllerLldpProfileMedLocationServiceStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-MedLocationService-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sys_location_id"
		if _, ok := i["sys-location-id"]; ok {
			v := flattenSwitchControllerLldpProfileMedLocationServiceSysLocationId(i["sys-location-id"], d, pre_append)
			tmp["sys_location_id"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-MedLocationService-SysLocationId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerLldpProfileMedLocationServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedLocationServiceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedLocationServiceSysLocationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerLldpProfileMedNetworkPolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assign_vlan"
		if _, ok := i["assign-vlan"]; ok {
			v := flattenSwitchControllerLldpProfileMedNetworkPolicyAssignVlan(i["assign-vlan"], d, pre_append)
			tmp["assign_vlan"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-MedNetworkPolicy-AssignVlan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := i["dscp"]; ok {
			v := flattenSwitchControllerLldpProfileMedNetworkPolicyDscp(i["dscp"], d, pre_append)
			tmp["dscp"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-MedNetworkPolicy-Dscp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSwitchControllerLldpProfileMedNetworkPolicyName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-MedNetworkPolicy-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenSwitchControllerLldpProfileMedNetworkPolicyPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-MedNetworkPolicy-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSwitchControllerLldpProfileMedNetworkPolicyStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-MedNetworkPolicy-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := i["vlan"]; ok {
			v := flattenSwitchControllerLldpProfileMedNetworkPolicyVlan(i["vlan"], d, pre_append)
			tmp["vlan"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-MedNetworkPolicy-Vlan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_intf"
		if _, ok := i["vlan-intf"]; ok {
			v := flattenSwitchControllerLldpProfileMedNetworkPolicyVlanIntf(i["vlan-intf"], d, pre_append)
			tmp["vlan_intf"] = fortiAPISubPartPatch(v, "SwitchControllerLldpProfile-MedNetworkPolicy-VlanIntf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyAssignVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyVlanIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerLldpProfileMedTlvs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerLldpProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerLldpProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("n8021_tlvs", flattenSwitchControllerLldpProfile8021Tlvs(o["802.1-tlvs"], d, "n8021_tlvs")); err != nil {
		if vv, ok := fortiAPIPatch(o["802.1-tlvs"], "SwitchControllerLldpProfile-8021Tlvs"); ok {
			if err = d.Set("n8021_tlvs", vv); err != nil {
				return fmt.Errorf("Error reading n8021_tlvs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n8021_tlvs: %v", err)
		}
	}

	if err = d.Set("n8023_tlvs", flattenSwitchControllerLldpProfile8023Tlvs(o["802.3-tlvs"], d, "n8023_tlvs")); err != nil {
		if vv, ok := fortiAPIPatch(o["802.3-tlvs"], "SwitchControllerLldpProfile-8023Tlvs"); ok {
			if err = d.Set("n8023_tlvs", vv); err != nil {
				return fmt.Errorf("Error reading n8023_tlvs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n8023_tlvs: %v", err)
		}
	}

	if err = d.Set("auto_isl", flattenSwitchControllerLldpProfileAutoIsl(o["auto-isl"], d, "auto_isl")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl"], "SwitchControllerLldpProfile-AutoIsl"); ok {
			if err = d.Set("auto_isl", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth", flattenSwitchControllerLldpProfileAutoIslAuth(o["auto-isl-auth"], d, "auto_isl_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-auth"], "SwitchControllerLldpProfile-AutoIslAuth"); ok {
			if err = d.Set("auto_isl_auth", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_auth: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_encrypt", flattenSwitchControllerLldpProfileAutoIslAuthEncrypt(o["auto-isl-auth-encrypt"], d, "auto_isl_auth_encrypt")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-auth-encrypt"], "SwitchControllerLldpProfile-AutoIslAuthEncrypt"); ok {
			if err = d.Set("auto_isl_auth_encrypt", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_auth_encrypt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_auth_encrypt: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_identity", flattenSwitchControllerLldpProfileAutoIslAuthIdentity(o["auto-isl-auth-identity"], d, "auto_isl_auth_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-auth-identity"], "SwitchControllerLldpProfile-AutoIslAuthIdentity"); ok {
			if err = d.Set("auto_isl_auth_identity", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_auth_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_auth_identity: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_macsec_profile", flattenSwitchControllerLldpProfileAutoIslAuthMacsecProfile(o["auto-isl-auth-macsec-profile"], d, "auto_isl_auth_macsec_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-auth-macsec-profile"], "SwitchControllerLldpProfile-AutoIslAuthMacsecProfile"); ok {
			if err = d.Set("auto_isl_auth_macsec_profile", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_auth_macsec_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_auth_macsec_profile: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_reauth", flattenSwitchControllerLldpProfileAutoIslAuthReauth(o["auto-isl-auth-reauth"], d, "auto_isl_auth_reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-auth-reauth"], "SwitchControllerLldpProfile-AutoIslAuthReauth"); ok {
			if err = d.Set("auto_isl_auth_reauth", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_auth_reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_auth_reauth: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_user", flattenSwitchControllerLldpProfileAutoIslAuthUser(o["auto-isl-auth-user"], d, "auto_isl_auth_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-auth-user"], "SwitchControllerLldpProfile-AutoIslAuthUser"); ok {
			if err = d.Set("auto_isl_auth_user", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_auth_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_auth_user: %v", err)
		}
	}

	if err = d.Set("auto_isl_hello_timer", flattenSwitchControllerLldpProfileAutoIslHelloTimer(o["auto-isl-hello-timer"], d, "auto_isl_hello_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-hello-timer"], "SwitchControllerLldpProfile-AutoIslHelloTimer"); ok {
			if err = d.Set("auto_isl_hello_timer", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_hello_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_hello_timer: %v", err)
		}
	}

	if err = d.Set("auto_isl_port_group", flattenSwitchControllerLldpProfileAutoIslPortGroup(o["auto-isl-port-group"], d, "auto_isl_port_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-port-group"], "SwitchControllerLldpProfile-AutoIslPortGroup"); ok {
			if err = d.Set("auto_isl_port_group", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_port_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_port_group: %v", err)
		}
	}

	if err = d.Set("auto_isl_receive_timeout", flattenSwitchControllerLldpProfileAutoIslReceiveTimeout(o["auto-isl-receive-timeout"], d, "auto_isl_receive_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-receive-timeout"], "SwitchControllerLldpProfile-AutoIslReceiveTimeout"); ok {
			if err = d.Set("auto_isl_receive_timeout", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_receive_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_receive_timeout: %v", err)
		}
	}

	if err = d.Set("auto_mclag_icl", flattenSwitchControllerLldpProfileAutoMclagIcl(o["auto-mclag-icl"], d, "auto_mclag_icl")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-mclag-icl"], "SwitchControllerLldpProfile-AutoMclagIcl"); ok {
			if err = d.Set("auto_mclag_icl", vv); err != nil {
				return fmt.Errorf("Error reading auto_mclag_icl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_mclag_icl: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_tlvs", flattenSwitchControllerLldpProfileCustomTlvs(o["custom-tlvs"], d, "custom_tlvs")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-tlvs"], "SwitchControllerLldpProfile-CustomTlvs"); ok {
				if err = d.Set("custom_tlvs", vv); err != nil {
					return fmt.Errorf("Error reading custom_tlvs: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_tlvs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_tlvs"); ok {
			if err = d.Set("custom_tlvs", flattenSwitchControllerLldpProfileCustomTlvs(o["custom-tlvs"], d, "custom_tlvs")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-tlvs"], "SwitchControllerLldpProfile-CustomTlvs"); ok {
					if err = d.Set("custom_tlvs", vv); err != nil {
						return fmt.Errorf("Error reading custom_tlvs: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_tlvs: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("med_location_service", flattenSwitchControllerLldpProfileMedLocationService(o["med-location-service"], d, "med_location_service")); err != nil {
			if vv, ok := fortiAPIPatch(o["med-location-service"], "SwitchControllerLldpProfile-MedLocationService"); ok {
				if err = d.Set("med_location_service", vv); err != nil {
					return fmt.Errorf("Error reading med_location_service: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading med_location_service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("med_location_service"); ok {
			if err = d.Set("med_location_service", flattenSwitchControllerLldpProfileMedLocationService(o["med-location-service"], d, "med_location_service")); err != nil {
				if vv, ok := fortiAPIPatch(o["med-location-service"], "SwitchControllerLldpProfile-MedLocationService"); ok {
					if err = d.Set("med_location_service", vv); err != nil {
						return fmt.Errorf("Error reading med_location_service: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading med_location_service: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("med_network_policy", flattenSwitchControllerLldpProfileMedNetworkPolicy(o["med-network-policy"], d, "med_network_policy")); err != nil {
			if vv, ok := fortiAPIPatch(o["med-network-policy"], "SwitchControllerLldpProfile-MedNetworkPolicy"); ok {
				if err = d.Set("med_network_policy", vv); err != nil {
					return fmt.Errorf("Error reading med_network_policy: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading med_network_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("med_network_policy"); ok {
			if err = d.Set("med_network_policy", flattenSwitchControllerLldpProfileMedNetworkPolicy(o["med-network-policy"], d, "med_network_policy")); err != nil {
				if vv, ok := fortiAPIPatch(o["med-network-policy"], "SwitchControllerLldpProfile-MedNetworkPolicy"); ok {
					if err = d.Set("med_network_policy", vv); err != nil {
						return fmt.Errorf("Error reading med_network_policy: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading med_network_policy: %v", err)
				}
			}
		}
	}

	if err = d.Set("med_tlvs", flattenSwitchControllerLldpProfileMedTlvs(o["med-tlvs"], d, "med_tlvs")); err != nil {
		if vv, ok := fortiAPIPatch(o["med-tlvs"], "SwitchControllerLldpProfile-MedTlvs"); ok {
			if err = d.Set("med_tlvs", vv); err != nil {
				return fmt.Errorf("Error reading med_tlvs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading med_tlvs: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerLldpProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerLldpProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerLldpProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerLldpProfile8021Tlvs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerLldpProfile8023Tlvs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerLldpProfileAutoIsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslAuthEncrypt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslAuthIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslAuthMacsecProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslAuthReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslAuthUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslHelloTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslPortGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslReceiveTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoMclagIcl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "information_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["information-string"], _ = expandSwitchControllerLldpProfileCustomTlvsInformationString(d, i["information_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSwitchControllerLldpProfileCustomTlvsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["oui"], _ = expandSwitchControllerLldpProfileCustomTlvsOui(d, i["oui"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtype"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subtype"], _ = expandSwitchControllerLldpProfileCustomTlvsSubtype(d, i["subtype"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerLldpProfileCustomTlvsInformationString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvsOui(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvsSubtype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedLocationService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSwitchControllerLldpProfileMedLocationServiceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSwitchControllerLldpProfileMedLocationServiceStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sys_location_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sys-location-id"], _ = expandSwitchControllerLldpProfileMedLocationServiceSysLocationId(d, i["sys_location_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerLldpProfileMedLocationServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedLocationServiceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedLocationServiceSysLocationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assign_vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["assign-vlan"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyAssignVlan(d, i["assign_vlan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyDscp(d, i["dscp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyVlan(d, i["vlan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_intf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-intf"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyVlanIntf(d, i["vlan_intf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyAssignVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyVlanIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerLldpProfileMedTlvs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerLldpProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerLldpProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n8021_tlvs"); ok || d.HasChange("n8021_tlvs") {
		t, err := expandSwitchControllerLldpProfile8021Tlvs(d, v, "n8021_tlvs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802.1-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("n8023_tlvs"); ok || d.HasChange("n8023_tlvs") {
		t, err := expandSwitchControllerLldpProfile8023Tlvs(d, v, "n8023_tlvs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802.3-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl"); ok || d.HasChange("auto_isl") {
		t, err := expandSwitchControllerLldpProfileAutoIsl(d, v, "auto_isl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth"); ok || d.HasChange("auto_isl_auth") {
		t, err := expandSwitchControllerLldpProfileAutoIslAuth(d, v, "auto_isl_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_encrypt"); ok || d.HasChange("auto_isl_auth_encrypt") {
		t, err := expandSwitchControllerLldpProfileAutoIslAuthEncrypt(d, v, "auto_isl_auth_encrypt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-encrypt"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_identity"); ok || d.HasChange("auto_isl_auth_identity") {
		t, err := expandSwitchControllerLldpProfileAutoIslAuthIdentity(d, v, "auto_isl_auth_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-identity"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_macsec_profile"); ok || d.HasChange("auto_isl_auth_macsec_profile") {
		t, err := expandSwitchControllerLldpProfileAutoIslAuthMacsecProfile(d, v, "auto_isl_auth_macsec_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-macsec-profile"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_reauth"); ok || d.HasChange("auto_isl_auth_reauth") {
		t, err := expandSwitchControllerLldpProfileAutoIslAuthReauth(d, v, "auto_isl_auth_reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-reauth"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_user"); ok || d.HasChange("auto_isl_auth_user") {
		t, err := expandSwitchControllerLldpProfileAutoIslAuthUser(d, v, "auto_isl_auth_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-user"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_hello_timer"); ok || d.HasChange("auto_isl_hello_timer") {
		t, err := expandSwitchControllerLldpProfileAutoIslHelloTimer(d, v, "auto_isl_hello_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-hello-timer"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_port_group"); ok || d.HasChange("auto_isl_port_group") {
		t, err := expandSwitchControllerLldpProfileAutoIslPortGroup(d, v, "auto_isl_port_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-port-group"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_receive_timeout"); ok || d.HasChange("auto_isl_receive_timeout") {
		t, err := expandSwitchControllerLldpProfileAutoIslReceiveTimeout(d, v, "auto_isl_receive_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-receive-timeout"] = t
		}
	}

	if v, ok := d.GetOk("auto_mclag_icl"); ok || d.HasChange("auto_mclag_icl") {
		t, err := expandSwitchControllerLldpProfileAutoMclagIcl(d, v, "auto_mclag_icl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-mclag-icl"] = t
		}
	}

	if v, ok := d.GetOk("custom_tlvs"); ok || d.HasChange("custom_tlvs") {
		t, err := expandSwitchControllerLldpProfileCustomTlvs(d, v, "custom_tlvs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("med_location_service"); ok || d.HasChange("med_location_service") {
		t, err := expandSwitchControllerLldpProfileMedLocationService(d, v, "med_location_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-location-service"] = t
		}
	}

	if v, ok := d.GetOk("med_network_policy"); ok || d.HasChange("med_network_policy") {
		t, err := expandSwitchControllerLldpProfileMedNetworkPolicy(d, v, "med_network_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-network-policy"] = t
		}
	}

	if v, ok := d.GetOk("med_tlvs"); ok || d.HasChange("med_tlvs") {
		t, err := expandSwitchControllerLldpProfileMedTlvs(d, v, "med_tlvs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerLldpProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
