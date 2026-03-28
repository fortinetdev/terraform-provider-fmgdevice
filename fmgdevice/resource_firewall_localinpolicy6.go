// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure user defined IPv6 local-in policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallLocalInPolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallLocalInPolicy6Create,
		Read:   resourceFirewallLocalInPolicy6Read,
		Update: resourceFirewallLocalInPolicy6Update,
		Delete: resourceFirewallLocalInPolicy6Delete,

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
			"internet_service6_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceFirewallLocalInPolicy6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallLocalInPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallLocalInPolicy6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("policyid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallLocalInPolicy6(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallLocalInPolicy6(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallLocalInPolicy6 resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallLocalInPolicy6(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallLocalInPolicy6 resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallLocalInPolicy6Read(d, m)
}

func resourceFirewallLocalInPolicy6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallLocalInPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallLocalInPolicy6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallLocalInPolicy6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallLocalInPolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallLocalInPolicy6Read(d, m)
}

func resourceFirewallLocalInPolicy6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallLocalInPolicy6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallLocalInPolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallLocalInPolicy6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallLocalInPolicy6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallLocalInPolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallLocalInPolicy6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallLocalInPolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallLocalInPolicy6Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicy6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicy6Dstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicy6DstaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicy6InternetService6Src(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicy6InternetService6SrcCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicy6InternetService6SrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicy6InternetService6SrcFortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicy6InternetService6SrcGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicy6InternetService6SrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicy6InternetService6SrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicy6Intf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicy6Logtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicy6Policyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicy6Schedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicy6Service(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicy6ServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicy6Srcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallLocalInPolicy6SrcaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicy6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicy6Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLocalInPolicy6VirtualPatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallLocalInPolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenFirewallLocalInPolicy6Action(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "FirewallLocalInPolicy6-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallLocalInPolicy6Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallLocalInPolicy6-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenFirewallLocalInPolicy6Dstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "FirewallLocalInPolicy6-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenFirewallLocalInPolicy6DstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr-negate"], "FirewallLocalInPolicy6-DstaddrNegate"); ok {
			if err = d.Set("dstaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("internet_service6_src", flattenFirewallLocalInPolicy6InternetService6Src(o["internet-service6-src"], d, "internet_service6_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src"], "FirewallLocalInPolicy6-InternetService6Src"); ok {
			if err = d.Set("internet_service6_src", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_custom", flattenFirewallLocalInPolicy6InternetService6SrcCustom(o["internet-service6-src-custom"], d, "internet_service6_src_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-custom"], "FirewallLocalInPolicy6-InternetService6SrcCustom"); ok {
			if err = d.Set("internet_service6_src_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_custom: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_custom_group", flattenFirewallLocalInPolicy6InternetService6SrcCustomGroup(o["internet-service6-src-custom-group"], d, "internet_service6_src_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-custom-group"], "FirewallLocalInPolicy6-InternetService6SrcCustomGroup"); ok {
			if err = d.Set("internet_service6_src_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_fortiguard", flattenFirewallLocalInPolicy6InternetService6SrcFortiguard(o["internet-service6-src-fortiguard"], d, "internet_service6_src_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-fortiguard"], "FirewallLocalInPolicy6-InternetService6SrcFortiguard"); ok {
			if err = d.Set("internet_service6_src_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_group", flattenFirewallLocalInPolicy6InternetService6SrcGroup(o["internet-service6-src-group"], d, "internet_service6_src_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-group"], "FirewallLocalInPolicy6-InternetService6SrcGroup"); ok {
			if err = d.Set("internet_service6_src_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_group: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_name", flattenFirewallLocalInPolicy6InternetService6SrcName(o["internet-service6-src-name"], d, "internet_service6_src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-name"], "FirewallLocalInPolicy6-InternetService6SrcName"); ok {
			if err = d.Set("internet_service6_src_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_name: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_negate", flattenFirewallLocalInPolicy6InternetService6SrcNegate(o["internet-service6-src-negate"], d, "internet_service6_src_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service6-src-negate"], "FirewallLocalInPolicy6-InternetService6SrcNegate"); ok {
			if err = d.Set("internet_service6_src_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service6_src_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service6_src_negate: %v", err)
		}
	}

	if err = d.Set("intf", flattenFirewallLocalInPolicy6Intf(o["intf"], d, "intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["intf"], "FirewallLocalInPolicy6-Intf"); ok {
			if err = d.Set("intf", vv); err != nil {
				return fmt.Errorf("Error reading intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intf: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallLocalInPolicy6Logtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "FirewallLocalInPolicy6-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("policyid", flattenFirewallLocalInPolicy6Policyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "FirewallLocalInPolicy6-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("schedule", flattenFirewallLocalInPolicy6Schedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "FirewallLocalInPolicy6-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallLocalInPolicy6Service(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "FirewallLocalInPolicy6-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenFirewallLocalInPolicy6ServiceNegate(o["service-negate"], d, "service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-negate"], "FirewallLocalInPolicy6-ServiceNegate"); ok {
			if err = d.Set("service_negate", vv); err != nil {
				return fmt.Errorf("Error reading service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenFirewallLocalInPolicy6Srcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "FirewallLocalInPolicy6-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenFirewallLocalInPolicy6SrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr-negate"], "FirewallLocalInPolicy6-SrcaddrNegate"); ok {
			if err = d.Set("srcaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallLocalInPolicy6Status(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallLocalInPolicy6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallLocalInPolicy6Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallLocalInPolicy6-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("virtual_patch", flattenFirewallLocalInPolicy6VirtualPatch(o["virtual-patch"], d, "virtual_patch")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-patch"], "FirewallLocalInPolicy6-VirtualPatch"); ok {
			if err = d.Set("virtual_patch", vv); err != nil {
				return fmt.Errorf("Error reading virtual_patch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_patch: %v", err)
		}
	}

	return nil
}

func flattenFirewallLocalInPolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallLocalInPolicy6Action(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicy6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicy6Dstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicy6DstaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicy6InternetService6Src(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicy6InternetService6SrcCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicy6InternetService6SrcCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicy6InternetService6SrcFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicy6InternetService6SrcGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicy6InternetService6SrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicy6InternetService6SrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicy6Intf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicy6Logtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicy6Policyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicy6Schedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicy6Service(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicy6ServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicy6Srcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallLocalInPolicy6SrcaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicy6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicy6Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLocalInPolicy6VirtualPatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallLocalInPolicy6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandFirewallLocalInPolicy6Action(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallLocalInPolicy6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallLocalInPolicy6Dstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok || d.HasChange("dstaddr_negate") {
		t, err := expandFirewallLocalInPolicy6DstaddrNegate(d, v, "dstaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src"); ok || d.HasChange("internet_service6_src") {
		t, err := expandFirewallLocalInPolicy6InternetService6Src(d, v, "internet_service6_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_custom"); ok || d.HasChange("internet_service6_src_custom") {
		t, err := expandFirewallLocalInPolicy6InternetService6SrcCustom(d, v, "internet_service6_src_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_custom_group"); ok || d.HasChange("internet_service6_src_custom_group") {
		t, err := expandFirewallLocalInPolicy6InternetService6SrcCustomGroup(d, v, "internet_service6_src_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_fortiguard"); ok || d.HasChange("internet_service6_src_fortiguard") {
		t, err := expandFirewallLocalInPolicy6InternetService6SrcFortiguard(d, v, "internet_service6_src_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_group"); ok || d.HasChange("internet_service6_src_group") {
		t, err := expandFirewallLocalInPolicy6InternetService6SrcGroup(d, v, "internet_service6_src_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_name"); ok || d.HasChange("internet_service6_src_name") {
		t, err := expandFirewallLocalInPolicy6InternetService6SrcName(d, v, "internet_service6_src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_negate"); ok || d.HasChange("internet_service6_src_negate") {
		t, err := expandFirewallLocalInPolicy6InternetService6SrcNegate(d, v, "internet_service6_src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("intf"); ok || d.HasChange("intf") {
		t, err := expandFirewallLocalInPolicy6Intf(d, v, "intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok || d.HasChange("logtraffic") {
		t, err := expandFirewallLocalInPolicy6Logtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandFirewallLocalInPolicy6Policyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandFirewallLocalInPolicy6Schedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallLocalInPolicy6Service(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok || d.HasChange("service_negate") {
		t, err := expandFirewallLocalInPolicy6ServiceNegate(d, v, "service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallLocalInPolicy6Srcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok || d.HasChange("srcaddr_negate") {
		t, err := expandFirewallLocalInPolicy6SrcaddrNegate(d, v, "srcaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallLocalInPolicy6Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallLocalInPolicy6Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("virtual_patch"); ok || d.HasChange("virtual_patch") {
		t, err := expandFirewallLocalInPolicy6VirtualPatch(d, v, "virtual_patch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-patch"] = t
		}
	}

	return &obj, nil
}
