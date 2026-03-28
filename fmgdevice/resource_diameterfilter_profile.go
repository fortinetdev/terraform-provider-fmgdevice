// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Diameter filter profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDiameterFilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiameterFilterProfileCreate,
		Read:   resourceDiameterFilterProfileRead,
		Update: resourceDiameterFilterProfileUpdate,
		Delete: resourceDiameterFilterProfileDelete,

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
			"cmd_flags_reserve_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"command_code_invalid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"command_code_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message_length_invalid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"missing_request_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor_all_messages": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"protocol_version_invalid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_error_flag_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"track_requests_answers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDiameterFilterProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDiameterFilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating DiameterFilterProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDiameterFilterProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDiameterFilterProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DiameterFilterProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDiameterFilterProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DiameterFilterProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceDiameterFilterProfileRead(d, m)
}

func resourceDiameterFilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDiameterFilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating DiameterFilterProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDiameterFilterProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DiameterFilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDiameterFilterProfileRead(d, m)
}

func resourceDiameterFilterProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteDiameterFilterProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DiameterFilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDiameterFilterProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDiameterFilterProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DiameterFilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDiameterFilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DiameterFilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenDiameterFilterProfileCmdFlagsReserveSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDiameterFilterProfileCommandCodeInvalid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDiameterFilterProfileCommandCodeRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDiameterFilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDiameterFilterProfileLogPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDiameterFilterProfileMessageLengthInvalid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDiameterFilterProfileMissingRequestAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDiameterFilterProfileMonitorAllMessages(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDiameterFilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDiameterFilterProfileProtocolVersionInvalid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDiameterFilterProfileRequestErrorFlagSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDiameterFilterProfileTrackRequestsAnswers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDiameterFilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cmd_flags_reserve_set", flattenDiameterFilterProfileCmdFlagsReserveSet(o["cmd-flags-reserve-set"], d, "cmd_flags_reserve_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["cmd-flags-reserve-set"], "DiameterFilterProfile-CmdFlagsReserveSet"); ok {
			if err = d.Set("cmd_flags_reserve_set", vv); err != nil {
				return fmt.Errorf("Error reading cmd_flags_reserve_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cmd_flags_reserve_set: %v", err)
		}
	}

	if err = d.Set("command_code_invalid", flattenDiameterFilterProfileCommandCodeInvalid(o["command-code-invalid"], d, "command_code_invalid")); err != nil {
		if vv, ok := fortiAPIPatch(o["command-code-invalid"], "DiameterFilterProfile-CommandCodeInvalid"); ok {
			if err = d.Set("command_code_invalid", vv); err != nil {
				return fmt.Errorf("Error reading command_code_invalid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command_code_invalid: %v", err)
		}
	}

	if err = d.Set("command_code_range", flattenDiameterFilterProfileCommandCodeRange(o["command-code-range"], d, "command_code_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["command-code-range"], "DiameterFilterProfile-CommandCodeRange"); ok {
			if err = d.Set("command_code_range", vv); err != nil {
				return fmt.Errorf("Error reading command_code_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command_code_range: %v", err)
		}
	}

	if err = d.Set("comment", flattenDiameterFilterProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "DiameterFilterProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenDiameterFilterProfileLogPacket(o["log-packet"], d, "log_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-packet"], "DiameterFilterProfile-LogPacket"); ok {
			if err = d.Set("log_packet", vv); err != nil {
				return fmt.Errorf("Error reading log_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if err = d.Set("message_length_invalid", flattenDiameterFilterProfileMessageLengthInvalid(o["message-length-invalid"], d, "message_length_invalid")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-length-invalid"], "DiameterFilterProfile-MessageLengthInvalid"); ok {
			if err = d.Set("message_length_invalid", vv); err != nil {
				return fmt.Errorf("Error reading message_length_invalid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_length_invalid: %v", err)
		}
	}

	if err = d.Set("missing_request_action", flattenDiameterFilterProfileMissingRequestAction(o["missing-request-action"], d, "missing_request_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["missing-request-action"], "DiameterFilterProfile-MissingRequestAction"); ok {
			if err = d.Set("missing_request_action", vv); err != nil {
				return fmt.Errorf("Error reading missing_request_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading missing_request_action: %v", err)
		}
	}

	if err = d.Set("monitor_all_messages", flattenDiameterFilterProfileMonitorAllMessages(o["monitor-all-messages"], d, "monitor_all_messages")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-all-messages"], "DiameterFilterProfile-MonitorAllMessages"); ok {
			if err = d.Set("monitor_all_messages", vv); err != nil {
				return fmt.Errorf("Error reading monitor_all_messages: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_all_messages: %v", err)
		}
	}

	if err = d.Set("name", flattenDiameterFilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DiameterFilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol_version_invalid", flattenDiameterFilterProfileProtocolVersionInvalid(o["protocol-version-invalid"], d, "protocol_version_invalid")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol-version-invalid"], "DiameterFilterProfile-ProtocolVersionInvalid"); ok {
			if err = d.Set("protocol_version_invalid", vv); err != nil {
				return fmt.Errorf("Error reading protocol_version_invalid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol_version_invalid: %v", err)
		}
	}

	if err = d.Set("request_error_flag_set", flattenDiameterFilterProfileRequestErrorFlagSet(o["request-error-flag-set"], d, "request_error_flag_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["request-error-flag-set"], "DiameterFilterProfile-RequestErrorFlagSet"); ok {
			if err = d.Set("request_error_flag_set", vv); err != nil {
				return fmt.Errorf("Error reading request_error_flag_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request_error_flag_set: %v", err)
		}
	}

	if err = d.Set("track_requests_answers", flattenDiameterFilterProfileTrackRequestsAnswers(o["track-requests-answers"], d, "track_requests_answers")); err != nil {
		if vv, ok := fortiAPIPatch(o["track-requests-answers"], "DiameterFilterProfile-TrackRequestsAnswers"); ok {
			if err = d.Set("track_requests_answers", vv); err != nil {
				return fmt.Errorf("Error reading track_requests_answers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading track_requests_answers: %v", err)
		}
	}

	return nil
}

func flattenDiameterFilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDiameterFilterProfileCmdFlagsReserveSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileCommandCodeInvalid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileCommandCodeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileLogPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileMessageLengthInvalid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileMissingRequestAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileMonitorAllMessages(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileProtocolVersionInvalid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileRequestErrorFlagSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileTrackRequestsAnswers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDiameterFilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cmd_flags_reserve_set"); ok || d.HasChange("cmd_flags_reserve_set") {
		t, err := expandDiameterFilterProfileCmdFlagsReserveSet(d, v, "cmd_flags_reserve_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmd-flags-reserve-set"] = t
		}
	}

	if v, ok := d.GetOk("command_code_invalid"); ok || d.HasChange("command_code_invalid") {
		t, err := expandDiameterFilterProfileCommandCodeInvalid(d, v, "command_code_invalid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-code-invalid"] = t
		}
	}

	if v, ok := d.GetOk("command_code_range"); ok || d.HasChange("command_code_range") {
		t, err := expandDiameterFilterProfileCommandCodeRange(d, v, "command_code_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-code-range"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandDiameterFilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok || d.HasChange("log_packet") {
		t, err := expandDiameterFilterProfileLogPacket(d, v, "log_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("message_length_invalid"); ok || d.HasChange("message_length_invalid") {
		t, err := expandDiameterFilterProfileMessageLengthInvalid(d, v, "message_length_invalid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-length-invalid"] = t
		}
	}

	if v, ok := d.GetOk("missing_request_action"); ok || d.HasChange("missing_request_action") {
		t, err := expandDiameterFilterProfileMissingRequestAction(d, v, "missing_request_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["missing-request-action"] = t
		}
	}

	if v, ok := d.GetOk("monitor_all_messages"); ok || d.HasChange("monitor_all_messages") {
		t, err := expandDiameterFilterProfileMonitorAllMessages(d, v, "monitor_all_messages")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-all-messages"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDiameterFilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol_version_invalid"); ok || d.HasChange("protocol_version_invalid") {
		t, err := expandDiameterFilterProfileProtocolVersionInvalid(d, v, "protocol_version_invalid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol-version-invalid"] = t
		}
	}

	if v, ok := d.GetOk("request_error_flag_set"); ok || d.HasChange("request_error_flag_set") {
		t, err := expandDiameterFilterProfileRequestErrorFlagSet(d, v, "request_error_flag_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-error-flag-set"] = t
		}
	}

	if v, ok := d.GetOk("track_requests_answers"); ok || d.HasChange("track_requests_answers") {
		t, err := expandDiameterFilterProfileTrackRequestsAnswers(d, v, "track_requests_answers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["track-requests-answers"] = t
		}
	}

	return &obj, nil
}
