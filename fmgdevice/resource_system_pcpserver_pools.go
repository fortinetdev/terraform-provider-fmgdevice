// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure PCP pools.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPcpServerPools() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPcpServerPoolsCreate,
		Read:   resourceSystemPcpServerPoolsRead,
		Update: resourceSystemPcpServerPoolsUpdate,
		Delete: resourceSystemPcpServerPoolsDelete,

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
			"fosid": &schema.Schema{
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
				ForceNew: true,
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
	}
}

func resourceSystemPcpServerPoolsCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemPcpServerPools(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemPcpServerPools resource while getting object: %v", err)
	}

	_, err = c.CreateSystemPcpServerPools(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemPcpServerPools resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemPcpServerPoolsRead(d, m)
}

func resourceSystemPcpServerPoolsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemPcpServerPools(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemPcpServerPools resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemPcpServerPools(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemPcpServerPools resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemPcpServerPoolsRead(d, m)
}

func resourceSystemPcpServerPoolsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemPcpServerPools(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPcpServerPools resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPcpServerPoolsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemPcpServerPools(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemPcpServerPools resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPcpServerPools(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemPcpServerPools resource from API: %v", err)
	}
	return nil
}

func flattenSystemPcpServerPoolsAllowOpcode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPcpServerPoolsAnnouncementCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsArpReply2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsClientMappingLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsClientSubnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPcpServerPoolsDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsExtIntf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPcpServerPoolsExtip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsExtport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsIntlIntf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPcpServerPoolsMappingFilterLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsMaximalLifetime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsMinimalLifetime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsMulticastAnnouncement2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsRecycleDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsThirdParty2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsThirdPartySubnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemPcpServerPools(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("allow_opcode", flattenSystemPcpServerPoolsAllowOpcode2edl(o["allow-opcode"], d, "allow_opcode")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-opcode"], "SystemPcpServerPools-AllowOpcode"); ok {
			if err = d.Set("allow_opcode", vv); err != nil {
				return fmt.Errorf("Error reading allow_opcode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_opcode: %v", err)
		}
	}

	if err = d.Set("announcement_count", flattenSystemPcpServerPoolsAnnouncementCount2edl(o["announcement-count"], d, "announcement_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["announcement-count"], "SystemPcpServerPools-AnnouncementCount"); ok {
			if err = d.Set("announcement_count", vv); err != nil {
				return fmt.Errorf("Error reading announcement_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading announcement_count: %v", err)
		}
	}

	if err = d.Set("arp_reply", flattenSystemPcpServerPoolsArpReply2edl(o["arp-reply"], d, "arp_reply")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-reply"], "SystemPcpServerPools-ArpReply"); ok {
			if err = d.Set("arp_reply", vv); err != nil {
				return fmt.Errorf("Error reading arp_reply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("client_mapping_limit", flattenSystemPcpServerPoolsClientMappingLimit2edl(o["client-mapping-limit"], d, "client_mapping_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-mapping-limit"], "SystemPcpServerPools-ClientMappingLimit"); ok {
			if err = d.Set("client_mapping_limit", vv); err != nil {
				return fmt.Errorf("Error reading client_mapping_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_mapping_limit: %v", err)
		}
	}

	if err = d.Set("client_subnet", flattenSystemPcpServerPoolsClientSubnet2edl(o["client-subnet"], d, "client_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-subnet"], "SystemPcpServerPools-ClientSubnet"); ok {
			if err = d.Set("client_subnet", vv); err != nil {
				return fmt.Errorf("Error reading client_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_subnet: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemPcpServerPoolsDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemPcpServerPools-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("ext_intf", flattenSystemPcpServerPoolsExtIntf2edl(o["ext-intf"], d, "ext_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["ext-intf"], "SystemPcpServerPools-ExtIntf"); ok {
			if err = d.Set("ext_intf", vv); err != nil {
				return fmt.Errorf("Error reading ext_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ext_intf: %v", err)
		}
	}

	if err = d.Set("extip", flattenSystemPcpServerPoolsExtip2edl(o["extip"], d, "extip")); err != nil {
		if vv, ok := fortiAPIPatch(o["extip"], "SystemPcpServerPools-Extip"); ok {
			if err = d.Set("extip", vv); err != nil {
				return fmt.Errorf("Error reading extip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extip: %v", err)
		}
	}

	if err = d.Set("extport", flattenSystemPcpServerPoolsExtport2edl(o["extport"], d, "extport")); err != nil {
		if vv, ok := fortiAPIPatch(o["extport"], "SystemPcpServerPools-Extport"); ok {
			if err = d.Set("extport", vv); err != nil {
				return fmt.Errorf("Error reading extport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extport: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemPcpServerPoolsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemPcpServerPools-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("intl_intf", flattenSystemPcpServerPoolsIntlIntf2edl(o["intl-intf"], d, "intl_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["intl-intf"], "SystemPcpServerPools-IntlIntf"); ok {
			if err = d.Set("intl_intf", vv); err != nil {
				return fmt.Errorf("Error reading intl_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intl_intf: %v", err)
		}
	}

	if err = d.Set("mapping_filter_limit", flattenSystemPcpServerPoolsMappingFilterLimit2edl(o["mapping-filter-limit"], d, "mapping_filter_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["mapping-filter-limit"], "SystemPcpServerPools-MappingFilterLimit"); ok {
			if err = d.Set("mapping_filter_limit", vv); err != nil {
				return fmt.Errorf("Error reading mapping_filter_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mapping_filter_limit: %v", err)
		}
	}

	if err = d.Set("maximal_lifetime", flattenSystemPcpServerPoolsMaximalLifetime2edl(o["maximal-lifetime"], d, "maximal_lifetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximal-lifetime"], "SystemPcpServerPools-MaximalLifetime"); ok {
			if err = d.Set("maximal_lifetime", vv); err != nil {
				return fmt.Errorf("Error reading maximal_lifetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximal_lifetime: %v", err)
		}
	}

	if err = d.Set("minimal_lifetime", flattenSystemPcpServerPoolsMinimalLifetime2edl(o["minimal-lifetime"], d, "minimal_lifetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["minimal-lifetime"], "SystemPcpServerPools-MinimalLifetime"); ok {
			if err = d.Set("minimal_lifetime", vv); err != nil {
				return fmt.Errorf("Error reading minimal_lifetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minimal_lifetime: %v", err)
		}
	}

	if err = d.Set("multicast_announcement", flattenSystemPcpServerPoolsMulticastAnnouncement2edl(o["multicast-announcement"], d, "multicast_announcement")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-announcement"], "SystemPcpServerPools-MulticastAnnouncement"); ok {
			if err = d.Set("multicast_announcement", vv); err != nil {
				return fmt.Errorf("Error reading multicast_announcement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_announcement: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemPcpServerPoolsName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemPcpServerPools-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("recycle_delay", flattenSystemPcpServerPoolsRecycleDelay2edl(o["recycle-delay"], d, "recycle_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["recycle-delay"], "SystemPcpServerPools-RecycleDelay"); ok {
			if err = d.Set("recycle_delay", vv); err != nil {
				return fmt.Errorf("Error reading recycle_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading recycle_delay: %v", err)
		}
	}

	if err = d.Set("third_party", flattenSystemPcpServerPoolsThirdParty2edl(o["third-party"], d, "third_party")); err != nil {
		if vv, ok := fortiAPIPatch(o["third-party"], "SystemPcpServerPools-ThirdParty"); ok {
			if err = d.Set("third_party", vv); err != nil {
				return fmt.Errorf("Error reading third_party: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading third_party: %v", err)
		}
	}

	if err = d.Set("third_party_subnet", flattenSystemPcpServerPoolsThirdPartySubnet2edl(o["third-party-subnet"], d, "third_party_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["third-party-subnet"], "SystemPcpServerPools-ThirdPartySubnet"); ok {
			if err = d.Set("third_party_subnet", vv); err != nil {
				return fmt.Errorf("Error reading third_party_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading third_party_subnet: %v", err)
		}
	}

	return nil
}

func flattenSystemPcpServerPoolsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemPcpServerPoolsAllowOpcode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPcpServerPoolsAnnouncementCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsArpReply2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsClientMappingLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsClientSubnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPcpServerPoolsDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsExtIntf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPcpServerPoolsExtip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsExtport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsIntlIntf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPcpServerPoolsMappingFilterLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsMaximalLifetime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsMinimalLifetime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsMulticastAnnouncement2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsRecycleDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsThirdParty2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsThirdPartySubnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemPcpServerPools(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allow_opcode"); ok || d.HasChange("allow_opcode") {
		t, err := expandSystemPcpServerPoolsAllowOpcode2edl(d, v, "allow_opcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-opcode"] = t
		}
	}

	if v, ok := d.GetOk("announcement_count"); ok || d.HasChange("announcement_count") {
		t, err := expandSystemPcpServerPoolsAnnouncementCount2edl(d, v, "announcement_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["announcement-count"] = t
		}
	}

	if v, ok := d.GetOk("arp_reply"); ok || d.HasChange("arp_reply") {
		t, err := expandSystemPcpServerPoolsArpReply2edl(d, v, "arp_reply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("client_mapping_limit"); ok || d.HasChange("client_mapping_limit") {
		t, err := expandSystemPcpServerPoolsClientMappingLimit2edl(d, v, "client_mapping_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-mapping-limit"] = t
		}
	}

	if v, ok := d.GetOk("client_subnet"); ok || d.HasChange("client_subnet") {
		t, err := expandSystemPcpServerPoolsClientSubnet2edl(d, v, "client_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-subnet"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemPcpServerPoolsDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("ext_intf"); ok || d.HasChange("ext_intf") {
		t, err := expandSystemPcpServerPoolsExtIntf2edl(d, v, "ext_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-intf"] = t
		}
	}

	if v, ok := d.GetOk("extip"); ok || d.HasChange("extip") {
		t, err := expandSystemPcpServerPoolsExtip2edl(d, v, "extip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extip"] = t
		}
	}

	if v, ok := d.GetOk("extport"); ok || d.HasChange("extport") {
		t, err := expandSystemPcpServerPoolsExtport2edl(d, v, "extport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extport"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemPcpServerPoolsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("intl_intf"); ok || d.HasChange("intl_intf") {
		t, err := expandSystemPcpServerPoolsIntlIntf2edl(d, v, "intl_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intl-intf"] = t
		}
	}

	if v, ok := d.GetOk("mapping_filter_limit"); ok || d.HasChange("mapping_filter_limit") {
		t, err := expandSystemPcpServerPoolsMappingFilterLimit2edl(d, v, "mapping_filter_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapping-filter-limit"] = t
		}
	}

	if v, ok := d.GetOk("maximal_lifetime"); ok || d.HasChange("maximal_lifetime") {
		t, err := expandSystemPcpServerPoolsMaximalLifetime2edl(d, v, "maximal_lifetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximal-lifetime"] = t
		}
	}

	if v, ok := d.GetOk("minimal_lifetime"); ok || d.HasChange("minimal_lifetime") {
		t, err := expandSystemPcpServerPoolsMinimalLifetime2edl(d, v, "minimal_lifetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimal-lifetime"] = t
		}
	}

	if v, ok := d.GetOk("multicast_announcement"); ok || d.HasChange("multicast_announcement") {
		t, err := expandSystemPcpServerPoolsMulticastAnnouncement2edl(d, v, "multicast_announcement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-announcement"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemPcpServerPoolsName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("recycle_delay"); ok || d.HasChange("recycle_delay") {
		t, err := expandSystemPcpServerPoolsRecycleDelay2edl(d, v, "recycle_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recycle-delay"] = t
		}
	}

	if v, ok := d.GetOk("third_party"); ok || d.HasChange("third_party") {
		t, err := expandSystemPcpServerPoolsThirdParty2edl(d, v, "third_party")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["third-party"] = t
		}
	}

	if v, ok := d.GetOk("third_party_subnet"); ok || d.HasChange("third_party_subnet") {
		t, err := expandSystemPcpServerPoolsThirdPartySubnet2edl(d, v, "third_party_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["third-party-subnet"] = t
		}
	}

	return &obj, nil
}
