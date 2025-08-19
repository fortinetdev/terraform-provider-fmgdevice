// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IP address management services.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIpam() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpamUpdate,
		Read:   resourceSystemIpamRead,
		Update: resourceSystemIpamUpdate,
		Delete: resourceSystemIpamDelete,

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
			"pool_subnet": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"automatic_conflict_resolution": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"manage_lan_addresses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"manage_lan_extension_addresses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"manage_ssid_addresses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pools": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"exclude": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"exclude_subnet": &schema.Schema{
										Type:     schema.TypeList,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subnet": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"require_subnet_size_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"device": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dhcp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pool": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
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

func resourceSystemIpamUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemIpam(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpam resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIpam(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpam resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemIpam")

	return resourceSystemIpamRead(d, m)
}

func resourceSystemIpamDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemIpam(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIpam resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpamRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemIpam(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpam resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpam(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpam resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpamPoolSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemIpamAutomaticConflictResolution(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamManageLanAddresses(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamManageLanExtensionAddresses(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamManageSsidAddresses(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamPools(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemIpamPoolsDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SystemIpam-Pools-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude"
		if _, ok := i["exclude"]; ok {
			v := flattenSystemIpamPoolsExclude(i["exclude"], d, pre_append)
			tmp["exclude"] = fortiAPISubPartPatch(v, "SystemIpam-Pools-Exclude")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemIpamPoolsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemIpam-Pools-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenSystemIpamPoolsSubnet(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "SystemIpam-Pools-Subnet")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemIpamPoolsDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamPoolsExclude(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		if _, ok := i["ID"]; ok {
			v := flattenSystemIpamPoolsExcludeId(i["ID"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemIpamPools-Exclude-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude_subnet"
		if _, ok := i["exclude-subnet"]; ok {
			v := flattenSystemIpamPoolsExcludeExcludeSubnet(i["exclude-subnet"], d, pre_append)
			tmp["exclude_subnet"] = fortiAPISubPartPatch(v, "SystemIpamPools-Exclude-ExcludeSubnet")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemIpamPoolsExcludeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamPoolsExcludeExcludeSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemIpamPoolsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamPoolsSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemIpamRequireSubnetSizeMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamRules(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemIpamRulesDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SystemIpam-Rules-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := i["device"]; ok {
			v := flattenSystemIpamRulesDevice(i["device"], d, pre_append)
			tmp["device"] = fortiAPISubPartPatch(v, "SystemIpam-Rules-Device")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp"
		if _, ok := i["dhcp"]; ok {
			v := flattenSystemIpamRulesDhcp(i["dhcp"], d, pre_append)
			tmp["dhcp"] = fortiAPISubPartPatch(v, "SystemIpam-Rules-Dhcp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystemIpamRulesInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystemIpam-Rules-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemIpamRulesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemIpam-Rules-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pool"
		if _, ok := i["pool"]; ok {
			v := flattenSystemIpamRulesPool(i["pool"], d, pre_append)
			tmp["pool"] = fortiAPISubPartPatch(v, "SystemIpam-Rules-Pool")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenSystemIpamRulesRole(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "SystemIpam-Rules-Role")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemIpamRulesDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamRulesDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemIpamRulesDhcp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamRulesInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemIpamRulesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamRulesPool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemIpamRulesRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemIpam(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("pool_subnet", flattenSystemIpamPoolSubnet(o["pool-subnet"], d, "pool_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["pool-subnet"], "SystemIpam-PoolSubnet"); ok {
			if err = d.Set("pool_subnet", vv); err != nil {
				return fmt.Errorf("Error reading pool_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pool_subnet: %v", err)
		}
	}

	if err = d.Set("automatic_conflict_resolution", flattenSystemIpamAutomaticConflictResolution(o["automatic-conflict-resolution"], d, "automatic_conflict_resolution")); err != nil {
		if vv, ok := fortiAPIPatch(o["automatic-conflict-resolution"], "SystemIpam-AutomaticConflictResolution"); ok {
			if err = d.Set("automatic_conflict_resolution", vv); err != nil {
				return fmt.Errorf("Error reading automatic_conflict_resolution: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading automatic_conflict_resolution: %v", err)
		}
	}

	if err = d.Set("manage_lan_addresses", flattenSystemIpamManageLanAddresses(o["manage-lan-addresses"], d, "manage_lan_addresses")); err != nil {
		if vv, ok := fortiAPIPatch(o["manage-lan-addresses"], "SystemIpam-ManageLanAddresses"); ok {
			if err = d.Set("manage_lan_addresses", vv); err != nil {
				return fmt.Errorf("Error reading manage_lan_addresses: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading manage_lan_addresses: %v", err)
		}
	}

	if err = d.Set("manage_lan_extension_addresses", flattenSystemIpamManageLanExtensionAddresses(o["manage-lan-extension-addresses"], d, "manage_lan_extension_addresses")); err != nil {
		if vv, ok := fortiAPIPatch(o["manage-lan-extension-addresses"], "SystemIpam-ManageLanExtensionAddresses"); ok {
			if err = d.Set("manage_lan_extension_addresses", vv); err != nil {
				return fmt.Errorf("Error reading manage_lan_extension_addresses: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading manage_lan_extension_addresses: %v", err)
		}
	}

	if err = d.Set("manage_ssid_addresses", flattenSystemIpamManageSsidAddresses(o["manage-ssid-addresses"], d, "manage_ssid_addresses")); err != nil {
		if vv, ok := fortiAPIPatch(o["manage-ssid-addresses"], "SystemIpam-ManageSsidAddresses"); ok {
			if err = d.Set("manage_ssid_addresses", vv); err != nil {
				return fmt.Errorf("Error reading manage_ssid_addresses: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading manage_ssid_addresses: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("pools", flattenSystemIpamPools(o["pools"], d, "pools")); err != nil {
			if vv, ok := fortiAPIPatch(o["pools"], "SystemIpam-Pools"); ok {
				if err = d.Set("pools", vv); err != nil {
					return fmt.Errorf("Error reading pools: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pools: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pools"); ok {
			if err = d.Set("pools", flattenSystemIpamPools(o["pools"], d, "pools")); err != nil {
				if vv, ok := fortiAPIPatch(o["pools"], "SystemIpam-Pools"); ok {
					if err = d.Set("pools", vv); err != nil {
						return fmt.Errorf("Error reading pools: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pools: %v", err)
				}
			}
		}
	}

	if err = d.Set("require_subnet_size_match", flattenSystemIpamRequireSubnetSizeMatch(o["require-subnet-size-match"], d, "require_subnet_size_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["require-subnet-size-match"], "SystemIpam-RequireSubnetSizeMatch"); ok {
			if err = d.Set("require_subnet_size_match", vv); err != nil {
				return fmt.Errorf("Error reading require_subnet_size_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading require_subnet_size_match: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rules", flattenSystemIpamRules(o["rules"], d, "rules")); err != nil {
			if vv, ok := fortiAPIPatch(o["rules"], "SystemIpam-Rules"); ok {
				if err = d.Set("rules", vv); err != nil {
					return fmt.Errorf("Error reading rules: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rules"); ok {
			if err = d.Set("rules", flattenSystemIpamRules(o["rules"], d, "rules")); err != nil {
				if vv, ok := fortiAPIPatch(o["rules"], "SystemIpam-Rules"); ok {
					if err = d.Set("rules", vv); err != nil {
						return fmt.Errorf("Error reading rules: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rules: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenSystemIpamServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "SystemIpam-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemIpamStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemIpam-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemIpamFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemIpamPoolSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemIpamAutomaticConflictResolution(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamManageLanAddresses(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamManageLanExtensionAddresses(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamManageSsidAddresses(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamPools(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["description"], _ = expandSystemIpamPoolsDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSystemIpamPoolsExclude(d, i["exclude"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["exclude"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemIpamPoolsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet"], _ = expandSystemIpamPoolsSubnet(d, i["subnet"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemIpamPoolsDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamPoolsExclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["ID"], _ = expandSystemIpamPoolsExcludeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude_subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exclude-subnet"], _ = expandSystemIpamPoolsExcludeExcludeSubnet(d, i["exclude_subnet"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemIpamPoolsExcludeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamPoolsExcludeExcludeSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemIpamPoolsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamPoolsSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemIpamRequireSubnetSizeMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamRules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["description"], _ = expandSystemIpamRulesDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device"], _ = expandSystemIpamRulesDevice(d, i["device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp"], _ = expandSystemIpamRulesDhcp(d, i["dhcp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystemIpamRulesInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemIpamRulesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pool"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pool"], _ = expandSystemIpamRulesPool(d, i["pool"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandSystemIpamRulesRole(d, i["role"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemIpamRulesDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamRulesDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemIpamRulesDhcp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamRulesInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemIpamRulesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamRulesPool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemIpamRulesRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIpam(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("pool_subnet"); ok || d.HasChange("pool_subnet") {
		t, err := expandSystemIpamPoolSubnet(d, v, "pool_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pool-subnet"] = t
		}
	}

	if v, ok := d.GetOk("automatic_conflict_resolution"); ok || d.HasChange("automatic_conflict_resolution") {
		t, err := expandSystemIpamAutomaticConflictResolution(d, v, "automatic_conflict_resolution")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["automatic-conflict-resolution"] = t
		}
	}

	if v, ok := d.GetOk("manage_lan_addresses"); ok || d.HasChange("manage_lan_addresses") {
		t, err := expandSystemIpamManageLanAddresses(d, v, "manage_lan_addresses")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["manage-lan-addresses"] = t
		}
	}

	if v, ok := d.GetOk("manage_lan_extension_addresses"); ok || d.HasChange("manage_lan_extension_addresses") {
		t, err := expandSystemIpamManageLanExtensionAddresses(d, v, "manage_lan_extension_addresses")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["manage-lan-extension-addresses"] = t
		}
	}

	if v, ok := d.GetOk("manage_ssid_addresses"); ok || d.HasChange("manage_ssid_addresses") {
		t, err := expandSystemIpamManageSsidAddresses(d, v, "manage_ssid_addresses")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["manage-ssid-addresses"] = t
		}
	}

	if v, ok := d.GetOk("pools"); ok || d.HasChange("pools") {
		t, err := expandSystemIpamPools(d, v, "pools")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pools"] = t
		}
	}

	if v, ok := d.GetOk("require_subnet_size_match"); ok || d.HasChange("require_subnet_size_match") {
		t, err := expandSystemIpamRequireSubnetSizeMatch(d, v, "require_subnet_size_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["require-subnet-size-match"] = t
		}
	}

	if v, ok := d.GetOk("rules"); ok || d.HasChange("rules") {
		t, err := expandSystemIpamRules(d, v, "rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rules"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandSystemIpamServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemIpamStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
