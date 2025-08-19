// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Advertised IPv6 DNSS list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemInterfaceIpv6Ip6DnsslList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceIpv6Ip6DnsslListCreate,
		Read:   resourceSystemInterfaceIpv6Ip6DnsslListRead,
		Update: resourceSystemInterfaceIpv6Ip6DnsslListUpdate,
		Delete: resourceSystemInterfaceIpv6Ip6DnsslListDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dnssl_life_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemInterfaceIpv6Ip6DnsslListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemInterfaceIpv6Ip6DnsslList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceIpv6Ip6DnsslList resource while getting object: %v", err)
	}

	_, err = c.CreateSystemInterfaceIpv6Ip6DnsslList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceIpv6Ip6DnsslList resource: %v", err)
	}

	d.SetId(getStringKey(d, "domain"))

	return resourceSystemInterfaceIpv6Ip6DnsslListRead(d, m)
}

func resourceSystemInterfaceIpv6Ip6DnsslListUpdate(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemInterfaceIpv6Ip6DnsslList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6Ip6DnsslList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceIpv6Ip6DnsslList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6Ip6DnsslList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "domain"))

	return resourceSystemInterfaceIpv6Ip6DnsslListRead(d, m)
}

func resourceSystemInterfaceIpv6Ip6DnsslListDelete(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemInterfaceIpv6Ip6DnsslList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceIpv6Ip6DnsslList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceIpv6Ip6DnsslListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	var_interface := d.Get("interface").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if var_interface == "" {
		var_interface = importOptionChecking(m.(*FortiClient).Cfg, "interface")
		if var_interface == "" {
			return fmt.Errorf("Parameter interface is missing")
		}
		if err = d.Set("interface", var_interface); err != nil {
			return fmt.Errorf("Error set params interface: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	o, err := c.ReadSystemInterfaceIpv6Ip6DnsslList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6Ip6DnsslList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceIpv6Ip6DnsslList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6Ip6DnsslList resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceIpv6Ip6DnsslListDnsslLifeTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DnsslListDomain3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterfaceIpv6Ip6DnsslList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dnssl_life_time", flattenSystemInterfaceIpv6Ip6DnsslListDnsslLifeTime3rdl(o["dnssl-life-time"], d, "dnssl_life_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnssl-life-time"], "SystemInterfaceIpv6Ip6DnsslList-DnsslLifeTime"); ok {
			if err = d.Set("dnssl_life_time", vv); err != nil {
				return fmt.Errorf("Error reading dnssl_life_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnssl_life_time: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemInterfaceIpv6Ip6DnsslListDomain3rdl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "SystemInterfaceIpv6Ip6DnsslList-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfaceIpv6Ip6DnsslListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceIpv6Ip6DnsslListDnsslLifeTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DnsslListDomain3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceIpv6Ip6DnsslList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dnssl_life_time"); ok || d.HasChange("dnssl_life_time") {
		t, err := expandSystemInterfaceIpv6Ip6DnsslListDnsslLifeTime3rdl(d, v, "dnssl_life_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnssl-life-time"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandSystemInterfaceIpv6Ip6DnsslListDomain3rdl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	return &obj, nil
}
