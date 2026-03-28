// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPS custom signature.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpsCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsCustomCreate,
		Read:   resourceIpsCustomRead,
		Update: resourceIpsCustomUpdate,
		Delete: resourceIpsCustomDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": &schema.Schema{
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
			"log_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"os": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sig_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceIpsCustomCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIpsCustom(d)
	if err != nil {
		return fmt.Errorf("Error creating IpsCustom resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("tag")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIpsCustom(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIpsCustom(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IpsCustom resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateIpsCustom(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IpsCustom resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "tag"))

	return resourceIpsCustomRead(d, m)
}

func resourceIpsCustomUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIpsCustom(d)
	if err != nil {
		return fmt.Errorf("Error updating IpsCustom resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIpsCustom(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IpsCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "tag"))

	return resourceIpsCustomRead(d, m)
}

func resourceIpsCustomDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteIpsCustom(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IpsCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsCustomRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadIpsCustom(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IpsCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsCustom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IpsCustom resource from API: %v", err)
	}
	return nil
}

func flattenIpsCustomAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsCustomApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsCustomComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsCustomLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsCustomLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsCustomLogPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsCustomOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsCustomProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsCustomRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsCustomSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsCustomSigName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsCustomSignature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsCustomStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsCustomTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIpsCustom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenIpsCustomAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "IpsCustom-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("application", flattenIpsCustomApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "IpsCustom-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("comment", flattenIpsCustomComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "IpsCustom-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("location", flattenIpsCustomLocation(o["location"], d, "location")); err != nil {
		if vv, ok := fortiAPIPatch(o["location"], "IpsCustom-Location"); ok {
			if err = d.Set("location", vv); err != nil {
				return fmt.Errorf("Error reading location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("log", flattenIpsCustomLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "IpsCustom-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenIpsCustomLogPacket(o["log-packet"], d, "log_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-packet"], "IpsCustom-LogPacket"); ok {
			if err = d.Set("log_packet", vv); err != nil {
				return fmt.Errorf("Error reading log_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if err = d.Set("os", flattenIpsCustomOs(o["os"], d, "os")); err != nil {
		if vv, ok := fortiAPIPatch(o["os"], "IpsCustom-Os"); ok {
			if err = d.Set("os", vv); err != nil {
				return fmt.Errorf("Error reading os: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("protocol", flattenIpsCustomProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "IpsCustom-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("rule_id", flattenIpsCustomRuleId(o["rule-id"], d, "rule_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["rule-id"], "IpsCustom-RuleId"); ok {
			if err = d.Set("rule_id", vv); err != nil {
				return fmt.Errorf("Error reading rule_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rule_id: %v", err)
		}
	}

	if err = d.Set("severity", flattenIpsCustomSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "IpsCustom-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("sig_name", flattenIpsCustomSigName(o["sig-name"], d, "sig_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["sig-name"], "IpsCustom-SigName"); ok {
			if err = d.Set("sig_name", vv); err != nil {
				return fmt.Errorf("Error reading sig_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sig_name: %v", err)
		}
	}

	if err = d.Set("signature", flattenIpsCustomSignature(o["signature"], d, "signature")); err != nil {
		if vv, ok := fortiAPIPatch(o["signature"], "IpsCustom-Signature"); ok {
			if err = d.Set("signature", vv); err != nil {
				return fmt.Errorf("Error reading signature: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signature: %v", err)
		}
	}

	if err = d.Set("status", flattenIpsCustomStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "IpsCustom-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tag", flattenIpsCustomTag(o["tag"], d, "tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag"], "IpsCustom-Tag"); ok {
			if err = d.Set("tag", vv); err != nil {
				return fmt.Errorf("Error reading tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag: %v", err)
		}
	}

	return nil
}

func flattenIpsCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIpsCustomAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsCustomComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsCustomLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomLogPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsCustomProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomSigName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIpsCustom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandIpsCustomAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandIpsCustomApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandIpsCustomComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok || d.HasChange("location") {
		t, err := expandIpsCustomLocation(d, v, "location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandIpsCustomLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok || d.HasChange("log_packet") {
		t, err := expandIpsCustomLogPacket(d, v, "log_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok || d.HasChange("os") {
		t, err := expandIpsCustomOs(d, v, "os")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandIpsCustomProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("rule_id"); ok || d.HasChange("rule_id") {
		t, err := expandIpsCustomRuleId(d, v, "rule_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule-id"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandIpsCustomSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("sig_name"); ok || d.HasChange("sig_name") {
		t, err := expandIpsCustomSigName(d, v, "sig_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sig-name"] = t
		}
	}

	if v, ok := d.GetOk("signature"); ok || d.HasChange("signature") {
		t, err := expandIpsCustomSignature(d, v, "signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandIpsCustomStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tag"); ok || d.HasChange("tag") {
		t, err := expandIpsCustomTag(d, v, "tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag"] = t
		}
	}

	return &obj, nil
}
