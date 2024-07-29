// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Settings for local disk logging.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogDiskSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogDiskSettingUpdate,
		Read:   resourceLogDiskSettingRead,
		Update: resourceLogDiskSettingUpdate,
		Delete: resourceLogDiskSettingDelete,

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
			"diskfull": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dlp_archive_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"full_final_warning_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"full_first_warning_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"full_second_warning_threshold": &schema.Schema{
				Type:     schema.TypeInt,
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
			"ips_archive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_log_file_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_policy_packet_capture_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maximum_log_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"report_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"roll_day": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"roll_schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"roll_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_delete_files": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_destination": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_ssl_conn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uploaddir": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uploadip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uploadpass": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"uploadport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uploadsched": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uploadtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uploadtype": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"uploaduser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceLogDiskSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogDiskSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogDiskSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogDiskSetting(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating LogDiskSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogDiskSetting")

	return resourceLogDiskSettingRead(d, m)
}

func resourceLogDiskSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogDiskSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting LogDiskSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogDiskSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogDiskSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogDiskSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogDiskSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogDiskSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogDiskSettingDiskfull(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingDlpArchiveQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingFullFinalWarningThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingFullFirstWarningThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingFullSecondWarningThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogDiskSettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingIpsArchive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingLogQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingMaxLogFileSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingMaxPolicyPacketCaptureSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingMaximumLogAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingReportQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingRollDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogDiskSettingRollSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingRollTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadDeleteFiles(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadDestination(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadSslConn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploaddir(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadsched(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadtype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogDiskSettingUploaduser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogDiskSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("diskfull", flattenLogDiskSettingDiskfull(o["diskfull"], d, "diskfull")); err != nil {
		if vv, ok := fortiAPIPatch(o["diskfull"], "LogDiskSetting-Diskfull"); ok {
			if err = d.Set("diskfull", vv); err != nil {
				return fmt.Errorf("Error reading diskfull: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diskfull: %v", err)
		}
	}

	if err = d.Set("dlp_archive_quota", flattenLogDiskSettingDlpArchiveQuota(o["dlp-archive-quota"], d, "dlp_archive_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-archive-quota"], "LogDiskSetting-DlpArchiveQuota"); ok {
			if err = d.Set("dlp_archive_quota", vv); err != nil {
				return fmt.Errorf("Error reading dlp_archive_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_archive_quota: %v", err)
		}
	}

	if err = d.Set("full_final_warning_threshold", flattenLogDiskSettingFullFinalWarningThreshold(o["full-final-warning-threshold"], d, "full_final_warning_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["full-final-warning-threshold"], "LogDiskSetting-FullFinalWarningThreshold"); ok {
			if err = d.Set("full_final_warning_threshold", vv); err != nil {
				return fmt.Errorf("Error reading full_final_warning_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading full_final_warning_threshold: %v", err)
		}
	}

	if err = d.Set("full_first_warning_threshold", flattenLogDiskSettingFullFirstWarningThreshold(o["full-first-warning-threshold"], d, "full_first_warning_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["full-first-warning-threshold"], "LogDiskSetting-FullFirstWarningThreshold"); ok {
			if err = d.Set("full_first_warning_threshold", vv); err != nil {
				return fmt.Errorf("Error reading full_first_warning_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading full_first_warning_threshold: %v", err)
		}
	}

	if err = d.Set("full_second_warning_threshold", flattenLogDiskSettingFullSecondWarningThreshold(o["full-second-warning-threshold"], d, "full_second_warning_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["full-second-warning-threshold"], "LogDiskSetting-FullSecondWarningThreshold"); ok {
			if err = d.Set("full_second_warning_threshold", vv); err != nil {
				return fmt.Errorf("Error reading full_second_warning_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading full_second_warning_threshold: %v", err)
		}
	}

	if err = d.Set("interface", flattenLogDiskSettingInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "LogDiskSetting-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenLogDiskSettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "LogDiskSetting-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ips_archive", flattenLogDiskSettingIpsArchive(o["ips-archive"], d, "ips_archive")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-archive"], "LogDiskSetting-IpsArchive"); ok {
			if err = d.Set("ips_archive", vv); err != nil {
				return fmt.Errorf("Error reading ips_archive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_archive: %v", err)
		}
	}

	if err = d.Set("log_quota", flattenLogDiskSettingLogQuota(o["log-quota"], d, "log_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-quota"], "LogDiskSetting-LogQuota"); ok {
			if err = d.Set("log_quota", vv); err != nil {
				return fmt.Errorf("Error reading log_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_quota: %v", err)
		}
	}

	if err = d.Set("max_log_file_size", flattenLogDiskSettingMaxLogFileSize(o["max-log-file-size"], d, "max_log_file_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-file-size"], "LogDiskSetting-MaxLogFileSize"); ok {
			if err = d.Set("max_log_file_size", vv); err != nil {
				return fmt.Errorf("Error reading max_log_file_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_file_size: %v", err)
		}
	}

	if err = d.Set("max_policy_packet_capture_size", flattenLogDiskSettingMaxPolicyPacketCaptureSize(o["max-policy-packet-capture-size"], d, "max_policy_packet_capture_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-policy-packet-capture-size"], "LogDiskSetting-MaxPolicyPacketCaptureSize"); ok {
			if err = d.Set("max_policy_packet_capture_size", vv); err != nil {
				return fmt.Errorf("Error reading max_policy_packet_capture_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_policy_packet_capture_size: %v", err)
		}
	}

	if err = d.Set("maximum_log_age", flattenLogDiskSettingMaximumLogAge(o["maximum-log-age"], d, "maximum_log_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-log-age"], "LogDiskSetting-MaximumLogAge"); ok {
			if err = d.Set("maximum_log_age", vv); err != nil {
				return fmt.Errorf("Error reading maximum_log_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_log_age: %v", err)
		}
	}

	if err = d.Set("report_quota", flattenLogDiskSettingReportQuota(o["report-quota"], d, "report_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["report-quota"], "LogDiskSetting-ReportQuota"); ok {
			if err = d.Set("report_quota", vv); err != nil {
				return fmt.Errorf("Error reading report_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading report_quota: %v", err)
		}
	}

	if err = d.Set("roll_day", flattenLogDiskSettingRollDay(o["roll-day"], d, "roll_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["roll-day"], "LogDiskSetting-RollDay"); ok {
			if err = d.Set("roll_day", vv); err != nil {
				return fmt.Errorf("Error reading roll_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading roll_day: %v", err)
		}
	}

	if err = d.Set("roll_schedule", flattenLogDiskSettingRollSchedule(o["roll-schedule"], d, "roll_schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["roll-schedule"], "LogDiskSetting-RollSchedule"); ok {
			if err = d.Set("roll_schedule", vv); err != nil {
				return fmt.Errorf("Error reading roll_schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading roll_schedule: %v", err)
		}
	}

	if err = d.Set("roll_time", flattenLogDiskSettingRollTime(o["roll-time"], d, "roll_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["roll-time"], "LogDiskSetting-RollTime"); ok {
			if err = d.Set("roll_time", vv); err != nil {
				return fmt.Errorf("Error reading roll_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading roll_time: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogDiskSettingSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "LogDiskSetting-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("status", flattenLogDiskSettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "LogDiskSetting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upload", flattenLogDiskSettingUpload(o["upload"], d, "upload")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload"], "LogDiskSetting-Upload"); ok {
			if err = d.Set("upload", vv); err != nil {
				return fmt.Errorf("Error reading upload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload: %v", err)
		}
	}

	if err = d.Set("upload_delete_files", flattenLogDiskSettingUploadDeleteFiles(o["upload-delete-files"], d, "upload_delete_files")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-delete-files"], "LogDiskSetting-UploadDeleteFiles"); ok {
			if err = d.Set("upload_delete_files", vv); err != nil {
				return fmt.Errorf("Error reading upload_delete_files: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_delete_files: %v", err)
		}
	}

	if err = d.Set("upload_destination", flattenLogDiskSettingUploadDestination(o["upload-destination"], d, "upload_destination")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-destination"], "LogDiskSetting-UploadDestination"); ok {
			if err = d.Set("upload_destination", vv); err != nil {
				return fmt.Errorf("Error reading upload_destination: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_destination: %v", err)
		}
	}

	if err = d.Set("upload_ssl_conn", flattenLogDiskSettingUploadSslConn(o["upload-ssl-conn"], d, "upload_ssl_conn")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-ssl-conn"], "LogDiskSetting-UploadSslConn"); ok {
			if err = d.Set("upload_ssl_conn", vv); err != nil {
				return fmt.Errorf("Error reading upload_ssl_conn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_ssl_conn: %v", err)
		}
	}

	if err = d.Set("uploaddir", flattenLogDiskSettingUploaddir(o["uploaddir"], d, "uploaddir")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploaddir"], "LogDiskSetting-Uploaddir"); ok {
			if err = d.Set("uploaddir", vv); err != nil {
				return fmt.Errorf("Error reading uploaddir: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploaddir: %v", err)
		}
	}

	if err = d.Set("uploadip", flattenLogDiskSettingUploadip(o["uploadip"], d, "uploadip")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploadip"], "LogDiskSetting-Uploadip"); ok {
			if err = d.Set("uploadip", vv); err != nil {
				return fmt.Errorf("Error reading uploadip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploadip: %v", err)
		}
	}

	if err = d.Set("uploadport", flattenLogDiskSettingUploadport(o["uploadport"], d, "uploadport")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploadport"], "LogDiskSetting-Uploadport"); ok {
			if err = d.Set("uploadport", vv); err != nil {
				return fmt.Errorf("Error reading uploadport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploadport: %v", err)
		}
	}

	if err = d.Set("uploadsched", flattenLogDiskSettingUploadsched(o["uploadsched"], d, "uploadsched")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploadsched"], "LogDiskSetting-Uploadsched"); ok {
			if err = d.Set("uploadsched", vv); err != nil {
				return fmt.Errorf("Error reading uploadsched: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploadsched: %v", err)
		}
	}

	if err = d.Set("uploadtime", flattenLogDiskSettingUploadtime(o["uploadtime"], d, "uploadtime")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploadtime"], "LogDiskSetting-Uploadtime"); ok {
			if err = d.Set("uploadtime", vv); err != nil {
				return fmt.Errorf("Error reading uploadtime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploadtime: %v", err)
		}
	}

	if err = d.Set("uploadtype", flattenLogDiskSettingUploadtype(o["uploadtype"], d, "uploadtype")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploadtype"], "LogDiskSetting-Uploadtype"); ok {
			if err = d.Set("uploadtype", vv); err != nil {
				return fmt.Errorf("Error reading uploadtype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploadtype: %v", err)
		}
	}

	if err = d.Set("uploaduser", flattenLogDiskSettingUploaduser(o["uploaduser"], d, "uploaduser")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploaduser"], "LogDiskSetting-Uploaduser"); ok {
			if err = d.Set("uploaduser", vv); err != nil {
				return fmt.Errorf("Error reading uploaduser: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploaduser: %v", err)
		}
	}

	return nil
}

func flattenLogDiskSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogDiskSettingDiskfull(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingDlpArchiveQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingFullFinalWarningThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingFullFirstWarningThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingFullSecondWarningThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogDiskSettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingIpsArchive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingLogQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingMaxLogFileSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingMaxPolicyPacketCaptureSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingMaximumLogAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingReportQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingRollDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogDiskSettingRollSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingRollTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadDeleteFiles(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadDestination(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadSslConn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploaddir(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadpass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogDiskSettingUploadport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadsched(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadtype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogDiskSettingUploaduser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogDiskSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("diskfull"); ok || d.HasChange("diskfull") {
		t, err := expandLogDiskSettingDiskfull(d, v, "diskfull")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diskfull"] = t
		}
	}

	if v, ok := d.GetOk("dlp_archive_quota"); ok || d.HasChange("dlp_archive_quota") {
		t, err := expandLogDiskSettingDlpArchiveQuota(d, v, "dlp_archive_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-archive-quota"] = t
		}
	}

	if v, ok := d.GetOk("full_final_warning_threshold"); ok || d.HasChange("full_final_warning_threshold") {
		t, err := expandLogDiskSettingFullFinalWarningThreshold(d, v, "full_final_warning_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-final-warning-threshold"] = t
		}
	}

	if v, ok := d.GetOk("full_first_warning_threshold"); ok || d.HasChange("full_first_warning_threshold") {
		t, err := expandLogDiskSettingFullFirstWarningThreshold(d, v, "full_first_warning_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-first-warning-threshold"] = t
		}
	}

	if v, ok := d.GetOk("full_second_warning_threshold"); ok || d.HasChange("full_second_warning_threshold") {
		t, err := expandLogDiskSettingFullSecondWarningThreshold(d, v, "full_second_warning_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-second-warning-threshold"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandLogDiskSettingInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandLogDiskSettingInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ips_archive"); ok || d.HasChange("ips_archive") {
		t, err := expandLogDiskSettingIpsArchive(d, v, "ips_archive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-archive"] = t
		}
	}

	if v, ok := d.GetOk("log_quota"); ok || d.HasChange("log_quota") {
		t, err := expandLogDiskSettingLogQuota(d, v, "log_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-quota"] = t
		}
	}

	if v, ok := d.GetOk("max_log_file_size"); ok || d.HasChange("max_log_file_size") {
		t, err := expandLogDiskSettingMaxLogFileSize(d, v, "max_log_file_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-file-size"] = t
		}
	}

	if v, ok := d.GetOk("max_policy_packet_capture_size"); ok || d.HasChange("max_policy_packet_capture_size") {
		t, err := expandLogDiskSettingMaxPolicyPacketCaptureSize(d, v, "max_policy_packet_capture_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-policy-packet-capture-size"] = t
		}
	}

	if v, ok := d.GetOk("maximum_log_age"); ok || d.HasChange("maximum_log_age") {
		t, err := expandLogDiskSettingMaximumLogAge(d, v, "maximum_log_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-log-age"] = t
		}
	}

	if v, ok := d.GetOk("report_quota"); ok || d.HasChange("report_quota") {
		t, err := expandLogDiskSettingReportQuota(d, v, "report_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-quota"] = t
		}
	}

	if v, ok := d.GetOk("roll_day"); ok || d.HasChange("roll_day") {
		t, err := expandLogDiskSettingRollDay(d, v, "roll_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roll-day"] = t
		}
	}

	if v, ok := d.GetOk("roll_schedule"); ok || d.HasChange("roll_schedule") {
		t, err := expandLogDiskSettingRollSchedule(d, v, "roll_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roll-schedule"] = t
		}
	}

	if v, ok := d.GetOk("roll_time"); ok || d.HasChange("roll_time") {
		t, err := expandLogDiskSettingRollTime(d, v, "roll_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roll-time"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandLogDiskSettingSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandLogDiskSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("upload"); ok || d.HasChange("upload") {
		t, err := expandLogDiskSettingUpload(d, v, "upload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload"] = t
		}
	}

	if v, ok := d.GetOk("upload_delete_files"); ok || d.HasChange("upload_delete_files") {
		t, err := expandLogDiskSettingUploadDeleteFiles(d, v, "upload_delete_files")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-delete-files"] = t
		}
	}

	if v, ok := d.GetOk("upload_destination"); ok || d.HasChange("upload_destination") {
		t, err := expandLogDiskSettingUploadDestination(d, v, "upload_destination")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-destination"] = t
		}
	}

	if v, ok := d.GetOk("upload_ssl_conn"); ok || d.HasChange("upload_ssl_conn") {
		t, err := expandLogDiskSettingUploadSslConn(d, v, "upload_ssl_conn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-ssl-conn"] = t
		}
	}

	if v, ok := d.GetOk("uploaddir"); ok || d.HasChange("uploaddir") {
		t, err := expandLogDiskSettingUploaddir(d, v, "uploaddir")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploaddir"] = t
		}
	}

	if v, ok := d.GetOk("uploadip"); ok || d.HasChange("uploadip") {
		t, err := expandLogDiskSettingUploadip(d, v, "uploadip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadip"] = t
		}
	}

	if v, ok := d.GetOk("uploadpass"); ok || d.HasChange("uploadpass") {
		t, err := expandLogDiskSettingUploadpass(d, v, "uploadpass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadpass"] = t
		}
	}

	if v, ok := d.GetOk("uploadport"); ok || d.HasChange("uploadport") {
		t, err := expandLogDiskSettingUploadport(d, v, "uploadport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadport"] = t
		}
	}

	if v, ok := d.GetOk("uploadsched"); ok || d.HasChange("uploadsched") {
		t, err := expandLogDiskSettingUploadsched(d, v, "uploadsched")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadsched"] = t
		}
	}

	if v, ok := d.GetOk("uploadtime"); ok || d.HasChange("uploadtime") {
		t, err := expandLogDiskSettingUploadtime(d, v, "uploadtime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadtime"] = t
		}
	}

	if v, ok := d.GetOk("uploadtype"); ok || d.HasChange("uploadtype") {
		t, err := expandLogDiskSettingUploadtype(d, v, "uploadtype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadtype"] = t
		}
	}

	if v, ok := d.GetOk("uploaduser"); ok || d.HasChange("uploaduser") {
		t, err := expandLogDiskSettingUploaduser(d, v, "uploaduser")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploaduser"] = t
		}
	}

	return &obj, nil
}
