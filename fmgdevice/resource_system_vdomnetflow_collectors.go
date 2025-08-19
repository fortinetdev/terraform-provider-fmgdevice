// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Netflow collectors.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVdomNetflowCollectors() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomNetflowCollectorsCreate,
		Read:   resourceSystemVdomNetflowCollectorsRead,
		Update: resourceSystemVdomNetflowCollectorsUpdate,
		Delete: resourceSystemVdomNetflowCollectorsDelete,

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
			"collector_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"collector_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemVdomNetflowCollectorsCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemVdomNetflowCollectors(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomNetflowCollectors resource while getting object: %v", err)
	}

	_, err = c.CreateSystemVdomNetflowCollectors(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomNetflowCollectors resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemVdomNetflowCollectorsRead(d, m)
}

func resourceSystemVdomNetflowCollectorsUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemVdomNetflowCollectors(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomNetflowCollectors resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVdomNetflowCollectors(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomNetflowCollectors resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemVdomNetflowCollectorsRead(d, m)
}

func resourceSystemVdomNetflowCollectorsDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemVdomNetflowCollectors(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomNetflowCollectors resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomNetflowCollectorsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVdomNetflowCollectors(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomNetflowCollectors resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomNetflowCollectors(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomNetflowCollectors resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomNetflowCollectorsCollectorIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorsCollectorPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorsInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVdomNetflowCollectorsInterfaceSelectMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorsSourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorsSourceIpInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVdomNetflowCollectorsVrfSelect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemVdomNetflowCollectors(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("collector_ip", flattenSystemVdomNetflowCollectorsCollectorIp2edl(o["collector-ip"], d, "collector_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["collector-ip"], "SystemVdomNetflowCollectors-CollectorIp"); ok {
			if err = d.Set("collector_ip", vv); err != nil {
				return fmt.Errorf("Error reading collector_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading collector_ip: %v", err)
		}
	}

	if err = d.Set("collector_port", flattenSystemVdomNetflowCollectorsCollectorPort2edl(o["collector-port"], d, "collector_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["collector-port"], "SystemVdomNetflowCollectors-CollectorPort"); ok {
			if err = d.Set("collector_port", vv); err != nil {
				return fmt.Errorf("Error reading collector_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading collector_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemVdomNetflowCollectorsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemVdomNetflowCollectors-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemVdomNetflowCollectorsInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemVdomNetflowCollectors-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemVdomNetflowCollectorsInterfaceSelectMethod2edl(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystemVdomNetflowCollectors-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemVdomNetflowCollectorsSourceIp2edl(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystemVdomNetflowCollectors-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenSystemVdomNetflowCollectorsSourceIpInterface2edl(o["source-ip-interface"], d, "source_ip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-interface"], "SystemVdomNetflowCollectors-SourceIpInterface"); ok {
			if err = d.Set("source_ip_interface", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenSystemVdomNetflowCollectorsVrfSelect2edl(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "SystemVdomNetflowCollectors-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomNetflowCollectorsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVdomNetflowCollectorsCollectorIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorsCollectorPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorsInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVdomNetflowCollectorsInterfaceSelectMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorsSourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorsSourceIpInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVdomNetflowCollectorsVrfSelect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdomNetflowCollectors(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("collector_ip"); ok || d.HasChange("collector_ip") {
		t, err := expandSystemVdomNetflowCollectorsCollectorIp2edl(d, v, "collector_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-ip"] = t
		}
	}

	if v, ok := d.GetOk("collector_port"); ok || d.HasChange("collector_port") {
		t, err := expandSystemVdomNetflowCollectorsCollectorPort2edl(d, v, "collector_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemVdomNetflowCollectorsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemVdomNetflowCollectorsInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystemVdomNetflowCollectorsInterfaceSelectMethod2edl(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystemVdomNetflowCollectorsSourceIp2edl(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok || d.HasChange("source_ip_interface") {
		t, err := expandSystemVdomNetflowCollectorsSourceIpInterface2edl(d, v, "source_ip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-interface"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandSystemVdomNetflowCollectorsVrfSelect2edl(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
