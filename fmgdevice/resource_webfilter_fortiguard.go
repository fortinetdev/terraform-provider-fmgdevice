// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard Web Filter service.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterFortiguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterFortiguardUpdate,
		Read:   resourceWebfilterFortiguardRead,
		Update: resourceWebfilterFortiguardUpdate,
		Delete: resourceWebfilterFortiguardDelete,

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
			"cache_mem_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cache_mem_permille": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cache_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_prefix_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"close_ports": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"embed_image": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ovrd_auth_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ovrd_auth_port_http": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ovrd_auth_port_https": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ovrd_auth_port_https_flow": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ovrd_auth_port_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"request_packet_size_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"warn_auth_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebfilterFortiguardUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectWebfilterFortiguard(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFortiguard resource while getting object: %v", err)
	}

	_, err = c.UpdateWebfilterFortiguard(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFortiguard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WebfilterFortiguard")

	return resourceWebfilterFortiguardRead(d, m)
}

func resourceWebfilterFortiguardDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteWebfilterFortiguard(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterFortiguardRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadWebfilterFortiguard(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFortiguard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterFortiguard(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFortiguard resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterFortiguardCacheMemPercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFortiguardCacheMemPermille(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFortiguardCacheMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFortiguardCachePrefixMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFortiguardClosePorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFortiguardEmbedImage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFortiguardOvrdAuthHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFortiguardOvrdAuthPortHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFortiguardOvrdAuthPortHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFortiguardOvrdAuthPortHttpsFlow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFortiguardOvrdAuthPortWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFortiguardRequestPacketSizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFortiguardWarnAuthHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterFortiguard(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cache_mem_percent", flattenWebfilterFortiguardCacheMemPercent(o["cache-mem-percent"], d, "cache_mem_percent")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-mem-percent"], "WebfilterFortiguard-CacheMemPercent"); ok {
			if err = d.Set("cache_mem_percent", vv); err != nil {
				return fmt.Errorf("Error reading cache_mem_percent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_mem_percent: %v", err)
		}
	}

	if err = d.Set("cache_mem_permille", flattenWebfilterFortiguardCacheMemPermille(o["cache-mem-permille"], d, "cache_mem_permille")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-mem-permille"], "WebfilterFortiguard-CacheMemPermille"); ok {
			if err = d.Set("cache_mem_permille", vv); err != nil {
				return fmt.Errorf("Error reading cache_mem_permille: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_mem_permille: %v", err)
		}
	}

	if err = d.Set("cache_mode", flattenWebfilterFortiguardCacheMode(o["cache-mode"], d, "cache_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-mode"], "WebfilterFortiguard-CacheMode"); ok {
			if err = d.Set("cache_mode", vv); err != nil {
				return fmt.Errorf("Error reading cache_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_mode: %v", err)
		}
	}

	if err = d.Set("cache_prefix_match", flattenWebfilterFortiguardCachePrefixMatch(o["cache-prefix-match"], d, "cache_prefix_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-prefix-match"], "WebfilterFortiguard-CachePrefixMatch"); ok {
			if err = d.Set("cache_prefix_match", vv); err != nil {
				return fmt.Errorf("Error reading cache_prefix_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_prefix_match: %v", err)
		}
	}

	if err = d.Set("close_ports", flattenWebfilterFortiguardClosePorts(o["close-ports"], d, "close_ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["close-ports"], "WebfilterFortiguard-ClosePorts"); ok {
			if err = d.Set("close_ports", vv); err != nil {
				return fmt.Errorf("Error reading close_ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading close_ports: %v", err)
		}
	}

	if err = d.Set("embed_image", flattenWebfilterFortiguardEmbedImage(o["embed-image"], d, "embed_image")); err != nil {
		if vv, ok := fortiAPIPatch(o["embed-image"], "WebfilterFortiguard-EmbedImage"); ok {
			if err = d.Set("embed_image", vv); err != nil {
				return fmt.Errorf("Error reading embed_image: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading embed_image: %v", err)
		}
	}

	if err = d.Set("ovrd_auth_https", flattenWebfilterFortiguardOvrdAuthHttps(o["ovrd-auth-https"], d, "ovrd_auth_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-auth-https"], "WebfilterFortiguard-OvrdAuthHttps"); ok {
			if err = d.Set("ovrd_auth_https", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_auth_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_auth_https: %v", err)
		}
	}

	if err = d.Set("ovrd_auth_port_http", flattenWebfilterFortiguardOvrdAuthPortHttp(o["ovrd-auth-port-http"], d, "ovrd_auth_port_http")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-auth-port-http"], "WebfilterFortiguard-OvrdAuthPortHttp"); ok {
			if err = d.Set("ovrd_auth_port_http", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_auth_port_http: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_auth_port_http: %v", err)
		}
	}

	if err = d.Set("ovrd_auth_port_https", flattenWebfilterFortiguardOvrdAuthPortHttps(o["ovrd-auth-port-https"], d, "ovrd_auth_port_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-auth-port-https"], "WebfilterFortiguard-OvrdAuthPortHttps"); ok {
			if err = d.Set("ovrd_auth_port_https", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_auth_port_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_auth_port_https: %v", err)
		}
	}

	if err = d.Set("ovrd_auth_port_https_flow", flattenWebfilterFortiguardOvrdAuthPortHttpsFlow(o["ovrd-auth-port-https-flow"], d, "ovrd_auth_port_https_flow")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-auth-port-https-flow"], "WebfilterFortiguard-OvrdAuthPortHttpsFlow"); ok {
			if err = d.Set("ovrd_auth_port_https_flow", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_auth_port_https_flow: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_auth_port_https_flow: %v", err)
		}
	}

	if err = d.Set("ovrd_auth_port_warning", flattenWebfilterFortiguardOvrdAuthPortWarning(o["ovrd-auth-port-warning"], d, "ovrd_auth_port_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-auth-port-warning"], "WebfilterFortiguard-OvrdAuthPortWarning"); ok {
			if err = d.Set("ovrd_auth_port_warning", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_auth_port_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_auth_port_warning: %v", err)
		}
	}

	if err = d.Set("request_packet_size_limit", flattenWebfilterFortiguardRequestPacketSizeLimit(o["request-packet-size-limit"], d, "request_packet_size_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["request-packet-size-limit"], "WebfilterFortiguard-RequestPacketSizeLimit"); ok {
			if err = d.Set("request_packet_size_limit", vv); err != nil {
				return fmt.Errorf("Error reading request_packet_size_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading request_packet_size_limit: %v", err)
		}
	}

	if err = d.Set("warn_auth_https", flattenWebfilterFortiguardWarnAuthHttps(o["warn-auth-https"], d, "warn_auth_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["warn-auth-https"], "WebfilterFortiguard-WarnAuthHttps"); ok {
			if err = d.Set("warn_auth_https", vv); err != nil {
				return fmt.Errorf("Error reading warn_auth_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading warn_auth_https: %v", err)
		}
	}

	return nil
}

func flattenWebfilterFortiguardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterFortiguardCacheMemPercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardCacheMemPermille(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardCacheMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardCachePrefixMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardClosePorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardEmbedImage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardOvrdAuthHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardOvrdAuthPortHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardOvrdAuthPortHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardOvrdAuthPortHttpsFlow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardOvrdAuthPortWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardRequestPacketSizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardWarnAuthHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterFortiguard(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cache_mem_percent"); ok || d.HasChange("cache_mem_percent") {
		t, err := expandWebfilterFortiguardCacheMemPercent(d, v, "cache_mem_percent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-mem-percent"] = t
		}
	}

	if v, ok := d.GetOk("cache_mem_permille"); ok || d.HasChange("cache_mem_permille") {
		t, err := expandWebfilterFortiguardCacheMemPermille(d, v, "cache_mem_permille")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-mem-permille"] = t
		}
	}

	if v, ok := d.GetOk("cache_mode"); ok || d.HasChange("cache_mode") {
		t, err := expandWebfilterFortiguardCacheMode(d, v, "cache_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-mode"] = t
		}
	}

	if v, ok := d.GetOk("cache_prefix_match"); ok || d.HasChange("cache_prefix_match") {
		t, err := expandWebfilterFortiguardCachePrefixMatch(d, v, "cache_prefix_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-prefix-match"] = t
		}
	}

	if v, ok := d.GetOk("close_ports"); ok || d.HasChange("close_ports") {
		t, err := expandWebfilterFortiguardClosePorts(d, v, "close_ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["close-ports"] = t
		}
	}

	if v, ok := d.GetOk("embed_image"); ok || d.HasChange("embed_image") {
		t, err := expandWebfilterFortiguardEmbedImage(d, v, "embed_image")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["embed-image"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_auth_https"); ok || d.HasChange("ovrd_auth_https") {
		t, err := expandWebfilterFortiguardOvrdAuthHttps(d, v, "ovrd_auth_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-auth-https"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_auth_port_http"); ok || d.HasChange("ovrd_auth_port_http") {
		t, err := expandWebfilterFortiguardOvrdAuthPortHttp(d, v, "ovrd_auth_port_http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-auth-port-http"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_auth_port_https"); ok || d.HasChange("ovrd_auth_port_https") {
		t, err := expandWebfilterFortiguardOvrdAuthPortHttps(d, v, "ovrd_auth_port_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-auth-port-https"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_auth_port_https_flow"); ok || d.HasChange("ovrd_auth_port_https_flow") {
		t, err := expandWebfilterFortiguardOvrdAuthPortHttpsFlow(d, v, "ovrd_auth_port_https_flow")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-auth-port-https-flow"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_auth_port_warning"); ok || d.HasChange("ovrd_auth_port_warning") {
		t, err := expandWebfilterFortiguardOvrdAuthPortWarning(d, v, "ovrd_auth_port_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-auth-port-warning"] = t
		}
	}

	if v, ok := d.GetOk("request_packet_size_limit"); ok || d.HasChange("request_packet_size_limit") {
		t, err := expandWebfilterFortiguardRequestPacketSizeLimit(d, v, "request_packet_size_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-packet-size-limit"] = t
		}
	}

	if v, ok := d.GetOk("warn_auth_https"); ok || d.HasChange("warn_auth_https") {
		t, err := expandWebfilterFortiguardWarnAuthHttps(d, v, "warn_auth_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warn-auth-https"] = t
		}
	}

	return &obj, nil
}
