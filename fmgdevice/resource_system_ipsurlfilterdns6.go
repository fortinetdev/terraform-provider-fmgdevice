// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPS URL filter IPv6 DNS servers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIpsUrlfilterDns6() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpsUrlfilterDns6Create,
		Read:   resourceSystemIpsUrlfilterDns6Read,
		Update: resourceSystemIpsUrlfilterDns6Update,
		Delete: resourceSystemIpsUrlfilterDns6Delete,

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
			"address6": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemIpsUrlfilterDns6Create(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemIpsUrlfilterDns6(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpsUrlfilterDns6 resource while getting object: %v", err)
	}

	_, err = c.CreateSystemIpsUrlfilterDns6(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpsUrlfilterDns6 resource: %v", err)
	}

	d.SetId(getStringKey(d, "address6"))

	return resourceSystemIpsUrlfilterDns6Read(d, m)
}

func resourceSystemIpsUrlfilterDns6Update(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemIpsUrlfilterDns6(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpsUrlfilterDns6 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIpsUrlfilterDns6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpsUrlfilterDns6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "address6"))

	return resourceSystemIpsUrlfilterDns6Read(d, m)
}

func resourceSystemIpsUrlfilterDns6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemIpsUrlfilterDns6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIpsUrlfilterDns6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpsUrlfilterDns6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemIpsUrlfilterDns6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpsUrlfilterDns6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpsUrlfilterDns6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpsUrlfilterDns6 resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpsUrlfilterDns6Address6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpsUrlfilterDns6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemIpsUrlfilterDns6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address6", flattenSystemIpsUrlfilterDns6Address6(o["address6"], d, "address6")); err != nil {
		if vv, ok := fortiAPIPatch(o["address6"], "SystemIpsUrlfilterDns6-Address6"); ok {
			if err = d.Set("address6", vv); err != nil {
				return fmt.Errorf("Error reading address6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address6: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemIpsUrlfilterDns6Status(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemIpsUrlfilterDns6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemIpsUrlfilterDns6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemIpsUrlfilterDns6Address6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpsUrlfilterDns6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIpsUrlfilterDns6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address6"); ok || d.HasChange("address6") {
		t, err := expandSystemIpsUrlfilterDns6Address6(d, v, "address6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address6"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemIpsUrlfilterDns6Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
