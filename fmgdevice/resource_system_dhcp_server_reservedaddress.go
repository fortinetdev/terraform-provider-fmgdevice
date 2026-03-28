// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Options for the DHCP server to assign IP settings to specific MAC addresses.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDhcpServerReservedAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcpServerReservedAddressCreate,
		Read:   resourceSystemDhcpServerReservedAddressRead,
		Update: resourceSystemDhcpServerReservedAddressUpdate,
		Delete: resourceSystemDhcpServerReservedAddressDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"circuit_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"circuit_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemDhcpServerReservedAddressCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcpServerReservedAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcpServerReservedAddress resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemDhcpServerReservedAddress(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemDhcpServerReservedAddress(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemDhcpServerReservedAddress resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemDhcpServerReservedAddress(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemDhcpServerReservedAddress resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpServerReservedAddressRead(d, m)
}

func resourceSystemDhcpServerReservedAddressUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcpServerReservedAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpServerReservedAddress resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemDhcpServerReservedAddress(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpServerReservedAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpServerReservedAddressRead(d, m)
}

func resourceSystemDhcpServerReservedAddressDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemDhcpServerReservedAddress(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcpServerReservedAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcpServerReservedAddressRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDhcpServerReservedAddress(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemDhcpServerReservedAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcpServerReservedAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcpServerReservedAddress resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcpServerReservedAddressAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressCircuitId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressCircuitIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressRemoteId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressRemoteIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDhcpServerReservedAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenSystemDhcpServerReservedAddressAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SystemDhcpServerReservedAddress-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("circuit_id", flattenSystemDhcpServerReservedAddressCircuitId2edl(o["circuit-id"], d, "circuit_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["circuit-id"], "SystemDhcpServerReservedAddress-CircuitId"); ok {
			if err = d.Set("circuit_id", vv); err != nil {
				return fmt.Errorf("Error reading circuit_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading circuit_id: %v", err)
		}
	}

	if err = d.Set("circuit_id_type", flattenSystemDhcpServerReservedAddressCircuitIdType2edl(o["circuit-id-type"], d, "circuit_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["circuit-id-type"], "SystemDhcpServerReservedAddress-CircuitIdType"); ok {
			if err = d.Set("circuit_id_type", vv); err != nil {
				return fmt.Errorf("Error reading circuit_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading circuit_id_type: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemDhcpServerReservedAddressDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemDhcpServerReservedAddress-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemDhcpServerReservedAddressId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemDhcpServerReservedAddress-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemDhcpServerReservedAddressIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemDhcpServerReservedAddress-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mac", flattenSystemDhcpServerReservedAddressMac2edl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "SystemDhcpServerReservedAddress-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("remote_id", flattenSystemDhcpServerReservedAddressRemoteId2edl(o["remote-id"], d, "remote_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-id"], "SystemDhcpServerReservedAddress-RemoteId"); ok {
			if err = d.Set("remote_id", vv); err != nil {
				return fmt.Errorf("Error reading remote_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_id: %v", err)
		}
	}

	if err = d.Set("remote_id_type", flattenSystemDhcpServerReservedAddressRemoteIdType2edl(o["remote-id-type"], d, "remote_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-id-type"], "SystemDhcpServerReservedAddress-RemoteIdType"); ok {
			if err = d.Set("remote_id_type", vv); err != nil {
				return fmt.Errorf("Error reading remote_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_id_type: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemDhcpServerReservedAddressType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemDhcpServerReservedAddress-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSystemDhcpServerReservedAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDhcpServerReservedAddressAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressCircuitId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressCircuitIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressRemoteId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressRemoteIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDhcpServerReservedAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSystemDhcpServerReservedAddressAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("circuit_id"); ok || d.HasChange("circuit_id") {
		t, err := expandSystemDhcpServerReservedAddressCircuitId2edl(d, v, "circuit_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["circuit-id"] = t
		}
	}

	if v, ok := d.GetOk("circuit_id_type"); ok || d.HasChange("circuit_id_type") {
		t, err := expandSystemDhcpServerReservedAddressCircuitIdType2edl(d, v, "circuit_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["circuit-id-type"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemDhcpServerReservedAddressDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemDhcpServerReservedAddressId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemDhcpServerReservedAddressIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandSystemDhcpServerReservedAddressMac2edl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("remote_id"); ok || d.HasChange("remote_id") {
		t, err := expandSystemDhcpServerReservedAddressRemoteId2edl(d, v, "remote_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_id_type"); ok || d.HasChange("remote_id_type") {
		t, err := expandSystemDhcpServerReservedAddressRemoteIdType2edl(d, v, "remote_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-id-type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemDhcpServerReservedAddressType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
