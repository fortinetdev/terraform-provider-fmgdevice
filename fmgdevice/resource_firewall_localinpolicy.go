// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure user defined IPv4 local-in policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallLocalInPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallLocalInPolicyCreate,
		Read:   resourceFirewallLocalInPolicyRead,
		Update: resourceFirewallLocalInPolicyUpdate,
		Delete: resourceFirewallLocalInPolicyDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"ha_mgmt_intf_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"intf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"schedule": &schema.Schema{
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
			"virtual_patch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallLocalInPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallLocalInPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallLocalInPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("policyid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallLocalInPolicy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallLocalInPolicy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallLocalInPolicy resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallLocalInPolicy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallLocalInPolicy resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallLocalInPolicyRead(d, m)
}

func resourceFirewallLocalInPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallLocalInPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallLocalInPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallLocalInPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallLocalInPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallLocalInPolicyRead(d, m)
}

func resourceFirewallLocalInPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallLocalInPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallLocalInPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallLocalInPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallLocalInPolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallLocalInPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallLocalInPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallLocalInPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallLocalInPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicyDstaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicyHaMgmtIntfOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicyInternetServiceSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicyInternetServiceSrcCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicyInternetServiceSrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicyInternetServiceSrcFortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicyInternetServiceSrcGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicyInternetServiceSrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicyInternetServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicyIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicyLogtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicyServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicySrcaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicyVirtualPatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallLocalInPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenFirewallLocalInPolicyAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "FirewallLocalInPolicy-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallLocalInPolicyComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallLocalInPolicy-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenFirewallLocalInPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "FirewallLocalInPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenFirewallLocalInPolicyDstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr-negate"], "FirewallLocalInPolicy-DstaddrNegate"); ok {
			if err = d.Set("dstaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("ha_mgmt_intf_only", flattenFirewallLocalInPolicyHaMgmtIntfOnly(o["ha-mgmt-intf-only"], d, "ha_mgmt_intf_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-mgmt-intf-only"], "FirewallLocalInPolicy-HaMgmtIntfOnly"); ok {
			if err = d.Set("ha_mgmt_intf_only", vv); err != nil {
				return fmt.Errorf("Error reading ha_mgmt_intf_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_mgmt_intf_only: %v", err)
		}
	}

	if err = d.Set("internet_service_src", flattenFirewallLocalInPolicyInternetServiceSrc(o["internet-service-src"], d, "internet_service_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src"], "FirewallLocalInPolicy-InternetServiceSrc"); ok {
			if err = d.Set("internet_service_src", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom", flattenFirewallLocalInPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom"], "FirewallLocalInPolicy-InternetServiceSrcCustom"); ok {
			if err = d.Set("internet_service_src_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom_group", flattenFirewallLocalInPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom-group"], "FirewallLocalInPolicy-InternetServiceSrcCustomGroup"); ok {
			if err = d.Set("internet_service_src_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_fortiguard", flattenFirewallLocalInPolicyInternetServiceSrcFortiguard(o["internet-service-src-fortiguard"], d, "internet_service_src_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-fortiguard"], "FirewallLocalInPolicy-InternetServiceSrcFortiguard"); ok {
			if err = d.Set("internet_service_src_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service_src_group", flattenFirewallLocalInPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-group"], "FirewallLocalInPolicy-InternetServiceSrcGroup"); ok {
			if err = d.Set("internet_service_src_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_name", flattenFirewallLocalInPolicyInternetServiceSrcName(o["internet-service-src-name"], d, "internet_service_src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-name"], "FirewallLocalInPolicy-InternetServiceSrcName"); ok {
			if err = d.Set("internet_service_src_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_name: %v", err)
		}
	}

	if err = d.Set("internet_service_src_negate", flattenFirewallLocalInPolicyInternetServiceSrcNegate(o["internet-service-src-negate"], d, "internet_service_src_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-negate"], "FirewallLocalInPolicy-InternetServiceSrcNegate"); ok {
			if err = d.Set("internet_service_src_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
		}
	}

	if err = d.Set("intf", flattenFirewallLocalInPolicyIntf(o["intf"], d, "intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["intf"], "FirewallLocalInPolicy-Intf"); ok {
			if err = d.Set("intf", vv); err != nil {
				return fmt.Errorf("Error reading intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intf: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallLocalInPolicyLogtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "FirewallLocalInPolicy-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("policyid", flattenFirewallLocalInPolicyPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "FirewallLocalInPolicy-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("schedule", flattenFirewallLocalInPolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "FirewallLocalInPolicy-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallLocalInPolicyService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "FirewallLocalInPolicy-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenFirewallLocalInPolicyServiceNegate(o["service-negate"], d, "service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-negate"], "FirewallLocalInPolicy-ServiceNegate"); ok {
			if err = d.Set("service_negate", vv); err != nil {
				return fmt.Errorf("Error reading service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenFirewallLocalInPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "FirewallLocalInPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenFirewallLocalInPolicySrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr-negate"], "FirewallLocalInPolicy-SrcaddrNegate"); ok {
			if err = d.Set("srcaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallLocalInPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallLocalInPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallLocalInPolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallLocalInPolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("virtual_patch", flattenFirewallLocalInPolicyVirtualPatch(o["virtual-patch"], d, "virtual_patch")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-patch"], "FirewallLocalInPolicy-VirtualPatch"); ok {
			if err = d.Set("virtual_patch", vv); err != nil {
				return fmt.Errorf("Error reading virtual_patch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_patch: %v", err)
		}
	}

	return nil
}

func flattenFirewallLocalInPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallLocalInPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicyDstaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicyHaMgmtIntfOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicyInternetServiceSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicyInternetServiceSrcFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicyInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicyInternetServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicyIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicyLogtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicySchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicyServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicySrcaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicyVirtualPatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallLocalInPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandFirewallLocalInPolicyAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallLocalInPolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallLocalInPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok || d.HasChange("dstaddr_negate") {
		t, err := expandFirewallLocalInPolicyDstaddrNegate(d, v, "dstaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("ha_mgmt_intf_only"); ok || d.HasChange("ha_mgmt_intf_only") {
		t, err := expandFirewallLocalInPolicyHaMgmtIntfOnly(d, v, "ha_mgmt_intf_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-mgmt-intf-only"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src"); ok || d.HasChange("internet_service_src") {
		t, err := expandFirewallLocalInPolicyInternetServiceSrc(d, v, "internet_service_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom"); ok || d.HasChange("internet_service_src_custom") {
		t, err := expandFirewallLocalInPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom_group"); ok || d.HasChange("internet_service_src_custom_group") {
		t, err := expandFirewallLocalInPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_fortiguard"); ok || d.HasChange("internet_service_src_fortiguard") {
		t, err := expandFirewallLocalInPolicyInternetServiceSrcFortiguard(d, v, "internet_service_src_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_group"); ok || d.HasChange("internet_service_src_group") {
		t, err := expandFirewallLocalInPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_name"); ok || d.HasChange("internet_service_src_name") {
		t, err := expandFirewallLocalInPolicyInternetServiceSrcName(d, v, "internet_service_src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_negate"); ok || d.HasChange("internet_service_src_negate") {
		t, err := expandFirewallLocalInPolicyInternetServiceSrcNegate(d, v, "internet_service_src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("intf"); ok || d.HasChange("intf") {
		t, err := expandFirewallLocalInPolicyIntf(d, v, "intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok || d.HasChange("logtraffic") {
		t, err := expandFirewallLocalInPolicyLogtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandFirewallLocalInPolicyPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandFirewallLocalInPolicySchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallLocalInPolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok || d.HasChange("service_negate") {
		t, err := expandFirewallLocalInPolicyServiceNegate(d, v, "service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallLocalInPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok || d.HasChange("srcaddr_negate") {
		t, err := expandFirewallLocalInPolicySrcaddrNegate(d, v, "srcaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallLocalInPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallLocalInPolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("virtual_patch"); ok || d.HasChange("virtual_patch") {
		t, err := expandFirewallLocalInPolicyVirtualPatch(d, v, "virtual_patch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-patch"] = t
		}
	}

	return &obj, nil
}
