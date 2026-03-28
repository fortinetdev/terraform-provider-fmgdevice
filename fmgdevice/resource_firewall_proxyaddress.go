// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure web proxy address.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProxyAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProxyAddressCreate,
		Read:   resourceFirewallProxyAddressRead,
		Update: resourceFirewallProxyAddressUpdate,
		Delete: resourceFirewallProxyAddressDelete,

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
			"_image_base64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
				Elem:     &schema.Schema{Type: schema.TypeString},
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
							Computed: true,
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
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"post_arg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url_list": &schema.Schema{
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

func resourceFirewallProxyAddressCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProxyAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallProxyAddress resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallProxyAddress(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallProxyAddress(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallProxyAddress resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallProxyAddress(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallProxyAddress resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallProxyAddressRead(d, m)
}

func resourceFirewallProxyAddressUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProxyAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProxyAddress resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallProxyAddress(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProxyAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallProxyAddressRead(d, m)
}

func resourceFirewallProxyAddressDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallProxyAddress(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProxyAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProxyAddressRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallProxyAddress(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallProxyAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProxyAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProxyAddress resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProxyAddressImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddressCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddressColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressHeaderGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallProxyAddressHeaderGroupCaseSensitivity(i["case-sensitivity"], d, pre_append)
			tmp["case_sensitivity"] = fortiAPISubPartPatch(v, "FirewallProxyAddress-HeaderGroup-CaseSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenFirewallProxyAddressHeaderGroupHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "FirewallProxyAddress-HeaderGroup-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenFirewallProxyAddressHeaderGroupHeaderName(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "FirewallProxyAddress-HeaderGroup-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallProxyAddressHeaderGroupId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallProxyAddress-HeaderGroup-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallProxyAddressHeaderGroupCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressHeaderGroupHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressHeaderGroupHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressHeaderGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddressHostRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressQuery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressReferrer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressTagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallProxyAddressTaggingCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "FirewallProxyAddress-Tagging-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallProxyAddressTaggingName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallProxyAddress-Tagging-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenFirewallProxyAddressTaggingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "FirewallProxyAddress-Tagging-Tags")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallProxyAddressTaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddressTaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressTaggingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressUa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProxyAddressUaMaxVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressUaMinVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressPostArg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProxyAddressUrlList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallProxyAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_image_base64", flattenFirewallProxyAddressImageBase64(o["_image-base64"], d, "_image_base64")); err != nil {
		if vv, ok := fortiAPIPatch(o["_image-base64"], "FirewallProxyAddress-ImageBase64"); ok {
			if err = d.Set("_image_base64", vv); err != nil {
				return fmt.Errorf("Error reading _image_base64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _image_base64: %v", err)
		}
	}

	if err = d.Set("application", flattenFirewallProxyAddressApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "FirewallProxyAddress-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("case_sensitivity", flattenFirewallProxyAddressCaseSensitivity(o["case-sensitivity"], d, "case_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-sensitivity"], "FirewallProxyAddress-CaseSensitivity"); ok {
			if err = d.Set("case_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading case_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_sensitivity: %v", err)
		}
	}

	if err = d.Set("category", flattenFirewallProxyAddressCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "FirewallProxyAddress-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallProxyAddressColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallProxyAddress-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallProxyAddressComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallProxyAddress-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("header", flattenFirewallProxyAddressHeader(o["header"], d, "header")); err != nil {
		if vv, ok := fortiAPIPatch(o["header"], "FirewallProxyAddress-Header"); ok {
			if err = d.Set("header", vv); err != nil {
				return fmt.Errorf("Error reading header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("header_group", flattenFirewallProxyAddressHeaderGroup(o["header-group"], d, "header_group")); err != nil {
			if vv, ok := fortiAPIPatch(o["header-group"], "FirewallProxyAddress-HeaderGroup"); ok {
				if err = d.Set("header_group", vv); err != nil {
					return fmt.Errorf("Error reading header_group: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading header_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("header_group"); ok {
			if err = d.Set("header_group", flattenFirewallProxyAddressHeaderGroup(o["header-group"], d, "header_group")); err != nil {
				if vv, ok := fortiAPIPatch(o["header-group"], "FirewallProxyAddress-HeaderGroup"); ok {
					if err = d.Set("header_group", vv); err != nil {
						return fmt.Errorf("Error reading header_group: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading header_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("header_name", flattenFirewallProxyAddressHeaderName(o["header-name"], d, "header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-name"], "FirewallProxyAddress-HeaderName"); ok {
			if err = d.Set("header_name", vv); err != nil {
				return fmt.Errorf("Error reading header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("host", flattenFirewallProxyAddressHost(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "FirewallProxyAddress-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("host_regex", flattenFirewallProxyAddressHostRegex(o["host-regex"], d, "host_regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-regex"], "FirewallProxyAddress-HostRegex"); ok {
			if err = d.Set("host_regex", vv); err != nil {
				return fmt.Errorf("Error reading host_regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_regex: %v", err)
		}
	}

	if err = d.Set("method", flattenFirewallProxyAddressMethod(o["method"], d, "method")); err != nil {
		if vv, ok := fortiAPIPatch(o["method"], "FirewallProxyAddress-Method"); ok {
			if err = d.Set("method", vv); err != nil {
				return fmt.Errorf("Error reading method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallProxyAddressName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallProxyAddress-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("path", flattenFirewallProxyAddressPath(o["path"], d, "path")); err != nil {
		if vv, ok := fortiAPIPatch(o["path"], "FirewallProxyAddress-Path"); ok {
			if err = d.Set("path", vv); err != nil {
				return fmt.Errorf("Error reading path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading path: %v", err)
		}
	}

	if err = d.Set("query", flattenFirewallProxyAddressQuery(o["query"], d, "query")); err != nil {
		if vv, ok := fortiAPIPatch(o["query"], "FirewallProxyAddress-Query"); ok {
			if err = d.Set("query", vv); err != nil {
				return fmt.Errorf("Error reading query: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query: %v", err)
		}
	}

	if err = d.Set("referrer", flattenFirewallProxyAddressReferrer(o["referrer"], d, "referrer")); err != nil {
		if vv, ok := fortiAPIPatch(o["referrer"], "FirewallProxyAddress-Referrer"); ok {
			if err = d.Set("referrer", vv); err != nil {
				return fmt.Errorf("Error reading referrer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading referrer: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tagging", flattenFirewallProxyAddressTagging(o["tagging"], d, "tagging")); err != nil {
			if vv, ok := fortiAPIPatch(o["tagging"], "FirewallProxyAddress-Tagging"); ok {
				if err = d.Set("tagging", vv); err != nil {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tagging: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tagging"); ok {
			if err = d.Set("tagging", flattenFirewallProxyAddressTagging(o["tagging"], d, "tagging")); err != nil {
				if vv, ok := fortiAPIPatch(o["tagging"], "FirewallProxyAddress-Tagging"); ok {
					if err = d.Set("tagging", vv); err != nil {
						return fmt.Errorf("Error reading tagging: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenFirewallProxyAddressType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallProxyAddress-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ua", flattenFirewallProxyAddressUa(o["ua"], d, "ua")); err != nil {
		if vv, ok := fortiAPIPatch(o["ua"], "FirewallProxyAddress-Ua"); ok {
			if err = d.Set("ua", vv); err != nil {
				return fmt.Errorf("Error reading ua: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ua: %v", err)
		}
	}

	if err = d.Set("ua_max_ver", flattenFirewallProxyAddressUaMaxVer(o["ua-max-ver"], d, "ua_max_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["ua-max-ver"], "FirewallProxyAddress-UaMaxVer"); ok {
			if err = d.Set("ua_max_ver", vv); err != nil {
				return fmt.Errorf("Error reading ua_max_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ua_max_ver: %v", err)
		}
	}

	if err = d.Set("ua_min_ver", flattenFirewallProxyAddressUaMinVer(o["ua-min-ver"], d, "ua_min_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["ua-min-ver"], "FirewallProxyAddress-UaMinVer"); ok {
			if err = d.Set("ua_min_ver", vv); err != nil {
				return fmt.Errorf("Error reading ua_min_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ua_min_ver: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallProxyAddressUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallProxyAddress-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("post_arg", flattenFirewallProxyAddressPostArg(o["post-arg"], d, "post_arg")); err != nil {
		if vv, ok := fortiAPIPatch(o["post-arg"], "FirewallProxyAddress-PostArg"); ok {
			if err = d.Set("post_arg", vv); err != nil {
				return fmt.Errorf("Error reading post_arg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading post_arg: %v", err)
		}
	}

	if err = d.Set("url_list", flattenFirewallProxyAddressUrlList(o["url-list"], d, "url_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-list"], "FirewallProxyAddress-UrlList"); ok {
			if err = d.Set("url_list", vv); err != nil {
				return fmt.Errorf("Error reading url_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_list: %v", err)
		}
	}

	return nil
}

func flattenFirewallProxyAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProxyAddressImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddressCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddressColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHeaderGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["case-sensitivity"], _ = expandFirewallProxyAddressHeaderGroupCaseSensitivity(d, i["case_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandFirewallProxyAddressHeaderGroupHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandFirewallProxyAddressHeaderGroupHeaderName(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallProxyAddressHeaderGroupId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallProxyAddressHeaderGroupCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHeaderGroupHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHeaderGroupHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHeaderGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddressHostRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddressName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressQuery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressReferrer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressTagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["category"], _ = expandFirewallProxyAddressTaggingCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallProxyAddressTaggingName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandFirewallProxyAddressTaggingTags(d, i["tags"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallProxyAddressTaggingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddressTaggingName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressTaggingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressUa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProxyAddressUaMaxVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressUaMinVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressPostArg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressUrlList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallProxyAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_image_base64"); ok || d.HasChange("_image_base64") {
		t, err := expandFirewallProxyAddressImageBase64(d, v, "_image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_image-base64"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandFirewallProxyAddressApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("case_sensitivity"); ok || d.HasChange("case_sensitivity") {
		t, err := expandFirewallProxyAddressCaseSensitivity(d, v, "case_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandFirewallProxyAddressCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallProxyAddressColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallProxyAddressComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok || d.HasChange("header") {
		t, err := expandFirewallProxyAddressHeader(d, v, "header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("header_group"); ok || d.HasChange("header_group") {
		t, err := expandFirewallProxyAddressHeaderGroup(d, v, "header_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-group"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok || d.HasChange("header_name") {
		t, err := expandFirewallProxyAddressHeaderName(d, v, "header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandFirewallProxyAddressHost(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("host_regex"); ok || d.HasChange("host_regex") {
		t, err := expandFirewallProxyAddressHostRegex(d, v, "host_regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-regex"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandFirewallProxyAddressMethod(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallProxyAddressName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("path"); ok || d.HasChange("path") {
		t, err := expandFirewallProxyAddressPath(d, v, "path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["path"] = t
		}
	}

	if v, ok := d.GetOk("query"); ok || d.HasChange("query") {
		t, err := expandFirewallProxyAddressQuery(d, v, "query")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query"] = t
		}
	}

	if v, ok := d.GetOk("referrer"); ok || d.HasChange("referrer") {
		t, err := expandFirewallProxyAddressReferrer(d, v, "referrer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["referrer"] = t
		}
	}

	if v, ok := d.GetOk("tagging"); ok || d.HasChange("tagging") {
		t, err := expandFirewallProxyAddressTagging(d, v, "tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tagging"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallProxyAddressType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("ua"); ok || d.HasChange("ua") {
		t, err := expandFirewallProxyAddressUa(d, v, "ua")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ua"] = t
		}
	}

	if v, ok := d.GetOk("ua_max_ver"); ok || d.HasChange("ua_max_ver") {
		t, err := expandFirewallProxyAddressUaMaxVer(d, v, "ua_max_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ua-max-ver"] = t
		}
	}

	if v, ok := d.GetOk("ua_min_ver"); ok || d.HasChange("ua_min_ver") {
		t, err := expandFirewallProxyAddressUaMinVer(d, v, "ua_min_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ua-min-ver"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallProxyAddressUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("post_arg"); ok || d.HasChange("post_arg") {
		t, err := expandFirewallProxyAddressPostArg(d, v, "post_arg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["post-arg"] = t
		}
	}

	if v, ok := d.GetOk("url_list"); ok || d.HasChange("url_list") {
		t, err := expandFirewallProxyAddressUrlList(d, v, "url_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-list"] = t
		}
	}

	return &obj, nil
}
