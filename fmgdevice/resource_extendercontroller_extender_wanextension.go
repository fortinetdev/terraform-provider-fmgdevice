// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device ExtenderControllerExtenderWanExtension

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtenderControllerExtenderWanExtension() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtenderControllerExtenderWanExtensionUpdate,
		Read:   resourceExtenderControllerExtenderWanExtensionRead,
		Update: resourceExtenderControllerExtenderWanExtensionUpdate,
		Delete: resourceExtenderControllerExtenderWanExtensionDelete,

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
			"extender": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"modem1_extension": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"modem2_extension": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceExtenderControllerExtenderWanExtensionUpdate(d *schema.ResourceData, m interface{}) error {
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
	extender := d.Get("extender").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender"] = extender

	obj, err := getObjectExtenderControllerExtenderWanExtension(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerExtenderWanExtension resource while getting object: %v", err)
	}

	_, err = c.UpdateExtenderControllerExtenderWanExtension(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerExtenderWanExtension resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ExtenderControllerExtenderWanExtension")

	return resourceExtenderControllerExtenderWanExtensionRead(d, m)
}

func resourceExtenderControllerExtenderWanExtensionDelete(d *schema.ResourceData, m interface{}) error {
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
	extender := d.Get("extender").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender"] = extender

	err = c.DeleteExtenderControllerExtenderWanExtension(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ExtenderControllerExtenderWanExtension resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtenderControllerExtenderWanExtensionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	extender := d.Get("extender").(string)
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
	if extender == "" {
		extender = importOptionChecking(m.(*FortiClient).Cfg, "extender")
		if extender == "" {
			return fmt.Errorf("Parameter extender is missing")
		}
		if err = d.Set("extender", extender); err != nil {
			return fmt.Errorf("Error set params extender: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender"] = extender

	o, err := c.ReadExtenderControllerExtenderWanExtension(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerExtenderWanExtension resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtenderControllerExtenderWanExtension(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerExtenderWanExtension resource from API: %v", err)
	}
	return nil
}

func flattenExtenderControllerExtenderWanExtensionModem1Extension2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderWanExtensionModem2Extension2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectExtenderControllerExtenderWanExtension(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("modem1_extension", flattenExtenderControllerExtenderWanExtensionModem1Extension2edl(o["modem1-extension"], d, "modem1_extension")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem1-extension"], "ExtenderControllerExtenderWanExtension-Modem1Extension"); ok {
			if err = d.Set("modem1_extension", vv); err != nil {
				return fmt.Errorf("Error reading modem1_extension: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem1_extension: %v", err)
		}
	}

	if err = d.Set("modem2_extension", flattenExtenderControllerExtenderWanExtensionModem2Extension2edl(o["modem2-extension"], d, "modem2_extension")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem2-extension"], "ExtenderControllerExtenderWanExtension-Modem2Extension"); ok {
			if err = d.Set("modem2_extension", vv); err != nil {
				return fmt.Errorf("Error reading modem2_extension: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem2_extension: %v", err)
		}
	}

	return nil
}

func flattenExtenderControllerExtenderWanExtensionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtenderControllerExtenderWanExtensionModem1Extension2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderWanExtensionModem2Extension2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectExtenderControllerExtenderWanExtension(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("modem1_extension"); ok || d.HasChange("modem1_extension") {
		t, err := expandExtenderControllerExtenderWanExtensionModem1Extension2edl(d, v, "modem1_extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem1-extension"] = t
		}
	}

	if v, ok := d.GetOk("modem2_extension"); ok || d.HasChange("modem2_extension") {
		t, err := expandExtenderControllerExtenderWanExtensionModem2Extension2edl(d, v, "modem2_extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem2-extension"] = t
		}
	}

	return &obj, nil
}
