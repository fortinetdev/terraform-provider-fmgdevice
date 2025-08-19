// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch global settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerGlobalUpdate,
		Read:   resourceSwitchControllerGlobalRead,
		Update: resourceSwitchControllerGlobalUpdate,
		Delete: resourceSwitchControllerGlobalDelete,

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
			"allow_multiple_interfaces": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bounce_quarantined_link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_command": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_entry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"command_name": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"default_virtual_switch_vlan": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_circuit_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_remote_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dhcp_server_access_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snoop_client_db_exp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcp_snoop_client_req": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snoop_db_per_port_learn_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"disable_discovery": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fips_enforce": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firmware_provision_on_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_image_push": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_mac_limit_violations": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_aging_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mac_event_logging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_retention_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mac_violation_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"quarantine_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sn_dns_resolution": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_on_deauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_user_device": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vlan_all_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_optimization": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceSwitchControllerGlobalUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerGlobal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerGlobal")

	return resourceSwitchControllerGlobalRead(d, m)
}

func resourceSwitchControllerGlobalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerGlobal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerGlobal resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerGlobalAllowMultipleInterfaces(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalBounceQuarantinedLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalCustomCommand(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_entry"
		if _, ok := i["command-entry"]; ok {
			v := flattenSwitchControllerGlobalCustomCommandCommandEntry(i["command-entry"], d, pre_append)
			tmp["command_entry"] = fortiAPISubPartPatch(v, "SwitchControllerGlobal-CustomCommand-CommandEntry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := i["command-name"]; ok {
			v := flattenSwitchControllerGlobalCustomCommandCommandName(i["command-name"], d, pre_append)
			tmp["command_name"] = fortiAPISubPartPatch(v, "SwitchControllerGlobal-CustomCommand-CommandName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerGlobalCustomCommandCommandEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalCustomCommandCommandName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerGlobalDefaultVirtualSwitchVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerGlobalDhcpOption82CircuitId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerGlobalDhcpOption82Format(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDhcpOption82RemoteId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerGlobalDhcpServerAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDhcpSnoopClientDbExp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDhcpSnoopClientReq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDhcpSnoopDbPerPortLearnLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDisableDiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerGlobalFipsEnforce(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalFirmwareProvisionOnAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalHttpsImagePush(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalLogMacLimitViolations(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalMacAgingInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalMacEventLogging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalMacRetentionPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalMacViolationTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalQuarantineMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalSnDnsResolution(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalSwitchOnDeauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalUpdateUserDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerGlobalVlanAllMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalVlanIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerGlobalVlanOptimization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("allow_multiple_interfaces", flattenSwitchControllerGlobalAllowMultipleInterfaces(o["allow-multiple-interfaces"], d, "allow_multiple_interfaces")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-multiple-interfaces"], "SwitchControllerGlobal-AllowMultipleInterfaces"); ok {
			if err = d.Set("allow_multiple_interfaces", vv); err != nil {
				return fmt.Errorf("Error reading allow_multiple_interfaces: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_multiple_interfaces: %v", err)
		}
	}

	if err = d.Set("bounce_quarantined_link", flattenSwitchControllerGlobalBounceQuarantinedLink(o["bounce-quarantined-link"], d, "bounce_quarantined_link")); err != nil {
		if vv, ok := fortiAPIPatch(o["bounce-quarantined-link"], "SwitchControllerGlobal-BounceQuarantinedLink"); ok {
			if err = d.Set("bounce_quarantined_link", vv); err != nil {
				return fmt.Errorf("Error reading bounce_quarantined_link: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bounce_quarantined_link: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_command", flattenSwitchControllerGlobalCustomCommand(o["custom-command"], d, "custom_command")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-command"], "SwitchControllerGlobal-CustomCommand"); ok {
				if err = d.Set("custom_command", vv); err != nil {
					return fmt.Errorf("Error reading custom_command: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_command: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_command"); ok {
			if err = d.Set("custom_command", flattenSwitchControllerGlobalCustomCommand(o["custom-command"], d, "custom_command")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-command"], "SwitchControllerGlobal-CustomCommand"); ok {
					if err = d.Set("custom_command", vv); err != nil {
						return fmt.Errorf("Error reading custom_command: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_command: %v", err)
				}
			}
		}
	}

	if err = d.Set("default_virtual_switch_vlan", flattenSwitchControllerGlobalDefaultVirtualSwitchVlan(o["default-virtual-switch-vlan"], d, "default_virtual_switch_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-virtual-switch-vlan"], "SwitchControllerGlobal-DefaultVirtualSwitchVlan"); ok {
			if err = d.Set("default_virtual_switch_vlan", vv); err != nil {
				return fmt.Errorf("Error reading default_virtual_switch_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_virtual_switch_vlan: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_circuit_id", flattenSwitchControllerGlobalDhcpOption82CircuitId(o["dhcp-option82-circuit-id"], d, "dhcp_option82_circuit_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option82-circuit-id"], "SwitchControllerGlobal-DhcpOption82CircuitId"); ok {
			if err = d.Set("dhcp_option82_circuit_id", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option82_circuit_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option82_circuit_id: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_format", flattenSwitchControllerGlobalDhcpOption82Format(o["dhcp-option82-format"], d, "dhcp_option82_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option82-format"], "SwitchControllerGlobal-DhcpOption82Format"); ok {
			if err = d.Set("dhcp_option82_format", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option82_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option82_format: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_remote_id", flattenSwitchControllerGlobalDhcpOption82RemoteId(o["dhcp-option82-remote-id"], d, "dhcp_option82_remote_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-option82-remote-id"], "SwitchControllerGlobal-DhcpOption82RemoteId"); ok {
			if err = d.Set("dhcp_option82_remote_id", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_option82_remote_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_option82_remote_id: %v", err)
		}
	}

	if err = d.Set("dhcp_server_access_list", flattenSwitchControllerGlobalDhcpServerAccessList(o["dhcp-server-access-list"], d, "dhcp_server_access_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-server-access-list"], "SwitchControllerGlobal-DhcpServerAccessList"); ok {
			if err = d.Set("dhcp_server_access_list", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_server_access_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_server_access_list: %v", err)
		}
	}

	if err = d.Set("dhcp_snoop_client_db_exp", flattenSwitchControllerGlobalDhcpSnoopClientDbExp(o["dhcp-snoop-client-db-exp"], d, "dhcp_snoop_client_db_exp")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-snoop-client-db-exp"], "SwitchControllerGlobal-DhcpSnoopClientDbExp"); ok {
			if err = d.Set("dhcp_snoop_client_db_exp", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_snoop_client_db_exp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_snoop_client_db_exp: %v", err)
		}
	}

	if err = d.Set("dhcp_snoop_client_req", flattenSwitchControllerGlobalDhcpSnoopClientReq(o["dhcp-snoop-client-req"], d, "dhcp_snoop_client_req")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-snoop-client-req"], "SwitchControllerGlobal-DhcpSnoopClientReq"); ok {
			if err = d.Set("dhcp_snoop_client_req", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_snoop_client_req: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_snoop_client_req: %v", err)
		}
	}

	if err = d.Set("dhcp_snoop_db_per_port_learn_limit", flattenSwitchControllerGlobalDhcpSnoopDbPerPortLearnLimit(o["dhcp-snoop-db-per-port-learn-limit"], d, "dhcp_snoop_db_per_port_learn_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-snoop-db-per-port-learn-limit"], "SwitchControllerGlobal-DhcpSnoopDbPerPortLearnLimit"); ok {
			if err = d.Set("dhcp_snoop_db_per_port_learn_limit", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_snoop_db_per_port_learn_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_snoop_db_per_port_learn_limit: %v", err)
		}
	}

	if err = d.Set("disable_discovery", flattenSwitchControllerGlobalDisableDiscovery(o["disable-discovery"], d, "disable_discovery")); err != nil {
		if vv, ok := fortiAPIPatch(o["disable-discovery"], "SwitchControllerGlobal-DisableDiscovery"); ok {
			if err = d.Set("disable_discovery", vv); err != nil {
				return fmt.Errorf("Error reading disable_discovery: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disable_discovery: %v", err)
		}
	}

	if err = d.Set("fips_enforce", flattenSwitchControllerGlobalFipsEnforce(o["fips-enforce"], d, "fips_enforce")); err != nil {
		if vv, ok := fortiAPIPatch(o["fips-enforce"], "SwitchControllerGlobal-FipsEnforce"); ok {
			if err = d.Set("fips_enforce", vv); err != nil {
				return fmt.Errorf("Error reading fips_enforce: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fips_enforce: %v", err)
		}
	}

	if err = d.Set("firmware_provision_on_authorization", flattenSwitchControllerGlobalFirmwareProvisionOnAuthorization(o["firmware-provision-on-authorization"], d, "firmware_provision_on_authorization")); err != nil {
		if vv, ok := fortiAPIPatch(o["firmware-provision-on-authorization"], "SwitchControllerGlobal-FirmwareProvisionOnAuthorization"); ok {
			if err = d.Set("firmware_provision_on_authorization", vv); err != nil {
				return fmt.Errorf("Error reading firmware_provision_on_authorization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firmware_provision_on_authorization: %v", err)
		}
	}

	if err = d.Set("https_image_push", flattenSwitchControllerGlobalHttpsImagePush(o["https-image-push"], d, "https_image_push")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-image-push"], "SwitchControllerGlobal-HttpsImagePush"); ok {
			if err = d.Set("https_image_push", vv); err != nil {
				return fmt.Errorf("Error reading https_image_push: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_image_push: %v", err)
		}
	}

	if err = d.Set("log_mac_limit_violations", flattenSwitchControllerGlobalLogMacLimitViolations(o["log-mac-limit-violations"], d, "log_mac_limit_violations")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-mac-limit-violations"], "SwitchControllerGlobal-LogMacLimitViolations"); ok {
			if err = d.Set("log_mac_limit_violations", vv); err != nil {
				return fmt.Errorf("Error reading log_mac_limit_violations: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_mac_limit_violations: %v", err)
		}
	}

	if err = d.Set("mac_aging_interval", flattenSwitchControllerGlobalMacAgingInterval(o["mac-aging-interval"], d, "mac_aging_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-aging-interval"], "SwitchControllerGlobal-MacAgingInterval"); ok {
			if err = d.Set("mac_aging_interval", vv); err != nil {
				return fmt.Errorf("Error reading mac_aging_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_aging_interval: %v", err)
		}
	}

	if err = d.Set("mac_event_logging", flattenSwitchControllerGlobalMacEventLogging(o["mac-event-logging"], d, "mac_event_logging")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-event-logging"], "SwitchControllerGlobal-MacEventLogging"); ok {
			if err = d.Set("mac_event_logging", vv); err != nil {
				return fmt.Errorf("Error reading mac_event_logging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_event_logging: %v", err)
		}
	}

	if err = d.Set("mac_retention_period", flattenSwitchControllerGlobalMacRetentionPeriod(o["mac-retention-period"], d, "mac_retention_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-retention-period"], "SwitchControllerGlobal-MacRetentionPeriod"); ok {
			if err = d.Set("mac_retention_period", vv); err != nil {
				return fmt.Errorf("Error reading mac_retention_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_retention_period: %v", err)
		}
	}

	if err = d.Set("mac_violation_timer", flattenSwitchControllerGlobalMacViolationTimer(o["mac-violation-timer"], d, "mac_violation_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-violation-timer"], "SwitchControllerGlobal-MacViolationTimer"); ok {
			if err = d.Set("mac_violation_timer", vv); err != nil {
				return fmt.Errorf("Error reading mac_violation_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_violation_timer: %v", err)
		}
	}

	if err = d.Set("quarantine_mode", flattenSwitchControllerGlobalQuarantineMode(o["quarantine-mode"], d, "quarantine_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-mode"], "SwitchControllerGlobal-QuarantineMode"); ok {
			if err = d.Set("quarantine_mode", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_mode: %v", err)
		}
	}

	if err = d.Set("sn_dns_resolution", flattenSwitchControllerGlobalSnDnsResolution(o["sn-dns-resolution"], d, "sn_dns_resolution")); err != nil {
		if vv, ok := fortiAPIPatch(o["sn-dns-resolution"], "SwitchControllerGlobal-SnDnsResolution"); ok {
			if err = d.Set("sn_dns_resolution", vv); err != nil {
				return fmt.Errorf("Error reading sn_dns_resolution: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sn_dns_resolution: %v", err)
		}
	}

	if err = d.Set("switch_on_deauth", flattenSwitchControllerGlobalSwitchOnDeauth(o["switch-on-deauth"], d, "switch_on_deauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-on-deauth"], "SwitchControllerGlobal-SwitchOnDeauth"); ok {
			if err = d.Set("switch_on_deauth", vv); err != nil {
				return fmt.Errorf("Error reading switch_on_deauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_on_deauth: %v", err)
		}
	}

	if err = d.Set("update_user_device", flattenSwitchControllerGlobalUpdateUserDevice(o["update-user-device"], d, "update_user_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-user-device"], "SwitchControllerGlobal-UpdateUserDevice"); ok {
			if err = d.Set("update_user_device", vv); err != nil {
				return fmt.Errorf("Error reading update_user_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_user_device: %v", err)
		}
	}

	if err = d.Set("vlan_all_mode", flattenSwitchControllerGlobalVlanAllMode(o["vlan-all-mode"], d, "vlan_all_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-all-mode"], "SwitchControllerGlobal-VlanAllMode"); ok {
			if err = d.Set("vlan_all_mode", vv); err != nil {
				return fmt.Errorf("Error reading vlan_all_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_all_mode: %v", err)
		}
	}

	if err = d.Set("vlan_identity", flattenSwitchControllerGlobalVlanIdentity(o["vlan-identity"], d, "vlan_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-identity"], "SwitchControllerGlobal-VlanIdentity"); ok {
			if err = d.Set("vlan_identity", vv); err != nil {
				return fmt.Errorf("Error reading vlan_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_identity: %v", err)
		}
	}

	if err = d.Set("vlan_optimization", flattenSwitchControllerGlobalVlanOptimization(o["vlan-optimization"], d, "vlan_optimization")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-optimization"], "SwitchControllerGlobal-VlanOptimization"); ok {
			if err = d.Set("vlan_optimization", vv); err != nil {
				return fmt.Errorf("Error reading vlan_optimization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_optimization: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerGlobalAllowMultipleInterfaces(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalBounceQuarantinedLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalCustomCommand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_entry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["command-entry"], _ = expandSwitchControllerGlobalCustomCommandCommandEntry(d, i["command_entry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["command-name"], _ = expandSwitchControllerGlobalCustomCommandCommandName(d, i["command_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerGlobalCustomCommandCommandEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalCustomCommandCommandName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerGlobalDefaultVirtualSwitchVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerGlobalDhcpOption82CircuitId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerGlobalDhcpOption82Format(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDhcpOption82RemoteId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerGlobalDhcpServerAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDhcpSnoopClientDbExp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDhcpSnoopClientReq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDhcpSnoopDbPerPortLearnLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDisableDiscovery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerGlobalFipsEnforce(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalFirmwareProvisionOnAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalHttpsImagePush(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalLogMacLimitViolations(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalMacAgingInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalMacEventLogging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalMacRetentionPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalMacViolationTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalQuarantineMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalSnDnsResolution(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalSwitchOnDeauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalUpdateUserDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerGlobalVlanAllMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalVlanIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalVlanOptimization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allow_multiple_interfaces"); ok || d.HasChange("allow_multiple_interfaces") {
		t, err := expandSwitchControllerGlobalAllowMultipleInterfaces(d, v, "allow_multiple_interfaces")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-multiple-interfaces"] = t
		}
	}

	if v, ok := d.GetOk("bounce_quarantined_link"); ok || d.HasChange("bounce_quarantined_link") {
		t, err := expandSwitchControllerGlobalBounceQuarantinedLink(d, v, "bounce_quarantined_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bounce-quarantined-link"] = t
		}
	}

	if v, ok := d.GetOk("custom_command"); ok || d.HasChange("custom_command") {
		t, err := expandSwitchControllerGlobalCustomCommand(d, v, "custom_command")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-command"] = t
		}
	}

	if v, ok := d.GetOk("default_virtual_switch_vlan"); ok || d.HasChange("default_virtual_switch_vlan") {
		t, err := expandSwitchControllerGlobalDefaultVirtualSwitchVlan(d, v, "default_virtual_switch_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-virtual-switch-vlan"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_circuit_id"); ok || d.HasChange("dhcp_option82_circuit_id") {
		t, err := expandSwitchControllerGlobalDhcpOption82CircuitId(d, v, "dhcp_option82_circuit_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-circuit-id"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_format"); ok || d.HasChange("dhcp_option82_format") {
		t, err := expandSwitchControllerGlobalDhcpOption82Format(d, v, "dhcp_option82_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-format"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_remote_id"); ok || d.HasChange("dhcp_option82_remote_id") {
		t, err := expandSwitchControllerGlobalDhcpOption82RemoteId(d, v, "dhcp_option82_remote_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-remote-id"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server_access_list"); ok || d.HasChange("dhcp_server_access_list") {
		t, err := expandSwitchControllerGlobalDhcpServerAccessList(d, v, "dhcp_server_access_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server-access-list"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snoop_client_db_exp"); ok || d.HasChange("dhcp_snoop_client_db_exp") {
		t, err := expandSwitchControllerGlobalDhcpSnoopClientDbExp(d, v, "dhcp_snoop_client_db_exp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snoop-client-db-exp"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snoop_client_req"); ok || d.HasChange("dhcp_snoop_client_req") {
		t, err := expandSwitchControllerGlobalDhcpSnoopClientReq(d, v, "dhcp_snoop_client_req")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snoop-client-req"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snoop_db_per_port_learn_limit"); ok || d.HasChange("dhcp_snoop_db_per_port_learn_limit") {
		t, err := expandSwitchControllerGlobalDhcpSnoopDbPerPortLearnLimit(d, v, "dhcp_snoop_db_per_port_learn_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snoop-db-per-port-learn-limit"] = t
		}
	}

	if v, ok := d.GetOk("disable_discovery"); ok || d.HasChange("disable_discovery") {
		t, err := expandSwitchControllerGlobalDisableDiscovery(d, v, "disable_discovery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disable-discovery"] = t
		}
	}

	if v, ok := d.GetOk("fips_enforce"); ok || d.HasChange("fips_enforce") {
		t, err := expandSwitchControllerGlobalFipsEnforce(d, v, "fips_enforce")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fips-enforce"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision_on_authorization"); ok || d.HasChange("firmware_provision_on_authorization") {
		t, err := expandSwitchControllerGlobalFirmwareProvisionOnAuthorization(d, v, "firmware_provision_on_authorization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision-on-authorization"] = t
		}
	}

	if v, ok := d.GetOk("https_image_push"); ok || d.HasChange("https_image_push") {
		t, err := expandSwitchControllerGlobalHttpsImagePush(d, v, "https_image_push")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-image-push"] = t
		}
	}

	if v, ok := d.GetOk("log_mac_limit_violations"); ok || d.HasChange("log_mac_limit_violations") {
		t, err := expandSwitchControllerGlobalLogMacLimitViolations(d, v, "log_mac_limit_violations")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-mac-limit-violations"] = t
		}
	}

	if v, ok := d.GetOk("mac_aging_interval"); ok || d.HasChange("mac_aging_interval") {
		t, err := expandSwitchControllerGlobalMacAgingInterval(d, v, "mac_aging_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-aging-interval"] = t
		}
	}

	if v, ok := d.GetOk("mac_event_logging"); ok || d.HasChange("mac_event_logging") {
		t, err := expandSwitchControllerGlobalMacEventLogging(d, v, "mac_event_logging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-event-logging"] = t
		}
	}

	if v, ok := d.GetOk("mac_retention_period"); ok || d.HasChange("mac_retention_period") {
		t, err := expandSwitchControllerGlobalMacRetentionPeriod(d, v, "mac_retention_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-retention-period"] = t
		}
	}

	if v, ok := d.GetOk("mac_violation_timer"); ok || d.HasChange("mac_violation_timer") {
		t, err := expandSwitchControllerGlobalMacViolationTimer(d, v, "mac_violation_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-violation-timer"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_mode"); ok || d.HasChange("quarantine_mode") {
		t, err := expandSwitchControllerGlobalQuarantineMode(d, v, "quarantine_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-mode"] = t
		}
	}

	if v, ok := d.GetOk("sn_dns_resolution"); ok || d.HasChange("sn_dns_resolution") {
		t, err := expandSwitchControllerGlobalSnDnsResolution(d, v, "sn_dns_resolution")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sn-dns-resolution"] = t
		}
	}

	if v, ok := d.GetOk("switch_on_deauth"); ok || d.HasChange("switch_on_deauth") {
		t, err := expandSwitchControllerGlobalSwitchOnDeauth(d, v, "switch_on_deauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-on-deauth"] = t
		}
	}

	if v, ok := d.GetOk("update_user_device"); ok || d.HasChange("update_user_device") {
		t, err := expandSwitchControllerGlobalUpdateUserDevice(d, v, "update_user_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-user-device"] = t
		}
	}

	if v, ok := d.GetOk("vlan_all_mode"); ok || d.HasChange("vlan_all_mode") {
		t, err := expandSwitchControllerGlobalVlanAllMode(d, v, "vlan_all_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-all-mode"] = t
		}
	}

	if v, ok := d.GetOk("vlan_identity"); ok || d.HasChange("vlan_identity") {
		t, err := expandSwitchControllerGlobalVlanIdentity(d, v, "vlan_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-identity"] = t
		}
	}

	if v, ok := d.GetOk("vlan_optimization"); ok || d.HasChange("vlan_optimization") {
		t, err := expandSwitchControllerGlobalVlanOptimization(d, v, "vlan_optimization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-optimization"] = t
		}
	}

	return &obj, nil
}
