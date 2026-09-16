// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Fabric-wide objects shared by non-root nodes.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCsfSharedObjects() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCsfSharedObjectsCreate,
		Read:   resourceSystemCsfSharedObjectsRead,
		Update: resourceSystemCsfSharedObjectsUpdate,
		Delete: resourceSystemCsfSharedObjectsDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"objects": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keys": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"pathname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"trusted_list_entry": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceSystemCsfSharedObjectsCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemCsfSharedObjects(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCsfSharedObjects resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemCsfSharedObjects(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemCsfSharedObjects(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemCsfSharedObjects resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemCsfSharedObjects(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemCsfSharedObjects resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemCsfSharedObjectsRead(d, m)
}

func resourceSystemCsfSharedObjectsUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemCsfSharedObjects(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsfSharedObjects resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemCsfSharedObjects(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsfSharedObjects resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemCsfSharedObjectsRead(d, m)
}

func resourceSystemCsfSharedObjectsDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteSystemCsfSharedObjects(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCsfSharedObjects resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCsfSharedObjectsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCsfSharedObjects(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemCsfSharedObjects resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCsfSharedObjects(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsfSharedObjects resource from API: %v", err)
	}
	return nil
}

func flattenSystemCsfSharedObjectsName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfSharedObjectsObjects2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keys"
		if _, ok := i["keys"]; ok {
			v := flattenSystemCsfSharedObjectsObjectsKeys2edl(i["keys"], d, pre_append)
			tmp["keys"] = fortiAPISubPartPatch(v, "SystemCsfSharedObjects-Objects-Keys")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pathname"
		if _, ok := i["pathname"]; ok {
			v := flattenSystemCsfSharedObjectsObjectsPathname2edl(i["pathname"], d, pre_append)
			tmp["pathname"] = fortiAPISubPartPatch(v, "SystemCsfSharedObjects-Objects-Pathname")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemCsfSharedObjectsObjectsKeys2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemCsfSharedObjectsObjectsKeysName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemCsfSharedObjectsObjects-Keys-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemCsfSharedObjectsObjectsKeysName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenSystemCsfSharedObjectsObjectsPathname2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfSharedObjectsTrustedListEntry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemCsfSharedObjects(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenSystemCsfSharedObjectsName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemCsfSharedObjects-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("objects", flattenSystemCsfSharedObjectsObjects2edl(o["objects"], d, "objects")); err != nil {
			if vv, ok := fortiAPIPatch(o["objects"], "SystemCsfSharedObjects-Objects"); ok {
				if err = d.Set("objects", vv); err != nil {
					return fmt.Errorf("Error reading objects: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading objects: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("objects"); ok {
			if err = d.Set("objects", flattenSystemCsfSharedObjectsObjects2edl(o["objects"], d, "objects")); err != nil {
				if vv, ok := fortiAPIPatch(o["objects"], "SystemCsfSharedObjects-Objects"); ok {
					if err = d.Set("objects", vv); err != nil {
						return fmt.Errorf("Error reading objects: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading objects: %v", err)
				}
			}
		}
	}

	if err = d.Set("trusted_list_entry", flattenSystemCsfSharedObjectsTrustedListEntry2edl(o["trusted-list-entry"], d, "trusted_list_entry")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusted-list-entry"], "SystemCsfSharedObjects-TrustedListEntry"); ok {
			if err = d.Set("trusted_list_entry", vv); err != nil {
				return fmt.Errorf("Error reading trusted_list_entry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusted_list_entry: %v", err)
		}
	}

	return nil
}

func flattenSystemCsfSharedObjectsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCsfSharedObjectsName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfSharedObjectsObjects2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keys"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSystemCsfSharedObjectsObjectsKeys2edl(d, i["keys"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["keys"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pathname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pathname"], _ = expandSystemCsfSharedObjectsObjectsPathname2edl(d, i["pathname"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemCsfSharedObjectsObjectsKeys2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemCsfSharedObjectsObjectsKeysName2edl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemCsfSharedObjectsObjectsKeysName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfSharedObjectsObjectsPathname2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfSharedObjectsTrustedListEntry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemCsfSharedObjects(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemCsfSharedObjectsName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("objects"); ok || d.HasChange("objects") {
		t, err := expandSystemCsfSharedObjectsObjects2edl(d, v, "objects")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["objects"] = t
		}
	}

	if v, ok := d.GetOk("trusted_list_entry"); ok || d.HasChange("trusted_list_entry") {
		t, err := expandSystemCsfSharedObjectsTrustedListEntry2edl(d, v, "trusted_list_entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-list-entry"] = t
		}
	}

	return &obj, nil
}
