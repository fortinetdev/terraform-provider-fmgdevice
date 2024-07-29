// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: WTP, FortiAP, or AP platform.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWtpProfilePlatform() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWtpProfilePlatformUpdate,
		Read:   resourceWirelessControllerWtpProfilePlatformRead,
		Update: resourceWirelessControllerWtpProfilePlatformUpdate,
		Delete: resourceWirelessControllerWtpProfilePlatformDelete,

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
			"wtp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"_local_platform_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddscan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerWtpProfilePlatformUpdate(d *schema.ResourceData, m interface{}) error {
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
	wtp_profile := d.Get("wtp_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp_profile"] = wtp_profile

	obj, err := getObjectWirelessControllerWtpProfilePlatform(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfilePlatform resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerWtpProfilePlatform(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfilePlatform resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerWtpProfilePlatform")

	return resourceWirelessControllerWtpProfilePlatformRead(d, m)
}

func resourceWirelessControllerWtpProfilePlatformDelete(d *schema.ResourceData, m interface{}) error {
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
	wtp_profile := d.Get("wtp_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp_profile"] = wtp_profile

	err = c.DeleteWirelessControllerWtpProfilePlatform(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWtpProfilePlatform resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpProfilePlatformRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	wtp_profile := d.Get("wtp_profile").(string)
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
	if wtp_profile == "" {
		wtp_profile = importOptionChecking(m.(*FortiClient).Cfg, "wtp_profile")
		if wtp_profile == "" {
			return fmt.Errorf("Parameter wtp_profile is missing")
		}
		if err = d.Set("wtp_profile", wtp_profile); err != nil {
			return fmt.Errorf("Error set params wtp_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp_profile"] = wtp_profile

	o, err := c.ReadWirelessControllerWtpProfilePlatform(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfilePlatform resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtpProfilePlatform(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfilePlatform resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpProfilePlatformLocalPlatformStr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfilePlatformDdscan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfilePlatformMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfilePlatformType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerWtpProfilePlatform(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("_local_platform_str", flattenWirelessControllerWtpProfilePlatformLocalPlatformStr2edl(o["_local_platform_str"], d, "_local_platform_str")); err != nil {
		if vv, ok := fortiAPIPatch(o["_local_platform_str"], "WirelessControllerWtpProfilePlatform-LocalPlatformStr"); ok {
			if err = d.Set("_local_platform_str", vv); err != nil {
				return fmt.Errorf("Error reading _local_platform_str: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _local_platform_str: %v", err)
		}
	}

	if err = d.Set("ddscan", flattenWirelessControllerWtpProfilePlatformDdscan2edl(o["ddscan"], d, "ddscan")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddscan"], "WirelessControllerWtpProfilePlatform-Ddscan"); ok {
			if err = d.Set("ddscan", vv); err != nil {
				return fmt.Errorf("Error reading ddscan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddscan: %v", err)
		}
	}

	if err = d.Set("mode", flattenWirelessControllerWtpProfilePlatformMode2edl(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "WirelessControllerWtpProfilePlatform-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("type", flattenWirelessControllerWtpProfilePlatformType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "WirelessControllerWtpProfilePlatform-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpProfilePlatformFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWtpProfilePlatformLocalPlatformStr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfilePlatformDdscan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfilePlatformMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfilePlatformType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWtpProfilePlatform(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_local_platform_str"); ok || d.HasChange("_local_platform_str") {
		t, err := expandWirelessControllerWtpProfilePlatformLocalPlatformStr2edl(d, v, "_local_platform_str")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_local_platform_str"] = t
		}
	}

	if v, ok := d.GetOk("ddscan"); ok || d.HasChange("ddscan") {
		t, err := expandWirelessControllerWtpProfilePlatformDdscan2edl(d, v, "ddscan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddscan"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandWirelessControllerWtpProfilePlatformMode2edl(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandWirelessControllerWtpProfilePlatformType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
