// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure ICAP forwarded request headers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIcapProfileIcapHeaders() *schema.Resource {
	return &schema.Resource{
		Create: resourceIcapProfileIcapHeadersCreate,
		Read:   resourceIcapProfileIcapHeadersRead,
		Update: resourceIcapProfileIcapHeadersUpdate,
		Delete: resourceIcapProfileIcapHeadersDelete,

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
			"base64_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sesson_info_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceIcapProfileIcapHeadersCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIcapProfileIcapHeaders(d)
	if err != nil {
		return fmt.Errorf("Error creating IcapProfileIcapHeaders resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIcapProfileIcapHeaders(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIcapProfileIcapHeaders(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IcapProfileIcapHeaders resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateIcapProfileIcapHeaders(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IcapProfileIcapHeaders resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceIcapProfileIcapHeadersRead(d, m)
			} else {
				return fmt.Errorf("Error creating IcapProfileIcapHeaders resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceIcapProfileIcapHeadersRead(d, m)
}

func resourceIcapProfileIcapHeadersUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIcapProfileIcapHeaders(d)
	if err != nil {
		return fmt.Errorf("Error updating IcapProfileIcapHeaders resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIcapProfileIcapHeaders(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IcapProfileIcapHeaders resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceIcapProfileIcapHeadersRead(d, m)
}

func resourceIcapProfileIcapHeadersDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteIcapProfileIcapHeaders(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IcapProfileIcapHeaders resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIcapProfileIcapHeadersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadIcapProfileIcapHeaders(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IcapProfileIcapHeaders resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIcapProfileIcapHeaders(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IcapProfileIcapHeaders resource from API: %v", err)
	}
	return nil
}

func flattenIcapProfileIcapHeadersBase64Encoding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersContent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersHttpHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersSessonInfoType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIcapProfileIcapHeaders(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("base64_encoding", flattenIcapProfileIcapHeadersBase64Encoding2edl(o["base64-encoding"], d, "base64_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["base64-encoding"], "IcapProfileIcapHeaders-Base64Encoding"); ok {
			if err = d.Set("base64_encoding", vv); err != nil {
				return fmt.Errorf("Error reading base64_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading base64_encoding: %v", err)
		}
	}

	if err = d.Set("content", flattenIcapProfileIcapHeadersContent2edl(o["content"], d, "content")); err != nil {
		if vv, ok := fortiAPIPatch(o["content"], "IcapProfileIcapHeaders-Content"); ok {
			if err = d.Set("content", vv); err != nil {
				return fmt.Errorf("Error reading content: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content: %v", err)
		}
	}

	if err = d.Set("fosid", flattenIcapProfileIcapHeadersId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "IcapProfileIcapHeaders-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenIcapProfileIcapHeadersName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "IcapProfileIcapHeaders-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("http_header", flattenIcapProfileIcapHeadersHttpHeader2edl(o["http-header"], d, "http_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-header"], "IcapProfileIcapHeaders-HttpHeader"); ok {
			if err = d.Set("http_header", vv); err != nil {
				return fmt.Errorf("Error reading http_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_header: %v", err)
		}
	}

	if err = d.Set("sesson_info_type", flattenIcapProfileIcapHeadersSessonInfoType2edl(o["sesson-info-type"], d, "sesson_info_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["sesson-info-type"], "IcapProfileIcapHeaders-SessonInfoType"); ok {
			if err = d.Set("sesson_info_type", vv); err != nil {
				return fmt.Errorf("Error reading sesson_info_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sesson_info_type: %v", err)
		}
	}

	if err = d.Set("source", flattenIcapProfileIcapHeadersSource2edl(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "IcapProfileIcapHeaders-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	return nil
}

func flattenIcapProfileIcapHeadersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIcapProfileIcapHeadersBase64Encoding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersContent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersHttpHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersSessonInfoType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIcapProfileIcapHeaders(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("base64_encoding"); ok || d.HasChange("base64_encoding") {
		t, err := expandIcapProfileIcapHeadersBase64Encoding2edl(d, v, "base64_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base64-encoding"] = t
		}
	}

	if v, ok := d.GetOk("content"); ok || d.HasChange("content") {
		t, err := expandIcapProfileIcapHeadersContent2edl(d, v, "content")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandIcapProfileIcapHeadersId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandIcapProfileIcapHeadersName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("http_header"); ok || d.HasChange("http_header") {
		t, err := expandIcapProfileIcapHeadersHttpHeader2edl(d, v, "http_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-header"] = t
		}
	}

	if v, ok := d.GetOk("sesson_info_type"); ok || d.HasChange("sesson_info_type") {
		t, err := expandIcapProfileIcapHeadersSessonInfoType2edl(d, v, "sesson_info_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sesson-info-type"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandIcapProfileIcapHeadersSource2edl(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	return &obj, nil
}
