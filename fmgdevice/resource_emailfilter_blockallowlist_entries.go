// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Anti-spam block/allow entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterBlockAllowListEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterBlockAllowListEntriesCreate,
		Read:   resourceEmailfilterBlockAllowListEntriesRead,
		Update: resourceEmailfilterBlockAllowListEntriesUpdate,
		Delete: resourceEmailfilterBlockAllowListEntriesDelete,

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
			"block_allow_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceEmailfilterBlockAllowListEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	block_allow_list := d.Get("block_allow_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["block_allow_list"] = block_allow_list

	obj, err := getObjectEmailfilterBlockAllowListEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating EmailfilterBlockAllowListEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadEmailfilterBlockAllowListEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateEmailfilterBlockAllowListEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating EmailfilterBlockAllowListEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateEmailfilterBlockAllowListEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating EmailfilterBlockAllowListEntries resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterBlockAllowListEntriesRead(d, m)
}

func resourceEmailfilterBlockAllowListEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	block_allow_list := d.Get("block_allow_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["block_allow_list"] = block_allow_list

	obj, err := getObjectEmailfilterBlockAllowListEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterBlockAllowListEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEmailfilterBlockAllowListEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterBlockAllowListEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterBlockAllowListEntriesRead(d, m)
}

func resourceEmailfilterBlockAllowListEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	block_allow_list := d.Get("block_allow_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["block_allow_list"] = block_allow_list

	wsParams["adom"] = adomv

	err = c.DeleteEmailfilterBlockAllowListEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterBlockAllowListEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterBlockAllowListEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	block_allow_list := d.Get("block_allow_list").(string)
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
	if block_allow_list == "" {
		block_allow_list = importOptionChecking(m.(*FortiClient).Cfg, "block_allow_list")
		if block_allow_list == "" {
			return fmt.Errorf("Parameter block_allow_list is missing")
		}
		if err = d.Set("block_allow_list", block_allow_list); err != nil {
			return fmt.Errorf("Error set params block_allow_list: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["block_allow_list"] = block_allow_list

	o, err := c.ReadEmailfilterBlockAllowListEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EmailfilterBlockAllowListEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterBlockAllowListEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterBlockAllowListEntries resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterBlockAllowListEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesIp4Subnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterBlockAllowListEntriesIp6Subnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesPatternType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterBlockAllowListEntriesEmailPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEmailfilterBlockAllowListEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenEmailfilterBlockAllowListEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "EmailfilterBlockAllowListEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("addr_type", flattenEmailfilterBlockAllowListEntriesAddrType2edl(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "EmailfilterBlockAllowListEntries-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenEmailfilterBlockAllowListEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "EmailfilterBlockAllowListEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip4_subnet", flattenEmailfilterBlockAllowListEntriesIp4Subnet2edl(o["ip4-subnet"], d, "ip4_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip4-subnet"], "EmailfilterBlockAllowListEntries-Ip4Subnet"); ok {
			if err = d.Set("ip4_subnet", vv); err != nil {
				return fmt.Errorf("Error reading ip4_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip4_subnet: %v", err)
		}
	}

	if err = d.Set("ip6_subnet", flattenEmailfilterBlockAllowListEntriesIp6Subnet2edl(o["ip6-subnet"], d, "ip6_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-subnet"], "EmailfilterBlockAllowListEntries-Ip6Subnet"); ok {
			if err = d.Set("ip6_subnet", vv); err != nil {
				return fmt.Errorf("Error reading ip6_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_subnet: %v", err)
		}
	}

	if err = d.Set("pattern", flattenEmailfilterBlockAllowListEntriesPattern2edl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "EmailfilterBlockAllowListEntries-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("pattern_type", flattenEmailfilterBlockAllowListEntriesPatternType2edl(o["pattern-type"], d, "pattern_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern-type"], "EmailfilterBlockAllowListEntries-PatternType"); ok {
			if err = d.Set("pattern_type", vv); err != nil {
				return fmt.Errorf("Error reading pattern_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern_type: %v", err)
		}
	}

	if err = d.Set("status", flattenEmailfilterBlockAllowListEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "EmailfilterBlockAllowListEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenEmailfilterBlockAllowListEntriesType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "EmailfilterBlockAllowListEntries-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("email_pattern", flattenEmailfilterBlockAllowListEntriesEmailPattern2edl(o["email-pattern"], d, "email_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-pattern"], "EmailfilterBlockAllowListEntries-EmailPattern"); ok {
			if err = d.Set("email_pattern", vv); err != nil {
				return fmt.Errorf("Error reading email_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_pattern: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterBlockAllowListEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEmailfilterBlockAllowListEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesIp4Subnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandEmailfilterBlockAllowListEntriesIp6Subnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesPatternType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBlockAllowListEntriesEmailPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterBlockAllowListEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandEmailfilterBlockAllowListEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandEmailfilterBlockAllowListEntriesAddrType2edl(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandEmailfilterBlockAllowListEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip4_subnet"); ok || d.HasChange("ip4_subnet") {
		t, err := expandEmailfilterBlockAllowListEntriesIp4Subnet2edl(d, v, "ip4_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip4-subnet"] = t
		}
	}

	if v, ok := d.GetOk("ip6_subnet"); ok || d.HasChange("ip6_subnet") {
		t, err := expandEmailfilterBlockAllowListEntriesIp6Subnet2edl(d, v, "ip6_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-subnet"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandEmailfilterBlockAllowListEntriesPattern2edl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("pattern_type"); ok || d.HasChange("pattern_type") {
		t, err := expandEmailfilterBlockAllowListEntriesPatternType2edl(d, v, "pattern_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern-type"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandEmailfilterBlockAllowListEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandEmailfilterBlockAllowListEntriesType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("email_pattern"); ok || d.HasChange("email_pattern") {
		t, err := expandEmailfilterBlockAllowListEntriesEmailPattern2edl(d, v, "email_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-pattern"] = t
		}
	}

	return &obj, nil
}
