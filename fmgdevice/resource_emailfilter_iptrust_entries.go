// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Spam filter trusted IP addresses.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterIptrustEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterIptrustEntriesCreate,
		Read:   resourceEmailfilterIptrustEntriesRead,
		Update: resourceEmailfilterIptrustEntriesUpdate,
		Delete: resourceEmailfilterIptrustEntriesDelete,

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
			"iptrust": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceEmailfilterIptrustEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	iptrust := d.Get("iptrust").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["iptrust"] = iptrust

	obj, err := getObjectEmailfilterIptrustEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating EmailfilterIptrustEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadEmailfilterIptrustEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateEmailfilterIptrustEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating EmailfilterIptrustEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateEmailfilterIptrustEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating EmailfilterIptrustEntries resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterIptrustEntriesRead(d, m)
}

func resourceEmailfilterIptrustEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	iptrust := d.Get("iptrust").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["iptrust"] = iptrust

	obj, err := getObjectEmailfilterIptrustEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterIptrustEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEmailfilterIptrustEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterIptrustEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceEmailfilterIptrustEntriesRead(d, m)
}

func resourceEmailfilterIptrustEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	iptrust := d.Get("iptrust").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["iptrust"] = iptrust

	wsParams["adom"] = adomv

	err = c.DeleteEmailfilterIptrustEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterIptrustEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterIptrustEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	iptrust := d.Get("iptrust").(string)
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
	if iptrust == "" {
		iptrust = importOptionChecking(m.(*FortiClient).Cfg, "iptrust")
		if iptrust == "" {
			return fmt.Errorf("Parameter iptrust is missing")
		}
		if err = d.Set("iptrust", iptrust); err != nil {
			return fmt.Errorf("Error set params iptrust: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["iptrust"] = iptrust

	o, err := c.ReadEmailfilterIptrustEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EmailfilterIptrustEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterIptrustEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterIptrustEntries resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterIptrustEntriesAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterIptrustEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterIptrustEntriesIp4Subnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterIptrustEntriesIp6Subnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterIptrustEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEmailfilterIptrustEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_type", flattenEmailfilterIptrustEntriesAddrType2edl(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "EmailfilterIptrustEntries-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenEmailfilterIptrustEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "EmailfilterIptrustEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip4_subnet", flattenEmailfilterIptrustEntriesIp4Subnet2edl(o["ip4-subnet"], d, "ip4_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip4-subnet"], "EmailfilterIptrustEntries-Ip4Subnet"); ok {
			if err = d.Set("ip4_subnet", vv); err != nil {
				return fmt.Errorf("Error reading ip4_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip4_subnet: %v", err)
		}
	}

	if err = d.Set("ip6_subnet", flattenEmailfilterIptrustEntriesIp6Subnet2edl(o["ip6-subnet"], d, "ip6_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-subnet"], "EmailfilterIptrustEntries-Ip6Subnet"); ok {
			if err = d.Set("ip6_subnet", vv); err != nil {
				return fmt.Errorf("Error reading ip6_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_subnet: %v", err)
		}
	}

	if err = d.Set("status", flattenEmailfilterIptrustEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "EmailfilterIptrustEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterIptrustEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEmailfilterIptrustEntriesAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterIptrustEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterIptrustEntriesIp4Subnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandEmailfilterIptrustEntriesIp6Subnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterIptrustEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterIptrustEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandEmailfilterIptrustEntriesAddrType2edl(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandEmailfilterIptrustEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip4_subnet"); ok || d.HasChange("ip4_subnet") {
		t, err := expandEmailfilterIptrustEntriesIp4Subnet2edl(d, v, "ip4_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip4-subnet"] = t
		}
	}

	if v, ok := d.GetOk("ip6_subnet"); ok || d.HasChange("ip6_subnet") {
		t, err := expandEmailfilterIptrustEntriesIp6Subnet2edl(d, v, "ip6_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-subnet"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandEmailfilterIptrustEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
