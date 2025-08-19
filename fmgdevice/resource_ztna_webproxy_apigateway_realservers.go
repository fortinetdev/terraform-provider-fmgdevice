// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Select the real servers that this Access Proxy will distribute traffic to.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaWebProxyApiGatewayRealservers() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebProxyApiGatewayRealserversCreate,
		Read:   resourceZtnaWebProxyApiGatewayRealserversRead,
		Update: resourceZtnaWebProxyApiGatewayRealserversUpdate,
		Delete: resourceZtnaWebProxyApiGatewayRealserversDelete,

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
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"health_check_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"holddown_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"translate_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceZtnaWebProxyApiGatewayRealserversCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectZtnaWebProxyApiGatewayRealservers(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebProxyApiGatewayRealservers resource while getting object: %v", err)
	}

	_, err = c.CreateZtnaWebProxyApiGatewayRealservers(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebProxyApiGatewayRealservers resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceZtnaWebProxyApiGatewayRealserversRead(d, m)
}

func resourceZtnaWebProxyApiGatewayRealserversUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectZtnaWebProxyApiGatewayRealservers(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxyApiGatewayRealservers resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaWebProxyApiGatewayRealservers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxyApiGatewayRealservers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceZtnaWebProxyApiGatewayRealserversRead(d, m)
}

func resourceZtnaWebProxyApiGatewayRealserversDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteZtnaWebProxyApiGatewayRealservers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebProxyApiGatewayRealservers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebProxyApiGatewayRealserversRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaWebProxyApiGatewayRealservers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxyApiGatewayRealservers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebProxyApiGatewayRealservers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxyApiGatewayRealservers resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebProxyApiGatewayRealserversAddrType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebProxyApiGatewayRealserversHealthCheck3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversHealthCheckProto3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversHolddownInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversHttpHost3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversTranslateHost3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversWeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectZtnaWebProxyApiGatewayRealservers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_type", flattenZtnaWebProxyApiGatewayRealserversAddrType3rdl(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "ZtnaWebProxyApiGatewayRealservers-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("address", flattenZtnaWebProxyApiGatewayRealserversAddress3rdl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ZtnaWebProxyApiGatewayRealservers-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("health_check", flattenZtnaWebProxyApiGatewayRealserversHealthCheck3rdl(o["health-check"], d, "health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check"], "ZtnaWebProxyApiGatewayRealservers-HealthCheck"); ok {
			if err = d.Set("health_check", vv); err != nil {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("health_check_proto", flattenZtnaWebProxyApiGatewayRealserversHealthCheckProto3rdl(o["health-check-proto"], d, "health_check_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-proto"], "ZtnaWebProxyApiGatewayRealservers-HealthCheckProto"); ok {
			if err = d.Set("health_check_proto", vv); err != nil {
				return fmt.Errorf("Error reading health_check_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_proto: %v", err)
		}
	}

	if err = d.Set("holddown_interval", flattenZtnaWebProxyApiGatewayRealserversHolddownInterval3rdl(o["holddown-interval"], d, "holddown_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["holddown-interval"], "ZtnaWebProxyApiGatewayRealservers-HolddownInterval"); ok {
			if err = d.Set("holddown_interval", vv); err != nil {
				return fmt.Errorf("Error reading holddown_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading holddown_interval: %v", err)
		}
	}

	if err = d.Set("http_host", flattenZtnaWebProxyApiGatewayRealserversHttpHost3rdl(o["http-host"], d, "http_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-host"], "ZtnaWebProxyApiGatewayRealservers-HttpHost"); ok {
			if err = d.Set("http_host", vv); err != nil {
				return fmt.Errorf("Error reading http_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_host: %v", err)
		}
	}

	if err = d.Set("fosid", flattenZtnaWebProxyApiGatewayRealserversId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ZtnaWebProxyApiGatewayRealservers-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenZtnaWebProxyApiGatewayRealserversIp3rdl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ZtnaWebProxyApiGatewayRealservers-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("port", flattenZtnaWebProxyApiGatewayRealserversPort3rdl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ZtnaWebProxyApiGatewayRealservers-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("status", flattenZtnaWebProxyApiGatewayRealserversStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ZtnaWebProxyApiGatewayRealservers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("translate_host", flattenZtnaWebProxyApiGatewayRealserversTranslateHost3rdl(o["translate-host"], d, "translate_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["translate-host"], "ZtnaWebProxyApiGatewayRealservers-TranslateHost"); ok {
			if err = d.Set("translate_host", vv); err != nil {
				return fmt.Errorf("Error reading translate_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading translate_host: %v", err)
		}
	}

	if err = d.Set("weight", flattenZtnaWebProxyApiGatewayRealserversWeight3rdl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ZtnaWebProxyApiGatewayRealservers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenZtnaWebProxyApiGatewayRealserversFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaWebProxyApiGatewayRealserversAddrType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebProxyApiGatewayRealserversHealthCheck3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversHealthCheckProto3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversHolddownInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversHttpHost3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversTranslateHost3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversWeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaWebProxyApiGatewayRealservers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandZtnaWebProxyApiGatewayRealserversAddrType3rdl(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandZtnaWebProxyApiGatewayRealserversAddress3rdl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandZtnaWebProxyApiGatewayRealserversHealthCheck3rdl(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("health_check_proto"); ok || d.HasChange("health_check_proto") {
		t, err := expandZtnaWebProxyApiGatewayRealserversHealthCheckProto3rdl(d, v, "health_check_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-proto"] = t
		}
	}

	if v, ok := d.GetOk("holddown_interval"); ok || d.HasChange("holddown_interval") {
		t, err := expandZtnaWebProxyApiGatewayRealserversHolddownInterval3rdl(d, v, "holddown_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holddown-interval"] = t
		}
	}

	if v, ok := d.GetOk("http_host"); ok || d.HasChange("http_host") {
		t, err := expandZtnaWebProxyApiGatewayRealserversHttpHost3rdl(d, v, "http_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-host"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandZtnaWebProxyApiGatewayRealserversId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandZtnaWebProxyApiGatewayRealserversIp3rdl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandZtnaWebProxyApiGatewayRealserversPort3rdl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandZtnaWebProxyApiGatewayRealserversStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("translate_host"); ok || d.HasChange("translate_host") {
		t, err := expandZtnaWebProxyApiGatewayRealserversTranslateHost3rdl(d, v, "translate_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["translate-host"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandZtnaWebProxyApiGatewayRealserversWeight3rdl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
