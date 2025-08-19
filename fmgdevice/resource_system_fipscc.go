// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FIPS-CC mode.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFipsCc() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFipsCcUpdate,
		Read:   resourceSystemFipsCcRead,
		Update: resourceSystemFipsCcUpdate,
		Delete: resourceSystemFipsCcDelete,

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
			"entropy_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_generation_self_test": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"self_test_period": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSystemFipsCcUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemFipsCc(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFipsCc resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFipsCc(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemFipsCc resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemFipsCc")

	return resourceSystemFipsCcRead(d, m)
}

func resourceSystemFipsCcDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemFipsCc(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFipsCc resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFipsCcRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemFipsCc(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemFipsCc resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFipsCc(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFipsCc resource from API: %v", err)
	}
	return nil
}

func flattenSystemFipsCcEntropyToken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFipsCcKeyGenerationSelfTest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFipsCcSelfTestPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFipsCcStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFipsCc(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("entropy_token", flattenSystemFipsCcEntropyToken(o["entropy-token"], d, "entropy_token")); err != nil {
		if vv, ok := fortiAPIPatch(o["entropy-token"], "SystemFipsCc-EntropyToken"); ok {
			if err = d.Set("entropy_token", vv); err != nil {
				return fmt.Errorf("Error reading entropy_token: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading entropy_token: %v", err)
		}
	}

	if err = d.Set("key_generation_self_test", flattenSystemFipsCcKeyGenerationSelfTest(o["key-generation-self-test"], d, "key_generation_self_test")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-generation-self-test"], "SystemFipsCc-KeyGenerationSelfTest"); ok {
			if err = d.Set("key_generation_self_test", vv); err != nil {
				return fmt.Errorf("Error reading key_generation_self_test: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_generation_self_test: %v", err)
		}
	}

	if err = d.Set("self_test_period", flattenSystemFipsCcSelfTestPeriod(o["self-test-period"], d, "self_test_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["self-test-period"], "SystemFipsCc-SelfTestPeriod"); ok {
			if err = d.Set("self_test_period", vv); err != nil {
				return fmt.Errorf("Error reading self_test_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading self_test_period: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemFipsCcStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemFipsCc-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemFipsCcFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFipsCcEntropyToken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFipsCcKeyGenerationSelfTest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFipsCcSelfTestPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFipsCcStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFipsCc(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("entropy_token"); ok || d.HasChange("entropy_token") {
		t, err := expandSystemFipsCcEntropyToken(d, v, "entropy_token")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entropy-token"] = t
		}
	}

	if v, ok := d.GetOk("key_generation_self_test"); ok || d.HasChange("key_generation_self_test") {
		t, err := expandSystemFipsCcKeyGenerationSelfTest(d, v, "key_generation_self_test")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-generation-self-test"] = t
		}
	}

	if v, ok := d.GetOk("self_test_period"); ok || d.HasChange("self_test_period") {
		t, err := expandSystemFipsCcSelfTestPeriod(d, v, "self_test_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["self-test-period"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemFipsCcStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
