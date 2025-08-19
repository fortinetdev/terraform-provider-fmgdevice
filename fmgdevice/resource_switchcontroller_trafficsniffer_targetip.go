// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Sniffer IPs to filter.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerTrafficSnifferTargetIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerTrafficSnifferTargetIpCreate,
		Read:   resourceSwitchControllerTrafficSnifferTargetIpRead,
		Update: resourceSwitchControllerTrafficSnifferTargetIpUpdate,
		Delete: resourceSwitchControllerTrafficSnifferTargetIpDelete,

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
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"src_entry_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerTrafficSnifferTargetIpCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerTrafficSnifferTargetIp(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerTrafficSnifferTargetIp resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerTrafficSnifferTargetIp(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerTrafficSnifferTargetIp resource: %v", err)
	}

	d.SetId(getStringKey(d, "ip"))

	return resourceSwitchControllerTrafficSnifferTargetIpRead(d, m)
}

func resourceSwitchControllerTrafficSnifferTargetIpUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerTrafficSnifferTargetIp(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficSnifferTargetIp resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerTrafficSnifferTargetIp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficSnifferTargetIp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "ip"))

	return resourceSwitchControllerTrafficSnifferTargetIpRead(d, m)
}

func resourceSwitchControllerTrafficSnifferTargetIpDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerTrafficSnifferTargetIp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerTrafficSnifferTargetIp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerTrafficSnifferTargetIpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerTrafficSnifferTargetIp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficSnifferTargetIp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerTrafficSnifferTargetIp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficSnifferTargetIp resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerTrafficSnifferTargetIpDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetIpDstEntryId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetIpIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetIpSrcEntryId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerTrafficSnifferTargetIp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("description", flattenSwitchControllerTrafficSnifferTargetIpDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerTrafficSnifferTargetIp-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("dst_entry_id", flattenSwitchControllerTrafficSnifferTargetIpDstEntryId2edl(o["dst-entry-id"], d, "dst_entry_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-entry-id"], "SwitchControllerTrafficSnifferTargetIp-DstEntryId"); ok {
			if err = d.Set("dst_entry_id", vv); err != nil {
				return fmt.Errorf("Error reading dst_entry_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_entry_id: %v", err)
		}
	}

	if err = d.Set("ip", flattenSwitchControllerTrafficSnifferTargetIpIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SwitchControllerTrafficSnifferTargetIp-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("src_entry_id", flattenSwitchControllerTrafficSnifferTargetIpSrcEntryId2edl(o["src-entry-id"], d, "src_entry_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-entry-id"], "SwitchControllerTrafficSnifferTargetIp-SrcEntryId"); ok {
			if err = d.Set("src_entry_id", vv); err != nil {
				return fmt.Errorf("Error reading src_entry_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_entry_id: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerTrafficSnifferTargetIpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerTrafficSnifferTargetIpDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetIpDstEntryId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetIpIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetIpSrcEntryId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerTrafficSnifferTargetIp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerTrafficSnifferTargetIpDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dst_entry_id"); ok || d.HasChange("dst_entry_id") {
		t, err := expandSwitchControllerTrafficSnifferTargetIpDstEntryId2edl(d, v, "dst_entry_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-entry-id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSwitchControllerTrafficSnifferTargetIpIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("src_entry_id"); ok || d.HasChange("src_entry_id") {
		t, err := expandSwitchControllerTrafficSnifferTargetIpSrcEntryId2edl(d, v, "src_entry_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-entry-id"] = t
		}
	}

	return &obj, nil
}
