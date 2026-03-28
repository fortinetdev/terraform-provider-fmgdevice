// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> OCSP server configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnCertificateOcspServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateOcspServerCreate,
		Read:   resourceVpnCertificateOcspServerRead,
		Update: resourceVpnCertificateOcspServerUpdate,
		Delete: resourceVpnCertificateOcspServerDelete,

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
			"cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"secondary_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"secondary_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"unavail_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVpnCertificateOcspServerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateOcspServer(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateOcspServer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVpnCertificateOcspServer(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVpnCertificateOcspServer(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VpnCertificateOcspServer resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateVpnCertificateOcspServer(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VpnCertificateOcspServer resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnCertificateOcspServerRead(d, m)
}

func resourceVpnCertificateOcspServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateOcspServer(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateOcspServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVpnCertificateOcspServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateOcspServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnCertificateOcspServerRead(d, m)
}

func resourceVpnCertificateOcspServerDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnCertificateOcspServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateOcspServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateOcspServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnCertificateOcspServer(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnCertificateOcspServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateOcspServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateOcspServer resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateOcspServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateOcspServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerSecondaryCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnCertificateOcspServerSecondaryUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerUnavailAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnCertificateOcspServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cert", flattenVpnCertificateOcspServerCert(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "VpnCertificateOcspServer-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnCertificateOcspServerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnCertificateOcspServer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("secondary_cert", flattenVpnCertificateOcspServerSecondaryCert(o["secondary-cert"], d, "secondary_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-cert"], "VpnCertificateOcspServer-SecondaryCert"); ok {
			if err = d.Set("secondary_cert", vv); err != nil {
				return fmt.Errorf("Error reading secondary_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_cert: %v", err)
		}
	}

	if err = d.Set("secondary_url", flattenVpnCertificateOcspServerSecondaryUrl(o["secondary-url"], d, "secondary_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-url"], "VpnCertificateOcspServer-SecondaryUrl"); ok {
			if err = d.Set("secondary_url", vv); err != nil {
				return fmt.Errorf("Error reading secondary_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_url: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenVpnCertificateOcspServerSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "VpnCertificateOcspServer-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("unavail_action", flattenVpnCertificateOcspServerUnavailAction(o["unavail-action"], d, "unavail_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["unavail-action"], "VpnCertificateOcspServer-UnavailAction"); ok {
			if err = d.Set("unavail_action", vv); err != nil {
				return fmt.Errorf("Error reading unavail_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unavail_action: %v", err)
		}
	}

	if err = d.Set("url", flattenVpnCertificateOcspServerUrl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "VpnCertificateOcspServer-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateOcspServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnCertificateOcspServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateOcspServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerSecondaryCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnCertificateOcspServerSecondaryUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerUnavailAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateOcspServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandVpnCertificateOcspServerCert(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnCertificateOcspServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("secondary_cert"); ok || d.HasChange("secondary_cert") {
		t, err := expandVpnCertificateOcspServerSecondaryCert(d, v, "secondary_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-cert"] = t
		}
	}

	if v, ok := d.GetOk("secondary_url"); ok || d.HasChange("secondary_url") {
		t, err := expandVpnCertificateOcspServerSecondaryUrl(d, v, "secondary_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-url"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandVpnCertificateOcspServerSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("unavail_action"); ok || d.HasChange("unavail_action") {
		t, err := expandVpnCertificateOcspServerUnavailAction(d, v, "unavail_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unavail-action"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandVpnCertificateOcspServerUrl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
