// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiAP regions (for floor plans and maps).

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerRegion() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerRegionCreate,
		Read:   resourceWirelessControllerRegionRead,
		Update: resourceWirelessControllerRegionUpdate,
		Delete: resourceWirelessControllerRegionDelete,

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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"grayscale": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"opacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerRegionCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerRegion(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerRegion resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerRegion(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerRegion resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerRegionRead(d, m)
}

func resourceWirelessControllerRegionUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerRegion(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerRegion resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerRegion(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerRegion resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerRegionRead(d, m)
}

func resourceWirelessControllerRegionDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerRegion(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerRegion resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerRegionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerRegion(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerRegion resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerRegion(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerRegion resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerRegionComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerRegionGrayscale(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerRegionImageType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerRegionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerRegionOpacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerRegion(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comments", flattenWirelessControllerRegionComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "WirelessControllerRegion-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("grayscale", flattenWirelessControllerRegionGrayscale(o["grayscale"], d, "grayscale")); err != nil {
		if vv, ok := fortiAPIPatch(o["grayscale"], "WirelessControllerRegion-Grayscale"); ok {
			if err = d.Set("grayscale", vv); err != nil {
				return fmt.Errorf("Error reading grayscale: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading grayscale: %v", err)
		}
	}

	if err = d.Set("image_type", flattenWirelessControllerRegionImageType(o["image-type"], d, "image_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-type"], "WirelessControllerRegion-ImageType"); ok {
			if err = d.Set("image_type", vv); err != nil {
				return fmt.Errorf("Error reading image_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_type: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerRegionName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerRegion-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("opacity", flattenWirelessControllerRegionOpacity(o["opacity"], d, "opacity")); err != nil {
		if vv, ok := fortiAPIPatch(o["opacity"], "WirelessControllerRegion-Opacity"); ok {
			if err = d.Set("opacity", vv); err != nil {
				return fmt.Errorf("Error reading opacity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading opacity: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerRegionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerRegionComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerRegionGrayscale(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerRegionImageType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerRegionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerRegionOpacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerRegion(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandWirelessControllerRegionComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("grayscale"); ok || d.HasChange("grayscale") {
		t, err := expandWirelessControllerRegionGrayscale(d, v, "grayscale")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["grayscale"] = t
		}
	}

	if v, ok := d.GetOk("image_type"); ok || d.HasChange("image_type") {
		t, err := expandWirelessControllerRegionImageType(d, v, "image_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerRegionName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("opacity"); ok || d.HasChange("opacity") {
		t, err := expandWirelessControllerRegionOpacity(d, v, "opacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["opacity"] = t
		}
	}

	return &obj, nil
}
