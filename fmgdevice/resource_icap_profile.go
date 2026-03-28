// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure ICAP profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIcapProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceIcapProfileCreate,
		Read:   resourceIcapProfileRead,
		Update: resourceIcapProfileUpdate,
		Delete: resourceIcapProfileDelete,

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
			"n204_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n204_size_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"chunk_encap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extension_feature": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"file_transfer": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"file_transfer_failure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_transfer_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_transfer_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"icap_block_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icap_headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base64_encoding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
				},
			},
			"methods": &schema.Schema{
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
			"ocr_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preview": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preview_data_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_failure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"respmod_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"respmod_forward_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header_group": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"case_sensitivity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"header": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"header_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"host": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"http_resp_status_code": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_failure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_req_hdr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"scan_progress_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"streaming_content_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"allow_204_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scan_oversize_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scan_size_limit": &schema.Schema{
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

func resourceIcapProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIcapProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating IcapProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIcapProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIcapProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IcapProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateIcapProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IcapProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceIcapProfileRead(d, m)
}

func resourceIcapProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIcapProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating IcapProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIcapProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IcapProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceIcapProfileRead(d, m)
}

func resourceIcapProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteIcapProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IcapProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIcapProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadIcapProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IcapProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIcapProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IcapProfile resource from API: %v", err)
	}
	return nil
}

func flattenIcapProfile204Response(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfile204SizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileChunkEncap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileExtensionFeature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapProfileFileTransfer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapProfileFileTransferFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileFileTransferPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileFileTransferServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapProfileIcapBlockLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeaders(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := i["base64-encoding"]; ok {
			v := flattenIcapProfileIcapHeadersBase64Encoding(i["base64-encoding"], d, pre_append)
			tmp["base64_encoding"] = fortiAPISubPartPatch(v, "IcapProfile-IcapHeaders-Base64Encoding")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := i["content"]; ok {
			v := flattenIcapProfileIcapHeadersContent(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "IcapProfile-IcapHeaders-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenIcapProfileIcapHeadersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "IcapProfile-IcapHeaders-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenIcapProfileIcapHeadersName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "IcapProfile-IcapHeaders-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_header"
		if _, ok := i["http-header"]; ok {
			v := flattenIcapProfileIcapHeadersHttpHeader(i["http-header"], d, pre_append)
			tmp["http_header"] = fortiAPISubPartPatch(v, "IcapProfile-IcapHeaders-HttpHeader")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sesson_info_type"
		if _, ok := i["sesson-info-type"]; ok {
			v := flattenIcapProfileIcapHeadersSessonInfoType(i["sesson-info-type"], d, pre_append)
			tmp["sesson_info_type"] = fortiAPISubPartPatch(v, "IcapProfile-IcapHeaders-SessonInfoType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {
			v := flattenIcapProfileIcapHeadersSource(i["source"], d, pre_append)
			tmp["source"] = fortiAPISubPartPatch(v, "IcapProfile-IcapHeaders-Source")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenIcapProfileIcapHeadersBase64Encoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersHttpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersSessonInfoType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileMethods(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileOcrOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfilePreview(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfilePreviewDataLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapProfileRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRequestFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRequestPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRequestServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapProfileRespmodDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRules(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenIcapProfileRespmodForwardRulesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "IcapProfile-RespmodForwardRules-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_group"
		if _, ok := i["header-group"]; ok {
			v := flattenIcapProfileRespmodForwardRulesHeaderGroup(i["header-group"], d, pre_append)
			tmp["header_group"] = fortiAPISubPartPatch(v, "IcapProfile-RespmodForwardRules-HeaderGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenIcapProfileRespmodForwardRulesHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "IcapProfile-RespmodForwardRules-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_resp_status_code"
		if _, ok := i["http-resp-status-code"]; ok {
			v := flattenIcapProfileRespmodForwardRulesHttpRespStatusCode(i["http-resp-status-code"], d, pre_append)
			tmp["http_resp_status_code"] = fortiAPISubPartPatch(v, "IcapProfile-RespmodForwardRules-HttpRespStatusCode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenIcapProfileRespmodForwardRulesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "IcapProfile-RespmodForwardRules-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenIcapProfileRespmodForwardRulesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRulesHeaderGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {
			v := flattenIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity(i["case-sensitivity"], d, pre_append)
			tmp["case_sensitivity"] = fortiAPISubPartPatch(v, "IcapProfileRespmodForwardRules-HeaderGroup-CaseSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenIcapProfileRespmodForwardRulesHeaderGroupHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "IcapProfileRespmodForwardRules-HeaderGroup-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenIcapProfileRespmodForwardRulesHeaderGroupHeaderName(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "IcapProfileRespmodForwardRules-HeaderGroup-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenIcapProfileRespmodForwardRulesHeaderGroupId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "IcapProfileRespmodForwardRules-HeaderGroup-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRulesHeaderGroupHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRulesHeaderGroupHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRulesHeaderGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRulesHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapProfileRespmodForwardRulesHttpRespStatusCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenIcapProfileRespmodForwardRulesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileResponseFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileResponsePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileResponseReqHdr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileResponseServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapProfileScanProgressInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileStreamingContentBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileAllow204Response(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileScanOversizeLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileScanSizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIcapProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("n204_response", flattenIcapProfile204Response(o["204-response"], d, "n204_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["204-response"], "IcapProfile-204Response"); ok {
			if err = d.Set("n204_response", vv); err != nil {
				return fmt.Errorf("Error reading n204_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n204_response: %v", err)
		}
	}

	if err = d.Set("n204_size_limit", flattenIcapProfile204SizeLimit(o["204-size-limit"], d, "n204_size_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["204-size-limit"], "IcapProfile-204SizeLimit"); ok {
			if err = d.Set("n204_size_limit", vv); err != nil {
				return fmt.Errorf("Error reading n204_size_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n204_size_limit: %v", err)
		}
	}

	if err = d.Set("chunk_encap", flattenIcapProfileChunkEncap(o["chunk-encap"], d, "chunk_encap")); err != nil {
		if vv, ok := fortiAPIPatch(o["chunk-encap"], "IcapProfile-ChunkEncap"); ok {
			if err = d.Set("chunk_encap", vv); err != nil {
				return fmt.Errorf("Error reading chunk_encap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading chunk_encap: %v", err)
		}
	}

	if err = d.Set("comment", flattenIcapProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "IcapProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("extension_feature", flattenIcapProfileExtensionFeature(o["extension-feature"], d, "extension_feature")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension-feature"], "IcapProfile-ExtensionFeature"); ok {
			if err = d.Set("extension_feature", vv); err != nil {
				return fmt.Errorf("Error reading extension_feature: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension_feature: %v", err)
		}
	}

	if err = d.Set("file_transfer", flattenIcapProfileFileTransfer(o["file-transfer"], d, "file_transfer")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-transfer"], "IcapProfile-FileTransfer"); ok {
			if err = d.Set("file_transfer", vv); err != nil {
				return fmt.Errorf("Error reading file_transfer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_transfer: %v", err)
		}
	}

	if err = d.Set("file_transfer_failure", flattenIcapProfileFileTransferFailure(o["file-transfer-failure"], d, "file_transfer_failure")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-transfer-failure"], "IcapProfile-FileTransferFailure"); ok {
			if err = d.Set("file_transfer_failure", vv); err != nil {
				return fmt.Errorf("Error reading file_transfer_failure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_transfer_failure: %v", err)
		}
	}

	if err = d.Set("file_transfer_path", flattenIcapProfileFileTransferPath(o["file-transfer-path"], d, "file_transfer_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-transfer-path"], "IcapProfile-FileTransferPath"); ok {
			if err = d.Set("file_transfer_path", vv); err != nil {
				return fmt.Errorf("Error reading file_transfer_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_transfer_path: %v", err)
		}
	}

	if err = d.Set("file_transfer_server", flattenIcapProfileFileTransferServer(o["file-transfer-server"], d, "file_transfer_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-transfer-server"], "IcapProfile-FileTransferServer"); ok {
			if err = d.Set("file_transfer_server", vv); err != nil {
				return fmt.Errorf("Error reading file_transfer_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_transfer_server: %v", err)
		}
	}

	if err = d.Set("icap_block_log", flattenIcapProfileIcapBlockLog(o["icap-block-log"], d, "icap_block_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["icap-block-log"], "IcapProfile-IcapBlockLog"); ok {
			if err = d.Set("icap_block_log", vv); err != nil {
				return fmt.Errorf("Error reading icap_block_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icap_block_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("icap_headers", flattenIcapProfileIcapHeaders(o["icap-headers"], d, "icap_headers")); err != nil {
			if vv, ok := fortiAPIPatch(o["icap-headers"], "IcapProfile-IcapHeaders"); ok {
				if err = d.Set("icap_headers", vv); err != nil {
					return fmt.Errorf("Error reading icap_headers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading icap_headers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("icap_headers"); ok {
			if err = d.Set("icap_headers", flattenIcapProfileIcapHeaders(o["icap-headers"], d, "icap_headers")); err != nil {
				if vv, ok := fortiAPIPatch(o["icap-headers"], "IcapProfile-IcapHeaders"); ok {
					if err = d.Set("icap_headers", vv); err != nil {
						return fmt.Errorf("Error reading icap_headers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading icap_headers: %v", err)
				}
			}
		}
	}

	if err = d.Set("methods", flattenIcapProfileMethods(o["methods"], d, "methods")); err != nil {
		if vv, ok := fortiAPIPatch(o["methods"], "IcapProfile-Methods"); ok {
			if err = d.Set("methods", vv); err != nil {
				return fmt.Errorf("Error reading methods: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading methods: %v", err)
		}
	}

	if err = d.Set("name", flattenIcapProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "IcapProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ocr_only", flattenIcapProfileOcrOnly(o["ocr-only"], d, "ocr_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["ocr-only"], "IcapProfile-OcrOnly"); ok {
			if err = d.Set("ocr_only", vv); err != nil {
				return fmt.Errorf("Error reading ocr_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ocr_only: %v", err)
		}
	}

	if err = d.Set("preview", flattenIcapProfilePreview(o["preview"], d, "preview")); err != nil {
		if vv, ok := fortiAPIPatch(o["preview"], "IcapProfile-Preview"); ok {
			if err = d.Set("preview", vv); err != nil {
				return fmt.Errorf("Error reading preview: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preview: %v", err)
		}
	}

	if err = d.Set("preview_data_length", flattenIcapProfilePreviewDataLength(o["preview-data-length"], d, "preview_data_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["preview-data-length"], "IcapProfile-PreviewDataLength"); ok {
			if err = d.Set("preview_data_length", vv); err != nil {
				return fmt.Errorf("Error reading preview_data_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preview_data_length: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenIcapProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "IcapProfile-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("request", flattenIcapProfileRequest(o["request"], d, "request")); err != nil {
		if vv, ok := fortiAPIPatch(o["request"], "IcapProfile-Request"); ok {
			if err = d.Set("request", vv); err != nil {
				return fmt.Errorf("Error reading request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request: %v", err)
		}
	}

	if err = d.Set("request_failure", flattenIcapProfileRequestFailure(o["request-failure"], d, "request_failure")); err != nil {
		if vv, ok := fortiAPIPatch(o["request-failure"], "IcapProfile-RequestFailure"); ok {
			if err = d.Set("request_failure", vv); err != nil {
				return fmt.Errorf("Error reading request_failure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request_failure: %v", err)
		}
	}

	if err = d.Set("request_path", flattenIcapProfileRequestPath(o["request-path"], d, "request_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["request-path"], "IcapProfile-RequestPath"); ok {
			if err = d.Set("request_path", vv); err != nil {
				return fmt.Errorf("Error reading request_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request_path: %v", err)
		}
	}

	if err = d.Set("request_server", flattenIcapProfileRequestServer(o["request-server"], d, "request_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["request-server"], "IcapProfile-RequestServer"); ok {
			if err = d.Set("request_server", vv); err != nil {
				return fmt.Errorf("Error reading request_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request_server: %v", err)
		}
	}

	if err = d.Set("respmod_default_action", flattenIcapProfileRespmodDefaultAction(o["respmod-default-action"], d, "respmod_default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["respmod-default-action"], "IcapProfile-RespmodDefaultAction"); ok {
			if err = d.Set("respmod_default_action", vv); err != nil {
				return fmt.Errorf("Error reading respmod_default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading respmod_default_action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("respmod_forward_rules", flattenIcapProfileRespmodForwardRules(o["respmod-forward-rules"], d, "respmod_forward_rules")); err != nil {
			if vv, ok := fortiAPIPatch(o["respmod-forward-rules"], "IcapProfile-RespmodForwardRules"); ok {
				if err = d.Set("respmod_forward_rules", vv); err != nil {
					return fmt.Errorf("Error reading respmod_forward_rules: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading respmod_forward_rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("respmod_forward_rules"); ok {
			if err = d.Set("respmod_forward_rules", flattenIcapProfileRespmodForwardRules(o["respmod-forward-rules"], d, "respmod_forward_rules")); err != nil {
				if vv, ok := fortiAPIPatch(o["respmod-forward-rules"], "IcapProfile-RespmodForwardRules"); ok {
					if err = d.Set("respmod_forward_rules", vv); err != nil {
						return fmt.Errorf("Error reading respmod_forward_rules: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading respmod_forward_rules: %v", err)
				}
			}
		}
	}

	if err = d.Set("response", flattenIcapProfileResponse(o["response"], d, "response")); err != nil {
		if vv, ok := fortiAPIPatch(o["response"], "IcapProfile-Response"); ok {
			if err = d.Set("response", vv); err != nil {
				return fmt.Errorf("Error reading response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading response: %v", err)
		}
	}

	if err = d.Set("response_failure", flattenIcapProfileResponseFailure(o["response-failure"], d, "response_failure")); err != nil {
		if vv, ok := fortiAPIPatch(o["response-failure"], "IcapProfile-ResponseFailure"); ok {
			if err = d.Set("response_failure", vv); err != nil {
				return fmt.Errorf("Error reading response_failure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading response_failure: %v", err)
		}
	}

	if err = d.Set("response_path", flattenIcapProfileResponsePath(o["response-path"], d, "response_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["response-path"], "IcapProfile-ResponsePath"); ok {
			if err = d.Set("response_path", vv); err != nil {
				return fmt.Errorf("Error reading response_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading response_path: %v", err)
		}
	}

	if err = d.Set("response_req_hdr", flattenIcapProfileResponseReqHdr(o["response-req-hdr"], d, "response_req_hdr")); err != nil {
		if vv, ok := fortiAPIPatch(o["response-req-hdr"], "IcapProfile-ResponseReqHdr"); ok {
			if err = d.Set("response_req_hdr", vv); err != nil {
				return fmt.Errorf("Error reading response_req_hdr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading response_req_hdr: %v", err)
		}
	}

	if err = d.Set("response_server", flattenIcapProfileResponseServer(o["response-server"], d, "response_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["response-server"], "IcapProfile-ResponseServer"); ok {
			if err = d.Set("response_server", vv); err != nil {
				return fmt.Errorf("Error reading response_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading response_server: %v", err)
		}
	}

	if err = d.Set("scan_progress_interval", flattenIcapProfileScanProgressInterval(o["scan-progress-interval"], d, "scan_progress_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-progress-interval"], "IcapProfile-ScanProgressInterval"); ok {
			if err = d.Set("scan_progress_interval", vv); err != nil {
				return fmt.Errorf("Error reading scan_progress_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_progress_interval: %v", err)
		}
	}

	if err = d.Set("streaming_content_bypass", flattenIcapProfileStreamingContentBypass(o["streaming-content-bypass"], d, "streaming_content_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["streaming-content-bypass"], "IcapProfile-StreamingContentBypass"); ok {
			if err = d.Set("streaming_content_bypass", vv); err != nil {
				return fmt.Errorf("Error reading streaming_content_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading streaming_content_bypass: %v", err)
		}
	}

	if err = d.Set("timeout", flattenIcapProfileTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "IcapProfile-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("allow_204_response", flattenIcapProfileAllow204Response(o["allow-204-response"], d, "allow_204_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-204-response"], "IcapProfile-Allow204Response"); ok {
			if err = d.Set("allow_204_response", vv); err != nil {
				return fmt.Errorf("Error reading allow_204_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_204_response: %v", err)
		}
	}

	if err = d.Set("scan_oversize_log", flattenIcapProfileScanOversizeLog(o["scan-oversize-log"], d, "scan_oversize_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-oversize-log"], "IcapProfile-ScanOversizeLog"); ok {
			if err = d.Set("scan_oversize_log", vv); err != nil {
				return fmt.Errorf("Error reading scan_oversize_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_oversize_log: %v", err)
		}
	}

	if err = d.Set("scan_size_limit", flattenIcapProfileScanSizeLimit(o["scan-size-limit"], d, "scan_size_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-size-limit"], "IcapProfile-ScanSizeLimit"); ok {
			if err = d.Set("scan_size_limit", vv); err != nil {
				return fmt.Errorf("Error reading scan_size_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_size_limit: %v", err)
		}
	}

	return nil
}

func flattenIcapProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIcapProfile204Response(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfile204SizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileChunkEncap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileExtensionFeature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapProfileFileTransfer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapProfileFileTransferFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileFileTransferPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileFileTransferServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapProfileIcapBlockLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeaders(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["base64-encoding"], _ = expandIcapProfileIcapHeadersBase64Encoding(d, i["base64_encoding"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["content"], _ = expandIcapProfileIcapHeadersContent(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandIcapProfileIcapHeadersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandIcapProfileIcapHeadersName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-header"], _ = expandIcapProfileIcapHeadersHttpHeader(d, i["http_header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sesson_info_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sesson-info-type"], _ = expandIcapProfileIcapHeadersSessonInfoType(d, i["sesson_info_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source"], _ = expandIcapProfileIcapHeadersSource(d, i["source"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandIcapProfileIcapHeadersBase64Encoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersHttpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersSessonInfoType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileMethods(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileOcrOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfilePreview(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfilePreviewDataLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapProfileRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRequestFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRequestPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRequestServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapProfileRespmodDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandIcapProfileRespmodForwardRulesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandIcapProfileRespmodForwardRulesHeaderGroup(d, i["header_group"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["header-group"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandIcapProfileRespmodForwardRulesHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_resp_status_code"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-resp-status-code"], _ = expandIcapProfileRespmodForwardRulesHttpRespStatusCode(d, i["http_resp_status_code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandIcapProfileRespmodForwardRulesName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandIcapProfileRespmodForwardRulesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRulesHeaderGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitivity"], _ = expandIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity(d, i["case_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandIcapProfileRespmodForwardRulesHeaderGroupHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandIcapProfileRespmodForwardRulesHeaderGroupHeaderName(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandIcapProfileRespmodForwardRulesHeaderGroupId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRulesHeaderGroupHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRulesHeaderGroupHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRulesHeaderGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRulesHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapProfileRespmodForwardRulesHttpRespStatusCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandIcapProfileRespmodForwardRulesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponseFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponsePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponseReqHdr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponseServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapProfileScanProgressInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileStreamingContentBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileAllow204Response(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileScanOversizeLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileScanSizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIcapProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n204_response"); ok || d.HasChange("n204_response") {
		t, err := expandIcapProfile204Response(d, v, "n204_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["204-response"] = t
		}
	}

	if v, ok := d.GetOk("n204_size_limit"); ok || d.HasChange("n204_size_limit") {
		t, err := expandIcapProfile204SizeLimit(d, v, "n204_size_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["204-size-limit"] = t
		}
	}

	if v, ok := d.GetOk("chunk_encap"); ok || d.HasChange("chunk_encap") {
		t, err := expandIcapProfileChunkEncap(d, v, "chunk_encap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chunk-encap"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandIcapProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("extension_feature"); ok || d.HasChange("extension_feature") {
		t, err := expandIcapProfileExtensionFeature(d, v, "extension_feature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-feature"] = t
		}
	}

	if v, ok := d.GetOk("file_transfer"); ok || d.HasChange("file_transfer") {
		t, err := expandIcapProfileFileTransfer(d, v, "file_transfer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-transfer"] = t
		}
	}

	if v, ok := d.GetOk("file_transfer_failure"); ok || d.HasChange("file_transfer_failure") {
		t, err := expandIcapProfileFileTransferFailure(d, v, "file_transfer_failure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-transfer-failure"] = t
		}
	}

	if v, ok := d.GetOk("file_transfer_path"); ok || d.HasChange("file_transfer_path") {
		t, err := expandIcapProfileFileTransferPath(d, v, "file_transfer_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-transfer-path"] = t
		}
	}

	if v, ok := d.GetOk("file_transfer_server"); ok || d.HasChange("file_transfer_server") {
		t, err := expandIcapProfileFileTransferServer(d, v, "file_transfer_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-transfer-server"] = t
		}
	}

	if v, ok := d.GetOk("icap_block_log"); ok || d.HasChange("icap_block_log") {
		t, err := expandIcapProfileIcapBlockLog(d, v, "icap_block_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-block-log"] = t
		}
	}

	if v, ok := d.GetOk("icap_headers"); ok || d.HasChange("icap_headers") {
		t, err := expandIcapProfileIcapHeaders(d, v, "icap_headers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-headers"] = t
		}
	}

	if v, ok := d.GetOk("methods"); ok || d.HasChange("methods") {
		t, err := expandIcapProfileMethods(d, v, "methods")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["methods"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandIcapProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ocr_only"); ok || d.HasChange("ocr_only") {
		t, err := expandIcapProfileOcrOnly(d, v, "ocr_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocr-only"] = t
		}
	}

	if v, ok := d.GetOk("preview"); ok || d.HasChange("preview") {
		t, err := expandIcapProfilePreview(d, v, "preview")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preview"] = t
		}
	}

	if v, ok := d.GetOk("preview_data_length"); ok || d.HasChange("preview_data_length") {
		t, err := expandIcapProfilePreviewDataLength(d, v, "preview_data_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preview-data-length"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandIcapProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("request"); ok || d.HasChange("request") {
		t, err := expandIcapProfileRequest(d, v, "request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request"] = t
		}
	}

	if v, ok := d.GetOk("request_failure"); ok || d.HasChange("request_failure") {
		t, err := expandIcapProfileRequestFailure(d, v, "request_failure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-failure"] = t
		}
	}

	if v, ok := d.GetOk("request_path"); ok || d.HasChange("request_path") {
		t, err := expandIcapProfileRequestPath(d, v, "request_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-path"] = t
		}
	}

	if v, ok := d.GetOk("request_server"); ok || d.HasChange("request_server") {
		t, err := expandIcapProfileRequestServer(d, v, "request_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-server"] = t
		}
	}

	if v, ok := d.GetOk("respmod_default_action"); ok || d.HasChange("respmod_default_action") {
		t, err := expandIcapProfileRespmodDefaultAction(d, v, "respmod_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["respmod-default-action"] = t
		}
	}

	if v, ok := d.GetOk("respmod_forward_rules"); ok || d.HasChange("respmod_forward_rules") {
		t, err := expandIcapProfileRespmodForwardRules(d, v, "respmod_forward_rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["respmod-forward-rules"] = t
		}
	}

	if v, ok := d.GetOk("response"); ok || d.HasChange("response") {
		t, err := expandIcapProfileResponse(d, v, "response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response"] = t
		}
	}

	if v, ok := d.GetOk("response_failure"); ok || d.HasChange("response_failure") {
		t, err := expandIcapProfileResponseFailure(d, v, "response_failure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-failure"] = t
		}
	}

	if v, ok := d.GetOk("response_path"); ok || d.HasChange("response_path") {
		t, err := expandIcapProfileResponsePath(d, v, "response_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-path"] = t
		}
	}

	if v, ok := d.GetOk("response_req_hdr"); ok || d.HasChange("response_req_hdr") {
		t, err := expandIcapProfileResponseReqHdr(d, v, "response_req_hdr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-req-hdr"] = t
		}
	}

	if v, ok := d.GetOk("response_server"); ok || d.HasChange("response_server") {
		t, err := expandIcapProfileResponseServer(d, v, "response_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-server"] = t
		}
	}

	if v, ok := d.GetOk("scan_progress_interval"); ok || d.HasChange("scan_progress_interval") {
		t, err := expandIcapProfileScanProgressInterval(d, v, "scan_progress_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-progress-interval"] = t
		}
	}

	if v, ok := d.GetOk("streaming_content_bypass"); ok || d.HasChange("streaming_content_bypass") {
		t, err := expandIcapProfileStreamingContentBypass(d, v, "streaming_content_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["streaming-content-bypass"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandIcapProfileTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("allow_204_response"); ok || d.HasChange("allow_204_response") {
		t, err := expandIcapProfileAllow204Response(d, v, "allow_204_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-204-response"] = t
		}
	}

	if v, ok := d.GetOk("scan_oversize_log"); ok || d.HasChange("scan_oversize_log") {
		t, err := expandIcapProfileScanOversizeLog(d, v, "scan_oversize_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-oversize-log"] = t
		}
	}

	if v, ok := d.GetOk("scan_size_limit"); ok || d.HasChange("scan_size_limit") {
		t, err := expandIcapProfileScanSizeLimit(d, v, "scan_size_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-size-limit"] = t
		}
	}

	return &obj, nil
}
