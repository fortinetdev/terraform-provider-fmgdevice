// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AntiSpam IP trust.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterIptrust() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterIptrustCreate,
		Read:   resourceEmailfilterIptrustRead,
		Update: resourceEmailfilterIptrustUpdate,
		Delete: resourceEmailfilterIptrustDelete,

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
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"status": &schema.Schema{
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

func resourceEmailfilterIptrustCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEmailfilterIptrust(d)
	if err != nil {
		return fmt.Errorf("Error creating EmailfilterIptrust resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadEmailfilterIptrust(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateEmailfilterIptrust(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating EmailfilterIptrust resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateEmailfilterIptrust(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating EmailfilterIptrust resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterIptrustRead(d, m)
}

func resourceEmailfilterIptrustUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEmailfilterIptrust(d)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterIptrust resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEmailfilterIptrust(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterIptrust resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterIptrustRead(d, m)
}

func resourceEmailfilterIptrustDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteEmailfilterIptrust(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterIptrust resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterIptrustRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadEmailfilterIptrust(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EmailfilterIptrust resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterIptrust(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterIptrust resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterIptrustComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterIptrustEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {
			v := flattenEmailfilterIptrustEntriesAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "EmailfilterIptrust-Entries-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenEmailfilterIptrustEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "EmailfilterIptrust-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip4_subnet"
		if _, ok := i["ip4-subnet"]; ok {
			v := flattenEmailfilterIptrustEntriesIp4Subnet(i["ip4-subnet"], d, pre_append)
			tmp["ip4_subnet"] = fortiAPISubPartPatch(v, "EmailfilterIptrust-Entries-Ip4Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_subnet"
		if _, ok := i["ip6-subnet"]; ok {
			v := flattenEmailfilterIptrustEntriesIp6Subnet(i["ip6-subnet"], d, pre_append)
			tmp["ip6_subnet"] = fortiAPISubPartPatch(v, "EmailfilterIptrust-Entries-Ip6Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenEmailfilterIptrustEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "EmailfilterIptrust-Entries-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenEmailfilterIptrustEntriesAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterIptrustEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterIptrustEntriesIp4Subnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterIptrustEntriesIp6Subnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterIptrustEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterIptrustId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterIptrustName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEmailfilterIptrust(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenEmailfilterIptrustComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "EmailfilterIptrust-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenEmailfilterIptrustEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "EmailfilterIptrust-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenEmailfilterIptrustEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "EmailfilterIptrust-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenEmailfilterIptrustId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "EmailfilterIptrust-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenEmailfilterIptrustName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "EmailfilterIptrust-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterIptrustFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEmailfilterIptrustComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterIptrustEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-type"], _ = expandEmailfilterIptrustEntriesAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandEmailfilterIptrustEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip4_subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip4-subnet"], _ = expandEmailfilterIptrustEntriesIp4Subnet(d, i["ip4_subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6-subnet"], _ = expandEmailfilterIptrustEntriesIp6Subnet(d, i["ip6_subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandEmailfilterIptrustEntriesStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandEmailfilterIptrustEntriesAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterIptrustEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterIptrustEntriesIp4Subnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandEmailfilterIptrustEntriesIp6Subnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterIptrustEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterIptrustId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterIptrustName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterIptrust(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandEmailfilterIptrustComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandEmailfilterIptrustEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandEmailfilterIptrustId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandEmailfilterIptrustName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
