// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device WirelessControllerVapDynamicMappingIp6PrefixList

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerVapDynamicMappingIp6PrefixList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerVapDynamicMappingIp6PrefixListCreate,
		Read:   resourceWirelessControllerVapDynamicMappingIp6PrefixListRead,
		Update: resourceWirelessControllerVapDynamicMappingIp6PrefixListUpdate,
		Delete: resourceWirelessControllerVapDynamicMappingIp6PrefixListDelete,

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
			"vap": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dynamic_mapping_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dynamic_mapping_vdom": &schema.Schema{
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

func resourceWirelessControllerVapDynamicMappingIp6PrefixListCreate(d *schema.ResourceData, m interface{}) error {
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
	vap := d.Get("vap").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["device"] = device_name
	paradict["vap"] = vap
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	obj, err := getObjectWirelessControllerVapDynamicMappingIp6PrefixList(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerVapDynamicMappingIp6PrefixList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWirelessControllerVapDynamicMappingIp6PrefixList(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWirelessControllerVapDynamicMappingIp6PrefixList(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WirelessControllerVapDynamicMappingIp6PrefixList resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWirelessControllerVapDynamicMappingIp6PrefixList(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WirelessControllerVapDynamicMappingIp6PrefixList resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, ""))

	return resourceWirelessControllerVapDynamicMappingIp6PrefixListRead(d, m)
}

func resourceWirelessControllerVapDynamicMappingIp6PrefixListUpdate(d *schema.ResourceData, m interface{}) error {
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
	vap := d.Get("vap").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["device"] = device_name
	paradict["vap"] = vap
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	obj, err := getObjectWirelessControllerVapDynamicMappingIp6PrefixList(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVapDynamicMappingIp6PrefixList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWirelessControllerVapDynamicMappingIp6PrefixList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVapDynamicMappingIp6PrefixList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceWirelessControllerVapDynamicMappingIp6PrefixListRead(d, m)
}

func resourceWirelessControllerVapDynamicMappingIp6PrefixListDelete(d *schema.ResourceData, m interface{}) error {
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
	vap := d.Get("vap").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["device"] = device_name
	paradict["vap"] = vap
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	wsParams["adom"] = adomv

	err = c.DeleteWirelessControllerVapDynamicMappingIp6PrefixList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerVapDynamicMappingIp6PrefixList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerVapDynamicMappingIp6PrefixListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	vap := d.Get("vap").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
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
	if dynamic_mapping_name == "" {
		dynamic_mapping_name = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_mapping_name")
		if dynamic_mapping_name == "" {
			return fmt.Errorf("Parameter dynamic_mapping_name is missing")
		}
		if err = d.Set("dynamic_mapping_name", dynamic_mapping_name); err != nil {
			return fmt.Errorf("Error set params dynamic_mapping_name: %v", err)
		}
	}
	if dynamic_mapping_vdom == "" {
		dynamic_mapping_vdom = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_mapping_vdom")
		if dynamic_mapping_vdom == "" {
			return fmt.Errorf("Parameter dynamic_mapping_vdom is missing")
		}
		if err = d.Set("dynamic_mapping_vdom", dynamic_mapping_vdom); err != nil {
			return fmt.Errorf("Error set params dynamic_mapping_vdom: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vap"] = vap
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	o, err := c.ReadWirelessControllerVapDynamicMappingIp6PrefixList(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WirelessControllerVapDynamicMappingIp6PrefixList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerVapDynamicMappingIp6PrefixList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVapDynamicMappingIp6PrefixList resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerVapDynamicMappingIp6PrefixListAutonomousFlag3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIp6PrefixListDnssl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingIp6PrefixListOnlinkFlag3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIp6PrefixListPreferredLifeTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIp6PrefixListPrefix3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicMappingIp6PrefixListRdnss3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerVapDynamicMappingIp6PrefixListValidLifeTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerVapDynamicMappingIp6PrefixList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("autonomous_flag", flattenWirelessControllerVapDynamicMappingIp6PrefixListAutonomousFlag3rdl(o["autonomous-flag"], d, "autonomous_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["autonomous-flag"], "WirelessControllerVapDynamicMappingIp6PrefixList-AutonomousFlag"); ok {
			if err = d.Set("autonomous_flag", vv); err != nil {
				return fmt.Errorf("Error reading autonomous_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading autonomous_flag: %v", err)
		}
	}

	if err = d.Set("dnssl", flattenWirelessControllerVapDynamicMappingIp6PrefixListDnssl3rdl(o["dnssl"], d, "dnssl")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnssl"], "WirelessControllerVapDynamicMappingIp6PrefixList-Dnssl"); ok {
			if err = d.Set("dnssl", vv); err != nil {
				return fmt.Errorf("Error reading dnssl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnssl: %v", err)
		}
	}

	if err = d.Set("onlink_flag", flattenWirelessControllerVapDynamicMappingIp6PrefixListOnlinkFlag3rdl(o["onlink-flag"], d, "onlink_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["onlink-flag"], "WirelessControllerVapDynamicMappingIp6PrefixList-OnlinkFlag"); ok {
			if err = d.Set("onlink_flag", vv); err != nil {
				return fmt.Errorf("Error reading onlink_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading onlink_flag: %v", err)
		}
	}

	if err = d.Set("preferred_life_time", flattenWirelessControllerVapDynamicMappingIp6PrefixListPreferredLifeTime3rdl(o["preferred-life-time"], d, "preferred_life_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-life-time"], "WirelessControllerVapDynamicMappingIp6PrefixList-PreferredLifeTime"); ok {
			if err = d.Set("preferred_life_time", vv); err != nil {
				return fmt.Errorf("Error reading preferred_life_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_life_time: %v", err)
		}
	}

	if err = d.Set("prefix", flattenWirelessControllerVapDynamicMappingIp6PrefixListPrefix3rdl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "WirelessControllerVapDynamicMappingIp6PrefixList-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("rdnss", flattenWirelessControllerVapDynamicMappingIp6PrefixListRdnss3rdl(o["rdnss"], d, "rdnss")); err != nil {
		if vv, ok := fortiAPIPatch(o["rdnss"], "WirelessControllerVapDynamicMappingIp6PrefixList-Rdnss"); ok {
			if err = d.Set("rdnss", vv); err != nil {
				return fmt.Errorf("Error reading rdnss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rdnss: %v", err)
		}
	}

	if err = d.Set("valid_life_time", flattenWirelessControllerVapDynamicMappingIp6PrefixListValidLifeTime3rdl(o["valid-life-time"], d, "valid_life_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["valid-life-time"], "WirelessControllerVapDynamicMappingIp6PrefixList-ValidLifeTime"); ok {
			if err = d.Set("valid_life_time", vv); err != nil {
				return fmt.Errorf("Error reading valid_life_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading valid_life_time: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerVapDynamicMappingIp6PrefixListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerVapDynamicMappingIp6PrefixListAutonomousFlag3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIp6PrefixListDnssl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingIp6PrefixListOnlinkFlag3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIp6PrefixListPreferredLifeTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIp6PrefixListPrefix3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicMappingIp6PrefixListRdnss3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerVapDynamicMappingIp6PrefixListValidLifeTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerVapDynamicMappingIp6PrefixList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("autonomous_flag"); ok || d.HasChange("autonomous_flag") {
		t, err := expandWirelessControllerVapDynamicMappingIp6PrefixListAutonomousFlag3rdl(d, v, "autonomous_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["autonomous-flag"] = t
		}
	}

	if v, ok := d.GetOk("dnssl"); ok || d.HasChange("dnssl") {
		t, err := expandWirelessControllerVapDynamicMappingIp6PrefixListDnssl3rdl(d, v, "dnssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnssl"] = t
		}
	}

	if v, ok := d.GetOk("onlink_flag"); ok || d.HasChange("onlink_flag") {
		t, err := expandWirelessControllerVapDynamicMappingIp6PrefixListOnlinkFlag3rdl(d, v, "onlink_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["onlink-flag"] = t
		}
	}

	if v, ok := d.GetOk("preferred_life_time"); ok || d.HasChange("preferred_life_time") {
		t, err := expandWirelessControllerVapDynamicMappingIp6PrefixListPreferredLifeTime3rdl(d, v, "preferred_life_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-life-time"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandWirelessControllerVapDynamicMappingIp6PrefixListPrefix3rdl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("rdnss"); ok || d.HasChange("rdnss") {
		t, err := expandWirelessControllerVapDynamicMappingIp6PrefixListRdnss3rdl(d, v, "rdnss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rdnss"] = t
		}
	}

	if v, ok := d.GetOk("valid_life_time"); ok || d.HasChange("valid_life_time") {
		t, err := expandWirelessControllerVapDynamicMappingIp6PrefixListValidLifeTime3rdl(d, v, "valid_life_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["valid-life-time"] = t
		}
	}

	return &obj, nil
}
