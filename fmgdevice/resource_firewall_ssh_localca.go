// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> SSH proxy local CA.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSshLocalCa() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSshLocalCaCreate,
		Read:   resourceFirewallSshLocalCaRead,
		Update: resourceFirewallSshLocalCaUpdate,
		Delete: resourceFirewallSshLocalCaDelete,

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
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_key": &schema.Schema{
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

func resourceFirewallSshLocalCaCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallSshLocalCa(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSshLocalCa resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallSshLocalCa(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallSshLocalCa(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallSshLocalCa resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallSshLocalCa(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallSshLocalCa resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallSshLocalCaRead(d, m)
}

func resourceFirewallSshLocalCaUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallSshLocalCa(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSshLocalCa resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallSshLocalCa(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSshLocalCa resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallSshLocalCaRead(d, m)
}

func resourceFirewallSshLocalCaDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallSshLocalCa(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSshLocalCa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSshLocalCaRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallSshLocalCa(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallSshLocalCa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSshLocalCa(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSshLocalCa resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSshLocalCaName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshLocalCaPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshLocalCaPublicKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshLocalCaSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSshLocalCa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallSshLocalCaName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallSshLocalCa-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("private_key", flattenFirewallSshLocalCaPrivateKey(o["private-key"], d, "private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-key"], "FirewallSshLocalCa-PrivateKey"); ok {
			if err = d.Set("private_key", vv); err != nil {
				return fmt.Errorf("Error reading private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_key: %v", err)
		}
	}

	if err = d.Set("public_key", flattenFirewallSshLocalCaPublicKey(o["public-key"], d, "public_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["public-key"], "FirewallSshLocalCa-PublicKey"); ok {
			if err = d.Set("public_key", vv); err != nil {
				return fmt.Errorf("Error reading public_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading public_key: %v", err)
		}
	}

	if err = d.Set("source", flattenFirewallSshLocalCaSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "FirewallSshLocalCa-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	return nil
}

func flattenFirewallSshLocalCaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSshLocalCaName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshLocalCaPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSshLocalCaPrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshLocalCaPublicKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshLocalCaSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSshLocalCa(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallSshLocalCaName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandFirewallSshLocalCaPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok || d.HasChange("private_key") {
		t, err := expandFirewallSshLocalCaPrivateKey(d, v, "private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	}

	if v, ok := d.GetOk("public_key"); ok || d.HasChange("public_key") {
		t, err := expandFirewallSshLocalCaPublicKey(d, v, "public_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-key"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandFirewallSshLocalCaSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	return &obj, nil
}
