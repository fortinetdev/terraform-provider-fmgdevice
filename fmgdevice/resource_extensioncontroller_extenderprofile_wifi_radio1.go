// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Radio-1 config for Wi-Fi 2.4GHz

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderProfileWifiRadio1() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderProfileWifiRadio1Update,
		Read:   resourceExtensionControllerExtenderProfileWifiRadio1Read,
		Update: resourceExtensionControllerExtenderProfileWifiRadio1Update,
		Delete: resourceExtensionControllerExtenderProfileWifiRadio1Delete,

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
			"n80211d": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"band": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"beacon_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bss_color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bss_color_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"channel": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"extension_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guard_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lan_ext_vap": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"local_vaps": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"max_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"operating_standard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"power_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radio_id": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceExtensionControllerExtenderProfileWifiRadio1Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderProfileWifiRadio1(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileWifiRadio1 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateExtensionControllerExtenderProfileWifiRadio1(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileWifiRadio1 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ExtensionControllerExtenderProfileWifiRadio1")

	return resourceExtensionControllerExtenderProfileWifiRadio1Read(d, m)
}

func resourceExtensionControllerExtenderProfileWifiRadio1Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteExtensionControllerExtenderProfileWifiRadio1(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderProfileWifiRadio1 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderProfileWifiRadio1Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerExtenderProfileWifiRadio1(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileWifiRadio1 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderProfileWifiRadio1(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileWifiRadio1 resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderProfileWifiRadio180211D3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Band3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Bandwidth3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1BeaconInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1BssColor3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1BssColorMode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Channel3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio1ExtensionChannel3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1GuardInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1LanExtVap3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio1LocalVaps3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio1MaxClients3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Mode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1OperatingStandard3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1PowerLevel3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1RadioId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Status3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtenderProfileWifiRadio1(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("n80211d", flattenExtensionControllerExtenderProfileWifiRadio180211D3rdl(o["80211d"], d, "n80211d")); err != nil {
		if vv, ok := fortiAPIPatch(o["80211d"], "ExtensionControllerExtenderProfileWifiRadio1-80211D"); ok {
			if err = d.Set("n80211d", vv); err != nil {
				return fmt.Errorf("Error reading n80211d: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n80211d: %v", err)
		}
	}

	if err = d.Set("band", flattenExtensionControllerExtenderProfileWifiRadio1Band3rdl(o["band"], d, "band")); err != nil {
		if vv, ok := fortiAPIPatch(o["band"], "ExtensionControllerExtenderProfileWifiRadio1-Band"); ok {
			if err = d.Set("band", vv); err != nil {
				return fmt.Errorf("Error reading band: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading band: %v", err)
		}
	}

	if err = d.Set("bandwidth", flattenExtensionControllerExtenderProfileWifiRadio1Bandwidth3rdl(o["bandwidth"], d, "bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth"], "ExtensionControllerExtenderProfileWifiRadio1-Bandwidth"); ok {
			if err = d.Set("bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth: %v", err)
		}
	}

	if err = d.Set("beacon_interval", flattenExtensionControllerExtenderProfileWifiRadio1BeaconInterval3rdl(o["beacon-interval"], d, "beacon_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["beacon-interval"], "ExtensionControllerExtenderProfileWifiRadio1-BeaconInterval"); ok {
			if err = d.Set("beacon_interval", vv); err != nil {
				return fmt.Errorf("Error reading beacon_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading beacon_interval: %v", err)
		}
	}

	if err = d.Set("bss_color", flattenExtensionControllerExtenderProfileWifiRadio1BssColor3rdl(o["bss-color"], d, "bss_color")); err != nil {
		if vv, ok := fortiAPIPatch(o["bss-color"], "ExtensionControllerExtenderProfileWifiRadio1-BssColor"); ok {
			if err = d.Set("bss_color", vv); err != nil {
				return fmt.Errorf("Error reading bss_color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bss_color: %v", err)
		}
	}

	if err = d.Set("bss_color_mode", flattenExtensionControllerExtenderProfileWifiRadio1BssColorMode3rdl(o["bss-color-mode"], d, "bss_color_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["bss-color-mode"], "ExtensionControllerExtenderProfileWifiRadio1-BssColorMode"); ok {
			if err = d.Set("bss_color_mode", vv); err != nil {
				return fmt.Errorf("Error reading bss_color_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bss_color_mode: %v", err)
		}
	}

	if err = d.Set("channel", flattenExtensionControllerExtenderProfileWifiRadio1Channel3rdl(o["channel"], d, "channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["channel"], "ExtensionControllerExtenderProfileWifiRadio1-Channel"); ok {
			if err = d.Set("channel", vv); err != nil {
				return fmt.Errorf("Error reading channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading channel: %v", err)
		}
	}

	if err = d.Set("extension_channel", flattenExtensionControllerExtenderProfileWifiRadio1ExtensionChannel3rdl(o["extension-channel"], d, "extension_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension-channel"], "ExtensionControllerExtenderProfileWifiRadio1-ExtensionChannel"); ok {
			if err = d.Set("extension_channel", vv); err != nil {
				return fmt.Errorf("Error reading extension_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension_channel: %v", err)
		}
	}

	if err = d.Set("guard_interval", flattenExtensionControllerExtenderProfileWifiRadio1GuardInterval3rdl(o["guard-interval"], d, "guard_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["guard-interval"], "ExtensionControllerExtenderProfileWifiRadio1-GuardInterval"); ok {
			if err = d.Set("guard_interval", vv); err != nil {
				return fmt.Errorf("Error reading guard_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guard_interval: %v", err)
		}
	}

	if err = d.Set("lan_ext_vap", flattenExtensionControllerExtenderProfileWifiRadio1LanExtVap3rdl(o["lan-ext-vap"], d, "lan_ext_vap")); err != nil {
		if vv, ok := fortiAPIPatch(o["lan-ext-vap"], "ExtensionControllerExtenderProfileWifiRadio1-LanExtVap"); ok {
			if err = d.Set("lan_ext_vap", vv); err != nil {
				return fmt.Errorf("Error reading lan_ext_vap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lan_ext_vap: %v", err)
		}
	}

	if err = d.Set("local_vaps", flattenExtensionControllerExtenderProfileWifiRadio1LocalVaps3rdl(o["local-vaps"], d, "local_vaps")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-vaps"], "ExtensionControllerExtenderProfileWifiRadio1-LocalVaps"); ok {
			if err = d.Set("local_vaps", vv); err != nil {
				return fmt.Errorf("Error reading local_vaps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_vaps: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenExtensionControllerExtenderProfileWifiRadio1MaxClients3rdl(o["max-clients"], d, "max_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-clients"], "ExtensionControllerExtenderProfileWifiRadio1-MaxClients"); ok {
			if err = d.Set("max_clients", vv); err != nil {
				return fmt.Errorf("Error reading max_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("mode", flattenExtensionControllerExtenderProfileWifiRadio1Mode3rdl(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "ExtensionControllerExtenderProfileWifiRadio1-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("operating_standard", flattenExtensionControllerExtenderProfileWifiRadio1OperatingStandard3rdl(o["operating-standard"], d, "operating_standard")); err != nil {
		if vv, ok := fortiAPIPatch(o["operating-standard"], "ExtensionControllerExtenderProfileWifiRadio1-OperatingStandard"); ok {
			if err = d.Set("operating_standard", vv); err != nil {
				return fmt.Errorf("Error reading operating_standard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading operating_standard: %v", err)
		}
	}

	if err = d.Set("power_level", flattenExtensionControllerExtenderProfileWifiRadio1PowerLevel3rdl(o["power-level"], d, "power_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["power-level"], "ExtensionControllerExtenderProfileWifiRadio1-PowerLevel"); ok {
			if err = d.Set("power_level", vv); err != nil {
				return fmt.Errorf("Error reading power_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading power_level: %v", err)
		}
	}

	if err = d.Set("radio_id", flattenExtensionControllerExtenderProfileWifiRadio1RadioId3rdl(o["radio-id"], d, "radio_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-id"], "ExtensionControllerExtenderProfileWifiRadio1-RadioId"); ok {
			if err = d.Set("radio_id", vv); err != nil {
				return fmt.Errorf("Error reading radio_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_id: %v", err)
		}
	}

	if err = d.Set("status", flattenExtensionControllerExtenderProfileWifiRadio1Status3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ExtensionControllerExtenderProfileWifiRadio1-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerExtenderProfileWifiRadio1FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerExtenderProfileWifiRadio180211D3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Band3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Bandwidth3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1BeaconInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1BssColor3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1BssColorMode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Channel3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileWifiRadio1ExtensionChannel3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1GuardInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1LanExtVap3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileWifiRadio1LocalVaps3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileWifiRadio1MaxClients3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Mode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1OperatingStandard3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1PowerLevel3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1RadioId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Status3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtenderProfileWifiRadio1(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n80211d"); ok || d.HasChange("n80211d") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio180211D3rdl(d, v, "n80211d")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["80211d"] = t
		}
	}

	if v, ok := d.GetOk("band"); ok || d.HasChange("band") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1Band3rdl(d, v, "band")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["band"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth"); ok || d.HasChange("bandwidth") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1Bandwidth3rdl(d, v, "bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("beacon_interval"); ok || d.HasChange("beacon_interval") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1BeaconInterval3rdl(d, v, "beacon_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-interval"] = t
		}
	}

	if v, ok := d.GetOk("bss_color"); ok || d.HasChange("bss_color") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1BssColor3rdl(d, v, "bss_color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-color"] = t
		}
	}

	if v, ok := d.GetOk("bss_color_mode"); ok || d.HasChange("bss_color_mode") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1BssColorMode3rdl(d, v, "bss_color_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-color-mode"] = t
		}
	}

	if v, ok := d.GetOk("channel"); ok || d.HasChange("channel") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1Channel3rdl(d, v, "channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["channel"] = t
		}
	}

	if v, ok := d.GetOk("extension_channel"); ok || d.HasChange("extension_channel") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1ExtensionChannel3rdl(d, v, "extension_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-channel"] = t
		}
	}

	if v, ok := d.GetOk("guard_interval"); ok || d.HasChange("guard_interval") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1GuardInterval3rdl(d, v, "guard_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guard-interval"] = t
		}
	}

	if v, ok := d.GetOk("lan_ext_vap"); ok || d.HasChange("lan_ext_vap") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1LanExtVap3rdl(d, v, "lan_ext_vap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-ext-vap"] = t
		}
	}

	if v, ok := d.GetOk("local_vaps"); ok || d.HasChange("local_vaps") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1LocalVaps3rdl(d, v, "local_vaps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-vaps"] = t
		}
	}

	if v, ok := d.GetOk("max_clients"); ok || d.HasChange("max_clients") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1MaxClients3rdl(d, v, "max_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1Mode3rdl(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("operating_standard"); ok || d.HasChange("operating_standard") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1OperatingStandard3rdl(d, v, "operating_standard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["operating-standard"] = t
		}
	}

	if v, ok := d.GetOk("power_level"); ok || d.HasChange("power_level") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1PowerLevel3rdl(d, v, "power_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["power-level"] = t
		}
	}

	if v, ok := d.GetOk("radio_id"); ok || d.HasChange("radio_id") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1RadioId3rdl(d, v, "radio_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1Status3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
