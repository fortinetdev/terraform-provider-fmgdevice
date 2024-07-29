// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Name list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpVenueNameValueList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpVenueNameValueListCreate,
		Read:   resourceWirelessControllerHotspot20AnqpVenueNameValueListRead,
		Update: resourceWirelessControllerHotspot20AnqpVenueNameValueListUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpVenueNameValueListDelete,

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
			"anqp_venue_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"lang": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20AnqpVenueNameValueListCreate(d *schema.ResourceData, m interface{}) error {
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
	anqp_venue_name := d.Get("anqp_venue_name").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_venue_name"] = anqp_venue_name

	obj, err := getObjectWirelessControllerHotspot20AnqpVenueNameValueList(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpVenueNameValueList resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20AnqpVenueNameValueList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpVenueNameValueList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20AnqpVenueNameValueListRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpVenueNameValueListUpdate(d *schema.ResourceData, m interface{}) error {
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
	anqp_venue_name := d.Get("anqp_venue_name").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_venue_name"] = anqp_venue_name

	obj, err := getObjectWirelessControllerHotspot20AnqpVenueNameValueList(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpVenueNameValueList resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20AnqpVenueNameValueList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpVenueNameValueList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20AnqpVenueNameValueListRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpVenueNameValueListDelete(d *schema.ResourceData, m interface{}) error {
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
	anqp_venue_name := d.Get("anqp_venue_name").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_venue_name"] = anqp_venue_name

	err = c.DeleteWirelessControllerHotspot20AnqpVenueNameValueList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpVenueNameValueList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpVenueNameValueListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	anqp_venue_name := d.Get("anqp_venue_name").(string)
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
	if anqp_venue_name == "" {
		anqp_venue_name = importOptionChecking(m.(*FortiClient).Cfg, "anqp_venue_name")
		if anqp_venue_name == "" {
			return fmt.Errorf("Parameter anqp_venue_name is missing")
		}
		if err = d.Set("anqp_venue_name", anqp_venue_name); err != nil {
			return fmt.Errorf("Error set params anqp_venue_name: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_venue_name"] = anqp_venue_name

	o, err := c.ReadWirelessControllerHotspot20AnqpVenueNameValueList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpVenueNameValueList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpVenueNameValueList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpVenueNameValueList resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpVenueNameValueListIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpVenueNameValueListLang2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpVenueNameValueListValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpVenueNameValueList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("index", flattenWirelessControllerHotspot20AnqpVenueNameValueListIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "WirelessControllerHotspot20AnqpVenueNameValueList-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("lang", flattenWirelessControllerHotspot20AnqpVenueNameValueListLang2edl(o["lang"], d, "lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["lang"], "WirelessControllerHotspot20AnqpVenueNameValueList-Lang"); ok {
			if err = d.Set("lang", vv); err != nil {
				return fmt.Errorf("Error reading lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lang: %v", err)
		}
	}

	if err = d.Set("value", flattenWirelessControllerHotspot20AnqpVenueNameValueListValue2edl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "WirelessControllerHotspot20AnqpVenueNameValueList-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpVenueNameValueListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20AnqpVenueNameValueListIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpVenueNameValueListLang2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpVenueNameValueListValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpVenueNameValueList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandWirelessControllerHotspot20AnqpVenueNameValueListIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("lang"); ok || d.HasChange("lang") {
		t, err := expandWirelessControllerHotspot20AnqpVenueNameValueListLang2edl(d, v, "lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lang"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandWirelessControllerHotspot20AnqpVenueNameValueListValue2edl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
