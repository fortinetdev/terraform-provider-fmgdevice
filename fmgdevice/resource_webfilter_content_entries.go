// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure banned word entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterContentEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterContentEntriesCreate,
		Read:   resourceWebfilterContentEntriesRead,
		Update: resourceWebfilterContentEntriesUpdate,
		Delete: resourceWebfilterContentEntriesDelete,

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
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lang": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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
		},
	}
}

func resourceWebfilterContentEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	content := d.Get("content").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["content"] = content

	obj, err := getObjectWebfilterContentEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterContentEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebfilterContentEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebfilterContentEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebfilterContentEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebfilterContentEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebfilterContentEntries resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebfilterContentEntriesRead(d, m)
}

func resourceWebfilterContentEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	content := d.Get("content").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["content"] = content

	obj, err := getObjectWebfilterContentEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterContentEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterContentEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterContentEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebfilterContentEntriesRead(d, m)
}

func resourceWebfilterContentEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	content := d.Get("content").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["content"] = content

	wsParams["adom"] = adomv

	err = c.DeleteWebfilterContentEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterContentEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterContentEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	content := d.Get("content").(string)
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
	if content == "" {
		content = importOptionChecking(m.(*FortiClient).Cfg, "content")
		if content == "" {
			return fmt.Errorf("Parameter content is missing")
		}
		if err = d.Set("content", content); err != nil {
			return fmt.Errorf("Error set params content: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["content"] = content

	o, err := c.ReadWebfilterContentEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterContentEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterContentEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterContentEntries resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterContentEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterContentEntriesLang2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterContentEntriesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterContentEntriesPatternType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterContentEntriesScore2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterContentEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterContentEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenWebfilterContentEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "WebfilterContentEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("lang", flattenWebfilterContentEntriesLang2edl(o["lang"], d, "lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["lang"], "WebfilterContentEntries-Lang"); ok {
			if err = d.Set("lang", vv); err != nil {
				return fmt.Errorf("Error reading lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lang: %v", err)
		}
	}

	if err = d.Set("name", flattenWebfilterContentEntriesName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebfilterContentEntries-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pattern_type", flattenWebfilterContentEntriesPatternType2edl(o["pattern-type"], d, "pattern_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern-type"], "WebfilterContentEntries-PatternType"); ok {
			if err = d.Set("pattern_type", vv); err != nil {
				return fmt.Errorf("Error reading pattern_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern_type: %v", err)
		}
	}

	if err = d.Set("score", flattenWebfilterContentEntriesScore2edl(o["score"], d, "score")); err != nil {
		if vv, ok := fortiAPIPatch(o["score"], "WebfilterContentEntries-Score"); ok {
			if err = d.Set("score", vv); err != nil {
				return fmt.Errorf("Error reading score: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading score: %v", err)
		}
	}

	if err = d.Set("status", flattenWebfilterContentEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WebfilterContentEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenWebfilterContentEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterContentEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterContentEntriesLang2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterContentEntriesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterContentEntriesPatternType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterContentEntriesScore2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterContentEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterContentEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandWebfilterContentEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("lang"); ok || d.HasChange("lang") {
		t, err := expandWebfilterContentEntriesLang2edl(d, v, "lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lang"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebfilterContentEntriesName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pattern_type"); ok || d.HasChange("pattern_type") {
		t, err := expandWebfilterContentEntriesPatternType2edl(d, v, "pattern_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern-type"] = t
		}
	}

	if v, ok := d.GetOk("score"); ok || d.HasChange("score") {
		t, err := expandWebfilterContentEntriesScore2edl(d, v, "score")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["score"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWebfilterContentEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
