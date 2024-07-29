// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: NAI list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpNaiRealmNaiListCreate,
		Read:   resourceWirelessControllerHotspot20AnqpNaiRealmNaiListRead,
		Update: resourceWirelessControllerHotspot20AnqpNaiRealmNaiListUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpNaiRealmNaiListDelete,

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
			"anqp_nai_realm": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListCreate(d *schema.ResourceData, m interface{}) error {
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
	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_nai_realm"] = anqp_nai_realm

	obj, err := getObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpNaiRealmNaiList resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20AnqpNaiRealmNaiList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpNaiRealmNaiList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20AnqpNaiRealmNaiListRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListUpdate(d *schema.ResourceData, m interface{}) error {
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
	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_nai_realm"] = anqp_nai_realm

	obj, err := getObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpNaiRealmNaiList resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20AnqpNaiRealmNaiList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpNaiRealmNaiList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20AnqpNaiRealmNaiListRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListDelete(d *schema.ResourceData, m interface{}) error {
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
	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_nai_realm"] = anqp_nai_realm

	err = c.DeleteWirelessControllerHotspot20AnqpNaiRealmNaiList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpNaiRealmNaiList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
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
	if anqp_nai_realm == "" {
		anqp_nai_realm = importOptionChecking(m.(*FortiClient).Cfg, "anqp_nai_realm")
		if anqp_nai_realm == "" {
			return fmt.Errorf("Parameter anqp_nai_realm is missing")
		}
		if err = d.Set("anqp_nai_realm", anqp_nai_realm); err != nil {
			return fmt.Errorf("Error set params anqp_nai_realm: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_nai_realm"] = anqp_nai_realm

	o, err := c.ReadWirelessControllerHotspot20AnqpNaiRealmNaiList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpNaiRealmNaiList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpNaiRealmNaiList resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam2edl(i["auth-param"], d, pre_append)
			tmp["auth_param"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod-AuthParam")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex2edl(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := i["method"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod2edl(i["method"], d, pre_append)
			tmp["method"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod-Method")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex2edl(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "val"
		if _, ok := i["val"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal2edl(i["val"], d, pre_append)
			tmp["val"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Val")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("eap_method", flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod2edl(o["eap-method"], d, "eap_method")); err != nil {
			if vv, ok := fortiAPIPatch(o["eap-method"], "WirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod"); ok {
				if err = d.Set("eap_method", vv); err != nil {
					return fmt.Errorf("Error reading eap_method: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading eap_method: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("eap_method"); ok {
			if err = d.Set("eap_method", flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod2edl(o["eap-method"], d, "eap_method")); err != nil {
				if vv, ok := fortiAPIPatch(o["eap-method"], "WirelessControllerHotspot20AnqpNaiRealmNaiList-EapMethod"); ok {
					if err = d.Set("eap_method", vv); err != nil {
						return fmt.Errorf("Error reading eap_method: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading eap_method: %v", err)
				}
			}
		}
	}

	if err = d.Set("encoding", flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding2edl(o["encoding"], d, "encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["encoding"], "WirelessControllerHotspot20AnqpNaiRealmNaiList-Encoding"); ok {
			if err = d.Set("encoding", vv); err != nil {
				return fmt.Errorf("Error reading encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encoding: %v", err)
		}
	}

	if err = d.Set("nai_realm", flattenWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm2edl(o["nai-realm"], d, "nai_realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["nai-realm"], "WirelessControllerHotspot20AnqpNaiRealmNaiList-NaiRealm"); ok {
			if err = d.Set("nai_realm", vv); err != nil {
				return fmt.Errorf("Error reading nai_realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nai_realm: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerHotspot20AnqpNaiRealmNaiListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerHotspot20AnqpNaiRealmNaiList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam2edl(d, i["auth_param"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["auth-param"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex2edl(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["method"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod2edl(d, i["method"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex2edl(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "val"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["val"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal2edl(d, i["val"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpNaiRealmNaiList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("eap_method"); ok || d.HasChange("eap_method") {
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod2edl(d, v, "eap_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-method"] = t
		}
	}

	if v, ok := d.GetOk("encoding"); ok || d.HasChange("encoding") {
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEncoding2edl(d, v, "encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encoding"] = t
		}
	}

	if v, ok := d.GetOk("nai_realm"); ok || d.HasChange("nai_realm") {
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListNaiRealm2edl(d, v, "nai_realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nai-realm"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
