// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure decrypted traffic mirror.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallDecryptedTrafficMirror() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallDecryptedTrafficMirrorCreate,
		Read:   resourceFirewallDecryptedTrafficMirrorRead,
		Update: resourceFirewallDecryptedTrafficMirrorUpdate,
		Delete: resourceFirewallDecryptedTrafficMirrorDelete,

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
			"dstmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallDecryptedTrafficMirrorCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectFirewallDecryptedTrafficMirror(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallDecryptedTrafficMirror resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallDecryptedTrafficMirror(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallDecryptedTrafficMirror(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallDecryptedTrafficMirror resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallDecryptedTrafficMirror(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallDecryptedTrafficMirror resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallDecryptedTrafficMirrorRead(d, m)
}

func resourceFirewallDecryptedTrafficMirrorUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectFirewallDecryptedTrafficMirror(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDecryptedTrafficMirror resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallDecryptedTrafficMirror(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDecryptedTrafficMirror resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallDecryptedTrafficMirrorRead(d, m)
}

func resourceFirewallDecryptedTrafficMirrorDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteFirewallDecryptedTrafficMirror(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallDecryptedTrafficMirror resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallDecryptedTrafficMirrorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadFirewallDecryptedTrafficMirror(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallDecryptedTrafficMirror resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallDecryptedTrafficMirror(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallDecryptedTrafficMirror resource from API: %v", err)
	}
	return nil
}

func flattenFirewallDecryptedTrafficMirrorDstmac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDecryptedTrafficMirrorInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallDecryptedTrafficMirrorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDecryptedTrafficMirrorTrafficSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDecryptedTrafficMirrorTrafficType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallDecryptedTrafficMirror(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dstmac", flattenFirewallDecryptedTrafficMirrorDstmac(o["dstmac"], d, "dstmac")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstmac"], "FirewallDecryptedTrafficMirror-Dstmac"); ok {
			if err = d.Set("dstmac", vv); err != nil {
				return fmt.Errorf("Error reading dstmac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstmac: %v", err)
		}
	}

	if err = d.Set("interface", flattenFirewallDecryptedTrafficMirrorInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "FirewallDecryptedTrafficMirror-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallDecryptedTrafficMirrorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallDecryptedTrafficMirror-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("traffic_source", flattenFirewallDecryptedTrafficMirrorTrafficSource(o["traffic-source"], d, "traffic_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-source"], "FirewallDecryptedTrafficMirror-TrafficSource"); ok {
			if err = d.Set("traffic_source", vv); err != nil {
				return fmt.Errorf("Error reading traffic_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_source: %v", err)
		}
	}

	if err = d.Set("traffic_type", flattenFirewallDecryptedTrafficMirrorTrafficType(o["traffic-type"], d, "traffic_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-type"], "FirewallDecryptedTrafficMirror-TrafficType"); ok {
			if err = d.Set("traffic_type", vv); err != nil {
				return fmt.Errorf("Error reading traffic_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_type: %v", err)
		}
	}

	return nil
}

func flattenFirewallDecryptedTrafficMirrorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallDecryptedTrafficMirrorDstmac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDecryptedTrafficMirrorInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallDecryptedTrafficMirrorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDecryptedTrafficMirrorTrafficSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDecryptedTrafficMirrorTrafficType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallDecryptedTrafficMirror(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dstmac"); ok || d.HasChange("dstmac") {
		t, err := expandFirewallDecryptedTrafficMirrorDstmac(d, v, "dstmac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstmac"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandFirewallDecryptedTrafficMirrorInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallDecryptedTrafficMirrorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("traffic_source"); ok || d.HasChange("traffic_source") {
		t, err := expandFirewallDecryptedTrafficMirrorTrafficSource(d, v, "traffic_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-source"] = t
		}
	}

	if v, ok := d.GetOk("traffic_type"); ok || d.HasChange("traffic_type") {
		t, err := expandFirewallDecryptedTrafficMirrorTrafficType(d, v, "traffic_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-type"] = t
		}
	}

	return &obj, nil
}
