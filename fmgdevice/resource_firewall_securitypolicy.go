// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure NGFW IPv4/IPv6 application policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSecurityPolicyCreate,
		Read:   resourceFirewallSecurityPolicyRead,
		Update: resourceFirewallSecurityPolicyUpdate,
		Delete: resourceFirewallSecurityPolicyDelete,

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
			"_policy_block": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"app_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"application_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"cifs_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"casb_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dlp_sensor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"diameter_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dlp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"enforce_default_app_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fsso_groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"global_label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"icap_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_src_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service6_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service6_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service6_src_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ips_voip_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"learning_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"profile_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"profile_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sctp_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"send_deny_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"telemetry_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"url_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"videofilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"virtual_patch_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"voip_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallSecurityPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallSecurityPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSecurityPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("policyid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallSecurityPolicy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallSecurityPolicy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallSecurityPolicy resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallSecurityPolicy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallSecurityPolicy resource: %v", err)
		}

		if v != nil && v["policyid"] != nil {
			if vidn, ok := v["policyid"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallSecurityPolicyRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallSecurityPolicy resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallSecurityPolicyRead(d, m)
}

func resourceFirewallSecurityPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallSecurityPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSecurityPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallSecurityPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSecurityPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallSecurityPolicyRead(d, m)
}

func resourceFirewallSecurityPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallSecurityPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSecurityPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSecurityPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallSecurityPolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallSecurityPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSecurityPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSecurityPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSecurityPolicyPolicyBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyAppCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyAppGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallSecurityPolicyApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyAvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyCifsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyCasbProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyDiameterFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyDlpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyDnsfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyDstaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyDstaddr6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyEnforceDefaultAppPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyFileFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyFssoGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyGlobalLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyIcapProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetServiceFortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceSrcCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetServiceSrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetServiceSrcFortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetServiceSrcGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetServiceSrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6Custom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetService6CustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetService6Fortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetService6Group(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetService6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetService6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6Src(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6SrcCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetService6SrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetService6SrcFortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetService6SrcGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetService6SrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyInternetService6SrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyIpsVoipFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyLearningMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyLogtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyNat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyNat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyProfileGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicySctpFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicySendDenyPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicySrcaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicySrcaddr6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicySshFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicySslSshProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyTelemetryProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyUrlCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyUtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyVideofilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyVirtualPatchProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyVoipProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSecurityPolicyWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallSecurityPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("_policy_block", flattenFirewallSecurityPolicyPolicyBlock(o["_policy_block"], d, "_policy_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["_policy_block"], "FirewallSecurityPolicy-PolicyBlock"); ok {
			if err = d.Set("_policy_block", vv); err != nil {
				return fmt.Errorf("Error reading _policy_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _policy_block: %v", err)
		}
	}

	if err = d.Set("action", flattenFirewallSecurityPolicyAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "FirewallSecurityPolicy-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("app_category", flattenFirewallSecurityPolicyAppCategory(o["app-category"], d, "app_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-category"], "FirewallSecurityPolicy-AppCategory"); ok {
			if err = d.Set("app_category", vv); err != nil {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_category: %v", err)
		}
	}

	if err = d.Set("app_group", flattenFirewallSecurityPolicyAppGroup(o["app-group"], d, "app_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-group"], "FirewallSecurityPolicy-AppGroup"); ok {
			if err = d.Set("app_group", vv); err != nil {
				return fmt.Errorf("Error reading app_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_group: %v", err)
		}
	}

	if err = d.Set("application", flattenFirewallSecurityPolicyApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "FirewallSecurityPolicy-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallSecurityPolicyApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "FirewallSecurityPolicy-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallSecurityPolicyAvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile"], "FirewallSecurityPolicy-AvProfile"); ok {
			if err = d.Set("av_profile", vv); err != nil {
				return fmt.Errorf("Error reading av_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("cifs_profile", flattenFirewallSecurityPolicyCifsProfile(o["cifs-profile"], d, "cifs_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["cifs-profile"], "FirewallSecurityPolicy-CifsProfile"); ok {
			if err = d.Set("cifs_profile", vv); err != nil {
				return fmt.Errorf("Error reading cifs_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("casb_profile", flattenFirewallSecurityPolicyCasbProfile(o["casb-profile"], d, "casb_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["casb-profile"], "FirewallSecurityPolicy-CasbProfile"); ok {
			if err = d.Set("casb_profile", vv); err != nil {
				return fmt.Errorf("Error reading casb_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading casb_profile: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallSecurityPolicyComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallSecurityPolicy-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallSecurityPolicyDlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor"], "FirewallSecurityPolicy-DlpSensor"); ok {
			if err = d.Set("dlp_sensor", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("diameter_filter_profile", flattenFirewallSecurityPolicyDiameterFilterProfile(o["diameter-filter-profile"], d, "diameter_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["diameter-filter-profile"], "FirewallSecurityPolicy-DiameterFilterProfile"); ok {
			if err = d.Set("diameter_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading diameter_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diameter_filter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenFirewallSecurityPolicyDlpProfile(o["dlp-profile"], d, "dlp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-profile"], "FirewallSecurityPolicy-DlpProfile"); ok {
			if err = d.Set("dlp_profile", vv); err != nil {
				return fmt.Errorf("Error reading dlp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenFirewallSecurityPolicyDnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnsfilter-profile"], "FirewallSecurityPolicy-DnsfilterProfile"); ok {
			if err = d.Set("dnsfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenFirewallSecurityPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "FirewallSecurityPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenFirewallSecurityPolicyDstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr-negate"], "FirewallSecurityPolicy-DstaddrNegate"); ok {
			if err = d.Set("dstaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenFirewallSecurityPolicyDstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "FirewallSecurityPolicy-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("dstaddr6_negate", flattenFirewallSecurityPolicyDstaddr6Negate(o["dstaddr6-negate"], d, "dstaddr6_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6-negate"], "FirewallSecurityPolicy-Dstaddr6Negate"); ok {
			if err = d.Set("dstaddr6_negate", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6_negate: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenFirewallSecurityPolicyDstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "FirewallSecurityPolicy-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallSecurityPolicyEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile"], "FirewallSecurityPolicy-EmailfilterProfile"); ok {
			if err = d.Set("emailfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("enforce_default_app_port", flattenFirewallSecurityPolicyEnforceDefaultAppPort(o["enforce-default-app-port"], d, "enforce_default_app_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-default-app-port"], "FirewallSecurityPolicy-EnforceDefaultAppPort"); ok {
			if err = d.Set("enforce_default_app_port", vv); err != nil {
				return fmt.Errorf("Error reading enforce_default_app_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_default_app_port: %v", err)
		}
	}

	if err = d.Set("file_filter_profile", flattenFirewallSecurityPolicyFileFilterProfile(o["file-filter-profile"], d, "file_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-filter-profile"], "FirewallSecurityPolicy-FileFilterProfile"); ok {
			if err = d.Set("file_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading file_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_filter_profile: %v", err)
		}
	}

	if err = d.Set("fsso_groups", flattenFirewallSecurityPolicyFssoGroups(o["fsso-groups"], d, "fsso_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-groups"], "FirewallSecurityPolicy-FssoGroups"); ok {
			if err = d.Set("fsso_groups", vv); err != nil {
				return fmt.Errorf("Error reading fsso_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_groups: %v", err)
		}
	}

	if err = d.Set("global_label", flattenFirewallSecurityPolicyGlobalLabel(o["global-label"], d, "global_label")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-label"], "FirewallSecurityPolicy-GlobalLabel"); ok {
			if err = d.Set("global_label", vv); err != nil {
				return fmt.Errorf("Error reading global_label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_label: %v", err)
		}
	}

	if err = d.Set("groups", flattenFirewallSecurityPolicyGroups(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "FirewallSecurityPolicy-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenFirewallSecurityPolicyIcapProfile(o["icap-profile"], d, "icap_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["icap-profile"], "FirewallSecurityPolicy-IcapProfile"); ok {
			if err = d.Set("icap_profile", vv); err != nil {
				return fmt.Errorf("Error reading icap_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenFirewallSecurityPolicyInternetService(o["internet-service"], d, "internet_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service"], "FirewallSecurityPolicy-InternetService"); ok {
			if err = d.Set("internet_service", vv); err != nil {
				return fmt.Errorf("Error reading internet_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", flattenFirewallSecurityPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom"], "FirewallSecurityPolicy-InternetServiceCustom"); ok {
			if err = d.Set("internet_service_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_custom_group", flattenFirewallSecurityPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom-group"], "FirewallSecurityPolicy-InternetServiceCustomGroup"); ok {
			if err = d.Set("internet_service_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_fortiguard", flattenFirewallSecurityPolicyInternetServiceFortiguard(o["internet-service-fortiguard"], d, "internet_service_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-fortiguard"], "FirewallSecurityPolicy-InternetServiceFortiguard"); ok {
			if err = d.Set("internet_service_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service_group", flattenFirewallSecurityPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-group"], "FirewallSecurityPolicy-InternetServiceGroup"); ok {
			if err = d.Set("internet_service_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_group: %v", err)
		}
	}

	if err = d.Set("internet_service_name", flattenFirewallSecurityPolicyInternetServiceName(o["internet-service-name"], d, "internet_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-name"], "FirewallSecurityPolicy-InternetServiceName"); ok {
			if err = d.Set("internet_service_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_name: %v", err)
		}
	}

	if err = d.Set("internet_service_negate", flattenFirewallSecurityPolicyInternetServiceNegate(o["internet-service-negate"], d, "internet_service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-negate"], "FirewallSecurityPolicy-InternetServiceNegate"); ok {
			if err = d.Set("internet_service_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service_src", flattenFirewallSecurityPolicyInternetServiceSrc(o["internet-service-src"], d, "internet_service_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src"], "FirewallSecurityPolicy-InternetServiceSrc"); ok {
			if err = d.Set("internet_service_src", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom", flattenFirewallSecurityPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom"], "FirewallSecurityPolicy-InternetServiceSrcCustom"); ok {
			if err = d.Set("internet_service_src_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom_group", flattenFirewallSecurityPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom-group"], "FirewallSecurityPolicy-InternetServiceSrcCustomGroup"); ok {
			if err = d.Set("internet_service_src_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_fortiguard", flattenFirewallSecurityPolicyInternetServiceSrcFortiguard(o["internet-service-src-fortiguard"], d, "internet_service_src_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-fortiguard"], "FirewallSecurityPolicy-InternetServiceSrcFortiguard"); ok {
			if err = d.Set("internet_service_src_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service_src_group", flattenFirewallSecurityPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-group"], "FirewallSecurityPolicy-InternetServiceSrcGroup"); ok {
			if err = d.Set("internet_service_src_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_name", flattenFirewallSecurityPolicyInternetServiceSrcName(o["internet-service-src-name"], d, "internet_service_src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-name"], "FirewallSecurityPolicy-InternetServiceSrcName"); ok {
			if err = d.Set("internet_service_src_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_name: %v", err)
		}
	}

	if err = d.Set("internet_service_src_negate", flattenFirewallSecurityPolicyInternetServiceSrcNegate(o["internet-service-src-negate"], d, "internet_service_src_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-negate"], "FirewallSecurityPolicy-InternetServiceSrcNegate"); ok {
			if err = d.Set("internet_service_src_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
		}
	}

	if err = d.Set("internet_service6", flattenFirewallSecurityPolicyInternetService6(o["internet-service6"], d, "internet_service6")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6"], "FirewallSecurityPolicy-InternetService6"); ok {
			if err = d.Set("internet_service6", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6: %v", err)
		}
	}

	if err = d.Set("internet_service6_custom", flattenFirewallSecurityPolicyInternetService6Custom(o["internet-service6-custom"], d, "internet_service6_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-custom"], "FirewallSecurityPolicy-InternetService6Custom"); ok {
			if err = d.Set("internet_service6_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_custom: %v", err)
		}
	}

	if err = d.Set("internet_service6_custom_group", flattenFirewallSecurityPolicyInternetService6CustomGroup(o["internet-service6-custom-group"], d, "internet_service6_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-custom-group"], "FirewallSecurityPolicy-InternetService6CustomGroup"); ok {
			if err = d.Set("internet_service6_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_fortiguard", flattenFirewallSecurityPolicyInternetService6Fortiguard(o["internet-service6-fortiguard"], d, "internet_service6_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-fortiguard"], "FirewallSecurityPolicy-InternetService6Fortiguard"); ok {
			if err = d.Set("internet_service6_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service6_group", flattenFirewallSecurityPolicyInternetService6Group(o["internet-service6-group"], d, "internet_service6_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-group"], "FirewallSecurityPolicy-InternetService6Group"); ok {
			if err = d.Set("internet_service6_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_name", flattenFirewallSecurityPolicyInternetService6Name(o["internet-service6-name"], d, "internet_service6_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-name"], "FirewallSecurityPolicy-InternetService6Name"); ok {
			if err = d.Set("internet_service6_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_name: %v", err)
		}
	}

	if err = d.Set("internet_service6_negate", flattenFirewallSecurityPolicyInternetService6Negate(o["internet-service6-negate"], d, "internet_service6_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-negate"], "FirewallSecurityPolicy-InternetService6Negate"); ok {
			if err = d.Set("internet_service6_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_negate: %v", err)
		}
	}

	if err = d.Set("internet_service6_src", flattenFirewallSecurityPolicyInternetService6Src(o["internet-service6-src"], d, "internet_service6_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src"], "FirewallSecurityPolicy-InternetService6Src"); ok {
			if err = d.Set("internet_service6_src", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_custom", flattenFirewallSecurityPolicyInternetService6SrcCustom(o["internet-service6-src-custom"], d, "internet_service6_src_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-custom"], "FirewallSecurityPolicy-InternetService6SrcCustom"); ok {
			if err = d.Set("internet_service6_src_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_custom: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_custom_group", flattenFirewallSecurityPolicyInternetService6SrcCustomGroup(o["internet-service6-src-custom-group"], d, "internet_service6_src_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-custom-group"], "FirewallSecurityPolicy-InternetService6SrcCustomGroup"); ok {
			if err = d.Set("internet_service6_src_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_fortiguard", flattenFirewallSecurityPolicyInternetService6SrcFortiguard(o["internet-service6-src-fortiguard"], d, "internet_service6_src_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-fortiguard"], "FirewallSecurityPolicy-InternetService6SrcFortiguard"); ok {
			if err = d.Set("internet_service6_src_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_group", flattenFirewallSecurityPolicyInternetService6SrcGroup(o["internet-service6-src-group"], d, "internet_service6_src_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-group"], "FirewallSecurityPolicy-InternetService6SrcGroup"); ok {
			if err = d.Set("internet_service6_src_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_name", flattenFirewallSecurityPolicyInternetService6SrcName(o["internet-service6-src-name"], d, "internet_service6_src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-name"], "FirewallSecurityPolicy-InternetService6SrcName"); ok {
			if err = d.Set("internet_service6_src_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_name: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_negate", flattenFirewallSecurityPolicyInternetService6SrcNegate(o["internet-service6-src-negate"], d, "internet_service6_src_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-negate"], "FirewallSecurityPolicy-InternetService6SrcNegate"); ok {
			if err = d.Set("internet_service6_src_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_negate: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallSecurityPolicyIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "FirewallSecurityPolicy-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("ips_voip_filter", flattenFirewallSecurityPolicyIpsVoipFilter(o["ips-voip-filter"], d, "ips_voip_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-voip-filter"], "FirewallSecurityPolicy-IpsVoipFilter"); ok {
			if err = d.Set("ips_voip_filter", vv); err != nil {
				return fmt.Errorf("Error reading ips_voip_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_voip_filter: %v", err)
		}
	}

	if err = d.Set("learning_mode", flattenFirewallSecurityPolicyLearningMode(o["learning-mode"], d, "learning_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["learning-mode"], "FirewallSecurityPolicy-LearningMode"); ok {
			if err = d.Set("learning_mode", vv); err != nil {
				return fmt.Errorf("Error reading learning_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learning_mode: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallSecurityPolicyLogtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "FirewallSecurityPolicy-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallSecurityPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallSecurityPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat46", flattenFirewallSecurityPolicyNat46(o["nat46"], d, "nat46")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46"], "FirewallSecurityPolicy-Nat46"); ok {
			if err = d.Set("nat46", vv); err != nil {
				return fmt.Errorf("Error reading nat46: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("nat64", flattenFirewallSecurityPolicyNat64(o["nat64"], d, "nat64")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat64"], "FirewallSecurityPolicy-Nat64"); ok {
			if err = d.Set("nat64", vv); err != nil {
				return fmt.Errorf("Error reading nat64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("policyid", flattenFirewallSecurityPolicyPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "FirewallSecurityPolicy-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("profile_group", flattenFirewallSecurityPolicyProfileGroup(o["profile-group"], d, "profile_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-group"], "FirewallSecurityPolicy-ProfileGroup"); ok {
			if err = d.Set("profile_group", vv); err != nil {
				return fmt.Errorf("Error reading profile_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenFirewallSecurityPolicyProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-protocol-options"], "FirewallSecurityPolicy-ProfileProtocolOptions"); ok {
			if err = d.Set("profile_protocol_options", vv); err != nil {
				return fmt.Errorf("Error reading profile_protocol_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenFirewallSecurityPolicyProfileType(o["profile-type"], d, "profile_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-type"], "FirewallSecurityPolicy-ProfileType"); ok {
			if err = d.Set("profile_type", vv); err != nil {
				return fmt.Errorf("Error reading profile_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	if err = d.Set("schedule", flattenFirewallSecurityPolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "FirewallSecurityPolicy-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("sctp_filter_profile", flattenFirewallSecurityPolicySctpFilterProfile(o["sctp-filter-profile"], d, "sctp_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-filter-profile"], "FirewallSecurityPolicy-SctpFilterProfile"); ok {
			if err = d.Set("sctp_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
		}
	}

	if err = d.Set("send_deny_packet", flattenFirewallSecurityPolicySendDenyPacket(o["send-deny-packet"], d, "send_deny_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-deny-packet"], "FirewallSecurityPolicy-SendDenyPacket"); ok {
			if err = d.Set("send_deny_packet", vv); err != nil {
				return fmt.Errorf("Error reading send_deny_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_deny_packet: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallSecurityPolicyService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "FirewallSecurityPolicy-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenFirewallSecurityPolicyServiceNegate(o["service-negate"], d, "service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-negate"], "FirewallSecurityPolicy-ServiceNegate"); ok {
			if err = d.Set("service_negate", vv); err != nil {
				return fmt.Errorf("Error reading service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenFirewallSecurityPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "FirewallSecurityPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenFirewallSecurityPolicySrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr-negate"], "FirewallSecurityPolicy-SrcaddrNegate"); ok {
			if err = d.Set("srcaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenFirewallSecurityPolicySrcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "FirewallSecurityPolicy-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcaddr6_negate", flattenFirewallSecurityPolicySrcaddr6Negate(o["srcaddr6-negate"], d, "srcaddr6_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6-negate"], "FirewallSecurityPolicy-Srcaddr6Negate"); ok {
			if err = d.Set("srcaddr6_negate", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6_negate: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenFirewallSecurityPolicySrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "FirewallSecurityPolicy-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenFirewallSecurityPolicySshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-filter-profile"], "FirewallSecurityPolicy-SshFilterProfile"); ok {
			if err = d.Set("ssh_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenFirewallSecurityPolicySslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-ssh-profile"], "FirewallSecurityPolicy-SslSshProfile"); ok {
			if err = d.Set("ssl_ssh_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallSecurityPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallSecurityPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("telemetry_profile", flattenFirewallSecurityPolicyTelemetryProfile(o["telemetry-profile"], d, "telemetry_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["telemetry-profile"], "FirewallSecurityPolicy-TelemetryProfile"); ok {
			if err = d.Set("telemetry_profile", vv); err != nil {
				return fmt.Errorf("Error reading telemetry_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading telemetry_profile: %v", err)
		}
	}

	if err = d.Set("url_category", flattenFirewallSecurityPolicyUrlCategory(o["url-category"], d, "url_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-category"], "FirewallSecurityPolicy-UrlCategory"); ok {
			if err = d.Set("url_category", vv); err != nil {
				return fmt.Errorf("Error reading url_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_category: %v", err)
		}
	}

	if err = d.Set("users", flattenFirewallSecurityPolicyUsers(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "FirewallSecurityPolicy-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenFirewallSecurityPolicyUtmStatus(o["utm-status"], d, "utm_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-status"], "FirewallSecurityPolicy-UtmStatus"); ok {
			if err = d.Set("utm_status", vv); err != nil {
				return fmt.Errorf("Error reading utm_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallSecurityPolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallSecurityPolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("videofilter_profile", flattenFirewallSecurityPolicyVideofilterProfile(o["videofilter-profile"], d, "videofilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["videofilter-profile"], "FirewallSecurityPolicy-VideofilterProfile"); ok {
			if err = d.Set("videofilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading videofilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading videofilter_profile: %v", err)
		}
	}

	if err = d.Set("virtual_patch_profile", flattenFirewallSecurityPolicyVirtualPatchProfile(o["virtual-patch-profile"], d, "virtual_patch_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-patch-profile"], "FirewallSecurityPolicy-VirtualPatchProfile"); ok {
			if err = d.Set("virtual_patch_profile", vv); err != nil {
				return fmt.Errorf("Error reading virtual_patch_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_patch_profile: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenFirewallSecurityPolicyVoipProfile(o["voip-profile"], d, "voip_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip-profile"], "FirewallSecurityPolicy-VoipProfile"); ok {
			if err = d.Set("voip_profile", vv); err != nil {
				return fmt.Errorf("Error reading voip_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallSecurityPolicyWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "FirewallSecurityPolicy-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	return nil
}

func flattenFirewallSecurityPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSecurityPolicyPolicyBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyAppCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyAppGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyAvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyCifsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyCasbProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyDiameterFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyDlpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyDnsfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyDstaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyDstaddr6Negate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyEmailfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyEnforceDefaultAppPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyFileFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyFssoGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyGlobalLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyIcapProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetServiceFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetServiceSrcFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6Custom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetService6CustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetService6Fortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetService6Group(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetService6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetService6Negate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6Src(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6SrcCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetService6SrcCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetService6SrcFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetService6SrcGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetService6SrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyInternetService6SrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyIpsVoipFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyLearningMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyLogtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyNat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyNat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyProfileGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicySctpFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicySendDenyPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicySrcaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicySrcaddr6Negate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicySshFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicySslSshProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyTelemetryProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyUrlCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyUtmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyVideofilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyVirtualPatchProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyVoipProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSecurityPolicyWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallSecurityPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_policy_block"); ok || d.HasChange("_policy_block") {
		t, err := expandFirewallSecurityPolicyPolicyBlock(d, v, "_policy_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_policy_block"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandFirewallSecurityPolicyAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("app_category"); ok || d.HasChange("app_category") {
		t, err := expandFirewallSecurityPolicyAppCategory(d, v, "app_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("app_group"); ok || d.HasChange("app_group") {
		t, err := expandFirewallSecurityPolicyAppGroup(d, v, "app_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-group"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandFirewallSecurityPolicyApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandFirewallSecurityPolicyApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok || d.HasChange("av_profile") {
		t, err := expandFirewallSecurityPolicyAvProfile(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("cifs_profile"); ok || d.HasChange("cifs_profile") {
		t, err := expandFirewallSecurityPolicyCifsProfile(d, v, "cifs_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs-profile"] = t
		}
	}

	if v, ok := d.GetOk("casb_profile"); ok || d.HasChange("casb_profile") {
		t, err := expandFirewallSecurityPolicyCasbProfile(d, v, "casb_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallSecurityPolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok || d.HasChange("dlp_sensor") {
		t, err := expandFirewallSecurityPolicyDlpSensor(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("diameter_filter_profile"); ok || d.HasChange("diameter_filter_profile") {
		t, err := expandFirewallSecurityPolicyDiameterFilterProfile(d, v, "diameter_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diameter-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile"); ok || d.HasChange("dlp_profile") {
		t, err := expandFirewallSecurityPolicyDlpProfile(d, v, "dlp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok || d.HasChange("dnsfilter_profile") {
		t, err := expandFirewallSecurityPolicyDnsfilterProfile(d, v, "dnsfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallSecurityPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok || d.HasChange("dstaddr_negate") {
		t, err := expandFirewallSecurityPolicyDstaddrNegate(d, v, "dstaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandFirewallSecurityPolicyDstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6_negate"); ok || d.HasChange("dstaddr6_negate") {
		t, err := expandFirewallSecurityPolicyDstaddr6Negate(d, v, "dstaddr6_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandFirewallSecurityPolicyDstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok || d.HasChange("emailfilter_profile") {
		t, err := expandFirewallSecurityPolicyEmailfilterProfile(d, v, "emailfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("enforce_default_app_port"); ok || d.HasChange("enforce_default_app_port") {
		t, err := expandFirewallSecurityPolicyEnforceDefaultAppPort(d, v, "enforce_default_app_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-default-app-port"] = t
		}
	}

	if v, ok := d.GetOk("file_filter_profile"); ok || d.HasChange("file_filter_profile") {
		t, err := expandFirewallSecurityPolicyFileFilterProfile(d, v, "file_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("fsso_groups"); ok || d.HasChange("fsso_groups") {
		t, err := expandFirewallSecurityPolicyFssoGroups(d, v, "fsso_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-groups"] = t
		}
	}

	if v, ok := d.GetOk("global_label"); ok || d.HasChange("global_label") {
		t, err := expandFirewallSecurityPolicyGlobalLabel(d, v, "global_label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-label"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandFirewallSecurityPolicyGroups(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("icap_profile"); ok || d.HasChange("icap_profile") {
		t, err := expandFirewallSecurityPolicyIcapProfile(d, v, "icap_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok || d.HasChange("internet_service") {
		t, err := expandFirewallSecurityPolicyInternetService(d, v, "internet_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok || d.HasChange("internet_service_custom") {
		t, err := expandFirewallSecurityPolicyInternetServiceCustom(d, v, "internet_service_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok || d.HasChange("internet_service_custom_group") {
		t, err := expandFirewallSecurityPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_fortiguard"); ok || d.HasChange("internet_service_fortiguard") {
		t, err := expandFirewallSecurityPolicyInternetServiceFortiguard(d, v, "internet_service_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok || d.HasChange("internet_service_group") {
		t, err := expandFirewallSecurityPolicyInternetServiceGroup(d, v, "internet_service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok || d.HasChange("internet_service_name") {
		t, err := expandFirewallSecurityPolicyInternetServiceName(d, v, "internet_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_negate"); ok || d.HasChange("internet_service_negate") {
		t, err := expandFirewallSecurityPolicyInternetServiceNegate(d, v, "internet_service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src"); ok || d.HasChange("internet_service_src") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrc(d, v, "internet_service_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom"); ok || d.HasChange("internet_service_src_custom") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom_group"); ok || d.HasChange("internet_service_src_custom_group") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_fortiguard"); ok || d.HasChange("internet_service_src_fortiguard") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcFortiguard(d, v, "internet_service_src_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_group"); ok || d.HasChange("internet_service_src_group") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_name"); ok || d.HasChange("internet_service_src_name") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcName(d, v, "internet_service_src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_negate"); ok || d.HasChange("internet_service_src_negate") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcNegate(d, v, "internet_service_src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6"); ok || d.HasChange("internet_service6") {
		t, err := expandFirewallSecurityPolicyInternetService6(d, v, "internet_service6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_custom"); ok || d.HasChange("internet_service6_custom") {
		t, err := expandFirewallSecurityPolicyInternetService6Custom(d, v, "internet_service6_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_custom_group"); ok || d.HasChange("internet_service6_custom_group") {
		t, err := expandFirewallSecurityPolicyInternetService6CustomGroup(d, v, "internet_service6_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_fortiguard"); ok || d.HasChange("internet_service6_fortiguard") {
		t, err := expandFirewallSecurityPolicyInternetService6Fortiguard(d, v, "internet_service6_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_group"); ok || d.HasChange("internet_service6_group") {
		t, err := expandFirewallSecurityPolicyInternetService6Group(d, v, "internet_service6_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_name"); ok || d.HasChange("internet_service6_name") {
		t, err := expandFirewallSecurityPolicyInternetService6Name(d, v, "internet_service6_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_negate"); ok || d.HasChange("internet_service6_negate") {
		t, err := expandFirewallSecurityPolicyInternetService6Negate(d, v, "internet_service6_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src"); ok || d.HasChange("internet_service6_src") {
		t, err := expandFirewallSecurityPolicyInternetService6Src(d, v, "internet_service6_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_custom"); ok || d.HasChange("internet_service6_src_custom") {
		t, err := expandFirewallSecurityPolicyInternetService6SrcCustom(d, v, "internet_service6_src_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_custom_group"); ok || d.HasChange("internet_service6_src_custom_group") {
		t, err := expandFirewallSecurityPolicyInternetService6SrcCustomGroup(d, v, "internet_service6_src_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_fortiguard"); ok || d.HasChange("internet_service6_src_fortiguard") {
		t, err := expandFirewallSecurityPolicyInternetService6SrcFortiguard(d, v, "internet_service6_src_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_group"); ok || d.HasChange("internet_service6_src_group") {
		t, err := expandFirewallSecurityPolicyInternetService6SrcGroup(d, v, "internet_service6_src_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_name"); ok || d.HasChange("internet_service6_src_name") {
		t, err := expandFirewallSecurityPolicyInternetService6SrcName(d, v, "internet_service6_src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_negate"); ok || d.HasChange("internet_service6_src_negate") {
		t, err := expandFirewallSecurityPolicyInternetService6SrcNegate(d, v, "internet_service6_src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandFirewallSecurityPolicyIpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ips_voip_filter"); ok || d.HasChange("ips_voip_filter") {
		t, err := expandFirewallSecurityPolicyIpsVoipFilter(d, v, "ips_voip_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-voip-filter"] = t
		}
	}

	if v, ok := d.GetOk("learning_mode"); ok || d.HasChange("learning_mode") {
		t, err := expandFirewallSecurityPolicyLearningMode(d, v, "learning_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learning-mode"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok || d.HasChange("logtraffic") {
		t, err := expandFirewallSecurityPolicyLogtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallSecurityPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat46"); ok || d.HasChange("nat46") {
		t, err := expandFirewallSecurityPolicyNat46(d, v, "nat46")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok || d.HasChange("nat64") {
		t, err := expandFirewallSecurityPolicyNat64(d, v, "nat64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandFirewallSecurityPolicyPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("profile_group"); ok || d.HasChange("profile_group") {
		t, err := expandFirewallSecurityPolicyProfileGroup(d, v, "profile_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-group"] = t
		}
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok || d.HasChange("profile_protocol_options") {
		t, err := expandFirewallSecurityPolicyProfileProtocolOptions(d, v, "profile_protocol_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("profile_type"); ok || d.HasChange("profile_type") {
		t, err := expandFirewallSecurityPolicyProfileType(d, v, "profile_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandFirewallSecurityPolicySchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("sctp_filter_profile"); ok || d.HasChange("sctp_filter_profile") {
		t, err := expandFirewallSecurityPolicySctpFilterProfile(d, v, "sctp_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("send_deny_packet"); ok || d.HasChange("send_deny_packet") {
		t, err := expandFirewallSecurityPolicySendDenyPacket(d, v, "send_deny_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-deny-packet"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallSecurityPolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok || d.HasChange("service_negate") {
		t, err := expandFirewallSecurityPolicyServiceNegate(d, v, "service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallSecurityPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok || d.HasChange("srcaddr_negate") {
		t, err := expandFirewallSecurityPolicySrcaddrNegate(d, v, "srcaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandFirewallSecurityPolicySrcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6_negate"); ok || d.HasChange("srcaddr6_negate") {
		t, err := expandFirewallSecurityPolicySrcaddr6Negate(d, v, "srcaddr6_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandFirewallSecurityPolicySrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok || d.HasChange("ssh_filter_profile") {
		t, err := expandFirewallSecurityPolicySshFilterProfile(d, v, "ssh_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok || d.HasChange("ssl_ssh_profile") {
		t, err := expandFirewallSecurityPolicySslSshProfile(d, v, "ssl_ssh_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallSecurityPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("telemetry_profile"); ok || d.HasChange("telemetry_profile") {
		t, err := expandFirewallSecurityPolicyTelemetryProfile(d, v, "telemetry_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["telemetry-profile"] = t
		}
	}

	if v, ok := d.GetOk("url_category"); ok || d.HasChange("url_category") {
		t, err := expandFirewallSecurityPolicyUrlCategory(d, v, "url_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-category"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandFirewallSecurityPolicyUsers(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok || d.HasChange("utm_status") {
		t, err := expandFirewallSecurityPolicyUtmStatus(d, v, "utm_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallSecurityPolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("videofilter_profile"); ok || d.HasChange("videofilter_profile") {
		t, err := expandFirewallSecurityPolicyVideofilterProfile(d, v, "videofilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["videofilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("virtual_patch_profile"); ok || d.HasChange("virtual_patch_profile") {
		t, err := expandFirewallSecurityPolicyVirtualPatchProfile(d, v, "virtual_patch_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-patch-profile"] = t
		}
	}

	if v, ok := d.GetOk("voip_profile"); ok || d.HasChange("voip_profile") {
		t, err := expandFirewallSecurityPolicyVoipProfile(d, v, "voip_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandFirewallSecurityPolicyWebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	return &obj, nil
}
