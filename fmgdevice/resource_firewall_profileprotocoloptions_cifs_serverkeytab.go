// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Server keytab.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProfileProtocolOptionsCifsServerKeytab() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProfileProtocolOptionsCifsServerKeytabCreate,
		Read:   resourceFirewallProfileProtocolOptionsCifsServerKeytabRead,
		Update: resourceFirewallProfileProtocolOptionsCifsServerKeytabUpdate,
		Delete: resourceFirewallProfileProtocolOptionsCifsServerKeytabDelete,

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
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"keytab": &schema.Schema{
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
			"principal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallProfileProtocolOptionsCifsServerKeytabCreate(d *schema.ResourceData, m interface{}) error {
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
	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	obj, err := getObjectFirewallProfileProtocolOptionsCifsServerKeytab(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallProfileProtocolOptionsCifsServerKeytab resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("principal")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallProfileProtocolOptionsCifsServerKeytab(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallProfileProtocolOptionsCifsServerKeytab(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallProfileProtocolOptionsCifsServerKeytab resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallProfileProtocolOptionsCifsServerKeytab(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallProfileProtocolOptionsCifsServerKeytab resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "principal"))

	return resourceFirewallProfileProtocolOptionsCifsServerKeytabRead(d, m)
}

func resourceFirewallProfileProtocolOptionsCifsServerKeytabUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	obj, err := getObjectFirewallProfileProtocolOptionsCifsServerKeytab(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptionsCifsServerKeytab resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallProfileProtocolOptionsCifsServerKeytab(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptionsCifsServerKeytab resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "principal"))

	return resourceFirewallProfileProtocolOptionsCifsServerKeytabRead(d, m)
}

func resourceFirewallProfileProtocolOptionsCifsServerKeytabDelete(d *schema.ResourceData, m interface{}) error {
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
	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	wsParams["adom"] = adomv

	err = c.DeleteFirewallProfileProtocolOptionsCifsServerKeytab(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProfileProtocolOptionsCifsServerKeytab resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProfileProtocolOptionsCifsServerKeytabRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile_protocol_options := d.Get("profile_protocol_options").(string)
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
	if profile_protocol_options == "" {
		profile_protocol_options = importOptionChecking(m.(*FortiClient).Cfg, "profile_protocol_options")
		if profile_protocol_options == "" {
			return fmt.Errorf("Parameter profile_protocol_options is missing")
		}
		if err = d.Set("profile_protocol_options", profile_protocol_options); err != nil {
			return fmt.Errorf("Error set params profile_protocol_options: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	o, err := c.ReadFirewallProfileProtocolOptionsCifsServerKeytab(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallProfileProtocolOptionsCifsServerKeytab resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProfileProtocolOptionsCifsServerKeytab(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProfileProtocolOptionsCifsServerKeytab resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProfileProtocolOptionsCifsServerKeytabKeytab3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsServerKeytabPrincipal3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallProfileProtocolOptionsCifsServerKeytab(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("keytab", flattenFirewallProfileProtocolOptionsCifsServerKeytabKeytab3rdl(o["keytab"], d, "keytab")); err != nil {
		if vv, ok := fortiAPIPatch(o["keytab"], "FirewallProfileProtocolOptionsCifsServerKeytab-Keytab"); ok {
			if err = d.Set("keytab", vv); err != nil {
				return fmt.Errorf("Error reading keytab: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keytab: %v", err)
		}
	}

	if err = d.Set("principal", flattenFirewallProfileProtocolOptionsCifsServerKeytabPrincipal3rdl(o["principal"], d, "principal")); err != nil {
		if vv, ok := fortiAPIPatch(o["principal"], "FirewallProfileProtocolOptionsCifsServerKeytab-Principal"); ok {
			if err = d.Set("principal", vv); err != nil {
				return fmt.Errorf("Error reading principal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading principal: %v", err)
		}
	}

	return nil
}

func flattenFirewallProfileProtocolOptionsCifsServerKeytabFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProfileProtocolOptionsCifsServerKeytabKeytab3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsServerKeytabPassword3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsServerKeytabPrincipal3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProfileProtocolOptionsCifsServerKeytab(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("keytab"); ok || d.HasChange("keytab") {
		t, err := expandFirewallProfileProtocolOptionsCifsServerKeytabKeytab3rdl(d, v, "keytab")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keytab"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandFirewallProfileProtocolOptionsCifsServerKeytabPassword3rdl(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("principal"); ok || d.HasChange("principal") {
		t, err := expandFirewallProfileProtocolOptionsCifsServerKeytabPrincipal3rdl(d, v, "principal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["principal"] = t
		}
	}

	return &obj, nil
}
