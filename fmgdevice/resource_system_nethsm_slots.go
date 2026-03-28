// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device SystemNethsmSlots

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNethsmSlots() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNethsmSlotsCreate,
		Read:   resourceSystemNethsmSlotsRead,
		Update: resourceSystemNethsmSlotsUpdate,
		Delete: resourceSystemNethsmSlotsDelete,

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
			"for_ha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
		},
	}
}

func resourceSystemNethsmSlotsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemNethsmSlots(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemNethsmSlots resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemNethsmSlots(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemNethsmSlots(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemNethsmSlots resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemNethsmSlots(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemNethsmSlots resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemNethsmSlotsRead(d, m)
}

func resourceSystemNethsmSlotsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemNethsmSlots(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNethsmSlots resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemNethsmSlots(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemNethsmSlots resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemNethsmSlotsRead(d, m)
}

func resourceSystemNethsmSlotsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemNethsmSlots(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNethsmSlots resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNethsmSlotsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemNethsmSlots(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemNethsmSlots resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNethsmSlots(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNethsmSlots resource from API: %v", err)
	}
	return nil
}

func flattenSystemNethsmSlotsForHa2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmSlotsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmSlotsName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNethsmSlots(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("for_ha", flattenSystemNethsmSlotsForHa2edl(o["for-ha"], d, "for_ha")); err != nil {
		if vv, ok := fortiAPIPatch(o["for-ha"], "SystemNethsmSlots-ForHa"); ok {
			if err = d.Set("for_ha", vv); err != nil {
				return fmt.Errorf("Error reading for_ha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading for_ha: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemNethsmSlotsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemNethsmSlots-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemNethsmSlotsName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemNethsmSlots-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemNethsmSlotsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNethsmSlotsForHa2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmSlotsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmSlotsName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmSlotsPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemNethsmSlots(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("for_ha"); ok || d.HasChange("for_ha") {
		t, err := expandSystemNethsmSlotsForHa2edl(d, v, "for_ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["for-ha"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemNethsmSlotsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemNethsmSlotsName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemNethsmSlotsPassword2edl(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	return &obj, nil
}
