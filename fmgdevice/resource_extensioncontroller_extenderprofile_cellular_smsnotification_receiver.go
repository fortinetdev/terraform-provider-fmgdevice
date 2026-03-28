// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SMS notification receiver list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderProfileCellularSmsNotificationReceiver() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderProfileCellularSmsNotificationReceiverCreate,
		Read:   resourceExtensionControllerExtenderProfileCellularSmsNotificationReceiverRead,
		Update: resourceExtensionControllerExtenderProfileCellularSmsNotificationReceiverUpdate,
		Delete: resourceExtensionControllerExtenderProfileCellularSmsNotificationReceiverDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

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
			"alert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"phone_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceExtensionControllerExtenderProfileCellularSmsNotificationReceiverCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadExtensionControllerExtenderProfileCellularSmsNotificationReceiver(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateExtensionControllerExtenderProfileCellularSmsNotificationReceiver(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateExtensionControllerExtenderProfileCellularSmsNotificationReceiver(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerExtenderProfileCellularSmsNotificationReceiverRead(d, m)
}

func resourceExtensionControllerExtenderProfileCellularSmsNotificationReceiverUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateExtensionControllerExtenderProfileCellularSmsNotificationReceiver(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerExtenderProfileCellularSmsNotificationReceiverRead(d, m)
}

func resourceExtensionControllerExtenderProfileCellularSmsNotificationReceiverDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteExtensionControllerExtenderProfileCellularSmsNotificationReceiver(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderProfileCellularSmsNotificationReceiverRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerExtenderProfileCellularSmsNotificationReceiver(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("alert", flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert4thl(o["alert"], d, "alert")); err != nil {
		if vv, ok := fortiAPIPatch(o["alert"], "ExtensionControllerExtenderProfileCellularSmsNotificationReceiver-Alert"); ok {
			if err = d.Set("alert", vv); err != nil {
				return fmt.Errorf("Error reading alert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alert: %v", err)
		}
	}

	if err = d.Set("name", flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverName4thl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ExtensionControllerExtenderProfileCellularSmsNotificationReceiver-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("phone_number", flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber4thl(o["phone-number"], d, "phone_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["phone-number"], "ExtensionControllerExtenderProfileCellularSmsNotificationReceiver-PhoneNumber"); ok {
			if err = d.Set("phone_number", vv); err != nil {
				return fmt.Errorf("Error reading phone_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phone_number: %v", err)
		}
	}

	if err = d.Set("status", flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus4thl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ExtensionControllerExtenderProfileCellularSmsNotificationReceiver-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("alert"); ok || d.HasChange("alert") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert4thl(d, v, "alert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alert"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverName4thl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("phone_number"); ok || d.HasChange("phone_number") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber4thl(d, v, "phone_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phone-number"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus4thl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
