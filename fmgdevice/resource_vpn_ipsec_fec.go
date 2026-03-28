// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Forward Error Correction (FEC) mapping profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnIpsecFec() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecFecCreate,
		Read:   resourceVpnIpsecFecRead,
		Update: resourceVpnIpsecFecUpdate,
		Delete: resourceVpnIpsecFecDelete,

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
			"mappings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bandwidth_bi_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bandwidth_down_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bandwidth_up_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"base": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"latency_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"packet_loss_threshold": &schema.Schema{
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
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceVpnIpsecFecCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnIpsecFec(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecFec resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVpnIpsecFec(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVpnIpsecFec(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VpnIpsecFec resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateVpnIpsecFec(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VpnIpsecFec resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnIpsecFecRead(d, m)
}

func resourceVpnIpsecFecUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnIpsecFec(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecFec resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVpnIpsecFec(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecFec resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnIpsecFecRead(d, m)
}

func resourceVpnIpsecFecDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnIpsecFec(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecFec resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecFecRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnIpsecFec(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnIpsecFec resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecFec(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecFec resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecFecMappings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_bi_threshold"
		if _, ok := i["bandwidth-bi-threshold"]; ok {
			v := flattenVpnIpsecFecMappingsBandwidthBiThreshold(i["bandwidth-bi-threshold"], d, pre_append)
			tmp["bandwidth_bi_threshold"] = fortiAPISubPartPatch(v, "VpnIpsecFec-Mappings-BandwidthBiThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_down_threshold"
		if _, ok := i["bandwidth-down-threshold"]; ok {
			v := flattenVpnIpsecFecMappingsBandwidthDownThreshold(i["bandwidth-down-threshold"], d, pre_append)
			tmp["bandwidth_down_threshold"] = fortiAPISubPartPatch(v, "VpnIpsecFec-Mappings-BandwidthDownThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_up_threshold"
		if _, ok := i["bandwidth-up-threshold"]; ok {
			v := flattenVpnIpsecFecMappingsBandwidthUpThreshold(i["bandwidth-up-threshold"], d, pre_append)
			tmp["bandwidth_up_threshold"] = fortiAPISubPartPatch(v, "VpnIpsecFec-Mappings-BandwidthUpThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base"
		if _, ok := i["base"]; ok {
			v := flattenVpnIpsecFecMappingsBase(i["base"], d, pre_append)
			tmp["base"] = fortiAPISubPartPatch(v, "VpnIpsecFec-Mappings-Base")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenVpnIpsecFecMappingsLatencyThreshold(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "VpnIpsecFec-Mappings-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_threshold"
		if _, ok := i["packet-loss-threshold"]; ok {
			v := flattenVpnIpsecFecMappingsPacketLossThreshold(i["packet-loss-threshold"], d, pre_append)
			tmp["packet_loss_threshold"] = fortiAPISubPartPatch(v, "VpnIpsecFec-Mappings-PacketLossThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redundant"
		if _, ok := i["redundant"]; ok {
			v := flattenVpnIpsecFecMappingsRedundant(i["redundant"], d, pre_append)
			tmp["redundant"] = fortiAPISubPartPatch(v, "VpnIpsecFec-Mappings-Redundant")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seqno"
		if _, ok := i["seqno"]; ok {
			v := flattenVpnIpsecFecMappingsSeqno(i["seqno"], d, pre_append)
			tmp["seqno"] = fortiAPISubPartPatch(v, "VpnIpsecFec-Mappings-Seqno")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnIpsecFecMappingsBandwidthBiThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBandwidthDownThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBandwidthUpThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsPacketLossThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsRedundant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsSeqno(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnIpsecFec(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("mappings", flattenVpnIpsecFecMappings(o["mappings"], d, "mappings")); err != nil {
			if vv, ok := fortiAPIPatch(o["mappings"], "VpnIpsecFec-Mappings"); ok {
				if err = d.Set("mappings", vv); err != nil {
					return fmt.Errorf("Error reading mappings: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mappings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mappings"); ok {
			if err = d.Set("mappings", flattenVpnIpsecFecMappings(o["mappings"], d, "mappings")); err != nil {
				if vv, ok := fortiAPIPatch(o["mappings"], "VpnIpsecFec-Mappings"); ok {
					if err = d.Set("mappings", vv); err != nil {
						return fmt.Errorf("Error reading mappings: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mappings: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenVpnIpsecFecName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnIpsecFec-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecFecFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnIpsecFecMappings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_bi_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bandwidth-bi-threshold"], _ = expandVpnIpsecFecMappingsBandwidthBiThreshold(d, i["bandwidth_bi_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_down_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bandwidth-down-threshold"], _ = expandVpnIpsecFecMappingsBandwidthDownThreshold(d, i["bandwidth_down_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_up_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bandwidth-up-threshold"], _ = expandVpnIpsecFecMappingsBandwidthUpThreshold(d, i["bandwidth_up_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["base"], _ = expandVpnIpsecFecMappingsBase(d, i["base"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandVpnIpsecFecMappingsLatencyThreshold(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-loss-threshold"], _ = expandVpnIpsecFecMappingsPacketLossThreshold(d, i["packet_loss_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redundant"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["redundant"], _ = expandVpnIpsecFecMappingsRedundant(d, i["redundant"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seqno"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seqno"], _ = expandVpnIpsecFecMappingsSeqno(d, i["seqno"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnIpsecFecMappingsBandwidthBiThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBandwidthDownThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBandwidthUpThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsLatencyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsPacketLossThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsRedundant(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsSeqno(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecFec(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mappings"); ok || d.HasChange("mappings") {
		t, err := expandVpnIpsecFecMappings(d, v, "mappings")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappings"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnIpsecFecName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
