// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard services.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFortiguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFortiguardUpdate,
		Read:   resourceSystemFortiguardRead,
		Update: resourceSystemFortiguardUpdate,
		Delete: resourceSystemFortiguardDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"device_name": &schema.Schema{
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
			"antispam_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antispam_cache_mpermille": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"antispam_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"antispam_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"antispam_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antispam_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"antispam_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"anycast_sdns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anycast_sdns_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_firmware_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_firmware_upgrade_day": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auto_firmware_upgrade_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_firmware_upgrade_end_hour": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_firmware_upgrade_start_hour": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_join_forticloud": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_server_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortiguard_anycast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiguard_anycast_source": &schema.Schema{
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
			"load_balance_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_cache_mpermille": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"outbreak_prevention_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"persistent_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"proxy_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"proxy_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sandbox_inline_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sandbox_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sdns_options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sdns_server_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sdns_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"service_account_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscribe_update_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_build_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_dldb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_extdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_ffdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_server_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_uwdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"webfilter_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"webfilter_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"webfilter_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"webfilter_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"antispam_cache_mpercent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dlp_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dlp_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fnbi_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fnbi_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gui_prompt_auto_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ia_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ia_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"outbreak_prevention_cache_mpercent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"videofilter_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"videofilter_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemFortiguardUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemFortiguard(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortiguard resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemFortiguard(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortiguard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemFortiguard")

	return resourceSystemFortiguardRead(d, m)
}

func resourceSystemFortiguardDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteSystemFortiguard(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortiguardRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemFortiguard(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemFortiguard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFortiguard(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortiguard resource from API: %v", err)
	}
	return nil
}

func flattenSystemFortiguardFdsLicenseExpiringDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCacheMpermille(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAnycastSdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAnycastSdnsServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAutoFirmwareUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAutoFirmwareUpgradeDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFortiguardAutoFirmwareUpgradeDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAutoFirmwareUpgradeEndHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAutoFirmwareUpgradeStartHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAutoJoinForticloud(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardDdnsServerIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardDdnsServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardFortiguardAnycast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardFortiguardAnycastSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFortiguardInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardLoadBalanceServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionCacheMpermille(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardPersistentConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardProxyServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardProxyServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardProxyUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardSandboxInlineScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardSandboxRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardSdnsOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFortiguardSdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFortiguardSdnsServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardServiceAccountId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardSubscribeUpdateNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateBuildProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateDldb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateExtdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateFfdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateServerLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateUwdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFortiguardVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCacheMpercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardDlpExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardDlpLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardFnbiExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardFnbiLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardGuiPromptAutoUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardIaExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardIaLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionCacheMpercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardVideofilterExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardVideofilterLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFortiguard(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fds_license_expiring_days", flattenSystemFortiguardFdsLicenseExpiringDays(o["FDS-license-expiring-days"], d, "fds_license_expiring_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["FDS-license-expiring-days"], "SystemFortiguard-FdsLicenseExpiringDays"); ok {
			if err = d.Set("fds_license_expiring_days", vv); err != nil {
				return fmt.Errorf("Error reading fds_license_expiring_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_license_expiring_days: %v", err)
		}
	}

	if err = d.Set("antispam_cache", flattenSystemFortiguardAntispamCache(o["antispam-cache"], d, "antispam_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-cache"], "SystemFortiguard-AntispamCache"); ok {
			if err = d.Set("antispam_cache", vv); err != nil {
				return fmt.Errorf("Error reading antispam_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_cache: %v", err)
		}
	}

	if err = d.Set("antispam_cache_mpermille", flattenSystemFortiguardAntispamCacheMpermille(o["antispam-cache-mpermille"], d, "antispam_cache_mpermille")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-cache-mpermille"], "SystemFortiguard-AntispamCacheMpermille"); ok {
			if err = d.Set("antispam_cache_mpermille", vv); err != nil {
				return fmt.Errorf("Error reading antispam_cache_mpermille: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_cache_mpermille: %v", err)
		}
	}

	if err = d.Set("antispam_cache_ttl", flattenSystemFortiguardAntispamCacheTtl(o["antispam-cache-ttl"], d, "antispam_cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-cache-ttl"], "SystemFortiguard-AntispamCacheTtl"); ok {
			if err = d.Set("antispam_cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading antispam_cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_cache_ttl: %v", err)
		}
	}

	if err = d.Set("antispam_expiration", flattenSystemFortiguardAntispamExpiration(o["antispam-expiration"], d, "antispam_expiration")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-expiration"], "SystemFortiguard-AntispamExpiration"); ok {
			if err = d.Set("antispam_expiration", vv); err != nil {
				return fmt.Errorf("Error reading antispam_expiration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_expiration: %v", err)
		}
	}

	if err = d.Set("antispam_force_off", flattenSystemFortiguardAntispamForceOff(o["antispam-force-off"], d, "antispam_force_off")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-force-off"], "SystemFortiguard-AntispamForceOff"); ok {
			if err = d.Set("antispam_force_off", vv); err != nil {
				return fmt.Errorf("Error reading antispam_force_off: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_force_off: %v", err)
		}
	}

	if err = d.Set("antispam_license", flattenSystemFortiguardAntispamLicense(o["antispam-license"], d, "antispam_license")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-license"], "SystemFortiguard-AntispamLicense"); ok {
			if err = d.Set("antispam_license", vv); err != nil {
				return fmt.Errorf("Error reading antispam_license: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_license: %v", err)
		}
	}

	if err = d.Set("antispam_timeout", flattenSystemFortiguardAntispamTimeout(o["antispam-timeout"], d, "antispam_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-timeout"], "SystemFortiguard-AntispamTimeout"); ok {
			if err = d.Set("antispam_timeout", vv); err != nil {
				return fmt.Errorf("Error reading antispam_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_timeout: %v", err)
		}
	}

	if err = d.Set("anycast_sdns_server_ip", flattenSystemFortiguardAnycastSdnsServerIp(o["anycast-sdns-server-ip"], d, "anycast_sdns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["anycast-sdns-server-ip"], "SystemFortiguard-AnycastSdnsServerIp"); ok {
			if err = d.Set("anycast_sdns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading anycast_sdns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anycast_sdns_server_ip: %v", err)
		}
	}

	if err = d.Set("anycast_sdns_server_port", flattenSystemFortiguardAnycastSdnsServerPort(o["anycast-sdns-server-port"], d, "anycast_sdns_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["anycast-sdns-server-port"], "SystemFortiguard-AnycastSdnsServerPort"); ok {
			if err = d.Set("anycast_sdns_server_port", vv); err != nil {
				return fmt.Errorf("Error reading anycast_sdns_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anycast_sdns_server_port: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade", flattenSystemFortiguardAutoFirmwareUpgrade(o["auto-firmware-upgrade"], d, "auto_firmware_upgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-firmware-upgrade"], "SystemFortiguard-AutoFirmwareUpgrade"); ok {
			if err = d.Set("auto_firmware_upgrade", vv); err != nil {
				return fmt.Errorf("Error reading auto_firmware_upgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_firmware_upgrade: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade_day", flattenSystemFortiguardAutoFirmwareUpgradeDay(o["auto-firmware-upgrade-day"], d, "auto_firmware_upgrade_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-firmware-upgrade-day"], "SystemFortiguard-AutoFirmwareUpgradeDay"); ok {
			if err = d.Set("auto_firmware_upgrade_day", vv); err != nil {
				return fmt.Errorf("Error reading auto_firmware_upgrade_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_firmware_upgrade_day: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade_delay", flattenSystemFortiguardAutoFirmwareUpgradeDelay(o["auto-firmware-upgrade-delay"], d, "auto_firmware_upgrade_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-firmware-upgrade-delay"], "SystemFortiguard-AutoFirmwareUpgradeDelay"); ok {
			if err = d.Set("auto_firmware_upgrade_delay", vv); err != nil {
				return fmt.Errorf("Error reading auto_firmware_upgrade_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_firmware_upgrade_delay: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade_end_hour", flattenSystemFortiguardAutoFirmwareUpgradeEndHour(o["auto-firmware-upgrade-end-hour"], d, "auto_firmware_upgrade_end_hour")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-firmware-upgrade-end-hour"], "SystemFortiguard-AutoFirmwareUpgradeEndHour"); ok {
			if err = d.Set("auto_firmware_upgrade_end_hour", vv); err != nil {
				return fmt.Errorf("Error reading auto_firmware_upgrade_end_hour: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_firmware_upgrade_end_hour: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade_start_hour", flattenSystemFortiguardAutoFirmwareUpgradeStartHour(o["auto-firmware-upgrade-start-hour"], d, "auto_firmware_upgrade_start_hour")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-firmware-upgrade-start-hour"], "SystemFortiguard-AutoFirmwareUpgradeStartHour"); ok {
			if err = d.Set("auto_firmware_upgrade_start_hour", vv); err != nil {
				return fmt.Errorf("Error reading auto_firmware_upgrade_start_hour: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_firmware_upgrade_start_hour: %v", err)
		}
	}

	if err = d.Set("auto_join_forticloud", flattenSystemFortiguardAutoJoinForticloud(o["auto-join-forticloud"], d, "auto_join_forticloud")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-join-forticloud"], "SystemFortiguard-AutoJoinForticloud"); ok {
			if err = d.Set("auto_join_forticloud", vv); err != nil {
				return fmt.Errorf("Error reading auto_join_forticloud: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_join_forticloud: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenSystemFortiguardDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-ip"], "SystemFortiguard-DdnsServerIp"); ok {
			if err = d.Set("ddns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip6", flattenSystemFortiguardDdnsServerIp6(o["ddns-server-ip6"], d, "ddns_server_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-ip6"], "SystemFortiguard-DdnsServerIp6"); ok {
			if err = d.Set("ddns_server_ip6", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_ip6: %v", err)
		}
	}

	if err = d.Set("ddns_server_port", flattenSystemFortiguardDdnsServerPort(o["ddns-server-port"], d, "ddns_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-port"], "SystemFortiguard-DdnsServerPort"); ok {
			if err = d.Set("ddns_server_port", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_port: %v", err)
		}
	}

	if err = d.Set("fortiguard_anycast", flattenSystemFortiguardFortiguardAnycast(o["fortiguard-anycast"], d, "fortiguard_anycast")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-anycast"], "SystemFortiguard-FortiguardAnycast"); ok {
			if err = d.Set("fortiguard_anycast", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_anycast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_anycast: %v", err)
		}
	}

	if err = d.Set("fortiguard_anycast_source", flattenSystemFortiguardFortiguardAnycastSource(o["fortiguard-anycast-source"], d, "fortiguard_anycast_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-anycast-source"], "SystemFortiguard-FortiguardAnycastSource"); ok {
			if err = d.Set("fortiguard_anycast_source", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_anycast_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_anycast_source: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemFortiguardInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemFortiguard-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemFortiguardInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystemFortiguard-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("load_balance_servers", flattenSystemFortiguardLoadBalanceServers(o["load-balance-servers"], d, "load_balance_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balance-servers"], "SystemFortiguard-LoadBalanceServers"); ok {
			if err = d.Set("load_balance_servers", vv); err != nil {
				return fmt.Errorf("Error reading load_balance_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balance_servers: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache", flattenSystemFortiguardOutbreakPreventionCache(o["outbreak-prevention-cache"], d, "outbreak_prevention_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-cache"], "SystemFortiguard-OutbreakPreventionCache"); ok {
			if err = d.Set("outbreak_prevention_cache", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_cache: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache_mpermille", flattenSystemFortiguardOutbreakPreventionCacheMpermille(o["outbreak-prevention-cache-mpermille"], d, "outbreak_prevention_cache_mpermille")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-cache-mpermille"], "SystemFortiguard-OutbreakPreventionCacheMpermille"); ok {
			if err = d.Set("outbreak_prevention_cache_mpermille", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_cache_mpermille: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_cache_mpermille: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache_ttl", flattenSystemFortiguardOutbreakPreventionCacheTtl(o["outbreak-prevention-cache-ttl"], d, "outbreak_prevention_cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-cache-ttl"], "SystemFortiguard-OutbreakPreventionCacheTtl"); ok {
			if err = d.Set("outbreak_prevention_cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_cache_ttl: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_expiration", flattenSystemFortiguardOutbreakPreventionExpiration(o["outbreak-prevention-expiration"], d, "outbreak_prevention_expiration")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-expiration"], "SystemFortiguard-OutbreakPreventionExpiration"); ok {
			if err = d.Set("outbreak_prevention_expiration", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_expiration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_expiration: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_force_off", flattenSystemFortiguardOutbreakPreventionForceOff(o["outbreak-prevention-force-off"], d, "outbreak_prevention_force_off")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-force-off"], "SystemFortiguard-OutbreakPreventionForceOff"); ok {
			if err = d.Set("outbreak_prevention_force_off", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_force_off: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_force_off: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_license", flattenSystemFortiguardOutbreakPreventionLicense(o["outbreak-prevention-license"], d, "outbreak_prevention_license")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-license"], "SystemFortiguard-OutbreakPreventionLicense"); ok {
			if err = d.Set("outbreak_prevention_license", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_license: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_license: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_timeout", flattenSystemFortiguardOutbreakPreventionTimeout(o["outbreak-prevention-timeout"], d, "outbreak_prevention_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-timeout"], "SystemFortiguard-OutbreakPreventionTimeout"); ok {
			if err = d.Set("outbreak_prevention_timeout", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_timeout: %v", err)
		}
	}

	if err = d.Set("persistent_connection", flattenSystemFortiguardPersistentConnection(o["persistent-connection"], d, "persistent_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["persistent-connection"], "SystemFortiguard-PersistentConnection"); ok {
			if err = d.Set("persistent_connection", vv); err != nil {
				return fmt.Errorf("Error reading persistent_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading persistent_connection: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemFortiguardPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemFortiguard-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemFortiguardProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemFortiguard-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("proxy_server_ip", flattenSystemFortiguardProxyServerIp(o["proxy-server-ip"], d, "proxy_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-server-ip"], "SystemFortiguard-ProxyServerIp"); ok {
			if err = d.Set("proxy_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading proxy_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_server_ip: %v", err)
		}
	}

	if err = d.Set("proxy_server_port", flattenSystemFortiguardProxyServerPort(o["proxy-server-port"], d, "proxy_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-server-port"], "SystemFortiguard-ProxyServerPort"); ok {
			if err = d.Set("proxy_server_port", vv); err != nil {
				return fmt.Errorf("Error reading proxy_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_server_port: %v", err)
		}
	}

	if err = d.Set("proxy_username", flattenSystemFortiguardProxyUsername(o["proxy-username"], d, "proxy_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-username"], "SystemFortiguard-ProxyUsername"); ok {
			if err = d.Set("proxy_username", vv); err != nil {
				return fmt.Errorf("Error reading proxy_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_username: %v", err)
		}
	}

	if err = d.Set("sandbox_inline_scan", flattenSystemFortiguardSandboxInlineScan(o["sandbox-inline-scan"], d, "sandbox_inline_scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["sandbox-inline-scan"], "SystemFortiguard-SandboxInlineScan"); ok {
			if err = d.Set("sandbox_inline_scan", vv); err != nil {
				return fmt.Errorf("Error reading sandbox_inline_scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sandbox_inline_scan: %v", err)
		}
	}

	if err = d.Set("sandbox_region", flattenSystemFortiguardSandboxRegion(o["sandbox-region"], d, "sandbox_region")); err != nil {
		if vv, ok := fortiAPIPatch(o["sandbox-region"], "SystemFortiguard-SandboxRegion"); ok {
			if err = d.Set("sandbox_region", vv); err != nil {
				return fmt.Errorf("Error reading sandbox_region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sandbox_region: %v", err)
		}
	}

	if err = d.Set("sdns_options", flattenSystemFortiguardSdnsOptions(o["sdns-options"], d, "sdns_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdns-options"], "SystemFortiguard-SdnsOptions"); ok {
			if err = d.Set("sdns_options", vv); err != nil {
				return fmt.Errorf("Error reading sdns_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdns_options: %v", err)
		}
	}

	if err = d.Set("sdns_server_ip", flattenSystemFortiguardSdnsServerIp(o["sdns-server-ip"], d, "sdns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdns-server-ip"], "SystemFortiguard-SdnsServerIp"); ok {
			if err = d.Set("sdns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading sdns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdns_server_ip: %v", err)
		}
	}

	if err = d.Set("sdns_server_port", flattenSystemFortiguardSdnsServerPort(o["sdns-server-port"], d, "sdns_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdns-server-port"], "SystemFortiguard-SdnsServerPort"); ok {
			if err = d.Set("sdns_server_port", vv); err != nil {
				return fmt.Errorf("Error reading sdns_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdns_server_port: %v", err)
		}
	}

	if err = d.Set("service_account_id", flattenSystemFortiguardServiceAccountId(o["service-account-id"], d, "service_account_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-account-id"], "SystemFortiguard-ServiceAccountId"); ok {
			if err = d.Set("service_account_id", vv); err != nil {
				return fmt.Errorf("Error reading service_account_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_account_id: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemFortiguardSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystemFortiguard-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenSystemFortiguardSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip6"], "SystemFortiguard-SourceIp6"); ok {
			if err = d.Set("source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("subscribe_update_notification", flattenSystemFortiguardSubscribeUpdateNotification(o["subscribe-update-notification"], d, "subscribe_update_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["subscribe-update-notification"], "SystemFortiguard-SubscribeUpdateNotification"); ok {
			if err = d.Set("subscribe_update_notification", vv); err != nil {
				return fmt.Errorf("Error reading subscribe_update_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subscribe_update_notification: %v", err)
		}
	}

	if err = d.Set("update_build_proxy", flattenSystemFortiguardUpdateBuildProxy(o["update-build-proxy"], d, "update_build_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-build-proxy"], "SystemFortiguard-UpdateBuildProxy"); ok {
			if err = d.Set("update_build_proxy", vv); err != nil {
				return fmt.Errorf("Error reading update_build_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_build_proxy: %v", err)
		}
	}

	if err = d.Set("update_dldb", flattenSystemFortiguardUpdateDldb(o["update-dldb"], d, "update_dldb")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-dldb"], "SystemFortiguard-UpdateDldb"); ok {
			if err = d.Set("update_dldb", vv); err != nil {
				return fmt.Errorf("Error reading update_dldb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_dldb: %v", err)
		}
	}

	if err = d.Set("update_extdb", flattenSystemFortiguardUpdateExtdb(o["update-extdb"], d, "update_extdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-extdb"], "SystemFortiguard-UpdateExtdb"); ok {
			if err = d.Set("update_extdb", vv); err != nil {
				return fmt.Errorf("Error reading update_extdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_extdb: %v", err)
		}
	}

	if err = d.Set("update_ffdb", flattenSystemFortiguardUpdateFfdb(o["update-ffdb"], d, "update_ffdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-ffdb"], "SystemFortiguard-UpdateFfdb"); ok {
			if err = d.Set("update_ffdb", vv); err != nil {
				return fmt.Errorf("Error reading update_ffdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_ffdb: %v", err)
		}
	}

	if err = d.Set("update_server_location", flattenSystemFortiguardUpdateServerLocation(o["update-server-location"], d, "update_server_location")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-server-location"], "SystemFortiguard-UpdateServerLocation"); ok {
			if err = d.Set("update_server_location", vv); err != nil {
				return fmt.Errorf("Error reading update_server_location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_server_location: %v", err)
		}
	}

	if err = d.Set("update_uwdb", flattenSystemFortiguardUpdateUwdb(o["update-uwdb"], d, "update_uwdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-uwdb"], "SystemFortiguard-UpdateUwdb"); ok {
			if err = d.Set("update_uwdb", vv); err != nil {
				return fmt.Errorf("Error reading update_uwdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_uwdb: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemFortiguardVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemFortiguard-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenSystemFortiguardVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "SystemFortiguard-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	if err = d.Set("webfilter_cache", flattenSystemFortiguardWebfilterCache(o["webfilter-cache"], d, "webfilter_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-cache"], "SystemFortiguard-WebfilterCache"); ok {
			if err = d.Set("webfilter_cache", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_cache: %v", err)
		}
	}

	if err = d.Set("webfilter_cache_ttl", flattenSystemFortiguardWebfilterCacheTtl(o["webfilter-cache-ttl"], d, "webfilter_cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-cache-ttl"], "SystemFortiguard-WebfilterCacheTtl"); ok {
			if err = d.Set("webfilter_cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_cache_ttl: %v", err)
		}
	}

	if err = d.Set("webfilter_expiration", flattenSystemFortiguardWebfilterExpiration(o["webfilter-expiration"], d, "webfilter_expiration")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-expiration"], "SystemFortiguard-WebfilterExpiration"); ok {
			if err = d.Set("webfilter_expiration", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_expiration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_expiration: %v", err)
		}
	}

	if err = d.Set("webfilter_force_off", flattenSystemFortiguardWebfilterForceOff(o["webfilter-force-off"], d, "webfilter_force_off")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-force-off"], "SystemFortiguard-WebfilterForceOff"); ok {
			if err = d.Set("webfilter_force_off", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_force_off: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_force_off: %v", err)
		}
	}

	if err = d.Set("webfilter_license", flattenSystemFortiguardWebfilterLicense(o["webfilter-license"], d, "webfilter_license")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-license"], "SystemFortiguard-WebfilterLicense"); ok {
			if err = d.Set("webfilter_license", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_license: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_license: %v", err)
		}
	}

	if err = d.Set("webfilter_timeout", flattenSystemFortiguardWebfilterTimeout(o["webfilter-timeout"], d, "webfilter_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-timeout"], "SystemFortiguard-WebfilterTimeout"); ok {
			if err = d.Set("webfilter_timeout", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_timeout: %v", err)
		}
	}

	if err = d.Set("antispam_cache_mpercent", flattenSystemFortiguardAntispamCacheMpercent(o["antispam-cache-mpercent"], d, "antispam_cache_mpercent")); err != nil {
		if vv, ok := fortiAPIPatch(o["antispam-cache-mpercent"], "SystemFortiguard-AntispamCacheMpercent"); ok {
			if err = d.Set("antispam_cache_mpercent", vv); err != nil {
				return fmt.Errorf("Error reading antispam_cache_mpercent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antispam_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("dlp_expiration", flattenSystemFortiguardDlpExpiration(o["dlp-expiration"], d, "dlp_expiration")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-expiration"], "SystemFortiguard-DlpExpiration"); ok {
			if err = d.Set("dlp_expiration", vv); err != nil {
				return fmt.Errorf("Error reading dlp_expiration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_expiration: %v", err)
		}
	}

	if err = d.Set("dlp_license", flattenSystemFortiguardDlpLicense(o["dlp-license"], d, "dlp_license")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-license"], "SystemFortiguard-DlpLicense"); ok {
			if err = d.Set("dlp_license", vv); err != nil {
				return fmt.Errorf("Error reading dlp_license: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_license: %v", err)
		}
	}

	if err = d.Set("fnbi_expiration", flattenSystemFortiguardFnbiExpiration(o["fnbi-expiration"], d, "fnbi_expiration")); err != nil {
		if vv, ok := fortiAPIPatch(o["fnbi-expiration"], "SystemFortiguard-FnbiExpiration"); ok {
			if err = d.Set("fnbi_expiration", vv); err != nil {
				return fmt.Errorf("Error reading fnbi_expiration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fnbi_expiration: %v", err)
		}
	}

	if err = d.Set("fnbi_license", flattenSystemFortiguardFnbiLicense(o["fnbi-license"], d, "fnbi_license")); err != nil {
		if vv, ok := fortiAPIPatch(o["fnbi-license"], "SystemFortiguard-FnbiLicense"); ok {
			if err = d.Set("fnbi_license", vv); err != nil {
				return fmt.Errorf("Error reading fnbi_license: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fnbi_license: %v", err)
		}
	}

	if err = d.Set("gui_prompt_auto_upgrade", flattenSystemFortiguardGuiPromptAutoUpgrade(o["gui-prompt-auto-upgrade"], d, "gui_prompt_auto_upgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-prompt-auto-upgrade"], "SystemFortiguard-GuiPromptAutoUpgrade"); ok {
			if err = d.Set("gui_prompt_auto_upgrade", vv); err != nil {
				return fmt.Errorf("Error reading gui_prompt_auto_upgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_prompt_auto_upgrade: %v", err)
		}
	}

	if err = d.Set("ia_expiration", flattenSystemFortiguardIaExpiration(o["ia-expiration"], d, "ia_expiration")); err != nil {
		if vv, ok := fortiAPIPatch(o["ia-expiration"], "SystemFortiguard-IaExpiration"); ok {
			if err = d.Set("ia_expiration", vv); err != nil {
				return fmt.Errorf("Error reading ia_expiration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ia_expiration: %v", err)
		}
	}

	if err = d.Set("ia_license", flattenSystemFortiguardIaLicense(o["ia-license"], d, "ia_license")); err != nil {
		if vv, ok := fortiAPIPatch(o["ia-license"], "SystemFortiguard-IaLicense"); ok {
			if err = d.Set("ia_license", vv); err != nil {
				return fmt.Errorf("Error reading ia_license: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ia_license: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache_mpercent", flattenSystemFortiguardOutbreakPreventionCacheMpercent(o["outbreak-prevention-cache-mpercent"], d, "outbreak_prevention_cache_mpercent")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-cache-mpercent"], "SystemFortiguard-OutbreakPreventionCacheMpercent"); ok {
			if err = d.Set("outbreak_prevention_cache_mpercent", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_cache_mpercent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("videofilter_expiration", flattenSystemFortiguardVideofilterExpiration(o["videofilter-expiration"], d, "videofilter_expiration")); err != nil {
		if vv, ok := fortiAPIPatch(o["videofilter-expiration"], "SystemFortiguard-VideofilterExpiration"); ok {
			if err = d.Set("videofilter_expiration", vv); err != nil {
				return fmt.Errorf("Error reading videofilter_expiration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading videofilter_expiration: %v", err)
		}
	}

	if err = d.Set("videofilter_license", flattenSystemFortiguardVideofilterLicense(o["videofilter-license"], d, "videofilter_license")); err != nil {
		if vv, ok := fortiAPIPatch(o["videofilter-license"], "SystemFortiguard-VideofilterLicense"); ok {
			if err = d.Set("videofilter_license", vv); err != nil {
				return fmt.Errorf("Error reading videofilter_license: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading videofilter_license: %v", err)
		}
	}

	return nil
}

func flattenSystemFortiguardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFortiguardFdsLicenseExpiringDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCacheMpermille(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamForceOff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAnycastSdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAnycastSdnsServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAutoFirmwareUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAutoFirmwareUpgradeDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFortiguardAutoFirmwareUpgradeDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAutoFirmwareUpgradeEndHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAutoFirmwareUpgradeStartHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAutoJoinForticloud(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardDdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardDdnsServerIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardDdnsServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardFortiguardAnycast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardFortiguardAnycastSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFortiguardInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardLoadBalanceServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionCacheMpermille(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionForceOff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardPersistentConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardProxyPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFortiguardProxyServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardProxyServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardProxyUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSandboxInlineScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSandboxRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSdnsOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFortiguardSdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFortiguardSdnsServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardServiceAccountId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSubscribeUpdateNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateBuildProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateDldb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateExtdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateFfdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateServerLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateUwdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFortiguardVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterForceOff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCacheMpercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardDlpExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardDlpLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardFnbiExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardFnbiLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardGuiPromptAutoUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardIaExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardIaLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionCacheMpercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardVideofilterExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardVideofilterLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFortiguard(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fds_license_expiring_days"); ok || d.HasChange("fds_license_expiring_days") {
		t, err := expandSystemFortiguardFdsLicenseExpiringDays(d, v, "fds_license_expiring_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FDS-license-expiring-days"] = t
		}
	}

	if v, ok := d.GetOk("antispam_cache"); ok || d.HasChange("antispam_cache") {
		t, err := expandSystemFortiguardAntispamCache(d, v, "antispam_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-cache"] = t
		}
	}

	if v, ok := d.GetOk("antispam_cache_mpermille"); ok || d.HasChange("antispam_cache_mpermille") {
		t, err := expandSystemFortiguardAntispamCacheMpermille(d, v, "antispam_cache_mpermille")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-cache-mpermille"] = t
		}
	}

	if v, ok := d.GetOk("antispam_cache_ttl"); ok || d.HasChange("antispam_cache_ttl") {
		t, err := expandSystemFortiguardAntispamCacheTtl(d, v, "antispam_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("antispam_expiration"); ok || d.HasChange("antispam_expiration") {
		t, err := expandSystemFortiguardAntispamExpiration(d, v, "antispam_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-expiration"] = t
		}
	}

	if v, ok := d.GetOk("antispam_force_off"); ok || d.HasChange("antispam_force_off") {
		t, err := expandSystemFortiguardAntispamForceOff(d, v, "antispam_force_off")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-force-off"] = t
		}
	}

	if v, ok := d.GetOk("antispam_license"); ok || d.HasChange("antispam_license") {
		t, err := expandSystemFortiguardAntispamLicense(d, v, "antispam_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-license"] = t
		}
	}

	if v, ok := d.GetOk("antispam_timeout"); ok || d.HasChange("antispam_timeout") {
		t, err := expandSystemFortiguardAntispamTimeout(d, v, "antispam_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-timeout"] = t
		}
	}

	if v, ok := d.GetOk("anycast_sdns_server_ip"); ok || d.HasChange("anycast_sdns_server_ip") {
		t, err := expandSystemFortiguardAnycastSdnsServerIp(d, v, "anycast_sdns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anycast-sdns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("anycast_sdns_server_port"); ok || d.HasChange("anycast_sdns_server_port") {
		t, err := expandSystemFortiguardAnycastSdnsServerPort(d, v, "anycast_sdns_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anycast-sdns-server-port"] = t
		}
	}

	if v, ok := d.GetOk("auto_firmware_upgrade"); ok || d.HasChange("auto_firmware_upgrade") {
		t, err := expandSystemFortiguardAutoFirmwareUpgrade(d, v, "auto_firmware_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-firmware-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("auto_firmware_upgrade_day"); ok || d.HasChange("auto_firmware_upgrade_day") {
		t, err := expandSystemFortiguardAutoFirmwareUpgradeDay(d, v, "auto_firmware_upgrade_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-firmware-upgrade-day"] = t
		}
	}

	if v, ok := d.GetOk("auto_firmware_upgrade_delay"); ok || d.HasChange("auto_firmware_upgrade_delay") {
		t, err := expandSystemFortiguardAutoFirmwareUpgradeDelay(d, v, "auto_firmware_upgrade_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-firmware-upgrade-delay"] = t
		}
	}

	if v, ok := d.GetOk("auto_firmware_upgrade_end_hour"); ok || d.HasChange("auto_firmware_upgrade_end_hour") {
		t, err := expandSystemFortiguardAutoFirmwareUpgradeEndHour(d, v, "auto_firmware_upgrade_end_hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-firmware-upgrade-end-hour"] = t
		}
	}

	if v, ok := d.GetOk("auto_firmware_upgrade_start_hour"); ok || d.HasChange("auto_firmware_upgrade_start_hour") {
		t, err := expandSystemFortiguardAutoFirmwareUpgradeStartHour(d, v, "auto_firmware_upgrade_start_hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-firmware-upgrade-start-hour"] = t
		}
	}

	if v, ok := d.GetOk("auto_join_forticloud"); ok || d.HasChange("auto_join_forticloud") {
		t, err := expandSystemFortiguardAutoJoinForticloud(d, v, "auto_join_forticloud")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-join-forticloud"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok || d.HasChange("ddns_server_ip") {
		t, err := expandSystemFortiguardDdnsServerIp(d, v, "ddns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip6"); ok || d.HasChange("ddns_server_ip6") {
		t, err := expandSystemFortiguardDdnsServerIp6(d, v, "ddns_server_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip6"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_port"); ok || d.HasChange("ddns_server_port") {
		t, err := expandSystemFortiguardDdnsServerPort(d, v, "ddns_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-port"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_anycast"); ok || d.HasChange("fortiguard_anycast") {
		t, err := expandSystemFortiguardFortiguardAnycast(d, v, "fortiguard_anycast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-anycast"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_anycast_source"); ok || d.HasChange("fortiguard_anycast_source") {
		t, err := expandSystemFortiguardFortiguardAnycastSource(d, v, "fortiguard_anycast_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-anycast-source"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemFortiguardInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystemFortiguardInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_servers"); ok || d.HasChange("load_balance_servers") {
		t, err := expandSystemFortiguardLoadBalanceServers(d, v, "load_balance_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-servers"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache"); ok || d.HasChange("outbreak_prevention_cache") {
		t, err := expandSystemFortiguardOutbreakPreventionCache(d, v, "outbreak_prevention_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-cache"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache_mpermille"); ok || d.HasChange("outbreak_prevention_cache_mpermille") {
		t, err := expandSystemFortiguardOutbreakPreventionCacheMpermille(d, v, "outbreak_prevention_cache_mpermille")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-cache-mpermille"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache_ttl"); ok || d.HasChange("outbreak_prevention_cache_ttl") {
		t, err := expandSystemFortiguardOutbreakPreventionCacheTtl(d, v, "outbreak_prevention_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_expiration"); ok || d.HasChange("outbreak_prevention_expiration") {
		t, err := expandSystemFortiguardOutbreakPreventionExpiration(d, v, "outbreak_prevention_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-expiration"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_force_off"); ok || d.HasChange("outbreak_prevention_force_off") {
		t, err := expandSystemFortiguardOutbreakPreventionForceOff(d, v, "outbreak_prevention_force_off")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-force-off"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_license"); ok || d.HasChange("outbreak_prevention_license") {
		t, err := expandSystemFortiguardOutbreakPreventionLicense(d, v, "outbreak_prevention_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-license"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_timeout"); ok || d.HasChange("outbreak_prevention_timeout") {
		t, err := expandSystemFortiguardOutbreakPreventionTimeout(d, v, "outbreak_prevention_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-timeout"] = t
		}
	}

	if v, ok := d.GetOk("persistent_connection"); ok || d.HasChange("persistent_connection") {
		t, err := expandSystemFortiguardPersistentConnection(d, v, "persistent_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistent-connection"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemFortiguardPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemFortiguardProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("proxy_password"); ok || d.HasChange("proxy_password") {
		t, err := expandSystemFortiguardProxyPassword(d, v, "proxy_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-password"] = t
		}
	}

	if v, ok := d.GetOk("proxy_server_ip"); ok || d.HasChange("proxy_server_ip") {
		t, err := expandSystemFortiguardProxyServerIp(d, v, "proxy_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("proxy_server_port"); ok || d.HasChange("proxy_server_port") {
		t, err := expandSystemFortiguardProxyServerPort(d, v, "proxy_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-server-port"] = t
		}
	}

	if v, ok := d.GetOk("proxy_username"); ok || d.HasChange("proxy_username") {
		t, err := expandSystemFortiguardProxyUsername(d, v, "proxy_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-username"] = t
		}
	}

	if v, ok := d.GetOk("sandbox_inline_scan"); ok || d.HasChange("sandbox_inline_scan") {
		t, err := expandSystemFortiguardSandboxInlineScan(d, v, "sandbox_inline_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sandbox-inline-scan"] = t
		}
	}

	if v, ok := d.GetOk("sandbox_region"); ok || d.HasChange("sandbox_region") {
		t, err := expandSystemFortiguardSandboxRegion(d, v, "sandbox_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sandbox-region"] = t
		}
	}

	if v, ok := d.GetOk("sdns_options"); ok || d.HasChange("sdns_options") {
		t, err := expandSystemFortiguardSdnsOptions(d, v, "sdns_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-options"] = t
		}
	}

	if v, ok := d.GetOk("sdns_server_ip"); ok || d.HasChange("sdns_server_ip") {
		t, err := expandSystemFortiguardSdnsServerIp(d, v, "sdns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("sdns_server_port"); ok || d.HasChange("sdns_server_port") {
		t, err := expandSystemFortiguardSdnsServerPort(d, v, "sdns_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-server-port"] = t
		}
	}

	if v, ok := d.GetOk("service_account_id"); ok || d.HasChange("service_account_id") {
		t, err := expandSystemFortiguardServiceAccountId(d, v, "service_account_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-account-id"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystemFortiguardSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok || d.HasChange("source_ip6") {
		t, err := expandSystemFortiguardSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("subscribe_update_notification"); ok || d.HasChange("subscribe_update_notification") {
		t, err := expandSystemFortiguardSubscribeUpdateNotification(d, v, "subscribe_update_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subscribe-update-notification"] = t
		}
	}

	if v, ok := d.GetOk("update_build_proxy"); ok || d.HasChange("update_build_proxy") {
		t, err := expandSystemFortiguardUpdateBuildProxy(d, v, "update_build_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-build-proxy"] = t
		}
	}

	if v, ok := d.GetOk("update_dldb"); ok || d.HasChange("update_dldb") {
		t, err := expandSystemFortiguardUpdateDldb(d, v, "update_dldb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-dldb"] = t
		}
	}

	if v, ok := d.GetOk("update_extdb"); ok || d.HasChange("update_extdb") {
		t, err := expandSystemFortiguardUpdateExtdb(d, v, "update_extdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-extdb"] = t
		}
	}

	if v, ok := d.GetOk("update_ffdb"); ok || d.HasChange("update_ffdb") {
		t, err := expandSystemFortiguardUpdateFfdb(d, v, "update_ffdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-ffdb"] = t
		}
	}

	if v, ok := d.GetOk("update_server_location"); ok || d.HasChange("update_server_location") {
		t, err := expandSystemFortiguardUpdateServerLocation(d, v, "update_server_location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-server-location"] = t
		}
	}

	if v, ok := d.GetOk("update_uwdb"); ok || d.HasChange("update_uwdb") {
		t, err := expandSystemFortiguardUpdateUwdb(d, v, "update_uwdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-uwdb"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemFortiguardVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandSystemFortiguardVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_cache"); ok || d.HasChange("webfilter_cache") {
		t, err := expandSystemFortiguardWebfilterCache(d, v, "webfilter_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-cache"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_cache_ttl"); ok || d.HasChange("webfilter_cache_ttl") {
		t, err := expandSystemFortiguardWebfilterCacheTtl(d, v, "webfilter_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_expiration"); ok || d.HasChange("webfilter_expiration") {
		t, err := expandSystemFortiguardWebfilterExpiration(d, v, "webfilter_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-expiration"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_force_off"); ok || d.HasChange("webfilter_force_off") {
		t, err := expandSystemFortiguardWebfilterForceOff(d, v, "webfilter_force_off")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-force-off"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_license"); ok || d.HasChange("webfilter_license") {
		t, err := expandSystemFortiguardWebfilterLicense(d, v, "webfilter_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-license"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_timeout"); ok || d.HasChange("webfilter_timeout") {
		t, err := expandSystemFortiguardWebfilterTimeout(d, v, "webfilter_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-timeout"] = t
		}
	}

	if v, ok := d.GetOk("antispam_cache_mpercent"); ok || d.HasChange("antispam_cache_mpercent") {
		t, err := expandSystemFortiguardAntispamCacheMpercent(d, v, "antispam_cache_mpercent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-cache-mpercent"] = t
		}
	}

	if v, ok := d.GetOk("dlp_expiration"); ok || d.HasChange("dlp_expiration") {
		t, err := expandSystemFortiguardDlpExpiration(d, v, "dlp_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-expiration"] = t
		}
	}

	if v, ok := d.GetOk("dlp_license"); ok || d.HasChange("dlp_license") {
		t, err := expandSystemFortiguardDlpLicense(d, v, "dlp_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-license"] = t
		}
	}

	if v, ok := d.GetOk("fnbi_expiration"); ok || d.HasChange("fnbi_expiration") {
		t, err := expandSystemFortiguardFnbiExpiration(d, v, "fnbi_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fnbi-expiration"] = t
		}
	}

	if v, ok := d.GetOk("fnbi_license"); ok || d.HasChange("fnbi_license") {
		t, err := expandSystemFortiguardFnbiLicense(d, v, "fnbi_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fnbi-license"] = t
		}
	}

	if v, ok := d.GetOk("gui_prompt_auto_upgrade"); ok || d.HasChange("gui_prompt_auto_upgrade") {
		t, err := expandSystemFortiguardGuiPromptAutoUpgrade(d, v, "gui_prompt_auto_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-prompt-auto-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("ia_expiration"); ok || d.HasChange("ia_expiration") {
		t, err := expandSystemFortiguardIaExpiration(d, v, "ia_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ia-expiration"] = t
		}
	}

	if v, ok := d.GetOk("ia_license"); ok || d.HasChange("ia_license") {
		t, err := expandSystemFortiguardIaLicense(d, v, "ia_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ia-license"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache_mpercent"); ok || d.HasChange("outbreak_prevention_cache_mpercent") {
		t, err := expandSystemFortiguardOutbreakPreventionCacheMpercent(d, v, "outbreak_prevention_cache_mpercent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-cache-mpercent"] = t
		}
	}

	if v, ok := d.GetOk("videofilter_expiration"); ok || d.HasChange("videofilter_expiration") {
		t, err := expandSystemFortiguardVideofilterExpiration(d, v, "videofilter_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["videofilter-expiration"] = t
		}
	}

	if v, ok := d.GetOk("videofilter_license"); ok || d.HasChange("videofilter_license") {
		t, err := expandSystemFortiguardVideofilterLicense(d, v, "videofilter_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["videofilter-license"] = t
		}
	}

	return &obj, nil
}
