// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch IGMP snooping global settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchIgmpSnooping() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchIgmpSnoopingUpdate,
		Read:   resourceSwitchControllerManagedSwitchIgmpSnoopingRead,
		Update: resourceSwitchControllerManagedSwitchIgmpSnoopingUpdate,
		Delete: resourceSwitchControllerManagedSwitchIgmpSnoopingDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"aging_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"flood_unknown_multicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlans": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"querier": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"querier_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"vlan_name": &schema.Schema{
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

func resourceSwitchControllerManagedSwitchIgmpSnoopingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchIgmpSnooping(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchIgmpSnooping resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchIgmpSnooping(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchIgmpSnooping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerManagedSwitchIgmpSnooping")

	return resourceSwitchControllerManagedSwitchIgmpSnoopingRead(d, m)
}

func resourceSwitchControllerManagedSwitchIgmpSnoopingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerManagedSwitchIgmpSnooping(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchIgmpSnooping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchIgmpSnoopingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	managed_switch := d.Get("managed_switch").(string)
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
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	o, err := c.ReadSwitchControllerManagedSwitchIgmpSnooping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchIgmpSnooping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchIgmpSnooping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchIgmpSnooping resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingAgingTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingFloodUnknownMulticast2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingLocalOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlans2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy"
		if _, ok := i["proxy"]; ok {
			v := flattenSwitchControllerManagedSwitchIgmpSnoopingVlansProxy2edl(i["proxy"], d, pre_append)
			tmp["proxy"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIgmpSnooping-Vlans-Proxy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "querier"
		if _, ok := i["querier"]; ok {
			v := flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier2edl(i["querier"], d, pre_append)
			tmp["querier"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIgmpSnooping-Vlans-Querier")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "querier_addr"
		if _, ok := i["querier-addr"]; ok {
			v := flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr2edl(i["querier-addr"], d, pre_append)
			tmp["querier_addr"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIgmpSnooping-Vlans-QuerierAddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVersion2edl(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIgmpSnooping-Vlans-Version")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := i["vlan-name"]; ok {
			v := flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName2edl(i["vlan-name"], d, pre_append)
			tmp["vlan_name"] = fortiAPISubPartPatch(v, "SwitchControllerManagedSwitchIgmpSnooping-Vlans-VlanName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansProxy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerManagedSwitchIgmpSnooping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("aging_time", flattenSwitchControllerManagedSwitchIgmpSnoopingAgingTime2edl(o["aging-time"], d, "aging_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["aging-time"], "SwitchControllerManagedSwitchIgmpSnooping-AgingTime"); ok {
			if err = d.Set("aging_time", vv); err != nil {
				return fmt.Errorf("Error reading aging_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aging_time: %v", err)
		}
	}

	if err = d.Set("flood_unknown_multicast", flattenSwitchControllerManagedSwitchIgmpSnoopingFloodUnknownMulticast2edl(o["flood-unknown-multicast"], d, "flood_unknown_multicast")); err != nil {
		if vv, ok := fortiAPIPatch(o["flood-unknown-multicast"], "SwitchControllerManagedSwitchIgmpSnooping-FloodUnknownMulticast"); ok {
			if err = d.Set("flood_unknown_multicast", vv); err != nil {
				return fmt.Errorf("Error reading flood_unknown_multicast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flood_unknown_multicast: %v", err)
		}
	}

	if err = d.Set("local_override", flattenSwitchControllerManagedSwitchIgmpSnoopingLocalOverride2edl(o["local-override"], d, "local_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-override"], "SwitchControllerManagedSwitchIgmpSnooping-LocalOverride"); ok {
			if err = d.Set("local_override", vv); err != nil {
				return fmt.Errorf("Error reading local_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_override: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vlans", flattenSwitchControllerManagedSwitchIgmpSnoopingVlans2edl(o["vlans"], d, "vlans")); err != nil {
			if vv, ok := fortiAPIPatch(o["vlans"], "SwitchControllerManagedSwitchIgmpSnooping-Vlans"); ok {
				if err = d.Set("vlans", vv); err != nil {
					return fmt.Errorf("Error reading vlans: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vlans: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vlans"); ok {
			if err = d.Set("vlans", flattenSwitchControllerManagedSwitchIgmpSnoopingVlans2edl(o["vlans"], d, "vlans")); err != nil {
				if vv, ok := fortiAPIPatch(o["vlans"], "SwitchControllerManagedSwitchIgmpSnooping-Vlans"); ok {
					if err = d.Set("vlans", vv); err != nil {
						return fmt.Errorf("Error reading vlans: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vlans: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchIgmpSnoopingAgingTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingFloodUnknownMulticast2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingLocalOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlans2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["proxy"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansProxy2edl(d, i["proxy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "querier"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["querier"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier2edl(d, i["querier"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "querier_addr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["querier-addr"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr2edl(d, i["querier_addr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["version"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansVersion2edl(d, i["version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName2edl(d, i["vlan_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansProxy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerManagedSwitchIgmpSnooping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aging_time"); ok || d.HasChange("aging_time") {
		t, err := expandSwitchControllerManagedSwitchIgmpSnoopingAgingTime2edl(d, v, "aging_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aging-time"] = t
		}
	}

	if v, ok := d.GetOk("flood_unknown_multicast"); ok || d.HasChange("flood_unknown_multicast") {
		t, err := expandSwitchControllerManagedSwitchIgmpSnoopingFloodUnknownMulticast2edl(d, v, "flood_unknown_multicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flood-unknown-multicast"] = t
		}
	}

	if v, ok := d.GetOk("local_override"); ok || d.HasChange("local_override") {
		t, err := expandSwitchControllerManagedSwitchIgmpSnoopingLocalOverride2edl(d, v, "local_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-override"] = t
		}
	}

	if v, ok := d.GetOk("vlans"); ok || d.HasChange("vlans") {
		t, err := expandSwitchControllerManagedSwitchIgmpSnoopingVlans2edl(d, v, "vlans")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlans"] = t
		}
	}

	return &obj, nil
}
