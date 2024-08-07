// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure a list of hashes to be exempt from AV scanning.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAntivirusExemptList() *schema.Resource {
	return &schema.Resource{
		Create: resourceAntivirusExemptListCreate,
		Read:   resourceAntivirusExemptListRead,
		Update: resourceAntivirusExemptListUpdate,
		Delete: resourceAntivirusExemptListDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hash_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAntivirusExemptListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectAntivirusExemptList(d)
	if err != nil {
		return fmt.Errorf("Error creating AntivirusExemptList resource while getting object: %v", err)
	}

	_, err = c.CreateAntivirusExemptList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating AntivirusExemptList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceAntivirusExemptListRead(d, m)
}

func resourceAntivirusExemptListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectAntivirusExemptList(d)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusExemptList resource while getting object: %v", err)
	}

	_, err = c.UpdateAntivirusExemptList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusExemptList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceAntivirusExemptListRead(d, m)
}

func resourceAntivirusExemptListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteAntivirusExemptList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting AntivirusExemptList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusExemptListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadAntivirusExemptList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusExemptList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusExemptList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusExemptList resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusExemptListComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusExemptListHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusExemptListHashType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusExemptListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusExemptListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAntivirusExemptList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenAntivirusExemptListComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "AntivirusExemptList-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("hash", flattenAntivirusExemptListHash(o["hash"], d, "hash")); err != nil {
		if vv, ok := fortiAPIPatch(o["hash"], "AntivirusExemptList-Hash"); ok {
			if err = d.Set("hash", vv); err != nil {
				return fmt.Errorf("Error reading hash: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hash: %v", err)
		}
	}

	if err = d.Set("hash_type", flattenAntivirusExemptListHashType(o["hash-type"], d, "hash_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["hash-type"], "AntivirusExemptList-HashType"); ok {
			if err = d.Set("hash_type", vv); err != nil {
				return fmt.Errorf("Error reading hash_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hash_type: %v", err)
		}
	}

	if err = d.Set("name", flattenAntivirusExemptListName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "AntivirusExemptList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenAntivirusExemptListStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "AntivirusExemptList-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenAntivirusExemptListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAntivirusExemptListComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusExemptListHash(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusExemptListHashType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusExemptListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusExemptListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusExemptList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandAntivirusExemptListComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("hash"); ok || d.HasChange("hash") {
		t, err := expandAntivirusExemptListHash(d, v, "hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash"] = t
		}
	}

	if v, ok := d.GetOk("hash_type"); ok || d.HasChange("hash_type") {
		t, err := expandAntivirusExemptListHashType(d, v, "hash_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandAntivirusExemptListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandAntivirusExemptListStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
