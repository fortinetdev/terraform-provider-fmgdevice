// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DHCP prefix configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDhcp6ServerPrefixRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcp6ServerPrefixRangeCreate,
		Read:   resourceSystemDhcp6ServerPrefixRangeRead,
		Update: resourceSystemDhcp6ServerPrefixRangeUpdate,
		Delete: resourceSystemDhcp6ServerPrefixRangeDelete,

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
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"end_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"prefix_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDhcp6ServerPrefixRangeCreate(d *schema.ResourceData, m interface{}) error {
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
	server := d.Get("server").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["server"] = server

	obj, err := getObjectSystemDhcp6ServerPrefixRange(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcp6ServerPrefixRange resource while getting object: %v", err)
	}

	_, err = c.CreateSystemDhcp6ServerPrefixRange(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemDhcp6ServerPrefixRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcp6ServerPrefixRangeRead(d, m)
}

func resourceSystemDhcp6ServerPrefixRangeUpdate(d *schema.ResourceData, m interface{}) error {
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
	server := d.Get("server").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["server"] = server

	obj, err := getObjectSystemDhcp6ServerPrefixRange(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcp6ServerPrefixRange resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDhcp6ServerPrefixRange(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcp6ServerPrefixRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcp6ServerPrefixRangeRead(d, m)
}

func resourceSystemDhcp6ServerPrefixRangeDelete(d *schema.ResourceData, m interface{}) error {
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
	server := d.Get("server").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["server"] = server

	err = c.DeleteSystemDhcp6ServerPrefixRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcp6ServerPrefixRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcp6ServerPrefixRangeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	server := d.Get("server").(string)
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
	if server == "" {
		server = importOptionChecking(m.(*FortiClient).Cfg, "server")
		if server == "" {
			return fmt.Errorf("Parameter server is missing")
		}
		if err = d.Set("server", server); err != nil {
			return fmt.Errorf("Error set params server: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["server"] = server

	o, err := c.ReadSystemDhcp6ServerPrefixRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcp6ServerPrefixRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcp6ServerPrefixRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcp6ServerPrefixRange resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcp6ServerPrefixRangeEndPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRangePrefixLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRangeStartPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDhcp6ServerPrefixRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("end_prefix", flattenSystemDhcp6ServerPrefixRangeEndPrefix2edl(o["end-prefix"], d, "end_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-prefix"], "SystemDhcp6ServerPrefixRange-EndPrefix"); ok {
			if err = d.Set("end_prefix", vv); err != nil {
				return fmt.Errorf("Error reading end_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_prefix: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemDhcp6ServerPrefixRangeId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemDhcp6ServerPrefixRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("prefix_length", flattenSystemDhcp6ServerPrefixRangePrefixLength2edl(o["prefix-length"], d, "prefix_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-length"], "SystemDhcp6ServerPrefixRange-PrefixLength"); ok {
			if err = d.Set("prefix_length", vv); err != nil {
				return fmt.Errorf("Error reading prefix_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_length: %v", err)
		}
	}

	if err = d.Set("start_prefix", flattenSystemDhcp6ServerPrefixRangeStartPrefix2edl(o["start-prefix"], d, "start_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-prefix"], "SystemDhcp6ServerPrefixRange-StartPrefix"); ok {
			if err = d.Set("start_prefix", vv); err != nil {
				return fmt.Errorf("Error reading start_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_prefix: %v", err)
		}
	}

	return nil
}

func flattenSystemDhcp6ServerPrefixRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDhcp6ServerPrefixRangeEndPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRangePrefixLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRangeStartPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDhcp6ServerPrefixRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_prefix"); ok || d.HasChange("end_prefix") {
		t, err := expandSystemDhcp6ServerPrefixRangeEndPrefix2edl(d, v, "end_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-prefix"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemDhcp6ServerPrefixRangeId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("prefix_length"); ok || d.HasChange("prefix_length") {
		t, err := expandSystemDhcp6ServerPrefixRangePrefixLength2edl(d, v, "prefix_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-length"] = t
		}
	}

	if v, ok := d.GetOk("start_prefix"); ok || d.HasChange("start_prefix") {
		t, err := expandSystemDhcp6ServerPrefixRangeStartPrefix2edl(d, v, "start_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-prefix"] = t
		}
	}

	return &obj, nil
}
