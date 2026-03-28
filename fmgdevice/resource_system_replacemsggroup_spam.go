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

func resourceSystemReplacemsgGroupSpam() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReplacemsgGroupSpamCreate,
		Read:   resourceSystemReplacemsgGroupSpamRead,
		Update: resourceSystemReplacemsgGroupSpamUpdate,
		Delete: resourceSystemReplacemsgGroupSpamDelete,

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
			"msg_type": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemReplacemsgGroupSpamCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemReplacemsgGroupSpam(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemReplacemsgGroupSpam resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("msg_type")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemReplacemsgGroupSpam(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemReplacemsgGroupSpam(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemReplacemsgGroupSpam resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemReplacemsgGroupSpam(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemReplacemsgGroupSpam resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "msg_type"))

	return resourceSystemReplacemsgGroupSpamRead(d, m)
}

func resourceSystemReplacemsgGroupSpamUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemReplacemsgGroupSpam(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgGroupSpam resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemReplacemsgGroupSpam(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgGroupSpam resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "msg_type"))

	return resourceSystemReplacemsgGroupSpamRead(d, m)
}

func resourceSystemReplacemsgGroupSpamDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemReplacemsgGroupSpam(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReplacemsgGroupSpam resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgGroupSpamRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemReplacemsgGroupSpam(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemReplacemsgGroupSpam resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReplacemsgGroupSpam(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgGroupSpam resource from API: %v", err)
	}
	return nil
}

func flattenSystemReplacemsgGroupSpamBuffer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSpamFormat2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSpamHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSpamMsgType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemReplacemsgGroupSpam(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("buffer", flattenSystemReplacemsgGroupSpamBuffer2edl(o["buffer"], d, "buffer")); err != nil {
		if vv, ok := fortiAPIPatch(o["buffer"], "SystemReplacemsgGroupSpam-Buffer"); ok {
			if err = d.Set("buffer", vv); err != nil {
				return fmt.Errorf("Error reading buffer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading buffer: %v", err)
		}
	}

	if err = d.Set("format", flattenSystemReplacemsgGroupSpamFormat2edl(o["format"], d, "format")); err != nil {
		if vv, ok := fortiAPIPatch(o["format"], "SystemReplacemsgGroupSpam-Format"); ok {
			if err = d.Set("format", vv); err != nil {
				return fmt.Errorf("Error reading format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("header", flattenSystemReplacemsgGroupSpamHeader2edl(o["header"], d, "header")); err != nil {
		if vv, ok := fortiAPIPatch(o["header"], "SystemReplacemsgGroupSpam-Header"); ok {
			if err = d.Set("header", vv); err != nil {
				return fmt.Errorf("Error reading header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("msg_type", flattenSystemReplacemsgGroupSpamMsgType2edl(o["msg-type"], d, "msg_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["msg-type"], "SystemReplacemsgGroupSpam-MsgType"); ok {
			if err = d.Set("msg_type", vv); err != nil {
				return fmt.Errorf("Error reading msg_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading msg_type: %v", err)
		}
	}

	return nil
}

func flattenSystemReplacemsgGroupSpamFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemReplacemsgGroupSpamBuffer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSpamFormat2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSpamHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSpamMsgType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReplacemsgGroupSpam(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("buffer"); ok || d.HasChange("buffer") {
		t, err := expandSystemReplacemsgGroupSpamBuffer2edl(d, v, "buffer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["buffer"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok || d.HasChange("format") {
		t, err := expandSystemReplacemsgGroupSpamFormat2edl(d, v, "format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok || d.HasChange("header") {
		t, err := expandSystemReplacemsgGroupSpamHeader2edl(d, v, "header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("msg_type"); ok || d.HasChange("msg_type") {
		t, err := expandSystemReplacemsgGroupSpamMsgType2edl(d, v, "msg_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msg-type"] = t
		}
	}

	return &obj, nil
}
