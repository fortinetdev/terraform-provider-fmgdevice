// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure virtual domain.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVdom() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomCreate,
		Read:   resourceSystemVdomRead,
		Update: resourceSystemVdomUpdate,
		Delete: resourceSystemVdomDelete,

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
			"flag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"short_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcluster_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemVdomCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemVdom(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdom resource while getting object: %v", err)
	}

	_, err = c.CreateSystemVdom(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdom resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemVdomRead(d, m)
}

func resourceSystemVdomUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemVdom(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdom resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVdom(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemVdomRead(d, m)
}

func resourceSystemVdomDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemVdom(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVdom(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdom resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomShortName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomVclusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemVdom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("flag", flattenSystemVdomFlag(o["flag"], d, "flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["flag"], "SystemVdom-Flag"); ok {
			if err = d.Set("flag", vv); err != nil {
				return fmt.Errorf("Error reading flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flag: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemVdomName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemVdom-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("short_name", flattenSystemVdomShortName(o["short-name"], d, "short_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["short-name"], "SystemVdom-ShortName"); ok {
			if err = d.Set("short_name", vv); err != nil {
				return fmt.Errorf("Error reading short_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading short_name: %v", err)
		}
	}

	if err = d.Set("vcluster_id", flattenSystemVdomVclusterId(o["vcluster-id"], d, "vcluster_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vcluster-id"], "SystemVdom-VclusterId"); ok {
			if err = d.Set("vcluster_id", vv); err != nil {
				return fmt.Errorf("Error reading vcluster_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vcluster_id: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVdomFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomShortName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomVclusterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("flag"); ok || d.HasChange("flag") {
		t, err := expandSystemVdomFlag(d, v, "flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flag"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemVdomName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("short_name"); ok || d.HasChange("short_name") {
		t, err := expandSystemVdomShortName(d, v, "short_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["short-name"] = t
		}
	}

	if v, ok := d.GetOk("vcluster_id"); ok || d.HasChange("vcluster_id") {
		t, err := expandSystemVdomVclusterId(d, v, "vcluster_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcluster-id"] = t
		}
	}

	return &obj, nil
}
