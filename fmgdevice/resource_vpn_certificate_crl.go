// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Certificate Revocation List as a PEM file.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnCertificateCrl() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateCrlCreate,
		Read:   resourceVpnCertificateCrlRead,
		Update: resourceVpnCertificateCrlUpdate,
		Delete: resourceVpnCertificateCrlDelete,

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
			"crl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ldap_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ldap_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scep_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"scep_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"update_vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnCertificateCrlCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectVpnCertificateCrl(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateCrl resource while getting object: %v", err)
	}

	_, err = c.CreateVpnCertificateCrl(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateCrl resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnCertificateCrlRead(d, m)
}

func resourceVpnCertificateCrlUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectVpnCertificateCrl(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateCrl resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnCertificateCrl(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateCrl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnCertificateCrlRead(d, m)
}

func resourceVpnCertificateCrlDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteVpnCertificateCrl(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateCrl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateCrlRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnCertificateCrl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateCrl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateCrl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateCrl resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateCrlCrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCrlHttpUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCrlLastUpdated(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCrlLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateCrlLdapUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCrlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCrlRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCrlScepCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateCrlScepUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCrlSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCrlSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCrlUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCrlUpdateVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectVpnCertificateCrl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("crl", flattenVpnCertificateCrlCrl(o["crl"], d, "crl")); err != nil {
		if vv, ok := fortiAPIPatch(o["crl"], "VpnCertificateCrl-Crl"); ok {
			if err = d.Set("crl", vv); err != nil {
				return fmt.Errorf("Error reading crl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading crl: %v", err)
		}
	}

	if err = d.Set("http_url", flattenVpnCertificateCrlHttpUrl(o["http-url"], d, "http_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-url"], "VpnCertificateCrl-HttpUrl"); ok {
			if err = d.Set("http_url", vv); err != nil {
				return fmt.Errorf("Error reading http_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_url: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenVpnCertificateCrlLastUpdated(o["last-updated"], d, "last_updated")); err != nil {
		if vv, ok := fortiAPIPatch(o["last-updated"], "VpnCertificateCrl-LastUpdated"); ok {
			if err = d.Set("last_updated", vv); err != nil {
				return fmt.Errorf("Error reading last_updated: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenVpnCertificateCrlLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "VpnCertificateCrl-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("ldap_username", flattenVpnCertificateCrlLdapUsername(o["ldap-username"], d, "ldap_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-username"], "VpnCertificateCrl-LdapUsername"); ok {
			if err = d.Set("ldap_username", vv); err != nil {
				return fmt.Errorf("Error reading ldap_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_username: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnCertificateCrlName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnCertificateCrl-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("range", flattenVpnCertificateCrlRange(o["range"], d, "range")); err != nil {
		if vv, ok := fortiAPIPatch(o["range"], "VpnCertificateCrl-Range"); ok {
			if err = d.Set("range", vv); err != nil {
				return fmt.Errorf("Error reading range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("scep_cert", flattenVpnCertificateCrlScepCert(o["scep-cert"], d, "scep_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["scep-cert"], "VpnCertificateCrl-ScepCert"); ok {
			if err = d.Set("scep_cert", vv); err != nil {
				return fmt.Errorf("Error reading scep_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scep_cert: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenVpnCertificateCrlScepUrl(o["scep-url"], d, "scep_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["scep-url"], "VpnCertificateCrl-ScepUrl"); ok {
			if err = d.Set("scep_url", vv); err != nil {
				return fmt.Errorf("Error reading scep_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("source", flattenVpnCertificateCrlSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "VpnCertificateCrl-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenVpnCertificateCrlSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "VpnCertificateCrl-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenVpnCertificateCrlUpdateInterval(o["update-interval"], d, "update_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-interval"], "VpnCertificateCrl-UpdateInterval"); ok {
			if err = d.Set("update_interval", vv); err != nil {
				return fmt.Errorf("Error reading update_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("update_vdom", flattenVpnCertificateCrlUpdateVdom(o["update-vdom"], d, "update_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-vdom"], "VpnCertificateCrl-UpdateVdom"); ok {
			if err = d.Set("update_vdom", vv); err != nil {
				return fmt.Errorf("Error reading update_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_vdom: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateCrlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnCertificateCrlCrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlHttpUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlLastUpdated(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlLdapPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateCrlLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateCrlLdapUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlScepCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateCrlScepUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlUpdateInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlUpdateVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectVpnCertificateCrl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("crl"); ok || d.HasChange("crl") {
		t, err := expandVpnCertificateCrlCrl(d, v, "crl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crl"] = t
		}
	}

	if v, ok := d.GetOk("http_url"); ok || d.HasChange("http_url") {
		t, err := expandVpnCertificateCrlHttpUrl(d, v, "http_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-url"] = t
		}
	}

	if v, ok := d.GetOk("last_updated"); ok || d.HasChange("last_updated") {
		t, err := expandVpnCertificateCrlLastUpdated(d, v, "last_updated")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-updated"] = t
		}
	}

	if v, ok := d.GetOk("ldap_password"); ok || d.HasChange("ldap_password") {
		t, err := expandVpnCertificateCrlLdapPassword(d, v, "ldap_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-password"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandVpnCertificateCrlLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("ldap_username"); ok || d.HasChange("ldap_username") {
		t, err := expandVpnCertificateCrlLdapUsername(d, v, "ldap_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-username"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnCertificateCrlName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok || d.HasChange("range") {
		t, err := expandVpnCertificateCrlRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("scep_cert"); ok || d.HasChange("scep_cert") {
		t, err := expandVpnCertificateCrlScepCert(d, v, "scep_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-cert"] = t
		}
	}

	if v, ok := d.GetOk("scep_url"); ok || d.HasChange("scep_url") {
		t, err := expandVpnCertificateCrlScepUrl(d, v, "scep_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandVpnCertificateCrlSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandVpnCertificateCrlSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok || d.HasChange("update_interval") {
		t, err := expandVpnCertificateCrlUpdateInterval(d, v, "update_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	if v, ok := d.GetOk("update_vdom"); ok || d.HasChange("update_vdom") {
		t, err := expandVpnCertificateCrlUpdateVdom(d, v, "update_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-vdom"] = t
		}
	}

	return &obj, nil
}
