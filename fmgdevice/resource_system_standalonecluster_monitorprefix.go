// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure a list of routing prefixes to monitor.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemStandaloneClusterMonitorPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemStandaloneClusterMonitorPrefixCreate,
		Read:   resourceSystemStandaloneClusterMonitorPrefixRead,
		Update: resourceSystemStandaloneClusterMonitorPrefixUpdate,
		Delete: resourceSystemStandaloneClusterMonitorPrefixDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemStandaloneClusterMonitorPrefixCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemStandaloneClusterMonitorPrefix(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemStandaloneClusterMonitorPrefix resource while getting object: %v", err)
	}

	_, err = c.CreateSystemStandaloneClusterMonitorPrefix(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemStandaloneClusterMonitorPrefix resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemStandaloneClusterMonitorPrefixRead(d, m)
}

func resourceSystemStandaloneClusterMonitorPrefixUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemStandaloneClusterMonitorPrefix(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneClusterMonitorPrefix resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemStandaloneClusterMonitorPrefix(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneClusterMonitorPrefix resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemStandaloneClusterMonitorPrefixRead(d, m)
}

func resourceSystemStandaloneClusterMonitorPrefixDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemStandaloneClusterMonitorPrefix(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemStandaloneClusterMonitorPrefix resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStandaloneClusterMonitorPrefixRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemStandaloneClusterMonitorPrefix(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneClusterMonitorPrefix resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemStandaloneClusterMonitorPrefix(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneClusterMonitorPrefix resource from API: %v", err)
	}
	return nil
}

func flattenSystemStandaloneClusterMonitorPrefixId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterMonitorPrefixPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterMonitorPrefixVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterMonitorPrefixVrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemStandaloneClusterMonitorPrefix(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemStandaloneClusterMonitorPrefixId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemStandaloneClusterMonitorPrefix-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("prefix", flattenSystemStandaloneClusterMonitorPrefixPrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "SystemStandaloneClusterMonitorPrefix-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemStandaloneClusterMonitorPrefixVdom2edl(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemStandaloneClusterMonitorPrefix-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("vrf", flattenSystemStandaloneClusterMonitorPrefixVrf2edl(o["vrf"], d, "vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf"], "SystemStandaloneClusterMonitorPrefix-Vrf"); ok {
			if err = d.Set("vrf", vv); err != nil {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	return nil
}

func flattenSystemStandaloneClusterMonitorPrefixFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemStandaloneClusterMonitorPrefixId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterMonitorPrefixPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemStandaloneClusterMonitorPrefixVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterMonitorPrefixVrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemStandaloneClusterMonitorPrefix(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemStandaloneClusterMonitorPrefixId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandSystemStandaloneClusterMonitorPrefixPrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemStandaloneClusterMonitorPrefixVdom2edl(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		t, err := expandSystemStandaloneClusterMonitorPrefixVrf2edl(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	return &obj, nil
}
