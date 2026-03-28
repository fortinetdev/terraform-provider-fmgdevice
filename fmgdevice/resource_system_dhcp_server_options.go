// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DHCP options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDhcpServerOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcpServerOptionsCreate,
		Read:   resourceSystemDhcpServerOptionsRead,
		Update: resourceSystemDhcpServerOptionsUpdate,
		Delete: resourceSystemDhcpServerOptionsDelete,

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
			"code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uci_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uci_string": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vci_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceSystemDhcpServerOptionsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcpServerOptions(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcpServerOptions resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemDhcpServerOptions(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemDhcpServerOptions(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemDhcpServerOptions resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemDhcpServerOptions(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemDhcpServerOptions resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpServerOptionsRead(d, m)
}

func resourceSystemDhcpServerOptionsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcpServerOptions(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpServerOptions resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemDhcpServerOptions(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpServerOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpServerOptionsRead(d, m)
}

func resourceSystemDhcpServerOptionsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemDhcpServerOptions(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcpServerOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcpServerOptionsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDhcpServerOptions(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemDhcpServerOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcpServerOptions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcpServerOptions resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcpServerOptionsCode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerOptionsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsUciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsUciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpServerOptionsValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemDhcpServerOptions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("code", flattenSystemDhcpServerOptionsCode2edl(o["code"], d, "code")); err != nil {
		if vv, ok := fortiAPIPatch(o["code"], "SystemDhcpServerOptions-Code"); ok {
			if err = d.Set("code", vv); err != nil {
				return fmt.Errorf("Error reading code: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading code: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemDhcpServerOptionsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemDhcpServerOptions-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemDhcpServerOptionsIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemDhcpServerOptions-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemDhcpServerOptionsType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemDhcpServerOptions-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uci_match", flattenSystemDhcpServerOptionsUciMatch2edl(o["uci-match"], d, "uci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["uci-match"], "SystemDhcpServerOptions-UciMatch"); ok {
			if err = d.Set("uci_match", vv); err != nil {
				return fmt.Errorf("Error reading uci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uci_match: %v", err)
		}
	}

	if err = d.Set("uci_string", flattenSystemDhcpServerOptionsUciString2edl(o["uci-string"], d, "uci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["uci-string"], "SystemDhcpServerOptions-UciString"); ok {
			if err = d.Set("uci_string", vv); err != nil {
				return fmt.Errorf("Error reading uci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uci_string: %v", err)
		}
	}

	if err = d.Set("value", flattenSystemDhcpServerOptionsValue2edl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "SystemDhcpServerOptions-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenSystemDhcpServerOptionsVciMatch2edl(o["vci-match"], d, "vci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-match"], "SystemDhcpServerOptions-VciMatch"); ok {
			if err = d.Set("vci_match", vv); err != nil {
				return fmt.Errorf("Error reading vci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

	if err = d.Set("vci_string", flattenSystemDhcpServerOptionsVciString2edl(o["vci-string"], d, "vci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-string"], "SystemDhcpServerOptions-VciString"); ok {
			if err = d.Set("vci_string", vv); err != nil {
				return fmt.Errorf("Error reading vci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_string: %v", err)
		}
	}

	return nil
}

func flattenSystemDhcpServerOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDhcpServerOptionsCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerOptionsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsUciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsUciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpServerOptionsValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemDhcpServerOptions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("code"); ok || d.HasChange("code") {
		t, err := expandSystemDhcpServerOptionsCode2edl(d, v, "code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["code"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemDhcpServerOptionsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemDhcpServerOptionsIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemDhcpServerOptionsType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uci_match"); ok || d.HasChange("uci_match") {
		t, err := expandSystemDhcpServerOptionsUciMatch2edl(d, v, "uci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uci-match"] = t
		}
	}

	if v, ok := d.GetOk("uci_string"); ok || d.HasChange("uci_string") {
		t, err := expandSystemDhcpServerOptionsUciString2edl(d, v, "uci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uci-string"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandSystemDhcpServerOptionsValue2edl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok || d.HasChange("vci_match") {
		t, err := expandSystemDhcpServerOptionsVciMatch2edl(d, v, "vci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok || d.HasChange("vci_string") {
		t, err := expandSystemDhcpServerOptionsVciString2edl(d, v, "vci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	return &obj, nil
}
