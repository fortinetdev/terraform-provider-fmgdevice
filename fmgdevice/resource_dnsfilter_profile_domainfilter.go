// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Domain filter settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDnsfilterProfileDomainFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsfilterProfileDomainFilterUpdate,
		Read:   resourceDnsfilterProfileDomainFilterRead,
		Update: resourceDnsfilterProfileDomainFilterUpdate,
		Delete: resourceDnsfilterProfileDomainFilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain_filter_table": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDnsfilterProfileDomainFilterUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectDnsfilterProfileDomainFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating DnsfilterProfileDomainFilter resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDnsfilterProfileDomainFilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DnsfilterProfileDomainFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DnsfilterProfileDomainFilter")

	return resourceDnsfilterProfileDomainFilterRead(d, m)
}

func resourceDnsfilterProfileDomainFilterDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteDnsfilterProfileDomainFilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DnsfilterProfileDomainFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDnsfilterProfileDomainFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadDnsfilterProfileDomainFilter(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DnsfilterProfileDomainFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDnsfilterProfileDomainFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DnsfilterProfileDomainFilter resource from API: %v", err)
	}
	return nil
}

func flattenDnsfilterProfileDomainFilterDomainFilterTable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectDnsfilterProfileDomainFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("domain_filter_table", flattenDnsfilterProfileDomainFilterDomainFilterTable2edl(o["domain-filter-table"], d, "domain_filter_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-filter-table"], "DnsfilterProfileDomainFilter-DomainFilterTable"); ok {
			if err = d.Set("domain_filter_table", vv); err != nil {
				return fmt.Errorf("Error reading domain_filter_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_filter_table: %v", err)
		}
	}

	return nil
}

func flattenDnsfilterProfileDomainFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDnsfilterProfileDomainFilterDomainFilterTable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectDnsfilterProfileDomainFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("domain_filter_table"); ok || d.HasChange("domain_filter_table") {
		t, err := expandDnsfilterProfileDomainFilterDomainFilterTable2edl(d, v, "domain_filter_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-filter-table"] = t
		}
	}

	return &obj, nil
}
