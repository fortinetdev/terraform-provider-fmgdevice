// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Global FortiAnalyzer Cloud settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogFortianalyzerCloudSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzerCloudSettingUpdate,
		Read:   resourceLogFortianalyzerCloudSettingRead,
		Update: resourceLogFortianalyzerCloudSettingUpdate,
		Delete: resourceLogFortianalyzerCloudSettingDelete,

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
			"__change_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"access_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"certificate_verification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conn_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hmac_algorithm": &schema.Schema{
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
			},
			"ips_archive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_log_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"monitor_failure_retry_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"monitor_keepalive_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"preshared_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Computed: true,
			},
			"upload_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogFortianalyzerCloudSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogFortianalyzerCloudSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerCloudSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortianalyzerCloudSetting(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerCloudSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogFortianalyzerCloudSetting")

	return resourceLogFortianalyzerCloudSettingRead(d, m)
}

func resourceLogFortianalyzerCloudSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogFortianalyzerCloudSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting LogFortianalyzerCloudSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzerCloudSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogFortianalyzerCloudSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzerCloudSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzerCloudSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzerCloudSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzerCloudSettingChangeIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingAccessConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogFortianalyzerCloudSettingCertificateVerification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingConnTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingHmacAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogFortianalyzerCloudSettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingIpsArchive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingMaxLogRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingMonitorFailureRetryPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingMonitorKeepalivePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingPresharedKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogFortianalyzerCloudSettingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingUploadDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingUploadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingUploadOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudSettingUploadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogFortianalyzerCloudSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("__change_ip", flattenLogFortianalyzerCloudSettingChangeIp(o["__change_ip"], d, "__change_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["__change_ip"], "LogFortianalyzerCloudSetting-ChangeIp"); ok {
			if err = d.Set("__change_ip", vv); err != nil {
				return fmt.Errorf("Error reading __change_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading __change_ip: %v", err)
		}
	}

	if err = d.Set("access_config", flattenLogFortianalyzerCloudSettingAccessConfig(o["access-config"], d, "access_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-config"], "LogFortianalyzerCloudSetting-AccessConfig"); ok {
			if err = d.Set("access_config", vv); err != nil {
				return fmt.Errorf("Error reading access_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_config: %v", err)
		}
	}

	if err = d.Set("certificate", flattenLogFortianalyzerCloudSettingCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "LogFortianalyzerCloudSetting-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("certificate_verification", flattenLogFortianalyzerCloudSettingCertificateVerification(o["certificate-verification"], d, "certificate_verification")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate-verification"], "LogFortianalyzerCloudSetting-CertificateVerification"); ok {
			if err = d.Set("certificate_verification", vv); err != nil {
				return fmt.Errorf("Error reading certificate_verification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate_verification: %v", err)
		}
	}

	if err = d.Set("conn_timeout", flattenLogFortianalyzerCloudSettingConnTimeout(o["conn-timeout"], d, "conn_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn-timeout"], "LogFortianalyzerCloudSetting-ConnTimeout"); ok {
			if err = d.Set("conn_timeout", vv); err != nil {
				return fmt.Errorf("Error reading conn_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_timeout: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogFortianalyzerCloudSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["enc-algorithm"], "LogFortianalyzerCloudSetting-EncAlgorithm"); ok {
			if err = d.Set("enc_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading enc_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("hmac_algorithm", flattenLogFortianalyzerCloudSettingHmacAlgorithm(o["hmac-algorithm"], d, "hmac_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["hmac-algorithm"], "LogFortianalyzerCloudSetting-HmacAlgorithm"); ok {
			if err = d.Set("hmac_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading hmac_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hmac_algorithm: %v", err)
		}
	}

	if err = d.Set("interface", flattenLogFortianalyzerCloudSettingInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "LogFortianalyzerCloudSetting-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenLogFortianalyzerCloudSettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "LogFortianalyzerCloudSetting-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ips_archive", flattenLogFortianalyzerCloudSettingIpsArchive(o["ips-archive"], d, "ips_archive")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-archive"], "LogFortianalyzerCloudSetting-IpsArchive"); ok {
			if err = d.Set("ips_archive", vv); err != nil {
				return fmt.Errorf("Error reading ips_archive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_archive: %v", err)
		}
	}

	if err = d.Set("max_log_rate", flattenLogFortianalyzerCloudSettingMaxLogRate(o["max-log-rate"], d, "max_log_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-rate"], "LogFortianalyzerCloudSetting-MaxLogRate"); ok {
			if err = d.Set("max_log_rate", vv); err != nil {
				return fmt.Errorf("Error reading max_log_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_rate: %v", err)
		}
	}

	if err = d.Set("monitor_failure_retry_period", flattenLogFortianalyzerCloudSettingMonitorFailureRetryPeriod(o["monitor-failure-retry-period"], d, "monitor_failure_retry_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-failure-retry-period"], "LogFortianalyzerCloudSetting-MonitorFailureRetryPeriod"); ok {
			if err = d.Set("monitor_failure_retry_period", vv); err != nil {
				return fmt.Errorf("Error reading monitor_failure_retry_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_failure_retry_period: %v", err)
		}
	}

	if err = d.Set("monitor_keepalive_period", flattenLogFortianalyzerCloudSettingMonitorKeepalivePeriod(o["monitor-keepalive-period"], d, "monitor_keepalive_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-keepalive-period"], "LogFortianalyzerCloudSetting-MonitorKeepalivePeriod"); ok {
			if err = d.Set("monitor_keepalive_period", vv); err != nil {
				return fmt.Errorf("Error reading monitor_keepalive_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_keepalive_period: %v", err)
		}
	}

	if err = d.Set("preshared_key", flattenLogFortianalyzerCloudSettingPresharedKey(o["preshared-key"], d, "preshared_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["preshared-key"], "LogFortianalyzerCloudSetting-PresharedKey"); ok {
			if err = d.Set("preshared_key", vv); err != nil {
				return fmt.Errorf("Error reading preshared_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preshared_key: %v", err)
		}
	}

	if err = d.Set("priority", flattenLogFortianalyzerCloudSettingPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "LogFortianalyzerCloudSetting-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("serial", flattenLogFortianalyzerCloudSettingSerial(o["serial"], d, "serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial"], "LogFortianalyzerCloudSetting-Serial"); ok {
			if err = d.Set("serial", vv); err != nil {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogFortianalyzerCloudSettingSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "LogFortianalyzerCloudSetting-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenLogFortianalyzerCloudSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "LogFortianalyzerCloudSetting-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("status", flattenLogFortianalyzerCloudSettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "LogFortianalyzerCloudSetting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upload_day", flattenLogFortianalyzerCloudSettingUploadDay(o["upload-day"], d, "upload_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-day"], "LogFortianalyzerCloudSetting-UploadDay"); ok {
			if err = d.Set("upload_day", vv); err != nil {
				return fmt.Errorf("Error reading upload_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_day: %v", err)
		}
	}

	if err = d.Set("upload_interval", flattenLogFortianalyzerCloudSettingUploadInterval(o["upload-interval"], d, "upload_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-interval"], "LogFortianalyzerCloudSetting-UploadInterval"); ok {
			if err = d.Set("upload_interval", vv); err != nil {
				return fmt.Errorf("Error reading upload_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_interval: %v", err)
		}
	}

	if err = d.Set("upload_option", flattenLogFortianalyzerCloudSettingUploadOption(o["upload-option"], d, "upload_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-option"], "LogFortianalyzerCloudSetting-UploadOption"); ok {
			if err = d.Set("upload_option", vv); err != nil {
				return fmt.Errorf("Error reading upload_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_option: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenLogFortianalyzerCloudSettingUploadTime(o["upload-time"], d, "upload_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-time"], "LogFortianalyzerCloudSetting-UploadTime"); ok {
			if err = d.Set("upload_time", vv); err != nil {
				return fmt.Errorf("Error reading upload_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzerCloudSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogFortianalyzerCloudSettingChangeIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingAccessConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogFortianalyzerCloudSettingCertificateVerification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingConnTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingHmacAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogFortianalyzerCloudSettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingIpsArchive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingMaxLogRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingMonitorFailureRetryPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingMonitorKeepalivePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingPresharedKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogFortianalyzerCloudSettingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingUploadDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingUploadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingUploadOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudSettingUploadTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzerCloudSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("__change_ip"); ok || d.HasChange("__change_ip") {
		t, err := expandLogFortianalyzerCloudSettingChangeIp(d, v, "__change_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["__change_ip"] = t
		}
	}

	if v, ok := d.GetOk("access_config"); ok || d.HasChange("access_config") {
		t, err := expandLogFortianalyzerCloudSettingAccessConfig(d, v, "access_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-config"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandLogFortianalyzerCloudSettingCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("certificate_verification"); ok || d.HasChange("certificate_verification") {
		t, err := expandLogFortianalyzerCloudSettingCertificateVerification(d, v, "certificate_verification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate-verification"] = t
		}
	}

	if v, ok := d.GetOk("conn_timeout"); ok || d.HasChange("conn_timeout") {
		t, err := expandLogFortianalyzerCloudSettingConnTimeout(d, v, "conn_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-timeout"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok || d.HasChange("enc_algorithm") {
		t, err := expandLogFortianalyzerCloudSettingEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("hmac_algorithm"); ok || d.HasChange("hmac_algorithm") {
		t, err := expandLogFortianalyzerCloudSettingHmacAlgorithm(d, v, "hmac_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hmac-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandLogFortianalyzerCloudSettingInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandLogFortianalyzerCloudSettingInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ips_archive"); ok || d.HasChange("ips_archive") {
		t, err := expandLogFortianalyzerCloudSettingIpsArchive(d, v, "ips_archive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-archive"] = t
		}
	}

	if v, ok := d.GetOk("max_log_rate"); ok || d.HasChange("max_log_rate") {
		t, err := expandLogFortianalyzerCloudSettingMaxLogRate(d, v, "max_log_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-rate"] = t
		}
	}

	if v, ok := d.GetOk("monitor_failure_retry_period"); ok || d.HasChange("monitor_failure_retry_period") {
		t, err := expandLogFortianalyzerCloudSettingMonitorFailureRetryPeriod(d, v, "monitor_failure_retry_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-failure-retry-period"] = t
		}
	}

	if v, ok := d.GetOk("monitor_keepalive_period"); ok || d.HasChange("monitor_keepalive_period") {
		t, err := expandLogFortianalyzerCloudSettingMonitorKeepalivePeriod(d, v, "monitor_keepalive_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-keepalive-period"] = t
		}
	}

	if v, ok := d.GetOk("preshared_key"); ok || d.HasChange("preshared_key") {
		t, err := expandLogFortianalyzerCloudSettingPresharedKey(d, v, "preshared_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preshared-key"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandLogFortianalyzerCloudSettingPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		t, err := expandLogFortianalyzerCloudSettingSerial(d, v, "serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandLogFortianalyzerCloudSettingSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok || d.HasChange("ssl_min_proto_version") {
		t, err := expandLogFortianalyzerCloudSettingSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandLogFortianalyzerCloudSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("upload_day"); ok || d.HasChange("upload_day") {
		t, err := expandLogFortianalyzerCloudSettingUploadDay(d, v, "upload_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-day"] = t
		}
	}

	if v, ok := d.GetOk("upload_interval"); ok || d.HasChange("upload_interval") {
		t, err := expandLogFortianalyzerCloudSettingUploadInterval(d, v, "upload_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-interval"] = t
		}
	}

	if v, ok := d.GetOk("upload_option"); ok || d.HasChange("upload_option") {
		t, err := expandLogFortianalyzerCloudSettingUploadOption(d, v, "upload_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-option"] = t
		}
	}

	if v, ok := d.GetOk("upload_time"); ok || d.HasChange("upload_time") {
		t, err := expandLogFortianalyzerCloudSettingUploadTime(d, v, "upload_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-time"] = t
		}
	}

	return &obj, nil
}
