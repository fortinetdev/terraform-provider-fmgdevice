// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: PIM sparse-mode global settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticastPimSmGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticastPimSmGlobalUpdate,
		Read:   resourceRouterMulticastPimSmGlobalRead,
		Update: resourceRouterMulticastPimSmGlobalUpdate,
		Delete: resourceRouterMulticastPimSmGlobalDelete,

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
			"accept_register_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"accept_source_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"bsr_allow_quick_refresh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bsr_candidate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bsr_hash": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bsr_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"bsr_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cisco_crp_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cisco_ignore_rp_set_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cisco_register_checksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cisco_register_checksum_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"join_prune_holdtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"message_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"null_register_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pim_use_sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"register_rate_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"register_rp_reachability": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"register_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"register_source_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"register_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"register_supression": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rp_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"rp_register_keepalive": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"spt_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spt_threshold_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssm_range": &schema.Schema{
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

func resourceRouterMulticastPimSmGlobalUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterMulticastPimSmGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastPimSmGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticastPimSmGlobal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastPimSmGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterMulticastPimSmGlobal")

	return resourceRouterMulticastPimSmGlobalRead(d, m)
}

func resourceRouterMulticastPimSmGlobalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterMulticastPimSmGlobal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticastPimSmGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastPimSmGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterMulticastPimSmGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastPimSmGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticastPimSmGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastPimSmGlobal resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticastPimSmGlobalAcceptRegisterList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalAcceptSourceList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalBsrAllowQuickRefresh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalBsrCandidate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalBsrHash2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalBsrInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalBsrPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalCiscoCrpPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalCiscoRegisterChecksum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalJoinPruneHoldtime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalMessageInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalNullRegisterRetries2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalPimUseSdwan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterRateLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterRpReachability2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterSourceInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalRegisterSourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterSupression2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRpAddress2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {
			v := flattenRouterMulticastPimSmGlobalRpAddressGroup2edl(i["group"], d, pre_append)
			tmp["group"] = fortiAPISubPartPatch(v, "RouterMulticastPimSmGlobal-RpAddress-Group")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterMulticastPimSmGlobalRpAddressId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterMulticastPimSmGlobal-RpAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := i["ip-address"]; ok {
			v := flattenRouterMulticastPimSmGlobalRpAddressIpAddress2edl(i["ip-address"], d, pre_append)
			tmp["ip_address"] = fortiAPISubPartPatch(v, "RouterMulticastPimSmGlobal-RpAddress-IpAddress")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterMulticastPimSmGlobalRpAddressGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalRpAddressId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRpAddressIpAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRpRegisterKeepalive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalSptThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalSptThresholdGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalSsm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalSsmRange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectRouterMulticastPimSmGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("accept_register_list", flattenRouterMulticastPimSmGlobalAcceptRegisterList2edl(o["accept-register-list"], d, "accept_register_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["accept-register-list"], "RouterMulticastPimSmGlobal-AcceptRegisterList"); ok {
			if err = d.Set("accept_register_list", vv); err != nil {
				return fmt.Errorf("Error reading accept_register_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accept_register_list: %v", err)
		}
	}

	if err = d.Set("accept_source_list", flattenRouterMulticastPimSmGlobalAcceptSourceList2edl(o["accept-source-list"], d, "accept_source_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["accept-source-list"], "RouterMulticastPimSmGlobal-AcceptSourceList"); ok {
			if err = d.Set("accept_source_list", vv); err != nil {
				return fmt.Errorf("Error reading accept_source_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accept_source_list: %v", err)
		}
	}

	if err = d.Set("bsr_allow_quick_refresh", flattenRouterMulticastPimSmGlobalBsrAllowQuickRefresh2edl(o["bsr-allow-quick-refresh"], d, "bsr_allow_quick_refresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-allow-quick-refresh"], "RouterMulticastPimSmGlobal-BsrAllowQuickRefresh"); ok {
			if err = d.Set("bsr_allow_quick_refresh", vv); err != nil {
				return fmt.Errorf("Error reading bsr_allow_quick_refresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_allow_quick_refresh: %v", err)
		}
	}

	if err = d.Set("bsr_candidate", flattenRouterMulticastPimSmGlobalBsrCandidate2edl(o["bsr-candidate"], d, "bsr_candidate")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-candidate"], "RouterMulticastPimSmGlobal-BsrCandidate"); ok {
			if err = d.Set("bsr_candidate", vv); err != nil {
				return fmt.Errorf("Error reading bsr_candidate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_candidate: %v", err)
		}
	}

	if err = d.Set("bsr_hash", flattenRouterMulticastPimSmGlobalBsrHash2edl(o["bsr-hash"], d, "bsr_hash")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-hash"], "RouterMulticastPimSmGlobal-BsrHash"); ok {
			if err = d.Set("bsr_hash", vv); err != nil {
				return fmt.Errorf("Error reading bsr_hash: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_hash: %v", err)
		}
	}

	if err = d.Set("bsr_interface", flattenRouterMulticastPimSmGlobalBsrInterface2edl(o["bsr-interface"], d, "bsr_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-interface"], "RouterMulticastPimSmGlobal-BsrInterface"); ok {
			if err = d.Set("bsr_interface", vv); err != nil {
				return fmt.Errorf("Error reading bsr_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_interface: %v", err)
		}
	}

	if err = d.Set("bsr_priority", flattenRouterMulticastPimSmGlobalBsrPriority2edl(o["bsr-priority"], d, "bsr_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["bsr-priority"], "RouterMulticastPimSmGlobal-BsrPriority"); ok {
			if err = d.Set("bsr_priority", vv); err != nil {
				return fmt.Errorf("Error reading bsr_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bsr_priority: %v", err)
		}
	}

	if err = d.Set("cisco_crp_prefix", flattenRouterMulticastPimSmGlobalCiscoCrpPrefix2edl(o["cisco-crp-prefix"], d, "cisco_crp_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["cisco-crp-prefix"], "RouterMulticastPimSmGlobal-CiscoCrpPrefix"); ok {
			if err = d.Set("cisco_crp_prefix", vv); err != nil {
				return fmt.Errorf("Error reading cisco_crp_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cisco_crp_prefix: %v", err)
		}
	}

	if err = d.Set("cisco_ignore_rp_set_priority", flattenRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority2edl(o["cisco-ignore-rp-set-priority"], d, "cisco_ignore_rp_set_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["cisco-ignore-rp-set-priority"], "RouterMulticastPimSmGlobal-CiscoIgnoreRpSetPriority"); ok {
			if err = d.Set("cisco_ignore_rp_set_priority", vv); err != nil {
				return fmt.Errorf("Error reading cisco_ignore_rp_set_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cisco_ignore_rp_set_priority: %v", err)
		}
	}

	if err = d.Set("cisco_register_checksum", flattenRouterMulticastPimSmGlobalCiscoRegisterChecksum2edl(o["cisco-register-checksum"], d, "cisco_register_checksum")); err != nil {
		if vv, ok := fortiAPIPatch(o["cisco-register-checksum"], "RouterMulticastPimSmGlobal-CiscoRegisterChecksum"); ok {
			if err = d.Set("cisco_register_checksum", vv); err != nil {
				return fmt.Errorf("Error reading cisco_register_checksum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cisco_register_checksum: %v", err)
		}
	}

	if err = d.Set("cisco_register_checksum_group", flattenRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup2edl(o["cisco-register-checksum-group"], d, "cisco_register_checksum_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["cisco-register-checksum-group"], "RouterMulticastPimSmGlobal-CiscoRegisterChecksumGroup"); ok {
			if err = d.Set("cisco_register_checksum_group", vv); err != nil {
				return fmt.Errorf("Error reading cisco_register_checksum_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cisco_register_checksum_group: %v", err)
		}
	}

	if err = d.Set("join_prune_holdtime", flattenRouterMulticastPimSmGlobalJoinPruneHoldtime2edl(o["join-prune-holdtime"], d, "join_prune_holdtime")); err != nil {
		if vv, ok := fortiAPIPatch(o["join-prune-holdtime"], "RouterMulticastPimSmGlobal-JoinPruneHoldtime"); ok {
			if err = d.Set("join_prune_holdtime", vv); err != nil {
				return fmt.Errorf("Error reading join_prune_holdtime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading join_prune_holdtime: %v", err)
		}
	}

	if err = d.Set("message_interval", flattenRouterMulticastPimSmGlobalMessageInterval2edl(o["message-interval"], d, "message_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-interval"], "RouterMulticastPimSmGlobal-MessageInterval"); ok {
			if err = d.Set("message_interval", vv); err != nil {
				return fmt.Errorf("Error reading message_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_interval: %v", err)
		}
	}

	if err = d.Set("null_register_retries", flattenRouterMulticastPimSmGlobalNullRegisterRetries2edl(o["null-register-retries"], d, "null_register_retries")); err != nil {
		if vv, ok := fortiAPIPatch(o["null-register-retries"], "RouterMulticastPimSmGlobal-NullRegisterRetries"); ok {
			if err = d.Set("null_register_retries", vv); err != nil {
				return fmt.Errorf("Error reading null_register_retries: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading null_register_retries: %v", err)
		}
	}

	if err = d.Set("pim_use_sdwan", flattenRouterMulticastPimSmGlobalPimUseSdwan2edl(o["pim-use-sdwan"], d, "pim_use_sdwan")); err != nil {
		if vv, ok := fortiAPIPatch(o["pim-use-sdwan"], "RouterMulticastPimSmGlobal-PimUseSdwan"); ok {
			if err = d.Set("pim_use_sdwan", vv); err != nil {
				return fmt.Errorf("Error reading pim_use_sdwan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pim_use_sdwan: %v", err)
		}
	}

	if err = d.Set("register_rate_limit", flattenRouterMulticastPimSmGlobalRegisterRateLimit2edl(o["register-rate-limit"], d, "register_rate_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-rate-limit"], "RouterMulticastPimSmGlobal-RegisterRateLimit"); ok {
			if err = d.Set("register_rate_limit", vv); err != nil {
				return fmt.Errorf("Error reading register_rate_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_rate_limit: %v", err)
		}
	}

	if err = d.Set("register_rp_reachability", flattenRouterMulticastPimSmGlobalRegisterRpReachability2edl(o["register-rp-reachability"], d, "register_rp_reachability")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-rp-reachability"], "RouterMulticastPimSmGlobal-RegisterRpReachability"); ok {
			if err = d.Set("register_rp_reachability", vv); err != nil {
				return fmt.Errorf("Error reading register_rp_reachability: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_rp_reachability: %v", err)
		}
	}

	if err = d.Set("register_source", flattenRouterMulticastPimSmGlobalRegisterSource2edl(o["register-source"], d, "register_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-source"], "RouterMulticastPimSmGlobal-RegisterSource"); ok {
			if err = d.Set("register_source", vv); err != nil {
				return fmt.Errorf("Error reading register_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_source: %v", err)
		}
	}

	if err = d.Set("register_source_interface", flattenRouterMulticastPimSmGlobalRegisterSourceInterface2edl(o["register-source-interface"], d, "register_source_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-source-interface"], "RouterMulticastPimSmGlobal-RegisterSourceInterface"); ok {
			if err = d.Set("register_source_interface", vv); err != nil {
				return fmt.Errorf("Error reading register_source_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_source_interface: %v", err)
		}
	}

	if err = d.Set("register_source_ip", flattenRouterMulticastPimSmGlobalRegisterSourceIp2edl(o["register-source-ip"], d, "register_source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-source-ip"], "RouterMulticastPimSmGlobal-RegisterSourceIp"); ok {
			if err = d.Set("register_source_ip", vv); err != nil {
				return fmt.Errorf("Error reading register_source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_source_ip: %v", err)
		}
	}

	if err = d.Set("register_supression", flattenRouterMulticastPimSmGlobalRegisterSupression2edl(o["register-supression"], d, "register_supression")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-supression"], "RouterMulticastPimSmGlobal-RegisterSupression"); ok {
			if err = d.Set("register_supression", vv); err != nil {
				return fmt.Errorf("Error reading register_supression: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_supression: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rp_address", flattenRouterMulticastPimSmGlobalRpAddress2edl(o["rp-address"], d, "rp_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["rp-address"], "RouterMulticastPimSmGlobal-RpAddress"); ok {
				if err = d.Set("rp_address", vv); err != nil {
					return fmt.Errorf("Error reading rp_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rp_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rp_address"); ok {
			if err = d.Set("rp_address", flattenRouterMulticastPimSmGlobalRpAddress2edl(o["rp-address"], d, "rp_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["rp-address"], "RouterMulticastPimSmGlobal-RpAddress"); ok {
					if err = d.Set("rp_address", vv); err != nil {
						return fmt.Errorf("Error reading rp_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rp_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("rp_register_keepalive", flattenRouterMulticastPimSmGlobalRpRegisterKeepalive2edl(o["rp-register-keepalive"], d, "rp_register_keepalive")); err != nil {
		if vv, ok := fortiAPIPatch(o["rp-register-keepalive"], "RouterMulticastPimSmGlobal-RpRegisterKeepalive"); ok {
			if err = d.Set("rp_register_keepalive", vv); err != nil {
				return fmt.Errorf("Error reading rp_register_keepalive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rp_register_keepalive: %v", err)
		}
	}

	if err = d.Set("spt_threshold", flattenRouterMulticastPimSmGlobalSptThreshold2edl(o["spt-threshold"], d, "spt_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["spt-threshold"], "RouterMulticastPimSmGlobal-SptThreshold"); ok {
			if err = d.Set("spt_threshold", vv); err != nil {
				return fmt.Errorf("Error reading spt_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spt_threshold: %v", err)
		}
	}

	if err = d.Set("spt_threshold_group", flattenRouterMulticastPimSmGlobalSptThresholdGroup2edl(o["spt-threshold-group"], d, "spt_threshold_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["spt-threshold-group"], "RouterMulticastPimSmGlobal-SptThresholdGroup"); ok {
			if err = d.Set("spt_threshold_group", vv); err != nil {
				return fmt.Errorf("Error reading spt_threshold_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spt_threshold_group: %v", err)
		}
	}

	if err = d.Set("ssm", flattenRouterMulticastPimSmGlobalSsm2edl(o["ssm"], d, "ssm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssm"], "RouterMulticastPimSmGlobal-Ssm"); ok {
			if err = d.Set("ssm", vv); err != nil {
				return fmt.Errorf("Error reading ssm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssm: %v", err)
		}
	}

	if err = d.Set("ssm_range", flattenRouterMulticastPimSmGlobalSsmRange2edl(o["ssm-range"], d, "ssm_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssm-range"], "RouterMulticastPimSmGlobal-SsmRange"); ok {
			if err = d.Set("ssm_range", vv); err != nil {
				return fmt.Errorf("Error reading ssm_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssm_range: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticastPimSmGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticastPimSmGlobalAcceptRegisterList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalAcceptSourceList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalBsrAllowQuickRefresh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalBsrCandidate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalBsrHash2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalBsrInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalBsrPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalCiscoCrpPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalCiscoRegisterChecksum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalJoinPruneHoldtime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalMessageInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalNullRegisterRetries2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalPimUseSdwan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterRateLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterRpReachability2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterSourceInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalRegisterSourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterSupression2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRpAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group"], _ = expandRouterMulticastPimSmGlobalRpAddressGroup2edl(d, i["group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterMulticastPimSmGlobalRpAddressId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-address"], _ = expandRouterMulticastPimSmGlobalRpAddressIpAddress2edl(d, i["ip_address"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterMulticastPimSmGlobalRpAddressGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalRpAddressId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRpAddressIpAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRpRegisterKeepalive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalSptThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalSptThresholdGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalSsm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalSsmRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectRouterMulticastPimSmGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accept_register_list"); ok || d.HasChange("accept_register_list") {
		t, err := expandRouterMulticastPimSmGlobalAcceptRegisterList2edl(d, v, "accept_register_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accept-register-list"] = t
		}
	}

	if v, ok := d.GetOk("accept_source_list"); ok || d.HasChange("accept_source_list") {
		t, err := expandRouterMulticastPimSmGlobalAcceptSourceList2edl(d, v, "accept_source_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accept-source-list"] = t
		}
	}

	if v, ok := d.GetOk("bsr_allow_quick_refresh"); ok || d.HasChange("bsr_allow_quick_refresh") {
		t, err := expandRouterMulticastPimSmGlobalBsrAllowQuickRefresh2edl(d, v, "bsr_allow_quick_refresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-allow-quick-refresh"] = t
		}
	}

	if v, ok := d.GetOk("bsr_candidate"); ok || d.HasChange("bsr_candidate") {
		t, err := expandRouterMulticastPimSmGlobalBsrCandidate2edl(d, v, "bsr_candidate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-candidate"] = t
		}
	}

	if v, ok := d.GetOk("bsr_hash"); ok || d.HasChange("bsr_hash") {
		t, err := expandRouterMulticastPimSmGlobalBsrHash2edl(d, v, "bsr_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-hash"] = t
		}
	}

	if v, ok := d.GetOk("bsr_interface"); ok || d.HasChange("bsr_interface") {
		t, err := expandRouterMulticastPimSmGlobalBsrInterface2edl(d, v, "bsr_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-interface"] = t
		}
	}

	if v, ok := d.GetOk("bsr_priority"); ok || d.HasChange("bsr_priority") {
		t, err := expandRouterMulticastPimSmGlobalBsrPriority2edl(d, v, "bsr_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bsr-priority"] = t
		}
	}

	if v, ok := d.GetOk("cisco_crp_prefix"); ok || d.HasChange("cisco_crp_prefix") {
		t, err := expandRouterMulticastPimSmGlobalCiscoCrpPrefix2edl(d, v, "cisco_crp_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cisco-crp-prefix"] = t
		}
	}

	if v, ok := d.GetOk("cisco_ignore_rp_set_priority"); ok || d.HasChange("cisco_ignore_rp_set_priority") {
		t, err := expandRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority2edl(d, v, "cisco_ignore_rp_set_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cisco-ignore-rp-set-priority"] = t
		}
	}

	if v, ok := d.GetOk("cisco_register_checksum"); ok || d.HasChange("cisco_register_checksum") {
		t, err := expandRouterMulticastPimSmGlobalCiscoRegisterChecksum2edl(d, v, "cisco_register_checksum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cisco-register-checksum"] = t
		}
	}

	if v, ok := d.GetOk("cisco_register_checksum_group"); ok || d.HasChange("cisco_register_checksum_group") {
		t, err := expandRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup2edl(d, v, "cisco_register_checksum_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cisco-register-checksum-group"] = t
		}
	}

	if v, ok := d.GetOk("join_prune_holdtime"); ok || d.HasChange("join_prune_holdtime") {
		t, err := expandRouterMulticastPimSmGlobalJoinPruneHoldtime2edl(d, v, "join_prune_holdtime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["join-prune-holdtime"] = t
		}
	}

	if v, ok := d.GetOk("message_interval"); ok || d.HasChange("message_interval") {
		t, err := expandRouterMulticastPimSmGlobalMessageInterval2edl(d, v, "message_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-interval"] = t
		}
	}

	if v, ok := d.GetOk("null_register_retries"); ok || d.HasChange("null_register_retries") {
		t, err := expandRouterMulticastPimSmGlobalNullRegisterRetries2edl(d, v, "null_register_retries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["null-register-retries"] = t
		}
	}

	if v, ok := d.GetOk("pim_use_sdwan"); ok || d.HasChange("pim_use_sdwan") {
		t, err := expandRouterMulticastPimSmGlobalPimUseSdwan2edl(d, v, "pim_use_sdwan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pim-use-sdwan"] = t
		}
	}

	if v, ok := d.GetOk("register_rate_limit"); ok || d.HasChange("register_rate_limit") {
		t, err := expandRouterMulticastPimSmGlobalRegisterRateLimit2edl(d, v, "register_rate_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-rate-limit"] = t
		}
	}

	if v, ok := d.GetOk("register_rp_reachability"); ok || d.HasChange("register_rp_reachability") {
		t, err := expandRouterMulticastPimSmGlobalRegisterRpReachability2edl(d, v, "register_rp_reachability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-rp-reachability"] = t
		}
	}

	if v, ok := d.GetOk("register_source"); ok || d.HasChange("register_source") {
		t, err := expandRouterMulticastPimSmGlobalRegisterSource2edl(d, v, "register_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-source"] = t
		}
	}

	if v, ok := d.GetOk("register_source_interface"); ok || d.HasChange("register_source_interface") {
		t, err := expandRouterMulticastPimSmGlobalRegisterSourceInterface2edl(d, v, "register_source_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-source-interface"] = t
		}
	}

	if v, ok := d.GetOk("register_source_ip"); ok || d.HasChange("register_source_ip") {
		t, err := expandRouterMulticastPimSmGlobalRegisterSourceIp2edl(d, v, "register_source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-source-ip"] = t
		}
	}

	if v, ok := d.GetOk("register_supression"); ok || d.HasChange("register_supression") {
		t, err := expandRouterMulticastPimSmGlobalRegisterSupression2edl(d, v, "register_supression")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-supression"] = t
		}
	}

	if v, ok := d.GetOk("rp_address"); ok || d.HasChange("rp_address") {
		t, err := expandRouterMulticastPimSmGlobalRpAddress2edl(d, v, "rp_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rp-address"] = t
		}
	}

	if v, ok := d.GetOk("rp_register_keepalive"); ok || d.HasChange("rp_register_keepalive") {
		t, err := expandRouterMulticastPimSmGlobalRpRegisterKeepalive2edl(d, v, "rp_register_keepalive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rp-register-keepalive"] = t
		}
	}

	if v, ok := d.GetOk("spt_threshold"); ok || d.HasChange("spt_threshold") {
		t, err := expandRouterMulticastPimSmGlobalSptThreshold2edl(d, v, "spt_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spt-threshold"] = t
		}
	}

	if v, ok := d.GetOk("spt_threshold_group"); ok || d.HasChange("spt_threshold_group") {
		t, err := expandRouterMulticastPimSmGlobalSptThresholdGroup2edl(d, v, "spt_threshold_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spt-threshold-group"] = t
		}
	}

	if v, ok := d.GetOk("ssm"); ok || d.HasChange("ssm") {
		t, err := expandRouterMulticastPimSmGlobalSsm2edl(d, v, "ssm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssm"] = t
		}
	}

	if v, ok := d.GetOk("ssm_range"); ok || d.HasChange("ssm_range") {
		t, err := expandRouterMulticastPimSmGlobalSsmRange2edl(d, v, "ssm_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssm-range"] = t
		}
	}

	return &obj, nil
}
