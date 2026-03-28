// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure system interface on FortiSwitch.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchSystemInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchSystemInterfaceCreate,
		Read:   resourceSwitchControllerManagedSwitchSystemInterfaceRead,
		Update: resourceSwitchControllerManagedSwitchSystemInterfaceUpdate,
		Delete: resourceSwitchControllerManagedSwitchSystemInterfaceDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchSystemInterfaceCreate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectSwitchControllerManagedSwitchSystemInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchSystemInterface resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSwitchControllerManagedSwitchSystemInterface(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSwitchControllerManagedSwitchSystemInterface(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SwitchControllerManagedSwitchSystemInterface resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSwitchControllerManagedSwitchSystemInterface(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SwitchControllerManagedSwitchSystemInterface resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerManagedSwitchSystemInterfaceRead(d, m)
}

func resourceSwitchControllerManagedSwitchSystemInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectSwitchControllerManagedSwitchSystemInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSystemInterface resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSwitchControllerManagedSwitchSystemInterface(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSystemInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerManagedSwitchSystemInterfaceRead(d, m)
}

func resourceSwitchControllerManagedSwitchSystemInterfaceDelete(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	wsParams["adom"] = adomv

	err = c.DeleteSwitchControllerManagedSwitchSystemInterface(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchSystemInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchSystemInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerManagedSwitchSystemInterface(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSystemInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchSystemInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSystemInterface resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchSystemInterfaceAllowaccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchSystemInterfaceInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchSystemInterfaceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchSystemInterfaceMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemInterfaceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemInterfaceStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemInterfaceSwitchId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchSystemInterfaceType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemInterfaceVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchSystemInterfaceVrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerManagedSwitchSystemInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("allowaccess", flattenSwitchControllerManagedSwitchSystemInterfaceAllowaccess2edl(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "SwitchControllerManagedSwitchSystemInterface-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("interface", flattenSwitchControllerManagedSwitchSystemInterfaceInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SwitchControllerManagedSwitchSystemInterface-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip", flattenSwitchControllerManagedSwitchSystemInterfaceIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SwitchControllerManagedSwitchSystemInterface-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mode", flattenSwitchControllerManagedSwitchSystemInterfaceMode2edl(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SwitchControllerManagedSwitchSystemInterface-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerManagedSwitchSystemInterfaceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerManagedSwitchSystemInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchControllerManagedSwitchSystemInterfaceStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SwitchControllerManagedSwitchSystemInterface-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("switch_id", flattenSwitchControllerManagedSwitchSystemInterfaceSwitchId2edl(o["switch-id"], d, "switch_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-id"], "SwitchControllerManagedSwitchSystemInterface-SwitchId"); ok {
			if err = d.Set("switch_id", vv); err != nil {
				return fmt.Errorf("Error reading switch_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_id: %v", err)
		}
	}

	if err = d.Set("type", flattenSwitchControllerManagedSwitchSystemInterfaceType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SwitchControllerManagedSwitchSystemInterface-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSwitchControllerManagedSwitchSystemInterfaceVlan2edl(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "SwitchControllerManagedSwitchSystemInterface-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	if err = d.Set("vrf", flattenSwitchControllerManagedSwitchSystemInterfaceVrf2edl(o["vrf"], d, "vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf"], "SwitchControllerManagedSwitchSystemInterface-Vrf"); ok {
			if err = d.Set("vrf", vv); err != nil {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchSystemInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchSystemInterfaceAllowaccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchSystemInterfaceInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchSystemInterfaceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerManagedSwitchSystemInterfaceMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemInterfaceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemInterfaceStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemInterfaceSwitchId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchSystemInterfaceType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemInterfaceVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchSystemInterfaceVrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerManagedSwitchSystemInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandSwitchControllerManagedSwitchSystemInterfaceAllowaccess2edl(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSwitchControllerManagedSwitchSystemInterfaceInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSwitchControllerManagedSwitchSystemInterfaceIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSwitchControllerManagedSwitchSystemInterfaceMode2edl(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerManagedSwitchSystemInterfaceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSwitchControllerManagedSwitchSystemInterfaceStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("switch_id"); ok || d.HasChange("switch_id") {
		t, err := expandSwitchControllerManagedSwitchSystemInterfaceSwitchId2edl(d, v, "switch_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSwitchControllerManagedSwitchSystemInterfaceType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandSwitchControllerManagedSwitchSystemInterfaceVlan2edl(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		t, err := expandSwitchControllerManagedSwitchSystemInterfaceVrf2edl(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	return &obj, nil
}
