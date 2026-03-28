// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> PPID filters list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSctpFilterProfilePpidFilters() *schema.Resource {
	return &schema.Resource{
		Create: resourceSctpFilterProfilePpidFiltersCreate,
		Read:   resourceSctpFilterProfilePpidFiltersRead,
		Update: resourceSctpFilterProfilePpidFiltersUpdate,
		Delete: resourceSctpFilterProfilePpidFiltersDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ppid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSctpFilterProfilePpidFiltersCreate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectSctpFilterProfilePpidFilters(d)
	if err != nil {
		return fmt.Errorf("Error creating SctpFilterProfilePpidFilters resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSctpFilterProfilePpidFilters(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSctpFilterProfilePpidFilters(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SctpFilterProfilePpidFilters resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateSctpFilterProfilePpidFilters(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SctpFilterProfilePpidFilters resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceSctpFilterProfilePpidFiltersRead(d, m)
			} else {
				return fmt.Errorf("Error creating SctpFilterProfilePpidFilters resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSctpFilterProfilePpidFiltersRead(d, m)
}

func resourceSctpFilterProfilePpidFiltersUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectSctpFilterProfilePpidFilters(d)
	if err != nil {
		return fmt.Errorf("Error updating SctpFilterProfilePpidFilters resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSctpFilterProfilePpidFilters(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SctpFilterProfilePpidFilters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSctpFilterProfilePpidFiltersRead(d, m)
}

func resourceSctpFilterProfilePpidFiltersDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteSctpFilterProfilePpidFilters(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SctpFilterProfilePpidFilters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSctpFilterProfilePpidFiltersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadSctpFilterProfilePpidFilters(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SctpFilterProfilePpidFilters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSctpFilterProfilePpidFilters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SctpFilterProfilePpidFilters resource from API: %v", err)
	}
	return nil
}

func flattenSctpFilterProfilePpidFiltersAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSctpFilterProfilePpidFiltersComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSctpFilterProfilePpidFiltersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSctpFilterProfilePpidFiltersPpid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSctpFilterProfilePpidFilters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenSctpFilterProfilePpidFiltersAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SctpFilterProfilePpidFilters-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("comment", flattenSctpFilterProfilePpidFiltersComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "SctpFilterProfilePpidFilters-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSctpFilterProfilePpidFiltersId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SctpFilterProfilePpidFilters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ppid", flattenSctpFilterProfilePpidFiltersPpid2edl(o["ppid"], d, "ppid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppid"], "SctpFilterProfilePpidFilters-Ppid"); ok {
			if err = d.Set("ppid", vv); err != nil {
				return fmt.Errorf("Error reading ppid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppid: %v", err)
		}
	}

	return nil
}

func flattenSctpFilterProfilePpidFiltersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSctpFilterProfilePpidFiltersAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfilePpidFiltersComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfilePpidFiltersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfilePpidFiltersPpid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSctpFilterProfilePpidFilters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSctpFilterProfilePpidFiltersAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandSctpFilterProfilePpidFiltersComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSctpFilterProfilePpidFiltersId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ppid"); ok || d.HasChange("ppid") {
		t, err := expandSctpFilterProfilePpidFiltersPpid2edl(d, v, "ppid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppid"] = t
		}
	}

	return &obj, nil
}
