// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: EAP Methods.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodCreate,
		Read:   resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodRead,
		Update: resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodDelete,

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
			"nai_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
				ForceNew: true,
				Optional: true,
			},
			"method": &schema.Schema{
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

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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
	nai_list := d.Get("nai_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_nai_realm"] = anqp_nai_realm
	paradict["nai_list"] = nai_list

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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
	nai_list := d.Get("nai_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_nai_realm"] = anqp_nai_realm
	paradict["nai_list"] = nai_list

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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
	nai_list := d.Get("nai_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_nai_realm"] = anqp_nai_realm
	paradict["nai_list"] = nai_list

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	nai_list := d.Get("nai_list").(string)
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
	if nai_list == "" {
		nai_list = importOptionChecking(m.(*FortiClient).Cfg, "nai_list")
		if nai_list == "" {
			return fmt.Errorf("Parameter nai_list is missing")
		}
		if err = d.Set("nai_list", nai_list); err != nil {
			return fmt.Errorf("Error set params nai_list: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_nai_realm"] = anqp_nai_realm
	paradict["nai_list"] = nai_list

	o, err := c.ReadWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex3rdl(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "val"
		if _, ok := i["val"]; ok {
			v := flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal3rdl(i["val"], d, pre_append)
			tmp["val"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam-Val")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("auth_param", flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam3rdl(o["auth-param"], d, "auth_param")); err != nil {
			if vv, ok := fortiAPIPatch(o["auth-param"], "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam"); ok {
				if err = d.Set("auth_param", vv); err != nil {
					return fmt.Errorf("Error reading auth_param: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading auth_param: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("auth_param"); ok {
			if err = d.Set("auth_param", flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam3rdl(o["auth-param"], d, "auth_param")); err != nil {
				if vv, ok := fortiAPIPatch(o["auth-param"], "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-AuthParam"); ok {
					if err = d.Set("auth_param", vv); err != nil {
						return fmt.Errorf("Error reading auth_param: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading auth_param: %v", err)
				}
			}
		}
	}

	if err = d.Set("index", flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex3rdl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("method", flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod3rdl(o["method"], d, "method")); err != nil {
		if vv, ok := fortiAPIPatch(o["method"], "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod-Method"); ok {
			if err = d.Set("method", vv); err != nil {
				return fmt.Errorf("Error reading method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex3rdl(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "val"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["val"], _ = expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal3rdl(d, i["val"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethod(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_param"); ok || d.HasChange("auth_param") {
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam3rdl(d, v, "auth_param")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-param"] = t
		}
	}

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodIndex3rdl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodMethod3rdl(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	return &obj, nil
}
