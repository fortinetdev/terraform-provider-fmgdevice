// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure DNS domain filter profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDnsfilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsfilterProfileCreate,
		Read:   resourceDnsfilterProfileRead,
		Update: resourceDnsfilterProfileUpdate,
		Delete: resourceDnsfilterProfileDelete,

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
			"block_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_botnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_translation": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dst": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dst6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"netmask": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"src": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"domain_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_filter_table": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"external_ip_blocklist": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ftgd_dns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"log_all_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"redirect_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirect_portal6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"safe_search": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdns_domain_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdns_ftgd_err_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strip_ech": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transparent_dns_database": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"youtube_restrict": &schema.Schema{
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

func resourceDnsfilterProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDnsfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating DnsfilterProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDnsfilterProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDnsfilterProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DnsfilterProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDnsfilterProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DnsfilterProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceDnsfilterProfileRead(d, m)
}

func resourceDnsfilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDnsfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating DnsfilterProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDnsfilterProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DnsfilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDnsfilterProfileRead(d, m)
}

func resourceDnsfilterProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteDnsfilterProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DnsfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDnsfilterProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDnsfilterProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DnsfilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDnsfilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DnsfilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenDnsfilterProfileBlockAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileBlockBotnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslation(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenDnsfilterProfileDnsTranslationAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "DnsfilterProfile-DnsTranslation-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			v := flattenDnsfilterProfileDnsTranslationDst(i["dst"], d, pre_append)
			tmp["dst"] = fortiAPISubPartPatch(v, "DnsfilterProfile-DnsTranslation-Dst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := i["dst6"]; ok {
			v := flattenDnsfilterProfileDnsTranslationDst6(i["dst6"], d, pre_append)
			tmp["dst6"] = fortiAPISubPartPatch(v, "DnsfilterProfile-DnsTranslation-Dst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenDnsfilterProfileDnsTranslationId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "DnsfilterProfile-DnsTranslation-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "netmask"
		if _, ok := i["netmask"]; ok {
			v := flattenDnsfilterProfileDnsTranslationNetmask(i["netmask"], d, pre_append)
			tmp["netmask"] = fortiAPISubPartPatch(v, "DnsfilterProfile-DnsTranslation-Netmask")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenDnsfilterProfileDnsTranslationPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "DnsfilterProfile-DnsTranslation-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := i["src"]; ok {
			v := flattenDnsfilterProfileDnsTranslationSrc(i["src"], d, pre_append)
			tmp["src"] = fortiAPISubPartPatch(v, "DnsfilterProfile-DnsTranslation-Src")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := i["src6"]; ok {
			v := flattenDnsfilterProfileDnsTranslationSrc6(i["src6"], d, pre_append)
			tmp["src6"] = fortiAPISubPartPatch(v, "DnsfilterProfile-DnsTranslation-Src6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenDnsfilterProfileDnsTranslationStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "DnsfilterProfile-DnsTranslation-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDnsfilterProfileDnsTranslationAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationDst6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationSrc6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDomainFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "domain_filter_table"
	if _, ok := i["domain-filter-table"]; ok {
		result["domain_filter_table"] = flattenDnsfilterProfileDomainFilterDomainFilterTable(i["domain-filter-table"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenDnsfilterProfileDomainFilterDomainFilterTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDnsfilterProfileExternalIpBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDnsfilterProfileFtgdDns(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenDnsfilterProfileFtgdDnsFilters(i["filters"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenDnsfilterProfileFtgdDnsOptions(i["options"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenDnsfilterProfileFtgdDnsFilters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenDnsfilterProfileFtgdDnsFiltersAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "DnsfilterProfileFtgdDns-Filters-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenDnsfilterProfileFtgdDnsFiltersCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "DnsfilterProfileFtgdDns-Filters-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenDnsfilterProfileFtgdDnsFiltersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "DnsfilterProfileFtgdDns-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenDnsfilterProfileFtgdDnsFiltersLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "DnsfilterProfileFtgdDns-Filters-Log")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDnsfilterProfileFtgdDnsFiltersAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileFtgdDnsFiltersCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDnsfilterProfileFtgdDnsFiltersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileFtgdDnsFiltersLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileFtgdDnsOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDnsfilterProfileLogAllDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileRedirectPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileRedirectPortal6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileSafeSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileSdnsDomainLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileSdnsFtgdErrLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileStripEch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileTransparentDnsDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDnsfilterProfileYoutubeRestrict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDnsfilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("block_action", flattenDnsfilterProfileBlockAction(o["block-action"], d, "block_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-action"], "DnsfilterProfile-BlockAction"); ok {
			if err = d.Set("block_action", vv); err != nil {
				return fmt.Errorf("Error reading block_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_action: %v", err)
		}
	}

	if err = d.Set("block_botnet", flattenDnsfilterProfileBlockBotnet(o["block-botnet"], d, "block_botnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-botnet"], "DnsfilterProfile-BlockBotnet"); ok {
			if err = d.Set("block_botnet", vv); err != nil {
				return fmt.Errorf("Error reading block_botnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_botnet: %v", err)
		}
	}

	if err = d.Set("comment", flattenDnsfilterProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "DnsfilterProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dns_translation", flattenDnsfilterProfileDnsTranslation(o["dns-translation"], d, "dns_translation")); err != nil {
			if vv, ok := fortiAPIPatch(o["dns-translation"], "DnsfilterProfile-DnsTranslation"); ok {
				if err = d.Set("dns_translation", vv); err != nil {
					return fmt.Errorf("Error reading dns_translation: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dns_translation: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dns_translation"); ok {
			if err = d.Set("dns_translation", flattenDnsfilterProfileDnsTranslation(o["dns-translation"], d, "dns_translation")); err != nil {
				if vv, ok := fortiAPIPatch(o["dns-translation"], "DnsfilterProfile-DnsTranslation"); ok {
					if err = d.Set("dns_translation", vv); err != nil {
						return fmt.Errorf("Error reading dns_translation: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dns_translation: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("domain_filter", flattenDnsfilterProfileDomainFilter(o["domain-filter"], d, "domain_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["domain-filter"], "DnsfilterProfile-DomainFilter"); ok {
				if err = d.Set("domain_filter", vv); err != nil {
					return fmt.Errorf("Error reading domain_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading domain_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("domain_filter"); ok {
			if err = d.Set("domain_filter", flattenDnsfilterProfileDomainFilter(o["domain-filter"], d, "domain_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["domain-filter"], "DnsfilterProfile-DomainFilter"); ok {
					if err = d.Set("domain_filter", vv); err != nil {
						return fmt.Errorf("Error reading domain_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading domain_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("external_ip_blocklist", flattenDnsfilterProfileExternalIpBlocklist(o["external-ip-blocklist"], d, "external_ip_blocklist")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-ip-blocklist"], "DnsfilterProfile-ExternalIpBlocklist"); ok {
			if err = d.Set("external_ip_blocklist", vv); err != nil {
				return fmt.Errorf("Error reading external_ip_blocklist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_ip_blocklist: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ftgd_dns", flattenDnsfilterProfileFtgdDns(o["ftgd-dns"], d, "ftgd_dns")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftgd-dns"], "DnsfilterProfile-FtgdDns"); ok {
				if err = d.Set("ftgd_dns", vv); err != nil {
					return fmt.Errorf("Error reading ftgd_dns: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftgd_dns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftgd_dns"); ok {
			if err = d.Set("ftgd_dns", flattenDnsfilterProfileFtgdDns(o["ftgd-dns"], d, "ftgd_dns")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftgd-dns"], "DnsfilterProfile-FtgdDns"); ok {
					if err = d.Set("ftgd_dns", vv); err != nil {
						return fmt.Errorf("Error reading ftgd_dns: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftgd_dns: %v", err)
				}
			}
		}
	}

	if err = d.Set("log_all_domain", flattenDnsfilterProfileLogAllDomain(o["log-all-domain"], d, "log_all_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-all-domain"], "DnsfilterProfile-LogAllDomain"); ok {
			if err = d.Set("log_all_domain", vv); err != nil {
				return fmt.Errorf("Error reading log_all_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_all_domain: %v", err)
		}
	}

	if err = d.Set("name", flattenDnsfilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DnsfilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("redirect_portal", flattenDnsfilterProfileRedirectPortal(o["redirect-portal"], d, "redirect_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-portal"], "DnsfilterProfile-RedirectPortal"); ok {
			if err = d.Set("redirect_portal", vv); err != nil {
				return fmt.Errorf("Error reading redirect_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_portal: %v", err)
		}
	}

	if err = d.Set("redirect_portal6", flattenDnsfilterProfileRedirectPortal6(o["redirect-portal6"], d, "redirect_portal6")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-portal6"], "DnsfilterProfile-RedirectPortal6"); ok {
			if err = d.Set("redirect_portal6", vv); err != nil {
				return fmt.Errorf("Error reading redirect_portal6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_portal6: %v", err)
		}
	}

	if err = d.Set("safe_search", flattenDnsfilterProfileSafeSearch(o["safe-search"], d, "safe_search")); err != nil {
		if vv, ok := fortiAPIPatch(o["safe-search"], "DnsfilterProfile-SafeSearch"); ok {
			if err = d.Set("safe_search", vv); err != nil {
				return fmt.Errorf("Error reading safe_search: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading safe_search: %v", err)
		}
	}

	if err = d.Set("sdns_domain_log", flattenDnsfilterProfileSdnsDomainLog(o["sdns-domain-log"], d, "sdns_domain_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdns-domain-log"], "DnsfilterProfile-SdnsDomainLog"); ok {
			if err = d.Set("sdns_domain_log", vv); err != nil {
				return fmt.Errorf("Error reading sdns_domain_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdns_domain_log: %v", err)
		}
	}

	if err = d.Set("sdns_ftgd_err_log", flattenDnsfilterProfileSdnsFtgdErrLog(o["sdns-ftgd-err-log"], d, "sdns_ftgd_err_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdns-ftgd-err-log"], "DnsfilterProfile-SdnsFtgdErrLog"); ok {
			if err = d.Set("sdns_ftgd_err_log", vv); err != nil {
				return fmt.Errorf("Error reading sdns_ftgd_err_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdns_ftgd_err_log: %v", err)
		}
	}

	if err = d.Set("strip_ech", flattenDnsfilterProfileStripEch(o["strip-ech"], d, "strip_ech")); err != nil {
		if vv, ok := fortiAPIPatch(o["strip-ech"], "DnsfilterProfile-StripEch"); ok {
			if err = d.Set("strip_ech", vv); err != nil {
				return fmt.Errorf("Error reading strip_ech: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strip_ech: %v", err)
		}
	}

	if err = d.Set("transparent_dns_database", flattenDnsfilterProfileTransparentDnsDatabase(o["transparent-dns-database"], d, "transparent_dns_database")); err != nil {
		if vv, ok := fortiAPIPatch(o["transparent-dns-database"], "DnsfilterProfile-TransparentDnsDatabase"); ok {
			if err = d.Set("transparent_dns_database", vv); err != nil {
				return fmt.Errorf("Error reading transparent_dns_database: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transparent_dns_database: %v", err)
		}
	}

	if err = d.Set("youtube_restrict", flattenDnsfilterProfileYoutubeRestrict(o["youtube-restrict"], d, "youtube_restrict")); err != nil {
		if vv, ok := fortiAPIPatch(o["youtube-restrict"], "DnsfilterProfile-YoutubeRestrict"); ok {
			if err = d.Set("youtube_restrict", vv); err != nil {
				return fmt.Errorf("Error reading youtube_restrict: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading youtube_restrict: %v", err)
		}
	}

	return nil
}

func flattenDnsfilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDnsfilterProfileBlockAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileBlockBotnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-type"], _ = expandDnsfilterProfileDnsTranslationAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst"], _ = expandDnsfilterProfileDnsTranslationDst(d, i["dst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst6"], _ = expandDnsfilterProfileDnsTranslationDst6(d, i["dst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandDnsfilterProfileDnsTranslationId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "netmask"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["netmask"], _ = expandDnsfilterProfileDnsTranslationNetmask(d, i["netmask"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandDnsfilterProfileDnsTranslationPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src"], _ = expandDnsfilterProfileDnsTranslationSrc(d, i["src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src6"], _ = expandDnsfilterProfileDnsTranslationSrc6(d, i["src6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandDnsfilterProfileDnsTranslationStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDnsfilterProfileDnsTranslationAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationDst6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationSrc6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDomainFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "domain_filter_table"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["domain-filter-table"], _ = expandDnsfilterProfileDomainFilterDomainFilterTable(d, i["domain_filter_table"], pre_append)
	}

	return result, nil
}

func expandDnsfilterProfileDomainFilterDomainFilterTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDnsfilterProfileExternalIpBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDnsfilterProfileFtgdDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandDnsfilterProfileFtgdDnsFilters(d, i["filters"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["filters"] = t
		}
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandDnsfilterProfileFtgdDnsOptions(d, i["options"], pre_append)
	}

	return result, nil
}

func expandDnsfilterProfileFtgdDnsFilters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandDnsfilterProfileFtgdDnsFiltersAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandDnsfilterProfileFtgdDnsFiltersCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandDnsfilterProfileFtgdDnsFiltersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandDnsfilterProfileFtgdDnsFiltersLog(d, i["log"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDnsfilterProfileFtgdDnsFiltersAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDnsFiltersCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDnsfilterProfileFtgdDnsFiltersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDnsFiltersLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDnsOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDnsfilterProfileLogAllDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileRedirectPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileRedirectPortal6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileSafeSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileSdnsDomainLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileSdnsFtgdErrLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileStripEch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileTransparentDnsDatabase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDnsfilterProfileYoutubeRestrict(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDnsfilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("block_action"); ok || d.HasChange("block_action") {
		t, err := expandDnsfilterProfileBlockAction(d, v, "block_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-action"] = t
		}
	}

	if v, ok := d.GetOk("block_botnet"); ok || d.HasChange("block_botnet") {
		t, err := expandDnsfilterProfileBlockBotnet(d, v, "block_botnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-botnet"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandDnsfilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dns_translation"); ok || d.HasChange("dns_translation") {
		t, err := expandDnsfilterProfileDnsTranslation(d, v, "dns_translation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-translation"] = t
		}
	}

	if v, ok := d.GetOk("domain_filter"); ok || d.HasChange("domain_filter") {
		t, err := expandDnsfilterProfileDomainFilter(d, v, "domain_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-filter"] = t
		}
	}

	if v, ok := d.GetOk("external_ip_blocklist"); ok || d.HasChange("external_ip_blocklist") {
		t, err := expandDnsfilterProfileExternalIpBlocklist(d, v, "external_ip_blocklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-ip-blocklist"] = t
		}
	}

	if v, ok := d.GetOk("ftgd_dns"); ok || d.HasChange("ftgd_dns") {
		t, err := expandDnsfilterProfileFtgdDns(d, v, "ftgd_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftgd-dns"] = t
		}
	}

	if v, ok := d.GetOk("log_all_domain"); ok || d.HasChange("log_all_domain") {
		t, err := expandDnsfilterProfileLogAllDomain(d, v, "log_all_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-all-domain"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDnsfilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("redirect_portal"); ok || d.HasChange("redirect_portal") {
		t, err := expandDnsfilterProfileRedirectPortal(d, v, "redirect_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-portal"] = t
		}
	}

	if v, ok := d.GetOk("redirect_portal6"); ok || d.HasChange("redirect_portal6") {
		t, err := expandDnsfilterProfileRedirectPortal6(d, v, "redirect_portal6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-portal6"] = t
		}
	}

	if v, ok := d.GetOk("safe_search"); ok || d.HasChange("safe_search") {
		t, err := expandDnsfilterProfileSafeSearch(d, v, "safe_search")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safe-search"] = t
		}
	}

	if v, ok := d.GetOk("sdns_domain_log"); ok || d.HasChange("sdns_domain_log") {
		t, err := expandDnsfilterProfileSdnsDomainLog(d, v, "sdns_domain_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-domain-log"] = t
		}
	}

	if v, ok := d.GetOk("sdns_ftgd_err_log"); ok || d.HasChange("sdns_ftgd_err_log") {
		t, err := expandDnsfilterProfileSdnsFtgdErrLog(d, v, "sdns_ftgd_err_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-ftgd-err-log"] = t
		}
	}

	if v, ok := d.GetOk("strip_ech"); ok || d.HasChange("strip_ech") {
		t, err := expandDnsfilterProfileStripEch(d, v, "strip_ech")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-ech"] = t
		}
	}

	if v, ok := d.GetOk("transparent_dns_database"); ok || d.HasChange("transparent_dns_database") {
		t, err := expandDnsfilterProfileTransparentDnsDatabase(d, v, "transparent_dns_database")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transparent-dns-database"] = t
		}
	}

	if v, ok := d.GetOk("youtube_restrict"); ok || d.HasChange("youtube_restrict") {
		t, err := expandDnsfilterProfileYoutubeRestrict(d, v, "youtube_restrict")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-restrict"] = t
		}
	}

	return &obj, nil
}
