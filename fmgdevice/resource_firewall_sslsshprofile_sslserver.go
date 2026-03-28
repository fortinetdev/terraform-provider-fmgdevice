// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> SSL server settings used for client certificate request.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSslSshProfileSslServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSslSshProfileSslServerCreate,
		Read:   resourceFirewallSslSshProfileSslServerRead,
		Update: resourceFirewallSslSshProfileSslServerUpdate,
		Delete: resourceFirewallSslSshProfileSslServerDelete,

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
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ftps_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"imaps_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pop3s_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smtps_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_other_client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallSslSshProfileSslServerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallSslSshProfileSslServer(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSslSshProfileSslServer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallSslSshProfileSslServer(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallSslSshProfileSslServer(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallSslSshProfileSslServer resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallSslSshProfileSslServer(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallSslSshProfileSslServer resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallSslSshProfileSslServerRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallSslSshProfileSslServer resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallSslSshProfileSslServerRead(d, m)
}

func resourceFirewallSslSshProfileSslServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallSslSshProfileSslServer(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileSslServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallSslSshProfileSslServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfileSslServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallSslSshProfileSslServerRead(d, m)
}

func resourceFirewallSslSshProfileSslServerDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallSslSshProfileSslServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSslSshProfileSslServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslSshProfileSslServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallSslSshProfileSslServer(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallSslSshProfileSslServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslSshProfileSslServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSshProfileSslServer resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslSshProfileSslServerFtpsClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerHttpsClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerImapsClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerPop3SClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerSmtpsClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerSslOtherClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSslSshProfileSslServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ftps_client_certificate", flattenFirewallSslSshProfileSslServerFtpsClientCertificate2edl(o["ftps-client-certificate"], d, "ftps_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftps-client-certificate"], "FirewallSslSshProfileSslServer-FtpsClientCertificate"); ok {
			if err = d.Set("ftps_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ftps_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftps_client_certificate: %v", err)
		}
	}

	if err = d.Set("https_client_certificate", flattenFirewallSslSshProfileSslServerHttpsClientCertificate2edl(o["https-client-certificate"], d, "https_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-client-certificate"], "FirewallSslSshProfileSslServer-HttpsClientCertificate"); ok {
			if err = d.Set("https_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading https_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_client_certificate: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallSslSshProfileSslServerId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallSslSshProfileSslServer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("imaps_client_certificate", flattenFirewallSslSshProfileSslServerImapsClientCertificate2edl(o["imaps-client-certificate"], d, "imaps_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["imaps-client-certificate"], "FirewallSslSshProfileSslServer-ImapsClientCertificate"); ok {
			if err = d.Set("imaps_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading imaps_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading imaps_client_certificate: %v", err)
		}
	}

	if err = d.Set("ip", flattenFirewallSslSshProfileSslServerIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FirewallSslSshProfileSslServer-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("pop3s_client_certificate", flattenFirewallSslSshProfileSslServerPop3SClientCertificate2edl(o["pop3s-client-certificate"], d, "pop3s_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["pop3s-client-certificate"], "FirewallSslSshProfileSslServer-Pop3SClientCertificate"); ok {
			if err = d.Set("pop3s_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading pop3s_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pop3s_client_certificate: %v", err)
		}
	}

	if err = d.Set("smtps_client_certificate", flattenFirewallSslSshProfileSslServerSmtpsClientCertificate2edl(o["smtps-client-certificate"], d, "smtps_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["smtps-client-certificate"], "FirewallSslSshProfileSslServer-SmtpsClientCertificate"); ok {
			if err = d.Set("smtps_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading smtps_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smtps_client_certificate: %v", err)
		}
	}

	if err = d.Set("ssl_other_client_certificate", flattenFirewallSslSshProfileSslServerSslOtherClientCertificate2edl(o["ssl-other-client-certificate"], d, "ssl_other_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-other-client-certificate"], "FirewallSslSshProfileSslServer-SslOtherClientCertificate"); ok {
			if err = d.Set("ssl_other_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_other_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_other_client_certificate: %v", err)
		}
	}

	return nil
}

func flattenFirewallSslSshProfileSslServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSslSshProfileSslServerFtpsClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerHttpsClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerImapsClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerPop3SClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerSmtpsClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerSslOtherClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSslSshProfileSslServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ftps_client_certificate"); ok || d.HasChange("ftps_client_certificate") {
		t, err := expandFirewallSslSshProfileSslServerFtpsClientCertificate2edl(d, v, "ftps_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftps-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("https_client_certificate"); ok || d.HasChange("https_client_certificate") {
		t, err := expandFirewallSslSshProfileSslServerHttpsClientCertificate2edl(d, v, "https_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallSslSshProfileSslServerId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("imaps_client_certificate"); ok || d.HasChange("imaps_client_certificate") {
		t, err := expandFirewallSslSshProfileSslServerImapsClientCertificate2edl(d, v, "imaps_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imaps-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandFirewallSslSshProfileSslServerIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("pop3s_client_certificate"); ok || d.HasChange("pop3s_client_certificate") {
		t, err := expandFirewallSslSshProfileSslServerPop3SClientCertificate2edl(d, v, "pop3s_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3s-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("smtps_client_certificate"); ok || d.HasChange("smtps_client_certificate") {
		t, err := expandFirewallSslSshProfileSslServerSmtpsClientCertificate2edl(d, v, "smtps_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtps-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_other_client_certificate"); ok || d.HasChange("ssl_other_client_certificate") {
		t, err := expandFirewallSslSshProfileSslServerSslOtherClientCertificate2edl(d, v, "ssl_other_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-other-client-certificate"] = t
		}
	}

	return &obj, nil
}
