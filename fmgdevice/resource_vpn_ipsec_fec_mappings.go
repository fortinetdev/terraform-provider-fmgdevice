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
				ForceNew: true,
				Optional: true,
				Computed: true,
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

	_, err = c.UpdateVpnIpsecFecMappings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecFecMappings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "seqno")))

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

func flattenVpnIpsecFecMappingsBandwidthDownThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBandwidthUpThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsLatencyThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsPacketLossThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsRedundant2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsSeqno2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnIpsecFecMappings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bandwidth_bi_threshold", flattenVpnIpsecFecMappingsBandwidthBiThreshold2edl(o["bandwidth-bi-threshold"], d, "bandwidth_bi_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-bi-threshold"], "VpnIpsecFecMappings-BandwidthBiThreshold"); ok {
			if err = d.Set("bandwidth_bi_threshold", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_bi_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_bi_threshold: %v", err)
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

	if err = d.Set("bandwidth_up_threshold", flattenVpnIpsecFecMappingsBandwidthUpThreshold2edl(o["bandwidth-up-threshold"], d, "bandwidth_up_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-up-threshold"], "VpnIpsecFecMappings-BandwidthUpThreshold"); ok {
			if err = d.Set("bandwidth_up_threshold", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_up_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_up_threshold: %v", err)
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

	if err = d.Set("packet_loss_threshold", flattenVpnIpsecFecMappingsPacketLossThreshold2edl(o["packet-loss-threshold"], d, "packet_loss_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-loss-threshold"], "VpnIpsecFecMappings-PacketLossThreshold"); ok {
			if err = d.Set("packet_loss_threshold", vv); err != nil {
				return fmt.Errorf("Error reading packet_loss_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_loss_threshold: %v", err)
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

func expandVpnIpsecFecMappingsBandwidthDownThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBandwidthUpThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsLatencyThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsPacketLossThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsRedundant2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsSeqno2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

	if v, ok := d.GetOk("bandwidth_down_threshold"); ok || d.HasChange("bandwidth_down_threshold") {
		t, err := expandVpnIpsecFecMappingsBandwidthDownThreshold2edl(d, v, "bandwidth_down_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-down-threshold"] = t
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

	if v, ok := d.GetOk("packet_loss_threshold"); ok || d.HasChange("packet_loss_threshold") {
		t, err := expandVpnIpsecFecMappingsPacketLossThreshold2edl(d, v, "packet_loss_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-loss-threshold"] = t
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

	return &obj, nil
}
