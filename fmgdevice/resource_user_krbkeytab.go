// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Kerberos keytab entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserKrbKeytab() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserKrbKeytabCreate,
		Read:   resourceUserKrbKeytabRead,
		Update: resourceUserKrbKeytabUpdate,
		Delete: resourceUserKrbKeytabDelete,

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
			"keytab": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"pac_data": &schema.Schema{
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
			"principal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceUserKrbKeytabCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserKrbKeytab(d)
	if err != nil {
		return fmt.Errorf("Error creating UserKrbKeytab resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserKrbKeytab(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserKrbKeytab(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserKrbKeytab resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserKrbKeytab(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserKrbKeytab resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserKrbKeytabRead(d, m)
}

func resourceUserKrbKeytabUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserKrbKeytab(d)
	if err != nil {
		return fmt.Errorf("Error updating UserKrbKeytab resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserKrbKeytab(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserKrbKeytab resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserKrbKeytabRead(d, m)
}

func resourceUserKrbKeytabDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserKrbKeytab(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserKrbKeytab resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserKrbKeytabRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserKrbKeytab(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserKrbKeytab resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserKrbKeytab(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserKrbKeytab resource from API: %v", err)
	}
	return nil
}

func flattenUserKrbKeytabKeytab(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserKrbKeytabLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserKrbKeytabName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserKrbKeytabPacData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserKrbKeytabPrincipal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserKrbKeytab(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("keytab", flattenUserKrbKeytabKeytab(o["keytab"], d, "keytab")); err != nil {
		if vv, ok := fortiAPIPatch(o["keytab"], "UserKrbKeytab-Keytab"); ok {
			if err = d.Set("keytab", vv); err != nil {
				return fmt.Errorf("Error reading keytab: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keytab: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserKrbKeytabLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "UserKrbKeytab-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("name", flattenUserKrbKeytabName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserKrbKeytab-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pac_data", flattenUserKrbKeytabPacData(o["pac-data"], d, "pac_data")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-data"], "UserKrbKeytab-PacData"); ok {
			if err = d.Set("pac_data", vv); err != nil {
				return fmt.Errorf("Error reading pac_data: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_data: %v", err)
		}
	}

	if err = d.Set("principal", flattenUserKrbKeytabPrincipal(o["principal"], d, "principal")); err != nil {
		if vv, ok := fortiAPIPatch(o["principal"], "UserKrbKeytab-Principal"); ok {
			if err = d.Set("principal", vv); err != nil {
				return fmt.Errorf("Error reading principal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading principal: %v", err)
		}
	}

	return nil
}

func flattenUserKrbKeytabFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserKrbKeytabKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserKrbKeytabLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserKrbKeytabName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserKrbKeytabPacData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserKrbKeytabPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserKrbKeytabPrincipal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserKrbKeytab(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("keytab"); ok || d.HasChange("keytab") {
		t, err := expandUserKrbKeytabKeytab(d, v, "keytab")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keytab"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandUserKrbKeytabLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserKrbKeytabName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pac_data"); ok || d.HasChange("pac_data") {
		t, err := expandUserKrbKeytabPacData(d, v, "pac_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-data"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandUserKrbKeytabPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("principal"); ok || d.HasChange("principal") {
		t, err := expandUserKrbKeytabPrincipal(d, v, "principal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["principal"] = t
		}
	}

	return &obj, nil
}
