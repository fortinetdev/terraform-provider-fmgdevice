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

func resourceFirewallAccessProxy6ApiGateway6Realservers() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxy6ApiGateway6RealserversCreate,
		Read:   resourceFirewallAccessProxy6ApiGateway6RealserversRead,
		Update: resourceFirewallAccessProxy6ApiGateway6RealserversUpdate,
		Delete: resourceFirewallAccessProxy6ApiGateway6RealserversDelete,

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
			"api_gateway6": &schema.Schema{
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

func resourceFirewallAccessProxy6ApiGateway6RealserversCreate(d *schema.ResourceData, m interface{}) error {
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
	api_gateway6 := d.Get("api_gateway6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy6"] = access_proxy6
	paradict["api_gateway6"] = api_gateway6

	obj, err := getObjectFirewallAccessProxy6ApiGateway6Realservers(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxy6ApiGateway6Realservers resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAccessProxy6ApiGateway6Realservers(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAccessProxy6ApiGateway6Realservers(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAccessProxy6ApiGateway6Realservers resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallAccessProxy6ApiGateway6Realservers(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAccessProxy6ApiGateway6Realservers resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallAccessProxy6ApiGateway6RealserversRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallAccessProxy6ApiGateway6Realservers resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallAccessProxy6ApiGateway6RealserversRead(d, m)
}

func resourceFirewallAccessProxy6ApiGateway6RealserversUpdate(d *schema.ResourceData, m interface{}) error {
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
	api_gateway6 := d.Get("api_gateway6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy6"] = access_proxy6
	paradict["api_gateway6"] = api_gateway6

	obj, err := getObjectFirewallAccessProxy6ApiGateway6Realservers(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy6ApiGateway6Realservers resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAccessProxy6ApiGateway6Realservers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy6ApiGateway6Realservers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallAccessProxy6ApiGateway6RealserversRead(d, m)
}

func resourceFirewallAccessProxy6ApiGateway6RealserversDelete(d *schema.ResourceData, m interface{}) error {
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
	api_gateway6 := d.Get("api_gateway6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy6"] = access_proxy6
	paradict["api_gateway6"] = api_gateway6

	wsParams["adom"] = adomv

	err = c.DeleteFirewallAccessProxy6ApiGateway6Realservers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxy6ApiGateway6Realservers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxy6ApiGateway6RealserversRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	access_proxy6 := d.Get("access_proxy6").(string)
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
	if access_proxy6 == "" {
		access_proxy6 = importOptionChecking(m.(*FortiClient).Cfg, "access_proxy6")
		if access_proxy6 == "" {
			return fmt.Errorf("Parameter access_proxy6 is missing")
		}
		if err = d.Set("access_proxy6", access_proxy6); err != nil {
			return fmt.Errorf("Error set params access_proxy6: %v", err)
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
	paradict["access_proxy6"] = access_proxy6
	paradict["api_gateway6"] = api_gateway6

	o, err := c.ReadFirewallAccessProxy6ApiGateway6Realservers(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAccessProxy6ApiGateway6Realservers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxy6ApiGateway6Realservers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxy6ApiGateway6Realservers resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxy6ApiGateway6RealserversAddrType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6RealserversDomain3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversExternalAuth3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversHealthCheck3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversHealthCheckProto3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversHolddownInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversHttpHost3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversMappedport3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6RealserversPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversSshClientCert3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKey3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKeyValidation3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversTranslateHost3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversTunnelEncryption3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversVerifyCert3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversWeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAccessProxy6ApiGateway6Realservers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_type", flattenFirewallAccessProxy6ApiGateway6RealserversAddrType3rdl(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "FirewallAccessProxy6ApiGateway6Realservers-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("address", flattenFirewallAccessProxy6ApiGateway6RealserversAddress3rdl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "FirewallAccessProxy6ApiGateway6Realservers-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("domain", flattenFirewallAccessProxy6ApiGateway6RealserversDomain3rdl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "FirewallAccessProxy6ApiGateway6Realservers-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("external_auth", flattenFirewallAccessProxy6ApiGateway6RealserversExternalAuth3rdl(o["external-auth"], d, "external_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-auth"], "FirewallAccessProxy6ApiGateway6Realservers-ExternalAuth"); ok {
			if err = d.Set("external_auth", vv); err != nil {
				return fmt.Errorf("Error reading external_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_auth: %v", err)
		}
	}

	if err = d.Set("health_check", flattenFirewallAccessProxy6ApiGateway6RealserversHealthCheck3rdl(o["health-check"], d, "health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check"], "FirewallAccessProxy6ApiGateway6Realservers-HealthCheck"); ok {
			if err = d.Set("health_check", vv); err != nil {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("health_check_proto", flattenFirewallAccessProxy6ApiGateway6RealserversHealthCheckProto3rdl(o["health-check-proto"], d, "health_check_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-proto"], "FirewallAccessProxy6ApiGateway6Realservers-HealthCheckProto"); ok {
			if err = d.Set("health_check_proto", vv); err != nil {
				return fmt.Errorf("Error reading health_check_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_proto: %v", err)
		}
	}

	if err = d.Set("holddown_interval", flattenFirewallAccessProxy6ApiGateway6RealserversHolddownInterval3rdl(o["holddown-interval"], d, "holddown_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["holddown-interval"], "FirewallAccessProxy6ApiGateway6Realservers-HolddownInterval"); ok {
			if err = d.Set("holddown_interval", vv); err != nil {
				return fmt.Errorf("Error reading holddown_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading holddown_interval: %v", err)
		}
	}

	if err = d.Set("http_host", flattenFirewallAccessProxy6ApiGateway6RealserversHttpHost3rdl(o["http-host"], d, "http_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-host"], "FirewallAccessProxy6ApiGateway6Realservers-HttpHost"); ok {
			if err = d.Set("http_host", vv); err != nil {
				return fmt.Errorf("Error reading http_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_host: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallAccessProxy6ApiGateway6RealserversId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallAccessProxy6ApiGateway6Realservers-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenFirewallAccessProxy6ApiGateway6RealserversIp3rdl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FirewallAccessProxy6ApiGateway6Realservers-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mappedport", flattenFirewallAccessProxy6ApiGateway6RealserversMappedport3rdl(o["mappedport"], d, "mappedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["mappedport"], "FirewallAccessProxy6ApiGateway6Realservers-Mappedport"); ok {
			if err = d.Set("mappedport", vv); err != nil {
				return fmt.Errorf("Error reading mappedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mappedport: %v", err)
		}
	}

	if err = d.Set("port", flattenFirewallAccessProxy6ApiGateway6RealserversPort3rdl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FirewallAccessProxy6ApiGateway6Realservers-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("ssh_client_cert", flattenFirewallAccessProxy6ApiGateway6RealserversSshClientCert3rdl(o["ssh-client-cert"], d, "ssh_client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-client-cert"], "FirewallAccessProxy6ApiGateway6Realservers-SshClientCert"); ok {
			if err = d.Set("ssh_client_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssh_client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_client_cert: %v", err)
		}
	}

	if err = d.Set("ssh_host_key", flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKey3rdl(o["ssh-host-key"], d, "ssh_host_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-host-key"], "FirewallAccessProxy6ApiGateway6Realservers-SshHostKey"); ok {
			if err = d.Set("ssh_host_key", vv); err != nil {
				return fmt.Errorf("Error reading ssh_host_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_host_key: %v", err)
		}
	}

	if err = d.Set("ssh_host_key_validation", flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKeyValidation3rdl(o["ssh-host-key-validation"], d, "ssh_host_key_validation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-host-key-validation"], "FirewallAccessProxy6ApiGateway6Realservers-SshHostKeyValidation"); ok {
			if err = d.Set("ssh_host_key_validation", vv); err != nil {
				return fmt.Errorf("Error reading ssh_host_key_validation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_host_key_validation: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallAccessProxy6ApiGateway6RealserversStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallAccessProxy6ApiGateway6Realservers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("translate_host", flattenFirewallAccessProxy6ApiGateway6RealserversTranslateHost3rdl(o["translate-host"], d, "translate_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["translate-host"], "FirewallAccessProxy6ApiGateway6Realservers-TranslateHost"); ok {
			if err = d.Set("translate_host", vv); err != nil {
				return fmt.Errorf("Error reading translate_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading translate_host: %v", err)
		}
	}

	if err = d.Set("tunnel_encryption", flattenFirewallAccessProxy6ApiGateway6RealserversTunnelEncryption3rdl(o["tunnel-encryption"], d, "tunnel_encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-encryption"], "FirewallAccessProxy6ApiGateway6Realservers-TunnelEncryption"); ok {
			if err = d.Set("tunnel_encryption", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_encryption: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallAccessProxy6ApiGateway6RealserversType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallAccessProxy6ApiGateway6Realservers-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("verify_cert", flattenFirewallAccessProxy6ApiGateway6RealserversVerifyCert3rdl(o["verify-cert"], d, "verify_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-cert"], "FirewallAccessProxy6ApiGateway6Realservers-VerifyCert"); ok {
			if err = d.Set("verify_cert", vv); err != nil {
				return fmt.Errorf("Error reading verify_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_cert: %v", err)
		}
	}

	if err = d.Set("weight", flattenFirewallAccessProxy6ApiGateway6RealserversWeight3rdl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "FirewallAccessProxy6ApiGateway6Realservers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenFirewallAccessProxy6ApiGateway6RealserversFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAccessProxy6ApiGateway6RealserversAddrType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversDomain3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversExternalAuth3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversHealthCheck3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversHealthCheckProto3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversHolddownInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversHttpHost3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversMappedport3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversSshClientCert3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversSshHostKey3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversSshHostKeyValidation3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversTranslateHost3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversTunnelEncryption3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversVerifyCert3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversWeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAccessProxy6ApiGateway6Realservers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversAddrType3rdl(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversAddress3rdl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversDomain3rdl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("external_auth"); ok || d.HasChange("external_auth") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversExternalAuth3rdl(d, v, "external_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-auth"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversHealthCheck3rdl(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("health_check_proto"); ok || d.HasChange("health_check_proto") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversHealthCheckProto3rdl(d, v, "health_check_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-proto"] = t
		}
	}

	if v, ok := d.GetOk("holddown_interval"); ok || d.HasChange("holddown_interval") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversHolddownInterval3rdl(d, v, "holddown_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holddown-interval"] = t
		}
	}

	if v, ok := d.GetOk("http_host"); ok || d.HasChange("http_host") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversHttpHost3rdl(d, v, "http_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-host"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversIp3rdl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mappedport"); ok || d.HasChange("mappedport") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversMappedport3rdl(d, v, "mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversPort3rdl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("ssh_client_cert"); ok || d.HasChange("ssh_client_cert") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversSshClientCert3rdl(d, v, "ssh_client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-client-cert"] = t
		}
	}

	if v, ok := d.GetOk("ssh_host_key"); ok || d.HasChange("ssh_host_key") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversSshHostKey3rdl(d, v, "ssh_host_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-host-key"] = t
		}
	}

	if v, ok := d.GetOk("ssh_host_key_validation"); ok || d.HasChange("ssh_host_key_validation") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversSshHostKeyValidation3rdl(d, v, "ssh_host_key_validation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-host-key-validation"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("translate_host"); ok || d.HasChange("translate_host") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversTranslateHost3rdl(d, v, "translate_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["translate-host"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_encryption"); ok || d.HasChange("tunnel_encryption") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversTunnelEncryption3rdl(d, v, "tunnel_encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-encryption"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("verify_cert"); ok || d.HasChange("verify_cert") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversVerifyCert3rdl(d, v, "verify_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-cert"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandFirewallAccessProxy6ApiGateway6RealserversWeight3rdl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
