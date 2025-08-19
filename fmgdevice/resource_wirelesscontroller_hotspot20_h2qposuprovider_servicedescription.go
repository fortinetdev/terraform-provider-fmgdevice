// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSU service name.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20H2QpOsuProviderServiceDescription() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionCreate,
		Read:   resourceWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionRead,
		Update: resourceWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionUpdate,
		Delete: resourceWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionDelete,

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
			"h2qp_osu_provider": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"lang": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionCreate(d *schema.ResourceData, m interface{}) error {
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
	h2qp_osu_provider := d.Get("h2qp_osu_provider").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["h2qp_osu_provider"] = h2qp_osu_provider

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpOsuProviderServiceDescription resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20H2QpOsuProviderServiceDescription(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpOsuProviderServiceDescription resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "service_id")))

	return resourceWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionUpdate(d *schema.ResourceData, m interface{}) error {
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
	h2qp_osu_provider := d.Get("h2qp_osu_provider").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["h2qp_osu_provider"] = h2qp_osu_provider

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpOsuProviderServiceDescription resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20H2QpOsuProviderServiceDescription(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpOsuProviderServiceDescription resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "service_id")))

	return resourceWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionDelete(d *schema.ResourceData, m interface{}) error {
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
	h2qp_osu_provider := d.Get("h2qp_osu_provider").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["h2qp_osu_provider"] = h2qp_osu_provider

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerHotspot20H2QpOsuProviderServiceDescription(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20H2QpOsuProviderServiceDescription resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	h2qp_osu_provider := d.Get("h2qp_osu_provider").(string)
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
	if h2qp_osu_provider == "" {
		h2qp_osu_provider = importOptionChecking(m.(*FortiClient).Cfg, "h2qp_osu_provider")
		if h2qp_osu_provider == "" {
			return fmt.Errorf("Parameter h2qp_osu_provider is missing")
		}
		if err = d.Set("h2qp_osu_provider", h2qp_osu_provider); err != nil {
			return fmt.Errorf("Error set params h2qp_osu_provider: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["h2qp_osu_provider"] = h2qp_osu_provider

	o, err := c.ReadWirelessControllerHotspot20H2QpOsuProviderServiceDescription(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpOsuProviderServiceDescription resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpOsuProviderServiceDescription resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionLang2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("lang", flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionLang2edl(o["lang"], d, "lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["lang"], "WirelessControllerHotspot20H2QpOsuProviderServiceDescription-Lang"); ok {
			if err = d.Set("lang", vv); err != nil {
				return fmt.Errorf("Error reading lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lang: %v", err)
		}
	}

	if err = d.Set("service_description", flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceDescription2edl(o["service-description"], d, "service_description")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-description"], "WirelessControllerHotspot20H2QpOsuProviderServiceDescription-ServiceDescription"); ok {
			if err = d.Set("service_description", vv); err != nil {
				return fmt.Errorf("Error reading service_description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_description: %v", err)
		}
	}

	if err = d.Set("service_id", flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceId2edl(o["service-id"], d, "service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-id"], "WirelessControllerHotspot20H2QpOsuProviderServiceDescription-ServiceId"); ok {
			if err = d.Set("service_id", vv); err != nil {
				return fmt.Errorf("Error reading service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_id: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionLang2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("lang"); ok || d.HasChange("lang") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionLang2edl(d, v, "lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lang"] = t
		}
	}

	if v, ok := d.GetOk("service_description"); ok || d.HasChange("service_description") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceDescription2edl(d, v, "service_description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-description"] = t
		}
	}

	if v, ok := d.GetOk("service_id"); ok || d.HasChange("service_id") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceId2edl(d, v, "service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-id"] = t
		}
	}

	return &obj, nil
}
