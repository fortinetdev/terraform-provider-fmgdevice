// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device SystemSpanPort

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSpanPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSpanPortCreate,
		Read:   resourceSystemSpanPortRead,
		Update: resourceSystemSpanPortUpdate,
		Delete: resourceSystemSpanPortDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"span_dest_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"span_direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"span_source_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceSystemSpanPortCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemSpanPort(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSpanPort resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemSpanPort(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemSpanPort(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemSpanPort resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateSystemSpanPort(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemSpanPort resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceSystemSpanPortRead(d, m)
			} else {
				return fmt.Errorf("Error creating SystemSpanPort resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSpanPortRead(d, m)
}

func resourceSystemSpanPortUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemSpanPort(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpanPort resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSpanPort(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpanPort resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSpanPortRead(d, m)
}

func resourceSystemSpanPortDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteSystemSpanPort(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSpanPort resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSpanPortRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSpanPort(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemSpanPort resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSpanPort(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpanPort resource from API: %v", err)
	}
	return nil
}

func flattenSystemSpanPortId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpanPortSpanDestPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSpanPortSpanDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpanPortSpanSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSpanPortStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSpanPort(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemSpanPortId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSpanPort-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("span_dest_port", flattenSystemSpanPortSpanDestPort(o["span-dest-port"], d, "span_dest_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["span-dest-port"], "SystemSpanPort-SpanDestPort"); ok {
			if err = d.Set("span_dest_port", vv); err != nil {
				return fmt.Errorf("Error reading span_dest_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading span_dest_port: %v", err)
		}
	}

	if err = d.Set("span_direction", flattenSystemSpanPortSpanDirection(o["span-direction"], d, "span_direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["span-direction"], "SystemSpanPort-SpanDirection"); ok {
			if err = d.Set("span_direction", vv); err != nil {
				return fmt.Errorf("Error reading span_direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading span_direction: %v", err)
		}
	}

	if err = d.Set("span_source_port", flattenSystemSpanPortSpanSourcePort(o["span-source-port"], d, "span_source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["span-source-port"], "SystemSpanPort-SpanSourcePort"); ok {
			if err = d.Set("span_source_port", vv); err != nil {
				return fmt.Errorf("Error reading span_source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading span_source_port: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSpanPortStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSpanPort-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemSpanPortFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSpanPortId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpanPortSpanDestPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSpanPortSpanDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpanPortSpanSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSpanPortStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSpanPort(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSpanPortId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("span_dest_port"); ok || d.HasChange("span_dest_port") {
		t, err := expandSystemSpanPortSpanDestPort(d, v, "span_dest_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span-dest-port"] = t
		}
	}

	if v, ok := d.GetOk("span_direction"); ok || d.HasChange("span_direction") {
		t, err := expandSystemSpanPortSpanDirection(d, v, "span_direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span-direction"] = t
		}
	}

	if v, ok := d.GetOk("span_source_port"); ok || d.HasChange("span_source_port") {
		t, err := expandSystemSpanPortSpanSourcePort(d, v, "span_source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span-source-port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSpanPortStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
