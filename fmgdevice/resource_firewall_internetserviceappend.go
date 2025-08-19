// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure additional port mappings for Internet Services.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceAppend() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceAppendUpdate,
		Read:   resourceFirewallInternetServiceAppendRead,
		Update: resourceFirewallInternetServiceAppendUpdate,
		Delete: resourceFirewallInternetServiceAppendDelete,

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
			"addr_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"append_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceFirewallInternetServiceAppendUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectFirewallInternetServiceAppend(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceAppend resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallInternetServiceAppend(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceAppend resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallInternetServiceAppend")

	return resourceFirewallInternetServiceAppendRead(d, m)
}

func resourceFirewallInternetServiceAppendDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallInternetServiceAppend(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceAppend resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceAppendRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallInternetServiceAppend(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceAppend resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceAppend(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceAppend resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceAppendAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAppendAppendPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceAppendMatchPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceAppend(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_mode", flattenFirewallInternetServiceAppendAddrMode(o["addr-mode"], d, "addr_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-mode"], "FirewallInternetServiceAppend-AddrMode"); ok {
			if err = d.Set("addr_mode", vv); err != nil {
				return fmt.Errorf("Error reading addr_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("append_port", flattenFirewallInternetServiceAppendAppendPort(o["append-port"], d, "append_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["append-port"], "FirewallInternetServiceAppend-AppendPort"); ok {
			if err = d.Set("append_port", vv); err != nil {
				return fmt.Errorf("Error reading append_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading append_port: %v", err)
		}
	}

	if err = d.Set("match_port", flattenFirewallInternetServiceAppendMatchPort(o["match-port"], d, "match_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-port"], "FirewallInternetServiceAppend-MatchPort"); ok {
			if err = d.Set("match_port", vv); err != nil {
				return fmt.Errorf("Error reading match_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_port: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceAppendFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInternetServiceAppendAddrMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAppendAppendPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAppendMatchPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceAppend(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_mode"); ok || d.HasChange("addr_mode") {
		t, err := expandFirewallInternetServiceAppendAddrMode(d, v, "addr_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("append_port"); ok || d.HasChange("append_port") {
		t, err := expandFirewallInternetServiceAppendAppendPort(d, v, "append_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["append-port"] = t
		}
	}

	if v, ok := d.GetOk("match_port"); ok || d.HasChange("match_port") {
		t, err := expandFirewallInternetServiceAppendMatchPort(d, v, "match_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-port"] = t
		}
	}

	return &obj, nil
}
