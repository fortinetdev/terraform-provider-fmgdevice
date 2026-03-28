// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptProfileHttp() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptProfileHttpUpdate,
		Read:   resourceWanoptProfileHttpRead,
		Update: resourceWanoptProfileHttpUpdate,
		Delete: resourceWanoptProfileHttpDelete,

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
			"log_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prefer_chunking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol_opt": &schema.Schema{
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

func resourceWanoptProfileHttpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanoptProfileHttp(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptProfileHttp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWanoptProfileHttp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WanoptProfileHttp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WanoptProfileHttp")

	return resourceWanoptProfileHttpRead(d, m)
}

func resourceWanoptProfileHttpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWanoptProfileHttp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptProfileHttp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptProfileHttpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptProfileHttp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WanoptProfileHttp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptProfileHttp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptProfileHttp resource from API: %v", err)
	}
	return nil
}

func flattenWanoptProfileHttpByteCaching2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpLogTraffic2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpPreferChunking2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpProtocolOpt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpSecureTunnel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpSsl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpTunnelSharing2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptProfileHttp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("byte_caching", flattenWanoptProfileHttpByteCaching2edl(o["byte-caching"], d, "byte_caching")); err != nil {
		if vv, ok := fortiAPIPatch(o["byte-caching"], "WanoptProfileHttp-ByteCaching"); ok {
			if err = d.Set("byte_caching", vv); err != nil {
				return fmt.Errorf("Error reading byte_caching: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading byte_caching: %v", err)
		}
	}

	if err = d.Set("log_traffic", flattenWanoptProfileHttpLogTraffic2edl(o["log-traffic"], d, "log_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-traffic"], "WanoptProfileHttp-LogTraffic"); ok {
			if err = d.Set("log_traffic", vv); err != nil {
				return fmt.Errorf("Error reading log_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_traffic: %v", err)
		}
	}

	if err = d.Set("prefer_chunking", flattenWanoptProfileHttpPreferChunking2edl(o["prefer-chunking"], d, "prefer_chunking")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefer-chunking"], "WanoptProfileHttp-PreferChunking"); ok {
			if err = d.Set("prefer_chunking", vv); err != nil {
				return fmt.Errorf("Error reading prefer_chunking: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefer_chunking: %v", err)
		}
	}

	if err = d.Set("protocol_opt", flattenWanoptProfileHttpProtocolOpt2edl(o["protocol-opt"], d, "protocol_opt")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol-opt"], "WanoptProfileHttp-ProtocolOpt"); ok {
			if err = d.Set("protocol_opt", vv); err != nil {
				return fmt.Errorf("Error reading protocol_opt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol_opt: %v", err)
		}
	}

	if err = d.Set("secure_tunnel", flattenWanoptProfileHttpSecureTunnel2edl(o["secure-tunnel"], d, "secure_tunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-tunnel"], "WanoptProfileHttp-SecureTunnel"); ok {
			if err = d.Set("secure_tunnel", vv); err != nil {
				return fmt.Errorf("Error reading secure_tunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_tunnel: %v", err)
		}
	}

	if err = d.Set("ssl", flattenWanoptProfileHttpSsl2edl(o["ssl"], d, "ssl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl"], "WanoptProfileHttp-Ssl"); ok {
			if err = d.Set("ssl", vv); err != nil {
				return fmt.Errorf("Error reading ssl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("status", flattenWanoptProfileHttpStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WanoptProfileHttp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tunnel_sharing", flattenWanoptProfileHttpTunnelSharing2edl(o["tunnel-sharing"], d, "tunnel_sharing")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-sharing"], "WanoptProfileHttp-TunnelSharing"); ok {
			if err = d.Set("tunnel_sharing", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_sharing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_sharing: %v", err)
		}
	}

	return nil
}

func flattenWanoptProfileHttpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptProfileHttpByteCaching2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpLogTraffic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpPreferChunking2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpProtocolOpt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpSecureTunnel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpSsl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpTunnelSharing2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptProfileHttp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("byte_caching"); ok || d.HasChange("byte_caching") {
		t, err := expandWanoptProfileHttpByteCaching2edl(d, v, "byte_caching")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["byte-caching"] = t
		}
	}

	if v, ok := d.GetOk("log_traffic"); ok || d.HasChange("log_traffic") {
		t, err := expandWanoptProfileHttpLogTraffic2edl(d, v, "log_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-traffic"] = t
		}
	}

	if v, ok := d.GetOk("prefer_chunking"); ok || d.HasChange("prefer_chunking") {
		t, err := expandWanoptProfileHttpPreferChunking2edl(d, v, "prefer_chunking")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefer-chunking"] = t
		}
	}

	if v, ok := d.GetOk("protocol_opt"); ok || d.HasChange("protocol_opt") {
		t, err := expandWanoptProfileHttpProtocolOpt2edl(d, v, "protocol_opt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol-opt"] = t
		}
	}

	if v, ok := d.GetOk("secure_tunnel"); ok || d.HasChange("secure_tunnel") {
		t, err := expandWanoptProfileHttpSecureTunnel2edl(d, v, "secure_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok || d.HasChange("ssl") {
		t, err := expandWanoptProfileHttpSsl2edl(d, v, "ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWanoptProfileHttpStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_sharing"); ok || d.HasChange("tunnel_sharing") {
		t, err := expandWanoptProfileHttpTunnelSharing2edl(d, v, "tunnel_sharing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-sharing"] = t
		}
	}

	return &obj, nil
}
