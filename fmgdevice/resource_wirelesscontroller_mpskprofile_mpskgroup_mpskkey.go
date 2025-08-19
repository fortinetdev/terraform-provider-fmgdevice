// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: List of multiple PSK entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerMpskProfileMpskGroupMpskKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerMpskProfileMpskGroupMpskKeyCreate,
		Read:   resourceWirelessControllerMpskProfileMpskGroupMpskKeyRead,
		Update: resourceWirelessControllerMpskProfileMpskGroupMpskKeyUpdate,
		Delete: resourceWirelessControllerMpskProfileMpskGroupMpskKeyDelete,

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
			"mpsk_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
				ForceNew: true,
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
	}
}

func resourceWirelessControllerMpskProfileMpskGroupMpskKeyCreate(d *schema.ResourceData, m interface{}) error {
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
	mpsk_group := d.Get("mpsk_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mpsk_profile"] = mpsk_profile
	paradict["mpsk_group"] = mpsk_group

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerMpskProfileMpskGroupMpskKey(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerMpskProfileMpskGroupMpskKey resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerMpskProfileMpskGroupMpskKey(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerMpskProfileMpskGroupMpskKey resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerMpskProfileMpskGroupMpskKeyRead(d, m)
}

func resourceWirelessControllerMpskProfileMpskGroupMpskKeyUpdate(d *schema.ResourceData, m interface{}) error {
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
	mpsk_group := d.Get("mpsk_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mpsk_profile"] = mpsk_profile
	paradict["mpsk_group"] = mpsk_group

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerMpskProfileMpskGroupMpskKey(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerMpskProfileMpskGroupMpskKey resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerMpskProfileMpskGroupMpskKey(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerMpskProfileMpskGroupMpskKey resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerMpskProfileMpskGroupMpskKeyRead(d, m)
}

func resourceWirelessControllerMpskProfileMpskGroupMpskKeyDelete(d *schema.ResourceData, m interface{}) error {
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
	mpsk_group := d.Get("mpsk_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mpsk_profile"] = mpsk_profile
	paradict["mpsk_group"] = mpsk_group

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerMpskProfileMpskGroupMpskKey(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerMpskProfileMpskGroupMpskKey resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerMpskProfileMpskGroupMpskKeyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	mpsk_profile := d.Get("mpsk_profile").(string)
	mpsk_group := d.Get("mpsk_group").(string)
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
	if mpsk_group == "" {
		mpsk_group = importOptionChecking(m.(*FortiClient).Cfg, "mpsk_group")
		if mpsk_group == "" {
			return fmt.Errorf("Parameter mpsk_group is missing")
		}
		if err = d.Set("mpsk_group", mpsk_group); err != nil {
			return fmt.Errorf("Error set params mpsk_group: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mpsk_profile"] = mpsk_profile
	paradict["mpsk_group"] = mpsk_group

	o, err := c.ReadWirelessControllerMpskProfileMpskGroupMpskKey(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerMpskProfileMpskGroupMpskKey resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerMpskProfileMpskGroupMpskKey(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerMpskProfileMpskGroupMpskKey resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyComment3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyKeyType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyMac3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePk3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerMpskProfileMpskGroupMpskKey(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenWirelessControllerMpskProfileMpskGroupMpskKeyComment3rdl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerMpskProfileMpskGroupMpskKey-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("concurrent_client_limit_type", flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType3rdl(o["concurrent-client-limit-type"], d, "concurrent_client_limit_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["concurrent-client-limit-type"], "WirelessControllerMpskProfileMpskGroupMpskKey-ConcurrentClientLimitType"); ok {
			if err = d.Set("concurrent_client_limit_type", vv); err != nil {
				return fmt.Errorf("Error reading concurrent_client_limit_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading concurrent_client_limit_type: %v", err)
		}
	}

	if err = d.Set("concurrent_clients", flattenWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients3rdl(o["concurrent-clients"], d, "concurrent_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["concurrent-clients"], "WirelessControllerMpskProfileMpskGroupMpskKey-ConcurrentClients"); ok {
			if err = d.Set("concurrent_clients", vv); err != nil {
				return fmt.Errorf("Error reading concurrent_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading concurrent_clients: %v", err)
		}
	}

	if err = d.Set("key_type", flattenWirelessControllerMpskProfileMpskGroupMpskKeyKeyType3rdl(o["key-type"], d, "key_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-type"], "WirelessControllerMpskProfileMpskGroupMpskKey-KeyType"); ok {
			if err = d.Set("key_type", vv); err != nil {
				return fmt.Errorf("Error reading key_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_type: %v", err)
		}
	}

	if err = d.Set("mac", flattenWirelessControllerMpskProfileMpskGroupMpskKeyMac3rdl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "WirelessControllerMpskProfileMpskGroupMpskKey-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("mpsk_schedules", flattenWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules3rdl(o["mpsk-schedules"], d, "mpsk_schedules")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-schedules"], "WirelessControllerMpskProfileMpskGroupMpskKey-MpskSchedules"); ok {
			if err = d.Set("mpsk_schedules", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_schedules: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_schedules: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerMpskProfileMpskGroupMpskKeyName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerMpskProfileMpskGroupMpskKey-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("sae_pk", flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePk3rdl(o["sae-pk"], d, "sae_pk")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-pk"], "WirelessControllerMpskProfileMpskGroupMpskKey-SaePk"); ok {
			if err = d.Set("sae_pk", vv); err != nil {
				return fmt.Errorf("Error reading sae_pk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_pk: %v", err)
		}
	}

	if err = d.Set("sae_private_key", flattenWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey3rdl(o["sae-private-key"], d, "sae_private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["sae-private-key"], "WirelessControllerMpskProfileMpskGroupMpskKey-SaePrivateKey"); ok {
			if err = d.Set("sae_private_key", vv); err != nil {
				return fmt.Errorf("Error reading sae_private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sae_private_key: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerMpskProfileMpskGroupMpskKeyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyComment3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyKeyType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyMac3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeyPmk3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeySaePassword3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeySaePk3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerMpskProfileMpskGroupMpskKey(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKeyComment3rdl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("concurrent_client_limit_type"); ok || d.HasChange("concurrent_client_limit_type") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType3rdl(d, v, "concurrent_client_limit_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["concurrent-client-limit-type"] = t
		}
	}

	if v, ok := d.GetOk("concurrent_clients"); ok || d.HasChange("concurrent_clients") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients3rdl(d, v, "concurrent_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["concurrent-clients"] = t
		}
	}

	if v, ok := d.GetOk("key_type"); ok || d.HasChange("key_type") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKeyKeyType3rdl(d, v, "key_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-type"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKeyMac3rdl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_schedules"); ok || d.HasChange("mpsk_schedules") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules3rdl(d, v, "mpsk_schedules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-schedules"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKeyName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("passphrase"); ok || d.HasChange("passphrase") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase3rdl(d, v, "passphrase")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passphrase"] = t
		}
	}

	if v, ok := d.GetOk("pmk"); ok || d.HasChange("pmk") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKeyPmk3rdl(d, v, "pmk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmk"] = t
		}
	}

	if v, ok := d.GetOk("sae_password"); ok || d.HasChange("sae_password") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKeySaePassword3rdl(d, v, "sae_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-password"] = t
		}
	}

	if v, ok := d.GetOk("sae_pk"); ok || d.HasChange("sae_pk") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKeySaePk3rdl(d, v, "sae_pk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-pk"] = t
		}
	}

	if v, ok := d.GetOk("sae_private_key"); ok || d.HasChange("sae_private_key") {
		t, err := expandWirelessControllerMpskProfileMpskGroupMpskKeySaePrivateKey3rdl(d, v, "sae_private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-private-key"] = t
		}
	}

	return &obj, nil
}
