// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> FortiGuard DNS domain filters.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDnsfilterProfileFtgdDnsFilters() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsfilterProfileFtgdDnsFiltersCreate,
		Read:   resourceDnsfilterProfileFtgdDnsFiltersRead,
		Update: resourceDnsfilterProfileFtgdDnsFiltersUpdate,
		Delete: resourceDnsfilterProfileFtgdDnsFiltersDelete,

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
			"category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDnsfilterProfileFtgdDnsFiltersCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDnsfilterProfileFtgdDnsFilters(d)
	if err != nil {
		return fmt.Errorf("Error creating DnsfilterProfileFtgdDnsFilters resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDnsfilterProfileFtgdDnsFilters(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDnsfilterProfileFtgdDnsFilters(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DnsfilterProfileFtgdDnsFilters resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateDnsfilterProfileFtgdDnsFilters(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DnsfilterProfileFtgdDnsFilters resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceDnsfilterProfileFtgdDnsFiltersRead(d, m)
			} else {
				return fmt.Errorf("Error creating DnsfilterProfileFtgdDnsFilters resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDnsfilterProfileFtgdDnsFiltersRead(d, m)
}

func resourceDnsfilterProfileFtgdDnsFiltersUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDnsfilterProfileFtgdDnsFilters(d)
	if err != nil {
		return fmt.Errorf("Error updating DnsfilterProfileFtgdDnsFilters resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDnsfilterProfileFtgdDnsFilters(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DnsfilterProfileFtgdDnsFilters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDnsfilterProfileFtgdDnsFiltersRead(d, m)
}

func resourceDnsfilterProfileFtgdDnsFiltersDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteDnsfilterProfileFtgdDnsFilters(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DnsfilterProfileFtgdDnsFilters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDnsfilterProfileFtgdDnsFiltersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDnsfilterProfileFtgdDnsFilters(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DnsfilterProfileFtgdDnsFilters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDnsfilterProfileFtgdDnsFilters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DnsfilterProfileFtgdDnsFilters resource from API: %v", err)
	}
	return nil
}

func flattenDnsfilterProfileFtgdDnsFiltersAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileFtgdDnsFiltersCategory3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDnsfilterProfileFtgdDnsFiltersId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileFtgdDnsFiltersLog3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDnsfilterProfileFtgdDnsFilters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenDnsfilterProfileFtgdDnsFiltersAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "DnsfilterProfileFtgdDnsFilters-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("category", flattenDnsfilterProfileFtgdDnsFiltersCategory3rdl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "DnsfilterProfileFtgdDnsFilters-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("fosid", flattenDnsfilterProfileFtgdDnsFiltersId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "DnsfilterProfileFtgdDnsFilters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("log", flattenDnsfilterProfileFtgdDnsFiltersLog3rdl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "DnsfilterProfileFtgdDnsFilters-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	return nil
}

func flattenDnsfilterProfileFtgdDnsFiltersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDnsfilterProfileFtgdDnsFiltersAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDnsFiltersCategory3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDnsfilterProfileFtgdDnsFiltersId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDnsFiltersLog3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDnsfilterProfileFtgdDnsFilters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandDnsfilterProfileFtgdDnsFiltersAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandDnsfilterProfileFtgdDnsFiltersCategory3rdl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandDnsfilterProfileFtgdDnsFiltersId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandDnsfilterProfileFtgdDnsFiltersLog3rdl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	return &obj, nil
}
