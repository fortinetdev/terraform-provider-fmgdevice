// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Port ranges in the custom entry.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceAdditionEntryPortRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceAdditionEntryPortRangeCreate,
		Read:   resourceFirewallInternetServiceAdditionEntryPortRangeRead,
		Update: resourceFirewallInternetServiceAdditionEntryPortRangeUpdate,
		Delete: resourceFirewallInternetServiceAdditionEntryPortRangeDelete,

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
			"internet_service_addition": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"entry": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"end_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceFirewallInternetServiceAdditionEntryPortRangeCreate(d *schema.ResourceData, m interface{}) error {
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
	internet_service_addition := d.Get("internet_service_addition").(string)
	entry := d.Get("entry").(string)
	paradict["device"] = device_name
	paradict["internet_service_addition"] = internet_service_addition
	paradict["entry"] = entry

	obj, err := getObjectFirewallInternetServiceAdditionEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceAdditionEntryPortRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallInternetServiceAdditionEntryPortRange(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallInternetServiceAdditionEntryPortRange(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallInternetServiceAdditionEntryPortRange resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallInternetServiceAdditionEntryPortRange(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallInternetServiceAdditionEntryPortRange resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallInternetServiceAdditionEntryPortRangeRead(d, m)
}

func resourceFirewallInternetServiceAdditionEntryPortRangeUpdate(d *schema.ResourceData, m interface{}) error {
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
	internet_service_addition := d.Get("internet_service_addition").(string)
	entry := d.Get("entry").(string)
	paradict["device"] = device_name
	paradict["internet_service_addition"] = internet_service_addition
	paradict["entry"] = entry

	obj, err := getObjectFirewallInternetServiceAdditionEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceAdditionEntryPortRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallInternetServiceAdditionEntryPortRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceAdditionEntryPortRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallInternetServiceAdditionEntryPortRangeRead(d, m)
}

func resourceFirewallInternetServiceAdditionEntryPortRangeDelete(d *schema.ResourceData, m interface{}) error {
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
	internet_service_addition := d.Get("internet_service_addition").(string)
	entry := d.Get("entry").(string)
	paradict["device"] = device_name
	paradict["internet_service_addition"] = internet_service_addition
	paradict["entry"] = entry

	wsParams["adom"] = adomv

	err = c.DeleteFirewallInternetServiceAdditionEntryPortRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceAdditionEntryPortRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceAdditionEntryPortRangeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	internet_service_addition := d.Get("internet_service_addition").(string)
	entry := d.Get("entry").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if internet_service_addition == "" {
		internet_service_addition = importOptionChecking(m.(*FortiClient).Cfg, "internet_service_addition")
		if internet_service_addition == "" {
			return fmt.Errorf("Parameter internet_service_addition is missing")
		}
		if err = d.Set("internet_service_addition", internet_service_addition); err != nil {
			return fmt.Errorf("Error set params internet_service_addition: %v", err)
		}
	}
	if entry == "" {
		entry = importOptionChecking(m.(*FortiClient).Cfg, "entry")
		if entry == "" {
			return fmt.Errorf("Parameter entry is missing")
		}
		if err = d.Set("entry", entry); err != nil {
			return fmt.Errorf("Error set params entry: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["internet_service_addition"] = internet_service_addition
	paradict["entry"] = entry

	o, err := c.ReadFirewallInternetServiceAdditionEntryPortRange(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallInternetServiceAdditionEntryPortRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceAdditionEntryPortRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceAdditionEntryPortRange resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceAdditionEntryPortRangeEndPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntryPortRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntryPortRangeStartPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceAdditionEntryPortRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("end_port", flattenFirewallInternetServiceAdditionEntryPortRangeEndPort3rdl(o["end-port"], d, "end_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-port"], "FirewallInternetServiceAdditionEntryPortRange-EndPort"); ok {
			if err = d.Set("end_port", vv); err != nil {
				return fmt.Errorf("Error reading end_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallInternetServiceAdditionEntryPortRangeId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallInternetServiceAdditionEntryPortRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_port", flattenFirewallInternetServiceAdditionEntryPortRangeStartPort3rdl(o["start-port"], d, "start_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-port"], "FirewallInternetServiceAdditionEntryPortRange-StartPort"); ok {
			if err = d.Set("start_port", vv); err != nil {
				return fmt.Errorf("Error reading start_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceAdditionEntryPortRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInternetServiceAdditionEntryPortRangeEndPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntryPortRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntryPortRangeStartPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceAdditionEntryPortRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_port"); ok || d.HasChange("end_port") {
		t, err := expandFirewallInternetServiceAdditionEntryPortRangeEndPort3rdl(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallInternetServiceAdditionEntryPortRangeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok || d.HasChange("start_port") {
		t, err := expandFirewallInternetServiceAdditionEntryPortRangeStartPort3rdl(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	return &obj, nil
}
