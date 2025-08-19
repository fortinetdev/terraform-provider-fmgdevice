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

func resourceZtnaWebProxyApiGateway6Quic() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebProxyApiGateway6QuicUpdate,
		Read:   resourceZtnaWebProxyApiGateway6QuicRead,
		Update: resourceZtnaWebProxyApiGateway6QuicUpdate,
		Delete: resourceZtnaWebProxyApiGateway6QuicDelete,

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
			"api_gateway6": &schema.Schema{
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

func resourceZtnaWebProxyApiGateway6QuicUpdate(d *schema.ResourceData, m interface{}) error {
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
	api_gateway6 := d.Get("api_gateway6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway6"] = api_gateway6

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaWebProxyApiGateway6Quic(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxyApiGateway6Quic resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaWebProxyApiGateway6Quic(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxyApiGateway6Quic resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ZtnaWebProxyApiGateway6Quic")

	return resourceZtnaWebProxyApiGateway6QuicRead(d, m)
}

func resourceZtnaWebProxyApiGateway6QuicDelete(d *schema.ResourceData, m interface{}) error {
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
	api_gateway6 := d.Get("api_gateway6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway6"] = api_gateway6

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteZtnaWebProxyApiGateway6Quic(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebProxyApiGateway6Quic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebProxyApiGateway6QuicRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	web_proxy := d.Get("web_proxy").(string)
	api_gateway6 := d.Get("api_gateway6").(string)
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
	if api_gateway6 == "" {
		api_gateway6 = importOptionChecking(m.(*FortiClient).Cfg, "api_gateway6")
		if api_gateway6 == "" {
			return fmt.Errorf("Parameter api_gateway6 is missing")
		}
		if err = d.Set("api_gateway6", api_gateway6); err != nil {
			return fmt.Errorf("Error set params api_gateway6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway6"] = api_gateway6

	o, err := c.ReadZtnaWebProxyApiGateway6Quic(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxyApiGateway6Quic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebProxyApiGateway6Quic(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxyApiGateway6Quic resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebProxyApiGateway6QuicAckDelayExponent3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicActiveMigration3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicGreaseQuicBit3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicMaxAckDelay3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicMaxIdleTimeout3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectZtnaWebProxyApiGateway6Quic(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ack_delay_exponent", flattenZtnaWebProxyApiGateway6QuicAckDelayExponent3rdl(o["ack-delay-exponent"], d, "ack_delay_exponent")); err != nil {
		if vv, ok := fortiAPIPatch(o["ack-delay-exponent"], "ZtnaWebProxyApiGateway6Quic-AckDelayExponent"); ok {
			if err = d.Set("ack_delay_exponent", vv); err != nil {
				return fmt.Errorf("Error reading ack_delay_exponent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ack_delay_exponent: %v", err)
		}
	}

	if err = d.Set("active_connection_id_limit", flattenZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit3rdl(o["active-connection-id-limit"], d, "active_connection_id_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-connection-id-limit"], "ZtnaWebProxyApiGateway6Quic-ActiveConnectionIdLimit"); ok {
			if err = d.Set("active_connection_id_limit", vv); err != nil {
				return fmt.Errorf("Error reading active_connection_id_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_connection_id_limit: %v", err)
		}
	}

	if err = d.Set("active_migration", flattenZtnaWebProxyApiGateway6QuicActiveMigration3rdl(o["active-migration"], d, "active_migration")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-migration"], "ZtnaWebProxyApiGateway6Quic-ActiveMigration"); ok {
			if err = d.Set("active_migration", vv); err != nil {
				return fmt.Errorf("Error reading active_migration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_migration: %v", err)
		}
	}

	if err = d.Set("grease_quic_bit", flattenZtnaWebProxyApiGateway6QuicGreaseQuicBit3rdl(o["grease-quic-bit"], d, "grease_quic_bit")); err != nil {
		if vv, ok := fortiAPIPatch(o["grease-quic-bit"], "ZtnaWebProxyApiGateway6Quic-GreaseQuicBit"); ok {
			if err = d.Set("grease_quic_bit", vv); err != nil {
				return fmt.Errorf("Error reading grease_quic_bit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading grease_quic_bit: %v", err)
		}
	}

	if err = d.Set("max_ack_delay", flattenZtnaWebProxyApiGateway6QuicMaxAckDelay3rdl(o["max-ack-delay"], d, "max_ack_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-ack-delay"], "ZtnaWebProxyApiGateway6Quic-MaxAckDelay"); ok {
			if err = d.Set("max_ack_delay", vv); err != nil {
				return fmt.Errorf("Error reading max_ack_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_ack_delay: %v", err)
		}
	}

	if err = d.Set("max_datagram_frame_size", flattenZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize3rdl(o["max-datagram-frame-size"], d, "max_datagram_frame_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-datagram-frame-size"], "ZtnaWebProxyApiGateway6Quic-MaxDatagramFrameSize"); ok {
			if err = d.Set("max_datagram_frame_size", vv); err != nil {
				return fmt.Errorf("Error reading max_datagram_frame_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_datagram_frame_size: %v", err)
		}
	}

	if err = d.Set("max_idle_timeout", flattenZtnaWebProxyApiGateway6QuicMaxIdleTimeout3rdl(o["max-idle-timeout"], d, "max_idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-idle-timeout"], "ZtnaWebProxyApiGateway6Quic-MaxIdleTimeout"); ok {
			if err = d.Set("max_idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading max_idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_idle_timeout: %v", err)
		}
	}

	if err = d.Set("max_udp_payload_size", flattenZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize3rdl(o["max-udp-payload-size"], d, "max_udp_payload_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-udp-payload-size"], "ZtnaWebProxyApiGateway6Quic-MaxUdpPayloadSize"); ok {
			if err = d.Set("max_udp_payload_size", vv); err != nil {
				return fmt.Errorf("Error reading max_udp_payload_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_udp_payload_size: %v", err)
		}
	}

	return nil
}

func flattenZtnaWebProxyApiGateway6QuicFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaWebProxyApiGateway6QuicAckDelayExponent3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicActiveMigration3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicGreaseQuicBit3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxAckDelay3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxIdleTimeout3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaWebProxyApiGateway6Quic(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ack_delay_exponent"); ok || d.HasChange("ack_delay_exponent") {
		t, err := expandZtnaWebProxyApiGateway6QuicAckDelayExponent3rdl(d, v, "ack_delay_exponent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ack-delay-exponent"] = t
		}
	}

	if v, ok := d.GetOk("active_connection_id_limit"); ok || d.HasChange("active_connection_id_limit") {
		t, err := expandZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit3rdl(d, v, "active_connection_id_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-connection-id-limit"] = t
		}
	}

	if v, ok := d.GetOk("active_migration"); ok || d.HasChange("active_migration") {
		t, err := expandZtnaWebProxyApiGateway6QuicActiveMigration3rdl(d, v, "active_migration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-migration"] = t
		}
	}

	if v, ok := d.GetOk("grease_quic_bit"); ok || d.HasChange("grease_quic_bit") {
		t, err := expandZtnaWebProxyApiGateway6QuicGreaseQuicBit3rdl(d, v, "grease_quic_bit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["grease-quic-bit"] = t
		}
	}

	if v, ok := d.GetOk("max_ack_delay"); ok || d.HasChange("max_ack_delay") {
		t, err := expandZtnaWebProxyApiGateway6QuicMaxAckDelay3rdl(d, v, "max_ack_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-ack-delay"] = t
		}
	}

	if v, ok := d.GetOk("max_datagram_frame_size"); ok || d.HasChange("max_datagram_frame_size") {
		t, err := expandZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize3rdl(d, v, "max_datagram_frame_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-datagram-frame-size"] = t
		}
	}

	if v, ok := d.GetOk("max_idle_timeout"); ok || d.HasChange("max_idle_timeout") {
		t, err := expandZtnaWebProxyApiGateway6QuicMaxIdleTimeout3rdl(d, v, "max_idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("max_udp_payload_size"); ok || d.HasChange("max_udp_payload_size") {
		t, err := expandZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize3rdl(d, v, "max_udp_payload_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-udp-payload-size"] = t
		}
	}

	return &obj, nil
}
