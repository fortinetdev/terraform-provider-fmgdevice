// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Spam filter mime header content.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterMheaderEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterMheaderEntriesCreate,
		Read:   resourceEmailfilterMheaderEntriesRead,
		Update: resourceEmailfilterMheaderEntriesUpdate,
		Delete: resourceEmailfilterMheaderEntriesDelete,

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
			"mheader": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fieldbody": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fieldname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"pattern_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceEmailfilterMheaderEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	mheader := d.Get("mheader").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mheader"] = mheader

	obj, err := getObjectEmailfilterMheaderEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating EmailfilterMheaderEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadEmailfilterMheaderEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateEmailfilterMheaderEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating EmailfilterMheaderEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateEmailfilterMheaderEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating EmailfilterMheaderEntries resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterMheaderEntriesRead(d, m)
}

func resourceEmailfilterMheaderEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	mheader := d.Get("mheader").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mheader"] = mheader

	obj, err := getObjectEmailfilterMheaderEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterMheaderEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEmailfilterMheaderEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterMheaderEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterMheaderEntriesRead(d, m)
}

func resourceEmailfilterMheaderEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	mheader := d.Get("mheader").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mheader"] = mheader

	wsParams["adom"] = adomv

	err = c.DeleteEmailfilterMheaderEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterMheaderEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterMheaderEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	mheader := d.Get("mheader").(string)
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
	if mheader == "" {
		mheader = importOptionChecking(m.(*FortiClient).Cfg, "mheader")
		if mheader == "" {
			return fmt.Errorf("Parameter mheader is missing")
		}
		if err = d.Set("mheader", mheader); err != nil {
			return fmt.Errorf("Error set params mheader: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mheader"] = mheader

	o, err := c.ReadEmailfilterMheaderEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EmailfilterMheaderEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterMheaderEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterMheaderEntries resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterMheaderEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesFieldbody2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesFieldname2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesPatternType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEmailfilterMheaderEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenEmailfilterMheaderEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "EmailfilterMheaderEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fieldbody", flattenEmailfilterMheaderEntriesFieldbody2edl(o["fieldbody"], d, "fieldbody")); err != nil {
		if vv, ok := fortiAPIPatch(o["fieldbody"], "EmailfilterMheaderEntries-Fieldbody"); ok {
			if err = d.Set("fieldbody", vv); err != nil {
				return fmt.Errorf("Error reading fieldbody: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fieldbody: %v", err)
		}
	}

	if err = d.Set("fieldname", flattenEmailfilterMheaderEntriesFieldname2edl(o["fieldname"], d, "fieldname")); err != nil {
		if vv, ok := fortiAPIPatch(o["fieldname"], "EmailfilterMheaderEntries-Fieldname"); ok {
			if err = d.Set("fieldname", vv); err != nil {
				return fmt.Errorf("Error reading fieldname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fieldname: %v", err)
		}
	}

	if err = d.Set("fosid", flattenEmailfilterMheaderEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "EmailfilterMheaderEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("pattern_type", flattenEmailfilterMheaderEntriesPatternType2edl(o["pattern-type"], d, "pattern_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern-type"], "EmailfilterMheaderEntries-PatternType"); ok {
			if err = d.Set("pattern_type", vv); err != nil {
				return fmt.Errorf("Error reading pattern_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern_type: %v", err)
		}
	}

	if err = d.Set("status", flattenEmailfilterMheaderEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "EmailfilterMheaderEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterMheaderEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEmailfilterMheaderEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesFieldbody2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesFieldname2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesPatternType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterMheaderEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandEmailfilterMheaderEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fieldbody"); ok || d.HasChange("fieldbody") {
		t, err := expandEmailfilterMheaderEntriesFieldbody2edl(d, v, "fieldbody")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fieldbody"] = t
		}
	}

	if v, ok := d.GetOk("fieldname"); ok || d.HasChange("fieldname") {
		t, err := expandEmailfilterMheaderEntriesFieldname2edl(d, v, "fieldname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fieldname"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandEmailfilterMheaderEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("pattern_type"); ok || d.HasChange("pattern_type") {
		t, err := expandEmailfilterMheaderEntriesPatternType2edl(d, v, "pattern_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern-type"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandEmailfilterMheaderEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
