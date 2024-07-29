// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPS group parameters.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpsDecoderParameter() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsDecoderParameterUpdate,
		Read:   resourceIpsDecoderParameterRead,
		Update: resourceIpsDecoderParameterUpdate,
		Delete: resourceIpsDecoderParameterDelete,

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
			"decoder": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceIpsDecoderParameterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	decoder := d.Get("decoder").(string)
	paradict["device"] = device_name
	paradict["decoder"] = decoder

	obj, err := getObjectIpsDecoderParameter(d)
	if err != nil {
		return fmt.Errorf("Error updating IpsDecoderParameter resource while getting object: %v", err)
	}

	_, err = c.UpdateIpsDecoderParameter(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating IpsDecoderParameter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("IpsDecoderParameter")

	return resourceIpsDecoderParameterRead(d, m)
}

func resourceIpsDecoderParameterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	decoder := d.Get("decoder").(string)
	paradict["device"] = device_name
	paradict["decoder"] = decoder

	err = c.DeleteIpsDecoderParameter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting IpsDecoderParameter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsDecoderParameterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	decoder := d.Get("decoder").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if decoder == "" {
		decoder = importOptionChecking(m.(*FortiClient).Cfg, "decoder")
		if decoder == "" {
			return fmt.Errorf("Parameter decoder is missing")
		}
		if err = d.Set("decoder", decoder); err != nil {
			return fmt.Errorf("Error set params decoder: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["decoder"] = decoder

	o, err := c.ReadIpsDecoderParameter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading IpsDecoderParameter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsDecoderParameter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IpsDecoderParameter resource from API: %v", err)
	}
	return nil
}

func flattenIpsDecoderParameterName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsDecoderParameterValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIpsDecoderParameter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenIpsDecoderParameterName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "IpsDecoderParameter-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenIpsDecoderParameterValue2edl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "IpsDecoderParameter-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenIpsDecoderParameterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIpsDecoderParameterName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsDecoderParameterValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIpsDecoderParameter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandIpsDecoderParameterName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandIpsDecoderParameterValue2edl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
