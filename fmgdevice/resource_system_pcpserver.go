// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure PCP server information.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPcpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPcpServerUpdate,
		Read:   resourceSystemPcpServerRead,
		Update: resourceSystemPcpServerUpdate,
		Delete: resourceSystemPcpServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"pools": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_opcode": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"announcement_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"arp_reply": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_mapping_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"client_subnet": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ext_intf": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"extip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"extport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"intl_intf": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mapping_filter_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"maximal_lifetime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"minimal_lifetime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"multicast_announcement": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"recycle_delay": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"third_party": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"third_party_subnet": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"status": &schema.Schema{
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

func resourceSystemPcpServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemPcpServer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemPcpServer resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemPcpServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemPcpServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemPcpServer")

	return resourceSystemPcpServerRead(d, m)
}

func resourceSystemPcpServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemPcpServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPcpServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPcpServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemPcpServer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemPcpServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPcpServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemPcpServer resource from API: %v", err)
	}
	return nil
}

func flattenSystemPcpServerPools(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_opcode"
		if _, ok := i["allow-opcode"]; ok {
			v := flattenSystemPcpServerPoolsAllowOpcode(i["allow-opcode"], d, pre_append)
			tmp["allow_opcode"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-AllowOpcode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "announcement_count"
		if _, ok := i["announcement-count"]; ok {
			v := flattenSystemPcpServerPoolsAnnouncementCount(i["announcement-count"], d, pre_append)
			tmp["announcement_count"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-AnnouncementCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := i["arp-reply"]; ok {
			v := flattenSystemPcpServerPoolsArpReply(i["arp-reply"], d, pre_append)
			tmp["arp_reply"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-ArpReply")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_mapping_limit"
		if _, ok := i["client-mapping-limit"]; ok {
			v := flattenSystemPcpServerPoolsClientMappingLimit(i["client-mapping-limit"], d, pre_append)
			tmp["client_mapping_limit"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-ClientMappingLimit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_subnet"
		if _, ok := i["client-subnet"]; ok {
			v := flattenSystemPcpServerPoolsClientSubnet(i["client-subnet"], d, pre_append)
			tmp["client_subnet"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-ClientSubnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenSystemPcpServerPoolsDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ext_intf"
		if _, ok := i["ext-intf"]; ok {
			v := flattenSystemPcpServerPoolsExtIntf(i["ext-intf"], d, pre_append)
			tmp["ext_intf"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-ExtIntf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := i["extip"]; ok {
			v := flattenSystemPcpServerPoolsExtip(i["extip"], d, pre_append)
			tmp["extip"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-Extip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := i["extport"]; ok {
			v := flattenSystemPcpServerPoolsExtport(i["extport"], d, pre_append)
			tmp["extport"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-Extport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemPcpServerPoolsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intl_intf"
		if _, ok := i["intl-intf"]; ok {
			v := flattenSystemPcpServerPoolsIntlIntf(i["intl-intf"], d, pre_append)
			tmp["intl_intf"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-IntlIntf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mapping_filter_limit"
		if _, ok := i["mapping-filter-limit"]; ok {
			v := flattenSystemPcpServerPoolsMappingFilterLimit(i["mapping-filter-limit"], d, pre_append)
			tmp["mapping_filter_limit"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-MappingFilterLimit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximal_lifetime"
		if _, ok := i["maximal-lifetime"]; ok {
			v := flattenSystemPcpServerPoolsMaximalLifetime(i["maximal-lifetime"], d, pre_append)
			tmp["maximal_lifetime"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-MaximalLifetime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimal_lifetime"
		if _, ok := i["minimal-lifetime"]; ok {
			v := flattenSystemPcpServerPoolsMinimalLifetime(i["minimal-lifetime"], d, pre_append)
			tmp["minimal_lifetime"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-MinimalLifetime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_announcement"
		if _, ok := i["multicast-announcement"]; ok {
			v := flattenSystemPcpServerPoolsMulticastAnnouncement(i["multicast-announcement"], d, pre_append)
			tmp["multicast_announcement"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-MulticastAnnouncement")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemPcpServerPoolsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recycle_delay"
		if _, ok := i["recycle-delay"]; ok {
			v := flattenSystemPcpServerPoolsRecycleDelay(i["recycle-delay"], d, pre_append)
			tmp["recycle_delay"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-RecycleDelay")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "third_party"
		if _, ok := i["third-party"]; ok {
			v := flattenSystemPcpServerPoolsThirdParty(i["third-party"], d, pre_append)
			tmp["third_party"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-ThirdParty")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "third_party_subnet"
		if _, ok := i["third-party-subnet"]; ok {
			v := flattenSystemPcpServerPoolsThirdPartySubnet(i["third-party-subnet"], d, pre_append)
			tmp["third_party_subnet"] = fortiAPISubPartPatch(v, "SystemPcpServer-Pools-ThirdPartySubnet")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemPcpServerPoolsAllowOpcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPcpServerPoolsAnnouncementCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsClientMappingLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsClientSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPcpServerPoolsDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsExtIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPcpServerPoolsExtip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsExtport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsIntlIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPcpServerPoolsMappingFilterLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsMaximalLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsMinimalLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsMulticastAnnouncement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsRecycleDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsThirdParty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsThirdPartySubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPcpServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemPcpServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("pools", flattenSystemPcpServerPools(o["pools"], d, "pools")); err != nil {
			if vv, ok := fortiAPIPatch(o["pools"], "SystemPcpServer-Pools"); ok {
				if err = d.Set("pools", vv); err != nil {
					return fmt.Errorf("Error reading pools: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pools: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pools"); ok {
			if err = d.Set("pools", flattenSystemPcpServerPools(o["pools"], d, "pools")); err != nil {
				if vv, ok := fortiAPIPatch(o["pools"], "SystemPcpServer-Pools"); ok {
					if err = d.Set("pools", vv); err != nil {
						return fmt.Errorf("Error reading pools: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pools: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenSystemPcpServerStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemPcpServer-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemPcpServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemPcpServerPools(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_opcode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allow-opcode"], _ = expandSystemPcpServerPoolsAllowOpcode(d, i["allow_opcode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "announcement_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["announcement-count"], _ = expandSystemPcpServerPoolsAnnouncementCount(d, i["announcement_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["arp-reply"], _ = expandSystemPcpServerPoolsArpReply(d, i["arp_reply"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_mapping_limit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-mapping-limit"], _ = expandSystemPcpServerPoolsClientMappingLimit(d, i["client_mapping_limit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-subnet"], _ = expandSystemPcpServerPoolsClientSubnet(d, i["client_subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandSystemPcpServerPoolsDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ext_intf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ext-intf"], _ = expandSystemPcpServerPoolsExtIntf(d, i["ext_intf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extip"], _ = expandSystemPcpServerPoolsExtip(d, i["extip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extport"], _ = expandSystemPcpServerPoolsExtport(d, i["extport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemPcpServerPoolsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intl_intf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["intl-intf"], _ = expandSystemPcpServerPoolsIntlIntf(d, i["intl_intf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mapping_filter_limit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mapping-filter-limit"], _ = expandSystemPcpServerPoolsMappingFilterLimit(d, i["mapping_filter_limit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximal_lifetime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximal-lifetime"], _ = expandSystemPcpServerPoolsMaximalLifetime(d, i["maximal_lifetime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimal_lifetime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["minimal-lifetime"], _ = expandSystemPcpServerPoolsMinimalLifetime(d, i["minimal_lifetime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_announcement"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["multicast-announcement"], _ = expandSystemPcpServerPoolsMulticastAnnouncement(d, i["multicast_announcement"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemPcpServerPoolsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recycle_delay"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["recycle-delay"], _ = expandSystemPcpServerPoolsRecycleDelay(d, i["recycle_delay"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "third_party"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["third-party"], _ = expandSystemPcpServerPoolsThirdParty(d, i["third_party"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "third_party_subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["third-party-subnet"], _ = expandSystemPcpServerPoolsThirdPartySubnet(d, i["third_party_subnet"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemPcpServerPoolsAllowOpcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPcpServerPoolsAnnouncementCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsClientMappingLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsClientSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPcpServerPoolsDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsExtIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPcpServerPoolsExtip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsExtport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsIntlIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPcpServerPoolsMappingFilterLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsMaximalLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsMinimalLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsMulticastAnnouncement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsRecycleDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsThirdParty(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsThirdPartySubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPcpServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPcpServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("pools"); ok || d.HasChange("pools") {
		t, err := expandSystemPcpServerPools(d, v, "pools")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pools"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemPcpServerStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
