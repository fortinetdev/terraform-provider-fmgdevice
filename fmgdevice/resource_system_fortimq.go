// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiMQ settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFortimq() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFortimqUpdate,
		Read:   resourceSystemFortimqRead,
		Update: resourceSystemFortimqUpdate,
		Delete: resourceSystemFortimqDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"ocsp_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"publish_metadata": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemFortimqUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemFortimq(d, false)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortimq resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemFortimq(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortimq resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemFortimq")

	return resourceSystemFortimqRead(d, m)
}

func resourceSystemFortimqDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemFortimq(d, true)

	if err != nil {
		return fmt.Errorf("Error updating SystemFortimq resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemFortimq(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing SystemFortimq resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortimqRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemFortimq(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemFortimq resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFortimq(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortimq resource from API: %v", err)
	}
	return nil
}

func flattenSystemFortimqOcspCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortimqPublishMetadata(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortimqStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFortimq(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ocsp_check", flattenSystemFortimqOcspCheck(o["ocsp-check"], d, "ocsp_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["ocsp-check"], "SystemFortimq-OcspCheck"); ok {
			if err = d.Set("ocsp_check", vv); err != nil {
				return fmt.Errorf("Error reading ocsp_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ocsp_check: %v", err)
		}
	}

	if err = d.Set("publish_metadata", flattenSystemFortimqPublishMetadata(o["publish-metadata"], d, "publish_metadata")); err != nil {
		if vv, ok := fortiAPIPatch(o["publish-metadata"], "SystemFortimq-PublishMetadata"); ok {
			if err = d.Set("publish_metadata", vv); err != nil {
				return fmt.Errorf("Error reading publish_metadata: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading publish_metadata: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemFortimqStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemFortimq-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemFortimqFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFortimqOcspCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortimqPublishMetadata(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortimqStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFortimq(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ocsp_check"); ok || d.HasChange("ocsp_check") {
		t, err := expandSystemFortimqOcspCheck(d, v, "ocsp_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-check"] = t
		}
	}

	if v, ok := d.GetOk("publish_metadata"); ok || d.HasChange("publish_metadata") {
		t, err := expandSystemFortimqPublishMetadata(d, v, "publish_metadata")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["publish-metadata"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemFortimqStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
