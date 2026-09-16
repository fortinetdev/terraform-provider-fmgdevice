// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CMDB table entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCsfSharedObjectsObjects() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCsfSharedObjectsObjectsCreate,
		Read:   resourceSystemCsfSharedObjectsObjectsRead,
		Update: resourceSystemCsfSharedObjectsObjectsUpdate,
		Delete: resourceSystemCsfSharedObjectsObjectsDelete,

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
			"shared_objects": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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

func resourceSystemCsfSharedObjectsObjectsCreate(d *schema.ResourceData, m interface{}) error {
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
	shared_objects := d.Get("shared_objects").(string)
	paradict["device"] = device_name
	paradict["shared_objects"] = shared_objects

	obj, err := getObjectSystemCsfSharedObjectsObjects(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCsfSharedObjectsObjects resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("pathname")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemCsfSharedObjectsObjects(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemCsfSharedObjectsObjects(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemCsfSharedObjectsObjects resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemCsfSharedObjectsObjects(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemCsfSharedObjectsObjects resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "pathname"))

	return resourceSystemCsfSharedObjectsObjectsRead(d, m)
}

func resourceSystemCsfSharedObjectsObjectsUpdate(d *schema.ResourceData, m interface{}) error {
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
	shared_objects := d.Get("shared_objects").(string)
	paradict["device"] = device_name
	paradict["shared_objects"] = shared_objects

	obj, err := getObjectSystemCsfSharedObjectsObjects(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsfSharedObjectsObjects resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemCsfSharedObjectsObjects(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsfSharedObjectsObjects resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "pathname"))

	return resourceSystemCsfSharedObjectsObjectsRead(d, m)
}

func resourceSystemCsfSharedObjectsObjectsDelete(d *schema.ResourceData, m interface{}) error {
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
	shared_objects := d.Get("shared_objects").(string)
	paradict["device"] = device_name
	paradict["shared_objects"] = shared_objects

	wsParams["adom"] = adomv

	err = c.DeleteSystemCsfSharedObjectsObjects(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCsfSharedObjectsObjects resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCsfSharedObjectsObjectsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	shared_objects := d.Get("shared_objects").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if shared_objects == "" {
		shared_objects = importOptionChecking(m.(*FortiClient).Cfg, "shared_objects")
		if shared_objects == "" {
			return fmt.Errorf("Parameter shared_objects is missing")
		}
		if err = d.Set("shared_objects", shared_objects); err != nil {
			return fmt.Errorf("Error set params shared_objects: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["shared_objects"] = shared_objects

	o, err := c.ReadSystemCsfSharedObjectsObjects(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemCsfSharedObjectsObjects resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCsfSharedObjectsObjects(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsfSharedObjectsObjects resource from API: %v", err)
	}
	return nil
}

func flattenSystemCsfSharedObjectsObjectsKeys3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemCsfSharedObjectsObjectsKeysName3rdl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemCsfSharedObjectsObjects-Keys-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemCsfSharedObjectsObjectsKeysName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenSystemCsfSharedObjectsObjectsPathname3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemCsfSharedObjectsObjects(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("keys", flattenSystemCsfSharedObjectsObjectsKeys3rdl(o["keys"], d, "keys")); err != nil {
			if vv, ok := fortiAPIPatch(o["keys"], "SystemCsfSharedObjectsObjects-Keys"); ok {
				if err = d.Set("keys", vv); err != nil {
					return fmt.Errorf("Error reading keys: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading keys: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("keys"); ok {
			if err = d.Set("keys", flattenSystemCsfSharedObjectsObjectsKeys3rdl(o["keys"], d, "keys")); err != nil {
				if vv, ok := fortiAPIPatch(o["keys"], "SystemCsfSharedObjectsObjects-Keys"); ok {
					if err = d.Set("keys", vv); err != nil {
						return fmt.Errorf("Error reading keys: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading keys: %v", err)
				}
			}
		}
	}

	if err = d.Set("pathname", flattenSystemCsfSharedObjectsObjectsPathname3rdl(o["pathname"], d, "pathname")); err != nil {
		if vv, ok := fortiAPIPatch(o["pathname"], "SystemCsfSharedObjectsObjects-Pathname"); ok {
			if err = d.Set("pathname", vv); err != nil {
				return fmt.Errorf("Error reading pathname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pathname: %v", err)
		}
	}

	return nil
}

func flattenSystemCsfSharedObjectsObjectsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCsfSharedObjectsObjectsKeys3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemCsfSharedObjectsObjectsKeysName3rdl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemCsfSharedObjectsObjectsKeysName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfSharedObjectsObjectsPathname3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCsfSharedObjectsObjects(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("keys"); ok || d.HasChange("keys") {
		t, err := expandSystemCsfSharedObjectsObjectsKeys3rdl(d, v, "keys")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keys"] = t
		}
	}

	if v, ok := d.GetOk("pathname"); ok || d.HasChange("pathname") {
		t, err := expandSystemCsfSharedObjectsObjectsPathname3rdl(d, v, "pathname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pathname"] = t
		}
	}

	return &obj, nil
}
