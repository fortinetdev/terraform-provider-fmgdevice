// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Local certificates whose keys are stored on HSM.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCertificateHsmLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateHsmLocalCreate,
		Read:   resourceCertificateHsmLocalRead,
		Update: resourceCertificateHsmLocalUpdate,
		Delete: resourceCertificateHsmLocalDelete,

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
			"api_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gch_cloud_service_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"gch_cryptokey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gch_cryptokey_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gch_cryptokey_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gch_keyring": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gch_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gch_project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gch_url": &schema.Schema{
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
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tmp_cert_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCertificateHsmLocalCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectCertificateHsmLocal(d)
	if err != nil {
		return fmt.Errorf("Error creating CertificateHsmLocal resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCertificateHsmLocal(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCertificateHsmLocal(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CertificateHsmLocal resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateCertificateHsmLocal(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CertificateHsmLocal resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceCertificateHsmLocalRead(d, m)
}

func resourceCertificateHsmLocalUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectCertificateHsmLocal(d)
	if err != nil {
		return fmt.Errorf("Error updating CertificateHsmLocal resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCertificateHsmLocal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CertificateHsmLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCertificateHsmLocalRead(d, m)
}

func resourceCertificateHsmLocalDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteCertificateHsmLocal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CertificateHsmLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCertificateHsmLocalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadCertificateHsmLocal(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CertificateHsmLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCertificateHsmLocal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CertificateHsmLocal resource from API: %v", err)
	}
	return nil
}

func flattenCertificateHsmLocalApiVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalGchCloudServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCertificateHsmLocalGchCryptokey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalGchCryptokeyAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalGchCryptokeyVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalGchKeyring(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalGchLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalGchProject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalGchUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalTmpCertFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateHsmLocalVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCertificateHsmLocal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("api_version", flattenCertificateHsmLocalApiVersion(o["api-version"], d, "api_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["api-version"], "CertificateHsmLocal-ApiVersion"); ok {
			if err = d.Set("api_version", vv); err != nil {
				return fmt.Errorf("Error reading api_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading api_version: %v", err)
		}
	}

	if err = d.Set("certificate", flattenCertificateHsmLocalCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "CertificateHsmLocal-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("comments", flattenCertificateHsmLocalComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "CertificateHsmLocal-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("gch_cloud_service_name", flattenCertificateHsmLocalGchCloudServiceName(o["gch-cloud-service-name"], d, "gch_cloud_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-cloud-service-name"], "CertificateHsmLocal-GchCloudServiceName"); ok {
			if err = d.Set("gch_cloud_service_name", vv); err != nil {
				return fmt.Errorf("Error reading gch_cloud_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_cloud_service_name: %v", err)
		}
	}

	if err = d.Set("gch_cryptokey", flattenCertificateHsmLocalGchCryptokey(o["gch-cryptokey"], d, "gch_cryptokey")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-cryptokey"], "CertificateHsmLocal-GchCryptokey"); ok {
			if err = d.Set("gch_cryptokey", vv); err != nil {
				return fmt.Errorf("Error reading gch_cryptokey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_cryptokey: %v", err)
		}
	}

	if err = d.Set("gch_cryptokey_algorithm", flattenCertificateHsmLocalGchCryptokeyAlgorithm(o["gch-cryptokey-algorithm"], d, "gch_cryptokey_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-cryptokey-algorithm"], "CertificateHsmLocal-GchCryptokeyAlgorithm"); ok {
			if err = d.Set("gch_cryptokey_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading gch_cryptokey_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_cryptokey_algorithm: %v", err)
		}
	}

	if err = d.Set("gch_cryptokey_version", flattenCertificateHsmLocalGchCryptokeyVersion(o["gch-cryptokey-version"], d, "gch_cryptokey_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-cryptokey-version"], "CertificateHsmLocal-GchCryptokeyVersion"); ok {
			if err = d.Set("gch_cryptokey_version", vv); err != nil {
				return fmt.Errorf("Error reading gch_cryptokey_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_cryptokey_version: %v", err)
		}
	}

	if err = d.Set("gch_keyring", flattenCertificateHsmLocalGchKeyring(o["gch-keyring"], d, "gch_keyring")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-keyring"], "CertificateHsmLocal-GchKeyring"); ok {
			if err = d.Set("gch_keyring", vv); err != nil {
				return fmt.Errorf("Error reading gch_keyring: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_keyring: %v", err)
		}
	}

	if err = d.Set("gch_location", flattenCertificateHsmLocalGchLocation(o["gch-location"], d, "gch_location")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-location"], "CertificateHsmLocal-GchLocation"); ok {
			if err = d.Set("gch_location", vv); err != nil {
				return fmt.Errorf("Error reading gch_location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_location: %v", err)
		}
	}

	if err = d.Set("gch_project", flattenCertificateHsmLocalGchProject(o["gch-project"], d, "gch_project")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-project"], "CertificateHsmLocal-GchProject"); ok {
			if err = d.Set("gch_project", vv); err != nil {
				return fmt.Errorf("Error reading gch_project: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_project: %v", err)
		}
	}

	if err = d.Set("gch_url", flattenCertificateHsmLocalGchUrl(o["gch-url"], d, "gch_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-url"], "CertificateHsmLocal-GchUrl"); ok {
			if err = d.Set("gch_url", vv); err != nil {
				return fmt.Errorf("Error reading gch_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_url: %v", err)
		}
	}

	if err = d.Set("name", flattenCertificateHsmLocalName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CertificateHsmLocal-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("range", flattenCertificateHsmLocalRange(o["range"], d, "range")); err != nil {
		if vv, ok := fortiAPIPatch(o["range"], "CertificateHsmLocal-Range"); ok {
			if err = d.Set("range", vv); err != nil {
				return fmt.Errorf("Error reading range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenCertificateHsmLocalSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "CertificateHsmLocal-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("tmp_cert_file", flattenCertificateHsmLocalTmpCertFile(o["tmp-cert-file"], d, "tmp_cert_file")); err != nil {
		if vv, ok := fortiAPIPatch(o["tmp-cert-file"], "CertificateHsmLocal-TmpCertFile"); ok {
			if err = d.Set("tmp_cert_file", vv); err != nil {
				return fmt.Errorf("Error reading tmp_cert_file: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tmp_cert_file: %v", err)
		}
	}

	if err = d.Set("vendor", flattenCertificateHsmLocalVendor(o["vendor"], d, "vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["vendor"], "CertificateHsmLocal-Vendor"); ok {
			if err = d.Set("vendor", vv); err != nil {
				return fmt.Errorf("Error reading vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	return nil
}

func flattenCertificateHsmLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCertificateHsmLocalApiVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalGchCloudServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCertificateHsmLocalGchCryptokey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalGchCryptokeyAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalGchCryptokeyVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalGchKeyring(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalGchLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalGchProject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalGchUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalTmpCertFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateHsmLocalVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCertificateHsmLocal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("api_version"); ok || d.HasChange("api_version") {
		t, err := expandCertificateHsmLocalApiVersion(d, v, "api_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-version"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandCertificateHsmLocalCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandCertificateHsmLocalComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("gch_cloud_service_name"); ok || d.HasChange("gch_cloud_service_name") {
		t, err := expandCertificateHsmLocalGchCloudServiceName(d, v, "gch_cloud_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-cloud-service-name"] = t
		}
	}

	if v, ok := d.GetOk("gch_cryptokey"); ok || d.HasChange("gch_cryptokey") {
		t, err := expandCertificateHsmLocalGchCryptokey(d, v, "gch_cryptokey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-cryptokey"] = t
		}
	}

	if v, ok := d.GetOk("gch_cryptokey_algorithm"); ok || d.HasChange("gch_cryptokey_algorithm") {
		t, err := expandCertificateHsmLocalGchCryptokeyAlgorithm(d, v, "gch_cryptokey_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-cryptokey-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("gch_cryptokey_version"); ok || d.HasChange("gch_cryptokey_version") {
		t, err := expandCertificateHsmLocalGchCryptokeyVersion(d, v, "gch_cryptokey_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-cryptokey-version"] = t
		}
	}

	if v, ok := d.GetOk("gch_keyring"); ok || d.HasChange("gch_keyring") {
		t, err := expandCertificateHsmLocalGchKeyring(d, v, "gch_keyring")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-keyring"] = t
		}
	}

	if v, ok := d.GetOk("gch_location"); ok || d.HasChange("gch_location") {
		t, err := expandCertificateHsmLocalGchLocation(d, v, "gch_location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-location"] = t
		}
	}

	if v, ok := d.GetOk("gch_project"); ok || d.HasChange("gch_project") {
		t, err := expandCertificateHsmLocalGchProject(d, v, "gch_project")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-project"] = t
		}
	}

	if v, ok := d.GetOk("gch_url"); ok || d.HasChange("gch_url") {
		t, err := expandCertificateHsmLocalGchUrl(d, v, "gch_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-url"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCertificateHsmLocalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok || d.HasChange("range") {
		t, err := expandCertificateHsmLocalRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandCertificateHsmLocalSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("tmp_cert_file"); ok || d.HasChange("tmp_cert_file") {
		t, err := expandCertificateHsmLocalTmpCertFile(d, v, "tmp_cert_file")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tmp-cert-file"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok || d.HasChange("vendor") {
		t, err := expandCertificateHsmLocalVendor(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	return &obj, nil
}
