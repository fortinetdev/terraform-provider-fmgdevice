// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NSX-T service chain.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNsxtServiceChain() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtServiceChainCreate,
		Read:   resourceNsxtServiceChainRead,
		Update: resourceNsxtServiceChainUpdate,
		Delete: resourceNsxtServiceChainDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_index": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"reverse_index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"vd": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
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

func resourceNsxtServiceChainCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectNsxtServiceChain(d)
	if err != nil {
		return fmt.Errorf("Error creating NsxtServiceChain resource while getting object: %v", err)
	}

	_, err = c.CreateNsxtServiceChain(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating NsxtServiceChain resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceNsxtServiceChainRead(d, m)
}

func resourceNsxtServiceChainUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectNsxtServiceChain(d)
	if err != nil {
		return fmt.Errorf("Error updating NsxtServiceChain resource while getting object: %v", err)
	}

	_, err = c.UpdateNsxtServiceChain(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating NsxtServiceChain resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceNsxtServiceChainRead(d, m)
}

func resourceNsxtServiceChainDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteNsxtServiceChain(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting NsxtServiceChain resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceNsxtServiceChainRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadNsxtServiceChain(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading NsxtServiceChain resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectNsxtServiceChain(d, o)
	if err != nil {
		return fmt.Errorf("Error reading NsxtServiceChain resource from API: %v", err)
	}
	return nil
}

func flattenNsxtServiceChainId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenNsxtServiceChainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenNsxtServiceChainServiceIndex(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenNsxtServiceChainServiceIndexId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "NsxtServiceChain-ServiceIndex-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenNsxtServiceChainServiceIndexName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "NsxtServiceChain-ServiceIndex-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reverse_index"
		if _, ok := i["reverse-index"]; ok {
			v := flattenNsxtServiceChainServiceIndexReverseIndex(i["reverse-index"], d, pre_append)
			tmp["reverse_index"] = fortiAPISubPartPatch(v, "NsxtServiceChain-ServiceIndex-ReverseIndex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vd"
		if _, ok := i["vd"]; ok {
			v := flattenNsxtServiceChainServiceIndexVd(i["vd"], d, pre_append)
			tmp["vd"] = fortiAPISubPartPatch(v, "NsxtServiceChain-ServiceIndex-Vd")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenNsxtServiceChainServiceIndexId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenNsxtServiceChainServiceIndexName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenNsxtServiceChainServiceIndexReverseIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenNsxtServiceChainServiceIndexVd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectNsxtServiceChain(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fosid", flattenNsxtServiceChainId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "NsxtServiceChain-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenNsxtServiceChainName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "NsxtServiceChain-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service_index", flattenNsxtServiceChainServiceIndex(o["service-index"], d, "service_index")); err != nil {
			if vv, ok := fortiAPIPatch(o["service-index"], "NsxtServiceChain-ServiceIndex"); ok {
				if err = d.Set("service_index", vv); err != nil {
					return fmt.Errorf("Error reading service_index: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading service_index: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service_index"); ok {
			if err = d.Set("service_index", flattenNsxtServiceChainServiceIndex(o["service-index"], d, "service_index")); err != nil {
				if vv, ok := fortiAPIPatch(o["service-index"], "NsxtServiceChain-ServiceIndex"); ok {
					if err = d.Set("service_index", vv); err != nil {
						return fmt.Errorf("Error reading service_index: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading service_index: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenNsxtServiceChainFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandNsxtServiceChainId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainServiceIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandNsxtServiceChainServiceIndexId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandNsxtServiceChainServiceIndexName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reverse_index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["reverse-index"], _ = expandNsxtServiceChainServiceIndexReverseIndex(d, i["reverse_index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vd"], _ = expandNsxtServiceChainServiceIndexVd(d, i["vd"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandNsxtServiceChainServiceIndexId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainServiceIndexName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainServiceIndexReverseIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainServiceIndexVd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectNsxtServiceChain(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandNsxtServiceChainId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandNsxtServiceChainName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("service_index"); ok || d.HasChange("service_index") {
		t, err := expandNsxtServiceChainServiceIndex(d, v, "service_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-index"] = t
		}
	}

	return &obj, nil
}
