// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Remote certificate as a PEM file.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCertificateRemote() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateRemoteCreate,
		Read:   resourceCertificateRemoteRead,
		Update: resourceCertificateRemoteUpdate,
		Delete: resourceCertificateRemoteDelete,

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

func resourceCertificateRemoteCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectCertificateRemote(d)
	if err != nil {
		return fmt.Errorf("Error creating CertificateRemote resource while getting object: %v", err)
	}

	_, err = c.CreateCertificateRemote(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating CertificateRemote resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceCertificateRemoteRead(d, m)
}

func resourceCertificateRemoteUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectCertificateRemote(d)
	if err != nil {
		return fmt.Errorf("Error updating CertificateRemote resource while getting object: %v", err)
	}

	_, err = c.UpdateCertificateRemote(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating CertificateRemote resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCertificateRemoteRead(d, m)
}

func resourceCertificateRemoteDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteCertificateRemote(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting CertificateRemote resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCertificateRemoteRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCertificateRemote(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading CertificateRemote resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCertificateRemote(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CertificateRemote resource from API: %v", err)
	}
	return nil
}

func flattenCertificateRemoteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateRemoteRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateRemoteRemote(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateRemoteSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCertificateRemote(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenCertificateRemoteName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CertificateRemote-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("range", flattenCertificateRemoteRange(o["range"], d, "range")); err != nil {
		if vv, ok := fortiAPIPatch(o["range"], "CertificateRemote-Range"); ok {
			if err = d.Set("range", vv); err != nil {
				return fmt.Errorf("Error reading range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("remote", flattenCertificateRemoteRemote(o["remote"], d, "remote")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote"], "CertificateRemote-Remote"); ok {
			if err = d.Set("remote", vv); err != nil {
				return fmt.Errorf("Error reading remote: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote: %v", err)
		}
	}

	if err = d.Set("source", flattenCertificateRemoteSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "CertificateRemote-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	return nil
}

func flattenCertificateRemoteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCertificateRemoteName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateRemoteRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateRemoteRemote(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateRemoteSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCertificateRemote(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCertificateRemoteName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok || d.HasChange("range") {
		t, err := expandCertificateRemoteRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("remote"); ok || d.HasChange("remote") {
		t, err := expandCertificateRemoteRemote(d, v, "remote")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandCertificateRemoteSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	return &obj, nil
}
