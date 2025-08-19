// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Advertised IPv6 RDNSS list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemInterfaceIpv6Ip6RdnssList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceIpv6Ip6RdnssListCreate,
		Read:   resourceSystemInterfaceIpv6Ip6RdnssListRead,
		Update: resourceSystemInterfaceIpv6Ip6RdnssListUpdate,
		Delete: resourceSystemInterfaceIpv6Ip6RdnssListDelete,

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
			"rdnss": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"rdnss_life_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemInterfaceIpv6Ip6RdnssListCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemInterfaceIpv6Ip6RdnssList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceIpv6Ip6RdnssList resource while getting object: %v", err)
	}

	_, err = c.CreateSystemInterfaceIpv6Ip6RdnssList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceIpv6Ip6RdnssList resource: %v", err)
	}

	d.SetId(getStringKey(d, "rdnss"))

	return resourceSystemInterfaceIpv6Ip6RdnssListRead(d, m)
}

func resourceSystemInterfaceIpv6Ip6RdnssListUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemInterfaceIpv6Ip6RdnssList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6Ip6RdnssList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceIpv6Ip6RdnssList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6Ip6RdnssList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "rdnss"))

	return resourceSystemInterfaceIpv6Ip6RdnssListRead(d, m)
}

func resourceSystemInterfaceIpv6Ip6RdnssListDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemInterfaceIpv6Ip6RdnssList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceIpv6Ip6RdnssList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceIpv6Ip6RdnssListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemInterfaceIpv6Ip6RdnssList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6Ip6RdnssList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceIpv6Ip6RdnssList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6Ip6RdnssList resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceIpv6Ip6RdnssListRdnss3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6RdnssListRdnssLifeTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterfaceIpv6Ip6RdnssList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("rdnss", flattenSystemInterfaceIpv6Ip6RdnssListRdnss3rdl(o["rdnss"], d, "rdnss")); err != nil {
		if vv, ok := fortiAPIPatch(o["rdnss"], "SystemInterfaceIpv6Ip6RdnssList-Rdnss"); ok {
			if err = d.Set("rdnss", vv); err != nil {
				return fmt.Errorf("Error reading rdnss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rdnss: %v", err)
		}
	}

	if err = d.Set("rdnss_life_time", flattenSystemInterfaceIpv6Ip6RdnssListRdnssLifeTime3rdl(o["rdnss-life-time"], d, "rdnss_life_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["rdnss-life-time"], "SystemInterfaceIpv6Ip6RdnssList-RdnssLifeTime"); ok {
			if err = d.Set("rdnss_life_time", vv); err != nil {
				return fmt.Errorf("Error reading rdnss_life_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rdnss_life_time: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfaceIpv6Ip6RdnssListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceIpv6Ip6RdnssListRdnss3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6RdnssListRdnssLifeTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceIpv6Ip6RdnssList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("rdnss"); ok || d.HasChange("rdnss") {
		t, err := expandSystemInterfaceIpv6Ip6RdnssListRdnss3rdl(d, v, "rdnss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rdnss"] = t
		}
	}

	if v, ok := d.GetOk("rdnss_life_time"); ok || d.HasChange("rdnss_life_time") {
		t, err := expandSystemInterfaceIpv6Ip6RdnssListRdnssLifeTime3rdl(d, v, "rdnss_life_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rdnss-life-time"] = t
		}
	}

	return &obj, nil
}
