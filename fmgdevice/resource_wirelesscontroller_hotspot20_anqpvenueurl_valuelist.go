// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: URL list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpVenueUrlValueList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpVenueUrlValueListCreate,
		Read:   resourceWirelessControllerHotspot20AnqpVenueUrlValueListRead,
		Update: resourceWirelessControllerHotspot20AnqpVenueUrlValueListUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpVenueUrlValueListDelete,

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
			"anqp_venue_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20AnqpVenueUrlValueListCreate(d *schema.ResourceData, m interface{}) error {
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
	anqp_venue_url := d.Get("anqp_venue_url").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_venue_url"] = anqp_venue_url

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20AnqpVenueUrlValueList(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpVenueUrlValueList resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20AnqpVenueUrlValueList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpVenueUrlValueList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20AnqpVenueUrlValueListRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpVenueUrlValueListUpdate(d *schema.ResourceData, m interface{}) error {
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
	anqp_venue_url := d.Get("anqp_venue_url").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_venue_url"] = anqp_venue_url

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20AnqpVenueUrlValueList(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpVenueUrlValueList resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20AnqpVenueUrlValueList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpVenueUrlValueList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20AnqpVenueUrlValueListRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpVenueUrlValueListDelete(d *schema.ResourceData, m interface{}) error {
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
	anqp_venue_url := d.Get("anqp_venue_url").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_venue_url"] = anqp_venue_url

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerHotspot20AnqpVenueUrlValueList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpVenueUrlValueList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpVenueUrlValueListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	anqp_venue_url := d.Get("anqp_venue_url").(string)
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
	if anqp_venue_url == "" {
		anqp_venue_url = importOptionChecking(m.(*FortiClient).Cfg, "anqp_venue_url")
		if anqp_venue_url == "" {
			return fmt.Errorf("Parameter anqp_venue_url is missing")
		}
		if err = d.Set("anqp_venue_url", anqp_venue_url); err != nil {
			return fmt.Errorf("Error set params anqp_venue_url: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_venue_url"] = anqp_venue_url

	o, err := c.ReadWirelessControllerHotspot20AnqpVenueUrlValueList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpVenueUrlValueList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpVenueUrlValueList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpVenueUrlValueList resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpVenueUrlValueListIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpVenueUrlValueListNumber2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpVenueUrlValueListValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpVenueUrlValueList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("index", flattenWirelessControllerHotspot20AnqpVenueUrlValueListIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "WirelessControllerHotspot20AnqpVenueUrlValueList-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("number", flattenWirelessControllerHotspot20AnqpVenueUrlValueListNumber2edl(o["number"], d, "number")); err != nil {
		if vv, ok := fortiAPIPatch(o["number"], "WirelessControllerHotspot20AnqpVenueUrlValueList-Number"); ok {
			if err = d.Set("number", vv); err != nil {
				return fmt.Errorf("Error reading number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading number: %v", err)
		}
	}

	if err = d.Set("value", flattenWirelessControllerHotspot20AnqpVenueUrlValueListValue2edl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "WirelessControllerHotspot20AnqpVenueUrlValueList-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpVenueUrlValueListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20AnqpVenueUrlValueListIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpVenueUrlValueListNumber2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpVenueUrlValueListValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpVenueUrlValueList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandWirelessControllerHotspot20AnqpVenueUrlValueListIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("number"); ok || d.HasChange("number") {
		t, err := expandWirelessControllerHotspot20AnqpVenueUrlValueListNumber2edl(d, v, "number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["number"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandWirelessControllerHotspot20AnqpVenueUrlValueListValue2edl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
