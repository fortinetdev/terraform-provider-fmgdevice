// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Mobile Country Code and Mobile Network Code configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20Anqp3GppCellularMccMncList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20Anqp3GppCellularMccMncListCreate,
		Read:   resourceWirelessControllerHotspot20Anqp3GppCellularMccMncListRead,
		Update: resourceWirelessControllerHotspot20Anqp3GppCellularMccMncListUpdate,
		Delete: resourceWirelessControllerHotspot20Anqp3GppCellularMccMncListDelete,

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
			"anqp_3gpp_cellular": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"mcc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mnc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20Anqp3GppCellularMccMncListCreate(d *schema.ResourceData, m interface{}) error {
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
	anqp_3gpp_cellular := d.Get("anqp_3gpp_cellular").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_3gpp_cellular"] = anqp_3gpp_cellular

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20Anqp3GppCellularMccMncList resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20Anqp3GppCellularMccMncList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20Anqp3GppCellularMccMncList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWirelessControllerHotspot20Anqp3GppCellularMccMncListRead(d, m)
}

func resourceWirelessControllerHotspot20Anqp3GppCellularMccMncListUpdate(d *schema.ResourceData, m interface{}) error {
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
	anqp_3gpp_cellular := d.Get("anqp_3gpp_cellular").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_3gpp_cellular"] = anqp_3gpp_cellular

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20Anqp3GppCellularMccMncList resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20Anqp3GppCellularMccMncList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20Anqp3GppCellularMccMncList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWirelessControllerHotspot20Anqp3GppCellularMccMncListRead(d, m)
}

func resourceWirelessControllerHotspot20Anqp3GppCellularMccMncListDelete(d *schema.ResourceData, m interface{}) error {
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
	anqp_3gpp_cellular := d.Get("anqp_3gpp_cellular").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_3gpp_cellular"] = anqp_3gpp_cellular

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerHotspot20Anqp3GppCellularMccMncList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20Anqp3GppCellularMccMncList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20Anqp3GppCellularMccMncListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	anqp_3gpp_cellular := d.Get("anqp_3gpp_cellular").(string)
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
	if anqp_3gpp_cellular == "" {
		anqp_3gpp_cellular = importOptionChecking(m.(*FortiClient).Cfg, "anqp_3gpp_cellular")
		if anqp_3gpp_cellular == "" {
			return fmt.Errorf("Parameter anqp_3gpp_cellular is missing")
		}
		if err = d.Set("anqp_3gpp_cellular", anqp_3gpp_cellular); err != nil {
			return fmt.Errorf("Error set params anqp_3gpp_cellular: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_3gpp_cellular"] = anqp_3gpp_cellular

	o, err := c.ReadWirelessControllerHotspot20Anqp3GppCellularMccMncList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20Anqp3GppCellularMccMncList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20Anqp3GppCellularMccMncList resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WirelessControllerHotspot20Anqp3GppCellularMccMncList-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("mcc", flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc2edl(o["mcc"], d, "mcc")); err != nil {
		if vv, ok := fortiAPIPatch(o["mcc"], "WirelessControllerHotspot20Anqp3GppCellularMccMncList-Mcc"); ok {
			if err = d.Set("mcc", vv); err != nil {
				return fmt.Errorf("Error reading mcc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mcc: %v", err)
		}
	}

	if err = d.Set("mnc", flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc2edl(o["mnc"], d, "mnc")); err != nil {
		if vv, ok := fortiAPIPatch(o["mnc"], "WirelessControllerHotspot20Anqp3GppCellularMccMncList-Mnc"); ok {
			if err = d.Set("mnc", vv); err != nil {
				return fmt.Errorf("Error reading mnc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mnc: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20Anqp3GppCellularMccMncListId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWirelessControllerHotspot20Anqp3GppCellularMccMncListId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("mcc"); ok || d.HasChange("mcc") {
		t, err := expandWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc2edl(d, v, "mcc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcc"] = t
		}
	}

	if v, ok := d.GetOk("mnc"); ok || d.HasChange("mnc") {
		t, err := expandWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc2edl(d, v, "mnc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mnc"] = t
		}
	}

	return &obj, nil
}
