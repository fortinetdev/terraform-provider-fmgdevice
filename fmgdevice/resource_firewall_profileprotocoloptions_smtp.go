// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure SMTP protocol options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProfileProtocolOptionsSmtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProfileProtocolOptionsSmtpUpdate,
		Read:   resourceFirewallProfileProtocolOptionsSmtpRead,
		Update: resourceFirewallProfileProtocolOptionsSmtpUpdate,
		Delete: resourceFirewallProfileProtocolOptionsSmtpDelete,

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
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"inspect_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"oversize_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"proxy_after_tcp_handshake": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_bzip2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_busy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_offloaded": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uncompressed_nest_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uncompressed_oversize_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallProfileProtocolOptionsSmtpUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	obj, err := getObjectFirewallProfileProtocolOptionsSmtp(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptionsSmtp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallProfileProtocolOptionsSmtp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptionsSmtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallProfileProtocolOptionsSmtp")

	return resourceFirewallProfileProtocolOptionsSmtpRead(d, m)
}

func resourceFirewallProfileProtocolOptionsSmtpDelete(d *schema.ResourceData, m interface{}) error {
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
	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	wsParams["adom"] = adomv

	err = c.DeleteFirewallProfileProtocolOptionsSmtp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProfileProtocolOptionsSmtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProfileProtocolOptionsSmtpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile_protocol_options := d.Get("profile_protocol_options").(string)
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
	if profile_protocol_options == "" {
		profile_protocol_options = importOptionChecking(m.(*FortiClient).Cfg, "profile_protocol_options")
		if profile_protocol_options == "" {
			return fmt.Errorf("Parameter profile_protocol_options is missing")
		}
		if err = d.Set("profile_protocol_options", profile_protocol_options); err != nil {
			return fmt.Errorf("Error set params profile_protocol_options: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	o, err := c.ReadFirewallProfileProtocolOptionsSmtp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallProfileProtocolOptionsSmtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProfileProtocolOptionsSmtp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProfileProtocolOptionsSmtp resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProfileProtocolOptionsSmtpInspectAll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsSmtpOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpPorts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpScanBzip22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpServerBusy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpSslOffloaded2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpUncompressedNestLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallProfileProtocolOptionsSmtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("inspect_all", flattenFirewallProfileProtocolOptionsSmtpInspectAll2edl(o["inspect-all"], d, "inspect_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["inspect-all"], "FirewallProfileProtocolOptionsSmtp-InspectAll"); ok {
			if err = d.Set("inspect_all", vv); err != nil {
				return fmt.Errorf("Error reading inspect_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inspect_all: %v", err)
		}
	}

	if err = d.Set("options", flattenFirewallProfileProtocolOptionsSmtpOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "FirewallProfileProtocolOptionsSmtp-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("oversize_limit", flattenFirewallProfileProtocolOptionsSmtpOversizeLimit2edl(o["oversize-limit"], d, "oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["oversize-limit"], "FirewallProfileProtocolOptionsSmtp-OversizeLimit"); ok {
			if err = d.Set("oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oversize_limit: %v", err)
		}
	}

	if err = d.Set("ports", flattenFirewallProfileProtocolOptionsSmtpPorts2edl(o["ports"], d, "ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports"], "FirewallProfileProtocolOptionsSmtp-Ports"); ok {
			if err = d.Set("ports", vv); err != nil {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("proxy_after_tcp_handshake", flattenFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake2edl(o["proxy-after-tcp-handshake"], d, "proxy_after_tcp_handshake")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-after-tcp-handshake"], "FirewallProfileProtocolOptionsSmtp-ProxyAfterTcpHandshake"); ok {
			if err = d.Set("proxy_after_tcp_handshake", vv); err != nil {
				return fmt.Errorf("Error reading proxy_after_tcp_handshake: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_after_tcp_handshake: %v", err)
		}
	}

	if err = d.Set("scan_bzip2", flattenFirewallProfileProtocolOptionsSmtpScanBzip22edl(o["scan-bzip2"], d, "scan_bzip2")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-bzip2"], "FirewallProfileProtocolOptionsSmtp-ScanBzip2"); ok {
			if err = d.Set("scan_bzip2", vv); err != nil {
				return fmt.Errorf("Error reading scan_bzip2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_bzip2: %v", err)
		}
	}

	if err = d.Set("server_busy", flattenFirewallProfileProtocolOptionsSmtpServerBusy2edl(o["server-busy"], d, "server_busy")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-busy"], "FirewallProfileProtocolOptionsSmtp-ServerBusy"); ok {
			if err = d.Set("server_busy", vv); err != nil {
				return fmt.Errorf("Error reading server_busy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_busy: %v", err)
		}
	}

	if err = d.Set("ssl_offloaded", flattenFirewallProfileProtocolOptionsSmtpSslOffloaded2edl(o["ssl-offloaded"], d, "ssl_offloaded")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-offloaded"], "FirewallProfileProtocolOptionsSmtp-SslOffloaded"); ok {
			if err = d.Set("ssl_offloaded", vv); err != nil {
				return fmt.Errorf("Error reading ssl_offloaded: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_offloaded: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallProfileProtocolOptionsSmtpStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallProfileProtocolOptionsSmtp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("uncompressed_nest_limit", flattenFirewallProfileProtocolOptionsSmtpUncompressedNestLimit2edl(o["uncompressed-nest-limit"], d, "uncompressed_nest_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-nest-limit"], "FirewallProfileProtocolOptionsSmtp-UncompressedNestLimit"); ok {
			if err = d.Set("uncompressed_nest_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
		}
	}

	if err = d.Set("uncompressed_oversize_limit", flattenFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit2edl(o["uncompressed-oversize-limit"], d, "uncompressed_oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-oversize-limit"], "FirewallProfileProtocolOptionsSmtp-UncompressedOversizeLimit"); ok {
			if err = d.Set("uncompressed_oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
		}
	}

	return nil
}

func flattenFirewallProfileProtocolOptionsSmtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProfileProtocolOptionsSmtpInspectAll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsSmtpOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpPorts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpScanBzip22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpServerBusy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpSslOffloaded2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpUncompressedNestLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProfileProtocolOptionsSmtp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("inspect_all"); ok || d.HasChange("inspect_all") {
		t, err := expandFirewallProfileProtocolOptionsSmtpInspectAll2edl(d, v, "inspect_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspect-all"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandFirewallProfileProtocolOptionsSmtpOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("oversize_limit"); ok || d.HasChange("oversize_limit") {
		t, err := expandFirewallProfileProtocolOptionsSmtpOversizeLimit2edl(d, v, "oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oversize-limit"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandFirewallProfileProtocolOptionsSmtpPorts2edl(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("proxy_after_tcp_handshake"); ok || d.HasChange("proxy_after_tcp_handshake") {
		t, err := expandFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake2edl(d, v, "proxy_after_tcp_handshake")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-after-tcp-handshake"] = t
		}
	}

	if v, ok := d.GetOk("scan_bzip2"); ok || d.HasChange("scan_bzip2") {
		t, err := expandFirewallProfileProtocolOptionsSmtpScanBzip22edl(d, v, "scan_bzip2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-bzip2"] = t
		}
	}

	if v, ok := d.GetOk("server_busy"); ok || d.HasChange("server_busy") {
		t, err := expandFirewallProfileProtocolOptionsSmtpServerBusy2edl(d, v, "server_busy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-busy"] = t
		}
	}

	if v, ok := d.GetOk("ssl_offloaded"); ok || d.HasChange("ssl_offloaded") {
		t, err := expandFirewallProfileProtocolOptionsSmtpSslOffloaded2edl(d, v, "ssl_offloaded")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-offloaded"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallProfileProtocolOptionsSmtpStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_nest_limit"); ok || d.HasChange("uncompressed_nest_limit") {
		t, err := expandFirewallProfileProtocolOptionsSmtpUncompressedNestLimit2edl(d, v, "uncompressed_nest_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-nest-limit"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_oversize_limit"); ok || d.HasChange("uncompressed_oversize_limit") {
		t, err := expandFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit2edl(d, v, "uncompressed_oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-oversize-limit"] = t
		}
	}

	return &obj, nil
}
