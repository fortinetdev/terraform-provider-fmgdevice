// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure SSL options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSslSshProfileSsl() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSslSshProfileSslUpdate,
		Read:   resourceFirewallSslSshProfileSslRead,
		Update: resourceFirewallSslSshProfileSslUpdate,
		Delete: resourceFirewallSslSshProfileSslDelete,

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
			"cert_probe_failure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert_validation_failure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert_validation_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encrypted_client_hello": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expired_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inspect_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_allowed_ssl_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revoked_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sni_server_cert_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"unsupported_ssl_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unsupported_ssl_negotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unsupported_ssl_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"untrusted_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallSslSshProfileSslUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallSslSshProfileSsl(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileSsl resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallSslSshProfileSsl(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileSsl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallSslSshProfileSsl")

	return resourceFirewallSslSshProfileSslRead(d, m)
}

func resourceFirewallSslSshProfileSslDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallSslSshProfileSsl(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSslSshProfileSsl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslSshProfileSslRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallSslSshProfileSsl(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallSslSshProfileSsl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslSshProfileSsl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSshProfileSsl resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslSshProfileSslCertProbeFailure2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslCertValidationFailure2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslCertValidationTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslClientCertificateU2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslEncryptedClientHello2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExpiredServerCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslInspectAll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslMinAllowedSslVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslRevokedServerCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslSniServerCertCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUnsupportedSslCipher2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUnsupportedSslNegotiation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUnsupportedSslVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUntrustedServerCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSslSshProfileSsl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cert_probe_failure", flattenFirewallSslSshProfileSslCertProbeFailure2edl(o["cert-probe-failure"], d, "cert_probe_failure")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-probe-failure"], "FirewallSslSshProfileSsl-CertProbeFailure"); ok {
			if err = d.Set("cert_probe_failure", vv); err != nil {
				return fmt.Errorf("Error reading cert_probe_failure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_probe_failure: %v", err)
		}
	}

	if err = d.Set("cert_validation_failure", flattenFirewallSslSshProfileSslCertValidationFailure2edl(o["cert-validation-failure"], d, "cert_validation_failure")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-validation-failure"], "FirewallSslSshProfileSsl-CertValidationFailure"); ok {
			if err = d.Set("cert_validation_failure", vv); err != nil {
				return fmt.Errorf("Error reading cert_validation_failure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_validation_failure: %v", err)
		}
	}

	if err = d.Set("cert_validation_timeout", flattenFirewallSslSshProfileSslCertValidationTimeout2edl(o["cert-validation-timeout"], d, "cert_validation_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-validation-timeout"], "FirewallSslSshProfileSsl-CertValidationTimeout"); ok {
			if err = d.Set("cert_validation_timeout", vv); err != nil {
				return fmt.Errorf("Error reading cert_validation_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_validation_timeout: %v", err)
		}
	}

	if err = d.Set("client_certificate", flattenFirewallSslSshProfileSslClientCertificateU2edl(o["client-certificate"], d, "client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-certificate"], "FirewallSslSshProfileSsl-ClientCertificate"); ok {
			if err = d.Set("client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_certificate: %v", err)
		}
	}

	if err = d.Set("encrypted_client_hello", flattenFirewallSslSshProfileSslEncryptedClientHello2edl(o["encrypted-client-hello"], d, "encrypted_client_hello")); err != nil {
		if vv, ok := fortiAPIPatch(o["encrypted-client-hello"], "FirewallSslSshProfileSsl-EncryptedClientHello"); ok {
			if err = d.Set("encrypted_client_hello", vv); err != nil {
				return fmt.Errorf("Error reading encrypted_client_hello: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encrypted_client_hello: %v", err)
		}
	}

	if err = d.Set("expired_server_cert", flattenFirewallSslSshProfileSslExpiredServerCert2edl(o["expired-server-cert"], d, "expired_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["expired-server-cert"], "FirewallSslSshProfileSsl-ExpiredServerCert"); ok {
			if err = d.Set("expired_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading expired_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expired_server_cert: %v", err)
		}
	}

	if err = d.Set("inspect_all", flattenFirewallSslSshProfileSslInspectAll2edl(o["inspect-all"], d, "inspect_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["inspect-all"], "FirewallSslSshProfileSsl-InspectAll"); ok {
			if err = d.Set("inspect_all", vv); err != nil {
				return fmt.Errorf("Error reading inspect_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inspect_all: %v", err)
		}
	}

	if err = d.Set("min_allowed_ssl_version", flattenFirewallSslSshProfileSslMinAllowedSslVersion2edl(o["min-allowed-ssl-version"], d, "min_allowed_ssl_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-allowed-ssl-version"], "FirewallSslSshProfileSsl-MinAllowedSslVersion"); ok {
			if err = d.Set("min_allowed_ssl_version", vv); err != nil {
				return fmt.Errorf("Error reading min_allowed_ssl_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_allowed_ssl_version: %v", err)
		}
	}

	if err = d.Set("revoked_server_cert", flattenFirewallSslSshProfileSslRevokedServerCert2edl(o["revoked-server-cert"], d, "revoked_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["revoked-server-cert"], "FirewallSslSshProfileSsl-RevokedServerCert"); ok {
			if err = d.Set("revoked_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading revoked_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading revoked_server_cert: %v", err)
		}
	}

	if err = d.Set("sni_server_cert_check", flattenFirewallSslSshProfileSslSniServerCertCheck2edl(o["sni-server-cert-check"], d, "sni_server_cert_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["sni-server-cert-check"], "FirewallSslSshProfileSsl-SniServerCertCheck"); ok {
			if err = d.Set("sni_server_cert_check", vv); err != nil {
				return fmt.Errorf("Error reading sni_server_cert_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sni_server_cert_check: %v", err)
		}
	}

	if err = d.Set("unsupported_ssl_cipher", flattenFirewallSslSshProfileSslUnsupportedSslCipher2edl(o["unsupported-ssl-cipher"], d, "unsupported_ssl_cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsupported-ssl-cipher"], "FirewallSslSshProfileSsl-UnsupportedSslCipher"); ok {
			if err = d.Set("unsupported_ssl_cipher", vv); err != nil {
				return fmt.Errorf("Error reading unsupported_ssl_cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsupported_ssl_cipher: %v", err)
		}
	}

	if err = d.Set("unsupported_ssl_negotiation", flattenFirewallSslSshProfileSslUnsupportedSslNegotiation2edl(o["unsupported-ssl-negotiation"], d, "unsupported_ssl_negotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsupported-ssl-negotiation"], "FirewallSslSshProfileSsl-UnsupportedSslNegotiation"); ok {
			if err = d.Set("unsupported_ssl_negotiation", vv); err != nil {
				return fmt.Errorf("Error reading unsupported_ssl_negotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsupported_ssl_negotiation: %v", err)
		}
	}

	if err = d.Set("unsupported_ssl_version", flattenFirewallSslSshProfileSslUnsupportedSslVersion2edl(o["unsupported-ssl-version"], d, "unsupported_ssl_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsupported-ssl-version"], "FirewallSslSshProfileSsl-UnsupportedSslVersion"); ok {
			if err = d.Set("unsupported_ssl_version", vv); err != nil {
				return fmt.Errorf("Error reading unsupported_ssl_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsupported_ssl_version: %v", err)
		}
	}

	if err = d.Set("untrusted_server_cert", flattenFirewallSslSshProfileSslUntrustedServerCert2edl(o["untrusted-server-cert"], d, "untrusted_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["untrusted-server-cert"], "FirewallSslSshProfileSsl-UntrustedServerCert"); ok {
			if err = d.Set("untrusted_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading untrusted_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading untrusted_server_cert: %v", err)
		}
	}

	return nil
}

func flattenFirewallSslSshProfileSslFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSslSshProfileSslCertProbeFailure2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslCertValidationFailure2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslCertValidationTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslClientCertificateU2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslEncryptedClientHello2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExpiredServerCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslInspectAll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslMinAllowedSslVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslRevokedServerCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslSniServerCertCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUnsupportedSslCipher2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUnsupportedSslNegotiation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUnsupportedSslVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUntrustedServerCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSslSshProfileSsl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cert_probe_failure"); ok || d.HasChange("cert_probe_failure") {
		t, err := expandFirewallSslSshProfileSslCertProbeFailure2edl(d, v, "cert_probe_failure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-probe-failure"] = t
		}
	}

	if v, ok := d.GetOk("cert_validation_failure"); ok || d.HasChange("cert_validation_failure") {
		t, err := expandFirewallSslSshProfileSslCertValidationFailure2edl(d, v, "cert_validation_failure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-validation-failure"] = t
		}
	}

	if v, ok := d.GetOk("cert_validation_timeout"); ok || d.HasChange("cert_validation_timeout") {
		t, err := expandFirewallSslSshProfileSslCertValidationTimeout2edl(d, v, "cert_validation_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-validation-timeout"] = t
		}
	}

	if v, ok := d.GetOk("client_certificate"); ok || d.HasChange("client_certificate") {
		t, err := expandFirewallSslSshProfileSslClientCertificateU2edl(d, v, "client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("encrypted_client_hello"); ok || d.HasChange("encrypted_client_hello") {
		t, err := expandFirewallSslSshProfileSslEncryptedClientHello2edl(d, v, "encrypted_client_hello")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encrypted-client-hello"] = t
		}
	}

	if v, ok := d.GetOk("expired_server_cert"); ok || d.HasChange("expired_server_cert") {
		t, err := expandFirewallSslSshProfileSslExpiredServerCert2edl(d, v, "expired_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expired-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("inspect_all"); ok || d.HasChange("inspect_all") {
		t, err := expandFirewallSslSshProfileSslInspectAll2edl(d, v, "inspect_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspect-all"] = t
		}
	}

	if v, ok := d.GetOk("min_allowed_ssl_version"); ok || d.HasChange("min_allowed_ssl_version") {
		t, err := expandFirewallSslSshProfileSslMinAllowedSslVersion2edl(d, v, "min_allowed_ssl_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-allowed-ssl-version"] = t
		}
	}

	if v, ok := d.GetOk("revoked_server_cert"); ok || d.HasChange("revoked_server_cert") {
		t, err := expandFirewallSslSshProfileSslRevokedServerCert2edl(d, v, "revoked_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revoked-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("sni_server_cert_check"); ok || d.HasChange("sni_server_cert_check") {
		t, err := expandFirewallSslSshProfileSslSniServerCertCheck2edl(d, v, "sni_server_cert_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni-server-cert-check"] = t
		}
	}

	if v, ok := d.GetOk("unsupported_ssl_cipher"); ok || d.HasChange("unsupported_ssl_cipher") {
		t, err := expandFirewallSslSshProfileSslUnsupportedSslCipher2edl(d, v, "unsupported_ssl_cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsupported-ssl-cipher"] = t
		}
	}

	if v, ok := d.GetOk("unsupported_ssl_negotiation"); ok || d.HasChange("unsupported_ssl_negotiation") {
		t, err := expandFirewallSslSshProfileSslUnsupportedSslNegotiation2edl(d, v, "unsupported_ssl_negotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsupported-ssl-negotiation"] = t
		}
	}

	if v, ok := d.GetOk("unsupported_ssl_version"); ok || d.HasChange("unsupported_ssl_version") {
		t, err := expandFirewallSslSshProfileSslUnsupportedSslVersion2edl(d, v, "unsupported_ssl_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsupported-ssl-version"] = t
		}
	}

	if v, ok := d.GetOk("untrusted_server_cert"); ok || d.HasChange("untrusted_server_cert") {
		t, err := expandFirewallSslSshProfileSslUntrustedServerCert2edl(d, v, "untrusted_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untrusted-server-cert"] = t
		}
	}

	return &obj, nil
}
