// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> DLP dictionary entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpDictionaryEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpDictionaryEntriesCreate,
		Read:   resourceDlpDictionaryEntriesRead,
		Update: resourceDlpDictionaryEntriesUpdate,
		Delete: resourceDlpDictionaryEntriesDelete,

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
			"dictionary": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ignore_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"repeat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDlpDictionaryEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	dictionary := d.Get("dictionary").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["dictionary"] = dictionary

	obj, err := getObjectDlpDictionaryEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpDictionaryEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDlpDictionaryEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDlpDictionaryEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DlpDictionaryEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDlpDictionaryEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DlpDictionaryEntries resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDlpDictionaryEntriesRead(d, m)
}

func resourceDlpDictionaryEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	dictionary := d.Get("dictionary").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["dictionary"] = dictionary

	obj, err := getObjectDlpDictionaryEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpDictionaryEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpDictionaryEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpDictionaryEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDlpDictionaryEntriesRead(d, m)
}

func resourceDlpDictionaryEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	dictionary := d.Get("dictionary").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["dictionary"] = dictionary

	wsParams["adom"] = adomv

	err = c.DeleteDlpDictionaryEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpDictionaryEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpDictionaryEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	dictionary := d.Get("dictionary").(string)
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
	if dictionary == "" {
		dictionary = importOptionChecking(m.(*FortiClient).Cfg, "dictionary")
		if dictionary == "" {
			return fmt.Errorf("Parameter dictionary is missing")
		}
		if err = d.Set("dictionary", dictionary); err != nil {
			return fmt.Errorf("Error set params dictionary: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["dictionary"] = dictionary

	o, err := c.ReadDlpDictionaryEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpDictionaryEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpDictionaryEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpDictionaryEntries resource from API: %v", err)
	}
	return nil
}

func flattenDlpDictionaryEntriesComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesIgnoreCase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesRepeat2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectDlpDictionaryEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenDlpDictionaryEntriesComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "DlpDictionaryEntries-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("fosid", flattenDlpDictionaryEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "DlpDictionaryEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ignore_case", flattenDlpDictionaryEntriesIgnoreCase2edl(o["ignore-case"], d, "ignore_case")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-case"], "DlpDictionaryEntries-IgnoreCase"); ok {
			if err = d.Set("ignore_case", vv); err != nil {
				return fmt.Errorf("Error reading ignore_case: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_case: %v", err)
		}
	}

	if err = d.Set("pattern", flattenDlpDictionaryEntriesPattern2edl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "DlpDictionaryEntries-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("repeat", flattenDlpDictionaryEntriesRepeat2edl(o["repeat"], d, "repeat")); err != nil {
		if vv, ok := fortiAPIPatch(o["repeat"], "DlpDictionaryEntries-Repeat"); ok {
			if err = d.Set("repeat", vv); err != nil {
				return fmt.Errorf("Error reading repeat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading repeat: %v", err)
		}
	}

	if err = d.Set("status", flattenDlpDictionaryEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "DlpDictionaryEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenDlpDictionaryEntriesType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "DlpDictionaryEntries-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenDlpDictionaryEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpDictionaryEntriesComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesIgnoreCase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesRepeat2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectDlpDictionaryEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandDlpDictionaryEntriesComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandDlpDictionaryEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ignore_case"); ok || d.HasChange("ignore_case") {
		t, err := expandDlpDictionaryEntriesIgnoreCase2edl(d, v, "ignore_case")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-case"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandDlpDictionaryEntriesPattern2edl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("repeat"); ok || d.HasChange("repeat") {
		t, err := expandDlpDictionaryEntriesRepeat2edl(d, v, "repeat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["repeat"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandDlpDictionaryEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandDlpDictionaryEntriesType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
