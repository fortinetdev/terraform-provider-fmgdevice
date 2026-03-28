// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Address block and allow lists.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfileAddressList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileAddressListUpdate,
		Read:   resourceWafProfileAddressListRead,
		Update: resourceWafProfileAddressListUpdate,
		Delete: resourceWafProfileAddressListDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"blocked_address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"blocked_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted_address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWafProfileAddressListUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectWafProfileAddressList(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileAddressList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfileAddressList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileAddressList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WafProfileAddressList")

	return resourceWafProfileAddressListRead(d, m)
}

func resourceWafProfileAddressListDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteWafProfileAddressList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfileAddressList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileAddressListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadWafProfileAddressList(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfileAddressList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfileAddressList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfileAddressList resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileAddressListBlockedAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileAddressListBlockedLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileAddressListSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileAddressListStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileAddressListTrustedAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWafProfileAddressList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("blocked_address", flattenWafProfileAddressListBlockedAddress2edl(o["blocked-address"], d, "blocked_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["blocked-address"], "WafProfileAddressList-BlockedAddress"); ok {
			if err = d.Set("blocked_address", vv); err != nil {
				return fmt.Errorf("Error reading blocked_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blocked_address: %v", err)
		}
	}

	if err = d.Set("blocked_log", flattenWafProfileAddressListBlockedLog2edl(o["blocked-log"], d, "blocked_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["blocked-log"], "WafProfileAddressList-BlockedLog"); ok {
			if err = d.Set("blocked_log", vv); err != nil {
				return fmt.Errorf("Error reading blocked_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blocked_log: %v", err)
		}
	}

	if err = d.Set("severity", flattenWafProfileAddressListSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "WafProfileAddressList-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenWafProfileAddressListStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WafProfileAddressList-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trusted_address", flattenWafProfileAddressListTrustedAddress2edl(o["trusted-address"], d, "trusted_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusted-address"], "WafProfileAddressList-TrustedAddress"); ok {
			if err = d.Set("trusted_address", vv); err != nil {
				return fmt.Errorf("Error reading trusted_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusted_address: %v", err)
		}
	}

	return nil
}

func flattenWafProfileAddressListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileAddressListBlockedAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileAddressListBlockedLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileAddressListSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileAddressListStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileAddressListTrustedAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWafProfileAddressList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("blocked_address"); ok || d.HasChange("blocked_address") {
		t, err := expandWafProfileAddressListBlockedAddress2edl(d, v, "blocked_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blocked-address"] = t
		}
	}

	if v, ok := d.GetOk("blocked_log"); ok || d.HasChange("blocked_log") {
		t, err := expandWafProfileAddressListBlockedLog2edl(d, v, "blocked_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blocked-log"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandWafProfileAddressListSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWafProfileAddressListStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trusted_address"); ok || d.HasChange("trusted_address") {
		t, err := expandWafProfileAddressListTrustedAddress2edl(d, v, "trusted_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-address"] = t
		}
	}

	return &obj, nil
}
