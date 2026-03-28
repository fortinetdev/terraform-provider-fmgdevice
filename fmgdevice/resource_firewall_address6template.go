// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv6 address templates.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAddress6Template() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAddress6TemplateCreate,
		Read:   resourceFirewallAddress6TemplateRead,
		Update: resourceFirewallAddress6TemplateUpdate,
		Delete: resourceFirewallAddress6TemplateDelete,

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
			"_image_base64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_segment": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bits": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"exclusive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
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
					},
				},
			},
			"subnet_segment_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uuid": &schema.Schema{
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

func resourceFirewallAddress6TemplateCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAddress6Template(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAddress6Template resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAddress6Template(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAddress6Template(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAddress6Template resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAddress6Template(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAddress6Template resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAddress6TemplateRead(d, m)
}

func resourceFirewallAddress6TemplateUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAddress6Template(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6Template resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAddress6Template(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6Template resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAddress6TemplateRead(d, m)
}

func resourceFirewallAddress6TemplateDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallAddress6Template(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAddress6Template resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddress6TemplateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallAddress6Template(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAddress6Template resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAddress6Template(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddress6Template resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAddress6TemplateImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegment(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bits"
		if _, ok := i["bits"]; ok {
			v := flattenFirewallAddress6TemplateSubnetSegmentBits(i["bits"], d, pre_append)
			tmp["bits"] = fortiAPISubPartPatch(v, "FirewallAddress6Template-SubnetSegment-Bits")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusive"
		if _, ok := i["exclusive"]; ok {
			v := flattenFirewallAddress6TemplateSubnetSegmentExclusive(i["exclusive"], d, pre_append)
			tmp["exclusive"] = fortiAPISubPartPatch(v, "FirewallAddress6Template-SubnetSegment-Exclusive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallAddress6TemplateSubnetSegmentId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallAddress6Template-SubnetSegment-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallAddress6TemplateSubnetSegmentName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddress6Template-SubnetSegment-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
		if _, ok := i["values"]; ok {
			v := flattenFirewallAddress6TemplateSubnetSegmentValues(i["values"], d, pre_append)
			tmp["values"] = fortiAPISubPartPatch(v, "FirewallAddress6Template-SubnetSegment-Values")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddress6TemplateSubnetSegmentBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentExclusive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentValues(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAddress6TemplateSubnetSegmentValuesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddress6TemplateSubnetSegment-Values-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenFirewallAddress6TemplateSubnetSegmentValuesValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "FirewallAddress6TemplateSubnetSegment-Values-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddress6TemplateSubnetSegmentValuesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentValuesValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAddress6Template(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_image_base64", flattenFirewallAddress6TemplateImageBase64(o["_image-base64"], d, "_image_base64")); err != nil {
		if vv, ok := fortiAPIPatch(o["_image-base64"], "FirewallAddress6Template-ImageBase64"); ok {
			if err = d.Set("_image_base64", vv); err != nil {
				return fmt.Errorf("Error reading _image_base64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _image_base64: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenFirewallAddress6TemplateFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "FirewallAddress6Template-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("ip6", flattenFirewallAddress6TemplateIp6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "FirewallAddress6Template-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallAddress6TemplateName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallAddress6Template-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("subnet_segment", flattenFirewallAddress6TemplateSubnetSegment(o["subnet-segment"], d, "subnet_segment")); err != nil {
			if vv, ok := fortiAPIPatch(o["subnet-segment"], "FirewallAddress6Template-SubnetSegment"); ok {
				if err = d.Set("subnet_segment", vv); err != nil {
					return fmt.Errorf("Error reading subnet_segment: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading subnet_segment: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("subnet_segment"); ok {
			if err = d.Set("subnet_segment", flattenFirewallAddress6TemplateSubnetSegment(o["subnet-segment"], d, "subnet_segment")); err != nil {
				if vv, ok := fortiAPIPatch(o["subnet-segment"], "FirewallAddress6Template-SubnetSegment"); ok {
					if err = d.Set("subnet_segment", vv); err != nil {
						return fmt.Errorf("Error reading subnet_segment: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading subnet_segment: %v", err)
				}
			}
		}
	}

	if err = d.Set("subnet_segment_count", flattenFirewallAddress6TemplateSubnetSegmentCount(o["subnet-segment-count"], d, "subnet_segment_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet-segment-count"], "FirewallAddress6Template-SubnetSegmentCount"); ok {
			if err = d.Set("subnet_segment_count", vv); err != nil {
				return fmt.Errorf("Error reading subnet_segment_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet_segment_count: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallAddress6TemplateUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallAddress6Template-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenFirewallAddress6TemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAddress6TemplateImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bits"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bits"], _ = expandFirewallAddress6TemplateSubnetSegmentBits(d, i["bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exclusive"], _ = expandFirewallAddress6TemplateSubnetSegmentExclusive(d, i["exclusive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallAddress6TemplateSubnetSegmentId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallAddress6TemplateSubnetSegmentName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAddress6TemplateSubnetSegmentValues(d, i["values"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["values"] = t
			}
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6TemplateSubnetSegmentBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentExclusive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentValues(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallAddress6TemplateSubnetSegmentValuesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandFirewallAddress6TemplateSubnetSegmentValuesValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6TemplateSubnetSegmentValuesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentValuesValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAddress6Template(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_image_base64"); ok || d.HasChange("_image_base64") {
		t, err := expandFirewallAddress6TemplateImageBase64(d, v, "_image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_image-base64"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandFirewallAddress6TemplateFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok || d.HasChange("ip6") {
		t, err := expandFirewallAddress6TemplateIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallAddress6TemplateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("subnet_segment"); ok || d.HasChange("subnet_segment") {
		t, err := expandFirewallAddress6TemplateSubnetSegment(d, v, "subnet_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-segment"] = t
		}
	}

	if v, ok := d.GetOk("subnet_segment_count"); ok || d.HasChange("subnet_segment_count") {
		t, err := expandFirewallAddress6TemplateSubnetSegmentCount(d, v, "subnet_segment_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-segment-count"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallAddress6TemplateUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
