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

func resourceFirewallProxyAddress6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProxyAddress6Create,
		Read:   resourceFirewallProxyAddress6Read,
		Update: resourceFirewallProxyAddress6Update,
		Delete: resourceFirewallProxyAddress6Delete,

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
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"case_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"header_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"case_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"header_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"host_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"post_arg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"referrer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tagging": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ua": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ua_max_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ua_min_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
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

func resourceFirewallProxyAddress6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProxyAddress6(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallProxyAddress6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallProxyAddress6(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallProxyAddress6(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallProxyAddress6 resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallProxyAddress6(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallProxyAddress6 resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallProxyAddress6Read(d, m)
}

func resourceFirewallProxyAddress6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProxyAddress6(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProxyAddress6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallProxyAddress6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProxyAddress6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallProxyAddress6Read(d, m)
}

func resourceFirewallProxyAddress6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallProxyAddress6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProxyAddress6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProxyAddress6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallProxyAddress6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallProxyAddress6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProxyAddress6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProxyAddress6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProxyAddress6Application(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddress6CaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6Category(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProxyAddress6Color(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6Comment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6Header(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6HeaderGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {
			v := flattenFirewallProxyAddress6HeaderGroupCaseSensitivity(i["case-sensitivity"], d, pre_append)
			tmp["case_sensitivity"] = fortiAPISubPartPatch(v, "FirewallProxyAddress6-HeaderGroup-CaseSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenFirewallProxyAddress6HeaderGroupHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "FirewallProxyAddress6-HeaderGroup-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenFirewallProxyAddress6HeaderGroupHeaderName(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "FirewallProxyAddress6-HeaderGroup-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallProxyAddress6HeaderGroupId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallProxyAddress6-HeaderGroup-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallProxyAddress6HeaderGroupCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6HeaderGroupHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6HeaderGroupHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6HeaderGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6HeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6Host(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddress6HostRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6Method(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddress6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6Path(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6PostArg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6Query(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6Referrer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6Tagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenFirewallProxyAddress6TaggingCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "FirewallProxyAddress6-Tagging-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallProxyAddress6TaggingName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallProxyAddress6-Tagging-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenFirewallProxyAddress6TaggingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "FirewallProxyAddress6-Tagging-Tags")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallProxyAddress6TaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddress6TaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6TaggingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddress6Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6Ua(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddress6UaMaxVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6UaMinVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddress6UrlList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddress6Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallProxyAddress6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("application", flattenFirewallProxyAddress6Application(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "FirewallProxyAddress6-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("case_sensitivity", flattenFirewallProxyAddress6CaseSensitivity(o["case-sensitivity"], d, "case_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-sensitivity"], "FirewallProxyAddress6-CaseSensitivity"); ok {
			if err = d.Set("case_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading case_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_sensitivity: %v", err)
		}
	}

	if err = d.Set("category", flattenFirewallProxyAddress6Category(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "FirewallProxyAddress6-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallProxyAddress6Color(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallProxyAddress6-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallProxyAddress6Comment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallProxyAddress6-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("header", flattenFirewallProxyAddress6Header(o["header"], d, "header")); err != nil {
		if vv, ok := fortiAPIPatch(o["header"], "FirewallProxyAddress6-Header"); ok {
			if err = d.Set("header", vv); err != nil {
				return fmt.Errorf("Error reading header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("header_group", flattenFirewallProxyAddress6HeaderGroup(o["header-group"], d, "header_group")); err != nil {
			if vv, ok := fortiAPIPatch(o["header-group"], "FirewallProxyAddress6-HeaderGroup"); ok {
				if err = d.Set("header_group", vv); err != nil {
					return fmt.Errorf("Error reading header_group: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading header_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("header_group"); ok {
			if err = d.Set("header_group", flattenFirewallProxyAddress6HeaderGroup(o["header-group"], d, "header_group")); err != nil {
				if vv, ok := fortiAPIPatch(o["header-group"], "FirewallProxyAddress6-HeaderGroup"); ok {
					if err = d.Set("header_group", vv); err != nil {
						return fmt.Errorf("Error reading header_group: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading header_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("header_name", flattenFirewallProxyAddress6HeaderName(o["header-name"], d, "header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-name"], "FirewallProxyAddress6-HeaderName"); ok {
			if err = d.Set("header_name", vv); err != nil {
				return fmt.Errorf("Error reading header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("host", flattenFirewallProxyAddress6Host(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "FirewallProxyAddress6-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("host_regex", flattenFirewallProxyAddress6HostRegex(o["host-regex"], d, "host_regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-regex"], "FirewallProxyAddress6-HostRegex"); ok {
			if err = d.Set("host_regex", vv); err != nil {
				return fmt.Errorf("Error reading host_regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_regex: %v", err)
		}
	}

	if err = d.Set("method", flattenFirewallProxyAddress6Method(o["method"], d, "method")); err != nil {
		if vv, ok := fortiAPIPatch(o["method"], "FirewallProxyAddress6-Method"); ok {
			if err = d.Set("method", vv); err != nil {
				return fmt.Errorf("Error reading method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallProxyAddress6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallProxyAddress6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("path", flattenFirewallProxyAddress6Path(o["path"], d, "path")); err != nil {
		if vv, ok := fortiAPIPatch(o["path"], "FirewallProxyAddress6-Path"); ok {
			if err = d.Set("path", vv); err != nil {
				return fmt.Errorf("Error reading path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading path: %v", err)
		}
	}

	if err = d.Set("post_arg", flattenFirewallProxyAddress6PostArg(o["post-arg"], d, "post_arg")); err != nil {
		if vv, ok := fortiAPIPatch(o["post-arg"], "FirewallProxyAddress6-PostArg"); ok {
			if err = d.Set("post_arg", vv); err != nil {
				return fmt.Errorf("Error reading post_arg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading post_arg: %v", err)
		}
	}

	if err = d.Set("query", flattenFirewallProxyAddress6Query(o["query"], d, "query")); err != nil {
		if vv, ok := fortiAPIPatch(o["query"], "FirewallProxyAddress6-Query"); ok {
			if err = d.Set("query", vv); err != nil {
				return fmt.Errorf("Error reading query: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query: %v", err)
		}
	}

	if err = d.Set("referrer", flattenFirewallProxyAddress6Referrer(o["referrer"], d, "referrer")); err != nil {
		if vv, ok := fortiAPIPatch(o["referrer"], "FirewallProxyAddress6-Referrer"); ok {
			if err = d.Set("referrer", vv); err != nil {
				return fmt.Errorf("Error reading referrer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading referrer: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tagging", flattenFirewallProxyAddress6Tagging(o["tagging"], d, "tagging")); err != nil {
			if vv, ok := fortiAPIPatch(o["tagging"], "FirewallProxyAddress6-Tagging"); ok {
				if err = d.Set("tagging", vv); err != nil {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tagging: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tagging"); ok {
			if err = d.Set("tagging", flattenFirewallProxyAddress6Tagging(o["tagging"], d, "tagging")); err != nil {
				if vv, ok := fortiAPIPatch(o["tagging"], "FirewallProxyAddress6-Tagging"); ok {
					if err = d.Set("tagging", vv); err != nil {
						return fmt.Errorf("Error reading tagging: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenFirewallProxyAddress6Type(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallProxyAddress6-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ua", flattenFirewallProxyAddress6Ua(o["ua"], d, "ua")); err != nil {
		if vv, ok := fortiAPIPatch(o["ua"], "FirewallProxyAddress6-Ua"); ok {
			if err = d.Set("ua", vv); err != nil {
				return fmt.Errorf("Error reading ua: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ua: %v", err)
		}
	}

	if err = d.Set("ua_max_ver", flattenFirewallProxyAddress6UaMaxVer(o["ua-max-ver"], d, "ua_max_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["ua-max-ver"], "FirewallProxyAddress6-UaMaxVer"); ok {
			if err = d.Set("ua_max_ver", vv); err != nil {
				return fmt.Errorf("Error reading ua_max_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ua_max_ver: %v", err)
		}
	}

	if err = d.Set("ua_min_ver", flattenFirewallProxyAddress6UaMinVer(o["ua-min-ver"], d, "ua_min_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["ua-min-ver"], "FirewallProxyAddress6-UaMinVer"); ok {
			if err = d.Set("ua_min_ver", vv); err != nil {
				return fmt.Errorf("Error reading ua_min_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ua_min_ver: %v", err)
		}
	}

	if err = d.Set("url_list", flattenFirewallProxyAddress6UrlList(o["url-list"], d, "url_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-list"], "FirewallProxyAddress6-UrlList"); ok {
			if err = d.Set("url_list", vv); err != nil {
				return fmt.Errorf("Error reading url_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_list: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallProxyAddress6Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallProxyAddress6-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenFirewallProxyAddress6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProxyAddress6Application(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddress6CaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6Category(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddress6Color(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6Comment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6Header(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6HeaderGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitivity"], _ = expandFirewallProxyAddress6HeaderGroupCaseSensitivity(d, i["case_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandFirewallProxyAddress6HeaderGroupHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandFirewallProxyAddress6HeaderGroupHeaderName(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallProxyAddress6HeaderGroupId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallProxyAddress6HeaderGroupCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6HeaderGroupHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6HeaderGroupHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6HeaderGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6HeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6Host(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddress6HostRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6Method(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddress6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6Path(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6PostArg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6Query(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6Referrer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6Tagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandFirewallProxyAddress6TaggingCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallProxyAddress6TaggingName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandFirewallProxyAddress6TaggingTags(d, i["tags"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallProxyAddress6TaggingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddress6TaggingName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6TaggingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddress6Type(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6Ua(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddress6UaMaxVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6UaMinVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddress6UrlList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddress6Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProxyAddress6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandFirewallProxyAddress6Application(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("case_sensitivity"); ok || d.HasChange("case_sensitivity") {
		t, err := expandFirewallProxyAddress6CaseSensitivity(d, v, "case_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandFirewallProxyAddress6Category(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallProxyAddress6Color(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallProxyAddress6Comment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok || d.HasChange("header") {
		t, err := expandFirewallProxyAddress6Header(d, v, "header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("header_group"); ok || d.HasChange("header_group") {
		t, err := expandFirewallProxyAddress6HeaderGroup(d, v, "header_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-group"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok || d.HasChange("header_name") {
		t, err := expandFirewallProxyAddress6HeaderName(d, v, "header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandFirewallProxyAddress6Host(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("host_regex"); ok || d.HasChange("host_regex") {
		t, err := expandFirewallProxyAddress6HostRegex(d, v, "host_regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-regex"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandFirewallProxyAddress6Method(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallProxyAddress6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("path"); ok || d.HasChange("path") {
		t, err := expandFirewallProxyAddress6Path(d, v, "path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["path"] = t
		}
	}

	if v, ok := d.GetOk("post_arg"); ok || d.HasChange("post_arg") {
		t, err := expandFirewallProxyAddress6PostArg(d, v, "post_arg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["post-arg"] = t
		}
	}

	if v, ok := d.GetOk("query"); ok || d.HasChange("query") {
		t, err := expandFirewallProxyAddress6Query(d, v, "query")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query"] = t
		}
	}

	if v, ok := d.GetOk("referrer"); ok || d.HasChange("referrer") {
		t, err := expandFirewallProxyAddress6Referrer(d, v, "referrer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["referrer"] = t
		}
	}

	if v, ok := d.GetOk("tagging"); ok || d.HasChange("tagging") {
		t, err := expandFirewallProxyAddress6Tagging(d, v, "tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tagging"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallProxyAddress6Type(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("ua"); ok || d.HasChange("ua") {
		t, err := expandFirewallProxyAddress6Ua(d, v, "ua")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ua"] = t
		}
	}

	if v, ok := d.GetOk("ua_max_ver"); ok || d.HasChange("ua_max_ver") {
		t, err := expandFirewallProxyAddress6UaMaxVer(d, v, "ua_max_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ua-max-ver"] = t
		}
	}

	if v, ok := d.GetOk("ua_min_ver"); ok || d.HasChange("ua_min_ver") {
		t, err := expandFirewallProxyAddress6UaMinVer(d, v, "ua_min_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ua-min-ver"] = t
		}
	}

	if v, ok := d.GetOk("url_list"); ok || d.HasChange("url_list") {
		t, err := expandFirewallProxyAddress6UrlList(d, v, "url_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-list"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallProxyAddress6Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
