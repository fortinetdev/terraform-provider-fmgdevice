// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Port ranges in the custom entry.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceExtensionEntryPortRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceExtensionEntryPortRangeCreate,
		Read:   resourceFirewallInternetServiceExtensionEntryPortRangeRead,
		Update: resourceFirewallInternetServiceExtensionEntryPortRangeUpdate,
		Delete: resourceFirewallInternetServiceExtensionEntryPortRangeDelete,

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
			"internet_service_extension": &schema.Schema{
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

func resourceFirewallInternetServiceExtensionEntryPortRangeCreate(d *schema.ResourceData, m interface{}) error {
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
	internet_service_extension := d.Get("internet_service_extension").(string)
	entry := d.Get("entry").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["internet_service_extension"] = internet_service_extension
	paradict["entry"] = entry

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectFirewallInternetServiceExtensionEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceExtensionEntryPortRange resource while getting object: %v", err)
	}

	_, err = c.CreateFirewallInternetServiceExtensionEntryPortRange(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceExtensionEntryPortRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallInternetServiceExtensionEntryPortRangeRead(d, m)
}

func resourceFirewallInternetServiceExtensionEntryPortRangeUpdate(d *schema.ResourceData, m interface{}) error {
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
	internet_service_extension := d.Get("internet_service_extension").(string)
	entry := d.Get("entry").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["internet_service_extension"] = internet_service_extension
	paradict["entry"] = entry

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectFirewallInternetServiceExtensionEntryPortRange(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceExtensionEntryPortRange resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallInternetServiceExtensionEntryPortRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceExtensionEntryPortRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallInternetServiceExtensionEntryPortRangeRead(d, m)
}

func resourceFirewallInternetServiceExtensionEntryPortRangeDelete(d *schema.ResourceData, m interface{}) error {
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
	internet_service_extension := d.Get("internet_service_extension").(string)
	entry := d.Get("entry").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["internet_service_extension"] = internet_service_extension
	paradict["entry"] = entry

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteFirewallInternetServiceExtensionEntryPortRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceExtensionEntryPortRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceExtensionEntryPortRangeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	internet_service_extension := d.Get("internet_service_extension").(string)
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
	if device_vdom == "" {
		device_vdom = importOptionChecking(m.(*FortiClient).Cfg, "device_vdom")
		if device_vdom == "" {
			return fmt.Errorf("Parameter device_vdom is missing")
		}
		if err = d.Set("device_vdom", device_vdom); err != nil {
			return fmt.Errorf("Error set params device_vdom: %v", err)
		}
	}
	if internet_service_extension == "" {
		internet_service_extension = importOptionChecking(m.(*FortiClient).Cfg, "internet_service_extension")
		if internet_service_extension == "" {
			return fmt.Errorf("Parameter internet_service_extension is missing")
		}
		if err = d.Set("internet_service_extension", internet_service_extension); err != nil {
			return fmt.Errorf("Error set params internet_service_extension: %v", err)
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
	paradict["vdom"] = device_vdom
	paradict["internet_service_extension"] = internet_service_extension
	paradict["entry"] = entry

	o, err := c.ReadFirewallInternetServiceExtensionEntryPortRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceExtensionEntryPortRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceExtensionEntryPortRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceExtensionEntryPortRange resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceExtensionEntryPortRangeEndPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtensionEntryPortRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtensionEntryPortRangeStartPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceExtensionEntryPortRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("end_port", flattenFirewallInternetServiceExtensionEntryPortRangeEndPort3rdl(o["end-port"], d, "end_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-port"], "FirewallInternetServiceExtensionEntryPortRange-EndPort"); ok {
			if err = d.Set("end_port", vv); err != nil {
				return fmt.Errorf("Error reading end_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallInternetServiceExtensionEntryPortRangeId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallInternetServiceExtensionEntryPortRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_port", flattenFirewallInternetServiceExtensionEntryPortRangeStartPort3rdl(o["start-port"], d, "start_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-port"], "FirewallInternetServiceExtensionEntryPortRange-StartPort"); ok {
			if err = d.Set("start_port", vv); err != nil {
				return fmt.Errorf("Error reading start_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceExtensionEntryPortRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInternetServiceExtensionEntryPortRangeEndPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionEntryPortRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionEntryPortRangeStartPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceExtensionEntryPortRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_port"); ok || d.HasChange("end_port") {
		t, err := expandFirewallInternetServiceExtensionEntryPortRangeEndPort3rdl(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallInternetServiceExtensionEntryPortRangeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok || d.HasChange("start_port") {
		t, err := expandFirewallInternetServiceExtensionEntryPortRangeStartPort3rdl(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	return &obj, nil
}
