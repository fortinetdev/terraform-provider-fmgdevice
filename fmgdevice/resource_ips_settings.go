// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPS VDOM parameter.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpsSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsSettingsUpdate,
		Read:   resourceIpsSettingsRead,
		Update: resourceIpsSettingsUpdate,
		Delete: resourceIpsSettingsDelete,

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
			"ha_session_pickup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_packet_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"packet_log_history": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"packet_log_memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"packet_log_post_attack": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"proxy_inline_ips": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIpsSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectIpsSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating IpsSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateIpsSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IpsSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("IpsSettings")

	return resourceIpsSettingsRead(d, m)
}

func resourceIpsSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteIpsSettings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IpsSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadIpsSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading IpsSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IpsSettings resource from API: %v", err)
	}
	return nil
}

func flattenIpsSettingsHaSessionPickup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSettingsIpsPacketQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSettingsPacketLogHistory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSettingsPacketLogMemory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSettingsPacketLogPostAttack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSettingsProxyInlineIps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIpsSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ha_session_pickup", flattenIpsSettingsHaSessionPickup(o["ha-session-pickup"], d, "ha_session_pickup")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-session-pickup"], "IpsSettings-HaSessionPickup"); ok {
			if err = d.Set("ha_session_pickup", vv); err != nil {
				return fmt.Errorf("Error reading ha_session_pickup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_session_pickup: %v", err)
		}
	}

	if err = d.Set("ips_packet_quota", flattenIpsSettingsIpsPacketQuota(o["ips-packet-quota"], d, "ips_packet_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-packet-quota"], "IpsSettings-IpsPacketQuota"); ok {
			if err = d.Set("ips_packet_quota", vv); err != nil {
				return fmt.Errorf("Error reading ips_packet_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_packet_quota: %v", err)
		}
	}

	if err = d.Set("packet_log_history", flattenIpsSettingsPacketLogHistory(o["packet-log-history"], d, "packet_log_history")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-log-history"], "IpsSettings-PacketLogHistory"); ok {
			if err = d.Set("packet_log_history", vv); err != nil {
				return fmt.Errorf("Error reading packet_log_history: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_log_history: %v", err)
		}
	}

	if err = d.Set("packet_log_memory", flattenIpsSettingsPacketLogMemory(o["packet-log-memory"], d, "packet_log_memory")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-log-memory"], "IpsSettings-PacketLogMemory"); ok {
			if err = d.Set("packet_log_memory", vv); err != nil {
				return fmt.Errorf("Error reading packet_log_memory: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_log_memory: %v", err)
		}
	}

	if err = d.Set("packet_log_post_attack", flattenIpsSettingsPacketLogPostAttack(o["packet-log-post-attack"], d, "packet_log_post_attack")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-log-post-attack"], "IpsSettings-PacketLogPostAttack"); ok {
			if err = d.Set("packet_log_post_attack", vv); err != nil {
				return fmt.Errorf("Error reading packet_log_post_attack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_log_post_attack: %v", err)
		}
	}

	if err = d.Set("proxy_inline_ips", flattenIpsSettingsProxyInlineIps(o["proxy-inline-ips"], d, "proxy_inline_ips")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-inline-ips"], "IpsSettings-ProxyInlineIps"); ok {
			if err = d.Set("proxy_inline_ips", vv); err != nil {
				return fmt.Errorf("Error reading proxy_inline_ips: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_inline_ips: %v", err)
		}
	}

	return nil
}

func flattenIpsSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIpsSettingsHaSessionPickup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSettingsIpsPacketQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSettingsPacketLogHistory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSettingsPacketLogMemory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSettingsPacketLogPostAttack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSettingsProxyInlineIps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIpsSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ha_session_pickup"); ok || d.HasChange("ha_session_pickup") {
		t, err := expandIpsSettingsHaSessionPickup(d, v, "ha_session_pickup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-session-pickup"] = t
		}
	}

	if v, ok := d.GetOk("ips_packet_quota"); ok || d.HasChange("ips_packet_quota") {
		t, err := expandIpsSettingsIpsPacketQuota(d, v, "ips_packet_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-packet-quota"] = t
		}
	}

	if v, ok := d.GetOk("packet_log_history"); ok || d.HasChange("packet_log_history") {
		t, err := expandIpsSettingsPacketLogHistory(d, v, "packet_log_history")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-log-history"] = t
		}
	}

	if v, ok := d.GetOk("packet_log_memory"); ok || d.HasChange("packet_log_memory") {
		t, err := expandIpsSettingsPacketLogMemory(d, v, "packet_log_memory")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-log-memory"] = t
		}
	}

	if v, ok := d.GetOk("packet_log_post_attack"); ok || d.HasChange("packet_log_post_attack") {
		t, err := expandIpsSettingsPacketLogPostAttack(d, v, "packet_log_post_attack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-log-post-attack"] = t
		}
	}

	if v, ok := d.GetOk("proxy_inline_ips"); ok || d.HasChange("proxy_inline_ips") {
		t, err := expandIpsSettingsProxyInlineIps(d, v, "proxy_inline_ips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-inline-ips"] = t
		}
	}

	return &obj, nil
}
