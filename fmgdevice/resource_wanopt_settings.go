// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WAN optimization settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptSettingsUpdate,
		Read:   resourceWanoptSettingsRead,
		Update: resourceWanoptSettingsUpdate,
		Delete: resourceWanoptSettingsDelete,

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
			"auto_detect_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_optimization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWanoptSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanoptSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateWanoptSettings(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WanoptSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WanoptSettings")

	return resourceWanoptSettingsRead(d, m)
}

func resourceWanoptSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWanoptSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WanoptSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptSettings resource from API: %v", err)
	}
	return nil
}

func flattenWanoptSettingsAutoDetectAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptSettingsHostId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptSettingsTunnelOptimization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptSettingsTunnelSslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auto_detect_algorithm", flattenWanoptSettingsAutoDetectAlgorithm(o["auto-detect-algorithm"], d, "auto_detect_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-detect-algorithm"], "WanoptSettings-AutoDetectAlgorithm"); ok {
			if err = d.Set("auto_detect_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading auto_detect_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_detect_algorithm: %v", err)
		}
	}

	if err = d.Set("host_id", flattenWanoptSettingsHostId(o["host-id"], d, "host_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-id"], "WanoptSettings-HostId"); ok {
			if err = d.Set("host_id", vv); err != nil {
				return fmt.Errorf("Error reading host_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_id: %v", err)
		}
	}

	if err = d.Set("tunnel_optimization", flattenWanoptSettingsTunnelOptimization(o["tunnel-optimization"], d, "tunnel_optimization")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-optimization"], "WanoptSettings-TunnelOptimization"); ok {
			if err = d.Set("tunnel_optimization", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_optimization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_optimization: %v", err)
		}
	}

	if err = d.Set("tunnel_ssl_algorithm", flattenWanoptSettingsTunnelSslAlgorithm(o["tunnel-ssl-algorithm"], d, "tunnel_ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-ssl-algorithm"], "WanoptSettings-TunnelSslAlgorithm"); ok {
			if err = d.Set("tunnel_ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_ssl_algorithm: %v", err)
		}
	}

	return nil
}

func flattenWanoptSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptSettingsAutoDetectAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptSettingsHostId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptSettingsTunnelOptimization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptSettingsTunnelSslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_detect_algorithm"); ok || d.HasChange("auto_detect_algorithm") {
		t, err := expandWanoptSettingsAutoDetectAlgorithm(d, v, "auto_detect_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-detect-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("host_id"); ok || d.HasChange("host_id") {
		t, err := expandWanoptSettingsHostId(d, v, "host_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-id"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_optimization"); ok || d.HasChange("tunnel_optimization") {
		t, err := expandWanoptSettingsTunnelOptimization(d, v, "tunnel_optimization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-optimization"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_ssl_algorithm"); ok || d.HasChange("tunnel_ssl_algorithm") {
		t, err := expandWanoptSettingsTunnelSslAlgorithm(d, v, "tunnel_ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-ssl-algorithm"] = t
		}
	}

	return &obj, nil
}
