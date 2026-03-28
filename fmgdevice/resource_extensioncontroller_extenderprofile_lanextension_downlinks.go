// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Config FortiExtender downlink interface for LAN extension.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderProfileLanExtensionDownlinks() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderProfileLanExtensionDownlinksCreate,
		Read:   resourceExtensionControllerExtenderProfileLanExtensionDownlinksRead,
		Update: resourceExtensionControllerExtenderProfileLanExtensionDownlinksUpdate,
		Delete: resourceExtensionControllerExtenderProfileLanExtensionDownlinksDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pvid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vap": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vids": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceExtensionControllerExtenderProfileLanExtensionDownlinksCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderProfileLanExtensionDownlinks(d)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtenderProfileLanExtensionDownlinks resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadExtensionControllerExtenderProfileLanExtensionDownlinks(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateExtensionControllerExtenderProfileLanExtensionDownlinks(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ExtensionControllerExtenderProfileLanExtensionDownlinks resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateExtensionControllerExtenderProfileLanExtensionDownlinks(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ExtensionControllerExtenderProfileLanExtensionDownlinks resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerExtenderProfileLanExtensionDownlinksRead(d, m)
}

func resourceExtensionControllerExtenderProfileLanExtensionDownlinksUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderProfileLanExtensionDownlinks(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileLanExtensionDownlinks resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateExtensionControllerExtenderProfileLanExtensionDownlinks(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileLanExtensionDownlinks resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerExtenderProfileLanExtensionDownlinksRead(d, m)
}

func resourceExtensionControllerExtenderProfileLanExtensionDownlinksDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteExtensionControllerExtenderProfileLanExtensionDownlinks(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderProfileLanExtensionDownlinks resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderProfileLanExtensionDownlinksRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerExtenderProfileLanExtensionDownlinks(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileLanExtensionDownlinks resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderProfileLanExtensionDownlinks(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileLanExtensionDownlinks resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksPvid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksVap3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksVids3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func refreshObjectExtensionControllerExtenderProfileLanExtensionDownlinks(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenExtensionControllerExtenderProfileLanExtensionDownlinksName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ExtensionControllerExtenderProfileLanExtensionDownlinks-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenExtensionControllerExtenderProfileLanExtensionDownlinksPort3rdl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ExtensionControllerExtenderProfileLanExtensionDownlinks-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("pvid", flattenExtensionControllerExtenderProfileLanExtensionDownlinksPvid3rdl(o["pvid"], d, "pvid")); err != nil {
		if vv, ok := fortiAPIPatch(o["pvid"], "ExtensionControllerExtenderProfileLanExtensionDownlinks-Pvid"); ok {
			if err = d.Set("pvid", vv); err != nil {
				return fmt.Errorf("Error reading pvid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pvid: %v", err)
		}
	}

	if err = d.Set("type", flattenExtensionControllerExtenderProfileLanExtensionDownlinksType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ExtensionControllerExtenderProfileLanExtensionDownlinks-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("vap", flattenExtensionControllerExtenderProfileLanExtensionDownlinksVap3rdl(o["vap"], d, "vap")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap"], "ExtensionControllerExtenderProfileLanExtensionDownlinks-Vap"); ok {
			if err = d.Set("vap", vv); err != nil {
				return fmt.Errorf("Error reading vap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap: %v", err)
		}
	}

	if err = d.Set("vids", flattenExtensionControllerExtenderProfileLanExtensionDownlinksVids3rdl(o["vids"], d, "vids")); err != nil {
		if vv, ok := fortiAPIPatch(o["vids"], "ExtensionControllerExtenderProfileLanExtensionDownlinks-Vids"); ok {
			if err = d.Set("vids", vv); err != nil {
				return fmt.Errorf("Error reading vids: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vids: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksPvid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksVap3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksVids3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func getObjectExtensionControllerExtenderProfileLanExtensionDownlinks(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionDownlinksName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionDownlinksPort3rdl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("pvid"); ok || d.HasChange("pvid") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionDownlinksPvid3rdl(d, v, "pvid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pvid"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionDownlinksType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("vap"); ok || d.HasChange("vap") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionDownlinksVap3rdl(d, v, "vap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap"] = t
		}
	}

	if v, ok := d.GetOk("vids"); ok || d.HasChange("vids") {
		t, err := expandExtensionControllerExtenderProfileLanExtensionDownlinksVids3rdl(d, v, "vids")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vids"] = t
		}
	}

	return &obj, nil
}
