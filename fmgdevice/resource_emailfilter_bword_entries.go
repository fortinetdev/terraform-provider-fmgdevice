// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Spam filter banned word.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterBwordEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterBwordEntriesCreate,
		Read:   resourceEmailfilterBwordEntriesRead,
		Update: resourceEmailfilterBwordEntriesUpdate,
		Delete: resourceEmailfilterBwordEntriesDelete,

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
			"bword": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
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
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pattern_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"score": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"where": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceEmailfilterBwordEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	bword := d.Get("bword").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["bword"] = bword

	obj, err := getObjectEmailfilterBwordEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating EmailfilterBwordEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadEmailfilterBwordEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateEmailfilterBwordEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating EmailfilterBwordEntries resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateEmailfilterBwordEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating EmailfilterBwordEntries resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceEmailfilterBwordEntriesRead(d, m)
			} else {
				return fmt.Errorf("Error creating EmailfilterBwordEntries resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterBwordEntriesRead(d, m)
}

func resourceEmailfilterBwordEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	bword := d.Get("bword").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["bword"] = bword

	obj, err := getObjectEmailfilterBwordEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterBwordEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEmailfilterBwordEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterBwordEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterBwordEntriesRead(d, m)
}

func resourceEmailfilterBwordEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	bword := d.Get("bword").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["bword"] = bword

	wsParams["adom"] = adomv

	err = c.DeleteEmailfilterBwordEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterBwordEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterBwordEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	bword := d.Get("bword").(string)
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
	if bword == "" {
		bword = importOptionChecking(m.(*FortiClient).Cfg, "bword")
		if bword == "" {
			return fmt.Errorf("Parameter bword is missing")
		}
		if err = d.Set("bword", bword); err != nil {
			return fmt.Errorf("Error set params bword: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["bword"] = bword

	o, err := c.ReadEmailfilterBwordEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EmailfilterBwordEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterBwordEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterBwordEntries resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterBwordEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesLanguage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesPatternType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesScore2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesWhere2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEmailfilterBwordEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenEmailfilterBwordEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "EmailfilterBwordEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fosid", flattenEmailfilterBwordEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "EmailfilterBwordEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("language", flattenEmailfilterBwordEntriesLanguage2edl(o["language"], d, "language")); err != nil {
		if vv, ok := fortiAPIPatch(o["language"], "EmailfilterBwordEntries-Language"); ok {
			if err = d.Set("language", vv); err != nil {
				return fmt.Errorf("Error reading language: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading language: %v", err)
		}
	}

	if err = d.Set("pattern", flattenEmailfilterBwordEntriesPattern2edl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "EmailfilterBwordEntries-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("pattern_type", flattenEmailfilterBwordEntriesPatternType2edl(o["pattern-type"], d, "pattern_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern-type"], "EmailfilterBwordEntries-PatternType"); ok {
			if err = d.Set("pattern_type", vv); err != nil {
				return fmt.Errorf("Error reading pattern_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern_type: %v", err)
		}
	}

	if err = d.Set("score", flattenEmailfilterBwordEntriesScore2edl(o["score"], d, "score")); err != nil {
		if vv, ok := fortiAPIPatch(o["score"], "EmailfilterBwordEntries-Score"); ok {
			if err = d.Set("score", vv); err != nil {
				return fmt.Errorf("Error reading score: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading score: %v", err)
		}
	}

	if err = d.Set("status", flattenEmailfilterBwordEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "EmailfilterBwordEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("where", flattenEmailfilterBwordEntriesWhere2edl(o["where"], d, "where")); err != nil {
		if vv, ok := fortiAPIPatch(o["where"], "EmailfilterBwordEntries-Where"); ok {
			if err = d.Set("where", vv); err != nil {
				return fmt.Errorf("Error reading where: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading where: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterBwordEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEmailfilterBwordEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesLanguage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesPatternType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesScore2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesWhere2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterBwordEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandEmailfilterBwordEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandEmailfilterBwordEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("language"); ok || d.HasChange("language") {
		t, err := expandEmailfilterBwordEntriesLanguage2edl(d, v, "language")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["language"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandEmailfilterBwordEntriesPattern2edl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("pattern_type"); ok || d.HasChange("pattern_type") {
		t, err := expandEmailfilterBwordEntriesPatternType2edl(d, v, "pattern_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern-type"] = t
		}
	}

	if v, ok := d.GetOk("score"); ok || d.HasChange("score") {
		t, err := expandEmailfilterBwordEntriesScore2edl(d, v, "score")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["score"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandEmailfilterBwordEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("where"); ok || d.HasChange("where") {
		t, err := expandEmailfilterBwordEntriesWhere2edl(d, v, "where")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["where"] = t
		}
	}

	return &obj, nil
}
