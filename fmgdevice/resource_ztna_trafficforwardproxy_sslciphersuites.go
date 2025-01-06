// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SSL/TLS cipher suites acceptable from a client, ordered by priority.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaTrafficForwardProxySslCipherSuites() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaTrafficForwardProxySslCipherSuitesCreate,
		Read:   resourceZtnaTrafficForwardProxySslCipherSuitesRead,
		Update: resourceZtnaTrafficForwardProxySslCipherSuitesUpdate,
		Delete: resourceZtnaTrafficForwardProxySslCipherSuitesDelete,

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
			"traffic_forward_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"versions": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceZtnaTrafficForwardProxySslCipherSuitesCreate(d *schema.ResourceData, m interface{}) error {
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
	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	obj, err := getObjectZtnaTrafficForwardProxySslCipherSuites(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaTrafficForwardProxySslCipherSuites resource while getting object: %v", err)
	}

	v, err := c.CreateZtnaTrafficForwardProxySslCipherSuites(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ZtnaTrafficForwardProxySslCipherSuites resource: %v", err)
	}

	if v != nil && v["priority"] != nil {
		if vidn, ok := v["priority"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceZtnaTrafficForwardProxySslCipherSuitesRead(d, m)
		} else {
			return fmt.Errorf("Error creating ZtnaTrafficForwardProxySslCipherSuites resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "priority")))

	return resourceZtnaTrafficForwardProxySslCipherSuitesRead(d, m)
}

func resourceZtnaTrafficForwardProxySslCipherSuitesUpdate(d *schema.ResourceData, m interface{}) error {
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
	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	obj, err := getObjectZtnaTrafficForwardProxySslCipherSuites(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxySslCipherSuites resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaTrafficForwardProxySslCipherSuites(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxySslCipherSuites resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "priority")))

	return resourceZtnaTrafficForwardProxySslCipherSuitesRead(d, m)
}

func resourceZtnaTrafficForwardProxySslCipherSuitesDelete(d *schema.ResourceData, m interface{}) error {
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
	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	err = c.DeleteZtnaTrafficForwardProxySslCipherSuites(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaTrafficForwardProxySslCipherSuites resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaTrafficForwardProxySslCipherSuitesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
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
	if traffic_forward_proxy == "" {
		traffic_forward_proxy = importOptionChecking(m.(*FortiClient).Cfg, "traffic_forward_proxy")
		if traffic_forward_proxy == "" {
			return fmt.Errorf("Parameter traffic_forward_proxy is missing")
		}
		if err = d.Set("traffic_forward_proxy", traffic_forward_proxy); err != nil {
			return fmt.Errorf("Error set params traffic_forward_proxy: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	o, err := c.ReadZtnaTrafficForwardProxySslCipherSuites(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxySslCipherSuites resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaTrafficForwardProxySslCipherSuites(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxySslCipherSuites resource from API: %v", err)
	}
	return nil
}

func flattenZtnaTrafficForwardProxySslCipherSuitesCipher2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslCipherSuitesPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslCipherSuitesVersions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectZtnaTrafficForwardProxySslCipherSuites(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cipher", flattenZtnaTrafficForwardProxySslCipherSuitesCipher2edl(o["cipher"], d, "cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["cipher"], "ZtnaTrafficForwardProxySslCipherSuites-Cipher"); ok {
			if err = d.Set("cipher", vv); err != nil {
				return fmt.Errorf("Error reading cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cipher: %v", err)
		}
	}

	if err = d.Set("priority", flattenZtnaTrafficForwardProxySslCipherSuitesPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "ZtnaTrafficForwardProxySslCipherSuites-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("versions", flattenZtnaTrafficForwardProxySslCipherSuitesVersions2edl(o["versions"], d, "versions")); err != nil {
		if vv, ok := fortiAPIPatch(o["versions"], "ZtnaTrafficForwardProxySslCipherSuites-Versions"); ok {
			if err = d.Set("versions", vv); err != nil {
				return fmt.Errorf("Error reading versions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading versions: %v", err)
		}
	}

	return nil
}

func flattenZtnaTrafficForwardProxySslCipherSuitesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaTrafficForwardProxySslCipherSuitesCipher2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslCipherSuitesPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslCipherSuitesVersions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectZtnaTrafficForwardProxySslCipherSuites(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cipher"); ok || d.HasChange("cipher") {
		t, err := expandZtnaTrafficForwardProxySslCipherSuitesCipher2edl(d, v, "cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cipher"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandZtnaTrafficForwardProxySslCipherSuitesPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("versions"); ok || d.HasChange("versions") {
		t, err := expandZtnaTrafficForwardProxySslCipherSuitesVersions2edl(d, v, "versions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["versions"] = t
		}
	}

	return &obj, nil
}
