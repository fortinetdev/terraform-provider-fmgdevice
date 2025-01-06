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
			"db_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectDlpSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateDlpSettings(obj, mkey, paradict)
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

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteDlpSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting DlpSettings resource: %v", err)
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

func flattenDlpSettingsDbMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	if err = d.Set("db_mode", flattenDlpSettingsDbMode(o["db-mode"], d, "db_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["db-mode"], "DlpSettings-DbMode"); ok {
			if err = d.Set("db_mode", vv); err != nil {
				return fmt.Errorf("Error reading db_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading db_mode: %v", err)
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

func expandDlpSettingsDbMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsStorageDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectDlpSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
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

	if v, ok := d.GetOk("db_mode"); ok || d.HasChange("db_mode") {
		t, err := expandDlpSettingsDbMode(d, v, "db_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["db-mode"] = t
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
