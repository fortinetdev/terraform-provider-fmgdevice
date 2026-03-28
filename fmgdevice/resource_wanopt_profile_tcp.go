// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptProfileTcp() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptProfileTcpUpdate,
		Read:   resourceWanoptProfileTcpRead,
		Update: resourceWanoptProfileTcpUpdate,
		Delete: resourceWanoptProfileTcpDelete,

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
			"byte_caching": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"byte_caching_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secure_tunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tunnel_sharing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWanoptProfileTcpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanoptProfileTcp(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptProfileTcp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWanoptProfileTcp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WanoptProfileTcp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WanoptProfileTcp")

	return resourceWanoptProfileTcpRead(d, m)
}

func resourceWanoptProfileTcpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWanoptProfileTcp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptProfileTcp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptProfileTcpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptProfileTcp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WanoptProfileTcp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptProfileTcp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptProfileTcp resource from API: %v", err)
	}
	return nil
}

func flattenWanoptProfileTcpByteCaching2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpByteCachingOpt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpLogTraffic2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpSecureTunnel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpSsl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpSslPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWanoptProfileTcpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpTunnelSharing2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptProfileTcp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("byte_caching", flattenWanoptProfileTcpByteCaching2edl(o["byte-caching"], d, "byte_caching")); err != nil {
		if vv, ok := fortiAPIPatch(o["byte-caching"], "WanoptProfileTcp-ByteCaching"); ok {
			if err = d.Set("byte_caching", vv); err != nil {
				return fmt.Errorf("Error reading byte_caching: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading byte_caching: %v", err)
		}
	}

	if err = d.Set("byte_caching_opt", flattenWanoptProfileTcpByteCachingOpt2edl(o["byte-caching-opt"], d, "byte_caching_opt")); err != nil {
		if vv, ok := fortiAPIPatch(o["byte-caching-opt"], "WanoptProfileTcp-ByteCachingOpt"); ok {
			if err = d.Set("byte_caching_opt", vv); err != nil {
				return fmt.Errorf("Error reading byte_caching_opt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading byte_caching_opt: %v", err)
		}
	}

	if err = d.Set("log_traffic", flattenWanoptProfileTcpLogTraffic2edl(o["log-traffic"], d, "log_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-traffic"], "WanoptProfileTcp-LogTraffic"); ok {
			if err = d.Set("log_traffic", vv); err != nil {
				return fmt.Errorf("Error reading log_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_traffic: %v", err)
		}
	}

	if err = d.Set("port", flattenWanoptProfileTcpPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "WanoptProfileTcp-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("secure_tunnel", flattenWanoptProfileTcpSecureTunnel2edl(o["secure-tunnel"], d, "secure_tunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-tunnel"], "WanoptProfileTcp-SecureTunnel"); ok {
			if err = d.Set("secure_tunnel", vv); err != nil {
				return fmt.Errorf("Error reading secure_tunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_tunnel: %v", err)
		}
	}

	if err = d.Set("ssl", flattenWanoptProfileTcpSsl2edl(o["ssl"], d, "ssl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl"], "WanoptProfileTcp-Ssl"); ok {
			if err = d.Set("ssl", vv); err != nil {
				return fmt.Errorf("Error reading ssl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("ssl_port", flattenWanoptProfileTcpSslPort2edl(o["ssl-port"], d, "ssl_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-port"], "WanoptProfileTcp-SslPort"); ok {
			if err = d.Set("ssl_port", vv); err != nil {
				return fmt.Errorf("Error reading ssl_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_port: %v", err)
		}
	}

	if err = d.Set("status", flattenWanoptProfileTcpStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WanoptProfileTcp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tunnel_sharing", flattenWanoptProfileTcpTunnelSharing2edl(o["tunnel-sharing"], d, "tunnel_sharing")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-sharing"], "WanoptProfileTcp-TunnelSharing"); ok {
			if err = d.Set("tunnel_sharing", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_sharing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_sharing: %v", err)
		}
	}

	return nil
}

func flattenWanoptProfileTcpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptProfileTcpByteCaching2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpByteCachingOpt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpLogTraffic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpSecureTunnel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpSsl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpSslPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWanoptProfileTcpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpTunnelSharing2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptProfileTcp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("byte_caching"); ok || d.HasChange("byte_caching") {
		t, err := expandWanoptProfileTcpByteCaching2edl(d, v, "byte_caching")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["byte-caching"] = t
		}
	}

	if v, ok := d.GetOk("byte_caching_opt"); ok || d.HasChange("byte_caching_opt") {
		t, err := expandWanoptProfileTcpByteCachingOpt2edl(d, v, "byte_caching_opt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["byte-caching-opt"] = t
		}
	}

	if v, ok := d.GetOk("log_traffic"); ok || d.HasChange("log_traffic") {
		t, err := expandWanoptProfileTcpLogTraffic2edl(d, v, "log_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-traffic"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandWanoptProfileTcpPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("secure_tunnel"); ok || d.HasChange("secure_tunnel") {
		t, err := expandWanoptProfileTcpSecureTunnel2edl(d, v, "secure_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok || d.HasChange("ssl") {
		t, err := expandWanoptProfileTcpSsl2edl(d, v, "ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("ssl_port"); ok || d.HasChange("ssl_port") {
		t, err := expandWanoptProfileTcpSslPort2edl(d, v, "ssl_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWanoptProfileTcpStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_sharing"); ok || d.HasChange("tunnel_sharing") {
		t, err := expandWanoptProfileTcpTunnelSharing2edl(d, v, "tunnel_sharing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-sharing"] = t
		}
	}

	return &obj, nil
}
