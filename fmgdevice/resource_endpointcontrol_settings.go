// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure endpoint control settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEndpointControlSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceEndpointControlSettingsUpdate,
		Read:   resourceEndpointControlSettingsRead,
		Update: resourceEndpointControlSettingsUpdate,
		Delete: resourceEndpointControlSettingsDelete,

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
			"forticlient_disconnect_unsupported_client": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forticlient_keepalive_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"forticlient_sys_update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"forticlient_user_avatar": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceEndpointControlSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEndpointControlSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateEndpointControlSettings(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("EndpointControlSettings")

	return resourceEndpointControlSettingsRead(d, m)
}

func resourceEndpointControlSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteEndpointControlSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting EndpointControlSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEndpointControlSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadEndpointControlSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEndpointControlSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlSettings resource from API: %v", err)
	}
	return nil
}

func flattenEndpointControlSettingsForticlientDisconnectUnsupportedClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientKeepaliveInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientSysUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientUserAvatar(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlSettingsOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEndpointControlSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("forticlient_disconnect_unsupported_client", flattenEndpointControlSettingsForticlientDisconnectUnsupportedClient(o["forticlient-disconnect-unsupported-client"], d, "forticlient_disconnect_unsupported_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-disconnect-unsupported-client"], "EndpointControlSettings-ForticlientDisconnectUnsupportedClient"); ok {
			if err = d.Set("forticlient_disconnect_unsupported_client", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_disconnect_unsupported_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_disconnect_unsupported_client: %v", err)
		}
	}

	if err = d.Set("forticlient_keepalive_interval", flattenEndpointControlSettingsForticlientKeepaliveInterval(o["forticlient-keepalive-interval"], d, "forticlient_keepalive_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-keepalive-interval"], "EndpointControlSettings-ForticlientKeepaliveInterval"); ok {
			if err = d.Set("forticlient_keepalive_interval", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_keepalive_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_keepalive_interval: %v", err)
		}
	}

	if err = d.Set("forticlient_sys_update_interval", flattenEndpointControlSettingsForticlientSysUpdateInterval(o["forticlient-sys-update-interval"], d, "forticlient_sys_update_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-sys-update-interval"], "EndpointControlSettings-ForticlientSysUpdateInterval"); ok {
			if err = d.Set("forticlient_sys_update_interval", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_sys_update_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_sys_update_interval: %v", err)
		}
	}

	if err = d.Set("forticlient_user_avatar", flattenEndpointControlSettingsForticlientUserAvatar(o["forticlient-user-avatar"], d, "forticlient_user_avatar")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-user-avatar"], "EndpointControlSettings-ForticlientUserAvatar"); ok {
			if err = d.Set("forticlient_user_avatar", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_user_avatar: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_user_avatar: %v", err)
		}
	}

	if err = d.Set("override", flattenEndpointControlSettingsOverride(o["override"], d, "override")); err != nil {
		if vv, ok := fortiAPIPatch(o["override"], "EndpointControlSettings-Override"); ok {
			if err = d.Set("override", vv); err != nil {
				return fmt.Errorf("Error reading override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	return nil
}

func flattenEndpointControlSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEndpointControlSettingsForticlientDisconnectUnsupportedClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientKeepaliveInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientSysUpdateInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientUserAvatar(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEndpointControlSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("forticlient_disconnect_unsupported_client"); ok || d.HasChange("forticlient_disconnect_unsupported_client") {
		t, err := expandEndpointControlSettingsForticlientDisconnectUnsupportedClient(d, v, "forticlient_disconnect_unsupported_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-disconnect-unsupported-client"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_keepalive_interval"); ok || d.HasChange("forticlient_keepalive_interval") {
		t, err := expandEndpointControlSettingsForticlientKeepaliveInterval(d, v, "forticlient_keepalive_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-keepalive-interval"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_sys_update_interval"); ok || d.HasChange("forticlient_sys_update_interval") {
		t, err := expandEndpointControlSettingsForticlientSysUpdateInterval(d, v, "forticlient_sys_update_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-sys-update-interval"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_user_avatar"); ok || d.HasChange("forticlient_user_avatar") {
		t, err := expandEndpointControlSettingsForticlientUserAvatar(d, v, "forticlient_user_avatar")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-user-avatar"] = t
		}
	}

	if v, ok := d.GetOk("override"); ok || d.HasChange("override") {
		t, err := expandEndpointControlSettingsOverride(d, v, "override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	return &obj, nil
}
