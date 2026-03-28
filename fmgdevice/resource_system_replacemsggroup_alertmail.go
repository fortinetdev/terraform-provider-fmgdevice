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

func resourceSystemReplacemsgGroupAlertmail() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReplacemsgGroupAlertmailCreate,
		Read:   resourceSystemReplacemsgGroupAlertmailRead,
		Update: resourceSystemReplacemsgGroupAlertmailUpdate,
		Delete: resourceSystemReplacemsgGroupAlertmailDelete,

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

func resourceSystemReplacemsgGroupAlertmailCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemReplacemsgGroupAlertmail(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemReplacemsgGroupAlertmail resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("msg_type")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemReplacemsgGroupAlertmail(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemReplacemsgGroupAlertmail(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemReplacemsgGroupAlertmail resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemReplacemsgGroupAlertmail(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemReplacemsgGroupAlertmail resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "msg_type"))

	return resourceSystemReplacemsgGroupAlertmailRead(d, m)
}

func resourceSystemReplacemsgGroupAlertmailUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemReplacemsgGroupAlertmail(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgGroupAlertmail resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemReplacemsgGroupAlertmail(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgGroupAlertmail resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "msg_type"))

	return resourceSystemReplacemsgGroupAlertmailRead(d, m)
}

func resourceSystemReplacemsgGroupAlertmailDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemReplacemsgGroupAlertmail(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReplacemsgGroupAlertmail resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgGroupAlertmailRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemReplacemsgGroupAlertmail(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemReplacemsgGroupAlertmail resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReplacemsgGroupAlertmail(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgGroupAlertmail resource from API: %v", err)
	}
	return nil
}

func flattenSystemReplacemsgGroupAlertmailBuffer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmailFormat2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmailHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmailId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmailMsgType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemReplacemsgGroupAlertmail(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("buffer", flattenSystemReplacemsgGroupAlertmailBuffer2edl(o["buffer"], d, "buffer")); err != nil {
		if vv, ok := fortiAPIPatch(o["buffer"], "SystemReplacemsgGroupAlertmail-Buffer"); ok {
			if err = d.Set("buffer", vv); err != nil {
				return fmt.Errorf("Error reading buffer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading buffer: %v", err)
		}
	}

	if err = d.Set("format", flattenSystemReplacemsgGroupAlertmailFormat2edl(o["format"], d, "format")); err != nil {
		if vv, ok := fortiAPIPatch(o["format"], "SystemReplacemsgGroupAlertmail-Format"); ok {
			if err = d.Set("format", vv); err != nil {
				return fmt.Errorf("Error reading format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("header", flattenSystemReplacemsgGroupAlertmailHeader2edl(o["header"], d, "header")); err != nil {
		if vv, ok := fortiAPIPatch(o["header"], "SystemReplacemsgGroupAlertmail-Header"); ok {
			if err = d.Set("header", vv); err != nil {
				return fmt.Errorf("Error reading header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemReplacemsgGroupAlertmailId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemReplacemsgGroupAlertmail-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("msg_type", flattenSystemReplacemsgGroupAlertmailMsgType2edl(o["msg-type"], d, "msg_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["msg-type"], "SystemReplacemsgGroupAlertmail-MsgType"); ok {
			if err = d.Set("msg_type", vv); err != nil {
				return fmt.Errorf("Error reading msg_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading msg_type: %v", err)
		}
	}

	return nil
}

func flattenSystemReplacemsgGroupAlertmailFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemReplacemsgGroupAlertmailBuffer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmailFormat2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmailHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmailId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmailMsgType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReplacemsgGroupAlertmail(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("buffer"); ok || d.HasChange("buffer") {
		t, err := expandSystemReplacemsgGroupAlertmailBuffer2edl(d, v, "buffer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["buffer"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok || d.HasChange("format") {
		t, err := expandSystemReplacemsgGroupAlertmailFormat2edl(d, v, "format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok || d.HasChange("header") {
		t, err := expandSystemReplacemsgGroupAlertmailHeader2edl(d, v, "header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemReplacemsgGroupAlertmailId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("msg_type"); ok || d.HasChange("msg_type") {
		t, err := expandSystemReplacemsgGroupAlertmailMsgType2edl(d, v, "msg_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msg-type"] = t
		}
	}

	return &obj, nil
}
