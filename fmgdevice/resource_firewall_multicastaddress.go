// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure multicast addresses.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallMulticastAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallMulticastAddressCreate,
		Read:   resourceFirewallMulticastAddressRead,
		Update: resourceFirewallMulticastAddressUpdate,
		Delete: resourceFirewallMulticastAddressDelete,

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
			"associated_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"tagging": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"type": &schema.Schema{
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

func resourceFirewallMulticastAddressCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallMulticastAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallMulticastAddress resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallMulticastAddress(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallMulticastAddress(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallMulticastAddress resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallMulticastAddress(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallMulticastAddress resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallMulticastAddressRead(d, m)
}

func resourceFirewallMulticastAddressUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallMulticastAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallMulticastAddress resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallMulticastAddress(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallMulticastAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallMulticastAddressRead(d, m)
}

func resourceFirewallMulticastAddressDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallMulticastAddress(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallMulticastAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallMulticastAddressRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallMulticastAddress(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallMulticastAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallMulticastAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallMulticastAddress resource from API: %v", err)
	}
	return nil
}

func flattenFirewallMulticastAddressAssociatedInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallMulticastAddressColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastAddressComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastAddressEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastAddressStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastAddressSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallMulticastAddressTagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenFirewallMulticastAddressTaggingCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "FirewallMulticastAddress-Tagging-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallMulticastAddressTaggingName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallMulticastAddress-Tagging-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenFirewallMulticastAddressTaggingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "FirewallMulticastAddress-Tagging-Tags")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallMulticastAddressTaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallMulticastAddressTaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastAddressTaggingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallMulticastAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallMulticastAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("associated_interface", flattenFirewallMulticastAddressAssociatedInterface(o["associated-interface"], d, "associated_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["associated-interface"], "FirewallMulticastAddress-AssociatedInterface"); ok {
			if err = d.Set("associated_interface", vv); err != nil {
				return fmt.Errorf("Error reading associated_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading associated_interface: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallMulticastAddressColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallMulticastAddress-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallMulticastAddressComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallMulticastAddress-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("end_ip", flattenFirewallMulticastAddressEndIp(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "FirewallMulticastAddress-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallMulticastAddressName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallMulticastAddress-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenFirewallMulticastAddressStartIp(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "FirewallMulticastAddress-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("subnet", flattenFirewallMulticastAddressSubnet(o["subnet"], d, "subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet"], "FirewallMulticastAddress-Subnet"); ok {
			if err = d.Set("subnet", vv); err != nil {
				return fmt.Errorf("Error reading subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tagging", flattenFirewallMulticastAddressTagging(o["tagging"], d, "tagging")); err != nil {
			if vv, ok := fortiAPIPatch(o["tagging"], "FirewallMulticastAddress-Tagging"); ok {
				if err = d.Set("tagging", vv); err != nil {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tagging: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tagging"); ok {
			if err = d.Set("tagging", flattenFirewallMulticastAddressTagging(o["tagging"], d, "tagging")); err != nil {
				if vv, ok := fortiAPIPatch(o["tagging"], "FirewallMulticastAddress-Tagging"); ok {
					if err = d.Set("tagging", vv); err != nil {
						return fmt.Errorf("Error reading tagging: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenFirewallMulticastAddressType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallMulticastAddress-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenFirewallMulticastAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallMulticastAddressAssociatedInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallMulticastAddressColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastAddressComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastAddressEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastAddressName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastAddressStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastAddressSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandFirewallMulticastAddressTagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandFirewallMulticastAddressTaggingCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallMulticastAddressTaggingName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandFirewallMulticastAddressTaggingTags(d, i["tags"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallMulticastAddressTaggingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallMulticastAddressTaggingName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastAddressTaggingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallMulticastAddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallMulticastAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("associated_interface"); ok || d.HasChange("associated_interface") {
		t, err := expandFirewallMulticastAddressAssociatedInterface(d, v, "associated_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["associated-interface"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallMulticastAddressColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallMulticastAddressComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandFirewallMulticastAddressEndIp(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallMulticastAddressName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandFirewallMulticastAddressStartIp(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("subnet"); ok || d.HasChange("subnet") {
		t, err := expandFirewallMulticastAddressSubnet(d, v, "subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet"] = t
		}
	}

	if v, ok := d.GetOk("tagging"); ok || d.HasChange("tagging") {
		t, err := expandFirewallMulticastAddressTagging(d, v, "tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tagging"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallMulticastAddressType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
