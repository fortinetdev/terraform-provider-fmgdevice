// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> File filter rules.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFileFilterProfileRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceFileFilterProfileRulesCreate,
		Read:   resourceFileFilterProfileRulesRead,
		Update: resourceFileFilterProfileRulesUpdate,
		Delete: resourceFileFilterProfileRulesDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password_protected": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFileFilterProfileRulesCreate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectFileFilterProfileRules(d)
	if err != nil {
		return fmt.Errorf("Error creating FileFilterProfileRules resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFileFilterProfileRules(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFileFilterProfileRules(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FileFilterProfileRules resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFileFilterProfileRules(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FileFilterProfileRules resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFileFilterProfileRulesRead(d, m)
}

func resourceFileFilterProfileRulesUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectFileFilterProfileRules(d)
	if err != nil {
		return fmt.Errorf("Error updating FileFilterProfileRules resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFileFilterProfileRules(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FileFilterProfileRules resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFileFilterProfileRulesRead(d, m)
}

func resourceFileFilterProfileRulesDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteFileFilterProfileRules(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FileFilterProfileRules resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFileFilterProfileRulesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadFileFilterProfileRules(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FileFilterProfileRules resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFileFilterProfileRules(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FileFilterProfileRules resource from API: %v", err)
	}
	return nil
}

func flattenFileFilterProfileRulesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFileFilterProfileRulesComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFileFilterProfileRulesDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFileFilterProfileRulesFileType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFileFilterProfileRulesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFileFilterProfileRulesPasswordProtected2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFileFilterProfileRulesProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFileFilterProfileRules(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenFileFilterProfileRulesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "FileFilterProfileRules-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("comment", flattenFileFilterProfileRulesComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FileFilterProfileRules-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("direction", flattenFileFilterProfileRulesDirection2edl(o["direction"], d, "direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["direction"], "FileFilterProfileRules-Direction"); ok {
			if err = d.Set("direction", vv); err != nil {
				return fmt.Errorf("Error reading direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("file_type", flattenFileFilterProfileRulesFileType2edl(o["file-type"], d, "file_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-type"], "FileFilterProfileRules-FileType"); ok {
			if err = d.Set("file_type", vv); err != nil {
				return fmt.Errorf("Error reading file_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_type: %v", err)
		}
	}

	if err = d.Set("name", flattenFileFilterProfileRulesName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FileFilterProfileRules-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("password_protected", flattenFileFilterProfileRulesPasswordProtected2edl(o["password-protected"], d, "password_protected")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-protected"], "FileFilterProfileRules-PasswordProtected"); ok {
			if err = d.Set("password_protected", vv); err != nil {
				return fmt.Errorf("Error reading password_protected: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_protected: %v", err)
		}
	}

	if err = d.Set("protocol", flattenFileFilterProfileRulesProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "FileFilterProfileRules-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	return nil
}

func flattenFileFilterProfileRulesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFileFilterProfileRulesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFileFilterProfileRulesComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFileFilterProfileRulesDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFileFilterProfileRulesFileType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFileFilterProfileRulesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFileFilterProfileRulesPasswordProtected2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFileFilterProfileRulesProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFileFilterProfileRules(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandFileFilterProfileRulesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFileFilterProfileRulesComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok || d.HasChange("direction") {
		t, err := expandFileFilterProfileRulesDirection2edl(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("file_type"); ok || d.HasChange("file_type") {
		t, err := expandFileFilterProfileRulesFileType2edl(d, v, "file_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFileFilterProfileRulesName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password_protected"); ok || d.HasChange("password_protected") {
		t, err := expandFileFilterProfileRulesPasswordProtected2edl(d, v, "password_protected")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-protected"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandFileFilterProfileRulesProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	return &obj, nil
}
