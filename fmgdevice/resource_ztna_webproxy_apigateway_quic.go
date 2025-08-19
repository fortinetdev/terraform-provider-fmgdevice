// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: QUIC setting.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaWebProxyApiGatewayQuic() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebProxyApiGatewayQuicUpdate,
		Read:   resourceZtnaWebProxyApiGatewayQuicRead,
		Update: resourceZtnaWebProxyApiGatewayQuicUpdate,
		Delete: resourceZtnaWebProxyApiGatewayQuicDelete,

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
			"web_proxy": &schema.Schema{
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

func resourceZtnaWebProxyApiGatewayQuicUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	web_proxy := d.Get("web_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway"] = api_gateway

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaWebProxyApiGatewayQuic(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxyApiGatewayQuic resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaWebProxyApiGatewayQuic(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxyApiGatewayQuic resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ZtnaWebProxyApiGatewayQuic")

	return resourceZtnaWebProxyApiGatewayQuicRead(d, m)
}

func resourceZtnaWebProxyApiGatewayQuicDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	web_proxy := d.Get("web_proxy").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway"] = api_gateway

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteZtnaWebProxyApiGatewayQuic(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebProxyApiGatewayQuic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebProxyApiGatewayQuicRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	web_proxy := d.Get("web_proxy").(string)
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
	if web_proxy == "" {
		web_proxy = importOptionChecking(m.(*FortiClient).Cfg, "web_proxy")
		if web_proxy == "" {
			return fmt.Errorf("Parameter web_proxy is missing")
		}
		if err = d.Set("web_proxy", web_proxy); err != nil {
			return fmt.Errorf("Error set params web_proxy: %v", err)
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
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway"] = api_gateway

	o, err := c.ReadZtnaWebProxyApiGatewayQuic(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxyApiGatewayQuic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebProxyApiGatewayQuic(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxyApiGatewayQuic resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebProxyApiGatewayQuicAckDelayExponent3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicActiveConnectionIdLimit3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicActiveMigration3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicGreaseQuicBit3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicMaxAckDelay3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicMaxDatagramFrameSize3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicMaxIdleTimeout3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicMaxUdpPayloadSize3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectZtnaWebProxyApiGatewayQuic(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ack_delay_exponent", flattenZtnaWebProxyApiGatewayQuicAckDelayExponent3rdl(o["ack-delay-exponent"], d, "ack_delay_exponent")); err != nil {
		if vv, ok := fortiAPIPatch(o["ack-delay-exponent"], "ZtnaWebProxyApiGatewayQuic-AckDelayExponent"); ok {
			if err = d.Set("ack_delay_exponent", vv); err != nil {
				return fmt.Errorf("Error reading ack_delay_exponent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ack_delay_exponent: %v", err)
		}
	}

	if err = d.Set("active_connection_id_limit", flattenZtnaWebProxyApiGatewayQuicActiveConnectionIdLimit3rdl(o["active-connection-id-limit"], d, "active_connection_id_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-connection-id-limit"], "ZtnaWebProxyApiGatewayQuic-ActiveConnectionIdLimit"); ok {
			if err = d.Set("active_connection_id_limit", vv); err != nil {
				return fmt.Errorf("Error reading active_connection_id_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_connection_id_limit: %v", err)
		}
	}

	if err = d.Set("active_migration", flattenZtnaWebProxyApiGatewayQuicActiveMigration3rdl(o["active-migration"], d, "active_migration")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-migration"], "ZtnaWebProxyApiGatewayQuic-ActiveMigration"); ok {
			if err = d.Set("active_migration", vv); err != nil {
				return fmt.Errorf("Error reading active_migration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_migration: %v", err)
		}
	}

	if err = d.Set("grease_quic_bit", flattenZtnaWebProxyApiGatewayQuicGreaseQuicBit3rdl(o["grease-quic-bit"], d, "grease_quic_bit")); err != nil {
		if vv, ok := fortiAPIPatch(o["grease-quic-bit"], "ZtnaWebProxyApiGatewayQuic-GreaseQuicBit"); ok {
			if err = d.Set("grease_quic_bit", vv); err != nil {
				return fmt.Errorf("Error reading grease_quic_bit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading grease_quic_bit: %v", err)
		}
	}

	if err = d.Set("max_ack_delay", flattenZtnaWebProxyApiGatewayQuicMaxAckDelay3rdl(o["max-ack-delay"], d, "max_ack_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-ack-delay"], "ZtnaWebProxyApiGatewayQuic-MaxAckDelay"); ok {
			if err = d.Set("max_ack_delay", vv); err != nil {
				return fmt.Errorf("Error reading max_ack_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_ack_delay: %v", err)
		}
	}

	if err = d.Set("max_datagram_frame_size", flattenZtnaWebProxyApiGatewayQuicMaxDatagramFrameSize3rdl(o["max-datagram-frame-size"], d, "max_datagram_frame_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-datagram-frame-size"], "ZtnaWebProxyApiGatewayQuic-MaxDatagramFrameSize"); ok {
			if err = d.Set("max_datagram_frame_size", vv); err != nil {
				return fmt.Errorf("Error reading max_datagram_frame_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_datagram_frame_size: %v", err)
		}
	}

	if err = d.Set("max_idle_timeout", flattenZtnaWebProxyApiGatewayQuicMaxIdleTimeout3rdl(o["max-idle-timeout"], d, "max_idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-idle-timeout"], "ZtnaWebProxyApiGatewayQuic-MaxIdleTimeout"); ok {
			if err = d.Set("max_idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading max_idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_idle_timeout: %v", err)
		}
	}

	if err = d.Set("max_udp_payload_size", flattenZtnaWebProxyApiGatewayQuicMaxUdpPayloadSize3rdl(o["max-udp-payload-size"], d, "max_udp_payload_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-udp-payload-size"], "ZtnaWebProxyApiGatewayQuic-MaxUdpPayloadSize"); ok {
			if err = d.Set("max_udp_payload_size", vv); err != nil {
				return fmt.Errorf("Error reading max_udp_payload_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_udp_payload_size: %v", err)
		}
	}

	return nil
}

func flattenZtnaWebProxyApiGatewayQuicFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaWebProxyApiGatewayQuicAckDelayExponent3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicActiveConnectionIdLimit3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicActiveMigration3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicGreaseQuicBit3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicMaxAckDelay3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicMaxDatagramFrameSize3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicMaxIdleTimeout3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicMaxUdpPayloadSize3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaWebProxyApiGatewayQuic(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ack_delay_exponent"); ok || d.HasChange("ack_delay_exponent") {
		t, err := expandZtnaWebProxyApiGatewayQuicAckDelayExponent3rdl(d, v, "ack_delay_exponent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ack-delay-exponent"] = t
		}
	}

	if v, ok := d.GetOk("active_connection_id_limit"); ok || d.HasChange("active_connection_id_limit") {
		t, err := expandZtnaWebProxyApiGatewayQuicActiveConnectionIdLimit3rdl(d, v, "active_connection_id_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-connection-id-limit"] = t
		}
	}

	if v, ok := d.GetOk("active_migration"); ok || d.HasChange("active_migration") {
		t, err := expandZtnaWebProxyApiGatewayQuicActiveMigration3rdl(d, v, "active_migration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-migration"] = t
		}
	}

	if v, ok := d.GetOk("grease_quic_bit"); ok || d.HasChange("grease_quic_bit") {
		t, err := expandZtnaWebProxyApiGatewayQuicGreaseQuicBit3rdl(d, v, "grease_quic_bit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["grease-quic-bit"] = t
		}
	}

	if v, ok := d.GetOk("max_ack_delay"); ok || d.HasChange("max_ack_delay") {
		t, err := expandZtnaWebProxyApiGatewayQuicMaxAckDelay3rdl(d, v, "max_ack_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-ack-delay"] = t
		}
	}

	if v, ok := d.GetOk("max_datagram_frame_size"); ok || d.HasChange("max_datagram_frame_size") {
		t, err := expandZtnaWebProxyApiGatewayQuicMaxDatagramFrameSize3rdl(d, v, "max_datagram_frame_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-datagram-frame-size"] = t
		}
	}

	if v, ok := d.GetOk("max_idle_timeout"); ok || d.HasChange("max_idle_timeout") {
		t, err := expandZtnaWebProxyApiGatewayQuicMaxIdleTimeout3rdl(d, v, "max_idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("max_udp_payload_size"); ok || d.HasChange("max_udp_payload_size") {
		t, err := expandZtnaWebProxyApiGatewayQuicMaxUdpPayloadSize3rdl(d, v, "max_udp_payload_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-udp-payload-size"] = t
		}
	}

	return &obj, nil
}
