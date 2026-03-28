// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure SFTP and SCP protocol options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProfileProtocolOptionsSsh() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProfileProtocolOptionsSshUpdate,
		Read:   resourceFirewallProfileProtocolOptionsSshRead,
		Update: resourceFirewallProfileProtocolOptionsSshUpdate,
		Delete: resourceFirewallProfileProtocolOptionsSshDelete,

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
			"comfort_amount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comfort_interval": &schema.Schema{
				Type:     schema.TypeInt,
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
			"scan_bzip2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_offloaded": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stream_based_uncompressed_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_window_maximum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_window_minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_window_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_window_type": &schema.Schema{
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
			"explicit_ftp_tls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallProfileProtocolOptionsSshUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProfileProtocolOptionsSsh(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptionsSsh resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallProfileProtocolOptionsSsh(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptionsSsh resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallProfileProtocolOptionsSsh")

	return resourceFirewallProfileProtocolOptionsSshRead(d, m)
}

func resourceFirewallProfileProtocolOptionsSshDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallProfileProtocolOptionsSsh(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProfileProtocolOptionsSsh resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProfileProtocolOptionsSshRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallProfileProtocolOptionsSsh(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallProfileProtocolOptionsSsh resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProfileProtocolOptionsSsh(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProfileProtocolOptionsSsh resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProfileProtocolOptionsSshComfortAmount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshComfortInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsSshOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshScanBzip22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshSslOffloaded2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshTcpWindowMaximum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshTcpWindowMinimum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshTcpWindowSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshTcpWindowType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshUncompressedNestLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshUncompressedOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSshExplicitFtpTls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallProfileProtocolOptionsSsh(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comfort_amount", flattenFirewallProfileProtocolOptionsSshComfortAmount2edl(o["comfort-amount"], d, "comfort_amount")); err != nil {
		if vv, ok := fortiAPIPatch(o["comfort-amount"], "FirewallProfileProtocolOptionsSsh-ComfortAmount"); ok {
			if err = d.Set("comfort_amount", vv); err != nil {
				return fmt.Errorf("Error reading comfort_amount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comfort_amount: %v", err)
		}
	}

	if err = d.Set("comfort_interval", flattenFirewallProfileProtocolOptionsSshComfortInterval2edl(o["comfort-interval"], d, "comfort_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["comfort-interval"], "FirewallProfileProtocolOptionsSsh-ComfortInterval"); ok {
			if err = d.Set("comfort_interval", vv); err != nil {
				return fmt.Errorf("Error reading comfort_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comfort_interval: %v", err)
		}
	}

	if err = d.Set("options", flattenFirewallProfileProtocolOptionsSshOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "FirewallProfileProtocolOptionsSsh-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("oversize_limit", flattenFirewallProfileProtocolOptionsSshOversizeLimit2edl(o["oversize-limit"], d, "oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["oversize-limit"], "FirewallProfileProtocolOptionsSsh-OversizeLimit"); ok {
			if err = d.Set("oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oversize_limit: %v", err)
		}
	}

	if err = d.Set("scan_bzip2", flattenFirewallProfileProtocolOptionsSshScanBzip22edl(o["scan-bzip2"], d, "scan_bzip2")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-bzip2"], "FirewallProfileProtocolOptionsSsh-ScanBzip2"); ok {
			if err = d.Set("scan_bzip2", vv); err != nil {
				return fmt.Errorf("Error reading scan_bzip2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_bzip2: %v", err)
		}
	}

	if err = d.Set("ssl_offloaded", flattenFirewallProfileProtocolOptionsSshSslOffloaded2edl(o["ssl-offloaded"], d, "ssl_offloaded")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-offloaded"], "FirewallProfileProtocolOptionsSsh-SslOffloaded"); ok {
			if err = d.Set("ssl_offloaded", vv); err != nil {
				return fmt.Errorf("Error reading ssl_offloaded: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_offloaded: %v", err)
		}
	}

	if err = d.Set("stream_based_uncompressed_limit", flattenFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit2edl(o["stream-based-uncompressed-limit"], d, "stream_based_uncompressed_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["stream-based-uncompressed-limit"], "FirewallProfileProtocolOptionsSsh-StreamBasedUncompressedLimit"); ok {
			if err = d.Set("stream_based_uncompressed_limit", vv); err != nil {
				return fmt.Errorf("Error reading stream_based_uncompressed_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stream_based_uncompressed_limit: %v", err)
		}
	}

	if err = d.Set("tcp_window_maximum", flattenFirewallProfileProtocolOptionsSshTcpWindowMaximum2edl(o["tcp-window-maximum"], d, "tcp_window_maximum")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-maximum"], "FirewallProfileProtocolOptionsSsh-TcpWindowMaximum"); ok {
			if err = d.Set("tcp_window_maximum", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_maximum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_maximum: %v", err)
		}
	}

	if err = d.Set("tcp_window_minimum", flattenFirewallProfileProtocolOptionsSshTcpWindowMinimum2edl(o["tcp-window-minimum"], d, "tcp_window_minimum")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-minimum"], "FirewallProfileProtocolOptionsSsh-TcpWindowMinimum"); ok {
			if err = d.Set("tcp_window_minimum", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_minimum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_minimum: %v", err)
		}
	}

	if err = d.Set("tcp_window_size", flattenFirewallProfileProtocolOptionsSshTcpWindowSize2edl(o["tcp-window-size"], d, "tcp_window_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-size"], "FirewallProfileProtocolOptionsSsh-TcpWindowSize"); ok {
			if err = d.Set("tcp_window_size", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_size: %v", err)
		}
	}

	if err = d.Set("tcp_window_type", flattenFirewallProfileProtocolOptionsSshTcpWindowType2edl(o["tcp-window-type"], d, "tcp_window_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-type"], "FirewallProfileProtocolOptionsSsh-TcpWindowType"); ok {
			if err = d.Set("tcp_window_type", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_type: %v", err)
		}
	}

	if err = d.Set("uncompressed_nest_limit", flattenFirewallProfileProtocolOptionsSshUncompressedNestLimit2edl(o["uncompressed-nest-limit"], d, "uncompressed_nest_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-nest-limit"], "FirewallProfileProtocolOptionsSsh-UncompressedNestLimit"); ok {
			if err = d.Set("uncompressed_nest_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
		}
	}

	if err = d.Set("uncompressed_oversize_limit", flattenFirewallProfileProtocolOptionsSshUncompressedOversizeLimit2edl(o["uncompressed-oversize-limit"], d, "uncompressed_oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-oversize-limit"], "FirewallProfileProtocolOptionsSsh-UncompressedOversizeLimit"); ok {
			if err = d.Set("uncompressed_oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
		}
	}

	if err = d.Set("explicit_ftp_tls", flattenFirewallProfileProtocolOptionsSshExplicitFtpTls2edl(o["explicit-ftp-tls"], d, "explicit_ftp_tls")); err != nil {
		if vv, ok := fortiAPIPatch(o["explicit-ftp-tls"], "FirewallProfileProtocolOptionsSsh-ExplicitFtpTls"); ok {
			if err = d.Set("explicit_ftp_tls", vv); err != nil {
				return fmt.Errorf("Error reading explicit_ftp_tls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading explicit_ftp_tls: %v", err)
		}
	}

	return nil
}

func flattenFirewallProfileProtocolOptionsSshFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProfileProtocolOptionsSshComfortAmount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshComfortInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsSshOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshScanBzip22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshSslOffloaded2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshTcpWindowMaximum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshTcpWindowMinimum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshTcpWindowSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshTcpWindowType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshUncompressedNestLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshUncompressedOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSshExplicitFtpTls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProfileProtocolOptionsSsh(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comfort_amount"); ok || d.HasChange("comfort_amount") {
		t, err := expandFirewallProfileProtocolOptionsSshComfortAmount2edl(d, v, "comfort_amount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comfort-amount"] = t
		}
	}

	if v, ok := d.GetOk("comfort_interval"); ok || d.HasChange("comfort_interval") {
		t, err := expandFirewallProfileProtocolOptionsSshComfortInterval2edl(d, v, "comfort_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comfort-interval"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandFirewallProfileProtocolOptionsSshOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("oversize_limit"); ok || d.HasChange("oversize_limit") {
		t, err := expandFirewallProfileProtocolOptionsSshOversizeLimit2edl(d, v, "oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oversize-limit"] = t
		}
	}

	if v, ok := d.GetOk("scan_bzip2"); ok || d.HasChange("scan_bzip2") {
		t, err := expandFirewallProfileProtocolOptionsSshScanBzip22edl(d, v, "scan_bzip2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-bzip2"] = t
		}
	}

	if v, ok := d.GetOk("ssl_offloaded"); ok || d.HasChange("ssl_offloaded") {
		t, err := expandFirewallProfileProtocolOptionsSshSslOffloaded2edl(d, v, "ssl_offloaded")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-offloaded"] = t
		}
	}

	if v, ok := d.GetOk("stream_based_uncompressed_limit"); ok || d.HasChange("stream_based_uncompressed_limit") {
		t, err := expandFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit2edl(d, v, "stream_based_uncompressed_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stream-based-uncompressed-limit"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_maximum"); ok || d.HasChange("tcp_window_maximum") {
		t, err := expandFirewallProfileProtocolOptionsSshTcpWindowMaximum2edl(d, v, "tcp_window_maximum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-maximum"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_minimum"); ok || d.HasChange("tcp_window_minimum") {
		t, err := expandFirewallProfileProtocolOptionsSshTcpWindowMinimum2edl(d, v, "tcp_window_minimum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-minimum"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_size"); ok || d.HasChange("tcp_window_size") {
		t, err := expandFirewallProfileProtocolOptionsSshTcpWindowSize2edl(d, v, "tcp_window_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-size"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_type"); ok || d.HasChange("tcp_window_type") {
		t, err := expandFirewallProfileProtocolOptionsSshTcpWindowType2edl(d, v, "tcp_window_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-type"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_nest_limit"); ok || d.HasChange("uncompressed_nest_limit") {
		t, err := expandFirewallProfileProtocolOptionsSshUncompressedNestLimit2edl(d, v, "uncompressed_nest_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-nest-limit"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_oversize_limit"); ok || d.HasChange("uncompressed_oversize_limit") {
		t, err := expandFirewallProfileProtocolOptionsSshUncompressedOversizeLimit2edl(d, v, "uncompressed_oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-oversize-limit"] = t
		}
	}

	if v, ok := d.GetOk("explicit_ftp_tls"); ok || d.HasChange("explicit_ftp_tls") {
		t, err := expandFirewallProfileProtocolOptionsSshExplicitFtpTls2edl(d, v, "explicit_ftp_tls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["explicit-ftp-tls"] = t
		}
	}

	return &obj, nil
}
