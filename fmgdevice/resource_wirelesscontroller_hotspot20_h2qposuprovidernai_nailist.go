// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSU NAI list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20H2QpOsuProviderNaiNaiList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20H2QpOsuProviderNaiNaiListCreate,
		Read:   resourceWirelessControllerHotspot20H2QpOsuProviderNaiNaiListRead,
		Update: resourceWirelessControllerHotspot20H2QpOsuProviderNaiNaiListUpdate,
		Delete: resourceWirelessControllerHotspot20H2QpOsuProviderNaiNaiListDelete,

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
			"h2qp_osu_provider_nai": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"osu_nai": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20H2QpOsuProviderNaiNaiListCreate(d *schema.ResourceData, m interface{}) error {
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
	h2qp_osu_provider_nai := d.Get("h2qp_osu_provider_nai").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["h2qp_osu_provider_nai"] = h2qp_osu_provider_nai

	obj, err := getObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20H2QpOsuProviderNaiNaiListRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpOsuProviderNaiNaiListUpdate(d *schema.ResourceData, m interface{}) error {
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
	h2qp_osu_provider_nai := d.Get("h2qp_osu_provider_nai").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["h2qp_osu_provider_nai"] = h2qp_osu_provider_nai

	obj, err := getObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20H2QpOsuProviderNaiNaiListRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpOsuProviderNaiNaiListDelete(d *schema.ResourceData, m interface{}) error {
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
	h2qp_osu_provider_nai := d.Get("h2qp_osu_provider_nai").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["h2qp_osu_provider_nai"] = h2qp_osu_provider_nai

	err = c.DeleteWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2QpOsuProviderNaiNaiListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	h2qp_osu_provider_nai := d.Get("h2qp_osu_provider_nai").(string)
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
	if h2qp_osu_provider_nai == "" {
		h2qp_osu_provider_nai = importOptionChecking(m.(*FortiClient).Cfg, "h2qp_osu_provider_nai")
		if h2qp_osu_provider_nai == "" {
			return fmt.Errorf("Parameter h2qp_osu_provider_nai is missing")
		}
		if err = d.Set("h2qp_osu_provider_nai", h2qp_osu_provider_nai); err != nil {
			return fmt.Errorf("Error set params h2qp_osu_provider_nai: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["h2qp_osu_provider_nai"] = h2qp_osu_provider_nai

	o, err := c.ReadWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20H2QpOsuProviderNaiNaiListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderNaiNaiListOsuNai2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerHotspot20H2QpOsuProviderNaiNaiListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerHotspot20H2QpOsuProviderNaiNaiList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("osu_nai", flattenWirelessControllerHotspot20H2QpOsuProviderNaiNaiListOsuNai2edl(o["osu-nai"], d, "osu_nai")); err != nil {
		if vv, ok := fortiAPIPatch(o["osu-nai"], "WirelessControllerHotspot20H2QpOsuProviderNaiNaiList-OsuNai"); ok {
			if err = d.Set("osu_nai", vv); err != nil {
				return fmt.Errorf("Error reading osu_nai: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading osu_nai: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20H2QpOsuProviderNaiNaiListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20H2QpOsuProviderNaiNaiListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderNaiNaiListOsuNai2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderNaiNaiListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("osu_nai"); ok || d.HasChange("osu_nai") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderNaiNaiListOsuNai2edl(d, v, "osu_nai")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osu-nai"] = t
		}
	}

	return &obj, nil
}
