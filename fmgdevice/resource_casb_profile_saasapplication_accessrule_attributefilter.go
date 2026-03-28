// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> CASB profile attribute filter.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbProfileSaasApplicationAccessRuleAttributeFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbProfileSaasApplicationAccessRuleAttributeFilterCreate,
		Read:   resourceCasbProfileSaasApplicationAccessRuleAttributeFilterRead,
		Update: resourceCasbProfileSaasApplicationAccessRuleAttributeFilterUpdate,
		Delete: resourceCasbProfileSaasApplicationAccessRuleAttributeFilterDelete,

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
			"saas_application": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"access_rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_match": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceCasbProfileSaasApplicationAccessRuleAttributeFilterCreate(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	access_rule := d.Get("access_rule").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["access_rule"] = access_rule

	obj, err := getObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbProfileSaasApplicationAccessRuleAttributeFilter resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbProfileSaasApplicationAccessRuleAttributeFilter(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbProfileSaasApplicationAccessRuleAttributeFilter(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbProfileSaasApplicationAccessRuleAttributeFilter resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateCasbProfileSaasApplicationAccessRuleAttributeFilter(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbProfileSaasApplicationAccessRuleAttributeFilter resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceCasbProfileSaasApplicationAccessRuleAttributeFilterRead(d, m)
}

func resourceCasbProfileSaasApplicationAccessRuleAttributeFilterUpdate(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	access_rule := d.Get("access_rule").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["access_rule"] = access_rule

	obj, err := getObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfileSaasApplicationAccessRuleAttributeFilter resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbProfileSaasApplicationAccessRuleAttributeFilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfileSaasApplicationAccessRuleAttributeFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceCasbProfileSaasApplicationAccessRuleAttributeFilterRead(d, m)
}

func resourceCasbProfileSaasApplicationAccessRuleAttributeFilterDelete(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	access_rule := d.Get("access_rule").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["access_rule"] = access_rule

	wsParams["adom"] = adomv

	err = c.DeleteCasbProfileSaasApplicationAccessRuleAttributeFilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbProfileSaasApplicationAccessRuleAttributeFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbProfileSaasApplicationAccessRuleAttributeFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
	saas_application := d.Get("saas_application").(string)
	access_rule := d.Get("access_rule").(string)
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
	if saas_application == "" {
		saas_application = importOptionChecking(m.(*FortiClient).Cfg, "saas_application")
		if saas_application == "" {
			return fmt.Errorf("Parameter saas_application is missing")
		}
		if err = d.Set("saas_application", saas_application); err != nil {
			return fmt.Errorf("Error set params saas_application: %v", err)
		}
	}
	if access_rule == "" {
		access_rule = importOptionChecking(m.(*FortiClient).Cfg, "access_rule")
		if access_rule == "" {
			return fmt.Errorf("Parameter access_rule is missing")
		}
		if err = d.Set("access_rule", access_rule); err != nil {
			return fmt.Errorf("Error set params access_rule: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["access_rule"] = access_rule

	o, err := c.ReadCasbProfileSaasApplicationAccessRuleAttributeFilter(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbProfileSaasApplicationAccessRuleAttributeFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbProfileSaasApplicationAccessRuleAttributeFilter resource from API: %v", err)
	}
	return nil
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAction4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAction4thl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "CasbProfileSaasApplicationAccessRuleAttributeFilter-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("attribute_match", flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch4thl(o["attribute-match"], d, "attribute_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["attribute-match"], "CasbProfileSaasApplicationAccessRuleAttributeFilter-AttributeMatch"); ok {
			if err = d.Set("attribute_match", vv); err != nil {
				return fmt.Errorf("Error reading attribute_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attribute_match: %v", err)
		}
	}

	if err = d.Set("fosid", flattenCasbProfileSaasApplicationAccessRuleAttributeFilterId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "CasbProfileSaasApplicationAccessRuleAttributeFilter-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterAction4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandCasbProfileSaasApplicationAccessRuleAttributeFilterAction4thl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("attribute_match"); ok || d.HasChange("attribute_match") {
		t, err := expandCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch4thl(d, v, "attribute_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-match"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandCasbProfileSaasApplicationAccessRuleAttributeFilterId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
