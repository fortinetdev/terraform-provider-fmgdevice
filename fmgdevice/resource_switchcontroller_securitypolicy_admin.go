// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure fortiswitch's admin security-policy.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSecurityPolicyAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSecurityPolicyAdminCreate,
		Read:   resourceSwitchControllerSecurityPolicyAdminRead,
		Update: resourceSwitchControllerSecurityPolicyAdminUpdate,
		Delete: resourceSwitchControllerSecurityPolicyAdminDelete,

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
			"auto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6_trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost10": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost3": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost4": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost5": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost6": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost7": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost8": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost9": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerSecurityPolicyAdminCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerSecurityPolicyAdmin(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSecurityPolicyAdmin resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSwitchControllerSecurityPolicyAdmin(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSwitchControllerSecurityPolicyAdmin(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SwitchControllerSecurityPolicyAdmin resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSwitchControllerSecurityPolicyAdmin(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SwitchControllerSecurityPolicyAdmin resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerSecurityPolicyAdminRead(d, m)
}

func resourceSwitchControllerSecurityPolicyAdminUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerSecurityPolicyAdmin(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicyAdmin resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSwitchControllerSecurityPolicyAdmin(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicyAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerSecurityPolicyAdminRead(d, m)
}

func resourceSwitchControllerSecurityPolicyAdminDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerSecurityPolicyAdmin(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSecurityPolicyAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSecurityPolicyAdminRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerSecurityPolicyAdmin(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicyAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSecurityPolicyAdmin(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicyAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSecurityPolicyAdminAuto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerSecurityPolicyAdmin(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auto", flattenSwitchControllerSecurityPolicyAdminAuto(o["auto"], d, "auto")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto"], "SwitchControllerSecurityPolicyAdmin-Auto"); ok {
			if err = d.Set("auto", vv); err != nil {
				return fmt.Errorf("Error reading auto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost1", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost1(o["ip6-trusthost1"], d, "ip6_trusthost1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost1"], "SwitchControllerSecurityPolicyAdmin-Ip6Trusthost1"); ok {
			if err = d.Set("ip6_trusthost1", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost1: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost10", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost10(o["ip6-trusthost10"], d, "ip6_trusthost10")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost10"], "SwitchControllerSecurityPolicyAdmin-Ip6Trusthost10"); ok {
			if err = d.Set("ip6_trusthost10", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost10: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost2", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost2(o["ip6-trusthost2"], d, "ip6_trusthost2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost2"], "SwitchControllerSecurityPolicyAdmin-Ip6Trusthost2"); ok {
			if err = d.Set("ip6_trusthost2", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost2: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost3", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost3(o["ip6-trusthost3"], d, "ip6_trusthost3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost3"], "SwitchControllerSecurityPolicyAdmin-Ip6Trusthost3"); ok {
			if err = d.Set("ip6_trusthost3", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost3: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost4", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost4(o["ip6-trusthost4"], d, "ip6_trusthost4")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost4"], "SwitchControllerSecurityPolicyAdmin-Ip6Trusthost4"); ok {
			if err = d.Set("ip6_trusthost4", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost4: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost5", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost5(o["ip6-trusthost5"], d, "ip6_trusthost5")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost5"], "SwitchControllerSecurityPolicyAdmin-Ip6Trusthost5"); ok {
			if err = d.Set("ip6_trusthost5", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost5: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost6", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost6(o["ip6-trusthost6"], d, "ip6_trusthost6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost6"], "SwitchControllerSecurityPolicyAdmin-Ip6Trusthost6"); ok {
			if err = d.Set("ip6_trusthost6", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost6: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost7", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost7(o["ip6-trusthost7"], d, "ip6_trusthost7")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost7"], "SwitchControllerSecurityPolicyAdmin-Ip6Trusthost7"); ok {
			if err = d.Set("ip6_trusthost7", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost7: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost8", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost8(o["ip6-trusthost8"], d, "ip6_trusthost8")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost8"], "SwitchControllerSecurityPolicyAdmin-Ip6Trusthost8"); ok {
			if err = d.Set("ip6_trusthost8", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost8: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost9", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost9(o["ip6-trusthost9"], d, "ip6_trusthost9")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost9"], "SwitchControllerSecurityPolicyAdmin-Ip6Trusthost9"); ok {
			if err = d.Set("ip6_trusthost9", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost9: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerSecurityPolicyAdminName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerSecurityPolicyAdmin-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("trusthost1", flattenSwitchControllerSecurityPolicyAdminTrusthost1(o["trusthost1"], d, "trusthost1")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost1"], "SwitchControllerSecurityPolicyAdmin-Trusthost1"); ok {
			if err = d.Set("trusthost1", vv); err != nil {
				return fmt.Errorf("Error reading trusthost1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost1: %v", err)
		}
	}

	if err = d.Set("trusthost10", flattenSwitchControllerSecurityPolicyAdminTrusthost10(o["trusthost10"], d, "trusthost10")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost10"], "SwitchControllerSecurityPolicyAdmin-Trusthost10"); ok {
			if err = d.Set("trusthost10", vv); err != nil {
				return fmt.Errorf("Error reading trusthost10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost10: %v", err)
		}
	}

	if err = d.Set("trusthost2", flattenSwitchControllerSecurityPolicyAdminTrusthost2(o["trusthost2"], d, "trusthost2")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost2"], "SwitchControllerSecurityPolicyAdmin-Trusthost2"); ok {
			if err = d.Set("trusthost2", vv); err != nil {
				return fmt.Errorf("Error reading trusthost2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost2: %v", err)
		}
	}

	if err = d.Set("trusthost3", flattenSwitchControllerSecurityPolicyAdminTrusthost3(o["trusthost3"], d, "trusthost3")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost3"], "SwitchControllerSecurityPolicyAdmin-Trusthost3"); ok {
			if err = d.Set("trusthost3", vv); err != nil {
				return fmt.Errorf("Error reading trusthost3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost3: %v", err)
		}
	}

	if err = d.Set("trusthost4", flattenSwitchControllerSecurityPolicyAdminTrusthost4(o["trusthost4"], d, "trusthost4")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost4"], "SwitchControllerSecurityPolicyAdmin-Trusthost4"); ok {
			if err = d.Set("trusthost4", vv); err != nil {
				return fmt.Errorf("Error reading trusthost4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost4: %v", err)
		}
	}

	if err = d.Set("trusthost5", flattenSwitchControllerSecurityPolicyAdminTrusthost5(o["trusthost5"], d, "trusthost5")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost5"], "SwitchControllerSecurityPolicyAdmin-Trusthost5"); ok {
			if err = d.Set("trusthost5", vv); err != nil {
				return fmt.Errorf("Error reading trusthost5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost5: %v", err)
		}
	}

	if err = d.Set("trusthost6", flattenSwitchControllerSecurityPolicyAdminTrusthost6(o["trusthost6"], d, "trusthost6")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost6"], "SwitchControllerSecurityPolicyAdmin-Trusthost6"); ok {
			if err = d.Set("trusthost6", vv); err != nil {
				return fmt.Errorf("Error reading trusthost6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost6: %v", err)
		}
	}

	if err = d.Set("trusthost7", flattenSwitchControllerSecurityPolicyAdminTrusthost7(o["trusthost7"], d, "trusthost7")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost7"], "SwitchControllerSecurityPolicyAdmin-Trusthost7"); ok {
			if err = d.Set("trusthost7", vv); err != nil {
				return fmt.Errorf("Error reading trusthost7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost7: %v", err)
		}
	}

	if err = d.Set("trusthost8", flattenSwitchControllerSecurityPolicyAdminTrusthost8(o["trusthost8"], d, "trusthost8")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost8"], "SwitchControllerSecurityPolicyAdmin-Trusthost8"); ok {
			if err = d.Set("trusthost8", vv); err != nil {
				return fmt.Errorf("Error reading trusthost8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost8: %v", err)
		}
	}

	if err = d.Set("trusthost9", flattenSwitchControllerSecurityPolicyAdminTrusthost9(o["trusthost9"], d, "trusthost9")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost9"], "SwitchControllerSecurityPolicyAdmin-Trusthost9"); ok {
			if err = d.Set("trusthost9", vv); err != nil {
				return fmt.Errorf("Error reading trusthost9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost9: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSecurityPolicyAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSecurityPolicyAdminAuto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectSwitchControllerSecurityPolicyAdmin(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto"); ok || d.HasChange("auto") {
		t, err := expandSwitchControllerSecurityPolicyAdminAuto(d, v, "auto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost1"); ok || d.HasChange("ip6_trusthost1") {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost1(d, v, "ip6_trusthost1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost10"); ok || d.HasChange("ip6_trusthost10") {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost10(d, v, "ip6_trusthost10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost2"); ok || d.HasChange("ip6_trusthost2") {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost2(d, v, "ip6_trusthost2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost3"); ok || d.HasChange("ip6_trusthost3") {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost3(d, v, "ip6_trusthost3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost4"); ok || d.HasChange("ip6_trusthost4") {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost4(d, v, "ip6_trusthost4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost5"); ok || d.HasChange("ip6_trusthost5") {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost5(d, v, "ip6_trusthost5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost6"); ok || d.HasChange("ip6_trusthost6") {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost6(d, v, "ip6_trusthost6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost7"); ok || d.HasChange("ip6_trusthost7") {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost7(d, v, "ip6_trusthost7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost8"); ok || d.HasChange("ip6_trusthost8") {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost8(d, v, "ip6_trusthost8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost9"); ok || d.HasChange("ip6_trusthost9") {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost9(d, v, "ip6_trusthost9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerSecurityPolicyAdminName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("trusthost1"); ok || d.HasChange("trusthost1") {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost1(d, v, "trusthost1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("trusthost10"); ok || d.HasChange("trusthost10") {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost10(d, v, "trusthost10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("trusthost2"); ok || d.HasChange("trusthost2") {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost2(d, v, "trusthost2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("trusthost3"); ok || d.HasChange("trusthost3") {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost3(d, v, "trusthost3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("trusthost4"); ok || d.HasChange("trusthost4") {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost4(d, v, "trusthost4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("trusthost5"); ok || d.HasChange("trusthost5") {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost5(d, v, "trusthost5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("trusthost6"); ok || d.HasChange("trusthost6") {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost6(d, v, "trusthost6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("trusthost7"); ok || d.HasChange("trusthost7") {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost7(d, v, "trusthost7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("trusthost8"); ok || d.HasChange("trusthost8") {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost8(d, v, "trusthost8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("trusthost9"); ok || d.HasChange("trusthost9") {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost9(d, v, "trusthost9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost9"] = t
		}
	}

	return &obj, nil
}
