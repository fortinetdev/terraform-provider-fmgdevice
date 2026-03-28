// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Replacement message table entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemReplacemsgGroupNacQuar() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReplacemsgGroupNacQuarCreate,
		Read:   resourceSystemReplacemsgGroupNacQuarRead,
		Update: resourceSystemReplacemsgGroupNacQuarUpdate,
		Delete: resourceSystemReplacemsgGroupNacQuarDelete,

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
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"buffer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"msg_type": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemReplacemsgGroupNacQuarCreate(d *schema.ResourceData, m interface{}) error {
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
	replacemsg_group := d.Get("replacemsg_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["replacemsg_group"] = replacemsg_group

	obj, err := getObjectSystemReplacemsgGroupNacQuar(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemReplacemsgGroupNacQuar resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("msg_type")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemReplacemsgGroupNacQuar(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemReplacemsgGroupNacQuar(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemReplacemsgGroupNacQuar resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemReplacemsgGroupNacQuar(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemReplacemsgGroupNacQuar resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "msg_type"))

	return resourceSystemReplacemsgGroupNacQuarRead(d, m)
}

func resourceSystemReplacemsgGroupNacQuarUpdate(d *schema.ResourceData, m interface{}) error {
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
	replacemsg_group := d.Get("replacemsg_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["replacemsg_group"] = replacemsg_group

	obj, err := getObjectSystemReplacemsgGroupNacQuar(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgGroupNacQuar resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemReplacemsgGroupNacQuar(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgGroupNacQuar resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "msg_type"))

	return resourceSystemReplacemsgGroupNacQuarRead(d, m)
}

func resourceSystemReplacemsgGroupNacQuarDelete(d *schema.ResourceData, m interface{}) error {
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
	replacemsg_group := d.Get("replacemsg_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["replacemsg_group"] = replacemsg_group

	wsParams["adom"] = adomv

	err = c.DeleteSystemReplacemsgGroupNacQuar(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReplacemsgGroupNacQuar resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgGroupNacQuarRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	replacemsg_group := d.Get("replacemsg_group").(string)
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
	if replacemsg_group == "" {
		replacemsg_group = importOptionChecking(m.(*FortiClient).Cfg, "replacemsg_group")
		if replacemsg_group == "" {
			return fmt.Errorf("Parameter replacemsg_group is missing")
		}
		if err = d.Set("replacemsg_group", replacemsg_group); err != nil {
			return fmt.Errorf("Error set params replacemsg_group: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["replacemsg_group"] = replacemsg_group

	o, err := c.ReadSystemReplacemsgGroupNacQuar(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemReplacemsgGroupNacQuar resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReplacemsgGroupNacQuar(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgGroupNacQuar resource from API: %v", err)
	}
	return nil
}

func flattenSystemReplacemsgGroupNacQuarBuffer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuarFormat2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuarHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuarId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuarMsgType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemReplacemsgGroupNacQuar(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("buffer", flattenSystemReplacemsgGroupNacQuarBuffer2edl(o["buffer"], d, "buffer")); err != nil {
		if vv, ok := fortiAPIPatch(o["buffer"], "SystemReplacemsgGroupNacQuar-Buffer"); ok {
			if err = d.Set("buffer", vv); err != nil {
				return fmt.Errorf("Error reading buffer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading buffer: %v", err)
		}
	}

	if err = d.Set("format", flattenSystemReplacemsgGroupNacQuarFormat2edl(o["format"], d, "format")); err != nil {
		if vv, ok := fortiAPIPatch(o["format"], "SystemReplacemsgGroupNacQuar-Format"); ok {
			if err = d.Set("format", vv); err != nil {
				return fmt.Errorf("Error reading format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("header", flattenSystemReplacemsgGroupNacQuarHeader2edl(o["header"], d, "header")); err != nil {
		if vv, ok := fortiAPIPatch(o["header"], "SystemReplacemsgGroupNacQuar-Header"); ok {
			if err = d.Set("header", vv); err != nil {
				return fmt.Errorf("Error reading header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemReplacemsgGroupNacQuarId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemReplacemsgGroupNacQuar-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("msg_type", flattenSystemReplacemsgGroupNacQuarMsgType2edl(o["msg-type"], d, "msg_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["msg-type"], "SystemReplacemsgGroupNacQuar-MsgType"); ok {
			if err = d.Set("msg_type", vv); err != nil {
				return fmt.Errorf("Error reading msg_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading msg_type: %v", err)
		}
	}

	return nil
}

func flattenSystemReplacemsgGroupNacQuarFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemReplacemsgGroupNacQuarBuffer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuarFormat2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuarHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuarId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuarMsgType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReplacemsgGroupNacQuar(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("buffer"); ok || d.HasChange("buffer") {
		t, err := expandSystemReplacemsgGroupNacQuarBuffer2edl(d, v, "buffer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["buffer"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok || d.HasChange("format") {
		t, err := expandSystemReplacemsgGroupNacQuarFormat2edl(d, v, "format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok || d.HasChange("header") {
		t, err := expandSystemReplacemsgGroupNacQuarHeader2edl(d, v, "header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemReplacemsgGroupNacQuarId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("msg_type"); ok || d.HasChange("msg_type") {
		t, err := expandSystemReplacemsgGroupNacQuarMsgType2edl(d, v, "msg_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msg-type"] = t
		}
	}

	return &obj, nil
}
