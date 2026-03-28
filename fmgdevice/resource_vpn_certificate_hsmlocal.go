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

func resourceVpnCertificateHsmLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateHsmLocalCreate,
		Read:   resourceVpnCertificateHsmLocalRead,
		Update: resourceVpnCertificateHsmLocalUpdate,
		Delete: resourceVpnCertificateHsmLocalDelete,

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

func resourceVpnCertificateHsmLocalCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateHsmLocal(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateHsmLocal resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVpnCertificateHsmLocal(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVpnCertificateHsmLocal(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VpnCertificateHsmLocal resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateVpnCertificateHsmLocal(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VpnCertificateHsmLocal resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnCertificateHsmLocalRead(d, m)
}

func resourceVpnCertificateHsmLocalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateHsmLocal(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateHsmLocal resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVpnCertificateHsmLocal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateHsmLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnCertificateHsmLocalRead(d, m)
}

func resourceVpnCertificateHsmLocalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnCertificateHsmLocal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateHsmLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateHsmLocalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnCertificateHsmLocal(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnCertificateHsmLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateHsmLocal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateHsmLocal resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateHsmLocalApiVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalGchCloudServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateHsmLocalGchCryptokey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalGchCryptokeyAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalGchCryptokeyVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalGchKeyring(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalGchLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalGchProject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalGchUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalTmpCertFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateHsmLocalVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnCertificateHsmLocal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("api_version", flattenVpnCertificateHsmLocalApiVersion(o["api-version"], d, "api_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["api-version"], "VpnCertificateHsmLocal-ApiVersion"); ok {
			if err = d.Set("api_version", vv); err != nil {
				return fmt.Errorf("Error reading api_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading api_version: %v", err)
		}
	}

	if err = d.Set("certificate", flattenVpnCertificateHsmLocalCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "VpnCertificateHsmLocal-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnCertificateHsmLocalComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "VpnCertificateHsmLocal-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("gch_cloud_service_name", flattenVpnCertificateHsmLocalGchCloudServiceName(o["gch-cloud-service-name"], d, "gch_cloud_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-cloud-service-name"], "VpnCertificateHsmLocal-GchCloudServiceName"); ok {
			if err = d.Set("gch_cloud_service_name", vv); err != nil {
				return fmt.Errorf("Error reading gch_cloud_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_cloud_service_name: %v", err)
		}
	}

	if err = d.Set("gch_cryptokey", flattenVpnCertificateHsmLocalGchCryptokey(o["gch-cryptokey"], d, "gch_cryptokey")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-cryptokey"], "VpnCertificateHsmLocal-GchCryptokey"); ok {
			if err = d.Set("gch_cryptokey", vv); err != nil {
				return fmt.Errorf("Error reading gch_cryptokey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_cryptokey: %v", err)
		}
	}

	if err = d.Set("gch_cryptokey_algorithm", flattenVpnCertificateHsmLocalGchCryptokeyAlgorithm(o["gch-cryptokey-algorithm"], d, "gch_cryptokey_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-cryptokey-algorithm"], "VpnCertificateHsmLocal-GchCryptokeyAlgorithm"); ok {
			if err = d.Set("gch_cryptokey_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading gch_cryptokey_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_cryptokey_algorithm: %v", err)
		}
	}

	if err = d.Set("gch_cryptokey_version", flattenVpnCertificateHsmLocalGchCryptokeyVersion(o["gch-cryptokey-version"], d, "gch_cryptokey_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-cryptokey-version"], "VpnCertificateHsmLocal-GchCryptokeyVersion"); ok {
			if err = d.Set("gch_cryptokey_version", vv); err != nil {
				return fmt.Errorf("Error reading gch_cryptokey_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_cryptokey_version: %v", err)
		}
	}

	if err = d.Set("gch_keyring", flattenVpnCertificateHsmLocalGchKeyring(o["gch-keyring"], d, "gch_keyring")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-keyring"], "VpnCertificateHsmLocal-GchKeyring"); ok {
			if err = d.Set("gch_keyring", vv); err != nil {
				return fmt.Errorf("Error reading gch_keyring: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_keyring: %v", err)
		}
	}

	if err = d.Set("gch_location", flattenVpnCertificateHsmLocalGchLocation(o["gch-location"], d, "gch_location")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-location"], "VpnCertificateHsmLocal-GchLocation"); ok {
			if err = d.Set("gch_location", vv); err != nil {
				return fmt.Errorf("Error reading gch_location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_location: %v", err)
		}
	}

	if err = d.Set("gch_project", flattenVpnCertificateHsmLocalGchProject(o["gch-project"], d, "gch_project")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-project"], "VpnCertificateHsmLocal-GchProject"); ok {
			if err = d.Set("gch_project", vv); err != nil {
				return fmt.Errorf("Error reading gch_project: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_project: %v", err)
		}
	}

	if err = d.Set("gch_url", flattenVpnCertificateHsmLocalGchUrl(o["gch-url"], d, "gch_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-url"], "VpnCertificateHsmLocal-GchUrl"); ok {
			if err = d.Set("gch_url", vv); err != nil {
				return fmt.Errorf("Error reading gch_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_url: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnCertificateHsmLocalName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnCertificateHsmLocal-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("range", flattenVpnCertificateHsmLocalRange(o["range"], d, "range")); err != nil {
		if vv, ok := fortiAPIPatch(o["range"], "VpnCertificateHsmLocal-Range"); ok {
			if err = d.Set("range", vv); err != nil {
				return fmt.Errorf("Error reading range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenVpnCertificateHsmLocalSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "VpnCertificateHsmLocal-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("tmp_cert_file", flattenVpnCertificateHsmLocalTmpCertFile(o["tmp-cert-file"], d, "tmp_cert_file")); err != nil {
		if vv, ok := fortiAPIPatch(o["tmp-cert-file"], "VpnCertificateHsmLocal-TmpCertFile"); ok {
			if err = d.Set("tmp_cert_file", vv); err != nil {
				return fmt.Errorf("Error reading tmp_cert_file: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tmp_cert_file: %v", err)
		}
	}

	if err = d.Set("vendor", flattenVpnCertificateHsmLocalVendor(o["vendor"], d, "vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["vendor"], "VpnCertificateHsmLocal-Vendor"); ok {
			if err = d.Set("vendor", vv); err != nil {
				return fmt.Errorf("Error reading vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateHsmLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnCertificateHsmLocalApiVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalGchCloudServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateHsmLocalGchCryptokey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalGchCryptokeyAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalGchCryptokeyVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalGchKeyring(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalGchLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalGchProject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalGchUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalTmpCertFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateHsmLocalVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateHsmLocal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("api_version"); ok || d.HasChange("api_version") {
		t, err := expandVpnCertificateHsmLocalApiVersion(d, v, "api_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-version"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandVpnCertificateHsmLocalCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandVpnCertificateHsmLocalComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("gch_cloud_service_name"); ok || d.HasChange("gch_cloud_service_name") {
		t, err := expandVpnCertificateHsmLocalGchCloudServiceName(d, v, "gch_cloud_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-cloud-service-name"] = t
		}
	}

	if v, ok := d.GetOk("gch_cryptokey"); ok || d.HasChange("gch_cryptokey") {
		t, err := expandVpnCertificateHsmLocalGchCryptokey(d, v, "gch_cryptokey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-cryptokey"] = t
		}
	}

	if v, ok := d.GetOk("gch_cryptokey_algorithm"); ok || d.HasChange("gch_cryptokey_algorithm") {
		t, err := expandVpnCertificateHsmLocalGchCryptokeyAlgorithm(d, v, "gch_cryptokey_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-cryptokey-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("gch_cryptokey_version"); ok || d.HasChange("gch_cryptokey_version") {
		t, err := expandVpnCertificateHsmLocalGchCryptokeyVersion(d, v, "gch_cryptokey_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-cryptokey-version"] = t
		}
	}

	if v, ok := d.GetOk("gch_keyring"); ok || d.HasChange("gch_keyring") {
		t, err := expandVpnCertificateHsmLocalGchKeyring(d, v, "gch_keyring")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-keyring"] = t
		}
	}

	if v, ok := d.GetOk("gch_location"); ok || d.HasChange("gch_location") {
		t, err := expandVpnCertificateHsmLocalGchLocation(d, v, "gch_location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-location"] = t
		}
	}

	if v, ok := d.GetOk("gch_project"); ok || d.HasChange("gch_project") {
		t, err := expandVpnCertificateHsmLocalGchProject(d, v, "gch_project")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-project"] = t
		}
	}

	if v, ok := d.GetOk("gch_url"); ok || d.HasChange("gch_url") {
		t, err := expandVpnCertificateHsmLocalGchUrl(d, v, "gch_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-url"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnCertificateHsmLocalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok || d.HasChange("range") {
		t, err := expandVpnCertificateHsmLocalRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandVpnCertificateHsmLocalSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("tmp_cert_file"); ok || d.HasChange("tmp_cert_file") {
		t, err := expandVpnCertificateHsmLocalTmpCertFile(d, v, "tmp_cert_file")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tmp-cert-file"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok || d.HasChange("vendor") {
		t, err := expandVpnCertificateHsmLocalVendor(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	return &obj, nil
}
