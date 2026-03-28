// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> QUIC setting.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallVip6Quic() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVip6QuicUpdate,
		Read:   resourceFirewallVip6QuicRead,
		Update: resourceFirewallVip6QuicUpdate,
		Delete: resourceFirewallVip6QuicDelete,

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
			"vip6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ack_delay_exponent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"active_connection_id_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"active_migration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"grease_quic_bit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_ack_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_datagram_frame_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_udp_payload_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallVip6QuicUpdate(d *schema.ResourceData, m interface{}) error {
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
	vip6 := d.Get("vip6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip6"] = vip6

	obj, err := getObjectFirewallVip6Quic(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip6Quic resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallVip6Quic(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip6Quic resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallVip6Quic")

	return resourceFirewallVip6QuicRead(d, m)
}

func resourceFirewallVip6QuicDelete(d *schema.ResourceData, m interface{}) error {
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
	vip6 := d.Get("vip6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip6"] = vip6

	wsParams["adom"] = adomv

	err = c.DeleteFirewallVip6Quic(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVip6Quic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVip6QuicRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	vip6 := d.Get("vip6").(string)
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
	if vip6 == "" {
		vip6 = importOptionChecking(m.(*FortiClient).Cfg, "vip6")
		if vip6 == "" {
			return fmt.Errorf("Parameter vip6 is missing")
		}
		if err = d.Set("vip6", vip6); err != nil {
			return fmt.Errorf("Error set params vip6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip6"] = vip6

	o, err := c.ReadFirewallVip6Quic(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallVip6Quic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVip6Quic(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip6Quic resource from API: %v", err)
	}
	return nil
}

func flattenFirewallVip6QuicAckDelayExponent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicActiveConnectionIdLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicActiveMigration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicGreaseQuicBit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicMaxAckDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicMaxDatagramFrameSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicMaxIdleTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6QuicMaxUdpPayloadSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallVip6Quic(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ack_delay_exponent", flattenFirewallVip6QuicAckDelayExponent2edl(o["ack-delay-exponent"], d, "ack_delay_exponent")); err != nil {
		if vv, ok := fortiAPIPatch(o["ack-delay-exponent"], "FirewallVip6Quic-AckDelayExponent"); ok {
			if err = d.Set("ack_delay_exponent", vv); err != nil {
				return fmt.Errorf("Error reading ack_delay_exponent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ack_delay_exponent: %v", err)
		}
	}

	if err = d.Set("active_connection_id_limit", flattenFirewallVip6QuicActiveConnectionIdLimit2edl(o["active-connection-id-limit"], d, "active_connection_id_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-connection-id-limit"], "FirewallVip6Quic-ActiveConnectionIdLimit"); ok {
			if err = d.Set("active_connection_id_limit", vv); err != nil {
				return fmt.Errorf("Error reading active_connection_id_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_connection_id_limit: %v", err)
		}
	}

	if err = d.Set("active_migration", flattenFirewallVip6QuicActiveMigration2edl(o["active-migration"], d, "active_migration")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-migration"], "FirewallVip6Quic-ActiveMigration"); ok {
			if err = d.Set("active_migration", vv); err != nil {
				return fmt.Errorf("Error reading active_migration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_migration: %v", err)
		}
	}

	if err = d.Set("grease_quic_bit", flattenFirewallVip6QuicGreaseQuicBit2edl(o["grease-quic-bit"], d, "grease_quic_bit")); err != nil {
		if vv, ok := fortiAPIPatch(o["grease-quic-bit"], "FirewallVip6Quic-GreaseQuicBit"); ok {
			if err = d.Set("grease_quic_bit", vv); err != nil {
				return fmt.Errorf("Error reading grease_quic_bit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading grease_quic_bit: %v", err)
		}
	}

	if err = d.Set("max_ack_delay", flattenFirewallVip6QuicMaxAckDelay2edl(o["max-ack-delay"], d, "max_ack_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-ack-delay"], "FirewallVip6Quic-MaxAckDelay"); ok {
			if err = d.Set("max_ack_delay", vv); err != nil {
				return fmt.Errorf("Error reading max_ack_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_ack_delay: %v", err)
		}
	}

	if err = d.Set("max_datagram_frame_size", flattenFirewallVip6QuicMaxDatagramFrameSize2edl(o["max-datagram-frame-size"], d, "max_datagram_frame_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-datagram-frame-size"], "FirewallVip6Quic-MaxDatagramFrameSize"); ok {
			if err = d.Set("max_datagram_frame_size", vv); err != nil {
				return fmt.Errorf("Error reading max_datagram_frame_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_datagram_frame_size: %v", err)
		}
	}

	if err = d.Set("max_idle_timeout", flattenFirewallVip6QuicMaxIdleTimeout2edl(o["max-idle-timeout"], d, "max_idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-idle-timeout"], "FirewallVip6Quic-MaxIdleTimeout"); ok {
			if err = d.Set("max_idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading max_idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_idle_timeout: %v", err)
		}
	}

	if err = d.Set("max_udp_payload_size", flattenFirewallVip6QuicMaxUdpPayloadSize2edl(o["max-udp-payload-size"], d, "max_udp_payload_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-udp-payload-size"], "FirewallVip6Quic-MaxUdpPayloadSize"); ok {
			if err = d.Set("max_udp_payload_size", vv); err != nil {
				return fmt.Errorf("Error reading max_udp_payload_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_udp_payload_size: %v", err)
		}
	}

	return nil
}

func flattenFirewallVip6QuicFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallVip6QuicAckDelayExponent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicActiveConnectionIdLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicActiveMigration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicGreaseQuicBit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicMaxAckDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicMaxDatagramFrameSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicMaxIdleTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6QuicMaxUdpPayloadSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallVip6Quic(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ack_delay_exponent"); ok || d.HasChange("ack_delay_exponent") {
		t, err := expandFirewallVip6QuicAckDelayExponent2edl(d, v, "ack_delay_exponent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ack-delay-exponent"] = t
		}
	}

	if v, ok := d.GetOk("active_connection_id_limit"); ok || d.HasChange("active_connection_id_limit") {
		t, err := expandFirewallVip6QuicActiveConnectionIdLimit2edl(d, v, "active_connection_id_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-connection-id-limit"] = t
		}
	}

	if v, ok := d.GetOk("active_migration"); ok || d.HasChange("active_migration") {
		t, err := expandFirewallVip6QuicActiveMigration2edl(d, v, "active_migration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-migration"] = t
		}
	}

	if v, ok := d.GetOk("grease_quic_bit"); ok || d.HasChange("grease_quic_bit") {
		t, err := expandFirewallVip6QuicGreaseQuicBit2edl(d, v, "grease_quic_bit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["grease-quic-bit"] = t
		}
	}

	if v, ok := d.GetOk("max_ack_delay"); ok || d.HasChange("max_ack_delay") {
		t, err := expandFirewallVip6QuicMaxAckDelay2edl(d, v, "max_ack_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-ack-delay"] = t
		}
	}

	if v, ok := d.GetOk("max_datagram_frame_size"); ok || d.HasChange("max_datagram_frame_size") {
		t, err := expandFirewallVip6QuicMaxDatagramFrameSize2edl(d, v, "max_datagram_frame_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-datagram-frame-size"] = t
		}
	}

	if v, ok := d.GetOk("max_idle_timeout"); ok || d.HasChange("max_idle_timeout") {
		t, err := expandFirewallVip6QuicMaxIdleTimeout2edl(d, v, "max_idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("max_udp_payload_size"); ok || d.HasChange("max_udp_payload_size") {
		t, err := expandFirewallVip6QuicMaxUdpPayloadSize2edl(d, v, "max_udp_payload_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-udp-payload-size"] = t
		}
	}

	return &obj, nil
}
