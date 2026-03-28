// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device WebProxyImplicitProxyRuleProxyServers

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyImplicitProxyRuleProxyServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyImplicitProxyRuleProxyServersCreate,
		Read:   resourceWebProxyImplicitProxyRuleProxyServersRead,
		Update: resourceWebProxyImplicitProxyRuleProxyServersUpdate,
		Delete: resourceWebProxyImplicitProxyRuleProxyServersDelete,

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
			"implicit_proxy_rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebProxyImplicitProxyRuleProxyServersCreate(d *schema.ResourceData, m interface{}) error {
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
	implicit_proxy_rule := d.Get("implicit_proxy_rule").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["implicit_proxy_rule"] = implicit_proxy_rule

	obj, err := getObjectWebProxyImplicitProxyRuleProxyServers(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyImplicitProxyRuleProxyServers resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebProxyImplicitProxyRuleProxyServers(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebProxyImplicitProxyRuleProxyServers(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebProxyImplicitProxyRuleProxyServers resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateWebProxyImplicitProxyRuleProxyServers(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebProxyImplicitProxyRuleProxyServers resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceWebProxyImplicitProxyRuleProxyServersRead(d, m)
			} else {
				return fmt.Errorf("Error creating WebProxyImplicitProxyRuleProxyServers resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebProxyImplicitProxyRuleProxyServersRead(d, m)
}

func resourceWebProxyImplicitProxyRuleProxyServersUpdate(d *schema.ResourceData, m interface{}) error {
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
	implicit_proxy_rule := d.Get("implicit_proxy_rule").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["implicit_proxy_rule"] = implicit_proxy_rule

	obj, err := getObjectWebProxyImplicitProxyRuleProxyServers(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyImplicitProxyRuleProxyServers resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebProxyImplicitProxyRuleProxyServers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyImplicitProxyRuleProxyServers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebProxyImplicitProxyRuleProxyServersRead(d, m)
}

func resourceWebProxyImplicitProxyRuleProxyServersDelete(d *schema.ResourceData, m interface{}) error {
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
	implicit_proxy_rule := d.Get("implicit_proxy_rule").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["implicit_proxy_rule"] = implicit_proxy_rule

	wsParams["adom"] = adomv

	err = c.DeleteWebProxyImplicitProxyRuleProxyServers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyImplicitProxyRuleProxyServers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyImplicitProxyRuleProxyServersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	implicit_proxy_rule := d.Get("implicit_proxy_rule").(string)
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
	if implicit_proxy_rule == "" {
		implicit_proxy_rule = importOptionChecking(m.(*FortiClient).Cfg, "implicit_proxy_rule")
		if implicit_proxy_rule == "" {
			return fmt.Errorf("Parameter implicit_proxy_rule is missing")
		}
		if err = d.Set("implicit_proxy_rule", implicit_proxy_rule); err != nil {
			return fmt.Errorf("Error set params implicit_proxy_rule: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["implicit_proxy_rule"] = implicit_proxy_rule

	o, err := c.ReadWebProxyImplicitProxyRuleProxyServers(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebProxyImplicitProxyRuleProxyServers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyImplicitProxyRuleProxyServers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyImplicitProxyRuleProxyServers resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyImplicitProxyRuleProxyServersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyImplicitProxyRuleProxyServersIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyImplicitProxyRuleProxyServersPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyImplicitProxyRuleProxyServersWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyImplicitProxyRuleProxyServers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenWebProxyImplicitProxyRuleProxyServersId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WebProxyImplicitProxyRuleProxyServers-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenWebProxyImplicitProxyRuleProxyServersIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "WebProxyImplicitProxyRuleProxyServers-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("port", flattenWebProxyImplicitProxyRuleProxyServersPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "WebProxyImplicitProxyRuleProxyServers-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("weight", flattenWebProxyImplicitProxyRuleProxyServersWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "WebProxyImplicitProxyRuleProxyServers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenWebProxyImplicitProxyRuleProxyServersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyImplicitProxyRuleProxyServersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyImplicitProxyRuleProxyServersIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyImplicitProxyRuleProxyServersPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyImplicitProxyRuleProxyServersWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyImplicitProxyRuleProxyServers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWebProxyImplicitProxyRuleProxyServersId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandWebProxyImplicitProxyRuleProxyServersIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandWebProxyImplicitProxyRuleProxyServersPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandWebProxyImplicitProxyRuleProxyServersWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
