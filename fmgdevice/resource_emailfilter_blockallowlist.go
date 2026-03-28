// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure anti-spam block/allow list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterBlockAllowList() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterBlockAllowListCreate,
		Read:   resourceEmailfilterBlockAllowListRead,
		Update: resourceEmailfilterBlockAllowListUpdate,
		Delete: resourceEmailfilterBlockAllowListDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip4_subnet": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ip6_subnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pattern_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"email_pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceEmailfilterBlockAllowListCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEmailfilterBlockAllowList(d)
	if err != nil {
		return fmt.Errorf("Error creating EmailfilterBlockAllowList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadEmailfilterBlockAllowList(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateEmailfilterBlockAllowList(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating EmailfilterBlockAllowList resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateEmailfilterBlockAllowList(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating EmailfilterBlockAllowList resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceEmailfilterBlockAllowListRead(d, m)
			} else {
				return fmt.Errorf("Error creating EmailfilterBlockAllowList resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterBlockAllowListRead(d, m)
}

func resourceEmailfilterBlockAllowListUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEmailfilterBlockAllowList(d)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterBlockAllowList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEmailfilterBlockAllowList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterBlockAllowList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterBlockAllowListRead(d, m)
}

func resourceEmailfilterBlockAllowListDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteEmailfilterBlockAllowList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterBlockAllowList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterBlockAllowListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadEmailfilterBlockAllowList(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EmailfilterBlockAllowList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterBlockAllowList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterBlockAllowList resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterBlockAllowListComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenEmailfilterBlockAllowListEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "EmailfilterBlockAllowList-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {
			v := flattenEmailfilterBlockAllowListEntriesAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "EmailfilterBlockAllowList-Entries-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenEmailfilterBlockAllowListEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "EmailfilterBlockAllowList-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip4_subnet"
		if _, ok := i["ip4-subnet"]; ok {
			v := flattenEmailfilterBlockAllowListEntriesIp4Subnet(i["ip4-subnet"], d, pre_append)
			tmp["ip4_subnet"] = fortiAPISubPartPatch(v, "EmailfilterBlockAllowList-Entries-Ip4Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_subnet"
		if _, ok := i["ip6-subnet"]; ok {
			v := flattenEmailfilterBlockAllowListEntriesIp6Subnet(i["ip6-subnet"], d, pre_append)
			tmp["ip6_subnet"] = fortiAPISubPartPatch(v, "EmailfilterBlockAllowList-Entries-Ip6Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenEmailfilterBlockAllowListEntriesPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "EmailfilterBlockAllowList-Entries-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_type"
		if _, ok := i["pattern-type"]; ok {
			v := flattenEmailfilterBlockAllowListEntriesPatternType(i["pattern-type"], d, pre_append)
			tmp["pattern_type"] = fortiAPISubPartPatch(v, "EmailfilterBlockAllowList-Entries-PatternType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenEmailfilterBlockAllowListEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "EmailfilterBlockAllowList-Entries-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenEmailfilterBlockAllowListEntriesType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "EmailfilterBlockAllowList-Entries-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email_pattern"
		if _, ok := i["email-pattern"]; ok {
			v := flattenEmailfilterBlockAllowListEntriesEmailPattern(i["email-pattern"], d, pre_append)
			tmp["email_pattern"] = fortiAPISubPartPatch(v, "EmailfilterBlockAllowList-Entries-EmailPattern")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenEmailfilterBlockAllowListEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesIp4Subnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterBlockAllowListEntriesIp6Subnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesPatternType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesEmailPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEmailfilterBlockAllowList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenEmailfilterBlockAllowListComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "EmailfilterBlockAllowList-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenEmailfilterBlockAllowListEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "EmailfilterBlockAllowList-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenEmailfilterBlockAllowListEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "EmailfilterBlockAllowList-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenEmailfilterBlockAllowListId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "EmailfilterBlockAllowList-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenEmailfilterBlockAllowListName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "EmailfilterBlockAllowList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterBlockAllowListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEmailfilterBlockAllowListComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandEmailfilterBlockAllowListEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-type"], _ = expandEmailfilterBlockAllowListEntriesAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandEmailfilterBlockAllowListEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip4_subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip4-subnet"], _ = expandEmailfilterBlockAllowListEntriesIp4Subnet(d, i["ip4_subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6-subnet"], _ = expandEmailfilterBlockAllowListEntriesIp6Subnet(d, i["ip6_subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandEmailfilterBlockAllowListEntriesPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern-type"], _ = expandEmailfilterBlockAllowListEntriesPatternType(d, i["pattern_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandEmailfilterBlockAllowListEntriesStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandEmailfilterBlockAllowListEntriesType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email_pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["email-pattern"], _ = expandEmailfilterBlockAllowListEntriesEmailPattern(d, i["email_pattern"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandEmailfilterBlockAllowListEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesIp4Subnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandEmailfilterBlockAllowListEntriesIp6Subnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesPatternType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesEmailPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterBlockAllowList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandEmailfilterBlockAllowListComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandEmailfilterBlockAllowListEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandEmailfilterBlockAllowListId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandEmailfilterBlockAllowListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
