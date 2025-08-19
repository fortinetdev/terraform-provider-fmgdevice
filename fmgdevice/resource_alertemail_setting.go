// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure alert email settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAlertemailSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertemailSettingUpdate,
		Read:   resourceAlertemailSettingRead,
		Update: resourceAlertemailSettingUpdate,
		Delete: resourceAlertemailSettingDelete,

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
			"fds_license_expiring_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fds_license_expiring_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fds_update_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fips_cc_errors": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsso_disconnect_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_errors_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppp_errors_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_login_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alert_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"amc_interface_bypass_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antivirus_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"configuration_changes_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"critical_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"debug_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"email_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"emergency_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"error_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"filter_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firewall_authentication_failure_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiguard_log_quota_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"information_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"local_disk_usage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_disk_usage_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mailto1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mailto2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mailto3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"notification_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_authentication_errors_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"violation_traffic_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"warning_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"webfilter_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAlertemailSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectAlertemailSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating AlertemailSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateAlertemailSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AlertemailSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("AlertemailSetting")

	return resourceAlertemailSettingRead(d, m)
}

func resourceAlertemailSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteAlertemailSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AlertemailSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAlertemailSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAlertemailSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading AlertemailSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAlertemailSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AlertemailSetting resource from API: %v", err)
	}
	return nil
}

func flattenAlertemailSettingFdsLicenseExpiringDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingFdsLicenseExpiringWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingFdsUpdateLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingFipsCcErrors(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingFssoDisconnectLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingHaLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingIpsLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingIpsecErrorsLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingPppErrorsLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingAdminLoginLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingAlertInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingAmcInterfaceBypassMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingAntivirusLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingConfigurationChangesLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingCriticalInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingDebugInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingEmailInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingEmergencyInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingErrorInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingFilterMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingFirewallAuthenticationFailureLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingFortiguardLogQuotaWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingInformationInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingLocalDiskUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingLogDiskUsageWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingMailto1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingMailto2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingMailto3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingNotificationInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingSshLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingSslvpnAuthenticationErrorsLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingViolationTrafficLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingWarningInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAlertemailSettingWebfilterLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAlertemailSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fds_license_expiring_days", flattenAlertemailSettingFdsLicenseExpiringDays(o["FDS-license-expiring-days"], d, "fds_license_expiring_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["FDS-license-expiring-days"], "AlertemailSetting-FdsLicenseExpiringDays"); ok {
			if err = d.Set("fds_license_expiring_days", vv); err != nil {
				return fmt.Errorf("Error reading fds_license_expiring_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_license_expiring_days: %v", err)
		}
	}

	if err = d.Set("fds_license_expiring_warning", flattenAlertemailSettingFdsLicenseExpiringWarning(o["FDS-license-expiring-warning"], d, "fds_license_expiring_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["FDS-license-expiring-warning"], "AlertemailSetting-FdsLicenseExpiringWarning"); ok {
			if err = d.Set("fds_license_expiring_warning", vv); err != nil {
				return fmt.Errorf("Error reading fds_license_expiring_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_license_expiring_warning: %v", err)
		}
	}

	if err = d.Set("fds_update_logs", flattenAlertemailSettingFdsUpdateLogs(o["FDS-update-logs"], d, "fds_update_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["FDS-update-logs"], "AlertemailSetting-FdsUpdateLogs"); ok {
			if err = d.Set("fds_update_logs", vv); err != nil {
				return fmt.Errorf("Error reading fds_update_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_update_logs: %v", err)
		}
	}

	if err = d.Set("fips_cc_errors", flattenAlertemailSettingFipsCcErrors(o["FIPS-CC-errors"], d, "fips_cc_errors")); err != nil {
		if vv, ok := fortiAPIPatch(o["FIPS-CC-errors"], "AlertemailSetting-FipsCcErrors"); ok {
			if err = d.Set("fips_cc_errors", vv); err != nil {
				return fmt.Errorf("Error reading fips_cc_errors: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fips_cc_errors: %v", err)
		}
	}

	if err = d.Set("fsso_disconnect_logs", flattenAlertemailSettingFssoDisconnectLogs(o["FSSO-disconnect-logs"], d, "fsso_disconnect_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["FSSO-disconnect-logs"], "AlertemailSetting-FssoDisconnectLogs"); ok {
			if err = d.Set("fsso_disconnect_logs", vv); err != nil {
				return fmt.Errorf("Error reading fsso_disconnect_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_disconnect_logs: %v", err)
		}
	}

	if err = d.Set("ha_logs", flattenAlertemailSettingHaLogs(o["HA-logs"], d, "ha_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["HA-logs"], "AlertemailSetting-HaLogs"); ok {
			if err = d.Set("ha_logs", vv); err != nil {
				return fmt.Errorf("Error reading ha_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_logs: %v", err)
		}
	}

	if err = d.Set("ips_logs", flattenAlertemailSettingIpsLogs(o["IPS-logs"], d, "ips_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["IPS-logs"], "AlertemailSetting-IpsLogs"); ok {
			if err = d.Set("ips_logs", vv); err != nil {
				return fmt.Errorf("Error reading ips_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_logs: %v", err)
		}
	}

	if err = d.Set("ipsec_errors_logs", flattenAlertemailSettingIpsecErrorsLogs(o["IPsec-errors-logs"], d, "ipsec_errors_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["IPsec-errors-logs"], "AlertemailSetting-IpsecErrorsLogs"); ok {
			if err = d.Set("ipsec_errors_logs", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_errors_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_errors_logs: %v", err)
		}
	}

	if err = d.Set("ppp_errors_logs", flattenAlertemailSettingPppErrorsLogs(o["PPP-errors-logs"], d, "ppp_errors_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["PPP-errors-logs"], "AlertemailSetting-PppErrorsLogs"); ok {
			if err = d.Set("ppp_errors_logs", vv); err != nil {
				return fmt.Errorf("Error reading ppp_errors_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppp_errors_logs: %v", err)
		}
	}

	if err = d.Set("admin_login_logs", flattenAlertemailSettingAdminLoginLogs(o["admin-login-logs"], d, "admin_login_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-login-logs"], "AlertemailSetting-AdminLoginLogs"); ok {
			if err = d.Set("admin_login_logs", vv); err != nil {
				return fmt.Errorf("Error reading admin_login_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_login_logs: %v", err)
		}
	}

	if err = d.Set("alert_interval", flattenAlertemailSettingAlertInterval(o["alert-interval"], d, "alert_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["alert-interval"], "AlertemailSetting-AlertInterval"); ok {
			if err = d.Set("alert_interval", vv); err != nil {
				return fmt.Errorf("Error reading alert_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alert_interval: %v", err)
		}
	}

	if err = d.Set("amc_interface_bypass_mode", flattenAlertemailSettingAmcInterfaceBypassMode(o["amc-interface-bypass-mode"], d, "amc_interface_bypass_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["amc-interface-bypass-mode"], "AlertemailSetting-AmcInterfaceBypassMode"); ok {
			if err = d.Set("amc_interface_bypass_mode", vv); err != nil {
				return fmt.Errorf("Error reading amc_interface_bypass_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading amc_interface_bypass_mode: %v", err)
		}
	}

	if err = d.Set("antivirus_logs", flattenAlertemailSettingAntivirusLogs(o["antivirus-logs"], d, "antivirus_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["antivirus-logs"], "AlertemailSetting-AntivirusLogs"); ok {
			if err = d.Set("antivirus_logs", vv); err != nil {
				return fmt.Errorf("Error reading antivirus_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antivirus_logs: %v", err)
		}
	}

	if err = d.Set("configuration_changes_logs", flattenAlertemailSettingConfigurationChangesLogs(o["configuration-changes-logs"], d, "configuration_changes_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["configuration-changes-logs"], "AlertemailSetting-ConfigurationChangesLogs"); ok {
			if err = d.Set("configuration_changes_logs", vv); err != nil {
				return fmt.Errorf("Error reading configuration_changes_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading configuration_changes_logs: %v", err)
		}
	}

	if err = d.Set("critical_interval", flattenAlertemailSettingCriticalInterval(o["critical-interval"], d, "critical_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["critical-interval"], "AlertemailSetting-CriticalInterval"); ok {
			if err = d.Set("critical_interval", vv); err != nil {
				return fmt.Errorf("Error reading critical_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading critical_interval: %v", err)
		}
	}

	if err = d.Set("debug_interval", flattenAlertemailSettingDebugInterval(o["debug-interval"], d, "debug_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["debug-interval"], "AlertemailSetting-DebugInterval"); ok {
			if err = d.Set("debug_interval", vv); err != nil {
				return fmt.Errorf("Error reading debug_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading debug_interval: %v", err)
		}
	}

	if err = d.Set("email_interval", flattenAlertemailSettingEmailInterval(o["email-interval"], d, "email_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-interval"], "AlertemailSetting-EmailInterval"); ok {
			if err = d.Set("email_interval", vv); err != nil {
				return fmt.Errorf("Error reading email_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_interval: %v", err)
		}
	}

	if err = d.Set("emergency_interval", flattenAlertemailSettingEmergencyInterval(o["emergency-interval"], d, "emergency_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["emergency-interval"], "AlertemailSetting-EmergencyInterval"); ok {
			if err = d.Set("emergency_interval", vv); err != nil {
				return fmt.Errorf("Error reading emergency_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emergency_interval: %v", err)
		}
	}

	if err = d.Set("error_interval", flattenAlertemailSettingErrorInterval(o["error-interval"], d, "error_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["error-interval"], "AlertemailSetting-ErrorInterval"); ok {
			if err = d.Set("error_interval", vv); err != nil {
				return fmt.Errorf("Error reading error_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading error_interval: %v", err)
		}
	}

	if err = d.Set("filter_mode", flattenAlertemailSettingFilterMode(o["filter-mode"], d, "filter_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-mode"], "AlertemailSetting-FilterMode"); ok {
			if err = d.Set("filter_mode", vv); err != nil {
				return fmt.Errorf("Error reading filter_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_mode: %v", err)
		}
	}

	if err = d.Set("firewall_authentication_failure_logs", flattenAlertemailSettingFirewallAuthenticationFailureLogs(o["firewall-authentication-failure-logs"], d, "firewall_authentication_failure_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["firewall-authentication-failure-logs"], "AlertemailSetting-FirewallAuthenticationFailureLogs"); ok {
			if err = d.Set("firewall_authentication_failure_logs", vv); err != nil {
				return fmt.Errorf("Error reading firewall_authentication_failure_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firewall_authentication_failure_logs: %v", err)
		}
	}

	if err = d.Set("fortiguard_log_quota_warning", flattenAlertemailSettingFortiguardLogQuotaWarning(o["fortiguard-log-quota-warning"], d, "fortiguard_log_quota_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-log-quota-warning"], "AlertemailSetting-FortiguardLogQuotaWarning"); ok {
			if err = d.Set("fortiguard_log_quota_warning", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_log_quota_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_log_quota_warning: %v", err)
		}
	}

	if err = d.Set("information_interval", flattenAlertemailSettingInformationInterval(o["information-interval"], d, "information_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["information-interval"], "AlertemailSetting-InformationInterval"); ok {
			if err = d.Set("information_interval", vv); err != nil {
				return fmt.Errorf("Error reading information_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading information_interval: %v", err)
		}
	}

	if err = d.Set("local_disk_usage", flattenAlertemailSettingLocalDiskUsage(o["local-disk-usage"], d, "local_disk_usage")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-disk-usage"], "AlertemailSetting-LocalDiskUsage"); ok {
			if err = d.Set("local_disk_usage", vv); err != nil {
				return fmt.Errorf("Error reading local_disk_usage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_disk_usage: %v", err)
		}
	}

	if err = d.Set("log_disk_usage_warning", flattenAlertemailSettingLogDiskUsageWarning(o["log-disk-usage-warning"], d, "log_disk_usage_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-disk-usage-warning"], "AlertemailSetting-LogDiskUsageWarning"); ok {
			if err = d.Set("log_disk_usage_warning", vv); err != nil {
				return fmt.Errorf("Error reading log_disk_usage_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_disk_usage_warning: %v", err)
		}
	}

	if err = d.Set("mailto1", flattenAlertemailSettingMailto1(o["mailto1"], d, "mailto1")); err != nil {
		if vv, ok := fortiAPIPatch(o["mailto1"], "AlertemailSetting-Mailto1"); ok {
			if err = d.Set("mailto1", vv); err != nil {
				return fmt.Errorf("Error reading mailto1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mailto1: %v", err)
		}
	}

	if err = d.Set("mailto2", flattenAlertemailSettingMailto2(o["mailto2"], d, "mailto2")); err != nil {
		if vv, ok := fortiAPIPatch(o["mailto2"], "AlertemailSetting-Mailto2"); ok {
			if err = d.Set("mailto2", vv); err != nil {
				return fmt.Errorf("Error reading mailto2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mailto2: %v", err)
		}
	}

	if err = d.Set("mailto3", flattenAlertemailSettingMailto3(o["mailto3"], d, "mailto3")); err != nil {
		if vv, ok := fortiAPIPatch(o["mailto3"], "AlertemailSetting-Mailto3"); ok {
			if err = d.Set("mailto3", vv); err != nil {
				return fmt.Errorf("Error reading mailto3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mailto3: %v", err)
		}
	}

	if err = d.Set("notification_interval", flattenAlertemailSettingNotificationInterval(o["notification-interval"], d, "notification_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["notification-interval"], "AlertemailSetting-NotificationInterval"); ok {
			if err = d.Set("notification_interval", vv); err != nil {
				return fmt.Errorf("Error reading notification_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notification_interval: %v", err)
		}
	}

	if err = d.Set("severity", flattenAlertemailSettingSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "AlertemailSetting-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("ssh_logs", flattenAlertemailSettingSshLogs(o["ssh-logs"], d, "ssh_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-logs"], "AlertemailSetting-SshLogs"); ok {
			if err = d.Set("ssh_logs", vv); err != nil {
				return fmt.Errorf("Error reading ssh_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_logs: %v", err)
		}
	}

	if err = d.Set("sslvpn_authentication_errors_logs", flattenAlertemailSettingSslvpnAuthenticationErrorsLogs(o["sslvpn-authentication-errors-logs"], d, "sslvpn_authentication_errors_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn-authentication-errors-logs"], "AlertemailSetting-SslvpnAuthenticationErrorsLogs"); ok {
			if err = d.Set("sslvpn_authentication_errors_logs", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn_authentication_errors_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn_authentication_errors_logs: %v", err)
		}
	}

	if err = d.Set("username", flattenAlertemailSettingUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "AlertemailSetting-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("violation_traffic_logs", flattenAlertemailSettingViolationTrafficLogs(o["violation-traffic-logs"], d, "violation_traffic_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["violation-traffic-logs"], "AlertemailSetting-ViolationTrafficLogs"); ok {
			if err = d.Set("violation_traffic_logs", vv); err != nil {
				return fmt.Errorf("Error reading violation_traffic_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading violation_traffic_logs: %v", err)
		}
	}

	if err = d.Set("warning_interval", flattenAlertemailSettingWarningInterval(o["warning-interval"], d, "warning_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["warning-interval"], "AlertemailSetting-WarningInterval"); ok {
			if err = d.Set("warning_interval", vv); err != nil {
				return fmt.Errorf("Error reading warning_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading warning_interval: %v", err)
		}
	}

	if err = d.Set("webfilter_logs", flattenAlertemailSettingWebfilterLogs(o["webfilter-logs"], d, "webfilter_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-logs"], "AlertemailSetting-WebfilterLogs"); ok {
			if err = d.Set("webfilter_logs", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_logs: %v", err)
		}
	}

	return nil
}

func flattenAlertemailSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAlertemailSettingFdsLicenseExpiringDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFdsLicenseExpiringWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFdsUpdateLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFipsCcErrors(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFssoDisconnectLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingHaLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingIpsLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingIpsecErrorsLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingPppErrorsLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingAdminLoginLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingAlertInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingAmcInterfaceBypassMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingAntivirusLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingConfigurationChangesLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingCriticalInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingDebugInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingEmailInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingEmergencyInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingErrorInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFilterMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFirewallAuthenticationFailureLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFortiguardLogQuotaWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingInformationInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingLocalDiskUsage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingLogDiskUsageWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingMailto1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingMailto2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingMailto3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingNotificationInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingSshLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingSslvpnAuthenticationErrorsLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingViolationTrafficLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingWarningInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingWebfilterLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAlertemailSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fds_license_expiring_days"); ok || d.HasChange("fds_license_expiring_days") {
		t, err := expandAlertemailSettingFdsLicenseExpiringDays(d, v, "fds_license_expiring_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FDS-license-expiring-days"] = t
		}
	}

	if v, ok := d.GetOk("fds_license_expiring_warning"); ok || d.HasChange("fds_license_expiring_warning") {
		t, err := expandAlertemailSettingFdsLicenseExpiringWarning(d, v, "fds_license_expiring_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FDS-license-expiring-warning"] = t
		}
	}

	if v, ok := d.GetOk("fds_update_logs"); ok || d.HasChange("fds_update_logs") {
		t, err := expandAlertemailSettingFdsUpdateLogs(d, v, "fds_update_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FDS-update-logs"] = t
		}
	}

	if v, ok := d.GetOk("fips_cc_errors"); ok || d.HasChange("fips_cc_errors") {
		t, err := expandAlertemailSettingFipsCcErrors(d, v, "fips_cc_errors")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FIPS-CC-errors"] = t
		}
	}

	if v, ok := d.GetOk("fsso_disconnect_logs"); ok || d.HasChange("fsso_disconnect_logs") {
		t, err := expandAlertemailSettingFssoDisconnectLogs(d, v, "fsso_disconnect_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FSSO-disconnect-logs"] = t
		}
	}

	if v, ok := d.GetOk("ha_logs"); ok || d.HasChange("ha_logs") {
		t, err := expandAlertemailSettingHaLogs(d, v, "ha_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["HA-logs"] = t
		}
	}

	if v, ok := d.GetOk("ips_logs"); ok || d.HasChange("ips_logs") {
		t, err := expandAlertemailSettingIpsLogs(d, v, "ips_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["IPS-logs"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_errors_logs"); ok || d.HasChange("ipsec_errors_logs") {
		t, err := expandAlertemailSettingIpsecErrorsLogs(d, v, "ipsec_errors_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["IPsec-errors-logs"] = t
		}
	}

	if v, ok := d.GetOk("ppp_errors_logs"); ok || d.HasChange("ppp_errors_logs") {
		t, err := expandAlertemailSettingPppErrorsLogs(d, v, "ppp_errors_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["PPP-errors-logs"] = t
		}
	}

	if v, ok := d.GetOk("admin_login_logs"); ok || d.HasChange("admin_login_logs") {
		t, err := expandAlertemailSettingAdminLoginLogs(d, v, "admin_login_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-login-logs"] = t
		}
	}

	if v, ok := d.GetOk("alert_interval"); ok || d.HasChange("alert_interval") {
		t, err := expandAlertemailSettingAlertInterval(d, v, "alert_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alert-interval"] = t
		}
	}

	if v, ok := d.GetOk("amc_interface_bypass_mode"); ok || d.HasChange("amc_interface_bypass_mode") {
		t, err := expandAlertemailSettingAmcInterfaceBypassMode(d, v, "amc_interface_bypass_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["amc-interface-bypass-mode"] = t
		}
	}

	if v, ok := d.GetOk("antivirus_logs"); ok || d.HasChange("antivirus_logs") {
		t, err := expandAlertemailSettingAntivirusLogs(d, v, "antivirus_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antivirus-logs"] = t
		}
	}

	if v, ok := d.GetOk("configuration_changes_logs"); ok || d.HasChange("configuration_changes_logs") {
		t, err := expandAlertemailSettingConfigurationChangesLogs(d, v, "configuration_changes_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["configuration-changes-logs"] = t
		}
	}

	if v, ok := d.GetOk("critical_interval"); ok || d.HasChange("critical_interval") {
		t, err := expandAlertemailSettingCriticalInterval(d, v, "critical_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["critical-interval"] = t
		}
	}

	if v, ok := d.GetOk("debug_interval"); ok || d.HasChange("debug_interval") {
		t, err := expandAlertemailSettingDebugInterval(d, v, "debug_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["debug-interval"] = t
		}
	}

	if v, ok := d.GetOk("email_interval"); ok || d.HasChange("email_interval") {
		t, err := expandAlertemailSettingEmailInterval(d, v, "email_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-interval"] = t
		}
	}

	if v, ok := d.GetOk("emergency_interval"); ok || d.HasChange("emergency_interval") {
		t, err := expandAlertemailSettingEmergencyInterval(d, v, "emergency_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emergency-interval"] = t
		}
	}

	if v, ok := d.GetOk("error_interval"); ok || d.HasChange("error_interval") {
		t, err := expandAlertemailSettingErrorInterval(d, v, "error_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["error-interval"] = t
		}
	}

	if v, ok := d.GetOk("filter_mode"); ok || d.HasChange("filter_mode") {
		t, err := expandAlertemailSettingFilterMode(d, v, "filter_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-mode"] = t
		}
	}

	if v, ok := d.GetOk("firewall_authentication_failure_logs"); ok || d.HasChange("firewall_authentication_failure_logs") {
		t, err := expandAlertemailSettingFirewallAuthenticationFailureLogs(d, v, "firewall_authentication_failure_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-authentication-failure-logs"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_log_quota_warning"); ok || d.HasChange("fortiguard_log_quota_warning") {
		t, err := expandAlertemailSettingFortiguardLogQuotaWarning(d, v, "fortiguard_log_quota_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-log-quota-warning"] = t
		}
	}

	if v, ok := d.GetOk("information_interval"); ok || d.HasChange("information_interval") {
		t, err := expandAlertemailSettingInformationInterval(d, v, "information_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["information-interval"] = t
		}
	}

	if v, ok := d.GetOk("local_disk_usage"); ok || d.HasChange("local_disk_usage") {
		t, err := expandAlertemailSettingLocalDiskUsage(d, v, "local_disk_usage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-disk-usage"] = t
		}
	}

	if v, ok := d.GetOk("log_disk_usage_warning"); ok || d.HasChange("log_disk_usage_warning") {
		t, err := expandAlertemailSettingLogDiskUsageWarning(d, v, "log_disk_usage_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-disk-usage-warning"] = t
		}
	}

	if v, ok := d.GetOk("mailto1"); ok || d.HasChange("mailto1") {
		t, err := expandAlertemailSettingMailto1(d, v, "mailto1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mailto1"] = t
		}
	}

	if v, ok := d.GetOk("mailto2"); ok || d.HasChange("mailto2") {
		t, err := expandAlertemailSettingMailto2(d, v, "mailto2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mailto2"] = t
		}
	}

	if v, ok := d.GetOk("mailto3"); ok || d.HasChange("mailto3") {
		t, err := expandAlertemailSettingMailto3(d, v, "mailto3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mailto3"] = t
		}
	}

	if v, ok := d.GetOk("notification_interval"); ok || d.HasChange("notification_interval") {
		t, err := expandAlertemailSettingNotificationInterval(d, v, "notification_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notification-interval"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandAlertemailSettingSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("ssh_logs"); ok || d.HasChange("ssh_logs") {
		t, err := expandAlertemailSettingSshLogs(d, v, "ssh_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-logs"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_authentication_errors_logs"); ok || d.HasChange("sslvpn_authentication_errors_logs") {
		t, err := expandAlertemailSettingSslvpnAuthenticationErrorsLogs(d, v, "sslvpn_authentication_errors_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-authentication-errors-logs"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandAlertemailSettingUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("violation_traffic_logs"); ok || d.HasChange("violation_traffic_logs") {
		t, err := expandAlertemailSettingViolationTrafficLogs(d, v, "violation_traffic_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["violation-traffic-logs"] = t
		}
	}

	if v, ok := d.GetOk("warning_interval"); ok || d.HasChange("warning_interval") {
		t, err := expandAlertemailSettingWarningInterval(d, v, "warning_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warning-interval"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_logs"); ok || d.HasChange("webfilter_logs") {
		t, err := expandAlertemailSettingWebfilterLogs(d, v, "webfilter_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-logs"] = t
		}
	}

	return &obj, nil
}
