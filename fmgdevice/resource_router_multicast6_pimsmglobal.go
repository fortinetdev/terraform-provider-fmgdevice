// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: PIM sparse-mode global settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticast6PimSmGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticast6PimSmGlobalUpdate,
		Read:   resourceRouterMulticast6PimSmGlobalRead,
		Update: resourceRouterMulticast6PimSmGlobalUpdate,
		Delete: resourceRouterMulticast6PimSmGlobalDelete,

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
			"register_rate_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rp_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
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

func resourceRouterMulticast6PimSmGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectRouterMulticast6PimSmGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6PimSmGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticast6PimSmGlobal(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6PimSmGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterMulticast6PimSmGlobal")

	return resourceRouterMulticast6PimSmGlobalRead(d, m)
}

func resourceRouterMulticast6PimSmGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteRouterMulticast6PimSmGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticast6PimSmGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticast6PimSmGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterMulticast6PimSmGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast6PimSmGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticast6PimSmGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast6PimSmGlobal resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticast6PimSmGlobalRegisterRateLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalRpAddress2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterMulticast6PimSmGlobalRpAddressId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterMulticast6PimSmGlobal-RpAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := i["ip6-address"]; ok {
			v := flattenRouterMulticast6PimSmGlobalRpAddressIp6Address2edl(i["ip6-address"], d, pre_append)
			tmp["ip6_address"] = fortiAPISubPartPatch(v, "RouterMulticast6PimSmGlobal-RpAddress-Ip6Address")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterMulticast6PimSmGlobalRpAddressId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalRpAddressIp6Address2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticast6PimSmGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("register_rate_limit", flattenRouterMulticast6PimSmGlobalRegisterRateLimit2edl(o["register-rate-limit"], d, "register_rate_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-rate-limit"], "RouterMulticast6PimSmGlobal-RegisterRateLimit"); ok {
			if err = d.Set("register_rate_limit", vv); err != nil {
				return fmt.Errorf("Error reading register_rate_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_rate_limit: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rp_address", flattenRouterMulticast6PimSmGlobalRpAddress2edl(o["rp-address"], d, "rp_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["rp-address"], "RouterMulticast6PimSmGlobal-RpAddress"); ok {
				if err = d.Set("rp_address", vv); err != nil {
					return fmt.Errorf("Error reading rp_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rp_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rp_address"); ok {
			if err = d.Set("rp_address", flattenRouterMulticast6PimSmGlobalRpAddress2edl(o["rp-address"], d, "rp_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["rp-address"], "RouterMulticast6PimSmGlobal-RpAddress"); ok {
					if err = d.Set("rp_address", vv); err != nil {
						return fmt.Errorf("Error reading rp_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rp_address: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterMulticast6PimSmGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticast6PimSmGlobalRegisterRateLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalRpAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterMulticast6PimSmGlobalRpAddressId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6-address"], _ = expandRouterMulticast6PimSmGlobalRpAddressIp6Address2edl(d, i["ip6_address"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterMulticast6PimSmGlobalRpAddressId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalRpAddressIp6Address2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticast6PimSmGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("register_rate_limit"); ok || d.HasChange("register_rate_limit") {
		t, err := expandRouterMulticast6PimSmGlobalRegisterRateLimit2edl(d, v, "register_rate_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-rate-limit"] = t
		}
	}

	if v, ok := d.GetOk("rp_address"); ok || d.HasChange("rp_address") {
		t, err := expandRouterMulticast6PimSmGlobalRpAddress2edl(d, v, "rp_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rp-address"] = t
		}
	}

	return &obj, nil
}
