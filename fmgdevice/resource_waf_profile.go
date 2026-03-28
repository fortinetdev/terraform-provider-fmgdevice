// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Web application firewall configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileCreate,
		Read:   resourceWafProfileRead,
		Update: resourceWafProfileUpdate,
		Delete: resourceWafProfileDelete,

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
			"address_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"blocked_address": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"blocked_log": &schema.Schema{
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
						"trusted_address": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"constraint": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_allowed_methods": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"method_policy": &schema.Schema{
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
									"allowed_methods": &schema.Schema{
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
									"pattern": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"regex": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"signature": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credit_card_detection_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"custom_signature": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"case_sensitivity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"pattern": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									"target": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"disabled_signature": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"disabled_sub_class": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"main_class": &schema.Schema{
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
									"id": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
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
					},
				},
			},
			"url_access": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_pattern": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"negate": &schema.Schema{
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
									"srcaddr": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"address": &schema.Schema{
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

func resourceWafProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWafProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WafProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWafProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWafProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WafProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWafProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WafProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWafProfileRead(d, m)
}

func resourceWafProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWafProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWafProfileRead(d, m)
}

func resourceWafProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWafProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWafProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfile resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileAddressList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "blocked_address"
	if _, ok := i["blocked-address"]; ok {
		result["blocked_address"] = flattenWafProfileAddressListBlockedAddress(i["blocked-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "blocked_log"
	if _, ok := i["blocked-log"]; ok {
		result["blocked_log"] = flattenWafProfileAddressListBlockedLog(i["blocked-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileAddressListSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileAddressListStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "trusted_address"
	if _, ok := i["trusted-address"]; ok {
		result["trusted_address"] = flattenWafProfileAddressListTrustedAddress(i["trusted-address"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileAddressListBlockedAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileAddressListBlockedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileAddressListSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileAddressListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileAddressListTrustedAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraint(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "content_length"
	if _, ok := i["content-length"]; ok {
		result["content_length"] = flattenWafProfileConstraintContentLength(i["content-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "exception"
	if _, ok := i["exception"]; ok {
		result["exception"] = flattenWafProfileConstraintException(i["exception"], d, pre_append)
	}

	pre_append = pre + ".0." + "header_length"
	if _, ok := i["header-length"]; ok {
		result["header_length"] = flattenWafProfileConstraintHeaderLength(i["header-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "hostname"
	if _, ok := i["hostname"]; ok {
		result["hostname"] = flattenWafProfileConstraintHostname(i["hostname"], d, pre_append)
	}

	pre_append = pre + ".0." + "line_length"
	if _, ok := i["line-length"]; ok {
		result["line_length"] = flattenWafProfileConstraintLineLength(i["line-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed"
	if _, ok := i["malformed"]; ok {
		result["malformed"] = flattenWafProfileConstraintMalformed(i["malformed"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_cookie"
	if _, ok := i["max-cookie"]; ok {
		result["max_cookie"] = flattenWafProfileConstraintMaxCookie(i["max-cookie"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_header_line"
	if _, ok := i["max-header-line"]; ok {
		result["max_header_line"] = flattenWafProfileConstraintMaxHeaderLine(i["max-header-line"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := i["max-range-segment"]; ok {
		result["max_range_segment"] = flattenWafProfileConstraintMaxRangeSegment(i["max-range-segment"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_url_param"
	if _, ok := i["max-url-param"]; ok {
		result["max_url_param"] = flattenWafProfileConstraintMaxUrlParam(i["max-url-param"], d, pre_append)
	}

	pre_append = pre + ".0." + "method"
	if _, ok := i["method"]; ok {
		result["method"] = flattenWafProfileConstraintMethod(i["method"], d, pre_append)
	}

	pre_append = pre + ".0." + "param_length"
	if _, ok := i["param-length"]; ok {
		result["param_length"] = flattenWafProfileConstraintParamLength(i["param-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "url_param_length"
	if _, ok := i["url-param-length"]; ok {
		result["url_param_length"] = flattenWafProfileConstraintUrlParamLength(i["url-param-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "version"
	if _, ok := i["version"]; ok {
		result["version"] = flattenWafProfileConstraintVersion(i["version"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintContentLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintContentLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintContentLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintContentLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintContentLengthSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintContentLengthStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintContentLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintException(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWafProfileConstraintExceptionAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_length"
		if _, ok := i["content-length"]; ok {
			v := flattenWafProfileConstraintExceptionContentLength(i["content-length"], d, pre_append)
			tmp["content_length"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-ContentLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_length"
		if _, ok := i["header-length"]; ok {
			v := flattenWafProfileConstraintExceptionHeaderLength(i["header-length"], d, pre_append)
			tmp["header_length"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-HeaderLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := i["hostname"]; ok {
			v := flattenWafProfileConstraintExceptionHostname(i["hostname"], d, pre_append)
			tmp["hostname"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Hostname")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWafProfileConstraintExceptionId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "line_length"
		if _, ok := i["line-length"]; ok {
			v := flattenWafProfileConstraintExceptionLineLength(i["line-length"], d, pre_append)
			tmp["line_length"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-LineLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "malformed"
		if _, ok := i["malformed"]; ok {
			v := flattenWafProfileConstraintExceptionMalformed(i["malformed"], d, pre_append)
			tmp["malformed"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Malformed")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_cookie"
		if _, ok := i["max-cookie"]; ok {
			v := flattenWafProfileConstraintExceptionMaxCookie(i["max-cookie"], d, pre_append)
			tmp["max_cookie"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-MaxCookie")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_header_line"
		if _, ok := i["max-header-line"]; ok {
			v := flattenWafProfileConstraintExceptionMaxHeaderLine(i["max-header-line"], d, pre_append)
			tmp["max_header_line"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-MaxHeaderLine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_range_segment"
		if _, ok := i["max-range-segment"]; ok {
			v := flattenWafProfileConstraintExceptionMaxRangeSegment(i["max-range-segment"], d, pre_append)
			tmp["max_range_segment"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-MaxRangeSegment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_url_param"
		if _, ok := i["max-url-param"]; ok {
			v := flattenWafProfileConstraintExceptionMaxUrlParam(i["max-url-param"], d, pre_append)
			tmp["max_url_param"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-MaxUrlParam")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := i["method"]; ok {
			v := flattenWafProfileConstraintExceptionMethod(i["method"], d, pre_append)
			tmp["method"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Method")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "param_length"
		if _, ok := i["param-length"]; ok {
			v := flattenWafProfileConstraintExceptionParamLength(i["param-length"], d, pre_append)
			tmp["param_length"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-ParamLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWafProfileConstraintExceptionPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenWafProfileConstraintExceptionRegex(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Regex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_param_length"
		if _, ok := i["url-param-length"]; ok {
			v := flattenWafProfileConstraintExceptionUrlParamLength(i["url-param-length"], d, pre_append)
			tmp["url_param_length"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-UrlParamLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			v := flattenWafProfileConstraintExceptionVersion(i["version"], d, pre_append)
			tmp["version"] = fortiAPISubPartPatch(v, "WafProfileConstraint-Exception-Version")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWafProfileConstraintExceptionAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileConstraintExceptionContentLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionHeaderLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionLineLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMalformed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxHeaderLine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxRangeSegment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxUrlParam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionParamLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionUrlParamLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintHeaderLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintHeaderLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintHeaderLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintHeaderLengthSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintHeaderLengthStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintHeaderLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHostname(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintHostnameAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintHostnameLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintHostnameSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintHostnameStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintHostnameAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHostnameLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHostnameSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHostnameStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintLineLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintLineLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintLineLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintLineLengthSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintLineLengthStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintLineLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMalformed(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMalformedAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMalformedLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMalformedSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMalformedStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMalformedAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMalformedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMalformedSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMalformedStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookie(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMaxCookieAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMaxCookieLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_cookie"
	if _, ok := i["max-cookie"]; ok {
		result["max_cookie"] = flattenWafProfileConstraintMaxCookieMaxCookie(i["max-cookie"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMaxCookieSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMaxCookieStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMaxCookieAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookieLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookieMaxCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookieSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookieStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLine(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMaxHeaderLineAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMaxHeaderLineLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_header_line"
	if _, ok := i["max-header-line"]; ok {
		result["max_header_line"] = flattenWafProfileConstraintMaxHeaderLineMaxHeaderLine(i["max-header-line"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMaxHeaderLineSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMaxHeaderLineStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMaxHeaderLineAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLineMaxHeaderLine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLineSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLineStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegment(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMaxRangeSegmentAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMaxRangeSegmentLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := i["max-range-segment"]; ok {
		result["max_range_segment"] = flattenWafProfileConstraintMaxRangeSegmentMaxRangeSegment(i["max-range-segment"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMaxRangeSegmentSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMaxRangeSegmentStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMaxRangeSegmentAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegmentLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegmentMaxRangeSegment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegmentSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegmentStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParam(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMaxUrlParamAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMaxUrlParamLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_url_param"
	if _, ok := i["max-url-param"]; ok {
		result["max_url_param"] = flattenWafProfileConstraintMaxUrlParamMaxUrlParam(i["max-url-param"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMaxUrlParamSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMaxUrlParamStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMaxUrlParamAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamMaxUrlParam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMethod(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMethodAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMethodLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMethodSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMethodStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMethodAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMethodLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMethodSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMethodStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintParamLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintParamLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintParamLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintParamLengthSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintParamLengthStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintParamLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintUrlParamLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintUrlParamLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintUrlParamLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintUrlParamLengthSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintUrlParamLengthStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintUrlParamLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintVersion(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintVersionAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintVersionLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintVersionSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintVersionStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintVersionAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintVersionLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintVersionSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintVersionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethod(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "default_allowed_methods"
	if _, ok := i["default-allowed-methods"]; ok {
		result["default_allowed_methods"] = flattenWafProfileMethodDefaultAllowedMethods(i["default-allowed-methods"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileMethodLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "method_policy"
	if _, ok := i["method-policy"]; ok {
		result["method_policy"] = flattenWafProfileMethodMethodPolicy(i["method-policy"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileMethodSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileMethodStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileMethodDefaultAllowedMethods(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileMethodLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWafProfileMethodMethodPolicyAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "WafProfileMethod-MethodPolicy-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_methods"
		if _, ok := i["allowed-methods"]; ok {
			v := flattenWafProfileMethodMethodPolicyAllowedMethods(i["allowed-methods"], d, pre_append)
			tmp["allowed_methods"] = fortiAPISubPartPatch(v, "WafProfileMethod-MethodPolicy-AllowedMethods")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWafProfileMethodMethodPolicyId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WafProfileMethod-MethodPolicy-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWafProfileMethodMethodPolicyPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WafProfileMethod-MethodPolicy-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenWafProfileMethodMethodPolicyRegex(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "WafProfileMethod-MethodPolicy-Regex")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWafProfileMethodMethodPolicyAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileMethodMethodPolicyAllowedMethods(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileMethodMethodPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicyPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicyRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignature(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "credit_card_detection_threshold"
	if _, ok := i["credit-card-detection-threshold"]; ok {
		result["credit_card_detection_threshold"] = flattenWafProfileSignatureCreditCardDetectionThreshold(i["credit-card-detection-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "custom_signature"
	if _, ok := i["custom-signature"]; ok {
		result["custom_signature"] = flattenWafProfileSignatureCustomSignature(i["custom-signature"], d, pre_append)
	}

	pre_append = pre + ".0." + "disabled_signature"
	if _, ok := i["disabled-signature"]; ok {
		result["disabled_signature"] = flattenWafProfileSignatureDisabledSignature(i["disabled-signature"], d, pre_append)
	}

	pre_append = pre + ".0." + "disabled_sub_class"
	if _, ok := i["disabled-sub-class"]; ok {
		result["disabled_sub_class"] = flattenWafProfileSignatureDisabledSubClass(i["disabled-sub-class"], d, pre_append)
	}

	pre_append = pre + ".0." + "main_class"
	if _, ok := i["main-class"]; ok {
		result["main_class"] = flattenWafProfileSignatureMainClass(i["main-class"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileSignatureCreditCardDetectionThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignature(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWafProfileSignatureCustomSignatureAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {
			v := flattenWafProfileSignatureCustomSignatureCaseSensitivity(i["case-sensitivity"], d, pre_append)
			tmp["case_sensitivity"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-CaseSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenWafProfileSignatureCustomSignatureDirection(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenWafProfileSignatureCustomSignatureLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWafProfileSignatureCustomSignatureName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWafProfileSignatureCustomSignaturePattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenWafProfileSignatureCustomSignatureSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWafProfileSignatureCustomSignatureStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenWafProfileSignatureCustomSignatureTarget(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Target")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWafProfileSignatureCustomSignatureAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignaturePattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileSignatureDisabledSignature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileSignatureDisabledSubClass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileSignatureMainClass(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileSignatureMainClassAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenWafProfileSignatureMainClassId(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileSignatureMainClassLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileSignatureMainClassSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileSignatureMainClassStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileSignatureMainClassAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileSignatureMainClassLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccess(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_pattern"
		if _, ok := i["access-pattern"]; ok {
			v := flattenWafProfileUrlAccessAccessPattern(i["access-pattern"], d, pre_append)
			tmp["access_pattern"] = fortiAPISubPartPatch(v, "WafProfile-UrlAccess-AccessPattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenWafProfileUrlAccessAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WafProfile-UrlAccess-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenWafProfileUrlAccessAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "WafProfile-UrlAccess-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWafProfileUrlAccessId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WafProfile-UrlAccess-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenWafProfileUrlAccessLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "WafProfile-UrlAccess-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenWafProfileUrlAccessSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "WafProfile-UrlAccess-Severity")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWafProfileUrlAccessAccessPattern(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWafProfileUrlAccessAccessPatternId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WafProfileUrlAccess-AccessPattern-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := i["negate"]; ok {
			v := flattenWafProfileUrlAccessAccessPatternNegate(i["negate"], d, pre_append)
			tmp["negate"] = fortiAPISubPartPatch(v, "WafProfileUrlAccess-AccessPattern-Negate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWafProfileUrlAccessAccessPatternPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WafProfileUrlAccess-AccessPattern-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenWafProfileUrlAccessAccessPatternRegex(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "WafProfileUrlAccess-AccessPattern-Regex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenWafProfileUrlAccessAccessPatternSrcaddr(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "WafProfileUrlAccess-AccessPattern-Srcaddr")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWafProfileUrlAccessAccessPatternId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileUrlAccessAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileUrlAccessId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWafProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("address_list", flattenWafProfileAddressList(o["address-list"], d, "address_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["address-list"], "WafProfile-AddressList"); ok {
				if err = d.Set("address_list", vv); err != nil {
					return fmt.Errorf("Error reading address_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading address_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("address_list"); ok {
			if err = d.Set("address_list", flattenWafProfileAddressList(o["address-list"], d, "address_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["address-list"], "WafProfile-AddressList"); ok {
					if err = d.Set("address_list", vv); err != nil {
						return fmt.Errorf("Error reading address_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading address_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenWafProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WafProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("constraint", flattenWafProfileConstraint(o["constraint"], d, "constraint")); err != nil {
			if vv, ok := fortiAPIPatch(o["constraint"], "WafProfile-Constraint"); ok {
				if err = d.Set("constraint", vv); err != nil {
					return fmt.Errorf("Error reading constraint: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading constraint: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("constraint"); ok {
			if err = d.Set("constraint", flattenWafProfileConstraint(o["constraint"], d, "constraint")); err != nil {
				if vv, ok := fortiAPIPatch(o["constraint"], "WafProfile-Constraint"); ok {
					if err = d.Set("constraint", vv); err != nil {
						return fmt.Errorf("Error reading constraint: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading constraint: %v", err)
				}
			}
		}
	}

	if err = d.Set("extended_log", flattenWafProfileExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "WafProfile-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("external", flattenWafProfileExternal(o["external"], d, "external")); err != nil {
		if vv, ok := fortiAPIPatch(o["external"], "WafProfile-External"); ok {
			if err = d.Set("external", vv); err != nil {
				return fmt.Errorf("Error reading external: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("method", flattenWafProfileMethod(o["method"], d, "method")); err != nil {
			if vv, ok := fortiAPIPatch(o["method"], "WafProfile-Method"); ok {
				if err = d.Set("method", vv); err != nil {
					return fmt.Errorf("Error reading method: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading method: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("method"); ok {
			if err = d.Set("method", flattenWafProfileMethod(o["method"], d, "method")); err != nil {
				if vv, ok := fortiAPIPatch(o["method"], "WafProfile-Method"); ok {
					if err = d.Set("method", vv); err != nil {
						return fmt.Errorf("Error reading method: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading method: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenWafProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WafProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("signature", flattenWafProfileSignature(o["signature"], d, "signature")); err != nil {
			if vv, ok := fortiAPIPatch(o["signature"], "WafProfile-Signature"); ok {
				if err = d.Set("signature", vv); err != nil {
					return fmt.Errorf("Error reading signature: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading signature: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("signature"); ok {
			if err = d.Set("signature", flattenWafProfileSignature(o["signature"], d, "signature")); err != nil {
				if vv, ok := fortiAPIPatch(o["signature"], "WafProfile-Signature"); ok {
					if err = d.Set("signature", vv); err != nil {
						return fmt.Errorf("Error reading signature: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading signature: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("url_access", flattenWafProfileUrlAccess(o["url-access"], d, "url_access")); err != nil {
			if vv, ok := fortiAPIPatch(o["url-access"], "WafProfile-UrlAccess"); ok {
				if err = d.Set("url_access", vv); err != nil {
					return fmt.Errorf("Error reading url_access: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading url_access: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("url_access"); ok {
			if err = d.Set("url_access", flattenWafProfileUrlAccess(o["url-access"], d, "url_access")); err != nil {
				if vv, ok := fortiAPIPatch(o["url-access"], "WafProfile-UrlAccess"); ok {
					if err = d.Set("url_access", vv); err != nil {
						return fmt.Errorf("Error reading url_access: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading url_access: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWafProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileAddressList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "blocked_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["blocked-address"], _ = expandWafProfileAddressListBlockedAddress(d, i["blocked_address"], pre_append)
	}
	pre_append = pre + ".0." + "blocked_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["blocked-log"], _ = expandWafProfileAddressListBlockedLog(d, i["blocked_log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileAddressListSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileAddressListStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "trusted_address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["trusted-address"], _ = expandWafProfileAddressListTrustedAddress(d, i["trusted_address"], pre_append)
	}

	return result, nil
}

func expandWafProfileAddressListBlockedAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileAddressListBlockedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileAddressListSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileAddressListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileAddressListTrustedAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "content_length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintContentLength(d, i["content_length"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["content-length"] = t
		}
	}
	pre_append = pre + ".0." + "exception"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintException(d, i["exception"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["exception"] = t
		}
	}
	pre_append = pre + ".0." + "header_length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintHeaderLength(d, i["header_length"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["header-length"] = t
		}
	}
	pre_append = pre + ".0." + "hostname"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintHostname(d, i["hostname"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["hostname"] = t
		}
	}
	pre_append = pre + ".0." + "line_length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintLineLength(d, i["line_length"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["line-length"] = t
		}
	}
	pre_append = pre + ".0." + "malformed"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintMalformed(d, i["malformed"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["malformed"] = t
		}
	}
	pre_append = pre + ".0." + "max_cookie"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintMaxCookie(d, i["max_cookie"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["max-cookie"] = t
		}
	}
	pre_append = pre + ".0." + "max_header_line"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintMaxHeaderLine(d, i["max_header_line"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["max-header-line"] = t
		}
	}
	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintMaxRangeSegment(d, i["max_range_segment"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["max-range-segment"] = t
		}
	}
	pre_append = pre + ".0." + "max_url_param"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintMaxUrlParam(d, i["max_url_param"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["max-url-param"] = t
		}
	}
	pre_append = pre + ".0." + "method"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintMethod(d, i["method"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["method"] = t
		}
	}
	pre_append = pre + ".0." + "param_length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintParamLength(d, i["param_length"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["param-length"] = t
		}
	}
	pre_append = pre + ".0." + "url_param_length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintUrlParamLength(d, i["url_param_length"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["url-param-length"] = t
		}
	}
	pre_append = pre + ".0." + "version"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileConstraintVersion(d, i["version"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["version"] = t
		}
	}

	return result, nil
}

func expandWafProfileConstraintContentLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintContentLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandWafProfileConstraintContentLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintContentLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintContentLengthSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintContentLengthStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintContentLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintException(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["address"], _ = expandWafProfileConstraintExceptionAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["content-length"], _ = expandWafProfileConstraintExceptionContentLength(d, i["content_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-length"], _ = expandWafProfileConstraintExceptionHeaderLength(d, i["header_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hostname"], _ = expandWafProfileConstraintExceptionHostname(d, i["hostname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWafProfileConstraintExceptionId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "line_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["line-length"], _ = expandWafProfileConstraintExceptionLineLength(d, i["line_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "malformed"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["malformed"], _ = expandWafProfileConstraintExceptionMalformed(d, i["malformed"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_cookie"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-cookie"], _ = expandWafProfileConstraintExceptionMaxCookie(d, i["max_cookie"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_header_line"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-header-line"], _ = expandWafProfileConstraintExceptionMaxHeaderLine(d, i["max_header_line"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_range_segment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-range-segment"], _ = expandWafProfileConstraintExceptionMaxRangeSegment(d, i["max_range_segment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_url_param"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-url-param"], _ = expandWafProfileConstraintExceptionMaxUrlParam(d, i["max_url_param"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["method"], _ = expandWafProfileConstraintExceptionMethod(d, i["method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "param_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["param-length"], _ = expandWafProfileConstraintExceptionParamLength(d, i["param_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWafProfileConstraintExceptionPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regex"], _ = expandWafProfileConstraintExceptionRegex(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_param_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url-param-length"], _ = expandWafProfileConstraintExceptionUrlParamLength(d, i["url_param_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["version"], _ = expandWafProfileConstraintExceptionVersion(d, i["version"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWafProfileConstraintExceptionAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileConstraintExceptionContentLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionHeaderLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionLineLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMalformed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxHeaderLine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxRangeSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxUrlParam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionParamLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionUrlParamLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintHeaderLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandWafProfileConstraintHeaderLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintHeaderLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintHeaderLengthSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintHeaderLengthStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintHeaderLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintHostnameAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintHostnameLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintHostnameSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintHostnameStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintHostnameAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHostnameLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHostnameSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHostnameStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintLineLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandWafProfileConstraintLineLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintLineLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintLineLengthSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintLineLengthStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintLineLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMalformed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintMalformedAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintMalformedLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintMalformedSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintMalformedStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMalformedAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMalformedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMalformedSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMalformedStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintMaxCookieAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintMaxCookieLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_cookie"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-cookie"], _ = expandWafProfileConstraintMaxCookieMaxCookie(d, i["max_cookie"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintMaxCookieSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintMaxCookieStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMaxCookieAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookieLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookieMaxCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookieSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookieStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintMaxHeaderLineAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintMaxHeaderLineLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_header_line"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-header-line"], _ = expandWafProfileConstraintMaxHeaderLineMaxHeaderLine(d, i["max_header_line"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintMaxHeaderLineSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintMaxHeaderLineStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMaxHeaderLineAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLineMaxHeaderLine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLineSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLineStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintMaxRangeSegmentAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintMaxRangeSegmentLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-range-segment"], _ = expandWafProfileConstraintMaxRangeSegmentMaxRangeSegment(d, i["max_range_segment"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintMaxRangeSegmentSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintMaxRangeSegmentStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMaxRangeSegmentAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegmentLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegmentMaxRangeSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegmentSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegmentStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintMaxUrlParamAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintMaxUrlParamLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "max_url_param"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-url-param"], _ = expandWafProfileConstraintMaxUrlParamMaxUrlParam(d, i["max_url_param"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintMaxUrlParamSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintMaxUrlParamStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMaxUrlParamAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamMaxUrlParam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintMethodAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintMethodLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintMethodSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintMethodStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMethodAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMethodLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMethodSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMethodStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintParamLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandWafProfileConstraintParamLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintParamLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintParamLengthSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintParamLengthStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintParamLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintUrlParamLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["length"], _ = expandWafProfileConstraintUrlParamLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintUrlParamLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintUrlParamLengthSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintUrlParamLengthStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintUrlParamLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileConstraintVersionAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileConstraintVersionLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileConstraintVersionSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileConstraintVersionStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintVersionAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintVersionLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintVersionSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintVersionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "default_allowed_methods"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-allowed-methods"], _ = expandWafProfileMethodDefaultAllowedMethods(d, i["default_allowed_methods"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileMethodLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "method_policy"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileMethodMethodPolicy(d, i["method_policy"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["method-policy"] = t
		}
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileMethodSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileMethodStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileMethodDefaultAllowedMethods(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileMethodLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["address"], _ = expandWafProfileMethodMethodPolicyAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_methods"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowed-methods"], _ = expandWafProfileMethodMethodPolicyAllowedMethods(d, i["allowed_methods"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWafProfileMethodMethodPolicyId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWafProfileMethodMethodPolicyPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regex"], _ = expandWafProfileMethodMethodPolicyRegex(d, i["regex"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWafProfileMethodMethodPolicyAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileMethodMethodPolicyAllowedMethods(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileMethodMethodPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicyPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicyRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "credit_card_detection_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["credit-card-detection-threshold"], _ = expandWafProfileSignatureCreditCardDetectionThreshold(d, i["credit_card_detection_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "custom_signature"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileSignatureCustomSignature(d, i["custom_signature"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["custom-signature"] = t
		}
	}
	pre_append = pre + ".0." + "disabled_signature"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disabled-signature"], _ = expandWafProfileSignatureDisabledSignature(d, i["disabled_signature"], pre_append)
	}
	pre_append = pre + ".0." + "disabled_sub_class"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disabled-sub-class"], _ = expandWafProfileSignatureDisabledSubClass(d, i["disabled_sub_class"], pre_append)
	}
	pre_append = pre + ".0." + "main_class"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWafProfileSignatureMainClass(d, i["main_class"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["main-class"] = t
		}
	}

	return result, nil
}

func expandWafProfileSignatureCreditCardDetectionThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandWafProfileSignatureCustomSignatureAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitivity"], _ = expandWafProfileSignatureCustomSignatureCaseSensitivity(d, i["case_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandWafProfileSignatureCustomSignatureDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandWafProfileSignatureCustomSignatureLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWafProfileSignatureCustomSignatureName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWafProfileSignatureCustomSignaturePattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandWafProfileSignatureCustomSignatureSeverity(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWafProfileSignatureCustomSignatureStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target"], _ = expandWafProfileSignatureCustomSignatureTarget(d, i["target"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWafProfileSignatureCustomSignatureAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignaturePattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileSignatureDisabledSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileSignatureDisabledSubClass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileSignatureMainClass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileSignatureMainClassAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandWafProfileSignatureMainClassId(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileSignatureMainClassLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileSignatureMainClassSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileSignatureMainClassStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileSignatureMainClassAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileSignatureMainClassLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWafProfileUrlAccessAccessPattern(d, i["access_pattern"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["access-pattern"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandWafProfileUrlAccessAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandWafProfileUrlAccessAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWafProfileUrlAccessId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandWafProfileUrlAccessLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandWafProfileUrlAccessSeverity(d, i["severity"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWafProfileUrlAccessAccessPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWafProfileUrlAccessAccessPatternId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["negate"], _ = expandWafProfileUrlAccessAccessPatternNegate(d, i["negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWafProfileUrlAccessAccessPatternPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regex"], _ = expandWafProfileUrlAccessAccessPatternRegex(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandWafProfileUrlAccessAccessPatternSrcaddr(d, i["srcaddr"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWafProfileUrlAccessAccessPatternId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileUrlAccessAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileUrlAccessId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWafProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address_list"); ok || d.HasChange("address_list") {
		t, err := expandWafProfileAddressList(d, v, "address_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-list"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWafProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("constraint"); ok || d.HasChange("constraint") {
		t, err := expandWafProfileConstraint(d, v, "constraint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["constraint"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandWafProfileExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("external"); ok || d.HasChange("external") {
		t, err := expandWafProfileExternal(d, v, "external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandWafProfileMethod(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWafProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("signature"); ok || d.HasChange("signature") {
		t, err := expandWafProfileSignature(d, v, "signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature"] = t
		}
	}

	if v, ok := d.GetOk("url_access"); ok || d.HasChange("url_access") {
		t, err := expandWafProfileUrlAccess(d, v, "url_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-access"] = t
		}
	}

	return &obj, nil
}
