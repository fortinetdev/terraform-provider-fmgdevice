// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Publicly accessible IP addresses for the FortiGSLB service.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallVipGslbPublicIps() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVipGslbPublicIpsCreate,
		Read:   resourceFirewallVipGslbPublicIpsRead,
		Update: resourceFirewallVipGslbPublicIpsUpdate,
		Delete: resourceFirewallVipGslbPublicIpsDelete,

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
			"vip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallVipGslbPublicIpsCreate(d *schema.ResourceData, m interface{}) error {
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
	vip := d.Get("vip").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip"] = vip

	obj, err := getObjectFirewallVipGslbPublicIps(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVipGslbPublicIps resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("index")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallVipGslbPublicIps(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallVipGslbPublicIps(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallVipGslbPublicIps resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallVipGslbPublicIps(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallVipGslbPublicIps resource: %v", err)
		}

		if v != nil && v["index"] != nil {
			if vidn, ok := v["index"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallVipGslbPublicIpsRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallVipGslbPublicIps resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceFirewallVipGslbPublicIpsRead(d, m)
}

func resourceFirewallVipGslbPublicIpsUpdate(d *schema.ResourceData, m interface{}) error {
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
	vip := d.Get("vip").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip"] = vip

	obj, err := getObjectFirewallVipGslbPublicIps(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVipGslbPublicIps resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallVipGslbPublicIps(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVipGslbPublicIps resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceFirewallVipGslbPublicIpsRead(d, m)
}

func resourceFirewallVipGslbPublicIpsDelete(d *schema.ResourceData, m interface{}) error {
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
	vip := d.Get("vip").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip"] = vip

	wsParams["adom"] = adomv

	err = c.DeleteFirewallVipGslbPublicIps(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVipGslbPublicIps resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVipGslbPublicIpsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	vip := d.Get("vip").(string)
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
	if vip == "" {
		vip = importOptionChecking(m.(*FortiClient).Cfg, "vip")
		if vip == "" {
			return fmt.Errorf("Parameter vip is missing")
		}
		if err = d.Set("vip", vip); err != nil {
			return fmt.Errorf("Error set params vip: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip"] = vip

	o, err := c.ReadFirewallVipGslbPublicIps(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallVipGslbPublicIps resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVipGslbPublicIps(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVipGslbPublicIps resource from API: %v", err)
	}
	return nil
}

func flattenFirewallVipGslbPublicIpsIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipGslbPublicIpsIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallVipGslbPublicIps(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("index", flattenFirewallVipGslbPublicIpsIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "FirewallVipGslbPublicIps-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("ip", flattenFirewallVipGslbPublicIpsIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FirewallVipGslbPublicIps-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	return nil
}

func flattenFirewallVipGslbPublicIpsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallVipGslbPublicIpsIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipGslbPublicIpsIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallVipGslbPublicIps(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandFirewallVipGslbPublicIpsIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandFirewallVipGslbPublicIpsIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	return &obj, nil
}
