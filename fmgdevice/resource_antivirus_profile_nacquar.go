// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure AntiVirus quarantine settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAntivirusProfileNacQuar() *schema.Resource {
	return &schema.Resource{
		Create: resourceAntivirusProfileNacQuarUpdate,
		Read:   resourceAntivirusProfileNacQuarRead,
		Update: resourceAntivirusProfileNacQuarUpdate,
		Delete: resourceAntivirusProfileNacQuarDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"infected": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAntivirusProfileNacQuarUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectAntivirusProfileNacQuar(d)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusProfileNacQuar resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateAntivirusProfileNacQuar(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusProfileNacQuar resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("AntivirusProfileNacQuar")

	return resourceAntivirusProfileNacQuarRead(d, m)
}

func resourceAntivirusProfileNacQuarDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteAntivirusProfileNacQuar(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AntivirusProfileNacQuar resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusProfileNacQuarRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadAntivirusProfileNacQuar(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading AntivirusProfileNacQuar resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusProfileNacQuar(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusProfileNacQuar resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusProfileNacQuarExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNacQuarInfected2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNacQuarLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAntivirusProfileNacQuar(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("expiry", flattenAntivirusProfileNacQuarExpiry2edl(o["expiry"], d, "expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["expiry"], "AntivirusProfileNacQuar-Expiry"); ok {
			if err = d.Set("expiry", vv); err != nil {
				return fmt.Errorf("Error reading expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expiry: %v", err)
		}
	}

	if err = d.Set("infected", flattenAntivirusProfileNacQuarInfected2edl(o["infected"], d, "infected")); err != nil {
		if vv, ok := fortiAPIPatch(o["infected"], "AntivirusProfileNacQuar-Infected"); ok {
			if err = d.Set("infected", vv); err != nil {
				return fmt.Errorf("Error reading infected: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading infected: %v", err)
		}
	}

	if err = d.Set("log", flattenAntivirusProfileNacQuarLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "AntivirusProfileNacQuar-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	return nil
}

func flattenAntivirusProfileNacQuarFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAntivirusProfileNacQuarExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNacQuarInfected2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNacQuarLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusProfileNacQuar(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("expiry"); ok || d.HasChange("expiry") {
		t, err := expandAntivirusProfileNacQuarExpiry2edl(d, v, "expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expiry"] = t
		}
	}

	if v, ok := d.GetOk("infected"); ok || d.HasChange("infected") {
		t, err := expandAntivirusProfileNacQuarInfected2edl(d, v, "infected")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["infected"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandAntivirusProfileNacQuarLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	return &obj, nil
}
