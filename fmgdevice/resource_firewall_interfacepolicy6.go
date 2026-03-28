// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv6 interface policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInterfacePolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInterfacePolicy6Create,
		Read:   resourceFirewallInterfacePolicy6Read,
		Update: resourceFirewallInterfacePolicy6Update,
		Delete: resourceFirewallInterfacePolicy6Delete,

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
			"dstaddr6": &schema.Schema{
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
			"service6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
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

func resourceFirewallInterfacePolicy6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInterfacePolicy6(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInterfacePolicy6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("policyid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallInterfacePolicy6(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallInterfacePolicy6(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallInterfacePolicy6 resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallInterfacePolicy6(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallInterfacePolicy6 resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallInterfacePolicy6Read(d, m)
}

func resourceFirewallInterfacePolicy6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInterfacePolicy6(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInterfacePolicy6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallInterfacePolicy6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInterfacePolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallInterfacePolicy6Read(d, m)
}

func resourceFirewallInterfacePolicy6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallInterfacePolicy6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInterfacePolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInterfacePolicy6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallInterfacePolicy6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallInterfacePolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInterfacePolicy6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInterfacePolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInterfacePolicy6AddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6ApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicy6ApplicationListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6AvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicy6AvProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6CasbProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicy6CasbProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6DlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicy6DlpSensorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6DlpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicy6DlpProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Dsri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Dstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicy6EmailfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicy6EmailfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicy6IpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicy6IpsSensorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Label(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Logtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Policyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Service6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicy6Srcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicy6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6WebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInterfacePolicy6WebfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallInterfacePolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address_type", flattenFirewallInterfacePolicy6AddressType(o["address-type"], d, "address_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-type"], "FirewallInterfacePolicy6-AddressType"); ok {
			if err = d.Set("address_type", vv); err != nil {
				return fmt.Errorf("Error reading address_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_type: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallInterfacePolicy6ApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "FirewallInterfacePolicy6-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("application_list_status", flattenFirewallInterfacePolicy6ApplicationListStatus(o["application-list-status"], d, "application_list_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list-status"], "FirewallInterfacePolicy6-ApplicationListStatus"); ok {
			if err = d.Set("application_list_status", vv); err != nil {
				return fmt.Errorf("Error reading application_list_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list_status: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallInterfacePolicy6AvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile"], "FirewallInterfacePolicy6-AvProfile"); ok {
			if err = d.Set("av_profile", vv); err != nil {
				return fmt.Errorf("Error reading av_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("av_profile_status", flattenFirewallInterfacePolicy6AvProfileStatus(o["av-profile-status"], d, "av_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile-status"], "FirewallInterfacePolicy6-AvProfileStatus"); ok {
			if err = d.Set("av_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading av_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile_status: %v", err)
		}
	}

	if err = d.Set("casb_profile", flattenFirewallInterfacePolicy6CasbProfile(o["casb-profile"], d, "casb_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["casb-profile"], "FirewallInterfacePolicy6-CasbProfile"); ok {
			if err = d.Set("casb_profile", vv); err != nil {
				return fmt.Errorf("Error reading casb_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading casb_profile: %v", err)
		}
	}

	if err = d.Set("casb_profile_status", flattenFirewallInterfacePolicy6CasbProfileStatus(o["casb-profile-status"], d, "casb_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["casb-profile-status"], "FirewallInterfacePolicy6-CasbProfileStatus"); ok {
			if err = d.Set("casb_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading casb_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading casb_profile_status: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallInterfacePolicy6Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallInterfacePolicy6-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallInterfacePolicy6DlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor"], "FirewallInterfacePolicy6-DlpSensor"); ok {
			if err = d.Set("dlp_sensor", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("dlp_sensor_status", flattenFirewallInterfacePolicy6DlpSensorStatus(o["dlp-sensor-status"], d, "dlp_sensor_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor-status"], "FirewallInterfacePolicy6-DlpSensorStatus"); ok {
			if err = d.Set("dlp_sensor_status", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor_status: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenFirewallInterfacePolicy6DlpProfile(o["dlp-profile"], d, "dlp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-profile"], "FirewallInterfacePolicy6-DlpProfile"); ok {
			if err = d.Set("dlp_profile", vv); err != nil {
				return fmt.Errorf("Error reading dlp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("dlp_profile_status", flattenFirewallInterfacePolicy6DlpProfileStatus(o["dlp-profile-status"], d, "dlp_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-profile-status"], "FirewallInterfacePolicy6-DlpProfileStatus"); ok {
			if err = d.Set("dlp_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading dlp_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_profile_status: %v", err)
		}
	}

	if err = d.Set("dsri", flattenFirewallInterfacePolicy6Dsri(o["dsri"], d, "dsri")); err != nil {
		if vv, ok := fortiAPIPatch(o["dsri"], "FirewallInterfacePolicy6-Dsri"); ok {
			if err = d.Set("dsri", vv); err != nil {
				return fmt.Errorf("Error reading dsri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenFirewallInterfacePolicy6Dstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "FirewallInterfacePolicy6-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallInterfacePolicy6EmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile"], "FirewallInterfacePolicy6-EmailfilterProfile"); ok {
			if err = d.Set("emailfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile_status", flattenFirewallInterfacePolicy6EmailfilterProfileStatus(o["emailfilter-profile-status"], d, "emailfilter_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile-status"], "FirewallInterfacePolicy6-EmailfilterProfileStatus"); ok {
			if err = d.Set("emailfilter_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile_status: %v", err)
		}
	}

	if err = d.Set("interface", flattenFirewallInterfacePolicy6Interface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "FirewallInterfacePolicy6-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallInterfacePolicy6IpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "FirewallInterfacePolicy6-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("ips_sensor_status", flattenFirewallInterfacePolicy6IpsSensorStatus(o["ips-sensor-status"], d, "ips_sensor_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor-status"], "FirewallInterfacePolicy6-IpsSensorStatus"); ok {
			if err = d.Set("ips_sensor_status", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor_status: %v", err)
		}
	}

	if err = d.Set("label", flattenFirewallInterfacePolicy6Label(o["label"], d, "label")); err != nil {
		if vv, ok := fortiAPIPatch(o["label"], "FirewallInterfacePolicy6-Label"); ok {
			if err = d.Set("label", vv); err != nil {
				return fmt.Errorf("Error reading label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallInterfacePolicy6Logtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "FirewallInterfacePolicy6-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("policyid", flattenFirewallInterfacePolicy6Policyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "FirewallInterfacePolicy6-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("service6", flattenFirewallInterfacePolicy6Service6(o["service6"], d, "service6")); err != nil {
		if vv, ok := fortiAPIPatch(o["service6"], "FirewallInterfacePolicy6-Service6"); ok {
			if err = d.Set("service6", vv); err != nil {
				return fmt.Errorf("Error reading service6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service6: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenFirewallInterfacePolicy6Srcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "FirewallInterfacePolicy6-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallInterfacePolicy6Status(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallInterfacePolicy6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallInterfacePolicy6Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallInterfacePolicy6-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallInterfacePolicy6WebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "FirewallInterfacePolicy6-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile_status", flattenFirewallInterfacePolicy6WebfilterProfileStatus(o["webfilter-profile-status"], d, "webfilter_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile-status"], "FirewallInterfacePolicy6-WebfilterProfileStatus"); ok {
			if err = d.Set("webfilter_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile_status: %v", err)
		}
	}

	return nil
}

func flattenFirewallInterfacePolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInterfacePolicy6AddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6ApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicy6ApplicationListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6AvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicy6AvProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6CasbProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicy6CasbProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6DlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicy6DlpSensorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6DlpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicy6DlpProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Dsri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Dstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicy6EmailfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicy6EmailfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicy6IpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicy6IpsSensorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Label(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Logtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Policyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Service6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicy6Srcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicy6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6WebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInterfacePolicy6WebfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInterfacePolicy6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address_type"); ok || d.HasChange("address_type") {
		t, err := expandFirewallInterfacePolicy6AddressType(d, v, "address_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-type"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandFirewallInterfacePolicy6ApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("application_list_status"); ok || d.HasChange("application_list_status") {
		t, err := expandFirewallInterfacePolicy6ApplicationListStatus(d, v, "application_list_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list-status"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok || d.HasChange("av_profile") {
		t, err := expandFirewallInterfacePolicy6AvProfile(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("av_profile_status"); ok || d.HasChange("av_profile_status") {
		t, err := expandFirewallInterfacePolicy6AvProfileStatus(d, v, "av_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("casb_profile"); ok || d.HasChange("casb_profile") {
		t, err := expandFirewallInterfacePolicy6CasbProfile(d, v, "casb_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile"] = t
		}
	}

	if v, ok := d.GetOk("casb_profile_status"); ok || d.HasChange("casb_profile_status") {
		t, err := expandFirewallInterfacePolicy6CasbProfileStatus(d, v, "casb_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallInterfacePolicy6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok || d.HasChange("dlp_sensor") {
		t, err := expandFirewallInterfacePolicy6DlpSensor(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor_status"); ok || d.HasChange("dlp_sensor_status") {
		t, err := expandFirewallInterfacePolicy6DlpSensorStatus(d, v, "dlp_sensor_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor-status"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile"); ok || d.HasChange("dlp_profile") {
		t, err := expandFirewallInterfacePolicy6DlpProfile(d, v, "dlp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile_status"); ok || d.HasChange("dlp_profile_status") {
		t, err := expandFirewallInterfacePolicy6DlpProfileStatus(d, v, "dlp_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("dsri"); ok || d.HasChange("dsri") {
		t, err := expandFirewallInterfacePolicy6Dsri(d, v, "dsri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsri"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandFirewallInterfacePolicy6Dstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok || d.HasChange("emailfilter_profile") {
		t, err := expandFirewallInterfacePolicy6EmailfilterProfile(d, v, "emailfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile_status"); ok || d.HasChange("emailfilter_profile_status") {
		t, err := expandFirewallInterfacePolicy6EmailfilterProfileStatus(d, v, "emailfilter_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandFirewallInterfacePolicy6Interface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandFirewallInterfacePolicy6IpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor_status"); ok || d.HasChange("ips_sensor_status") {
		t, err := expandFirewallInterfacePolicy6IpsSensorStatus(d, v, "ips_sensor_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor-status"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok || d.HasChange("label") {
		t, err := expandFirewallInterfacePolicy6Label(d, v, "label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok || d.HasChange("logtraffic") {
		t, err := expandFirewallInterfacePolicy6Logtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandFirewallInterfacePolicy6Policyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("service6"); ok || d.HasChange("service6") {
		t, err := expandFirewallInterfacePolicy6Service6(d, v, "service6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service6"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandFirewallInterfacePolicy6Srcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallInterfacePolicy6Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallInterfacePolicy6Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandFirewallInterfacePolicy6WebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile_status"); ok || d.HasChange("webfilter_profile_status") {
		t, err := expandFirewallInterfacePolicy6WebfilterProfileStatus(d, v, "webfilter_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile-status"] = t
		}
	}

	return &obj, nil
}
