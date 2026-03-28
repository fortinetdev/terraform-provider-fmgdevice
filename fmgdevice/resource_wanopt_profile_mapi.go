// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptProfileMapi() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptProfileMapiUpdate,
		Read:   resourceWanoptProfileMapiRead,
		Update: resourceWanoptProfileMapiUpdate,
		Delete: resourceWanoptProfileMapiDelete,

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
			"secure_tunnel": &schema.Schema{
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

func resourceWanoptProfileMapiUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanoptProfileMapi(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptProfileMapi resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWanoptProfileMapi(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WanoptProfileMapi resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WanoptProfileMapi")

	return resourceWanoptProfileMapiRead(d, m)
}

func resourceWanoptProfileMapiDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWanoptProfileMapi(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptProfileMapi resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptProfileMapiRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptProfileMapi(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WanoptProfileMapi resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptProfileMapi(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptProfileMapi resource from API: %v", err)
	}
	return nil
}

func flattenWanoptProfileMapiByteCaching2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileMapiLogTraffic2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileMapiSecureTunnel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileMapiStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileMapiTunnelSharing2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptProfileMapi(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("byte_caching", flattenWanoptProfileMapiByteCaching2edl(o["byte-caching"], d, "byte_caching")); err != nil {
		if vv, ok := fortiAPIPatch(o["byte-caching"], "WanoptProfileMapi-ByteCaching"); ok {
			if err = d.Set("byte_caching", vv); err != nil {
				return fmt.Errorf("Error reading byte_caching: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading byte_caching: %v", err)
		}
	}

	if err = d.Set("log_traffic", flattenWanoptProfileMapiLogTraffic2edl(o["log-traffic"], d, "log_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-traffic"], "WanoptProfileMapi-LogTraffic"); ok {
			if err = d.Set("log_traffic", vv); err != nil {
				return fmt.Errorf("Error reading log_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_traffic: %v", err)
		}
	}

	if err = d.Set("secure_tunnel", flattenWanoptProfileMapiSecureTunnel2edl(o["secure-tunnel"], d, "secure_tunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-tunnel"], "WanoptProfileMapi-SecureTunnel"); ok {
			if err = d.Set("secure_tunnel", vv); err != nil {
				return fmt.Errorf("Error reading secure_tunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_tunnel: %v", err)
		}
	}

	if err = d.Set("status", flattenWanoptProfileMapiStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WanoptProfileMapi-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tunnel_sharing", flattenWanoptProfileMapiTunnelSharing2edl(o["tunnel-sharing"], d, "tunnel_sharing")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-sharing"], "WanoptProfileMapi-TunnelSharing"); ok {
			if err = d.Set("tunnel_sharing", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_sharing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_sharing: %v", err)
		}
	}

	return nil
}

func flattenWanoptProfileMapiFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptProfileMapiByteCaching2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileMapiLogTraffic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileMapiSecureTunnel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileMapiStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileMapiTunnelSharing2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptProfileMapi(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("byte_caching"); ok || d.HasChange("byte_caching") {
		t, err := expandWanoptProfileMapiByteCaching2edl(d, v, "byte_caching")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["byte-caching"] = t
		}
	}

	if v, ok := d.GetOk("log_traffic"); ok || d.HasChange("log_traffic") {
		t, err := expandWanoptProfileMapiLogTraffic2edl(d, v, "log_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-traffic"] = t
		}
	}

	if v, ok := d.GetOk("secure_tunnel"); ok || d.HasChange("secure_tunnel") {
		t, err := expandWanoptProfileMapiSecureTunnel2edl(d, v, "secure_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWanoptProfileMapiStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_sharing"); ok || d.HasChange("tunnel_sharing") {
		t, err := expandWanoptProfileMapiTunnelSharing2edl(d, v, "tunnel_sharing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-sharing"] = t
		}
	}

	return &obj, nil
}
