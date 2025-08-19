// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration options for radio 4.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWtpRadio4() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWtpRadio4Update,
		Read:   resourceWirelessControllerWtpRadio4Read,
		Update: resourceWirelessControllerWtpRadio4Update,
		Delete: resourceWirelessControllerWtpRadio4Delete,

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
			"wtp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auto_power_high": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auto_power_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_power_low": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auto_power_target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"band": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"channel": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"drma_manual_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"override_analysis": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"override_band": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_txpower": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_vaps": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"power_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"power_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"power_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radio_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"spectrum_analysis": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vaps": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerWtpRadio4Update(d *schema.ResourceData, m interface{}) error {
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
	wtp := d.Get("wtp").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp"] = wtp

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerWtpRadio4(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpRadio4 resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerWtpRadio4(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpRadio4 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerWtpRadio4")

	return resourceWirelessControllerWtpRadio4Read(d, m)
}

func resourceWirelessControllerWtpRadio4Delete(d *schema.ResourceData, m interface{}) error {
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
	wtp := d.Get("wtp").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp"] = wtp

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerWtpRadio4(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWtpRadio4 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpRadio4Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	wtp := d.Get("wtp").(string)
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
	if wtp == "" {
		wtp = importOptionChecking(m.(*FortiClient).Cfg, "wtp")
		if wtp == "" {
			return fmt.Errorf("Parameter wtp is missing")
		}
		if err = d.Set("wtp", wtp); err != nil {
			return fmt.Errorf("Error set params wtp: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp"] = wtp

	o, err := c.ReadWirelessControllerWtpRadio4(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpRadio4 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtpRadio4(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpRadio4 resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpRadio4AutoPowerHigh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4AutoPowerLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4AutoPowerLow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4AutoPowerTarget2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Band2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenWirelessControllerWtpRadio4Channel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpRadio4DrmaManualMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideAnalysis2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideBand2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideChannel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideTxpower2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideVaps2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4PowerLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4PowerMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4PowerValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4RadioId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4SpectrumAnalysis2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4VapAll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap52edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap72edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap82edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vaps2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWirelessControllerWtpRadio4(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auto_power_high", flattenWirelessControllerWtpRadio4AutoPowerHigh2edl(o["auto-power-high"], d, "auto_power_high")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-power-high"], "WirelessControllerWtpRadio4-AutoPowerHigh"); ok {
			if err = d.Set("auto_power_high", vv); err != nil {
				return fmt.Errorf("Error reading auto_power_high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_power_high: %v", err)
		}
	}

	if err = d.Set("auto_power_level", flattenWirelessControllerWtpRadio4AutoPowerLevel2edl(o["auto-power-level"], d, "auto_power_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-power-level"], "WirelessControllerWtpRadio4-AutoPowerLevel"); ok {
			if err = d.Set("auto_power_level", vv); err != nil {
				return fmt.Errorf("Error reading auto_power_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_power_level: %v", err)
		}
	}

	if err = d.Set("auto_power_low", flattenWirelessControllerWtpRadio4AutoPowerLow2edl(o["auto-power-low"], d, "auto_power_low")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-power-low"], "WirelessControllerWtpRadio4-AutoPowerLow"); ok {
			if err = d.Set("auto_power_low", vv); err != nil {
				return fmt.Errorf("Error reading auto_power_low: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_power_low: %v", err)
		}
	}

	if err = d.Set("auto_power_target", flattenWirelessControllerWtpRadio4AutoPowerTarget2edl(o["auto-power-target"], d, "auto_power_target")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-power-target"], "WirelessControllerWtpRadio4-AutoPowerTarget"); ok {
			if err = d.Set("auto_power_target", vv); err != nil {
				return fmt.Errorf("Error reading auto_power_target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_power_target: %v", err)
		}
	}

	if err = d.Set("band", flattenWirelessControllerWtpRadio4Band2edl(o["band"], d, "band")); err != nil {
		if vv, ok := fortiAPIPatch(o["band"], "WirelessControllerWtpRadio4-Band"); ok {
			if err = d.Set("band", vv); err != nil {
				return fmt.Errorf("Error reading band: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading band: %v", err)
		}
	}

	if err = d.Set("channel", flattenWirelessControllerWtpRadio4Channel2edl(o["channel"], d, "channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["channel"], "WirelessControllerWtpRadio4-Channel"); ok {
			if err = d.Set("channel", vv); err != nil {
				return fmt.Errorf("Error reading channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading channel: %v", err)
		}
	}

	if err = d.Set("drma_manual_mode", flattenWirelessControllerWtpRadio4DrmaManualMode2edl(o["drma-manual-mode"], d, "drma_manual_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["drma-manual-mode"], "WirelessControllerWtpRadio4-DrmaManualMode"); ok {
			if err = d.Set("drma_manual_mode", vv); err != nil {
				return fmt.Errorf("Error reading drma_manual_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drma_manual_mode: %v", err)
		}
	}

	if err = d.Set("override_analysis", flattenWirelessControllerWtpRadio4OverrideAnalysis2edl(o["override-analysis"], d, "override_analysis")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-analysis"], "WirelessControllerWtpRadio4-OverrideAnalysis"); ok {
			if err = d.Set("override_analysis", vv); err != nil {
				return fmt.Errorf("Error reading override_analysis: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_analysis: %v", err)
		}
	}

	if err = d.Set("override_band", flattenWirelessControllerWtpRadio4OverrideBand2edl(o["override-band"], d, "override_band")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-band"], "WirelessControllerWtpRadio4-OverrideBand"); ok {
			if err = d.Set("override_band", vv); err != nil {
				return fmt.Errorf("Error reading override_band: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_band: %v", err)
		}
	}

	if err = d.Set("override_channel", flattenWirelessControllerWtpRadio4OverrideChannel2edl(o["override-channel"], d, "override_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-channel"], "WirelessControllerWtpRadio4-OverrideChannel"); ok {
			if err = d.Set("override_channel", vv); err != nil {
				return fmt.Errorf("Error reading override_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_channel: %v", err)
		}
	}

	if err = d.Set("override_txpower", flattenWirelessControllerWtpRadio4OverrideTxpower2edl(o["override-txpower"], d, "override_txpower")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-txpower"], "WirelessControllerWtpRadio4-OverrideTxpower"); ok {
			if err = d.Set("override_txpower", vv); err != nil {
				return fmt.Errorf("Error reading override_txpower: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_txpower: %v", err)
		}
	}

	if err = d.Set("override_vaps", flattenWirelessControllerWtpRadio4OverrideVaps2edl(o["override-vaps"], d, "override_vaps")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-vaps"], "WirelessControllerWtpRadio4-OverrideVaps"); ok {
			if err = d.Set("override_vaps", vv); err != nil {
				return fmt.Errorf("Error reading override_vaps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_vaps: %v", err)
		}
	}

	if err = d.Set("power_level", flattenWirelessControllerWtpRadio4PowerLevel2edl(o["power-level"], d, "power_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["power-level"], "WirelessControllerWtpRadio4-PowerLevel"); ok {
			if err = d.Set("power_level", vv); err != nil {
				return fmt.Errorf("Error reading power_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading power_level: %v", err)
		}
	}

	if err = d.Set("power_mode", flattenWirelessControllerWtpRadio4PowerMode2edl(o["power-mode"], d, "power_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["power-mode"], "WirelessControllerWtpRadio4-PowerMode"); ok {
			if err = d.Set("power_mode", vv); err != nil {
				return fmt.Errorf("Error reading power_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading power_mode: %v", err)
		}
	}

	if err = d.Set("power_value", flattenWirelessControllerWtpRadio4PowerValue2edl(o["power-value"], d, "power_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["power-value"], "WirelessControllerWtpRadio4-PowerValue"); ok {
			if err = d.Set("power_value", vv); err != nil {
				return fmt.Errorf("Error reading power_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading power_value: %v", err)
		}
	}

	if err = d.Set("radio_id", flattenWirelessControllerWtpRadio4RadioId2edl(o["radio-id"], d, "radio_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-id"], "WirelessControllerWtpRadio4-RadioId"); ok {
			if err = d.Set("radio_id", vv); err != nil {
				return fmt.Errorf("Error reading radio_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_id: %v", err)
		}
	}

	if err = d.Set("spectrum_analysis", flattenWirelessControllerWtpRadio4SpectrumAnalysis2edl(o["spectrum-analysis"], d, "spectrum_analysis")); err != nil {
		if vv, ok := fortiAPIPatch(o["spectrum-analysis"], "WirelessControllerWtpRadio4-SpectrumAnalysis"); ok {
			if err = d.Set("spectrum_analysis", vv); err != nil {
				return fmt.Errorf("Error reading spectrum_analysis: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spectrum_analysis: %v", err)
		}
	}

	if err = d.Set("vap_all", flattenWirelessControllerWtpRadio4VapAll2edl(o["vap-all"], d, "vap_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap-all"], "WirelessControllerWtpRadio4-VapAll"); ok {
			if err = d.Set("vap_all", vv); err != nil {
				return fmt.Errorf("Error reading vap_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap_all: %v", err)
		}
	}

	if err = d.Set("vap1", flattenWirelessControllerWtpRadio4Vap12edl(o["vap1"], d, "vap1")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap1"], "WirelessControllerWtpRadio4-Vap1"); ok {
			if err = d.Set("vap1", vv); err != nil {
				return fmt.Errorf("Error reading vap1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap1: %v", err)
		}
	}

	if err = d.Set("vap2", flattenWirelessControllerWtpRadio4Vap22edl(o["vap2"], d, "vap2")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap2"], "WirelessControllerWtpRadio4-Vap2"); ok {
			if err = d.Set("vap2", vv); err != nil {
				return fmt.Errorf("Error reading vap2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap2: %v", err)
		}
	}

	if err = d.Set("vap3", flattenWirelessControllerWtpRadio4Vap32edl(o["vap3"], d, "vap3")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap3"], "WirelessControllerWtpRadio4-Vap3"); ok {
			if err = d.Set("vap3", vv); err != nil {
				return fmt.Errorf("Error reading vap3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap3: %v", err)
		}
	}

	if err = d.Set("vap4", flattenWirelessControllerWtpRadio4Vap42edl(o["vap4"], d, "vap4")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap4"], "WirelessControllerWtpRadio4-Vap4"); ok {
			if err = d.Set("vap4", vv); err != nil {
				return fmt.Errorf("Error reading vap4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap4: %v", err)
		}
	}

	if err = d.Set("vap5", flattenWirelessControllerWtpRadio4Vap52edl(o["vap5"], d, "vap5")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap5"], "WirelessControllerWtpRadio4-Vap5"); ok {
			if err = d.Set("vap5", vv); err != nil {
				return fmt.Errorf("Error reading vap5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap5: %v", err)
		}
	}

	if err = d.Set("vap6", flattenWirelessControllerWtpRadio4Vap62edl(o["vap6"], d, "vap6")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap6"], "WirelessControllerWtpRadio4-Vap6"); ok {
			if err = d.Set("vap6", vv); err != nil {
				return fmt.Errorf("Error reading vap6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap6: %v", err)
		}
	}

	if err = d.Set("vap7", flattenWirelessControllerWtpRadio4Vap72edl(o["vap7"], d, "vap7")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap7"], "WirelessControllerWtpRadio4-Vap7"); ok {
			if err = d.Set("vap7", vv); err != nil {
				return fmt.Errorf("Error reading vap7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap7: %v", err)
		}
	}

	if err = d.Set("vap8", flattenWirelessControllerWtpRadio4Vap82edl(o["vap8"], d, "vap8")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap8"], "WirelessControllerWtpRadio4-Vap8"); ok {
			if err = d.Set("vap8", vv); err != nil {
				return fmt.Errorf("Error reading vap8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap8: %v", err)
		}
	}

	if err = d.Set("vaps", flattenWirelessControllerWtpRadio4Vaps2edl(o["vaps"], d, "vaps")); err != nil {
		if vv, ok := fortiAPIPatch(o["vaps"], "WirelessControllerWtpRadio4-Vaps"); ok {
			if err = d.Set("vaps", vv); err != nil {
				return fmt.Errorf("Error reading vaps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vaps: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpRadio4FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWtpRadio4AutoPowerHigh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4AutoPowerLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4AutoPowerLow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4AutoPowerTarget2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Band2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio4Channel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio4DrmaManualMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideAnalysis2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideBand2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideChannel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideTxpower2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideVaps2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4PowerLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4PowerMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4PowerValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4RadioId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4SpectrumAnalysis2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4VapAll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap52edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap72edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap82edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vaps2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWirelessControllerWtpRadio4(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_power_high"); ok || d.HasChange("auto_power_high") {
		t, err := expandWirelessControllerWtpRadio4AutoPowerHigh2edl(d, v, "auto_power_high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-power-high"] = t
		}
	}

	if v, ok := d.GetOk("auto_power_level"); ok || d.HasChange("auto_power_level") {
		t, err := expandWirelessControllerWtpRadio4AutoPowerLevel2edl(d, v, "auto_power_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-power-level"] = t
		}
	}

	if v, ok := d.GetOk("auto_power_low"); ok || d.HasChange("auto_power_low") {
		t, err := expandWirelessControllerWtpRadio4AutoPowerLow2edl(d, v, "auto_power_low")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-power-low"] = t
		}
	}

	if v, ok := d.GetOk("auto_power_target"); ok || d.HasChange("auto_power_target") {
		t, err := expandWirelessControllerWtpRadio4AutoPowerTarget2edl(d, v, "auto_power_target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-power-target"] = t
		}
	}

	if v, ok := d.GetOk("band"); ok || d.HasChange("band") {
		t, err := expandWirelessControllerWtpRadio4Band2edl(d, v, "band")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["band"] = t
		}
	}

	if v, ok := d.GetOk("channel"); ok || d.HasChange("channel") {
		t, err := expandWirelessControllerWtpRadio4Channel2edl(d, v, "channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["channel"] = t
		}
	}

	if v, ok := d.GetOk("drma_manual_mode"); ok || d.HasChange("drma_manual_mode") {
		t, err := expandWirelessControllerWtpRadio4DrmaManualMode2edl(d, v, "drma_manual_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drma-manual-mode"] = t
		}
	}

	if v, ok := d.GetOk("override_analysis"); ok || d.HasChange("override_analysis") {
		t, err := expandWirelessControllerWtpRadio4OverrideAnalysis2edl(d, v, "override_analysis")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-analysis"] = t
		}
	}

	if v, ok := d.GetOk("override_band"); ok || d.HasChange("override_band") {
		t, err := expandWirelessControllerWtpRadio4OverrideBand2edl(d, v, "override_band")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-band"] = t
		}
	}

	if v, ok := d.GetOk("override_channel"); ok || d.HasChange("override_channel") {
		t, err := expandWirelessControllerWtpRadio4OverrideChannel2edl(d, v, "override_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-channel"] = t
		}
	}

	if v, ok := d.GetOk("override_txpower"); ok || d.HasChange("override_txpower") {
		t, err := expandWirelessControllerWtpRadio4OverrideTxpower2edl(d, v, "override_txpower")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-txpower"] = t
		}
	}

	if v, ok := d.GetOk("override_vaps"); ok || d.HasChange("override_vaps") {
		t, err := expandWirelessControllerWtpRadio4OverrideVaps2edl(d, v, "override_vaps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-vaps"] = t
		}
	}

	if v, ok := d.GetOk("power_level"); ok || d.HasChange("power_level") {
		t, err := expandWirelessControllerWtpRadio4PowerLevel2edl(d, v, "power_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["power-level"] = t
		}
	}

	if v, ok := d.GetOk("power_mode"); ok || d.HasChange("power_mode") {
		t, err := expandWirelessControllerWtpRadio4PowerMode2edl(d, v, "power_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["power-mode"] = t
		}
	}

	if v, ok := d.GetOk("power_value"); ok || d.HasChange("power_value") {
		t, err := expandWirelessControllerWtpRadio4PowerValue2edl(d, v, "power_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["power-value"] = t
		}
	}

	if v, ok := d.GetOk("radio_id"); ok || d.HasChange("radio_id") {
		t, err := expandWirelessControllerWtpRadio4RadioId2edl(d, v, "radio_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-id"] = t
		}
	}

	if v, ok := d.GetOk("spectrum_analysis"); ok || d.HasChange("spectrum_analysis") {
		t, err := expandWirelessControllerWtpRadio4SpectrumAnalysis2edl(d, v, "spectrum_analysis")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spectrum-analysis"] = t
		}
	}

	if v, ok := d.GetOk("vap_all"); ok || d.HasChange("vap_all") {
		t, err := expandWirelessControllerWtpRadio4VapAll2edl(d, v, "vap_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap-all"] = t
		}
	}

	if v, ok := d.GetOk("vap1"); ok || d.HasChange("vap1") {
		t, err := expandWirelessControllerWtpRadio4Vap12edl(d, v, "vap1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap1"] = t
		}
	}

	if v, ok := d.GetOk("vap2"); ok || d.HasChange("vap2") {
		t, err := expandWirelessControllerWtpRadio4Vap22edl(d, v, "vap2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap2"] = t
		}
	}

	if v, ok := d.GetOk("vap3"); ok || d.HasChange("vap3") {
		t, err := expandWirelessControllerWtpRadio4Vap32edl(d, v, "vap3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap3"] = t
		}
	}

	if v, ok := d.GetOk("vap4"); ok || d.HasChange("vap4") {
		t, err := expandWirelessControllerWtpRadio4Vap42edl(d, v, "vap4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap4"] = t
		}
	}

	if v, ok := d.GetOk("vap5"); ok || d.HasChange("vap5") {
		t, err := expandWirelessControllerWtpRadio4Vap52edl(d, v, "vap5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap5"] = t
		}
	}

	if v, ok := d.GetOk("vap6"); ok || d.HasChange("vap6") {
		t, err := expandWirelessControllerWtpRadio4Vap62edl(d, v, "vap6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap6"] = t
		}
	}

	if v, ok := d.GetOk("vap7"); ok || d.HasChange("vap7") {
		t, err := expandWirelessControllerWtpRadio4Vap72edl(d, v, "vap7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap7"] = t
		}
	}

	if v, ok := d.GetOk("vap8"); ok || d.HasChange("vap8") {
		t, err := expandWirelessControllerWtpRadio4Vap82edl(d, v, "vap8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap8"] = t
		}
	}

	if v, ok := d.GetOk("vaps"); ok || d.HasChange("vaps") {
		t, err := expandWirelessControllerWtpRadio4Vaps2edl(d, v, "vaps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vaps"] = t
		}
	}

	return &obj, nil
}
