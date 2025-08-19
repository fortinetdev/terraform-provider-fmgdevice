// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemMobileTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemMobileTunnelCreate,
		Read:   resourceSystemMobileTunnelRead,
		Update: resourceSystemMobileTunnelUpdate,
		Delete: resourceSystemMobileTunnelDelete,

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
			"hash_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"home_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"home_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"n_mhae_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"n_mhae_key_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n_mhae_spi": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"reg_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reg_retry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"renew_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"roaming_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_mode": &schema.Schema{
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

func resourceSystemMobileTunnelCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemMobileTunnel(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemMobileTunnel resource while getting object: %v", err)
	}

	_, err = c.CreateSystemMobileTunnel(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemMobileTunnel resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemMobileTunnelRead(d, m)
}

func resourceSystemMobileTunnelUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemMobileTunnel(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemMobileTunnel resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemMobileTunnel(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemMobileTunnel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemMobileTunnelRead(d, m)
}

func resourceSystemMobileTunnelDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemMobileTunnel(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemMobileTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemMobileTunnelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemMobileTunnel(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemMobileTunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemMobileTunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemMobileTunnel resource from API: %v", err)
	}
	return nil
}

func flattenSystemMobileTunnelHashAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelHomeAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelHomeAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelNMhaeKeyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelNMhaeSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemMobileTunnelNetworkId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemMobileTunnel-Network-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystemMobileTunnelNetworkInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystemMobileTunnel-Network-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenSystemMobileTunnelNetworkPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "SystemMobileTunnel-Network-Prefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemMobileTunnelNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelNetworkInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemMobileTunnelNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemMobileTunnelRegInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelRegRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelRenewInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelRoamingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemMobileTunnelStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelTunnelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemMobileTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("hash_algorithm", flattenSystemMobileTunnelHashAlgorithm(o["hash-algorithm"], d, "hash_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["hash-algorithm"], "SystemMobileTunnel-HashAlgorithm"); ok {
			if err = d.Set("hash_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading hash_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hash_algorithm: %v", err)
		}
	}

	if err = d.Set("home_address", flattenSystemMobileTunnelHomeAddress(o["home-address"], d, "home_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["home-address"], "SystemMobileTunnel-HomeAddress"); ok {
			if err = d.Set("home_address", vv); err != nil {
				return fmt.Errorf("Error reading home_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading home_address: %v", err)
		}
	}

	if err = d.Set("home_agent", flattenSystemMobileTunnelHomeAgent(o["home-agent"], d, "home_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["home-agent"], "SystemMobileTunnel-HomeAgent"); ok {
			if err = d.Set("home_agent", vv); err != nil {
				return fmt.Errorf("Error reading home_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading home_agent: %v", err)
		}
	}

	if err = d.Set("lifetime", flattenSystemMobileTunnelLifetime(o["lifetime"], d, "lifetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["lifetime"], "SystemMobileTunnel-Lifetime"); ok {
			if err = d.Set("lifetime", vv); err != nil {
				return fmt.Errorf("Error reading lifetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lifetime: %v", err)
		}
	}

	if err = d.Set("n_mhae_key_type", flattenSystemMobileTunnelNMhaeKeyType(o["n-mhae-key-type"], d, "n_mhae_key_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["n-mhae-key-type"], "SystemMobileTunnel-NMhaeKeyType"); ok {
			if err = d.Set("n_mhae_key_type", vv); err != nil {
				return fmt.Errorf("Error reading n_mhae_key_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n_mhae_key_type: %v", err)
		}
	}

	if err = d.Set("n_mhae_spi", flattenSystemMobileTunnelNMhaeSpi(o["n-mhae-spi"], d, "n_mhae_spi")); err != nil {
		if vv, ok := fortiAPIPatch(o["n-mhae-spi"], "SystemMobileTunnel-NMhaeSpi"); ok {
			if err = d.Set("n_mhae_spi", vv); err != nil {
				return fmt.Errorf("Error reading n_mhae_spi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n_mhae_spi: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemMobileTunnelName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemMobileTunnel-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("network", flattenSystemMobileTunnelNetwork(o["network"], d, "network")); err != nil {
			if vv, ok := fortiAPIPatch(o["network"], "SystemMobileTunnel-Network"); ok {
				if err = d.Set("network", vv); err != nil {
					return fmt.Errorf("Error reading network: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading network: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network"); ok {
			if err = d.Set("network", flattenSystemMobileTunnelNetwork(o["network"], d, "network")); err != nil {
				if vv, ok := fortiAPIPatch(o["network"], "SystemMobileTunnel-Network"); ok {
					if err = d.Set("network", vv); err != nil {
						return fmt.Errorf("Error reading network: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading network: %v", err)
				}
			}
		}
	}

	if err = d.Set("reg_interval", flattenSystemMobileTunnelRegInterval(o["reg-interval"], d, "reg_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["reg-interval"], "SystemMobileTunnel-RegInterval"); ok {
			if err = d.Set("reg_interval", vv); err != nil {
				return fmt.Errorf("Error reading reg_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reg_interval: %v", err)
		}
	}

	if err = d.Set("reg_retry", flattenSystemMobileTunnelRegRetry(o["reg-retry"], d, "reg_retry")); err != nil {
		if vv, ok := fortiAPIPatch(o["reg-retry"], "SystemMobileTunnel-RegRetry"); ok {
			if err = d.Set("reg_retry", vv); err != nil {
				return fmt.Errorf("Error reading reg_retry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reg_retry: %v", err)
		}
	}

	if err = d.Set("renew_interval", flattenSystemMobileTunnelRenewInterval(o["renew-interval"], d, "renew_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["renew-interval"], "SystemMobileTunnel-RenewInterval"); ok {
			if err = d.Set("renew_interval", vv); err != nil {
				return fmt.Errorf("Error reading renew_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading renew_interval: %v", err)
		}
	}

	if err = d.Set("roaming_interface", flattenSystemMobileTunnelRoamingInterface(o["roaming-interface"], d, "roaming_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["roaming-interface"], "SystemMobileTunnel-RoamingInterface"); ok {
			if err = d.Set("roaming_interface", vv); err != nil {
				return fmt.Errorf("Error reading roaming_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading roaming_interface: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemMobileTunnelStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemMobileTunnel-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tunnel_mode", flattenSystemMobileTunnelTunnelMode(o["tunnel-mode"], d, "tunnel_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-mode"], "SystemMobileTunnel-TunnelMode"); ok {
			if err = d.Set("tunnel_mode", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_mode: %v", err)
		}
	}

	return nil
}

func flattenSystemMobileTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemMobileTunnelHashAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelHomeAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelHomeAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelNMhaeKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemMobileTunnelNMhaeKeyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelNMhaeSpi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemMobileTunnelNetworkId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystemMobileTunnelNetworkInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandSystemMobileTunnelNetworkPrefix(d, i["prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemMobileTunnelNetworkId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelNetworkInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemMobileTunnelNetworkPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemMobileTunnelRegInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelRegRetry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelRenewInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelRoamingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemMobileTunnelStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelTunnelMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemMobileTunnel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("hash_algorithm"); ok || d.HasChange("hash_algorithm") {
		t, err := expandSystemMobileTunnelHashAlgorithm(d, v, "hash_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("home_address"); ok || d.HasChange("home_address") {
		t, err := expandSystemMobileTunnelHomeAddress(d, v, "home_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["home-address"] = t
		}
	}

	if v, ok := d.GetOk("home_agent"); ok || d.HasChange("home_agent") {
		t, err := expandSystemMobileTunnelHomeAgent(d, v, "home_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["home-agent"] = t
		}
	}

	if v, ok := d.GetOk("lifetime"); ok || d.HasChange("lifetime") {
		t, err := expandSystemMobileTunnelLifetime(d, v, "lifetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lifetime"] = t
		}
	}

	if v, ok := d.GetOk("n_mhae_key"); ok || d.HasChange("n_mhae_key") {
		t, err := expandSystemMobileTunnelNMhaeKey(d, v, "n_mhae_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["n-mhae-key"] = t
		}
	}

	if v, ok := d.GetOk("n_mhae_key_type"); ok || d.HasChange("n_mhae_key_type") {
		t, err := expandSystemMobileTunnelNMhaeKeyType(d, v, "n_mhae_key_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["n-mhae-key-type"] = t
		}
	}

	if v, ok := d.GetOk("n_mhae_spi"); ok || d.HasChange("n_mhae_spi") {
		t, err := expandSystemMobileTunnelNMhaeSpi(d, v, "n_mhae_spi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["n-mhae-spi"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemMobileTunnelName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("network"); ok || d.HasChange("network") {
		t, err := expandSystemMobileTunnelNetwork(d, v, "network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network"] = t
		}
	}

	if v, ok := d.GetOk("reg_interval"); ok || d.HasChange("reg_interval") {
		t, err := expandSystemMobileTunnelRegInterval(d, v, "reg_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reg-interval"] = t
		}
	}

	if v, ok := d.GetOk("reg_retry"); ok || d.HasChange("reg_retry") {
		t, err := expandSystemMobileTunnelRegRetry(d, v, "reg_retry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reg-retry"] = t
		}
	}

	if v, ok := d.GetOk("renew_interval"); ok || d.HasChange("renew_interval") {
		t, err := expandSystemMobileTunnelRenewInterval(d, v, "renew_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["renew-interval"] = t
		}
	}

	if v, ok := d.GetOk("roaming_interface"); ok || d.HasChange("roaming_interface") {
		t, err := expandSystemMobileTunnelRoamingInterface(d, v, "roaming_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roaming-interface"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemMobileTunnelStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_mode"); ok || d.HasChange("tunnel_mode") {
		t, err := expandSystemMobileTunnelTunnelMode(d, v, "tunnel_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-mode"] = t
		}
	}

	return &obj, nil
}
