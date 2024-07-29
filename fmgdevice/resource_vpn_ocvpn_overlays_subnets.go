// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Internal subnets to register with OCVPN service.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnOcvpnOverlaysSubnets() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnOcvpnOverlaysSubnetsCreate,
		Read:   resourceVpnOcvpnOverlaysSubnetsRead,
		Update: resourceVpnOcvpnOverlaysSubnetsUpdate,
		Delete: resourceVpnOcvpnOverlaysSubnetsDelete,

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
			"overlays": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceVpnOcvpnOverlaysSubnetsCreate(d *schema.ResourceData, m interface{}) error {
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
	overlays := d.Get("overlays").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["overlays"] = overlays

	obj, err := getObjectVpnOcvpnOverlaysSubnets(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnOcvpnOverlaysSubnets resource while getting object: %v", err)
	}

	_, err = c.CreateVpnOcvpnOverlaysSubnets(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating VpnOcvpnOverlaysSubnets resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceVpnOcvpnOverlaysSubnetsRead(d, m)
}

func resourceVpnOcvpnOverlaysSubnetsUpdate(d *schema.ResourceData, m interface{}) error {
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
	overlays := d.Get("overlays").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["overlays"] = overlays

	obj, err := getObjectVpnOcvpnOverlaysSubnets(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnOcvpnOverlaysSubnets resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnOcvpnOverlaysSubnets(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating VpnOcvpnOverlaysSubnets resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceVpnOcvpnOverlaysSubnetsRead(d, m)
}

func resourceVpnOcvpnOverlaysSubnetsDelete(d *schema.ResourceData, m interface{}) error {
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
	overlays := d.Get("overlays").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["overlays"] = overlays

	err = c.DeleteVpnOcvpnOverlaysSubnets(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting VpnOcvpnOverlaysSubnets resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnOcvpnOverlaysSubnetsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	overlays := d.Get("overlays").(string)
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
	if overlays == "" {
		overlays = importOptionChecking(m.(*FortiClient).Cfg, "overlays")
		if overlays == "" {
			return fmt.Errorf("Parameter overlays is missing")
		}
		if err = d.Set("overlays", overlays); err != nil {
			return fmt.Errorf("Error set params overlays: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["overlays"] = overlays

	o, err := c.ReadVpnOcvpnOverlaysSubnets(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnOcvpnOverlaysSubnets resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnOcvpnOverlaysSubnets(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnOcvpnOverlaysSubnets resource from API: %v", err)
	}
	return nil
}

func flattenVpnOcvpnOverlaysSubnetsId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysSubnetsInterface3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnOcvpnOverlaysSubnetsSubnet3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnOcvpnOverlaysSubnetsType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnOcvpnOverlaysSubnets(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenVpnOcvpnOverlaysSubnetsId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "VpnOcvpnOverlaysSubnets-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenVpnOcvpnOverlaysSubnetsInterface3rdl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "VpnOcvpnOverlaysSubnets-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("subnet", flattenVpnOcvpnOverlaysSubnetsSubnet3rdl(o["subnet"], d, "subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet"], "VpnOcvpnOverlaysSubnets-Subnet"); ok {
			if err = d.Set("subnet", vv); err != nil {
				return fmt.Errorf("Error reading subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet: %v", err)
		}
	}

	if err = d.Set("type", flattenVpnOcvpnOverlaysSubnetsType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "VpnOcvpnOverlaysSubnets-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenVpnOcvpnOverlaysSubnetsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnOcvpnOverlaysSubnetsId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysSubnetsInterface3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnOcvpnOverlaysSubnetsSubnet3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandVpnOcvpnOverlaysSubnetsType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnOcvpnOverlaysSubnets(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandVpnOcvpnOverlaysSubnetsId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandVpnOcvpnOverlaysSubnetsInterface3rdl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("subnet"); ok || d.HasChange("subnet") {
		t, err := expandVpnOcvpnOverlaysSubnetsSubnet3rdl(d, v, "subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandVpnOcvpnOverlaysSubnetsType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
