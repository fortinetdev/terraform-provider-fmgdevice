// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure custom languages.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCustomLanguage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCustomLanguageCreate,
		Read:   resourceSystemCustomLanguageRead,
		Update: resourceSystemCustomLanguageUpdate,
		Delete: resourceSystemCustomLanguageDelete,

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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemCustomLanguageCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemCustomLanguage(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCustomLanguage resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemCustomLanguage(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemCustomLanguage(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemCustomLanguage resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemCustomLanguage(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemCustomLanguage resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemCustomLanguageRead(d, m)
}

func resourceSystemCustomLanguageUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemCustomLanguage(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCustomLanguage resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemCustomLanguage(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemCustomLanguage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemCustomLanguageRead(d, m)
}

func resourceSystemCustomLanguageDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteSystemCustomLanguage(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCustomLanguage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCustomLanguageRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCustomLanguage(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemCustomLanguage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCustomLanguage(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCustomLanguage resource from API: %v", err)
	}
	return nil
}

func flattenSystemCustomLanguageComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCustomLanguageFilename(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCustomLanguageName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemCustomLanguage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comments", flattenSystemCustomLanguageComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "SystemCustomLanguage-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("filename", flattenSystemCustomLanguageFilename(o["filename"], d, "filename")); err != nil {
		if vv, ok := fortiAPIPatch(o["filename"], "SystemCustomLanguage-Filename"); ok {
			if err = d.Set("filename", vv); err != nil {
				return fmt.Errorf("Error reading filename: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filename: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemCustomLanguageName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemCustomLanguage-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemCustomLanguageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCustomLanguageComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCustomLanguageFilename(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCustomLanguageName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCustomLanguage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandSystemCustomLanguageComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("filename"); ok || d.HasChange("filename") {
		t, err := expandSystemCustomLanguageFilename(d, v, "filename")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filename"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemCustomLanguageName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
