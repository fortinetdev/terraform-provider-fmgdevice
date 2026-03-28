// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure ICAP servers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIcapServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceIcapServerCreate,
		Read:   resourceIcapServerRead,
		Update: resourceIcapServerUpdate,
		Delete: resourceIcapServerDelete,

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
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"healthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"healthcheck_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIcapServerCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectIcapServer(d)
	if err != nil {
		return fmt.Errorf("Error creating IcapServer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIcapServer(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIcapServer(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IcapServer resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateIcapServer(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IcapServer resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceIcapServerRead(d, m)
}

func resourceIcapServerUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectIcapServer(d)
	if err != nil {
		return fmt.Errorf("Error updating IcapServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIcapServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IcapServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceIcapServerRead(d, m)
}

func resourceIcapServerDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteIcapServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IcapServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIcapServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadIcapServer(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IcapServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIcapServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IcapServer resource from API: %v", err)
	}
	return nil
}

func flattenIcapServerAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerHealthcheckService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerSslCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectIcapServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_type", flattenIcapServerAddrType(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "IcapServer-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenIcapServerFqdn(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "IcapServer-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("healthcheck", flattenIcapServerHealthcheck(o["healthcheck"], d, "healthcheck")); err != nil {
		if vv, ok := fortiAPIPatch(o["healthcheck"], "IcapServer-Healthcheck"); ok {
			if err = d.Set("healthcheck", vv); err != nil {
				return fmt.Errorf("Error reading healthcheck: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading healthcheck: %v", err)
		}
	}

	if err = d.Set("healthcheck_service", flattenIcapServerHealthcheckService(o["healthcheck-service"], d, "healthcheck_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["healthcheck-service"], "IcapServer-HealthcheckService"); ok {
			if err = d.Set("healthcheck_service", vv); err != nil {
				return fmt.Errorf("Error reading healthcheck_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading healthcheck_service: %v", err)
		}
	}

	if err = d.Set("ip_address", flattenIcapServerIpAddress(o["ip-address"], d, "ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-address"], "IcapServer-IpAddress"); ok {
			if err = d.Set("ip_address", vv); err != nil {
				return fmt.Errorf("Error reading ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenIcapServerIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-version"], "IcapServer-IpVersion"); ok {
			if err = d.Set("ip_version", vv); err != nil {
				return fmt.Errorf("Error reading ip_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("ip6_address", flattenIcapServerIp6Address(o["ip6-address"], d, "ip6_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-address"], "IcapServer-Ip6Address"); ok {
			if err = d.Set("ip6_address", vv); err != nil {
				return fmt.Errorf("Error reading ip6_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_address: %v", err)
		}
	}

	if err = d.Set("max_connections", flattenIcapServerMaxConnections(o["max-connections"], d, "max_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-connections"], "IcapServer-MaxConnections"); ok {
			if err = d.Set("max_connections", vv); err != nil {
				return fmt.Errorf("Error reading max_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_connections: %v", err)
		}
	}

	if err = d.Set("name", flattenIcapServerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "IcapServer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenIcapServerPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "IcapServer-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("secure", flattenIcapServerSecure(o["secure"], d, "secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure"], "IcapServer-Secure"); ok {
			if err = d.Set("secure", vv); err != nil {
				return fmt.Errorf("Error reading secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure: %v", err)
		}
	}

	if err = d.Set("ssl_cert", flattenIcapServerSslCert(o["ssl-cert"], d, "ssl_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-cert"], "IcapServer-SslCert"); ok {
			if err = d.Set("ssl_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssl_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_cert: %v", err)
		}
	}

	return nil
}

func flattenIcapServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIcapServerAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerHealthcheckService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerIpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerSslCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectIcapServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandIcapServerAddrType(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok || d.HasChange("fqdn") {
		t, err := expandIcapServerFqdn(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("healthcheck"); ok || d.HasChange("healthcheck") {
		t, err := expandIcapServerHealthcheck(d, v, "healthcheck")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["healthcheck"] = t
		}
	}

	if v, ok := d.GetOk("healthcheck_service"); ok || d.HasChange("healthcheck_service") {
		t, err := expandIcapServerHealthcheckService(d, v, "healthcheck_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["healthcheck-service"] = t
		}
	}

	if v, ok := d.GetOk("ip_address"); ok || d.HasChange("ip_address") {
		t, err := expandIcapServerIpAddress(d, v, "ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok || d.HasChange("ip_version") {
		t, err := expandIcapServerIpVersion(d, v, "ip_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("ip6_address"); ok || d.HasChange("ip6_address") {
		t, err := expandIcapServerIp6Address(d, v, "ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-address"] = t
		}
	}

	if v, ok := d.GetOk("max_connections"); ok || d.HasChange("max_connections") {
		t, err := expandIcapServerMaxConnections(d, v, "max_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-connections"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandIcapServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandIcapServerPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("secure"); ok || d.HasChange("secure") {
		t, err := expandIcapServerSecure(d, v, "secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cert"); ok || d.HasChange("ssl_cert") {
		t, err := expandIcapServerSslCert(d, v, "ssl_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cert"] = t
		}
	}

	return &obj, nil
}
