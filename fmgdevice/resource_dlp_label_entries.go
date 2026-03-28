// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> DLP label entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpLabelEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpLabelEntriesCreate,
		Read:   resourceDlpLabelEntriesRead,
		Update: resourceDlpLabelEntriesUpdate,
		Delete: resourceDlpLabelEntriesDelete,

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
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fortidata_label_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"guid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"mpip_label_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDlpLabelEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	label := d.Get("label").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["label"] = label

	obj, err := getObjectDlpLabelEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpLabelEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDlpLabelEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDlpLabelEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DlpLabelEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDlpLabelEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DlpLabelEntries resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDlpLabelEntriesRead(d, m)
}

func resourceDlpLabelEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	label := d.Get("label").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["label"] = label

	obj, err := getObjectDlpLabelEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpLabelEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpLabelEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpLabelEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDlpLabelEntriesRead(d, m)
}

func resourceDlpLabelEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	label := d.Get("label").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["label"] = label

	wsParams["adom"] = adomv

	err = c.DeleteDlpLabelEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpLabelEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpLabelEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	label := d.Get("label").(string)
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
	if label == "" {
		label = importOptionChecking(m.(*FortiClient).Cfg, "label")
		if label == "" {
			return fmt.Errorf("Parameter label is missing")
		}
		if err = d.Set("label", label); err != nil {
			return fmt.Errorf("Error set params label: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["label"] = label

	o, err := c.ReadDlpLabelEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpLabelEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpLabelEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpLabelEntries resource from API: %v", err)
	}
	return nil
}

func flattenDlpLabelEntriesFortidataLabelName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpLabelEntriesGuid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpLabelEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpLabelEntriesMpipLabelName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpLabelEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fortidata_label_name", flattenDlpLabelEntriesFortidataLabelName2edl(o["fortidata-label-name"], d, "fortidata_label_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortidata-label-name"], "DlpLabelEntries-FortidataLabelName"); ok {
			if err = d.Set("fortidata_label_name", vv); err != nil {
				return fmt.Errorf("Error reading fortidata_label_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortidata_label_name: %v", err)
		}
	}

	if err = d.Set("guid", flattenDlpLabelEntriesGuid2edl(o["guid"], d, "guid")); err != nil {
		if vv, ok := fortiAPIPatch(o["guid"], "DlpLabelEntries-Guid"); ok {
			if err = d.Set("guid", vv); err != nil {
				return fmt.Errorf("Error reading guid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guid: %v", err)
		}
	}

	if err = d.Set("fosid", flattenDlpLabelEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "DlpLabelEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("mpip_label_name", flattenDlpLabelEntriesMpipLabelName2edl(o["mpip-label-name"], d, "mpip_label_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpip-label-name"], "DlpLabelEntries-MpipLabelName"); ok {
			if err = d.Set("mpip_label_name", vv); err != nil {
				return fmt.Errorf("Error reading mpip_label_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpip_label_name: %v", err)
		}
	}

	return nil
}

func flattenDlpLabelEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpLabelEntriesFortidataLabelName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpLabelEntriesGuid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpLabelEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpLabelEntriesMpipLabelName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpLabelEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fortidata_label_name"); ok || d.HasChange("fortidata_label_name") {
		t, err := expandDlpLabelEntriesFortidataLabelName2edl(d, v, "fortidata_label_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortidata-label-name"] = t
		}
	}

	if v, ok := d.GetOk("guid"); ok || d.HasChange("guid") {
		t, err := expandDlpLabelEntriesGuid2edl(d, v, "guid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guid"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandDlpLabelEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("mpip_label_name"); ok || d.HasChange("mpip_label_name") {
		t, err := expandDlpLabelEntriesMpipLabelName2edl(d, v, "mpip_label_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpip-label-name"] = t
		}
	}

	return &obj, nil
}
