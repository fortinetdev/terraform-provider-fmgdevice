// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Diffie-Hellman group 5 (MODP-1536).

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIkeDhGroup5() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIkeDhGroup5Update,
		Read:   resourceSystemIkeDhGroup5Read,
		Update: resourceSystemIkeDhGroup5Update,
		Delete: resourceSystemIkeDhGroup5Delete,

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
				Optional: true,
			},
			"keypair_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keypair_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemIkeDhGroup5Update(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemIkeDhGroup5(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemIkeDhGroup5 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIkeDhGroup5(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemIkeDhGroup5 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemIkeDhGroup5")

	return resourceSystemIkeDhGroup5Read(d, m)
}

func resourceSystemIkeDhGroup5Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemIkeDhGroup5(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIkeDhGroup5 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIkeDhGroup5Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemIkeDhGroup5(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemIkeDhGroup5 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIkeDhGroup5(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemIkeDhGroup5 resource from API: %v", err)
	}
	return nil
}

func flattenSystemIkeDhGroup5Id2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup5KeypairCache2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup5KeypairCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup5Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemIkeDhGroup5(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemIkeDhGroup5Id2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemIkeDhGroup5-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("keypair_cache", flattenSystemIkeDhGroup5KeypairCache2edl(o["keypair-cache"], d, "keypair_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["keypair-cache"], "SystemIkeDhGroup5-KeypairCache"); ok {
			if err = d.Set("keypair_cache", vv); err != nil {
				return fmt.Errorf("Error reading keypair_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keypair_cache: %v", err)
		}
	}

	if err = d.Set("keypair_count", flattenSystemIkeDhGroup5KeypairCount2edl(o["keypair-count"], d, "keypair_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["keypair-count"], "SystemIkeDhGroup5-KeypairCount"); ok {
			if err = d.Set("keypair_count", vv); err != nil {
				return fmt.Errorf("Error reading keypair_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keypair_count: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemIkeDhGroup5Mode2edl(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemIkeDhGroup5-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	return nil
}

func flattenSystemIkeDhGroup5FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemIkeDhGroup5Id2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup5KeypairCache2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup5KeypairCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup5Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIkeDhGroup5(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemIkeDhGroup5Id2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("keypair_cache"); ok || d.HasChange("keypair_cache") {
		t, err := expandSystemIkeDhGroup5KeypairCache2edl(d, v, "keypair_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keypair-cache"] = t
		}
	}

	if v, ok := d.GetOk("keypair_count"); ok || d.HasChange("keypair_count") {
		t, err := expandSystemIkeDhGroup5KeypairCount2edl(d, v, "keypair_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keypair-count"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemIkeDhGroup5Mode2edl(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	return &obj, nil
}
