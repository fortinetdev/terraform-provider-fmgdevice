// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv4 interface policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInterfacePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInterfacePolicyCreate,
		Read:   resourceFirewallInterfacePolicyRead,
		Update: resourceFirewallInterfacePolicyUpdate,
		Delete: resourceFirewallInterfacePolicyDelete,

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
			"address_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"application_list_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"av_profile_status": &schema.Schema{
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
			"casb_profile_status": &schema.Schema{
				Type:     schema.TypeString,
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
			"dlp_sensor_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dlp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dlp_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dsri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
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
			"emailfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ips_sensor_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
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
			"uuid": &schema.Schema{
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
			"webfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallInterfacePolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInterfacePolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInterfacePolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("policyid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallInterfacePolicy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallInterfacePolicy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallInterfacePolicy resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallInterfacePolicy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallInterfacePolicy resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallInterfacePolicyRead(d, m)
}

func resourceFirewallInterfacePolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInterfacePolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInterfacePolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallInterfacePolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInterfacePolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallInterfacePolicyRead(d, m)
}

func resourceFirewallInterfacePolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallInterfacePolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInterfacePolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInterfacePolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallInterfacePolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallInterfacePolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInterfacePolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInterfacePolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInterfacePolicyAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicyApplicationListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyAvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicyAvProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyCasbProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicyCasbProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyDlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicyDlpSensorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyDlpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicyDlpProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyDsri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicyEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicyEmailfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicyIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicyIpsSensorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyLogtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicyWebfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallInterfacePolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address_type", flattenFirewallInterfacePolicyAddressType(o["address-type"], d, "address_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-type"], "FirewallInterfacePolicy-AddressType"); ok {
			if err = d.Set("address_type", vv); err != nil {
				return fmt.Errorf("Error reading address_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_type: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallInterfacePolicyApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "FirewallInterfacePolicy-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("application_list_status", flattenFirewallInterfacePolicyApplicationListStatus(o["application-list-status"], d, "application_list_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list-status"], "FirewallInterfacePolicy-ApplicationListStatus"); ok {
			if err = d.Set("application_list_status", vv); err != nil {
				return fmt.Errorf("Error reading application_list_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list_status: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallInterfacePolicyAvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile"], "FirewallInterfacePolicy-AvProfile"); ok {
			if err = d.Set("av_profile", vv); err != nil {
				return fmt.Errorf("Error reading av_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("av_profile_status", flattenFirewallInterfacePolicyAvProfileStatus(o["av-profile-status"], d, "av_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile-status"], "FirewallInterfacePolicy-AvProfileStatus"); ok {
			if err = d.Set("av_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading av_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile_status: %v", err)
		}
	}

	if err = d.Set("casb_profile", flattenFirewallInterfacePolicyCasbProfile(o["casb-profile"], d, "casb_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["casb-profile"], "FirewallInterfacePolicy-CasbProfile"); ok {
			if err = d.Set("casb_profile", vv); err != nil {
				return fmt.Errorf("Error reading casb_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading casb_profile: %v", err)
		}
	}

	if err = d.Set("casb_profile_status", flattenFirewallInterfacePolicyCasbProfileStatus(o["casb-profile-status"], d, "casb_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["casb-profile-status"], "FirewallInterfacePolicy-CasbProfileStatus"); ok {
			if err = d.Set("casb_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading casb_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading casb_profile_status: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallInterfacePolicyComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallInterfacePolicy-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallInterfacePolicyDlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor"], "FirewallInterfacePolicy-DlpSensor"); ok {
			if err = d.Set("dlp_sensor", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("dlp_sensor_status", flattenFirewallInterfacePolicyDlpSensorStatus(o["dlp-sensor-status"], d, "dlp_sensor_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor-status"], "FirewallInterfacePolicy-DlpSensorStatus"); ok {
			if err = d.Set("dlp_sensor_status", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor_status: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenFirewallInterfacePolicyDlpProfile(o["dlp-profile"], d, "dlp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-profile"], "FirewallInterfacePolicy-DlpProfile"); ok {
			if err = d.Set("dlp_profile", vv); err != nil {
				return fmt.Errorf("Error reading dlp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("dlp_profile_status", flattenFirewallInterfacePolicyDlpProfileStatus(o["dlp-profile-status"], d, "dlp_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-profile-status"], "FirewallInterfacePolicy-DlpProfileStatus"); ok {
			if err = d.Set("dlp_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading dlp_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_profile_status: %v", err)
		}
	}

	if err = d.Set("dsri", flattenFirewallInterfacePolicyDsri(o["dsri"], d, "dsri")); err != nil {
		if vv, ok := fortiAPIPatch(o["dsri"], "FirewallInterfacePolicy-Dsri"); ok {
			if err = d.Set("dsri", vv); err != nil {
				return fmt.Errorf("Error reading dsri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenFirewallInterfacePolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "FirewallInterfacePolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallInterfacePolicyEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile"], "FirewallInterfacePolicy-EmailfilterProfile"); ok {
			if err = d.Set("emailfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile_status", flattenFirewallInterfacePolicyEmailfilterProfileStatus(o["emailfilter-profile-status"], d, "emailfilter_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile-status"], "FirewallInterfacePolicy-EmailfilterProfileStatus"); ok {
			if err = d.Set("emailfilter_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile_status: %v", err)
		}
	}

	if err = d.Set("interface", flattenFirewallInterfacePolicyInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "FirewallInterfacePolicy-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallInterfacePolicyIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "FirewallInterfacePolicy-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("ips_sensor_status", flattenFirewallInterfacePolicyIpsSensorStatus(o["ips-sensor-status"], d, "ips_sensor_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor-status"], "FirewallInterfacePolicy-IpsSensorStatus"); ok {
			if err = d.Set("ips_sensor_status", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor_status: %v", err)
		}
	}

	if err = d.Set("label", flattenFirewallInterfacePolicyLabel(o["label"], d, "label")); err != nil {
		if vv, ok := fortiAPIPatch(o["label"], "FirewallInterfacePolicy-Label"); ok {
			if err = d.Set("label", vv); err != nil {
				return fmt.Errorf("Error reading label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallInterfacePolicyLogtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "FirewallInterfacePolicy-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("policyid", flattenFirewallInterfacePolicyPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "FirewallInterfacePolicy-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallInterfacePolicyService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "FirewallInterfacePolicy-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenFirewallInterfacePolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "FirewallInterfacePolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallInterfacePolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallInterfacePolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallInterfacePolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallInterfacePolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallInterfacePolicyWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "FirewallInterfacePolicy-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile_status", flattenFirewallInterfacePolicyWebfilterProfileStatus(o["webfilter-profile-status"], d, "webfilter_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile-status"], "FirewallInterfacePolicy-WebfilterProfileStatus"); ok {
			if err = d.Set("webfilter_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile_status: %v", err)
		}
	}

	return nil
}

func flattenFirewallInterfacePolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInterfacePolicyAddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicyApplicationListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyAvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicyAvProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyCasbProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicyCasbProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyDlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicyDlpSensorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyDlpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicyDlpProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyDsri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicyEmailfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicyEmailfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicyIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicyIpsSensorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyLogtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicyWebfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInterfacePolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address_type"); ok || d.HasChange("address_type") {
		t, err := expandFirewallInterfacePolicyAddressType(d, v, "address_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-type"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandFirewallInterfacePolicyApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("application_list_status"); ok || d.HasChange("application_list_status") {
		t, err := expandFirewallInterfacePolicyApplicationListStatus(d, v, "application_list_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list-status"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok || d.HasChange("av_profile") {
		t, err := expandFirewallInterfacePolicyAvProfile(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("av_profile_status"); ok || d.HasChange("av_profile_status") {
		t, err := expandFirewallInterfacePolicyAvProfileStatus(d, v, "av_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("casb_profile"); ok || d.HasChange("casb_profile") {
		t, err := expandFirewallInterfacePolicyCasbProfile(d, v, "casb_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile"] = t
		}
	}

	if v, ok := d.GetOk("casb_profile_status"); ok || d.HasChange("casb_profile_status") {
		t, err := expandFirewallInterfacePolicyCasbProfileStatus(d, v, "casb_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallInterfacePolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok || d.HasChange("dlp_sensor") {
		t, err := expandFirewallInterfacePolicyDlpSensor(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor_status"); ok || d.HasChange("dlp_sensor_status") {
		t, err := expandFirewallInterfacePolicyDlpSensorStatus(d, v, "dlp_sensor_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor-status"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile"); ok || d.HasChange("dlp_profile") {
		t, err := expandFirewallInterfacePolicyDlpProfile(d, v, "dlp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile_status"); ok || d.HasChange("dlp_profile_status") {
		t, err := expandFirewallInterfacePolicyDlpProfileStatus(d, v, "dlp_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("dsri"); ok || d.HasChange("dsri") {
		t, err := expandFirewallInterfacePolicyDsri(d, v, "dsri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsri"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallInterfacePolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok || d.HasChange("emailfilter_profile") {
		t, err := expandFirewallInterfacePolicyEmailfilterProfile(d, v, "emailfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile_status"); ok || d.HasChange("emailfilter_profile_status") {
		t, err := expandFirewallInterfacePolicyEmailfilterProfileStatus(d, v, "emailfilter_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandFirewallInterfacePolicyInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandFirewallInterfacePolicyIpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor_status"); ok || d.HasChange("ips_sensor_status") {
		t, err := expandFirewallInterfacePolicyIpsSensorStatus(d, v, "ips_sensor_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor-status"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok || d.HasChange("label") {
		t, err := expandFirewallInterfacePolicyLabel(d, v, "label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok || d.HasChange("logtraffic") {
		t, err := expandFirewallInterfacePolicyLogtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandFirewallInterfacePolicyPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallInterfacePolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallInterfacePolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallInterfacePolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallInterfacePolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandFirewallInterfacePolicyWebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile_status"); ok || d.HasChange("webfilter_profile_status") {
		t, err := expandFirewallInterfacePolicyWebfilterProfileStatus(d, v, "webfilter_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile-status"] = t
		}
	}

	return &obj, nil
}
