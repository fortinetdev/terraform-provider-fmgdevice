// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device WebProxyImplicitProxyRule

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyImplicitProxyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyImplicitProxyRuleCreate,
		Read:   resourceWebProxyImplicitProxyRuleRead,
		Update: resourceWebProxyImplicitProxyRuleUpdate,
		Delete: resourceWebProxyImplicitProxyRuleDelete,

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
			"bypass_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"load_balance_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"proxy_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWebProxyImplicitProxyRuleCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyImplicitProxyRule(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyImplicitProxyRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebProxyImplicitProxyRule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebProxyImplicitProxyRule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebProxyImplicitProxyRule resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebProxyImplicitProxyRule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebProxyImplicitProxyRule resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyImplicitProxyRuleRead(d, m)
}

func resourceWebProxyImplicitProxyRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyImplicitProxyRule(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyImplicitProxyRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebProxyImplicitProxyRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyImplicitProxyRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyImplicitProxyRuleRead(d, m)
}

func resourceWebProxyImplicitProxyRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebProxyImplicitProxyRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyImplicitProxyRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyImplicitProxyRuleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyImplicitProxyRule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebProxyImplicitProxyRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyImplicitProxyRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyImplicitProxyRule resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyImplicitProxyRuleBypassList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyImplicitProxyRuleLoadBalanceMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyImplicitProxyRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyImplicitProxyRuleProxyServers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWebProxyImplicitProxyRuleProxyServersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WebProxyImplicitProxyRule-ProxyServers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenWebProxyImplicitProxyRuleProxyServersIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "WebProxyImplicitProxyRule-ProxyServers-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenWebProxyImplicitProxyRuleProxyServersPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "WebProxyImplicitProxyRule-ProxyServers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenWebProxyImplicitProxyRuleProxyServersWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "WebProxyImplicitProxyRule-ProxyServers-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebProxyImplicitProxyRuleProxyServersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyImplicitProxyRuleProxyServersIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyImplicitProxyRuleProxyServersPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyImplicitProxyRuleProxyServersWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyImplicitProxyRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("bypass_list", flattenWebProxyImplicitProxyRuleBypassList(o["bypass-list"], d, "bypass_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["bypass-list"], "WebProxyImplicitProxyRule-BypassList"); ok {
			if err = d.Set("bypass_list", vv); err != nil {
				return fmt.Errorf("Error reading bypass_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bypass_list: %v", err)
		}
	}

	if err = d.Set("load_balance_method", flattenWebProxyImplicitProxyRuleLoadBalanceMethod(o["load-balance-method"], d, "load_balance_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balance-method"], "WebProxyImplicitProxyRule-LoadBalanceMethod"); ok {
			if err = d.Set("load_balance_method", vv); err != nil {
				return fmt.Errorf("Error reading load_balance_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balance_method: %v", err)
		}
	}

	if err = d.Set("name", flattenWebProxyImplicitProxyRuleName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebProxyImplicitProxyRule-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("proxy_servers", flattenWebProxyImplicitProxyRuleProxyServers(o["proxy-servers"], d, "proxy_servers")); err != nil {
			if vv, ok := fortiAPIPatch(o["proxy-servers"], "WebProxyImplicitProxyRule-ProxyServers"); ok {
				if err = d.Set("proxy_servers", vv); err != nil {
					return fmt.Errorf("Error reading proxy_servers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading proxy_servers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("proxy_servers"); ok {
			if err = d.Set("proxy_servers", flattenWebProxyImplicitProxyRuleProxyServers(o["proxy-servers"], d, "proxy_servers")); err != nil {
				if vv, ok := fortiAPIPatch(o["proxy-servers"], "WebProxyImplicitProxyRule-ProxyServers"); ok {
					if err = d.Set("proxy_servers", vv); err != nil {
						return fmt.Errorf("Error reading proxy_servers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading proxy_servers: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWebProxyImplicitProxyRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyImplicitProxyRuleBypassList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyImplicitProxyRuleLoadBalanceMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyImplicitProxyRuleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyImplicitProxyRuleProxyServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWebProxyImplicitProxyRuleProxyServersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandWebProxyImplicitProxyRuleProxyServersIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandWebProxyImplicitProxyRuleProxyServersPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandWebProxyImplicitProxyRuleProxyServersWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebProxyImplicitProxyRuleProxyServersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyImplicitProxyRuleProxyServersIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyImplicitProxyRuleProxyServersPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyImplicitProxyRuleProxyServersWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyImplicitProxyRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bypass_list"); ok || d.HasChange("bypass_list") {
		t, err := expandWebProxyImplicitProxyRuleBypassList(d, v, "bypass_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bypass-list"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_method"); ok || d.HasChange("load_balance_method") {
		t, err := expandWebProxyImplicitProxyRuleLoadBalanceMethod(d, v, "load_balance_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebProxyImplicitProxyRuleName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("proxy_servers"); ok || d.HasChange("proxy_servers") {
		t, err := expandWebProxyImplicitProxyRuleProxyServers(d, v, "proxy_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-servers"] = t
		}
	}

	return &obj, nil
}
