// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure central management.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCentralManagement() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCentralManagementUpdate,
		Read:   resourceSystemCentralManagementRead,
		Update: resourceSystemCentralManagementUpdate,
		Delete: resourceSystemCentralManagementDelete,

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
			"allow_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_push_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_push_firmware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_remote_firmware_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_remote_lte_firmware_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_remote_modem_firmware_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ca_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fmg_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg_source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg_update_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortigate_cloud_sso_default_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"include_default_servers": &schema.Schema{
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
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ltefw_upgrade_frequency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ltefw_upgrade_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modem_upgrade_frequency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"modem_upgrade_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"schedule_config_restore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule_script_restore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"server_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_address6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_elbc_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceSystemCentralManagementUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemCentralManagement(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagement resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCentralManagement(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagement resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemCentralManagement")

	return resourceSystemCentralManagementRead(d, m)
}

func resourceSystemCentralManagementDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemCentralManagement(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCentralManagement resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCentralManagementRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCentralManagement(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemCentralManagement resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCentralManagement(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCentralManagement resource from API: %v", err)
	}
	return nil
}

func flattenSystemCentralManagementAllowMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowPushConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowPushFirmware(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowRemoteFirmwareUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowRemoteLteFirmwareUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowRemoteModemFirmwareUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementFmg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCentralManagementFmgSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementFmgSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementFmgUpdatePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementFortigateCloudSsoDefaultProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCentralManagementIncludeDefaultServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCentralManagementInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementLtefwUpgradeFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementLtefwUpgradeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementModemUpgradeFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementModemUpgradeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementScheduleConfigRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementScheduleScriptRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCentralManagementServerList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {
			v := flattenSystemCentralManagementServerListAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "SystemCentralManagement-ServerList-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := i["fqdn"]; ok {
			v := flattenSystemCentralManagementServerListFqdn(i["fqdn"], d, pre_append)
			tmp["fqdn"] = fortiAPISubPartPatch(v, "SystemCentralManagement-ServerList-Fqdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemCentralManagementServerListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemCentralManagement-ServerList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address"
		if _, ok := i["server-address"]; ok {
			v := flattenSystemCentralManagementServerListServerAddress(i["server-address"], d, pre_append)
			tmp["server_address"] = fortiAPISubPartPatch(v, "SystemCentralManagement-ServerList-ServerAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address6"
		if _, ok := i["server-address6"]; ok {
			v := flattenSystemCentralManagementServerListServerAddress6(i["server-address6"], d, pre_append)
			tmp["server_address6"] = fortiAPISubPartPatch(v, "SystemCentralManagement-ServerList-ServerAddress6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := i["server-type"]; ok {
			v := flattenSystemCentralManagementServerListServerType(i["server-type"], d, pre_append)
			tmp["server_type"] = fortiAPISubPartPatch(v, "SystemCentralManagement-ServerList-ServerType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemCentralManagementServerListAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListServerAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListServerAddress6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCentralManagementType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementUseElbcVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemCentralManagement(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("allow_monitor", flattenSystemCentralManagementAllowMonitor(o["allow-monitor"], d, "allow_monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-monitor"], "SystemCentralManagement-AllowMonitor"); ok {
			if err = d.Set("allow_monitor", vv); err != nil {
				return fmt.Errorf("Error reading allow_monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_monitor: %v", err)
		}
	}

	if err = d.Set("allow_push_configuration", flattenSystemCentralManagementAllowPushConfiguration(o["allow-push-configuration"], d, "allow_push_configuration")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-push-configuration"], "SystemCentralManagement-AllowPushConfiguration"); ok {
			if err = d.Set("allow_push_configuration", vv); err != nil {
				return fmt.Errorf("Error reading allow_push_configuration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_push_configuration: %v", err)
		}
	}

	if err = d.Set("allow_push_firmware", flattenSystemCentralManagementAllowPushFirmware(o["allow-push-firmware"], d, "allow_push_firmware")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-push-firmware"], "SystemCentralManagement-AllowPushFirmware"); ok {
			if err = d.Set("allow_push_firmware", vv); err != nil {
				return fmt.Errorf("Error reading allow_push_firmware: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_push_firmware: %v", err)
		}
	}

	if err = d.Set("allow_remote_firmware_upgrade", flattenSystemCentralManagementAllowRemoteFirmwareUpgrade(o["allow-remote-firmware-upgrade"], d, "allow_remote_firmware_upgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-remote-firmware-upgrade"], "SystemCentralManagement-AllowRemoteFirmwareUpgrade"); ok {
			if err = d.Set("allow_remote_firmware_upgrade", vv); err != nil {
				return fmt.Errorf("Error reading allow_remote_firmware_upgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_remote_firmware_upgrade: %v", err)
		}
	}

	if err = d.Set("allow_remote_lte_firmware_upgrade", flattenSystemCentralManagementAllowRemoteLteFirmwareUpgrade(o["allow-remote-lte-firmware-upgrade"], d, "allow_remote_lte_firmware_upgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-remote-lte-firmware-upgrade"], "SystemCentralManagement-AllowRemoteLteFirmwareUpgrade"); ok {
			if err = d.Set("allow_remote_lte_firmware_upgrade", vv); err != nil {
				return fmt.Errorf("Error reading allow_remote_lte_firmware_upgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_remote_lte_firmware_upgrade: %v", err)
		}
	}

	if err = d.Set("allow_remote_modem_firmware_upgrade", flattenSystemCentralManagementAllowRemoteModemFirmwareUpgrade(o["allow-remote-modem-firmware-upgrade"], d, "allow_remote_modem_firmware_upgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-remote-modem-firmware-upgrade"], "SystemCentralManagement-AllowRemoteModemFirmwareUpgrade"); ok {
			if err = d.Set("allow_remote_modem_firmware_upgrade", vv); err != nil {
				return fmt.Errorf("Error reading allow_remote_modem_firmware_upgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_remote_modem_firmware_upgrade: %v", err)
		}
	}

	if err = d.Set("ca_cert", flattenSystemCentralManagementCaCert(o["ca-cert"], d, "ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca-cert"], "SystemCentralManagement-CaCert"); ok {
			if err = d.Set("ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenSystemCentralManagementEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["enc-algorithm"], "SystemCentralManagement-EncAlgorithm"); ok {
			if err = d.Set("enc_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading enc_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("fmg", flattenSystemCentralManagementFmg(o["fmg"], d, "fmg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmg"], "SystemCentralManagement-Fmg"); ok {
			if err = d.Set("fmg", vv); err != nil {
				return fmt.Errorf("Error reading fmg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmg: %v", err)
		}
	}

	if err = d.Set("fmg_source_ip", flattenSystemCentralManagementFmgSourceIp(o["fmg-source-ip"], d, "fmg_source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmg-source-ip"], "SystemCentralManagement-FmgSourceIp"); ok {
			if err = d.Set("fmg_source_ip", vv); err != nil {
				return fmt.Errorf("Error reading fmg_source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmg_source_ip: %v", err)
		}
	}

	if err = d.Set("fmg_source_ip6", flattenSystemCentralManagementFmgSourceIp6(o["fmg-source-ip6"], d, "fmg_source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmg-source-ip6"], "SystemCentralManagement-FmgSourceIp6"); ok {
			if err = d.Set("fmg_source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading fmg_source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmg_source_ip6: %v", err)
		}
	}

	if err = d.Set("fmg_update_port", flattenSystemCentralManagementFmgUpdatePort(o["fmg-update-port"], d, "fmg_update_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmg-update-port"], "SystemCentralManagement-FmgUpdatePort"); ok {
			if err = d.Set("fmg_update_port", vv); err != nil {
				return fmt.Errorf("Error reading fmg_update_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmg_update_port: %v", err)
		}
	}

	if err = d.Set("fortigate_cloud_sso_default_profile", flattenSystemCentralManagementFortigateCloudSsoDefaultProfile(o["fortigate-cloud-sso-default-profile"], d, "fortigate_cloud_sso_default_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortigate-cloud-sso-default-profile"], "SystemCentralManagement-FortigateCloudSsoDefaultProfile"); ok {
			if err = d.Set("fortigate_cloud_sso_default_profile", vv); err != nil {
				return fmt.Errorf("Error reading fortigate_cloud_sso_default_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortigate_cloud_sso_default_profile: %v", err)
		}
	}

	if err = d.Set("include_default_servers", flattenSystemCentralManagementIncludeDefaultServers(o["include-default-servers"], d, "include_default_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["include-default-servers"], "SystemCentralManagement-IncludeDefaultServers"); ok {
			if err = d.Set("include_default_servers", vv); err != nil {
				return fmt.Errorf("Error reading include_default_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading include_default_servers: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemCentralManagementInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemCentralManagement-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemCentralManagementInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystemCentralManagement-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("local_cert", flattenSystemCentralManagementLocalCert(o["local-cert"], d, "local_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-cert"], "SystemCentralManagement-LocalCert"); ok {
			if err = d.Set("local_cert", vv); err != nil {
				return fmt.Errorf("Error reading local_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}

	if err = d.Set("ltefw_upgrade_frequency", flattenSystemCentralManagementLtefwUpgradeFrequency(o["ltefw-upgrade-frequency"], d, "ltefw_upgrade_frequency")); err != nil {
		if vv, ok := fortiAPIPatch(o["ltefw-upgrade-frequency"], "SystemCentralManagement-LtefwUpgradeFrequency"); ok {
			if err = d.Set("ltefw_upgrade_frequency", vv); err != nil {
				return fmt.Errorf("Error reading ltefw_upgrade_frequency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ltefw_upgrade_frequency: %v", err)
		}
	}

	if err = d.Set("ltefw_upgrade_time", flattenSystemCentralManagementLtefwUpgradeTime(o["ltefw-upgrade-time"], d, "ltefw_upgrade_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["ltefw-upgrade-time"], "SystemCentralManagement-LtefwUpgradeTime"); ok {
			if err = d.Set("ltefw_upgrade_time", vv); err != nil {
				return fmt.Errorf("Error reading ltefw_upgrade_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ltefw_upgrade_time: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemCentralManagementMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemCentralManagement-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("modem_upgrade_frequency", flattenSystemCentralManagementModemUpgradeFrequency(o["modem-upgrade-frequency"], d, "modem_upgrade_frequency")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem-upgrade-frequency"], "SystemCentralManagement-ModemUpgradeFrequency"); ok {
			if err = d.Set("modem_upgrade_frequency", vv); err != nil {
				return fmt.Errorf("Error reading modem_upgrade_frequency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem_upgrade_frequency: %v", err)
		}
	}

	if err = d.Set("modem_upgrade_time", flattenSystemCentralManagementModemUpgradeTime(o["modem-upgrade-time"], d, "modem_upgrade_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem-upgrade-time"], "SystemCentralManagement-ModemUpgradeTime"); ok {
			if err = d.Set("modem_upgrade_time", vv); err != nil {
				return fmt.Errorf("Error reading modem_upgrade_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem_upgrade_time: %v", err)
		}
	}

	if err = d.Set("schedule_config_restore", flattenSystemCentralManagementScheduleConfigRestore(o["schedule-config-restore"], d, "schedule_config_restore")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule-config-restore"], "SystemCentralManagement-ScheduleConfigRestore"); ok {
			if err = d.Set("schedule_config_restore", vv); err != nil {
				return fmt.Errorf("Error reading schedule_config_restore: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule_config_restore: %v", err)
		}
	}

	if err = d.Set("schedule_script_restore", flattenSystemCentralManagementScheduleScriptRestore(o["schedule-script-restore"], d, "schedule_script_restore")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule-script-restore"], "SystemCentralManagement-ScheduleScriptRestore"); ok {
			if err = d.Set("schedule_script_restore", vv); err != nil {
				return fmt.Errorf("Error reading schedule_script_restore: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule_script_restore: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenSystemCentralManagementSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial-number"], "SystemCentralManagement-SerialNumber"); ok {
			if err = d.Set("serial_number", vv); err != nil {
				return fmt.Errorf("Error reading serial_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_list", flattenSystemCentralManagementServerList(o["server-list"], d, "server_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-list"], "SystemCentralManagement-ServerList"); ok {
				if err = d.Set("server_list", vv); err != nil {
					return fmt.Errorf("Error reading server_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_list"); ok {
			if err = d.Set("server_list", flattenSystemCentralManagementServerList(o["server-list"], d, "server_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-list"], "SystemCentralManagement-ServerList"); ok {
					if err = d.Set("server_list", vv); err != nil {
						return fmt.Errorf("Error reading server_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenSystemCentralManagementType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemCentralManagement-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("use_elbc_vdom", flattenSystemCentralManagementUseElbcVdom(o["use-elbc-vdom"], d, "use_elbc_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-elbc-vdom"], "SystemCentralManagement-UseElbcVdom"); ok {
			if err = d.Set("use_elbc_vdom", vv); err != nil {
				return fmt.Errorf("Error reading use_elbc_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_elbc_vdom: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemCentralManagementVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemCentralManagement-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenSystemCentralManagementFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCentralManagementAllowMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowPushConfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowPushFirmware(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowRemoteFirmwareUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowRemoteLteFirmwareUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowRemoteModemFirmwareUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCentralManagementFmgSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmgSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmgUpdatePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFortigateCloudSsoDefaultProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCentralManagementIncludeDefaultServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCentralManagementInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementLocalCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementLtefwUpgradeFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementLtefwUpgradeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementModemUpgradeFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementModemUpgradeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementScheduleConfigRestore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementScheduleScriptRestore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCentralManagementServerList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-type"], _ = expandSystemCentralManagementServerListAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fqdn"], _ = expandSystemCentralManagementServerListFqdn(d, i["fqdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemCentralManagementServerListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-address"], _ = expandSystemCentralManagementServerListServerAddress(d, i["server_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-address6"], _ = expandSystemCentralManagementServerListServerAddress6(d, i["server_address6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-type"], _ = expandSystemCentralManagementServerListServerType(d, i["server_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemCentralManagementServerListAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListServerAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListServerAddress6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCentralManagementType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementUseElbcVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemCentralManagement(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allow_monitor"); ok || d.HasChange("allow_monitor") {
		t, err := expandSystemCentralManagementAllowMonitor(d, v, "allow_monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-monitor"] = t
		}
	}

	if v, ok := d.GetOk("allow_push_configuration"); ok || d.HasChange("allow_push_configuration") {
		t, err := expandSystemCentralManagementAllowPushConfiguration(d, v, "allow_push_configuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-push-configuration"] = t
		}
	}

	if v, ok := d.GetOk("allow_push_firmware"); ok || d.HasChange("allow_push_firmware") {
		t, err := expandSystemCentralManagementAllowPushFirmware(d, v, "allow_push_firmware")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-push-firmware"] = t
		}
	}

	if v, ok := d.GetOk("allow_remote_firmware_upgrade"); ok || d.HasChange("allow_remote_firmware_upgrade") {
		t, err := expandSystemCentralManagementAllowRemoteFirmwareUpgrade(d, v, "allow_remote_firmware_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-remote-firmware-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("allow_remote_lte_firmware_upgrade"); ok || d.HasChange("allow_remote_lte_firmware_upgrade") {
		t, err := expandSystemCentralManagementAllowRemoteLteFirmwareUpgrade(d, v, "allow_remote_lte_firmware_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-remote-lte-firmware-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("allow_remote_modem_firmware_upgrade"); ok || d.HasChange("allow_remote_modem_firmware_upgrade") {
		t, err := expandSystemCentralManagementAllowRemoteModemFirmwareUpgrade(d, v, "allow_remote_modem_firmware_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-remote-modem-firmware-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok || d.HasChange("ca_cert") {
		t, err := expandSystemCentralManagementCaCert(d, v, "ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok || d.HasChange("enc_algorithm") {
		t, err := expandSystemCentralManagementEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("fmg"); ok || d.HasChange("fmg") {
		t, err := expandSystemCentralManagementFmg(d, v, "fmg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmg"] = t
		}
	}

	if v, ok := d.GetOk("fmg_source_ip"); ok || d.HasChange("fmg_source_ip") {
		t, err := expandSystemCentralManagementFmgSourceIp(d, v, "fmg_source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmg-source-ip"] = t
		}
	}

	if v, ok := d.GetOk("fmg_source_ip6"); ok || d.HasChange("fmg_source_ip6") {
		t, err := expandSystemCentralManagementFmgSourceIp6(d, v, "fmg_source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmg-source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("fmg_update_port"); ok || d.HasChange("fmg_update_port") {
		t, err := expandSystemCentralManagementFmgUpdatePort(d, v, "fmg_update_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmg-update-port"] = t
		}
	}

	if v, ok := d.GetOk("fortigate_cloud_sso_default_profile"); ok || d.HasChange("fortigate_cloud_sso_default_profile") {
		t, err := expandSystemCentralManagementFortigateCloudSsoDefaultProfile(d, v, "fortigate_cloud_sso_default_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortigate-cloud-sso-default-profile"] = t
		}
	}

	if v, ok := d.GetOk("include_default_servers"); ok || d.HasChange("include_default_servers") {
		t, err := expandSystemCentralManagementIncludeDefaultServers(d, v, "include_default_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-default-servers"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemCentralManagementInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystemCentralManagementInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("local_cert"); ok || d.HasChange("local_cert") {
		t, err := expandSystemCentralManagementLocalCert(d, v, "local_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cert"] = t
		}
	}

	if v, ok := d.GetOk("ltefw_upgrade_frequency"); ok || d.HasChange("ltefw_upgrade_frequency") {
		t, err := expandSystemCentralManagementLtefwUpgradeFrequency(d, v, "ltefw_upgrade_frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ltefw-upgrade-frequency"] = t
		}
	}

	if v, ok := d.GetOk("ltefw_upgrade_time"); ok || d.HasChange("ltefw_upgrade_time") {
		t, err := expandSystemCentralManagementLtefwUpgradeTime(d, v, "ltefw_upgrade_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ltefw-upgrade-time"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemCentralManagementMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("modem_upgrade_frequency"); ok || d.HasChange("modem_upgrade_frequency") {
		t, err := expandSystemCentralManagementModemUpgradeFrequency(d, v, "modem_upgrade_frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-upgrade-frequency"] = t
		}
	}

	if v, ok := d.GetOk("modem_upgrade_time"); ok || d.HasChange("modem_upgrade_time") {
		t, err := expandSystemCentralManagementModemUpgradeTime(d, v, "modem_upgrade_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-upgrade-time"] = t
		}
	}

	if v, ok := d.GetOk("schedule_config_restore"); ok || d.HasChange("schedule_config_restore") {
		t, err := expandSystemCentralManagementScheduleConfigRestore(d, v, "schedule_config_restore")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-config-restore"] = t
		}
	}

	if v, ok := d.GetOk("schedule_script_restore"); ok || d.HasChange("schedule_script_restore") {
		t, err := expandSystemCentralManagementScheduleScriptRestore(d, v, "schedule_script_restore")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-script-restore"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok || d.HasChange("serial_number") {
		t, err := expandSystemCentralManagementSerialNumber(d, v, "serial_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("server_list"); ok || d.HasChange("server_list") {
		t, err := expandSystemCentralManagementServerList(d, v, "server_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-list"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemCentralManagementType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("use_elbc_vdom"); ok || d.HasChange("use_elbc_vdom") {
		t, err := expandSystemCentralManagementUseElbcVdom(d, v, "use_elbc_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-elbc-vdom"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemCentralManagementVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
