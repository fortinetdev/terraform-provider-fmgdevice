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

func resourceSystemDhcpTemplateReservedAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcpTemplateReservedAddressCreate,
		Read:   resourceSystemDhcpTemplateReservedAddressRead,
		Update: resourceSystemDhcpTemplateReservedAddressUpdate,
		Delete: resourceSystemDhcpTemplateReservedAddressDelete,

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
			"template": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"circuit_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"circuit_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"ip_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDhcpTemplateReservedAddressCreate(d *schema.ResourceData, m interface{}) error {
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
	template := d.Get("template").(string)
	paradict["device"] = device_name
	paradict["template"] = template

	obj, err := getObjectSystemDhcpTemplateReservedAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcpTemplateReservedAddress resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemDhcpTemplateReservedAddress(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemDhcpTemplateReservedAddress(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemDhcpTemplateReservedAddress resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemDhcpTemplateReservedAddress(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemDhcpTemplateReservedAddress resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpTemplateReservedAddressRead(d, m)
}

func resourceSystemDhcpTemplateReservedAddressUpdate(d *schema.ResourceData, m interface{}) error {
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
	template := d.Get("template").(string)
	paradict["device"] = device_name
	paradict["template"] = template

	obj, err := getObjectSystemDhcpTemplateReservedAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpTemplateReservedAddress resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemDhcpTemplateReservedAddress(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpTemplateReservedAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpTemplateReservedAddressRead(d, m)
}

func resourceSystemDhcpTemplateReservedAddressDelete(d *schema.ResourceData, m interface{}) error {
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
	template := d.Get("template").(string)
	paradict["device"] = device_name
	paradict["template"] = template

	wsParams["adom"] = adomv

	err = c.DeleteSystemDhcpTemplateReservedAddress(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcpTemplateReservedAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcpTemplateReservedAddressRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	template := d.Get("template").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if template == "" {
		template = importOptionChecking(m.(*FortiClient).Cfg, "template")
		if template == "" {
			return fmt.Errorf("Parameter template is missing")
		}
		if err = d.Set("template", template); err != nil {
			return fmt.Errorf("Error set params template: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["template"] = template

	o, err := c.ReadSystemDhcpTemplateReservedAddress(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemDhcpTemplateReservedAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcpTemplateReservedAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcpTemplateReservedAddress resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcpTemplateReservedAddressAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressCircuitId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressCircuitIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressIpIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressRemoteId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressRemoteIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDhcpTemplateReservedAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenSystemDhcpTemplateReservedAddressAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SystemDhcpTemplateReservedAddress-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("circuit_id", flattenSystemDhcpTemplateReservedAddressCircuitId2edl(o["circuit-id"], d, "circuit_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["circuit-id"], "SystemDhcpTemplateReservedAddress-CircuitId"); ok {
			if err = d.Set("circuit_id", vv); err != nil {
				return fmt.Errorf("Error reading circuit_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading circuit_id: %v", err)
		}
	}

	if err = d.Set("circuit_id_type", flattenSystemDhcpTemplateReservedAddressCircuitIdType2edl(o["circuit-id-type"], d, "circuit_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["circuit-id-type"], "SystemDhcpTemplateReservedAddress-CircuitIdType"); ok {
			if err = d.Set("circuit_id_type", vv); err != nil {
				return fmt.Errorf("Error reading circuit_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading circuit_id_type: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemDhcpTemplateReservedAddressDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemDhcpTemplateReservedAddress-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemDhcpTemplateReservedAddressId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemDhcpTemplateReservedAddress-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_index", flattenSystemDhcpTemplateReservedAddressIpIndex2edl(o["ip-index"], d, "ip_index")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-index"], "SystemDhcpTemplateReservedAddress-IpIndex"); ok {
			if err = d.Set("ip_index", vv); err != nil {
				return fmt.Errorf("Error reading ip_index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_index: %v", err)
		}
	}

	if err = d.Set("mac", flattenSystemDhcpTemplateReservedAddressMac2edl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "SystemDhcpTemplateReservedAddress-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("remote_id", flattenSystemDhcpTemplateReservedAddressRemoteId2edl(o["remote-id"], d, "remote_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-id"], "SystemDhcpTemplateReservedAddress-RemoteId"); ok {
			if err = d.Set("remote_id", vv); err != nil {
				return fmt.Errorf("Error reading remote_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_id: %v", err)
		}
	}

	if err = d.Set("remote_id_type", flattenSystemDhcpTemplateReservedAddressRemoteIdType2edl(o["remote-id-type"], d, "remote_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-id-type"], "SystemDhcpTemplateReservedAddress-RemoteIdType"); ok {
			if err = d.Set("remote_id_type", vv); err != nil {
				return fmt.Errorf("Error reading remote_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_id_type: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemDhcpTemplateReservedAddressType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemDhcpTemplateReservedAddress-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSystemDhcpTemplateReservedAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDhcpTemplateReservedAddressAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressCircuitId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressCircuitIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressIpIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressRemoteId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressRemoteIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDhcpTemplateReservedAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSystemDhcpTemplateReservedAddressAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("circuit_id"); ok || d.HasChange("circuit_id") {
		t, err := expandSystemDhcpTemplateReservedAddressCircuitId2edl(d, v, "circuit_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["circuit-id"] = t
		}
	}

	if v, ok := d.GetOk("circuit_id_type"); ok || d.HasChange("circuit_id_type") {
		t, err := expandSystemDhcpTemplateReservedAddressCircuitIdType2edl(d, v, "circuit_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["circuit-id-type"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemDhcpTemplateReservedAddressDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemDhcpTemplateReservedAddressId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_index"); ok || d.HasChange("ip_index") {
		t, err := expandSystemDhcpTemplateReservedAddressIpIndex2edl(d, v, "ip_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-index"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandSystemDhcpTemplateReservedAddressMac2edl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("remote_id"); ok || d.HasChange("remote_id") {
		t, err := expandSystemDhcpTemplateReservedAddressRemoteId2edl(d, v, "remote_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_id_type"); ok || d.HasChange("remote_id_type") {
		t, err := expandSystemDhcpTemplateReservedAddressRemoteIdType2edl(d, v, "remote_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-id-type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemDhcpTemplateReservedAddressType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
