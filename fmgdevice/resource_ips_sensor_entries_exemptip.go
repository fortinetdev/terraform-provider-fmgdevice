// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Traffic from selected source or destination IP addresses is exempt from this signature.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpsSensorEntriesExemptIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsSensorEntriesExemptIpCreate,
		Read:   resourceIpsSensorEntriesExemptIpRead,
		Update: resourceIpsSensorEntriesExemptIpUpdate,
		Delete: resourceIpsSensorEntriesExemptIpDelete,

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
			"sensor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dst_ip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"src_ip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIpsSensorEntriesExemptIpCreate(d *schema.ResourceData, m interface{}) error {
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
	sensor := d.Get("sensor").(string)
	entries := d.Get("entries").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["sensor"] = sensor
	paradict["entries"] = entries

	obj, err := getObjectIpsSensorEntriesExemptIp(d)
	if err != nil {
		return fmt.Errorf("Error creating IpsSensorEntriesExemptIp resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIpsSensorEntriesExemptIp(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIpsSensorEntriesExemptIp(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IpsSensorEntriesExemptIp resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateIpsSensorEntriesExemptIp(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IpsSensorEntriesExemptIp resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceIpsSensorEntriesExemptIpRead(d, m)
			} else {
				return fmt.Errorf("Error creating IpsSensorEntriesExemptIp resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceIpsSensorEntriesExemptIpRead(d, m)
}

func resourceIpsSensorEntriesExemptIpUpdate(d *schema.ResourceData, m interface{}) error {
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
	sensor := d.Get("sensor").(string)
	entries := d.Get("entries").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["sensor"] = sensor
	paradict["entries"] = entries

	obj, err := getObjectIpsSensorEntriesExemptIp(d)
	if err != nil {
		return fmt.Errorf("Error updating IpsSensorEntriesExemptIp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIpsSensorEntriesExemptIp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IpsSensorEntriesExemptIp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceIpsSensorEntriesExemptIpRead(d, m)
}

func resourceIpsSensorEntriesExemptIpDelete(d *schema.ResourceData, m interface{}) error {
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
	sensor := d.Get("sensor").(string)
	entries := d.Get("entries").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["sensor"] = sensor
	paradict["entries"] = entries

	wsParams["adom"] = adomv

	err = c.DeleteIpsSensorEntriesExemptIp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IpsSensorEntriesExemptIp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsSensorEntriesExemptIpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	sensor := d.Get("sensor").(string)
	entries := d.Get("entries").(string)
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
	if sensor == "" {
		sensor = importOptionChecking(m.(*FortiClient).Cfg, "sensor")
		if sensor == "" {
			return fmt.Errorf("Parameter sensor is missing")
		}
		if err = d.Set("sensor", sensor); err != nil {
			return fmt.Errorf("Error set params sensor: %v", err)
		}
	}
	if entries == "" {
		entries = importOptionChecking(m.(*FortiClient).Cfg, "entries")
		if entries == "" {
			return fmt.Errorf("Parameter entries is missing")
		}
		if err = d.Set("entries", entries); err != nil {
			return fmt.Errorf("Error set params entries: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["sensor"] = sensor
	paradict["entries"] = entries

	o, err := c.ReadIpsSensorEntriesExemptIp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IpsSensorEntriesExemptIp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsSensorEntriesExemptIp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IpsSensorEntriesExemptIp resource from API: %v", err)
	}
	return nil
}

func flattenIpsSensorEntriesExemptIpDstIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesExemptIpId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesExemptIpSrcIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectIpsSensorEntriesExemptIp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dst_ip", flattenIpsSensorEntriesExemptIpDstIp3rdl(o["dst-ip"], d, "dst_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-ip"], "IpsSensorEntriesExemptIp-DstIp"); ok {
			if err = d.Set("dst_ip", vv); err != nil {
				return fmt.Errorf("Error reading dst_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenIpsSensorEntriesExemptIpId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "IpsSensorEntriesExemptIp-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("src_ip", flattenIpsSensorEntriesExemptIpSrcIp3rdl(o["src-ip"], d, "src_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-ip"], "IpsSensorEntriesExemptIp-SrcIp"); ok {
			if err = d.Set("src_ip", vv); err != nil {
				return fmt.Errorf("Error reading src_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_ip: %v", err)
		}
	}

	return nil
}

func flattenIpsSensorEntriesExemptIpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIpsSensorEntriesExemptIpDstIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandIpsSensorEntriesExemptIpId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesExemptIpSrcIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectIpsSensorEntriesExemptIp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dst_ip"); ok || d.HasChange("dst_ip") {
		t, err := expandIpsSensorEntriesExemptIpDstIp3rdl(d, v, "dst_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandIpsSensorEntriesExemptIpId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("src_ip"); ok || d.HasChange("src_ip") {
		t, err := expandIpsSensorEntriesExemptIpSrcIp3rdl(d, v, "src_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-ip"] = t
		}
	}

	return &obj, nil
}
