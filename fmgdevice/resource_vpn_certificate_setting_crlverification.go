// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CRL verification options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnCertificateSettingCrlVerification() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateSettingCrlVerificationUpdate,
		Read:   resourceVpnCertificateSettingCrlVerificationRead,
		Update: resourceVpnCertificateSettingCrlVerificationUpdate,
		Delete: resourceVpnCertificateSettingCrlVerificationDelete,

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
			"chain_crl_absence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"leaf_crl_absence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnCertificateSettingCrlVerificationUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectVpnCertificateSettingCrlVerification(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateSettingCrlVerification resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnCertificateSettingCrlVerification(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateSettingCrlVerification resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("VpnCertificateSettingCrlVerification")

	return resourceVpnCertificateSettingCrlVerificationRead(d, m)
}

func resourceVpnCertificateSettingCrlVerificationDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteVpnCertificateSettingCrlVerification(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateSettingCrlVerification resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateSettingCrlVerificationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnCertificateSettingCrlVerification(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateSettingCrlVerification resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateSettingCrlVerification(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateSettingCrlVerification resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateSettingCrlVerificationChainCrlAbsence2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCrlVerificationExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCrlVerificationLeafCrlAbsence2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnCertificateSettingCrlVerification(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("chain_crl_absence", flattenVpnCertificateSettingCrlVerificationChainCrlAbsence2edl(o["chain-crl-absence"], d, "chain_crl_absence")); err != nil {
		if vv, ok := fortiAPIPatch(o["chain-crl-absence"], "VpnCertificateSettingCrlVerification-ChainCrlAbsence"); ok {
			if err = d.Set("chain_crl_absence", vv); err != nil {
				return fmt.Errorf("Error reading chain_crl_absence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading chain_crl_absence: %v", err)
		}
	}

	if err = d.Set("expiry", flattenVpnCertificateSettingCrlVerificationExpiry2edl(o["expiry"], d, "expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["expiry"], "VpnCertificateSettingCrlVerification-Expiry"); ok {
			if err = d.Set("expiry", vv); err != nil {
				return fmt.Errorf("Error reading expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expiry: %v", err)
		}
	}

	if err = d.Set("leaf_crl_absence", flattenVpnCertificateSettingCrlVerificationLeafCrlAbsence2edl(o["leaf-crl-absence"], d, "leaf_crl_absence")); err != nil {
		if vv, ok := fortiAPIPatch(o["leaf-crl-absence"], "VpnCertificateSettingCrlVerification-LeafCrlAbsence"); ok {
			if err = d.Set("leaf_crl_absence", vv); err != nil {
				return fmt.Errorf("Error reading leaf_crl_absence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading leaf_crl_absence: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateSettingCrlVerificationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnCertificateSettingCrlVerificationChainCrlAbsence2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCrlVerificationExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCrlVerificationLeafCrlAbsence2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateSettingCrlVerification(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("chain_crl_absence"); ok || d.HasChange("chain_crl_absence") {
		t, err := expandVpnCertificateSettingCrlVerificationChainCrlAbsence2edl(d, v, "chain_crl_absence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chain-crl-absence"] = t
		}
	}

	if v, ok := d.GetOk("expiry"); ok || d.HasChange("expiry") {
		t, err := expandVpnCertificateSettingCrlVerificationExpiry2edl(d, v, "expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expiry"] = t
		}
	}

	if v, ok := d.GetOk("leaf_crl_absence"); ok || d.HasChange("leaf_crl_absence") {
		t, err := expandVpnCertificateSettingCrlVerificationLeafCrlAbsence2edl(d, v, "leaf_crl_absence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["leaf-crl-absence"] = t
		}
	}

	return &obj, nil
}
