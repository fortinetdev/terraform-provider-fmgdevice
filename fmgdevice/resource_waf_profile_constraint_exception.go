// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> HTTP constraint exception.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfileConstraintException() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileConstraintExceptionCreate,
		Read:   resourceWafProfileConstraintExceptionRead,
		Update: resourceWafProfileConstraintExceptionUpdate,
		Delete: resourceWafProfileConstraintExceptionDelete,

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
			"address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"content_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"line_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_header_line": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_range_segment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_url_param": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"param_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_param_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWafProfileConstraintExceptionCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWafProfileConstraintException(d)
	if err != nil {
		return fmt.Errorf("Error creating WafProfileConstraintException resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWafProfileConstraintException(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWafProfileConstraintException(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WafProfileConstraintException resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWafProfileConstraintException(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WafProfileConstraintException resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWafProfileConstraintExceptionRead(d, m)
}

func resourceWafProfileConstraintExceptionUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWafProfileConstraintException(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileConstraintException resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfileConstraintException(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileConstraintException resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWafProfileConstraintExceptionRead(d, m)
}

func resourceWafProfileConstraintExceptionDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWafProfileConstraintException(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfileConstraintException resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileConstraintExceptionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWafProfileConstraintException(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfileConstraintException resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfileConstraintException(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfileConstraintException resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileConstraintExceptionAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileConstraintExceptionContentLength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionHeaderLength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionHostname3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionLineLength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMalformed3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxCookie3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxHeaderLine3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxRangeSegment3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxUrlParam3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMethod3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionParamLength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionPattern3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionRegex3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionUrlParamLength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionVersion3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWafProfileConstraintException(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenWafProfileConstraintExceptionAddress3rdl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "WafProfileConstraintException-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("content_length", flattenWafProfileConstraintExceptionContentLength3rdl(o["content-length"], d, "content_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["content-length"], "WafProfileConstraintException-ContentLength"); ok {
			if err = d.Set("content_length", vv); err != nil {
				return fmt.Errorf("Error reading content_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content_length: %v", err)
		}
	}

	if err = d.Set("header_length", flattenWafProfileConstraintExceptionHeaderLength3rdl(o["header-length"], d, "header_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-length"], "WafProfileConstraintException-HeaderLength"); ok {
			if err = d.Set("header_length", vv); err != nil {
				return fmt.Errorf("Error reading header_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_length: %v", err)
		}
	}

	if err = d.Set("hostname", flattenWafProfileConstraintExceptionHostname3rdl(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "WafProfileConstraintException-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWafProfileConstraintExceptionId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WafProfileConstraintException-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("line_length", flattenWafProfileConstraintExceptionLineLength3rdl(o["line-length"], d, "line_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["line-length"], "WafProfileConstraintException-LineLength"); ok {
			if err = d.Set("line_length", vv); err != nil {
				return fmt.Errorf("Error reading line_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading line_length: %v", err)
		}
	}

	if err = d.Set("malformed", flattenWafProfileConstraintExceptionMalformed3rdl(o["malformed"], d, "malformed")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed"], "WafProfileConstraintException-Malformed"); ok {
			if err = d.Set("malformed", vv); err != nil {
				return fmt.Errorf("Error reading malformed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed: %v", err)
		}
	}

	if err = d.Set("max_cookie", flattenWafProfileConstraintExceptionMaxCookie3rdl(o["max-cookie"], d, "max_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-cookie"], "WafProfileConstraintException-MaxCookie"); ok {
			if err = d.Set("max_cookie", vv); err != nil {
				return fmt.Errorf("Error reading max_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_cookie: %v", err)
		}
	}

	if err = d.Set("max_header_line", flattenWafProfileConstraintExceptionMaxHeaderLine3rdl(o["max-header-line"], d, "max_header_line")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-header-line"], "WafProfileConstraintException-MaxHeaderLine"); ok {
			if err = d.Set("max_header_line", vv); err != nil {
				return fmt.Errorf("Error reading max_header_line: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_header_line: %v", err)
		}
	}

	if err = d.Set("max_range_segment", flattenWafProfileConstraintExceptionMaxRangeSegment3rdl(o["max-range-segment"], d, "max_range_segment")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-range-segment"], "WafProfileConstraintException-MaxRangeSegment"); ok {
			if err = d.Set("max_range_segment", vv); err != nil {
				return fmt.Errorf("Error reading max_range_segment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_range_segment: %v", err)
		}
	}

	if err = d.Set("max_url_param", flattenWafProfileConstraintExceptionMaxUrlParam3rdl(o["max-url-param"], d, "max_url_param")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-url-param"], "WafProfileConstraintException-MaxUrlParam"); ok {
			if err = d.Set("max_url_param", vv); err != nil {
				return fmt.Errorf("Error reading max_url_param: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_url_param: %v", err)
		}
	}

	if err = d.Set("method", flattenWafProfileConstraintExceptionMethod3rdl(o["method"], d, "method")); err != nil {
		if vv, ok := fortiAPIPatch(o["method"], "WafProfileConstraintException-Method"); ok {
			if err = d.Set("method", vv); err != nil {
				return fmt.Errorf("Error reading method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("param_length", flattenWafProfileConstraintExceptionParamLength3rdl(o["param-length"], d, "param_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["param-length"], "WafProfileConstraintException-ParamLength"); ok {
			if err = d.Set("param_length", vv); err != nil {
				return fmt.Errorf("Error reading param_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading param_length: %v", err)
		}
	}

	if err = d.Set("pattern", flattenWafProfileConstraintExceptionPattern3rdl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "WafProfileConstraintException-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("regex", flattenWafProfileConstraintExceptionRegex3rdl(o["regex"], d, "regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["regex"], "WafProfileConstraintException-Regex"); ok {
			if err = d.Set("regex", vv); err != nil {
				return fmt.Errorf("Error reading regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regex: %v", err)
		}
	}

	if err = d.Set("url_param_length", flattenWafProfileConstraintExceptionUrlParamLength3rdl(o["url-param-length"], d, "url_param_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-param-length"], "WafProfileConstraintException-UrlParamLength"); ok {
			if err = d.Set("url_param_length", vv); err != nil {
				return fmt.Errorf("Error reading url_param_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_param_length: %v", err)
		}
	}

	if err = d.Set("version", flattenWafProfileConstraintExceptionVersion3rdl(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "WafProfileConstraintException-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	return nil
}

func flattenWafProfileConstraintExceptionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileConstraintExceptionAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileConstraintExceptionContentLength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionHeaderLength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionHostname3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionLineLength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMalformed3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxCookie3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxHeaderLine3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxRangeSegment3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxUrlParam3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMethod3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionParamLength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionPattern3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionRegex3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionUrlParamLength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionVersion3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWafProfileConstraintException(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandWafProfileConstraintExceptionAddress3rdl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("content_length"); ok || d.HasChange("content_length") {
		t, err := expandWafProfileConstraintExceptionContentLength3rdl(d, v, "content_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-length"] = t
		}
	}

	if v, ok := d.GetOk("header_length"); ok || d.HasChange("header_length") {
		t, err := expandWafProfileConstraintExceptionHeaderLength3rdl(d, v, "header_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-length"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandWafProfileConstraintExceptionHostname3rdl(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWafProfileConstraintExceptionId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("line_length"); ok || d.HasChange("line_length") {
		t, err := expandWafProfileConstraintExceptionLineLength3rdl(d, v, "line_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["line-length"] = t
		}
	}

	if v, ok := d.GetOk("malformed"); ok || d.HasChange("malformed") {
		t, err := expandWafProfileConstraintExceptionMalformed3rdl(d, v, "malformed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed"] = t
		}
	}

	if v, ok := d.GetOk("max_cookie"); ok || d.HasChange("max_cookie") {
		t, err := expandWafProfileConstraintExceptionMaxCookie3rdl(d, v, "max_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cookie"] = t
		}
	}

	if v, ok := d.GetOk("max_header_line"); ok || d.HasChange("max_header_line") {
		t, err := expandWafProfileConstraintExceptionMaxHeaderLine3rdl(d, v, "max_header_line")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-header-line"] = t
		}
	}

	if v, ok := d.GetOk("max_range_segment"); ok || d.HasChange("max_range_segment") {
		t, err := expandWafProfileConstraintExceptionMaxRangeSegment3rdl(d, v, "max_range_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-range-segment"] = t
		}
	}

	if v, ok := d.GetOk("max_url_param"); ok || d.HasChange("max_url_param") {
		t, err := expandWafProfileConstraintExceptionMaxUrlParam3rdl(d, v, "max_url_param")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandWafProfileConstraintExceptionMethod3rdl(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("param_length"); ok || d.HasChange("param_length") {
		t, err := expandWafProfileConstraintExceptionParamLength3rdl(d, v, "param_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["param-length"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandWafProfileConstraintExceptionPattern3rdl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("regex"); ok || d.HasChange("regex") {
		t, err := expandWafProfileConstraintExceptionRegex3rdl(d, v, "regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regex"] = t
		}
	}

	if v, ok := d.GetOk("url_param_length"); ok || d.HasChange("url_param_length") {
		t, err := expandWafProfileConstraintExceptionUrlParamLength3rdl(d, v, "url_param_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-length"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandWafProfileConstraintExceptionVersion3rdl(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	return &obj, nil
}
