// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure proxy policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProxyPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProxyPolicyCreate,
		Read:   resourceFirewallProxyPolicyRead,
		Update: resourceFirewallProxyPolicyUpdate,
		Delete: resourceFirewallProxyPolicyDelete,

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
			"access_proxy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"access_proxy6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"block_notification": &schema.Schema{
				Type:     schema.TypeString,
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
			"decrypted_traffic_mirror": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"detect_https_in_http_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_ownership": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diameter_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"disclaimer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dlp_sensor": &schema.Schema{
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
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
			"file_filter_profile": &schema.Schema{
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
			"http_tunnel_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_sub_category": &schema.Schema{
				Type:     schema.TypeString,
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
				Computed: true,
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
			"internet_service6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
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
			"isolator_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_http_transaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"poolname": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Computed: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirect_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacemsg_override_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
			"service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
			"ssh_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
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
				Computed: true,
			},
			"telemetry_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"transparent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url_risk": &schema.Schema{
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
				Computed: true,
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
			"voip_profile": &schema.Schema{
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
			"waf_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"webcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webcache_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"webproxy_forward_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"webproxy_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ztna_ems_tag": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ztna_ems_tag_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ztna_proxy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ztna_tags_match_logic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallProxyPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProxyPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallProxyPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("policyid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallProxyPolicy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallProxyPolicy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallProxyPolicy resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallProxyPolicy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallProxyPolicy resource: %v", err)
		}

		if v != nil && v["policyid"] != nil {
			if vidn, ok := v["policyid"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallProxyPolicyRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallProxyPolicy resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallProxyPolicyRead(d, m)
}

func resourceFirewallProxyPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProxyPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProxyPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallProxyPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProxyPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallProxyPolicyRead(d, m)
}

func resourceFirewallProxyPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallProxyPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProxyPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProxyPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallProxyPolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallProxyPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProxyPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProxyPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProxyPolicyPolicyBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyAccessProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyAccessProxy6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyAvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyCifsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyBlockNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyCasbProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyDetectHttpsInHttpRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyDeviceOwnership(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyDiameterFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyDisclaimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyDlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyDlpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyDnsfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyDstaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyFileFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyGlobalLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyHttpTunnelAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyHttpsSubCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyIcapProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyInternetServiceFortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyInternetServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyInternetService6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyInternetService6Custom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyInternetService6CustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyInternetService6Fortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyInternetService6Group(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyInternetService6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyInternetService6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyIpsVoipFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyIsolatorServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyLogHttpTransaction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyLogtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyLogtrafficStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyPoolname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyProfileGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyReplacemsgOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicySctpFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicySessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicySrcaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicySrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicySshFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicySshPolicyRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicySslSshProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyTelemetryProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyTransparent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyUrlRisk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyUtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyVideofilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyVoipProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyVirtualPatchProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyWafProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyWebcache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyWebcacheHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyWebproxyForwardServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyWebproxyProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyZtnaEmsTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyZtnaEmsTagNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyPolicyZtnaProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyPolicyZtnaTagsMatchLogic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallProxyPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("_policy_block", flattenFirewallProxyPolicyPolicyBlock(o["_policy_block"], d, "_policy_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["_policy_block"], "FirewallProxyPolicy-PolicyBlock"); ok {
			if err = d.Set("_policy_block", vv); err != nil {
				return fmt.Errorf("Error reading _policy_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _policy_block: %v", err)
		}
	}

	if err = d.Set("access_proxy", flattenFirewallProxyPolicyAccessProxy(o["access-proxy"], d, "access_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-proxy"], "FirewallProxyPolicy-AccessProxy"); ok {
			if err = d.Set("access_proxy", vv); err != nil {
				return fmt.Errorf("Error reading access_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_proxy: %v", err)
		}
	}

	if err = d.Set("access_proxy6", flattenFirewallProxyPolicyAccessProxy6(o["access-proxy6"], d, "access_proxy6")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-proxy6"], "FirewallProxyPolicy-AccessProxy6"); ok {
			if err = d.Set("access_proxy6", vv); err != nil {
				return fmt.Errorf("Error reading access_proxy6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_proxy6: %v", err)
		}
	}

	if err = d.Set("action", flattenFirewallProxyPolicyAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "FirewallProxyPolicy-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallProxyPolicyApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "FirewallProxyPolicy-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallProxyPolicyAvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile"], "FirewallProxyPolicy-AvProfile"); ok {
			if err = d.Set("av_profile", vv); err != nil {
				return fmt.Errorf("Error reading av_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("cifs_profile", flattenFirewallProxyPolicyCifsProfile(o["cifs-profile"], d, "cifs_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["cifs-profile"], "FirewallProxyPolicy-CifsProfile"); ok {
			if err = d.Set("cifs_profile", vv); err != nil {
				return fmt.Errorf("Error reading cifs_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("block_notification", flattenFirewallProxyPolicyBlockNotification(o["block-notification"], d, "block_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-notification"], "FirewallProxyPolicy-BlockNotification"); ok {
			if err = d.Set("block_notification", vv); err != nil {
				return fmt.Errorf("Error reading block_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_notification: %v", err)
		}
	}

	if err = d.Set("casb_profile", flattenFirewallProxyPolicyCasbProfile(o["casb-profile"], d, "casb_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["casb-profile"], "FirewallProxyPolicy-CasbProfile"); ok {
			if err = d.Set("casb_profile", vv); err != nil {
				return fmt.Errorf("Error reading casb_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading casb_profile: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallProxyPolicyComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallProxyPolicy-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenFirewallProxyPolicyDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["decrypted-traffic-mirror"], "FirewallProxyPolicy-DecryptedTrafficMirror"); ok {
			if err = d.Set("decrypted_traffic_mirror", vv); err != nil {
				return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("detect_https_in_http_request", flattenFirewallProxyPolicyDetectHttpsInHttpRequest(o["detect-https-in-http-request"], d, "detect_https_in_http_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-https-in-http-request"], "FirewallProxyPolicy-DetectHttpsInHttpRequest"); ok {
			if err = d.Set("detect_https_in_http_request", vv); err != nil {
				return fmt.Errorf("Error reading detect_https_in_http_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_https_in_http_request: %v", err)
		}
	}

	if err = d.Set("device_ownership", flattenFirewallProxyPolicyDeviceOwnership(o["device-ownership"], d, "device_ownership")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-ownership"], "FirewallProxyPolicy-DeviceOwnership"); ok {
			if err = d.Set("device_ownership", vv); err != nil {
				return fmt.Errorf("Error reading device_ownership: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_ownership: %v", err)
		}
	}

	if err = d.Set("diameter_filter_profile", flattenFirewallProxyPolicyDiameterFilterProfile(o["diameter-filter-profile"], d, "diameter_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["diameter-filter-profile"], "FirewallProxyPolicy-DiameterFilterProfile"); ok {
			if err = d.Set("diameter_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading diameter_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diameter_filter_profile: %v", err)
		}
	}

	if err = d.Set("disclaimer", flattenFirewallProxyPolicyDisclaimer(o["disclaimer"], d, "disclaimer")); err != nil {
		if vv, ok := fortiAPIPatch(o["disclaimer"], "FirewallProxyPolicy-Disclaimer"); ok {
			if err = d.Set("disclaimer", vv); err != nil {
				return fmt.Errorf("Error reading disclaimer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disclaimer: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallProxyPolicyDlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor"], "FirewallProxyPolicy-DlpSensor"); ok {
			if err = d.Set("dlp_sensor", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenFirewallProxyPolicyDlpProfile(o["dlp-profile"], d, "dlp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-profile"], "FirewallProxyPolicy-DlpProfile"); ok {
			if err = d.Set("dlp_profile", vv); err != nil {
				return fmt.Errorf("Error reading dlp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenFirewallProxyPolicyDnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnsfilter-profile"], "FirewallProxyPolicy-DnsfilterProfile"); ok {
			if err = d.Set("dnsfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenFirewallProxyPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "FirewallProxyPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenFirewallProxyPolicyDstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr-negate"], "FirewallProxyPolicy-DstaddrNegate"); ok {
			if err = d.Set("dstaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenFirewallProxyPolicyDstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "FirewallProxyPolicy-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenFirewallProxyPolicyDstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "FirewallProxyPolicy-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallProxyPolicyEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile"], "FirewallProxyPolicy-EmailfilterProfile"); ok {
			if err = d.Set("emailfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("file_filter_profile", flattenFirewallProxyPolicyFileFilterProfile(o["file-filter-profile"], d, "file_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-filter-profile"], "FirewallProxyPolicy-FileFilterProfile"); ok {
			if err = d.Set("file_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading file_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_filter_profile: %v", err)
		}
	}

	if err = d.Set("global_label", flattenFirewallProxyPolicyGlobalLabel(o["global-label"], d, "global_label")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-label"], "FirewallProxyPolicy-GlobalLabel"); ok {
			if err = d.Set("global_label", vv); err != nil {
				return fmt.Errorf("Error reading global_label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_label: %v", err)
		}
	}

	if err = d.Set("groups", flattenFirewallProxyPolicyGroups(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "FirewallProxyPolicy-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("http_tunnel_auth", flattenFirewallProxyPolicyHttpTunnelAuth(o["http-tunnel-auth"], d, "http_tunnel_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-tunnel-auth"], "FirewallProxyPolicy-HttpTunnelAuth"); ok {
			if err = d.Set("http_tunnel_auth", vv); err != nil {
				return fmt.Errorf("Error reading http_tunnel_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_tunnel_auth: %v", err)
		}
	}

	if err = d.Set("https_sub_category", flattenFirewallProxyPolicyHttpsSubCategory(o["https-sub-category"], d, "https_sub_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-sub-category"], "FirewallProxyPolicy-HttpsSubCategory"); ok {
			if err = d.Set("https_sub_category", vv); err != nil {
				return fmt.Errorf("Error reading https_sub_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_sub_category: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenFirewallProxyPolicyIcapProfile(o["icap-profile"], d, "icap_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["icap-profile"], "FirewallProxyPolicy-IcapProfile"); ok {
			if err = d.Set("icap_profile", vv); err != nil {
				return fmt.Errorf("Error reading icap_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenFirewallProxyPolicyInternetService(o["internet-service"], d, "internet_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service"], "FirewallProxyPolicy-InternetService"); ok {
			if err = d.Set("internet_service", vv); err != nil {
				return fmt.Errorf("Error reading internet_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", flattenFirewallProxyPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom"], "FirewallProxyPolicy-InternetServiceCustom"); ok {
			if err = d.Set("internet_service_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_custom_group", flattenFirewallProxyPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom-group"], "FirewallProxyPolicy-InternetServiceCustomGroup"); ok {
			if err = d.Set("internet_service_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_fortiguard", flattenFirewallProxyPolicyInternetServiceFortiguard(o["internet-service-fortiguard"], d, "internet_service_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-fortiguard"], "FirewallProxyPolicy-InternetServiceFortiguard"); ok {
			if err = d.Set("internet_service_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service_group", flattenFirewallProxyPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-group"], "FirewallProxyPolicy-InternetServiceGroup"); ok {
			if err = d.Set("internet_service_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_group: %v", err)
		}
	}

	if err = d.Set("internet_service_name", flattenFirewallProxyPolicyInternetServiceName(o["internet-service-name"], d, "internet_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-name"], "FirewallProxyPolicy-InternetServiceName"); ok {
			if err = d.Set("internet_service_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_name: %v", err)
		}
	}

	if err = d.Set("internet_service_negate", flattenFirewallProxyPolicyInternetServiceNegate(o["internet-service-negate"], d, "internet_service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-negate"], "FirewallProxyPolicy-InternetServiceNegate"); ok {
			if err = d.Set("internet_service_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service6", flattenFirewallProxyPolicyInternetService6(o["internet-service6"], d, "internet_service6")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6"], "FirewallProxyPolicy-InternetService6"); ok {
			if err = d.Set("internet_service6", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6: %v", err)
		}
	}

	if err = d.Set("internet_service6_custom", flattenFirewallProxyPolicyInternetService6Custom(o["internet-service6-custom"], d, "internet_service6_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-custom"], "FirewallProxyPolicy-InternetService6Custom"); ok {
			if err = d.Set("internet_service6_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_custom: %v", err)
		}
	}

	if err = d.Set("internet_service6_custom_group", flattenFirewallProxyPolicyInternetService6CustomGroup(o["internet-service6-custom-group"], d, "internet_service6_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-custom-group"], "FirewallProxyPolicy-InternetService6CustomGroup"); ok {
			if err = d.Set("internet_service6_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_fortiguard", flattenFirewallProxyPolicyInternetService6Fortiguard(o["internet-service6-fortiguard"], d, "internet_service6_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-fortiguard"], "FirewallProxyPolicy-InternetService6Fortiguard"); ok {
			if err = d.Set("internet_service6_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service6_group", flattenFirewallProxyPolicyInternetService6Group(o["internet-service6-group"], d, "internet_service6_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-group"], "FirewallProxyPolicy-InternetService6Group"); ok {
			if err = d.Set("internet_service6_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_name", flattenFirewallProxyPolicyInternetService6Name(o["internet-service6-name"], d, "internet_service6_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-name"], "FirewallProxyPolicy-InternetService6Name"); ok {
			if err = d.Set("internet_service6_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_name: %v", err)
		}
	}

	if err = d.Set("internet_service6_negate", flattenFirewallProxyPolicyInternetService6Negate(o["internet-service6-negate"], d, "internet_service6_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-negate"], "FirewallProxyPolicy-InternetService6Negate"); ok {
			if err = d.Set("internet_service6_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_negate: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallProxyPolicyIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "FirewallProxyPolicy-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("ips_voip_filter", flattenFirewallProxyPolicyIpsVoipFilter(o["ips-voip-filter"], d, "ips_voip_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-voip-filter"], "FirewallProxyPolicy-IpsVoipFilter"); ok {
			if err = d.Set("ips_voip_filter", vv); err != nil {
				return fmt.Errorf("Error reading ips_voip_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_voip_filter: %v", err)
		}
	}

	if err = d.Set("isolator_server", flattenFirewallProxyPolicyIsolatorServer(o["isolator-server"], d, "isolator_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["isolator-server"], "FirewallProxyPolicy-IsolatorServer"); ok {
			if err = d.Set("isolator_server", vv); err != nil {
				return fmt.Errorf("Error reading isolator_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading isolator_server: %v", err)
		}
	}

	if err = d.Set("label", flattenFirewallProxyPolicyLabel(o["label"], d, "label")); err != nil {
		if vv, ok := fortiAPIPatch(o["label"], "FirewallProxyPolicy-Label"); ok {
			if err = d.Set("label", vv); err != nil {
				return fmt.Errorf("Error reading label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("log_http_transaction", flattenFirewallProxyPolicyLogHttpTransaction(o["log-http-transaction"], d, "log_http_transaction")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-http-transaction"], "FirewallProxyPolicy-LogHttpTransaction"); ok {
			if err = d.Set("log_http_transaction", vv); err != nil {
				return fmt.Errorf("Error reading log_http_transaction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_http_transaction: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallProxyPolicyLogtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "FirewallProxyPolicy-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("logtraffic_start", flattenFirewallProxyPolicyLogtrafficStart(o["logtraffic-start"], d, "logtraffic_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic-start"], "FirewallProxyPolicy-LogtrafficStart"); ok {
			if err = d.Set("logtraffic_start", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic_start: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallProxyPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallProxyPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("policyid", flattenFirewallProxyPolicyPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "FirewallProxyPolicy-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("poolname", flattenFirewallProxyPolicyPoolname(o["poolname"], d, "poolname")); err != nil {
		if vv, ok := fortiAPIPatch(o["poolname"], "FirewallProxyPolicy-Poolname"); ok {
			if err = d.Set("poolname", vv); err != nil {
				return fmt.Errorf("Error reading poolname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poolname: %v", err)
		}
	}

	if err = d.Set("profile_group", flattenFirewallProxyPolicyProfileGroup(o["profile-group"], d, "profile_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-group"], "FirewallProxyPolicy-ProfileGroup"); ok {
			if err = d.Set("profile_group", vv); err != nil {
				return fmt.Errorf("Error reading profile_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenFirewallProxyPolicyProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-protocol-options"], "FirewallProxyPolicy-ProfileProtocolOptions"); ok {
			if err = d.Set("profile_protocol_options", vv); err != nil {
				return fmt.Errorf("Error reading profile_protocol_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenFirewallProxyPolicyProfileType(o["profile-type"], d, "profile_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-type"], "FirewallProxyPolicy-ProfileType"); ok {
			if err = d.Set("profile_type", vv); err != nil {
				return fmt.Errorf("Error reading profile_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	if err = d.Set("proxy", flattenFirewallProxyPolicyProxy(o["proxy"], d, "proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy"], "FirewallProxyPolicy-Proxy"); ok {
			if err = d.Set("proxy", vv); err != nil {
				return fmt.Errorf("Error reading proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("redirect_url", flattenFirewallProxyPolicyRedirectUrl(o["redirect-url"], d, "redirect_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-url"], "FirewallProxyPolicy-RedirectUrl"); ok {
			if err = d.Set("redirect_url", vv); err != nil {
				return fmt.Errorf("Error reading redirect_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_url: %v", err)
		}
	}

	if err = d.Set("replacemsg_override_group", flattenFirewallProxyPolicyReplacemsgOverrideGroup(o["replacemsg-override-group"], d, "replacemsg_override_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-override-group"], "FirewallProxyPolicy-ReplacemsgOverrideGroup"); ok {
			if err = d.Set("replacemsg_override_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
		}
	}

	if err = d.Set("schedule", flattenFirewallProxyPolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "FirewallProxyPolicy-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("sctp_filter_profile", flattenFirewallProxyPolicySctpFilterProfile(o["sctp-filter-profile"], d, "sctp_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-filter-profile"], "FirewallProxyPolicy-SctpFilterProfile"); ok {
			if err = d.Set("sctp_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallProxyPolicyService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "FirewallProxyPolicy-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenFirewallProxyPolicyServiceNegate(o["service-negate"], d, "service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-negate"], "FirewallProxyPolicy-ServiceNegate"); ok {
			if err = d.Set("service_negate", vv); err != nil {
				return fmt.Errorf("Error reading service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("session_ttl", flattenFirewallProxyPolicySessionTtl(o["session-ttl"], d, "session_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-ttl"], "FirewallProxyPolicy-SessionTtl"); ok {
			if err = d.Set("session_ttl", vv); err != nil {
				return fmt.Errorf("Error reading session_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenFirewallProxyPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "FirewallProxyPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenFirewallProxyPolicySrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr-negate"], "FirewallProxyPolicy-SrcaddrNegate"); ok {
			if err = d.Set("srcaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenFirewallProxyPolicySrcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "FirewallProxyPolicy-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenFirewallProxyPolicySrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "FirewallProxyPolicy-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenFirewallProxyPolicySshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-filter-profile"], "FirewallProxyPolicy-SshFilterProfile"); ok {
			if err = d.Set("ssh_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("ssh_policy_redirect", flattenFirewallProxyPolicySshPolicyRedirect(o["ssh-policy-redirect"], d, "ssh_policy_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-policy-redirect"], "FirewallProxyPolicy-SshPolicyRedirect"); ok {
			if err = d.Set("ssh_policy_redirect", vv); err != nil {
				return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenFirewallProxyPolicySslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-ssh-profile"], "FirewallProxyPolicy-SslSshProfile"); ok {
			if err = d.Set("ssl_ssh_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallProxyPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallProxyPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("telemetry_profile", flattenFirewallProxyPolicyTelemetryProfile(o["telemetry-profile"], d, "telemetry_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["telemetry-profile"], "FirewallProxyPolicy-TelemetryProfile"); ok {
			if err = d.Set("telemetry_profile", vv); err != nil {
				return fmt.Errorf("Error reading telemetry_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading telemetry_profile: %v", err)
		}
	}

	if err = d.Set("transparent", flattenFirewallProxyPolicyTransparent(o["transparent"], d, "transparent")); err != nil {
		if vv, ok := fortiAPIPatch(o["transparent"], "FirewallProxyPolicy-Transparent"); ok {
			if err = d.Set("transparent", vv); err != nil {
				return fmt.Errorf("Error reading transparent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transparent: %v", err)
		}
	}

	if err = d.Set("url_risk", flattenFirewallProxyPolicyUrlRisk(o["url-risk"], d, "url_risk")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-risk"], "FirewallProxyPolicy-UrlRisk"); ok {
			if err = d.Set("url_risk", vv); err != nil {
				return fmt.Errorf("Error reading url_risk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_risk: %v", err)
		}
	}

	if err = d.Set("users", flattenFirewallProxyPolicyUsers(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "FirewallProxyPolicy-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenFirewallProxyPolicyUtmStatus(o["utm-status"], d, "utm_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-status"], "FirewallProxyPolicy-UtmStatus"); ok {
			if err = d.Set("utm_status", vv); err != nil {
				return fmt.Errorf("Error reading utm_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallProxyPolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallProxyPolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("videofilter_profile", flattenFirewallProxyPolicyVideofilterProfile(o["videofilter-profile"], d, "videofilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["videofilter-profile"], "FirewallProxyPolicy-VideofilterProfile"); ok {
			if err = d.Set("videofilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading videofilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading videofilter_profile: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenFirewallProxyPolicyVoipProfile(o["voip-profile"], d, "voip_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip-profile"], "FirewallProxyPolicy-VoipProfile"); ok {
			if err = d.Set("voip_profile", vv); err != nil {
				return fmt.Errorf("Error reading voip_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("virtual_patch_profile", flattenFirewallProxyPolicyVirtualPatchProfile(o["virtual-patch-profile"], d, "virtual_patch_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-patch-profile"], "FirewallProxyPolicy-VirtualPatchProfile"); ok {
			if err = d.Set("virtual_patch_profile", vv); err != nil {
				return fmt.Errorf("Error reading virtual_patch_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_patch_profile: %v", err)
		}
	}

	if err = d.Set("waf_profile", flattenFirewallProxyPolicyWafProfile(o["waf-profile"], d, "waf_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["waf-profile"], "FirewallProxyPolicy-WafProfile"); ok {
			if err = d.Set("waf_profile", vv); err != nil {
				return fmt.Errorf("Error reading waf_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading waf_profile: %v", err)
		}
	}

	if err = d.Set("webcache", flattenFirewallProxyPolicyWebcache(o["webcache"], d, "webcache")); err != nil {
		if vv, ok := fortiAPIPatch(o["webcache"], "FirewallProxyPolicy-Webcache"); ok {
			if err = d.Set("webcache", vv); err != nil {
				return fmt.Errorf("Error reading webcache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webcache: %v", err)
		}
	}

	if err = d.Set("webcache_https", flattenFirewallProxyPolicyWebcacheHttps(o["webcache-https"], d, "webcache_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["webcache-https"], "FirewallProxyPolicy-WebcacheHttps"); ok {
			if err = d.Set("webcache_https", vv); err != nil {
				return fmt.Errorf("Error reading webcache_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webcache_https: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallProxyPolicyWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "FirewallProxyPolicy-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("webproxy_forward_server", flattenFirewallProxyPolicyWebproxyForwardServer(o["webproxy-forward-server"], d, "webproxy_forward_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["webproxy-forward-server"], "FirewallProxyPolicy-WebproxyForwardServer"); ok {
			if err = d.Set("webproxy_forward_server", vv); err != nil {
				return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
		}
	}

	if err = d.Set("webproxy_profile", flattenFirewallProxyPolicyWebproxyProfile(o["webproxy-profile"], d, "webproxy_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webproxy-profile"], "FirewallProxyPolicy-WebproxyProfile"); ok {
			if err = d.Set("webproxy_profile", vv); err != nil {
				return fmt.Errorf("Error reading webproxy_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webproxy_profile: %v", err)
		}
	}

	if err = d.Set("ztna_ems_tag", flattenFirewallProxyPolicyZtnaEmsTag(o["ztna-ems-tag"], d, "ztna_ems_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-ems-tag"], "FirewallProxyPolicy-ZtnaEmsTag"); ok {
			if err = d.Set("ztna_ems_tag", vv); err != nil {
				return fmt.Errorf("Error reading ztna_ems_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_ems_tag: %v", err)
		}
	}

	if err = d.Set("ztna_ems_tag_negate", flattenFirewallProxyPolicyZtnaEmsTagNegate(o["ztna-ems-tag-negate"], d, "ztna_ems_tag_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-ems-tag-negate"], "FirewallProxyPolicy-ZtnaEmsTagNegate"); ok {
			if err = d.Set("ztna_ems_tag_negate", vv); err != nil {
				return fmt.Errorf("Error reading ztna_ems_tag_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_ems_tag_negate: %v", err)
		}
	}

	if err = d.Set("ztna_proxy", flattenFirewallProxyPolicyZtnaProxy(o["ztna-proxy"], d, "ztna_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-proxy"], "FirewallProxyPolicy-ZtnaProxy"); ok {
			if err = d.Set("ztna_proxy", vv); err != nil {
				return fmt.Errorf("Error reading ztna_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_proxy: %v", err)
		}
	}

	if err = d.Set("ztna_tags_match_logic", flattenFirewallProxyPolicyZtnaTagsMatchLogic(o["ztna-tags-match-logic"], d, "ztna_tags_match_logic")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-tags-match-logic"], "FirewallProxyPolicy-ZtnaTagsMatchLogic"); ok {
			if err = d.Set("ztna_tags_match_logic", vv); err != nil {
				return fmt.Errorf("Error reading ztna_tags_match_logic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_tags_match_logic: %v", err)
		}
	}

	return nil
}

func flattenFirewallProxyPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProxyPolicyPolicyBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyAccessProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyAccessProxy6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyAvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyCifsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyBlockNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyCasbProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyDecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyDetectHttpsInHttpRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyDeviceOwnership(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyDiameterFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyDisclaimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyDlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyDlpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyDnsfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyDstaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyEmailfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyFileFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyGlobalLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyHttpTunnelAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyHttpsSubCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyIcapProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyInternetService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyInternetServiceFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyInternetServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyInternetService6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyInternetService6Custom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyInternetService6CustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyInternetService6Fortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyInternetService6Group(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyInternetService6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyInternetService6Negate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyIpsVoipFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyIsolatorServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyLogHttpTransaction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyLogtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyLogtrafficStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyPoolname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyProfileGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyReplacemsgOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicySchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicySctpFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicySessionTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicySrcaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicySrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicySshFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicySshPolicyRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicySslSshProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyTelemetryProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyTransparent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyUrlRisk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyUtmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyVideofilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyVoipProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyVirtualPatchProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyWafProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyWebcache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyWebcacheHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyWebproxyForwardServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyWebproxyProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyZtnaEmsTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyZtnaEmsTagNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyPolicyZtnaProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyPolicyZtnaTagsMatchLogic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProxyPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_policy_block"); ok || d.HasChange("_policy_block") {
		t, err := expandFirewallProxyPolicyPolicyBlock(d, v, "_policy_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_policy_block"] = t
		}
	}

	if v, ok := d.GetOk("access_proxy"); ok || d.HasChange("access_proxy") {
		t, err := expandFirewallProxyPolicyAccessProxy(d, v, "access_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-proxy"] = t
		}
	}

	if v, ok := d.GetOk("access_proxy6"); ok || d.HasChange("access_proxy6") {
		t, err := expandFirewallProxyPolicyAccessProxy6(d, v, "access_proxy6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-proxy6"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandFirewallProxyPolicyAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandFirewallProxyPolicyApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok || d.HasChange("av_profile") {
		t, err := expandFirewallProxyPolicyAvProfile(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("cifs_profile"); ok || d.HasChange("cifs_profile") {
		t, err := expandFirewallProxyPolicyCifsProfile(d, v, "cifs_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs-profile"] = t
		}
	}

	if v, ok := d.GetOk("block_notification"); ok || d.HasChange("block_notification") {
		t, err := expandFirewallProxyPolicyBlockNotification(d, v, "block_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-notification"] = t
		}
	}

	if v, ok := d.GetOk("casb_profile"); ok || d.HasChange("casb_profile") {
		t, err := expandFirewallProxyPolicyCasbProfile(d, v, "casb_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallProxyPolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok || d.HasChange("decrypted_traffic_mirror") {
		t, err := expandFirewallProxyPolicyDecryptedTrafficMirror(d, v, "decrypted_traffic_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	}

	if v, ok := d.GetOk("detect_https_in_http_request"); ok || d.HasChange("detect_https_in_http_request") {
		t, err := expandFirewallProxyPolicyDetectHttpsInHttpRequest(d, v, "detect_https_in_http_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-https-in-http-request"] = t
		}
	}

	if v, ok := d.GetOk("device_ownership"); ok || d.HasChange("device_ownership") {
		t, err := expandFirewallProxyPolicyDeviceOwnership(d, v, "device_ownership")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-ownership"] = t
		}
	}

	if v, ok := d.GetOk("diameter_filter_profile"); ok || d.HasChange("diameter_filter_profile") {
		t, err := expandFirewallProxyPolicyDiameterFilterProfile(d, v, "diameter_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diameter-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("disclaimer"); ok || d.HasChange("disclaimer") {
		t, err := expandFirewallProxyPolicyDisclaimer(d, v, "disclaimer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disclaimer"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok || d.HasChange("dlp_sensor") {
		t, err := expandFirewallProxyPolicyDlpSensor(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile"); ok || d.HasChange("dlp_profile") {
		t, err := expandFirewallProxyPolicyDlpProfile(d, v, "dlp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok || d.HasChange("dnsfilter_profile") {
		t, err := expandFirewallProxyPolicyDnsfilterProfile(d, v, "dnsfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallProxyPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok || d.HasChange("dstaddr_negate") {
		t, err := expandFirewallProxyPolicyDstaddrNegate(d, v, "dstaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandFirewallProxyPolicyDstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandFirewallProxyPolicyDstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok || d.HasChange("emailfilter_profile") {
		t, err := expandFirewallProxyPolicyEmailfilterProfile(d, v, "emailfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("file_filter_profile"); ok || d.HasChange("file_filter_profile") {
		t, err := expandFirewallProxyPolicyFileFilterProfile(d, v, "file_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("global_label"); ok || d.HasChange("global_label") {
		t, err := expandFirewallProxyPolicyGlobalLabel(d, v, "global_label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-label"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandFirewallProxyPolicyGroups(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("http_tunnel_auth"); ok || d.HasChange("http_tunnel_auth") {
		t, err := expandFirewallProxyPolicyHttpTunnelAuth(d, v, "http_tunnel_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-tunnel-auth"] = t
		}
	}

	if v, ok := d.GetOk("https_sub_category"); ok || d.HasChange("https_sub_category") {
		t, err := expandFirewallProxyPolicyHttpsSubCategory(d, v, "https_sub_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-sub-category"] = t
		}
	}

	if v, ok := d.GetOk("icap_profile"); ok || d.HasChange("icap_profile") {
		t, err := expandFirewallProxyPolicyIcapProfile(d, v, "icap_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok || d.HasChange("internet_service") {
		t, err := expandFirewallProxyPolicyInternetService(d, v, "internet_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok || d.HasChange("internet_service_custom") {
		t, err := expandFirewallProxyPolicyInternetServiceCustom(d, v, "internet_service_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok || d.HasChange("internet_service_custom_group") {
		t, err := expandFirewallProxyPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_fortiguard"); ok || d.HasChange("internet_service_fortiguard") {
		t, err := expandFirewallProxyPolicyInternetServiceFortiguard(d, v, "internet_service_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok || d.HasChange("internet_service_group") {
		t, err := expandFirewallProxyPolicyInternetServiceGroup(d, v, "internet_service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok || d.HasChange("internet_service_name") {
		t, err := expandFirewallProxyPolicyInternetServiceName(d, v, "internet_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_negate"); ok || d.HasChange("internet_service_negate") {
		t, err := expandFirewallProxyPolicyInternetServiceNegate(d, v, "internet_service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6"); ok || d.HasChange("internet_service6") {
		t, err := expandFirewallProxyPolicyInternetService6(d, v, "internet_service6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_custom"); ok || d.HasChange("internet_service6_custom") {
		t, err := expandFirewallProxyPolicyInternetService6Custom(d, v, "internet_service6_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_custom_group"); ok || d.HasChange("internet_service6_custom_group") {
		t, err := expandFirewallProxyPolicyInternetService6CustomGroup(d, v, "internet_service6_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_fortiguard"); ok || d.HasChange("internet_service6_fortiguard") {
		t, err := expandFirewallProxyPolicyInternetService6Fortiguard(d, v, "internet_service6_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_group"); ok || d.HasChange("internet_service6_group") {
		t, err := expandFirewallProxyPolicyInternetService6Group(d, v, "internet_service6_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_name"); ok || d.HasChange("internet_service6_name") {
		t, err := expandFirewallProxyPolicyInternetService6Name(d, v, "internet_service6_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_negate"); ok || d.HasChange("internet_service6_negate") {
		t, err := expandFirewallProxyPolicyInternetService6Negate(d, v, "internet_service6_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-negate"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandFirewallProxyPolicyIpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ips_voip_filter"); ok || d.HasChange("ips_voip_filter") {
		t, err := expandFirewallProxyPolicyIpsVoipFilter(d, v, "ips_voip_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-voip-filter"] = t
		}
	}

	if v, ok := d.GetOk("isolator_server"); ok || d.HasChange("isolator_server") {
		t, err := expandFirewallProxyPolicyIsolatorServer(d, v, "isolator_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isolator-server"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok || d.HasChange("label") {
		t, err := expandFirewallProxyPolicyLabel(d, v, "label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	if v, ok := d.GetOk("log_http_transaction"); ok || d.HasChange("log_http_transaction") {
		t, err := expandFirewallProxyPolicyLogHttpTransaction(d, v, "log_http_transaction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-http-transaction"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok || d.HasChange("logtraffic") {
		t, err := expandFirewallProxyPolicyLogtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic_start"); ok || d.HasChange("logtraffic_start") {
		t, err := expandFirewallProxyPolicyLogtrafficStart(d, v, "logtraffic_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic-start"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallProxyPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandFirewallProxyPolicyPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("poolname"); ok || d.HasChange("poolname") {
		t, err := expandFirewallProxyPolicyPoolname(d, v, "poolname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname"] = t
		}
	}

	if v, ok := d.GetOk("profile_group"); ok || d.HasChange("profile_group") {
		t, err := expandFirewallProxyPolicyProfileGroup(d, v, "profile_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-group"] = t
		}
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok || d.HasChange("profile_protocol_options") {
		t, err := expandFirewallProxyPolicyProfileProtocolOptions(d, v, "profile_protocol_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("profile_type"); ok || d.HasChange("profile_type") {
		t, err := expandFirewallProxyPolicyProfileType(d, v, "profile_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	if v, ok := d.GetOk("proxy"); ok || d.HasChange("proxy") {
		t, err := expandFirewallProxyPolicyProxy(d, v, "proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	if v, ok := d.GetOk("redirect_url"); ok || d.HasChange("redirect_url") {
		t, err := expandFirewallProxyPolicyRedirectUrl(d, v, "redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_override_group"); ok || d.HasChange("replacemsg_override_group") {
		t, err := expandFirewallProxyPolicyReplacemsgOverrideGroup(d, v, "replacemsg_override_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-override-group"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandFirewallProxyPolicySchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("sctp_filter_profile"); ok || d.HasChange("sctp_filter_profile") {
		t, err := expandFirewallProxyPolicySctpFilterProfile(d, v, "sctp_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallProxyPolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok || d.HasChange("service_negate") {
		t, err := expandFirewallProxyPolicyServiceNegate(d, v, "service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("session_ttl"); ok || d.HasChange("session_ttl") {
		t, err := expandFirewallProxyPolicySessionTtl(d, v, "session_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallProxyPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok || d.HasChange("srcaddr_negate") {
		t, err := expandFirewallProxyPolicySrcaddrNegate(d, v, "srcaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandFirewallProxyPolicySrcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandFirewallProxyPolicySrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok || d.HasChange("ssh_filter_profile") {
		t, err := expandFirewallProxyPolicySshFilterProfile(d, v, "ssh_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("ssh_policy_redirect"); ok || d.HasChange("ssh_policy_redirect") {
		t, err := expandFirewallProxyPolicySshPolicyRedirect(d, v, "ssh_policy_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok || d.HasChange("ssl_ssh_profile") {
		t, err := expandFirewallProxyPolicySslSshProfile(d, v, "ssl_ssh_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallProxyPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("telemetry_profile"); ok || d.HasChange("telemetry_profile") {
		t, err := expandFirewallProxyPolicyTelemetryProfile(d, v, "telemetry_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["telemetry-profile"] = t
		}
	}

	if v, ok := d.GetOk("transparent"); ok || d.HasChange("transparent") {
		t, err := expandFirewallProxyPolicyTransparent(d, v, "transparent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transparent"] = t
		}
	}

	if v, ok := d.GetOk("url_risk"); ok || d.HasChange("url_risk") {
		t, err := expandFirewallProxyPolicyUrlRisk(d, v, "url_risk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-risk"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandFirewallProxyPolicyUsers(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok || d.HasChange("utm_status") {
		t, err := expandFirewallProxyPolicyUtmStatus(d, v, "utm_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallProxyPolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("videofilter_profile"); ok || d.HasChange("videofilter_profile") {
		t, err := expandFirewallProxyPolicyVideofilterProfile(d, v, "videofilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["videofilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("voip_profile"); ok || d.HasChange("voip_profile") {
		t, err := expandFirewallProxyPolicyVoipProfile(d, v, "voip_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	}

	if v, ok := d.GetOk("virtual_patch_profile"); ok || d.HasChange("virtual_patch_profile") {
		t, err := expandFirewallProxyPolicyVirtualPatchProfile(d, v, "virtual_patch_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-patch-profile"] = t
		}
	}

	if v, ok := d.GetOk("waf_profile"); ok || d.HasChange("waf_profile") {
		t, err := expandFirewallProxyPolicyWafProfile(d, v, "waf_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["waf-profile"] = t
		}
	}

	if v, ok := d.GetOk("webcache"); ok || d.HasChange("webcache") {
		t, err := expandFirewallProxyPolicyWebcache(d, v, "webcache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache"] = t
		}
	}

	if v, ok := d.GetOk("webcache_https"); ok || d.HasChange("webcache_https") {
		t, err := expandFirewallProxyPolicyWebcacheHttps(d, v, "webcache_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache-https"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandFirewallProxyPolicyWebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_forward_server"); ok || d.HasChange("webproxy_forward_server") {
		t, err := expandFirewallProxyPolicyWebproxyForwardServer(d, v, "webproxy_forward_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-forward-server"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_profile"); ok || d.HasChange("webproxy_profile") {
		t, err := expandFirewallProxyPolicyWebproxyProfile(d, v, "webproxy_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-profile"] = t
		}
	}

	if v, ok := d.GetOk("ztna_ems_tag"); ok || d.HasChange("ztna_ems_tag") {
		t, err := expandFirewallProxyPolicyZtnaEmsTag(d, v, "ztna_ems_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-ems-tag"] = t
		}
	}

	if v, ok := d.GetOk("ztna_ems_tag_negate"); ok || d.HasChange("ztna_ems_tag_negate") {
		t, err := expandFirewallProxyPolicyZtnaEmsTagNegate(d, v, "ztna_ems_tag_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-ems-tag-negate"] = t
		}
	}

	if v, ok := d.GetOk("ztna_proxy"); ok || d.HasChange("ztna_proxy") {
		t, err := expandFirewallProxyPolicyZtnaProxy(d, v, "ztna_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-proxy"] = t
		}
	}

	if v, ok := d.GetOk("ztna_tags_match_logic"); ok || d.HasChange("ztna_tags_match_logic") {
		t, err := expandFirewallProxyPolicyZtnaTagsMatchLogic(d, v, "ztna_tags_match_logic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-tags-match-logic"] = t
		}
	}

	return &obj, nil
}
