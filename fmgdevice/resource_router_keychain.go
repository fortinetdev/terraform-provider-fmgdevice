// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure key-chain.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterKeyChain() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterKeyChainCreate,
		Read:   resourceRouterKeyChainRead,
		Update: resourceRouterKeyChainUpdate,
		Delete: resourceRouterKeyChainDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"key": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_lifetime": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"key_string": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"send_lifetime": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterKeyChainCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterKeyChain(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterKeyChain resource while getting object: %v", err)
	}

	_, err = c.CreateRouterKeyChain(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterKeyChain resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterKeyChainRead(d, m)
}

func resourceRouterKeyChainUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterKeyChain(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterKeyChain resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterKeyChain(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterKeyChain resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterKeyChainRead(d, m)
}

func resourceRouterKeyChainDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteRouterKeyChain(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterKeyChain resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterKeyChainRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterKeyChain(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterKeyChain resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterKeyChain(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterKeyChain resource from API: %v", err)
	}
	return nil
}

func flattenRouterKeyChainKey(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_lifetime"
		if _, ok := i["accept-lifetime"]; ok {
			v := flattenRouterKeyChainKeyAcceptLifetime(i["accept-lifetime"], d, pre_append)
			tmp["accept_lifetime"] = fortiAPISubPartPatch(v, "RouterKeyChain-Key-AcceptLifetime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "algorithm"
		if _, ok := i["algorithm"]; ok {
			v := flattenRouterKeyChainKeyAlgorithm(i["algorithm"], d, pre_append)
			tmp["algorithm"] = fortiAPISubPartPatch(v, "RouterKeyChain-Key-Algorithm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterKeyChainKeyId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterKeyChain-Key-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_lifetime"
		if _, ok := i["send-lifetime"]; ok {
			v := flattenRouterKeyChainKeySendLifetime(i["send-lifetime"], d, pre_append)
			tmp["send_lifetime"] = fortiAPISubPartPatch(v, "RouterKeyChain-Key-SendLifetime")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterKeyChainKeyAcceptLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterKeyChainKeyAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterKeyChainKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterKeyChainKeySendLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterKeyChainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterKeyChain(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("key", flattenRouterKeyChainKey(o["key"], d, "key")); err != nil {
			if vv, ok := fortiAPIPatch(o["key"], "RouterKeyChain-Key"); ok {
				if err = d.Set("key", vv); err != nil {
					return fmt.Errorf("Error reading key: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading key: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("key"); ok {
			if err = d.Set("key", flattenRouterKeyChainKey(o["key"], d, "key")); err != nil {
				if vv, ok := fortiAPIPatch(o["key"], "RouterKeyChain-Key"); ok {
					if err = d.Set("key", vv); err != nil {
						return fmt.Errorf("Error reading key: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading key: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenRouterKeyChainName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterKeyChain-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenRouterKeyChainFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterKeyChainKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_lifetime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["accept-lifetime"], _ = expandRouterKeyChainKeyAcceptLifetime(d, i["accept_lifetime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "algorithm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["algorithm"], _ = expandRouterKeyChainKeyAlgorithm(d, i["algorithm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterKeyChainKeyId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-string"], _ = expandRouterKeyChainKeyKeyString(d, i["key_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_lifetime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-lifetime"], _ = expandRouterKeyChainKeySendLifetime(d, i["send_lifetime"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterKeyChainKeyAcceptLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeyAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeyKeyString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterKeyChainKeySendLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterKeyChain(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandRouterKeyChainKey(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterKeyChainName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
