// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DNS64.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDns64() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDns64Update,
		Read:   resourceSystemDns64Read,
		Update: resourceSystemDns64Update,
		Delete: resourceSystemDns64Delete,

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
			"always_synthesize_aaaa_record": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns64_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDns64Update(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemDns64(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns64 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDns64(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns64 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemDns64")

	return resourceSystemDns64Read(d, m)
}

func resourceSystemDns64Delete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteSystemDns64(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDns64 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDns64Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadSystemDns64(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns64 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDns64(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns64 resource from API: %v", err)
	}
	return nil
}

func flattenSystemDns64AlwaysSynthesizeAaaaRecord(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDns64Dns64Prefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDns64Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDns64(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("always_synthesize_aaaa_record", flattenSystemDns64AlwaysSynthesizeAaaaRecord(o["always-synthesize-aaaa-record"], d, "always_synthesize_aaaa_record")); err != nil {
		if vv, ok := fortiAPIPatch(o["always-synthesize-aaaa-record"], "SystemDns64-AlwaysSynthesizeAaaaRecord"); ok {
			if err = d.Set("always_synthesize_aaaa_record", vv); err != nil {
				return fmt.Errorf("Error reading always_synthesize_aaaa_record: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading always_synthesize_aaaa_record: %v", err)
		}
	}

	if err = d.Set("dns64_prefix", flattenSystemDns64Dns64Prefix(o["dns64-prefix"], d, "dns64_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns64-prefix"], "SystemDns64-Dns64Prefix"); ok {
			if err = d.Set("dns64_prefix", vv); err != nil {
				return fmt.Errorf("Error reading dns64_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns64_prefix: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemDns64Status(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemDns64-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemDns64FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDns64AlwaysSynthesizeAaaaRecord(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDns64Dns64Prefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDns64Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDns64(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("always_synthesize_aaaa_record"); ok || d.HasChange("always_synthesize_aaaa_record") {
		t, err := expandSystemDns64AlwaysSynthesizeAaaaRecord(d, v, "always_synthesize_aaaa_record")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["always-synthesize-aaaa-record"] = t
		}
	}

	if v, ok := d.GetOk("dns64_prefix"); ok || d.HasChange("dns64_prefix") {
		t, err := expandSystemDns64Dns64Prefix(d, v, "dns64_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns64-prefix"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemDns64Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
