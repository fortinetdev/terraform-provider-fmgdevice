// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Entries added to the Internet Service addition database.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceAdditionEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceAdditionEntryCreate,
		Read:   resourceFirewallInternetServiceAdditionEntryRead,
		Update: resourceFirewallInternetServiceAdditionEntryUpdate,
		Delete: resourceFirewallInternetServiceAdditionEntryDelete,

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
			"internet_service_addition": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"addr_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallInternetServiceAdditionEntryCreate(d *schema.ResourceData, m interface{}) error {
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
	internet_service_addition := d.Get("internet_service_addition").(string)
	paradict["device"] = device_name
	paradict["internet_service_addition"] = internet_service_addition

	obj, err := getObjectFirewallInternetServiceAdditionEntry(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceAdditionEntry resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallInternetServiceAdditionEntry(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallInternetServiceAdditionEntry(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallInternetServiceAdditionEntry resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallInternetServiceAdditionEntry(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallInternetServiceAdditionEntry resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallInternetServiceAdditionEntryRead(d, m)
}

func resourceFirewallInternetServiceAdditionEntryUpdate(d *schema.ResourceData, m interface{}) error {
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
	internet_service_addition := d.Get("internet_service_addition").(string)
	paradict["device"] = device_name
	paradict["internet_service_addition"] = internet_service_addition

	obj, err := getObjectFirewallInternetServiceAdditionEntry(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceAdditionEntry resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallInternetServiceAdditionEntry(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceAdditionEntry resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallInternetServiceAdditionEntryRead(d, m)
}

func resourceFirewallInternetServiceAdditionEntryDelete(d *schema.ResourceData, m interface{}) error {
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
	internet_service_addition := d.Get("internet_service_addition").(string)
	paradict["device"] = device_name
	paradict["internet_service_addition"] = internet_service_addition

	wsParams["adom"] = adomv

	err = c.DeleteFirewallInternetServiceAdditionEntry(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceAdditionEntry resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceAdditionEntryRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	internet_service_addition := d.Get("internet_service_addition").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if internet_service_addition == "" {
		internet_service_addition = importOptionChecking(m.(*FortiClient).Cfg, "internet_service_addition")
		if internet_service_addition == "" {
			return fmt.Errorf("Parameter internet_service_addition is missing")
		}
		if err = d.Set("internet_service_addition", internet_service_addition); err != nil {
			return fmt.Errorf("Error set params internet_service_addition: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["internet_service_addition"] = internet_service_addition

	o, err := c.ReadFirewallInternetServiceAdditionEntry(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallInternetServiceAdditionEntry resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceAdditionEntry(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceAdditionEntry resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceAdditionEntryAddrMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntryId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntryPortRange2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallInternetServiceAdditionEntryPortRangeEndPort2edl(i["end-port"], d, pre_append)
			tmp["end_port"] = fortiAPISubPartPatch(v, "FirewallInternetServiceAdditionEntry-PortRange-EndPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallInternetServiceAdditionEntryPortRangeId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallInternetServiceAdditionEntry-PortRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := i["start-port"]; ok {
			v := flattenFirewallInternetServiceAdditionEntryPortRangeStartPort2edl(i["start-port"], d, pre_append)
			tmp["start_port"] = fortiAPISubPartPatch(v, "FirewallInternetServiceAdditionEntry-PortRange-StartPort")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallInternetServiceAdditionEntryPortRangeEndPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntryPortRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntryPortRangeStartPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAdditionEntryProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceAdditionEntry(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("addr_mode", flattenFirewallInternetServiceAdditionEntryAddrMode2edl(o["addr-mode"], d, "addr_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-mode"], "FirewallInternetServiceAdditionEntry-AddrMode"); ok {
			if err = d.Set("addr_mode", vv); err != nil {
				return fmt.Errorf("Error reading addr_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallInternetServiceAdditionEntryId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallInternetServiceAdditionEntry-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("port_range", flattenFirewallInternetServiceAdditionEntryPortRange2edl(o["port-range"], d, "port_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["port-range"], "FirewallInternetServiceAdditionEntry-PortRange"); ok {
				if err = d.Set("port_range", vv); err != nil {
					return fmt.Errorf("Error reading port_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading port_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("port_range"); ok {
			if err = d.Set("port_range", flattenFirewallInternetServiceAdditionEntryPortRange2edl(o["port-range"], d, "port_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["port-range"], "FirewallInternetServiceAdditionEntry-PortRange"); ok {
					if err = d.Set("port_range", vv); err != nil {
						return fmt.Errorf("Error reading port_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading port_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("protocol", flattenFirewallInternetServiceAdditionEntryProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "FirewallInternetServiceAdditionEntry-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceAdditionEntryFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInternetServiceAdditionEntryAddrMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntryId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntryPortRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["end-port"], _ = expandFirewallInternetServiceAdditionEntryPortRangeEndPort2edl(d, i["end_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallInternetServiceAdditionEntryPortRangeId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-port"], _ = expandFirewallInternetServiceAdditionEntryPortRangeStartPort2edl(d, i["start_port"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceAdditionEntryPortRangeEndPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntryPortRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntryPortRangeStartPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAdditionEntryProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceAdditionEntry(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_mode"); ok || d.HasChange("addr_mode") {
		t, err := expandFirewallInternetServiceAdditionEntryAddrMode2edl(d, v, "addr_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallInternetServiceAdditionEntryId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("port_range"); ok || d.HasChange("port_range") {
		t, err := expandFirewallInternetServiceAdditionEntryPortRange2edl(d, v, "port_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-range"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandFirewallInternetServiceAdditionEntryProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	return &obj, nil
}
