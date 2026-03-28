// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> FortiGuard DNS Filter settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDnsfilterProfileFtgdDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsfilterProfileFtgdDnsUpdate,
		Read:   resourceDnsfilterProfileFtgdDnsRead,
		Update: resourceDnsfilterProfileFtgdDnsUpdate,
		Delete: resourceDnsfilterProfileFtgdDnsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"options": &schema.Schema{
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

func resourceDnsfilterProfileFtgdDnsUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectDnsfilterProfileFtgdDns(d)
	if err != nil {
		return fmt.Errorf("Error updating DnsfilterProfileFtgdDns resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDnsfilterProfileFtgdDns(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DnsfilterProfileFtgdDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DnsfilterProfileFtgdDns")

	return resourceDnsfilterProfileFtgdDnsRead(d, m)
}

func resourceDnsfilterProfileFtgdDnsDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteDnsfilterProfileFtgdDns(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DnsfilterProfileFtgdDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDnsfilterProfileFtgdDnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadDnsfilterProfileFtgdDns(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DnsfilterProfileFtgdDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDnsfilterProfileFtgdDns(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DnsfilterProfileFtgdDns resource from API: %v", err)
	}
	return nil
}

func flattenDnsfilterProfileFtgdDnsFilters2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenDnsfilterProfileFtgdDnsFiltersAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "DnsfilterProfileFtgdDns-Filters-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenDnsfilterProfileFtgdDnsFiltersCategory2edl(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "DnsfilterProfileFtgdDns-Filters-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenDnsfilterProfileFtgdDnsFiltersId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "DnsfilterProfileFtgdDns-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenDnsfilterProfileFtgdDnsFiltersLog2edl(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "DnsfilterProfileFtgdDns-Filters-Log")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDnsfilterProfileFtgdDnsFiltersAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileFtgdDnsFiltersCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDnsfilterProfileFtgdDnsFiltersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileFtgdDnsFiltersLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileFtgdDnsOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectDnsfilterProfileFtgdDns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("filters", flattenDnsfilterProfileFtgdDnsFilters2edl(o["filters"], d, "filters")); err != nil {
			if vv, ok := fortiAPIPatch(o["filters"], "DnsfilterProfileFtgdDns-Filters"); ok {
				if err = d.Set("filters", vv); err != nil {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filters"); ok {
			if err = d.Set("filters", flattenDnsfilterProfileFtgdDnsFilters2edl(o["filters"], d, "filters")); err != nil {
				if vv, ok := fortiAPIPatch(o["filters"], "DnsfilterProfileFtgdDns-Filters"); ok {
					if err = d.Set("filters", vv); err != nil {
						return fmt.Errorf("Error reading filters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			}
		}
	}

	if err = d.Set("options", flattenDnsfilterProfileFtgdDnsOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "DnsfilterProfileFtgdDns-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	return nil
}

func flattenDnsfilterProfileFtgdDnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDnsfilterProfileFtgdDnsFilters2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandDnsfilterProfileFtgdDnsFiltersAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandDnsfilterProfileFtgdDnsFiltersCategory2edl(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandDnsfilterProfileFtgdDnsFiltersId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandDnsfilterProfileFtgdDnsFiltersLog2edl(d, i["log"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDnsfilterProfileFtgdDnsFiltersAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDnsFiltersCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDnsfilterProfileFtgdDnsFiltersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDnsFiltersLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDnsOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectDnsfilterProfileFtgdDns(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("filters"); ok || d.HasChange("filters") {
		t, err := expandDnsfilterProfileFtgdDnsFilters2edl(d, v, "filters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filters"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandDnsfilterProfileFtgdDnsOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	return &obj, nil
}
