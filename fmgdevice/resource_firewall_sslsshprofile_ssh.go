// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure SSH options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSslSshProfileSsh() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSslSshProfileSshUpdate,
		Read:   resourceFirewallSslSshProfileSshRead,
		Update: resourceFirewallSslSshProfileSshUpdate,
		Delete: resourceFirewallSslSshProfileSshDelete,

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
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"inspect_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"proxy_after_tcp_handshake": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_tun_policy_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unsupported_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallSslSshProfileSshUpdate(d *schema.ResourceData, m interface{}) error {
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
	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	obj, err := getObjectFirewallSslSshProfileSsh(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileSsh resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallSslSshProfileSsh(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileSsh resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallSslSshProfileSsh")

	return resourceFirewallSslSshProfileSshRead(d, m)
}

func resourceFirewallSslSshProfileSshDelete(d *schema.ResourceData, m interface{}) error {
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
	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	wsParams["adom"] = adomv

	err = c.DeleteFirewallSslSshProfileSsh(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSslSshProfileSsh resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslSshProfileSshRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
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
	if ssl_ssh_profile == "" {
		ssl_ssh_profile = importOptionChecking(m.(*FortiClient).Cfg, "ssl_ssh_profile")
		if ssl_ssh_profile == "" {
			return fmt.Errorf("Parameter ssl_ssh_profile is missing")
		}
		if err = d.Set("ssl_ssh_profile", ssl_ssh_profile); err != nil {
			return fmt.Errorf("Error set params ssl_ssh_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	o, err := c.ReadFirewallSslSshProfileSsh(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallSslSshProfileSsh resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslSshProfileSsh(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSshProfileSsh resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslSshProfileSshInspectAll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshPorts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallSslSshProfileSshProxyAfterTcpHandshake2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshSshAlgorithm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshSshTunPolicyCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshUnsupportedVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSslSshProfileSsh(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("inspect_all", flattenFirewallSslSshProfileSshInspectAll2edl(o["inspect-all"], d, "inspect_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["inspect-all"], "FirewallSslSshProfileSsh-InspectAll"); ok {
			if err = d.Set("inspect_all", vv); err != nil {
				return fmt.Errorf("Error reading inspect_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inspect_all: %v", err)
		}
	}

	if err = d.Set("ports", flattenFirewallSslSshProfileSshPorts2edl(o["ports"], d, "ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports"], "FirewallSslSshProfileSsh-Ports"); ok {
			if err = d.Set("ports", vv); err != nil {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("proxy_after_tcp_handshake", flattenFirewallSslSshProfileSshProxyAfterTcpHandshake2edl(o["proxy-after-tcp-handshake"], d, "proxy_after_tcp_handshake")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-after-tcp-handshake"], "FirewallSslSshProfileSsh-ProxyAfterTcpHandshake"); ok {
			if err = d.Set("proxy_after_tcp_handshake", vv); err != nil {
				return fmt.Errorf("Error reading proxy_after_tcp_handshake: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_after_tcp_handshake: %v", err)
		}
	}

	if err = d.Set("ssh_algorithm", flattenFirewallSslSshProfileSshSshAlgorithm2edl(o["ssh-algorithm"], d, "ssh_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-algorithm"], "FirewallSslSshProfileSsh-SshAlgorithm"); ok {
			if err = d.Set("ssh_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading ssh_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_algorithm: %v", err)
		}
	}

	if err = d.Set("ssh_tun_policy_check", flattenFirewallSslSshProfileSshSshTunPolicyCheck2edl(o["ssh-tun-policy-check"], d, "ssh_tun_policy_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-tun-policy-check"], "FirewallSslSshProfileSsh-SshTunPolicyCheck"); ok {
			if err = d.Set("ssh_tun_policy_check", vv); err != nil {
				return fmt.Errorf("Error reading ssh_tun_policy_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_tun_policy_check: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallSslSshProfileSshStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallSslSshProfileSsh-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("unsupported_version", flattenFirewallSslSshProfileSshUnsupportedVersion2edl(o["unsupported-version"], d, "unsupported_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["unsupported-version"], "FirewallSslSshProfileSsh-UnsupportedVersion"); ok {
			if err = d.Set("unsupported_version", vv); err != nil {
				return fmt.Errorf("Error reading unsupported_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unsupported_version: %v", err)
		}
	}

	return nil
}

func flattenFirewallSslSshProfileSshFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSslSshProfileSshInspectAll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshPorts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallSslSshProfileSshProxyAfterTcpHandshake2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshSshAlgorithm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshSshTunPolicyCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshUnsupportedVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSslSshProfileSsh(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("inspect_all"); ok || d.HasChange("inspect_all") {
		t, err := expandFirewallSslSshProfileSshInspectAll2edl(d, v, "inspect_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspect-all"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandFirewallSslSshProfileSshPorts2edl(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("proxy_after_tcp_handshake"); ok || d.HasChange("proxy_after_tcp_handshake") {
		t, err := expandFirewallSslSshProfileSshProxyAfterTcpHandshake2edl(d, v, "proxy_after_tcp_handshake")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-after-tcp-handshake"] = t
		}
	}

	if v, ok := d.GetOk("ssh_algorithm"); ok || d.HasChange("ssh_algorithm") {
		t, err := expandFirewallSslSshProfileSshSshAlgorithm2edl(d, v, "ssh_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssh_tun_policy_check"); ok || d.HasChange("ssh_tun_policy_check") {
		t, err := expandFirewallSslSshProfileSshSshTunPolicyCheck2edl(d, v, "ssh_tun_policy_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-tun-policy-check"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallSslSshProfileSshStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("unsupported_version"); ok || d.HasChange("unsupported_version") {
		t, err := expandFirewallSslSshProfileSshUnsupportedVersion2edl(d, v, "unsupported_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsupported-version"] = t
		}
	}

	return &obj, nil
}
