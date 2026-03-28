// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> DLP sensor entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpSensorEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpSensorEntriesCreate,
		Read:   resourceDlpSensorEntriesRead,
		Update: resourceDlpSensorEntriesUpdate,
		Delete: resourceDlpSensorEntriesDelete,

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sensor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fmgcount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dictionary": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceDlpSensorEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	sensor := d.Get("sensor").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["sensor"] = sensor

	obj, err := getObjectDlpSensorEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpSensorEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDlpSensorEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDlpSensorEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DlpSensorEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDlpSensorEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DlpSensorEntries resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDlpSensorEntriesRead(d, m)
}

func resourceDlpSensorEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	sensor := d.Get("sensor").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["sensor"] = sensor

	obj, err := getObjectDlpSensorEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpSensorEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpSensorEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpSensorEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDlpSensorEntriesRead(d, m)
}

func resourceDlpSensorEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	sensor := d.Get("sensor").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["sensor"] = sensor

	wsParams["adom"] = adomv

	err = c.DeleteDlpSensorEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpSensorEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpSensorEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	sensor := d.Get("sensor").(string)
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
	if sensor == "" {
		sensor = importOptionChecking(m.(*FortiClient).Cfg, "sensor")
		if sensor == "" {
			return fmt.Errorf("Parameter sensor is missing")
		}
		if err = d.Set("sensor", sensor); err != nil {
			return fmt.Errorf("Error set params sensor: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["sensor"] = sensor

	o, err := c.ReadDlpSensorEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpSensorEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpSensorEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpSensorEntries resource from API: %v", err)
	}
	return nil
}

func flattenDlpSensorEntriesCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorEntriesDictionary2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpSensorEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpSensorEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fmgcount", flattenDlpSensorEntriesCount2edl(o["count"], d, "fmgcount")); err != nil {
		if vv, ok := fortiAPIPatch(o["count"], "DlpSensorEntries-Count"); ok {
			if err = d.Set("fmgcount", vv); err != nil {
				return fmt.Errorf("Error reading fmgcount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgcount: %v", err)
		}
	}

	if err = d.Set("dictionary", flattenDlpSensorEntriesDictionary2edl(o["dictionary"], d, "dictionary")); err != nil {
		if vv, ok := fortiAPIPatch(o["dictionary"], "DlpSensorEntries-Dictionary"); ok {
			if err = d.Set("dictionary", vv); err != nil {
				return fmt.Errorf("Error reading dictionary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dictionary: %v", err)
		}
	}

	if err = d.Set("fosid", flattenDlpSensorEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "DlpSensorEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenDlpSensorEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "DlpSensorEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenDlpSensorEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpSensorEntriesCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorEntriesDictionary2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpSensorEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpSensorEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgcount"); ok || d.HasChange("fmgcount") {
		t, err := expandDlpSensorEntriesCount2edl(d, v, "fmgcount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["count"] = t
		}
	}

	if v, ok := d.GetOk("dictionary"); ok || d.HasChange("dictionary") {
		t, err := expandDlpSensorEntriesDictionary2edl(d, v, "dictionary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dictionary"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandDlpSensorEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandDlpSensorEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
