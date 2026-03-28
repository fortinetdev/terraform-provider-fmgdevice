// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> IPv6 subnet segments.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAddress6TemplateSubnetSegment() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAddress6TemplateSubnetSegmentCreate,
		Read:   resourceFirewallAddress6TemplateSubnetSegmentRead,
		Update: resourceFirewallAddress6TemplateSubnetSegmentUpdate,
		Delete: resourceFirewallAddress6TemplateSubnetSegmentDelete,

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
			"address6_template": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bits": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"exclusive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"values": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallAddress6TemplateSubnetSegmentCreate(d *schema.ResourceData, m interface{}) error {
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
	address6_template := d.Get("address6_template").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6_template"] = address6_template

	obj, err := getObjectFirewallAddress6TemplateSubnetSegment(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAddress6TemplateSubnetSegment resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAddress6TemplateSubnetSegment(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAddress6TemplateSubnetSegment(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAddress6TemplateSubnetSegment resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallAddress6TemplateSubnetSegment(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAddress6TemplateSubnetSegment resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallAddress6TemplateSubnetSegmentRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallAddress6TemplateSubnetSegment resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallAddress6TemplateSubnetSegmentRead(d, m)
}

func resourceFirewallAddress6TemplateSubnetSegmentUpdate(d *schema.ResourceData, m interface{}) error {
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
	address6_template := d.Get("address6_template").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6_template"] = address6_template

	obj, err := getObjectFirewallAddress6TemplateSubnetSegment(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6TemplateSubnetSegment resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAddress6TemplateSubnetSegment(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6TemplateSubnetSegment resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallAddress6TemplateSubnetSegmentRead(d, m)
}

func resourceFirewallAddress6TemplateSubnetSegmentDelete(d *schema.ResourceData, m interface{}) error {
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
	address6_template := d.Get("address6_template").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6_template"] = address6_template

	wsParams["adom"] = adomv

	err = c.DeleteFirewallAddress6TemplateSubnetSegment(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAddress6TemplateSubnetSegment resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddress6TemplateSubnetSegmentRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	address6_template := d.Get("address6_template").(string)
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
	if address6_template == "" {
		address6_template = importOptionChecking(m.(*FortiClient).Cfg, "address6_template")
		if address6_template == "" {
			return fmt.Errorf("Parameter address6_template is missing")
		}
		if err = d.Set("address6_template", address6_template); err != nil {
			return fmt.Errorf("Error set params address6_template: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6_template"] = address6_template

	o, err := c.ReadFirewallAddress6TemplateSubnetSegment(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAddress6TemplateSubnetSegment resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAddress6TemplateSubnetSegment(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddress6TemplateSubnetSegment resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAddress6TemplateSubnetSegmentBits2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentExclusive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentValues2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallAddress6TemplateSubnetSegmentValuesName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddress6TemplateSubnetSegment-Values-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenFirewallAddress6TemplateSubnetSegmentValuesValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "FirewallAddress6TemplateSubnetSegment-Values-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddress6TemplateSubnetSegmentValuesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentValuesValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAddress6TemplateSubnetSegment(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("bits", flattenFirewallAddress6TemplateSubnetSegmentBits2edl(o["bits"], d, "bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["bits"], "FirewallAddress6TemplateSubnetSegment-Bits"); ok {
			if err = d.Set("bits", vv); err != nil {
				return fmt.Errorf("Error reading bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bits: %v", err)
		}
	}

	if err = d.Set("exclusive", flattenFirewallAddress6TemplateSubnetSegmentExclusive2edl(o["exclusive"], d, "exclusive")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclusive"], "FirewallAddress6TemplateSubnetSegment-Exclusive"); ok {
			if err = d.Set("exclusive", vv); err != nil {
				return fmt.Errorf("Error reading exclusive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclusive: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallAddress6TemplateSubnetSegmentId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallAddress6TemplateSubnetSegment-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallAddress6TemplateSubnetSegmentName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallAddress6TemplateSubnetSegment-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("values", flattenFirewallAddress6TemplateSubnetSegmentValues2edl(o["values"], d, "values")); err != nil {
			if vv, ok := fortiAPIPatch(o["values"], "FirewallAddress6TemplateSubnetSegment-Values"); ok {
				if err = d.Set("values", vv); err != nil {
					return fmt.Errorf("Error reading values: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading values: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("values"); ok {
			if err = d.Set("values", flattenFirewallAddress6TemplateSubnetSegmentValues2edl(o["values"], d, "values")); err != nil {
				if vv, ok := fortiAPIPatch(o["values"], "FirewallAddress6TemplateSubnetSegment-Values"); ok {
					if err = d.Set("values", vv); err != nil {
						return fmt.Errorf("Error reading values: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading values: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallAddress6TemplateSubnetSegmentFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAddress6TemplateSubnetSegmentBits2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentExclusive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentValues2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallAddress6TemplateSubnetSegmentValuesName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandFirewallAddress6TemplateSubnetSegmentValuesValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6TemplateSubnetSegmentValuesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentValuesValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAddress6TemplateSubnetSegment(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bits"); ok || d.HasChange("bits") {
		t, err := expandFirewallAddress6TemplateSubnetSegmentBits2edl(d, v, "bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bits"] = t
		}
	}

	if v, ok := d.GetOk("exclusive"); ok || d.HasChange("exclusive") {
		t, err := expandFirewallAddress6TemplateSubnetSegmentExclusive2edl(d, v, "exclusive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclusive"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallAddress6TemplateSubnetSegmentId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallAddress6TemplateSubnetSegmentName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("values"); ok || d.HasChange("values") {
		t, err := expandFirewallAddress6TemplateSubnetSegmentValues2edl(d, v, "values")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["values"] = t
		}
	}

	return &obj, nil
}
