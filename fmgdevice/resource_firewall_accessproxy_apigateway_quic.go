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

func resourceFirewallAccessProxyApiGatewayQuic() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxyApiGatewayQuicUpdate,
		Read:   resourceFirewallAccessProxyApiGatewayQuicRead,
		Update: resourceFirewallAccessProxyApiGatewayQuicUpdate,
		Delete: resourceFirewallAccessProxyApiGatewayQuicDelete,

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
			"access_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"api_gateway": &schema.Schema{
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

func resourceFirewallAccessProxyApiGatewayQuicUpdate(d *schema.ResourceData, m interface{}) error {
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
	access_proxy := d.Get("access_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy
	paradict["api_gateway"] = api_gateway

	obj, err := getObjectFirewallAccessProxyApiGatewayQuic(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxyApiGatewayQuic resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAccessProxyApiGatewayQuic(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxyApiGatewayQuic resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallAccessProxyApiGatewayQuic")

	return resourceFirewallAccessProxyApiGatewayQuicRead(d, m)
}

func resourceFirewallAccessProxyApiGatewayQuicDelete(d *schema.ResourceData, m interface{}) error {
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
	access_proxy := d.Get("access_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy
	paradict["api_gateway"] = api_gateway

	wsParams["adom"] = adomv

	err = c.DeleteFirewallAccessProxyApiGatewayQuic(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxyApiGatewayQuic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxyApiGatewayQuicRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	access_proxy := d.Get("access_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
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
	if access_proxy == "" {
		access_proxy = importOptionChecking(m.(*FortiClient).Cfg, "access_proxy")
		if access_proxy == "" {
			return fmt.Errorf("Parameter access_proxy is missing")
		}
		if err = d.Set("access_proxy", access_proxy); err != nil {
			return fmt.Errorf("Error set params access_proxy: %v", err)
		}
	}
	if api_gateway == "" {
		api_gateway = importOptionChecking(m.(*FortiClient).Cfg, "api_gateway")
		if api_gateway == "" {
			return fmt.Errorf("Parameter api_gateway is missing")
		}
		if err = d.Set("api_gateway", api_gateway); err != nil {
			return fmt.Errorf("Error set params api_gateway: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy
	paradict["api_gateway"] = api_gateway

	o, err := c.ReadFirewallAccessProxyApiGatewayQuic(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAccessProxyApiGatewayQuic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxyApiGatewayQuic(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxyApiGatewayQuic resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxyApiGatewayQuicAckDelayExponent3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicActiveConnectionIdLimit3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicActiveMigration3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicGreaseQuicBit3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicMaxAckDelay3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicMaxDatagramFrameSize3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicMaxIdleTimeout3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayQuicMaxUdpPayloadSize3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAccessProxyApiGatewayQuic(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ack_delay_exponent", flattenFirewallAccessProxyApiGatewayQuicAckDelayExponent3rdl(o["ack-delay-exponent"], d, "ack_delay_exponent")); err != nil {
		if vv, ok := fortiAPIPatch(o["ack-delay-exponent"], "FirewallAccessProxyApiGatewayQuic-AckDelayExponent"); ok {
			if err = d.Set("ack_delay_exponent", vv); err != nil {
				return fmt.Errorf("Error reading ack_delay_exponent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ack_delay_exponent: %v", err)
		}
	}

	if err = d.Set("active_connection_id_limit", flattenFirewallAccessProxyApiGatewayQuicActiveConnectionIdLimit3rdl(o["active-connection-id-limit"], d, "active_connection_id_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-connection-id-limit"], "FirewallAccessProxyApiGatewayQuic-ActiveConnectionIdLimit"); ok {
			if err = d.Set("active_connection_id_limit", vv); err != nil {
				return fmt.Errorf("Error reading active_connection_id_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_connection_id_limit: %v", err)
		}
	}

	if err = d.Set("active_migration", flattenFirewallAccessProxyApiGatewayQuicActiveMigration3rdl(o["active-migration"], d, "active_migration")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-migration"], "FirewallAccessProxyApiGatewayQuic-ActiveMigration"); ok {
			if err = d.Set("active_migration", vv); err != nil {
				return fmt.Errorf("Error reading active_migration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_migration: %v", err)
		}
	}

	if err = d.Set("grease_quic_bit", flattenFirewallAccessProxyApiGatewayQuicGreaseQuicBit3rdl(o["grease-quic-bit"], d, "grease_quic_bit")); err != nil {
		if vv, ok := fortiAPIPatch(o["grease-quic-bit"], "FirewallAccessProxyApiGatewayQuic-GreaseQuicBit"); ok {
			if err = d.Set("grease_quic_bit", vv); err != nil {
				return fmt.Errorf("Error reading grease_quic_bit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading grease_quic_bit: %v", err)
		}
	}

	if err = d.Set("max_ack_delay", flattenFirewallAccessProxyApiGatewayQuicMaxAckDelay3rdl(o["max-ack-delay"], d, "max_ack_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-ack-delay"], "FirewallAccessProxyApiGatewayQuic-MaxAckDelay"); ok {
			if err = d.Set("max_ack_delay", vv); err != nil {
				return fmt.Errorf("Error reading max_ack_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_ack_delay: %v", err)
		}
	}

	if err = d.Set("max_datagram_frame_size", flattenFirewallAccessProxyApiGatewayQuicMaxDatagramFrameSize3rdl(o["max-datagram-frame-size"], d, "max_datagram_frame_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-datagram-frame-size"], "FirewallAccessProxyApiGatewayQuic-MaxDatagramFrameSize"); ok {
			if err = d.Set("max_datagram_frame_size", vv); err != nil {
				return fmt.Errorf("Error reading max_datagram_frame_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_datagram_frame_size: %v", err)
		}
	}

	if err = d.Set("max_idle_timeout", flattenFirewallAccessProxyApiGatewayQuicMaxIdleTimeout3rdl(o["max-idle-timeout"], d, "max_idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-idle-timeout"], "FirewallAccessProxyApiGatewayQuic-MaxIdleTimeout"); ok {
			if err = d.Set("max_idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading max_idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_idle_timeout: %v", err)
		}
	}

	if err = d.Set("max_udp_payload_size", flattenFirewallAccessProxyApiGatewayQuicMaxUdpPayloadSize3rdl(o["max-udp-payload-size"], d, "max_udp_payload_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-udp-payload-size"], "FirewallAccessProxyApiGatewayQuic-MaxUdpPayloadSize"); ok {
			if err = d.Set("max_udp_payload_size", vv); err != nil {
				return fmt.Errorf("Error reading max_udp_payload_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_udp_payload_size: %v", err)
		}
	}

	return nil
}

func flattenFirewallAccessProxyApiGatewayQuicFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAccessProxyApiGatewayQuicAckDelayExponent3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicActiveConnectionIdLimit3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicActiveMigration3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicGreaseQuicBit3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicMaxAckDelay3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicMaxDatagramFrameSize3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicMaxIdleTimeout3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayQuicMaxUdpPayloadSize3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAccessProxyApiGatewayQuic(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ack_delay_exponent"); ok || d.HasChange("ack_delay_exponent") {
		t, err := expandFirewallAccessProxyApiGatewayQuicAckDelayExponent3rdl(d, v, "ack_delay_exponent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ack-delay-exponent"] = t
		}
	}

	if v, ok := d.GetOk("active_connection_id_limit"); ok || d.HasChange("active_connection_id_limit") {
		t, err := expandFirewallAccessProxyApiGatewayQuicActiveConnectionIdLimit3rdl(d, v, "active_connection_id_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-connection-id-limit"] = t
		}
	}

	if v, ok := d.GetOk("active_migration"); ok || d.HasChange("active_migration") {
		t, err := expandFirewallAccessProxyApiGatewayQuicActiveMigration3rdl(d, v, "active_migration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-migration"] = t
		}
	}

	if v, ok := d.GetOk("grease_quic_bit"); ok || d.HasChange("grease_quic_bit") {
		t, err := expandFirewallAccessProxyApiGatewayQuicGreaseQuicBit3rdl(d, v, "grease_quic_bit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["grease-quic-bit"] = t
		}
	}

	if v, ok := d.GetOk("max_ack_delay"); ok || d.HasChange("max_ack_delay") {
		t, err := expandFirewallAccessProxyApiGatewayQuicMaxAckDelay3rdl(d, v, "max_ack_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-ack-delay"] = t
		}
	}

	if v, ok := d.GetOk("max_datagram_frame_size"); ok || d.HasChange("max_datagram_frame_size") {
		t, err := expandFirewallAccessProxyApiGatewayQuicMaxDatagramFrameSize3rdl(d, v, "max_datagram_frame_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-datagram-frame-size"] = t
		}
	}

	if v, ok := d.GetOk("max_idle_timeout"); ok || d.HasChange("max_idle_timeout") {
		t, err := expandFirewallAccessProxyApiGatewayQuicMaxIdleTimeout3rdl(d, v, "max_idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("max_udp_payload_size"); ok || d.HasChange("max_udp_payload_size") {
		t, err := expandFirewallAccessProxyApiGatewayQuicMaxUdpPayloadSize3rdl(d, v, "max_udp_payload_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-udp-payload-size"] = t
		}
	}

	return &obj, nil
}
