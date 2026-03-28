// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> AntiPhishing entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterProfileAntiphishInspectionEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterProfileAntiphishInspectionEntriesCreate,
		Read:   resourceWebfilterProfileAntiphishInspectionEntriesRead,
		Update: resourceWebfilterProfileAntiphishInspectionEntriesUpdate,
		Delete: resourceWebfilterProfileAntiphishInspectionEntriesDelete,

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
			"fortiguard_category": &schema.Schema{
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
		},
	}
}

func resourceWebfilterProfileAntiphishInspectionEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterProfileAntiphishInspectionEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterProfileAntiphishInspectionEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebfilterProfileAntiphishInspectionEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebfilterProfileAntiphishInspectionEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebfilterProfileAntiphishInspectionEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebfilterProfileAntiphishInspectionEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebfilterProfileAntiphishInspectionEntries resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebfilterProfileAntiphishInspectionEntriesRead(d, m)
}

func resourceWebfilterProfileAntiphishInspectionEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterProfileAntiphishInspectionEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfileAntiphishInspectionEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterProfileAntiphishInspectionEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfileAntiphishInspectionEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebfilterProfileAntiphishInspectionEntriesRead(d, m)
}

func resourceWebfilterProfileAntiphishInspectionEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebfilterProfileAntiphishInspectionEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterProfileAntiphishInspectionEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterProfileAntiphishInspectionEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterProfileAntiphishInspectionEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterProfileAntiphishInspectionEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterProfileAntiphishInspectionEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterProfileAntiphishInspectionEntries resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterProfileAntiphishInspectionEntriesAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishInspectionEntriesFortiguardCategory3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileAntiphishInspectionEntriesName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterProfileAntiphishInspectionEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenWebfilterProfileAntiphishInspectionEntriesAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "WebfilterProfileAntiphishInspectionEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fortiguard_category", flattenWebfilterProfileAntiphishInspectionEntriesFortiguardCategory3rdl(o["fortiguard-category"], d, "fortiguard_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-category"], "WebfilterProfileAntiphishInspectionEntries-FortiguardCategory"); ok {
			if err = d.Set("fortiguard_category", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_category: %v", err)
		}
	}

	if err = d.Set("name", flattenWebfilterProfileAntiphishInspectionEntriesName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebfilterProfileAntiphishInspectionEntries-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenWebfilterProfileAntiphishInspectionEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterProfileAntiphishInspectionEntriesAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishInspectionEntriesFortiguardCategory3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileAntiphishInspectionEntriesName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterProfileAntiphishInspectionEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandWebfilterProfileAntiphishInspectionEntriesAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_category"); ok || d.HasChange("fortiguard_category") {
		t, err := expandWebfilterProfileAntiphishInspectionEntriesFortiguardCategory3rdl(d, v, "fortiguard_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-category"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebfilterProfileAntiphishInspectionEntriesName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
