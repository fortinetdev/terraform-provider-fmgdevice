// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IGMP snooping VLAN.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchIgmpSnoopingVlans() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchIgmpSnoopingVlansCreate,
		Read:   resourceSwitchControllerManagedSwitchIgmpSnoopingVlansRead,
		Update: resourceSwitchControllerManagedSwitchIgmpSnoopingVlansUpdate,
		Delete: resourceSwitchControllerManagedSwitchIgmpSnoopingVlansDelete,

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"querier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"querier_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchIgmpSnoopingVlansCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchIgmpSnoopingVlans(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchIgmpSnoopingVlans resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerManagedSwitchIgmpSnoopingVlans(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchIgmpSnoopingVlans resource: %v", err)
	}

	d.SetId(getStringKey(d, "vlan_name"))

	return resourceSwitchControllerManagedSwitchIgmpSnoopingVlansRead(d, m)
}

func resourceSwitchControllerManagedSwitchIgmpSnoopingVlansUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchIgmpSnoopingVlans(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchIgmpSnoopingVlans resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchIgmpSnoopingVlans(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchIgmpSnoopingVlans resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "vlan_name"))

	return resourceSwitchControllerManagedSwitchIgmpSnoopingVlansRead(d, m)
}

func resourceSwitchControllerManagedSwitchIgmpSnoopingVlansDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerManagedSwitchIgmpSnoopingVlans(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchIgmpSnoopingVlans resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchIgmpSnoopingVlansRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	managed_switch := d.Get("managed_switch").(string)
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
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	o, err := c.ReadSwitchControllerManagedSwitchIgmpSnoopingVlans(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchIgmpSnoopingVlans resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchIgmpSnoopingVlans(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchIgmpSnoopingVlans resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansProxy3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVersion3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchIgmpSnoopingVlans(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("proxy", flattenSwitchControllerManagedSwitchIgmpSnoopingVlansProxy3rdl(o["proxy"], d, "proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy"], "SwitchControllerManagedSwitchIgmpSnoopingVlans-Proxy"); ok {
			if err = d.Set("proxy", vv); err != nil {
				return fmt.Errorf("Error reading proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("querier", flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier3rdl(o["querier"], d, "querier")); err != nil {
		if vv, ok := fortiAPIPatch(o["querier"], "SwitchControllerManagedSwitchIgmpSnoopingVlans-Querier"); ok {
			if err = d.Set("querier", vv); err != nil {
				return fmt.Errorf("Error reading querier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading querier: %v", err)
		}
	}

	if err = d.Set("querier_addr", flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr3rdl(o["querier-addr"], d, "querier_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["querier-addr"], "SwitchControllerManagedSwitchIgmpSnoopingVlans-QuerierAddr"); ok {
			if err = d.Set("querier_addr", vv); err != nil {
				return fmt.Errorf("Error reading querier_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading querier_addr: %v", err)
		}
	}

	if err = d.Set("version", flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVersion3rdl(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "SwitchControllerManagedSwitchIgmpSnoopingVlans-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if err = d.Set("vlan_name", flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName3rdl(o["vlan-name"], d, "vlan_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-name"], "SwitchControllerManagedSwitchIgmpSnoopingVlans-VlanName"); ok {
			if err = d.Set("vlan_name", vv); err != nil {
				return fmt.Errorf("Error reading vlan_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_name: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansProxy3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansVersion3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchIgmpSnoopingVlans(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("proxy"); ok || d.HasChange("proxy") {
		t, err := expandSwitchControllerManagedSwitchIgmpSnoopingVlansProxy3rdl(d, v, "proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	if v, ok := d.GetOk("querier"); ok || d.HasChange("querier") {
		t, err := expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier3rdl(d, v, "querier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["querier"] = t
		}
	}

	if v, ok := d.GetOk("querier_addr"); ok || d.HasChange("querier_addr") {
		t, err := expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr3rdl(d, v, "querier_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["querier-addr"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandSwitchControllerManagedSwitchIgmpSnoopingVlansVersion3rdl(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	if v, ok := d.GetOk("vlan_name"); ok || d.HasChange("vlan_name") {
		t, err := expandSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName3rdl(d, v, "vlan_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-name"] = t
		}
	}

	return &obj, nil
}
