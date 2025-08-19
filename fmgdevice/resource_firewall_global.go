// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Global firewall settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallGlobalUpdate,
		Read:   resourceFirewallGlobalRead,
		Update: resourceFirewallGlobalUpdate,
		Delete: resourceFirewallGlobalDelete,

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
			"banned_ip_persistency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallGlobalUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectFirewallGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallGlobal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallGlobal")

	return resourceFirewallGlobalRead(d, m)
}

func resourceFirewallGlobalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallGlobal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallGlobal resource from API: %v", err)
	}
	return nil
}

func flattenFirewallGlobalBannedIpPersistency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("banned_ip_persistency", flattenFirewallGlobalBannedIpPersistency(o["banned-ip-persistency"], d, "banned_ip_persistency")); err != nil {
		if vv, ok := fortiAPIPatch(o["banned-ip-persistency"], "FirewallGlobal-BannedIpPersistency"); ok {
			if err = d.Set("banned_ip_persistency", vv); err != nil {
				return fmt.Errorf("Error reading banned_ip_persistency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading banned_ip_persistency: %v", err)
		}
	}

	return nil
}

func flattenFirewallGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallGlobalBannedIpPersistency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("banned_ip_persistency"); ok || d.HasChange("banned_ip_persistency") {
		t, err := expandFirewallGlobalBannedIpPersistency(d, v, "banned_ip_persistency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banned-ip-persistency"] = t
		}
	}

	return &obj, nil
}
