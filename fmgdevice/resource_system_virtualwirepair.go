// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure virtual wire pairs.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVirtualWirePair() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVirtualWirePairCreate,
		Read:   resourceSystemVirtualWirePairRead,
		Update: resourceSystemVirtualWirePairUpdate,
		Delete: resourceSystemVirtualWirePairDelete,

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
			"member": &schema.Schema{
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
			"outer_vlan_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"poweroff_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"poweron_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wildcard_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemVirtualWirePairCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemVirtualWirePair(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVirtualWirePair resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemVirtualWirePair(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemVirtualWirePair(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemVirtualWirePair resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemVirtualWirePair(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemVirtualWirePair resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemVirtualWirePairRead(d, m)
}

func resourceSystemVirtualWirePairUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemVirtualWirePair(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVirtualWirePair resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemVirtualWirePair(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemVirtualWirePair resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemVirtualWirePairRead(d, m)
}

func resourceSystemVirtualWirePairDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteSystemVirtualWirePair(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVirtualWirePair resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVirtualWirePairRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVirtualWirePair(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemVirtualWirePair resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVirtualWirePair(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVirtualWirePair resource from API: %v", err)
	}
	return nil
}

func flattenSystemVirtualWirePairMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVirtualWirePairName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVirtualWirePairOuterVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemVirtualWirePairPoweroffBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVirtualWirePairPoweronBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVirtualWirePairVlanFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVirtualWirePairWildcardVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemVirtualWirePair(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("member", flattenSystemVirtualWirePairMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "SystemVirtualWirePair-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemVirtualWirePairName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemVirtualWirePair-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("outer_vlan_id", flattenSystemVirtualWirePairOuterVlanId(o["outer-vlan-id"], d, "outer_vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["outer-vlan-id"], "SystemVirtualWirePair-OuterVlanId"); ok {
			if err = d.Set("outer_vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading outer_vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outer_vlan_id: %v", err)
		}
	}

	if err = d.Set("poweroff_bypass", flattenSystemVirtualWirePairPoweroffBypass(o["poweroff-bypass"], d, "poweroff_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["poweroff-bypass"], "SystemVirtualWirePair-PoweroffBypass"); ok {
			if err = d.Set("poweroff_bypass", vv); err != nil {
				return fmt.Errorf("Error reading poweroff_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poweroff_bypass: %v", err)
		}
	}

	if err = d.Set("poweron_bypass", flattenSystemVirtualWirePairPoweronBypass(o["poweron-bypass"], d, "poweron_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["poweron-bypass"], "SystemVirtualWirePair-PoweronBypass"); ok {
			if err = d.Set("poweron_bypass", vv); err != nil {
				return fmt.Errorf("Error reading poweron_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poweron_bypass: %v", err)
		}
	}

	if err = d.Set("vlan_filter", flattenSystemVirtualWirePairVlanFilter(o["vlan-filter"], d, "vlan_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-filter"], "SystemVirtualWirePair-VlanFilter"); ok {
			if err = d.Set("vlan_filter", vv); err != nil {
				return fmt.Errorf("Error reading vlan_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_filter: %v", err)
		}
	}

	if err = d.Set("wildcard_vlan", flattenSystemVirtualWirePairWildcardVlan(o["wildcard-vlan"], d, "wildcard_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard-vlan"], "SystemVirtualWirePair-WildcardVlan"); ok {
			if err = d.Set("wildcard_vlan", vv); err != nil {
				return fmt.Errorf("Error reading wildcard_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard_vlan: %v", err)
		}
	}

	return nil
}

func flattenSystemVirtualWirePairFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVirtualWirePairMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVirtualWirePairName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWirePairOuterVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemVirtualWirePairPoweroffBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWirePairPoweronBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWirePairVlanFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWirePairWildcardVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVirtualWirePair(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandSystemVirtualWirePairMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemVirtualWirePairName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("outer_vlan_id"); ok || d.HasChange("outer_vlan_id") {
		t, err := expandSystemVirtualWirePairOuterVlanId(d, v, "outer_vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outer-vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("poweroff_bypass"); ok || d.HasChange("poweroff_bypass") {
		t, err := expandSystemVirtualWirePairPoweroffBypass(d, v, "poweroff_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poweroff-bypass"] = t
		}
	}

	if v, ok := d.GetOk("poweron_bypass"); ok || d.HasChange("poweron_bypass") {
		t, err := expandSystemVirtualWirePairPoweronBypass(d, v, "poweron_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poweron-bypass"] = t
		}
	}

	if v, ok := d.GetOk("vlan_filter"); ok || d.HasChange("vlan_filter") {
		t, err := expandSystemVirtualWirePairVlanFilter(d, v, "vlan_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-filter"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_vlan"); ok || d.HasChange("wildcard_vlan") {
		t, err := expandSystemVirtualWirePairWildcardVlan(d, v, "wildcard_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-vlan"] = t
		}
	}

	return &obj, nil
}
