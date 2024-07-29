// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Sniffer MACs to filter.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerTrafficSnifferTargetMac() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerTrafficSnifferTargetMacCreate,
		Read:   resourceSwitchControllerTrafficSnifferTargetMacRead,
		Update: resourceSwitchControllerTrafficSnifferTargetMacUpdate,
		Delete: resourceSwitchControllerTrafficSnifferTargetMacDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_entry_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"src_entry_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerTrafficSnifferTargetMacCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSwitchControllerTrafficSnifferTargetMac(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerTrafficSnifferTargetMac resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerTrafficSnifferTargetMac(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerTrafficSnifferTargetMac resource: %v", err)
	}

	d.SetId(getStringKey(d, "mac"))

	return resourceSwitchControllerTrafficSnifferTargetMacRead(d, m)
}

func resourceSwitchControllerTrafficSnifferTargetMacUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSwitchControllerTrafficSnifferTargetMac(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficSnifferTargetMac resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerTrafficSnifferTargetMac(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficSnifferTargetMac resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "mac"))

	return resourceSwitchControllerTrafficSnifferTargetMacRead(d, m)
}

func resourceSwitchControllerTrafficSnifferTargetMacDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteSwitchControllerTrafficSnifferTargetMac(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerTrafficSnifferTargetMac resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerTrafficSnifferTargetMacRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerTrafficSnifferTargetMac(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficSnifferTargetMac resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerTrafficSnifferTargetMac(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficSnifferTargetMac resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerTrafficSnifferTargetMacDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetMacDstEntryId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetMacMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return case_insensitive(v, d.Get(pre))
}

func flattenSwitchControllerTrafficSnifferTargetMacSrcEntryId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerTrafficSnifferTargetMac(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("description", flattenSwitchControllerTrafficSnifferTargetMacDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerTrafficSnifferTargetMac-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("dst_entry_id", flattenSwitchControllerTrafficSnifferTargetMacDstEntryId2edl(o["dst-entry-id"], d, "dst_entry_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-entry-id"], "SwitchControllerTrafficSnifferTargetMac-DstEntryId"); ok {
			if err = d.Set("dst_entry_id", vv); err != nil {
				return fmt.Errorf("Error reading dst_entry_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_entry_id: %v", err)
		}
	}

	if err = d.Set("mac", flattenSwitchControllerTrafficSnifferTargetMacMac2edl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "SwitchControllerTrafficSnifferTargetMac-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("src_entry_id", flattenSwitchControllerTrafficSnifferTargetMacSrcEntryId2edl(o["src-entry-id"], d, "src_entry_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-entry-id"], "SwitchControllerTrafficSnifferTargetMac-SrcEntryId"); ok {
			if err = d.Set("src_entry_id", vv); err != nil {
				return fmt.Errorf("Error reading src_entry_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_entry_id: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerTrafficSnifferTargetMacFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerTrafficSnifferTargetMacDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetMacDstEntryId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetMacMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetMacSrcEntryId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerTrafficSnifferTargetMac(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerTrafficSnifferTargetMacDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dst_entry_id"); ok || d.HasChange("dst_entry_id") {
		t, err := expandSwitchControllerTrafficSnifferTargetMacDstEntryId2edl(d, v, "dst_entry_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-entry-id"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandSwitchControllerTrafficSnifferTargetMacMac2edl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("src_entry_id"); ok || d.HasChange("src_entry_id") {
		t, err := expandSwitchControllerTrafficSnifferTargetMacSrcEntryId2edl(d, v, "src_entry_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-entry-id"] = t
		}
	}

	return &obj, nil
}
