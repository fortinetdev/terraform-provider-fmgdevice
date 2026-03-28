// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> CASB user activity tenant extraction filters.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbUserActivityMatchTenantExtractionFilters() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbUserActivityMatchTenantExtractionFiltersCreate,
		Read:   resourceCasbUserActivityMatchTenantExtractionFiltersRead,
		Update: resourceCasbUserActivityMatchTenantExtractionFiltersUpdate,
		Delete: resourceCasbUserActivityMatchTenantExtractionFiltersDelete,

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
			"user_activity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"body_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"header_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"place": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCasbUserActivityMatchTenantExtractionFiltersCreate(d *schema.ResourceData, m interface{}) error {
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
	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	obj, err := getObjectCasbUserActivityMatchTenantExtractionFilters(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbUserActivityMatchTenantExtractionFilters resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbUserActivityMatchTenantExtractionFilters(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbUserActivityMatchTenantExtractionFilters(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbUserActivityMatchTenantExtractionFilters resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateCasbUserActivityMatchTenantExtractionFilters(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbUserActivityMatchTenantExtractionFilters resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceCasbUserActivityMatchTenantExtractionFiltersRead(d, m)
}

func resourceCasbUserActivityMatchTenantExtractionFiltersUpdate(d *schema.ResourceData, m interface{}) error {
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
	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	obj, err := getObjectCasbUserActivityMatchTenantExtractionFilters(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivityMatchTenantExtractionFilters resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbUserActivityMatchTenantExtractionFilters(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivityMatchTenantExtractionFilters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceCasbUserActivityMatchTenantExtractionFiltersRead(d, m)
}

func resourceCasbUserActivityMatchTenantExtractionFiltersDelete(d *schema.ResourceData, m interface{}) error {
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
	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	wsParams["adom"] = adomv

	err = c.DeleteCasbUserActivityMatchTenantExtractionFilters(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbUserActivityMatchTenantExtractionFilters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbUserActivityMatchTenantExtractionFiltersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
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
	if user_activity == "" {
		user_activity = importOptionChecking(m.(*FortiClient).Cfg, "user_activity")
		if user_activity == "" {
			return fmt.Errorf("Parameter user_activity is missing")
		}
		if err = d.Set("user_activity", user_activity); err != nil {
			return fmt.Errorf("Error set params user_activity: %v", err)
		}
	}
	if match == "" {
		match = importOptionChecking(m.(*FortiClient).Cfg, "match")
		if match == "" {
			return fmt.Errorf("Parameter match is missing")
		}
		if err = d.Set("match", match); err != nil {
			return fmt.Errorf("Error set params match: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	o, err := c.ReadCasbUserActivityMatchTenantExtractionFilters(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbUserActivityMatchTenantExtractionFilters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbUserActivityMatchTenantExtractionFilters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbUserActivityMatchTenantExtractionFilters resource from API: %v", err)
	}
	return nil
}

func flattenCasbUserActivityMatchTenantExtractionFiltersBodyType4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersDirection4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersHeaderName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersPlace4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbUserActivityMatchTenantExtractionFilters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("body_type", flattenCasbUserActivityMatchTenantExtractionFiltersBodyType4thl(o["body-type"], d, "body_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["body-type"], "CasbUserActivityMatchTenantExtractionFilters-BodyType"); ok {
			if err = d.Set("body_type", vv); err != nil {
				return fmt.Errorf("Error reading body_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading body_type: %v", err)
		}
	}

	if err = d.Set("direction", flattenCasbUserActivityMatchTenantExtractionFiltersDirection4thl(o["direction"], d, "direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["direction"], "CasbUserActivityMatchTenantExtractionFilters-Direction"); ok {
			if err = d.Set("direction", vv); err != nil {
				return fmt.Errorf("Error reading direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("header_name", flattenCasbUserActivityMatchTenantExtractionFiltersHeaderName4thl(o["header-name"], d, "header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-name"], "CasbUserActivityMatchTenantExtractionFilters-HeaderName"); ok {
			if err = d.Set("header_name", vv); err != nil {
				return fmt.Errorf("Error reading header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenCasbUserActivityMatchTenantExtractionFiltersId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "CasbUserActivityMatchTenantExtractionFilters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("place", flattenCasbUserActivityMatchTenantExtractionFiltersPlace4thl(o["place"], d, "place")); err != nil {
		if vv, ok := fortiAPIPatch(o["place"], "CasbUserActivityMatchTenantExtractionFilters-Place"); ok {
			if err = d.Set("place", vv); err != nil {
				return fmt.Errorf("Error reading place: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading place: %v", err)
		}
	}

	return nil
}

func flattenCasbUserActivityMatchTenantExtractionFiltersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbUserActivityMatchTenantExtractionFiltersBodyType4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersDirection4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersHeaderName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersPlace4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbUserActivityMatchTenantExtractionFilters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("body_type"); ok || d.HasChange("body_type") {
		t, err := expandCasbUserActivityMatchTenantExtractionFiltersBodyType4thl(d, v, "body_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["body-type"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok || d.HasChange("direction") {
		t, err := expandCasbUserActivityMatchTenantExtractionFiltersDirection4thl(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok || d.HasChange("header_name") {
		t, err := expandCasbUserActivityMatchTenantExtractionFiltersHeaderName4thl(d, v, "header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandCasbUserActivityMatchTenantExtractionFiltersId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("place"); ok || d.HasChange("place") {
		t, err := expandCasbUserActivityMatchTenantExtractionFiltersPlace4thl(d, v, "place")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["place"] = t
		}
	}

	return &obj, nil
}
