// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Differentiated Services Code Point (DSCP) exceptions.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20QosMapDscpExcept() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20QosMapDscpExceptCreate,
		Read:   resourceWirelessControllerHotspot20QosMapDscpExceptRead,
		Update: resourceWirelessControllerHotspot20QosMapDscpExceptUpdate,
		Delete: resourceWirelessControllerHotspot20QosMapDscpExceptDelete,

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
			"qos_map": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dscp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"up": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20QosMapDscpExceptCreate(d *schema.ResourceData, m interface{}) error {
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
	qos_map := d.Get("qos_map").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["qos_map"] = qos_map

	obj, err := getObjectWirelessControllerHotspot20QosMapDscpExcept(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20QosMapDscpExcept resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20QosMapDscpExcept(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20QosMapDscpExcept resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20QosMapDscpExceptRead(d, m)
}

func resourceWirelessControllerHotspot20QosMapDscpExceptUpdate(d *schema.ResourceData, m interface{}) error {
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
	qos_map := d.Get("qos_map").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["qos_map"] = qos_map

	obj, err := getObjectWirelessControllerHotspot20QosMapDscpExcept(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20QosMapDscpExcept resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20QosMapDscpExcept(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20QosMapDscpExcept resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20QosMapDscpExceptRead(d, m)
}

func resourceWirelessControllerHotspot20QosMapDscpExceptDelete(d *schema.ResourceData, m interface{}) error {
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
	qos_map := d.Get("qos_map").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["qos_map"] = qos_map

	err = c.DeleteWirelessControllerHotspot20QosMapDscpExcept(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20QosMapDscpExcept resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20QosMapDscpExceptRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	qos_map := d.Get("qos_map").(string)
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
	if qos_map == "" {
		qos_map = importOptionChecking(m.(*FortiClient).Cfg, "qos_map")
		if qos_map == "" {
			return fmt.Errorf("Parameter qos_map is missing")
		}
		if err = d.Set("qos_map", qos_map); err != nil {
			return fmt.Errorf("Error set params qos_map: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["qos_map"] = qos_map

	o, err := c.ReadWirelessControllerHotspot20QosMapDscpExcept(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20QosMapDscpExcept resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20QosMapDscpExcept(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20QosMapDscpExcept resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20QosMapDscpExceptDscp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpExceptIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpExceptUp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20QosMapDscpExcept(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dscp", flattenWirelessControllerHotspot20QosMapDscpExceptDscp2edl(o["dscp"], d, "dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp"], "WirelessControllerHotspot20QosMapDscpExcept-Dscp"); ok {
			if err = d.Set("dscp", vv); err != nil {
				return fmt.Errorf("Error reading dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp: %v", err)
		}
	}

	if err = d.Set("index", flattenWirelessControllerHotspot20QosMapDscpExceptIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "WirelessControllerHotspot20QosMapDscpExcept-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("up", flattenWirelessControllerHotspot20QosMapDscpExceptUp2edl(o["up"], d, "up")); err != nil {
		if vv, ok := fortiAPIPatch(o["up"], "WirelessControllerHotspot20QosMapDscpExcept-Up"); ok {
			if err = d.Set("up", vv); err != nil {
				return fmt.Errorf("Error reading up: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading up: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20QosMapDscpExceptFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20QosMapDscpExceptDscp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpExceptIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpExceptUp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20QosMapDscpExcept(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dscp"); ok || d.HasChange("dscp") {
		t, err := expandWirelessControllerHotspot20QosMapDscpExceptDscp2edl(d, v, "dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp"] = t
		}
	}

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandWirelessControllerHotspot20QosMapDscpExceptIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("up"); ok || d.HasChange("up") {
		t, err := expandWirelessControllerHotspot20QosMapDscpExceptUp2edl(d, v, "up")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["up"] = t
		}
	}

	return &obj, nil
}
