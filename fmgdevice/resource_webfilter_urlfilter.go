// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure URL filter lists.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterUrlfilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterUrlfilterCreate,
		Read:   resourceWebfilterUrlfilterRead,
		Update: resourceWebfilterUrlfilterUpdate,
		Delete: resourceWebfilterUrlfilterDelete,

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
							Computed: true,
						},
						"antiphish_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dns_address_family": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"exempt": &schema.Schema{
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
						"referrer_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"web_proxy_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"include_subdomains": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_addr_block": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip4_mapped_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"one_arm_ips_urlfilter": &schema.Schema{
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

func resourceWebfilterUrlfilterCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterUrlfilter(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterUrlfilter resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebfilterUrlfilter(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebfilterUrlfilter(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebfilterUrlfilter resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebfilterUrlfilter(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebfilterUrlfilter resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebfilterUrlfilterRead(d, m)
}

func resourceWebfilterUrlfilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterUrlfilter(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterUrlfilter resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterUrlfilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterUrlfilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebfilterUrlfilterRead(d, m)
}

func resourceWebfilterUrlfilterDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebfilterUrlfilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterUrlfilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterUrlfilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterUrlfilter(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterUrlfilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterUrlfilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterUrlfilter resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterUrlfilterComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWebfilterUrlfilterEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WebfilterUrlfilter-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antiphish_action"
		if _, ok := i["antiphish-action"]; ok {
			v := flattenWebfilterUrlfilterEntriesAntiphishAction(i["antiphish-action"], d, pre_append)
			tmp["antiphish_action"] = fortiAPISubPartPatch(v, "WebfilterUrlfilter-Entries-AntiphishAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenWebfilterUrlfilterEntriesComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WebfilterUrlfilter-Entries-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_address_family"
		if _, ok := i["dns-address-family"]; ok {
			v := flattenWebfilterUrlfilterEntriesDnsAddressFamily(i["dns-address-family"], d, pre_append)
			tmp["dns_address_family"] = fortiAPISubPartPatch(v, "WebfilterUrlfilter-Entries-DnsAddressFamily")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt"
		if _, ok := i["exempt"]; ok {
			v := flattenWebfilterUrlfilterEntriesExempt(i["exempt"], d, pre_append)
			tmp["exempt"] = fortiAPISubPartPatch(v, "WebfilterUrlfilter-Entries-Exempt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWebfilterUrlfilterEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WebfilterUrlfilter-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "referrer_host"
		if _, ok := i["referrer-host"]; ok {
			v := flattenWebfilterUrlfilterEntriesReferrerHost(i["referrer-host"], d, pre_append)
			tmp["referrer_host"] = fortiAPISubPartPatch(v, "WebfilterUrlfilter-Entries-ReferrerHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWebfilterUrlfilterEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WebfilterUrlfilter-Entries-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenWebfilterUrlfilterEntriesType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "WebfilterUrlfilter-Entries-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			v := flattenWebfilterUrlfilterEntriesUrl(i["url"], d, pre_append)
			tmp["url"] = fortiAPISubPartPatch(v, "WebfilterUrlfilter-Entries-Url")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "web_proxy_profile"
		if _, ok := i["web-proxy-profile"]; ok {
			v := flattenWebfilterUrlfilterEntriesWebProxyProfile(i["web-proxy-profile"], d, pre_append)
			tmp["web_proxy_profile"] = fortiAPISubPartPatch(v, "WebfilterUrlfilter-Entries-WebProxyProfile")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebfilterUrlfilterEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesAntiphishAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesDnsAddressFamily(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesExempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterUrlfilterEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesReferrerHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesWebProxyProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterUrlfilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterIpAddrBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterIp4MappedIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterOneArmIpsUrlfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterUrlfilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenWebfilterUrlfilterComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WebfilterUrlfilter-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenWebfilterUrlfilterEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "WebfilterUrlfilter-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenWebfilterUrlfilterEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "WebfilterUrlfilter-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenWebfilterUrlfilterId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WebfilterUrlfilter-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("include_subdomains", flattenWebfilterUrlfilterIncludeSubdomains(o["include-subdomains"], d, "include_subdomains")); err != nil {
		if vv, ok := fortiAPIPatch(o["include-subdomains"], "WebfilterUrlfilter-IncludeSubdomains"); ok {
			if err = d.Set("include_subdomains", vv); err != nil {
				return fmt.Errorf("Error reading include_subdomains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading include_subdomains: %v", err)
		}
	}

	if err = d.Set("ip_addr_block", flattenWebfilterUrlfilterIpAddrBlock(o["ip-addr-block"], d, "ip_addr_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-addr-block"], "WebfilterUrlfilter-IpAddrBlock"); ok {
			if err = d.Set("ip_addr_block", vv); err != nil {
				return fmt.Errorf("Error reading ip_addr_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_addr_block: %v", err)
		}
	}

	if err = d.Set("ip4_mapped_ip6", flattenWebfilterUrlfilterIp4MappedIp6(o["ip4-mapped-ip6"], d, "ip4_mapped_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip4-mapped-ip6"], "WebfilterUrlfilter-Ip4MappedIp6"); ok {
			if err = d.Set("ip4_mapped_ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip4_mapped_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip4_mapped_ip6: %v", err)
		}
	}

	if err = d.Set("name", flattenWebfilterUrlfilterName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebfilterUrlfilter-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("one_arm_ips_urlfilter", flattenWebfilterUrlfilterOneArmIpsUrlfilter(o["one-arm-ips-urlfilter"], d, "one_arm_ips_urlfilter")); err != nil {
		if vv, ok := fortiAPIPatch(o["one-arm-ips-urlfilter"], "WebfilterUrlfilter-OneArmIpsUrlfilter"); ok {
			if err = d.Set("one_arm_ips_urlfilter", vv); err != nil {
				return fmt.Errorf("Error reading one_arm_ips_urlfilter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading one_arm_ips_urlfilter: %v", err)
		}
	}

	return nil
}

func flattenWebfilterUrlfilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterUrlfilterComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandWebfilterUrlfilterEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antiphish_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["antiphish-action"], _ = expandWebfilterUrlfilterEntriesAntiphishAction(d, i["antiphish_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandWebfilterUrlfilterEntriesComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_address_family"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-address-family"], _ = expandWebfilterUrlfilterEntriesDnsAddressFamily(d, i["dns_address_family"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exempt"], _ = expandWebfilterUrlfilterEntriesExempt(d, i["exempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWebfilterUrlfilterEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "referrer_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["referrer-host"], _ = expandWebfilterUrlfilterEntriesReferrerHost(d, i["referrer_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWebfilterUrlfilterEntriesStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandWebfilterUrlfilterEntriesType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url"], _ = expandWebfilterUrlfilterEntriesUrl(d, i["url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "web_proxy_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["web-proxy-profile"], _ = expandWebfilterUrlfilterEntriesWebProxyProfile(d, i["web_proxy_profile"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebfilterUrlfilterEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesAntiphishAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesDnsAddressFamily(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesExempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterUrlfilterEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesReferrerHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesWebProxyProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterUrlfilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterIpAddrBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterIp4MappedIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterOneArmIpsUrlfilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterUrlfilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWebfilterUrlfilterComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandWebfilterUrlfilterEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWebfilterUrlfilterId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("include_subdomains"); ok || d.HasChange("include_subdomains") {
		t, err := expandWebfilterUrlfilterIncludeSubdomains(d, v, "include_subdomains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ip_addr_block"); ok || d.HasChange("ip_addr_block") {
		t, err := expandWebfilterUrlfilterIpAddrBlock(d, v, "ip_addr_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-addr-block"] = t
		}
	}

	if v, ok := d.GetOk("ip4_mapped_ip6"); ok || d.HasChange("ip4_mapped_ip6") {
		t, err := expandWebfilterUrlfilterIp4MappedIp6(d, v, "ip4_mapped_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip4-mapped-ip6"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebfilterUrlfilterName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("one_arm_ips_urlfilter"); ok || d.HasChange("one_arm_ips_urlfilter") {
		t, err := expandWebfilterUrlfilterOneArmIpsUrlfilter(d, v, "one_arm_ips_urlfilter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["one-arm-ips-urlfilter"] = t
		}
	}

	return &obj, nil
}
