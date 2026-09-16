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

func resourceRouterMulticast6PimSmGlobalVrf() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticast6PimSmGlobalVrfCreate,
		Read:   resourceRouterMulticast6PimSmGlobalVrfRead,
		Update: resourceRouterMulticast6PimSmGlobalVrfUpdate,
		Delete: resourceRouterMulticast6PimSmGlobalVrfDelete,

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
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceRouterMulticast6PimSmGlobalVrfCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterMulticast6PimSmGlobalVrf(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticast6PimSmGlobalVrf resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("vrf")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterMulticast6PimSmGlobalVrf(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterMulticast6PimSmGlobalVrf(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterMulticast6PimSmGlobalVrf resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateRouterMulticast6PimSmGlobalVrf(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterMulticast6PimSmGlobalVrf resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "vrf")))

	return resourceRouterMulticast6PimSmGlobalVrfRead(d, m)
}

func resourceRouterMulticast6PimSmGlobalVrfUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterMulticast6PimSmGlobalVrf(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6PimSmGlobalVrf resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterMulticast6PimSmGlobalVrf(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6PimSmGlobalVrf resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "vrf")))

	return resourceRouterMulticast6PimSmGlobalVrfRead(d, m)
}

func resourceRouterMulticast6PimSmGlobalVrfDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterMulticast6PimSmGlobalVrf(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticast6PimSmGlobalVrf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticast6PimSmGlobalVrfRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterMulticast6PimSmGlobalVrf(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterMulticast6PimSmGlobalVrf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticast6PimSmGlobalVrf(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast6PimSmGlobalVrf resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfBsrCandidate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfBsrHash2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfBsrInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticast6PimSmGlobalVrfBsrPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddress2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterMulticast6PimSmGlobalVrfRpAddressGroup2edl(i["group"], d, pre_append)
			tmp["group"] = fortiAPISubPartPatch(v, "RouterMulticast6PimSmGlobalVrf-RpAddress-Group")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterMulticast6PimSmGlobalVrfRpAddressId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterMulticast6PimSmGlobalVrf-RpAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := i["ip6-address"]; ok {
			v := flattenRouterMulticast6PimSmGlobalVrfRpAddressIp6Address2edl(i["ip6-address"], d, pre_append)
			tmp["ip6_address"] = fortiAPISubPartPatch(v, "RouterMulticast6PimSmGlobalVrf-RpAddress-Ip6Address")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressIp6Address2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfVrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticast6PimSmGlobalVrf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("bsr_allow_quick_refresh", flattenRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh2edl(o["bsr-allow-quick-refresh"], d, "bsr_allow_quick_refresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-allow-quick-refresh"], "RouterMulticast6PimSmGlobalVrf-BsrAllowQuickRefresh"); ok {
			if err = d.Set("bsr_allow_quick_refresh", vv); err != nil {
				return fmt.Errorf("Error reading bsr_allow_quick_refresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_allow_quick_refresh: %v", err)
		}
	}

	if err = d.Set("bsr_candidate", flattenRouterMulticast6PimSmGlobalVrfBsrCandidate2edl(o["bsr-candidate"], d, "bsr_candidate")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-candidate"], "RouterMulticast6PimSmGlobalVrf-BsrCandidate"); ok {
			if err = d.Set("bsr_candidate", vv); err != nil {
				return fmt.Errorf("Error reading bsr_candidate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_candidate: %v", err)
		}
	}

	if err = d.Set("bsr_hash", flattenRouterMulticast6PimSmGlobalVrfBsrHash2edl(o["bsr-hash"], d, "bsr_hash")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-hash"], "RouterMulticast6PimSmGlobalVrf-BsrHash"); ok {
			if err = d.Set("bsr_hash", vv); err != nil {
				return fmt.Errorf("Error reading bsr_hash: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_hash: %v", err)
		}
	}

	if err = d.Set("bsr_interface", flattenRouterMulticast6PimSmGlobalVrfBsrInterface2edl(o["bsr-interface"], d, "bsr_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-interface"], "RouterMulticast6PimSmGlobalVrf-BsrInterface"); ok {
			if err = d.Set("bsr_interface", vv); err != nil {
				return fmt.Errorf("Error reading bsr_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_interface: %v", err)
		}
	}

	if err = d.Set("bsr_priority", flattenRouterMulticast6PimSmGlobalVrfBsrPriority2edl(o["bsr-priority"], d, "bsr_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-priority"], "RouterMulticast6PimSmGlobalVrf-BsrPriority"); ok {
			if err = d.Set("bsr_priority", vv); err != nil {
				return fmt.Errorf("Error reading bsr_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_priority: %v", err)
		}
	}

	if err = d.Set("cisco_crp_prefix", flattenRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix2edl(o["cisco-crp-prefix"], d, "cisco_crp_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["cisco-crp-prefix"], "RouterMulticast6PimSmGlobalVrf-CiscoCrpPrefix"); ok {
			if err = d.Set("cisco_crp_prefix", vv); err != nil {
				return fmt.Errorf("Error reading cisco_crp_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cisco_crp_prefix: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rp_address", flattenRouterMulticast6PimSmGlobalVrfRpAddress2edl(o["rp-address"], d, "rp_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["rp-address"], "RouterMulticast6PimSmGlobalVrf-RpAddress"); ok {
				if err = d.Set("rp_address", vv); err != nil {
					return fmt.Errorf("Error reading rp_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rp_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rp_address"); ok {
			if err = d.Set("rp_address", flattenRouterMulticast6PimSmGlobalVrfRpAddress2edl(o["rp-address"], d, "rp_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["rp-address"], "RouterMulticast6PimSmGlobalVrf-RpAddress"); ok {
					if err = d.Set("rp_address", vv); err != nil {
						return fmt.Errorf("Error reading rp_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rp_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("vrf", flattenRouterMulticast6PimSmGlobalVrfVrf2edl(o["vrf"], d, "vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf"], "RouterMulticast6PimSmGlobalVrf-Vrf"); ok {
			if err = d.Set("vrf", vv); err != nil {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticast6PimSmGlobalVrfFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrCandidate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrHash2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["group"], _ = expandRouterMulticast6PimSmGlobalVrfRpAddressGroup2edl(d, i["group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterMulticast6PimSmGlobalVrfRpAddressId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6-address"], _ = expandRouterMulticast6PimSmGlobalVrfRpAddressIp6Address2edl(d, i["ip6_address"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddressGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddressId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddressIp6Address2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfVrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticast6PimSmGlobalVrf(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bsr_allow_quick_refresh"); ok || d.HasChange("bsr_allow_quick_refresh") {
		t, err := expandRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh2edl(d, v, "bsr_allow_quick_refresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-allow-quick-refresh"] = t
		}
	}

	if v, ok := d.GetOk("bsr_candidate"); ok || d.HasChange("bsr_candidate") {
		t, err := expandRouterMulticast6PimSmGlobalVrfBsrCandidate2edl(d, v, "bsr_candidate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-candidate"] = t
		}
	}

	if v, ok := d.GetOk("bsr_hash"); ok || d.HasChange("bsr_hash") {
		t, err := expandRouterMulticast6PimSmGlobalVrfBsrHash2edl(d, v, "bsr_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-hash"] = t
		}
	}

	if v, ok := d.GetOk("bsr_interface"); ok || d.HasChange("bsr_interface") {
		t, err := expandRouterMulticast6PimSmGlobalVrfBsrInterface2edl(d, v, "bsr_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-interface"] = t
		}
	}

	if v, ok := d.GetOk("bsr_priority"); ok || d.HasChange("bsr_priority") {
		t, err := expandRouterMulticast6PimSmGlobalVrfBsrPriority2edl(d, v, "bsr_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-priority"] = t
		}
	}

	if v, ok := d.GetOk("cisco_crp_prefix"); ok || d.HasChange("cisco_crp_prefix") {
		t, err := expandRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix2edl(d, v, "cisco_crp_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cisco-crp-prefix"] = t
		}
	}

	if v, ok := d.GetOk("rp_address"); ok || d.HasChange("rp_address") {
		t, err := expandRouterMulticast6PimSmGlobalVrfRpAddress2edl(d, v, "rp_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rp-address"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		t, err := expandRouterMulticast6PimSmGlobalVrfVrf2edl(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	return &obj, nil
}
