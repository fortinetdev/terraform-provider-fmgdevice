// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Settings for individual Security Rating controls.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSecurityRatingControls() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSecurityRatingControlsCreate,
		Read:   resourceSystemSecurityRatingControlsRead,
		Update: resourceSystemSecurityRatingControlsUpdate,
		Delete: resourceSystemSecurityRatingControlsDelete,

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
			"display_insight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemSecurityRatingControlsCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSecurityRatingControls(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSecurityRatingControls resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSecurityRatingControls(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemSecurityRatingControls resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSecurityRatingControlsRead(d, m)
}

func resourceSystemSecurityRatingControlsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSecurityRatingControls(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSecurityRatingControls resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSecurityRatingControls(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSecurityRatingControls resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSecurityRatingControlsRead(d, m)
}

func resourceSystemSecurityRatingControlsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSecurityRatingControls(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSecurityRatingControls resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSecurityRatingControlsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSecurityRatingControls(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSecurityRatingControls resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSecurityRatingControls(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSecurityRatingControls resource from API: %v", err)
	}
	return nil
}

func flattenSystemSecurityRatingControlsDisplayInsight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSecurityRatingControlsDisplayReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSecurityRatingControlsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSecurityRatingControls(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("display_insight", flattenSystemSecurityRatingControlsDisplayInsight(o["display-insight"], d, "display_insight")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-insight"], "SystemSecurityRatingControls-DisplayInsight"); ok {
			if err = d.Set("display_insight", vv); err != nil {
				return fmt.Errorf("Error reading display_insight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_insight: %v", err)
		}
	}

	if err = d.Set("display_report", flattenSystemSecurityRatingControlsDisplayReport(o["display-report"], d, "display_report")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-report"], "SystemSecurityRatingControls-DisplayReport"); ok {
			if err = d.Set("display_report", vv); err != nil {
				return fmt.Errorf("Error reading display_report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_report: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSecurityRatingControlsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSecurityRatingControls-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemSecurityRatingControlsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSecurityRatingControlsDisplayInsight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSecurityRatingControlsDisplayReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSecurityRatingControlsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSecurityRatingControls(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("display_insight"); ok || d.HasChange("display_insight") {
		t, err := expandSystemSecurityRatingControlsDisplayInsight(d, v, "display_insight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-insight"] = t
		}
	}

	if v, ok := d.GetOk("display_report"); ok || d.HasChange("display_report") {
		t, err := expandSystemSecurityRatingControlsDisplayReport(d, v, "display_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-report"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSecurityRatingControlsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
