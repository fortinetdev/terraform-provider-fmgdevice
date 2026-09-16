// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FEC redundancy mapping table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnIpsecFecMappings() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecFecMappingsCreate,
		Read:   resourceVpnIpsecFecMappingsRead,
		Update: resourceVpnIpsecFecMappingsUpdate,
		Delete: resourceVpnIpsecFecMappingsDelete,

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
			"fec": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bandwidth_bi_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bandwidth_bi_threshold_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_down_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bandwidth_down_threshold_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_up_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bandwidth_up_threshold_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"base": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_threshold_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"packet_loss_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"packet_loss_threshold_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redundant": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"seqno": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"redundant": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"seqno": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tos": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tos_mask": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceVpnIpsecFecMappingsCreate(d *schema.ResourceData, m interface{}) error {
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
	fec := d.Get("fec").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fec"] = fec

	obj, err := getObjectVpnIpsecFecMappings(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecFecMappings resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("seqno")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVpnIpsecFecMappings(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVpnIpsecFecMappings(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VpnIpsecFecMappings resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateVpnIpsecFecMappings(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VpnIpsecFecMappings resource: %v", err)
		}

		if v != nil && v["seqno"] != nil {
			if vidn, ok := v["seqno"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceVpnIpsecFecMappingsRead(d, m)
			} else {
				return fmt.Errorf("Error creating VpnIpsecFecMappings resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "seqno")))

	return resourceVpnIpsecFecMappingsRead(d, m)
}

func resourceVpnIpsecFecMappingsUpdate(d *schema.ResourceData, m interface{}) error {
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
	fec := d.Get("fec").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fec"] = fec

	obj, err := getObjectVpnIpsecFecMappings(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecFecMappings resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	v, err := c.UpdateVpnIpsecFecMappings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecFecMappings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	if v != nil && v["seqno"] != nil {
		if vidn, ok := v["seqno"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceVpnIpsecFecMappingsRead(d, m)
		} else {
			return fmt.Errorf("Error updating VpnIpsecFecMappings resource: %v", err)
		}
	}

	return resourceVpnIpsecFecMappingsRead(d, m)
}

func resourceVpnIpsecFecMappingsDelete(d *schema.ResourceData, m interface{}) error {
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
	fec := d.Get("fec").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fec"] = fec

	wsParams["adom"] = adomv

	err = c.DeleteVpnIpsecFecMappings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecFecMappings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecFecMappingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	fec := d.Get("fec").(string)
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
	if fec == "" {
		fec = importOptionChecking(m.(*FortiClient).Cfg, "fec")
		if fec == "" {
			return fmt.Errorf("Parameter fec is missing")
		}
		if err = d.Set("fec", fec); err != nil {
			return fmt.Errorf("Error set params fec: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fec"] = fec

	o, err := c.ReadVpnIpsecFecMappings(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnIpsecFecMappings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecFecMappings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecFecMappings resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecFecMappingsBandwidthBiThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBandwidthBiThresholdNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBandwidthDownThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBandwidthDownThresholdNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBandwidthUpThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBandwidthUpThresholdNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsLatencyThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsLatencyThresholdNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsPacketLossThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsPacketLossThresholdNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsRedundant2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsSeqno2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsTos2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base"
		if _, ok := i["base"]; ok {
			v := flattenVpnIpsecFecMappingsTosBase2edl(i["base"], d, pre_append)
			tmp["base"] = fortiAPISubPartPatch(v, "VpnIpsecFecMappings-Tos-Base")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redundant"
		if _, ok := i["redundant"]; ok {
			v := flattenVpnIpsecFecMappingsTosRedundant2edl(i["redundant"], d, pre_append)
			tmp["redundant"] = fortiAPISubPartPatch(v, "VpnIpsecFecMappings-Tos-Redundant")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seqno"
		if _, ok := i["seqno"]; ok {
			v := flattenVpnIpsecFecMappingsTosSeqno2edl(i["seqno"], d, pre_append)
			tmp["seqno"] = fortiAPISubPartPatch(v, "VpnIpsecFecMappings-Tos-Seqno")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := i["tos"]; ok {
			v := flattenVpnIpsecFecMappingsTosTos2edl(i["tos"], d, pre_append)
			tmp["tos"] = fortiAPISubPartPatch(v, "VpnIpsecFecMappings-Tos-Tos")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := i["tos-mask"]; ok {
			v := flattenVpnIpsecFecMappingsTosTosMask2edl(i["tos-mask"], d, pre_append)
			tmp["tos_mask"] = fortiAPISubPartPatch(v, "VpnIpsecFecMappings-Tos-TosMask")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnIpsecFecMappingsTosBase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsTosRedundant2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsTosSeqno2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsTosTos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsTosTosMask2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnIpsecFecMappings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("bandwidth_bi_threshold", flattenVpnIpsecFecMappingsBandwidthBiThreshold2edl(o["bandwidth-bi-threshold"], d, "bandwidth_bi_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-bi-threshold"], "VpnIpsecFecMappings-BandwidthBiThreshold"); ok {
			if err = d.Set("bandwidth_bi_threshold", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_bi_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_bi_threshold: %v", err)
		}
	}

	if err = d.Set("bandwidth_bi_threshold_negate", flattenVpnIpsecFecMappingsBandwidthBiThresholdNegate2edl(o["bandwidth-bi-threshold-negate"], d, "bandwidth_bi_threshold_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-bi-threshold-negate"], "VpnIpsecFecMappings-BandwidthBiThresholdNegate"); ok {
			if err = d.Set("bandwidth_bi_threshold_negate", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_bi_threshold_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_bi_threshold_negate: %v", err)
		}
	}

	if err = d.Set("bandwidth_down_threshold", flattenVpnIpsecFecMappingsBandwidthDownThreshold2edl(o["bandwidth-down-threshold"], d, "bandwidth_down_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-down-threshold"], "VpnIpsecFecMappings-BandwidthDownThreshold"); ok {
			if err = d.Set("bandwidth_down_threshold", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_down_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_down_threshold: %v", err)
		}
	}

	if err = d.Set("bandwidth_down_threshold_negate", flattenVpnIpsecFecMappingsBandwidthDownThresholdNegate2edl(o["bandwidth-down-threshold-negate"], d, "bandwidth_down_threshold_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-down-threshold-negate"], "VpnIpsecFecMappings-BandwidthDownThresholdNegate"); ok {
			if err = d.Set("bandwidth_down_threshold_negate", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_down_threshold_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_down_threshold_negate: %v", err)
		}
	}

	if err = d.Set("bandwidth_up_threshold", flattenVpnIpsecFecMappingsBandwidthUpThreshold2edl(o["bandwidth-up-threshold"], d, "bandwidth_up_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-up-threshold"], "VpnIpsecFecMappings-BandwidthUpThreshold"); ok {
			if err = d.Set("bandwidth_up_threshold", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_up_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_up_threshold: %v", err)
		}
	}

	if err = d.Set("bandwidth_up_threshold_negate", flattenVpnIpsecFecMappingsBandwidthUpThresholdNegate2edl(o["bandwidth-up-threshold-negate"], d, "bandwidth_up_threshold_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-up-threshold-negate"], "VpnIpsecFecMappings-BandwidthUpThresholdNegate"); ok {
			if err = d.Set("bandwidth_up_threshold_negate", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_up_threshold_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_up_threshold_negate: %v", err)
		}
	}

	if err = d.Set("base", flattenVpnIpsecFecMappingsBase2edl(o["base"], d, "base")); err != nil {
		if vv, ok := fortiAPIPatch(o["base"], "VpnIpsecFecMappings-Base"); ok {
			if err = d.Set("base", vv); err != nil {
				return fmt.Errorf("Error reading base: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading base: %v", err)
		}
	}

	if err = d.Set("latency_threshold", flattenVpnIpsecFecMappingsLatencyThreshold2edl(o["latency-threshold"], d, "latency_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold"], "VpnIpsecFecMappings-LatencyThreshold"); ok {
			if err = d.Set("latency_threshold", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold: %v", err)
		}
	}

	if err = d.Set("latency_threshold_negate", flattenVpnIpsecFecMappingsLatencyThresholdNegate2edl(o["latency-threshold-negate"], d, "latency_threshold_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold-negate"], "VpnIpsecFecMappings-LatencyThresholdNegate"); ok {
			if err = d.Set("latency_threshold_negate", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold_negate: %v", err)
		}
	}

	if err = d.Set("packet_loss_threshold", flattenVpnIpsecFecMappingsPacketLossThreshold2edl(o["packet-loss-threshold"], d, "packet_loss_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-loss-threshold"], "VpnIpsecFecMappings-PacketLossThreshold"); ok {
			if err = d.Set("packet_loss_threshold", vv); err != nil {
				return fmt.Errorf("Error reading packet_loss_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_loss_threshold: %v", err)
		}
	}

	if err = d.Set("packet_loss_threshold_negate", flattenVpnIpsecFecMappingsPacketLossThresholdNegate2edl(o["packet-loss-threshold-negate"], d, "packet_loss_threshold_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-loss-threshold-negate"], "VpnIpsecFecMappings-PacketLossThresholdNegate"); ok {
			if err = d.Set("packet_loss_threshold_negate", vv); err != nil {
				return fmt.Errorf("Error reading packet_loss_threshold_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_loss_threshold_negate: %v", err)
		}
	}

	if err = d.Set("redundant", flattenVpnIpsecFecMappingsRedundant2edl(o["redundant"], d, "redundant")); err != nil {
		if vv, ok := fortiAPIPatch(o["redundant"], "VpnIpsecFecMappings-Redundant"); ok {
			if err = d.Set("redundant", vv); err != nil {
				return fmt.Errorf("Error reading redundant: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redundant: %v", err)
		}
	}

	if err = d.Set("seqno", flattenVpnIpsecFecMappingsSeqno2edl(o["seqno"], d, "seqno")); err != nil {
		if vv, ok := fortiAPIPatch(o["seqno"], "VpnIpsecFecMappings-Seqno"); ok {
			if err = d.Set("seqno", vv); err != nil {
				return fmt.Errorf("Error reading seqno: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seqno: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tos", flattenVpnIpsecFecMappingsTos2edl(o["tos"], d, "tos")); err != nil {
			if vv, ok := fortiAPIPatch(o["tos"], "VpnIpsecFecMappings-Tos"); ok {
				if err = d.Set("tos", vv); err != nil {
					return fmt.Errorf("Error reading tos: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tos"); ok {
			if err = d.Set("tos", flattenVpnIpsecFecMappingsTos2edl(o["tos"], d, "tos")); err != nil {
				if vv, ok := fortiAPIPatch(o["tos"], "VpnIpsecFecMappings-Tos"); ok {
					if err = d.Set("tos", vv); err != nil {
						return fmt.Errorf("Error reading tos: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tos: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenVpnIpsecFecMappingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnIpsecFecMappingsBandwidthBiThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBandwidthBiThresholdNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBandwidthDownThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBandwidthDownThresholdNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBandwidthUpThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBandwidthUpThresholdNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsLatencyThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsLatencyThresholdNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsPacketLossThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsPacketLossThresholdNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsRedundant2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsSeqno2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsTos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["base"], _ = expandVpnIpsecFecMappingsTosBase2edl(d, i["base"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redundant"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["redundant"], _ = expandVpnIpsecFecMappingsTosRedundant2edl(d, i["redundant"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seqno"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seqno"], _ = expandVpnIpsecFecMappingsTosSeqno2edl(d, i["seqno"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos"], _ = expandVpnIpsecFecMappingsTosTos2edl(d, i["tos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos-mask"], _ = expandVpnIpsecFecMappingsTosTosMask2edl(d, i["tos_mask"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnIpsecFecMappingsTosBase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsTosRedundant2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsTosSeqno2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsTosTos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsTosTosMask2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecFecMappings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bandwidth_bi_threshold"); ok || d.HasChange("bandwidth_bi_threshold") {
		t, err := expandVpnIpsecFecMappingsBandwidthBiThreshold2edl(d, v, "bandwidth_bi_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-bi-threshold"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_bi_threshold_negate"); ok || d.HasChange("bandwidth_bi_threshold_negate") {
		t, err := expandVpnIpsecFecMappingsBandwidthBiThresholdNegate2edl(d, v, "bandwidth_bi_threshold_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-bi-threshold-negate"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_down_threshold"); ok || d.HasChange("bandwidth_down_threshold") {
		t, err := expandVpnIpsecFecMappingsBandwidthDownThreshold2edl(d, v, "bandwidth_down_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-down-threshold"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_down_threshold_negate"); ok || d.HasChange("bandwidth_down_threshold_negate") {
		t, err := expandVpnIpsecFecMappingsBandwidthDownThresholdNegate2edl(d, v, "bandwidth_down_threshold_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-down-threshold-negate"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_up_threshold"); ok || d.HasChange("bandwidth_up_threshold") {
		t, err := expandVpnIpsecFecMappingsBandwidthUpThreshold2edl(d, v, "bandwidth_up_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-up-threshold"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_up_threshold_negate"); ok || d.HasChange("bandwidth_up_threshold_negate") {
		t, err := expandVpnIpsecFecMappingsBandwidthUpThresholdNegate2edl(d, v, "bandwidth_up_threshold_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-up-threshold-negate"] = t
		}
	}

	if v, ok := d.GetOk("base"); ok || d.HasChange("base") {
		t, err := expandVpnIpsecFecMappingsBase2edl(d, v, "base")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base"] = t
		}
	}

	if v, ok := d.GetOk("latency_threshold"); ok || d.HasChange("latency_threshold") {
		t, err := expandVpnIpsecFecMappingsLatencyThreshold2edl(d, v, "latency_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold"] = t
		}
	}

	if v, ok := d.GetOk("latency_threshold_negate"); ok || d.HasChange("latency_threshold_negate") {
		t, err := expandVpnIpsecFecMappingsLatencyThresholdNegate2edl(d, v, "latency_threshold_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold-negate"] = t
		}
	}

	if v, ok := d.GetOk("packet_loss_threshold"); ok || d.HasChange("packet_loss_threshold") {
		t, err := expandVpnIpsecFecMappingsPacketLossThreshold2edl(d, v, "packet_loss_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-loss-threshold"] = t
		}
	}

	if v, ok := d.GetOk("packet_loss_threshold_negate"); ok || d.HasChange("packet_loss_threshold_negate") {
		t, err := expandVpnIpsecFecMappingsPacketLossThresholdNegate2edl(d, v, "packet_loss_threshold_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-loss-threshold-negate"] = t
		}
	}

	if v, ok := d.GetOk("redundant"); ok || d.HasChange("redundant") {
		t, err := expandVpnIpsecFecMappingsRedundant2edl(d, v, "redundant")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant"] = t
		}
	}

	if v, ok := d.GetOk("seqno"); ok || d.HasChange("seqno") {
		t, err := expandVpnIpsecFecMappingsSeqno2edl(d, v, "seqno")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seqno"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandVpnIpsecFecMappingsTos2edl(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	return &obj, nil
}
