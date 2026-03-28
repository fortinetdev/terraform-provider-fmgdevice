// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> SSL/TLS cipher suites to offer to a server, ordered by priority.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAccessProxyApiGateway6SslCipherSuites() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxyApiGateway6SslCipherSuitesCreate,
		Read:   resourceFirewallAccessProxyApiGateway6SslCipherSuitesRead,
		Update: resourceFirewallAccessProxyApiGateway6SslCipherSuitesUpdate,
		Delete: resourceFirewallAccessProxyApiGateway6SslCipherSuitesDelete,

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
			"access_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"api_gateway6": &schema.Schema{
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

func resourceFirewallAccessProxyApiGateway6SslCipherSuitesCreate(d *schema.ResourceData, m interface{}) error {
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
	access_proxy := d.Get("access_proxy").(string)
	api_gateway6 := d.Get("api_gateway6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy
	paradict["api_gateway6"] = api_gateway6

	obj, err := getObjectFirewallAccessProxyApiGateway6SslCipherSuites(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxyApiGateway6SslCipherSuites resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("priority")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAccessProxyApiGateway6SslCipherSuites(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAccessProxyApiGateway6SslCipherSuites(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAccessProxyApiGateway6SslCipherSuites resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallAccessProxyApiGateway6SslCipherSuites(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAccessProxyApiGateway6SslCipherSuites resource: %v", err)
		}

		if v != nil && v["priority"] != nil {
			if vidn, ok := v["priority"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallAccessProxyApiGateway6SslCipherSuitesRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallAccessProxyApiGateway6SslCipherSuites resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "priority")))

	return resourceFirewallAccessProxyApiGateway6SslCipherSuitesRead(d, m)
}

func resourceFirewallAccessProxyApiGateway6SslCipherSuitesUpdate(d *schema.ResourceData, m interface{}) error {
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
	access_proxy := d.Get("access_proxy").(string)
	api_gateway6 := d.Get("api_gateway6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy
	paradict["api_gateway6"] = api_gateway6

	obj, err := getObjectFirewallAccessProxyApiGateway6SslCipherSuites(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxyApiGateway6SslCipherSuites resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAccessProxyApiGateway6SslCipherSuites(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxyApiGateway6SslCipherSuites resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "priority")))

	return resourceFirewallAccessProxyApiGateway6SslCipherSuitesRead(d, m)
}

func resourceFirewallAccessProxyApiGateway6SslCipherSuitesDelete(d *schema.ResourceData, m interface{}) error {
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
	access_proxy := d.Get("access_proxy").(string)
	api_gateway6 := d.Get("api_gateway6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy
	paradict["api_gateway6"] = api_gateway6

	wsParams["adom"] = adomv

	err = c.DeleteFirewallAccessProxyApiGateway6SslCipherSuites(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxyApiGateway6SslCipherSuites resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxyApiGateway6SslCipherSuitesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	access_proxy := d.Get("access_proxy").(string)
	api_gateway6 := d.Get("api_gateway6").(string)
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
	if access_proxy == "" {
		access_proxy = importOptionChecking(m.(*FortiClient).Cfg, "access_proxy")
		if access_proxy == "" {
			return fmt.Errorf("Parameter access_proxy is missing")
		}
		if err = d.Set("access_proxy", access_proxy); err != nil {
			return fmt.Errorf("Error set params access_proxy: %v", err)
		}
	}
	if api_gateway6 == "" {
		api_gateway6 = importOptionChecking(m.(*FortiClient).Cfg, "api_gateway6")
		if api_gateway6 == "" {
			return fmt.Errorf("Parameter api_gateway6 is missing")
		}
		if err = d.Set("api_gateway6", api_gateway6); err != nil {
			return fmt.Errorf("Error set params api_gateway6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy
	paradict["api_gateway6"] = api_gateway6

	o, err := c.ReadFirewallAccessProxyApiGateway6SslCipherSuites(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAccessProxyApiGateway6SslCipherSuites resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxyApiGateway6SslCipherSuites(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxyApiGateway6SslCipherSuites resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxyApiGateway6SslCipherSuitesCipher3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslCipherSuitesPriority3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslCipherSuitesVersions3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallAccessProxyApiGateway6SslCipherSuites(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cipher", flattenFirewallAccessProxyApiGateway6SslCipherSuitesCipher3rdl(o["cipher"], d, "cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["cipher"], "FirewallAccessProxyApiGateway6SslCipherSuites-Cipher"); ok {
			if err = d.Set("cipher", vv); err != nil {
				return fmt.Errorf("Error reading cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cipher: %v", err)
		}
	}

	if err = d.Set("priority", flattenFirewallAccessProxyApiGateway6SslCipherSuitesPriority3rdl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "FirewallAccessProxyApiGateway6SslCipherSuites-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("versions", flattenFirewallAccessProxyApiGateway6SslCipherSuitesVersions3rdl(o["versions"], d, "versions")); err != nil {
		if vv, ok := fortiAPIPatch(o["versions"], "FirewallAccessProxyApiGateway6SslCipherSuites-Versions"); ok {
			if err = d.Set("versions", vv); err != nil {
				return fmt.Errorf("Error reading versions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading versions: %v", err)
		}
	}

	return nil
}

func flattenFirewallAccessProxyApiGateway6SslCipherSuitesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAccessProxyApiGateway6SslCipherSuitesCipher3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslCipherSuitesPriority3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslCipherSuitesVersions3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallAccessProxyApiGateway6SslCipherSuites(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cipher"); ok || d.HasChange("cipher") {
		t, err := expandFirewallAccessProxyApiGateway6SslCipherSuitesCipher3rdl(d, v, "cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cipher"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandFirewallAccessProxyApiGateway6SslCipherSuitesPriority3rdl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("versions"); ok || d.HasChange("versions") {
		t, err := expandFirewallAccessProxyApiGateway6SslCipherSuitesVersions3rdl(d, v, "versions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["versions"] = t
		}
	}

	return &obj, nil
}
