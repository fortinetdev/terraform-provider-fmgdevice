// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpFpDocSource() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpFpDocSourceCreate,
		Read:   resourceDlpFpDocSourceRead,
		Update: resourceDlpFpDocSourceUpdate,
		Delete: resourceDlpFpDocSourceDelete,

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
			"date": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"file_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keep_modified": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remove_deleted": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_on_creation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_subdirectories": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sensitivity": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tod_hour": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tod_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weekday": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDlpFpDocSourceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectDlpFpDocSource(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpFpDocSource resource while getting object: %v", err)
	}

	_, err = c.CreateDlpFpDocSource(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating DlpFpDocSource resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceDlpFpDocSourceRead(d, m)
}

func resourceDlpFpDocSourceUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectDlpFpDocSource(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpFpDocSource resource while getting object: %v", err)
	}

	_, err = c.UpdateDlpFpDocSource(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpFpDocSource resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDlpFpDocSourceRead(d, m)
}

func resourceDlpFpDocSourceDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteDlpFpDocSource(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpFpDocSource resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpFpDocSourceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpFpDocSource(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading DlpFpDocSource resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpFpDocSource(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpFpDocSource resource from API: %v", err)
	}
	return nil
}

func flattenDlpFpDocSourceDate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceFilePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceFilePattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceKeepModified(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourcePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceRemoveDeleted(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceScanOnCreation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceScanSubdirectories(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpFpDocSourceServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceTodHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceTodMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceWeekday(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpFpDocSource(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("date", flattenDlpFpDocSourceDate(o["date"], d, "date")); err != nil {
		if vv, ok := fortiAPIPatch(o["date"], "DlpFpDocSource-Date"); ok {
			if err = d.Set("date", vv); err != nil {
				return fmt.Errorf("Error reading date: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading date: %v", err)
		}
	}

	if err = d.Set("file_path", flattenDlpFpDocSourceFilePath(o["file-path"], d, "file_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-path"], "DlpFpDocSource-FilePath"); ok {
			if err = d.Set("file_path", vv); err != nil {
				return fmt.Errorf("Error reading file_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_path: %v", err)
		}
	}

	if err = d.Set("file_pattern", flattenDlpFpDocSourceFilePattern(o["file-pattern"], d, "file_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-pattern"], "DlpFpDocSource-FilePattern"); ok {
			if err = d.Set("file_pattern", vv); err != nil {
				return fmt.Errorf("Error reading file_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_pattern: %v", err)
		}
	}

	if err = d.Set("keep_modified", flattenDlpFpDocSourceKeepModified(o["keep-modified"], d, "keep_modified")); err != nil {
		if vv, ok := fortiAPIPatch(o["keep-modified"], "DlpFpDocSource-KeepModified"); ok {
			if err = d.Set("keep_modified", vv); err != nil {
				return fmt.Errorf("Error reading keep_modified: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keep_modified: %v", err)
		}
	}

	if err = d.Set("name", flattenDlpFpDocSourceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DlpFpDocSource-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("period", flattenDlpFpDocSourcePeriod(o["period"], d, "period")); err != nil {
		if vv, ok := fortiAPIPatch(o["period"], "DlpFpDocSource-Period"); ok {
			if err = d.Set("period", vv); err != nil {
				return fmt.Errorf("Error reading period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading period: %v", err)
		}
	}

	if err = d.Set("remove_deleted", flattenDlpFpDocSourceRemoveDeleted(o["remove-deleted"], d, "remove_deleted")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-deleted"], "DlpFpDocSource-RemoveDeleted"); ok {
			if err = d.Set("remove_deleted", vv); err != nil {
				return fmt.Errorf("Error reading remove_deleted: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_deleted: %v", err)
		}
	}

	if err = d.Set("scan_on_creation", flattenDlpFpDocSourceScanOnCreation(o["scan-on-creation"], d, "scan_on_creation")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-on-creation"], "DlpFpDocSource-ScanOnCreation"); ok {
			if err = d.Set("scan_on_creation", vv); err != nil {
				return fmt.Errorf("Error reading scan_on_creation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_on_creation: %v", err)
		}
	}

	if err = d.Set("scan_subdirectories", flattenDlpFpDocSourceScanSubdirectories(o["scan-subdirectories"], d, "scan_subdirectories")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-subdirectories"], "DlpFpDocSource-ScanSubdirectories"); ok {
			if err = d.Set("scan_subdirectories", vv); err != nil {
				return fmt.Errorf("Error reading scan_subdirectories: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_subdirectories: %v", err)
		}
	}

	if err = d.Set("sensitivity", flattenDlpFpDocSourceSensitivity(o["sensitivity"], d, "sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["sensitivity"], "DlpFpDocSource-Sensitivity"); ok {
			if err = d.Set("sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sensitivity: %v", err)
		}
	}

	if err = d.Set("server", flattenDlpFpDocSourceServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "DlpFpDocSource-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_type", flattenDlpFpDocSourceServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "DlpFpDocSource-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("tod_hour", flattenDlpFpDocSourceTodHour(o["tod-hour"], d, "tod_hour")); err != nil {
		if vv, ok := fortiAPIPatch(o["tod-hour"], "DlpFpDocSource-TodHour"); ok {
			if err = d.Set("tod_hour", vv); err != nil {
				return fmt.Errorf("Error reading tod_hour: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tod_hour: %v", err)
		}
	}

	if err = d.Set("tod_min", flattenDlpFpDocSourceTodMin(o["tod-min"], d, "tod_min")); err != nil {
		if vv, ok := fortiAPIPatch(o["tod-min"], "DlpFpDocSource-TodMin"); ok {
			if err = d.Set("tod_min", vv); err != nil {
				return fmt.Errorf("Error reading tod_min: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tod_min: %v", err)
		}
	}

	if err = d.Set("username", flattenDlpFpDocSourceUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "DlpFpDocSource-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("vdom", flattenDlpFpDocSourceVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "DlpFpDocSource-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("weekday", flattenDlpFpDocSourceWeekday(o["weekday"], d, "weekday")); err != nil {
		if vv, ok := fortiAPIPatch(o["weekday"], "DlpFpDocSource-Weekday"); ok {
			if err = d.Set("weekday", vv); err != nil {
				return fmt.Errorf("Error reading weekday: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weekday: %v", err)
		}
	}

	return nil
}

func flattenDlpFpDocSourceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpFpDocSourceDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceFilePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceFilePattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceKeepModified(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourcePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpFpDocSourcePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceRemoveDeleted(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceScanOnCreation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceScanSubdirectories(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpFpDocSourceServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceTodHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceTodMin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceWeekday(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpFpDocSource(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("date"); ok || d.HasChange("date") {
		t, err := expandDlpFpDocSourceDate(d, v, "date")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["date"] = t
		}
	}

	if v, ok := d.GetOk("file_path"); ok || d.HasChange("file_path") {
		t, err := expandDlpFpDocSourceFilePath(d, v, "file_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-path"] = t
		}
	}

	if v, ok := d.GetOk("file_pattern"); ok || d.HasChange("file_pattern") {
		t, err := expandDlpFpDocSourceFilePattern(d, v, "file_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-pattern"] = t
		}
	}

	if v, ok := d.GetOk("keep_modified"); ok || d.HasChange("keep_modified") {
		t, err := expandDlpFpDocSourceKeepModified(d, v, "keep_modified")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keep-modified"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDlpFpDocSourceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandDlpFpDocSourcePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("period"); ok || d.HasChange("period") {
		t, err := expandDlpFpDocSourcePeriod(d, v, "period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["period"] = t
		}
	}

	if v, ok := d.GetOk("remove_deleted"); ok || d.HasChange("remove_deleted") {
		t, err := expandDlpFpDocSourceRemoveDeleted(d, v, "remove_deleted")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-deleted"] = t
		}
	}

	if v, ok := d.GetOk("scan_on_creation"); ok || d.HasChange("scan_on_creation") {
		t, err := expandDlpFpDocSourceScanOnCreation(d, v, "scan_on_creation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-on-creation"] = t
		}
	}

	if v, ok := d.GetOk("scan_subdirectories"); ok || d.HasChange("scan_subdirectories") {
		t, err := expandDlpFpDocSourceScanSubdirectories(d, v, "scan_subdirectories")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-subdirectories"] = t
		}
	}

	if v, ok := d.GetOk("sensitivity"); ok || d.HasChange("sensitivity") {
		t, err := expandDlpFpDocSourceSensitivity(d, v, "sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandDlpFpDocSourceServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandDlpFpDocSourceServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("tod_hour"); ok || d.HasChange("tod_hour") {
		t, err := expandDlpFpDocSourceTodHour(d, v, "tod_hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tod-hour"] = t
		}
	}

	if v, ok := d.GetOk("tod_min"); ok || d.HasChange("tod_min") {
		t, err := expandDlpFpDocSourceTodMin(d, v, "tod_min")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tod-min"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandDlpFpDocSourceUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandDlpFpDocSourceVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("weekday"); ok || d.HasChange("weekday") {
		t, err := expandDlpFpDocSourceWeekday(d, v, "weekday")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weekday"] = t
		}
	}

	return &obj, nil
}
