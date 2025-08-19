// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Icon list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20IconIconList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20IconIconListCreate,
		Read:   resourceWirelessControllerHotspot20IconIconListRead,
		Update: resourceWirelessControllerHotspot20IconIconListUpdate,
		Delete: resourceWirelessControllerHotspot20IconIconListDelete,

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
			"icon": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"height": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"lang": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"width": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20IconIconListCreate(d *schema.ResourceData, m interface{}) error {
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
	icon := d.Get("icon").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["icon"] = icon

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20IconIconList(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20IconIconList resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20IconIconList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20IconIconList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20IconIconListRead(d, m)
}

func resourceWirelessControllerHotspot20IconIconListUpdate(d *schema.ResourceData, m interface{}) error {
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
	icon := d.Get("icon").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["icon"] = icon

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20IconIconList(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20IconIconList resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20IconIconList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20IconIconList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20IconIconListRead(d, m)
}

func resourceWirelessControllerHotspot20IconIconListDelete(d *schema.ResourceData, m interface{}) error {
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
	icon := d.Get("icon").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["icon"] = icon

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerHotspot20IconIconList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20IconIconList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20IconIconListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	icon := d.Get("icon").(string)
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
	if icon == "" {
		icon = importOptionChecking(m.(*FortiClient).Cfg, "icon")
		if icon == "" {
			return fmt.Errorf("Parameter icon is missing")
		}
		if err = d.Set("icon", icon); err != nil {
			return fmt.Errorf("Error set params icon: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["icon"] = icon

	o, err := c.ReadWirelessControllerHotspot20IconIconList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20IconIconList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20IconIconList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20IconIconList resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20IconIconListFile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20IconIconListHeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20IconIconListLang2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20IconIconListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20IconIconListType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20IconIconListWidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20IconIconList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("file", flattenWirelessControllerHotspot20IconIconListFile2edl(o["file"], d, "file")); err != nil {
		if vv, ok := fortiAPIPatch(o["file"], "WirelessControllerHotspot20IconIconList-File"); ok {
			if err = d.Set("file", vv); err != nil {
				return fmt.Errorf("Error reading file: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file: %v", err)
		}
	}

	if err = d.Set("height", flattenWirelessControllerHotspot20IconIconListHeight2edl(o["height"], d, "height")); err != nil {
		if vv, ok := fortiAPIPatch(o["height"], "WirelessControllerHotspot20IconIconList-Height"); ok {
			if err = d.Set("height", vv); err != nil {
				return fmt.Errorf("Error reading height: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading height: %v", err)
		}
	}

	if err = d.Set("lang", flattenWirelessControllerHotspot20IconIconListLang2edl(o["lang"], d, "lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["lang"], "WirelessControllerHotspot20IconIconList-Lang"); ok {
			if err = d.Set("lang", vv); err != nil {
				return fmt.Errorf("Error reading lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lang: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerHotspot20IconIconListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerHotspot20IconIconList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenWirelessControllerHotspot20IconIconListType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "WirelessControllerHotspot20IconIconList-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("width", flattenWirelessControllerHotspot20IconIconListWidth2edl(o["width"], d, "width")); err != nil {
		if vv, ok := fortiAPIPatch(o["width"], "WirelessControllerHotspot20IconIconList-Width"); ok {
			if err = d.Set("width", vv); err != nil {
				return fmt.Errorf("Error reading width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading width: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20IconIconListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20IconIconListFile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20IconIconListHeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20IconIconListLang2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20IconIconListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20IconIconListType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20IconIconListWidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20IconIconList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("file"); ok || d.HasChange("file") {
		t, err := expandWirelessControllerHotspot20IconIconListFile2edl(d, v, "file")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file"] = t
		}
	}

	if v, ok := d.GetOk("height"); ok || d.HasChange("height") {
		t, err := expandWirelessControllerHotspot20IconIconListHeight2edl(d, v, "height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["height"] = t
		}
	}

	if v, ok := d.GetOk("lang"); ok || d.HasChange("lang") {
		t, err := expandWirelessControllerHotspot20IconIconListLang2edl(d, v, "lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lang"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerHotspot20IconIconListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandWirelessControllerHotspot20IconIconListType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("width"); ok || d.HasChange("width") {
		t, err := expandWirelessControllerHotspot20IconIconListWidth2edl(d, v, "width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["width"] = t
		}
	}

	return &obj, nil
}
