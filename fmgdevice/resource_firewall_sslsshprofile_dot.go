// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure DNS over TLS options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSslSshProfileDot() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSslSshProfileDotUpdate,
		Read:   resourceFirewallSslSshProfileDotRead,
		Update: resourceFirewallSslSshProfileDotUpdate,
		Delete: resourceFirewallSslSshProfileDotDelete,

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
			"cert_validation_failure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_validation_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expired_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_allowed_ssl_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_after_tcp_handshake": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revoked_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sni_server_cert_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_not_quic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
		},
	}
}

func resourceFirewallSslSshProfileDotUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallSslSshProfileDot(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileDot resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallSslSshProfileDot(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileDot resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallSslSshProfileDot")

	return resourceFirewallSslSshProfileDotRead(d, m)
}

func resourceFirewallSslSshProfileDotDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallSslSshProfileDot(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSslSshProfileDot resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslSshProfileDotRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallSslSshProfileDot(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallSslSshProfileDot resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslSshProfileDot(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSshProfileDot resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslSshProfileDotCertValidationFailure2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotCertValidationTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotExpiredServerCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotMinAllowedSslVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotProxyAfterTcpHandshake2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotQuic2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotRevokedServerCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotSniServerCertCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUdpNotQuic2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUnsupportedSslCipher2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUnsupportedSslNegotiation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUnsupportedSslVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUntrustedServerCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSslSshProfileDot(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cert_validation_failure", flattenFirewallSslSshProfileDotCertValidationFailure2edl(o["cert-validation-failure"], d, "cert_validation_failure")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-validation-failure"], "FirewallSslSshProfileDot-CertValidationFailure"); ok {
			if err = d.Set("cert_validation_failure", vv); err != nil {
				return fmt.Errorf("Error reading cert_validation_failure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_validation_failure: %v", err)
		}
	}

	if err = d.Set("cert_validation_timeout", flattenFirewallSslSshProfileDotCertValidationTimeout2edl(o["cert-validation-timeout"], d, "cert_validation_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-validation-timeout"], "FirewallSslSshProfileDot-CertValidationTimeout"); ok {
			if err = d.Set("cert_validation_timeout", vv); err != nil {
				return fmt.Errorf("Error reading cert_validation_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_validation_timeout: %v", err)
		}
	}

	if err = d.Set("client_certificate", flattenFirewallSslSshProfileDotClientCertificate2edl(o["client-certificate"], d, "client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-certificate"], "FirewallSslSshProfileDot-ClientCertificate"); ok {
			if err = d.Set("client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_certificate: %v", err)
		}
	}

	if err = d.Set("expired_server_cert", flattenFirewallSslSshProfileDotExpiredServerCert2edl(o["expired-server-cert"], d, "expired_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["expired-server-cert"], "FirewallSslSshProfileDot-ExpiredServerCert"); ok {
			if err = d.Set("expired_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading expired_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expired_server_cert: %v", err)
		}
	}

	if err = d.Set("min_allowed_ssl_version", flattenFirewallSslSshProfileDotMinAllowedSslVersion2edl(o["min-allowed-ssl-version"], d, "min_allowed_ssl_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-allowed-ssl-version"], "FirewallSslSshProfileDot-MinAllowedSslVersion"); ok {
			if err = d.Set("min_allowed_ssl_version", vv); err != nil {
				return fmt.Errorf("Error reading min_allowed_ssl_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_allowed_ssl_version: %v", err)
		}
	}

	if err = d.Set("proxy_after_tcp_handshake", flattenFirewallSslSshProfileDotProxyAfterTcpHandshake2edl(o["proxy-after-tcp-handshake"], d, "proxy_after_tcp_handshake")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-after-tcp-handshake"], "FirewallSslSshProfileDot-ProxyAfterTcpHandshake"); ok {
			if err = d.Set("proxy_after_tcp_handshake", vv); err != nil {
				return fmt.Errorf("Error reading proxy_after_tcp_handshake: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_after_tcp_handshake: %v", err)
		}
	}

	if err = d.Set("quic", flattenFirewallSslSshProfileDotQuic2edl(o["quic"], d, "quic")); err != nil {
		if vv, ok := fortiAPIPatch(o["quic"], "FirewallSslSshProfileDot-Quic"); ok {
			if err = d.Set("quic", vv); err != nil {
				return fmt.Errorf("Error reading quic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quic: %v", err)
		}
	}

	if err = d.Set("revoked_server_cert", flattenFirewallSslSshProfileDotRevokedServerCert2edl(o["revoked-server-cert"], d, "revoked_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["revoked-server-cert"], "FirewallSslSshProfileDot-RevokedServerCert"); ok {
			if err = d.Set("revoked_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading revoked_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading revoked_server_cert: %v", err)
		}
	}

	if err = d.Set("sni_server_cert_check", flattenFirewallSslSshProfileDotSniServerCertCheck2edl(o["sni-server-cert-check"], d, "sni_server_cert_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["sni-server-cert-check"], "FirewallSslSshProfileDot-SniServerCertCheck"); ok {
			if err = d.Set("sni_server_cert_check", vv); err != nil {
				return fmt.Errorf("Error reading sni_server_cert_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sni_server_cert_check: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallSslSshProfileDotStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallSslSshProfileDot-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("udp_not_quic", flattenFirewallSslSshProfileDotUdpNotQuic2edl(o["udp-not-quic"], d, "udp_not_quic")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-not-quic"], "FirewallSslSshProfileDot-UdpNotQuic"); ok {
			if err = d.Set("udp_not_quic", vv); err != nil {
				return fmt.Errorf("Error reading udp_not_quic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_not_quic: %v", err)
		}
	}

	if err = d.Set("unsupported_ssl_cipher", flattenFirewallSslSshProfileDotUnsupportedSslCipher2edl(o["unsupported-ssl-cipher"], d, "unsupported_ssl_cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsupported-ssl-cipher"], "FirewallSslSshProfileDot-UnsupportedSslCipher"); ok {
			if err = d.Set("unsupported_ssl_cipher", vv); err != nil {
				return fmt.Errorf("Error reading unsupported_ssl_cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsupported_ssl_cipher: %v", err)
		}
	}

	if err = d.Set("unsupported_ssl_negotiation", flattenFirewallSslSshProfileDotUnsupportedSslNegotiation2edl(o["unsupported-ssl-negotiation"], d, "unsupported_ssl_negotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsupported-ssl-negotiation"], "FirewallSslSshProfileDot-UnsupportedSslNegotiation"); ok {
			if err = d.Set("unsupported_ssl_negotiation", vv); err != nil {
				return fmt.Errorf("Error reading unsupported_ssl_negotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsupported_ssl_negotiation: %v", err)
		}
	}

	if err = d.Set("unsupported_ssl_version", flattenFirewallSslSshProfileDotUnsupportedSslVersion2edl(o["unsupported-ssl-version"], d, "unsupported_ssl_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsupported-ssl-version"], "FirewallSslSshProfileDot-UnsupportedSslVersion"); ok {
			if err = d.Set("unsupported_ssl_version", vv); err != nil {
				return fmt.Errorf("Error reading unsupported_ssl_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsupported_ssl_version: %v", err)
		}
	}

	if err = d.Set("untrusted_server_cert", flattenFirewallSslSshProfileDotUntrustedServerCert2edl(o["untrusted-server-cert"], d, "untrusted_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["untrusted-server-cert"], "FirewallSslSshProfileDot-UntrustedServerCert"); ok {
			if err = d.Set("untrusted_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading untrusted_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading untrusted_server_cert: %v", err)
		}
	}

	return nil
}

func flattenFirewallSslSshProfileDotFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSslSshProfileDotCertValidationFailure2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotCertValidationTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotExpiredServerCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotMinAllowedSslVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotProxyAfterTcpHandshake2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotQuic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotRevokedServerCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotSniServerCertCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUdpNotQuic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUnsupportedSslCipher2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUnsupportedSslNegotiation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUnsupportedSslVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUntrustedServerCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSslSshProfileDot(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cert_validation_failure"); ok || d.HasChange("cert_validation_failure") {
		t, err := expandFirewallSslSshProfileDotCertValidationFailure2edl(d, v, "cert_validation_failure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-validation-failure"] = t
		}
	}

	if v, ok := d.GetOk("cert_validation_timeout"); ok || d.HasChange("cert_validation_timeout") {
		t, err := expandFirewallSslSshProfileDotCertValidationTimeout2edl(d, v, "cert_validation_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-validation-timeout"] = t
		}
	}

	if v, ok := d.GetOk("client_certificate"); ok || d.HasChange("client_certificate") {
		t, err := expandFirewallSslSshProfileDotClientCertificate2edl(d, v, "client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("expired_server_cert"); ok || d.HasChange("expired_server_cert") {
		t, err := expandFirewallSslSshProfileDotExpiredServerCert2edl(d, v, "expired_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expired-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("min_allowed_ssl_version"); ok || d.HasChange("min_allowed_ssl_version") {
		t, err := expandFirewallSslSshProfileDotMinAllowedSslVersion2edl(d, v, "min_allowed_ssl_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-allowed-ssl-version"] = t
		}
	}

	if v, ok := d.GetOk("proxy_after_tcp_handshake"); ok || d.HasChange("proxy_after_tcp_handshake") {
		t, err := expandFirewallSslSshProfileDotProxyAfterTcpHandshake2edl(d, v, "proxy_after_tcp_handshake")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-after-tcp-handshake"] = t
		}
	}

	if v, ok := d.GetOk("quic"); ok || d.HasChange("quic") {
		t, err := expandFirewallSslSshProfileDotQuic2edl(d, v, "quic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic"] = t
		}
	}

	if v, ok := d.GetOk("revoked_server_cert"); ok || d.HasChange("revoked_server_cert") {
		t, err := expandFirewallSslSshProfileDotRevokedServerCert2edl(d, v, "revoked_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revoked-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("sni_server_cert_check"); ok || d.HasChange("sni_server_cert_check") {
		t, err := expandFirewallSslSshProfileDotSniServerCertCheck2edl(d, v, "sni_server_cert_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni-server-cert-check"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallSslSshProfileDotStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("udp_not_quic"); ok || d.HasChange("udp_not_quic") {
		t, err := expandFirewallSslSshProfileDotUdpNotQuic2edl(d, v, "udp_not_quic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-not-quic"] = t
		}
	}

	if v, ok := d.GetOk("unsupported_ssl_cipher"); ok || d.HasChange("unsupported_ssl_cipher") {
		t, err := expandFirewallSslSshProfileDotUnsupportedSslCipher2edl(d, v, "unsupported_ssl_cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsupported-ssl-cipher"] = t
		}
	}

	if v, ok := d.GetOk("unsupported_ssl_negotiation"); ok || d.HasChange("unsupported_ssl_negotiation") {
		t, err := expandFirewallSslSshProfileDotUnsupportedSslNegotiation2edl(d, v, "unsupported_ssl_negotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsupported-ssl-negotiation"] = t
		}
	}

	if v, ok := d.GetOk("unsupported_ssl_version"); ok || d.HasChange("unsupported_ssl_version") {
		t, err := expandFirewallSslSshProfileDotUnsupportedSslVersion2edl(d, v, "unsupported_ssl_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsupported-ssl-version"] = t
		}
	}

	if v, ok := d.GetOk("untrusted_server_cert"); ok || d.HasChange("untrusted_server_cert") {
		t, err := expandFirewallSslSshProfileDotUntrustedServerCert2edl(d, v, "untrusted_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untrusted-server-cert"] = t
		}
	}

	return &obj, nil
}
