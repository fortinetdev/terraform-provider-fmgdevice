// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DLP exact-data-match column types.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpExactDataMatchColumns() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpExactDataMatchColumnsCreate,
		Read:   resourceDlpExactDataMatchColumnsRead,
		Update: resourceDlpExactDataMatchColumnsUpdate,
		Delete: resourceDlpExactDataMatchColumnsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"exact_data_match": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"optional": &schema.Schema{
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

func resourceDlpExactDataMatchColumnsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	exact_data_match := d.Get("exact_data_match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["exact_data_match"] = exact_data_match

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectDlpExactDataMatchColumns(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpExactDataMatchColumns resource while getting object: %v", err)
	}

	_, err = c.CreateDlpExactDataMatchColumns(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating DlpExactDataMatchColumns resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceDlpExactDataMatchColumnsRead(d, m)
}

func resourceDlpExactDataMatchColumnsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	exact_data_match := d.Get("exact_data_match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["exact_data_match"] = exact_data_match

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectDlpExactDataMatchColumns(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpExactDataMatchColumns resource while getting object: %v", err)
	}

	_, err = c.UpdateDlpExactDataMatchColumns(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpExactDataMatchColumns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceDlpExactDataMatchColumnsRead(d, m)
}

func resourceDlpExactDataMatchColumnsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	exact_data_match := d.Get("exact_data_match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["exact_data_match"] = exact_data_match

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteDlpExactDataMatchColumns(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpExactDataMatchColumns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpExactDataMatchColumnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	exact_data_match := d.Get("exact_data_match").(string)
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
	if exact_data_match == "" {
		exact_data_match = importOptionChecking(m.(*FortiClient).Cfg, "exact_data_match")
		if exact_data_match == "" {
			return fmt.Errorf("Parameter exact_data_match is missing")
		}
		if err = d.Set("exact_data_match", exact_data_match); err != nil {
			return fmt.Errorf("Error set params exact_data_match: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["exact_data_match"] = exact_data_match

	o, err := c.ReadDlpExactDataMatchColumns(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading DlpExactDataMatchColumns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpExactDataMatchColumns(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpExactDataMatchColumns resource from API: %v", err)
	}
	return nil
}

func flattenDlpExactDataMatchColumnsIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpExactDataMatchColumnsOptional2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpExactDataMatchColumnsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectDlpExactDataMatchColumns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("index", flattenDlpExactDataMatchColumnsIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "DlpExactDataMatchColumns-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("optional", flattenDlpExactDataMatchColumnsOptional2edl(o["optional"], d, "optional")); err != nil {
		if vv, ok := fortiAPIPatch(o["optional"], "DlpExactDataMatchColumns-Optional"); ok {
			if err = d.Set("optional", vv); err != nil {
				return fmt.Errorf("Error reading optional: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading optional: %v", err)
		}
	}

	if err = d.Set("type", flattenDlpExactDataMatchColumnsType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "DlpExactDataMatchColumns-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenDlpExactDataMatchColumnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpExactDataMatchColumnsIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpExactDataMatchColumnsOptional2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpExactDataMatchColumnsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectDlpExactDataMatchColumns(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandDlpExactDataMatchColumnsIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("optional"); ok || d.HasChange("optional") {
		t, err := expandDlpExactDataMatchColumnsOptional2edl(d, v, "optional")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["optional"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandDlpExactDataMatchColumnsType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
