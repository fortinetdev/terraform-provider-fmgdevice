// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure system probe response.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemProbeResponse() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemProbeResponseUpdate,
		Read:   resourceSystemProbeResponseRead,
		Update: resourceSystemProbeResponseUpdate,
		Delete: resourceSystemProbeResponseDelete,

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
			"http_probe_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
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
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ttl_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemProbeResponseUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemProbeResponse(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemProbeResponse resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemProbeResponse(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemProbeResponse resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemProbeResponse")

	return resourceSystemProbeResponseRead(d, m)
}

func resourceSystemProbeResponseDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemProbeResponse(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemProbeResponse resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemProbeResponseRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemProbeResponse(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemProbeResponse resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemProbeResponse(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemProbeResponse resource from API: %v", err)
	}
	return nil
}

func flattenSystemProbeResponseHttpProbeValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemProbeResponseMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemProbeResponsePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemProbeResponseSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemProbeResponseTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemProbeResponseTtlMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemProbeResponse(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("http_probe_value", flattenSystemProbeResponseHttpProbeValue(o["http-probe-value"], d, "http_probe_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-probe-value"], "SystemProbeResponse-HttpProbeValue"); ok {
			if err = d.Set("http_probe_value", vv); err != nil {
				return fmt.Errorf("Error reading http_probe_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_probe_value: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemProbeResponseMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemProbeResponse-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemProbeResponsePort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemProbeResponse-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenSystemProbeResponseSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-mode"], "SystemProbeResponse-SecurityMode"); ok {
			if err = d.Set("security_mode", vv); err != nil {
				return fmt.Errorf("Error reading security_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("timeout", flattenSystemProbeResponseTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "SystemProbeResponse-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("ttl_mode", flattenSystemProbeResponseTtlMode(o["ttl-mode"], d, "ttl_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ttl-mode"], "SystemProbeResponse-TtlMode"); ok {
			if err = d.Set("ttl_mode", vv); err != nil {
				return fmt.Errorf("Error reading ttl_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ttl_mode: %v", err)
		}
	}

	return nil
}

func flattenSystemProbeResponseFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemProbeResponseHttpProbeValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemProbeResponseMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemProbeResponsePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemProbeResponsePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemProbeResponseSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemProbeResponseTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemProbeResponseTtlMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemProbeResponse(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("http_probe_value"); ok || d.HasChange("http_probe_value") {
		t, err := expandSystemProbeResponseHttpProbeValue(d, v, "http_probe_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-probe-value"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemProbeResponseMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemProbeResponsePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemProbeResponsePort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok || d.HasChange("security_mode") {
		t, err := expandSystemProbeResponseSecurityMode(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandSystemProbeResponseTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("ttl_mode"); ok || d.HasChange("ttl_mode") {
		t, err := expandSystemProbeResponseTtlMode(d, v, "ttl_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl-mode"] = t
		}
	}

	return &obj, nil
}
