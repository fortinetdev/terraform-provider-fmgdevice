// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device SystemNethsmPartitions

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNethsmPartitions() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNethsmPartitionsCreate,
		Read:   resourceSystemNethsmPartitionsRead,
		Update: resourceSystemNethsmPartitionsUpdate,
		Delete: resourceSystemNethsmPartitionsDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"pkcs11_pin": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"slot_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemNethsmPartitionsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemNethsmPartitions(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemNethsmPartitions resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemNethsmPartitions(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemNethsmPartitions(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemNethsmPartitions resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemNethsmPartitions(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemNethsmPartitions resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemNethsmPartitionsRead(d, m)
}

func resourceSystemNethsmPartitionsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemNethsmPartitions(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNethsmPartitions resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemNethsmPartitions(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemNethsmPartitions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemNethsmPartitionsRead(d, m)
}

func resourceSystemNethsmPartitionsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemNethsmPartitions(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNethsmPartitions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNethsmPartitionsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemNethsmPartitions(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemNethsmPartitions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNethsmPartitions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNethsmPartitions resource from API: %v", err)
	}
	return nil
}

func flattenSystemNethsmPartitionsName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmPartitionsSlotId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNethsmPartitions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemNethsmPartitionsName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemNethsmPartitions-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("slot_id", flattenSystemNethsmPartitionsSlotId2edl(o["slot-id"], d, "slot_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["slot-id"], "SystemNethsmPartitions-SlotId"); ok {
			if err = d.Set("slot_id", vv); err != nil {
				return fmt.Errorf("Error reading slot_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slot_id: %v", err)
		}
	}

	return nil
}

func flattenSystemNethsmPartitionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNethsmPartitionsName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmPartitionsPkcs11Pin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemNethsmPartitionsSlotId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNethsmPartitions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemNethsmPartitionsName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pkcs11_pin"); ok || d.HasChange("pkcs11_pin") {
		t, err := expandSystemNethsmPartitionsPkcs11Pin2edl(d, v, "pkcs11_pin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pkcs11-pin"] = t
		}
	}

	if v, ok := d.GetOk("slot_id"); ok || d.HasChange("slot_id") {
		t, err := expandSystemNethsmPartitionsSlotId2edl(d, v, "slot_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slot-id"] = t
		}
	}

	return &obj, nil
}
