// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ZTNA traffic forward proxy reverse service.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaTrafficForwardProxyReverseService() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaTrafficForwardProxyReverseServiceUpdate,
		Read:   resourceZtnaTrafficForwardProxyReverseServiceRead,
		Update: resourceZtnaTrafficForwardProxyReverseServiceUpdate,
		Delete: resourceZtnaTrafficForwardProxyReverseServiceDelete,

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
			"remote_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"certificate": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"health_check_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_max_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trusted_server_ca": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceZtnaTrafficForwardProxyReverseServiceUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaTrafficForwardProxyReverseService(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxyReverseService resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaTrafficForwardProxyReverseService(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxyReverseService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ZtnaTrafficForwardProxyReverseService")

	return resourceZtnaTrafficForwardProxyReverseServiceRead(d, m)
}

func resourceZtnaTrafficForwardProxyReverseServiceDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteZtnaTrafficForwardProxyReverseService(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaTrafficForwardProxyReverseService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaTrafficForwardProxyReverseServiceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaTrafficForwardProxyReverseService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxyReverseService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaTrafficForwardProxyReverseService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxyReverseService resource from API: %v", err)
	}
	return nil
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenZtnaTrafficForwardProxyReverseServiceRemoteServersAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxyReverseService-RemoteServers-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := i["certificate"]; ok {
			v := flattenZtnaTrafficForwardProxyReverseServiceRemoteServersCertificate(i["certificate"], d, pre_append)
			tmp["certificate"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxyReverseService-RemoteServers-Certificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_interval"
		if _, ok := i["health-check-interval"]; ok {
			v := flattenZtnaTrafficForwardProxyReverseServiceRemoteServersHealthCheckInterval(i["health-check-interval"], d, pre_append)
			tmp["health_check_interval"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxyReverseService-RemoteServers-HealthCheckInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenZtnaTrafficForwardProxyReverseServiceRemoteServersName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxyReverseService-RemoteServers-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenZtnaTrafficForwardProxyReverseServiceRemoteServersPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxyReverseService-RemoteServers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {
			v := flattenZtnaTrafficForwardProxyReverseServiceRemoteServersSslMaxVersion(i["ssl-max-version"], d, pre_append)
			tmp["ssl_max_version"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxyReverseService-RemoteServers-SslMaxVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenZtnaTrafficForwardProxyReverseServiceRemoteServersStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxyReverseService-RemoteServers-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trusted_server_ca"
		if _, ok := i["trusted-server-ca"]; ok {
			v := flattenZtnaTrafficForwardProxyReverseServiceRemoteServersTrustedServerCa(i["trusted-server-ca"], d, pre_append)
			tmp["trusted_server_ca"] = fortiAPISubPartPatch(v, "ZtnaTrafficForwardProxyReverseService-RemoteServers-TrustedServerCa")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersTrustedServerCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectZtnaTrafficForwardProxyReverseService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("remote_servers", flattenZtnaTrafficForwardProxyReverseServiceRemoteServers(o["remote-servers"], d, "remote_servers")); err != nil {
			if vv, ok := fortiAPIPatch(o["remote-servers"], "ZtnaTrafficForwardProxyReverseService-RemoteServers"); ok {
				if err = d.Set("remote_servers", vv); err != nil {
					return fmt.Errorf("Error reading remote_servers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading remote_servers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("remote_servers"); ok {
			if err = d.Set("remote_servers", flattenZtnaTrafficForwardProxyReverseServiceRemoteServers(o["remote-servers"], d, "remote_servers")); err != nil {
				if vv, ok := fortiAPIPatch(o["remote-servers"], "ZtnaTrafficForwardProxyReverseService-RemoteServers"); ok {
					if err = d.Set("remote_servers", vv); err != nil {
						return fmt.Errorf("Error reading remote_servers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading remote_servers: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenZtnaTrafficForwardProxyReverseServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["certificate"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersCertificate(d, i["certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check-interval"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersHealthCheckInterval(d, i["health_check_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-max-version"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersSslMaxVersion(d, i["ssl_max_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trusted_server_ca"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["trusted-server-ca"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersTrustedServerCa(d, i["trusted_server_ca"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersTrustedServerCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectZtnaTrafficForwardProxyReverseService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("remote_servers"); ok || d.HasChange("remote_servers") {
		t, err := expandZtnaTrafficForwardProxyReverseServiceRemoteServers(d, v, "remote_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-servers"] = t
		}
	}

	return &obj, nil
}
