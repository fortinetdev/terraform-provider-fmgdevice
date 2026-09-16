// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Designate logical storage for DLP fingerprint database.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpSettingsUpdate,
		Read:   resourceDlpSettingsRead,
		Update: resourceDlpSettingsUpdate,
		Delete: resourceDlpSettingsDelete,

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
			"cache_mem_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"chunk_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"config_builder_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocr": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"storage_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDlpSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpSettings(d, false)
	if err != nil {
		return fmt.Errorf("Error updating DlpSettings resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DlpSettings")

	return resourceDlpSettingsRead(d, m)
}

func resourceDlpSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpSettings(d, true)

	if err != nil {
		return fmt.Errorf("Error updating DlpSettings resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing DlpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpSettings(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpSettings resource from API: %v", err)
	}
	return nil
}

func flattenDlpSettingsCacheMemPercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsChunkSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsConfigBuilderTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsDbMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsOcr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "confidence"
	if _, ok := i["confidence"]; ok {
		result["confidence"] = flattenDlpSettingsOcrConfidence(i["confidence"], d, pre_append)
	}

	pre_append = pre + ".0." + "filetype_ignore_list"
	if _, ok := i["filetype-ignore-list"]; ok {
		result["filetype_ignore_list"] = flattenDlpSettingsOcrFiletypeIgnoreList(i["filetype-ignore-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_file_size"
	if _, ok := i["max-file-size"]; ok {
		result["max_file_size"] = flattenDlpSettingsOcrMaxFileSize(i["max-file-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan"
	if _, ok := i["scan"]; ok {
		result["scan"] = flattenDlpSettingsOcrScan(i["scan"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenDlpSettingsOcrConfidence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsOcrFiletypeIgnoreList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpSettingsOcrMaxFileSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsOcrScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsStorageDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectDlpSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cache_mem_percent", flattenDlpSettingsCacheMemPercent(o["cache-mem-percent"], d, "cache_mem_percent")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-mem-percent"], "DlpSettings-CacheMemPercent"); ok {
			if err = d.Set("cache_mem_percent", vv); err != nil {
				return fmt.Errorf("Error reading cache_mem_percent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_mem_percent: %v", err)
		}
	}

	if err = d.Set("chunk_size", flattenDlpSettingsChunkSize(o["chunk-size"], d, "chunk_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["chunk-size"], "DlpSettings-ChunkSize"); ok {
			if err = d.Set("chunk_size", vv); err != nil {
				return fmt.Errorf("Error reading chunk_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading chunk_size: %v", err)
		}
	}

	if err = d.Set("config_builder_timeout", flattenDlpSettingsConfigBuilderTimeout(o["config-builder-timeout"], d, "config_builder_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["config-builder-timeout"], "DlpSettings-ConfigBuilderTimeout"); ok {
			if err = d.Set("config_builder_timeout", vv); err != nil {
				return fmt.Errorf("Error reading config_builder_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading config_builder_timeout: %v", err)
		}
	}

	if err = d.Set("db_mode", flattenDlpSettingsDbMode(o["db-mode"], d, "db_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["db-mode"], "DlpSettings-DbMode"); ok {
			if err = d.Set("db_mode", vv); err != nil {
				return fmt.Errorf("Error reading db_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading db_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ocr", flattenDlpSettingsOcr(o["ocr"], d, "ocr")); err != nil {
			if vv, ok := fortiAPIPatch(o["ocr"], "DlpSettings-Ocr"); ok {
				if err = d.Set("ocr", vv); err != nil {
					return fmt.Errorf("Error reading ocr: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ocr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ocr"); ok {
			if err = d.Set("ocr", flattenDlpSettingsOcr(o["ocr"], d, "ocr")); err != nil {
				if vv, ok := fortiAPIPatch(o["ocr"], "DlpSettings-Ocr"); ok {
					if err = d.Set("ocr", vv); err != nil {
						return fmt.Errorf("Error reading ocr: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ocr: %v", err)
				}
			}
		}
	}

	if err = d.Set("size", flattenDlpSettingsSize(o["size"], d, "size")); err != nil {
		if vv, ok := fortiAPIPatch(o["size"], "DlpSettings-Size"); ok {
			if err = d.Set("size", vv); err != nil {
				return fmt.Errorf("Error reading size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading size: %v", err)
		}
	}

	if err = d.Set("storage_device", flattenDlpSettingsStorageDevice(o["storage-device"], d, "storage_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["storage-device"], "DlpSettings-StorageDevice"); ok {
			if err = d.Set("storage_device", vv); err != nil {
				return fmt.Errorf("Error reading storage_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading storage_device: %v", err)
		}
	}

	return nil
}

func flattenDlpSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpSettingsCacheMemPercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsChunkSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsConfigBuilderTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsDbMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsOcr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "confidence"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["confidence"], _ = expandDlpSettingsOcrConfidence(d, i["confidence"], pre_append)
	}
	pre_append = pre + ".0." + "filetype_ignore_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["filetype-ignore-list"], _ = expandDlpSettingsOcrFiletypeIgnoreList(d, i["filetype_ignore_list"], pre_append)
	}
	pre_append = pre + ".0." + "max_file_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-file-size"], _ = expandDlpSettingsOcrMaxFileSize(d, i["max_file_size"], pre_append)
	}
	pre_append = pre + ".0." + "scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan"], _ = expandDlpSettingsOcrScan(d, i["scan"], pre_append)
	}

	return result, nil
}

func expandDlpSettingsOcrConfidence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsOcrFiletypeIgnoreList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpSettingsOcrMaxFileSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsOcrScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsStorageDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectDlpSettings(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cache_mem_percent"); ok || d.HasChange("cache_mem_percent") {
		t, err := expandDlpSettingsCacheMemPercent(d, v, "cache_mem_percent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-mem-percent"] = t
		}
	}

	if v, ok := d.GetOk("chunk_size"); ok || d.HasChange("chunk_size") {
		t, err := expandDlpSettingsChunkSize(d, v, "chunk_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chunk-size"] = t
		}
	}

	if v, ok := d.GetOk("config_builder_timeout"); ok || d.HasChange("config_builder_timeout") {
		t, err := expandDlpSettingsConfigBuilderTimeout(d, v, "config_builder_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["config-builder-timeout"] = t
		}
	}

	if v, ok := d.GetOk("db_mode"); ok || d.HasChange("db_mode") {
		t, err := expandDlpSettingsDbMode(d, v, "db_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["db-mode"] = t
		}
	}

	if v, ok := d.GetOk("ocr"); ok || d.HasChange("ocr") {
		t, err := expandDlpSettingsOcr(d, v, "ocr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocr"] = t
		}
	}

	if v, ok := d.GetOk("size"); ok || d.HasChange("size") {
		t, err := expandDlpSettingsSize(d, v, "size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["size"] = t
		}
	}

	if v, ok := d.GetOk("storage_device"); ok || d.HasChange("storage_device") {
		t, err := expandDlpSettingsStorageDevice(d, v, "storage_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["storage-device"] = t
		}
	}

	return &obj, nil
}
