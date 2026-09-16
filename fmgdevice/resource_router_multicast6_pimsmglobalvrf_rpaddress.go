// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Statically configured RP addresses.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticast6PimSmGlobalVrfRpAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticast6PimSmGlobalVrfRpAddressCreate,
		Read:   resourceRouterMulticast6PimSmGlobalVrfRpAddressRead,
		Update: resourceRouterMulticast6PimSmGlobalVrfRpAddressUpdate,
		Delete: resourceRouterMulticast6PimSmGlobalVrfRpAddressDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

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
			"ip6_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceRouterMulticast6PimSmGlobalVrfRpAddressCreate(d *schema.ResourceData, m interface{}) error {
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
	pim_sm_global_vrf := d.Get("pim_sm_global_vrf").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["pim_sm_global_vrf"] = pim_sm_global_vrf

	obj, err := getObjectRouterMulticast6PimSmGlobalVrfRpAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticast6PimSmGlobalVrfRpAddress resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterMulticast6PimSmGlobalVrfRpAddress(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterMulticast6PimSmGlobalVrfRpAddress(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterMulticast6PimSmGlobalVrfRpAddress resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateRouterMulticast6PimSmGlobalVrfRpAddress(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterMulticast6PimSmGlobalVrfRpAddress resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterMulticast6PimSmGlobalVrfRpAddressRead(d, m)
}

func resourceRouterMulticast6PimSmGlobalVrfRpAddressUpdate(d *schema.ResourceData, m interface{}) error {
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
	pim_sm_global_vrf := d.Get("pim_sm_global_vrf").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["pim_sm_global_vrf"] = pim_sm_global_vrf

	obj, err := getObjectRouterMulticast6PimSmGlobalVrfRpAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6PimSmGlobalVrfRpAddress resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterMulticast6PimSmGlobalVrfRpAddress(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6PimSmGlobalVrfRpAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterMulticast6PimSmGlobalVrfRpAddressRead(d, m)
}

func resourceRouterMulticast6PimSmGlobalVrfRpAddressDelete(d *schema.ResourceData, m interface{}) error {
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
	pim_sm_global_vrf := d.Get("pim_sm_global_vrf").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["pim_sm_global_vrf"] = pim_sm_global_vrf

	wsParams["adom"] = adomv

	err = c.DeleteRouterMulticast6PimSmGlobalVrfRpAddress(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticast6PimSmGlobalVrfRpAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticast6PimSmGlobalVrfRpAddressRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterMulticast6PimSmGlobalVrfRpAddress(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterMulticast6PimSmGlobalVrfRpAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticast6PimSmGlobalVrfRpAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast6PimSmGlobalVrfRpAddress resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressGroup3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressIp6Address3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticast6PimSmGlobalVrfRpAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("group", flattenRouterMulticast6PimSmGlobalVrfRpAddressGroup3rdl(o["group"], d, "group")); err != nil {
		if vv, ok := fortiAPIPatch(o["group"], "RouterMulticast6PimSmGlobalVrfRpAddress-Group"); ok {
			if err = d.Set("group", vv); err != nil {
				return fmt.Errorf("Error reading group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterMulticast6PimSmGlobalVrfRpAddressId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterMulticast6PimSmGlobalVrfRpAddress-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip6_address", flattenRouterMulticast6PimSmGlobalVrfRpAddressIp6Address3rdl(o["ip6-address"], d, "ip6_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-address"], "RouterMulticast6PimSmGlobalVrfRpAddress-Ip6Address"); ok {
			if err = d.Set("ip6_address", vv); err != nil {
				return fmt.Errorf("Error reading ip6_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_address: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticast6PimSmGlobalVrfRpAddressGroup3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddressId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddressIp6Address3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticast6PimSmGlobalVrfRpAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("group"); ok || d.HasChange("group") {
		t, err := expandRouterMulticast6PimSmGlobalVrfRpAddressGroup3rdl(d, v, "group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterMulticast6PimSmGlobalVrfRpAddressId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip6_address"); ok || d.HasChange("ip6_address") {
		t, err := expandRouterMulticast6PimSmGlobalVrfRpAddressIp6Address3rdl(d, v, "ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-address"] = t
		}
	}

	return &obj, nil
}
