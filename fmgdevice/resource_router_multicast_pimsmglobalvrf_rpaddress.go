// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Statically configure RP addresses.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticastPimSmGlobalVrfRpAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticastPimSmGlobalVrfRpAddressCreate,
		Read:   resourceRouterMulticastPimSmGlobalVrfRpAddressRead,
		Update: resourceRouterMulticastPimSmGlobalVrfRpAddressUpdate,
		Delete: resourceRouterMulticastPimSmGlobalVrfRpAddressDelete,

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
			"pim_sm_global_vrf": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterMulticastPimSmGlobalVrfRpAddressCreate(d *schema.ResourceData, m interface{}) error {
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
	pim_sm_global_vrf := d.Get("pim_sm_global_vrf").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["pim_sm_global_vrf"] = pim_sm_global_vrf

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterMulticastPimSmGlobalVrfRpAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastPimSmGlobalVrfRpAddress resource while getting object: %v", err)
	}

	_, err = c.CreateRouterMulticastPimSmGlobalVrfRpAddress(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastPimSmGlobalVrfRpAddress resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterMulticastPimSmGlobalVrfRpAddressRead(d, m)
}

func resourceRouterMulticastPimSmGlobalVrfRpAddressUpdate(d *schema.ResourceData, m interface{}) error {
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
	pim_sm_global_vrf := d.Get("pim_sm_global_vrf").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["pim_sm_global_vrf"] = pim_sm_global_vrf

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterMulticastPimSmGlobalVrfRpAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastPimSmGlobalVrfRpAddress resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticastPimSmGlobalVrfRpAddress(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastPimSmGlobalVrfRpAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterMulticastPimSmGlobalVrfRpAddressRead(d, m)
}

func resourceRouterMulticastPimSmGlobalVrfRpAddressDelete(d *schema.ResourceData, m interface{}) error {
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
	pim_sm_global_vrf := d.Get("pim_sm_global_vrf").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["pim_sm_global_vrf"] = pim_sm_global_vrf

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteRouterMulticastPimSmGlobalVrfRpAddress(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticastPimSmGlobalVrfRpAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastPimSmGlobalVrfRpAddressRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	pim_sm_global_vrf := d.Get("pim_sm_global_vrf").(string)
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
	if pim_sm_global_vrf == "" {
		pim_sm_global_vrf = importOptionChecking(m.(*FortiClient).Cfg, "pim_sm_global_vrf")
		if pim_sm_global_vrf == "" {
			return fmt.Errorf("Parameter pim_sm_global_vrf is missing")
		}
		if err = d.Set("pim_sm_global_vrf", pim_sm_global_vrf); err != nil {
			return fmt.Errorf("Error set params pim_sm_global_vrf: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["pim_sm_global_vrf"] = pim_sm_global_vrf

	o, err := c.ReadRouterMulticastPimSmGlobalVrfRpAddress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastPimSmGlobalVrfRpAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticastPimSmGlobalVrfRpAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastPimSmGlobalVrfRpAddress resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticastPimSmGlobalVrfRpAddressGroup3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastPimSmGlobalVrfRpAddressId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfRpAddressIpAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticastPimSmGlobalVrfRpAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("group", flattenRouterMulticastPimSmGlobalVrfRpAddressGroup3rdl(o["group"], d, "group")); err != nil {
		if vv, ok := fortiAPIPatch(o["group"], "RouterMulticastPimSmGlobalVrfRpAddress-Group"); ok {
			if err = d.Set("group", vv); err != nil {
				return fmt.Errorf("Error reading group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterMulticastPimSmGlobalVrfRpAddressId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterMulticastPimSmGlobalVrfRpAddress-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_address", flattenRouterMulticastPimSmGlobalVrfRpAddressIpAddress3rdl(o["ip-address"], d, "ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-address"], "RouterMulticastPimSmGlobalVrfRpAddress-IpAddress"); ok {
			if err = d.Set("ip_address", vv); err != nil {
				return fmt.Errorf("Error reading ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticastPimSmGlobalVrfRpAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticastPimSmGlobalVrfRpAddressGroup3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastPimSmGlobalVrfRpAddressId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfRpAddressIpAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticastPimSmGlobalVrfRpAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("group"); ok || d.HasChange("group") {
		t, err := expandRouterMulticastPimSmGlobalVrfRpAddressGroup3rdl(d, v, "group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterMulticastPimSmGlobalVrfRpAddressId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_address"); ok || d.HasChange("ip_address") {
		t, err := expandRouterMulticastPimSmGlobalVrfRpAddressIpAddress3rdl(d, v, "ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	return &obj, nil
}
