// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> WAF HTTP protocol restrictions.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfileConstraint() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileConstraintUpdate,
		Read:   resourceWafProfileConstraintRead,
		Update: resourceWafProfileConstraintUpdate,
		Delete: resourceWafProfileConstraintDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"content_length": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"exception": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
				},
			},
			"header_length": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"line_length": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"malformed": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"max_cookie": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_cookie": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"max_header_line": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_header_line": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"max_range_segment": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_range_segment": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"max_url_param": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_url_param": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"method": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"param_length": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"url_param_length": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"version": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWafProfileConstraintUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWafProfileConstraint(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileConstraint resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfileConstraint(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileConstraint resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WafProfileConstraint")

	return resourceWafProfileConstraintRead(d, m)
}

func resourceWafProfileConstraintDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWafProfileConstraint(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfileConstraint resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileConstraintRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWafProfileConstraint(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfileConstraint resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfileConstraint(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfileConstraint resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileConstraintContentLength2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintContentLengthAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintContentLengthLength2edl(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintContentLengthLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintContentLengthSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintContentLengthStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintContentLengthAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLengthLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLengthLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLengthSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLengthStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintException2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenWafProfileConstraintExceptionAddress2edl(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_length"
		if _, ok := i["content-length"]; ok {
			v := flattenWafProfileConstraintExceptionContentLength2edl(i["content-length"], d, pre_append)
			tmp["content_length"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-ContentLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_length"
		if _, ok := i["header-length"]; ok {
			v := flattenWafProfileConstraintExceptionHeaderLength2edl(i["header-length"], d, pre_append)
			tmp["header_length"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-HeaderLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := i["hostname"]; ok {
			v := flattenWafProfileConstraintExceptionHostname2edl(i["hostname"], d, pre_append)
			tmp["hostname"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Hostname")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWafProfileConstraintExceptionId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "line_length"
		if _, ok := i["line-length"]; ok {
			v := flattenWafProfileConstraintExceptionLineLength2edl(i["line-length"], d, pre_append)
			tmp["line_length"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-LineLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "malformed"
		if _, ok := i["malformed"]; ok {
			v := flattenWafProfileConstraintExceptionMalformed2edl(i["malformed"], d, pre_append)
			tmp["malformed"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Malformed")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_cookie"
		if _, ok := i["max-cookie"]; ok {
			v := flattenWafProfileConstraintExceptionMaxCookie2edl(i["max-cookie"], d, pre_append)
			tmp["max_cookie"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-MaxCookie")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_header_line"
		if _, ok := i["max-header-line"]; ok {
			v := flattenWafProfileConstraintExceptionMaxHeaderLine2edl(i["max-header-line"], d, pre_append)
			tmp["max_header_line"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-MaxHeaderLine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_range_segment"
		if _, ok := i["max-range-segment"]; ok {
			v := flattenWafProfileConstraintExceptionMaxRangeSegment2edl(i["max-range-segment"], d, pre_append)
			tmp["max_range_segment"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-MaxRangeSegment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_url_param"
		if _, ok := i["max-url-param"]; ok {
			v := flattenWafProfileConstraintExceptionMaxUrlParam2edl(i["max-url-param"], d, pre_append)
			tmp["max_url_param"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-MaxUrlParam")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := i["method"]; ok {
			v := flattenWafProfileConstraintExceptionMethod2edl(i["method"], d, pre_append)
			tmp["method"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Method")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "param_length"
		if _, ok := i["param-length"]; ok {
			v := flattenWafProfileConstraintExceptionParamLength2edl(i["param-length"], d, pre_append)
			tmp["param_length"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-ParamLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWafProfileConstraintExceptionPattern2edl(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenWafProfileConstraintExceptionRegex2edl(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Regex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_param_length"
		if _, ok := i["url-param-length"]; ok {
			v := flattenWafProfileConstraintExceptionUrlParamLength2edl(i["url-param-length"], d, pre_append)
			tmp["url_param_length"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-UrlParamLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenWafProfileConstraintExceptionVersion2edl(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Version")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWafProfileConstraintExceptionAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileConstraintExceptionContentLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionHeaderLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionHostname2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionLineLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMalformed2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxCookie2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxHeaderLine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxRangeSegment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxUrlParam2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionParamLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionRegex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionUrlParamLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLength2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintHeaderLengthAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintHeaderLengthLength2edl(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintHeaderLengthLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintHeaderLengthSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintHeaderLengthStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintHeaderLengthAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHostname2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintHostnameAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintHostnameLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintHostnameSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintHostnameStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintHostnameAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHostnameLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHostnameSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHostnameStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLength2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintLineLengthAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintLineLengthLength2edl(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintLineLengthLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintLineLengthSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintLineLengthStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintLineLengthAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLengthLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLengthLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLengthSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLengthStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMalformed2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMalformedAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMalformedLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMalformedSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMalformedStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMalformedAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMalformedLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMalformedSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMalformedStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookie2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMaxCookieAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMaxCookieLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_cookie"
	if _, ok := i["max-cookie"]; ok {
		result["max_cookie"] = flattenWafProfileConstraintMaxCookieMaxCookie2edl(i["max-cookie"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMaxCookieSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMaxCookieStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMaxCookieAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookieLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookieMaxCookie2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookieSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookieStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLine2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMaxHeaderLineAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMaxHeaderLineLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_header_line"
	if _, ok := i["max-header-line"]; ok {
		result["max_header_line"] = flattenWafProfileConstraintMaxHeaderLineMaxHeaderLine2edl(i["max-header-line"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMaxHeaderLineSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMaxHeaderLineStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMaxHeaderLineAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLineLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLineMaxHeaderLine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLineSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLineStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegment2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMaxRangeSegmentAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMaxRangeSegmentLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := i["max-range-segment"]; ok {
		result["max_range_segment"] = flattenWafProfileConstraintMaxRangeSegmentMaxRangeSegment2edl(i["max-range-segment"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMaxRangeSegmentSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMaxRangeSegmentStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMaxRangeSegmentAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegmentLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegmentMaxRangeSegment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegmentSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegmentStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParam2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMaxUrlParamAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMaxUrlParamLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_url_param"
	if _, ok := i["max-url-param"]; ok {
		result["max_url_param"] = flattenWafProfileConstraintMaxUrlParamMaxUrlParam2edl(i["max-url-param"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMaxUrlParamSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMaxUrlParamStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMaxUrlParamAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamMaxUrlParam2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMethod2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMethodAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMethodLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMethodSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMethodStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMethodAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMethodLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMethodSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMethodStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLength2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintParamLengthAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintParamLengthLength2edl(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintParamLengthLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintParamLengthSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintParamLengthStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintParamLengthAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLengthLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLengthLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLengthSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLengthStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLength2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintUrlParamLengthAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintUrlParamLengthLength2edl(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintUrlParamLengthLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintUrlParamLengthSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintUrlParamLengthStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintUrlParamLengthAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLengthLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLengthLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLengthSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLengthStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintVersion2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintVersionAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintVersionLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintVersionSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintVersionStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintVersionAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintVersionLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintVersionSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintVersionStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWafProfileConstraint(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("content_length", flattenWafProfileConstraintContentLength2edl(o["content-length"], d, "content_length")); err != nil {
			if vv, ok := fortiAPIPatch(o["content-length"], "WafProfileConstraint-ContentLength"); ok {
				if err = d.Set("content_length", vv); err != nil {
					return fmt.Errorf("Error reading content_length: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading content_length: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("content_length"); ok {
			if err = d.Set("content_length", flattenWafProfileConstraintContentLength2edl(o["content-length"], d, "content_length")); err != nil {
				if vv, ok := fortiAPIPatch(o["content-length"], "WafProfileConstraint-ContentLength"); ok {
					if err = d.Set("content_length", vv); err != nil {
						return fmt.Errorf("Error reading content_length: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading content_length: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("exception", flattenWafProfileConstraintException2edl(o["exception"], d, "exception")); err != nil {
			if vv, ok := fortiAPIPatch(o["exception"], "WafProfileConstraint-Exception"); ok {
				if err = d.Set("exception", vv); err != nil {
					return fmt.Errorf("Error reading exception: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading exception: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("exception"); ok {
			if err = d.Set("exception", flattenWafProfileConstraintException2edl(o["exception"], d, "exception")); err != nil {
				if vv, ok := fortiAPIPatch(o["exception"], "WafProfileConstraint-Exception"); ok {
					if err = d.Set("exception", vv); err != nil {
						return fmt.Errorf("Error reading exception: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading exception: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("header_length", flattenWafProfileConstraintHeaderLength2edl(o["header-length"], d, "header_length")); err != nil {
			if vv, ok := fortiAPIPatch(o["header-length"], "WafProfileConstraint-HeaderLength"); ok {
				if err = d.Set("header_length", vv); err != nil {
					return fmt.Errorf("Error reading header_length: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading header_length: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("header_length"); ok {
			if err = d.Set("header_length", flattenWafProfileConstraintHeaderLength2edl(o["header-length"], d, "header_length")); err != nil {
				if vv, ok := fortiAPIPatch(o["header-length"], "WafProfileConstraint-HeaderLength"); ok {
					if err = d.Set("header_length", vv); err != nil {
						return fmt.Errorf("Error reading header_length: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading header_length: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("hostname", flattenWafProfileConstraintHostname2edl(o["hostname"], d, "hostname")); err != nil {
			if vv, ok := fortiAPIPatch(o["hostname"], "WafProfileConstraint-Hostname"); ok {
				if err = d.Set("hostname", vv); err != nil {
					return fmt.Errorf("Error reading hostname: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hostname"); ok {
			if err = d.Set("hostname", flattenWafProfileConstraintHostname2edl(o["hostname"], d, "hostname")); err != nil {
				if vv, ok := fortiAPIPatch(o["hostname"], "WafProfileConstraint-Hostname"); ok {
					if err = d.Set("hostname", vv); err != nil {
						return fmt.Errorf("Error reading hostname: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading hostname: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("line_length", flattenWafProfileConstraintLineLength2edl(o["line-length"], d, "line_length")); err != nil {
			if vv, ok := fortiAPIPatch(o["line-length"], "WafProfileConstraint-LineLength"); ok {
				if err = d.Set("line_length", vv); err != nil {
					return fmt.Errorf("Error reading line_length: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading line_length: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("line_length"); ok {
			if err = d.Set("line_length", flattenWafProfileConstraintLineLength2edl(o["line-length"], d, "line_length")); err != nil {
				if vv, ok := fortiAPIPatch(o["line-length"], "WafProfileConstraint-LineLength"); ok {
					if err = d.Set("line_length", vv); err != nil {
						return fmt.Errorf("Error reading line_length: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading line_length: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("malformed", flattenWafProfileConstraintMalformed2edl(o["malformed"], d, "malformed")); err != nil {
			if vv, ok := fortiAPIPatch(o["malformed"], "WafProfileConstraint-Malformed"); ok {
				if err = d.Set("malformed", vv); err != nil {
					return fmt.Errorf("Error reading malformed: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading malformed: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("malformed"); ok {
			if err = d.Set("malformed", flattenWafProfileConstraintMalformed2edl(o["malformed"], d, "malformed")); err != nil {
				if vv, ok := fortiAPIPatch(o["malformed"], "WafProfileConstraint-Malformed"); ok {
					if err = d.Set("malformed", vv); err != nil {
						return fmt.Errorf("Error reading malformed: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading malformed: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("max_cookie", flattenWafProfileConstraintMaxCookie2edl(o["max-cookie"], d, "max_cookie")); err != nil {
			if vv, ok := fortiAPIPatch(o["max-cookie"], "WafProfileConstraint-MaxCookie"); ok {
				if err = d.Set("max_cookie", vv); err != nil {
					return fmt.Errorf("Error reading max_cookie: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading max_cookie: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("max_cookie"); ok {
			if err = d.Set("max_cookie", flattenWafProfileConstraintMaxCookie2edl(o["max-cookie"], d, "max_cookie")); err != nil {
				if vv, ok := fortiAPIPatch(o["max-cookie"], "WafProfileConstraint-MaxCookie"); ok {
					if err = d.Set("max_cookie", vv); err != nil {
						return fmt.Errorf("Error reading max_cookie: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading max_cookie: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("max_header_line", flattenWafProfileConstraintMaxHeaderLine2edl(o["max-header-line"], d, "max_header_line")); err != nil {
			if vv, ok := fortiAPIPatch(o["max-header-line"], "WafProfileConstraint-MaxHeaderLine"); ok {
				if err = d.Set("max_header_line", vv); err != nil {
					return fmt.Errorf("Error reading max_header_line: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading max_header_line: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("max_header_line"); ok {
			if err = d.Set("max_header_line", flattenWafProfileConstraintMaxHeaderLine2edl(o["max-header-line"], d, "max_header_line")); err != nil {
				if vv, ok := fortiAPIPatch(o["max-header-line"], "WafProfileConstraint-MaxHeaderLine"); ok {
					if err = d.Set("max_header_line", vv); err != nil {
						return fmt.Errorf("Error reading max_header_line: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading max_header_line: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("max_range_segment", flattenWafProfileConstraintMaxRangeSegment2edl(o["max-range-segment"], d, "max_range_segment")); err != nil {
			if vv, ok := fortiAPIPatch(o["max-range-segment"], "WafProfileConstraint-MaxRangeSegment"); ok {
				if err = d.Set("max_range_segment", vv); err != nil {
					return fmt.Errorf("Error reading max_range_segment: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading max_range_segment: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("max_range_segment"); ok {
			if err = d.Set("max_range_segment", flattenWafProfileConstraintMaxRangeSegment2edl(o["max-range-segment"], d, "max_range_segment")); err != nil {
				if vv, ok := fortiAPIPatch(o["max-range-segment"], "WafProfileConstraint-MaxRangeSegment"); ok {
					if err = d.Set("max_range_segment", vv); err != nil {
						return fmt.Errorf("Error reading max_range_segment: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading max_range_segment: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("max_url_param", flattenWafProfileConstraintMaxUrlParam2edl(o["max-url-param"], d, "max_url_param")); err != nil {
			if vv, ok := fortiAPIPatch(o["max-url-param"], "WafProfileConstraint-MaxUrlParam"); ok {
				if err = d.Set("max_url_param", vv); err != nil {
					return fmt.Errorf("Error reading max_url_param: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading max_url_param: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("max_url_param"); ok {
			if err = d.Set("max_url_param", flattenWafProfileConstraintMaxUrlParam2edl(o["max-url-param"], d, "max_url_param")); err != nil {
				if vv, ok := fortiAPIPatch(o["max-url-param"], "WafProfileConstraint-MaxUrlParam"); ok {
					if err = d.Set("max_url_param", vv); err != nil {
						return fmt.Errorf("Error reading max_url_param: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading max_url_param: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("method", flattenWafProfileConstraintMethod2edl(o["method"], d, "method")); err != nil {
			if vv, ok := fortiAPIPatch(o["method"], "WafProfileConstraint-Method"); ok {
				if err = d.Set("method", vv); err != nil {
					return fmt.Errorf("Error reading method: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading method: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("method"); ok {
			if err = d.Set("method", flattenWafProfileConstraintMethod2edl(o["method"], d, "method")); err != nil {
				if vv, ok := fortiAPIPatch(o["method"], "WafProfileConstraint-Method"); ok {
					if err = d.Set("method", vv); err != nil {
						return fmt.Errorf("Error reading method: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading method: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("param_length", flattenWafProfileConstraintParamLength2edl(o["param-length"], d, "param_length")); err != nil {
			if vv, ok := fortiAPIPatch(o["param-length"], "WafProfileConstraint-ParamLength"); ok {
				if err = d.Set("param_length", vv); err != nil {
					return fmt.Errorf("Error reading param_length: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading param_length: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("param_length"); ok {
			if err = d.Set("param_length", flattenWafProfileConstraintParamLength2edl(o["param-length"], d, "param_length")); err != nil {
				if vv, ok := fortiAPIPatch(o["param-length"], "WafProfileConstraint-ParamLength"); ok {
					if err = d.Set("param_length", vv); err != nil {
						return fmt.Errorf("Error reading param_length: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading param_length: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("url_param_length", flattenWafProfileConstraintUrlParamLength2edl(o["url-param-length"], d, "url_param_length")); err != nil {
			if vv, ok := fortiAPIPatch(o["url-param-length"], "WafProfileConstraint-UrlParamLength"); ok {
				if err = d.Set("url_param_length", vv); err != nil {
					return fmt.Errorf("Error reading url_param_length: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading url_param_length: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("url_param_length"); ok {
			if err = d.Set("url_param_length", flattenWafProfileConstraintUrlParamLength2edl(o["url-param-length"], d, "url_param_length")); err != nil {
				if vv, ok := fortiAPIPatch(o["url-param-length"], "WafProfileConstraint-UrlParamLength"); ok {
					if err = d.Set("url_param_length", vv); err != nil {
						return fmt.Errorf("Error reading url_param_length: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading url_param_length: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("version", flattenWafProfileConstraintVersion2edl(o["version"], d, "version")); err != nil {
			if vv, ok := fortiAPIPatch(o["version"], "WafProfileConstraint-Version"); ok {
				if err = d.Set("version", vv); err != nil {
					return fmt.Errorf("Error reading version: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading version: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("version"); ok {
			if err = d.Set("version", flattenWafProfileConstraintVersion2edl(o["version"], d, "version")); err != nil {
				if vv, ok := fortiAPIPatch(o["version"], "WafProfileConstraint-Version"); ok {
					if err = d.Set("version", vv); err != nil {
						return fmt.Errorf("Error reading version: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading version: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWafProfileConstraintFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileConstraintContentLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintContentLengthAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandWafProfileConstraintContentLengthLength2edl(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintContentLengthLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintContentLengthSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintContentLengthStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintContentLengthAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLengthLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLengthLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLengthSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLengthStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintException2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandWafProfileConstraintExceptionAddress2edl(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["content-length"], _ = expandWafProfileConstraintExceptionContentLength2edl(d, i["content_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-length"], _ = expandWafProfileConstraintExceptionHeaderLength2edl(d, i["header_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hostname"], _ = expandWafProfileConstraintExceptionHostname2edl(d, i["hostname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWafProfileConstraintExceptionId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "line_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["line-length"], _ = expandWafProfileConstraintExceptionLineLength2edl(d, i["line_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "malformed"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["malformed"], _ = expandWafProfileConstraintExceptionMalformed2edl(d, i["malformed"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_cookie"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-cookie"], _ = expandWafProfileConstraintExceptionMaxCookie2edl(d, i["max_cookie"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_header_line"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-header-line"], _ = expandWafProfileConstraintExceptionMaxHeaderLine2edl(d, i["max_header_line"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_range_segment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-range-segment"], _ = expandWafProfileConstraintExceptionMaxRangeSegment2edl(d, i["max_range_segment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_url_param"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-url-param"], _ = expandWafProfileConstraintExceptionMaxUrlParam2edl(d, i["max_url_param"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["method"], _ = expandWafProfileConstraintExceptionMethod2edl(d, i["method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "param_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["param-length"], _ = expandWafProfileConstraintExceptionParamLength2edl(d, i["param_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWafProfileConstraintExceptionPattern2edl(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regex"], _ = expandWafProfileConstraintExceptionRegex2edl(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_param_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-param-length"], _ = expandWafProfileConstraintExceptionUrlParamLength2edl(d, i["url_param_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["version"], _ = expandWafProfileConstraintExceptionVersion2edl(d, i["version"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWafProfileConstraintExceptionAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileConstraintExceptionContentLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionHeaderLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionHostname2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionLineLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMalformed2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxCookie2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxHeaderLine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxRangeSegment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxUrlParam2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionParamLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionRegex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionUrlParamLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintHeaderLengthAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandWafProfileConstraintHeaderLengthLength2edl(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintHeaderLengthLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintHeaderLengthSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintHeaderLengthStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintHeaderLengthAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHostname2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintHostnameAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintHostnameLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintHostnameSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintHostnameStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintHostnameAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHostnameLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHostnameSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHostnameStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintLineLengthAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandWafProfileConstraintLineLengthLength2edl(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintLineLengthLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintLineLengthSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintLineLengthStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintLineLengthAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLengthLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLengthLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLengthSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLengthStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMalformed2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintMalformedAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintMalformedLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintMalformedSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintMalformedStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMalformedAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMalformedLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMalformedSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMalformedStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookie2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintMaxCookieAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintMaxCookieLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_cookie"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-cookie"], _ = expandWafProfileConstraintMaxCookieMaxCookie2edl(d, i["max_cookie"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintMaxCookieSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintMaxCookieStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMaxCookieAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookieLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookieMaxCookie2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookieSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookieStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintMaxHeaderLineAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintMaxHeaderLineLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_header_line"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-header-line"], _ = expandWafProfileConstraintMaxHeaderLineMaxHeaderLine2edl(d, i["max_header_line"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintMaxHeaderLineSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintMaxHeaderLineStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMaxHeaderLineAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLineLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLineMaxHeaderLine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLineSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLineStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintMaxRangeSegmentAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintMaxRangeSegmentLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-range-segment"], _ = expandWafProfileConstraintMaxRangeSegmentMaxRangeSegment2edl(d, i["max_range_segment"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintMaxRangeSegmentSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintMaxRangeSegmentStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMaxRangeSegmentAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegmentLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegmentMaxRangeSegment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegmentSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegmentStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParam2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintMaxUrlParamAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintMaxUrlParamLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_url_param"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-url-param"], _ = expandWafProfileConstraintMaxUrlParamMaxUrlParam2edl(d, i["max_url_param"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintMaxUrlParamSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintMaxUrlParamStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMaxUrlParamAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamMaxUrlParam2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintMethodAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintMethodLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintMethodSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintMethodStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMethodAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMethodLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMethodSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMethodStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintParamLengthAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandWafProfileConstraintParamLengthLength2edl(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintParamLengthLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintParamLengthSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintParamLengthStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintParamLengthAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLengthLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLengthLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLengthSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLengthStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintUrlParamLengthAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandWafProfileConstraintUrlParamLengthLength2edl(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintUrlParamLengthLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintUrlParamLengthSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintUrlParamLengthStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintUrlParamLengthAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLengthLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLengthLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLengthSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLengthStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintVersionAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintVersionLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintVersionSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintVersionStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintVersionAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintVersionLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintVersionSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintVersionStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWafProfileConstraint(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("content_length"); ok || d.HasChange("content_length") {
		t, err := expandWafProfileConstraintContentLength2edl(d, v, "content_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-length"] = t
		}
	}

	if v, ok := d.GetOk("exception"); ok || d.HasChange("exception") {
		t, err := expandWafProfileConstraintException2edl(d, v, "exception")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception"] = t
		}
	}

	if v, ok := d.GetOk("header_length"); ok || d.HasChange("header_length") {
		t, err := expandWafProfileConstraintHeaderLength2edl(d, v, "header_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-length"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandWafProfileConstraintHostname2edl(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("line_length"); ok || d.HasChange("line_length") {
		t, err := expandWafProfileConstraintLineLength2edl(d, v, "line_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["line-length"] = t
		}
	}

	if v, ok := d.GetOk("malformed"); ok || d.HasChange("malformed") {
		t, err := expandWafProfileConstraintMalformed2edl(d, v, "malformed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed"] = t
		}
	}

	if v, ok := d.GetOk("max_cookie"); ok || d.HasChange("max_cookie") {
		t, err := expandWafProfileConstraintMaxCookie2edl(d, v, "max_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cookie"] = t
		}
	}

	if v, ok := d.GetOk("max_header_line"); ok || d.HasChange("max_header_line") {
		t, err := expandWafProfileConstraintMaxHeaderLine2edl(d, v, "max_header_line")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-header-line"] = t
		}
	}

	if v, ok := d.GetOk("max_range_segment"); ok || d.HasChange("max_range_segment") {
		t, err := expandWafProfileConstraintMaxRangeSegment2edl(d, v, "max_range_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-range-segment"] = t
		}
	}

	if v, ok := d.GetOk("max_url_param"); ok || d.HasChange("max_url_param") {
		t, err := expandWafProfileConstraintMaxUrlParam2edl(d, v, "max_url_param")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandWafProfileConstraintMethod2edl(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("param_length"); ok || d.HasChange("param_length") {
		t, err := expandWafProfileConstraintParamLength2edl(d, v, "param_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["param-length"] = t
		}
	}

	if v, ok := d.GetOk("url_param_length"); ok || d.HasChange("url_param_length") {
		t, err := expandWafProfileConstraintUrlParamLength2edl(d, v, "url_param_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-length"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandWafProfileConstraintVersion2edl(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	return &obj, nil
}
