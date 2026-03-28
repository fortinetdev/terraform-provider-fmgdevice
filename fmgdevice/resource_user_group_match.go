// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Group matches.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserGroupMatch() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserGroupMatchCreate,
		Read:   resourceUserGroupMatchRead,
		Update: resourceUserGroupMatchUpdate,
		Delete: resourceUserGroupMatchDelete,

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
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"_gui_meta": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserGroupMatchCreate(d *schema.ResourceData, m interface{}) error {
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
	group := d.Get("group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["group"] = group

	obj, err := getObjectUserGroupMatch(d)
	if err != nil {
		return fmt.Errorf("Error creating UserGroupMatch resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserGroupMatch(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserGroupMatch(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserGroupMatch resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserGroupMatch(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserGroupMatch resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceUserGroupMatchRead(d, m)
}

func resourceUserGroupMatchUpdate(d *schema.ResourceData, m interface{}) error {
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
	group := d.Get("group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["group"] = group

	obj, err := getObjectUserGroupMatch(d)
	if err != nil {
		return fmt.Errorf("Error updating UserGroupMatch resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserGroupMatch(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserGroupMatch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceUserGroupMatchRead(d, m)
}

func resourceUserGroupMatchDelete(d *schema.ResourceData, m interface{}) error {
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
	group := d.Get("group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["group"] = group

	wsParams["adom"] = adomv

	err = c.DeleteUserGroupMatch(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserGroupMatch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserGroupMatchRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	group := d.Get("group").(string)
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
	if group == "" {
		group = importOptionChecking(m.(*FortiClient).Cfg, "group")
		if group == "" {
			return fmt.Errorf("Parameter group is missing")
		}
		if err = d.Set("group", group); err != nil {
			return fmt.Errorf("Error set params group: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["group"] = group

	o, err := c.ReadUserGroupMatch(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserGroupMatch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserGroupMatch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserGroupMatch resource from API: %v", err)
	}
	return nil
}

func flattenUserGroupMatchGuiMeta2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupMatchGroupName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupMatchId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupMatchServerName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectUserGroupMatch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("_gui_meta", flattenUserGroupMatchGuiMeta2edl(o["_gui_meta"], d, "_gui_meta")); err != nil {
		if vv, ok := fortiAPIPatch(o["_gui_meta"], "UserGroupMatch-GuiMeta"); ok {
			if err = d.Set("_gui_meta", vv); err != nil {
				return fmt.Errorf("Error reading _gui_meta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _gui_meta: %v", err)
		}
	}

	if err = d.Set("group_name", flattenUserGroupMatchGroupName2edl(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "UserGroupMatch-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserGroupMatchId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "UserGroupMatch-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("server_name", flattenUserGroupMatchServerName2edl(o["server-name"], d, "server_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-name"], "UserGroupMatch-ServerName"); ok {
			if err = d.Set("server_name", vv); err != nil {
				return fmt.Errorf("Error reading server_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	return nil
}

func flattenUserGroupMatchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserGroupMatchGuiMeta2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMatchGroupName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMatchId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMatchServerName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectUserGroupMatch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_gui_meta"); ok || d.HasChange("_gui_meta") {
		t, err := expandUserGroupMatchGuiMeta2edl(d, v, "_gui_meta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_gui_meta"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandUserGroupMatchGroupName2edl(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandUserGroupMatchId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok || d.HasChange("server_name") {
		t, err := expandUserGroupMatchServerName2edl(d, v, "server_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	return &obj, nil
}
