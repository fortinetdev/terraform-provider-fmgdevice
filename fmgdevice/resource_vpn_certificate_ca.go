// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CA certificate.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnCertificateCa() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateCaCreate,
		Read:   resourceVpnCertificateCaRead,
		Update: resourceVpnCertificateCaUpdate,
		Delete: resourceVpnCertificateCaDelete,

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
			"_private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_update_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auto_update_days_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ca": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ca_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"est_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_ca": &schema.Schema{
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
			"non_fabric_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obsolete": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeString,
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
			"ssl_inspection_trusted": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnCertificateCaCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateCa(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateCa resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVpnCertificateCa(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVpnCertificateCa(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VpnCertificateCa resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateVpnCertificateCa(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VpnCertificateCa resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnCertificateCaRead(d, m)
}

func resourceVpnCertificateCaUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateCa(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateCa resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVpnCertificateCa(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateCa resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnCertificateCaRead(d, m)
}

func resourceVpnCertificateCaDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnCertificateCa(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateCa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateCaRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnCertificateCa(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnCertificateCa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateCa(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateCa resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateCaPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaAutoUpdateDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaAutoUpdateDaysWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaCaIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaEstUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaFabricCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaLastUpdated(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaNonFabricName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaObsolete(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaScepUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaSslInspectionTrusted(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnCertificateCa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("_private_key", flattenVpnCertificateCaPrivateKey(o["_private_key"], d, "_private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["_private_key"], "VpnCertificateCa-PrivateKey"); ok {
			if err = d.Set("_private_key", vv); err != nil {
				return fmt.Errorf("Error reading _private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _private_key: %v", err)
		}
	}

	if err = d.Set("auto_update_days", flattenVpnCertificateCaAutoUpdateDays(o["auto-update-days"], d, "auto_update_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-update-days"], "VpnCertificateCa-AutoUpdateDays"); ok {
			if err = d.Set("auto_update_days", vv); err != nil {
				return fmt.Errorf("Error reading auto_update_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_update_days: %v", err)
		}
	}

	if err = d.Set("auto_update_days_warning", flattenVpnCertificateCaAutoUpdateDaysWarning(o["auto-update-days-warning"], d, "auto_update_days_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-update-days-warning"], "VpnCertificateCa-AutoUpdateDaysWarning"); ok {
			if err = d.Set("auto_update_days_warning", vv); err != nil {
				return fmt.Errorf("Error reading auto_update_days_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_update_days_warning: %v", err)
		}
	}

	if err = d.Set("ca", flattenVpnCertificateCaCa(o["ca"], d, "ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca"], "VpnCertificateCa-Ca"); ok {
			if err = d.Set("ca", vv); err != nil {
				return fmt.Errorf("Error reading ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("ca_identifier", flattenVpnCertificateCaCaIdentifier(o["ca-identifier"], d, "ca_identifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca-identifier"], "VpnCertificateCa-CaIdentifier"); ok {
			if err = d.Set("ca_identifier", vv); err != nil {
				return fmt.Errorf("Error reading ca_identifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca_identifier: %v", err)
		}
	}

	if err = d.Set("est_url", flattenVpnCertificateCaEstUrl(o["est-url"], d, "est_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-url"], "VpnCertificateCa-EstUrl"); ok {
			if err = d.Set("est_url", vv); err != nil {
				return fmt.Errorf("Error reading est_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_url: %v", err)
		}
	}

	if err = d.Set("fabric_ca", flattenVpnCertificateCaFabricCa(o["fabric-ca"], d, "fabric_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-ca"], "VpnCertificateCa-FabricCa"); ok {
			if err = d.Set("fabric_ca", vv); err != nil {
				return fmt.Errorf("Error reading fabric_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_ca: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenVpnCertificateCaLastUpdated(o["last-updated"], d, "last_updated")); err != nil {
		if vv, ok := fortiAPIPatch(o["last-updated"], "VpnCertificateCa-LastUpdated"); ok {
			if err = d.Set("last_updated", vv); err != nil {
				return fmt.Errorf("Error reading last_updated: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnCertificateCaName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnCertificateCa-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("non_fabric_name", flattenVpnCertificateCaNonFabricName(o["non-fabric-name"], d, "non_fabric_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["non-fabric-name"], "VpnCertificateCa-NonFabricName"); ok {
			if err = d.Set("non_fabric_name", vv); err != nil {
				return fmt.Errorf("Error reading non_fabric_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading non_fabric_name: %v", err)
		}
	}

	if err = d.Set("obsolete", flattenVpnCertificateCaObsolete(o["obsolete"], d, "obsolete")); err != nil {
		if vv, ok := fortiAPIPatch(o["obsolete"], "VpnCertificateCa-Obsolete"); ok {
			if err = d.Set("obsolete", vv); err != nil {
				return fmt.Errorf("Error reading obsolete: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obsolete: %v", err)
		}
	}

	if err = d.Set("range", flattenVpnCertificateCaRange(o["range"], d, "range")); err != nil {
		if vv, ok := fortiAPIPatch(o["range"], "VpnCertificateCa-Range"); ok {
			if err = d.Set("range", vv); err != nil {
				return fmt.Errorf("Error reading range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenVpnCertificateCaScepUrl(o["scep-url"], d, "scep_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["scep-url"], "VpnCertificateCa-ScepUrl"); ok {
			if err = d.Set("scep_url", vv); err != nil {
				return fmt.Errorf("Error reading scep_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("source", flattenVpnCertificateCaSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "VpnCertificateCa-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenVpnCertificateCaSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "VpnCertificateCa-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ssl_inspection_trusted", flattenVpnCertificateCaSslInspectionTrusted(o["ssl-inspection-trusted"], d, "ssl_inspection_trusted")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-inspection-trusted"], "VpnCertificateCa-SslInspectionTrusted"); ok {
			if err = d.Set("ssl_inspection_trusted", vv); err != nil {
				return fmt.Errorf("Error reading ssl_inspection_trusted: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_inspection_trusted: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateCaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnCertificateCaPrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaAutoUpdateDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaAutoUpdateDaysWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaCaIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaEstUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaFabricCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaLastUpdated(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaNonFabricName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaObsolete(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaScepUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaSslInspectionTrusted(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateCa(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_private_key"); ok || d.HasChange("_private_key") {
		t, err := expandVpnCertificateCaPrivateKey(d, v, "_private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_private_key"] = t
		}
	}

	if v, ok := d.GetOk("auto_update_days"); ok || d.HasChange("auto_update_days") {
		t, err := expandVpnCertificateCaAutoUpdateDays(d, v, "auto_update_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-update-days"] = t
		}
	}

	if v, ok := d.GetOk("auto_update_days_warning"); ok || d.HasChange("auto_update_days_warning") {
		t, err := expandVpnCertificateCaAutoUpdateDaysWarning(d, v, "auto_update_days_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-update-days-warning"] = t
		}
	}

	if v, ok := d.GetOk("ca"); ok || d.HasChange("ca") {
		t, err := expandVpnCertificateCaCa(d, v, "ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("ca_identifier"); ok || d.HasChange("ca_identifier") {
		t, err := expandVpnCertificateCaCaIdentifier(d, v, "ca_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-identifier"] = t
		}
	}

	if v, ok := d.GetOk("est_url"); ok || d.HasChange("est_url") {
		t, err := expandVpnCertificateCaEstUrl(d, v, "est_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-url"] = t
		}
	}

	if v, ok := d.GetOk("fabric_ca"); ok || d.HasChange("fabric_ca") {
		t, err := expandVpnCertificateCaFabricCa(d, v, "fabric_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-ca"] = t
		}
	}

	if v, ok := d.GetOk("last_updated"); ok || d.HasChange("last_updated") {
		t, err := expandVpnCertificateCaLastUpdated(d, v, "last_updated")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-updated"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnCertificateCaName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("non_fabric_name"); ok || d.HasChange("non_fabric_name") {
		t, err := expandVpnCertificateCaNonFabricName(d, v, "non_fabric_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["non-fabric-name"] = t
		}
	}

	if v, ok := d.GetOk("obsolete"); ok || d.HasChange("obsolete") {
		t, err := expandVpnCertificateCaObsolete(d, v, "obsolete")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obsolete"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok || d.HasChange("range") {
		t, err := expandVpnCertificateCaRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("scep_url"); ok || d.HasChange("scep_url") {
		t, err := expandVpnCertificateCaScepUrl(d, v, "scep_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandVpnCertificateCaSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandVpnCertificateCaSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("ssl_inspection_trusted"); ok || d.HasChange("ssl_inspection_trusted") {
		t, err := expandVpnCertificateCaSslInspectionTrusted(d, v, "ssl_inspection_trusted")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-inspection-trusted"] = t
		}
	}

	return &obj, nil
}
