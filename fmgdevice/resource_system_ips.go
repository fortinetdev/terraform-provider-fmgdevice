// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPS system settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIps() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpsUpdate,
		Read:   resourceSystemIpsRead,
		Update: resourceSystemIpsUpdate,
		Delete: resourceSystemIpsDelete,

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
			"override_signature_hold_by_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature_hold_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemIpsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemIps(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemIps resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIps(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemIps resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemIps")

	return resourceSystemIpsRead(d, m)
}

func resourceSystemIpsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemIps(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIps resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemIps(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemIps resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIps(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemIps resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpsOverrideSignatureHoldById(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpsSignatureHoldTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemIps(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("override_signature_hold_by_id", flattenSystemIpsOverrideSignatureHoldById(o["override-signature-hold-by-id"], d, "override_signature_hold_by_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-signature-hold-by-id"], "SystemIps-OverrideSignatureHoldById"); ok {
			if err = d.Set("override_signature_hold_by_id", vv); err != nil {
				return fmt.Errorf("Error reading override_signature_hold_by_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_signature_hold_by_id: %v", err)
		}
	}

	if err = d.Set("signature_hold_time", flattenSystemIpsSignatureHoldTime(o["signature-hold-time"], d, "signature_hold_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["signature-hold-time"], "SystemIps-SignatureHoldTime"); ok {
			if err = d.Set("signature_hold_time", vv); err != nil {
				return fmt.Errorf("Error reading signature_hold_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signature_hold_time: %v", err)
		}
	}

	return nil
}

func flattenSystemIpsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemIpsOverrideSignatureHoldById(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpsSignatureHoldTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIps(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("override_signature_hold_by_id"); ok || d.HasChange("override_signature_hold_by_id") {
		t, err := expandSystemIpsOverrideSignatureHoldById(d, v, "override_signature_hold_by_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-signature-hold-by-id"] = t
		}
	}

	if v, ok := d.GetOk("signature_hold_time"); ok || d.HasChange("signature_hold_time") {
		t, err := expandSystemIpsSignatureHoldTime(d, v, "signature_hold_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature-hold-time"] = t
		}
	}

	return &obj, nil
}
