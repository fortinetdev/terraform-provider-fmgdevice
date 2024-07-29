// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Speed test schedule for each interface.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSpeedTestSchedule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSpeedTestScheduleCreate,
		Read:   resourceSystemSpeedTestScheduleRead,
		Update: resourceSystemSpeedTestScheduleUpdate,
		Delete: resourceSystemSpeedTestScheduleDelete,

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
			"ctrl_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"diffserv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedules": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_inbandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_inbandwidth_maximum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"update_inbandwidth_minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"update_outbandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_outbandwidth_maximum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"update_outbandwidth_minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"update_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSpeedTestScheduleCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemSpeedTestSchedule(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSpeedTestSchedule resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSpeedTestSchedule(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSpeedTestSchedule resource: %v", err)
	}

	d.SetId(getStringKey(d, "interface"))

	return resourceSystemSpeedTestScheduleRead(d, m)
}

func resourceSystemSpeedTestScheduleUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemSpeedTestSchedule(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestSchedule resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSpeedTestSchedule(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestSchedule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "interface"))

	return resourceSystemSpeedTestScheduleRead(d, m)
}

func resourceSystemSpeedTestScheduleDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteSystemSpeedTestSchedule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSpeedTestSchedule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSpeedTestScheduleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadSystemSpeedTestSchedule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpeedTestSchedule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSpeedTestSchedule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpeedTestSchedule resource from API: %v", err)
	}
	return nil
}

func flattenSystemSpeedTestScheduleCtrlPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleDiffserv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleDynamicServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleSchedules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSpeedTestScheduleServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateInbandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateInbandwidthMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateInbandwidthMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateOutbandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateOutbandwidthMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateOutbandwidthMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSpeedTestSchedule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ctrl_port", flattenSystemSpeedTestScheduleCtrlPort(o["ctrl-port"], d, "ctrl_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ctrl-port"], "SystemSpeedTestSchedule-CtrlPort"); ok {
			if err = d.Set("ctrl_port", vv); err != nil {
				return fmt.Errorf("Error reading ctrl_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ctrl_port: %v", err)
		}
	}

	if err = d.Set("diffserv", flattenSystemSpeedTestScheduleDiffserv(o["diffserv"], d, "diffserv")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv"], "SystemSpeedTestSchedule-Diffserv"); ok {
			if err = d.Set("diffserv", vv); err != nil {
				return fmt.Errorf("Error reading diffserv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv: %v", err)
		}
	}

	if err = d.Set("dynamic_server", flattenSystemSpeedTestScheduleDynamicServer(o["dynamic-server"], d, "dynamic_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-server"], "SystemSpeedTestSchedule-DynamicServer"); ok {
			if err = d.Set("dynamic_server", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_server: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemSpeedTestScheduleInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemSpeedTestSchedule-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemSpeedTestScheduleMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemSpeedTestSchedule-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("schedules", flattenSystemSpeedTestScheduleSchedules(o["schedules"], d, "schedules")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedules"], "SystemSpeedTestSchedule-Schedules"); ok {
			if err = d.Set("schedules", vv); err != nil {
				return fmt.Errorf("Error reading schedules: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedules: %v", err)
		}
	}

	if err = d.Set("server_name", flattenSystemSpeedTestScheduleServerName(o["server-name"], d, "server_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-name"], "SystemSpeedTestSchedule-ServerName"); ok {
			if err = d.Set("server_name", vv); err != nil {
				return fmt.Errorf("Error reading server_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	if err = d.Set("server_port", flattenSystemSpeedTestScheduleServerPort(o["server-port"], d, "server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-port"], "SystemSpeedTestSchedule-ServerPort"); ok {
			if err = d.Set("server_port", vv); err != nil {
				return fmt.Errorf("Error reading server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSpeedTestScheduleStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSpeedTestSchedule-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("update_inbandwidth", flattenSystemSpeedTestScheduleUpdateInbandwidth(o["update-inbandwidth"], d, "update_inbandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-inbandwidth"], "SystemSpeedTestSchedule-UpdateInbandwidth"); ok {
			if err = d.Set("update_inbandwidth", vv); err != nil {
				return fmt.Errorf("Error reading update_inbandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_inbandwidth: %v", err)
		}
	}

	if err = d.Set("update_inbandwidth_maximum", flattenSystemSpeedTestScheduleUpdateInbandwidthMaximum(o["update-inbandwidth-maximum"], d, "update_inbandwidth_maximum")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-inbandwidth-maximum"], "SystemSpeedTestSchedule-UpdateInbandwidthMaximum"); ok {
			if err = d.Set("update_inbandwidth_maximum", vv); err != nil {
				return fmt.Errorf("Error reading update_inbandwidth_maximum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_inbandwidth_maximum: %v", err)
		}
	}

	if err = d.Set("update_inbandwidth_minimum", flattenSystemSpeedTestScheduleUpdateInbandwidthMinimum(o["update-inbandwidth-minimum"], d, "update_inbandwidth_minimum")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-inbandwidth-minimum"], "SystemSpeedTestSchedule-UpdateInbandwidthMinimum"); ok {
			if err = d.Set("update_inbandwidth_minimum", vv); err != nil {
				return fmt.Errorf("Error reading update_inbandwidth_minimum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_inbandwidth_minimum: %v", err)
		}
	}

	if err = d.Set("update_outbandwidth", flattenSystemSpeedTestScheduleUpdateOutbandwidth(o["update-outbandwidth"], d, "update_outbandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-outbandwidth"], "SystemSpeedTestSchedule-UpdateOutbandwidth"); ok {
			if err = d.Set("update_outbandwidth", vv); err != nil {
				return fmt.Errorf("Error reading update_outbandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_outbandwidth: %v", err)
		}
	}

	if err = d.Set("update_outbandwidth_maximum", flattenSystemSpeedTestScheduleUpdateOutbandwidthMaximum(o["update-outbandwidth-maximum"], d, "update_outbandwidth_maximum")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-outbandwidth-maximum"], "SystemSpeedTestSchedule-UpdateOutbandwidthMaximum"); ok {
			if err = d.Set("update_outbandwidth_maximum", vv); err != nil {
				return fmt.Errorf("Error reading update_outbandwidth_maximum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_outbandwidth_maximum: %v", err)
		}
	}

	if err = d.Set("update_outbandwidth_minimum", flattenSystemSpeedTestScheduleUpdateOutbandwidthMinimum(o["update-outbandwidth-minimum"], d, "update_outbandwidth_minimum")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-outbandwidth-minimum"], "SystemSpeedTestSchedule-UpdateOutbandwidthMinimum"); ok {
			if err = d.Set("update_outbandwidth_minimum", vv); err != nil {
				return fmt.Errorf("Error reading update_outbandwidth_minimum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_outbandwidth_minimum: %v", err)
		}
	}

	if err = d.Set("update_shaper", flattenSystemSpeedTestScheduleUpdateShaper(o["update-shaper"], d, "update_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-shaper"], "SystemSpeedTestSchedule-UpdateShaper"); ok {
			if err = d.Set("update_shaper", vv); err != nil {
				return fmt.Errorf("Error reading update_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_shaper: %v", err)
		}
	}

	return nil
}

func flattenSystemSpeedTestScheduleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSpeedTestScheduleCtrlPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleDiffserv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleDynamicServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleSchedules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSpeedTestScheduleServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateInbandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateInbandwidthMaximum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateInbandwidthMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateOutbandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateOutbandwidthMaximum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateOutbandwidthMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSpeedTestSchedule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ctrl_port"); ok || d.HasChange("ctrl_port") {
		t, err := expandSystemSpeedTestScheduleCtrlPort(d, v, "ctrl_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ctrl-port"] = t
		}
	}

	if v, ok := d.GetOk("diffserv"); ok || d.HasChange("diffserv") {
		t, err := expandSystemSpeedTestScheduleDiffserv(d, v, "diffserv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_server"); ok || d.HasChange("dynamic_server") {
		t, err := expandSystemSpeedTestScheduleDynamicServer(d, v, "dynamic_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-server"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemSpeedTestScheduleInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemSpeedTestScheduleMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("schedules"); ok || d.HasChange("schedules") {
		t, err := expandSystemSpeedTestScheduleSchedules(d, v, "schedules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedules"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok || d.HasChange("server_name") {
		t, err := expandSystemSpeedTestScheduleServerName(d, v, "server_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	if v, ok := d.GetOk("server_port"); ok || d.HasChange("server_port") {
		t, err := expandSystemSpeedTestScheduleServerPort(d, v, "server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSpeedTestScheduleStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("update_inbandwidth"); ok || d.HasChange("update_inbandwidth") {
		t, err := expandSystemSpeedTestScheduleUpdateInbandwidth(d, v, "update_inbandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-inbandwidth"] = t
		}
	}

	if v, ok := d.GetOk("update_inbandwidth_maximum"); ok || d.HasChange("update_inbandwidth_maximum") {
		t, err := expandSystemSpeedTestScheduleUpdateInbandwidthMaximum(d, v, "update_inbandwidth_maximum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-inbandwidth-maximum"] = t
		}
	}

	if v, ok := d.GetOk("update_inbandwidth_minimum"); ok || d.HasChange("update_inbandwidth_minimum") {
		t, err := expandSystemSpeedTestScheduleUpdateInbandwidthMinimum(d, v, "update_inbandwidth_minimum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-inbandwidth-minimum"] = t
		}
	}

	if v, ok := d.GetOk("update_outbandwidth"); ok || d.HasChange("update_outbandwidth") {
		t, err := expandSystemSpeedTestScheduleUpdateOutbandwidth(d, v, "update_outbandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-outbandwidth"] = t
		}
	}

	if v, ok := d.GetOk("update_outbandwidth_maximum"); ok || d.HasChange("update_outbandwidth_maximum") {
		t, err := expandSystemSpeedTestScheduleUpdateOutbandwidthMaximum(d, v, "update_outbandwidth_maximum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-outbandwidth-maximum"] = t
		}
	}

	if v, ok := d.GetOk("update_outbandwidth_minimum"); ok || d.HasChange("update_outbandwidth_minimum") {
		t, err := expandSystemSpeedTestScheduleUpdateOutbandwidthMinimum(d, v, "update_outbandwidth_minimum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-outbandwidth-minimum"] = t
		}
	}

	if v, ok := d.GetOk("update_shaper"); ok || d.HasChange("update_shaper") {
		t, err := expandSystemSpeedTestScheduleUpdateShaper(d, v, "update_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-shaper"] = t
		}
	}

	return &obj, nil
}
