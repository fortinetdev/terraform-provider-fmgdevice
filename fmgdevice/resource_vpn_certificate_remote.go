// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Remote certificate as a PEM file.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnCertificateRemote() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateRemoteCreate,
		Read:   resourceVpnCertificateRemoteRead,
		Update: resourceVpnCertificateRemoteUpdate,
		Delete: resourceVpnCertificateRemoteDelete,

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
			"remote": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnCertificateRemoteCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateRemote(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateRemote resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVpnCertificateRemote(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVpnCertificateRemote(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VpnCertificateRemote resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateVpnCertificateRemote(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VpnCertificateRemote resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnCertificateRemoteRead(d, m)
}

func resourceVpnCertificateRemoteUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateRemote(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateRemote resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVpnCertificateRemote(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateRemote resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnCertificateRemoteRead(d, m)
}

func resourceVpnCertificateRemoteDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnCertificateRemote(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateRemote resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateRemoteRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnCertificateRemote(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnCertificateRemote resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateRemote(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateRemote resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateRemoteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateRemoteRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateRemoteRemote(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateRemoteSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnCertificateRemote(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenVpnCertificateRemoteName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnCertificateRemote-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("range", flattenVpnCertificateRemoteRange(o["range"], d, "range")); err != nil {
		if vv, ok := fortiAPIPatch(o["range"], "VpnCertificateRemote-Range"); ok {
			if err = d.Set("range", vv); err != nil {
				return fmt.Errorf("Error reading range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("remote", flattenVpnCertificateRemoteRemote(o["remote"], d, "remote")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote"], "VpnCertificateRemote-Remote"); ok {
			if err = d.Set("remote", vv); err != nil {
				return fmt.Errorf("Error reading remote: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote: %v", err)
		}
	}

	if err = d.Set("source", flattenVpnCertificateRemoteSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "VpnCertificateRemote-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateRemoteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnCertificateRemoteName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateRemoteRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateRemoteRemote(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateRemoteSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateRemote(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnCertificateRemoteName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok || d.HasChange("range") {
		t, err := expandVpnCertificateRemoteRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("remote"); ok || d.HasChange("remote") {
		t, err := expandVpnCertificateRemoteRemote(d, v, "remote")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandVpnCertificateRemoteSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	return &obj, nil
}
