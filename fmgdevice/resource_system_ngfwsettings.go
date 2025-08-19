// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPS NGFW policy-mode VDOM settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNgfwSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNgfwSettingsUpdate,
		Read:   resourceSystemNgfwSettingsRead,
		Update: resourceSystemNgfwSettingsUpdate,
		Delete: resourceSystemNgfwSettingsDelete,

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
			"match_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_halfopen_match_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_match_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemNgfwSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemNgfwSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNgfwSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNgfwSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemNgfwSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemNgfwSettings")

	return resourceSystemNgfwSettingsRead(d, m)
}

func resourceSystemNgfwSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemNgfwSettings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNgfwSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNgfwSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemNgfwSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemNgfwSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNgfwSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNgfwSettings resource from API: %v", err)
	}
	return nil
}

func flattenSystemNgfwSettingsMatchTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNgfwSettingsTcpHalfopenMatchTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNgfwSettingsTcpMatchTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNgfwSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("match_timeout", flattenSystemNgfwSettingsMatchTimeout(o["match-timeout"], d, "match_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-timeout"], "SystemNgfwSettings-MatchTimeout"); ok {
			if err = d.Set("match_timeout", vv); err != nil {
				return fmt.Errorf("Error reading match_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_timeout: %v", err)
		}
	}

	if err = d.Set("tcp_halfopen_match_timeout", flattenSystemNgfwSettingsTcpHalfopenMatchTimeout(o["tcp-halfopen-match-timeout"], d, "tcp_halfopen_match_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-halfopen-match-timeout"], "SystemNgfwSettings-TcpHalfopenMatchTimeout"); ok {
			if err = d.Set("tcp_halfopen_match_timeout", vv); err != nil {
				return fmt.Errorf("Error reading tcp_halfopen_match_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_halfopen_match_timeout: %v", err)
		}
	}

	if err = d.Set("tcp_match_timeout", flattenSystemNgfwSettingsTcpMatchTimeout(o["tcp-match-timeout"], d, "tcp_match_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-match-timeout"], "SystemNgfwSettings-TcpMatchTimeout"); ok {
			if err = d.Set("tcp_match_timeout", vv); err != nil {
				return fmt.Errorf("Error reading tcp_match_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_match_timeout: %v", err)
		}
	}

	return nil
}

func flattenSystemNgfwSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNgfwSettingsMatchTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNgfwSettingsTcpHalfopenMatchTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNgfwSettingsTcpMatchTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNgfwSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("match_timeout"); ok || d.HasChange("match_timeout") {
		t, err := expandSystemNgfwSettingsMatchTimeout(d, v, "match_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-timeout"] = t
		}
	}

	if v, ok := d.GetOk("tcp_halfopen_match_timeout"); ok || d.HasChange("tcp_halfopen_match_timeout") {
		t, err := expandSystemNgfwSettingsTcpHalfopenMatchTimeout(d, v, "tcp_halfopen_match_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-halfopen-match-timeout"] = t
		}
	}

	if v, ok := d.GetOk("tcp_match_timeout"); ok || d.HasChange("tcp_match_timeout") {
		t, err := expandSystemNgfwSettingsTcpMatchTimeout(d, v, "tcp_match_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-match-timeout"] = t
		}
	}

	return &obj, nil
}
