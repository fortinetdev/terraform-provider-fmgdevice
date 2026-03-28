// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> ClientHelloOuter SNIs to be blocked.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSslSshProfileEchOuterSni() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSslSshProfileEchOuterSniCreate,
		Read:   resourceFirewallSslSshProfileEchOuterSniRead,
		Update: resourceFirewallSslSshProfileEchOuterSniUpdate,
		Delete: resourceFirewallSslSshProfileEchOuterSniDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

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
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sni": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallSslSshProfileEchOuterSniCreate(d *schema.ResourceData, m interface{}) error {
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
	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	obj, err := getObjectFirewallSslSshProfileEchOuterSni(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSslSshProfileEchOuterSni resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallSslSshProfileEchOuterSni(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallSslSshProfileEchOuterSni(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallSslSshProfileEchOuterSni resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallSslSshProfileEchOuterSni(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallSslSshProfileEchOuterSni resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallSslSshProfileEchOuterSniRead(d, m)
}

func resourceFirewallSslSshProfileEchOuterSniUpdate(d *schema.ResourceData, m interface{}) error {
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
	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	obj, err := getObjectFirewallSslSshProfileEchOuterSni(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileEchOuterSni resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallSslSshProfileEchOuterSni(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileEchOuterSni resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallSslSshProfileEchOuterSniRead(d, m)
}

func resourceFirewallSslSshProfileEchOuterSniDelete(d *schema.ResourceData, m interface{}) error {
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
	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	wsParams["adom"] = adomv

	err = c.DeleteFirewallSslSshProfileEchOuterSni(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSslSshProfileEchOuterSni resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslSshProfileEchOuterSniRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
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
	if ssl_ssh_profile == "" {
		ssl_ssh_profile = importOptionChecking(m.(*FortiClient).Cfg, "ssl_ssh_profile")
		if ssl_ssh_profile == "" {
			return fmt.Errorf("Parameter ssl_ssh_profile is missing")
		}
		if err = d.Set("ssl_ssh_profile", ssl_ssh_profile); err != nil {
			return fmt.Errorf("Error set params ssl_ssh_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	o, err := c.ReadFirewallSslSshProfileEchOuterSni(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallSslSshProfileEchOuterSni resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslSshProfileEchOuterSni(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSshProfileEchOuterSni resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslSshProfileEchOuterSniName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileEchOuterSniSni2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSslSshProfileEchOuterSni(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallSslSshProfileEchOuterSniName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallSslSshProfileEchOuterSni-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("sni", flattenFirewallSslSshProfileEchOuterSniSni2edl(o["sni"], d, "sni")); err != nil {
		if vv, ok := fortiAPIPatch(o["sni"], "FirewallSslSshProfileEchOuterSni-Sni"); ok {
			if err = d.Set("sni", vv); err != nil {
				return fmt.Errorf("Error reading sni: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sni: %v", err)
		}
	}

	return nil
}

func flattenFirewallSslSshProfileEchOuterSniFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSslSshProfileEchOuterSniName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileEchOuterSniSni2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSslSshProfileEchOuterSni(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallSslSshProfileEchOuterSniName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("sni"); ok || d.HasChange("sni") {
		t, err := expandFirewallSslSshProfileEchOuterSniSni2edl(d, v, "sni")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni"] = t
		}
	}

	return &obj, nil
}
