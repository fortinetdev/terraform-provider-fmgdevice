// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure general log settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSettingUpdate,
		Read:   resourceLogSettingRead,
		Update: resourceLogSettingUpdate,
		Delete: resourceLogSettingDelete,

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
			"anonymization_hash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"brief_traffic_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_log_fields": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"daemon_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expolicy_implicit_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_utm_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"faz_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiview_weekly_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fwpolicy_implicit_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwpolicy6_implicit_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_in_allow": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_in_deny_broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_in_deny_unicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_in_policy_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_out": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_out_ioc_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_invalid_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_policy_comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_user_in_upper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"long_live_session_stat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"neighbor_event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resolve_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resolve_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rest_api_get": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rest_api_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syslog_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_anonymize": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogSetting")

	return resourceLogSettingRead(d, m)
}

func resourceLogSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogSettingAnonymizationHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingBriefTrafficFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingCustomLogFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogSettingDaemonLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingExpolicyImplicitLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingExtendedUtmLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingFazOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingFortiviewWeeklyData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingFwpolicyImplicitLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingFwpolicy6ImplicitLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingLocalInAllow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingLocalInDenyBroadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingLocalInDenyUnicast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingLocalInPolicyLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingLocalOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingLocalOutIocDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingLogInvalidPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingLogPolicyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingLogPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingLogUserInUpper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingLongLiveSessionStat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingNeighborEvent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingResolveIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingResolvePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingRestApiGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingRestApiSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingSyslogOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSettingUserAnonymize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("anonymization_hash", flattenLogSettingAnonymizationHash(o["anonymization-hash"], d, "anonymization_hash")); err != nil {
		if vv, ok := fortiAPIPatch(o["anonymization-hash"], "LogSetting-AnonymizationHash"); ok {
			if err = d.Set("anonymization_hash", vv); err != nil {
				return fmt.Errorf("Error reading anonymization_hash: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anonymization_hash: %v", err)
		}
	}

	if err = d.Set("brief_traffic_format", flattenLogSettingBriefTrafficFormat(o["brief-traffic-format"], d, "brief_traffic_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["brief-traffic-format"], "LogSetting-BriefTrafficFormat"); ok {
			if err = d.Set("brief_traffic_format", vv); err != nil {
				return fmt.Errorf("Error reading brief_traffic_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading brief_traffic_format: %v", err)
		}
	}

	if err = d.Set("custom_log_fields", flattenLogSettingCustomLogFields(o["custom-log-fields"], d, "custom_log_fields")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-log-fields"], "LogSetting-CustomLogFields"); ok {
			if err = d.Set("custom_log_fields", vv); err != nil {
				return fmt.Errorf("Error reading custom_log_fields: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_log_fields: %v", err)
		}
	}

	if err = d.Set("daemon_log", flattenLogSettingDaemonLog(o["daemon-log"], d, "daemon_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["daemon-log"], "LogSetting-DaemonLog"); ok {
			if err = d.Set("daemon_log", vv); err != nil {
				return fmt.Errorf("Error reading daemon_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading daemon_log: %v", err)
		}
	}

	if err = d.Set("expolicy_implicit_log", flattenLogSettingExpolicyImplicitLog(o["expolicy-implicit-log"], d, "expolicy_implicit_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["expolicy-implicit-log"], "LogSetting-ExpolicyImplicitLog"); ok {
			if err = d.Set("expolicy_implicit_log", vv); err != nil {
				return fmt.Errorf("Error reading expolicy_implicit_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expolicy_implicit_log: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenLogSettingExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "LogSetting-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("extended_utm_log", flattenLogSettingExtendedUtmLog(o["extended-utm-log"], d, "extended_utm_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-utm-log"], "LogSetting-ExtendedUtmLog"); ok {
			if err = d.Set("extended_utm_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_utm_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_utm_log: %v", err)
		}
	}

	if err = d.Set("faz_override", flattenLogSettingFazOverride(o["faz-override"], d, "faz_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-override"], "LogSetting-FazOverride"); ok {
			if err = d.Set("faz_override", vv); err != nil {
				return fmt.Errorf("Error reading faz_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_override: %v", err)
		}
	}

	if err = d.Set("fortiview_weekly_data", flattenLogSettingFortiviewWeeklyData(o["fortiview-weekly-data"], d, "fortiview_weekly_data")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiview-weekly-data"], "LogSetting-FortiviewWeeklyData"); ok {
			if err = d.Set("fortiview_weekly_data", vv); err != nil {
				return fmt.Errorf("Error reading fortiview_weekly_data: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiview_weekly_data: %v", err)
		}
	}

	if err = d.Set("fwpolicy_implicit_log", flattenLogSettingFwpolicyImplicitLog(o["fwpolicy-implicit-log"], d, "fwpolicy_implicit_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwpolicy-implicit-log"], "LogSetting-FwpolicyImplicitLog"); ok {
			if err = d.Set("fwpolicy_implicit_log", vv); err != nil {
				return fmt.Errorf("Error reading fwpolicy_implicit_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwpolicy_implicit_log: %v", err)
		}
	}

	if err = d.Set("fwpolicy6_implicit_log", flattenLogSettingFwpolicy6ImplicitLog(o["fwpolicy6-implicit-log"], d, "fwpolicy6_implicit_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwpolicy6-implicit-log"], "LogSetting-Fwpolicy6ImplicitLog"); ok {
			if err = d.Set("fwpolicy6_implicit_log", vv); err != nil {
				return fmt.Errorf("Error reading fwpolicy6_implicit_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwpolicy6_implicit_log: %v", err)
		}
	}

	if err = d.Set("local_in_allow", flattenLogSettingLocalInAllow(o["local-in-allow"], d, "local_in_allow")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-in-allow"], "LogSetting-LocalInAllow"); ok {
			if err = d.Set("local_in_allow", vv); err != nil {
				return fmt.Errorf("Error reading local_in_allow: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_in_allow: %v", err)
		}
	}

	if err = d.Set("local_in_deny_broadcast", flattenLogSettingLocalInDenyBroadcast(o["local-in-deny-broadcast"], d, "local_in_deny_broadcast")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-in-deny-broadcast"], "LogSetting-LocalInDenyBroadcast"); ok {
			if err = d.Set("local_in_deny_broadcast", vv); err != nil {
				return fmt.Errorf("Error reading local_in_deny_broadcast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_in_deny_broadcast: %v", err)
		}
	}

	if err = d.Set("local_in_deny_unicast", flattenLogSettingLocalInDenyUnicast(o["local-in-deny-unicast"], d, "local_in_deny_unicast")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-in-deny-unicast"], "LogSetting-LocalInDenyUnicast"); ok {
			if err = d.Set("local_in_deny_unicast", vv); err != nil {
				return fmt.Errorf("Error reading local_in_deny_unicast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_in_deny_unicast: %v", err)
		}
	}

	if err = d.Set("local_in_policy_log", flattenLogSettingLocalInPolicyLog(o["local-in-policy-log"], d, "local_in_policy_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-in-policy-log"], "LogSetting-LocalInPolicyLog"); ok {
			if err = d.Set("local_in_policy_log", vv); err != nil {
				return fmt.Errorf("Error reading local_in_policy_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_in_policy_log: %v", err)
		}
	}

	if err = d.Set("local_out", flattenLogSettingLocalOut(o["local-out"], d, "local_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-out"], "LogSetting-LocalOut"); ok {
			if err = d.Set("local_out", vv); err != nil {
				return fmt.Errorf("Error reading local_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_out: %v", err)
		}
	}

	if err = d.Set("local_out_ioc_detection", flattenLogSettingLocalOutIocDetection(o["local-out-ioc-detection"], d, "local_out_ioc_detection")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-out-ioc-detection"], "LogSetting-LocalOutIocDetection"); ok {
			if err = d.Set("local_out_ioc_detection", vv); err != nil {
				return fmt.Errorf("Error reading local_out_ioc_detection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_out_ioc_detection: %v", err)
		}
	}

	if err = d.Set("log_invalid_packet", flattenLogSettingLogInvalidPacket(o["log-invalid-packet"], d, "log_invalid_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-invalid-packet"], "LogSetting-LogInvalidPacket"); ok {
			if err = d.Set("log_invalid_packet", vv); err != nil {
				return fmt.Errorf("Error reading log_invalid_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_invalid_packet: %v", err)
		}
	}

	if err = d.Set("log_policy_comment", flattenLogSettingLogPolicyComment(o["log-policy-comment"], d, "log_policy_comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-policy-comment"], "LogSetting-LogPolicyComment"); ok {
			if err = d.Set("log_policy_comment", vv); err != nil {
				return fmt.Errorf("Error reading log_policy_comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_policy_comment: %v", err)
		}
	}

	if err = d.Set("log_policy_name", flattenLogSettingLogPolicyName(o["log-policy-name"], d, "log_policy_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-policy-name"], "LogSetting-LogPolicyName"); ok {
			if err = d.Set("log_policy_name", vv); err != nil {
				return fmt.Errorf("Error reading log_policy_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_policy_name: %v", err)
		}
	}

	if err = d.Set("log_user_in_upper", flattenLogSettingLogUserInUpper(o["log-user-in-upper"], d, "log_user_in_upper")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-user-in-upper"], "LogSetting-LogUserInUpper"); ok {
			if err = d.Set("log_user_in_upper", vv); err != nil {
				return fmt.Errorf("Error reading log_user_in_upper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_user_in_upper: %v", err)
		}
	}

	if err = d.Set("long_live_session_stat", flattenLogSettingLongLiveSessionStat(o["long-live-session-stat"], d, "long_live_session_stat")); err != nil {
		if vv, ok := fortiAPIPatch(o["long-live-session-stat"], "LogSetting-LongLiveSessionStat"); ok {
			if err = d.Set("long_live_session_stat", vv); err != nil {
				return fmt.Errorf("Error reading long_live_session_stat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading long_live_session_stat: %v", err)
		}
	}

	if err = d.Set("neighbor_event", flattenLogSettingNeighborEvent(o["neighbor-event"], d, "neighbor_event")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-event"], "LogSetting-NeighborEvent"); ok {
			if err = d.Set("neighbor_event", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_event: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_event: %v", err)
		}
	}

	if err = d.Set("resolve_ip", flattenLogSettingResolveIp(o["resolve-ip"], d, "resolve_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["resolve-ip"], "LogSetting-ResolveIp"); ok {
			if err = d.Set("resolve_ip", vv); err != nil {
				return fmt.Errorf("Error reading resolve_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resolve_ip: %v", err)
		}
	}

	if err = d.Set("resolve_port", flattenLogSettingResolvePort(o["resolve-port"], d, "resolve_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["resolve-port"], "LogSetting-ResolvePort"); ok {
			if err = d.Set("resolve_port", vv); err != nil {
				return fmt.Errorf("Error reading resolve_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resolve_port: %v", err)
		}
	}

	if err = d.Set("rest_api_get", flattenLogSettingRestApiGet(o["rest-api-get"], d, "rest_api_get")); err != nil {
		if vv, ok := fortiAPIPatch(o["rest-api-get"], "LogSetting-RestApiGet"); ok {
			if err = d.Set("rest_api_get", vv); err != nil {
				return fmt.Errorf("Error reading rest_api_get: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rest_api_get: %v", err)
		}
	}

	if err = d.Set("rest_api_set", flattenLogSettingRestApiSet(o["rest-api-set"], d, "rest_api_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["rest-api-set"], "LogSetting-RestApiSet"); ok {
			if err = d.Set("rest_api_set", vv); err != nil {
				return fmt.Errorf("Error reading rest_api_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rest_api_set: %v", err)
		}
	}

	if err = d.Set("syslog_override", flattenLogSettingSyslogOverride(o["syslog-override"], d, "syslog_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["syslog-override"], "LogSetting-SyslogOverride"); ok {
			if err = d.Set("syslog_override", vv); err != nil {
				return fmt.Errorf("Error reading syslog_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading syslog_override: %v", err)
		}
	}

	if err = d.Set("user_anonymize", flattenLogSettingUserAnonymize(o["user-anonymize"], d, "user_anonymize")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-anonymize"], "LogSetting-UserAnonymize"); ok {
			if err = d.Set("user_anonymize", vv); err != nil {
				return fmt.Errorf("Error reading user_anonymize: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_anonymize: %v", err)
		}
	}

	return nil
}

func flattenLogSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogSettingAnonymizationHash(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingBriefTrafficFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingCustomLogFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogSettingDaemonLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingExpolicyImplicitLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingExtendedUtmLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingFazOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingFortiviewWeeklyData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingFwpolicyImplicitLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingFwpolicy6ImplicitLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLocalInAllow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLocalInDenyBroadcast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLocalInDenyUnicast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLocalInPolicyLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLocalOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLocalOutIocDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLogInvalidPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLogPolicyComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLogPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLogUserInUpper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLongLiveSessionStat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingNeighborEvent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingResolveIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingResolvePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingRestApiGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingRestApiSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingSyslogOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSettingUserAnonymize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anonymization_hash"); ok || d.HasChange("anonymization_hash") {
		t, err := expandLogSettingAnonymizationHash(d, v, "anonymization_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anonymization-hash"] = t
		}
	}

	if v, ok := d.GetOk("brief_traffic_format"); ok || d.HasChange("brief_traffic_format") {
		t, err := expandLogSettingBriefTrafficFormat(d, v, "brief_traffic_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["brief-traffic-format"] = t
		}
	}

	if v, ok := d.GetOk("custom_log_fields"); ok || d.HasChange("custom_log_fields") {
		t, err := expandLogSettingCustomLogFields(d, v, "custom_log_fields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-log-fields"] = t
		}
	}

	if v, ok := d.GetOk("daemon_log"); ok || d.HasChange("daemon_log") {
		t, err := expandLogSettingDaemonLog(d, v, "daemon_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["daemon-log"] = t
		}
	}

	if v, ok := d.GetOk("expolicy_implicit_log"); ok || d.HasChange("expolicy_implicit_log") {
		t, err := expandLogSettingExpolicyImplicitLog(d, v, "expolicy_implicit_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expolicy-implicit-log"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandLogSettingExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("extended_utm_log"); ok || d.HasChange("extended_utm_log") {
		t, err := expandLogSettingExtendedUtmLog(d, v, "extended_utm_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-utm-log"] = t
		}
	}

	if v, ok := d.GetOk("faz_override"); ok || d.HasChange("faz_override") {
		t, err := expandLogSettingFazOverride(d, v, "faz_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-override"] = t
		}
	}

	if v, ok := d.GetOk("fortiview_weekly_data"); ok || d.HasChange("fortiview_weekly_data") {
		t, err := expandLogSettingFortiviewWeeklyData(d, v, "fortiview_weekly_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiview-weekly-data"] = t
		}
	}

	if v, ok := d.GetOk("fwpolicy_implicit_log"); ok || d.HasChange("fwpolicy_implicit_log") {
		t, err := expandLogSettingFwpolicyImplicitLog(d, v, "fwpolicy_implicit_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwpolicy-implicit-log"] = t
		}
	}

	if v, ok := d.GetOk("fwpolicy6_implicit_log"); ok || d.HasChange("fwpolicy6_implicit_log") {
		t, err := expandLogSettingFwpolicy6ImplicitLog(d, v, "fwpolicy6_implicit_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwpolicy6-implicit-log"] = t
		}
	}

	if v, ok := d.GetOk("local_in_allow"); ok || d.HasChange("local_in_allow") {
		t, err := expandLogSettingLocalInAllow(d, v, "local_in_allow")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-in-allow"] = t
		}
	}

	if v, ok := d.GetOk("local_in_deny_broadcast"); ok || d.HasChange("local_in_deny_broadcast") {
		t, err := expandLogSettingLocalInDenyBroadcast(d, v, "local_in_deny_broadcast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-in-deny-broadcast"] = t
		}
	}

	if v, ok := d.GetOk("local_in_deny_unicast"); ok || d.HasChange("local_in_deny_unicast") {
		t, err := expandLogSettingLocalInDenyUnicast(d, v, "local_in_deny_unicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-in-deny-unicast"] = t
		}
	}

	if v, ok := d.GetOk("local_in_policy_log"); ok || d.HasChange("local_in_policy_log") {
		t, err := expandLogSettingLocalInPolicyLog(d, v, "local_in_policy_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-in-policy-log"] = t
		}
	}

	if v, ok := d.GetOk("local_out"); ok || d.HasChange("local_out") {
		t, err := expandLogSettingLocalOut(d, v, "local_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-out"] = t
		}
	}

	if v, ok := d.GetOk("local_out_ioc_detection"); ok || d.HasChange("local_out_ioc_detection") {
		t, err := expandLogSettingLocalOutIocDetection(d, v, "local_out_ioc_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-out-ioc-detection"] = t
		}
	}

	if v, ok := d.GetOk("log_invalid_packet"); ok || d.HasChange("log_invalid_packet") {
		t, err := expandLogSettingLogInvalidPacket(d, v, "log_invalid_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-invalid-packet"] = t
		}
	}

	if v, ok := d.GetOk("log_policy_comment"); ok || d.HasChange("log_policy_comment") {
		t, err := expandLogSettingLogPolicyComment(d, v, "log_policy_comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-policy-comment"] = t
		}
	}

	if v, ok := d.GetOk("log_policy_name"); ok || d.HasChange("log_policy_name") {
		t, err := expandLogSettingLogPolicyName(d, v, "log_policy_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-policy-name"] = t
		}
	}

	if v, ok := d.GetOk("log_user_in_upper"); ok || d.HasChange("log_user_in_upper") {
		t, err := expandLogSettingLogUserInUpper(d, v, "log_user_in_upper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-user-in-upper"] = t
		}
	}

	if v, ok := d.GetOk("long_live_session_stat"); ok || d.HasChange("long_live_session_stat") {
		t, err := expandLogSettingLongLiveSessionStat(d, v, "long_live_session_stat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["long-live-session-stat"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_event"); ok || d.HasChange("neighbor_event") {
		t, err := expandLogSettingNeighborEvent(d, v, "neighbor_event")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-event"] = t
		}
	}

	if v, ok := d.GetOk("resolve_ip"); ok || d.HasChange("resolve_ip") {
		t, err := expandLogSettingResolveIp(d, v, "resolve_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resolve-ip"] = t
		}
	}

	if v, ok := d.GetOk("resolve_port"); ok || d.HasChange("resolve_port") {
		t, err := expandLogSettingResolvePort(d, v, "resolve_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resolve-port"] = t
		}
	}

	if v, ok := d.GetOk("rest_api_get"); ok || d.HasChange("rest_api_get") {
		t, err := expandLogSettingRestApiGet(d, v, "rest_api_get")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-api-get"] = t
		}
	}

	if v, ok := d.GetOk("rest_api_set"); ok || d.HasChange("rest_api_set") {
		t, err := expandLogSettingRestApiSet(d, v, "rest_api_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-api-set"] = t
		}
	}

	if v, ok := d.GetOk("syslog_override"); ok || d.HasChange("syslog_override") {
		t, err := expandLogSettingSyslogOverride(d, v, "syslog_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syslog-override"] = t
		}
	}

	if v, ok := d.GetOk("user_anonymize"); ok || d.HasChange("user_anonymize") {
		t, err := expandLogSettingUserAnonymize(d, v, "user_anonymize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-anonymize"] = t
		}
	}

	return &obj, nil
}
