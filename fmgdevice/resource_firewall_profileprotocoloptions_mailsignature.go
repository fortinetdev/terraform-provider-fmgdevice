// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Mail signature.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProfileProtocolOptionsMailSignature() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProfileProtocolOptionsMailSignatureUpdate,
		Read:   resourceFirewallProfileProtocolOptionsMailSignatureRead,
		Update: resourceFirewallProfileProtocolOptionsMailSignatureUpdate,
		Delete: resourceFirewallProfileProtocolOptionsMailSignatureDelete,

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
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"signature": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceFirewallProfileProtocolOptionsMailSignatureUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	obj, err := getObjectFirewallProfileProtocolOptionsMailSignature(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptionsMailSignature resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallProfileProtocolOptionsMailSignature(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptionsMailSignature resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallProfileProtocolOptionsMailSignature")

	return resourceFirewallProfileProtocolOptionsMailSignatureRead(d, m)
}

func resourceFirewallProfileProtocolOptionsMailSignatureDelete(d *schema.ResourceData, m interface{}) error {
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
	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	wsParams["adom"] = adomv

	err = c.DeleteFirewallProfileProtocolOptionsMailSignature(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProfileProtocolOptionsMailSignature resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProfileProtocolOptionsMailSignatureRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile_protocol_options := d.Get("profile_protocol_options").(string)
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
	if profile_protocol_options == "" {
		profile_protocol_options = importOptionChecking(m.(*FortiClient).Cfg, "profile_protocol_options")
		if profile_protocol_options == "" {
			return fmt.Errorf("Parameter profile_protocol_options is missing")
		}
		if err = d.Set("profile_protocol_options", profile_protocol_options); err != nil {
			return fmt.Errorf("Error set params profile_protocol_options: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	o, err := c.ReadFirewallProfileProtocolOptionsMailSignature(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallProfileProtocolOptionsMailSignature resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProfileProtocolOptionsMailSignature(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProfileProtocolOptionsMailSignature resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProfileProtocolOptionsMailSignatureSignature2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMailSignatureStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallProfileProtocolOptionsMailSignature(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("signature", flattenFirewallProfileProtocolOptionsMailSignatureSignature2edl(o["signature"], d, "signature")); err != nil {
		if vv, ok := fortiAPIPatch(o["signature"], "FirewallProfileProtocolOptionsMailSignature-Signature"); ok {
			if err = d.Set("signature", vv); err != nil {
				return fmt.Errorf("Error reading signature: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signature: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallProfileProtocolOptionsMailSignatureStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallProfileProtocolOptionsMailSignature-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFirewallProfileProtocolOptionsMailSignatureFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProfileProtocolOptionsMailSignatureSignature2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMailSignatureStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProfileProtocolOptionsMailSignature(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("signature"); ok || d.HasChange("signature") {
		t, err := expandFirewallProfileProtocolOptionsMailSignatureSignature2edl(d, v, "signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallProfileProtocolOptionsMailSignatureStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
