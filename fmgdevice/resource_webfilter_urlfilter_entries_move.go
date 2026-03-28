// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> URL filter entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterUrlfilterEntriesMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterUrlfilterEntriesMoveUpdate,
		Read:   resourceWebfilterUrlfilterEntriesMoveRead,
		Update: resourceWebfilterUrlfilterEntriesMoveUpdate,
		Delete: resourceWebfilterUrlfilterEntriesMoveDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"state_pos": &schema.Schema{
				Type:     schema.TypeString,
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
			"urlfilter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceWebfilterUrlfilterEntriesMoveUpdate(d *schema.ResourceData, m interface{}) error {
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
	urlfilter := d.Get("urlfilter").(string)
	entries := d.Get("entries").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["urlfilter"] = urlfilter
	paradict["entries"] = entries

	target := d.Get("target").(string)
	obj, err := getObjectWebfilterUrlfilterEntriesMove(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterUrlfilterEntriesMove resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterUrlfilterEntriesMove(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterUrlfilterEntriesMove resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WebfilterUrlfilterEntriesMove" + "_" + device_name + "_" + device_vdom + "_" + urlfilter + "_" + entries + "_" + target)

	return resourceWebfilterUrlfilterEntriesMoveRead(d, m)
}

func resourceWebfilterUrlfilterEntriesMoveDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceWebfilterUrlfilterEntriesMoveRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	sid := d.Get("entries").(string)
	did := d.Get("target").(string)
	action := d.Get("option").(string)

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	urlfilter := d.Get("urlfilter").(string)
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
	if urlfilter == "" {
		urlfilter = importOptionChecking(m.(*FortiClient).Cfg, "urlfilter")
		if urlfilter == "" {
			return fmt.Errorf("Parameter urlfilter is missing")
		}
		if err = d.Set("urlfilter", urlfilter); err != nil {
			return fmt.Errorf("Error set params urlfilter: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["urlfilter"] = urlfilter

	o, err := c.ReadWebfilterUrlfilterEntriesMove(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterUrlfilterEntriesMove resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if o != nil {
		i := 1
		now_sid := -1
		now_did := -1

		for _, z := range o {
			if v, ok := z.(map[string]interface{}); ok {
				if _, ok := v["id"]; !ok {
					return fmt.Errorf("Error reading WebfilterUrlfilterEntriesMove resource: id doesn't exist.")
				}

				vid := fmt.Sprintf("%v", v["id"])

				if vid == sid {
					now_sid = i
				}

				if vid == did {
					now_did = i
				}
			} else {
				return fmt.Errorf("Error reading WebfilterUrlfilterEntriesMove resource: no valid map string.")
			}

			i += 1
		}

		state_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_pos = "id(" + sid + ") and target(" + did + ") were deleted"
			} else if now_sid == -1 {
				state_pos = "id(" + sid + ") was deleted"
			} else if now_did == -1 {
				state_pos = "target(" + did + ") was deleted"
			}
		} else {
			bconsistent := true
			if action == "before" {
				if now_sid != now_did-1 {
					bconsistent = false
				}
			}

			if action == "after" {
				if now_sid != now_did+1 {
					bconsistent = false
				}
			}

			if bconsistent == false {
				relative_pos := now_sid - now_did

				if relative_pos > 0 {
					state_pos = "id(" + sid + ") is " + strconv.Itoa(relative_pos) + " behind target(" + did + ")"
				} else {
					state_pos = "id(" + sid + ") is " + strconv.Itoa(-relative_pos) + " ahead of target(" + did + ")"
				}
			}
		}

		d.Set("state_pos", state_pos)
	}

	return nil
}

func flattenWebfilterUrlfilterEntriesMoveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterUrlfilterEntriesMoveTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesMoveOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterUrlfilterEntriesMove(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandWebfilterUrlfilterEntriesMoveTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandWebfilterUrlfilterEntriesMoveOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
