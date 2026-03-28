// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyRedirectProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyRedirectProfileCreate,
		Read:   resourceWebProxyRedirectProfileRead,
		Update: resourceWebProxyRedirectProfileUpdate,
		Delete: resourceWebProxyRedirectProfileDelete,

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
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"redirect_code": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"redirect_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"url": &schema.Schema{
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

func resourceWebProxyRedirectProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyRedirectProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyRedirectProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebProxyRedirectProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebProxyRedirectProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebProxyRedirectProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebProxyRedirectProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebProxyRedirectProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyRedirectProfileRead(d, m)
}

func resourceWebProxyRedirectProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyRedirectProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyRedirectProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebProxyRedirectProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyRedirectProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyRedirectProfileRead(d, m)
}

func resourceWebProxyRedirectProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebProxyRedirectProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyRedirectProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyRedirectProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyRedirectProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebProxyRedirectProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyRedirectProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyRedirectProfile resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyRedirectProfileEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWebProxyRedirectProfileEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WebProxyRedirectProfile-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redirect_code"
		if _, ok := i["redirect-code"]; ok {
			v := flattenWebProxyRedirectProfileEntriesRedirectCode(i["redirect-code"], d, pre_append)
			tmp["redirect_code"] = fortiAPISubPartPatch(v, "WebProxyRedirectProfile-Entries-RedirectCode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redirect_url"
		if _, ok := i["redirect-url"]; ok {
			v := flattenWebProxyRedirectProfileEntriesRedirectUrl(i["redirect-url"], d, pre_append)
			tmp["redirect_url"] = fortiAPISubPartPatch(v, "WebProxyRedirectProfile-Entries-RedirectUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenWebProxyRedirectProfileEntriesType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "WebProxyRedirectProfile-Entries-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			v := flattenWebProxyRedirectProfileEntriesUrl(i["url"], d, pre_append)
			tmp["url"] = fortiAPISubPartPatch(v, "WebProxyRedirectProfile-Entries-Url")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebProxyRedirectProfileEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyRedirectProfileEntriesRedirectCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyRedirectProfileEntriesRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyRedirectProfileEntriesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyRedirectProfileEntriesUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyRedirectProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyRedirectProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("entries", flattenWebProxyRedirectProfileEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "WebProxyRedirectProfile-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenWebProxyRedirectProfileEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "WebProxyRedirectProfile-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenWebProxyRedirectProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebProxyRedirectProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenWebProxyRedirectProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyRedirectProfileEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWebProxyRedirectProfileEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redirect_code"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["redirect-code"], _ = expandWebProxyRedirectProfileEntriesRedirectCode(d, i["redirect_code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redirect_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["redirect-url"], _ = expandWebProxyRedirectProfileEntriesRedirectUrl(d, i["redirect_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandWebProxyRedirectProfileEntriesType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url"], _ = expandWebProxyRedirectProfileEntriesUrl(d, i["url"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebProxyRedirectProfileEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyRedirectProfileEntriesRedirectCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyRedirectProfileEntriesRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyRedirectProfileEntriesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyRedirectProfileEntriesUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyRedirectProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyRedirectProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandWebProxyRedirectProfileEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebProxyRedirectProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
