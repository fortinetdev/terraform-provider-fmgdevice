// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 SNMP managers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSnmpCommunityHosts6() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpCommunityHosts6Create,
		Read:   resourceSystemSnmpCommunityHosts6Read,
		Update: resourceSystemSnmpCommunityHosts6Update,
		Delete: resourceSystemSnmpCommunityHosts6Delete,

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
			"community": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ha_direct": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemSnmpCommunityHosts6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	community := d.Get("community").(string)
	paradict["device"] = device_name
	paradict["community"] = community

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemSnmpCommunityHosts6(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunityHosts6 resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSnmpCommunityHosts6(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunityHosts6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSnmpCommunityHosts6Read(d, m)
}

func resourceSystemSnmpCommunityHosts6Update(d *schema.ResourceData, m interface{}) error {
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
	community := d.Get("community").(string)
	paradict["device"] = device_name
	paradict["community"] = community

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemSnmpCommunityHosts6(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunityHosts6 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSnmpCommunityHosts6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunityHosts6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSnmpCommunityHosts6Read(d, m)
}

func resourceSystemSnmpCommunityHosts6Delete(d *schema.ResourceData, m interface{}) error {
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
	community := d.Get("community").(string)
	paradict["device"] = device_name
	paradict["community"] = community

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemSnmpCommunityHosts6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpCommunityHosts6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpCommunityHosts6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	community := d.Get("community").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if community == "" {
		community = importOptionChecking(m.(*FortiClient).Cfg, "community")
		if community == "" {
			return fmt.Errorf("Parameter community is missing")
		}
		if err = d.Set("community", community); err != nil {
			return fmt.Errorf("Error set params community: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["community"] = community

	o, err := c.ReadSystemSnmpCommunityHosts6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunityHosts6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpCommunityHosts6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunityHosts6 resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpCommunityHosts6HaDirect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6HostType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6Id2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6Interface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSnmpCommunityHosts6InterfaceSelectMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6Ipv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6SourceIpv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6VrfSelect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSnmpCommunityHosts6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ha_direct", flattenSystemSnmpCommunityHosts6HaDirect2edl(o["ha-direct"], d, "ha_direct")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-direct"], "SystemSnmpCommunityHosts6-HaDirect"); ok {
			if err = d.Set("ha_direct", vv); err != nil {
				return fmt.Errorf("Error reading ha_direct: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_direct: %v", err)
		}
	}

	if err = d.Set("host_type", flattenSystemSnmpCommunityHosts6HostType2edl(o["host-type"], d, "host_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-type"], "SystemSnmpCommunityHosts6-HostType"); ok {
			if err = d.Set("host_type", vv); err != nil {
				return fmt.Errorf("Error reading host_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemSnmpCommunityHosts6Id2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSnmpCommunityHosts6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemSnmpCommunityHosts6Interface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemSnmpCommunityHosts6-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemSnmpCommunityHosts6InterfaceSelectMethod2edl(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystemSnmpCommunityHosts6-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenSystemSnmpCommunityHosts6Ipv62edl(o["ipv6"], d, "ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6"], "SystemSnmpCommunityHosts6-Ipv6"); ok {
			if err = d.Set("ipv6", vv); err != nil {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("source_ipv6", flattenSystemSnmpCommunityHosts6SourceIpv62edl(o["source-ipv6"], d, "source_ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ipv6"], "SystemSnmpCommunityHosts6-SourceIpv6"); ok {
			if err = d.Set("source_ipv6", vv); err != nil {
				return fmt.Errorf("Error reading source_ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ipv6: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenSystemSnmpCommunityHosts6VrfSelect2edl(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "SystemSnmpCommunityHosts6-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpCommunityHosts6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSnmpCommunityHosts6HaDirect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6HostType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6Id2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6Interface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSnmpCommunityHosts6InterfaceSelectMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6Ipv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6SourceIpv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6VrfSelect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpCommunityHosts6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ha_direct"); ok || d.HasChange("ha_direct") {
		t, err := expandSystemSnmpCommunityHosts6HaDirect2edl(d, v, "ha_direct")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-direct"] = t
		}
	}

	if v, ok := d.GetOk("host_type"); ok || d.HasChange("host_type") {
		t, err := expandSystemSnmpCommunityHosts6HostType2edl(d, v, "host_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSnmpCommunityHosts6Id2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemSnmpCommunityHosts6Interface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystemSnmpCommunityHosts6InterfaceSelectMethod2edl(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok || d.HasChange("ipv6") {
		t, err := expandSystemSnmpCommunityHosts6Ipv62edl(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("source_ipv6"); ok || d.HasChange("source_ipv6") {
		t, err := expandSystemSnmpCommunityHosts6SourceIpv62edl(d, v, "source_ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ipv6"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandSystemSnmpCommunityHosts6VrfSelect2edl(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
