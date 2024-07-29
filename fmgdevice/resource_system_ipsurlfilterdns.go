// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPS URL filter DNS servers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIpsUrlfilterDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpsUrlfilterDnsCreate,
		Read:   resourceSystemIpsUrlfilterDnsRead,
		Update: resourceSystemIpsUrlfilterDnsUpdate,
		Delete: resourceSystemIpsUrlfilterDnsDelete,

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
			"address": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ipv6_capability": &schema.Schema{
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
	}
}

func resourceSystemIpsUrlfilterDnsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemIpsUrlfilterDns(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpsUrlfilterDns resource while getting object: %v", err)
	}

	_, err = c.CreateSystemIpsUrlfilterDns(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemIpsUrlfilterDns resource: %v", err)
	}

	d.SetId(getStringKey(d, "address"))

	return resourceSystemIpsUrlfilterDnsRead(d, m)
}

func resourceSystemIpsUrlfilterDnsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemIpsUrlfilterDns(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpsUrlfilterDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIpsUrlfilterDns(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpsUrlfilterDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "address"))

	return resourceSystemIpsUrlfilterDnsRead(d, m)
}

func resourceSystemIpsUrlfilterDnsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemIpsUrlfilterDns(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIpsUrlfilterDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpsUrlfilterDnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemIpsUrlfilterDns(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpsUrlfilterDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpsUrlfilterDns(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpsUrlfilterDns resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpsUrlfilterDnsAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpsUrlfilterDnsIpv6Capability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpsUrlfilterDnsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemIpsUrlfilterDns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenSystemIpsUrlfilterDnsAddress(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "SystemIpsUrlfilterDns-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("ipv6_capability", flattenSystemIpsUrlfilterDnsIpv6Capability(o["ipv6-capability"], d, "ipv6_capability")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-capability"], "SystemIpsUrlfilterDns-Ipv6Capability"); ok {
			if err = d.Set("ipv6_capability", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_capability: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_capability: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemIpsUrlfilterDnsStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemIpsUrlfilterDns-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemIpsUrlfilterDnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemIpsUrlfilterDnsAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpsUrlfilterDnsIpv6Capability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpsUrlfilterDnsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIpsUrlfilterDns(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandSystemIpsUrlfilterDnsAddress(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_capability"); ok || d.HasChange("ipv6_capability") {
		t, err := expandSystemIpsUrlfilterDnsIpv6Capability(d, v, "ipv6_capability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-capability"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemIpsUrlfilterDnsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
