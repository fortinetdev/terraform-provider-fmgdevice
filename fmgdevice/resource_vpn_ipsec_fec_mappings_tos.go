// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FEC redundancy mapping table for specific type of service (TOS).

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnIpsecFecMappingsTos() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecFecMappingsTosCreate,
		Read:   resourceVpnIpsecFecMappingsTosRead,
		Update: resourceVpnIpsecFecMappingsTosUpdate,
		Delete: resourceVpnIpsecFecMappingsTosDelete,

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
			"fec": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mappings": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"base": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"redundant": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"seqno": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVpnIpsecFecMappingsTosCreate(d *schema.ResourceData, m interface{}) error {
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
	fec := d.Get("fec").(string)
	mappings := d.Get("mappings").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fec"] = fec
	paradict["mappings"] = mappings

	obj, err := getObjectVpnIpsecFecMappingsTos(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecFecMappingsTos resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("seqno")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVpnIpsecFecMappingsTos(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVpnIpsecFecMappingsTos(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VpnIpsecFecMappingsTos resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateVpnIpsecFecMappingsTos(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VpnIpsecFecMappingsTos resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "seqno")))

	return resourceVpnIpsecFecMappingsTosRead(d, m)
}

func resourceVpnIpsecFecMappingsTosUpdate(d *schema.ResourceData, m interface{}) error {
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
	fec := d.Get("fec").(string)
	mappings := d.Get("mappings").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fec"] = fec
	paradict["mappings"] = mappings

	obj, err := getObjectVpnIpsecFecMappingsTos(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecFecMappingsTos resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVpnIpsecFecMappingsTos(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecFecMappingsTos resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "seqno")))

	return resourceVpnIpsecFecMappingsTosRead(d, m)
}

func resourceVpnIpsecFecMappingsTosDelete(d *schema.ResourceData, m interface{}) error {
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
	fec := d.Get("fec").(string)
	mappings := d.Get("mappings").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fec"] = fec
	paradict["mappings"] = mappings

	wsParams["adom"] = adomv

	err = c.DeleteVpnIpsecFecMappingsTos(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecFecMappingsTos resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecFecMappingsTosRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	fec := d.Get("fec").(string)
	mappings := d.Get("mappings").(string)
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
	if fec == "" {
		fec = importOptionChecking(m.(*FortiClient).Cfg, "fec")
		if fec == "" {
			return fmt.Errorf("Parameter fec is missing")
		}
		if err = d.Set("fec", fec); err != nil {
			return fmt.Errorf("Error set params fec: %v", err)
		}
	}
	if mappings == "" {
		mappings = importOptionChecking(m.(*FortiClient).Cfg, "mappings")
		if mappings == "" {
			return fmt.Errorf("Parameter mappings is missing")
		}
		if err = d.Set("mappings", mappings); err != nil {
			return fmt.Errorf("Error set params mappings: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fec"] = fec
	paradict["mappings"] = mappings

	o, err := c.ReadVpnIpsecFecMappingsTos(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnIpsecFecMappingsTos resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecFecMappingsTos(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecFecMappingsTos resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecFecMappingsTosBase3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsTosRedundant3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsTosSeqno3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsTosTos3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsTosTosMask3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnIpsecFecMappingsTos(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("base", flattenVpnIpsecFecMappingsTosBase3rdl(o["base"], d, "base")); err != nil {
		if vv, ok := fortiAPIPatch(o["base"], "VpnIpsecFecMappingsTos-Base"); ok {
			if err = d.Set("base", vv); err != nil {
				return fmt.Errorf("Error reading base: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading base: %v", err)
		}
	}

	if err = d.Set("redundant", flattenVpnIpsecFecMappingsTosRedundant3rdl(o["redundant"], d, "redundant")); err != nil {
		if vv, ok := fortiAPIPatch(o["redundant"], "VpnIpsecFecMappingsTos-Redundant"); ok {
			if err = d.Set("redundant", vv); err != nil {
				return fmt.Errorf("Error reading redundant: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redundant: %v", err)
		}
	}

	if err = d.Set("seqno", flattenVpnIpsecFecMappingsTosSeqno3rdl(o["seqno"], d, "seqno")); err != nil {
		if vv, ok := fortiAPIPatch(o["seqno"], "VpnIpsecFecMappingsTos-Seqno"); ok {
			if err = d.Set("seqno", vv); err != nil {
				return fmt.Errorf("Error reading seqno: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seqno: %v", err)
		}
	}

	if err = d.Set("tos", flattenVpnIpsecFecMappingsTosTos3rdl(o["tos"], d, "tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos"], "VpnIpsecFecMappingsTos-Tos"); ok {
			if err = d.Set("tos", vv); err != nil {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenVpnIpsecFecMappingsTosTosMask3rdl(o["tos-mask"], d, "tos_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-mask"], "VpnIpsecFecMappingsTos-TosMask"); ok {
			if err = d.Set("tos_mask", vv); err != nil {
				return fmt.Errorf("Error reading tos_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecFecMappingsTosFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnIpsecFecMappingsTosBase3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsTosRedundant3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsTosSeqno3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsTosTos3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsTosTosMask3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecFecMappingsTos(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("base"); ok || d.HasChange("base") {
		t, err := expandVpnIpsecFecMappingsTosBase3rdl(d, v, "base")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base"] = t
		}
	}

	if v, ok := d.GetOk("redundant"); ok || d.HasChange("redundant") {
		t, err := expandVpnIpsecFecMappingsTosRedundant3rdl(d, v, "redundant")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant"] = t
		}
	}

	if v, ok := d.GetOk("seqno"); ok || d.HasChange("seqno") {
		t, err := expandVpnIpsecFecMappingsTosSeqno3rdl(d, v, "seqno")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seqno"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandVpnIpsecFecMappingsTosTos3rdl(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok || d.HasChange("tos_mask") {
		t, err := expandVpnIpsecFecMappingsTosTosMask3rdl(d, v, "tos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	return &obj, nil
}
