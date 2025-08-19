// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Maps between IP-DSCP value to COS queue.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerQosIpDscpMapMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerQosIpDscpMapMapCreate,
		Read:   resourceSwitchControllerQosIpDscpMapMapRead,
		Update: resourceSwitchControllerQosIpDscpMapMapUpdate,
		Delete: resourceSwitchControllerQosIpDscpMapMapDelete,

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
			"ip_dscp_map": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cos_queue": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"diffserv": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_precedence": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerQosIpDscpMapMapCreate(d *schema.ResourceData, m interface{}) error {
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
	ip_dscp_map := d.Get("ip_dscp_map").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ip_dscp_map"] = ip_dscp_map

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerQosIpDscpMapMap(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosIpDscpMapMap resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerQosIpDscpMapMap(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosIpDscpMapMap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerQosIpDscpMapMapRead(d, m)
}

func resourceSwitchControllerQosIpDscpMapMapUpdate(d *schema.ResourceData, m interface{}) error {
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
	ip_dscp_map := d.Get("ip_dscp_map").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ip_dscp_map"] = ip_dscp_map

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerQosIpDscpMapMap(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosIpDscpMapMap resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerQosIpDscpMapMap(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosIpDscpMapMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerQosIpDscpMapMapRead(d, m)
}

func resourceSwitchControllerQosIpDscpMapMapDelete(d *schema.ResourceData, m interface{}) error {
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
	ip_dscp_map := d.Get("ip_dscp_map").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ip_dscp_map"] = ip_dscp_map

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerQosIpDscpMapMap(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerQosIpDscpMapMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerQosIpDscpMapMapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	ip_dscp_map := d.Get("ip_dscp_map").(string)
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
	if ip_dscp_map == "" {
		ip_dscp_map = importOptionChecking(m.(*FortiClient).Cfg, "ip_dscp_map")
		if ip_dscp_map == "" {
			return fmt.Errorf("Parameter ip_dscp_map is missing")
		}
		if err = d.Set("ip_dscp_map", ip_dscp_map); err != nil {
			return fmt.Errorf("Error set params ip_dscp_map: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ip_dscp_map"] = ip_dscp_map

	o, err := c.ReadSwitchControllerQosIpDscpMapMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosIpDscpMapMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerQosIpDscpMapMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosIpDscpMapMap resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerQosIpDscpMapMapCosQueue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosIpDscpMapMapDiffserv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerQosIpDscpMapMapIpPrecedence2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerQosIpDscpMapMapName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosIpDscpMapMapValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerQosIpDscpMapMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cos_queue", flattenSwitchControllerQosIpDscpMapMapCosQueue2edl(o["cos-queue"], d, "cos_queue")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos-queue"], "SwitchControllerQosIpDscpMapMap-CosQueue"); ok {
			if err = d.Set("cos_queue", vv); err != nil {
				return fmt.Errorf("Error reading cos_queue: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos_queue: %v", err)
		}
	}

	if err = d.Set("diffserv", flattenSwitchControllerQosIpDscpMapMapDiffserv2edl(o["diffserv"], d, "diffserv")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv"], "SwitchControllerQosIpDscpMapMap-Diffserv"); ok {
			if err = d.Set("diffserv", vv); err != nil {
				return fmt.Errorf("Error reading diffserv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv: %v", err)
		}
	}

	if err = d.Set("ip_precedence", flattenSwitchControllerQosIpDscpMapMapIpPrecedence2edl(o["ip-precedence"], d, "ip_precedence")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-precedence"], "SwitchControllerQosIpDscpMapMap-IpPrecedence"); ok {
			if err = d.Set("ip_precedence", vv); err != nil {
				return fmt.Errorf("Error reading ip_precedence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_precedence: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerQosIpDscpMapMapName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerQosIpDscpMapMap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenSwitchControllerQosIpDscpMapMapValue2edl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "SwitchControllerQosIpDscpMapMap-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerQosIpDscpMapMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerQosIpDscpMapMapCosQueue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosIpDscpMapMapDiffserv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerQosIpDscpMapMapIpPrecedence2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerQosIpDscpMapMapName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosIpDscpMapMapValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerQosIpDscpMapMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cos_queue"); ok || d.HasChange("cos_queue") {
		t, err := expandSwitchControllerQosIpDscpMapMapCosQueue2edl(d, v, "cos_queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-queue"] = t
		}
	}

	if v, ok := d.GetOk("diffserv"); ok || d.HasChange("diffserv") {
		t, err := expandSwitchControllerQosIpDscpMapMapDiffserv2edl(d, v, "diffserv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv"] = t
		}
	}

	if v, ok := d.GetOk("ip_precedence"); ok || d.HasChange("ip_precedence") {
		t, err := expandSwitchControllerQosIpDscpMapMapIpPrecedence2edl(d, v, "ip_precedence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-precedence"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerQosIpDscpMapMapName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandSwitchControllerQosIpDscpMapMapValue2edl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
