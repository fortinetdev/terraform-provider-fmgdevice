// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure web proxy profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyProfileCreate,
		Read:   resourceWebProxyProfileRead,
		Update: resourceWebProxyProfileUpdate,
		Delete: resourceWebProxyProfileDelete,

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
			"header_client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_front_end_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_via_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_via_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_x_authenticated_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_x_authenticated_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_x_forwarded_client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_x_forwarded_for": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
				},
			},
			"log_header_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"strip_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_cache_object_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWebProxyProfileCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectWebProxyProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebProxyProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebProxyProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebProxyProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebProxyProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebProxyProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyProfileRead(d, m)
}

func resourceWebProxyProfileUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectWebProxyProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebProxyProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyProfileRead(d, m)
}

func resourceWebProxyProfileDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteWebProxyProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadWebProxyProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebProxyProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyProfile resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyProfileHeaderClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderFrontEndHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderViaRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderViaResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderXAuthenticatedGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderXAuthenticatedUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderXForwardedClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderXForwardedFor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaders(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenWebProxyProfileHeadersAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WebProxyProfile-Headers-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_option"
		if _, ok := i["add-option"]; ok {
			v := flattenWebProxyProfileHeadersAddOption(i["add-option"], d, pre_append)
			tmp["add_option"] = fortiAPISubPartPatch(v, "WebProxyProfile-Headers-AddOption")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := i["base64-encoding"]; ok {
			v := flattenWebProxyProfileHeadersBase64Encoding(i["base64-encoding"], d, pre_append)
			tmp["base64_encoding"] = fortiAPISubPartPatch(v, "WebProxyProfile-Headers-Base64Encoding")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := i["content"]; ok {
			v := flattenWebProxyProfileHeadersContent(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "WebProxyProfile-Headers-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {
			v := flattenWebProxyProfileHeadersDstaddr(i["dstaddr"], d, pre_append)
			tmp["dstaddr"] = fortiAPISubPartPatch(v, "WebProxyProfile-Headers-Dstaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := i["dstaddr6"]; ok {
			v := flattenWebProxyProfileHeadersDstaddr6(i["dstaddr6"], d, pre_append)
			tmp["dstaddr6"] = fortiAPISubPartPatch(v, "WebProxyProfile-Headers-Dstaddr6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWebProxyProfileHeadersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WebProxyProfile-Headers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWebProxyProfileHeadersName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WebProxyProfile-Headers-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenWebProxyProfileHeadersProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "WebProxyProfile-Headers-Protocol")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebProxyProfileHeadersAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersAddOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersBase64Encoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyProfileHeadersDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyProfileHeadersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyProfileLogHeaderChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileStripEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileMaxCacheObjectSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("header_client_ip", flattenWebProxyProfileHeaderClientIp(o["header-client-ip"], d, "header_client_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-client-ip"], "WebProxyProfile-HeaderClientIp"); ok {
			if err = d.Set("header_client_ip", vv); err != nil {
				return fmt.Errorf("Error reading header_client_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_client_ip: %v", err)
		}
	}

	if err = d.Set("header_front_end_https", flattenWebProxyProfileHeaderFrontEndHttps(o["header-front-end-https"], d, "header_front_end_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-front-end-https"], "WebProxyProfile-HeaderFrontEndHttps"); ok {
			if err = d.Set("header_front_end_https", vv); err != nil {
				return fmt.Errorf("Error reading header_front_end_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_front_end_https: %v", err)
		}
	}

	if err = d.Set("header_via_request", flattenWebProxyProfileHeaderViaRequest(o["header-via-request"], d, "header_via_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-via-request"], "WebProxyProfile-HeaderViaRequest"); ok {
			if err = d.Set("header_via_request", vv); err != nil {
				return fmt.Errorf("Error reading header_via_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_via_request: %v", err)
		}
	}

	if err = d.Set("header_via_response", flattenWebProxyProfileHeaderViaResponse(o["header-via-response"], d, "header_via_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-via-response"], "WebProxyProfile-HeaderViaResponse"); ok {
			if err = d.Set("header_via_response", vv); err != nil {
				return fmt.Errorf("Error reading header_via_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_via_response: %v", err)
		}
	}

	if err = d.Set("header_x_authenticated_groups", flattenWebProxyProfileHeaderXAuthenticatedGroups(o["header-x-authenticated-groups"], d, "header_x_authenticated_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-x-authenticated-groups"], "WebProxyProfile-HeaderXAuthenticatedGroups"); ok {
			if err = d.Set("header_x_authenticated_groups", vv); err != nil {
				return fmt.Errorf("Error reading header_x_authenticated_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_x_authenticated_groups: %v", err)
		}
	}

	if err = d.Set("header_x_authenticated_user", flattenWebProxyProfileHeaderXAuthenticatedUser(o["header-x-authenticated-user"], d, "header_x_authenticated_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-x-authenticated-user"], "WebProxyProfile-HeaderXAuthenticatedUser"); ok {
			if err = d.Set("header_x_authenticated_user", vv); err != nil {
				return fmt.Errorf("Error reading header_x_authenticated_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_x_authenticated_user: %v", err)
		}
	}

	if err = d.Set("header_x_forwarded_client_cert", flattenWebProxyProfileHeaderXForwardedClientCert(o["header-x-forwarded-client-cert"], d, "header_x_forwarded_client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-x-forwarded-client-cert"], "WebProxyProfile-HeaderXForwardedClientCert"); ok {
			if err = d.Set("header_x_forwarded_client_cert", vv); err != nil {
				return fmt.Errorf("Error reading header_x_forwarded_client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_x_forwarded_client_cert: %v", err)
		}
	}

	if err = d.Set("header_x_forwarded_for", flattenWebProxyProfileHeaderXForwardedFor(o["header-x-forwarded-for"], d, "header_x_forwarded_for")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-x-forwarded-for"], "WebProxyProfile-HeaderXForwardedFor"); ok {
			if err = d.Set("header_x_forwarded_for", vv); err != nil {
				return fmt.Errorf("Error reading header_x_forwarded_for: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_x_forwarded_for: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("headers", flattenWebProxyProfileHeaders(o["headers"], d, "headers")); err != nil {
			if vv, ok := fortiAPIPatch(o["headers"], "WebProxyProfile-Headers"); ok {
				if err = d.Set("headers", vv); err != nil {
					return fmt.Errorf("Error reading headers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading headers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("headers"); ok {
			if err = d.Set("headers", flattenWebProxyProfileHeaders(o["headers"], d, "headers")); err != nil {
				if vv, ok := fortiAPIPatch(o["headers"], "WebProxyProfile-Headers"); ok {
					if err = d.Set("headers", vv); err != nil {
						return fmt.Errorf("Error reading headers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading headers: %v", err)
				}
			}
		}
	}

	if err = d.Set("log_header_change", flattenWebProxyProfileLogHeaderChange(o["log-header-change"], d, "log_header_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-header-change"], "WebProxyProfile-LogHeaderChange"); ok {
			if err = d.Set("log_header_change", vv); err != nil {
				return fmt.Errorf("Error reading log_header_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_header_change: %v", err)
		}
	}

	if err = d.Set("name", flattenWebProxyProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebProxyProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("strip_encoding", flattenWebProxyProfileStripEncoding(o["strip-encoding"], d, "strip_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["strip-encoding"], "WebProxyProfile-StripEncoding"); ok {
			if err = d.Set("strip_encoding", vv); err != nil {
				return fmt.Errorf("Error reading strip_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strip_encoding: %v", err)
		}
	}

	if err = d.Set("max_cache_object_size", flattenWebProxyProfileMaxCacheObjectSize(o["max-cache-object-size"], d, "max_cache_object_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-cache-object-size"], "WebProxyProfile-MaxCacheObjectSize"); ok {
			if err = d.Set("max_cache_object_size", vv); err != nil {
				return fmt.Errorf("Error reading max_cache_object_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_cache_object_size: %v", err)
		}
	}

	return nil
}

func flattenWebProxyProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyProfileHeaderClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderFrontEndHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderViaRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderViaResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderXAuthenticatedGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderXAuthenticatedUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderXForwardedClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderXForwardedFor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaders(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandWebProxyProfileHeadersAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_option"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-option"], _ = expandWebProxyProfileHeadersAddOption(d, i["add_option"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["base64-encoding"], _ = expandWebProxyProfileHeadersBase64Encoding(d, i["base64_encoding"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["content"], _ = expandWebProxyProfileHeadersContent(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr"], _ = expandWebProxyProfileHeadersDstaddr(d, i["dstaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr6"], _ = expandWebProxyProfileHeadersDstaddr6(d, i["dstaddr6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWebProxyProfileHeadersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWebProxyProfileHeadersName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandWebProxyProfileHeadersProtocol(d, i["protocol"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebProxyProfileHeadersAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersAddOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersBase64Encoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyProfileHeadersDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyProfileHeadersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyProfileLogHeaderChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileStripEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileMaxCacheObjectSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("header_client_ip"); ok || d.HasChange("header_client_ip") {
		t, err := expandWebProxyProfileHeaderClientIp(d, v, "header_client_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-client-ip"] = t
		}
	}

	if v, ok := d.GetOk("header_front_end_https"); ok || d.HasChange("header_front_end_https") {
		t, err := expandWebProxyProfileHeaderFrontEndHttps(d, v, "header_front_end_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-front-end-https"] = t
		}
	}

	if v, ok := d.GetOk("header_via_request"); ok || d.HasChange("header_via_request") {
		t, err := expandWebProxyProfileHeaderViaRequest(d, v, "header_via_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-via-request"] = t
		}
	}

	if v, ok := d.GetOk("header_via_response"); ok || d.HasChange("header_via_response") {
		t, err := expandWebProxyProfileHeaderViaResponse(d, v, "header_via_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-via-response"] = t
		}
	}

	if v, ok := d.GetOk("header_x_authenticated_groups"); ok || d.HasChange("header_x_authenticated_groups") {
		t, err := expandWebProxyProfileHeaderXAuthenticatedGroups(d, v, "header_x_authenticated_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-authenticated-groups"] = t
		}
	}

	if v, ok := d.GetOk("header_x_authenticated_user"); ok || d.HasChange("header_x_authenticated_user") {
		t, err := expandWebProxyProfileHeaderXAuthenticatedUser(d, v, "header_x_authenticated_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-authenticated-user"] = t
		}
	}

	if v, ok := d.GetOk("header_x_forwarded_client_cert"); ok || d.HasChange("header_x_forwarded_client_cert") {
		t, err := expandWebProxyProfileHeaderXForwardedClientCert(d, v, "header_x_forwarded_client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-forwarded-client-cert"] = t
		}
	}

	if v, ok := d.GetOk("header_x_forwarded_for"); ok || d.HasChange("header_x_forwarded_for") {
		t, err := expandWebProxyProfileHeaderXForwardedFor(d, v, "header_x_forwarded_for")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-forwarded-for"] = t
		}
	}

	if v, ok := d.GetOk("headers"); ok || d.HasChange("headers") {
		t, err := expandWebProxyProfileHeaders(d, v, "headers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["headers"] = t
		}
	}

	if v, ok := d.GetOk("log_header_change"); ok || d.HasChange("log_header_change") {
		t, err := expandWebProxyProfileLogHeaderChange(d, v, "log_header_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-header-change"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebProxyProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("strip_encoding"); ok || d.HasChange("strip_encoding") {
		t, err := expandWebProxyProfileStripEncoding(d, v, "strip_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-encoding"] = t
		}
	}

	if v, ok := d.GetOk("max_cache_object_size"); ok || d.HasChange("max_cache_object_size") {
		t, err := expandWebProxyProfileMaxCacheObjectSize(d, v, "max_cache_object_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cache-object-size"] = t
		}
	}

	return &obj, nil
}
