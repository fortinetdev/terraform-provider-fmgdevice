// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Local keys and certificates.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnCertificateLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateLocalCreate,
		Read:   resourceVpnCertificateLocalRead,
		Update: resourceVpnCertificateLocalUpdate,
		Delete: resourceVpnCertificateLocalDelete,

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
			"acme_ca_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acme_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"acme_email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"acme_renew_window": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"acme_rsa_key_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_regenerate_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auto_regenerate_days_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ca_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cmp_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cmp_regeneration_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmp_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cmp_server_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"csr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enroll_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"est_ca_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"est_client_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"est_http_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"est_http_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"est_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"est_server_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"est_srp_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"est_srp_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ike_localid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ike_localid_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"name_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_key_retain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scep_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
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
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tmp_cert_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVpnCertificateLocalCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateLocal(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateLocal resource while getting object: %v", err)
	}

	_, err = c.CreateVpnCertificateLocal(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateLocal resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnCertificateLocalRead(d, m)
}

func resourceVpnCertificateLocalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateLocal(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateLocal resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnCertificateLocal(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnCertificateLocalRead(d, m)
}

func resourceVpnCertificateLocalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnCertificateLocal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateLocalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnCertificateLocal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateLocal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateLocal resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateLocalAcmeCaUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalAcmeDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalAcmeEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalAcmeRenewWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalAcmeRsaKeySize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalAutoRegenerateDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalAutoRegenerateDaysWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalCaIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalCmpPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalCmpRegenerationMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalCmpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalCmpServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateLocalComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalCsr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalEnrollProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstCaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateLocalEstHttpPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstHttpUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateLocalEstSrpPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstSrpUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalIkeLocalid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalIkeLocalidType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalLastUpdated(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalNameEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalPrivateKeyRetain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalScepUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateLocalTmpCertFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnCertificateLocal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("acme_ca_url", flattenVpnCertificateLocalAcmeCaUrl(o["acme-ca-url"], d, "acme_ca_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["acme-ca-url"], "VpnCertificateLocal-AcmeCaUrl"); ok {
			if err = d.Set("acme_ca_url", vv); err != nil {
				return fmt.Errorf("Error reading acme_ca_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acme_ca_url: %v", err)
		}
	}

	if err = d.Set("acme_domain", flattenVpnCertificateLocalAcmeDomain(o["acme-domain"], d, "acme_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["acme-domain"], "VpnCertificateLocal-AcmeDomain"); ok {
			if err = d.Set("acme_domain", vv); err != nil {
				return fmt.Errorf("Error reading acme_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acme_domain: %v", err)
		}
	}

	if err = d.Set("acme_email", flattenVpnCertificateLocalAcmeEmail(o["acme-email"], d, "acme_email")); err != nil {
		if vv, ok := fortiAPIPatch(o["acme-email"], "VpnCertificateLocal-AcmeEmail"); ok {
			if err = d.Set("acme_email", vv); err != nil {
				return fmt.Errorf("Error reading acme_email: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acme_email: %v", err)
		}
	}

	if err = d.Set("acme_renew_window", flattenVpnCertificateLocalAcmeRenewWindow(o["acme-renew-window"], d, "acme_renew_window")); err != nil {
		if vv, ok := fortiAPIPatch(o["acme-renew-window"], "VpnCertificateLocal-AcmeRenewWindow"); ok {
			if err = d.Set("acme_renew_window", vv); err != nil {
				return fmt.Errorf("Error reading acme_renew_window: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acme_renew_window: %v", err)
		}
	}

	if err = d.Set("acme_rsa_key_size", flattenVpnCertificateLocalAcmeRsaKeySize(o["acme-rsa-key-size"], d, "acme_rsa_key_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["acme-rsa-key-size"], "VpnCertificateLocal-AcmeRsaKeySize"); ok {
			if err = d.Set("acme_rsa_key_size", vv); err != nil {
				return fmt.Errorf("Error reading acme_rsa_key_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acme_rsa_key_size: %v", err)
		}
	}

	if err = d.Set("auto_regenerate_days", flattenVpnCertificateLocalAutoRegenerateDays(o["auto-regenerate-days"], d, "auto_regenerate_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-regenerate-days"], "VpnCertificateLocal-AutoRegenerateDays"); ok {
			if err = d.Set("auto_regenerate_days", vv); err != nil {
				return fmt.Errorf("Error reading auto_regenerate_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_regenerate_days: %v", err)
		}
	}

	if err = d.Set("auto_regenerate_days_warning", flattenVpnCertificateLocalAutoRegenerateDaysWarning(o["auto-regenerate-days-warning"], d, "auto_regenerate_days_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-regenerate-days-warning"], "VpnCertificateLocal-AutoRegenerateDaysWarning"); ok {
			if err = d.Set("auto_regenerate_days_warning", vv); err != nil {
				return fmt.Errorf("Error reading auto_regenerate_days_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_regenerate_days_warning: %v", err)
		}
	}

	if err = d.Set("ca_identifier", flattenVpnCertificateLocalCaIdentifier(o["ca-identifier"], d, "ca_identifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca-identifier"], "VpnCertificateLocal-CaIdentifier"); ok {
			if err = d.Set("ca_identifier", vv); err != nil {
				return fmt.Errorf("Error reading ca_identifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca_identifier: %v", err)
		}
	}

	if err = d.Set("certificate", flattenVpnCertificateLocalCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "VpnCertificateLocal-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("cmp_path", flattenVpnCertificateLocalCmpPath(o["cmp-path"], d, "cmp_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["cmp-path"], "VpnCertificateLocal-CmpPath"); ok {
			if err = d.Set("cmp_path", vv); err != nil {
				return fmt.Errorf("Error reading cmp_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cmp_path: %v", err)
		}
	}

	if err = d.Set("cmp_regeneration_method", flattenVpnCertificateLocalCmpRegenerationMethod(o["cmp-regeneration-method"], d, "cmp_regeneration_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["cmp-regeneration-method"], "VpnCertificateLocal-CmpRegenerationMethod"); ok {
			if err = d.Set("cmp_regeneration_method", vv); err != nil {
				return fmt.Errorf("Error reading cmp_regeneration_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cmp_regeneration_method: %v", err)
		}
	}

	if err = d.Set("cmp_server", flattenVpnCertificateLocalCmpServer(o["cmp-server"], d, "cmp_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["cmp-server"], "VpnCertificateLocal-CmpServer"); ok {
			if err = d.Set("cmp_server", vv); err != nil {
				return fmt.Errorf("Error reading cmp_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cmp_server: %v", err)
		}
	}

	if err = d.Set("cmp_server_cert", flattenVpnCertificateLocalCmpServerCert(o["cmp-server-cert"], d, "cmp_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cmp-server-cert"], "VpnCertificateLocal-CmpServerCert"); ok {
			if err = d.Set("cmp_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading cmp_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cmp_server_cert: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnCertificateLocalComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "VpnCertificateLocal-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("csr", flattenVpnCertificateLocalCsr(o["csr"], d, "csr")); err != nil {
		if vv, ok := fortiAPIPatch(o["csr"], "VpnCertificateLocal-Csr"); ok {
			if err = d.Set("csr", vv); err != nil {
				return fmt.Errorf("Error reading csr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading csr: %v", err)
		}
	}

	if err = d.Set("enroll_protocol", flattenVpnCertificateLocalEnrollProtocol(o["enroll-protocol"], d, "enroll_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["enroll-protocol"], "VpnCertificateLocal-EnrollProtocol"); ok {
			if err = d.Set("enroll_protocol", vv); err != nil {
				return fmt.Errorf("Error reading enroll_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enroll_protocol: %v", err)
		}
	}

	if err = d.Set("est_ca_id", flattenVpnCertificateLocalEstCaId(o["est-ca-id"], d, "est_ca_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-ca-id"], "VpnCertificateLocal-EstCaId"); ok {
			if err = d.Set("est_ca_id", vv); err != nil {
				return fmt.Errorf("Error reading est_ca_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_ca_id: %v", err)
		}
	}

	if err = d.Set("est_client_cert", flattenVpnCertificateLocalEstClientCert(o["est-client-cert"], d, "est_client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-client-cert"], "VpnCertificateLocal-EstClientCert"); ok {
			if err = d.Set("est_client_cert", vv); err != nil {
				return fmt.Errorf("Error reading est_client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_client_cert: %v", err)
		}
	}

	if err = d.Set("est_http_password", flattenVpnCertificateLocalEstHttpPassword(o["est-http-password"], d, "est_http_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-http-password"], "VpnCertificateLocal-EstHttpPassword"); ok {
			if err = d.Set("est_http_password", vv); err != nil {
				return fmt.Errorf("Error reading est_http_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_http_password: %v", err)
		}
	}

	if err = d.Set("est_http_username", flattenVpnCertificateLocalEstHttpUsername(o["est-http-username"], d, "est_http_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-http-username"], "VpnCertificateLocal-EstHttpUsername"); ok {
			if err = d.Set("est_http_username", vv); err != nil {
				return fmt.Errorf("Error reading est_http_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_http_username: %v", err)
		}
	}

	if err = d.Set("est_server", flattenVpnCertificateLocalEstServer(o["est-server"], d, "est_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-server"], "VpnCertificateLocal-EstServer"); ok {
			if err = d.Set("est_server", vv); err != nil {
				return fmt.Errorf("Error reading est_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_server: %v", err)
		}
	}

	if err = d.Set("est_server_cert", flattenVpnCertificateLocalEstServerCert(o["est-server-cert"], d, "est_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-server-cert"], "VpnCertificateLocal-EstServerCert"); ok {
			if err = d.Set("est_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading est_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_server_cert: %v", err)
		}
	}

	if err = d.Set("est_srp_password", flattenVpnCertificateLocalEstSrpPassword(o["est-srp-password"], d, "est_srp_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-srp-password"], "VpnCertificateLocal-EstSrpPassword"); ok {
			if err = d.Set("est_srp_password", vv); err != nil {
				return fmt.Errorf("Error reading est_srp_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_srp_password: %v", err)
		}
	}

	if err = d.Set("est_srp_username", flattenVpnCertificateLocalEstSrpUsername(o["est-srp-username"], d, "est_srp_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-srp-username"], "VpnCertificateLocal-EstSrpUsername"); ok {
			if err = d.Set("est_srp_username", vv); err != nil {
				return fmt.Errorf("Error reading est_srp_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_srp_username: %v", err)
		}
	}

	if err = d.Set("ike_localid", flattenVpnCertificateLocalIkeLocalid(o["ike-localid"], d, "ike_localid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-localid"], "VpnCertificateLocal-IkeLocalid"); ok {
			if err = d.Set("ike_localid", vv); err != nil {
				return fmt.Errorf("Error reading ike_localid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_localid: %v", err)
		}
	}

	if err = d.Set("ike_localid_type", flattenVpnCertificateLocalIkeLocalidType(o["ike-localid-type"], d, "ike_localid_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-localid-type"], "VpnCertificateLocal-IkeLocalidType"); ok {
			if err = d.Set("ike_localid_type", vv); err != nil {
				return fmt.Errorf("Error reading ike_localid_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_localid_type: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenVpnCertificateLocalLastUpdated(o["last-updated"], d, "last_updated")); err != nil {
		if vv, ok := fortiAPIPatch(o["last-updated"], "VpnCertificateLocal-LastUpdated"); ok {
			if err = d.Set("last_updated", vv); err != nil {
				return fmt.Errorf("Error reading last_updated: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnCertificateLocalName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnCertificateLocal-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("name_encoding", flattenVpnCertificateLocalNameEncoding(o["name-encoding"], d, "name_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["name-encoding"], "VpnCertificateLocal-NameEncoding"); ok {
			if err = d.Set("name_encoding", vv); err != nil {
				return fmt.Errorf("Error reading name_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name_encoding: %v", err)
		}
	}

	if err = d.Set("private_key", flattenVpnCertificateLocalPrivateKey(o["private-key"], d, "private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-key"], "VpnCertificateLocal-PrivateKey"); ok {
			if err = d.Set("private_key", vv); err != nil {
				return fmt.Errorf("Error reading private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_key: %v", err)
		}
	}

	if err = d.Set("private_key_retain", flattenVpnCertificateLocalPrivateKeyRetain(o["private-key-retain"], d, "private_key_retain")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-key-retain"], "VpnCertificateLocal-PrivateKeyRetain"); ok {
			if err = d.Set("private_key_retain", vv); err != nil {
				return fmt.Errorf("Error reading private_key_retain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_key_retain: %v", err)
		}
	}

	if err = d.Set("range", flattenVpnCertificateLocalRange(o["range"], d, "range")); err != nil {
		if vv, ok := fortiAPIPatch(o["range"], "VpnCertificateLocal-Range"); ok {
			if err = d.Set("range", vv); err != nil {
				return fmt.Errorf("Error reading range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenVpnCertificateLocalScepUrl(o["scep-url"], d, "scep_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["scep-url"], "VpnCertificateLocal-ScepUrl"); ok {
			if err = d.Set("scep_url", vv); err != nil {
				return fmt.Errorf("Error reading scep_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("source", flattenVpnCertificateLocalSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "VpnCertificateLocal-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenVpnCertificateLocalSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "VpnCertificateLocal-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("state", flattenVpnCertificateLocalState(o["state"], d, "state")); err != nil {
		if vv, ok := fortiAPIPatch(o["state"], "VpnCertificateLocal-State"); ok {
			if err = d.Set("state", vv); err != nil {
				return fmt.Errorf("Error reading state: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading state: %v", err)
		}
	}

	if err = d.Set("tmp_cert_file", flattenVpnCertificateLocalTmpCertFile(o["tmp-cert-file"], d, "tmp_cert_file")); err != nil {
		if vv, ok := fortiAPIPatch(o["tmp-cert-file"], "VpnCertificateLocal-TmpCertFile"); ok {
			if err = d.Set("tmp_cert_file", vv); err != nil {
				return fmt.Errorf("Error reading tmp_cert_file: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tmp_cert_file: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnCertificateLocalAcmeCaUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAcmeDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAcmeEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAcmeRenewWindow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAcmeRsaKeySize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAutoRegenerateDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAutoRegenerateDaysWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCaIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCmpPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCmpRegenerationMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCmpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCmpServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateLocalComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCsr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEnrollProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstCaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateLocalEstHttpPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstHttpUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateLocalEstSrpPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstSrpUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalIkeLocalid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalIkeLocalidType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalLastUpdated(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalNameEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateLocalPrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalPrivateKeyRetain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalScepPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateLocalScepUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalTmpCertFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateLocal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("acme_ca_url"); ok || d.HasChange("acme_ca_url") {
		t, err := expandVpnCertificateLocalAcmeCaUrl(d, v, "acme_ca_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-ca-url"] = t
		}
	}

	if v, ok := d.GetOk("acme_domain"); ok || d.HasChange("acme_domain") {
		t, err := expandVpnCertificateLocalAcmeDomain(d, v, "acme_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-domain"] = t
		}
	}

	if v, ok := d.GetOk("acme_email"); ok || d.HasChange("acme_email") {
		t, err := expandVpnCertificateLocalAcmeEmail(d, v, "acme_email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-email"] = t
		}
	}

	if v, ok := d.GetOk("acme_renew_window"); ok || d.HasChange("acme_renew_window") {
		t, err := expandVpnCertificateLocalAcmeRenewWindow(d, v, "acme_renew_window")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-renew-window"] = t
		}
	}

	if v, ok := d.GetOk("acme_rsa_key_size"); ok || d.HasChange("acme_rsa_key_size") {
		t, err := expandVpnCertificateLocalAcmeRsaKeySize(d, v, "acme_rsa_key_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-rsa-key-size"] = t
		}
	}

	if v, ok := d.GetOk("auto_regenerate_days"); ok || d.HasChange("auto_regenerate_days") {
		t, err := expandVpnCertificateLocalAutoRegenerateDays(d, v, "auto_regenerate_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-regenerate-days"] = t
		}
	}

	if v, ok := d.GetOk("auto_regenerate_days_warning"); ok || d.HasChange("auto_regenerate_days_warning") {
		t, err := expandVpnCertificateLocalAutoRegenerateDaysWarning(d, v, "auto_regenerate_days_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-regenerate-days-warning"] = t
		}
	}

	if v, ok := d.GetOk("ca_identifier"); ok || d.HasChange("ca_identifier") {
		t, err := expandVpnCertificateLocalCaIdentifier(d, v, "ca_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-identifier"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandVpnCertificateLocalCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("cmp_path"); ok || d.HasChange("cmp_path") {
		t, err := expandVpnCertificateLocalCmpPath(d, v, "cmp_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-path"] = t
		}
	}

	if v, ok := d.GetOk("cmp_regeneration_method"); ok || d.HasChange("cmp_regeneration_method") {
		t, err := expandVpnCertificateLocalCmpRegenerationMethod(d, v, "cmp_regeneration_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-regeneration-method"] = t
		}
	}

	if v, ok := d.GetOk("cmp_server"); ok || d.HasChange("cmp_server") {
		t, err := expandVpnCertificateLocalCmpServer(d, v, "cmp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-server"] = t
		}
	}

	if v, ok := d.GetOk("cmp_server_cert"); ok || d.HasChange("cmp_server_cert") {
		t, err := expandVpnCertificateLocalCmpServerCert(d, v, "cmp_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandVpnCertificateLocalComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("csr"); ok || d.HasChange("csr") {
		t, err := expandVpnCertificateLocalCsr(d, v, "csr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csr"] = t
		}
	}

	if v, ok := d.GetOk("enroll_protocol"); ok || d.HasChange("enroll_protocol") {
		t, err := expandVpnCertificateLocalEnrollProtocol(d, v, "enroll_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enroll-protocol"] = t
		}
	}

	if v, ok := d.GetOk("est_ca_id"); ok || d.HasChange("est_ca_id") {
		t, err := expandVpnCertificateLocalEstCaId(d, v, "est_ca_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-ca-id"] = t
		}
	}

	if v, ok := d.GetOk("est_client_cert"); ok || d.HasChange("est_client_cert") {
		t, err := expandVpnCertificateLocalEstClientCert(d, v, "est_client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-client-cert"] = t
		}
	}

	if v, ok := d.GetOk("est_http_password"); ok || d.HasChange("est_http_password") {
		t, err := expandVpnCertificateLocalEstHttpPassword(d, v, "est_http_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-http-password"] = t
		}
	}

	if v, ok := d.GetOk("est_http_username"); ok || d.HasChange("est_http_username") {
		t, err := expandVpnCertificateLocalEstHttpUsername(d, v, "est_http_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-http-username"] = t
		}
	}

	if v, ok := d.GetOk("est_server"); ok || d.HasChange("est_server") {
		t, err := expandVpnCertificateLocalEstServer(d, v, "est_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-server"] = t
		}
	}

	if v, ok := d.GetOk("est_server_cert"); ok || d.HasChange("est_server_cert") {
		t, err := expandVpnCertificateLocalEstServerCert(d, v, "est_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("est_srp_password"); ok || d.HasChange("est_srp_password") {
		t, err := expandVpnCertificateLocalEstSrpPassword(d, v, "est_srp_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-srp-password"] = t
		}
	}

	if v, ok := d.GetOk("est_srp_username"); ok || d.HasChange("est_srp_username") {
		t, err := expandVpnCertificateLocalEstSrpUsername(d, v, "est_srp_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-srp-username"] = t
		}
	}

	if v, ok := d.GetOk("ike_localid"); ok || d.HasChange("ike_localid") {
		t, err := expandVpnCertificateLocalIkeLocalid(d, v, "ike_localid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-localid"] = t
		}
	}

	if v, ok := d.GetOk("ike_localid_type"); ok || d.HasChange("ike_localid_type") {
		t, err := expandVpnCertificateLocalIkeLocalidType(d, v, "ike_localid_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-localid-type"] = t
		}
	}

	if v, ok := d.GetOk("last_updated"); ok || d.HasChange("last_updated") {
		t, err := expandVpnCertificateLocalLastUpdated(d, v, "last_updated")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-updated"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnCertificateLocalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("name_encoding"); ok || d.HasChange("name_encoding") {
		t, err := expandVpnCertificateLocalNameEncoding(d, v, "name_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name-encoding"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandVpnCertificateLocalPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok || d.HasChange("private_key") {
		t, err := expandVpnCertificateLocalPrivateKey(d, v, "private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	}

	if v, ok := d.GetOk("private_key_retain"); ok || d.HasChange("private_key_retain") {
		t, err := expandVpnCertificateLocalPrivateKeyRetain(d, v, "private_key_retain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key-retain"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok || d.HasChange("range") {
		t, err := expandVpnCertificateLocalRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("scep_password"); ok || d.HasChange("scep_password") {
		t, err := expandVpnCertificateLocalScepPassword(d, v, "scep_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-password"] = t
		}
	}

	if v, ok := d.GetOk("scep_url"); ok || d.HasChange("scep_url") {
		t, err := expandVpnCertificateLocalScepUrl(d, v, "scep_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandVpnCertificateLocalSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandVpnCertificateLocalSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("state"); ok || d.HasChange("state") {
		t, err := expandVpnCertificateLocalState(d, v, "state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["state"] = t
		}
	}

	if v, ok := d.GetOk("tmp_cert_file"); ok || d.HasChange("tmp_cert_file") {
		t, err := expandVpnCertificateLocalTmpCertFile(d, v, "tmp_cert_file")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tmp-cert-file"] = t
		}
	}

	return &obj, nil
}
