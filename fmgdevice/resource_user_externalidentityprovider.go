// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure external identity provider.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserExternalIdentityProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserExternalIdentityProviderCreate,
		Read:   resourceUserExternalIdentityProviderRead,
		Update: resourceUserExternalIdentityProviderUpdate,
		Delete: resourceUserExternalIdentityProviderDelete,

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
			"group_attr_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_identity_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_attr_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceUserExternalIdentityProviderCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserExternalIdentityProvider(d)
	if err != nil {
		return fmt.Errorf("Error creating UserExternalIdentityProvider resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserExternalIdentityProvider(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserExternalIdentityProvider(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserExternalIdentityProvider resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserExternalIdentityProvider(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserExternalIdentityProvider resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserExternalIdentityProviderRead(d, m)
}

func resourceUserExternalIdentityProviderUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserExternalIdentityProvider(d)
	if err != nil {
		return fmt.Errorf("Error updating UserExternalIdentityProvider resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserExternalIdentityProvider(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserExternalIdentityProvider resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserExternalIdentityProviderRead(d, m)
}

func resourceUserExternalIdentityProviderDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserExternalIdentityProvider(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserExternalIdentityProvider resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserExternalIdentityProviderRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserExternalIdentityProvider(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserExternalIdentityProvider resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserExternalIdentityProvider(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserExternalIdentityProvider resource from API: %v", err)
	}
	return nil
}

func flattenUserExternalIdentityProviderGroupAttrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExternalIdentityProviderInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserExternalIdentityProviderInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExternalIdentityProviderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExternalIdentityProviderPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExternalIdentityProviderServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExternalIdentityProviderSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExternalIdentityProviderTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExternalIdentityProviderType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExternalIdentityProviderUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExternalIdentityProviderUserAttrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExternalIdentityProviderVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserExternalIdentityProviderVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserExternalIdentityProvider(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("group_attr_name", flattenUserExternalIdentityProviderGroupAttrName(o["group-attr-name"], d, "group_attr_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-attr-name"], "UserExternalIdentityProvider-GroupAttrName"); ok {
			if err = d.Set("group_attr_name", vv); err != nil {
				return fmt.Errorf("Error reading group_attr_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_attr_name: %v", err)
		}
	}

	if err = d.Set("interface", flattenUserExternalIdentityProviderInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "UserExternalIdentityProvider-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenUserExternalIdentityProviderInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "UserExternalIdentityProvider-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("name", flattenUserExternalIdentityProviderName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserExternalIdentityProvider-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenUserExternalIdentityProviderPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "UserExternalIdentityProvider-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("server_identity_check", flattenUserExternalIdentityProviderServerIdentityCheck(o["server-identity-check"], d, "server_identity_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-identity-check"], "UserExternalIdentityProvider-ServerIdentityCheck"); ok {
			if err = d.Set("server_identity_check", vv); err != nil {
				return fmt.Errorf("Error reading server_identity_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_identity_check: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserExternalIdentityProviderSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "UserExternalIdentityProvider-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("timeout", flattenUserExternalIdentityProviderTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "UserExternalIdentityProvider-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("type", flattenUserExternalIdentityProviderType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "UserExternalIdentityProvider-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("url", flattenUserExternalIdentityProviderUrl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "UserExternalIdentityProvider-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("user_attr_name", flattenUserExternalIdentityProviderUserAttrName(o["user-attr-name"], d, "user_attr_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-attr-name"], "UserExternalIdentityProvider-UserAttrName"); ok {
			if err = d.Set("user_attr_name", vv); err != nil {
				return fmt.Errorf("Error reading user_attr_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_attr_name: %v", err)
		}
	}

	if err = d.Set("version", flattenUserExternalIdentityProviderVersion(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "UserExternalIdentityProvider-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenUserExternalIdentityProviderVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "UserExternalIdentityProvider-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenUserExternalIdentityProviderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserExternalIdentityProviderGroupAttrName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExternalIdentityProviderInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserExternalIdentityProviderInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExternalIdentityProviderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExternalIdentityProviderPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExternalIdentityProviderServerIdentityCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExternalIdentityProviderSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExternalIdentityProviderTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExternalIdentityProviderType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExternalIdentityProviderUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExternalIdentityProviderUserAttrName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExternalIdentityProviderVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserExternalIdentityProviderVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserExternalIdentityProvider(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("group_attr_name"); ok || d.HasChange("group_attr_name") {
		t, err := expandUserExternalIdentityProviderGroupAttrName(d, v, "group_attr_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-attr-name"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandUserExternalIdentityProviderInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandUserExternalIdentityProviderInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserExternalIdentityProviderName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandUserExternalIdentityProviderPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("server_identity_check"); ok || d.HasChange("server_identity_check") {
		t, err := expandUserExternalIdentityProviderServerIdentityCheck(d, v, "server_identity_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandUserExternalIdentityProviderSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandUserExternalIdentityProviderTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandUserExternalIdentityProviderType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandUserExternalIdentityProviderUrl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("user_attr_name"); ok || d.HasChange("user_attr_name") {
		t, err := expandUserExternalIdentityProviderUserAttrName(d, v, "user_attr_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-attr-name"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandUserExternalIdentityProviderVersion(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandUserExternalIdentityProviderVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
