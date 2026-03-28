// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender controller report configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderProfileCellularControllerReport() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderProfileCellularControllerReportUpdate,
		Read:   resourceExtensionControllerExtenderProfileCellularControllerReportRead,
		Update: resourceExtensionControllerExtenderProfileCellularControllerReportUpdate,
		Delete: resourceExtensionControllerExtenderProfileCellularControllerReportDelete,

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
			"extender_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"signal_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceExtensionControllerExtenderProfileCellularControllerReportUpdate(d *schema.ResourceData, m interface{}) error {
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
	extender_profile := d.Get("extender_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectExtensionControllerExtenderProfileCellularControllerReport(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellularControllerReport resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateExtensionControllerExtenderProfileCellularControllerReport(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellularControllerReport resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ExtensionControllerExtenderProfileCellularControllerReport")

	return resourceExtensionControllerExtenderProfileCellularControllerReportRead(d, m)
}

func resourceExtensionControllerExtenderProfileCellularControllerReportDelete(d *schema.ResourceData, m interface{}) error {
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
	extender_profile := d.Get("extender_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	wsParams["adom"] = adomv

	err = c.DeleteExtensionControllerExtenderProfileCellularControllerReport(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderProfileCellularControllerReport resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderProfileCellularControllerReportRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	extender_profile := d.Get("extender_profile").(string)
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
	if extender_profile == "" {
		extender_profile = importOptionChecking(m.(*FortiClient).Cfg, "extender_profile")
		if extender_profile == "" {
			return fmt.Errorf("Parameter extender_profile is missing")
		}
		if err = d.Set("extender_profile", extender_profile); err != nil {
			return fmt.Errorf("Error set params extender_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	o, err := c.ReadExtensionControllerExtenderProfileCellularControllerReport(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileCellularControllerReport resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderProfileCellularControllerReport(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileCellularControllerReport resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderProfileCellularControllerReportInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularControllerReportSignalThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularControllerReportStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtenderProfileCellularControllerReport(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("interval", flattenExtensionControllerExtenderProfileCellularControllerReportInterval3rdl(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "ExtensionControllerExtenderProfileCellularControllerReport-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("signal_threshold", flattenExtensionControllerExtenderProfileCellularControllerReportSignalThreshold3rdl(o["signal-threshold"], d, "signal_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["signal-threshold"], "ExtensionControllerExtenderProfileCellularControllerReport-SignalThreshold"); ok {
			if err = d.Set("signal_threshold", vv); err != nil {
				return fmt.Errorf("Error reading signal_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signal_threshold: %v", err)
		}
	}

	if err = d.Set("status", flattenExtensionControllerExtenderProfileCellularControllerReportStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ExtensionControllerExtenderProfileCellularControllerReport-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerExtenderProfileCellularControllerReportFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerExtenderProfileCellularControllerReportInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReportSignalThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReportStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtenderProfileCellularControllerReport(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interval"); ok || d.HasChange("interval") {
		t, err := expandExtensionControllerExtenderProfileCellularControllerReportInterval3rdl(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("signal_threshold"); ok || d.HasChange("signal_threshold") {
		t, err := expandExtensionControllerExtenderProfileCellularControllerReportSignalThreshold3rdl(d, v, "signal_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-threshold"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandExtensionControllerExtenderProfileCellularControllerReportStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
