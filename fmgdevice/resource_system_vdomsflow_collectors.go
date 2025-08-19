// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: sFlow collectors.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVdomSflowCollectors() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomSflowCollectorsCreate,
		Read:   resourceSystemVdomSflowCollectorsRead,
		Update: resourceSystemVdomSflowCollectorsUpdate,
		Delete: resourceSystemVdomSflowCollectorsDelete,

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
				Computed: true,
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
				Computed: true,
			},
		},
	}
}

func resourceSystemVdomSflowCollectorsCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemVdomSflowCollectors(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomSflowCollectors resource while getting object: %v", err)
	}

	_, err = c.CreateSystemVdomSflowCollectors(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomSflowCollectors resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemVdomSflowCollectorsRead(d, m)
}

func resourceSystemVdomSflowCollectorsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemVdomSflowCollectors(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomSflowCollectors resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVdomSflowCollectors(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomSflowCollectors resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemVdomSflowCollectorsRead(d, m)
}

func resourceSystemVdomSflowCollectorsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemVdomSflowCollectors(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomSflowCollectors resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomSflowCollectorsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVdomSflowCollectors(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomSflowCollectors resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomSflowCollectors(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomSflowCollectors resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomSflowCollectorsCollectorIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomSflowCollectorsCollectorPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomSflowCollectorsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomSflowCollectorsInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVdomSflowCollectorsInterfaceSelectMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomSflowCollectorsSourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemVdomSflowCollectors(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("collector_ip", flattenSystemVdomSflowCollectorsCollectorIp2edl(o["collector-ip"], d, "collector_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["collector-ip"], "SystemVdomSflowCollectors-CollectorIp"); ok {
			if err = d.Set("collector_ip", vv); err != nil {
				return fmt.Errorf("Error reading collector_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading collector_ip: %v", err)
		}
	}

	if err = d.Set("collector_port", flattenSystemVdomSflowCollectorsCollectorPort2edl(o["collector-port"], d, "collector_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["collector-port"], "SystemVdomSflowCollectors-CollectorPort"); ok {
			if err = d.Set("collector_port", vv); err != nil {
				return fmt.Errorf("Error reading collector_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading collector_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemVdomSflowCollectorsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemVdomSflowCollectors-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemVdomSflowCollectorsInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemVdomSflowCollectors-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemVdomSflowCollectorsInterfaceSelectMethod2edl(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystemVdomSflowCollectors-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemVdomSflowCollectorsSourceIp2edl(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystemVdomSflowCollectors-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomSflowCollectorsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVdomSflowCollectorsCollectorIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomSflowCollectorsCollectorPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomSflowCollectorsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomSflowCollectorsInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVdomSflowCollectorsInterfaceSelectMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomSflowCollectorsSourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdomSflowCollectors(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("collector_ip"); ok || d.HasChange("collector_ip") {
		t, err := expandSystemVdomSflowCollectorsCollectorIp2edl(d, v, "collector_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-ip"] = t
		}
	}

	if v, ok := d.GetOk("collector_port"); ok || d.HasChange("collector_port") {
		t, err := expandSystemVdomSflowCollectorsCollectorPort2edl(d, v, "collector_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemVdomSflowCollectorsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemVdomSflowCollectorsInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystemVdomSflowCollectorsInterfaceSelectMethod2edl(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystemVdomSflowCollectorsSourceIp2edl(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	return &obj, nil
}
