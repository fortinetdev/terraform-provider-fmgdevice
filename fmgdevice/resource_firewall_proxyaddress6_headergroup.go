// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProxyAddress6HeaderGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProxyAddress6HeaderGroupCreate,
		Read:   resourceFirewallProxyAddress6HeaderGroupRead,
		Update: resourceFirewallProxyAddress6HeaderGroupUpdate,
		Delete: resourceFirewallProxyAddress6HeaderGroupDelete,

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
			"proxy_address6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"case_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"header_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceFirewallProxyAddress6HeaderGroupCreate(d *schema.ResourceData, m interface{}) error {
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
	proxy_address6 := d.Get("proxy_address6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["proxy_address6"] = proxy_address6

	obj, err := getObjectFirewallProxyAddress6HeaderGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallProxyAddress6HeaderGroup resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallProxyAddress6HeaderGroup(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallProxyAddress6HeaderGroup(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallProxyAddress6HeaderGroup resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallProxyAddress6HeaderGroup(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallProxyAddress6HeaderGroup resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallProxyAddress6HeaderGroupRead(d, m)
}

func resourceFirewallProxyAddress6HeaderGroupUpdate(d *schema.ResourceData, m interface{}) error {
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
	proxy_address6 := d.Get("proxy_address6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["proxy_address6"] = proxy_address6

	obj, err := getObjectFirewallProxyAddress6HeaderGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProxyAddress6HeaderGroup resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallProxyAddress6HeaderGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProxyAddress6HeaderGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallProxyAddress6HeaderGroupRead(d, m)
}

func resourceFirewallProxyAddress6HeaderGroupDelete(d *schema.ResourceData, m interface{}) error {
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
	proxy_address6 := d.Get("proxy_address6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["proxy_address6"] = proxy_address6

	wsParams["adom"] = adomv

	err = c.DeleteFirewallProxyAddress6HeaderGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProxyAddress6HeaderGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProxyAddress6HeaderGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	proxy_address6 := d.Get("proxy_address6").(string)
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
	if proxy_address6 == "" {
		proxy_address6 = importOptionChecking(m.(*FortiClient).Cfg, "proxy_address6")
		if proxy_address6 == "" {
			return fmt.Errorf("Parameter proxy_address6 is missing")
		}
		if err = d.Set("proxy_address6", proxy_address6); err != nil {
			return fmt.Errorf("Error set params proxy_address6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["proxy_address6"] = proxy_address6

	o, err := c.ReadFirewallProxyAddress6HeaderGroup(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallProxyAddress6HeaderGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProxyAddress6HeaderGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProxyAddress6HeaderGroup resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProxyAddress6HeaderGroupCaseSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6HeaderGroupHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6HeaderGroupHeaderName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6HeaderGroupId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallProxyAddress6HeaderGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("case_sensitivity", flattenFirewallProxyAddress6HeaderGroupCaseSensitivity2edl(o["case-sensitivity"], d, "case_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-sensitivity"], "FirewallProxyAddress6HeaderGroup-CaseSensitivity"); ok {
			if err = d.Set("case_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading case_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_sensitivity: %v", err)
		}
	}

	if err = d.Set("header", flattenFirewallProxyAddress6HeaderGroupHeader2edl(o["header"], d, "header")); err != nil {
		if vv, ok := fortiAPIPatch(o["header"], "FirewallProxyAddress6HeaderGroup-Header"); ok {
			if err = d.Set("header", vv); err != nil {
				return fmt.Errorf("Error reading header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("header_name", flattenFirewallProxyAddress6HeaderGroupHeaderName2edl(o["header-name"], d, "header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-name"], "FirewallProxyAddress6HeaderGroup-HeaderName"); ok {
			if err = d.Set("header_name", vv); err != nil {
				return fmt.Errorf("Error reading header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallProxyAddress6HeaderGroupId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallProxyAddress6HeaderGroup-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenFirewallProxyAddress6HeaderGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProxyAddress6HeaderGroupCaseSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6HeaderGroupHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6HeaderGroupHeaderName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6HeaderGroupId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProxyAddress6HeaderGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("case_sensitivity"); ok || d.HasChange("case_sensitivity") {
		t, err := expandFirewallProxyAddress6HeaderGroupCaseSensitivity2edl(d, v, "case_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok || d.HasChange("header") {
		t, err := expandFirewallProxyAddress6HeaderGroupHeader2edl(d, v, "header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok || d.HasChange("header_name") {
		t, err := expandFirewallProxyAddress6HeaderGroupHeaderName2edl(d, v, "header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallProxyAddress6HeaderGroupId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
