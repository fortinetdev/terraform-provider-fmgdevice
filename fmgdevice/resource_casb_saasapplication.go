// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure CASB SaaS application.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbSaasApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbSaasApplicationCreate,
		Read:   resourceCasbSaasApplicationRead,
		Update: resourceCasbSaasApplicationUpdate,
		Delete: resourceCasbSaasApplicationDelete,

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
			"casb_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domains": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"input_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"default": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fallback_input": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"required": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"output_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"optional": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"required": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceCasbSaasApplicationCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbSaasApplication(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbSaasApplication resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbSaasApplication(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbSaasApplication(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbSaasApplication resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateCasbSaasApplication(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbSaasApplication resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceCasbSaasApplicationRead(d, m)
}

func resourceCasbSaasApplicationUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbSaasApplication(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbSaasApplication resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbSaasApplication(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbSaasApplication resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCasbSaasApplicationRead(d, m)
}

func resourceCasbSaasApplicationDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteCasbSaasApplication(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbSaasApplication resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbSaasApplicationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbSaasApplication(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbSaasApplication resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbSaasApplication(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbSaasApplication resource from API: %v", err)
	}
	return nil
}

func flattenCasbSaasApplicationCasbName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationDomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbSaasApplicationInputAttributes(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attr_type"
		if _, ok := i["attr-type"]; ok {
			v := flattenCasbSaasApplicationInputAttributesAttrType(i["attr-type"], d, pre_append)
			tmp["attr_type"] = fortiAPISubPartPatch(v, "CasbSaasApplication-InputAttributes-AttrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := i["default"]; ok {
			v := flattenCasbSaasApplicationInputAttributesDefault(i["default"], d, pre_append)
			tmp["default"] = fortiAPISubPartPatch(v, "CasbSaasApplication-InputAttributes-Default")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenCasbSaasApplicationInputAttributesDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "CasbSaasApplication-InputAttributes-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fallback_input"
		if _, ok := i["fallback-input"]; ok {
			v := flattenCasbSaasApplicationInputAttributesFallbackInput(i["fallback-input"], d, pre_append)
			tmp["fallback_input"] = fortiAPISubPartPatch(v, "CasbSaasApplication-InputAttributes-FallbackInput")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbSaasApplicationInputAttributesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbSaasApplication-InputAttributes-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "required"
		if _, ok := i["required"]; ok {
			v := flattenCasbSaasApplicationInputAttributesRequired(i["required"], d, pre_append)
			tmp["required"] = fortiAPISubPartPatch(v, "CasbSaasApplication-InputAttributes-Required")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenCasbSaasApplicationInputAttributesType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "CasbSaasApplication-InputAttributes-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbSaasApplicationInputAttributesAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationInputAttributesDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationInputAttributesDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationInputAttributesFallbackInput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationInputAttributesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationInputAttributesRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationInputAttributesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationOutputAttributes(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attr_type"
		if _, ok := i["attr-type"]; ok {
			v := flattenCasbSaasApplicationOutputAttributesAttrType(i["attr-type"], d, pre_append)
			tmp["attr_type"] = fortiAPISubPartPatch(v, "CasbSaasApplication-OutputAttributes-AttrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenCasbSaasApplicationOutputAttributesDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "CasbSaasApplication-OutputAttributes-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbSaasApplicationOutputAttributesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbSaasApplication-OutputAttributes-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "optional"
		if _, ok := i["optional"]; ok {
			v := flattenCasbSaasApplicationOutputAttributesOptional(i["optional"], d, pre_append)
			tmp["optional"] = fortiAPISubPartPatch(v, "CasbSaasApplication-OutputAttributes-Optional")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "required"
		if _, ok := i["required"]; ok {
			v := flattenCasbSaasApplicationOutputAttributesRequired(i["required"], d, pre_append)
			tmp["required"] = fortiAPISubPartPatch(v, "CasbSaasApplication-OutputAttributes-Required")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenCasbSaasApplicationOutputAttributesType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "CasbSaasApplication-OutputAttributes-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbSaasApplicationOutputAttributesAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationOutputAttributesDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationOutputAttributesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationOutputAttributesOptional(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationOutputAttributesRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationOutputAttributesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbSaasApplicationUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbSaasApplication(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("casb_name", flattenCasbSaasApplicationCasbName(o["casb-name"], d, "casb_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["casb-name"], "CasbSaasApplication-CasbName"); ok {
			if err = d.Set("casb_name", vv); err != nil {
				return fmt.Errorf("Error reading casb_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading casb_name: %v", err)
		}
	}

	if err = d.Set("description", flattenCasbSaasApplicationDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "CasbSaasApplication-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("domains", flattenCasbSaasApplicationDomains(o["domains"], d, "domains")); err != nil {
		if vv, ok := fortiAPIPatch(o["domains"], "CasbSaasApplication-Domains"); ok {
			if err = d.Set("domains", vv); err != nil {
				return fmt.Errorf("Error reading domains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domains: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("input_attributes", flattenCasbSaasApplicationInputAttributes(o["input-attributes"], d, "input_attributes")); err != nil {
			if vv, ok := fortiAPIPatch(o["input-attributes"], "CasbSaasApplication-InputAttributes"); ok {
				if err = d.Set("input_attributes", vv); err != nil {
					return fmt.Errorf("Error reading input_attributes: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading input_attributes: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("input_attributes"); ok {
			if err = d.Set("input_attributes", flattenCasbSaasApplicationInputAttributes(o["input-attributes"], d, "input_attributes")); err != nil {
				if vv, ok := fortiAPIPatch(o["input-attributes"], "CasbSaasApplication-InputAttributes"); ok {
					if err = d.Set("input_attributes", vv); err != nil {
						return fmt.Errorf("Error reading input_attributes: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading input_attributes: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenCasbSaasApplicationName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CasbSaasApplication-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("output_attributes", flattenCasbSaasApplicationOutputAttributes(o["output-attributes"], d, "output_attributes")); err != nil {
			if vv, ok := fortiAPIPatch(o["output-attributes"], "CasbSaasApplication-OutputAttributes"); ok {
				if err = d.Set("output_attributes", vv); err != nil {
					return fmt.Errorf("Error reading output_attributes: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading output_attributes: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("output_attributes"); ok {
			if err = d.Set("output_attributes", flattenCasbSaasApplicationOutputAttributes(o["output-attributes"], d, "output_attributes")); err != nil {
				if vv, ok := fortiAPIPatch(o["output-attributes"], "CasbSaasApplication-OutputAttributes"); ok {
					if err = d.Set("output_attributes", vv); err != nil {
						return fmt.Errorf("Error reading output_attributes: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading output_attributes: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenCasbSaasApplicationStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "CasbSaasApplication-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenCasbSaasApplicationType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "CasbSaasApplication-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenCasbSaasApplicationUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "CasbSaasApplication-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenCasbSaasApplicationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbSaasApplicationCasbName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationDomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbSaasApplicationInputAttributes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attr-type"], _ = expandCasbSaasApplicationInputAttributesAttrType(d, i["attr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default"], _ = expandCasbSaasApplicationInputAttributesDefault(d, i["default"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandCasbSaasApplicationInputAttributesDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fallback_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fallback-input"], _ = expandCasbSaasApplicationInputAttributesFallbackInput(d, i["fallback_input"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbSaasApplicationInputAttributesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "required"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["required"], _ = expandCasbSaasApplicationInputAttributesRequired(d, i["required"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandCasbSaasApplicationInputAttributesType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbSaasApplicationInputAttributesAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationInputAttributesDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationInputAttributesDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationInputAttributesFallbackInput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationInputAttributesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationInputAttributesRequired(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationInputAttributesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationOutputAttributes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attr-type"], _ = expandCasbSaasApplicationOutputAttributesAttrType(d, i["attr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandCasbSaasApplicationOutputAttributesDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbSaasApplicationOutputAttributesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "optional"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["optional"], _ = expandCasbSaasApplicationOutputAttributesOptional(d, i["optional"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "required"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["required"], _ = expandCasbSaasApplicationOutputAttributesRequired(d, i["required"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandCasbSaasApplicationOutputAttributesType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbSaasApplicationOutputAttributesAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationOutputAttributesDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationOutputAttributesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationOutputAttributesOptional(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationOutputAttributesRequired(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationOutputAttributesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbSaasApplication(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("casb_name"); ok || d.HasChange("casb_name") {
		t, err := expandCasbSaasApplicationCasbName(d, v, "casb_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandCasbSaasApplicationDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("domains"); ok || d.HasChange("domains") {
		t, err := expandCasbSaasApplicationDomains(d, v, "domains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domains"] = t
		}
	}

	if v, ok := d.GetOk("input_attributes"); ok || d.HasChange("input_attributes") {
		t, err := expandCasbSaasApplicationInputAttributes(d, v, "input_attributes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-attributes"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCasbSaasApplicationName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("output_attributes"); ok || d.HasChange("output_attributes") {
		t, err := expandCasbSaasApplicationOutputAttributes(d, v, "output_attributes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output-attributes"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandCasbSaasApplicationStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandCasbSaasApplicationType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandCasbSaasApplicationUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
