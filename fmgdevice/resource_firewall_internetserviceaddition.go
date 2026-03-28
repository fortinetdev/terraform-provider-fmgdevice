// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Internet Services Addition.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceAddition() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceAdditionCreate,
		Read:   resourceFirewallInternetServiceAdditionRead,
		Update: resourceFirewallInternetServiceAdditionUpdate,
		Delete: resourceFirewallInternetServiceAdditionDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port_range": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"start_port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceFirewallInternetServiceAdditionCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectFirewallInternetServiceAddition(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceAddition resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallInternetServiceAddition(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallInternetServiceAddition(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallInternetServiceAddition resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallInternetServiceAddition(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallInternetServiceAddition resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "fosid"))

	return resourceFirewallInternetServiceAdditionRead(d, m)
}

func resourceFirewallInternetServiceAdditionUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectFirewallInternetServiceAddition(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceAddition resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallInternetServiceAddition(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceAddition resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "fosid"))

	return resourceFirewallInternetServiceAdditionRead(d, m)
}

func resourceFirewallInternetServiceAdditionDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteFirewallInternetServiceAddition(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceAddition resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceAdditionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallInternetServiceAddition(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallInternetServiceAddition resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceAddition(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceAddition resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceAdditionComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntry(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := i["addr-mode"]; ok {
			v := flattenFirewallInternetServiceAdditionEntryAddrMode(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "FirewallInternetServiceAddition-Entry-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallInternetServiceAdditionEntryId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallInternetServiceAddition-Entry-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_range"
		if _, ok := i["port-range"]; ok {
			v := flattenFirewallInternetServiceAdditionEntryPortRange(i["port-range"], d, pre_append)
			tmp["port_range"] = fortiAPISubPartPatch(v, "FirewallInternetServiceAddition-Entry-PortRange")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenFirewallInternetServiceAdditionEntryProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "FirewallInternetServiceAddition-Entry-Protocol")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallInternetServiceAdditionEntryAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntryPortRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := i["end-port"]; ok {
			v := flattenFirewallInternetServiceAdditionEntryPortRangeEndPort(i["end-port"], d, pre_append)
			tmp["end_port"] = fortiAPISubPartPatch(v, "FirewallInternetServiceAdditionEntry-PortRange-EndPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallInternetServiceAdditionEntryPortRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallInternetServiceAdditionEntry-PortRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := i["start-port"]; ok {
			v := flattenFirewallInternetServiceAdditionEntryPortRangeStartPort(i["start-port"], d, pre_append)
			tmp["start_port"] = fortiAPISubPartPatch(v, "FirewallInternetServiceAdditionEntry-PortRange-StartPort")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallInternetServiceAdditionEntryPortRangeEndPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntryPortRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntryPortRangeStartPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntryProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func refreshObjectFirewallInternetServiceAddition(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenFirewallInternetServiceAdditionComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallInternetServiceAddition-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entry", flattenFirewallInternetServiceAdditionEntry(o["entry"], d, "entry")); err != nil {
			if vv, ok := fortiAPIPatch(o["entry"], "FirewallInternetServiceAddition-Entry"); ok {
				if err = d.Set("entry", vv); err != nil {
					return fmt.Errorf("Error reading entry: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entry: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entry"); ok {
			if err = d.Set("entry", flattenFirewallInternetServiceAdditionEntry(o["entry"], d, "entry")); err != nil {
				if vv, ok := fortiAPIPatch(o["entry"], "FirewallInternetServiceAddition-Entry"); ok {
					if err = d.Set("entry", vv); err != nil {
						return fmt.Errorf("Error reading entry: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entry: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenFirewallInternetServiceAdditionId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallInternetServiceAddition-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceAdditionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInternetServiceAdditionComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-mode"], _ = expandFirewallInternetServiceAdditionEntryAddrMode(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallInternetServiceAdditionEntryId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallInternetServiceAdditionEntryPortRange(d, i["port_range"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["port-range"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandFirewallInternetServiceAdditionEntryProtocol(d, i["protocol"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceAdditionEntryAddrMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntryPortRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-port"], _ = expandFirewallInternetServiceAdditionEntryPortRangeEndPort(d, i["end_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallInternetServiceAdditionEntryPortRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-port"], _ = expandFirewallInternetServiceAdditionEntryPortRangeStartPort(d, i["start_port"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceAdditionEntryPortRangeEndPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntryPortRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntryPortRangeStartPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntryProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceAddition(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallInternetServiceAdditionComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entry"); ok || d.HasChange("entry") {
		t, err := expandFirewallInternetServiceAdditionEntry(d, v, "entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entry"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallInternetServiceAdditionId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
