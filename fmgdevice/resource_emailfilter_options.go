// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AntiSpam options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterOptionsUpdate,
		Read:   resourceEmailfilterOptionsRead,
		Update: resourceEmailfilterOptionsUpdate,
		Delete: resourceEmailfilterOptionsDelete,

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
			"dns_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceEmailfilterOptionsUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectEmailfilterOptions(d)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterOptions resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEmailfilterOptions(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("EmailfilterOptions")

	return resourceEmailfilterOptionsRead(d, m)
}

func resourceEmailfilterOptionsDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteEmailfilterOptions(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterOptionsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadEmailfilterOptions(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EmailfilterOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterOptions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterOptions resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterOptionsDnsTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEmailfilterOptions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dns_timeout", flattenEmailfilterOptionsDnsTimeout(o["dns-timeout"], d, "dns_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-timeout"], "EmailfilterOptions-DnsTimeout"); ok {
			if err = d.Set("dns_timeout", vv); err != nil {
				return fmt.Errorf("Error reading dns_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_timeout: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEmailfilterOptionsDnsTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterOptions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dns_timeout"); ok || d.HasChange("dns_timeout") {
		t, err := expandEmailfilterOptionsDnsTimeout(d, v, "dns_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-timeout"] = t
		}
	}

	return &obj, nil
}
