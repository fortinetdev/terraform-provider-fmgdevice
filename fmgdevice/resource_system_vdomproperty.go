// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure VDOM property.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVdomProperty() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomPropertyCreate,
		Read:   resourceSystemVdomPropertyRead,
		Update: resourceSystemVdomPropertyUpdate,
		Delete: resourceSystemVdomPropertyDelete,

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
			"custom_service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dialup_tunnel": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"firewall_address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"firewall_addrgrp": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"firewall_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"ipsec_phase1": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"ipsec_phase1_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"ipsec_phase2": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"ipsec_phase2_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"log_disk_quota": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"onetime_schedule": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"recurring_schedule": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"service_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"session": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"snmp_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sslvpn": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"user_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemVdomPropertyCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemVdomProperty(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomProperty resource while getting object: %v", err)
	}

	_, err = c.CreateSystemVdomProperty(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomProperty resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemVdomPropertyRead(d, m)
}

func resourceSystemVdomPropertyUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemVdomProperty(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomProperty resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVdomProperty(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomProperty resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemVdomPropertyRead(d, m)
}

func resourceSystemVdomPropertyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemVdomProperty(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomProperty resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomPropertyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVdomProperty(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomProperty resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomProperty(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomProperty resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomPropertyCustomService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyDialupTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyFirewallAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyFirewallAddrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyFirewallPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyIpsecPhase1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyIpsecPhase1Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyIpsecPhase2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyIpsecPhase2Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyLogDiskQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenSystemVdomPropertyOnetimeSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyRecurringSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertySession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertySnmpIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertySslvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVdomPropertyUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func refreshObjectSystemVdomProperty(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("custom_service", flattenSystemVdomPropertyCustomService(o["custom-service"], d, "custom_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-service"], "SystemVdomProperty-CustomService"); ok {
			if err = d.Set("custom_service", vv); err != nil {
				return fmt.Errorf("Error reading custom_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_service: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemVdomPropertyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemVdomProperty-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("dialup_tunnel", flattenSystemVdomPropertyDialupTunnel(o["dialup-tunnel"], d, "dialup_tunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["dialup-tunnel"], "SystemVdomProperty-DialupTunnel"); ok {
			if err = d.Set("dialup_tunnel", vv); err != nil {
				return fmt.Errorf("Error reading dialup_tunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dialup_tunnel: %v", err)
		}
	}

	if err = d.Set("firewall_address", flattenSystemVdomPropertyFirewallAddress(o["firewall-address"], d, "firewall_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["firewall-address"], "SystemVdomProperty-FirewallAddress"); ok {
			if err = d.Set("firewall_address", vv); err != nil {
				return fmt.Errorf("Error reading firewall_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firewall_address: %v", err)
		}
	}

	if err = d.Set("firewall_addrgrp", flattenSystemVdomPropertyFirewallAddrgrp(o["firewall-addrgrp"], d, "firewall_addrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["firewall-addrgrp"], "SystemVdomProperty-FirewallAddrgrp"); ok {
			if err = d.Set("firewall_addrgrp", vv); err != nil {
				return fmt.Errorf("Error reading firewall_addrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firewall_addrgrp: %v", err)
		}
	}

	if err = d.Set("firewall_policy", flattenSystemVdomPropertyFirewallPolicy(o["firewall-policy"], d, "firewall_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["firewall-policy"], "SystemVdomProperty-FirewallPolicy"); ok {
			if err = d.Set("firewall_policy", vv); err != nil {
				return fmt.Errorf("Error reading firewall_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firewall_policy: %v", err)
		}
	}

	if err = d.Set("ipsec_phase1", flattenSystemVdomPropertyIpsecPhase1(o["ipsec-phase1"], d, "ipsec_phase1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-phase1"], "SystemVdomProperty-IpsecPhase1"); ok {
			if err = d.Set("ipsec_phase1", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_phase1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_phase1: %v", err)
		}
	}

	if err = d.Set("ipsec_phase1_interface", flattenSystemVdomPropertyIpsecPhase1Interface(o["ipsec-phase1-interface"], d, "ipsec_phase1_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-phase1-interface"], "SystemVdomProperty-IpsecPhase1Interface"); ok {
			if err = d.Set("ipsec_phase1_interface", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_phase1_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_phase1_interface: %v", err)
		}
	}

	if err = d.Set("ipsec_phase2", flattenSystemVdomPropertyIpsecPhase2(o["ipsec-phase2"], d, "ipsec_phase2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-phase2"], "SystemVdomProperty-IpsecPhase2"); ok {
			if err = d.Set("ipsec_phase2", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_phase2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_phase2: %v", err)
		}
	}

	if err = d.Set("ipsec_phase2_interface", flattenSystemVdomPropertyIpsecPhase2Interface(o["ipsec-phase2-interface"], d, "ipsec_phase2_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-phase2-interface"], "SystemVdomProperty-IpsecPhase2Interface"); ok {
			if err = d.Set("ipsec_phase2_interface", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_phase2_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_phase2_interface: %v", err)
		}
	}

	if err = d.Set("log_disk_quota", flattenSystemVdomPropertyLogDiskQuota(o["log-disk-quota"], d, "log_disk_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-disk-quota"], "SystemVdomProperty-LogDiskQuota"); ok {
			if err = d.Set("log_disk_quota", vv); err != nil {
				return fmt.Errorf("Error reading log_disk_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_disk_quota: %v", err)
		}
	}

	if err = d.Set("onetime_schedule", flattenSystemVdomPropertyOnetimeSchedule(o["onetime-schedule"], d, "onetime_schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["onetime-schedule"], "SystemVdomProperty-OnetimeSchedule"); ok {
			if err = d.Set("onetime_schedule", vv); err != nil {
				return fmt.Errorf("Error reading onetime_schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading onetime_schedule: %v", err)
		}
	}

	if err = d.Set("proxy", flattenSystemVdomPropertyProxy(o["proxy"], d, "proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy"], "SystemVdomProperty-Proxy"); ok {
			if err = d.Set("proxy", vv); err != nil {
				return fmt.Errorf("Error reading proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("recurring_schedule", flattenSystemVdomPropertyRecurringSchedule(o["recurring-schedule"], d, "recurring_schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["recurring-schedule"], "SystemVdomProperty-RecurringSchedule"); ok {
			if err = d.Set("recurring_schedule", vv); err != nil {
				return fmt.Errorf("Error reading recurring_schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading recurring_schedule: %v", err)
		}
	}

	if err = d.Set("service_group", flattenSystemVdomPropertyServiceGroup(o["service-group"], d, "service_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-group"], "SystemVdomProperty-ServiceGroup"); ok {
			if err = d.Set("service_group", vv); err != nil {
				return fmt.Errorf("Error reading service_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_group: %v", err)
		}
	}

	if err = d.Set("session", flattenSystemVdomPropertySession(o["session"], d, "session")); err != nil {
		if vv, ok := fortiAPIPatch(o["session"], "SystemVdomProperty-Session"); ok {
			if err = d.Set("session", vv); err != nil {
				return fmt.Errorf("Error reading session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session: %v", err)
		}
	}

	if err = d.Set("snmp_index", flattenSystemVdomPropertySnmpIndex(o["snmp-index"], d, "snmp_index")); err != nil {
		if vv, ok := fortiAPIPatch(o["snmp-index"], "SystemVdomProperty-SnmpIndex"); ok {
			if err = d.Set("snmp_index", vv); err != nil {
				return fmt.Errorf("Error reading snmp_index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading snmp_index: %v", err)
		}
	}

	if err = d.Set("sslvpn", flattenSystemVdomPropertySslvpn(o["sslvpn"], d, "sslvpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sslvpn"], "SystemVdomProperty-Sslvpn"); ok {
			if err = d.Set("sslvpn", vv); err != nil {
				return fmt.Errorf("Error reading sslvpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sslvpn: %v", err)
		}
	}

	if err = d.Set("user", flattenSystemVdomPropertyUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "SystemVdomProperty-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("user_group", flattenSystemVdomPropertyUserGroup(o["user-group"], d, "user_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-group"], "SystemVdomProperty-UserGroup"); ok {
			if err = d.Set("user_group", vv); err != nil {
				return fmt.Errorf("Error reading user_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomPropertyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVdomPropertyCustomService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyDialupTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyFirewallAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyFirewallAddrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyFirewallPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyIpsecPhase1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyIpsecPhase1Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyIpsecPhase2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyIpsecPhase2Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyLogDiskQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandSystemVdomPropertyOnetimeSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyRecurringSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertySession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertySnmpIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertySslvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVdomPropertyUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func getObjectSystemVdomProperty(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("custom_service"); ok || d.HasChange("custom_service") {
		t, err := expandSystemVdomPropertyCustomService(d, v, "custom_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-service"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemVdomPropertyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dialup_tunnel"); ok || d.HasChange("dialup_tunnel") {
		t, err := expandSystemVdomPropertyDialupTunnel(d, v, "dialup_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dialup-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("firewall_address"); ok || d.HasChange("firewall_address") {
		t, err := expandSystemVdomPropertyFirewallAddress(d, v, "firewall_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-address"] = t
		}
	}

	if v, ok := d.GetOk("firewall_addrgrp"); ok || d.HasChange("firewall_addrgrp") {
		t, err := expandSystemVdomPropertyFirewallAddrgrp(d, v, "firewall_addrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-addrgrp"] = t
		}
	}

	if v, ok := d.GetOk("firewall_policy"); ok || d.HasChange("firewall_policy") {
		t, err := expandSystemVdomPropertyFirewallPolicy(d, v, "firewall_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-policy"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase1"); ok || d.HasChange("ipsec_phase1") {
		t, err := expandSystemVdomPropertyIpsecPhase1(d, v, "ipsec_phase1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase1"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase1_interface"); ok || d.HasChange("ipsec_phase1_interface") {
		t, err := expandSystemVdomPropertyIpsecPhase1Interface(d, v, "ipsec_phase1_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase1-interface"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase2"); ok || d.HasChange("ipsec_phase2") {
		t, err := expandSystemVdomPropertyIpsecPhase2(d, v, "ipsec_phase2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase2"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase2_interface"); ok || d.HasChange("ipsec_phase2_interface") {
		t, err := expandSystemVdomPropertyIpsecPhase2Interface(d, v, "ipsec_phase2_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase2-interface"] = t
		}
	}

	if v, ok := d.GetOk("log_disk_quota"); ok || d.HasChange("log_disk_quota") {
		t, err := expandSystemVdomPropertyLogDiskQuota(d, v, "log_disk_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-disk-quota"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemVdomPropertyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("onetime_schedule"); ok || d.HasChange("onetime_schedule") {
		t, err := expandSystemVdomPropertyOnetimeSchedule(d, v, "onetime_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["onetime-schedule"] = t
		}
	}

	if v, ok := d.GetOk("proxy"); ok || d.HasChange("proxy") {
		t, err := expandSystemVdomPropertyProxy(d, v, "proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	if v, ok := d.GetOk("recurring_schedule"); ok || d.HasChange("recurring_schedule") {
		t, err := expandSystemVdomPropertyRecurringSchedule(d, v, "recurring_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recurring-schedule"] = t
		}
	}

	if v, ok := d.GetOk("service_group"); ok || d.HasChange("service_group") {
		t, err := expandSystemVdomPropertyServiceGroup(d, v, "service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-group"] = t
		}
	}

	if v, ok := d.GetOk("session"); ok || d.HasChange("session") {
		t, err := expandSystemVdomPropertySession(d, v, "session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session"] = t
		}
	}

	if v, ok := d.GetOk("snmp_index"); ok || d.HasChange("snmp_index") {
		t, err := expandSystemVdomPropertySnmpIndex(d, v, "snmp_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp-index"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn"); ok || d.HasChange("sslvpn") {
		t, err := expandSystemVdomPropertySslvpn(d, v, "sslvpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandSystemVdomPropertyUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok || d.HasChange("user_group") {
		t, err := expandSystemVdomPropertyUserGroup(d, v, "user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	return &obj, nil
}
