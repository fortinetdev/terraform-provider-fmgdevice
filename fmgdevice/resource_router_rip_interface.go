// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: RIP interface configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterRipInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterRipInterfaceCreate,
		Read:   resourceRouterRipInterfaceRead,
		Update: resourceRouterRipInterfaceUpdate,
		Delete: resourceRouterRipInterfaceDelete,

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
			"auth_keychain": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_string": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"receive_version": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"send_version": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"send_version2_broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"split_horizon": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_horizon_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterRipInterfaceCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectRouterRipInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterRipInterface resource while getting object: %v", err)
	}

	_, err = c.CreateRouterRipInterface(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating RouterRipInterface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterRipInterfaceRead(d, m)
}

func resourceRouterRipInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectRouterRipInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterRipInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterRipInterface(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterRipInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterRipInterfaceRead(d, m)
}

func resourceRouterRipInterfaceDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteRouterRipInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterRipInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRipInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterRipInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterRipInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterRipInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterRipInterface resource from API: %v", err)
	}
	return nil
}

func flattenRouterRipInterfaceAuthKeychain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipInterfaceAuthMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipInterfaceFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipInterfaceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipInterfaceReceiveVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipInterfaceSendVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipInterfaceSendVersion2Broadcast2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipInterfaceSplitHorizon2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipInterfaceSplitHorizonStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterRipInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_keychain", flattenRouterRipInterfaceAuthKeychain2edl(o["auth-keychain"], d, "auth_keychain")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-keychain"], "RouterRipInterface-AuthKeychain"); ok {
			if err = d.Set("auth_keychain", vv); err != nil {
				return fmt.Errorf("Error reading auth_keychain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_keychain: %v", err)
		}
	}

	if err = d.Set("auth_mode", flattenRouterRipInterfaceAuthMode2edl(o["auth-mode"], d, "auth_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-mode"], "RouterRipInterface-AuthMode"); ok {
			if err = d.Set("auth_mode", vv); err != nil {
				return fmt.Errorf("Error reading auth_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_mode: %v", err)
		}
	}

	if err = d.Set("flags", flattenRouterRipInterfaceFlags2edl(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "RouterRipInterface-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterRipInterfaceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterRipInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("receive_version", flattenRouterRipInterfaceReceiveVersion2edl(o["receive-version"], d, "receive_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["receive-version"], "RouterRipInterface-ReceiveVersion"); ok {
			if err = d.Set("receive_version", vv); err != nil {
				return fmt.Errorf("Error reading receive_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading receive_version: %v", err)
		}
	}

	if err = d.Set("send_version", flattenRouterRipInterfaceSendVersion2edl(o["send-version"], d, "send_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-version"], "RouterRipInterface-SendVersion"); ok {
			if err = d.Set("send_version", vv); err != nil {
				return fmt.Errorf("Error reading send_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_version: %v", err)
		}
	}

	if err = d.Set("send_version2_broadcast", flattenRouterRipInterfaceSendVersion2Broadcast2edl(o["send-version2-broadcast"], d, "send_version2_broadcast")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-version2-broadcast"], "RouterRipInterface-SendVersion2Broadcast"); ok {
			if err = d.Set("send_version2_broadcast", vv); err != nil {
				return fmt.Errorf("Error reading send_version2_broadcast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_version2_broadcast: %v", err)
		}
	}

	if err = d.Set("split_horizon", flattenRouterRipInterfaceSplitHorizon2edl(o["split-horizon"], d, "split_horizon")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-horizon"], "RouterRipInterface-SplitHorizon"); ok {
			if err = d.Set("split_horizon", vv); err != nil {
				return fmt.Errorf("Error reading split_horizon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_horizon: %v", err)
		}
	}

	if err = d.Set("split_horizon_status", flattenRouterRipInterfaceSplitHorizonStatus2edl(o["split-horizon-status"], d, "split_horizon_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-horizon-status"], "RouterRipInterface-SplitHorizonStatus"); ok {
			if err = d.Set("split_horizon_status", vv); err != nil {
				return fmt.Errorf("Error reading split_horizon_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_horizon_status: %v", err)
		}
	}

	return nil
}

func flattenRouterRipInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterRipInterfaceAuthKeychain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipInterfaceAuthMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceAuthString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipInterfaceFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceReceiveVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipInterfaceSendVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipInterfaceSendVersion2Broadcast2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceSplitHorizon2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceSplitHorizonStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterRipInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_keychain"); ok || d.HasChange("auth_keychain") {
		t, err := expandRouterRipInterfaceAuthKeychain2edl(d, v, "auth_keychain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-keychain"] = t
		}
	}

	if v, ok := d.GetOk("auth_mode"); ok || d.HasChange("auth_mode") {
		t, err := expandRouterRipInterfaceAuthMode2edl(d, v, "auth_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-mode"] = t
		}
	}

	if v, ok := d.GetOk("auth_string"); ok || d.HasChange("auth_string") {
		t, err := expandRouterRipInterfaceAuthString2edl(d, v, "auth_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-string"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandRouterRipInterfaceFlags2edl(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterRipInterfaceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("receive_version"); ok || d.HasChange("receive_version") {
		t, err := expandRouterRipInterfaceReceiveVersion2edl(d, v, "receive_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["receive-version"] = t
		}
	}

	if v, ok := d.GetOk("send_version"); ok || d.HasChange("send_version") {
		t, err := expandRouterRipInterfaceSendVersion2edl(d, v, "send_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-version"] = t
		}
	}

	if v, ok := d.GetOk("send_version2_broadcast"); ok || d.HasChange("send_version2_broadcast") {
		t, err := expandRouterRipInterfaceSendVersion2Broadcast2edl(d, v, "send_version2_broadcast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-version2-broadcast"] = t
		}
	}

	if v, ok := d.GetOk("split_horizon"); ok || d.HasChange("split_horizon") {
		t, err := expandRouterRipInterfaceSplitHorizon2edl(d, v, "split_horizon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-horizon"] = t
		}
	}

	if v, ok := d.GetOk("split_horizon_status"); ok || d.HasChange("split_horizon_status") {
		t, err := expandRouterRipInterfaceSplitHorizonStatus2edl(d, v, "split_horizon_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-horizon-status"] = t
		}
	}

	return &obj, nil
}
