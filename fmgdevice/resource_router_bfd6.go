// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 BFD.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBfd6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBfd6Update,
		Read:   resourceRouterBfd6Read,
		Update: resourceRouterBfd6Update,
		Delete: resourceRouterBfd6Delete,

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
			"multihop_template": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bfd_desired_min_tx": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bfd_detect_mult": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bfd_required_min_rx": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dst": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"md5_key": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"src": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceRouterBfd6Update(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBfd6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBfd6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBfd6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterBfd6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterBfd6")

	return resourceRouterBfd6Read(d, m)
}

func resourceRouterBfd6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterBfd6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBfd6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBfd6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBfd6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBfd6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBfd6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBfd6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterBfd6MultihopTemplate(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := i["auth-mode"]; ok {
			v := flattenRouterBfd6MultihopTemplateAuthMode(i["auth-mode"], d, pre_append)
			tmp["auth_mode"] = fortiAPISubPartPatch(v, "RouterBfd6-MultihopTemplate-AuthMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_desired_min_tx"
		if _, ok := i["bfd-desired-min-tx"]; ok {
			v := flattenRouterBfd6MultihopTemplateBfdDesiredMinTx(i["bfd-desired-min-tx"], d, pre_append)
			tmp["bfd_desired_min_tx"] = fortiAPISubPartPatch(v, "RouterBfd6-MultihopTemplate-BfdDesiredMinTx")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_detect_mult"
		if _, ok := i["bfd-detect-mult"]; ok {
			v := flattenRouterBfd6MultihopTemplateBfdDetectMult(i["bfd-detect-mult"], d, pre_append)
			tmp["bfd_detect_mult"] = fortiAPISubPartPatch(v, "RouterBfd6-MultihopTemplate-BfdDetectMult")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_required_min_rx"
		if _, ok := i["bfd-required-min-rx"]; ok {
			v := flattenRouterBfd6MultihopTemplateBfdRequiredMinRx(i["bfd-required-min-rx"], d, pre_append)
			tmp["bfd_required_min_rx"] = fortiAPISubPartPatch(v, "RouterBfd6-MultihopTemplate-BfdRequiredMinRx")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			v := flattenRouterBfd6MultihopTemplateDst(i["dst"], d, pre_append)
			tmp["dst"] = fortiAPISubPartPatch(v, "RouterBfd6-MultihopTemplate-Dst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterBfd6MultihopTemplateId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterBfd6-MultihopTemplate-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := i["src"]; ok {
			v := flattenRouterBfd6MultihopTemplateSrc(i["src"], d, pre_append)
			tmp["src"] = fortiAPISubPartPatch(v, "RouterBfd6-MultihopTemplate-Src")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBfd6MultihopTemplateAuthMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateBfdDesiredMinTx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateBfdDetectMult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateBfdRequiredMinRx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6Neighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenRouterBfd6NeighborInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterBfd6-Neighbor-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := i["ip6-address"]; ok {
			v := flattenRouterBfd6NeighborIp6Address(i["ip6-address"], d, pre_append)
			tmp["ip6_address"] = fortiAPISubPartPatch(v, "RouterBfd6-Neighbor-Ip6Address")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBfd6NeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBfd6NeighborIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterBfd6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("multihop_template", flattenRouterBfd6MultihopTemplate(o["multihop-template"], d, "multihop_template")); err != nil {
			if vv, ok := fortiAPIPatch(o["multihop-template"], "RouterBfd6-MultihopTemplate"); ok {
				if err = d.Set("multihop_template", vv); err != nil {
					return fmt.Errorf("Error reading multihop_template: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading multihop_template: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("multihop_template"); ok {
			if err = d.Set("multihop_template", flattenRouterBfd6MultihopTemplate(o["multihop-template"], d, "multihop_template")); err != nil {
				if vv, ok := fortiAPIPatch(o["multihop-template"], "RouterBfd6-MultihopTemplate"); ok {
					if err = d.Set("multihop_template", vv); err != nil {
						return fmt.Errorf("Error reading multihop_template: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading multihop_template: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor", flattenRouterBfd6Neighbor(o["neighbor"], d, "neighbor")); err != nil {
			if vv, ok := fortiAPIPatch(o["neighbor"], "RouterBfd6-Neighbor"); ok {
				if err = d.Set("neighbor", vv); err != nil {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterBfd6Neighbor(o["neighbor"], d, "neighbor")); err != nil {
				if vv, ok := fortiAPIPatch(o["neighbor"], "RouterBfd6-Neighbor"); ok {
					if err = d.Set("neighbor", vv); err != nil {
						return fmt.Errorf("Error reading neighbor: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterBfd6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBfd6MultihopTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-mode"], _ = expandRouterBfd6MultihopTemplateAuthMode(d, i["auth_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_desired_min_tx"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bfd-desired-min-tx"], _ = expandRouterBfd6MultihopTemplateBfdDesiredMinTx(d, i["bfd_desired_min_tx"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_detect_mult"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bfd-detect-mult"], _ = expandRouterBfd6MultihopTemplateBfdDetectMult(d, i["bfd_detect_mult"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_required_min_rx"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bfd-required-min-rx"], _ = expandRouterBfd6MultihopTemplateBfdRequiredMinRx(d, i["bfd_required_min_rx"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst"], _ = expandRouterBfd6MultihopTemplateDst(d, i["dst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterBfd6MultihopTemplateId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["md5-key"], _ = expandRouterBfd6MultihopTemplateMd5Key(d, i["md5_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src"], _ = expandRouterBfd6MultihopTemplateSrc(d, i["src"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBfd6MultihopTemplateAuthMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateBfdDesiredMinTx(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateBfdDetectMult(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateBfdRequiredMinRx(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateMd5Key(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBfd6MultihopTemplateSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6Neighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandRouterBfd6NeighborInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6-address"], _ = expandRouterBfd6NeighborIp6Address(d, i["ip6_address"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBfd6NeighborInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBfd6NeighborIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBfd6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("multihop_template"); ok || d.HasChange("multihop_template") {
		t, err := expandRouterBfd6MultihopTemplate(d, v, "multihop_template")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multihop-template"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		t, err := expandRouterBfd6Neighbor(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	return &obj, nil
}
