// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSslSshProfileSslClientCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSslSshProfileSslClientCertificateUpdate,
		Read:   resourceFirewallSslSshProfileSslClientCertificateRead,
		Update: resourceFirewallSslSshProfileSslClientCertificateUpdate,
		Delete: resourceFirewallSslSshProfileSslClientCertificateDelete,

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
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"caname": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"keyring_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallSslSshProfileSslClientCertificateUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallSslSshProfileSslClientCertificate(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileSslClientCertificate resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallSslSshProfileSslClientCertificate(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileSslClientCertificate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallSslSshProfileSslClientCertificate")

	return resourceFirewallSslSshProfileSslClientCertificateRead(d, m)
}

func resourceFirewallSslSshProfileSslClientCertificateDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallSslSshProfileSslClientCertificate(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSslSshProfileSslClientCertificate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslSshProfileSslClientCertificateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallSslSshProfileSslClientCertificate(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallSslSshProfileSslClientCertificate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslSshProfileSslClientCertificate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSshProfileSslClientCertificate resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslSshProfileSslClientCertificateCaname2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslClientCertificateCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslClientCertificateKeyringList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslClientCertificateStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSslSshProfileSslClientCertificate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("caname", flattenFirewallSslSshProfileSslClientCertificateCaname2edl(o["caname"], d, "caname")); err != nil {
		if vv, ok := fortiAPIPatch(o["caname"], "FirewallSslSshProfileSslClientCertificate-Caname"); ok {
			if err = d.Set("caname", vv); err != nil {
				return fmt.Errorf("Error reading caname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading caname: %v", err)
		}
	}

	if err = d.Set("cert", flattenFirewallSslSshProfileSslClientCertificateCert2edl(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "FirewallSslSshProfileSslClientCertificate-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("keyring_list", flattenFirewallSslSshProfileSslClientCertificateKeyringList2edl(o["keyring-list"], d, "keyring_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["keyring-list"], "FirewallSslSshProfileSslClientCertificate-KeyringList"); ok {
			if err = d.Set("keyring_list", vv); err != nil {
				return fmt.Errorf("Error reading keyring_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keyring_list: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallSslSshProfileSslClientCertificateStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallSslSshProfileSslClientCertificate-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFirewallSslSshProfileSslClientCertificateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSslSshProfileSslClientCertificateCaname2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslClientCertificateCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslClientCertificateKeyringList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslClientCertificateStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSslSshProfileSslClientCertificate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("caname"); ok || d.HasChange("caname") {
		t, err := expandFirewallSslSshProfileSslClientCertificateCaname2edl(d, v, "caname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caname"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandFirewallSslSshProfileSslClientCertificateCert2edl(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("keyring_list"); ok || d.HasChange("keyring_list") {
		t, err := expandFirewallSslSshProfileSslClientCertificateKeyringList2edl(d, v, "keyring_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyring-list"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallSslSshProfileSslClientCertificateStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
