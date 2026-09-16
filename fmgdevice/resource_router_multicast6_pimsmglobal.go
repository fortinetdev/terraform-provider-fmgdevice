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
			"bsr_allow_quick_refresh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bsr_candidate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bsr_hash": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"cisco_ignore_rp_set_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pim_use_sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"spt_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spt_threshold_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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

	obj, err := getObjectRouterMulticast6PimSmGlobal(d, false)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6PimSmGlobal resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterMulticast6PimSmGlobal(obj, mkey, paradict, wsParams)
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

	obj, err := getObjectRouterMulticast6PimSmGlobal(d, true)

	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6PimSmGlobal resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterMulticast6PimSmGlobal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing RouterMulticast6PimSmGlobal resource: %v", err)
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
		d.SetId("")
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

func flattenRouterMulticast6PimSmGlobalBsrAllowQuickRefresh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalBsrCandidate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalBsrHash2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalBsrInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticast6PimSmGlobalBsrPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalCiscoCrpPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalPimUseSdwan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {
			v := flattenRouterMulticast6PimSmGlobalRpAddressGroup2edl(i["group"], d, pre_append)
			tmp["group"] = fortiAPISubPartPatch(v, "RouterMulticast6PimSmGlobal-RpAddress-Group")
		}

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

func flattenRouterMulticast6PimSmGlobalRpAddressGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticast6PimSmGlobalRpAddressId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalRpAddressIp6Address2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalSptThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalSptThresholdGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectRouterMulticast6PimSmGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("bsr_allow_quick_refresh", flattenRouterMulticast6PimSmGlobalBsrAllowQuickRefresh2edl(o["bsr-allow-quick-refresh"], d, "bsr_allow_quick_refresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-allow-quick-refresh"], "RouterMulticast6PimSmGlobal-BsrAllowQuickRefresh"); ok {
			if err = d.Set("bsr_allow_quick_refresh", vv); err != nil {
				return fmt.Errorf("Error reading bsr_allow_quick_refresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_allow_quick_refresh: %v", err)
		}
	}

	if err = d.Set("bsr_candidate", flattenRouterMulticast6PimSmGlobalBsrCandidate2edl(o["bsr-candidate"], d, "bsr_candidate")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-candidate"], "RouterMulticast6PimSmGlobal-BsrCandidate"); ok {
			if err = d.Set("bsr_candidate", vv); err != nil {
				return fmt.Errorf("Error reading bsr_candidate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_candidate: %v", err)
		}
	}

	if err = d.Set("bsr_hash", flattenRouterMulticast6PimSmGlobalBsrHash2edl(o["bsr-hash"], d, "bsr_hash")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-hash"], "RouterMulticast6PimSmGlobal-BsrHash"); ok {
			if err = d.Set("bsr_hash", vv); err != nil {
				return fmt.Errorf("Error reading bsr_hash: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_hash: %v", err)
		}
	}

	if err = d.Set("bsr_interface", flattenRouterMulticast6PimSmGlobalBsrInterface2edl(o["bsr-interface"], d, "bsr_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-interface"], "RouterMulticast6PimSmGlobal-BsrInterface"); ok {
			if err = d.Set("bsr_interface", vv); err != nil {
				return fmt.Errorf("Error reading bsr_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_interface: %v", err)
		}
	}

	if err = d.Set("bsr_priority", flattenRouterMulticast6PimSmGlobalBsrPriority2edl(o["bsr-priority"], d, "bsr_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-priority"], "RouterMulticast6PimSmGlobal-BsrPriority"); ok {
			if err = d.Set("bsr_priority", vv); err != nil {
				return fmt.Errorf("Error reading bsr_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_priority: %v", err)
		}
	}

	if err = d.Set("cisco_crp_prefix", flattenRouterMulticast6PimSmGlobalCiscoCrpPrefix2edl(o["cisco-crp-prefix"], d, "cisco_crp_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["cisco-crp-prefix"], "RouterMulticast6PimSmGlobal-CiscoCrpPrefix"); ok {
			if err = d.Set("cisco_crp_prefix", vv); err != nil {
				return fmt.Errorf("Error reading cisco_crp_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cisco_crp_prefix: %v", err)
		}
	}

	if err = d.Set("cisco_ignore_rp_set_priority", flattenRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority2edl(o["cisco-ignore-rp-set-priority"], d, "cisco_ignore_rp_set_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["cisco-ignore-rp-set-priority"], "RouterMulticast6PimSmGlobal-CiscoIgnoreRpSetPriority"); ok {
			if err = d.Set("cisco_ignore_rp_set_priority", vv); err != nil {
				return fmt.Errorf("Error reading cisco_ignore_rp_set_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cisco_ignore_rp_set_priority: %v", err)
		}
	}

	if err = d.Set("pim_use_sdwan", flattenRouterMulticast6PimSmGlobalPimUseSdwan2edl(o["pim-use-sdwan"], d, "pim_use_sdwan")); err != nil {
		if vv, ok := fortiAPIPatch(o["pim-use-sdwan"], "RouterMulticast6PimSmGlobal-PimUseSdwan"); ok {
			if err = d.Set("pim_use_sdwan", vv); err != nil {
				return fmt.Errorf("Error reading pim_use_sdwan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pim_use_sdwan: %v", err)
		}
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

	if err = d.Set("spt_threshold", flattenRouterMulticast6PimSmGlobalSptThreshold2edl(o["spt-threshold"], d, "spt_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["spt-threshold"], "RouterMulticast6PimSmGlobal-SptThreshold"); ok {
			if err = d.Set("spt_threshold", vv); err != nil {
				return fmt.Errorf("Error reading spt_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spt_threshold: %v", err)
		}
	}

	if err = d.Set("spt_threshold_group", flattenRouterMulticast6PimSmGlobalSptThresholdGroup2edl(o["spt-threshold-group"], d, "spt_threshold_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["spt-threshold-group"], "RouterMulticast6PimSmGlobal-SptThresholdGroup"); ok {
			if err = d.Set("spt_threshold_group", vv); err != nil {
				return fmt.Errorf("Error reading spt_threshold_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spt_threshold_group: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticast6PimSmGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticast6PimSmGlobalBsrAllowQuickRefresh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalBsrCandidate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalBsrHash2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalBsrInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticast6PimSmGlobalBsrPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalCiscoCrpPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalPimUseSdwan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group"], _ = expandRouterMulticast6PimSmGlobalRpAddressGroup2edl(d, i["group"], pre_append)
		}

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

func expandRouterMulticast6PimSmGlobalRpAddressGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticast6PimSmGlobalRpAddressId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalRpAddressIp6Address2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalSptThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalSptThresholdGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectRouterMulticast6PimSmGlobal(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bsr_allow_quick_refresh"); ok || d.HasChange("bsr_allow_quick_refresh") {
		t, err := expandRouterMulticast6PimSmGlobalBsrAllowQuickRefresh2edl(d, v, "bsr_allow_quick_refresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-allow-quick-refresh"] = t
		}
	}

	if v, ok := d.GetOk("bsr_candidate"); ok || d.HasChange("bsr_candidate") {
		t, err := expandRouterMulticast6PimSmGlobalBsrCandidate2edl(d, v, "bsr_candidate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-candidate"] = t
		}
	}

	if v, ok := d.GetOk("bsr_hash"); ok || d.HasChange("bsr_hash") {
		t, err := expandRouterMulticast6PimSmGlobalBsrHash2edl(d, v, "bsr_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-hash"] = t
		}
	}

	if v, ok := d.GetOk("bsr_interface"); ok || d.HasChange("bsr_interface") {
		t, err := expandRouterMulticast6PimSmGlobalBsrInterface2edl(d, v, "bsr_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-interface"] = t
		}
	}

	if v, ok := d.GetOk("bsr_priority"); ok || d.HasChange("bsr_priority") {
		t, err := expandRouterMulticast6PimSmGlobalBsrPriority2edl(d, v, "bsr_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-priority"] = t
		}
	}

	if v, ok := d.GetOk("cisco_crp_prefix"); ok || d.HasChange("cisco_crp_prefix") {
		t, err := expandRouterMulticast6PimSmGlobalCiscoCrpPrefix2edl(d, v, "cisco_crp_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cisco-crp-prefix"] = t
		}
	}

	if v, ok := d.GetOk("cisco_ignore_rp_set_priority"); ok || d.HasChange("cisco_ignore_rp_set_priority") {
		t, err := expandRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority2edl(d, v, "cisco_ignore_rp_set_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cisco-ignore-rp-set-priority"] = t
		}
	}

	if v, ok := d.GetOk("pim_use_sdwan"); ok || d.HasChange("pim_use_sdwan") {
		t, err := expandRouterMulticast6PimSmGlobalPimUseSdwan2edl(d, v, "pim_use_sdwan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pim-use-sdwan"] = t
		}
	}

	if v, ok := d.GetOk("register_rate_limit"); ok || d.HasChange("register_rate_limit") {
		t, err := expandRouterMulticast6PimSmGlobalRegisterRateLimit2edl(d, v, "register_rate_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-rate-limit"] = t
		}
	}

	if bemptysontable {
		obj["rp-address"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("rp_address"); ok || d.HasChange("rp_address") {
			t, err := expandRouterMulticast6PimSmGlobalRpAddress2edl(d, v, "rp_address")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rp-address"] = t
			}
		}
	}

	if v, ok := d.GetOk("spt_threshold"); ok || d.HasChange("spt_threshold") {
		t, err := expandRouterMulticast6PimSmGlobalSptThreshold2edl(d, v, "spt_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spt-threshold"] = t
		}
	}

	if v, ok := d.GetOk("spt_threshold_group"); ok || d.HasChange("spt_threshold_group") {
		t, err := expandRouterMulticast6PimSmGlobalSptThresholdGroup2edl(d, v, "spt_threshold_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spt-threshold-group"] = t
		}
	}

	return &obj, nil
}
