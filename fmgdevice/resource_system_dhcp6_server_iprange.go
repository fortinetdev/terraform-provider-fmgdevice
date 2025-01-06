// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DHCP IP range configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDhcp6ServerIpRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcp6ServerIpRangeCreate,
		Read:   resourceSystemDhcp6ServerIpRangeRead,
		Update: resourceSystemDhcp6ServerIpRangeUpdate,
		Delete: resourceSystemDhcp6ServerIpRangeDelete,

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
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vci_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vci_string": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDhcp6ServerIpRangeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcp6ServerIpRange(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcp6ServerIpRange resource while getting object: %v", err)
	}

	_, err = c.CreateSystemDhcp6ServerIpRange(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemDhcp6ServerIpRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcp6ServerIpRangeRead(d, m)
}

func resourceSystemDhcp6ServerIpRangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcp6ServerIpRange(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcp6ServerIpRange resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDhcp6ServerIpRange(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcp6ServerIpRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcp6ServerIpRangeRead(d, m)
}

func resourceSystemDhcp6ServerIpRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemDhcp6ServerIpRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcp6ServerIpRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcp6ServerIpRangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDhcp6ServerIpRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcp6ServerIpRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcp6ServerIpRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcp6ServerIpRange resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcp6ServerIpRangeEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemDhcp6ServerIpRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("end_ip", flattenSystemDhcp6ServerIpRangeEndIp2edl(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "SystemDhcp6ServerIpRange-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemDhcp6ServerIpRangeId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemDhcp6ServerIpRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenSystemDhcp6ServerIpRangeStartIp2edl(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "SystemDhcp6ServerIpRange-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenSystemDhcp6ServerIpRangeVciMatch2edl(o["vci-match"], d, "vci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-match"], "SystemDhcp6ServerIpRange-VciMatch"); ok {
			if err = d.Set("vci_match", vv); err != nil {
				return fmt.Errorf("Error reading vci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

	if err = d.Set("vci_string", flattenSystemDhcp6ServerIpRangeVciString2edl(o["vci-string"], d, "vci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-string"], "SystemDhcp6ServerIpRange-VciString"); ok {
			if err = d.Set("vci_string", vv); err != nil {
				return fmt.Errorf("Error reading vci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_string: %v", err)
		}
	}

	return nil
}

func flattenSystemDhcp6ServerIpRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDhcp6ServerIpRangeEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemDhcp6ServerIpRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandSystemDhcp6ServerIpRangeEndIp2edl(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemDhcp6ServerIpRangeId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandSystemDhcp6ServerIpRangeStartIp2edl(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok || d.HasChange("vci_match") {
		t, err := expandSystemDhcp6ServerIpRangeVciMatch2edl(d, v, "vci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok || d.HasChange("vci_string") {
		t, err := expandSystemDhcp6ServerIpRangeVciString2edl(d, v, "vci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	return &obj, nil
}
