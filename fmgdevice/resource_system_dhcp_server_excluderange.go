// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Exclude one or more ranges of IP addresses from being assigned to clients.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDhcpServerExcludeRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcpServerExcludeRangeCreate,
		Read:   resourceSystemDhcpServerExcludeRangeRead,
		Update: resourceSystemDhcpServerExcludeRangeUpdate,
		Delete: resourceSystemDhcpServerExcludeRangeDelete,

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
				Computed: true,
			},
			"lease_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uci_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uci_string": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceSystemDhcpServerExcludeRangeCreate(d *schema.ResourceData, m interface{}) error {
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
	server := d.Get("server").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["server"] = server

	obj, err := getObjectSystemDhcpServerExcludeRange(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcpServerExcludeRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemDhcpServerExcludeRange(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemDhcpServerExcludeRange(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemDhcpServerExcludeRange resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateSystemDhcpServerExcludeRange(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemDhcpServerExcludeRange resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceSystemDhcpServerExcludeRangeRead(d, m)
			} else {
				return fmt.Errorf("Error creating SystemDhcpServerExcludeRange resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpServerExcludeRangeRead(d, m)
}

func resourceSystemDhcpServerExcludeRangeUpdate(d *schema.ResourceData, m interface{}) error {
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
	server := d.Get("server").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["server"] = server

	obj, err := getObjectSystemDhcpServerExcludeRange(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpServerExcludeRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemDhcpServerExcludeRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpServerExcludeRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpServerExcludeRangeRead(d, m)
}

func resourceSystemDhcpServerExcludeRangeDelete(d *schema.ResourceData, m interface{}) error {
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
	server := d.Get("server").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["server"] = server

	wsParams["adom"] = adomv

	err = c.DeleteSystemDhcpServerExcludeRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcpServerExcludeRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcpServerExcludeRangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDhcpServerExcludeRange(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemDhcpServerExcludeRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcpServerExcludeRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcpServerExcludeRange resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcpServerExcludeRangeEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeLeaseTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeUciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeUciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerExcludeRangeVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemDhcpServerExcludeRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("end_ip", flattenSystemDhcpServerExcludeRangeEndIp2edl(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "SystemDhcpServerExcludeRange-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemDhcpServerExcludeRangeId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemDhcpServerExcludeRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenSystemDhcpServerExcludeRangeLeaseTime2edl(o["lease-time"], d, "lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["lease-time"], "SystemDhcpServerExcludeRange-LeaseTime"); ok {
			if err = d.Set("lease_time", vv); err != nil {
				return fmt.Errorf("Error reading lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenSystemDhcpServerExcludeRangeStartIp2edl(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "SystemDhcpServerExcludeRange-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("uci_match", flattenSystemDhcpServerExcludeRangeUciMatch2edl(o["uci-match"], d, "uci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["uci-match"], "SystemDhcpServerExcludeRange-UciMatch"); ok {
			if err = d.Set("uci_match", vv); err != nil {
				return fmt.Errorf("Error reading uci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uci_match: %v", err)
		}
	}

	if err = d.Set("uci_string", flattenSystemDhcpServerExcludeRangeUciString2edl(o["uci-string"], d, "uci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["uci-string"], "SystemDhcpServerExcludeRange-UciString"); ok {
			if err = d.Set("uci_string", vv); err != nil {
				return fmt.Errorf("Error reading uci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uci_string: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenSystemDhcpServerExcludeRangeVciMatch2edl(o["vci-match"], d, "vci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-match"], "SystemDhcpServerExcludeRange-VciMatch"); ok {
			if err = d.Set("vci_match", vv); err != nil {
				return fmt.Errorf("Error reading vci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

	if err = d.Set("vci_string", flattenSystemDhcpServerExcludeRangeVciString2edl(o["vci-string"], d, "vci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-string"], "SystemDhcpServerExcludeRange-VciString"); ok {
			if err = d.Set("vci_string", vv); err != nil {
				return fmt.Errorf("Error reading vci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_string: %v", err)
		}
	}

	return nil
}

func flattenSystemDhcpServerExcludeRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDhcpServerExcludeRangeEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeLeaseTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeUciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeUciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerExcludeRangeVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemDhcpServerExcludeRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandSystemDhcpServerExcludeRangeEndIp2edl(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemDhcpServerExcludeRangeId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok || d.HasChange("lease_time") {
		t, err := expandSystemDhcpServerExcludeRangeLeaseTime2edl(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandSystemDhcpServerExcludeRangeStartIp2edl(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("uci_match"); ok || d.HasChange("uci_match") {
		t, err := expandSystemDhcpServerExcludeRangeUciMatch2edl(d, v, "uci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uci-match"] = t
		}
	}

	if v, ok := d.GetOk("uci_string"); ok || d.HasChange("uci_string") {
		t, err := expandSystemDhcpServerExcludeRangeUciString2edl(d, v, "uci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uci-string"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok || d.HasChange("vci_match") {
		t, err := expandSystemDhcpServerExcludeRangeVciMatch2edl(d, v, "vci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok || d.HasChange("vci_string") {
		t, err := expandSystemDhcpServerExcludeRangeVciString2edl(d, v, "vci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	return &obj, nil
}
