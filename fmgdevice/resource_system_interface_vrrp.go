// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: VRRP configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemInterfaceVrrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceVrrpCreate,
		Read:   resourceSystemInterfaceVrrpRead,
		Update: resourceSystemInterfaceVrrpUpdate,
		Delete: resourceSystemInterfaceVrrpDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accept_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adv_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ignore_default_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preempt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxy_arp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrdst": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vrdst_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrgrp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"vrip": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceSystemInterfaceVrrpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	obj, err := getObjectSystemInterfaceVrrp(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceVrrp resource while getting object: %v", err)
	}

	_, err = c.CreateSystemInterfaceVrrp(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceVrrp resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "vrid")))

	return resourceSystemInterfaceVrrpRead(d, m)
}

func resourceSystemInterfaceVrrpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	obj, err := getObjectSystemInterfaceVrrp(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceVrrp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceVrrp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceVrrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "vrid")))

	return resourceSystemInterfaceVrrpRead(d, m)
}

func resourceSystemInterfaceVrrpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	err = c.DeleteSystemInterfaceVrrp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceVrrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceVrrpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	var_interface := d.Get("interface").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if var_interface == "" {
		var_interface = importOptionChecking(m.(*FortiClient).Cfg, "interface")
		if var_interface == "" {
			return fmt.Errorf("Parameter interface is missing")
		}
		if err = d.Set("interface", var_interface); err != nil {
			return fmt.Errorf("Error set params interface: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	o, err := c.ReadSystemInterfaceVrrp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceVrrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceVrrp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceVrrp resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceVrrpAcceptMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpAdvInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpIgnoreDefaultRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpPreempt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpProxyArp2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemInterfaceVrrpProxyArpId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemInterfaceVrrp-ProxyArp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemInterfaceVrrpProxyArpIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemInterfaceVrrp-ProxyArp-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemInterfaceVrrpProxyArpId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpProxyArpIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpStartTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpVrdst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceVrrpVrdstPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpVrgrp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpVrid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpVrip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterfaceVrrp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("accept_mode", flattenSystemInterfaceVrrpAcceptMode2edl(o["accept-mode"], d, "accept_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["accept-mode"], "SystemInterfaceVrrp-AcceptMode"); ok {
			if err = d.Set("accept_mode", vv); err != nil {
				return fmt.Errorf("Error reading accept_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accept_mode: %v", err)
		}
	}

	if err = d.Set("adv_interval", flattenSystemInterfaceVrrpAdvInterval2edl(o["adv-interval"], d, "adv_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-interval"], "SystemInterfaceVrrp-AdvInterval"); ok {
			if err = d.Set("adv_interval", vv); err != nil {
				return fmt.Errorf("Error reading adv_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_interval: %v", err)
		}
	}

	if err = d.Set("ignore_default_route", flattenSystemInterfaceVrrpIgnoreDefaultRoute2edl(o["ignore-default-route"], d, "ignore_default_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-default-route"], "SystemInterfaceVrrp-IgnoreDefaultRoute"); ok {
			if err = d.Set("ignore_default_route", vv); err != nil {
				return fmt.Errorf("Error reading ignore_default_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_default_route: %v", err)
		}
	}

	if err = d.Set("preempt", flattenSystemInterfaceVrrpPreempt2edl(o["preempt"], d, "preempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["preempt"], "SystemInterfaceVrrp-Preempt"); ok {
			if err = d.Set("preempt", vv); err != nil {
				return fmt.Errorf("Error reading preempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preempt: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemInterfaceVrrpPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemInterfaceVrrp-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("proxy_arp", flattenSystemInterfaceVrrpProxyArp2edl(o["proxy-arp"], d, "proxy_arp")); err != nil {
			if vv, ok := fortiAPIPatch(o["proxy-arp"], "SystemInterfaceVrrp-ProxyArp"); ok {
				if err = d.Set("proxy_arp", vv); err != nil {
					return fmt.Errorf("Error reading proxy_arp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading proxy_arp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("proxy_arp"); ok {
			if err = d.Set("proxy_arp", flattenSystemInterfaceVrrpProxyArp2edl(o["proxy-arp"], d, "proxy_arp")); err != nil {
				if vv, ok := fortiAPIPatch(o["proxy-arp"], "SystemInterfaceVrrp-ProxyArp"); ok {
					if err = d.Set("proxy_arp", vv); err != nil {
						return fmt.Errorf("Error reading proxy_arp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading proxy_arp: %v", err)
				}
			}
		}
	}

	if err = d.Set("start_time", flattenSystemInterfaceVrrpStartTime2edl(o["start-time"], d, "start_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-time"], "SystemInterfaceVrrp-StartTime"); ok {
			if err = d.Set("start_time", vv); err != nil {
				return fmt.Errorf("Error reading start_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_time: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemInterfaceVrrpStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemInterfaceVrrp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("version", flattenSystemInterfaceVrrpVersion2edl(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "SystemInterfaceVrrp-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if err = d.Set("vrdst", flattenSystemInterfaceVrrpVrdst2edl(o["vrdst"], d, "vrdst")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrdst"], "SystemInterfaceVrrp-Vrdst"); ok {
			if err = d.Set("vrdst", vv); err != nil {
				return fmt.Errorf("Error reading vrdst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrdst: %v", err)
		}
	}

	if err = d.Set("vrdst_priority", flattenSystemInterfaceVrrpVrdstPriority2edl(o["vrdst-priority"], d, "vrdst_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrdst-priority"], "SystemInterfaceVrrp-VrdstPriority"); ok {
			if err = d.Set("vrdst_priority", vv); err != nil {
				return fmt.Errorf("Error reading vrdst_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrdst_priority: %v", err)
		}
	}

	if err = d.Set("vrgrp", flattenSystemInterfaceVrrpVrgrp2edl(o["vrgrp"], d, "vrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrgrp"], "SystemInterfaceVrrp-Vrgrp"); ok {
			if err = d.Set("vrgrp", vv); err != nil {
				return fmt.Errorf("Error reading vrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrgrp: %v", err)
		}
	}

	if err = d.Set("vrid", flattenSystemInterfaceVrrpVrid2edl(o["vrid"], d, "vrid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrid"], "SystemInterfaceVrrp-Vrid"); ok {
			if err = d.Set("vrid", vv); err != nil {
				return fmt.Errorf("Error reading vrid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrid: %v", err)
		}
	}

	if err = d.Set("vrip", flattenSystemInterfaceVrrpVrip2edl(o["vrip"], d, "vrip")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrip"], "SystemInterfaceVrrp-Vrip"); ok {
			if err = d.Set("vrip", vv); err != nil {
				return fmt.Errorf("Error reading vrip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrip: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfaceVrrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceVrrpAcceptMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpAdvInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpIgnoreDefaultRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpPreempt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpProxyArp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemInterfaceVrrpProxyArpId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemInterfaceVrrpProxyArpIp2edl(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceVrrpProxyArpId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpProxyArpIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpStartTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpVrdst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceVrrpVrdstPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpVrgrp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpVrid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpVrip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceVrrp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accept_mode"); ok || d.HasChange("accept_mode") {
		t, err := expandSystemInterfaceVrrpAcceptMode2edl(d, v, "accept_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accept-mode"] = t
		}
	}

	if v, ok := d.GetOk("adv_interval"); ok || d.HasChange("adv_interval") {
		t, err := expandSystemInterfaceVrrpAdvInterval2edl(d, v, "adv_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-interval"] = t
		}
	}

	if v, ok := d.GetOk("ignore_default_route"); ok || d.HasChange("ignore_default_route") {
		t, err := expandSystemInterfaceVrrpIgnoreDefaultRoute2edl(d, v, "ignore_default_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-default-route"] = t
		}
	}

	if v, ok := d.GetOk("preempt"); ok || d.HasChange("preempt") {
		t, err := expandSystemInterfaceVrrpPreempt2edl(d, v, "preempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preempt"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemInterfaceVrrpPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("proxy_arp"); ok || d.HasChange("proxy_arp") {
		t, err := expandSystemInterfaceVrrpProxyArp2edl(d, v, "proxy_arp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-arp"] = t
		}
	}

	if v, ok := d.GetOk("start_time"); ok || d.HasChange("start_time") {
		t, err := expandSystemInterfaceVrrpStartTime2edl(d, v, "start_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-time"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemInterfaceVrrpStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandSystemInterfaceVrrpVersion2edl(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	if v, ok := d.GetOk("vrdst"); ok || d.HasChange("vrdst") {
		t, err := expandSystemInterfaceVrrpVrdst2edl(d, v, "vrdst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrdst"] = t
		}
	}

	if v, ok := d.GetOk("vrdst_priority"); ok || d.HasChange("vrdst_priority") {
		t, err := expandSystemInterfaceVrrpVrdstPriority2edl(d, v, "vrdst_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrdst-priority"] = t
		}
	}

	if v, ok := d.GetOk("vrgrp"); ok || d.HasChange("vrgrp") {
		t, err := expandSystemInterfaceVrrpVrgrp2edl(d, v, "vrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrgrp"] = t
		}
	}

	if v, ok := d.GetOk("vrid"); ok || d.HasChange("vrid") {
		t, err := expandSystemInterfaceVrrpVrid2edl(d, v, "vrid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrid"] = t
		}
	}

	if v, ok := d.GetOk("vrip"); ok || d.HasChange("vrip") {
		t, err := expandSystemInterfaceVrrpVrip2edl(d, v, "vrip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrip"] = t
		}
	}

	return &obj, nil
}
