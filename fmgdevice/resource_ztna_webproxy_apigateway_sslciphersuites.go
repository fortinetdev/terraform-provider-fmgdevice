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

func resourceZtnaWebProxyApiGatewaySslCipherSuites() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebProxyApiGatewaySslCipherSuitesCreate,
		Read:   resourceZtnaWebProxyApiGatewaySslCipherSuitesRead,
		Update: resourceZtnaWebProxyApiGatewaySslCipherSuitesUpdate,
		Delete: resourceZtnaWebProxyApiGatewaySslCipherSuitesDelete,

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
			"web_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"api_gateway": &schema.Schema{
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

func resourceZtnaWebProxyApiGatewaySslCipherSuitesCreate(d *schema.ResourceData, m interface{}) error {
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
	web_proxy := d.Get("web_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway"] = api_gateway

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaWebProxyApiGatewaySslCipherSuites(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebProxyApiGatewaySslCipherSuites resource while getting object: %v", err)
	}

	_, err = c.CreateZtnaWebProxyApiGatewaySslCipherSuites(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebProxyApiGatewaySslCipherSuites resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "priority")))

	return resourceZtnaWebProxyApiGatewaySslCipherSuitesRead(d, m)
}

func resourceZtnaWebProxyApiGatewaySslCipherSuitesUpdate(d *schema.ResourceData, m interface{}) error {
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
	web_proxy := d.Get("web_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway"] = api_gateway

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaWebProxyApiGatewaySslCipherSuites(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxyApiGatewaySslCipherSuites resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaWebProxyApiGatewaySslCipherSuites(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxyApiGatewaySslCipherSuites resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "priority")))

	return resourceZtnaWebProxyApiGatewaySslCipherSuitesRead(d, m)
}

func resourceZtnaWebProxyApiGatewaySslCipherSuitesDelete(d *schema.ResourceData, m interface{}) error {
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
	web_proxy := d.Get("web_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway"] = api_gateway

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteZtnaWebProxyApiGatewaySslCipherSuites(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebProxyApiGatewaySslCipherSuites resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebProxyApiGatewaySslCipherSuitesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	web_proxy := d.Get("web_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
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
	if web_proxy == "" {
		web_proxy = importOptionChecking(m.(*FortiClient).Cfg, "web_proxy")
		if web_proxy == "" {
			return fmt.Errorf("Parameter web_proxy is missing")
		}
		if err = d.Set("web_proxy", web_proxy); err != nil {
			return fmt.Errorf("Error set params web_proxy: %v", err)
		}
	}
	if api_gateway == "" {
		api_gateway = importOptionChecking(m.(*FortiClient).Cfg, "api_gateway")
		if api_gateway == "" {
			return fmt.Errorf("Parameter api_gateway is missing")
		}
		if err = d.Set("api_gateway", api_gateway); err != nil {
			return fmt.Errorf("Error set params api_gateway: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway"] = api_gateway

	o, err := c.ReadZtnaWebProxyApiGatewaySslCipherSuites(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxyApiGatewaySslCipherSuites resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebProxyApiGatewaySslCipherSuites(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxyApiGatewaySslCipherSuites resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebProxyApiGatewaySslCipherSuitesCipher3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslCipherSuitesPriority3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslCipherSuitesVersions3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectZtnaWebProxyApiGatewaySslCipherSuites(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cipher", flattenZtnaWebProxyApiGatewaySslCipherSuitesCipher3rdl(o["cipher"], d, "cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["cipher"], "ZtnaWebProxyApiGatewaySslCipherSuites-Cipher"); ok {
			if err = d.Set("cipher", vv); err != nil {
				return fmt.Errorf("Error reading cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cipher: %v", err)
		}
	}

	if err = d.Set("priority", flattenZtnaWebProxyApiGatewaySslCipherSuitesPriority3rdl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "ZtnaWebProxyApiGatewaySslCipherSuites-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("versions", flattenZtnaWebProxyApiGatewaySslCipherSuitesVersions3rdl(o["versions"], d, "versions")); err != nil {
		if vv, ok := fortiAPIPatch(o["versions"], "ZtnaWebProxyApiGatewaySslCipherSuites-Versions"); ok {
			if err = d.Set("versions", vv); err != nil {
				return fmt.Errorf("Error reading versions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading versions: %v", err)
		}
	}

	return nil
}

func flattenZtnaWebProxyApiGatewaySslCipherSuitesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaWebProxyApiGatewaySslCipherSuitesCipher3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslCipherSuitesPriority3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslCipherSuitesVersions3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectZtnaWebProxyApiGatewaySslCipherSuites(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cipher"); ok || d.HasChange("cipher") {
		t, err := expandZtnaWebProxyApiGatewaySslCipherSuitesCipher3rdl(d, v, "cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cipher"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandZtnaWebProxyApiGatewaySslCipherSuitesPriority3rdl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("versions"); ok || d.HasChange("versions") {
		t, err := expandZtnaWebProxyApiGatewaySslCipherSuitesVersions3rdl(d, v, "versions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["versions"] = t
		}
	}

	return &obj, nil
}
