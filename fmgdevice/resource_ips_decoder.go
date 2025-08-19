// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPS decoder.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpsDecoder() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsDecoderCreate,
		Read:   resourceIpsDecoderRead,
		Update: resourceIpsDecoderUpdate,
		Delete: resourceIpsDecoderDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"parameter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceIpsDecoderCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectIpsDecoder(d)
	if err != nil {
		return fmt.Errorf("Error creating IpsDecoder resource while getting object: %v", err)
	}

	_, err = c.CreateIpsDecoder(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating IpsDecoder resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceIpsDecoderRead(d, m)
}

func resourceIpsDecoderUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectIpsDecoder(d)
	if err != nil {
		return fmt.Errorf("Error updating IpsDecoder resource while getting object: %v", err)
	}

	_, err = c.UpdateIpsDecoder(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IpsDecoder resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceIpsDecoderRead(d, m)
}

func resourceIpsDecoderDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteIpsDecoder(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IpsDecoder resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsDecoderRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadIpsDecoder(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading IpsDecoder resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsDecoder(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IpsDecoder resource from API: %v", err)
	}
	return nil
}

func flattenIpsDecoderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsDecoderParameter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenIpsDecoderParameterName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "IpsDecoder-Parameter-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenIpsDecoderParameterValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "IpsDecoder-Parameter-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenIpsDecoderParameterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsDecoderParameterValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIpsDecoder(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenIpsDecoderName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "IpsDecoder-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("parameter", flattenIpsDecoderParameter(o["parameter"], d, "parameter")); err != nil {
			if vv, ok := fortiAPIPatch(o["parameter"], "IpsDecoder-Parameter"); ok {
				if err = d.Set("parameter", vv); err != nil {
					return fmt.Errorf("Error reading parameter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading parameter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("parameter"); ok {
			if err = d.Set("parameter", flattenIpsDecoderParameter(o["parameter"], d, "parameter")); err != nil {
				if vv, ok := fortiAPIPatch(o["parameter"], "IpsDecoder-Parameter"); ok {
					if err = d.Set("parameter", vv); err != nil {
						return fmt.Errorf("Error reading parameter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading parameter: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenIpsDecoderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIpsDecoderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsDecoderParameter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandIpsDecoderParameterName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandIpsDecoderParameterValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandIpsDecoderParameterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsDecoderParameterValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIpsDecoder(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandIpsDecoderName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("parameter"); ok || d.HasChange("parameter") {
		t, err := expandIpsDecoderParameter(d, v, "parameter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter"] = t
		}
	}

	return &obj, nil
}
