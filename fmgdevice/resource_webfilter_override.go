// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard Web Filter administrative overrides.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterOverrideCreate,
		Read:   resourceWebfilterOverrideRead,
		Update: resourceWebfilterOverrideUpdate,
		Delete: resourceWebfilterOverrideDelete,

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
			"expires": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"initiator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"new_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"old_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"user_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebfilterOverrideCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterOverride(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterOverride resource while getting object: %v", err)
	}

	v, err := c.CreateWebfilterOverride(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WebfilterOverride resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceWebfilterOverrideRead(d, m)
		} else {
			return fmt.Errorf("Error creating WebfilterOverride resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebfilterOverrideRead(d, m)
}

func resourceWebfilterOverrideUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterOverride(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterOverride resource while getting object: %v", err)
	}

	_, err = c.UpdateWebfilterOverride(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterOverride resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebfilterOverrideRead(d, m)
}

func resourceWebfilterOverrideDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebfilterOverride(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterOverrideRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterOverride(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterOverride(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterOverride resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterOverrideExpires(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenWebfilterOverrideId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideInitiator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideNewProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterOverrideOldProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterOverrideScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterOverrideUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWebfilterOverride(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("expires", flattenWebfilterOverrideExpires(o["expires"], d, "expires")); err != nil {
		if vv, ok := fortiAPIPatch(o["expires"], "WebfilterOverride-Expires"); ok {
			if err = d.Set("expires", vv); err != nil {
				return fmt.Errorf("Error reading expires: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expires: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWebfilterOverrideId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WebfilterOverride-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("initiator", flattenWebfilterOverrideInitiator(o["initiator"], d, "initiator")); err != nil {
		if vv, ok := fortiAPIPatch(o["initiator"], "WebfilterOverride-Initiator"); ok {
			if err = d.Set("initiator", vv); err != nil {
				return fmt.Errorf("Error reading initiator: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading initiator: %v", err)
		}
	}

	if err = d.Set("ip", flattenWebfilterOverrideIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "WebfilterOverride-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip6", flattenWebfilterOverrideIp6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "WebfilterOverride-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("new_profile", flattenWebfilterOverrideNewProfile(o["new-profile"], d, "new_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["new-profile"], "WebfilterOverride-NewProfile"); ok {
			if err = d.Set("new_profile", vv); err != nil {
				return fmt.Errorf("Error reading new_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading new_profile: %v", err)
		}
	}

	if err = d.Set("old_profile", flattenWebfilterOverrideOldProfile(o["old-profile"], d, "old_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["old-profile"], "WebfilterOverride-OldProfile"); ok {
			if err = d.Set("old_profile", vv); err != nil {
				return fmt.Errorf("Error reading old_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading old_profile: %v", err)
		}
	}

	if err = d.Set("scope", flattenWebfilterOverrideScope(o["scope"], d, "scope")); err != nil {
		if vv, ok := fortiAPIPatch(o["scope"], "WebfilterOverride-Scope"); ok {
			if err = d.Set("scope", vv); err != nil {
				return fmt.Errorf("Error reading scope: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scope: %v", err)
		}
	}

	if err = d.Set("status", flattenWebfilterOverrideStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WebfilterOverride-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("user", flattenWebfilterOverrideUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "WebfilterOverride-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("user_group", flattenWebfilterOverrideUserGroup(o["user-group"], d, "user_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-group"], "WebfilterOverride-UserGroup"); ok {
			if err = d.Set("user_group", vv); err != nil {
				return fmt.Errorf("Error reading user_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	return nil
}

func flattenWebfilterOverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterOverrideExpires(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideInitiator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideNewProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterOverrideOldProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterOverrideScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterOverrideUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWebfilterOverride(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("expires"); ok || d.HasChange("expires") {
		t, err := expandWebfilterOverrideExpires(d, v, "expires")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expires"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWebfilterOverrideId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("initiator"); ok || d.HasChange("initiator") {
		t, err := expandWebfilterOverrideInitiator(d, v, "initiator")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["initiator"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandWebfilterOverrideIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok || d.HasChange("ip6") {
		t, err := expandWebfilterOverrideIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("new_profile"); ok || d.HasChange("new_profile") {
		t, err := expandWebfilterOverrideNewProfile(d, v, "new_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["new-profile"] = t
		}
	}

	if v, ok := d.GetOk("old_profile"); ok || d.HasChange("old_profile") {
		t, err := expandWebfilterOverrideOldProfile(d, v, "old_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["old-profile"] = t
		}
	}

	if v, ok := d.GetOk("scope"); ok || d.HasChange("scope") {
		t, err := expandWebfilterOverrideScope(d, v, "scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWebfilterOverrideStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandWebfilterOverrideUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok || d.HasChange("user_group") {
		t, err := expandWebfilterOverrideUserGroup(d, v, "user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	return &obj, nil
}
