// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Access Proxy SSH client certificate.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAccessProxySshClientCert() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxySshClientCertCreate,
		Read:   resourceFirewallAccessProxySshClientCertRead,
		Update: resourceFirewallAccessProxySshClientCertUpdate,
		Delete: resourceFirewallAccessProxySshClientCertDelete,

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
			"auth_ca": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"cert_extension": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"critical": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"data": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"permit_agent_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_port_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_pty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_user_rc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_x11_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallAccessProxySshClientCertCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAccessProxySshClientCert(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxySshClientCert resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAccessProxySshClientCert(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAccessProxySshClientCert(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAccessProxySshClientCert resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAccessProxySshClientCert(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAccessProxySshClientCert resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAccessProxySshClientCertRead(d, m)
}

func resourceFirewallAccessProxySshClientCertUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAccessProxySshClientCert(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxySshClientCert resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAccessProxySshClientCert(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxySshClientCert resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAccessProxySshClientCertRead(d, m)
}

func resourceFirewallAccessProxySshClientCertDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallAccessProxySshClientCert(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxySshClientCert resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxySshClientCertRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallAccessProxySshClientCert(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAccessProxySshClientCert resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxySshClientCert(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxySshClientCert resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxySshClientCertAuthCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAccessProxySshClientCertCertExtension(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "critical"
		if _, ok := i["critical"]; ok {
			v := flattenFirewallAccessProxySshClientCertCertExtensionCritical(i["critical"], d, pre_append)
			tmp["critical"] = fortiAPISubPartPatch(v, "FirewallAccessProxySshClientCert-CertExtension-Critical")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := i["data"]; ok {
			v := flattenFirewallAccessProxySshClientCertCertExtensionData(i["data"], d, pre_append)
			tmp["data"] = fortiAPISubPartPatch(v, "FirewallAccessProxySshClientCert-CertExtension-Data")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallAccessProxySshClientCertCertExtensionName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAccessProxySshClientCert-CertExtension-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAccessProxySshClientCertCertExtensionType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAccessProxySshClientCert-CertExtension-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAccessProxySshClientCertCertExtensionCritical(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertCertExtensionData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertCertExtensionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertCertExtensionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertPermitAgentForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertPermitPortForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertPermitPty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertPermitUserRc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertPermitX11Forwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxySshClientCertSourceAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAccessProxySshClientCert(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auth_ca", flattenFirewallAccessProxySshClientCertAuthCa(o["auth-ca"], d, "auth_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-ca"], "FirewallAccessProxySshClientCert-AuthCa"); ok {
			if err = d.Set("auth_ca", vv); err != nil {
				return fmt.Errorf("Error reading auth_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_ca: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("cert_extension", flattenFirewallAccessProxySshClientCertCertExtension(o["cert-extension"], d, "cert_extension")); err != nil {
			if vv, ok := fortiAPIPatch(o["cert-extension"], "FirewallAccessProxySshClientCert-CertExtension"); ok {
				if err = d.Set("cert_extension", vv); err != nil {
					return fmt.Errorf("Error reading cert_extension: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cert_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cert_extension"); ok {
			if err = d.Set("cert_extension", flattenFirewallAccessProxySshClientCertCertExtension(o["cert-extension"], d, "cert_extension")); err != nil {
				if vv, ok := fortiAPIPatch(o["cert-extension"], "FirewallAccessProxySshClientCert-CertExtension"); ok {
					if err = d.Set("cert_extension", vv); err != nil {
						return fmt.Errorf("Error reading cert_extension: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cert_extension: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenFirewallAccessProxySshClientCertName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallAccessProxySshClientCert-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("permit_agent_forwarding", flattenFirewallAccessProxySshClientCertPermitAgentForwarding(o["permit-agent-forwarding"], d, "permit_agent_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-agent-forwarding"], "FirewallAccessProxySshClientCert-PermitAgentForwarding"); ok {
			if err = d.Set("permit_agent_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading permit_agent_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_agent_forwarding: %v", err)
		}
	}

	if err = d.Set("permit_port_forwarding", flattenFirewallAccessProxySshClientCertPermitPortForwarding(o["permit-port-forwarding"], d, "permit_port_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-port-forwarding"], "FirewallAccessProxySshClientCert-PermitPortForwarding"); ok {
			if err = d.Set("permit_port_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading permit_port_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_port_forwarding: %v", err)
		}
	}

	if err = d.Set("permit_pty", flattenFirewallAccessProxySshClientCertPermitPty(o["permit-pty"], d, "permit_pty")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-pty"], "FirewallAccessProxySshClientCert-PermitPty"); ok {
			if err = d.Set("permit_pty", vv); err != nil {
				return fmt.Errorf("Error reading permit_pty: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_pty: %v", err)
		}
	}

	if err = d.Set("permit_user_rc", flattenFirewallAccessProxySshClientCertPermitUserRc(o["permit-user-rc"], d, "permit_user_rc")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-user-rc"], "FirewallAccessProxySshClientCert-PermitUserRc"); ok {
			if err = d.Set("permit_user_rc", vv); err != nil {
				return fmt.Errorf("Error reading permit_user_rc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_user_rc: %v", err)
		}
	}

	if err = d.Set("permit_x11_forwarding", flattenFirewallAccessProxySshClientCertPermitX11Forwarding(o["permit-x11-forwarding"], d, "permit_x11_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-x11-forwarding"], "FirewallAccessProxySshClientCert-PermitX11Forwarding"); ok {
			if err = d.Set("permit_x11_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading permit_x11_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_x11_forwarding: %v", err)
		}
	}

	if err = d.Set("source_address", flattenFirewallAccessProxySshClientCertSourceAddress(o["source-address"], d, "source_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-address"], "FirewallAccessProxySshClientCert-SourceAddress"); ok {
			if err = d.Set("source_address", vv); err != nil {
				return fmt.Errorf("Error reading source_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_address: %v", err)
		}
	}

	return nil
}

func flattenFirewallAccessProxySshClientCertFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAccessProxySshClientCertAuthCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAccessProxySshClientCertCertExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "critical"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["critical"], _ = expandFirewallAccessProxySshClientCertCertExtensionCritical(d, i["critical"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["data"], _ = expandFirewallAccessProxySshClientCertCertExtensionData(d, i["data"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallAccessProxySshClientCertCertExtensionName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAccessProxySshClientCertCertExtensionType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxySshClientCertCertExtensionCritical(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertCertExtensionData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertCertExtensionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertCertExtensionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertPermitAgentForwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertPermitPortForwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertPermitPty(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertPermitUserRc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertPermitX11Forwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxySshClientCertSourceAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAccessProxySshClientCert(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_ca"); ok || d.HasChange("auth_ca") {
		t, err := expandFirewallAccessProxySshClientCertAuthCa(d, v, "auth_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ca"] = t
		}
	}

	if v, ok := d.GetOk("cert_extension"); ok || d.HasChange("cert_extension") {
		t, err := expandFirewallAccessProxySshClientCertCertExtension(d, v, "cert_extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-extension"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallAccessProxySshClientCertName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("permit_agent_forwarding"); ok || d.HasChange("permit_agent_forwarding") {
		t, err := expandFirewallAccessProxySshClientCertPermitAgentForwarding(d, v, "permit_agent_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-agent-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("permit_port_forwarding"); ok || d.HasChange("permit_port_forwarding") {
		t, err := expandFirewallAccessProxySshClientCertPermitPortForwarding(d, v, "permit_port_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-port-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("permit_pty"); ok || d.HasChange("permit_pty") {
		t, err := expandFirewallAccessProxySshClientCertPermitPty(d, v, "permit_pty")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-pty"] = t
		}
	}

	if v, ok := d.GetOk("permit_user_rc"); ok || d.HasChange("permit_user_rc") {
		t, err := expandFirewallAccessProxySshClientCertPermitUserRc(d, v, "permit_user_rc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-user-rc"] = t
		}
	}

	if v, ok := d.GetOk("permit_x11_forwarding"); ok || d.HasChange("permit_x11_forwarding") {
		t, err := expandFirewallAccessProxySshClientCertPermitX11Forwarding(d, v, "permit_x11_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-x11-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("source_address"); ok || d.HasChange("source_address") {
		t, err := expandFirewallAccessProxySshClientCertSourceAddress(d, v, "source_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address"] = t
		}
	}

	return &obj, nil
}
