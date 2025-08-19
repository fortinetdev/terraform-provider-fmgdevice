// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: List of multiple PSK groups.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerMpskProfileMpskGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerMpskProfileMpskGroupCreate,
		Read:   resourceWirelessControllerMpskProfileMpskGroupRead,
		Update: resourceWirelessControllerMpskProfileMpskGroupUpdate,
		Delete: resourceWirelessControllerMpskProfileMpskGroupDelete,

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
			"mpsk_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
				ForceNew: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerMpskProfileMpskGroupCreate(d *schema.ResourceData, m interface{}) error {
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
	mpsk_profile := d.Get("mpsk_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mpsk_profile"] = mpsk_profile

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerMpskProfileMpskGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerMpskProfileMpskGroup resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerMpskProfileMpskGroup(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerMpskProfileMpskGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerMpskProfileMpskGroupRead(d, m)
}

func resourceWirelessControllerMpskProfileMpskGroupUpdate(d *schema.ResourceData, m interface{}) error {
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
	mpsk_profile := d.Get("mpsk_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mpsk_profile"] = mpsk_profile

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerMpskProfileMpskGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerMpskProfileMpskGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerMpskProfileMpskGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerMpskProfileMpskGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerMpskProfileMpskGroupRead(d, m)
}

func resourceWirelessControllerMpskProfileMpskGroupDelete(d *schema.ResourceData, m interface{}) error {
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
	mpsk_profile := d.Get("mpsk_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mpsk_profile"] = mpsk_profile

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerMpskProfileMpskGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerMpskProfileMpskGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerMpskProfileMpskGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	mpsk_profile := d.Get("mpsk_profile").(string)
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
	if mpsk_profile == "" {
		mpsk_profile = importOptionChecking(m.(*FortiClient).Cfg, "mpsk_profile")
		if mpsk_profile == "" {
			return fmt.Errorf("Parameter mpsk_profile is missing")
		}
		if err = d.Set("mpsk_profile", mpsk_profile); err != nil {
			return fmt.Errorf("Error set params mpsk_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mpsk_profile"] = mpsk_profile

	o, err := c.ReadWirelessControllerMpskProfileMpskGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerMpskProfileMpskGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerMpskProfileMpskGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerMpskProfileMpskGroup resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKey2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyComment2edl(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_client_limit_type"
		if _, ok := i["concurrent-client-limit-type"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType2edl(i["concurrent-client-limit-type"], d, pre_append)
			tmp["concurrent_client_limit_type"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-ConcurrentClientLimitType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := i["concurrent-clients"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients2edl(i["concurrent-clients"], d, pre_append)
			tmp["concurrent_clients"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-ConcurrentClients")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_type"
		if _, ok := i["key-type"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyKeyType2edl(i["key-type"], d, pre_append)
			tmp["key_type"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-KeyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyMac2edl(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := i["mpsk-schedules"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules2edl(i["mpsk-schedules"], d, pre_append)
			tmp["mpsk_schedules"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-MpskSchedules")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeyName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_pk"
		if _, ok := i["sae-pk"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePk2edl(i["sae-pk"], d, pre_append)
			tmp["sae_pk"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-SaePk")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_private_key"
		if _, ok := i["sae-private-key"]; ok {
			v := flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey2edl(i["sae-private-key"], d, pre_append)
			tmp["sae_private_key"] = fortiAPISubPartPatch(v, "WirelessControllerMpskProfileMpskGroup-MpskKey-SaePrivateKey")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyKeyType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePk2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupVlanId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupVlanType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerMpskProfileMpskGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("mpsk_key", flattenWirelessControllerMpskProfileMpskGroupMpskKey2edl(o["mpsk-key"], d, "mpsk_key")); err != nil {
			if vv, ok := fortiAPIPatch(o["mpsk-key"], "WirelessControllerMpskProfileMpskGroup-MpskKey"); ok {
				if err = d.Set("mpsk_key", vv); err != nil {
					return fmt.Errorf("Error reading mpsk_key: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mpsk_key: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mpsk_key"); ok {
			if err = d.Set("mpsk_key", flattenWirelessControllerMpskProfileMpskGroupMpskKey2edl(o["mpsk-key"], d, "mpsk_key")); err != nil {
				if vv, ok := fortiAPIPatch(o["mpsk-key"], "WirelessControllerMpskProfileMpskGroup-MpskKey"); ok {
					if err = d.Set("mpsk_key", vv); err != nil {
						return fmt.Errorf("Error reading mpsk_key: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mpsk_key: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenWirelessControllerMpskProfileMpskGroupName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerMpskProfileMpskGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vlan_id", flattenWirelessControllerMpskProfileMpskGroupVlanId2edl(o["vlan-id"], d, "vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-id"], "WirelessControllerMpskProfileMpskGroup-VlanId"); ok {
			if err = d.Set("vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_id: %v", err)
		}
	}

	if err = d.Set("vlan_type", flattenWirelessControllerMpskProfileMpskGroupVlanType2edl(o["vlan-type"], d, "vlan_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-type"], "WirelessControllerMpskProfileMpskGroup-VlanType"); ok {
			if err = d.Set("vlan_type", vv); err != nil {
				return fmt.Errorf("Error reading vlan_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_type: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerMpskProfileMpskGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerMpskProfileMpskGroupMpskKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["comment"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyComment2edl(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_client_limit_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["concurrent-client-limit-type"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType2edl(d, i["concurrent_client_limit_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["concurrent-clients"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients2edl(d, i["concurrent_clients"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-type"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyKeyType2edl(d, i["key_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyMac2edl(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mpsk-schedules"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules2edl(d, i["mpsk_schedules"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passphrase"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase2edl(d, i["passphrase"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pmk"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pmk"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeyPmk2edl(d, i["pmk"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-password"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeySaePassword2edl(d, i["sae_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_pk"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-pk"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeySaePk2edl(d, i["sae_pk"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sae_private_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sae-private-key"], _ = expandWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey2edl(d, i["sae_private_key"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyKeyType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyPmk2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeySaePassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeySaePk2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupVlanId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupVlanType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerMpskProfileMpskGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mpsk_key"); ok || d.HasChange("mpsk_key") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKey2edl(d, v, "mpsk_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-key"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerMpskProfileMpskGroupName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vlan_id"); ok || d.HasChange("vlan_id") {
		t, err := expandWirelessControllerMpskProfileMpskGroupVlanId2edl(d, v, "vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("vlan_type"); ok || d.HasChange("vlan_type") {
		t, err := expandWirelessControllerMpskProfileMpskGroupVlanType2edl(d, v, "vlan_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-type"] = t
		}
	}

	return &obj, nil
}
