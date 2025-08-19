// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Network overlays to register with Overlay Controller VPN service.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnOcvpnOverlays() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnOcvpnOverlaysCreate,
		Read:   resourceVpnOcvpnOverlaysRead,
		Update: resourceVpnOcvpnOverlaysUpdate,
		Delete: resourceVpnOcvpnOverlaysDelete,

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
			"assign_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inter_overlay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"overlay_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"subnet": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
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

func resourceVpnOcvpnOverlaysCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnOcvpnOverlays(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnOcvpnOverlays resource while getting object: %v", err)
	}

	_, err = c.CreateVpnOcvpnOverlays(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating VpnOcvpnOverlays resource: %v", err)
	}

	d.SetId(getStringKey(d, "overlay_name"))

	return resourceVpnOcvpnOverlaysRead(d, m)
}

func resourceVpnOcvpnOverlaysUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnOcvpnOverlays(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnOcvpnOverlays resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnOcvpnOverlays(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnOcvpnOverlays resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "overlay_name"))

	return resourceVpnOcvpnOverlaysRead(d, m)
}

func resourceVpnOcvpnOverlaysDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnOcvpnOverlays(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnOcvpnOverlays resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnOcvpnOverlaysRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnOcvpnOverlays(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnOcvpnOverlays resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnOcvpnOverlays(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnOcvpnOverlays resource from API: %v", err)
	}
	return nil
}

func flattenVpnOcvpnOverlaysAssignIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysInterOverlay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysIpv4EndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysIpv4StartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysOverlayName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysSubnets2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenVpnOcvpnOverlaysSubnetsId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "VpnOcvpnOverlays-Subnets-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenVpnOcvpnOverlaysSubnetsInterface2edl(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "VpnOcvpnOverlays-Subnets-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenVpnOcvpnOverlaysSubnetsSubnet2edl(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "VpnOcvpnOverlays-Subnets-Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenVpnOcvpnOverlaysSubnetsType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "VpnOcvpnOverlays-Subnets-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnOcvpnOverlaysSubnetsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysSubnetsInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnOcvpnOverlaysSubnetsSubnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnOcvpnOverlaysSubnetsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnOcvpnOverlays(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("assign_ip", flattenVpnOcvpnOverlaysAssignIp2edl(o["assign-ip"], d, "assign_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["assign-ip"], "VpnOcvpnOverlays-AssignIp"); ok {
			if err = d.Set("assign_ip", vv); err != nil {
				return fmt.Errorf("Error reading assign_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assign_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenVpnOcvpnOverlaysId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "VpnOcvpnOverlays-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("inter_overlay", flattenVpnOcvpnOverlaysInterOverlay2edl(o["inter-overlay"], d, "inter_overlay")); err != nil {
		if vv, ok := fortiAPIPatch(o["inter-overlay"], "VpnOcvpnOverlays-InterOverlay"); ok {
			if err = d.Set("inter_overlay", vv); err != nil {
				return fmt.Errorf("Error reading inter_overlay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inter_overlay: %v", err)
		}
	}

	if err = d.Set("ipv4_end_ip", flattenVpnOcvpnOverlaysIpv4EndIp2edl(o["ipv4-end-ip"], d, "ipv4_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-end-ip"], "VpnOcvpnOverlays-Ipv4EndIp"); ok {
			if err = d.Set("ipv4_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_end_ip: %v", err)
		}
	}

	if err = d.Set("ipv4_start_ip", flattenVpnOcvpnOverlaysIpv4StartIp2edl(o["ipv4-start-ip"], d, "ipv4_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-start-ip"], "VpnOcvpnOverlays-Ipv4StartIp"); ok {
			if err = d.Set("ipv4_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_start_ip: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnOcvpnOverlaysName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnOcvpnOverlays-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("overlay_name", flattenVpnOcvpnOverlaysOverlayName2edl(o["overlay-name"], d, "overlay_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["overlay-name"], "VpnOcvpnOverlays-OverlayName"); ok {
			if err = d.Set("overlay_name", vv); err != nil {
				return fmt.Errorf("Error reading overlay_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overlay_name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("subnets", flattenVpnOcvpnOverlaysSubnets2edl(o["subnets"], d, "subnets")); err != nil {
			if vv, ok := fortiAPIPatch(o["subnets"], "VpnOcvpnOverlays-Subnets"); ok {
				if err = d.Set("subnets", vv); err != nil {
					return fmt.Errorf("Error reading subnets: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading subnets: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("subnets"); ok {
			if err = d.Set("subnets", flattenVpnOcvpnOverlaysSubnets2edl(o["subnets"], d, "subnets")); err != nil {
				if vv, ok := fortiAPIPatch(o["subnets"], "VpnOcvpnOverlays-Subnets"); ok {
					if err = d.Set("subnets", vv); err != nil {
						return fmt.Errorf("Error reading subnets: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading subnets: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenVpnOcvpnOverlaysFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnOcvpnOverlaysAssignIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysInterOverlay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysIpv4EndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysIpv4StartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysOverlayName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysSubnets2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandVpnOcvpnOverlaysSubnetsId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandVpnOcvpnOverlaysSubnetsInterface2edl(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet"], _ = expandVpnOcvpnOverlaysSubnetsSubnet2edl(d, i["subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandVpnOcvpnOverlaysSubnetsType2edl(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnOcvpnOverlaysSubnetsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysSubnetsInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnOcvpnOverlaysSubnetsSubnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandVpnOcvpnOverlaysSubnetsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnOcvpnOverlays(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("assign_ip"); ok || d.HasChange("assign_ip") {
		t, err := expandVpnOcvpnOverlaysAssignIp2edl(d, v, "assign_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandVpnOcvpnOverlaysId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("inter_overlay"); ok || d.HasChange("inter_overlay") {
		t, err := expandVpnOcvpnOverlaysInterOverlay2edl(d, v, "inter_overlay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-overlay"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_end_ip"); ok || d.HasChange("ipv4_end_ip") {
		t, err := expandVpnOcvpnOverlaysIpv4EndIp2edl(d, v, "ipv4_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_start_ip"); ok || d.HasChange("ipv4_start_ip") {
		t, err := expandVpnOcvpnOverlaysIpv4StartIp2edl(d, v, "ipv4_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnOcvpnOverlaysName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("overlay_name"); ok || d.HasChange("overlay_name") {
		t, err := expandVpnOcvpnOverlaysOverlayName2edl(d, v, "overlay_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overlay-name"] = t
		}
	}

	if v, ok := d.GetOk("subnets"); ok || d.HasChange("subnets") {
		t, err := expandVpnOcvpnOverlaysSubnets2edl(d, v, "subnets")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnets"] = t
		}
	}

	return &obj, nil
}
