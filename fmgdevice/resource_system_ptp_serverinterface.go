// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPtpServerInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPtpServerInterfaceCreate,
		Read:   resourceSystemPtpServerInterfaceRead,
		Update: resourceSystemPtpServerInterfaceUpdate,
		Delete: resourceSystemPtpServerInterfaceDelete,

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
			"delay_mechanism": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"server_interface_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemPtpServerInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemPtpServerInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemPtpServerInterface resource while getting object: %v", err)
	}

	_, err = c.CreateSystemPtpServerInterface(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemPtpServerInterface resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemPtpServerInterfaceRead(d, m)
}

func resourceSystemPtpServerInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemPtpServerInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemPtpServerInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemPtpServerInterface(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemPtpServerInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemPtpServerInterfaceRead(d, m)
}

func resourceSystemPtpServerInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemPtpServerInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPtpServerInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPtpServerInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemPtpServerInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemPtpServerInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPtpServerInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemPtpServerInterface resource from API: %v", err)
	}
	return nil
}

func flattenSystemPtpServerInterfaceDelayMechanism2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPtpServerInterfaceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPtpServerInterfaceServerInterfaceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemPtpServerInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("delay_mechanism", flattenSystemPtpServerInterfaceDelayMechanism2edl(o["delay-mechanism"], d, "delay_mechanism")); err != nil {
		if vv, ok := fortiAPIPatch(o["delay-mechanism"], "SystemPtpServerInterface-DelayMechanism"); ok {
			if err = d.Set("delay_mechanism", vv); err != nil {
				return fmt.Errorf("Error reading delay_mechanism: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delay_mechanism: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemPtpServerInterfaceId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemPtpServerInterface-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("server_interface_name", flattenSystemPtpServerInterfaceServerInterfaceName2edl(o["server-interface-name"], d, "server_interface_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-interface-name"], "SystemPtpServerInterface-ServerInterfaceName"); ok {
			if err = d.Set("server_interface_name", vv); err != nil {
				return fmt.Errorf("Error reading server_interface_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_interface_name: %v", err)
		}
	}

	return nil
}

func flattenSystemPtpServerInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemPtpServerInterfaceDelayMechanism2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpServerInterfaceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpServerInterfaceServerInterfaceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemPtpServerInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("delay_mechanism"); ok || d.HasChange("delay_mechanism") {
		t, err := expandSystemPtpServerInterfaceDelayMechanism2edl(d, v, "delay_mechanism")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delay-mechanism"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemPtpServerInterfaceId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("server_interface_name"); ok || d.HasChange("server_interface_name") {
		t, err := expandSystemPtpServerInterfaceServerInterfaceName2edl(d, v, "server_interface_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-interface-name"] = t
		}
	}

	return &obj, nil
}
