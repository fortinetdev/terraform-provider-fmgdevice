// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: BFD IPv6 multi-hop template table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBfd6MultihopTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBfd6MultihopTemplateCreate,
		Read:   resourceRouterBfd6MultihopTemplateRead,
		Update: resourceRouterBfd6MultihopTemplateUpdate,
		Delete: resourceRouterBfd6MultihopTemplateDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"auth_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd_desired_min_tx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bfd_detect_mult": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bfd_required_min_rx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"md5_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterBfd6MultihopTemplateCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	obj, err := getObjectRouterBfd6MultihopTemplate(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterBfd6MultihopTemplate resource while getting object: %v", err)
	}

	_, err = c.CreateRouterBfd6MultihopTemplate(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating RouterBfd6MultihopTemplate resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBfd6MultihopTemplateRead(d, m)
}

func resourceRouterBfd6MultihopTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	obj, err := getObjectRouterBfd6MultihopTemplate(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBfd6MultihopTemplate resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBfd6MultihopTemplate(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterBfd6MultihopTemplate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBfd6MultihopTemplateRead(d, m)
}

func resourceRouterBfd6MultihopTemplateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	err = c.DeleteRouterBfd6MultihopTemplate(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBfd6MultihopTemplate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBfd6MultihopTemplateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBfd6MultihopTemplate(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBfd6MultihopTemplate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBfd6MultihopTemplate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBfd6MultihopTemplate resource from API: %v", err)
	}
	return nil
}

func flattenRouterBfd6MultihopTemplateAuthMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateBfdDesiredMinTx2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateBfdDetectMult2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateBfdRequiredMinRx2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateDst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateSrc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterBfd6MultihopTemplate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_mode", flattenRouterBfd6MultihopTemplateAuthMode2edl(o["auth-mode"], d, "auth_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-mode"], "RouterBfd6MultihopTemplate-AuthMode"); ok {
			if err = d.Set("auth_mode", vv); err != nil {
				return fmt.Errorf("Error reading auth_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_mode: %v", err)
		}
	}

	if err = d.Set("bfd_desired_min_tx", flattenRouterBfd6MultihopTemplateBfdDesiredMinTx2edl(o["bfd-desired-min-tx"], d, "bfd_desired_min_tx")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd-desired-min-tx"], "RouterBfd6MultihopTemplate-BfdDesiredMinTx"); ok {
			if err = d.Set("bfd_desired_min_tx", vv); err != nil {
				return fmt.Errorf("Error reading bfd_desired_min_tx: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd_desired_min_tx: %v", err)
		}
	}

	if err = d.Set("bfd_detect_mult", flattenRouterBfd6MultihopTemplateBfdDetectMult2edl(o["bfd-detect-mult"], d, "bfd_detect_mult")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd-detect-mult"], "RouterBfd6MultihopTemplate-BfdDetectMult"); ok {
			if err = d.Set("bfd_detect_mult", vv); err != nil {
				return fmt.Errorf("Error reading bfd_detect_mult: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd_detect_mult: %v", err)
		}
	}

	if err = d.Set("bfd_required_min_rx", flattenRouterBfd6MultihopTemplateBfdRequiredMinRx2edl(o["bfd-required-min-rx"], d, "bfd_required_min_rx")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd-required-min-rx"], "RouterBfd6MultihopTemplate-BfdRequiredMinRx"); ok {
			if err = d.Set("bfd_required_min_rx", vv); err != nil {
				return fmt.Errorf("Error reading bfd_required_min_rx: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd_required_min_rx: %v", err)
		}
	}

	if err = d.Set("dst", flattenRouterBfd6MultihopTemplateDst2edl(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "RouterBfd6MultihopTemplate-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterBfd6MultihopTemplateId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterBfd6MultihopTemplate-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("src", flattenRouterBfd6MultihopTemplateSrc2edl(o["src"], d, "src")); err != nil {
		if vv, ok := fortiAPIPatch(o["src"], "RouterBfd6MultihopTemplate-Src"); ok {
			if err = d.Set("src", vv); err != nil {
				return fmt.Errorf("Error reading src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	return nil
}

func flattenRouterBfd6MultihopTemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBfd6MultihopTemplateAuthMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateBfdDesiredMinTx2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateBfdDetectMult2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateBfdRequiredMinRx2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateDst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateMd5Key2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterBfd6MultihopTemplateSrc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBfd6MultihopTemplate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_mode"); ok || d.HasChange("auth_mode") {
		t, err := expandRouterBfd6MultihopTemplateAuthMode2edl(d, v, "auth_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-mode"] = t
		}
	}

	if v, ok := d.GetOk("bfd_desired_min_tx"); ok || d.HasChange("bfd_desired_min_tx") {
		t, err := expandRouterBfd6MultihopTemplateBfdDesiredMinTx2edl(d, v, "bfd_desired_min_tx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-desired-min-tx"] = t
		}
	}

	if v, ok := d.GetOk("bfd_detect_mult"); ok || d.HasChange("bfd_detect_mult") {
		t, err := expandRouterBfd6MultihopTemplateBfdDetectMult2edl(d, v, "bfd_detect_mult")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-detect-mult"] = t
		}
	}

	if v, ok := d.GetOk("bfd_required_min_rx"); ok || d.HasChange("bfd_required_min_rx") {
		t, err := expandRouterBfd6MultihopTemplateBfdRequiredMinRx2edl(d, v, "bfd_required_min_rx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-required-min-rx"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandRouterBfd6MultihopTemplateDst2edl(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterBfd6MultihopTemplateId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("md5_key"); ok || d.HasChange("md5_key") {
		t, err := expandRouterBfd6MultihopTemplateMd5Key2edl(d, v, "md5_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["md5-key"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandRouterBfd6MultihopTemplateSrc2edl(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	return &obj, nil
}
