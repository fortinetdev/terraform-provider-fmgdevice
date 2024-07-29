// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure network access identifier (NAI) realm.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpNaiRealm() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpNaiRealmCreate,
		Read:   resourceWirelessControllerHotspot20AnqpNaiRealmRead,
		Update: resourceWirelessControllerHotspot20AnqpNaiRealmUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpNaiRealmDelete,

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
			"nai_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"eap_method": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_param": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"index": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"val": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"index": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"method": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"encoding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nai_realm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerHotspot20AnqpNaiRealmCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	obj, err := getObjectWirelessControllerHotspot20AnqpNaiRealm(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpNaiRealm resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20AnqpNaiRealm(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpNaiRealm resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20AnqpNaiRealmRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpNaiRealmUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	obj, err := getObjectWirelessControllerHotspot20AnqpNaiRealm(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpNaiRealm resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20AnqpNaiRealm(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpNaiRealm resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20AnqpNaiRealmRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpNaiRealmDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	err = c.DeleteWirelessControllerHotspot20AnqpNaiRealm(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpNaiRealm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpNaiRealmRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerHotspot20AnqpNaiRealm(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpNaiRealm resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpNaiRealm(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpNaiRealm resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eap_method"
		if _, ok := i["eap-method"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(i["eap-method"], d, pre_append)
			tmp["eap_method"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealm-NaiList-EapMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encoding"
		if _, ok := i["encoding"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding(i["encoding"], d, pre_append)
			tmp["encoding"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealm-NaiList-Encoding")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nai_realm"
		if _, ok := i["nai-realm"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm(i["nai-realm"], d, pre_append)
			tmp["nai_realm"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealm-NaiList-NaiRealm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealm-NaiList-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_param"
		if _, ok := i["auth-param"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(i["auth-param"], d, pre_append)
			tmp["auth_param"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod-AuthParam")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := i["method"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod(i["method"], d, pre_append)
			tmp["method"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod-Method")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "val"
		if _, ok := i["val"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal(i["val"], d, pre_append)
			tmp["val"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Val")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpNaiRealm(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("nai_list", flattenWirelessControllerHotspot20AnqpNaiRealmNaiList(o["nai-list"], d, "nai_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["nai-list"], "WirelessControllerHotspot20AnqpNaiRealm-NaiList"); ok {
				if err = d.Set("nai_list", vv); err != nil {
					return fmt.Errorf("Error reading nai_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nai_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nai_list"); ok {
			if err = d.Set("nai_list", flattenWirelessControllerHotspot20AnqpNaiRealmNaiList(o["nai-list"], d, "nai_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["nai-list"], "WirelessControllerHotspot20AnqpNaiRealm-NaiList"); ok {
					if err = d.Set("nai_list", vv); err != nil {
						return fmt.Errorf("Error reading nai_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nai_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenWirelessControllerHotspot20AnqpNaiRealmName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerHotspot20AnqpNaiRealm-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpNaiRealmFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "eap_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(d, i["eap_method"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["eap-method"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encoding"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["encoding"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding(d, i["encoding"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nai_realm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nai-realm"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm(d, i["nai_realm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_param"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d, i["auth_param"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["auth-param"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["method"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod(d, i["method"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "val"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["val"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal(d, i["val"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpNaiRealm(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("nai_list"); ok || d.HasChange("nai_list") {
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiList(d, v, "nai_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nai-list"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
