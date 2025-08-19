// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Global settings for remote syslog server.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogSyslogd2Setting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogd2SettingUpdate,
		Read:   resourceLogSyslogd2SettingRead,
		Update: resourceLogSyslogd2SettingUpdate,
		Delete: resourceLogSyslogd2SettingDelete,

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
			"certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"custom_field_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"facility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceLogSyslogd2SettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogSyslogd2Setting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd2Setting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogSyslogd2Setting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd2Setting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogSyslogd2Setting")

	return resourceLogSyslogd2SettingRead(d, m)
}

func resourceLogSyslogd2SettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogSyslogd2Setting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogSyslogd2Setting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogd2SettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogSyslogd2Setting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd2Setting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogd2Setting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd2Setting resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogd2SettingCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogSyslogd2SettingCustomFieldName(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom"
		if _, ok := i["custom"]; ok {
			v := flattenLogSyslogd2SettingCustomFieldNameCustom(i["custom"], d, pre_append)
			tmp["custom"] = fortiAPISubPartPatch(v, "LogSyslogd2Setting-CustomFieldName-Custom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenLogSyslogd2SettingCustomFieldNameId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "LogSyslogd2Setting-CustomFieldName-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenLogSyslogd2SettingCustomFieldNameName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "LogSyslogd2Setting-CustomFieldName-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenLogSyslogd2SettingCustomFieldNameCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingCustomFieldNameId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingCustomFieldNameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingFacility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogSyslogd2SettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingMaxLogRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingSourceIpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogSyslogd2SettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2SettingVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogSyslogd2Setting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("certificate", flattenLogSyslogd2SettingCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "LogSyslogd2Setting-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_field_name", flattenLogSyslogd2SettingCustomFieldName(o["custom-field-name"], d, "custom_field_name")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-field-name"], "LogSyslogd2Setting-CustomFieldName"); ok {
				if err = d.Set("custom_field_name", vv); err != nil {
					return fmt.Errorf("Error reading custom_field_name: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_field_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_field_name"); ok {
			if err = d.Set("custom_field_name", flattenLogSyslogd2SettingCustomFieldName(o["custom-field-name"], d, "custom_field_name")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-field-name"], "LogSyslogd2Setting-CustomFieldName"); ok {
					if err = d.Set("custom_field_name", vv); err != nil {
						return fmt.Errorf("Error reading custom_field_name: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_field_name: %v", err)
				}
			}
		}
	}

	if err = d.Set("enc_algorithm", flattenLogSyslogd2SettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["enc-algorithm"], "LogSyslogd2Setting-EncAlgorithm"); ok {
			if err = d.Set("enc_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading enc_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("facility", flattenLogSyslogd2SettingFacility(o["facility"], d, "facility")); err != nil {
		if vv, ok := fortiAPIPatch(o["facility"], "LogSyslogd2Setting-Facility"); ok {
			if err = d.Set("facility", vv); err != nil {
				return fmt.Errorf("Error reading facility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading facility: %v", err)
		}
	}

	if err = d.Set("format", flattenLogSyslogd2SettingFormat(o["format"], d, "format")); err != nil {
		if vv, ok := fortiAPIPatch(o["format"], "LogSyslogd2Setting-Format"); ok {
			if err = d.Set("format", vv); err != nil {
				return fmt.Errorf("Error reading format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("interface", flattenLogSyslogd2SettingInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "LogSyslogd2Setting-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenLogSyslogd2SettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "LogSyslogd2Setting-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("max_log_rate", flattenLogSyslogd2SettingMaxLogRate(o["max-log-rate"], d, "max_log_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-rate"], "LogSyslogd2Setting-MaxLogRate"); ok {
			if err = d.Set("max_log_rate", vv); err != nil {
				return fmt.Errorf("Error reading max_log_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_rate: %v", err)
		}
	}

	if err = d.Set("mode", flattenLogSyslogd2SettingMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "LogSyslogd2Setting-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("port", flattenLogSyslogd2SettingPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "LogSyslogd2Setting-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("priority", flattenLogSyslogd2SettingPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "LogSyslogd2Setting-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("server", flattenLogSyslogd2SettingServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "LogSyslogd2Setting-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogSyslogd2SettingSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "LogSyslogd2Setting-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenLogSyslogd2SettingSourceIpInterface(o["source-ip-interface"], d, "source_ip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-interface"], "LogSyslogd2Setting-SourceIpInterface"); ok {
			if err = d.Set("source_ip_interface", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenLogSyslogd2SettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "LogSyslogd2Setting-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("status", flattenLogSyslogd2SettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "LogSyslogd2Setting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenLogSyslogd2SettingVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "LogSyslogd2Setting-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogd2SettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogSyslogd2SettingCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogSyslogd2SettingCustomFieldName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["custom"], _ = expandLogSyslogd2SettingCustomFieldNameCustom(d, i["custom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandLogSyslogd2SettingCustomFieldNameId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandLogSyslogd2SettingCustomFieldNameName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandLogSyslogd2SettingCustomFieldNameCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingCustomFieldNameId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingCustomFieldNameName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingFacility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogSyslogd2SettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingMaxLogRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingSourceIpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogSyslogd2SettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2SettingVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogd2Setting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandLogSyslogd2SettingCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("custom_field_name"); ok || d.HasChange("custom_field_name") {
		t, err := expandLogSyslogd2SettingCustomFieldName(d, v, "custom_field_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-field-name"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok || d.HasChange("enc_algorithm") {
		t, err := expandLogSyslogd2SettingEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("facility"); ok || d.HasChange("facility") {
		t, err := expandLogSyslogd2SettingFacility(d, v, "facility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["facility"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok || d.HasChange("format") {
		t, err := expandLogSyslogd2SettingFormat(d, v, "format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandLogSyslogd2SettingInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandLogSyslogd2SettingInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("max_log_rate"); ok || d.HasChange("max_log_rate") {
		t, err := expandLogSyslogd2SettingMaxLogRate(d, v, "max_log_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-rate"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandLogSyslogd2SettingMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandLogSyslogd2SettingPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandLogSyslogd2SettingPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandLogSyslogd2SettingServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandLogSyslogd2SettingSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok || d.HasChange("source_ip_interface") {
		t, err := expandLogSyslogd2SettingSourceIpInterface(d, v, "source_ip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-interface"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok || d.HasChange("ssl_min_proto_version") {
		t, err := expandLogSyslogd2SettingSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandLogSyslogd2SettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandLogSyslogd2SettingVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
