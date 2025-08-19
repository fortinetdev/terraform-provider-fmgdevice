// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: BGP IPv6 VRF leaking table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgpVrf6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpVrf6Create,
		Read:   resourceRouterBgpVrf6Read,
		Update: resourceRouterBgpVrf6Update,
		Delete: resourceRouterBgpVrf6Delete,

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
			"export_rt": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"import_route_map": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"import_rt": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"leak_target": &schema.Schema{
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
						"route_map": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vrf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"rd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf": &schema.Schema{
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

func resourceRouterBgpVrf6Create(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBgpVrf6(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpVrf6 resource while getting object: %v", err)
	}

	_, err = c.CreateRouterBgpVrf6(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpVrf6 resource: %v", err)
	}

	d.SetId(getStringKey(d, "vrf"))

	return resourceRouterBgpVrf6Read(d, m)
}

func resourceRouterBgpVrf6Update(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterBgpVrf6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpVrf6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpVrf6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpVrf6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "vrf"))

	return resourceRouterBgpVrf6Read(d, m)
}

func resourceRouterBgpVrf6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterBgpVrf6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpVrf6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpVrf6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBgpVrf6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpVrf6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpVrf6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpVrf6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpVrf6ExportRt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrf6ImportRouteMap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrf6ImportRt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrf6LeakTarget2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterBgpVrf6LeakTargetInterface2edl(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterBgpVrf6-LeakTarget-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			v := flattenRouterBgpVrf6LeakTargetRouteMap2edl(i["route-map"], d, pre_append)
			tmp["route_map"] = fortiAPISubPartPatch(v, "RouterBgpVrf6-LeakTarget-RouteMap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenRouterBgpVrf6LeakTargetVrf2edl(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "RouterBgpVrf6-LeakTarget-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterBgpVrf6LeakTargetInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrf6LeakTargetRouteMap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpVrf6LeakTargetVrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrf6Rd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrf6Role2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpVrf6Vrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterBgpVrf6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("export_rt", flattenRouterBgpVrf6ExportRt2edl(o["export-rt"], d, "export_rt")); err != nil {
		if vv, ok := fortiAPIPatch(o["export-rt"], "RouterBgpVrf6-ExportRt"); ok {
			if err = d.Set("export_rt", vv); err != nil {
				return fmt.Errorf("Error reading export_rt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading export_rt: %v", err)
		}
	}

	if err = d.Set("import_route_map", flattenRouterBgpVrf6ImportRouteMap2edl(o["import-route-map"], d, "import_route_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["import-route-map"], "RouterBgpVrf6-ImportRouteMap"); ok {
			if err = d.Set("import_route_map", vv); err != nil {
				return fmt.Errorf("Error reading import_route_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading import_route_map: %v", err)
		}
	}

	if err = d.Set("import_rt", flattenRouterBgpVrf6ImportRt2edl(o["import-rt"], d, "import_rt")); err != nil {
		if vv, ok := fortiAPIPatch(o["import-rt"], "RouterBgpVrf6-ImportRt"); ok {
			if err = d.Set("import_rt", vv); err != nil {
				return fmt.Errorf("Error reading import_rt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading import_rt: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("leak_target", flattenRouterBgpVrf6LeakTarget2edl(o["leak-target"], d, "leak_target")); err != nil {
			if vv, ok := fortiAPIPatch(o["leak-target"], "RouterBgpVrf6-LeakTarget"); ok {
				if err = d.Set("leak_target", vv); err != nil {
					return fmt.Errorf("Error reading leak_target: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading leak_target: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("leak_target"); ok {
			if err = d.Set("leak_target", flattenRouterBgpVrf6LeakTarget2edl(o["leak-target"], d, "leak_target")); err != nil {
				if vv, ok := fortiAPIPatch(o["leak-target"], "RouterBgpVrf6-LeakTarget"); ok {
					if err = d.Set("leak_target", vv); err != nil {
						return fmt.Errorf("Error reading leak_target: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading leak_target: %v", err)
				}
			}
		}
	}

	if err = d.Set("rd", flattenRouterBgpVrf6Rd2edl(o["rd"], d, "rd")); err != nil {
		if vv, ok := fortiAPIPatch(o["rd"], "RouterBgpVrf6-Rd"); ok {
			if err = d.Set("rd", vv); err != nil {
				return fmt.Errorf("Error reading rd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rd: %v", err)
		}
	}

	if err = d.Set("role", flattenRouterBgpVrf6Role2edl(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "RouterBgpVrf6-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("vrf", flattenRouterBgpVrf6Vrf2edl(o["vrf"], d, "vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf"], "RouterBgpVrf6-Vrf"); ok {
			if err = d.Set("vrf", vv); err != nil {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	return nil
}

func flattenRouterBgpVrf6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBgpVrf6ExportRt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrf6ImportRouteMap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrf6ImportRt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrf6LeakTarget2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["interface"], _ = expandRouterBgpVrf6LeakTargetInterface2edl(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-map"], _ = expandRouterBgpVrf6LeakTargetRouteMap2edl(d, i["route_map"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandRouterBgpVrf6LeakTargetVrf2edl(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrf6LeakTargetInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrf6LeakTargetRouteMap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBgpVrf6LeakTargetVrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6Rd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6Role2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6Vrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgpVrf6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("export_rt"); ok || d.HasChange("export_rt") {
		t, err := expandRouterBgpVrf6ExportRt2edl(d, v, "export_rt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["export-rt"] = t
		}
	}

	if v, ok := d.GetOk("import_route_map"); ok || d.HasChange("import_route_map") {
		t, err := expandRouterBgpVrf6ImportRouteMap2edl(d, v, "import_route_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["import-route-map"] = t
		}
	}

	if v, ok := d.GetOk("import_rt"); ok || d.HasChange("import_rt") {
		t, err := expandRouterBgpVrf6ImportRt2edl(d, v, "import_rt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["import-rt"] = t
		}
	}

	if v, ok := d.GetOk("leak_target"); ok || d.HasChange("leak_target") {
		t, err := expandRouterBgpVrf6LeakTarget2edl(d, v, "leak_target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["leak-target"] = t
		}
	}

	if v, ok := d.GetOk("rd"); ok || d.HasChange("rd") {
		t, err := expandRouterBgpVrf6Rd2edl(d, v, "rd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rd"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandRouterBgpVrf6Role2edl(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		t, err := expandRouterBgpVrf6Vrf2edl(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	return &obj, nil
}
