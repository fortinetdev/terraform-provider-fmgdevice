// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure HTTP forwarded requests headers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyProfileHeaders() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyProfileHeadersCreate,
		Read:   resourceWebProxyProfileHeadersRead,
		Update: resourceWebProxyProfileHeadersUpdate,
		Delete: resourceWebProxyProfileHeadersDelete,

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
			"add_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"dstaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebProxyProfileHeadersCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyProfileHeaders(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyProfileHeaders resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebProxyProfileHeaders(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebProxyProfileHeaders(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebProxyProfileHeaders resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateWebProxyProfileHeaders(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebProxyProfileHeaders resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceWebProxyProfileHeadersRead(d, m)
			} else {
				return fmt.Errorf("Error creating WebProxyProfileHeaders resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebProxyProfileHeadersRead(d, m)
}

func resourceWebProxyProfileHeadersUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyProfileHeaders(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyProfileHeaders resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebProxyProfileHeaders(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyProfileHeaders resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebProxyProfileHeadersRead(d, m)
}

func resourceWebProxyProfileHeadersDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebProxyProfileHeaders(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyProfileHeaders resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyProfileHeadersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyProfileHeaders(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebProxyProfileHeaders resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyProfileHeaders(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyProfileHeaders resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyProfileHeadersAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersAddOption2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersBase64Encoding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersContent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersDstaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyProfileHeadersDstaddr62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyProfileHeadersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWebProxyProfileHeaders(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenWebProxyProfileHeadersAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "WebProxyProfileHeaders-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("add_option", flattenWebProxyProfileHeadersAddOption2edl(o["add-option"], d, "add_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-option"], "WebProxyProfileHeaders-AddOption"); ok {
			if err = d.Set("add_option", vv); err != nil {
				return fmt.Errorf("Error reading add_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_option: %v", err)
		}
	}

	if err = d.Set("base64_encoding", flattenWebProxyProfileHeadersBase64Encoding2edl(o["base64-encoding"], d, "base64_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["base64-encoding"], "WebProxyProfileHeaders-Base64Encoding"); ok {
			if err = d.Set("base64_encoding", vv); err != nil {
				return fmt.Errorf("Error reading base64_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading base64_encoding: %v", err)
		}
	}

	if err = d.Set("content", flattenWebProxyProfileHeadersContent2edl(o["content"], d, "content")); err != nil {
		if vv, ok := fortiAPIPatch(o["content"], "WebProxyProfileHeaders-Content"); ok {
			if err = d.Set("content", vv); err != nil {
				return fmt.Errorf("Error reading content: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenWebProxyProfileHeadersDstaddr2edl(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "WebProxyProfileHeaders-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenWebProxyProfileHeadersDstaddr62edl(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "WebProxyProfileHeaders-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWebProxyProfileHeadersId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WebProxyProfileHeaders-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenWebProxyProfileHeadersName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebProxyProfileHeaders-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenWebProxyProfileHeadersProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "WebProxyProfileHeaders-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	return nil
}

func flattenWebProxyProfileHeadersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyProfileHeadersAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersAddOption2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersBase64Encoding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersContent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersDstaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyProfileHeadersDstaddr62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyProfileHeadersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWebProxyProfileHeaders(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandWebProxyProfileHeadersAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("add_option"); ok || d.HasChange("add_option") {
		t, err := expandWebProxyProfileHeadersAddOption2edl(d, v, "add_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-option"] = t
		}
	}

	if v, ok := d.GetOk("base64_encoding"); ok || d.HasChange("base64_encoding") {
		t, err := expandWebProxyProfileHeadersBase64Encoding2edl(d, v, "base64_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base64-encoding"] = t
		}
	}

	if v, ok := d.GetOk("content"); ok || d.HasChange("content") {
		t, err := expandWebProxyProfileHeadersContent2edl(d, v, "content")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandWebProxyProfileHeadersDstaddr2edl(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandWebProxyProfileHeadersDstaddr62edl(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWebProxyProfileHeadersId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebProxyProfileHeadersName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandWebProxyProfileHeadersProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	return &obj, nil
}
