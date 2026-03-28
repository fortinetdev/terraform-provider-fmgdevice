// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> SIP.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVoipProfileSip() *schema.Resource {
	return &schema.Resource{
		Create: resourceVoipProfileSipUpdate,
		Read:   resourceVoipProfileSipRead,
		Update: resourceVoipProfileSipUpdate,
		Delete: resourceVoipProfileSipDelete,

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
			"ack_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ack_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_ack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_bye": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_cancel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_geo_red_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_invite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_long_lines": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_notify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_prack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_publish": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_refer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_register": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_subscribe": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_unknown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bye_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bye_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"call_id_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"call_keepalive": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cancel_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cancel_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contact_fixup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"content_type_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hnt_restrict_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hosted_nat_traversal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"info_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"info_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invite_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"invite_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_rtp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_call_summary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_violations": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_allow": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_call_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_contact": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_content_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_content_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_cseq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_expires": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_from": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_max_forwards": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_no_proxy_require": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_no_require": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_p_asserted_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_rack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_record_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_rseq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_sdp_a": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_sdp_b": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_sdp_c": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_sdp_i": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_sdp_k": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_sdp_m": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_sdp_o": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_sdp_r": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_sdp_s": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_sdp_t": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_sdp_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_sdp_z": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_to": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_header_via": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_request_line": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_body_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_dialogs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_idle_dialogs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_line_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"message_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"message_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat_port_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat_trace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"no_sdp_fixup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"notify_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"notify_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"open_contact_pinhole": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"open_record_route_pinhole": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"open_register_pinhole": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"open_via_pinhole": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"options_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"options_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prack_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"prack_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preserve_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"provisional_invite_expiry_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"publish_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"publish_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refer_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"refer_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"register_contact_trace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"register_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"register_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rfc2543_branch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_auth_client": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_auth_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_client_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_client_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_pfs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_send_empty_frags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_server_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"strict_register": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscribe_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"subscribe_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"update_rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVoipProfileSipUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVoipProfileSip(d)
	if err != nil {
		return fmt.Errorf("Error updating VoipProfileSip resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVoipProfileSip(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VoipProfileSip resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("VoipProfileSip")

	return resourceVoipProfileSipRead(d, m)
}

func resourceVoipProfileSipDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVoipProfileSip(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VoipProfileSip resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVoipProfileSipRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVoipProfileSip(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VoipProfileSip resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVoipProfileSip(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VoipProfileSip resource from API: %v", err)
	}
	return nil
}

func flattenVoipProfileSipAckRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipAckRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockAck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockBye2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockCancel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockGeoRedOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockInfo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockInvite2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockLongLines2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockMessage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockNotify2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockPrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockPublish2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockRefer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockRegister2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockSubscribe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockUnknown2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipBlockUpdate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipByeRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipByeRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipCallIdRegex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipCallKeepalive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipCancelRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipCancelRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipContactFixup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipContentTypeRegex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipHntRestrictSourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipHostedNatTraversal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipInfoRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipInfoRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipInviteRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipInviteRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipIpsRtp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipLogCallSummary2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipLogViolations2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderAllow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderCallId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderContact2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderContentLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderContentType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderCseq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderExpires2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderFrom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderMaxForwards2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderNoProxyRequire2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderNoRequire2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderPAssertedIdentity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderRack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderRecordRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderRseq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpA2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpB2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpC2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpI2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpK2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpM2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpO2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpR2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpS2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpT2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpZ2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderTo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderVia2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedRequestLine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMaxBodyLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMaxDialogs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMaxIdleDialogs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMaxLineLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMessageRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipMessageRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipNatPortRange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipNatTrace2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipNoSdpFixup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipNotifyRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipNotifyRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipOpenContactPinhole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipOpenRecordRoutePinhole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipOpenRegisterPinhole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipOpenViaPinhole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipOptionsRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipOptionsRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipPrackRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipPrackRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipPreserveOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipProvisionalInviteExpiryTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipPublishRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipPublishRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipReferRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipReferRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipRegisterContactTrace2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipRegisterRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipRegisterRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipRfc2543Branch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipRtp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipSslAlgorithm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipSslAuthClient2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVoipProfileSipSslAuthServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVoipProfileSipSslClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVoipProfileSipSslClientRenegotiation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipSslMaxVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipSslMinVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipSslMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipSslPfs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipSslSendEmptyFrags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipSslServerCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVoipProfileSipStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipStrictRegister2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipSubscribeRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipSubscribeRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipUnknownHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipUpdateRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSipUpdateRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVoipProfileSip(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ack_rate", flattenVoipProfileSipAckRate2edl(o["ack-rate"], d, "ack_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ack-rate"], "VoipProfileSip-AckRate"); ok {
			if err = d.Set("ack_rate", vv); err != nil {
				return fmt.Errorf("Error reading ack_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ack_rate: %v", err)
		}
	}

	if err = d.Set("ack_rate_track", flattenVoipProfileSipAckRateTrack2edl(o["ack-rate-track"], d, "ack_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["ack-rate-track"], "VoipProfileSip-AckRateTrack"); ok {
			if err = d.Set("ack_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading ack_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ack_rate_track: %v", err)
		}
	}

	if err = d.Set("block_ack", flattenVoipProfileSipBlockAck2edl(o["block-ack"], d, "block_ack")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-ack"], "VoipProfileSip-BlockAck"); ok {
			if err = d.Set("block_ack", vv); err != nil {
				return fmt.Errorf("Error reading block_ack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_ack: %v", err)
		}
	}

	if err = d.Set("block_bye", flattenVoipProfileSipBlockBye2edl(o["block-bye"], d, "block_bye")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-bye"], "VoipProfileSip-BlockBye"); ok {
			if err = d.Set("block_bye", vv); err != nil {
				return fmt.Errorf("Error reading block_bye: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_bye: %v", err)
		}
	}

	if err = d.Set("block_cancel", flattenVoipProfileSipBlockCancel2edl(o["block-cancel"], d, "block_cancel")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-cancel"], "VoipProfileSip-BlockCancel"); ok {
			if err = d.Set("block_cancel", vv); err != nil {
				return fmt.Errorf("Error reading block_cancel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_cancel: %v", err)
		}
	}

	if err = d.Set("block_geo_red_options", flattenVoipProfileSipBlockGeoRedOptions2edl(o["block-geo-red-options"], d, "block_geo_red_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-geo-red-options"], "VoipProfileSip-BlockGeoRedOptions"); ok {
			if err = d.Set("block_geo_red_options", vv); err != nil {
				return fmt.Errorf("Error reading block_geo_red_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_geo_red_options: %v", err)
		}
	}

	if err = d.Set("block_info", flattenVoipProfileSipBlockInfo2edl(o["block-info"], d, "block_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-info"], "VoipProfileSip-BlockInfo"); ok {
			if err = d.Set("block_info", vv); err != nil {
				return fmt.Errorf("Error reading block_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_info: %v", err)
		}
	}

	if err = d.Set("block_invite", flattenVoipProfileSipBlockInvite2edl(o["block-invite"], d, "block_invite")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-invite"], "VoipProfileSip-BlockInvite"); ok {
			if err = d.Set("block_invite", vv); err != nil {
				return fmt.Errorf("Error reading block_invite: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_invite: %v", err)
		}
	}

	if err = d.Set("block_long_lines", flattenVoipProfileSipBlockLongLines2edl(o["block-long-lines"], d, "block_long_lines")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-long-lines"], "VoipProfileSip-BlockLongLines"); ok {
			if err = d.Set("block_long_lines", vv); err != nil {
				return fmt.Errorf("Error reading block_long_lines: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_long_lines: %v", err)
		}
	}

	if err = d.Set("block_message", flattenVoipProfileSipBlockMessage2edl(o["block-message"], d, "block_message")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-message"], "VoipProfileSip-BlockMessage"); ok {
			if err = d.Set("block_message", vv); err != nil {
				return fmt.Errorf("Error reading block_message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_message: %v", err)
		}
	}

	if err = d.Set("block_notify", flattenVoipProfileSipBlockNotify2edl(o["block-notify"], d, "block_notify")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-notify"], "VoipProfileSip-BlockNotify"); ok {
			if err = d.Set("block_notify", vv); err != nil {
				return fmt.Errorf("Error reading block_notify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_notify: %v", err)
		}
	}

	if err = d.Set("block_options", flattenVoipProfileSipBlockOptions2edl(o["block-options"], d, "block_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-options"], "VoipProfileSip-BlockOptions"); ok {
			if err = d.Set("block_options", vv); err != nil {
				return fmt.Errorf("Error reading block_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_options: %v", err)
		}
	}

	if err = d.Set("block_prack", flattenVoipProfileSipBlockPrack2edl(o["block-prack"], d, "block_prack")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-prack"], "VoipProfileSip-BlockPrack"); ok {
			if err = d.Set("block_prack", vv); err != nil {
				return fmt.Errorf("Error reading block_prack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_prack: %v", err)
		}
	}

	if err = d.Set("block_publish", flattenVoipProfileSipBlockPublish2edl(o["block-publish"], d, "block_publish")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-publish"], "VoipProfileSip-BlockPublish"); ok {
			if err = d.Set("block_publish", vv); err != nil {
				return fmt.Errorf("Error reading block_publish: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_publish: %v", err)
		}
	}

	if err = d.Set("block_refer", flattenVoipProfileSipBlockRefer2edl(o["block-refer"], d, "block_refer")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-refer"], "VoipProfileSip-BlockRefer"); ok {
			if err = d.Set("block_refer", vv); err != nil {
				return fmt.Errorf("Error reading block_refer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_refer: %v", err)
		}
	}

	if err = d.Set("block_register", flattenVoipProfileSipBlockRegister2edl(o["block-register"], d, "block_register")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-register"], "VoipProfileSip-BlockRegister"); ok {
			if err = d.Set("block_register", vv); err != nil {
				return fmt.Errorf("Error reading block_register: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_register: %v", err)
		}
	}

	if err = d.Set("block_subscribe", flattenVoipProfileSipBlockSubscribe2edl(o["block-subscribe"], d, "block_subscribe")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-subscribe"], "VoipProfileSip-BlockSubscribe"); ok {
			if err = d.Set("block_subscribe", vv); err != nil {
				return fmt.Errorf("Error reading block_subscribe: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_subscribe: %v", err)
		}
	}

	if err = d.Set("block_unknown", flattenVoipProfileSipBlockUnknown2edl(o["block-unknown"], d, "block_unknown")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-unknown"], "VoipProfileSip-BlockUnknown"); ok {
			if err = d.Set("block_unknown", vv); err != nil {
				return fmt.Errorf("Error reading block_unknown: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_unknown: %v", err)
		}
	}

	if err = d.Set("block_update", flattenVoipProfileSipBlockUpdate2edl(o["block-update"], d, "block_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-update"], "VoipProfileSip-BlockUpdate"); ok {
			if err = d.Set("block_update", vv); err != nil {
				return fmt.Errorf("Error reading block_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_update: %v", err)
		}
	}

	if err = d.Set("bye_rate", flattenVoipProfileSipByeRate2edl(o["bye-rate"], d, "bye_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["bye-rate"], "VoipProfileSip-ByeRate"); ok {
			if err = d.Set("bye_rate", vv); err != nil {
				return fmt.Errorf("Error reading bye_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bye_rate: %v", err)
		}
	}

	if err = d.Set("bye_rate_track", flattenVoipProfileSipByeRateTrack2edl(o["bye-rate-track"], d, "bye_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["bye-rate-track"], "VoipProfileSip-ByeRateTrack"); ok {
			if err = d.Set("bye_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading bye_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bye_rate_track: %v", err)
		}
	}

	if err = d.Set("call_id_regex", flattenVoipProfileSipCallIdRegex2edl(o["call-id-regex"], d, "call_id_regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-id-regex"], "VoipProfileSip-CallIdRegex"); ok {
			if err = d.Set("call_id_regex", vv); err != nil {
				return fmt.Errorf("Error reading call_id_regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_id_regex: %v", err)
		}
	}

	if err = d.Set("call_keepalive", flattenVoipProfileSipCallKeepalive2edl(o["call-keepalive"], d, "call_keepalive")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-keepalive"], "VoipProfileSip-CallKeepalive"); ok {
			if err = d.Set("call_keepalive", vv); err != nil {
				return fmt.Errorf("Error reading call_keepalive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_keepalive: %v", err)
		}
	}

	if err = d.Set("cancel_rate", flattenVoipProfileSipCancelRate2edl(o["cancel-rate"], d, "cancel_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["cancel-rate"], "VoipProfileSip-CancelRate"); ok {
			if err = d.Set("cancel_rate", vv); err != nil {
				return fmt.Errorf("Error reading cancel_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cancel_rate: %v", err)
		}
	}

	if err = d.Set("cancel_rate_track", flattenVoipProfileSipCancelRateTrack2edl(o["cancel-rate-track"], d, "cancel_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["cancel-rate-track"], "VoipProfileSip-CancelRateTrack"); ok {
			if err = d.Set("cancel_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading cancel_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cancel_rate_track: %v", err)
		}
	}

	if err = d.Set("contact_fixup", flattenVoipProfileSipContactFixup2edl(o["contact-fixup"], d, "contact_fixup")); err != nil {
		if vv, ok := fortiAPIPatch(o["contact-fixup"], "VoipProfileSip-ContactFixup"); ok {
			if err = d.Set("contact_fixup", vv); err != nil {
				return fmt.Errorf("Error reading contact_fixup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading contact_fixup: %v", err)
		}
	}

	if err = d.Set("content_type_regex", flattenVoipProfileSipContentTypeRegex2edl(o["content-type-regex"], d, "content_type_regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["content-type-regex"], "VoipProfileSip-ContentTypeRegex"); ok {
			if err = d.Set("content_type_regex", vv); err != nil {
				return fmt.Errorf("Error reading content_type_regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content_type_regex: %v", err)
		}
	}

	if err = d.Set("hnt_restrict_source_ip", flattenVoipProfileSipHntRestrictSourceIp2edl(o["hnt-restrict-source-ip"], d, "hnt_restrict_source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["hnt-restrict-source-ip"], "VoipProfileSip-HntRestrictSourceIp"); ok {
			if err = d.Set("hnt_restrict_source_ip", vv); err != nil {
				return fmt.Errorf("Error reading hnt_restrict_source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hnt_restrict_source_ip: %v", err)
		}
	}

	if err = d.Set("hosted_nat_traversal", flattenVoipProfileSipHostedNatTraversal2edl(o["hosted-nat-traversal"], d, "hosted_nat_traversal")); err != nil {
		if vv, ok := fortiAPIPatch(o["hosted-nat-traversal"], "VoipProfileSip-HostedNatTraversal"); ok {
			if err = d.Set("hosted_nat_traversal", vv); err != nil {
				return fmt.Errorf("Error reading hosted_nat_traversal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hosted_nat_traversal: %v", err)
		}
	}

	if err = d.Set("info_rate", flattenVoipProfileSipInfoRate2edl(o["info-rate"], d, "info_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["info-rate"], "VoipProfileSip-InfoRate"); ok {
			if err = d.Set("info_rate", vv); err != nil {
				return fmt.Errorf("Error reading info_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading info_rate: %v", err)
		}
	}

	if err = d.Set("info_rate_track", flattenVoipProfileSipInfoRateTrack2edl(o["info-rate-track"], d, "info_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["info-rate-track"], "VoipProfileSip-InfoRateTrack"); ok {
			if err = d.Set("info_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading info_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading info_rate_track: %v", err)
		}
	}

	if err = d.Set("invite_rate", flattenVoipProfileSipInviteRate2edl(o["invite-rate"], d, "invite_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["invite-rate"], "VoipProfileSip-InviteRate"); ok {
			if err = d.Set("invite_rate", vv); err != nil {
				return fmt.Errorf("Error reading invite_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading invite_rate: %v", err)
		}
	}

	if err = d.Set("invite_rate_track", flattenVoipProfileSipInviteRateTrack2edl(o["invite-rate-track"], d, "invite_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["invite-rate-track"], "VoipProfileSip-InviteRateTrack"); ok {
			if err = d.Set("invite_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading invite_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading invite_rate_track: %v", err)
		}
	}

	if err = d.Set("ips_rtp", flattenVoipProfileSipIpsRtp2edl(o["ips-rtp"], d, "ips_rtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-rtp"], "VoipProfileSip-IpsRtp"); ok {
			if err = d.Set("ips_rtp", vv); err != nil {
				return fmt.Errorf("Error reading ips_rtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_rtp: %v", err)
		}
	}

	if err = d.Set("log_call_summary", flattenVoipProfileSipLogCallSummary2edl(o["log-call-summary"], d, "log_call_summary")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-call-summary"], "VoipProfileSip-LogCallSummary"); ok {
			if err = d.Set("log_call_summary", vv); err != nil {
				return fmt.Errorf("Error reading log_call_summary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_call_summary: %v", err)
		}
	}

	if err = d.Set("log_violations", flattenVoipProfileSipLogViolations2edl(o["log-violations"], d, "log_violations")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-violations"], "VoipProfileSip-LogViolations"); ok {
			if err = d.Set("log_violations", vv); err != nil {
				return fmt.Errorf("Error reading log_violations: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_violations: %v", err)
		}
	}

	if err = d.Set("malformed_header_allow", flattenVoipProfileSipMalformedHeaderAllow2edl(o["malformed-header-allow"], d, "malformed_header_allow")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-allow"], "VoipProfileSip-MalformedHeaderAllow"); ok {
			if err = d.Set("malformed_header_allow", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_allow: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_allow: %v", err)
		}
	}

	if err = d.Set("malformed_header_call_id", flattenVoipProfileSipMalformedHeaderCallId2edl(o["malformed-header-call-id"], d, "malformed_header_call_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-call-id"], "VoipProfileSip-MalformedHeaderCallId"); ok {
			if err = d.Set("malformed_header_call_id", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_call_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_call_id: %v", err)
		}
	}

	if err = d.Set("malformed_header_contact", flattenVoipProfileSipMalformedHeaderContact2edl(o["malformed-header-contact"], d, "malformed_header_contact")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-contact"], "VoipProfileSip-MalformedHeaderContact"); ok {
			if err = d.Set("malformed_header_contact", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_contact: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_contact: %v", err)
		}
	}

	if err = d.Set("malformed_header_content_length", flattenVoipProfileSipMalformedHeaderContentLength2edl(o["malformed-header-content-length"], d, "malformed_header_content_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-content-length"], "VoipProfileSip-MalformedHeaderContentLength"); ok {
			if err = d.Set("malformed_header_content_length", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_content_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_content_length: %v", err)
		}
	}

	if err = d.Set("malformed_header_content_type", flattenVoipProfileSipMalformedHeaderContentType2edl(o["malformed-header-content-type"], d, "malformed_header_content_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-content-type"], "VoipProfileSip-MalformedHeaderContentType"); ok {
			if err = d.Set("malformed_header_content_type", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_content_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_content_type: %v", err)
		}
	}

	if err = d.Set("malformed_header_cseq", flattenVoipProfileSipMalformedHeaderCseq2edl(o["malformed-header-cseq"], d, "malformed_header_cseq")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-cseq"], "VoipProfileSip-MalformedHeaderCseq"); ok {
			if err = d.Set("malformed_header_cseq", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_cseq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_cseq: %v", err)
		}
	}

	if err = d.Set("malformed_header_expires", flattenVoipProfileSipMalformedHeaderExpires2edl(o["malformed-header-expires"], d, "malformed_header_expires")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-expires"], "VoipProfileSip-MalformedHeaderExpires"); ok {
			if err = d.Set("malformed_header_expires", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_expires: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_expires: %v", err)
		}
	}

	if err = d.Set("malformed_header_from", flattenVoipProfileSipMalformedHeaderFrom2edl(o["malformed-header-from"], d, "malformed_header_from")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-from"], "VoipProfileSip-MalformedHeaderFrom"); ok {
			if err = d.Set("malformed_header_from", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_from: %v", err)
		}
	}

	if err = d.Set("malformed_header_max_forwards", flattenVoipProfileSipMalformedHeaderMaxForwards2edl(o["malformed-header-max-forwards"], d, "malformed_header_max_forwards")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-max-forwards"], "VoipProfileSip-MalformedHeaderMaxForwards"); ok {
			if err = d.Set("malformed_header_max_forwards", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_max_forwards: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_max_forwards: %v", err)
		}
	}

	if err = d.Set("malformed_header_no_proxy_require", flattenVoipProfileSipMalformedHeaderNoProxyRequire2edl(o["malformed-header-no-proxy-require"], d, "malformed_header_no_proxy_require")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-no-proxy-require"], "VoipProfileSip-MalformedHeaderNoProxyRequire"); ok {
			if err = d.Set("malformed_header_no_proxy_require", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_no_proxy_require: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_no_proxy_require: %v", err)
		}
	}

	if err = d.Set("malformed_header_no_require", flattenVoipProfileSipMalformedHeaderNoRequire2edl(o["malformed-header-no-require"], d, "malformed_header_no_require")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-no-require"], "VoipProfileSip-MalformedHeaderNoRequire"); ok {
			if err = d.Set("malformed_header_no_require", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_no_require: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_no_require: %v", err)
		}
	}

	if err = d.Set("malformed_header_p_asserted_identity", flattenVoipProfileSipMalformedHeaderPAssertedIdentity2edl(o["malformed-header-p-asserted-identity"], d, "malformed_header_p_asserted_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-p-asserted-identity"], "VoipProfileSip-MalformedHeaderPAssertedIdentity"); ok {
			if err = d.Set("malformed_header_p_asserted_identity", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_p_asserted_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_p_asserted_identity: %v", err)
		}
	}

	if err = d.Set("malformed_header_rack", flattenVoipProfileSipMalformedHeaderRack2edl(o["malformed-header-rack"], d, "malformed_header_rack")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-rack"], "VoipProfileSip-MalformedHeaderRack"); ok {
			if err = d.Set("malformed_header_rack", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_rack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_rack: %v", err)
		}
	}

	if err = d.Set("malformed_header_record_route", flattenVoipProfileSipMalformedHeaderRecordRoute2edl(o["malformed-header-record-route"], d, "malformed_header_record_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-record-route"], "VoipProfileSip-MalformedHeaderRecordRoute"); ok {
			if err = d.Set("malformed_header_record_route", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_record_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_record_route: %v", err)
		}
	}

	if err = d.Set("malformed_header_route", flattenVoipProfileSipMalformedHeaderRoute2edl(o["malformed-header-route"], d, "malformed_header_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-route"], "VoipProfileSip-MalformedHeaderRoute"); ok {
			if err = d.Set("malformed_header_route", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_route: %v", err)
		}
	}

	if err = d.Set("malformed_header_rseq", flattenVoipProfileSipMalformedHeaderRseq2edl(o["malformed-header-rseq"], d, "malformed_header_rseq")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-rseq"], "VoipProfileSip-MalformedHeaderRseq"); ok {
			if err = d.Set("malformed_header_rseq", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_rseq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_rseq: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_a", flattenVoipProfileSipMalformedHeaderSdpA2edl(o["malformed-header-sdp-a"], d, "malformed_header_sdp_a")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-a"], "VoipProfileSip-MalformedHeaderSdpA"); ok {
			if err = d.Set("malformed_header_sdp_a", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_a: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_a: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_b", flattenVoipProfileSipMalformedHeaderSdpB2edl(o["malformed-header-sdp-b"], d, "malformed_header_sdp_b")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-b"], "VoipProfileSip-MalformedHeaderSdpB"); ok {
			if err = d.Set("malformed_header_sdp_b", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_b: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_b: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_c", flattenVoipProfileSipMalformedHeaderSdpC2edl(o["malformed-header-sdp-c"], d, "malformed_header_sdp_c")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-c"], "VoipProfileSip-MalformedHeaderSdpC"); ok {
			if err = d.Set("malformed_header_sdp_c", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_c: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_c: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_i", flattenVoipProfileSipMalformedHeaderSdpI2edl(o["malformed-header-sdp-i"], d, "malformed_header_sdp_i")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-i"], "VoipProfileSip-MalformedHeaderSdpI"); ok {
			if err = d.Set("malformed_header_sdp_i", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_i: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_i: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_k", flattenVoipProfileSipMalformedHeaderSdpK2edl(o["malformed-header-sdp-k"], d, "malformed_header_sdp_k")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-k"], "VoipProfileSip-MalformedHeaderSdpK"); ok {
			if err = d.Set("malformed_header_sdp_k", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_k: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_k: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_m", flattenVoipProfileSipMalformedHeaderSdpM2edl(o["malformed-header-sdp-m"], d, "malformed_header_sdp_m")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-m"], "VoipProfileSip-MalformedHeaderSdpM"); ok {
			if err = d.Set("malformed_header_sdp_m", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_m: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_m: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_o", flattenVoipProfileSipMalformedHeaderSdpO2edl(o["malformed-header-sdp-o"], d, "malformed_header_sdp_o")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-o"], "VoipProfileSip-MalformedHeaderSdpO"); ok {
			if err = d.Set("malformed_header_sdp_o", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_o: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_o: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_r", flattenVoipProfileSipMalformedHeaderSdpR2edl(o["malformed-header-sdp-r"], d, "malformed_header_sdp_r")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-r"], "VoipProfileSip-MalformedHeaderSdpR"); ok {
			if err = d.Set("malformed_header_sdp_r", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_r: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_r: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_s", flattenVoipProfileSipMalformedHeaderSdpS2edl(o["malformed-header-sdp-s"], d, "malformed_header_sdp_s")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-s"], "VoipProfileSip-MalformedHeaderSdpS"); ok {
			if err = d.Set("malformed_header_sdp_s", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_s: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_s: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_t", flattenVoipProfileSipMalformedHeaderSdpT2edl(o["malformed-header-sdp-t"], d, "malformed_header_sdp_t")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-t"], "VoipProfileSip-MalformedHeaderSdpT"); ok {
			if err = d.Set("malformed_header_sdp_t", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_t: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_t: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_v", flattenVoipProfileSipMalformedHeaderSdpV2edl(o["malformed-header-sdp-v"], d, "malformed_header_sdp_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-v"], "VoipProfileSip-MalformedHeaderSdpV"); ok {
			if err = d.Set("malformed_header_sdp_v", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_v: %v", err)
		}
	}

	if err = d.Set("malformed_header_sdp_z", flattenVoipProfileSipMalformedHeaderSdpZ2edl(o["malformed-header-sdp-z"], d, "malformed_header_sdp_z")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-sdp-z"], "VoipProfileSip-MalformedHeaderSdpZ"); ok {
			if err = d.Set("malformed_header_sdp_z", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_sdp_z: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_sdp_z: %v", err)
		}
	}

	if err = d.Set("malformed_header_to", flattenVoipProfileSipMalformedHeaderTo2edl(o["malformed-header-to"], d, "malformed_header_to")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-to"], "VoipProfileSip-MalformedHeaderTo"); ok {
			if err = d.Set("malformed_header_to", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_to: %v", err)
		}
	}

	if err = d.Set("malformed_header_via", flattenVoipProfileSipMalformedHeaderVia2edl(o["malformed-header-via"], d, "malformed_header_via")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-header-via"], "VoipProfileSip-MalformedHeaderVia"); ok {
			if err = d.Set("malformed_header_via", vv); err != nil {
				return fmt.Errorf("Error reading malformed_header_via: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_header_via: %v", err)
		}
	}

	if err = d.Set("malformed_request_line", flattenVoipProfileSipMalformedRequestLine2edl(o["malformed-request-line"], d, "malformed_request_line")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-request-line"], "VoipProfileSip-MalformedRequestLine"); ok {
			if err = d.Set("malformed_request_line", vv); err != nil {
				return fmt.Errorf("Error reading malformed_request_line: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_request_line: %v", err)
		}
	}

	if err = d.Set("max_body_length", flattenVoipProfileSipMaxBodyLength2edl(o["max-body-length"], d, "max_body_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-body-length"], "VoipProfileSip-MaxBodyLength"); ok {
			if err = d.Set("max_body_length", vv); err != nil {
				return fmt.Errorf("Error reading max_body_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_body_length: %v", err)
		}
	}

	if err = d.Set("max_dialogs", flattenVoipProfileSipMaxDialogs2edl(o["max-dialogs"], d, "max_dialogs")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-dialogs"], "VoipProfileSip-MaxDialogs"); ok {
			if err = d.Set("max_dialogs", vv); err != nil {
				return fmt.Errorf("Error reading max_dialogs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_dialogs: %v", err)
		}
	}

	if err = d.Set("max_idle_dialogs", flattenVoipProfileSipMaxIdleDialogs2edl(o["max-idle-dialogs"], d, "max_idle_dialogs")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-idle-dialogs"], "VoipProfileSip-MaxIdleDialogs"); ok {
			if err = d.Set("max_idle_dialogs", vv); err != nil {
				return fmt.Errorf("Error reading max_idle_dialogs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_idle_dialogs: %v", err)
		}
	}

	if err = d.Set("max_line_length", flattenVoipProfileSipMaxLineLength2edl(o["max-line-length"], d, "max_line_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-line-length"], "VoipProfileSip-MaxLineLength"); ok {
			if err = d.Set("max_line_length", vv); err != nil {
				return fmt.Errorf("Error reading max_line_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_line_length: %v", err)
		}
	}

	if err = d.Set("message_rate", flattenVoipProfileSipMessageRate2edl(o["message-rate"], d, "message_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-rate"], "VoipProfileSip-MessageRate"); ok {
			if err = d.Set("message_rate", vv); err != nil {
				return fmt.Errorf("Error reading message_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_rate: %v", err)
		}
	}

	if err = d.Set("message_rate_track", flattenVoipProfileSipMessageRateTrack2edl(o["message-rate-track"], d, "message_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-rate-track"], "VoipProfileSip-MessageRateTrack"); ok {
			if err = d.Set("message_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading message_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_rate_track: %v", err)
		}
	}

	if err = d.Set("nat_port_range", flattenVoipProfileSipNatPortRange2edl(o["nat-port-range"], d, "nat_port_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-port-range"], "VoipProfileSip-NatPortRange"); ok {
			if err = d.Set("nat_port_range", vv); err != nil {
				return fmt.Errorf("Error reading nat_port_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_port_range: %v", err)
		}
	}

	if err = d.Set("nat_trace", flattenVoipProfileSipNatTrace2edl(o["nat-trace"], d, "nat_trace")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-trace"], "VoipProfileSip-NatTrace"); ok {
			if err = d.Set("nat_trace", vv); err != nil {
				return fmt.Errorf("Error reading nat_trace: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_trace: %v", err)
		}
	}

	if err = d.Set("no_sdp_fixup", flattenVoipProfileSipNoSdpFixup2edl(o["no-sdp-fixup"], d, "no_sdp_fixup")); err != nil {
		if vv, ok := fortiAPIPatch(o["no-sdp-fixup"], "VoipProfileSip-NoSdpFixup"); ok {
			if err = d.Set("no_sdp_fixup", vv); err != nil {
				return fmt.Errorf("Error reading no_sdp_fixup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading no_sdp_fixup: %v", err)
		}
	}

	if err = d.Set("notify_rate", flattenVoipProfileSipNotifyRate2edl(o["notify-rate"], d, "notify_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["notify-rate"], "VoipProfileSip-NotifyRate"); ok {
			if err = d.Set("notify_rate", vv); err != nil {
				return fmt.Errorf("Error reading notify_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notify_rate: %v", err)
		}
	}

	if err = d.Set("notify_rate_track", flattenVoipProfileSipNotifyRateTrack2edl(o["notify-rate-track"], d, "notify_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["notify-rate-track"], "VoipProfileSip-NotifyRateTrack"); ok {
			if err = d.Set("notify_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading notify_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notify_rate_track: %v", err)
		}
	}

	if err = d.Set("open_contact_pinhole", flattenVoipProfileSipOpenContactPinhole2edl(o["open-contact-pinhole"], d, "open_contact_pinhole")); err != nil {
		if vv, ok := fortiAPIPatch(o["open-contact-pinhole"], "VoipProfileSip-OpenContactPinhole"); ok {
			if err = d.Set("open_contact_pinhole", vv); err != nil {
				return fmt.Errorf("Error reading open_contact_pinhole: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading open_contact_pinhole: %v", err)
		}
	}

	if err = d.Set("open_record_route_pinhole", flattenVoipProfileSipOpenRecordRoutePinhole2edl(o["open-record-route-pinhole"], d, "open_record_route_pinhole")); err != nil {
		if vv, ok := fortiAPIPatch(o["open-record-route-pinhole"], "VoipProfileSip-OpenRecordRoutePinhole"); ok {
			if err = d.Set("open_record_route_pinhole", vv); err != nil {
				return fmt.Errorf("Error reading open_record_route_pinhole: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading open_record_route_pinhole: %v", err)
		}
	}

	if err = d.Set("open_register_pinhole", flattenVoipProfileSipOpenRegisterPinhole2edl(o["open-register-pinhole"], d, "open_register_pinhole")); err != nil {
		if vv, ok := fortiAPIPatch(o["open-register-pinhole"], "VoipProfileSip-OpenRegisterPinhole"); ok {
			if err = d.Set("open_register_pinhole", vv); err != nil {
				return fmt.Errorf("Error reading open_register_pinhole: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading open_register_pinhole: %v", err)
		}
	}

	if err = d.Set("open_via_pinhole", flattenVoipProfileSipOpenViaPinhole2edl(o["open-via-pinhole"], d, "open_via_pinhole")); err != nil {
		if vv, ok := fortiAPIPatch(o["open-via-pinhole"], "VoipProfileSip-OpenViaPinhole"); ok {
			if err = d.Set("open_via_pinhole", vv); err != nil {
				return fmt.Errorf("Error reading open_via_pinhole: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading open_via_pinhole: %v", err)
		}
	}

	if err = d.Set("options_rate", flattenVoipProfileSipOptionsRate2edl(o["options-rate"], d, "options_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["options-rate"], "VoipProfileSip-OptionsRate"); ok {
			if err = d.Set("options_rate", vv); err != nil {
				return fmt.Errorf("Error reading options_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options_rate: %v", err)
		}
	}

	if err = d.Set("options_rate_track", flattenVoipProfileSipOptionsRateTrack2edl(o["options-rate-track"], d, "options_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["options-rate-track"], "VoipProfileSip-OptionsRateTrack"); ok {
			if err = d.Set("options_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading options_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options_rate_track: %v", err)
		}
	}

	if err = d.Set("prack_rate", flattenVoipProfileSipPrackRate2edl(o["prack-rate"], d, "prack_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["prack-rate"], "VoipProfileSip-PrackRate"); ok {
			if err = d.Set("prack_rate", vv); err != nil {
				return fmt.Errorf("Error reading prack_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prack_rate: %v", err)
		}
	}

	if err = d.Set("prack_rate_track", flattenVoipProfileSipPrackRateTrack2edl(o["prack-rate-track"], d, "prack_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["prack-rate-track"], "VoipProfileSip-PrackRateTrack"); ok {
			if err = d.Set("prack_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading prack_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prack_rate_track: %v", err)
		}
	}

	if err = d.Set("preserve_override", flattenVoipProfileSipPreserveOverride2edl(o["preserve-override"], d, "preserve_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["preserve-override"], "VoipProfileSip-PreserveOverride"); ok {
			if err = d.Set("preserve_override", vv); err != nil {
				return fmt.Errorf("Error reading preserve_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preserve_override: %v", err)
		}
	}

	if err = d.Set("provisional_invite_expiry_time", flattenVoipProfileSipProvisionalInviteExpiryTime2edl(o["provisional-invite-expiry-time"], d, "provisional_invite_expiry_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["provisional-invite-expiry-time"], "VoipProfileSip-ProvisionalInviteExpiryTime"); ok {
			if err = d.Set("provisional_invite_expiry_time", vv); err != nil {
				return fmt.Errorf("Error reading provisional_invite_expiry_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading provisional_invite_expiry_time: %v", err)
		}
	}

	if err = d.Set("publish_rate", flattenVoipProfileSipPublishRate2edl(o["publish-rate"], d, "publish_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["publish-rate"], "VoipProfileSip-PublishRate"); ok {
			if err = d.Set("publish_rate", vv); err != nil {
				return fmt.Errorf("Error reading publish_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading publish_rate: %v", err)
		}
	}

	if err = d.Set("publish_rate_track", flattenVoipProfileSipPublishRateTrack2edl(o["publish-rate-track"], d, "publish_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["publish-rate-track"], "VoipProfileSip-PublishRateTrack"); ok {
			if err = d.Set("publish_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading publish_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading publish_rate_track: %v", err)
		}
	}

	if err = d.Set("refer_rate", flattenVoipProfileSipReferRate2edl(o["refer-rate"], d, "refer_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["refer-rate"], "VoipProfileSip-ReferRate"); ok {
			if err = d.Set("refer_rate", vv); err != nil {
				return fmt.Errorf("Error reading refer_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading refer_rate: %v", err)
		}
	}

	if err = d.Set("refer_rate_track", flattenVoipProfileSipReferRateTrack2edl(o["refer-rate-track"], d, "refer_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["refer-rate-track"], "VoipProfileSip-ReferRateTrack"); ok {
			if err = d.Set("refer_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading refer_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading refer_rate_track: %v", err)
		}
	}

	if err = d.Set("register_contact_trace", flattenVoipProfileSipRegisterContactTrace2edl(o["register-contact-trace"], d, "register_contact_trace")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-contact-trace"], "VoipProfileSip-RegisterContactTrace"); ok {
			if err = d.Set("register_contact_trace", vv); err != nil {
				return fmt.Errorf("Error reading register_contact_trace: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_contact_trace: %v", err)
		}
	}

	if err = d.Set("register_rate", flattenVoipProfileSipRegisterRate2edl(o["register-rate"], d, "register_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-rate"], "VoipProfileSip-RegisterRate"); ok {
			if err = d.Set("register_rate", vv); err != nil {
				return fmt.Errorf("Error reading register_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_rate: %v", err)
		}
	}

	if err = d.Set("register_rate_track", flattenVoipProfileSipRegisterRateTrack2edl(o["register-rate-track"], d, "register_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["register-rate-track"], "VoipProfileSip-RegisterRateTrack"); ok {
			if err = d.Set("register_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading register_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading register_rate_track: %v", err)
		}
	}

	if err = d.Set("rfc2543_branch", flattenVoipProfileSipRfc2543Branch2edl(o["rfc2543-branch"], d, "rfc2543_branch")); err != nil {
		if vv, ok := fortiAPIPatch(o["rfc2543-branch"], "VoipProfileSip-Rfc2543Branch"); ok {
			if err = d.Set("rfc2543_branch", vv); err != nil {
				return fmt.Errorf("Error reading rfc2543_branch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rfc2543_branch: %v", err)
		}
	}

	if err = d.Set("rtp", flattenVoipProfileSipRtp2edl(o["rtp"], d, "rtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["rtp"], "VoipProfileSip-Rtp"); ok {
			if err = d.Set("rtp", vv); err != nil {
				return fmt.Errorf("Error reading rtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rtp: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenVoipProfileSipSslAlgorithm2edl(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-algorithm"], "VoipProfileSip-SslAlgorithm"); ok {
			if err = d.Set("ssl_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssl_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_auth_client", flattenVoipProfileSipSslAuthClient2edl(o["ssl-auth-client"], d, "ssl_auth_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-auth-client"], "VoipProfileSip-SslAuthClient"); ok {
			if err = d.Set("ssl_auth_client", vv); err != nil {
				return fmt.Errorf("Error reading ssl_auth_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_auth_client: %v", err)
		}
	}

	if err = d.Set("ssl_auth_server", flattenVoipProfileSipSslAuthServer2edl(o["ssl-auth-server"], d, "ssl_auth_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-auth-server"], "VoipProfileSip-SslAuthServer"); ok {
			if err = d.Set("ssl_auth_server", vv); err != nil {
				return fmt.Errorf("Error reading ssl_auth_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_auth_server: %v", err)
		}
	}

	if err = d.Set("ssl_client_certificate", flattenVoipProfileSipSslClientCertificate2edl(o["ssl-client-certificate"], d, "ssl_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-certificate"], "VoipProfileSip-SslClientCertificate"); ok {
			if err = d.Set("ssl_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_certificate: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenVoipProfileSipSslClientRenegotiation2edl(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-client-renegotiation"], "VoipProfileSip-SslClientRenegotiation"); ok {
			if err = d.Set("ssl_client_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenVoipProfileSipSslMaxVersion2edl(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "VoipProfileSip-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenVoipProfileSipSslMinVersion2edl(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-version"], "VoipProfileSip-SslMinVersion"); ok {
			if err = d.Set("ssl_min_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_mode", flattenVoipProfileSipSslMode2edl(o["ssl-mode"], d, "ssl_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-mode"], "VoipProfileSip-SslMode"); ok {
			if err = d.Set("ssl_mode", vv); err != nil {
				return fmt.Errorf("Error reading ssl_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_mode: %v", err)
		}
	}

	if err = d.Set("ssl_pfs", flattenVoipProfileSipSslPfs2edl(o["ssl-pfs"], d, "ssl_pfs")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-pfs"], "VoipProfileSip-SslPfs"); ok {
			if err = d.Set("ssl_pfs", vv); err != nil {
				return fmt.Errorf("Error reading ssl_pfs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_pfs: %v", err)
		}
	}

	if err = d.Set("ssl_send_empty_frags", flattenVoipProfileSipSslSendEmptyFrags2edl(o["ssl-send-empty-frags"], d, "ssl_send_empty_frags")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-send-empty-frags"], "VoipProfileSip-SslSendEmptyFrags"); ok {
			if err = d.Set("ssl_send_empty_frags", vv); err != nil {
				return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
		}
	}

	if err = d.Set("ssl_server_certificate", flattenVoipProfileSipSslServerCertificate2edl(o["ssl-server-certificate"], d, "ssl_server_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-certificate"], "VoipProfileSip-SslServerCertificate"); ok {
			if err = d.Set("ssl_server_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_certificate: %v", err)
		}
	}

	if err = d.Set("status", flattenVoipProfileSipStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "VoipProfileSip-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("strict_register", flattenVoipProfileSipStrictRegister2edl(o["strict-register"], d, "strict_register")); err != nil {
		if vv, ok := fortiAPIPatch(o["strict-register"], "VoipProfileSip-StrictRegister"); ok {
			if err = d.Set("strict_register", vv); err != nil {
				return fmt.Errorf("Error reading strict_register: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strict_register: %v", err)
		}
	}

	if err = d.Set("subscribe_rate", flattenVoipProfileSipSubscribeRate2edl(o["subscribe-rate"], d, "subscribe_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["subscribe-rate"], "VoipProfileSip-SubscribeRate"); ok {
			if err = d.Set("subscribe_rate", vv); err != nil {
				return fmt.Errorf("Error reading subscribe_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subscribe_rate: %v", err)
		}
	}

	if err = d.Set("subscribe_rate_track", flattenVoipProfileSipSubscribeRateTrack2edl(o["subscribe-rate-track"], d, "subscribe_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["subscribe-rate-track"], "VoipProfileSip-SubscribeRateTrack"); ok {
			if err = d.Set("subscribe_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading subscribe_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subscribe_rate_track: %v", err)
		}
	}

	if err = d.Set("unknown_header", flattenVoipProfileSipUnknownHeader2edl(o["unknown-header"], d, "unknown_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-header"], "VoipProfileSip-UnknownHeader"); ok {
			if err = d.Set("unknown_header", vv); err != nil {
				return fmt.Errorf("Error reading unknown_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_header: %v", err)
		}
	}

	if err = d.Set("update_rate", flattenVoipProfileSipUpdateRate2edl(o["update-rate"], d, "update_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-rate"], "VoipProfileSip-UpdateRate"); ok {
			if err = d.Set("update_rate", vv); err != nil {
				return fmt.Errorf("Error reading update_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_rate: %v", err)
		}
	}

	if err = d.Set("update_rate_track", flattenVoipProfileSipUpdateRateTrack2edl(o["update-rate-track"], d, "update_rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-rate-track"], "VoipProfileSip-UpdateRateTrack"); ok {
			if err = d.Set("update_rate_track", vv); err != nil {
				return fmt.Errorf("Error reading update_rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_rate_track: %v", err)
		}
	}

	return nil
}

func flattenVoipProfileSipFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVoipProfileSipAckRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipAckRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockAck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockBye2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockCancel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockGeoRedOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockInfo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockInvite2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockLongLines2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockMessage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockNotify2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockPrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockPublish2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockRefer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockRegister2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockSubscribe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockUnknown2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockUpdate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipByeRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipByeRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipCallIdRegex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipCallKeepalive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipCancelRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipCancelRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipContactFixup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipContentTypeRegex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipHntRestrictSourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipHostedNatTraversal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipInfoRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipInfoRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipInviteRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipInviteRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipIpsRtp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipLogCallSummary2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipLogViolations2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderAllow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderCallId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderContact2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderContentLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderContentType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderCseq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderExpires2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderFrom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderMaxForwards2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderNoProxyRequire2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderNoRequire2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderPAssertedIdentity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderRack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderRecordRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderRseq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpA2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpB2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpC2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpI2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpK2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpM2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpO2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpR2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpS2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpT2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpZ2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderTo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderVia2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedRequestLine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMaxBodyLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMaxDialogs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMaxIdleDialogs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMaxLineLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMessageRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMessageRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipNatPortRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipNatTrace2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipNoSdpFixup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipNotifyRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipNotifyRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipOpenContactPinhole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipOpenRecordRoutePinhole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipOpenRegisterPinhole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipOpenViaPinhole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipOptionsRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipOptionsRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipPrackRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipPrackRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipPreserveOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipProvisionalInviteExpiryTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipPublishRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipPublishRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipReferRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipReferRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipRegisterContactTrace2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipRegisterRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipRegisterRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipRfc2543Branch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipRtp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslAlgorithm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslAuthClient2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVoipProfileSipSslAuthServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVoipProfileSipSslClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVoipProfileSipSslClientRenegotiation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslMaxVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslMinVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslPfs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslSendEmptyFrags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslServerCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVoipProfileSipStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipStrictRegister2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSubscribeRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSubscribeRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipUnknownHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipUpdateRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipUpdateRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVoipProfileSip(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ack_rate"); ok || d.HasChange("ack_rate") {
		t, err := expandVoipProfileSipAckRate2edl(d, v, "ack_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ack-rate"] = t
		}
	}

	if v, ok := d.GetOk("ack_rate_track"); ok || d.HasChange("ack_rate_track") {
		t, err := expandVoipProfileSipAckRateTrack2edl(d, v, "ack_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ack-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("block_ack"); ok || d.HasChange("block_ack") {
		t, err := expandVoipProfileSipBlockAck2edl(d, v, "block_ack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-ack"] = t
		}
	}

	if v, ok := d.GetOk("block_bye"); ok || d.HasChange("block_bye") {
		t, err := expandVoipProfileSipBlockBye2edl(d, v, "block_bye")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-bye"] = t
		}
	}

	if v, ok := d.GetOk("block_cancel"); ok || d.HasChange("block_cancel") {
		t, err := expandVoipProfileSipBlockCancel2edl(d, v, "block_cancel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-cancel"] = t
		}
	}

	if v, ok := d.GetOk("block_geo_red_options"); ok || d.HasChange("block_geo_red_options") {
		t, err := expandVoipProfileSipBlockGeoRedOptions2edl(d, v, "block_geo_red_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-geo-red-options"] = t
		}
	}

	if v, ok := d.GetOk("block_info"); ok || d.HasChange("block_info") {
		t, err := expandVoipProfileSipBlockInfo2edl(d, v, "block_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-info"] = t
		}
	}

	if v, ok := d.GetOk("block_invite"); ok || d.HasChange("block_invite") {
		t, err := expandVoipProfileSipBlockInvite2edl(d, v, "block_invite")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-invite"] = t
		}
	}

	if v, ok := d.GetOk("block_long_lines"); ok || d.HasChange("block_long_lines") {
		t, err := expandVoipProfileSipBlockLongLines2edl(d, v, "block_long_lines")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-long-lines"] = t
		}
	}

	if v, ok := d.GetOk("block_message"); ok || d.HasChange("block_message") {
		t, err := expandVoipProfileSipBlockMessage2edl(d, v, "block_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-message"] = t
		}
	}

	if v, ok := d.GetOk("block_notify"); ok || d.HasChange("block_notify") {
		t, err := expandVoipProfileSipBlockNotify2edl(d, v, "block_notify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-notify"] = t
		}
	}

	if v, ok := d.GetOk("block_options"); ok || d.HasChange("block_options") {
		t, err := expandVoipProfileSipBlockOptions2edl(d, v, "block_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-options"] = t
		}
	}

	if v, ok := d.GetOk("block_prack"); ok || d.HasChange("block_prack") {
		t, err := expandVoipProfileSipBlockPrack2edl(d, v, "block_prack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-prack"] = t
		}
	}

	if v, ok := d.GetOk("block_publish"); ok || d.HasChange("block_publish") {
		t, err := expandVoipProfileSipBlockPublish2edl(d, v, "block_publish")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-publish"] = t
		}
	}

	if v, ok := d.GetOk("block_refer"); ok || d.HasChange("block_refer") {
		t, err := expandVoipProfileSipBlockRefer2edl(d, v, "block_refer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-refer"] = t
		}
	}

	if v, ok := d.GetOk("block_register"); ok || d.HasChange("block_register") {
		t, err := expandVoipProfileSipBlockRegister2edl(d, v, "block_register")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-register"] = t
		}
	}

	if v, ok := d.GetOk("block_subscribe"); ok || d.HasChange("block_subscribe") {
		t, err := expandVoipProfileSipBlockSubscribe2edl(d, v, "block_subscribe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-subscribe"] = t
		}
	}

	if v, ok := d.GetOk("block_unknown"); ok || d.HasChange("block_unknown") {
		t, err := expandVoipProfileSipBlockUnknown2edl(d, v, "block_unknown")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-unknown"] = t
		}
	}

	if v, ok := d.GetOk("block_update"); ok || d.HasChange("block_update") {
		t, err := expandVoipProfileSipBlockUpdate2edl(d, v, "block_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-update"] = t
		}
	}

	if v, ok := d.GetOk("bye_rate"); ok || d.HasChange("bye_rate") {
		t, err := expandVoipProfileSipByeRate2edl(d, v, "bye_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bye-rate"] = t
		}
	}

	if v, ok := d.GetOk("bye_rate_track"); ok || d.HasChange("bye_rate_track") {
		t, err := expandVoipProfileSipByeRateTrack2edl(d, v, "bye_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bye-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("call_id_regex"); ok || d.HasChange("call_id_regex") {
		t, err := expandVoipProfileSipCallIdRegex2edl(d, v, "call_id_regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-id-regex"] = t
		}
	}

	if v, ok := d.GetOk("call_keepalive"); ok || d.HasChange("call_keepalive") {
		t, err := expandVoipProfileSipCallKeepalive2edl(d, v, "call_keepalive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-keepalive"] = t
		}
	}

	if v, ok := d.GetOk("cancel_rate"); ok || d.HasChange("cancel_rate") {
		t, err := expandVoipProfileSipCancelRate2edl(d, v, "cancel_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cancel-rate"] = t
		}
	}

	if v, ok := d.GetOk("cancel_rate_track"); ok || d.HasChange("cancel_rate_track") {
		t, err := expandVoipProfileSipCancelRateTrack2edl(d, v, "cancel_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cancel-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("contact_fixup"); ok || d.HasChange("contact_fixup") {
		t, err := expandVoipProfileSipContactFixup2edl(d, v, "contact_fixup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contact-fixup"] = t
		}
	}

	if v, ok := d.GetOk("content_type_regex"); ok || d.HasChange("content_type_regex") {
		t, err := expandVoipProfileSipContentTypeRegex2edl(d, v, "content_type_regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-type-regex"] = t
		}
	}

	if v, ok := d.GetOk("hnt_restrict_source_ip"); ok || d.HasChange("hnt_restrict_source_ip") {
		t, err := expandVoipProfileSipHntRestrictSourceIp2edl(d, v, "hnt_restrict_source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hnt-restrict-source-ip"] = t
		}
	}

	if v, ok := d.GetOk("hosted_nat_traversal"); ok || d.HasChange("hosted_nat_traversal") {
		t, err := expandVoipProfileSipHostedNatTraversal2edl(d, v, "hosted_nat_traversal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosted-nat-traversal"] = t
		}
	}

	if v, ok := d.GetOk("info_rate"); ok || d.HasChange("info_rate") {
		t, err := expandVoipProfileSipInfoRate2edl(d, v, "info_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["info-rate"] = t
		}
	}

	if v, ok := d.GetOk("info_rate_track"); ok || d.HasChange("info_rate_track") {
		t, err := expandVoipProfileSipInfoRateTrack2edl(d, v, "info_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["info-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("invite_rate"); ok || d.HasChange("invite_rate") {
		t, err := expandVoipProfileSipInviteRate2edl(d, v, "invite_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invite-rate"] = t
		}
	}

	if v, ok := d.GetOk("invite_rate_track"); ok || d.HasChange("invite_rate_track") {
		t, err := expandVoipProfileSipInviteRateTrack2edl(d, v, "invite_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invite-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("ips_rtp"); ok || d.HasChange("ips_rtp") {
		t, err := expandVoipProfileSipIpsRtp2edl(d, v, "ips_rtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-rtp"] = t
		}
	}

	if v, ok := d.GetOk("log_call_summary"); ok || d.HasChange("log_call_summary") {
		t, err := expandVoipProfileSipLogCallSummary2edl(d, v, "log_call_summary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-call-summary"] = t
		}
	}

	if v, ok := d.GetOk("log_violations"); ok || d.HasChange("log_violations") {
		t, err := expandVoipProfileSipLogViolations2edl(d, v, "log_violations")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-violations"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_allow"); ok || d.HasChange("malformed_header_allow") {
		t, err := expandVoipProfileSipMalformedHeaderAllow2edl(d, v, "malformed_header_allow")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-allow"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_call_id"); ok || d.HasChange("malformed_header_call_id") {
		t, err := expandVoipProfileSipMalformedHeaderCallId2edl(d, v, "malformed_header_call_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-call-id"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_contact"); ok || d.HasChange("malformed_header_contact") {
		t, err := expandVoipProfileSipMalformedHeaderContact2edl(d, v, "malformed_header_contact")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-contact"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_content_length"); ok || d.HasChange("malformed_header_content_length") {
		t, err := expandVoipProfileSipMalformedHeaderContentLength2edl(d, v, "malformed_header_content_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-content-length"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_content_type"); ok || d.HasChange("malformed_header_content_type") {
		t, err := expandVoipProfileSipMalformedHeaderContentType2edl(d, v, "malformed_header_content_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-content-type"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_cseq"); ok || d.HasChange("malformed_header_cseq") {
		t, err := expandVoipProfileSipMalformedHeaderCseq2edl(d, v, "malformed_header_cseq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-cseq"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_expires"); ok || d.HasChange("malformed_header_expires") {
		t, err := expandVoipProfileSipMalformedHeaderExpires2edl(d, v, "malformed_header_expires")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-expires"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_from"); ok || d.HasChange("malformed_header_from") {
		t, err := expandVoipProfileSipMalformedHeaderFrom2edl(d, v, "malformed_header_from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-from"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_max_forwards"); ok || d.HasChange("malformed_header_max_forwards") {
		t, err := expandVoipProfileSipMalformedHeaderMaxForwards2edl(d, v, "malformed_header_max_forwards")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-max-forwards"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_no_proxy_require"); ok || d.HasChange("malformed_header_no_proxy_require") {
		t, err := expandVoipProfileSipMalformedHeaderNoProxyRequire2edl(d, v, "malformed_header_no_proxy_require")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-no-proxy-require"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_no_require"); ok || d.HasChange("malformed_header_no_require") {
		t, err := expandVoipProfileSipMalformedHeaderNoRequire2edl(d, v, "malformed_header_no_require")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-no-require"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_p_asserted_identity"); ok || d.HasChange("malformed_header_p_asserted_identity") {
		t, err := expandVoipProfileSipMalformedHeaderPAssertedIdentity2edl(d, v, "malformed_header_p_asserted_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-p-asserted-identity"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_rack"); ok || d.HasChange("malformed_header_rack") {
		t, err := expandVoipProfileSipMalformedHeaderRack2edl(d, v, "malformed_header_rack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-rack"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_record_route"); ok || d.HasChange("malformed_header_record_route") {
		t, err := expandVoipProfileSipMalformedHeaderRecordRoute2edl(d, v, "malformed_header_record_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-record-route"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_route"); ok || d.HasChange("malformed_header_route") {
		t, err := expandVoipProfileSipMalformedHeaderRoute2edl(d, v, "malformed_header_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-route"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_rseq"); ok || d.HasChange("malformed_header_rseq") {
		t, err := expandVoipProfileSipMalformedHeaderRseq2edl(d, v, "malformed_header_rseq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-rseq"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_a"); ok || d.HasChange("malformed_header_sdp_a") {
		t, err := expandVoipProfileSipMalformedHeaderSdpA2edl(d, v, "malformed_header_sdp_a")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-a"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_b"); ok || d.HasChange("malformed_header_sdp_b") {
		t, err := expandVoipProfileSipMalformedHeaderSdpB2edl(d, v, "malformed_header_sdp_b")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-b"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_c"); ok || d.HasChange("malformed_header_sdp_c") {
		t, err := expandVoipProfileSipMalformedHeaderSdpC2edl(d, v, "malformed_header_sdp_c")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-c"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_i"); ok || d.HasChange("malformed_header_sdp_i") {
		t, err := expandVoipProfileSipMalformedHeaderSdpI2edl(d, v, "malformed_header_sdp_i")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-i"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_k"); ok || d.HasChange("malformed_header_sdp_k") {
		t, err := expandVoipProfileSipMalformedHeaderSdpK2edl(d, v, "malformed_header_sdp_k")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-k"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_m"); ok || d.HasChange("malformed_header_sdp_m") {
		t, err := expandVoipProfileSipMalformedHeaderSdpM2edl(d, v, "malformed_header_sdp_m")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-m"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_o"); ok || d.HasChange("malformed_header_sdp_o") {
		t, err := expandVoipProfileSipMalformedHeaderSdpO2edl(d, v, "malformed_header_sdp_o")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-o"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_r"); ok || d.HasChange("malformed_header_sdp_r") {
		t, err := expandVoipProfileSipMalformedHeaderSdpR2edl(d, v, "malformed_header_sdp_r")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-r"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_s"); ok || d.HasChange("malformed_header_sdp_s") {
		t, err := expandVoipProfileSipMalformedHeaderSdpS2edl(d, v, "malformed_header_sdp_s")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-s"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_t"); ok || d.HasChange("malformed_header_sdp_t") {
		t, err := expandVoipProfileSipMalformedHeaderSdpT2edl(d, v, "malformed_header_sdp_t")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-t"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_v"); ok || d.HasChange("malformed_header_sdp_v") {
		t, err := expandVoipProfileSipMalformedHeaderSdpV2edl(d, v, "malformed_header_sdp_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-v"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_sdp_z"); ok || d.HasChange("malformed_header_sdp_z") {
		t, err := expandVoipProfileSipMalformedHeaderSdpZ2edl(d, v, "malformed_header_sdp_z")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-sdp-z"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_to"); ok || d.HasChange("malformed_header_to") {
		t, err := expandVoipProfileSipMalformedHeaderTo2edl(d, v, "malformed_header_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-to"] = t
		}
	}

	if v, ok := d.GetOk("malformed_header_via"); ok || d.HasChange("malformed_header_via") {
		t, err := expandVoipProfileSipMalformedHeaderVia2edl(d, v, "malformed_header_via")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-header-via"] = t
		}
	}

	if v, ok := d.GetOk("malformed_request_line"); ok || d.HasChange("malformed_request_line") {
		t, err := expandVoipProfileSipMalformedRequestLine2edl(d, v, "malformed_request_line")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-request-line"] = t
		}
	}

	if v, ok := d.GetOk("max_body_length"); ok || d.HasChange("max_body_length") {
		t, err := expandVoipProfileSipMaxBodyLength2edl(d, v, "max_body_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-body-length"] = t
		}
	}

	if v, ok := d.GetOk("max_dialogs"); ok || d.HasChange("max_dialogs") {
		t, err := expandVoipProfileSipMaxDialogs2edl(d, v, "max_dialogs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-dialogs"] = t
		}
	}

	if v, ok := d.GetOk("max_idle_dialogs"); ok || d.HasChange("max_idle_dialogs") {
		t, err := expandVoipProfileSipMaxIdleDialogs2edl(d, v, "max_idle_dialogs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-idle-dialogs"] = t
		}
	}

	if v, ok := d.GetOk("max_line_length"); ok || d.HasChange("max_line_length") {
		t, err := expandVoipProfileSipMaxLineLength2edl(d, v, "max_line_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-line-length"] = t
		}
	}

	if v, ok := d.GetOk("message_rate"); ok || d.HasChange("message_rate") {
		t, err := expandVoipProfileSipMessageRate2edl(d, v, "message_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-rate"] = t
		}
	}

	if v, ok := d.GetOk("message_rate_track"); ok || d.HasChange("message_rate_track") {
		t, err := expandVoipProfileSipMessageRateTrack2edl(d, v, "message_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("nat_port_range"); ok || d.HasChange("nat_port_range") {
		t, err := expandVoipProfileSipNatPortRange2edl(d, v, "nat_port_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-port-range"] = t
		}
	}

	if v, ok := d.GetOk("nat_trace"); ok || d.HasChange("nat_trace") {
		t, err := expandVoipProfileSipNatTrace2edl(d, v, "nat_trace")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-trace"] = t
		}
	}

	if v, ok := d.GetOk("no_sdp_fixup"); ok || d.HasChange("no_sdp_fixup") {
		t, err := expandVoipProfileSipNoSdpFixup2edl(d, v, "no_sdp_fixup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["no-sdp-fixup"] = t
		}
	}

	if v, ok := d.GetOk("notify_rate"); ok || d.HasChange("notify_rate") {
		t, err := expandVoipProfileSipNotifyRate2edl(d, v, "notify_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify-rate"] = t
		}
	}

	if v, ok := d.GetOk("notify_rate_track"); ok || d.HasChange("notify_rate_track") {
		t, err := expandVoipProfileSipNotifyRateTrack2edl(d, v, "notify_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("open_contact_pinhole"); ok || d.HasChange("open_contact_pinhole") {
		t, err := expandVoipProfileSipOpenContactPinhole2edl(d, v, "open_contact_pinhole")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["open-contact-pinhole"] = t
		}
	}

	if v, ok := d.GetOk("open_record_route_pinhole"); ok || d.HasChange("open_record_route_pinhole") {
		t, err := expandVoipProfileSipOpenRecordRoutePinhole2edl(d, v, "open_record_route_pinhole")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["open-record-route-pinhole"] = t
		}
	}

	if v, ok := d.GetOk("open_register_pinhole"); ok || d.HasChange("open_register_pinhole") {
		t, err := expandVoipProfileSipOpenRegisterPinhole2edl(d, v, "open_register_pinhole")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["open-register-pinhole"] = t
		}
	}

	if v, ok := d.GetOk("open_via_pinhole"); ok || d.HasChange("open_via_pinhole") {
		t, err := expandVoipProfileSipOpenViaPinhole2edl(d, v, "open_via_pinhole")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["open-via-pinhole"] = t
		}
	}

	if v, ok := d.GetOk("options_rate"); ok || d.HasChange("options_rate") {
		t, err := expandVoipProfileSipOptionsRate2edl(d, v, "options_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options-rate"] = t
		}
	}

	if v, ok := d.GetOk("options_rate_track"); ok || d.HasChange("options_rate_track") {
		t, err := expandVoipProfileSipOptionsRateTrack2edl(d, v, "options_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("prack_rate"); ok || d.HasChange("prack_rate") {
		t, err := expandVoipProfileSipPrackRate2edl(d, v, "prack_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prack-rate"] = t
		}
	}

	if v, ok := d.GetOk("prack_rate_track"); ok || d.HasChange("prack_rate_track") {
		t, err := expandVoipProfileSipPrackRateTrack2edl(d, v, "prack_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prack-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("preserve_override"); ok || d.HasChange("preserve_override") {
		t, err := expandVoipProfileSipPreserveOverride2edl(d, v, "preserve_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preserve-override"] = t
		}
	}

	if v, ok := d.GetOk("provisional_invite_expiry_time"); ok || d.HasChange("provisional_invite_expiry_time") {
		t, err := expandVoipProfileSipProvisionalInviteExpiryTime2edl(d, v, "provisional_invite_expiry_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["provisional-invite-expiry-time"] = t
		}
	}

	if v, ok := d.GetOk("publish_rate"); ok || d.HasChange("publish_rate") {
		t, err := expandVoipProfileSipPublishRate2edl(d, v, "publish_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["publish-rate"] = t
		}
	}

	if v, ok := d.GetOk("publish_rate_track"); ok || d.HasChange("publish_rate_track") {
		t, err := expandVoipProfileSipPublishRateTrack2edl(d, v, "publish_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["publish-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("refer_rate"); ok || d.HasChange("refer_rate") {
		t, err := expandVoipProfileSipReferRate2edl(d, v, "refer_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["refer-rate"] = t
		}
	}

	if v, ok := d.GetOk("refer_rate_track"); ok || d.HasChange("refer_rate_track") {
		t, err := expandVoipProfileSipReferRateTrack2edl(d, v, "refer_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["refer-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("register_contact_trace"); ok || d.HasChange("register_contact_trace") {
		t, err := expandVoipProfileSipRegisterContactTrace2edl(d, v, "register_contact_trace")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-contact-trace"] = t
		}
	}

	if v, ok := d.GetOk("register_rate"); ok || d.HasChange("register_rate") {
		t, err := expandVoipProfileSipRegisterRate2edl(d, v, "register_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-rate"] = t
		}
	}

	if v, ok := d.GetOk("register_rate_track"); ok || d.HasChange("register_rate_track") {
		t, err := expandVoipProfileSipRegisterRateTrack2edl(d, v, "register_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["register-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("rfc2543_branch"); ok || d.HasChange("rfc2543_branch") {
		t, err := expandVoipProfileSipRfc2543Branch2edl(d, v, "rfc2543_branch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rfc2543-branch"] = t
		}
	}

	if v, ok := d.GetOk("rtp"); ok || d.HasChange("rtp") {
		t, err := expandVoipProfileSipRtp2edl(d, v, "rtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtp"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok || d.HasChange("ssl_algorithm") {
		t, err := expandVoipProfileSipSslAlgorithm2edl(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_auth_client"); ok || d.HasChange("ssl_auth_client") {
		t, err := expandVoipProfileSipSslAuthClient2edl(d, v, "ssl_auth_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-auth-client"] = t
		}
	}

	if v, ok := d.GetOk("ssl_auth_server"); ok || d.HasChange("ssl_auth_server") {
		t, err := expandVoipProfileSipSslAuthServer2edl(d, v, "ssl_auth_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-auth-server"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_certificate"); ok || d.HasChange("ssl_client_certificate") {
		t, err := expandVoipProfileSipSslClientCertificate2edl(d, v, "ssl_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok || d.HasChange("ssl_client_renegotiation") {
		t, err := expandVoipProfileSipSslClientRenegotiation2edl(d, v, "ssl_client_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandVoipProfileSipSslMaxVersion2edl(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok || d.HasChange("ssl_min_version") {
		t, err := expandVoipProfileSipSslMinVersion2edl(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mode"); ok || d.HasChange("ssl_mode") {
		t, err := expandVoipProfileSipSslMode2edl(d, v, "ssl_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl_pfs"); ok || d.HasChange("ssl_pfs") {
		t, err := expandVoipProfileSipSslPfs2edl(d, v, "ssl_pfs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-pfs"] = t
		}
	}

	if v, ok := d.GetOk("ssl_send_empty_frags"); ok || d.HasChange("ssl_send_empty_frags") {
		t, err := expandVoipProfileSipSslSendEmptyFrags2edl(d, v, "ssl_send_empty_frags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-send-empty-frags"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_certificate"); ok || d.HasChange("ssl_server_certificate") {
		t, err := expandVoipProfileSipSslServerCertificate2edl(d, v, "ssl_server_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-certificate"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandVoipProfileSipStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("strict_register"); ok || d.HasChange("strict_register") {
		t, err := expandVoipProfileSipStrictRegister2edl(d, v, "strict_register")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-register"] = t
		}
	}

	if v, ok := d.GetOk("subscribe_rate"); ok || d.HasChange("subscribe_rate") {
		t, err := expandVoipProfileSipSubscribeRate2edl(d, v, "subscribe_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subscribe-rate"] = t
		}
	}

	if v, ok := d.GetOk("subscribe_rate_track"); ok || d.HasChange("subscribe_rate_track") {
		t, err := expandVoipProfileSipSubscribeRateTrack2edl(d, v, "subscribe_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subscribe-rate-track"] = t
		}
	}

	if v, ok := d.GetOk("unknown_header"); ok || d.HasChange("unknown_header") {
		t, err := expandVoipProfileSipUnknownHeader2edl(d, v, "unknown_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-header"] = t
		}
	}

	if v, ok := d.GetOk("update_rate"); ok || d.HasChange("update_rate") {
		t, err := expandVoipProfileSipUpdateRate2edl(d, v, "update_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-rate"] = t
		}
	}

	if v, ok := d.GetOk("update_rate_track"); ok || d.HasChange("update_rate_track") {
		t, err := expandVoipProfileSipUpdateRateTrack2edl(d, v, "update_rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-rate-track"] = t
		}
	}

	return &obj, nil
}
