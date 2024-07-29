// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Organization identifier list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumOiList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpRoamingConsortiumOiListCreate,
		Read:   resourceWirelessControllerHotspot20AnqpRoamingConsortiumOiListRead,
		Update: resourceWirelessControllerHotspot20AnqpRoamingConsortiumOiListUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpRoamingConsortiumOiListDelete,

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
			"anqp_roaming_consortium": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"oi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumOiListCreate(d *schema.ResourceData, m interface{}) error {
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
	anqp_roaming_consortium := d.Get("anqp_roaming_consortium").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_roaming_consortium"] = anqp_roaming_consortium

	obj, err := getObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpRoamingConsortiumOiList resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20AnqpRoamingConsortiumOiList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpRoamingConsortiumOiList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20AnqpRoamingConsortiumOiListRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumOiListUpdate(d *schema.ResourceData, m interface{}) error {
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
	anqp_roaming_consortium := d.Get("anqp_roaming_consortium").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_roaming_consortium"] = anqp_roaming_consortium

	obj, err := getObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpRoamingConsortiumOiList resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20AnqpRoamingConsortiumOiList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpRoamingConsortiumOiList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20AnqpRoamingConsortiumOiListRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumOiListDelete(d *schema.ResourceData, m interface{}) error {
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
	anqp_roaming_consortium := d.Get("anqp_roaming_consortium").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_roaming_consortium"] = anqp_roaming_consortium

	err = c.DeleteWirelessControllerHotspot20AnqpRoamingConsortiumOiList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpRoamingConsortiumOiList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumOiListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	anqp_roaming_consortium := d.Get("anqp_roaming_consortium").(string)
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
	if anqp_roaming_consortium == "" {
		anqp_roaming_consortium = importOptionChecking(m.(*FortiClient).Cfg, "anqp_roaming_consortium")
		if anqp_roaming_consortium == "" {
			return fmt.Errorf("Parameter anqp_roaming_consortium is missing")
		}
		if err = d.Set("anqp_roaming_consortium", anqp_roaming_consortium); err != nil {
			return fmt.Errorf("Error set params anqp_roaming_consortium: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_roaming_consortium"] = anqp_roaming_consortium

	o, err := c.ReadWirelessControllerHotspot20AnqpRoamingConsortiumOiList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpRoamingConsortiumOiList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpRoamingConsortiumOiList resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerHotspot20AnqpRoamingConsortiumOiList-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("index", flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "WirelessControllerHotspot20AnqpRoamingConsortiumOiList-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("oi", flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi2edl(o["oi"], d, "oi")); err != nil {
		if vv, ok := fortiAPIPatch(o["oi"], "WirelessControllerHotspot20AnqpRoamingConsortiumOiList-Oi"); ok {
			if err = d.Set("oi", vv); err != nil {
				return fmt.Errorf("Error reading oi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oi: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("oi"); ok || d.HasChange("oi") {
		t, err := expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi2edl(d, v, "oi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oi"] = t
		}
	}

	return &obj, nil
}
