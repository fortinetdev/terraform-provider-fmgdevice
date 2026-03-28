// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device WirelessControllerVapIp6PrefixList

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerVapIp6PrefixList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerVapIp6PrefixListCreate,
		Read:   resourceWirelessControllerVapIp6PrefixListRead,
		Update: resourceWirelessControllerVapIp6PrefixListUpdate,
		Delete: resourceWirelessControllerVapIp6PrefixListDelete,

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
			"vap": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"autonomous_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dnssl": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"onlink_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"preferred_life_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rdnss": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"valid_life_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerVapIp6PrefixListCreate(d *schema.ResourceData, m interface{}) error {
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
	vap := d.Get("vap").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vap"] = vap

	obj, err := getObjectWirelessControllerVapIp6PrefixList(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerVapIp6PrefixList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWirelessControllerVapIp6PrefixList(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWirelessControllerVapIp6PrefixList(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WirelessControllerVapIp6PrefixList resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWirelessControllerVapIp6PrefixList(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WirelessControllerVapIp6PrefixList resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, ""))

	return resourceWirelessControllerVapIp6PrefixListRead(d, m)
}

func resourceWirelessControllerVapIp6PrefixListUpdate(d *schema.ResourceData, m interface{}) error {
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
	vap := d.Get("vap").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vap"] = vap

	obj, err := getObjectWirelessControllerVapIp6PrefixList(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVapIp6PrefixList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWirelessControllerVapIp6PrefixList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVapIp6PrefixList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceWirelessControllerVapIp6PrefixListRead(d, m)
}

func resourceWirelessControllerVapIp6PrefixListDelete(d *schema.ResourceData, m interface{}) error {
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
	vap := d.Get("vap").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vap"] = vap

	wsParams["adom"] = adomv

	err = c.DeleteWirelessControllerVapIp6PrefixList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerVapIp6PrefixList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerVapIp6PrefixListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	vap := d.Get("vap").(string)
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
	if vap == "" {
		vap = importOptionChecking(m.(*FortiClient).Cfg, "vap")
		if vap == "" {
			return fmt.Errorf("Parameter vap is missing")
		}
		if err = d.Set("vap", vap); err != nil {
			return fmt.Errorf("Error set params vap: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vap"] = vap

	o, err := c.ReadWirelessControllerVapIp6PrefixList(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WirelessControllerVapIp6PrefixList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerVapIp6PrefixList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVapIp6PrefixList resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerVapIp6PrefixListAutonomousFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIp6PrefixListDnssl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapIp6PrefixListOnlinkFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIp6PrefixListPreferredLifeTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIp6PrefixListPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIp6PrefixListRdnss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapIp6PrefixListValidLifeTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerVapIp6PrefixList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("autonomous_flag", flattenWirelessControllerVapIp6PrefixListAutonomousFlag2edl(o["autonomous-flag"], d, "autonomous_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["autonomous-flag"], "WirelessControllerVapIp6PrefixList-AutonomousFlag"); ok {
			if err = d.Set("autonomous_flag", vv); err != nil {
				return fmt.Errorf("Error reading autonomous_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading autonomous_flag: %v", err)
		}
	}

	if err = d.Set("dnssl", flattenWirelessControllerVapIp6PrefixListDnssl2edl(o["dnssl"], d, "dnssl")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnssl"], "WirelessControllerVapIp6PrefixList-Dnssl"); ok {
			if err = d.Set("dnssl", vv); err != nil {
				return fmt.Errorf("Error reading dnssl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnssl: %v", err)
		}
	}

	if err = d.Set("onlink_flag", flattenWirelessControllerVapIp6PrefixListOnlinkFlag2edl(o["onlink-flag"], d, "onlink_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["onlink-flag"], "WirelessControllerVapIp6PrefixList-OnlinkFlag"); ok {
			if err = d.Set("onlink_flag", vv); err != nil {
				return fmt.Errorf("Error reading onlink_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading onlink_flag: %v", err)
		}
	}

	if err = d.Set("preferred_life_time", flattenWirelessControllerVapIp6PrefixListPreferredLifeTime2edl(o["preferred-life-time"], d, "preferred_life_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-life-time"], "WirelessControllerVapIp6PrefixList-PreferredLifeTime"); ok {
			if err = d.Set("preferred_life_time", vv); err != nil {
				return fmt.Errorf("Error reading preferred_life_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_life_time: %v", err)
		}
	}

	if err = d.Set("prefix", flattenWirelessControllerVapIp6PrefixListPrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "WirelessControllerVapIp6PrefixList-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("rdnss", flattenWirelessControllerVapIp6PrefixListRdnss2edl(o["rdnss"], d, "rdnss")); err != nil {
		if vv, ok := fortiAPIPatch(o["rdnss"], "WirelessControllerVapIp6PrefixList-Rdnss"); ok {
			if err = d.Set("rdnss", vv); err != nil {
				return fmt.Errorf("Error reading rdnss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rdnss: %v", err)
		}
	}

	if err = d.Set("valid_life_time", flattenWirelessControllerVapIp6PrefixListValidLifeTime2edl(o["valid-life-time"], d, "valid_life_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["valid-life-time"], "WirelessControllerVapIp6PrefixList-ValidLifeTime"); ok {
			if err = d.Set("valid_life_time", vv); err != nil {
				return fmt.Errorf("Error reading valid_life_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading valid_life_time: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerVapIp6PrefixListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerVapIp6PrefixListAutonomousFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIp6PrefixListDnssl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapIp6PrefixListOnlinkFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIp6PrefixListPreferredLifeTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIp6PrefixListPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIp6PrefixListRdnss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapIp6PrefixListValidLifeTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerVapIp6PrefixList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("autonomous_flag"); ok || d.HasChange("autonomous_flag") {
		t, err := expandWirelessControllerVapIp6PrefixListAutonomousFlag2edl(d, v, "autonomous_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["autonomous-flag"] = t
		}
	}

	if v, ok := d.GetOk("dnssl"); ok || d.HasChange("dnssl") {
		t, err := expandWirelessControllerVapIp6PrefixListDnssl2edl(d, v, "dnssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnssl"] = t
		}
	}

	if v, ok := d.GetOk("onlink_flag"); ok || d.HasChange("onlink_flag") {
		t, err := expandWirelessControllerVapIp6PrefixListOnlinkFlag2edl(d, v, "onlink_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["onlink-flag"] = t
		}
	}

	if v, ok := d.GetOk("preferred_life_time"); ok || d.HasChange("preferred_life_time") {
		t, err := expandWirelessControllerVapIp6PrefixListPreferredLifeTime2edl(d, v, "preferred_life_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-life-time"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandWirelessControllerVapIp6PrefixListPrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("rdnss"); ok || d.HasChange("rdnss") {
		t, err := expandWirelessControllerVapIp6PrefixListRdnss2edl(d, v, "rdnss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rdnss"] = t
		}
	}

	if v, ok := d.GetOk("valid_life_time"); ok || d.HasChange("valid_life_time") {
		t, err := expandWirelessControllerVapIp6PrefixListValidLifeTime2edl(d, v, "valid_life_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["valid-life-time"] = t
		}
	}

	return &obj, nil
}
