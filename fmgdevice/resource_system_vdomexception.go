// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Global configuration objects that can be configured independently across different ha peers for all VDOMs or for the defined VDOM scope.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVdomException() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomExceptionCreate,
		Read:   resourceSystemVdomExceptionRead,
		Update: resourceSystemVdomExceptionUpdate,
		Delete: resourceSystemVdomExceptionDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemVdomExceptionCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemVdomException(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomException resource while getting object: %v", err)
	}

	_, err = c.CreateSystemVdomException(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomException resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemVdomExceptionRead(d, m)
}

func resourceSystemVdomExceptionUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemVdomException(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomException resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVdomException(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomException resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemVdomExceptionRead(d, m)
}

func resourceSystemVdomExceptionDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemVdomException(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomException resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomExceptionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVdomException(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomException resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomException(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomException resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomExceptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomExceptionObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomExceptionOid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomExceptionScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomExceptionVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemVdomException(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemVdomExceptionId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemVdomException-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("object", flattenSystemVdomExceptionObject(o["object"], d, "object")); err != nil {
		if vv, ok := fortiAPIPatch(o["object"], "SystemVdomException-Object"); ok {
			if err = d.Set("object", vv); err != nil {
				return fmt.Errorf("Error reading object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object: %v", err)
		}
	}

	if err = d.Set("oid", flattenSystemVdomExceptionOid(o["oid"], d, "oid")); err != nil {
		if vv, ok := fortiAPIPatch(o["oid"], "SystemVdomException-Oid"); ok {
			if err = d.Set("oid", vv); err != nil {
				return fmt.Errorf("Error reading oid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oid: %v", err)
		}
	}

	if err = d.Set("scope", flattenSystemVdomExceptionScope(o["scope"], d, "scope")); err != nil {
		if vv, ok := fortiAPIPatch(o["scope"], "SystemVdomException-Scope"); ok {
			if err = d.Set("scope", vv); err != nil {
				return fmt.Errorf("Error reading scope: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scope: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemVdomExceptionVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemVdomException-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomExceptionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVdomExceptionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomExceptionObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomExceptionOid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomExceptionScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomExceptionVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemVdomException(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemVdomExceptionId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("object"); ok || d.HasChange("object") {
		t, err := expandSystemVdomExceptionObject(d, v, "object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object"] = t
		}
	}

	if v, ok := d.GetOk("oid"); ok || d.HasChange("oid") {
		t, err := expandSystemVdomExceptionOid(d, v, "oid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oid"] = t
		}
	}

	if v, ok := d.GetOk("scope"); ok || d.HasChange("scope") {
		t, err := expandSystemVdomExceptionScope(d, v, "scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemVdomExceptionVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
