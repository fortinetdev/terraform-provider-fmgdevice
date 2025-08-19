// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure logging to FortiCloud.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogFortiguardSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortiguardSettingUpdate,
		Read:   resourceLogFortiguardSettingRead,
		Update: resourceLogFortiguardSettingUpdate,
		Delete: resourceLogFortiguardSettingDelete,

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
			"access_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conn_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
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
			"max_log_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceLogFortiguardSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogFortiguardSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogFortiguardSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortiguardSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogFortiguardSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogFortiguardSetting")

	return resourceLogFortiguardSettingRead(d, m)
}

func resourceLogFortiguardSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogFortiguardSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogFortiguardSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortiguardSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogFortiguardSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogFortiguardSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortiguardSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogFortiguardSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogFortiguardSettingAccessConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingConnTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogFortiguardSettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingMaxLogRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingUploadDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingUploadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingUploadOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingUploadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogFortiguardSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("access_config", flattenLogFortiguardSettingAccessConfig(o["access-config"], d, "access_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-config"], "LogFortiguardSetting-AccessConfig"); ok {
			if err = d.Set("access_config", vv); err != nil {
				return fmt.Errorf("Error reading access_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_config: %v", err)
		}
	}

	if err = d.Set("conn_timeout", flattenLogFortiguardSettingConnTimeout(o["conn-timeout"], d, "conn_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn-timeout"], "LogFortiguardSetting-ConnTimeout"); ok {
			if err = d.Set("conn_timeout", vv); err != nil {
				return fmt.Errorf("Error reading conn_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_timeout: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogFortiguardSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["enc-algorithm"], "LogFortiguardSetting-EncAlgorithm"); ok {
			if err = d.Set("enc_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading enc_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("interface", flattenLogFortiguardSettingInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "LogFortiguardSetting-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenLogFortiguardSettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "LogFortiguardSetting-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("max_log_rate", flattenLogFortiguardSettingMaxLogRate(o["max-log-rate"], d, "max_log_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-rate"], "LogFortiguardSetting-MaxLogRate"); ok {
			if err = d.Set("max_log_rate", vv); err != nil {
				return fmt.Errorf("Error reading max_log_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_rate: %v", err)
		}
	}

	if err = d.Set("priority", flattenLogFortiguardSettingPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "LogFortiguardSetting-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogFortiguardSettingSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "LogFortiguardSetting-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenLogFortiguardSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "LogFortiguardSetting-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("status", flattenLogFortiguardSettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "LogFortiguardSetting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upload_day", flattenLogFortiguardSettingUploadDay(o["upload-day"], d, "upload_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-day"], "LogFortiguardSetting-UploadDay"); ok {
			if err = d.Set("upload_day", vv); err != nil {
				return fmt.Errorf("Error reading upload_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_day: %v", err)
		}
	}

	if err = d.Set("upload_interval", flattenLogFortiguardSettingUploadInterval(o["upload-interval"], d, "upload_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-interval"], "LogFortiguardSetting-UploadInterval"); ok {
			if err = d.Set("upload_interval", vv); err != nil {
				return fmt.Errorf("Error reading upload_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_interval: %v", err)
		}
	}

	if err = d.Set("upload_option", flattenLogFortiguardSettingUploadOption(o["upload-option"], d, "upload_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-option"], "LogFortiguardSetting-UploadOption"); ok {
			if err = d.Set("upload_option", vv); err != nil {
				return fmt.Errorf("Error reading upload_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_option: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenLogFortiguardSettingUploadTime(o["upload-time"], d, "upload_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-time"], "LogFortiguardSetting-UploadTime"); ok {
			if err = d.Set("upload_time", vv); err != nil {
				return fmt.Errorf("Error reading upload_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenLogFortiguardSettingVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "LogFortiguardSetting-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenLogFortiguardSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogFortiguardSettingAccessConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingConnTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogFortiguardSettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingMaxLogRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingUploadDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingUploadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingUploadOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingUploadTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortiguardSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_config"); ok || d.HasChange("access_config") {
		t, err := expandLogFortiguardSettingAccessConfig(d, v, "access_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-config"] = t
		}
	}

	if v, ok := d.GetOk("conn_timeout"); ok || d.HasChange("conn_timeout") {
		t, err := expandLogFortiguardSettingConnTimeout(d, v, "conn_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-timeout"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok || d.HasChange("enc_algorithm") {
		t, err := expandLogFortiguardSettingEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandLogFortiguardSettingInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandLogFortiguardSettingInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("max_log_rate"); ok || d.HasChange("max_log_rate") {
		t, err := expandLogFortiguardSettingMaxLogRate(d, v, "max_log_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-rate"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandLogFortiguardSettingPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandLogFortiguardSettingSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok || d.HasChange("ssl_min_proto_version") {
		t, err := expandLogFortiguardSettingSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandLogFortiguardSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("upload_day"); ok || d.HasChange("upload_day") {
		t, err := expandLogFortiguardSettingUploadDay(d, v, "upload_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-day"] = t
		}
	}

	if v, ok := d.GetOk("upload_interval"); ok || d.HasChange("upload_interval") {
		t, err := expandLogFortiguardSettingUploadInterval(d, v, "upload_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-interval"] = t
		}
	}

	if v, ok := d.GetOk("upload_option"); ok || d.HasChange("upload_option") {
		t, err := expandLogFortiguardSettingUploadOption(d, v, "upload_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-option"] = t
		}
	}

	if v, ok := d.GetOk("upload_time"); ok || d.HasChange("upload_time") {
		t, err := expandLogFortiguardSettingUploadTime(d, v, "upload_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-time"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandLogFortiguardSettingVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
