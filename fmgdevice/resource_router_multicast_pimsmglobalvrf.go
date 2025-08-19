// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: per-VRF PIM sparse-mode global settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticastPimSmGlobalVrf() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticastPimSmGlobalVrfCreate,
		Read:   resourceRouterMulticastPimSmGlobalVrfRead,
		Update: resourceRouterMulticastPimSmGlobalVrfUpdate,
		Delete: resourceRouterMulticastPimSmGlobalVrfDelete,

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
			"bsr_allow_quick_refresh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bsr_candidate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bsr_hash": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bsr_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"bsr_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cisco_crp_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rp_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceRouterMulticastPimSmGlobalVrfCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterMulticastPimSmGlobalVrf(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastPimSmGlobalVrf resource while getting object: %v", err)
	}

	_, err = c.CreateRouterMulticastPimSmGlobalVrf(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastPimSmGlobalVrf resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "vrf")))

	return resourceRouterMulticastPimSmGlobalVrfRead(d, m)
}

func resourceRouterMulticastPimSmGlobalVrfUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterMulticastPimSmGlobalVrf(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastPimSmGlobalVrf resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticastPimSmGlobalVrf(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastPimSmGlobalVrf resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "vrf")))

	return resourceRouterMulticastPimSmGlobalVrfRead(d, m)
}

func resourceRouterMulticastPimSmGlobalVrfDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterMulticastPimSmGlobalVrf(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticastPimSmGlobalVrf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastPimSmGlobalVrfRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterMulticastPimSmGlobalVrf(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastPimSmGlobalVrf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticastPimSmGlobalVrf(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastPimSmGlobalVrf resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticastPimSmGlobalVrfBsrAllowQuickRefresh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfBsrCandidate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfBsrHash2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfBsrInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalVrfBsrPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfCiscoCrpPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfRpAddress2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {
			v := flattenRouterMulticastPimSmGlobalVrfRpAddressGroup2edl(i["group"], d, pre_append)
			tmp["group"] = fortiAPISubPartPatch(v, "RouterMulticastPimSmGlobalVrf-RpAddress-Group")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterMulticastPimSmGlobalVrfRpAddressId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterMulticastPimSmGlobalVrf-RpAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := i["ip-address"]; ok {
			v := flattenRouterMulticastPimSmGlobalVrfRpAddressIpAddress2edl(i["ip-address"], d, pre_append)
			tmp["ip_address"] = fortiAPISubPartPatch(v, "RouterMulticastPimSmGlobalVrf-RpAddress-IpAddress")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterMulticastPimSmGlobalVrfRpAddressGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalVrfRpAddressId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfRpAddressIpAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfVrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticastPimSmGlobalVrf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("bsr_allow_quick_refresh", flattenRouterMulticastPimSmGlobalVrfBsrAllowQuickRefresh2edl(o["bsr-allow-quick-refresh"], d, "bsr_allow_quick_refresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-allow-quick-refresh"], "RouterMulticastPimSmGlobalVrf-BsrAllowQuickRefresh"); ok {
			if err = d.Set("bsr_allow_quick_refresh", vv); err != nil {
				return fmt.Errorf("Error reading bsr_allow_quick_refresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_allow_quick_refresh: %v", err)
		}
	}

	if err = d.Set("bsr_candidate", flattenRouterMulticastPimSmGlobalVrfBsrCandidate2edl(o["bsr-candidate"], d, "bsr_candidate")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-candidate"], "RouterMulticastPimSmGlobalVrf-BsrCandidate"); ok {
			if err = d.Set("bsr_candidate", vv); err != nil {
				return fmt.Errorf("Error reading bsr_candidate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_candidate: %v", err)
		}
	}

	if err = d.Set("bsr_hash", flattenRouterMulticastPimSmGlobalVrfBsrHash2edl(o["bsr-hash"], d, "bsr_hash")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-hash"], "RouterMulticastPimSmGlobalVrf-BsrHash"); ok {
			if err = d.Set("bsr_hash", vv); err != nil {
				return fmt.Errorf("Error reading bsr_hash: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_hash: %v", err)
		}
	}

	if err = d.Set("bsr_interface", flattenRouterMulticastPimSmGlobalVrfBsrInterface2edl(o["bsr-interface"], d, "bsr_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-interface"], "RouterMulticastPimSmGlobalVrf-BsrInterface"); ok {
			if err = d.Set("bsr_interface", vv); err != nil {
				return fmt.Errorf("Error reading bsr_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_interface: %v", err)
		}
	}

	if err = d.Set("bsr_priority", flattenRouterMulticastPimSmGlobalVrfBsrPriority2edl(o["bsr-priority"], d, "bsr_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-priority"], "RouterMulticastPimSmGlobalVrf-BsrPriority"); ok {
			if err = d.Set("bsr_priority", vv); err != nil {
				return fmt.Errorf("Error reading bsr_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_priority: %v", err)
		}
	}

	if err = d.Set("cisco_crp_prefix", flattenRouterMulticastPimSmGlobalVrfCiscoCrpPrefix2edl(o["cisco-crp-prefix"], d, "cisco_crp_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["cisco-crp-prefix"], "RouterMulticastPimSmGlobalVrf-CiscoCrpPrefix"); ok {
			if err = d.Set("cisco_crp_prefix", vv); err != nil {
				return fmt.Errorf("Error reading cisco_crp_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cisco_crp_prefix: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rp_address", flattenRouterMulticastPimSmGlobalVrfRpAddress2edl(o["rp-address"], d, "rp_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["rp-address"], "RouterMulticastPimSmGlobalVrf-RpAddress"); ok {
				if err = d.Set("rp_address", vv); err != nil {
					return fmt.Errorf("Error reading rp_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rp_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rp_address"); ok {
			if err = d.Set("rp_address", flattenRouterMulticastPimSmGlobalVrfRpAddress2edl(o["rp-address"], d, "rp_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["rp-address"], "RouterMulticastPimSmGlobalVrf-RpAddress"); ok {
					if err = d.Set("rp_address", vv); err != nil {
						return fmt.Errorf("Error reading rp_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rp_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("vrf", flattenRouterMulticastPimSmGlobalVrfVrf2edl(o["vrf"], d, "vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf"], "RouterMulticastPimSmGlobalVrf-Vrf"); ok {
			if err = d.Set("vrf", vv); err != nil {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticastPimSmGlobalVrfFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticastPimSmGlobalVrfBsrAllowQuickRefresh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfBsrCandidate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfBsrHash2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfBsrInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalVrfBsrPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfCiscoCrpPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfRpAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group"], _ = expandRouterMulticastPimSmGlobalVrfRpAddressGroup2edl(d, i["group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterMulticastPimSmGlobalVrfRpAddressId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-address"], _ = expandRouterMulticastPimSmGlobalVrfRpAddressIpAddress2edl(d, i["ip_address"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterMulticastPimSmGlobalVrfRpAddressGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalVrfRpAddressId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfRpAddressIpAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfVrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticastPimSmGlobalVrf(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bsr_allow_quick_refresh"); ok || d.HasChange("bsr_allow_quick_refresh") {
		t, err := expandRouterMulticastPimSmGlobalVrfBsrAllowQuickRefresh2edl(d, v, "bsr_allow_quick_refresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-allow-quick-refresh"] = t
		}
	}

	if v, ok := d.GetOk("bsr_candidate"); ok || d.HasChange("bsr_candidate") {
		t, err := expandRouterMulticastPimSmGlobalVrfBsrCandidate2edl(d, v, "bsr_candidate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-candidate"] = t
		}
	}

	if v, ok := d.GetOk("bsr_hash"); ok || d.HasChange("bsr_hash") {
		t, err := expandRouterMulticastPimSmGlobalVrfBsrHash2edl(d, v, "bsr_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-hash"] = t
		}
	}

	if v, ok := d.GetOk("bsr_interface"); ok || d.HasChange("bsr_interface") {
		t, err := expandRouterMulticastPimSmGlobalVrfBsrInterface2edl(d, v, "bsr_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-interface"] = t
		}
	}

	if v, ok := d.GetOk("bsr_priority"); ok || d.HasChange("bsr_priority") {
		t, err := expandRouterMulticastPimSmGlobalVrfBsrPriority2edl(d, v, "bsr_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-priority"] = t
		}
	}

	if v, ok := d.GetOk("cisco_crp_prefix"); ok || d.HasChange("cisco_crp_prefix") {
		t, err := expandRouterMulticastPimSmGlobalVrfCiscoCrpPrefix2edl(d, v, "cisco_crp_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cisco-crp-prefix"] = t
		}
	}

	if v, ok := d.GetOk("rp_address"); ok || d.HasChange("rp_address") {
		t, err := expandRouterMulticastPimSmGlobalVrfRpAddress2edl(d, v, "rp_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rp-address"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		t, err := expandRouterMulticastPimSmGlobalVrfVrf2edl(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	return &obj, nil
}
