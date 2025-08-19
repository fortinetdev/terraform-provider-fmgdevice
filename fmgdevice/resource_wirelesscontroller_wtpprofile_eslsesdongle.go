// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ESL SES-imagotag dongle configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWtpProfileEslSesDongle() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWtpProfileEslSesDongleUpdate,
		Read:   resourceWirelessControllerWtpProfileEslSesDongleRead,
		Update: resourceWirelessControllerWtpProfileEslSesDongleUpdate,
		Delete: resourceWirelessControllerWtpProfileEslSesDongleDelete,

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
			"apc_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apc_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apc_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apc_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"coex_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compliance_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"esl_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"output_power": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scd_enable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_cert_verification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_fqdn_verification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerWtpProfileEslSesDongleUpdate(d *schema.ResourceData, m interface{}) error {
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
	wtp_profile := d.Get("wtp_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp_profile"] = wtp_profile

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerWtpProfileEslSesDongle(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfileEslSesDongle resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerWtpProfileEslSesDongle(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfileEslSesDongle resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerWtpProfileEslSesDongle")

	return resourceWirelessControllerWtpProfileEslSesDongleRead(d, m)
}

func resourceWirelessControllerWtpProfileEslSesDongleDelete(d *schema.ResourceData, m interface{}) error {
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
	wtp_profile := d.Get("wtp_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp_profile"] = wtp_profile

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerWtpProfileEslSesDongle(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWtpProfileEslSesDongle resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpProfileEslSesDongleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerWtpProfileEslSesDongle(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfileEslSesDongle resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtpProfileEslSesDongle(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfileEslSesDongle resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpProfileEslSesDongleApcAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleApcFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleApcIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleApcPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleCoexLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleComplianceLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleEslChannel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleOutputPower2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleScdEnable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleTlsCertVerification2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleTlsFqdnVerification2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerWtpProfileEslSesDongle(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("apc_addr_type", flattenWirelessControllerWtpProfileEslSesDongleApcAddrType2edl(o["apc-addr-type"], d, "apc_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["apc-addr-type"], "WirelessControllerWtpProfileEslSesDongle-ApcAddrType"); ok {
			if err = d.Set("apc_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading apc_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apc_addr_type: %v", err)
		}
	}

	if err = d.Set("apc_fqdn", flattenWirelessControllerWtpProfileEslSesDongleApcFqdn2edl(o["apc-fqdn"], d, "apc_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["apc-fqdn"], "WirelessControllerWtpProfileEslSesDongle-ApcFqdn"); ok {
			if err = d.Set("apc_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading apc_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apc_fqdn: %v", err)
		}
	}

	if err = d.Set("apc_ip", flattenWirelessControllerWtpProfileEslSesDongleApcIp2edl(o["apc-ip"], d, "apc_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["apc-ip"], "WirelessControllerWtpProfileEslSesDongle-ApcIp"); ok {
			if err = d.Set("apc_ip", vv); err != nil {
				return fmt.Errorf("Error reading apc_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apc_ip: %v", err)
		}
	}

	if err = d.Set("apc_port", flattenWirelessControllerWtpProfileEslSesDongleApcPort2edl(o["apc-port"], d, "apc_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["apc-port"], "WirelessControllerWtpProfileEslSesDongle-ApcPort"); ok {
			if err = d.Set("apc_port", vv); err != nil {
				return fmt.Errorf("Error reading apc_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apc_port: %v", err)
		}
	}

	if err = d.Set("coex_level", flattenWirelessControllerWtpProfileEslSesDongleCoexLevel2edl(o["coex-level"], d, "coex_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["coex-level"], "WirelessControllerWtpProfileEslSesDongle-CoexLevel"); ok {
			if err = d.Set("coex_level", vv); err != nil {
				return fmt.Errorf("Error reading coex_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading coex_level: %v", err)
		}
	}

	if err = d.Set("compliance_level", flattenWirelessControllerWtpProfileEslSesDongleComplianceLevel2edl(o["compliance-level"], d, "compliance_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["compliance-level"], "WirelessControllerWtpProfileEslSesDongle-ComplianceLevel"); ok {
			if err = d.Set("compliance_level", vv); err != nil {
				return fmt.Errorf("Error reading compliance_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading compliance_level: %v", err)
		}
	}

	if err = d.Set("esl_channel", flattenWirelessControllerWtpProfileEslSesDongleEslChannel2edl(o["esl-channel"], d, "esl_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["esl-channel"], "WirelessControllerWtpProfileEslSesDongle-EslChannel"); ok {
			if err = d.Set("esl_channel", vv); err != nil {
				return fmt.Errorf("Error reading esl_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading esl_channel: %v", err)
		}
	}

	if err = d.Set("output_power", flattenWirelessControllerWtpProfileEslSesDongleOutputPower2edl(o["output-power"], d, "output_power")); err != nil {
		if vv, ok := fortiAPIPatch(o["output-power"], "WirelessControllerWtpProfileEslSesDongle-OutputPower"); ok {
			if err = d.Set("output_power", vv); err != nil {
				return fmt.Errorf("Error reading output_power: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading output_power: %v", err)
		}
	}

	if err = d.Set("scd_enable", flattenWirelessControllerWtpProfileEslSesDongleScdEnable2edl(o["scd-enable"], d, "scd_enable")); err != nil {
		if vv, ok := fortiAPIPatch(o["scd-enable"], "WirelessControllerWtpProfileEslSesDongle-ScdEnable"); ok {
			if err = d.Set("scd_enable", vv); err != nil {
				return fmt.Errorf("Error reading scd_enable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scd_enable: %v", err)
		}
	}

	if err = d.Set("tls_cert_verification", flattenWirelessControllerWtpProfileEslSesDongleTlsCertVerification2edl(o["tls-cert-verification"], d, "tls_cert_verification")); err != nil {
		if vv, ok := fortiAPIPatch(o["tls-cert-verification"], "WirelessControllerWtpProfileEslSesDongle-TlsCertVerification"); ok {
			if err = d.Set("tls_cert_verification", vv); err != nil {
				return fmt.Errorf("Error reading tls_cert_verification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tls_cert_verification: %v", err)
		}
	}

	if err = d.Set("tls_fqdn_verification", flattenWirelessControllerWtpProfileEslSesDongleTlsFqdnVerification2edl(o["tls-fqdn-verification"], d, "tls_fqdn_verification")); err != nil {
		if vv, ok := fortiAPIPatch(o["tls-fqdn-verification"], "WirelessControllerWtpProfileEslSesDongle-TlsFqdnVerification"); ok {
			if err = d.Set("tls_fqdn_verification", vv); err != nil {
				return fmt.Errorf("Error reading tls_fqdn_verification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tls_fqdn_verification: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpProfileEslSesDongleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWtpProfileEslSesDongleApcAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleApcFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleApcIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleApcPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleCoexLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleComplianceLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleEslChannel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleOutputPower2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleScdEnable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleTlsCertVerification2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleTlsFqdnVerification2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWtpProfileEslSesDongle(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("apc_addr_type"); ok || d.HasChange("apc_addr_type") {
		t, err := expandWirelessControllerWtpProfileEslSesDongleApcAddrType2edl(d, v, "apc_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apc-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("apc_fqdn"); ok || d.HasChange("apc_fqdn") {
		t, err := expandWirelessControllerWtpProfileEslSesDongleApcFqdn2edl(d, v, "apc_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apc-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("apc_ip"); ok || d.HasChange("apc_ip") {
		t, err := expandWirelessControllerWtpProfileEslSesDongleApcIp2edl(d, v, "apc_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apc-ip"] = t
		}
	}

	if v, ok := d.GetOk("apc_port"); ok || d.HasChange("apc_port") {
		t, err := expandWirelessControllerWtpProfileEslSesDongleApcPort2edl(d, v, "apc_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apc-port"] = t
		}
	}

	if v, ok := d.GetOk("coex_level"); ok || d.HasChange("coex_level") {
		t, err := expandWirelessControllerWtpProfileEslSesDongleCoexLevel2edl(d, v, "coex_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["coex-level"] = t
		}
	}

	if v, ok := d.GetOk("compliance_level"); ok || d.HasChange("compliance_level") {
		t, err := expandWirelessControllerWtpProfileEslSesDongleComplianceLevel2edl(d, v, "compliance_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compliance-level"] = t
		}
	}

	if v, ok := d.GetOk("esl_channel"); ok || d.HasChange("esl_channel") {
		t, err := expandWirelessControllerWtpProfileEslSesDongleEslChannel2edl(d, v, "esl_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esl-channel"] = t
		}
	}

	if v, ok := d.GetOk("output_power"); ok || d.HasChange("output_power") {
		t, err := expandWirelessControllerWtpProfileEslSesDongleOutputPower2edl(d, v, "output_power")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output-power"] = t
		}
	}

	if v, ok := d.GetOk("scd_enable"); ok || d.HasChange("scd_enable") {
		t, err := expandWirelessControllerWtpProfileEslSesDongleScdEnable2edl(d, v, "scd_enable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scd-enable"] = t
		}
	}

	if v, ok := d.GetOk("tls_cert_verification"); ok || d.HasChange("tls_cert_verification") {
		t, err := expandWirelessControllerWtpProfileEslSesDongleTlsCertVerification2edl(d, v, "tls_cert_verification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-cert-verification"] = t
		}
	}

	if v, ok := d.GetOk("tls_fqdn_verification"); ok || d.HasChange("tls_fqdn_verification") {
		t, err := expandWirelessControllerWtpProfileEslSesDongleTlsFqdnVerification2edl(d, v, "tls_fqdn_verification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-fqdn-verification"] = t
		}
	}

	return &obj, nil
}
