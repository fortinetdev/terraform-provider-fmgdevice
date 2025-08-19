// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Differentiated Services Code Point (DSCP) ranges.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20QosMapDscpRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20QosMapDscpRangeCreate,
		Read:   resourceWirelessControllerHotspot20QosMapDscpRangeRead,
		Update: resourceWirelessControllerHotspot20QosMapDscpRangeUpdate,
		Delete: resourceWirelessControllerHotspot20QosMapDscpRangeDelete,

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
			"high": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"low": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"up": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20QosMapDscpRangeCreate(d *schema.ResourceData, m interface{}) error {
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
	qos_map := d.Get("qos_map").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["qos_map"] = qos_map

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20QosMapDscpRange(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20QosMapDscpRange resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20QosMapDscpRange(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20QosMapDscpRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20QosMapDscpRangeRead(d, m)
}

func resourceWirelessControllerHotspot20QosMapDscpRangeUpdate(d *schema.ResourceData, m interface{}) error {
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
	qos_map := d.Get("qos_map").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["qos_map"] = qos_map

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20QosMapDscpRange(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20QosMapDscpRange resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20QosMapDscpRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20QosMapDscpRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20QosMapDscpRangeRead(d, m)
}

func resourceWirelessControllerHotspot20QosMapDscpRangeDelete(d *schema.ResourceData, m interface{}) error {
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
	qos_map := d.Get("qos_map").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["qos_map"] = qos_map

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerHotspot20QosMapDscpRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20QosMapDscpRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20QosMapDscpRangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerHotspot20QosMapDscpRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20QosMapDscpRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20QosMapDscpRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20QosMapDscpRange resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20QosMapDscpRangeHigh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpRangeIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpRangeLow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpRangeUp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20QosMapDscpRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("high", flattenWirelessControllerHotspot20QosMapDscpRangeHigh2edl(o["high"], d, "high")); err != nil {
		if vv, ok := fortiAPIPatch(o["high"], "WirelessControllerHotspot20QosMapDscpRange-High"); ok {
			if err = d.Set("high", vv); err != nil {
				return fmt.Errorf("Error reading high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading high: %v", err)
		}
	}

	if err = d.Set("index", flattenWirelessControllerHotspot20QosMapDscpRangeIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "WirelessControllerHotspot20QosMapDscpRange-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("low", flattenWirelessControllerHotspot20QosMapDscpRangeLow2edl(o["low"], d, "low")); err != nil {
		if vv, ok := fortiAPIPatch(o["low"], "WirelessControllerHotspot20QosMapDscpRange-Low"); ok {
			if err = d.Set("low", vv); err != nil {
				return fmt.Errorf("Error reading low: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading low: %v", err)
		}
	}

	if err = d.Set("up", flattenWirelessControllerHotspot20QosMapDscpRangeUp2edl(o["up"], d, "up")); err != nil {
		if vv, ok := fortiAPIPatch(o["up"], "WirelessControllerHotspot20QosMapDscpRange-Up"); ok {
			if err = d.Set("up", vv); err != nil {
				return fmt.Errorf("Error reading up: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading up: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20QosMapDscpRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20QosMapDscpRangeHigh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpRangeIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpRangeLow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpRangeUp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20QosMapDscpRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("high"); ok || d.HasChange("high") {
		t, err := expandWirelessControllerHotspot20QosMapDscpRangeHigh2edl(d, v, "high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["high"] = t
		}
	}

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandWirelessControllerHotspot20QosMapDscpRangeIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("low"); ok || d.HasChange("low") {
		t, err := expandWirelessControllerHotspot20QosMapDscpRangeLow2edl(d, v, "low")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["low"] = t
		}
	}

	if v, ok := d.GetOk("up"); ok || d.HasChange("up") {
		t, err := expandWirelessControllerHotspot20QosMapDscpRangeUp2edl(d, v, "up")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["up"] = t
		}
	}

	return &obj, nil
}
