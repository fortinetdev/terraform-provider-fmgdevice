// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure MPSK profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerMpskProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerMpskProfileCreate,
		Read:   resourceWirelessControllerMpskProfileRead,
		Update: resourceWirelessControllerMpskProfileUpdate,
		Delete: resourceWirelessControllerMpskProfileDelete,

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
			"mpsk_concurrent_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mpsk_external_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mpsk_external_server_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mpsk_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mpsk_key": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"comment": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"concurrent_client_limit_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"concurrent_clients": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"key_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mpsk_schedules": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"passphrase": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"pmk": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"sae_password": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"sae_pk": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sae_private_key": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vlan_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mpsk_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ssid": &schema.Schema{
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

func resourceWirelessControllerMpskProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerMpskProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerMpskProfile resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerMpskProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerMpskProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerMpskProfileRead(d, m)
}

func resourceWirelessControllerMpskProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerMpskProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerMpskProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerMpskProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerMpskProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerMpskProfileRead(d, m)
}

func resourceWirelessControllerMpskProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerMpskProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerMpskProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerMpskProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerMpskProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerMpskProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerMpskProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerMpskProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerMpskProfileMpskConcurrentClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskExternalServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerMpskProfileMpskExternalServerAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_key"
		if _, ok := i["mpsk-key"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKey(i["mpsk-key"], d, pre_append)
			tmp["mpsk_key"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfile-MpskGroup-MpskKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfile-MpskGroup-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_id"
		if _, ok := i["vlan-id"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupVlanId(i["vlan-id"], d, pre_append)
			tmp["vlan_id"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfile-MpskGroup-VlanId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_type"
		if _, ok := i["vlan-type"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupVlanType(i["vlan-type"], d, pre_append)
			tmp["vlan_type"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfile-MpskGroup-VlanType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKey(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_client_limit_type"
		if _, ok := i["concurrent-client-limit-type"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType(i["concurrent-client-limit-type"], d, pre_append)
			tmp["concurrent_client_limit_type"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-ConcurrentClientLimitType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := i["concurrent-clients"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients(i["concurrent-clients"], d, pre_append)
			tmp["concurrent_clients"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-ConcurrentClients")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_type"
		if _, ok := i["key-type"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyKeyType(i["key-type"], d, pre_append)
			tmp["key_type"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-KeyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := i["mpsk-schedules"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(i["mpsk-schedules"], d, pre_append)
			tmp["mpsk_schedules"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-MpskSchedules")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_pk"
		if _, ok := i["sae-pk"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePk(i["sae-pk"], d, pre_append)
			tmp["sae_pk"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-SaePk")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_private_key"
		if _, ok := i["sae-private-key"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey(i["sae-private-key"], d, pre_append)
			tmp["sae_private_key"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-SaePrivateKey")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyKeyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupVlanType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerMpskProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("mpsk_concurrent_clients", flattenWirelessControllerMpskProfileMpskConcurrentClients(o["mpsk-concurrent-clients"], d, "mpsk_concurrent_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-concurrent-clients"], "WirelessControllerMpskProfile-MpskConcurrentClients"); ok {
			if err = d.Set("mpsk_concurrent_clients", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
		}
	}

	if err = d.Set("mpsk_external_server", flattenWirelessControllerMpskProfileMpskExternalServer(o["mpsk-external-server"], d, "mpsk_external_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-external-server"], "WirelessControllerMpskProfile-MpskExternalServer"); ok {
			if err = d.Set("mpsk_external_server", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_external_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_external_server: %v", err)
		}
	}

	if err = d.Set("mpsk_external_server_auth", flattenWirelessControllerMpskProfileMpskExternalServerAuth(o["mpsk-external-server-auth"], d, "mpsk_external_server_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-external-server-auth"], "WirelessControllerMpskProfile-MpskExternalServerAuth"); ok {
			if err = d.Set("mpsk_external_server_auth", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_external_server_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_external_server_auth: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mpsk_group", flattenWirelessControllerMpskProfileMpskGroup(o["mpsk-group"], d, "mpsk_group")); err != nil {
			if vv, ok := fortiAPIPatch(o["mpsk-group"], "WirelessControllerMpskProfile-MpskGroup"); ok {
				if err = d.Set("mpsk_group", vv); err != nil {
					return fmt.Errorf("Error reading mpsk_group: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mpsk_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mpsk_group"); ok {
			if err = d.Set("mpsk_group", flattenWirelessControllerMpskProfileMpskGroup(o["mpsk-group"], d, "mpsk_group")); err != nil {
				if vv, ok := fortiAPIPatch(o["mpsk-group"], "WirelessControllerMpskProfile-MpskGroup"); ok {
					if err = d.Set("mpsk_group", vv); err != nil {
						return fmt.Errorf("Error reading mpsk_group: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mpsk_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("mpsk_type", flattenWirelessControllerMpskProfileMpskType(o["mpsk-type"], d, "mpsk_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-type"], "WirelessControllerMpskProfile-MpskType"); ok {
			if err = d.Set("mpsk_type", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_type: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerMpskProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerMpskProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ssid", flattenWirelessControllerMpskProfileSsid(o["ssid"], d, "ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssid"], "WirelessControllerMpskProfile-Ssid"); ok {
			if err = d.Set("ssid", vv); err != nil {
				return fmt.Errorf("Error reading ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssid: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerMpskProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerMpskProfileMpskConcurrentClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskExternalServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskExternalServerAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWirelessControllerMpskProfileMpskGroupMpskKey(d, i["mpsk_key"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["mpsk-key"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWirelessControllerMpskProfileMpskGroupName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-id"], _ = expandWirelessControllerMpskProfileMpskGroupVlanId(d, i["vlan_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-type"], _ = expandWirelessControllerMpskProfileMpskGroupVlanType(d, i["vlan_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_client_limit_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["concurrent-client-limit-type"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType(d, i["concurrent_client_limit_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["concurrent-clients"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients(d, i["concurrent_clients"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-type"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyKeyType(d, i["key_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk-schedules"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(d, i["mpsk_schedules"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passphrase"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase(d, i["passphrase"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmk"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pmk"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyPmk(d, i["pmk"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-password"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeySaePassword(d, i["sae_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_pk"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-pk"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeySaePk(d, i["sae_pk"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_private_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-private-key"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey(d, i["sae_private_key"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyKeyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyPmk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeySaePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeySaePk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupVlanType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerMpskProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mpsk_concurrent_clients"); ok || d.HasChange("mpsk_concurrent_clients") {
		t, err := expandWirelessControllerMpskProfileMpskConcurrentClients(d, v, "mpsk_concurrent_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-concurrent-clients"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_external_server"); ok || d.HasChange("mpsk_external_server") {
		t, err := expandWirelessControllerMpskProfileMpskExternalServer(d, v, "mpsk_external_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-external-server"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_external_server_auth"); ok || d.HasChange("mpsk_external_server_auth") {
		t, err := expandWirelessControllerMpskProfileMpskExternalServerAuth(d, v, "mpsk_external_server_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-external-server-auth"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_group"); ok || d.HasChange("mpsk_group") {
		t, err := expandWirelessControllerMpskProfileMpskGroup(d, v, "mpsk_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-group"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_type"); ok || d.HasChange("mpsk_type") {
		t, err := expandWirelessControllerMpskProfileMpskType(d, v, "mpsk_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerMpskProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ssid"); ok || d.HasChange("ssid") {
		t, err := expandWirelessControllerMpskProfileSsid(d, v, "ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid"] = t
		}
	}

	return &obj, nil
}
