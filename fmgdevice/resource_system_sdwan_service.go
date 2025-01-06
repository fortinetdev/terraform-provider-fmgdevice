// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdwanService() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdwanServiceCreate,
		Read:   resourceSystemSdwanServiceRead,
		Update: resourceSystemSdwanServiceUpdate,
		Delete: resourceSystemSdwanServiceDelete,

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
			"addr_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"agent_exclusive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp_forward_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dscp_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp_reverse_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dst_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"end_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"end_src_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"hash_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"hold_down_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"input_device": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"input_device_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"input_zone": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_app_ctrl": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"internet_service_app_ctrl_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"internet_service_app_ctrl_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"jitter_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"link_cost_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"link_cost_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"load_balance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"minimum_sla_meet_members": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"packet_loss_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"passive_measurement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_members": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"priority_zone": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"quality_link": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_tag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"shortcut": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shortcut_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sla": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"sla_compare_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sla_stickiness": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"standalone_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"start_src_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tie_break": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_shortcut_sla": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"zone_mode": &schema.Schema{
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

func resourceSystemSdwanServiceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdwanService(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanService resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSdwanService(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanService resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSdwanServiceRead(d, m)
}

func resourceSystemSdwanServiceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdwanService(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanService resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSdwanService(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSdwanServiceRead(d, m)
}

func resourceSystemSdwanServiceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSdwanService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdwanService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanServiceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdwanService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdwanService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanService resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdwanServiceAddrMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceAgentExclusive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceBandwidthWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDefault2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDscpForward2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDscpForwardTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDscpReverse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDscpReverseTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceDstNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDst62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceEndPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceEndSrcPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceGroups2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceHashMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceHealthCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceHoldDownTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceInputDevice2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceInputDeviceNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceInputZone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceInternetService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceInternetServiceAppCtrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemSdwanServiceInternetServiceAppCtrlCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemSdwanServiceInternetServiceAppCtrlGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceInternetServiceCustom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceInternetServiceCustomGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceInternetServiceGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceInternetServiceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceJitterWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceLatencyWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceLinkCostFactor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceLinkCostThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceLoadBalance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceMinimumSlaMeetMembers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServicePacketLossWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServicePassiveMeasurement2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServicePriorityMembers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServicePriorityZone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceQualityLink2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceRole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceRouteTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceShortcut2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceShortcutPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceSla2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenSystemSdwanServiceSlaHealthCheck2edl(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "SystemSdwanService-Sla-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSdwanServiceSlaId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSdwanService-Sla-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceSlaHealthCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceSlaId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceSlaCompareMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceSlaStickiness2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceSrc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceSrcNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceSrc62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceStandaloneAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceStartPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceStartSrcPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceTieBreak2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceTos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceTosMask2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceUseShortcutSla2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceUsers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceZoneMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdwanService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("addr_mode", flattenSystemSdwanServiceAddrMode2edl(o["addr-mode"], d, "addr_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-mode"], "SystemSdwanService-AddrMode"); ok {
			if err = d.Set("addr_mode", vv); err != nil {
				return fmt.Errorf("Error reading addr_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("agent_exclusive", flattenSystemSdwanServiceAgentExclusive2edl(o["agent-exclusive"], d, "agent_exclusive")); err != nil {
		if vv, ok := fortiAPIPatch(o["agent-exclusive"], "SystemSdwanService-AgentExclusive"); ok {
			if err = d.Set("agent_exclusive", vv); err != nil {
				return fmt.Errorf("Error reading agent_exclusive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agent_exclusive: %v", err)
		}
	}

	if err = d.Set("bandwidth_weight", flattenSystemSdwanServiceBandwidthWeight2edl(o["bandwidth-weight"], d, "bandwidth_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-weight"], "SystemSdwanService-BandwidthWeight"); ok {
			if err = d.Set("bandwidth_weight", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_weight: %v", err)
		}
	}

	if err = d.Set("comment", flattenSystemSdwanServiceComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "SystemSdwanService-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("default", flattenSystemSdwanServiceDefault2edl(o["default"], d, "default")); err != nil {
		if vv, ok := fortiAPIPatch(o["default"], "SystemSdwanService-Default"); ok {
			if err = d.Set("default", vv); err != nil {
				return fmt.Errorf("Error reading default: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default: %v", err)
		}
	}

	if err = d.Set("dscp_forward", flattenSystemSdwanServiceDscpForward2edl(o["dscp-forward"], d, "dscp_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-forward"], "SystemSdwanService-DscpForward"); ok {
			if err = d.Set("dscp_forward", vv); err != nil {
				return fmt.Errorf("Error reading dscp_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_forward: %v", err)
		}
	}

	if err = d.Set("dscp_forward_tag", flattenSystemSdwanServiceDscpForwardTag2edl(o["dscp-forward-tag"], d, "dscp_forward_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-forward-tag"], "SystemSdwanService-DscpForwardTag"); ok {
			if err = d.Set("dscp_forward_tag", vv); err != nil {
				return fmt.Errorf("Error reading dscp_forward_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_forward_tag: %v", err)
		}
	}

	if err = d.Set("dscp_reverse", flattenSystemSdwanServiceDscpReverse2edl(o["dscp-reverse"], d, "dscp_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-reverse"], "SystemSdwanService-DscpReverse"); ok {
			if err = d.Set("dscp_reverse", vv); err != nil {
				return fmt.Errorf("Error reading dscp_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_reverse: %v", err)
		}
	}

	if err = d.Set("dscp_reverse_tag", flattenSystemSdwanServiceDscpReverseTag2edl(o["dscp-reverse-tag"], d, "dscp_reverse_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-reverse-tag"], "SystemSdwanService-DscpReverseTag"); ok {
			if err = d.Set("dscp_reverse_tag", vv); err != nil {
				return fmt.Errorf("Error reading dscp_reverse_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_reverse_tag: %v", err)
		}
	}

	if err = d.Set("dst", flattenSystemSdwanServiceDst2edl(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "SystemSdwanService-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("dst_negate", flattenSystemSdwanServiceDstNegate2edl(o["dst-negate"], d, "dst_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-negate"], "SystemSdwanService-DstNegate"); ok {
			if err = d.Set("dst_negate", vv); err != nil {
				return fmt.Errorf("Error reading dst_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_negate: %v", err)
		}
	}

	if err = d.Set("dst6", flattenSystemSdwanServiceDst62edl(o["dst6"], d, "dst6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst6"], "SystemSdwanService-Dst6"); ok {
			if err = d.Set("dst6", vv); err != nil {
				return fmt.Errorf("Error reading dst6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst6: %v", err)
		}
	}

	if err = d.Set("end_port", flattenSystemSdwanServiceEndPort2edl(o["end-port"], d, "end_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-port"], "SystemSdwanService-EndPort"); ok {
			if err = d.Set("end_port", vv); err != nil {
				return fmt.Errorf("Error reading end_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("end_src_port", flattenSystemSdwanServiceEndSrcPort2edl(o["end-src-port"], d, "end_src_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-src-port"], "SystemSdwanService-EndSrcPort"); ok {
			if err = d.Set("end_src_port", vv); err != nil {
				return fmt.Errorf("Error reading end_src_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_src_port: %v", err)
		}
	}

	if err = d.Set("gateway", flattenSystemSdwanServiceGateway2edl(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "SystemSdwanService-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("groups", flattenSystemSdwanServiceGroups2edl(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "SystemSdwanService-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("hash_mode", flattenSystemSdwanServiceHashMode2edl(o["hash-mode"], d, "hash_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["hash-mode"], "SystemSdwanService-HashMode"); ok {
			if err = d.Set("hash_mode", vv); err != nil {
				return fmt.Errorf("Error reading hash_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hash_mode: %v", err)
		}
	}

	if err = d.Set("health_check", flattenSystemSdwanServiceHealthCheck2edl(o["health-check"], d, "health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check"], "SystemSdwanService-HealthCheck"); ok {
			if err = d.Set("health_check", vv); err != nil {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("hold_down_time", flattenSystemSdwanServiceHoldDownTime2edl(o["hold-down-time"], d, "hold_down_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["hold-down-time"], "SystemSdwanService-HoldDownTime"); ok {
			if err = d.Set("hold_down_time", vv); err != nil {
				return fmt.Errorf("Error reading hold_down_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hold_down_time: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemSdwanServiceId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSdwanService-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("input_device", flattenSystemSdwanServiceInputDevice2edl(o["input-device"], d, "input_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["input-device"], "SystemSdwanService-InputDevice"); ok {
			if err = d.Set("input_device", vv); err != nil {
				return fmt.Errorf("Error reading input_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading input_device: %v", err)
		}
	}

	if err = d.Set("input_device_negate", flattenSystemSdwanServiceInputDeviceNegate2edl(o["input-device-negate"], d, "input_device_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["input-device-negate"], "SystemSdwanService-InputDeviceNegate"); ok {
			if err = d.Set("input_device_negate", vv); err != nil {
				return fmt.Errorf("Error reading input_device_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading input_device_negate: %v", err)
		}
	}

	if err = d.Set("input_zone", flattenSystemSdwanServiceInputZone2edl(o["input-zone"], d, "input_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["input-zone"], "SystemSdwanService-InputZone"); ok {
			if err = d.Set("input_zone", vv); err != nil {
				return fmt.Errorf("Error reading input_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading input_zone: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenSystemSdwanServiceInternetService2edl(o["internet-service"], d, "internet_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service"], "SystemSdwanService-InternetService"); ok {
			if err = d.Set("internet_service", vv); err != nil {
				return fmt.Errorf("Error reading internet_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_app_ctrl", flattenSystemSdwanServiceInternetServiceAppCtrl2edl(o["internet-service-app-ctrl"], d, "internet_service_app_ctrl")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-app-ctrl"], "SystemSdwanService-InternetServiceAppCtrl"); ok {
			if err = d.Set("internet_service_app_ctrl", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_app_ctrl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_app_ctrl: %v", err)
		}
	}

	if err = d.Set("internet_service_app_ctrl_category", flattenSystemSdwanServiceInternetServiceAppCtrlCategory2edl(o["internet-service-app-ctrl-category"], d, "internet_service_app_ctrl_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-app-ctrl-category"], "SystemSdwanService-InternetServiceAppCtrlCategory"); ok {
			if err = d.Set("internet_service_app_ctrl_category", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_app_ctrl_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_app_ctrl_category: %v", err)
		}
	}

	if err = d.Set("internet_service_app_ctrl_group", flattenSystemSdwanServiceInternetServiceAppCtrlGroup2edl(o["internet-service-app-ctrl-group"], d, "internet_service_app_ctrl_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-app-ctrl-group"], "SystemSdwanService-InternetServiceAppCtrlGroup"); ok {
			if err = d.Set("internet_service_app_ctrl_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_app_ctrl_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_app_ctrl_group: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", flattenSystemSdwanServiceInternetServiceCustom2edl(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom"], "SystemSdwanService-InternetServiceCustom"); ok {
			if err = d.Set("internet_service_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_custom_group", flattenSystemSdwanServiceInternetServiceCustomGroup2edl(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom-group"], "SystemSdwanService-InternetServiceCustomGroup"); ok {
			if err = d.Set("internet_service_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_group", flattenSystemSdwanServiceInternetServiceGroup2edl(o["internet-service-group"], d, "internet_service_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-group"], "SystemSdwanService-InternetServiceGroup"); ok {
			if err = d.Set("internet_service_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_group: %v", err)
		}
	}

	if err = d.Set("internet_service_name", flattenSystemSdwanServiceInternetServiceName2edl(o["internet-service-name"], d, "internet_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-name"], "SystemSdwanService-InternetServiceName"); ok {
			if err = d.Set("internet_service_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_name: %v", err)
		}
	}

	if err = d.Set("jitter_weight", flattenSystemSdwanServiceJitterWeight2edl(o["jitter-weight"], d, "jitter_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["jitter-weight"], "SystemSdwanService-JitterWeight"); ok {
			if err = d.Set("jitter_weight", vv); err != nil {
				return fmt.Errorf("Error reading jitter_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jitter_weight: %v", err)
		}
	}

	if err = d.Set("latency_weight", flattenSystemSdwanServiceLatencyWeight2edl(o["latency-weight"], d, "latency_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-weight"], "SystemSdwanService-LatencyWeight"); ok {
			if err = d.Set("latency_weight", vv); err != nil {
				return fmt.Errorf("Error reading latency_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_weight: %v", err)
		}
	}

	if err = d.Set("link_cost_factor", flattenSystemSdwanServiceLinkCostFactor2edl(o["link-cost-factor"], d, "link_cost_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-cost-factor"], "SystemSdwanService-LinkCostFactor"); ok {
			if err = d.Set("link_cost_factor", vv); err != nil {
				return fmt.Errorf("Error reading link_cost_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_cost_factor: %v", err)
		}
	}

	if err = d.Set("link_cost_threshold", flattenSystemSdwanServiceLinkCostThreshold2edl(o["link-cost-threshold"], d, "link_cost_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-cost-threshold"], "SystemSdwanService-LinkCostThreshold"); ok {
			if err = d.Set("link_cost_threshold", vv); err != nil {
				return fmt.Errorf("Error reading link_cost_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_cost_threshold: %v", err)
		}
	}

	if err = d.Set("load_balance", flattenSystemSdwanServiceLoadBalance2edl(o["load-balance"], d, "load_balance")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balance"], "SystemSdwanService-LoadBalance"); ok {
			if err = d.Set("load_balance", vv); err != nil {
				return fmt.Errorf("Error reading load_balance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balance: %v", err)
		}
	}

	if err = d.Set("minimum_sla_meet_members", flattenSystemSdwanServiceMinimumSlaMeetMembers2edl(o["minimum-sla-meet-members"], d, "minimum_sla_meet_members")); err != nil {
		if vv, ok := fortiAPIPatch(o["minimum-sla-meet-members"], "SystemSdwanService-MinimumSlaMeetMembers"); ok {
			if err = d.Set("minimum_sla_meet_members", vv); err != nil {
				return fmt.Errorf("Error reading minimum_sla_meet_members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minimum_sla_meet_members: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemSdwanServiceMode2edl(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemSdwanService-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSdwanServiceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSdwanService-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("packet_loss_weight", flattenSystemSdwanServicePacketLossWeight2edl(o["packet-loss-weight"], d, "packet_loss_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-loss-weight"], "SystemSdwanService-PacketLossWeight"); ok {
			if err = d.Set("packet_loss_weight", vv); err != nil {
				return fmt.Errorf("Error reading packet_loss_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_loss_weight: %v", err)
		}
	}

	if err = d.Set("passive_measurement", flattenSystemSdwanServicePassiveMeasurement2edl(o["passive-measurement"], d, "passive_measurement")); err != nil {
		if vv, ok := fortiAPIPatch(o["passive-measurement"], "SystemSdwanService-PassiveMeasurement"); ok {
			if err = d.Set("passive_measurement", vv); err != nil {
				return fmt.Errorf("Error reading passive_measurement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passive_measurement: %v", err)
		}
	}

	if err = d.Set("priority_members", flattenSystemSdwanServicePriorityMembers2edl(o["priority-members"], d, "priority_members")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-members"], "SystemSdwanService-PriorityMembers"); ok {
			if err = d.Set("priority_members", vv); err != nil {
				return fmt.Errorf("Error reading priority_members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_members: %v", err)
		}
	}

	if err = d.Set("priority_zone", flattenSystemSdwanServicePriorityZone2edl(o["priority-zone"], d, "priority_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-zone"], "SystemSdwanService-PriorityZone"); ok {
			if err = d.Set("priority_zone", vv); err != nil {
				return fmt.Errorf("Error reading priority_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_zone: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemSdwanServiceProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemSdwanService-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("quality_link", flattenSystemSdwanServiceQualityLink2edl(o["quality-link"], d, "quality_link")); err != nil {
		if vv, ok := fortiAPIPatch(o["quality-link"], "SystemSdwanService-QualityLink"); ok {
			if err = d.Set("quality_link", vv); err != nil {
				return fmt.Errorf("Error reading quality_link: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quality_link: %v", err)
		}
	}

	if err = d.Set("role", flattenSystemSdwanServiceRole2edl(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "SystemSdwanService-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("route_tag", flattenSystemSdwanServiceRouteTag2edl(o["route-tag"], d, "route_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-tag"], "SystemSdwanService-RouteTag"); ok {
			if err = d.Set("route_tag", vv); err != nil {
				return fmt.Errorf("Error reading route_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_tag: %v", err)
		}
	}

	if err = d.Set("shortcut", flattenSystemSdwanServiceShortcut2edl(o["shortcut"], d, "shortcut")); err != nil {
		if vv, ok := fortiAPIPatch(o["shortcut"], "SystemSdwanService-Shortcut"); ok {
			if err = d.Set("shortcut", vv); err != nil {
				return fmt.Errorf("Error reading shortcut: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shortcut: %v", err)
		}
	}

	if err = d.Set("shortcut_priority", flattenSystemSdwanServiceShortcutPriority2edl(o["shortcut-priority"], d, "shortcut_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["shortcut-priority"], "SystemSdwanService-ShortcutPriority"); ok {
			if err = d.Set("shortcut_priority", vv); err != nil {
				return fmt.Errorf("Error reading shortcut_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shortcut_priority: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sla", flattenSystemSdwanServiceSla2edl(o["sla"], d, "sla")); err != nil {
			if vv, ok := fortiAPIPatch(o["sla"], "SystemSdwanService-Sla"); ok {
				if err = d.Set("sla", vv); err != nil {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sla: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sla"); ok {
			if err = d.Set("sla", flattenSystemSdwanServiceSla2edl(o["sla"], d, "sla")); err != nil {
				if vv, ok := fortiAPIPatch(o["sla"], "SystemSdwanService-Sla"); ok {
					if err = d.Set("sla", vv); err != nil {
						return fmt.Errorf("Error reading sla: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			}
		}
	}

	if err = d.Set("sla_compare_method", flattenSystemSdwanServiceSlaCompareMethod2edl(o["sla-compare-method"], d, "sla_compare_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-compare-method"], "SystemSdwanService-SlaCompareMethod"); ok {
			if err = d.Set("sla_compare_method", vv); err != nil {
				return fmt.Errorf("Error reading sla_compare_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_compare_method: %v", err)
		}
	}

	if err = d.Set("sla_stickiness", flattenSystemSdwanServiceSlaStickiness2edl(o["sla-stickiness"], d, "sla_stickiness")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-stickiness"], "SystemSdwanService-SlaStickiness"); ok {
			if err = d.Set("sla_stickiness", vv); err != nil {
				return fmt.Errorf("Error reading sla_stickiness: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_stickiness: %v", err)
		}
	}

	if err = d.Set("src", flattenSystemSdwanServiceSrc2edl(o["src"], d, "src")); err != nil {
		if vv, ok := fortiAPIPatch(o["src"], "SystemSdwanService-Src"); ok {
			if err = d.Set("src", vv); err != nil {
				return fmt.Errorf("Error reading src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("src_negate", flattenSystemSdwanServiceSrcNegate2edl(o["src-negate"], d, "src_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-negate"], "SystemSdwanService-SrcNegate"); ok {
			if err = d.Set("src_negate", vv); err != nil {
				return fmt.Errorf("Error reading src_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_negate: %v", err)
		}
	}

	if err = d.Set("src6", flattenSystemSdwanServiceSrc62edl(o["src6"], d, "src6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src6"], "SystemSdwanService-Src6"); ok {
			if err = d.Set("src6", vv); err != nil {
				return fmt.Errorf("Error reading src6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src6: %v", err)
		}
	}

	if err = d.Set("standalone_action", flattenSystemSdwanServiceStandaloneAction2edl(o["standalone-action"], d, "standalone_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["standalone-action"], "SystemSdwanService-StandaloneAction"); ok {
			if err = d.Set("standalone_action", vv); err != nil {
				return fmt.Errorf("Error reading standalone_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading standalone_action: %v", err)
		}
	}

	if err = d.Set("start_port", flattenSystemSdwanServiceStartPort2edl(o["start-port"], d, "start_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-port"], "SystemSdwanService-StartPort"); ok {
			if err = d.Set("start_port", vv); err != nil {
				return fmt.Errorf("Error reading start_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	if err = d.Set("start_src_port", flattenSystemSdwanServiceStartSrcPort2edl(o["start-src-port"], d, "start_src_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-src-port"], "SystemSdwanService-StartSrcPort"); ok {
			if err = d.Set("start_src_port", vv); err != nil {
				return fmt.Errorf("Error reading start_src_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_src_port: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSdwanServiceStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSdwanService-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tie_break", flattenSystemSdwanServiceTieBreak2edl(o["tie-break"], d, "tie_break")); err != nil {
		if vv, ok := fortiAPIPatch(o["tie-break"], "SystemSdwanService-TieBreak"); ok {
			if err = d.Set("tie_break", vv); err != nil {
				return fmt.Errorf("Error reading tie_break: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tie_break: %v", err)
		}
	}

	if err = d.Set("tos", flattenSystemSdwanServiceTos2edl(o["tos"], d, "tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos"], "SystemSdwanService-Tos"); ok {
			if err = d.Set("tos", vv); err != nil {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenSystemSdwanServiceTosMask2edl(o["tos-mask"], d, "tos_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-mask"], "SystemSdwanService-TosMask"); ok {
			if err = d.Set("tos_mask", vv); err != nil {
				return fmt.Errorf("Error reading tos_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("use_shortcut_sla", flattenSystemSdwanServiceUseShortcutSla2edl(o["use-shortcut-sla"], d, "use_shortcut_sla")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-shortcut-sla"], "SystemSdwanService-UseShortcutSla"); ok {
			if err = d.Set("use_shortcut_sla", vv); err != nil {
				return fmt.Errorf("Error reading use_shortcut_sla: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_shortcut_sla: %v", err)
		}
	}

	if err = d.Set("users", flattenSystemSdwanServiceUsers2edl(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "SystemSdwanService-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	if err = d.Set("zone_mode", flattenSystemSdwanServiceZoneMode2edl(o["zone-mode"], d, "zone_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["zone-mode"], "SystemSdwanService-ZoneMode"); ok {
			if err = d.Set("zone_mode", vv); err != nil {
				return fmt.Errorf("Error reading zone_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading zone_mode: %v", err)
		}
	}

	return nil
}

func flattenSystemSdwanServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdwanServiceAddrMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceAgentExclusive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceBandwidthWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDefault2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDscpForward2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDscpForwardTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDscpReverse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDscpReverseTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceDstNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDst62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceEndPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceEndSrcPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceGroups2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceHashMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceHealthCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceHoldDownTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInputDevice2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInputDeviceNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInputZone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInternetServiceAppCtrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetServiceAppCtrlCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetServiceAppCtrlGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetServiceCustom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetServiceCustomGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetServiceGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetServiceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceJitterWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceLatencyWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceLinkCostFactor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceLinkCostThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceLoadBalance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceMinimumSlaMeetMembers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServicePacketLossWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServicePassiveMeasurement2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServicePriorityMembers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServicePriorityZone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceQualityLink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceRole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceRouteTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceShortcut2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceShortcutPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandSystemSdwanServiceSlaHealthCheck2edl(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSdwanServiceSlaId2edl(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceSlaHealthCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceSlaId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSlaCompareMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSlaStickiness2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSrc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceSrcNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSrc62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceStandaloneAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceStartPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceStartSrcPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceTieBreak2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceTos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceTosMask2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceUseShortcutSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceUsers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceZoneMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdwanService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_mode"); ok || d.HasChange("addr_mode") {
		t, err := expandSystemSdwanServiceAddrMode2edl(d, v, "addr_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("agent_exclusive"); ok || d.HasChange("agent_exclusive") {
		t, err := expandSystemSdwanServiceAgentExclusive2edl(d, v, "agent_exclusive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agent-exclusive"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_weight"); ok || d.HasChange("bandwidth_weight") {
		t, err := expandSystemSdwanServiceBandwidthWeight2edl(d, v, "bandwidth_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-weight"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandSystemSdwanServiceComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("default"); ok || d.HasChange("default") {
		t, err := expandSystemSdwanServiceDefault2edl(d, v, "default")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default"] = t
		}
	}

	if v, ok := d.GetOk("dscp_forward"); ok || d.HasChange("dscp_forward") {
		t, err := expandSystemSdwanServiceDscpForward2edl(d, v, "dscp_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-forward"] = t
		}
	}

	if v, ok := d.GetOk("dscp_forward_tag"); ok || d.HasChange("dscp_forward_tag") {
		t, err := expandSystemSdwanServiceDscpForwardTag2edl(d, v, "dscp_forward_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-forward-tag"] = t
		}
	}

	if v, ok := d.GetOk("dscp_reverse"); ok || d.HasChange("dscp_reverse") {
		t, err := expandSystemSdwanServiceDscpReverse2edl(d, v, "dscp_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-reverse"] = t
		}
	}

	if v, ok := d.GetOk("dscp_reverse_tag"); ok || d.HasChange("dscp_reverse_tag") {
		t, err := expandSystemSdwanServiceDscpReverseTag2edl(d, v, "dscp_reverse_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-reverse-tag"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandSystemSdwanServiceDst2edl(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("dst_negate"); ok || d.HasChange("dst_negate") {
		t, err := expandSystemSdwanServiceDstNegate2edl(d, v, "dst_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-negate"] = t
		}
	}

	if v, ok := d.GetOk("dst6"); ok || d.HasChange("dst6") {
		t, err := expandSystemSdwanServiceDst62edl(d, v, "dst6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst6"] = t
		}
	}

	if v, ok := d.GetOk("end_port"); ok || d.HasChange("end_port") {
		t, err := expandSystemSdwanServiceEndPort2edl(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("end_src_port"); ok || d.HasChange("end_src_port") {
		t, err := expandSystemSdwanServiceEndSrcPort2edl(d, v, "end_src_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-src-port"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandSystemSdwanServiceGateway2edl(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandSystemSdwanServiceGroups2edl(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("hash_mode"); ok || d.HasChange("hash_mode") {
		t, err := expandSystemSdwanServiceHashMode2edl(d, v, "hash_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash-mode"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandSystemSdwanServiceHealthCheck2edl(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("hold_down_time"); ok || d.HasChange("hold_down_time") {
		t, err := expandSystemSdwanServiceHoldDownTime2edl(d, v, "hold_down_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hold-down-time"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSdwanServiceId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("input_device"); ok || d.HasChange("input_device") {
		t, err := expandSystemSdwanServiceInputDevice2edl(d, v, "input_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device"] = t
		}
	}

	if v, ok := d.GetOk("input_device_negate"); ok || d.HasChange("input_device_negate") {
		t, err := expandSystemSdwanServiceInputDeviceNegate2edl(d, v, "input_device_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device-negate"] = t
		}
	}

	if v, ok := d.GetOk("input_zone"); ok || d.HasChange("input_zone") {
		t, err := expandSystemSdwanServiceInputZone2edl(d, v, "input_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-zone"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok || d.HasChange("internet_service") {
		t, err := expandSystemSdwanServiceInternetService2edl(d, v, "internet_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_app_ctrl"); ok || d.HasChange("internet_service_app_ctrl") {
		t, err := expandSystemSdwanServiceInternetServiceAppCtrl2edl(d, v, "internet_service_app_ctrl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-app-ctrl"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_app_ctrl_category"); ok || d.HasChange("internet_service_app_ctrl_category") {
		t, err := expandSystemSdwanServiceInternetServiceAppCtrlCategory2edl(d, v, "internet_service_app_ctrl_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-app-ctrl-category"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_app_ctrl_group"); ok || d.HasChange("internet_service_app_ctrl_group") {
		t, err := expandSystemSdwanServiceInternetServiceAppCtrlGroup2edl(d, v, "internet_service_app_ctrl_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-app-ctrl-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok || d.HasChange("internet_service_custom") {
		t, err := expandSystemSdwanServiceInternetServiceCustom2edl(d, v, "internet_service_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok || d.HasChange("internet_service_custom_group") {
		t, err := expandSystemSdwanServiceInternetServiceCustomGroup2edl(d, v, "internet_service_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok || d.HasChange("internet_service_group") {
		t, err := expandSystemSdwanServiceInternetServiceGroup2edl(d, v, "internet_service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok || d.HasChange("internet_service_name") {
		t, err := expandSystemSdwanServiceInternetServiceName2edl(d, v, "internet_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	}

	if v, ok := d.GetOk("jitter_weight"); ok || d.HasChange("jitter_weight") {
		t, err := expandSystemSdwanServiceJitterWeight2edl(d, v, "jitter_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jitter-weight"] = t
		}
	}

	if v, ok := d.GetOk("latency_weight"); ok || d.HasChange("latency_weight") {
		t, err := expandSystemSdwanServiceLatencyWeight2edl(d, v, "latency_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-weight"] = t
		}
	}

	if v, ok := d.GetOk("link_cost_factor"); ok || d.HasChange("link_cost_factor") {
		t, err := expandSystemSdwanServiceLinkCostFactor2edl(d, v, "link_cost_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost-factor"] = t
		}
	}

	if v, ok := d.GetOk("link_cost_threshold"); ok || d.HasChange("link_cost_threshold") {
		t, err := expandSystemSdwanServiceLinkCostThreshold2edl(d, v, "link_cost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("load_balance"); ok || d.HasChange("load_balance") {
		t, err := expandSystemSdwanServiceLoadBalance2edl(d, v, "load_balance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance"] = t
		}
	}

	if v, ok := d.GetOk("minimum_sla_meet_members"); ok || d.HasChange("minimum_sla_meet_members") {
		t, err := expandSystemSdwanServiceMinimumSlaMeetMembers2edl(d, v, "minimum_sla_meet_members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-sla-meet-members"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemSdwanServiceMode2edl(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSdwanServiceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("packet_loss_weight"); ok || d.HasChange("packet_loss_weight") {
		t, err := expandSystemSdwanServicePacketLossWeight2edl(d, v, "packet_loss_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-loss-weight"] = t
		}
	}

	if v, ok := d.GetOk("passive_measurement"); ok || d.HasChange("passive_measurement") {
		t, err := expandSystemSdwanServicePassiveMeasurement2edl(d, v, "passive_measurement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-measurement"] = t
		}
	}

	if v, ok := d.GetOk("priority_members"); ok || d.HasChange("priority_members") {
		t, err := expandSystemSdwanServicePriorityMembers2edl(d, v, "priority_members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-members"] = t
		}
	}

	if v, ok := d.GetOk("priority_zone"); ok || d.HasChange("priority_zone") {
		t, err := expandSystemSdwanServicePriorityZone2edl(d, v, "priority_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-zone"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemSdwanServiceProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quality_link"); ok || d.HasChange("quality_link") {
		t, err := expandSystemSdwanServiceQualityLink2edl(d, v, "quality_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quality-link"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandSystemSdwanServiceRole2edl(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("route_tag"); ok || d.HasChange("route_tag") {
		t, err := expandSystemSdwanServiceRouteTag2edl(d, v, "route_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-tag"] = t
		}
	}

	if v, ok := d.GetOk("shortcut"); ok || d.HasChange("shortcut") {
		t, err := expandSystemSdwanServiceShortcut2edl(d, v, "shortcut")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shortcut"] = t
		}
	}

	if v, ok := d.GetOk("shortcut_priority"); ok || d.HasChange("shortcut_priority") {
		t, err := expandSystemSdwanServiceShortcutPriority2edl(d, v, "shortcut_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shortcut-priority"] = t
		}
	}

	if v, ok := d.GetOk("sla"); ok || d.HasChange("sla") {
		t, err := expandSystemSdwanServiceSla2edl(d, v, "sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla"] = t
		}
	}

	if v, ok := d.GetOk("sla_compare_method"); ok || d.HasChange("sla_compare_method") {
		t, err := expandSystemSdwanServiceSlaCompareMethod2edl(d, v, "sla_compare_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-compare-method"] = t
		}
	}

	if v, ok := d.GetOk("sla_stickiness"); ok || d.HasChange("sla_stickiness") {
		t, err := expandSystemSdwanServiceSlaStickiness2edl(d, v, "sla_stickiness")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-stickiness"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandSystemSdwanServiceSrc2edl(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("src_negate"); ok || d.HasChange("src_negate") {
		t, err := expandSystemSdwanServiceSrcNegate2edl(d, v, "src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-negate"] = t
		}
	}

	if v, ok := d.GetOk("src6"); ok || d.HasChange("src6") {
		t, err := expandSystemSdwanServiceSrc62edl(d, v, "src6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src6"] = t
		}
	}

	if v, ok := d.GetOk("standalone_action"); ok || d.HasChange("standalone_action") {
		t, err := expandSystemSdwanServiceStandaloneAction2edl(d, v, "standalone_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["standalone-action"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok || d.HasChange("start_port") {
		t, err := expandSystemSdwanServiceStartPort2edl(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	if v, ok := d.GetOk("start_src_port"); ok || d.HasChange("start_src_port") {
		t, err := expandSystemSdwanServiceStartSrcPort2edl(d, v, "start_src_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-src-port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSdwanServiceStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tie_break"); ok || d.HasChange("tie_break") {
		t, err := expandSystemSdwanServiceTieBreak2edl(d, v, "tie_break")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tie-break"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandSystemSdwanServiceTos2edl(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok || d.HasChange("tos_mask") {
		t, err := expandSystemSdwanServiceTosMask2edl(d, v, "tos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("use_shortcut_sla"); ok || d.HasChange("use_shortcut_sla") {
		t, err := expandSystemSdwanServiceUseShortcutSla2edl(d, v, "use_shortcut_sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-shortcut-sla"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandSystemSdwanServiceUsers2edl(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("zone_mode"); ok || d.HasChange("zone_mode") {
		t, err := expandSystemSdwanServiceZoneMode2edl(d, v, "zone_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zone-mode"] = t
		}
	}

	return &obj, nil
}
