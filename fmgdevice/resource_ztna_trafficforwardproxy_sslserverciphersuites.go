// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SSL/TLS cipher suites to offer to a server, ordered by priority.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaTrafficForwardProxySslServerCipherSuites() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaTrafficForwardProxySslServerCipherSuitesCreate,
		Read:   resourceZtnaTrafficForwardProxySslServerCipherSuitesRead,
		Update: resourceZtnaTrafficForwardProxySslServerCipherSuitesUpdate,
		Delete: resourceZtnaTrafficForwardProxySslServerCipherSuitesDelete,

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

func resourceZtnaTrafficForwardProxySslServerCipherSuitesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaTrafficForwardProxySslServerCipherSuites(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaTrafficForwardProxySslServerCipherSuites resource while getting object: %v", err)
	}

	v, err := c.CreateZtnaTrafficForwardProxySslServerCipherSuites(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaTrafficForwardProxySslServerCipherSuites resource: %v", err)
	}

	if v != nil && v["priority"] != nil {
		if vidn, ok := v["priority"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceZtnaTrafficForwardProxySslServerCipherSuitesRead(d, m)
		} else {
			return fmt.Errorf("Error creating ZtnaTrafficForwardProxySslServerCipherSuites resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "priority")))

	return resourceZtnaTrafficForwardProxySslServerCipherSuitesRead(d, m)
}

func resourceZtnaTrafficForwardProxySslServerCipherSuitesUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaTrafficForwardProxySslServerCipherSuites(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxySslServerCipherSuites resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaTrafficForwardProxySslServerCipherSuites(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxySslServerCipherSuites resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "priority")))

	return resourceZtnaTrafficForwardProxySslServerCipherSuitesRead(d, m)
}

func resourceZtnaTrafficForwardProxySslServerCipherSuitesDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteZtnaTrafficForwardProxySslServerCipherSuites(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaTrafficForwardProxySslServerCipherSuites resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaTrafficForwardProxySslServerCipherSuitesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaTrafficForwardProxySslServerCipherSuites(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxySslServerCipherSuites resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaTrafficForwardProxySslServerCipherSuites(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxySslServerCipherSuites resource from API: %v", err)
	}
	return nil
}

func flattenZtnaTrafficForwardProxySslServerCipherSuitesCipher2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerCipherSuitesPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerCipherSuitesVersions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectZtnaTrafficForwardProxySslServerCipherSuites(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cipher", flattenZtnaTrafficForwardProxySslServerCipherSuitesCipher2edl(o["cipher"], d, "cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["cipher"], "ZtnaTrafficForwardProxySslServerCipherSuites-Cipher"); ok {
			if err = d.Set("cipher", vv); err != nil {
				return fmt.Errorf("Error reading cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cipher: %v", err)
		}
	}

	if err = d.Set("priority", flattenZtnaTrafficForwardProxySslServerCipherSuitesPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "ZtnaTrafficForwardProxySslServerCipherSuites-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("versions", flattenZtnaTrafficForwardProxySslServerCipherSuitesVersions2edl(o["versions"], d, "versions")); err != nil {
		if vv, ok := fortiAPIPatch(o["versions"], "ZtnaTrafficForwardProxySslServerCipherSuites-Versions"); ok {
			if err = d.Set("versions", vv); err != nil {
				return fmt.Errorf("Error reading versions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading versions: %v", err)
		}
	}

	return nil
}

func flattenZtnaTrafficForwardProxySslServerCipherSuitesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaTrafficForwardProxySslServerCipherSuitesCipher2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerCipherSuitesPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerCipherSuitesVersions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectZtnaTrafficForwardProxySslServerCipherSuites(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cipher"); ok || d.HasChange("cipher") {
		t, err := expandZtnaTrafficForwardProxySslServerCipherSuitesCipher2edl(d, v, "cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cipher"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandZtnaTrafficForwardProxySslServerCipherSuitesPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("versions"); ok || d.HasChange("versions") {
		t, err := expandZtnaTrafficForwardProxySslServerCipherSuitesVersions2edl(d, v, "versions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["versions"] = t
		}
	}

	return &obj, nil
}
