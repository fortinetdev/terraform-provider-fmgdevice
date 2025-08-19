// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure system PTP information.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPtpUpdate,
		Read:   resourceSystemPtpRead,
		Update: resourceSystemPtpUpdate,
		Delete: resourceSystemPtpDelete,

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
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delay_mechanism": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"server_interface_name": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"server_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemPtpUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemPtp(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemPtp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemPtp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemPtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemPtp")

	return resourceSystemPtpRead(d, m)
}

func resourceSystemPtpDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemPtp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPtpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemPtp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemPtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPtp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemPtp resource from API: %v", err)
	}
	return nil
}

func flattenSystemPtpDelayMechanism(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPtpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPtpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPtpRequestInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPtpServerInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delay_mechanism"
		if _, ok := i["delay-mechanism"]; ok {
			v := flattenSystemPtpServerInterfaceDelayMechanism(i["delay-mechanism"], d, pre_append)
			tmp["delay_mechanism"] = fortiAPISubPartPatch(v, "SystemPtp-ServerInterface-DelayMechanism")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemPtpServerInterfaceId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemPtp-ServerInterface-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_interface_name"
		if _, ok := i["server-interface-name"]; ok {
			v := flattenSystemPtpServerInterfaceServerInterfaceName(i["server-interface-name"], d, pre_append)
			tmp["server_interface_name"] = fortiAPISubPartPatch(v, "SystemPtp-ServerInterface-ServerInterfaceName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemPtpServerInterfaceDelayMechanism(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPtpServerInterfaceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPtpServerInterfaceServerInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPtpServerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemPtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("delay_mechanism", flattenSystemPtpDelayMechanism(o["delay-mechanism"], d, "delay_mechanism")); err != nil {
		if vv, ok := fortiAPIPatch(o["delay-mechanism"], "SystemPtp-DelayMechanism"); ok {
			if err = d.Set("delay_mechanism", vv); err != nil {
				return fmt.Errorf("Error reading delay_mechanism: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delay_mechanism: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemPtpInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemPtp-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemPtpMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemPtp-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("request_interval", flattenSystemPtpRequestInterval(o["request-interval"], d, "request_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["request-interval"], "SystemPtp-RequestInterval"); ok {
			if err = d.Set("request_interval", vv); err != nil {
				return fmt.Errorf("Error reading request_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request_interval: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_interface", flattenSystemPtpServerInterface(o["server-interface"], d, "server_interface")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-interface"], "SystemPtp-ServerInterface"); ok {
				if err = d.Set("server_interface", vv); err != nil {
					return fmt.Errorf("Error reading server_interface: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_interface"); ok {
			if err = d.Set("server_interface", flattenSystemPtpServerInterface(o["server-interface"], d, "server_interface")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-interface"], "SystemPtp-ServerInterface"); ok {
					if err = d.Set("server_interface", vv); err != nil {
						return fmt.Errorf("Error reading server_interface: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_mode", flattenSystemPtpServerMode(o["server-mode"], d, "server_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-mode"], "SystemPtp-ServerMode"); ok {
			if err = d.Set("server_mode", vv); err != nil {
				return fmt.Errorf("Error reading server_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_mode: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemPtpStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemPtp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemPtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemPtpDelayMechanism(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPtpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpRequestInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpServerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delay_mechanism"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["delay-mechanism"], _ = expandSystemPtpServerInterfaceDelayMechanism(d, i["delay_mechanism"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemPtpServerInterfaceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_interface_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-interface-name"], _ = expandSystemPtpServerInterfaceServerInterfaceName(d, i["server_interface_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemPtpServerInterfaceDelayMechanism(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpServerInterfaceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpServerInterfaceServerInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPtpServerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPtp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("delay_mechanism"); ok || d.HasChange("delay_mechanism") {
		t, err := expandSystemPtpDelayMechanism(d, v, "delay_mechanism")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delay-mechanism"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemPtpInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemPtpMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("request_interval"); ok || d.HasChange("request_interval") {
		t, err := expandSystemPtpRequestInterval(d, v, "request_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-interval"] = t
		}
	}

	if v, ok := d.GetOk("server_interface"); ok || d.HasChange("server_interface") {
		t, err := expandSystemPtpServerInterface(d, v, "server_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-interface"] = t
		}
	}

	if v, ok := d.GetOk("server_mode"); ok || d.HasChange("server_mode") {
		t, err := expandSystemPtpServerMode(d, v, "server_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-mode"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemPtpStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
