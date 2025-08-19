// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPS URL filter cache settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterIpsUrlfilterCacheSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterIpsUrlfilterCacheSettingUpdate,
		Read:   resourceWebfilterIpsUrlfilterCacheSettingRead,
		Update: resourceWebfilterIpsUrlfilterCacheSettingUpdate,
		Delete: resourceWebfilterIpsUrlfilterCacheSettingDelete,

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
			"dns_retry_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"extended_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceWebfilterIpsUrlfilterCacheSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWebfilterIpsUrlfilterCacheSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterCacheSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateWebfilterIpsUrlfilterCacheSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterCacheSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WebfilterIpsUrlfilterCacheSetting")

	return resourceWebfilterIpsUrlfilterCacheSettingRead(d, m)
}

func resourceWebfilterIpsUrlfilterCacheSettingDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWebfilterIpsUrlfilterCacheSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterIpsUrlfilterCacheSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterIpsUrlfilterCacheSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterIpsUrlfilterCacheSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterIpsUrlfilterCacheSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterIpsUrlfilterCacheSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterIpsUrlfilterCacheSetting resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterIpsUrlfilterCacheSettingDnsRetryInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterIpsUrlfilterCacheSettingExtendedTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterIpsUrlfilterCacheSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dns_retry_interval", flattenWebfilterIpsUrlfilterCacheSettingDnsRetryInterval(o["dns-retry-interval"], d, "dns_retry_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-retry-interval"], "WebfilterIpsUrlfilterCacheSetting-DnsRetryInterval"); ok {
			if err = d.Set("dns_retry_interval", vv); err != nil {
				return fmt.Errorf("Error reading dns_retry_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_retry_interval: %v", err)
		}
	}

	if err = d.Set("extended_ttl", flattenWebfilterIpsUrlfilterCacheSettingExtendedTtl(o["extended-ttl"], d, "extended_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-ttl"], "WebfilterIpsUrlfilterCacheSetting-ExtendedTtl"); ok {
			if err = d.Set("extended_ttl", vv); err != nil {
				return fmt.Errorf("Error reading extended_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_ttl: %v", err)
		}
	}

	return nil
}

func flattenWebfilterIpsUrlfilterCacheSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterIpsUrlfilterCacheSettingDnsRetryInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterCacheSettingExtendedTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterIpsUrlfilterCacheSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dns_retry_interval"); ok || d.HasChange("dns_retry_interval") {
		t, err := expandWebfilterIpsUrlfilterCacheSettingDnsRetryInterval(d, v, "dns_retry_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-retry-interval"] = t
		}
	}

	if v, ok := d.GetOk("extended_ttl"); ok || d.HasChange("extended_ttl") {
		t, err := expandWebfilterIpsUrlfilterCacheSettingExtendedTtl(d, v, "extended_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-ttl"] = t
		}
	}

	return &obj, nil
}
