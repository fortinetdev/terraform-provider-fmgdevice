// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure file patterns used by DLP blocking.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpFilepatternEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpFilepatternEntriesCreate,
		Read:   resourceDlpFilepatternEntriesRead,
		Update: resourceDlpFilepatternEntriesUpdate,
		Delete: resourceDlpFilepatternEntriesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"filepattern": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceDlpFilepatternEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	filepattern := d.Get("filepattern").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["filepattern"] = filepattern

	obj, err := getObjectDlpFilepatternEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpFilepatternEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("pattern")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDlpFilepatternEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDlpFilepatternEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DlpFilepatternEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDlpFilepatternEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DlpFilepatternEntries resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "pattern"))

	return resourceDlpFilepatternEntriesRead(d, m)
}

func resourceDlpFilepatternEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	filepattern := d.Get("filepattern").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["filepattern"] = filepattern

	obj, err := getObjectDlpFilepatternEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpFilepatternEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpFilepatternEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpFilepatternEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "pattern"))

	return resourceDlpFilepatternEntriesRead(d, m)
}

func resourceDlpFilepatternEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	filepattern := d.Get("filepattern").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["filepattern"] = filepattern

	wsParams["adom"] = adomv

	err = c.DeleteDlpFilepatternEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpFilepatternEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpFilepatternEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	filepattern := d.Get("filepattern").(string)
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
	if filepattern == "" {
		filepattern = importOptionChecking(m.(*FortiClient).Cfg, "filepattern")
		if filepattern == "" {
			return fmt.Errorf("Parameter filepattern is missing")
		}
		if err = d.Set("filepattern", filepattern); err != nil {
			return fmt.Errorf("Error set params filepattern: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["filepattern"] = filepattern

	o, err := c.ReadDlpFilepatternEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpFilepatternEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpFilepatternEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpFilepatternEntries resource from API: %v", err)
	}
	return nil
}

func flattenDlpFilepatternEntriesFileType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFilepatternEntriesFilterType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFilepatternEntriesPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpFilepatternEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("file_type", flattenDlpFilepatternEntriesFileType2edl(o["file-type"], d, "file_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-type"], "DlpFilepatternEntries-FileType"); ok {
			if err = d.Set("file_type", vv); err != nil {
				return fmt.Errorf("Error reading file_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_type: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenDlpFilepatternEntriesFilterType2edl(o["filter-type"], d, "filter_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-type"], "DlpFilepatternEntries-FilterType"); ok {
			if err = d.Set("filter_type", vv); err != nil {
				return fmt.Errorf("Error reading filter_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	if err = d.Set("pattern", flattenDlpFilepatternEntriesPattern2edl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "DlpFilepatternEntries-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	return nil
}

func flattenDlpFilepatternEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpFilepatternEntriesFileType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFilepatternEntriesFilterType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFilepatternEntriesPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpFilepatternEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("file_type"); ok || d.HasChange("file_type") {
		t, err := expandDlpFilepatternEntriesFileType2edl(d, v, "file_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-type"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok || d.HasChange("filter_type") {
		t, err := expandDlpFilepatternEntriesFilterType2edl(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandDlpFilepatternEntriesPattern2edl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	return &obj, nil
}
