// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure inter wireless controller operation.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerInterController() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerInterControllerUpdate,
		Read:   resourceWirelessControllerInterControllerRead,
		Update: resourceWirelessControllerInterControllerUpdate,
		Delete: resourceWirelessControllerInterControllerDelete,

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
			"fast_failover_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fast_failover_wait": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"inter_controller_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"inter_controller_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inter_controller_peer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"peer_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"peer_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"peer_priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"inter_controller_pri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l3_roaming": &schema.Schema{
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

func resourceWirelessControllerInterControllerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerInterController(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerInterController resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerInterController(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerInterController resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerInterController")

	return resourceWirelessControllerInterControllerRead(d, m)
}

func resourceWirelessControllerInterControllerDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerInterController(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerInterController resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerInterControllerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerInterController(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerInterController resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerInterController(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerInterController resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerInterControllerFastFailoverMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerFastFailoverWait(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPeer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWirelessControllerInterControllerInterControllerPeerId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerInterController-InterControllerPeer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_ip"
		if _, ok := i["peer-ip"]; ok {
			v := flattenWirelessControllerInterControllerInterControllerPeerPeerIp(i["peer-ip"], d, pre_append)
			tmp["peer_ip"] = fortiAPISubPartPatch(v, "WirelessControllerInterController-InterControllerPeer-PeerIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_port"
		if _, ok := i["peer-port"]; ok {
			v := flattenWirelessControllerInterControllerInterControllerPeerPeerPort(i["peer-port"], d, pre_append)
			tmp["peer_port"] = fortiAPISubPartPatch(v, "WirelessControllerInterController-InterControllerPeer-PeerPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_priority"
		if _, ok := i["peer-priority"]; ok {
			v := flattenWirelessControllerInterControllerInterControllerPeerPeerPriority(i["peer-priority"], d, pre_append)
			tmp["peer_priority"] = fortiAPISubPartPatch(v, "WirelessControllerInterController-InterControllerPeer-PeerPriority")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerInterControllerInterControllerPeerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPeerPeerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPeerPeerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPeerPeerPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerL3Roaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerInterController(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fast_failover_max", flattenWirelessControllerInterControllerFastFailoverMax(o["fast-failover-max"], d, "fast_failover_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["fast-failover-max"], "WirelessControllerInterController-FastFailoverMax"); ok {
			if err = d.Set("fast_failover_max", vv); err != nil {
				return fmt.Errorf("Error reading fast_failover_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fast_failover_max: %v", err)
		}
	}

	if err = d.Set("fast_failover_wait", flattenWirelessControllerInterControllerFastFailoverWait(o["fast-failover-wait"], d, "fast_failover_wait")); err != nil {
		if vv, ok := fortiAPIPatch(o["fast-failover-wait"], "WirelessControllerInterController-FastFailoverWait"); ok {
			if err = d.Set("fast_failover_wait", vv); err != nil {
				return fmt.Errorf("Error reading fast_failover_wait: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fast_failover_wait: %v", err)
		}
	}

	if err = d.Set("inter_controller_mode", flattenWirelessControllerInterControllerInterControllerMode(o["inter-controller-mode"], d, "inter_controller_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["inter-controller-mode"], "WirelessControllerInterController-InterControllerMode"); ok {
			if err = d.Set("inter_controller_mode", vv); err != nil {
				return fmt.Errorf("Error reading inter_controller_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inter_controller_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("inter_controller_peer", flattenWirelessControllerInterControllerInterControllerPeer(o["inter-controller-peer"], d, "inter_controller_peer")); err != nil {
			if vv, ok := fortiAPIPatch(o["inter-controller-peer"], "WirelessControllerInterController-InterControllerPeer"); ok {
				if err = d.Set("inter_controller_peer", vv); err != nil {
					return fmt.Errorf("Error reading inter_controller_peer: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading inter_controller_peer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("inter_controller_peer"); ok {
			if err = d.Set("inter_controller_peer", flattenWirelessControllerInterControllerInterControllerPeer(o["inter-controller-peer"], d, "inter_controller_peer")); err != nil {
				if vv, ok := fortiAPIPatch(o["inter-controller-peer"], "WirelessControllerInterController-InterControllerPeer"); ok {
					if err = d.Set("inter_controller_peer", vv); err != nil {
						return fmt.Errorf("Error reading inter_controller_peer: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading inter_controller_peer: %v", err)
				}
			}
		}
	}

	if err = d.Set("inter_controller_pri", flattenWirelessControllerInterControllerInterControllerPri(o["inter-controller-pri"], d, "inter_controller_pri")); err != nil {
		if vv, ok := fortiAPIPatch(o["inter-controller-pri"], "WirelessControllerInterController-InterControllerPri"); ok {
			if err = d.Set("inter_controller_pri", vv); err != nil {
				return fmt.Errorf("Error reading inter_controller_pri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inter_controller_pri: %v", err)
		}
	}

	if err = d.Set("l3_roaming", flattenWirelessControllerInterControllerL3Roaming(o["l3-roaming"], d, "l3_roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["l3-roaming"], "WirelessControllerInterController-L3Roaming"); ok {
			if err = d.Set("l3_roaming", vv); err != nil {
				return fmt.Errorf("Error reading l3_roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l3_roaming: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerInterControllerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerInterControllerFastFailoverMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerFastFailoverWait(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerInterControllerInterControllerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWirelessControllerInterControllerInterControllerPeerId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["peer-ip"], _ = expandWirelessControllerInterControllerInterControllerPeerPeerIp(d, i["peer_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["peer-port"], _ = expandWirelessControllerInterControllerInterControllerPeerPeerPort(d, i["peer_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["peer-priority"], _ = expandWirelessControllerInterControllerInterControllerPeerPeerPriority(d, i["peer_priority"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerInterControllerInterControllerPeerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPeerPeerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPeerPeerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPeerPeerPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerL3Roaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerInterController(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fast_failover_max"); ok || d.HasChange("fast_failover_max") {
		t, err := expandWirelessControllerInterControllerFastFailoverMax(d, v, "fast_failover_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-failover-max"] = t
		}
	}

	if v, ok := d.GetOk("fast_failover_wait"); ok || d.HasChange("fast_failover_wait") {
		t, err := expandWirelessControllerInterControllerFastFailoverWait(d, v, "fast_failover_wait")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-failover-wait"] = t
		}
	}

	if v, ok := d.GetOk("inter_controller_key"); ok || d.HasChange("inter_controller_key") {
		t, err := expandWirelessControllerInterControllerInterControllerKey(d, v, "inter_controller_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-controller-key"] = t
		}
	}

	if v, ok := d.GetOk("inter_controller_mode"); ok || d.HasChange("inter_controller_mode") {
		t, err := expandWirelessControllerInterControllerInterControllerMode(d, v, "inter_controller_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-controller-mode"] = t
		}
	}

	if v, ok := d.GetOk("inter_controller_peer"); ok || d.HasChange("inter_controller_peer") {
		t, err := expandWirelessControllerInterControllerInterControllerPeer(d, v, "inter_controller_peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-controller-peer"] = t
		}
	}

	if v, ok := d.GetOk("inter_controller_pri"); ok || d.HasChange("inter_controller_pri") {
		t, err := expandWirelessControllerInterControllerInterControllerPri(d, v, "inter_controller_pri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-controller-pri"] = t
		}
	}

	if v, ok := d.GetOk("l3_roaming"); ok || d.HasChange("l3_roaming") {
		t, err := expandWirelessControllerInterControllerL3Roaming(d, v, "l3_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-roaming"] = t
		}
	}

	return &obj, nil
}
