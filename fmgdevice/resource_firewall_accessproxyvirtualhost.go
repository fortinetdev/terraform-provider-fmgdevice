// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Access Proxy virtual hosts.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAccessProxyVirtualHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxyVirtualHostCreate,
		Read:   resourceFirewallAccessProxyVirtualHostRead,
		Update: resourceFirewallAccessProxyVirtualHostUpdate,
		Delete: resourceFirewallAccessProxyVirtualHostDelete,

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
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"empty_cert_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"user_agent_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallAccessProxyVirtualHostCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAccessProxyVirtualHost(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxyVirtualHost resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAccessProxyVirtualHost(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAccessProxyVirtualHost(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAccessProxyVirtualHost resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAccessProxyVirtualHost(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAccessProxyVirtualHost resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAccessProxyVirtualHostRead(d, m)
}

func resourceFirewallAccessProxyVirtualHostUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAccessProxyVirtualHost(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxyVirtualHost resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAccessProxyVirtualHost(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxyVirtualHost resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAccessProxyVirtualHostRead(d, m)
}

func resourceFirewallAccessProxyVirtualHostDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallAccessProxyVirtualHost(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxyVirtualHost resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxyVirtualHostRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallAccessProxyVirtualHost(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAccessProxyVirtualHost resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxyVirtualHost(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxyVirtualHost resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxyVirtualHostClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyVirtualHostEmptyCertAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyVirtualHostHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyVirtualHostHostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyVirtualHostName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyVirtualHostReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyVirtualHostSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxyVirtualHostUserAgentDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAccessProxyVirtualHost(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("client_cert", flattenFirewallAccessProxyVirtualHostClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "FirewallAccessProxyVirtualHost-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("empty_cert_action", flattenFirewallAccessProxyVirtualHostEmptyCertAction(o["empty-cert-action"], d, "empty_cert_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["empty-cert-action"], "FirewallAccessProxyVirtualHost-EmptyCertAction"); ok {
			if err = d.Set("empty_cert_action", vv); err != nil {
				return fmt.Errorf("Error reading empty_cert_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading empty_cert_action: %v", err)
		}
	}

	if err = d.Set("host", flattenFirewallAccessProxyVirtualHostHost(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "FirewallAccessProxyVirtualHost-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("host_type", flattenFirewallAccessProxyVirtualHostHostType(o["host-type"], d, "host_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-type"], "FirewallAccessProxyVirtualHost-HostType"); ok {
			if err = d.Set("host_type", vv); err != nil {
				return fmt.Errorf("Error reading host_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_type: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallAccessProxyVirtualHostName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallAccessProxyVirtualHost-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenFirewallAccessProxyVirtualHostReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "FirewallAccessProxyVirtualHost-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenFirewallAccessProxyVirtualHostSslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-certificate"], "FirewallAccessProxyVirtualHost-SslCertificate"); ok {
			if err = d.Set("ssl_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("user_agent_detect", flattenFirewallAccessProxyVirtualHostUserAgentDetect(o["user-agent-detect"], d, "user_agent_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-agent-detect"], "FirewallAccessProxyVirtualHost-UserAgentDetect"); ok {
			if err = d.Set("user_agent_detect", vv); err != nil {
				return fmt.Errorf("Error reading user_agent_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent_detect: %v", err)
		}
	}

	return nil
}

func flattenFirewallAccessProxyVirtualHostFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAccessProxyVirtualHostClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyVirtualHostEmptyCertAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyVirtualHostHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyVirtualHostHostType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyVirtualHostName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyVirtualHostReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyVirtualHostSslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxyVirtualHostUserAgentDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAccessProxyVirtualHost(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandFirewallAccessProxyVirtualHostClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("empty_cert_action"); ok || d.HasChange("empty_cert_action") {
		t, err := expandFirewallAccessProxyVirtualHostEmptyCertAction(d, v, "empty_cert_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandFirewallAccessProxyVirtualHostHost(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("host_type"); ok || d.HasChange("host_type") {
		t, err := expandFirewallAccessProxyVirtualHostHostType(d, v, "host_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallAccessProxyVirtualHostName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandFirewallAccessProxyVirtualHostReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandFirewallAccessProxyVirtualHostSslCertificate(d, v, "ssl_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("user_agent_detect"); ok || d.HasChange("user_agent_detect") {
		t, err := expandFirewallAccessProxyVirtualHostUserAgentDetect(d, v, "user_agent_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent-detect"] = t
		}
	}

	return &obj, nil
}
