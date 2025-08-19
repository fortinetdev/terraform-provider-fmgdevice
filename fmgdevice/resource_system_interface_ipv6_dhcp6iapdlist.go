// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DHCPv6 IA-PD list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemInterfaceIpv6Dhcp6IapdList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceIpv6Dhcp6IapdListCreate,
		Read:   resourceSystemInterfaceIpv6Dhcp6IapdListRead,
		Update: resourceSystemInterfaceIpv6Dhcp6IapdListUpdate,
		Delete: resourceSystemInterfaceIpv6Dhcp6IapdListDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"iaid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"prefix_hint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix_hint_plt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"prefix_hint_vlt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemInterfaceIpv6Dhcp6IapdListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemInterfaceIpv6Dhcp6IapdList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceIpv6Dhcp6IapdList resource while getting object: %v", err)
	}

	_, err = c.CreateSystemInterfaceIpv6Dhcp6IapdList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceIpv6Dhcp6IapdList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "iaid")))

	return resourceSystemInterfaceIpv6Dhcp6IapdListRead(d, m)
}

func resourceSystemInterfaceIpv6Dhcp6IapdListUpdate(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemInterfaceIpv6Dhcp6IapdList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6Dhcp6IapdList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceIpv6Dhcp6IapdList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6Dhcp6IapdList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "iaid")))

	return resourceSystemInterfaceIpv6Dhcp6IapdListRead(d, m)
}

func resourceSystemInterfaceIpv6Dhcp6IapdListDelete(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemInterfaceIpv6Dhcp6IapdList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceIpv6Dhcp6IapdList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceIpv6Dhcp6IapdListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	var_interface := d.Get("interface").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if var_interface == "" {
		var_interface = importOptionChecking(m.(*FortiClient).Cfg, "interface")
		if var_interface == "" {
			return fmt.Errorf("Parameter interface is missing")
		}
		if err = d.Set("interface", var_interface); err != nil {
			return fmt.Errorf("Error set params interface: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	o, err := c.ReadSystemInterfaceIpv6Dhcp6IapdList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6Dhcp6IapdList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceIpv6Dhcp6IapdList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6Dhcp6IapdList resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceIpv6Dhcp6IapdListIaid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6IapdListPrefixHint3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6IapdListPrefixHintPlt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6IapdListPrefixHintVlt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterfaceIpv6Dhcp6IapdList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("iaid", flattenSystemInterfaceIpv6Dhcp6IapdListIaid3rdl(o["iaid"], d, "iaid")); err != nil {
		if vv, ok := fortiAPIPatch(o["iaid"], "SystemInterfaceIpv6Dhcp6IapdList-Iaid"); ok {
			if err = d.Set("iaid", vv); err != nil {
				return fmt.Errorf("Error reading iaid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iaid: %v", err)
		}
	}

	if err = d.Set("prefix_hint", flattenSystemInterfaceIpv6Dhcp6IapdListPrefixHint3rdl(o["prefix-hint"], d, "prefix_hint")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-hint"], "SystemInterfaceIpv6Dhcp6IapdList-PrefixHint"); ok {
			if err = d.Set("prefix_hint", vv); err != nil {
				return fmt.Errorf("Error reading prefix_hint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_hint: %v", err)
		}
	}

	if err = d.Set("prefix_hint_plt", flattenSystemInterfaceIpv6Dhcp6IapdListPrefixHintPlt3rdl(o["prefix-hint-plt"], d, "prefix_hint_plt")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-hint-plt"], "SystemInterfaceIpv6Dhcp6IapdList-PrefixHintPlt"); ok {
			if err = d.Set("prefix_hint_plt", vv); err != nil {
				return fmt.Errorf("Error reading prefix_hint_plt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_hint_plt: %v", err)
		}
	}

	if err = d.Set("prefix_hint_vlt", flattenSystemInterfaceIpv6Dhcp6IapdListPrefixHintVlt3rdl(o["prefix-hint-vlt"], d, "prefix_hint_vlt")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-hint-vlt"], "SystemInterfaceIpv6Dhcp6IapdList-PrefixHintVlt"); ok {
			if err = d.Set("prefix_hint_vlt", vv); err != nil {
				return fmt.Errorf("Error reading prefix_hint_vlt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_hint_vlt: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfaceIpv6Dhcp6IapdListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceIpv6Dhcp6IapdListIaid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6IapdListPrefixHint3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6IapdListPrefixHintPlt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6IapdListPrefixHintVlt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceIpv6Dhcp6IapdList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("iaid"); ok || d.HasChange("iaid") {
		t, err := expandSystemInterfaceIpv6Dhcp6IapdListIaid3rdl(d, v, "iaid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iaid"] = t
		}
	}

	if v, ok := d.GetOk("prefix_hint"); ok || d.HasChange("prefix_hint") {
		t, err := expandSystemInterfaceIpv6Dhcp6IapdListPrefixHint3rdl(d, v, "prefix_hint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-hint"] = t
		}
	}

	if v, ok := d.GetOk("prefix_hint_plt"); ok || d.HasChange("prefix_hint_plt") {
		t, err := expandSystemInterfaceIpv6Dhcp6IapdListPrefixHintPlt3rdl(d, v, "prefix_hint_plt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-hint-plt"] = t
		}
	}

	if v, ok := d.GetOk("prefix_hint_vlt"); ok || d.HasChange("prefix_hint_vlt") {
		t, err := expandSystemInterfaceIpv6Dhcp6IapdListPrefixHintVlt3rdl(d, v, "prefix_hint_vlt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-hint-vlt"] = t
		}
	}

	return &obj, nil
}
