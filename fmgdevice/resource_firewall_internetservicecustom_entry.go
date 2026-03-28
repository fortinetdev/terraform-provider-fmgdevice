// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Entries added to the Internet Service database and custom database.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceCustomEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceCustomEntryCreate,
		Read:   resourceFirewallInternetServiceCustomEntryRead,
		Update: resourceFirewallInternetServiceCustomEntryUpdate,
		Delete: resourceFirewallInternetServiceCustomEntryDelete,

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
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"addr_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dst6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceFirewallInternetServiceCustomEntryCreate(d *schema.ResourceData, m interface{}) error {
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
	internet_service_custom := d.Get("internet_service_custom").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["internet_service_custom"] = internet_service_custom

	obj, err := getObjectFirewallInternetServiceCustomEntry(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceCustomEntry resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallInternetServiceCustomEntry(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallInternetServiceCustomEntry(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallInternetServiceCustomEntry resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallInternetServiceCustomEntry(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallInternetServiceCustomEntry resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallInternetServiceCustomEntryRead(d, m)
}

func resourceFirewallInternetServiceCustomEntryUpdate(d *schema.ResourceData, m interface{}) error {
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
	internet_service_custom := d.Get("internet_service_custom").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["internet_service_custom"] = internet_service_custom

	obj, err := getObjectFirewallInternetServiceCustomEntry(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceCustomEntry resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallInternetServiceCustomEntry(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceCustomEntry resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallInternetServiceCustomEntryRead(d, m)
}

func resourceFirewallInternetServiceCustomEntryDelete(d *schema.ResourceData, m interface{}) error {
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
	internet_service_custom := d.Get("internet_service_custom").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["internet_service_custom"] = internet_service_custom

	wsParams["adom"] = adomv

	err = c.DeleteFirewallInternetServiceCustomEntry(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceCustomEntry resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceCustomEntryRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	internet_service_custom := d.Get("internet_service_custom").(string)
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
	if internet_service_custom == "" {
		internet_service_custom = importOptionChecking(m.(*FortiClient).Cfg, "internet_service_custom")
		if internet_service_custom == "" {
			return fmt.Errorf("Parameter internet_service_custom is missing")
		}
		if err = d.Set("internet_service_custom", internet_service_custom); err != nil {
			return fmt.Errorf("Error set params internet_service_custom: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["internet_service_custom"] = internet_service_custom

	o, err := c.ReadFirewallInternetServiceCustomEntry(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallInternetServiceCustomEntry resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceCustomEntry(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceCustomEntry resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceCustomEntryAddrMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomEntryDst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInternetServiceCustomEntryDst62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInternetServiceCustomEntryId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomEntryPortRange2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallInternetServiceCustomEntryPortRangeEndPort2edl(i["end-port"], d, pre_append)
			tmp["end_port"] = fortiAPISubPartPatch(v, "FirewallInternetServiceCustomEntry-PortRange-EndPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallInternetServiceCustomEntryPortRangeId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallInternetServiceCustomEntry-PortRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := i["start-port"]; ok {
			v := flattenFirewallInternetServiceCustomEntryPortRangeStartPort2edl(i["start-port"], d, pre_append)
			tmp["start_port"] = fortiAPISubPartPatch(v, "FirewallInternetServiceCustomEntry-PortRange-StartPort")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallInternetServiceCustomEntryPortRangeEndPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomEntryPortRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomEntryPortRangeStartPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomEntryProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceCustomEntry(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("addr_mode", flattenFirewallInternetServiceCustomEntryAddrMode2edl(o["addr-mode"], d, "addr_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-mode"], "FirewallInternetServiceCustomEntry-AddrMode"); ok {
			if err = d.Set("addr_mode", vv); err != nil {
				return fmt.Errorf("Error reading addr_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("dst", flattenFirewallInternetServiceCustomEntryDst2edl(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "FirewallInternetServiceCustomEntry-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("dst6", flattenFirewallInternetServiceCustomEntryDst62edl(o["dst6"], d, "dst6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst6"], "FirewallInternetServiceCustomEntry-Dst6"); ok {
			if err = d.Set("dst6", vv); err != nil {
				return fmt.Errorf("Error reading dst6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst6: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallInternetServiceCustomEntryId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallInternetServiceCustomEntry-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("port_range", flattenFirewallInternetServiceCustomEntryPortRange2edl(o["port-range"], d, "port_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["port-range"], "FirewallInternetServiceCustomEntry-PortRange"); ok {
				if err = d.Set("port_range", vv); err != nil {
					return fmt.Errorf("Error reading port_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading port_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("port_range"); ok {
			if err = d.Set("port_range", flattenFirewallInternetServiceCustomEntryPortRange2edl(o["port-range"], d, "port_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["port-range"], "FirewallInternetServiceCustomEntry-PortRange"); ok {
					if err = d.Set("port_range", vv); err != nil {
						return fmt.Errorf("Error reading port_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading port_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("protocol", flattenFirewallInternetServiceCustomEntryProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "FirewallInternetServiceCustomEntry-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceCustomEntryFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInternetServiceCustomEntryAddrMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomEntryDst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInternetServiceCustomEntryDst62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInternetServiceCustomEntryId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomEntryPortRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["end-port"], _ = expandFirewallInternetServiceCustomEntryPortRangeEndPort2edl(d, i["end_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallInternetServiceCustomEntryPortRangeId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-port"], _ = expandFirewallInternetServiceCustomEntryPortRangeStartPort2edl(d, i["start_port"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceCustomEntryPortRangeEndPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomEntryPortRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomEntryPortRangeStartPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomEntryProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceCustomEntry(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_mode"); ok || d.HasChange("addr_mode") {
		t, err := expandFirewallInternetServiceCustomEntryAddrMode2edl(d, v, "addr_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandFirewallInternetServiceCustomEntryDst2edl(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("dst6"); ok || d.HasChange("dst6") {
		t, err := expandFirewallInternetServiceCustomEntryDst62edl(d, v, "dst6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst6"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallInternetServiceCustomEntryId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("port_range"); ok || d.HasChange("port_range") {
		t, err := expandFirewallInternetServiceCustomEntryPortRange2edl(d, v, "port_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-range"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandFirewallInternetServiceCustomEntryProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	return &obj, nil
}
