// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Port ranges in the definition entry.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceDefinitionEntryPortRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceDefinitionEntryPortRangeCreate,
		Read:   resourceFirewallInternetServiceDefinitionEntryPortRangeRead,
		Update: resourceFirewallInternetServiceDefinitionEntryPortRangeUpdate,
		Delete: resourceFirewallInternetServiceDefinitionEntryPortRangeDelete,

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
			"internet_service_definition": &schema.Schema{
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
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallInternetServiceDefinitionEntryPortRangeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	internet_service_definition := d.Get("internet_service_definition").(string)
	entry := d.Get("entry").(string)
	paradict["device"] = device_name
	paradict["internet_service_definition"] = internet_service_definition
	paradict["entry"] = entry

	obj, err := getObjectFirewallInternetServiceDefinitionEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceDefinitionEntryPortRange resource while getting object: %v", err)
	}

	_, err = c.CreateFirewallInternetServiceDefinitionEntryPortRange(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceDefinitionEntryPortRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallInternetServiceDefinitionEntryPortRangeRead(d, m)
}

func resourceFirewallInternetServiceDefinitionEntryPortRangeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	internet_service_definition := d.Get("internet_service_definition").(string)
	entry := d.Get("entry").(string)
	paradict["device"] = device_name
	paradict["internet_service_definition"] = internet_service_definition
	paradict["entry"] = entry

	obj, err := getObjectFirewallInternetServiceDefinitionEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceDefinitionEntryPortRange resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallInternetServiceDefinitionEntryPortRange(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceDefinitionEntryPortRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallInternetServiceDefinitionEntryPortRangeRead(d, m)
}

func resourceFirewallInternetServiceDefinitionEntryPortRangeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	internet_service_definition := d.Get("internet_service_definition").(string)
	entry := d.Get("entry").(string)
	paradict["device"] = device_name
	paradict["internet_service_definition"] = internet_service_definition
	paradict["entry"] = entry

	err = c.DeleteFirewallInternetServiceDefinitionEntryPortRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceDefinitionEntryPortRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceDefinitionEntryPortRangeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	internet_service_definition := d.Get("internet_service_definition").(string)
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
	if internet_service_definition == "" {
		internet_service_definition = importOptionChecking(m.(*FortiClient).Cfg, "internet_service_definition")
		if internet_service_definition == "" {
			return fmt.Errorf("Parameter internet_service_definition is missing")
		}
		if err = d.Set("internet_service_definition", internet_service_definition); err != nil {
			return fmt.Errorf("Error set params internet_service_definition: %v", err)
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
	paradict["internet_service_definition"] = internet_service_definition
	paradict["entry"] = entry

	o, err := c.ReadFirewallInternetServiceDefinitionEntryPortRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceDefinitionEntryPortRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceDefinitionEntryPortRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceDefinitionEntryPortRange resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceDefinitionEntryPortRangeEndPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceDefinitionEntryPortRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceDefinitionEntryPortRangeStartPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceDefinitionEntryPortRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("end_port", flattenFirewallInternetServiceDefinitionEntryPortRangeEndPort3rdl(o["end-port"], d, "end_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-port"], "FirewallInternetServiceDefinitionEntryPortRange-EndPort"); ok {
			if err = d.Set("end_port", vv); err != nil {
				return fmt.Errorf("Error reading end_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallInternetServiceDefinitionEntryPortRangeId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallInternetServiceDefinitionEntryPortRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_port", flattenFirewallInternetServiceDefinitionEntryPortRangeStartPort3rdl(o["start-port"], d, "start_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-port"], "FirewallInternetServiceDefinitionEntryPortRange-StartPort"); ok {
			if err = d.Set("start_port", vv); err != nil {
				return fmt.Errorf("Error reading start_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceDefinitionEntryPortRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInternetServiceDefinitionEntryPortRangeEndPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDefinitionEntryPortRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDefinitionEntryPortRangeStartPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceDefinitionEntryPortRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_port"); ok || d.HasChange("end_port") {
		t, err := expandFirewallInternetServiceDefinitionEntryPortRangeEndPort3rdl(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallInternetServiceDefinitionEntryPortRangeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok || d.HasChange("start_port") {
		t, err := expandFirewallInternetServiceDefinitionEntryPortRangeStartPort3rdl(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	return &obj, nil
}
