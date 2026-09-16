// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiGate interfaces added to the SD-WAN.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdwanMembers() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdwanMembersCreate,
		Read:   resourceSystemSdwanMembersRead,
		Update: resourceSystemSdwanMembersUpdate,
		Delete: resourceSystemSdwanMembersDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"billing_start_day": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"duplication_threshold_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"duplication_threshold_bibandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"duplication_threshold_dwbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"duplication_threshold_upbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ingress_spillover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"overage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"overage_cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"overage_volume_ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"overage_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"preferred_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"priority_in_sla": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority_out_sla": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority6": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"quota_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spillover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transport_group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"volume_ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSdwanMembersCreate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemSdwanMembers(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanMembers resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("seq_num")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemSdwanMembers(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemSdwanMembers(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemSdwanMembers resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemSdwanMembers(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemSdwanMembers resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "seq_num")))

	return resourceSystemSdwanMembersRead(d, m)
}

func resourceSystemSdwanMembersUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemSdwanMembers(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanMembers resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSdwanMembers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanMembers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "seq_num")))

	return resourceSystemSdwanMembersRead(d, m)
}

func resourceSystemSdwanMembersDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteSystemSdwanMembers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdwanMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanMembersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdwanMembers(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemSdwanMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdwanMembers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanMembers resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdwanMembersBillingStartDay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersCost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersDuplicationThresholdBandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersDuplicationThresholdBibandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersDuplicationThresholdDwbandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersDuplicationThresholdUpbandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersGateway62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersIngressSpilloverThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanMembersOverage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersOverageCost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersOverageVolumeRatio2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersOverageWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersPreferredSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersPriorityInSla2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersPriorityOutSla2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersPriority62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersQuotaLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersSeqNum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersSource62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersSpilloverThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersTransportGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersVolumeRatio2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersZone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemSdwanMembers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("billing_start_day", flattenSystemSdwanMembersBillingStartDay2edl(o["billing-start-day"], d, "billing_start_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["billing-start-day"], "SystemSdwanMembers-BillingStartDay"); ok {
			if err = d.Set("billing_start_day", vv); err != nil {
				return fmt.Errorf("Error reading billing_start_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading billing_start_day: %v", err)
		}
	}

	if err = d.Set("comment", flattenSystemSdwanMembersComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "SystemSdwanMembers-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("cost", flattenSystemSdwanMembersCost2edl(o["cost"], d, "cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["cost"], "SystemSdwanMembers-Cost"); ok {
			if err = d.Set("cost", vv); err != nil {
				return fmt.Errorf("Error reading cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("duplication_threshold_bandwidth", flattenSystemSdwanMembersDuplicationThresholdBandwidth2edl(o["duplication-threshold-bandwidth"], d, "duplication_threshold_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplication-threshold-bandwidth"], "SystemSdwanMembers-DuplicationThresholdBandwidth"); ok {
			if err = d.Set("duplication_threshold_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading duplication_threshold_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplication_threshold_bandwidth: %v", err)
		}
	}

	if err = d.Set("duplication_threshold_bibandwidth", flattenSystemSdwanMembersDuplicationThresholdBibandwidth2edl(o["duplication-threshold-bibandwidth"], d, "duplication_threshold_bibandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplication-threshold-bibandwidth"], "SystemSdwanMembers-DuplicationThresholdBibandwidth"); ok {
			if err = d.Set("duplication_threshold_bibandwidth", vv); err != nil {
				return fmt.Errorf("Error reading duplication_threshold_bibandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplication_threshold_bibandwidth: %v", err)
		}
	}

	if err = d.Set("duplication_threshold_dwbandwidth", flattenSystemSdwanMembersDuplicationThresholdDwbandwidth2edl(o["duplication-threshold-dwbandwidth"], d, "duplication_threshold_dwbandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplication-threshold-dwbandwidth"], "SystemSdwanMembers-DuplicationThresholdDwbandwidth"); ok {
			if err = d.Set("duplication_threshold_dwbandwidth", vv); err != nil {
				return fmt.Errorf("Error reading duplication_threshold_dwbandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplication_threshold_dwbandwidth: %v", err)
		}
	}

	if err = d.Set("duplication_threshold_upbandwidth", flattenSystemSdwanMembersDuplicationThresholdUpbandwidth2edl(o["duplication-threshold-upbandwidth"], d, "duplication_threshold_upbandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplication-threshold-upbandwidth"], "SystemSdwanMembers-DuplicationThresholdUpbandwidth"); ok {
			if err = d.Set("duplication_threshold_upbandwidth", vv); err != nil {
				return fmt.Errorf("Error reading duplication_threshold_upbandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplication_threshold_upbandwidth: %v", err)
		}
	}

	if err = d.Set("gateway", flattenSystemSdwanMembersGateway2edl(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "SystemSdwanMembers-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("gateway6", flattenSystemSdwanMembersGateway62edl(o["gateway6"], d, "gateway6")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway6"], "SystemSdwanMembers-Gateway6"); ok {
			if err = d.Set("gateway6", vv); err != nil {
				return fmt.Errorf("Error reading gateway6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway6: %v", err)
		}
	}

	if err = d.Set("ingress_spillover_threshold", flattenSystemSdwanMembersIngressSpilloverThreshold2edl(o["ingress-spillover-threshold"], d, "ingress_spillover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ingress-spillover-threshold"], "SystemSdwanMembers-IngressSpilloverThreshold"); ok {
			if err = d.Set("ingress_spillover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemSdwanMembersInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemSdwanMembers-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("overage", flattenSystemSdwanMembersOverage2edl(o["overage"], d, "overage")); err != nil {
		if vv, ok := fortiAPIPatch(o["overage"], "SystemSdwanMembers-Overage"); ok {
			if err = d.Set("overage", vv); err != nil {
				return fmt.Errorf("Error reading overage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overage: %v", err)
		}
	}

	if err = d.Set("overage_cost", flattenSystemSdwanMembersOverageCost2edl(o["overage-cost"], d, "overage_cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["overage-cost"], "SystemSdwanMembers-OverageCost"); ok {
			if err = d.Set("overage_cost", vv); err != nil {
				return fmt.Errorf("Error reading overage_cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overage_cost: %v", err)
		}
	}

	if err = d.Set("overage_volume_ratio", flattenSystemSdwanMembersOverageVolumeRatio2edl(o["overage-volume-ratio"], d, "overage_volume_ratio")); err != nil {
		if vv, ok := fortiAPIPatch(o["overage-volume-ratio"], "SystemSdwanMembers-OverageVolumeRatio"); ok {
			if err = d.Set("overage_volume_ratio", vv); err != nil {
				return fmt.Errorf("Error reading overage_volume_ratio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overage_volume_ratio: %v", err)
		}
	}

	if err = d.Set("overage_weight", flattenSystemSdwanMembersOverageWeight2edl(o["overage-weight"], d, "overage_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["overage-weight"], "SystemSdwanMembers-OverageWeight"); ok {
			if err = d.Set("overage_weight", vv); err != nil {
				return fmt.Errorf("Error reading overage_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overage_weight: %v", err)
		}
	}

	if err = d.Set("preferred_source", flattenSystemSdwanMembersPreferredSource2edl(o["preferred-source"], d, "preferred_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-source"], "SystemSdwanMembers-PreferredSource"); ok {
			if err = d.Set("preferred_source", vv); err != nil {
				return fmt.Errorf("Error reading preferred_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_source: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemSdwanMembersPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemSdwanMembers-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("priority_in_sla", flattenSystemSdwanMembersPriorityInSla2edl(o["priority-in-sla"], d, "priority_in_sla")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-in-sla"], "SystemSdwanMembers-PriorityInSla"); ok {
			if err = d.Set("priority_in_sla", vv); err != nil {
				return fmt.Errorf("Error reading priority_in_sla: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_in_sla: %v", err)
		}
	}

	if err = d.Set("priority_out_sla", flattenSystemSdwanMembersPriorityOutSla2edl(o["priority-out-sla"], d, "priority_out_sla")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-out-sla"], "SystemSdwanMembers-PriorityOutSla"); ok {
			if err = d.Set("priority_out_sla", vv); err != nil {
				return fmt.Errorf("Error reading priority_out_sla: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_out_sla: %v", err)
		}
	}

	if err = d.Set("priority6", flattenSystemSdwanMembersPriority62edl(o["priority6"], d, "priority6")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority6"], "SystemSdwanMembers-Priority6"); ok {
			if err = d.Set("priority6", vv); err != nil {
				return fmt.Errorf("Error reading priority6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority6: %v", err)
		}
	}

	if err = d.Set("quota_limit", flattenSystemSdwanMembersQuotaLimit2edl(o["quota-limit"], d, "quota_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["quota-limit"], "SystemSdwanMembers-QuotaLimit"); ok {
			if err = d.Set("quota_limit", vv); err != nil {
				return fmt.Errorf("Error reading quota_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quota_limit: %v", err)
		}
	}

	if err = d.Set("seq_num", flattenSystemSdwanMembersSeqNum2edl(o["seq-num"], d, "seq_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["seq-num"], "SystemSdwanMembers-SeqNum"); ok {
			if err = d.Set("seq_num", vv); err != nil {
				return fmt.Errorf("Error reading seq_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if err = d.Set("source", flattenSystemSdwanMembersSource2edl(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "SystemSdwanMembers-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("source6", flattenSystemSdwanMembersSource62edl(o["source6"], d, "source6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source6"], "SystemSdwanMembers-Source6"); ok {
			if err = d.Set("source6", vv); err != nil {
				return fmt.Errorf("Error reading source6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source6: %v", err)
		}
	}

	if err = d.Set("spillover_threshold", flattenSystemSdwanMembersSpilloverThreshold2edl(o["spillover-threshold"], d, "spillover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["spillover-threshold"], "SystemSdwanMembers-SpilloverThreshold"); ok {
			if err = d.Set("spillover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading spillover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spillover_threshold: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSdwanMembersStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSdwanMembers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("transport_group", flattenSystemSdwanMembersTransportGroup2edl(o["transport-group"], d, "transport_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["transport-group"], "SystemSdwanMembers-TransportGroup"); ok {
			if err = d.Set("transport_group", vv); err != nil {
				return fmt.Errorf("Error reading transport_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transport_group: %v", err)
		}
	}

	if err = d.Set("volume_ratio", flattenSystemSdwanMembersVolumeRatio2edl(o["volume-ratio"], d, "volume_ratio")); err != nil {
		if vv, ok := fortiAPIPatch(o["volume-ratio"], "SystemSdwanMembers-VolumeRatio"); ok {
			if err = d.Set("volume_ratio", vv); err != nil {
				return fmt.Errorf("Error reading volume_ratio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading volume_ratio: %v", err)
		}
	}

	if err = d.Set("weight", flattenSystemSdwanMembersWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "SystemSdwanMembers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("zone", flattenSystemSdwanMembersZone2edl(o["zone"], d, "zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["zone"], "SystemSdwanMembers-Zone"); ok {
			if err = d.Set("zone", vv); err != nil {
				return fmt.Errorf("Error reading zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading zone: %v", err)
		}
	}

	return nil
}

func flattenSystemSdwanMembersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdwanMembersBillingStartDay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersCost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersDuplicationThresholdBandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersDuplicationThresholdBibandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersDuplicationThresholdDwbandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersDuplicationThresholdUpbandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersGateway62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersIngressSpilloverThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanMembersOverage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersOverageCost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersOverageVolumeRatio2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersOverageWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersPreferredSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersPriorityInSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersPriorityOutSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersPriority62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersQuotaLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersSeqNum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersSource62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersSpilloverThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersTransportGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersVolumeRatio2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersZone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemSdwanMembers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("billing_start_day"); ok || d.HasChange("billing_start_day") {
		t, err := expandSystemSdwanMembersBillingStartDay2edl(d, v, "billing_start_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["billing-start-day"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandSystemSdwanMembersComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("cost"); ok || d.HasChange("cost") {
		t, err := expandSystemSdwanMembersCost2edl(d, v, "cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("duplication_threshold_bandwidth"); ok || d.HasChange("duplication_threshold_bandwidth") {
		t, err := expandSystemSdwanMembersDuplicationThresholdBandwidth2edl(d, v, "duplication_threshold_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-threshold-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("duplication_threshold_bibandwidth"); ok || d.HasChange("duplication_threshold_bibandwidth") {
		t, err := expandSystemSdwanMembersDuplicationThresholdBibandwidth2edl(d, v, "duplication_threshold_bibandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-threshold-bibandwidth"] = t
		}
	}

	if v, ok := d.GetOk("duplication_threshold_dwbandwidth"); ok || d.HasChange("duplication_threshold_dwbandwidth") {
		t, err := expandSystemSdwanMembersDuplicationThresholdDwbandwidth2edl(d, v, "duplication_threshold_dwbandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-threshold-dwbandwidth"] = t
		}
	}

	if v, ok := d.GetOk("duplication_threshold_upbandwidth"); ok || d.HasChange("duplication_threshold_upbandwidth") {
		t, err := expandSystemSdwanMembersDuplicationThresholdUpbandwidth2edl(d, v, "duplication_threshold_upbandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-threshold-upbandwidth"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandSystemSdwanMembersGateway2edl(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("gateway6"); ok || d.HasChange("gateway6") {
		t, err := expandSystemSdwanMembersGateway62edl(d, v, "gateway6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway6"] = t
		}
	}

	if v, ok := d.GetOk("ingress_spillover_threshold"); ok || d.HasChange("ingress_spillover_threshold") {
		t, err := expandSystemSdwanMembersIngressSpilloverThreshold2edl(d, v, "ingress_spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemSdwanMembersInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("overage"); ok || d.HasChange("overage") {
		t, err := expandSystemSdwanMembersOverage2edl(d, v, "overage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage"] = t
		}
	}

	if v, ok := d.GetOk("overage_cost"); ok || d.HasChange("overage_cost") {
		t, err := expandSystemSdwanMembersOverageCost2edl(d, v, "overage_cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage-cost"] = t
		}
	}

	if v, ok := d.GetOk("overage_volume_ratio"); ok || d.HasChange("overage_volume_ratio") {
		t, err := expandSystemSdwanMembersOverageVolumeRatio2edl(d, v, "overage_volume_ratio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage-volume-ratio"] = t
		}
	}

	if v, ok := d.GetOk("overage_weight"); ok || d.HasChange("overage_weight") {
		t, err := expandSystemSdwanMembersOverageWeight2edl(d, v, "overage_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage-weight"] = t
		}
	}

	if v, ok := d.GetOk("preferred_source"); ok || d.HasChange("preferred_source") {
		t, err := expandSystemSdwanMembersPreferredSource2edl(d, v, "preferred_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-source"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemSdwanMembersPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("priority_in_sla"); ok || d.HasChange("priority_in_sla") {
		t, err := expandSystemSdwanMembersPriorityInSla2edl(d, v, "priority_in_sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-in-sla"] = t
		}
	}

	if v, ok := d.GetOk("priority_out_sla"); ok || d.HasChange("priority_out_sla") {
		t, err := expandSystemSdwanMembersPriorityOutSla2edl(d, v, "priority_out_sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-out-sla"] = t
		}
	}

	if v, ok := d.GetOk("priority6"); ok || d.HasChange("priority6") {
		t, err := expandSystemSdwanMembersPriority62edl(d, v, "priority6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority6"] = t
		}
	}

	if v, ok := d.GetOk("quota_limit"); ok || d.HasChange("quota_limit") {
		t, err := expandSystemSdwanMembersQuotaLimit2edl(d, v, "quota_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quota-limit"] = t
		}
	}

	if v, ok := d.GetOk("seq_num"); ok || d.HasChange("seq_num") {
		t, err := expandSystemSdwanMembersSeqNum2edl(d, v, "seq_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandSystemSdwanMembersSource2edl(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source6"); ok || d.HasChange("source6") {
		t, err := expandSystemSdwanMembersSource62edl(d, v, "source6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source6"] = t
		}
	}

	if v, ok := d.GetOk("spillover_threshold"); ok || d.HasChange("spillover_threshold") {
		t, err := expandSystemSdwanMembersSpilloverThreshold2edl(d, v, "spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSdwanMembersStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("transport_group"); ok || d.HasChange("transport_group") {
		t, err := expandSystemSdwanMembersTransportGroup2edl(d, v, "transport_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transport-group"] = t
		}
	}

	if v, ok := d.GetOk("volume_ratio"); ok || d.HasChange("volume_ratio") {
		t, err := expandSystemSdwanMembersVolumeRatio2edl(d, v, "volume_ratio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["volume-ratio"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandSystemSdwanMembersWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("zone"); ok || d.HasChange("zone") {
		t, err := expandSystemSdwanMembersZone2edl(d, v, "zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zone"] = t
		}
	}

	return &obj, nil
}
