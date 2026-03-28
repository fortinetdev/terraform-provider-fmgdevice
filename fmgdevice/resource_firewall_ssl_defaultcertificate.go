// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device FirewallSslDefaultCertificate

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSslDefaultCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSslDefaultCertificateUpdate,
		Read:   resourceFirewallSslDefaultCertificateRead,
		Update: resourceFirewallSslDefaultCertificateUpdate,
		Delete: resourceFirewallSslDefaultCertificateDelete,

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
			"default_ca": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"default_server_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"default_untrusted_ca": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallSslDefaultCertificateUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectFirewallSslDefaultCertificate(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslDefaultCertificate resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallSslDefaultCertificate(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslDefaultCertificate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallSslDefaultCertificate")

	return resourceFirewallSslDefaultCertificateRead(d, m)
}

func resourceFirewallSslDefaultCertificateDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteFirewallSslDefaultCertificate(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSslDefaultCertificate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslDefaultCertificateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallSslDefaultCertificate(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallSslDefaultCertificate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslDefaultCertificate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslDefaultCertificate resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslDefaultCertificateDefaultCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslDefaultCertificateDefaultServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslDefaultCertificateDefaultUntrustedCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallSslDefaultCertificate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("default_ca", flattenFirewallSslDefaultCertificateDefaultCa(o["default-ca"], d, "default_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-ca"], "FirewallSslDefaultCertificate-DefaultCa"); ok {
			if err = d.Set("default_ca", vv); err != nil {
				return fmt.Errorf("Error reading default_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_ca: %v", err)
		}
	}

	if err = d.Set("default_server_cert", flattenFirewallSslDefaultCertificateDefaultServerCert(o["default-server-cert"], d, "default_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-server-cert"], "FirewallSslDefaultCertificate-DefaultServerCert"); ok {
			if err = d.Set("default_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading default_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_server_cert: %v", err)
		}
	}

	if err = d.Set("default_untrusted_ca", flattenFirewallSslDefaultCertificateDefaultUntrustedCa(o["default-untrusted-ca"], d, "default_untrusted_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-untrusted-ca"], "FirewallSslDefaultCertificate-DefaultUntrustedCa"); ok {
			if err = d.Set("default_untrusted_ca", vv); err != nil {
				return fmt.Errorf("Error reading default_untrusted_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_untrusted_ca: %v", err)
		}
	}

	return nil
}

func flattenFirewallSslDefaultCertificateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSslDefaultCertificateDefaultCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslDefaultCertificateDefaultServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslDefaultCertificateDefaultUntrustedCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallSslDefaultCertificate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_ca"); ok || d.HasChange("default_ca") {
		t, err := expandFirewallSslDefaultCertificateDefaultCa(d, v, "default_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-ca"] = t
		}
	}

	if v, ok := d.GetOk("default_server_cert"); ok || d.HasChange("default_server_cert") {
		t, err := expandFirewallSslDefaultCertificateDefaultServerCert(d, v, "default_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("default_untrusted_ca"); ok || d.HasChange("default_untrusted_ca") {
		t, err := expandFirewallSslDefaultCertificateDefaultUntrustedCa(d, v, "default_untrusted_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-untrusted-ca"] = t
		}
	}

	return &obj, nil
}
