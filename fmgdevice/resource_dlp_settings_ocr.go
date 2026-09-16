// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure settings for optical character recognition (OCR) conversion.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpSettingsOcr() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpSettingsOcrUpdate,
		Read:   resourceDlpSettingsOcrRead,
		Update: resourceDlpSettingsOcrUpdate,
		Delete: resourceDlpSettingsOcrDelete,

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
			"confidence": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"filetype_ignore_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"max_file_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDlpSettingsOcrUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpSettingsOcr(d, false)
	if err != nil {
		return fmt.Errorf("Error updating DlpSettingsOcr resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpSettingsOcr(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpSettingsOcr resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DlpSettingsOcr")

	return resourceDlpSettingsOcrRead(d, m)
}

func resourceDlpSettingsOcrDelete(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpSettingsOcr(d, true)

	if err != nil {
		return fmt.Errorf("Error updating DlpSettingsOcr resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpSettingsOcr(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing DlpSettingsOcr resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpSettingsOcrRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpSettingsOcr(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpSettingsOcr resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpSettingsOcr(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpSettingsOcr resource from API: %v", err)
	}
	return nil
}

func flattenDlpSettingsOcrConfidence2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsOcrFiletypeIgnoreList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpSettingsOcrMaxFileSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsOcrScan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpSettingsOcr(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("confidence", flattenDlpSettingsOcrConfidence2edl(o["confidence"], d, "confidence")); err != nil {
		if vv, ok := fortiAPIPatch(o["confidence"], "DlpSettingsOcr-Confidence"); ok {
			if err = d.Set("confidence", vv); err != nil {
				return fmt.Errorf("Error reading confidence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading confidence: %v", err)
		}
	}

	if err = d.Set("filetype_ignore_list", flattenDlpSettingsOcrFiletypeIgnoreList2edl(o["filetype-ignore-list"], d, "filetype_ignore_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["filetype-ignore-list"], "DlpSettingsOcr-FiletypeIgnoreList"); ok {
			if err = d.Set("filetype_ignore_list", vv); err != nil {
				return fmt.Errorf("Error reading filetype_ignore_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filetype_ignore_list: %v", err)
		}
	}

	if err = d.Set("max_file_size", flattenDlpSettingsOcrMaxFileSize2edl(o["max-file-size"], d, "max_file_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-file-size"], "DlpSettingsOcr-MaxFileSize"); ok {
			if err = d.Set("max_file_size", vv); err != nil {
				return fmt.Errorf("Error reading max_file_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_file_size: %v", err)
		}
	}

	if err = d.Set("scan", flattenDlpSettingsOcrScan2edl(o["scan"], d, "scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan"], "DlpSettingsOcr-Scan"); ok {
			if err = d.Set("scan", vv); err != nil {
				return fmt.Errorf("Error reading scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan: %v", err)
		}
	}

	return nil
}

func flattenDlpSettingsOcrFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpSettingsOcrConfidence2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsOcrFiletypeIgnoreList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpSettingsOcrMaxFileSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsOcrScan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpSettingsOcr(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("confidence"); ok || d.HasChange("confidence") {
		t, err := expandDlpSettingsOcrConfidence2edl(d, v, "confidence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["confidence"] = t
		}
	}

	if v, ok := d.GetOk("filetype_ignore_list"); ok || d.HasChange("filetype_ignore_list") {
		t, err := expandDlpSettingsOcrFiletypeIgnoreList2edl(d, v, "filetype_ignore_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filetype-ignore-list"] = t
		}
	}

	if v, ok := d.GetOk("max_file_size"); ok || d.HasChange("max_file_size") {
		t, err := expandDlpSettingsOcrMaxFileSize2edl(d, v, "max_file_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-file-size"] = t
		}
	}

	if v, ok := d.GetOk("scan"); ok || d.HasChange("scan") {
		t, err := expandDlpSettingsOcrScan2edl(d, v, "scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan"] = t
		}
	}

	return &obj, nil
}
