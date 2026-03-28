// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Select the real servers that this Access Proxy will distribute traffic to.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAccessProxy6ApiGatewayRealservers() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxy6ApiGatewayRealserversCreate,
		Read:   resourceFirewallAccessProxy6ApiGatewayRealserversRead,
		Update: resourceFirewallAccessProxy6ApiGatewayRealserversUpdate,
		Delete: resourceFirewallAccessProxy6ApiGatewayRealserversDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

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
			"access_proxy6": &schema.Schema{
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
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_auth": &schema.Schema{
				Type:     schema.TypeString,
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
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mappedport": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssh_client_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_host_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_host_key_validation": &schema.Schema{
				Type:     schema.TypeString,
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
			"tunnel_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verify_cert": &schema.Schema{
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

func resourceFirewallAccessProxy6ApiGatewayRealserversCreate(d *schema.ResourceData, m interface{}) error {
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
	access_proxy6 := d.Get("access_proxy6").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy6"] = access_proxy6
	paradict["api_gateway"] = api_gateway

	obj, err := getObjectFirewallAccessProxy6ApiGatewayRealservers(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxy6ApiGatewayRealservers resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAccessProxy6ApiGatewayRealservers(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAccessProxy6ApiGatewayRealservers(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAccessProxy6ApiGatewayRealservers resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallAccessProxy6ApiGatewayRealservers(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAccessProxy6ApiGatewayRealservers resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallAccessProxy6ApiGatewayRealserversRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallAccessProxy6ApiGatewayRealservers resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallAccessProxy6ApiGatewayRealserversRead(d, m)
}

func resourceFirewallAccessProxy6ApiGatewayRealserversUpdate(d *schema.ResourceData, m interface{}) error {
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
	access_proxy6 := d.Get("access_proxy6").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy6"] = access_proxy6
	paradict["api_gateway"] = api_gateway

	obj, err := getObjectFirewallAccessProxy6ApiGatewayRealservers(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy6ApiGatewayRealservers resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAccessProxy6ApiGatewayRealservers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy6ApiGatewayRealservers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallAccessProxy6ApiGatewayRealserversRead(d, m)
}

func resourceFirewallAccessProxy6ApiGatewayRealserversDelete(d *schema.ResourceData, m interface{}) error {
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
	access_proxy6 := d.Get("access_proxy6").(string)
	api_gateway := d.Get("api_gateway").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy6"] = access_proxy6
	paradict["api_gateway"] = api_gateway

	wsParams["adom"] = adomv

	err = c.DeleteFirewallAccessProxy6ApiGatewayRealservers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxy6ApiGatewayRealservers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxy6ApiGatewayRealserversRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	access_proxy6 := d.Get("access_proxy6").(string)
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
	if access_proxy6 == "" {
		access_proxy6 = importOptionChecking(m.(*FortiClient).Cfg, "access_proxy6")
		if access_proxy6 == "" {
			return fmt.Errorf("Parameter access_proxy6 is missing")
		}
		if err = d.Set("access_proxy6", access_proxy6); err != nil {
			return fmt.Errorf("Error set params access_proxy6: %v", err)
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
	paradict["access_proxy6"] = access_proxy6
	paradict["api_gateway"] = api_gateway

	o, err := c.ReadFirewallAccessProxy6ApiGatewayRealservers(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAccessProxy6ApiGatewayRealservers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxy6ApiGatewayRealservers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxy6ApiGatewayRealservers resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxy6ApiGatewayRealserversAddrType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversDomain3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversExternalAuth3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheck3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHolddownInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHttpHost3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversMappedport3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshClientCert3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKey3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversTranslateHost3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversVerifyCert3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversWeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAccessProxy6ApiGatewayRealservers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_type", flattenFirewallAccessProxy6ApiGatewayRealserversAddrType3rdl(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "FirewallAccessProxy6ApiGatewayRealservers-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("address", flattenFirewallAccessProxy6ApiGatewayRealserversAddress3rdl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "FirewallAccessProxy6ApiGatewayRealservers-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("domain", flattenFirewallAccessProxy6ApiGatewayRealserversDomain3rdl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "FirewallAccessProxy6ApiGatewayRealservers-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("external_auth", flattenFirewallAccessProxy6ApiGatewayRealserversExternalAuth3rdl(o["external-auth"], d, "external_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-auth"], "FirewallAccessProxy6ApiGatewayRealservers-ExternalAuth"); ok {
			if err = d.Set("external_auth", vv); err != nil {
				return fmt.Errorf("Error reading external_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_auth: %v", err)
		}
	}

	if err = d.Set("health_check", flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheck3rdl(o["health-check"], d, "health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check"], "FirewallAccessProxy6ApiGatewayRealservers-HealthCheck"); ok {
			if err = d.Set("health_check", vv); err != nil {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("health_check_proto", flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto3rdl(o["health-check-proto"], d, "health_check_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-proto"], "FirewallAccessProxy6ApiGatewayRealservers-HealthCheckProto"); ok {
			if err = d.Set("health_check_proto", vv); err != nil {
				return fmt.Errorf("Error reading health_check_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_proto: %v", err)
		}
	}

	if err = d.Set("holddown_interval", flattenFirewallAccessProxy6ApiGatewayRealserversHolddownInterval3rdl(o["holddown-interval"], d, "holddown_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["holddown-interval"], "FirewallAccessProxy6ApiGatewayRealservers-HolddownInterval"); ok {
			if err = d.Set("holddown_interval", vv); err != nil {
				return fmt.Errorf("Error reading holddown_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading holddown_interval: %v", err)
		}
	}

	if err = d.Set("http_host", flattenFirewallAccessProxy6ApiGatewayRealserversHttpHost3rdl(o["http-host"], d, "http_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-host"], "FirewallAccessProxy6ApiGatewayRealservers-HttpHost"); ok {
			if err = d.Set("http_host", vv); err != nil {
				return fmt.Errorf("Error reading http_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_host: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallAccessProxy6ApiGatewayRealserversId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallAccessProxy6ApiGatewayRealservers-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenFirewallAccessProxy6ApiGatewayRealserversIp3rdl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FirewallAccessProxy6ApiGatewayRealservers-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mappedport", flattenFirewallAccessProxy6ApiGatewayRealserversMappedport3rdl(o["mappedport"], d, "mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedport"], "FirewallAccessProxy6ApiGatewayRealservers-Mappedport"); ok {
			if err = d.Set("mappedport", vv); err != nil {
				return fmt.Errorf("Error reading mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedport: %v", err)
		}
	}

	if err = d.Set("port", flattenFirewallAccessProxy6ApiGatewayRealserversPort3rdl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FirewallAccessProxy6ApiGatewayRealservers-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("ssh_client_cert", flattenFirewallAccessProxy6ApiGatewayRealserversSshClientCert3rdl(o["ssh-client-cert"], d, "ssh_client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-client-cert"], "FirewallAccessProxy6ApiGatewayRealservers-SshClientCert"); ok {
			if err = d.Set("ssh_client_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssh_client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_client_cert: %v", err)
		}
	}

	if err = d.Set("ssh_host_key", flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKey3rdl(o["ssh-host-key"], d, "ssh_host_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-host-key"], "FirewallAccessProxy6ApiGatewayRealservers-SshHostKey"); ok {
			if err = d.Set("ssh_host_key", vv); err != nil {
				return fmt.Errorf("Error reading ssh_host_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_host_key: %v", err)
		}
	}

	if err = d.Set("ssh_host_key_validation", flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation3rdl(o["ssh-host-key-validation"], d, "ssh_host_key_validation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-host-key-validation"], "FirewallAccessProxy6ApiGatewayRealservers-SshHostKeyValidation"); ok {
			if err = d.Set("ssh_host_key_validation", vv); err != nil {
				return fmt.Errorf("Error reading ssh_host_key_validation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_host_key_validation: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallAccessProxy6ApiGatewayRealserversStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallAccessProxy6ApiGatewayRealservers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("translate_host", flattenFirewallAccessProxy6ApiGatewayRealserversTranslateHost3rdl(o["translate-host"], d, "translate_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["translate-host"], "FirewallAccessProxy6ApiGatewayRealservers-TranslateHost"); ok {
			if err = d.Set("translate_host", vv); err != nil {
				return fmt.Errorf("Error reading translate_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading translate_host: %v", err)
		}
	}

	if err = d.Set("tunnel_encryption", flattenFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption3rdl(o["tunnel-encryption"], d, "tunnel_encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-encryption"], "FirewallAccessProxy6ApiGatewayRealservers-TunnelEncryption"); ok {
			if err = d.Set("tunnel_encryption", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_encryption: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallAccessProxy6ApiGatewayRealserversType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallAccessProxy6ApiGatewayRealservers-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("verify_cert", flattenFirewallAccessProxy6ApiGatewayRealserversVerifyCert3rdl(o["verify-cert"], d, "verify_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-cert"], "FirewallAccessProxy6ApiGatewayRealservers-VerifyCert"); ok {
			if err = d.Set("verify_cert", vv); err != nil {
				return fmt.Errorf("Error reading verify_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_cert: %v", err)
		}
	}

	if err = d.Set("weight", flattenFirewallAccessProxy6ApiGatewayRealserversWeight3rdl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "FirewallAccessProxy6ApiGatewayRealservers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenFirewallAccessProxy6ApiGatewayRealserversFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAccessProxy6ApiGatewayRealserversAddrType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversDomain3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversExternalAuth3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHealthCheck3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHolddownInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHttpHost3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversMappedport3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshClientCert3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshHostKey3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversTranslateHost3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversVerifyCert3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversWeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAccessProxy6ApiGatewayRealservers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversAddrType3rdl(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversAddress3rdl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversDomain3rdl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("external_auth"); ok || d.HasChange("external_auth") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversExternalAuth3rdl(d, v, "external_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-auth"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversHealthCheck3rdl(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("health_check_proto"); ok || d.HasChange("health_check_proto") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto3rdl(d, v, "health_check_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-proto"] = t
		}
	}

	if v, ok := d.GetOk("holddown_interval"); ok || d.HasChange("holddown_interval") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversHolddownInterval3rdl(d, v, "holddown_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holddown-interval"] = t
		}
	}

	if v, ok := d.GetOk("http_host"); ok || d.HasChange("http_host") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversHttpHost3rdl(d, v, "http_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-host"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversIp3rdl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mappedport"); ok || d.HasChange("mappedport") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversMappedport3rdl(d, v, "mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversPort3rdl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("ssh_client_cert"); ok || d.HasChange("ssh_client_cert") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversSshClientCert3rdl(d, v, "ssh_client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-client-cert"] = t
		}
	}

	if v, ok := d.GetOk("ssh_host_key"); ok || d.HasChange("ssh_host_key") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversSshHostKey3rdl(d, v, "ssh_host_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-host-key"] = t
		}
	}

	if v, ok := d.GetOk("ssh_host_key_validation"); ok || d.HasChange("ssh_host_key_validation") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation3rdl(d, v, "ssh_host_key_validation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-host-key-validation"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("translate_host"); ok || d.HasChange("translate_host") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversTranslateHost3rdl(d, v, "translate_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["translate-host"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_encryption"); ok || d.HasChange("tunnel_encryption") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption3rdl(d, v, "tunnel_encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-encryption"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("verify_cert"); ok || d.HasChange("verify_cert") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversVerifyCert3rdl(d, v, "verify_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-cert"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandFirewallAccessProxy6ApiGatewayRealserversWeight3rdl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
