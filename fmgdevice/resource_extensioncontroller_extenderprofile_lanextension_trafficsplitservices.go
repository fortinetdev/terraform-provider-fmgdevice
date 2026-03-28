// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Config FortiExtender traffic split interface for LAN extension.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderProfileLanExtensionTrafficSplitServices() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesCreate,
		Read:   resourceExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesRead,
		Update: resourceExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesUpdate,
		Delete: resourceExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesDelete,

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
			"address": &schema.Schema{
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
			"service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vsdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesRead(d, m)
}

func resourceExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesRead(d, m)
}

func resourceExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress3rdl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ExtensionControllerExtenderProfileLanExtensionTrafficSplitServices-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("name", flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ExtensionControllerExtenderProfileLanExtensionTrafficSplitServices-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("service", flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService3rdl(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "ExtensionControllerExtenderProfileLanExtensionTrafficSplitServices-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("vsdb", flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb3rdl(o["vsdb"], d, "vsdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["vsdb"], "ExtensionControllerExtenderProfileLanExtensionTrafficSplitServices-Vsdb"); ok {
			if err = d.Set("vsdb", vv); err != nil {
				return fmt.Errorf("Error reading vsdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vsdb: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress3rdl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService3rdl(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("vsdb"); ok || d.HasChange("vsdb") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb3rdl(d, v, "vsdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vsdb"] = t
		}
	}

	return &obj, nil
}
