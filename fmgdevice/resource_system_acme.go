// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ACME client.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAcme() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAcmeUpdate,
		Read:   resourceSystemAcmeRead,
		Update: resourceSystemAcmeUpdate,
		Delete: resourceSystemAcmeDelete,

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
			"accounts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"privatekey": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"store_passphrase": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"use_ha_direct": &schema.Schema{
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

func resourceSystemAcmeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAcme(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAcme resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAcme(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAcme resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemAcme")

	return resourceSystemAcmeRead(d, m)
}

func resourceSystemAcmeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemAcme(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAcme resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAcmeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAcme(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAcme resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAcme(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAcme resource from API: %v", err)
	}
	return nil
}

func flattenSystemAcmeAccounts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ca_url"
		if _, ok := i["ca_url"]; ok {
			v := flattenSystemAcmeAccountsCaUrl(i["ca_url"], d, pre_append)
			tmp["ca_url"] = fortiAPISubPartPatch(v, "SystemAcme-Accounts-CaUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := i["email"]; ok {
			v := flattenSystemAcmeAccountsEmail(i["email"], d, pre_append)
			tmp["email"] = fortiAPISubPartPatch(v, "SystemAcme-Accounts-Email")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemAcmeAccountsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemAcme-Accounts-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "privatekey"
		if _, ok := i["privatekey"]; ok {
			v := flattenSystemAcmeAccountsPrivatekey(i["privatekey"], d, pre_append)
			tmp["privatekey"] = fortiAPISubPartPatch(v, "SystemAcme-Accounts-Privatekey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSystemAcmeAccountsStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SystemAcme-Accounts-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			v := flattenSystemAcmeAccountsUrl(i["url"], d, pre_append)
			tmp["url"] = fortiAPISubPartPatch(v, "SystemAcme-Accounts-Url")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAcmeAccountsCaUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeAccountsEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeAccountsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeAccountsPrivatekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeAccountsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeAccountsUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAcmeSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeUseHaDirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAcme(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("accounts", flattenSystemAcmeAccounts(o["accounts"], d, "accounts")); err != nil {
			if vv, ok := fortiAPIPatch(o["accounts"], "SystemAcme-Accounts"); ok {
				if err = d.Set("accounts", vv); err != nil {
					return fmt.Errorf("Error reading accounts: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading accounts: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("accounts"); ok {
			if err = d.Set("accounts", flattenSystemAcmeAccounts(o["accounts"], d, "accounts")); err != nil {
				if vv, ok := fortiAPIPatch(o["accounts"], "SystemAcme-Accounts"); ok {
					if err = d.Set("accounts", vv); err != nil {
						return fmt.Errorf("Error reading accounts: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading accounts: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface", flattenSystemAcmeInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemAcme-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemAcmeSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystemAcme-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenSystemAcmeSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip6"], "SystemAcme-SourceIp6"); ok {
			if err = d.Set("source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("use_ha_direct", flattenSystemAcmeUseHaDirect(o["use-ha-direct"], d, "use_ha_direct")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-ha-direct"], "SystemAcme-UseHaDirect"); ok {
			if err = d.Set("use_ha_direct", vv); err != nil {
				return fmt.Errorf("Error reading use_ha_direct: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_ha_direct: %v", err)
		}
	}

	return nil
}

func flattenSystemAcmeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAcmeAccounts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ca_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ca_url"], _ = expandSystemAcmeAccountsCaUrl(d, i["ca_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["email"], _ = expandSystemAcmeAccountsEmail(d, i["email"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemAcmeAccountsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "privatekey"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["privatekey"], _ = expandSystemAcmeAccountsPrivatekey(d, i["privatekey"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSystemAcmeAccountsStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url"], _ = expandSystemAcmeAccountsUrl(d, i["url"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAcmeAccountsCaUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeAccountsEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeAccountsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeAccountsPrivatekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeAccountsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeAccountsUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAcmeSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeStorePassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAcmeUseHaDirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAcme(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accounts"); ok || d.HasChange("accounts") {
		t, err := expandSystemAcmeAccounts(d, v, "accounts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accounts"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemAcmeInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystemAcmeSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok || d.HasChange("source_ip6") {
		t, err := expandSystemAcmeSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("store_passphrase"); ok || d.HasChange("store_passphrase") {
		t, err := expandSystemAcmeStorePassphrase(d, v, "store_passphrase")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["store-passphrase"] = t
		}
	}

	if v, ok := d.GetOk("use_ha_direct"); ok || d.HasChange("use_ha_direct") {
		t, err := expandSystemAcmeUseHaDirect(d, v, "use_ha_direct")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-ha-direct"] = t
		}
	}

	return &obj, nil
}
