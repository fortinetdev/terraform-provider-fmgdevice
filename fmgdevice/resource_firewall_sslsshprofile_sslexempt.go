// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Servers to exempt from SSL inspection.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSslSshProfileSslExempt() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSslSshProfileSslExemptCreate,
		Read:   resourceFirewallSslSshProfileSslExemptRead,
		Update: resourceFirewallSslSshProfileSslExemptUpdate,
		Delete: resourceFirewallSslSshProfileSslExemptDelete,

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
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"address6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fortiguard_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard_fqdn": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"finger_print_category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallSslSshProfileSslExemptCreate(d *schema.ResourceData, m interface{}) error {
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
	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	obj, err := getObjectFirewallSslSshProfileSslExempt(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSslSshProfileSslExempt resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallSslSshProfileSslExempt(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallSslSshProfileSslExempt(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallSslSshProfileSslExempt resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallSslSshProfileSslExempt(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallSslSshProfileSslExempt resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallSslSshProfileSslExemptRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallSslSshProfileSslExempt resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallSslSshProfileSslExemptRead(d, m)
}

func resourceFirewallSslSshProfileSslExemptUpdate(d *schema.ResourceData, m interface{}) error {
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
	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	obj, err := getObjectFirewallSslSshProfileSslExempt(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileSslExempt resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallSslSshProfileSslExempt(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileSslExempt resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallSslSshProfileSslExemptRead(d, m)
}

func resourceFirewallSslSshProfileSslExemptDelete(d *schema.ResourceData, m interface{}) error {
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
	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	wsParams["adom"] = adomv

	err = c.DeleteFirewallSslSshProfileSslExempt(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSslSshProfileSslExempt resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslSshProfileSslExemptRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
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
	if ssl_ssh_profile == "" {
		ssl_ssh_profile = importOptionChecking(m.(*FortiClient).Cfg, "ssl_ssh_profile")
		if ssl_ssh_profile == "" {
			return fmt.Errorf("Parameter ssl_ssh_profile is missing")
		}
		if err = d.Set("ssl_ssh_profile", ssl_ssh_profile); err != nil {
			return fmt.Errorf("Error set params ssl_ssh_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	o, err := c.ReadFirewallSslSshProfileSslExempt(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallSslSshProfileSslExempt resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslSshProfileSslExempt(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSshProfileSslExempt resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslSshProfileSslExemptAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslExemptAddress62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslExemptFortiguardCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslExemptId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptRegex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptWildcardFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSslSshProfileSslExemptFingerPrintCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSslSshProfileSslExempt(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenFirewallSslSshProfileSslExemptAddress2edl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "FirewallSslSshProfileSslExempt-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("address6", flattenFirewallSslSshProfileSslExemptAddress62edl(o["address6"], d, "address6")); err != nil {
		if vv, ok := fortiAPIPatch(o["address6"], "FirewallSslSshProfileSslExempt-Address6"); ok {
			if err = d.Set("address6", vv); err != nil {
				return fmt.Errorf("Error reading address6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address6: %v", err)
		}
	}

	if err = d.Set("fortiguard_category", flattenFirewallSslSshProfileSslExemptFortiguardCategory2edl(o["fortiguard-category"], d, "fortiguard_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-category"], "FirewallSslSshProfileSslExempt-FortiguardCategory"); ok {
			if err = d.Set("fortiguard_category", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_category: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallSslSshProfileSslExemptId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallSslSshProfileSslExempt-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("regex", flattenFirewallSslSshProfileSslExemptRegex2edl(o["regex"], d, "regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["regex"], "FirewallSslSshProfileSslExempt-Regex"); ok {
			if err = d.Set("regex", vv); err != nil {
				return fmt.Errorf("Error reading regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regex: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallSslSshProfileSslExemptType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallSslSshProfileSslExempt-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("wildcard_fqdn", flattenFirewallSslSshProfileSslExemptWildcardFqdn2edl(o["wildcard-fqdn"], d, "wildcard_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard-fqdn"], "FirewallSslSshProfileSslExempt-WildcardFqdn"); ok {
			if err = d.Set("wildcard_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
		}
	}

	if err = d.Set("finger_print_category", flattenFirewallSslSshProfileSslExemptFingerPrintCategory2edl(o["finger-print-category"], d, "finger_print_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["finger-print-category"], "FirewallSslSshProfileSslExempt-FingerPrintCategory"); ok {
			if err = d.Set("finger_print_category", vv); err != nil {
				return fmt.Errorf("Error reading finger_print_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading finger_print_category: %v", err)
		}
	}

	return nil
}

func flattenFirewallSslSshProfileSslExemptFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSslSshProfileSslExemptAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslExemptAddress62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslExemptFortiguardCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslExemptId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptRegex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptWildcardFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSslExemptFingerPrintCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSslSshProfileSslExempt(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandFirewallSslSshProfileSslExemptAddress2edl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("address6"); ok || d.HasChange("address6") {
		t, err := expandFirewallSslSshProfileSslExemptAddress62edl(d, v, "address6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address6"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_category"); ok || d.HasChange("fortiguard_category") {
		t, err := expandFirewallSslSshProfileSslExemptFortiguardCategory2edl(d, v, "fortiguard_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-category"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallSslSshProfileSslExemptId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("regex"); ok || d.HasChange("regex") {
		t, err := expandFirewallSslSshProfileSslExemptRegex2edl(d, v, "regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regex"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallSslSshProfileSslExemptType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_fqdn"); ok || d.HasChange("wildcard_fqdn") {
		t, err := expandFirewallSslSshProfileSslExemptWildcardFqdn2edl(d, v, "wildcard_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("finger_print_category"); ok || d.HasChange("finger_print_category") {
		t, err := expandFirewallSslSshProfileSslExemptFingerPrintCategory2edl(d, v, "finger_print_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["finger-print-category"] = t
		}
	}

	return &obj, nil
}
