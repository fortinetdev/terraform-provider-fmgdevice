// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NSX-T setting.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNsxtSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtSettingUpdate,
		Read:   resourceNsxtSettingRead,
		Update: resourceNsxtSettingUpdate,
		Delete: resourceNsxtSettingDelete,

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
			"liveness": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceNsxtSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectNsxtSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating NsxtSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateNsxtSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating NsxtSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("NsxtSetting")

	return resourceNsxtSettingRead(d, m)
}

func resourceNsxtSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteNsxtSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting NsxtSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceNsxtSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadNsxtSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading NsxtSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectNsxtSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading NsxtSetting resource from API: %v", err)
	}
	return nil
}

func flattenNsxtSettingLiveness(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenNsxtSettingService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectNsxtSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("liveness", flattenNsxtSettingLiveness(o["liveness"], d, "liveness")); err != nil {
		if vv, ok := fortiAPIPatch(o["liveness"], "NsxtSetting-Liveness"); ok {
			if err = d.Set("liveness", vv); err != nil {
				return fmt.Errorf("Error reading liveness: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading liveness: %v", err)
		}
	}

	if err = d.Set("service", flattenNsxtSettingService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "NsxtSetting-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	return nil
}

func flattenNsxtSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandNsxtSettingLiveness(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandNsxtSettingService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectNsxtSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("liveness"); ok || d.HasChange("liveness") {
		t, err := expandNsxtSettingLiveness(d, v, "liveness")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["liveness"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandNsxtSettingService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	return &obj, nil
}
