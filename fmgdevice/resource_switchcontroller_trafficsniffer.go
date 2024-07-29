// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch RSPAN/ERSPAN traffic sniffing parameters.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerTrafficSniffer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerTrafficSnifferUpdate,
		Read:   resourceSwitchControllerTrafficSnifferRead,
		Update: resourceSwitchControllerTrafficSnifferUpdate,
		Delete: resourceSwitchControllerTrafficSnifferDelete,

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
			"erspan_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dst_entry_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"src_entry_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"target_mac": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dst_entry_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"src_entry_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"target_port": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"in_ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"out_ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"switch_id": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
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

func resourceSwitchControllerTrafficSnifferUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerTrafficSniffer(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficSniffer resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerTrafficSniffer(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficSniffer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerTrafficSniffer")

	return resourceSwitchControllerTrafficSnifferRead(d, m)
}

func resourceSwitchControllerTrafficSnifferDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerTrafficSniffer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerTrafficSniffer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerTrafficSnifferRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerTrafficSniffer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficSniffer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerTrafficSniffer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficSniffer resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerTrafficSnifferErspanIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenSwitchControllerTrafficSnifferTargetIpDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SwitchControllerTrafficSniffer-TargetIp-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_entry_id"
		if _, ok := i["dst-entry-id"]; ok {
			v := flattenSwitchControllerTrafficSnifferTargetIpDstEntryId(i["dst-entry-id"], d, pre_append)
			tmp["dst_entry_id"] = fortiAPISubPartPatch(v, "SwitchControllerTrafficSniffer-TargetIp-DstEntryId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSwitchControllerTrafficSnifferTargetIpIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SwitchControllerTrafficSniffer-TargetIp-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_entry_id"
		if _, ok := i["src-entry-id"]; ok {
			v := flattenSwitchControllerTrafficSnifferTargetIpSrcEntryId(i["src-entry-id"], d, pre_append)
			tmp["src_entry_id"] = fortiAPISubPartPatch(v, "SwitchControllerTrafficSniffer-TargetIp-SrcEntryId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerTrafficSnifferTargetIpDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetIpDstEntryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetIpIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetIpSrcEntryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetMac(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenSwitchControllerTrafficSnifferTargetMacDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SwitchControllerTrafficSniffer-TargetMac-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_entry_id"
		if _, ok := i["dst-entry-id"]; ok {
			v := flattenSwitchControllerTrafficSnifferTargetMacDstEntryId(i["dst-entry-id"], d, pre_append)
			tmp["dst_entry_id"] = fortiAPISubPartPatch(v, "SwitchControllerTrafficSniffer-TargetMac-DstEntryId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenSwitchControllerTrafficSnifferTargetMacMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "SwitchControllerTrafficSniffer-TargetMac-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_entry_id"
		if _, ok := i["src-entry-id"]; ok {
			v := flattenSwitchControllerTrafficSnifferTargetMacSrcEntryId(i["src-entry-id"], d, pre_append)
			tmp["src_entry_id"] = fortiAPISubPartPatch(v, "SwitchControllerTrafficSniffer-TargetMac-SrcEntryId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerTrafficSnifferTargetMacDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetMacDstEntryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetMacMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return case_insensitive(v, d.Get(pre))
}

func flattenSwitchControllerTrafficSnifferTargetMacSrcEntryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetPort(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenSwitchControllerTrafficSnifferTargetPortDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SwitchControllerTrafficSniffer-TargetPort-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "in_ports"
		if _, ok := i["in-ports"]; ok {
			v := flattenSwitchControllerTrafficSnifferTargetPortInPorts(i["in-ports"], d, pre_append)
			tmp["in_ports"] = fortiAPISubPartPatch(v, "SwitchControllerTrafficSniffer-TargetPort-InPorts")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "out_ports"
		if _, ok := i["out-ports"]; ok {
			v := flattenSwitchControllerTrafficSnifferTargetPortOutPorts(i["out-ports"], d, pre_append)
			tmp["out_ports"] = fortiAPISubPartPatch(v, "SwitchControllerTrafficSniffer-TargetPort-OutPorts")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := i["switch-id"]; ok {
			v := flattenSwitchControllerTrafficSnifferTargetPortSwitchId(i["switch-id"], d, pre_append)
			tmp["switch_id"] = fortiAPISubPartPatch(v, "SwitchControllerTrafficSniffer-TargetPort-SwitchId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerTrafficSnifferTargetPortDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetPortInPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerTrafficSnifferTargetPortOutPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerTrafficSnifferTargetPortSwitchId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerTrafficSniffer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("erspan_ip", flattenSwitchControllerTrafficSnifferErspanIp(o["erspan-ip"], d, "erspan_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["erspan-ip"], "SwitchControllerTrafficSniffer-ErspanIp"); ok {
			if err = d.Set("erspan_ip", vv); err != nil {
				return fmt.Errorf("Error reading erspan_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading erspan_ip: %v", err)
		}
	}

	if err = d.Set("mode", flattenSwitchControllerTrafficSnifferMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SwitchControllerTrafficSniffer-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("target_ip", flattenSwitchControllerTrafficSnifferTargetIp(o["target-ip"], d, "target_ip")); err != nil {
			if vv, ok := fortiAPIPatch(o["target-ip"], "SwitchControllerTrafficSniffer-TargetIp"); ok {
				if err = d.Set("target_ip", vv); err != nil {
					return fmt.Errorf("Error reading target_ip: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading target_ip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("target_ip"); ok {
			if err = d.Set("target_ip", flattenSwitchControllerTrafficSnifferTargetIp(o["target-ip"], d, "target_ip")); err != nil {
				if vv, ok := fortiAPIPatch(o["target-ip"], "SwitchControllerTrafficSniffer-TargetIp"); ok {
					if err = d.Set("target_ip", vv); err != nil {
						return fmt.Errorf("Error reading target_ip: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading target_ip: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("target_mac", flattenSwitchControllerTrafficSnifferTargetMac(o["target-mac"], d, "target_mac")); err != nil {
			if vv, ok := fortiAPIPatch(o["target-mac"], "SwitchControllerTrafficSniffer-TargetMac"); ok {
				if err = d.Set("target_mac", vv); err != nil {
					return fmt.Errorf("Error reading target_mac: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading target_mac: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("target_mac"); ok {
			if err = d.Set("target_mac", flattenSwitchControllerTrafficSnifferTargetMac(o["target-mac"], d, "target_mac")); err != nil {
				if vv, ok := fortiAPIPatch(o["target-mac"], "SwitchControllerTrafficSniffer-TargetMac"); ok {
					if err = d.Set("target_mac", vv); err != nil {
						return fmt.Errorf("Error reading target_mac: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading target_mac: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("target_port", flattenSwitchControllerTrafficSnifferTargetPort(o["target-port"], d, "target_port")); err != nil {
			if vv, ok := fortiAPIPatch(o["target-port"], "SwitchControllerTrafficSniffer-TargetPort"); ok {
				if err = d.Set("target_port", vv); err != nil {
					return fmt.Errorf("Error reading target_port: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading target_port: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("target_port"); ok {
			if err = d.Set("target_port", flattenSwitchControllerTrafficSnifferTargetPort(o["target-port"], d, "target_port")); err != nil {
				if vv, ok := fortiAPIPatch(o["target-port"], "SwitchControllerTrafficSniffer-TargetPort"); ok {
					if err = d.Set("target_port", vv); err != nil {
						return fmt.Errorf("Error reading target_port: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading target_port: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerTrafficSnifferFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerTrafficSnifferErspanIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandSwitchControllerTrafficSnifferTargetIpDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_entry_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-entry-id"], _ = expandSwitchControllerTrafficSnifferTargetIpDstEntryId(d, i["dst_entry_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSwitchControllerTrafficSnifferTargetIpIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_entry_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-entry-id"], _ = expandSwitchControllerTrafficSnifferTargetIpSrcEntryId(d, i["src_entry_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerTrafficSnifferTargetIpDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetIpDstEntryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetIpIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetIpSrcEntryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandSwitchControllerTrafficSnifferTargetMacDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_entry_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-entry-id"], _ = expandSwitchControllerTrafficSnifferTargetMacDstEntryId(d, i["dst_entry_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandSwitchControllerTrafficSnifferTargetMacMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_entry_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-entry-id"], _ = expandSwitchControllerTrafficSnifferTargetMacSrcEntryId(d, i["src_entry_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerTrafficSnifferTargetMacDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetMacDstEntryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetMacMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetMacSrcEntryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandSwitchControllerTrafficSnifferTargetPortDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "in_ports"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["in-ports"], _ = expandSwitchControllerTrafficSnifferTargetPortInPorts(d, i["in_ports"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "out_ports"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["out-ports"], _ = expandSwitchControllerTrafficSnifferTargetPortOutPorts(d, i["out_ports"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["switch-id"], _ = expandSwitchControllerTrafficSnifferTargetPortSwitchId(d, i["switch_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerTrafficSnifferTargetPortDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetPortInPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerTrafficSnifferTargetPortOutPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerTrafficSnifferTargetPortSwitchId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerTrafficSniffer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("erspan_ip"); ok || d.HasChange("erspan_ip") {
		t, err := expandSwitchControllerTrafficSnifferErspanIp(d, v, "erspan_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["erspan-ip"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSwitchControllerTrafficSnifferMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("target_ip"); ok || d.HasChange("target_ip") {
		t, err := expandSwitchControllerTrafficSnifferTargetIp(d, v, "target_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-ip"] = t
		}
	}

	if v, ok := d.GetOk("target_mac"); ok || d.HasChange("target_mac") {
		t, err := expandSwitchControllerTrafficSnifferTargetMac(d, v, "target_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-mac"] = t
		}
	}

	if v, ok := d.GetOk("target_port"); ok || d.HasChange("target_port") {
		t, err := expandSwitchControllerTrafficSnifferTargetPort(d, v, "target_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-port"] = t
		}
	}

	return &obj, nil
}
