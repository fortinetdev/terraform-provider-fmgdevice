// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure software switch interfaces by grouping physical and WiFi interfaces.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSwitchInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSwitchInterfaceCreate,
		Read:   resourceSystemSwitchInterfaceRead,
		Update: resourceSystemSwitchInterfaceUpdate,
		Delete: resourceSystemSwitchInterfaceDelete,

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
			"intra_switch_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"span": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"span_dest_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"span_direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"span_source_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSwitchInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemSwitchInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSwitchInterface resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSwitchInterface(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemSwitchInterface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSwitchInterfaceRead(d, m)
}

func resourceSystemSwitchInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemSwitchInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSwitchInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSwitchInterface(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSwitchInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSwitchInterfaceRead(d, m)
}

func resourceSystemSwitchInterfaceDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemSwitchInterface(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSwitchInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSwitchInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemSwitchInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSwitchInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSwitchInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSwitchInterface resource from API: %v", err)
	}
	return nil
}

func flattenSystemSwitchInterfaceIntraSwitchPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceMacTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSwitchInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceSpan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceSpanDestPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSwitchInterfaceSpanDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceSpanSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSwitchInterfaceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemSwitchInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("intra_switch_policy", flattenSystemSwitchInterfaceIntraSwitchPolicy(o["intra-switch-policy"], d, "intra_switch_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["intra-switch-policy"], "SystemSwitchInterface-IntraSwitchPolicy"); ok {
			if err = d.Set("intra_switch_policy", vv); err != nil {
				return fmt.Errorf("Error reading intra_switch_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intra_switch_policy: %v", err)
		}
	}

	if err = d.Set("mac_ttl", flattenSystemSwitchInterfaceMacTtl(o["mac-ttl"], d, "mac_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-ttl"], "SystemSwitchInterface-MacTtl"); ok {
			if err = d.Set("mac_ttl", vv); err != nil {
				return fmt.Errorf("Error reading mac_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_ttl: %v", err)
		}
	}

	if err = d.Set("member", flattenSystemSwitchInterfaceMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "SystemSwitchInterface-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSwitchInterfaceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSwitchInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("span", flattenSystemSwitchInterfaceSpan(o["span"], d, "span")); err != nil {
		if vv, ok := fortiAPIPatch(o["span"], "SystemSwitchInterface-Span"); ok {
			if err = d.Set("span", vv); err != nil {
				return fmt.Errorf("Error reading span: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading span: %v", err)
		}
	}

	if err = d.Set("span_dest_port", flattenSystemSwitchInterfaceSpanDestPort(o["span-dest-port"], d, "span_dest_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["span-dest-port"], "SystemSwitchInterface-SpanDestPort"); ok {
			if err = d.Set("span_dest_port", vv); err != nil {
				return fmt.Errorf("Error reading span_dest_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading span_dest_port: %v", err)
		}
	}

	if err = d.Set("span_direction", flattenSystemSwitchInterfaceSpanDirection(o["span-direction"], d, "span_direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["span-direction"], "SystemSwitchInterface-SpanDirection"); ok {
			if err = d.Set("span_direction", vv); err != nil {
				return fmt.Errorf("Error reading span_direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading span_direction: %v", err)
		}
	}

	if err = d.Set("span_source_port", flattenSystemSwitchInterfaceSpanSourcePort(o["span-source-port"], d, "span_source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["span-source-port"], "SystemSwitchInterface-SpanSourcePort"); ok {
			if err = d.Set("span_source_port", vv); err != nil {
				return fmt.Errorf("Error reading span_source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading span_source_port: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemSwitchInterfaceType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemSwitchInterface-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemSwitchInterfaceVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemSwitchInterface-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenSystemSwitchInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSwitchInterfaceIntraSwitchPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceMacTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSwitchInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceSpan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceSpanDestPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSwitchInterfaceSpanDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceSpanSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSwitchInterfaceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemSwitchInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("intra_switch_policy"); ok || d.HasChange("intra_switch_policy") {
		t, err := expandSystemSwitchInterfaceIntraSwitchPolicy(d, v, "intra_switch_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intra-switch-policy"] = t
		}
	}

	if v, ok := d.GetOk("mac_ttl"); ok || d.HasChange("mac_ttl") {
		t, err := expandSystemSwitchInterfaceMacTtl(d, v, "mac_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-ttl"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandSystemSwitchInterfaceMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSwitchInterfaceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("span"); ok || d.HasChange("span") {
		t, err := expandSystemSwitchInterfaceSpan(d, v, "span")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span"] = t
		}
	}

	if v, ok := d.GetOk("span_dest_port"); ok || d.HasChange("span_dest_port") {
		t, err := expandSystemSwitchInterfaceSpanDestPort(d, v, "span_dest_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span-dest-port"] = t
		}
	}

	if v, ok := d.GetOk("span_direction"); ok || d.HasChange("span_direction") {
		t, err := expandSystemSwitchInterfaceSpanDirection(d, v, "span_direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span-direction"] = t
		}
	}

	if v, ok := d.GetOk("span_source_port"); ok || d.HasChange("span_source_port") {
		t, err := expandSystemSwitchInterfaceSpanSourcePort(d, v, "span_source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span-source-port"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemSwitchInterfaceType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemSwitchInterfaceVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
