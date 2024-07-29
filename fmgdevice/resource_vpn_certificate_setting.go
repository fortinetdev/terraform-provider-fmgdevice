// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: VPN certificate setting.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnCertificateSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateSettingUpdate,
		Read:   resourceVpnCertificateSettingRead,
		Update: resourceVpnCertificateSettingUpdate,
		Delete: resourceVpnCertificateSettingDelete,

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
			"cert_expire_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"certname_dsa1024": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"certname_dsa2048": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"certname_ecdsa256": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"certname_ecdsa384": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"certname_ecdsa521": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"certname_ed25519": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"certname_ed448": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"certname_rsa1024": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"certname_rsa2048": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"certname_rsa4096": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"check_ca_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_ca_chain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmp_key_usage_checking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmp_save_extra_certs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cn_allow_multi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cn_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crl_verification": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocsp_default_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ocsp_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocsp_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"proxy_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxy_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_ocsp_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strict_crl_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"strict_ocsp_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subject_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subject_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnCertificateSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnCertificateSetting(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("VpnCertificateSetting")

	return resourceVpnCertificateSettingRead(d, m)
}

func resourceVpnCertificateSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnCertificateSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnCertificateSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateSetting resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateSettingCertExpireWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameDsa1024(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateSettingCertnameDsa2048(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateSettingCertnameEcdsa256(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateSettingCertnameEcdsa384(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateSettingCertnameEcdsa521(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateSettingCertnameEd25519(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateSettingCertnameEd448(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateSettingCertnameRsa1024(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateSettingCertnameRsa2048(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateSettingCertnameRsa4096(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateSettingCheckCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCheckCaChain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCmpKeyUsageChecking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCmpSaveExtraCerts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCnAllowMulti(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCnMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCrlVerification(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "chain_crl_absence"
	if _, ok := i["chain-crl-absence"]; ok {
		result["chain_crl_absence"] = flattenVpnCertificateSettingCrlVerificationChainCrlAbsence(i["chain-crl-absence"], d, pre_append)
	}

	pre_append = pre + ".0." + "expiry"
	if _, ok := i["expiry"]; ok {
		result["expiry"] = flattenVpnCertificateSettingCrlVerificationExpiry(i["expiry"], d, pre_append)
	}

	pre_append = pre + ".0." + "leaf_crl_absence"
	if _, ok := i["leaf-crl-absence"]; ok {
		result["leaf_crl_absence"] = flattenVpnCertificateSettingCrlVerificationLeafCrlAbsence(i["leaf-crl-absence"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenVpnCertificateSettingCrlVerificationChainCrlAbsence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCrlVerificationExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCrlVerificationLeafCrlAbsence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateSettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingOcspDefaultServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateSettingOcspOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingOcspStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingProxyPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingProxyUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingSslOcspSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingStrictCrlCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingStrictOcspCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingSubjectMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingSubjectSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnCertificateSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cert_expire_warning", flattenVpnCertificateSettingCertExpireWarning(o["cert-expire-warning"], d, "cert_expire_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-expire-warning"], "VpnCertificateSetting-CertExpireWarning"); ok {
			if err = d.Set("cert_expire_warning", vv); err != nil {
				return fmt.Errorf("Error reading cert_expire_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_expire_warning: %v", err)
		}
	}

	if err = d.Set("certname_dsa1024", flattenVpnCertificateSettingCertnameDsa1024(o["certname-dsa1024"], d, "certname_dsa1024")); err != nil {
		if vv, ok := fortiAPIPatch(o["certname-dsa1024"], "VpnCertificateSetting-CertnameDsa1024"); ok {
			if err = d.Set("certname_dsa1024", vv); err != nil {
				return fmt.Errorf("Error reading certname_dsa1024: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certname_dsa1024: %v", err)
		}
	}

	if err = d.Set("certname_dsa2048", flattenVpnCertificateSettingCertnameDsa2048(o["certname-dsa2048"], d, "certname_dsa2048")); err != nil {
		if vv, ok := fortiAPIPatch(o["certname-dsa2048"], "VpnCertificateSetting-CertnameDsa2048"); ok {
			if err = d.Set("certname_dsa2048", vv); err != nil {
				return fmt.Errorf("Error reading certname_dsa2048: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certname_dsa2048: %v", err)
		}
	}

	if err = d.Set("certname_ecdsa256", flattenVpnCertificateSettingCertnameEcdsa256(o["certname-ecdsa256"], d, "certname_ecdsa256")); err != nil {
		if vv, ok := fortiAPIPatch(o["certname-ecdsa256"], "VpnCertificateSetting-CertnameEcdsa256"); ok {
			if err = d.Set("certname_ecdsa256", vv); err != nil {
				return fmt.Errorf("Error reading certname_ecdsa256: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certname_ecdsa256: %v", err)
		}
	}

	if err = d.Set("certname_ecdsa384", flattenVpnCertificateSettingCertnameEcdsa384(o["certname-ecdsa384"], d, "certname_ecdsa384")); err != nil {
		if vv, ok := fortiAPIPatch(o["certname-ecdsa384"], "VpnCertificateSetting-CertnameEcdsa384"); ok {
			if err = d.Set("certname_ecdsa384", vv); err != nil {
				return fmt.Errorf("Error reading certname_ecdsa384: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certname_ecdsa384: %v", err)
		}
	}

	if err = d.Set("certname_ecdsa521", flattenVpnCertificateSettingCertnameEcdsa521(o["certname-ecdsa521"], d, "certname_ecdsa521")); err != nil {
		if vv, ok := fortiAPIPatch(o["certname-ecdsa521"], "VpnCertificateSetting-CertnameEcdsa521"); ok {
			if err = d.Set("certname_ecdsa521", vv); err != nil {
				return fmt.Errorf("Error reading certname_ecdsa521: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certname_ecdsa521: %v", err)
		}
	}

	if err = d.Set("certname_ed25519", flattenVpnCertificateSettingCertnameEd25519(o["certname-ed25519"], d, "certname_ed25519")); err != nil {
		if vv, ok := fortiAPIPatch(o["certname-ed25519"], "VpnCertificateSetting-CertnameEd25519"); ok {
			if err = d.Set("certname_ed25519", vv); err != nil {
				return fmt.Errorf("Error reading certname_ed25519: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certname_ed25519: %v", err)
		}
	}

	if err = d.Set("certname_ed448", flattenVpnCertificateSettingCertnameEd448(o["certname-ed448"], d, "certname_ed448")); err != nil {
		if vv, ok := fortiAPIPatch(o["certname-ed448"], "VpnCertificateSetting-CertnameEd448"); ok {
			if err = d.Set("certname_ed448", vv); err != nil {
				return fmt.Errorf("Error reading certname_ed448: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certname_ed448: %v", err)
		}
	}

	if err = d.Set("certname_rsa1024", flattenVpnCertificateSettingCertnameRsa1024(o["certname-rsa1024"], d, "certname_rsa1024")); err != nil {
		if vv, ok := fortiAPIPatch(o["certname-rsa1024"], "VpnCertificateSetting-CertnameRsa1024"); ok {
			if err = d.Set("certname_rsa1024", vv); err != nil {
				return fmt.Errorf("Error reading certname_rsa1024: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certname_rsa1024: %v", err)
		}
	}

	if err = d.Set("certname_rsa2048", flattenVpnCertificateSettingCertnameRsa2048(o["certname-rsa2048"], d, "certname_rsa2048")); err != nil {
		if vv, ok := fortiAPIPatch(o["certname-rsa2048"], "VpnCertificateSetting-CertnameRsa2048"); ok {
			if err = d.Set("certname_rsa2048", vv); err != nil {
				return fmt.Errorf("Error reading certname_rsa2048: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certname_rsa2048: %v", err)
		}
	}

	if err = d.Set("certname_rsa4096", flattenVpnCertificateSettingCertnameRsa4096(o["certname-rsa4096"], d, "certname_rsa4096")); err != nil {
		if vv, ok := fortiAPIPatch(o["certname-rsa4096"], "VpnCertificateSetting-CertnameRsa4096"); ok {
			if err = d.Set("certname_rsa4096", vv); err != nil {
				return fmt.Errorf("Error reading certname_rsa4096: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certname_rsa4096: %v", err)
		}
	}

	if err = d.Set("check_ca_cert", flattenVpnCertificateSettingCheckCaCert(o["check-ca-cert"], d, "check_ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-ca-cert"], "VpnCertificateSetting-CheckCaCert"); ok {
			if err = d.Set("check_ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading check_ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_ca_cert: %v", err)
		}
	}

	if err = d.Set("check_ca_chain", flattenVpnCertificateSettingCheckCaChain(o["check-ca-chain"], d, "check_ca_chain")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-ca-chain"], "VpnCertificateSetting-CheckCaChain"); ok {
			if err = d.Set("check_ca_chain", vv); err != nil {
				return fmt.Errorf("Error reading check_ca_chain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_ca_chain: %v", err)
		}
	}

	if err = d.Set("cmp_key_usage_checking", flattenVpnCertificateSettingCmpKeyUsageChecking(o["cmp-key-usage-checking"], d, "cmp_key_usage_checking")); err != nil {
		if vv, ok := fortiAPIPatch(o["cmp-key-usage-checking"], "VpnCertificateSetting-CmpKeyUsageChecking"); ok {
			if err = d.Set("cmp_key_usage_checking", vv); err != nil {
				return fmt.Errorf("Error reading cmp_key_usage_checking: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cmp_key_usage_checking: %v", err)
		}
	}

	if err = d.Set("cmp_save_extra_certs", flattenVpnCertificateSettingCmpSaveExtraCerts(o["cmp-save-extra-certs"], d, "cmp_save_extra_certs")); err != nil {
		if vv, ok := fortiAPIPatch(o["cmp-save-extra-certs"], "VpnCertificateSetting-CmpSaveExtraCerts"); ok {
			if err = d.Set("cmp_save_extra_certs", vv); err != nil {
				return fmt.Errorf("Error reading cmp_save_extra_certs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cmp_save_extra_certs: %v", err)
		}
	}

	if err = d.Set("cn_allow_multi", flattenVpnCertificateSettingCnAllowMulti(o["cn-allow-multi"], d, "cn_allow_multi")); err != nil {
		if vv, ok := fortiAPIPatch(o["cn-allow-multi"], "VpnCertificateSetting-CnAllowMulti"); ok {
			if err = d.Set("cn_allow_multi", vv); err != nil {
				return fmt.Errorf("Error reading cn_allow_multi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cn_allow_multi: %v", err)
		}
	}

	if err = d.Set("cn_match", flattenVpnCertificateSettingCnMatch(o["cn-match"], d, "cn_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["cn-match"], "VpnCertificateSetting-CnMatch"); ok {
			if err = d.Set("cn_match", vv); err != nil {
				return fmt.Errorf("Error reading cn_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cn_match: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("crl_verification", flattenVpnCertificateSettingCrlVerification(o["crl-verification"], d, "crl_verification")); err != nil {
			if vv, ok := fortiAPIPatch(o["crl-verification"], "VpnCertificateSetting-CrlVerification"); ok {
				if err = d.Set("crl_verification", vv); err != nil {
					return fmt.Errorf("Error reading crl_verification: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading crl_verification: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("crl_verification"); ok {
			if err = d.Set("crl_verification", flattenVpnCertificateSettingCrlVerification(o["crl-verification"], d, "crl_verification")); err != nil {
				if vv, ok := fortiAPIPatch(o["crl-verification"], "VpnCertificateSetting-CrlVerification"); ok {
					if err = d.Set("crl_verification", vv); err != nil {
						return fmt.Errorf("Error reading crl_verification: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading crl_verification: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface", flattenVpnCertificateSettingInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "VpnCertificateSetting-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenVpnCertificateSettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "VpnCertificateSetting-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ocsp_default_server", flattenVpnCertificateSettingOcspDefaultServer(o["ocsp-default-server"], d, "ocsp_default_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ocsp-default-server"], "VpnCertificateSetting-OcspDefaultServer"); ok {
			if err = d.Set("ocsp_default_server", vv); err != nil {
				return fmt.Errorf("Error reading ocsp_default_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ocsp_default_server: %v", err)
		}
	}

	if err = d.Set("ocsp_option", flattenVpnCertificateSettingOcspOption(o["ocsp-option"], d, "ocsp_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["ocsp-option"], "VpnCertificateSetting-OcspOption"); ok {
			if err = d.Set("ocsp_option", vv); err != nil {
				return fmt.Errorf("Error reading ocsp_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ocsp_option: %v", err)
		}
	}

	if err = d.Set("ocsp_status", flattenVpnCertificateSettingOcspStatus(o["ocsp-status"], d, "ocsp_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ocsp-status"], "VpnCertificateSetting-OcspStatus"); ok {
			if err = d.Set("ocsp_status", vv); err != nil {
				return fmt.Errorf("Error reading ocsp_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ocsp_status: %v", err)
		}
	}

	if err = d.Set("proxy", flattenVpnCertificateSettingProxy(o["proxy"], d, "proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy"], "VpnCertificateSetting-Proxy"); ok {
			if err = d.Set("proxy", vv); err != nil {
				return fmt.Errorf("Error reading proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("proxy_port", flattenVpnCertificateSettingProxyPort(o["proxy-port"], d, "proxy_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-port"], "VpnCertificateSetting-ProxyPort"); ok {
			if err = d.Set("proxy_port", vv); err != nil {
				return fmt.Errorf("Error reading proxy_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_port: %v", err)
		}
	}

	if err = d.Set("proxy_username", flattenVpnCertificateSettingProxyUsername(o["proxy-username"], d, "proxy_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-username"], "VpnCertificateSetting-ProxyUsername"); ok {
			if err = d.Set("proxy_username", vv); err != nil {
				return fmt.Errorf("Error reading proxy_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_username: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenVpnCertificateSettingSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "VpnCertificateSetting-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenVpnCertificateSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "VpnCertificateSetting-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("ssl_ocsp_source_ip", flattenVpnCertificateSettingSslOcspSourceIp(o["ssl-ocsp-source-ip"], d, "ssl_ocsp_source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-ocsp-source-ip"], "VpnCertificateSetting-SslOcspSourceIp"); ok {
			if err = d.Set("ssl_ocsp_source_ip", vv); err != nil {
				return fmt.Errorf("Error reading ssl_ocsp_source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_ocsp_source_ip: %v", err)
		}
	}

	if err = d.Set("strict_crl_check", flattenVpnCertificateSettingStrictCrlCheck(o["strict-crl-check"], d, "strict_crl_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["strict-crl-check"], "VpnCertificateSetting-StrictCrlCheck"); ok {
			if err = d.Set("strict_crl_check", vv); err != nil {
				return fmt.Errorf("Error reading strict_crl_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strict_crl_check: %v", err)
		}
	}

	if err = d.Set("strict_ocsp_check", flattenVpnCertificateSettingStrictOcspCheck(o["strict-ocsp-check"], d, "strict_ocsp_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["strict-ocsp-check"], "VpnCertificateSetting-StrictOcspCheck"); ok {
			if err = d.Set("strict_ocsp_check", vv); err != nil {
				return fmt.Errorf("Error reading strict_ocsp_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strict_ocsp_check: %v", err)
		}
	}

	if err = d.Set("subject_match", flattenVpnCertificateSettingSubjectMatch(o["subject-match"], d, "subject_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["subject-match"], "VpnCertificateSetting-SubjectMatch"); ok {
			if err = d.Set("subject_match", vv); err != nil {
				return fmt.Errorf("Error reading subject_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subject_match: %v", err)
		}
	}

	if err = d.Set("subject_set", flattenVpnCertificateSettingSubjectSet(o["subject-set"], d, "subject_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["subject-set"], "VpnCertificateSetting-SubjectSet"); ok {
			if err = d.Set("subject_set", vv); err != nil {
				return fmt.Errorf("Error reading subject_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subject_set: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnCertificateSettingCertExpireWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameDsa1024(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingCertnameDsa2048(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingCertnameEcdsa256(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingCertnameEcdsa384(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingCertnameEcdsa521(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingCertnameEd25519(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingCertnameEd448(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingCertnameRsa1024(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingCertnameRsa2048(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingCertnameRsa4096(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingCheckCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCheckCaChain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCmpKeyUsageChecking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCmpSaveExtraCerts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCnAllowMulti(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCnMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCrlVerification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "chain_crl_absence"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["chain-crl-absence"], _ = expandVpnCertificateSettingCrlVerificationChainCrlAbsence(d, i["chain_crl_absence"], pre_append)
	}
	pre_append = pre + ".0." + "expiry"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["expiry"], _ = expandVpnCertificateSettingCrlVerificationExpiry(d, i["expiry"], pre_append)
	}
	pre_append = pre + ".0." + "leaf_crl_absence"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["leaf-crl-absence"], _ = expandVpnCertificateSettingCrlVerificationLeafCrlAbsence(d, i["leaf_crl_absence"], pre_append)
	}

	return result, nil
}

func expandVpnCertificateSettingCrlVerificationChainCrlAbsence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCrlVerificationExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCrlVerificationLeafCrlAbsence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingOcspDefaultServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingOcspOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingOcspStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingProxyPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateSettingProxyPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingProxyUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingSslOcspSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingStrictCrlCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingStrictOcspCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingSubjectMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingSubjectSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cert_expire_warning"); ok || d.HasChange("cert_expire_warning") {
		t, err := expandVpnCertificateSettingCertExpireWarning(d, v, "cert_expire_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-expire-warning"] = t
		}
	}

	if v, ok := d.GetOk("certname_dsa1024"); ok || d.HasChange("certname_dsa1024") {
		t, err := expandVpnCertificateSettingCertnameDsa1024(d, v, "certname_dsa1024")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-dsa1024"] = t
		}
	}

	if v, ok := d.GetOk("certname_dsa2048"); ok || d.HasChange("certname_dsa2048") {
		t, err := expandVpnCertificateSettingCertnameDsa2048(d, v, "certname_dsa2048")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-dsa2048"] = t
		}
	}

	if v, ok := d.GetOk("certname_ecdsa256"); ok || d.HasChange("certname_ecdsa256") {
		t, err := expandVpnCertificateSettingCertnameEcdsa256(d, v, "certname_ecdsa256")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-ecdsa256"] = t
		}
	}

	if v, ok := d.GetOk("certname_ecdsa384"); ok || d.HasChange("certname_ecdsa384") {
		t, err := expandVpnCertificateSettingCertnameEcdsa384(d, v, "certname_ecdsa384")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-ecdsa384"] = t
		}
	}

	if v, ok := d.GetOk("certname_ecdsa521"); ok || d.HasChange("certname_ecdsa521") {
		t, err := expandVpnCertificateSettingCertnameEcdsa521(d, v, "certname_ecdsa521")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-ecdsa521"] = t
		}
	}

	if v, ok := d.GetOk("certname_ed25519"); ok || d.HasChange("certname_ed25519") {
		t, err := expandVpnCertificateSettingCertnameEd25519(d, v, "certname_ed25519")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-ed25519"] = t
		}
	}

	if v, ok := d.GetOk("certname_ed448"); ok || d.HasChange("certname_ed448") {
		t, err := expandVpnCertificateSettingCertnameEd448(d, v, "certname_ed448")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-ed448"] = t
		}
	}

	if v, ok := d.GetOk("certname_rsa1024"); ok || d.HasChange("certname_rsa1024") {
		t, err := expandVpnCertificateSettingCertnameRsa1024(d, v, "certname_rsa1024")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-rsa1024"] = t
		}
	}

	if v, ok := d.GetOk("certname_rsa2048"); ok || d.HasChange("certname_rsa2048") {
		t, err := expandVpnCertificateSettingCertnameRsa2048(d, v, "certname_rsa2048")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-rsa2048"] = t
		}
	}

	if v, ok := d.GetOk("certname_rsa4096"); ok || d.HasChange("certname_rsa4096") {
		t, err := expandVpnCertificateSettingCertnameRsa4096(d, v, "certname_rsa4096")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-rsa4096"] = t
		}
	}

	if v, ok := d.GetOk("check_ca_cert"); ok || d.HasChange("check_ca_cert") {
		t, err := expandVpnCertificateSettingCheckCaCert(d, v, "check_ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("check_ca_chain"); ok || d.HasChange("check_ca_chain") {
		t, err := expandVpnCertificateSettingCheckCaChain(d, v, "check_ca_chain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-ca-chain"] = t
		}
	}

	if v, ok := d.GetOk("cmp_key_usage_checking"); ok || d.HasChange("cmp_key_usage_checking") {
		t, err := expandVpnCertificateSettingCmpKeyUsageChecking(d, v, "cmp_key_usage_checking")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-key-usage-checking"] = t
		}
	}

	if v, ok := d.GetOk("cmp_save_extra_certs"); ok || d.HasChange("cmp_save_extra_certs") {
		t, err := expandVpnCertificateSettingCmpSaveExtraCerts(d, v, "cmp_save_extra_certs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-save-extra-certs"] = t
		}
	}

	if v, ok := d.GetOk("cn_allow_multi"); ok || d.HasChange("cn_allow_multi") {
		t, err := expandVpnCertificateSettingCnAllowMulti(d, v, "cn_allow_multi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cn-allow-multi"] = t
		}
	}

	if v, ok := d.GetOk("cn_match"); ok || d.HasChange("cn_match") {
		t, err := expandVpnCertificateSettingCnMatch(d, v, "cn_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cn-match"] = t
		}
	}

	if v, ok := d.GetOk("crl_verification"); ok || d.HasChange("crl_verification") {
		t, err := expandVpnCertificateSettingCrlVerification(d, v, "crl_verification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crl-verification"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandVpnCertificateSettingInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandVpnCertificateSettingInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ocsp_default_server"); ok || d.HasChange("ocsp_default_server") {
		t, err := expandVpnCertificateSettingOcspDefaultServer(d, v, "ocsp_default_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-default-server"] = t
		}
	}

	if v, ok := d.GetOk("ocsp_option"); ok || d.HasChange("ocsp_option") {
		t, err := expandVpnCertificateSettingOcspOption(d, v, "ocsp_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-option"] = t
		}
	}

	if v, ok := d.GetOk("ocsp_status"); ok || d.HasChange("ocsp_status") {
		t, err := expandVpnCertificateSettingOcspStatus(d, v, "ocsp_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-status"] = t
		}
	}

	if v, ok := d.GetOk("proxy"); ok || d.HasChange("proxy") {
		t, err := expandVpnCertificateSettingProxy(d, v, "proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	if v, ok := d.GetOk("proxy_password"); ok || d.HasChange("proxy_password") {
		t, err := expandVpnCertificateSettingProxyPassword(d, v, "proxy_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-password"] = t
		}
	}

	if v, ok := d.GetOk("proxy_port"); ok || d.HasChange("proxy_port") {
		t, err := expandVpnCertificateSettingProxyPort(d, v, "proxy_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-port"] = t
		}
	}

	if v, ok := d.GetOk("proxy_username"); ok || d.HasChange("proxy_username") {
		t, err := expandVpnCertificateSettingProxyUsername(d, v, "proxy_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-username"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandVpnCertificateSettingSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok || d.HasChange("ssl_min_proto_version") {
		t, err := expandVpnCertificateSettingSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ocsp_source_ip"); ok || d.HasChange("ssl_ocsp_source_ip") {
		t, err := expandVpnCertificateSettingSslOcspSourceIp(d, v, "ssl_ocsp_source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ocsp-source-ip"] = t
		}
	}

	if v, ok := d.GetOk("strict_crl_check"); ok || d.HasChange("strict_crl_check") {
		t, err := expandVpnCertificateSettingStrictCrlCheck(d, v, "strict_crl_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-crl-check"] = t
		}
	}

	if v, ok := d.GetOk("strict_ocsp_check"); ok || d.HasChange("strict_ocsp_check") {
		t, err := expandVpnCertificateSettingStrictOcspCheck(d, v, "strict_ocsp_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-ocsp-check"] = t
		}
	}

	if v, ok := d.GetOk("subject_match"); ok || d.HasChange("subject_match") {
		t, err := expandVpnCertificateSettingSubjectMatch(d, v, "subject_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subject-match"] = t
		}
	}

	if v, ok := d.GetOk("subject_set"); ok || d.HasChange("subject_set") {
		t, err := expandVpnCertificateSettingSubjectSet(d, v, "subject_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subject-set"] = t
		}
	}

	return &obj, nil
}
