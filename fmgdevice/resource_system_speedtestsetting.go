// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure speed test setting.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSpeedTestSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSpeedTestSettingUpdate,
		Read:   resourceSystemSpeedTestSettingRead,
		Update: resourceSystemSpeedTestSettingUpdate,
		Delete: resourceSystemSpeedTestSettingDelete,

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
			"latency_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"multiple_tcp_stream": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSpeedTestSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSpeedTestSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSpeedTestSetting(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSpeedTestSetting")

	return resourceSystemSpeedTestSettingRead(d, m)
}

func resourceSystemSpeedTestSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSpeedTestSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSpeedTestSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSpeedTestSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSpeedTestSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpeedTestSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSpeedTestSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpeedTestSetting resource from API: %v", err)
	}
	return nil
}

func flattenSystemSpeedTestSettingLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestSettingMultipleTcpStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSpeedTestSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("latency_threshold", flattenSystemSpeedTestSettingLatencyThreshold(o["latency-threshold"], d, "latency_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold"], "SystemSpeedTestSetting-LatencyThreshold"); ok {
			if err = d.Set("latency_threshold", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold: %v", err)
		}
	}

	if err = d.Set("multiple_tcp_stream", flattenSystemSpeedTestSettingMultipleTcpStream(o["multiple-tcp-stream"], d, "multiple_tcp_stream")); err != nil {
		if vv, ok := fortiAPIPatch(o["multiple-tcp-stream"], "SystemSpeedTestSetting-MultipleTcpStream"); ok {
			if err = d.Set("multiple_tcp_stream", vv); err != nil {
				return fmt.Errorf("Error reading multiple_tcp_stream: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multiple_tcp_stream: %v", err)
		}
	}

	return nil
}

func flattenSystemSpeedTestSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSpeedTestSettingLatencyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestSettingMultipleTcpStream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSpeedTestSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("latency_threshold"); ok || d.HasChange("latency_threshold") {
		t, err := expandSystemSpeedTestSettingLatencyThreshold(d, v, "latency_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold"] = t
		}
	}

	if v, ok := d.GetOk("multiple_tcp_stream"); ok || d.HasChange("multiple_tcp_stream") {
		t, err := expandSystemSpeedTestSettingMultipleTcpStream(d, v, "multiple_tcp_stream")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple-tcp-stream"] = t
		}
	}

	return &obj, nil
}
