// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration method to edit FortiSwitch packet mirror.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchMirror() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchMirrorCreate,
		Read:   resourceSwitchControllerManagedSwitchMirrorRead,
		Update: resourceSwitchControllerManagedSwitchMirrorUpdate,
		Delete: resourceSwitchControllerManagedSwitchMirrorDelete,

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
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"src_egress": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"src_ingress": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switching_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchMirrorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectSwitchControllerManagedSwitchMirror(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchMirror resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerManagedSwitchMirror(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchMirror resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerManagedSwitchMirrorRead(d, m)
}

func resourceSwitchControllerManagedSwitchMirrorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectSwitchControllerManagedSwitchMirror(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchMirror resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchMirror(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchMirror resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerManagedSwitchMirrorRead(d, m)
}

func resourceSwitchControllerManagedSwitchMirrorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteSwitchControllerManagedSwitchMirror(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchMirror resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchMirrorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerManagedSwitchMirror(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchMirror resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchMirror(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchMirror resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchMirrorDst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirrorName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirrorSrcEgress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchMirrorSrcIngress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchMirrorStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirrorSwitchingPacket2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchMirror(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dst", flattenSwitchControllerManagedSwitchMirrorDst2edl(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "SwitchControllerManagedSwitchMirror-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerManagedSwitchMirrorName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerManagedSwitchMirror-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("src_egress", flattenSwitchControllerManagedSwitchMirrorSrcEgress2edl(o["src-egress"], d, "src_egress")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-egress"], "SwitchControllerManagedSwitchMirror-SrcEgress"); ok {
			if err = d.Set("src_egress", vv); err != nil {
				return fmt.Errorf("Error reading src_egress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_egress: %v", err)
		}
	}

	if err = d.Set("src_ingress", flattenSwitchControllerManagedSwitchMirrorSrcIngress2edl(o["src-ingress"], d, "src_ingress")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-ingress"], "SwitchControllerManagedSwitchMirror-SrcIngress"); ok {
			if err = d.Set("src_ingress", vv); err != nil {
				return fmt.Errorf("Error reading src_ingress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_ingress: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchControllerManagedSwitchMirrorStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SwitchControllerManagedSwitchMirror-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("switching_packet", flattenSwitchControllerManagedSwitchMirrorSwitchingPacket2edl(o["switching-packet"], d, "switching_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["switching-packet"], "SwitchControllerManagedSwitchMirror-SwitchingPacket"); ok {
			if err = d.Set("switching_packet", vv); err != nil {
				return fmt.Errorf("Error reading switching_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switching_packet: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchMirrorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchMirrorDst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirrorName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirrorSrcEgress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchMirrorSrcIngress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchMirrorStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirrorSwitchingPacket2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchMirror(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandSwitchControllerManagedSwitchMirrorDst2edl(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerManagedSwitchMirrorName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("src_egress"); ok || d.HasChange("src_egress") {
		t, err := expandSwitchControllerManagedSwitchMirrorSrcEgress2edl(d, v, "src_egress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-egress"] = t
		}
	}

	if v, ok := d.GetOk("src_ingress"); ok || d.HasChange("src_ingress") {
		t, err := expandSwitchControllerManagedSwitchMirrorSrcIngress2edl(d, v, "src_ingress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-ingress"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSwitchControllerManagedSwitchMirrorStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("switching_packet"); ok || d.HasChange("switching_packet") {
		t, err := expandSwitchControllerManagedSwitchMirrorSwitchingPacket2edl(d, v, "switching_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switching-packet"] = t
		}
	}

	return &obj, nil
}
