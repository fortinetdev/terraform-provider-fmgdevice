// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> SSL/TLS cipher suites acceptable from a client, ordered by priority.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallVip6SslCipherSuites() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVip6SslCipherSuitesCreate,
		Read:   resourceFirewallVip6SslCipherSuitesRead,
		Update: resourceFirewallVip6SslCipherSuitesUpdate,
		Delete: resourceFirewallVip6SslCipherSuitesDelete,

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
			"vip6": &schema.Schema{
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

func resourceFirewallVip6SslCipherSuitesCreate(d *schema.ResourceData, m interface{}) error {
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
	vip6 := d.Get("vip6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip6"] = vip6

	obj, err := getObjectFirewallVip6SslCipherSuites(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVip6SslCipherSuites resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("priority")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallVip6SslCipherSuites(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallVip6SslCipherSuites(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallVip6SslCipherSuites resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallVip6SslCipherSuites(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallVip6SslCipherSuites resource: %v", err)
		}

		if v != nil && v["priority"] != nil {
			if vidn, ok := v["priority"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallVip6SslCipherSuitesRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallVip6SslCipherSuites resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "priority")))

	return resourceFirewallVip6SslCipherSuitesRead(d, m)
}

func resourceFirewallVip6SslCipherSuitesUpdate(d *schema.ResourceData, m interface{}) error {
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
	vip6 := d.Get("vip6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip6"] = vip6

	obj, err := getObjectFirewallVip6SslCipherSuites(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip6SslCipherSuites resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallVip6SslCipherSuites(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip6SslCipherSuites resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "priority")))

	return resourceFirewallVip6SslCipherSuitesRead(d, m)
}

func resourceFirewallVip6SslCipherSuitesDelete(d *schema.ResourceData, m interface{}) error {
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
	vip6 := d.Get("vip6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip6"] = vip6

	wsParams["adom"] = adomv

	err = c.DeleteFirewallVip6SslCipherSuites(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVip6SslCipherSuites resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVip6SslCipherSuitesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	vip6 := d.Get("vip6").(string)
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
	if vip6 == "" {
		vip6 = importOptionChecking(m.(*FortiClient).Cfg, "vip6")
		if vip6 == "" {
			return fmt.Errorf("Parameter vip6 is missing")
		}
		if err = d.Set("vip6", vip6); err != nil {
			return fmt.Errorf("Error set params vip6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip6"] = vip6

	o, err := c.ReadFirewallVip6SslCipherSuites(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallVip6SslCipherSuites resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVip6SslCipherSuites(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip6SslCipherSuites resource from API: %v", err)
	}
	return nil
}

func flattenFirewallVip6SslCipherSuitesCipher2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslCipherSuitesPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6SslCipherSuitesVersions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallVip6SslCipherSuites(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cipher", flattenFirewallVip6SslCipherSuitesCipher2edl(o["cipher"], d, "cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["cipher"], "FirewallVip6SslCipherSuites-Cipher"); ok {
			if err = d.Set("cipher", vv); err != nil {
				return fmt.Errorf("Error reading cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cipher: %v", err)
		}
	}

	if err = d.Set("priority", flattenFirewallVip6SslCipherSuitesPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "FirewallVip6SslCipherSuites-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("versions", flattenFirewallVip6SslCipherSuitesVersions2edl(o["versions"], d, "versions")); err != nil {
		if vv, ok := fortiAPIPatch(o["versions"], "FirewallVip6SslCipherSuites-Versions"); ok {
			if err = d.Set("versions", vv); err != nil {
				return fmt.Errorf("Error reading versions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading versions: %v", err)
		}
	}

	return nil
}

func flattenFirewallVip6SslCipherSuitesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallVip6SslCipherSuitesCipher2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslCipherSuitesPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslCipherSuitesVersions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallVip6SslCipherSuites(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cipher"); ok || d.HasChange("cipher") {
		t, err := expandFirewallVip6SslCipherSuitesCipher2edl(d, v, "cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cipher"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandFirewallVip6SslCipherSuitesPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("versions"); ok || d.HasChange("versions") {
		t, err := expandFirewallVip6SslCipherSuitesVersions2edl(d, v, "versions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["versions"] = t
		}
	}

	return &obj, nil
}
