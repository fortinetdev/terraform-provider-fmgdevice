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

func resourceFirewallVipSslCipherSuites() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVipSslCipherSuitesCreate,
		Read:   resourceFirewallVipSslCipherSuitesRead,
		Update: resourceFirewallVipSslCipherSuitesUpdate,
		Delete: resourceFirewallVipSslCipherSuitesDelete,

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
			"vip": &schema.Schema{
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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceFirewallVipSslCipherSuitesCreate(d *schema.ResourceData, m interface{}) error {
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
	vip := d.Get("vip").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip"] = vip

	obj, err := getObjectFirewallVipSslCipherSuites(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVipSslCipherSuites resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("priority")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallVipSslCipherSuites(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallVipSslCipherSuites(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallVipSslCipherSuites resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallVipSslCipherSuites(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallVipSslCipherSuites resource: %v", err)
		}

		if v != nil && v["priority"] != nil {
			if vidn, ok := v["priority"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallVipSslCipherSuitesRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallVipSslCipherSuites resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "priority")))

	return resourceFirewallVipSslCipherSuitesRead(d, m)
}

func resourceFirewallVipSslCipherSuitesUpdate(d *schema.ResourceData, m interface{}) error {
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
	vip := d.Get("vip").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip"] = vip

	obj, err := getObjectFirewallVipSslCipherSuites(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVipSslCipherSuites resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallVipSslCipherSuites(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVipSslCipherSuites resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "priority")))

	return resourceFirewallVipSslCipherSuitesRead(d, m)
}

func resourceFirewallVipSslCipherSuitesDelete(d *schema.ResourceData, m interface{}) error {
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
	vip := d.Get("vip").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip"] = vip

	wsParams["adom"] = adomv

	err = c.DeleteFirewallVipSslCipherSuites(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVipSslCipherSuites resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVipSslCipherSuitesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	vip := d.Get("vip").(string)
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
	if vip == "" {
		vip = importOptionChecking(m.(*FortiClient).Cfg, "vip")
		if vip == "" {
			return fmt.Errorf("Parameter vip is missing")
		}
		if err = d.Set("vip", vip); err != nil {
			return fmt.Errorf("Error set params vip: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip"] = vip

	o, err := c.ReadFirewallVipSslCipherSuites(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallVipSslCipherSuites resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVipSslCipherSuites(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVipSslCipherSuites resource from API: %v", err)
	}
	return nil
}

func flattenFirewallVipSslCipherSuitesCipher2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslCipherSuitesPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipSslCipherSuitesVersions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipSslCipherSuitesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallVipSslCipherSuites(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cipher", flattenFirewallVipSslCipherSuitesCipher2edl(o["cipher"], d, "cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["cipher"], "FirewallVipSslCipherSuites-Cipher"); ok {
			if err = d.Set("cipher", vv); err != nil {
				return fmt.Errorf("Error reading cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cipher: %v", err)
		}
	}

	if err = d.Set("priority", flattenFirewallVipSslCipherSuitesPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "FirewallVipSslCipherSuites-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("versions", flattenFirewallVipSslCipherSuitesVersions2edl(o["versions"], d, "versions")); err != nil {
		if vv, ok := fortiAPIPatch(o["versions"], "FirewallVipSslCipherSuites-Versions"); ok {
			if err = d.Set("versions", vv); err != nil {
				return fmt.Errorf("Error reading versions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading versions: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallVipSslCipherSuitesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallVipSslCipherSuites-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenFirewallVipSslCipherSuitesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallVipSslCipherSuitesCipher2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslCipherSuitesPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslCipherSuitesVersions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipSslCipherSuitesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallVipSslCipherSuites(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cipher"); ok || d.HasChange("cipher") {
		t, err := expandFirewallVipSslCipherSuitesCipher2edl(d, v, "cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cipher"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandFirewallVipSslCipherSuitesPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("versions"); ok || d.HasChange("versions") {
		t, err := expandFirewallVipSslCipherSuitesVersions2edl(d, v, "versions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["versions"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallVipSslCipherSuitesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
