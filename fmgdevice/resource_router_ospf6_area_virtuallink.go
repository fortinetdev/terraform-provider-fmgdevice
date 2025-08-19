// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSPF6 virtual link configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspf6AreaVirtualLink() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspf6AreaVirtualLinkCreate,
		Read:   resourceRouterOspf6AreaVirtualLinkRead,
		Update: resourceRouterOspf6AreaVirtualLinkUpdate,
		Delete: resourceRouterOspf6AreaVirtualLinkDelete,

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
			"area": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dead_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipsec_auth_alg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec_enc_alg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec_keys": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_key": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"enc_key": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"spi": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"key_rollover_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"peer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retransmit_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"transmit_delay": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceRouterOspf6AreaVirtualLinkCreate(d *schema.ResourceData, m interface{}) error {
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
	area := d.Get("area").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterOspf6AreaVirtualLink(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6AreaVirtualLink resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspf6AreaVirtualLink(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6AreaVirtualLink resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterOspf6AreaVirtualLinkRead(d, m)
}

func resourceRouterOspf6AreaVirtualLinkUpdate(d *schema.ResourceData, m interface{}) error {
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
	area := d.Get("area").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterOspf6AreaVirtualLink(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6AreaVirtualLink resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf6AreaVirtualLink(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6AreaVirtualLink resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterOspf6AreaVirtualLinkRead(d, m)
}

func resourceRouterOspf6AreaVirtualLinkDelete(d *schema.ResourceData, m interface{}) error {
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
	area := d.Get("area").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteRouterOspf6AreaVirtualLink(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspf6AreaVirtualLink resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6AreaVirtualLinkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	area := d.Get("area").(string)
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
	if area == "" {
		area = importOptionChecking(m.(*FortiClient).Cfg, "area")
		if area == "" {
			return fmt.Errorf("Parameter area is missing")
		}
		if err = d.Set("area", area); err != nil {
			return fmt.Errorf("Error set params area: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area

	o, err := c.ReadRouterOspf6AreaVirtualLink(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6AreaVirtualLink resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf6AreaVirtualLink(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6AreaVirtualLink resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspf6AreaVirtualLinkAuthentication3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkDeadInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkHelloInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkIpsecAuthAlg3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkIpsecEncAlg3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkIpsecKeys3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := i["spi"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkIpsecKeysSpi3rdl(i["spi"], d, pre_append)
			tmp["spi"] = fortiAPISubPartPatch(v, "RouterOspf6AreaVirtualLink-IpsecKeys-Spi")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaVirtualLinkIpsecKeysSpi3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkKeyRolloverInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkPeer3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkRetransmitInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkTransmitDelay3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspf6AreaVirtualLink(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("authentication", flattenRouterOspf6AreaVirtualLinkAuthentication3rdl(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "RouterOspf6AreaVirtualLink-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("dead_interval", flattenRouterOspf6AreaVirtualLinkDeadInterval3rdl(o["dead-interval"], d, "dead_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["dead-interval"], "RouterOspf6AreaVirtualLink-DeadInterval"); ok {
			if err = d.Set("dead_interval", vv); err != nil {
				return fmt.Errorf("Error reading dead_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dead_interval: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenRouterOspf6AreaVirtualLinkHelloInterval3rdl(o["hello-interval"], d, "hello_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-interval"], "RouterOspf6AreaVirtualLink-HelloInterval"); ok {
			if err = d.Set("hello_interval", vv); err != nil {
				return fmt.Errorf("Error reading hello_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_interval: %v", err)
		}
	}

	if err = d.Set("ipsec_auth_alg", flattenRouterOspf6AreaVirtualLinkIpsecAuthAlg3rdl(o["ipsec-auth-alg"], d, "ipsec_auth_alg")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-auth-alg"], "RouterOspf6AreaVirtualLink-IpsecAuthAlg"); ok {
			if err = d.Set("ipsec_auth_alg", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_auth_alg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_auth_alg: %v", err)
		}
	}

	if err = d.Set("ipsec_enc_alg", flattenRouterOspf6AreaVirtualLinkIpsecEncAlg3rdl(o["ipsec-enc-alg"], d, "ipsec_enc_alg")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-enc-alg"], "RouterOspf6AreaVirtualLink-IpsecEncAlg"); ok {
			if err = d.Set("ipsec_enc_alg", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_enc_alg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_enc_alg: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipsec_keys", flattenRouterOspf6AreaVirtualLinkIpsecKeys3rdl(o["ipsec-keys"], d, "ipsec_keys")); err != nil {
			if vv, ok := fortiAPIPatch(o["ipsec-keys"], "RouterOspf6AreaVirtualLink-IpsecKeys"); ok {
				if err = d.Set("ipsec_keys", vv); err != nil {
					return fmt.Errorf("Error reading ipsec_keys: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ipsec_keys: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipsec_keys"); ok {
			if err = d.Set("ipsec_keys", flattenRouterOspf6AreaVirtualLinkIpsecKeys3rdl(o["ipsec-keys"], d, "ipsec_keys")); err != nil {
				if vv, ok := fortiAPIPatch(o["ipsec-keys"], "RouterOspf6AreaVirtualLink-IpsecKeys"); ok {
					if err = d.Set("ipsec_keys", vv); err != nil {
						return fmt.Errorf("Error reading ipsec_keys: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ipsec_keys: %v", err)
				}
			}
		}
	}

	if err = d.Set("key_rollover_interval", flattenRouterOspf6AreaVirtualLinkKeyRolloverInterval3rdl(o["key-rollover-interval"], d, "key_rollover_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-rollover-interval"], "RouterOspf6AreaVirtualLink-KeyRolloverInterval"); ok {
			if err = d.Set("key_rollover_interval", vv); err != nil {
				return fmt.Errorf("Error reading key_rollover_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_rollover_interval: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterOspf6AreaVirtualLinkName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterOspf6AreaVirtualLink-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("peer", flattenRouterOspf6AreaVirtualLinkPeer3rdl(o["peer"], d, "peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer"], "RouterOspf6AreaVirtualLink-Peer"); ok {
			if err = d.Set("peer", vv); err != nil {
				return fmt.Errorf("Error reading peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer: %v", err)
		}
	}

	if err = d.Set("retransmit_interval", flattenRouterOspf6AreaVirtualLinkRetransmitInterval3rdl(o["retransmit-interval"], d, "retransmit_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["retransmit-interval"], "RouterOspf6AreaVirtualLink-RetransmitInterval"); ok {
			if err = d.Set("retransmit_interval", vv); err != nil {
				return fmt.Errorf("Error reading retransmit_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retransmit_interval: %v", err)
		}
	}

	if err = d.Set("transmit_delay", flattenRouterOspf6AreaVirtualLinkTransmitDelay3rdl(o["transmit-delay"], d, "transmit_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["transmit-delay"], "RouterOspf6AreaVirtualLink-TransmitDelay"); ok {
			if err = d.Set("transmit_delay", vv); err != nil {
				return fmt.Errorf("Error reading transmit_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transmit_delay: %v", err)
		}
	}

	return nil
}

func flattenRouterOspf6AreaVirtualLinkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspf6AreaVirtualLinkAuthentication3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkDeadInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkHelloInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecAuthAlg3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecEncAlg3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeys3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-key"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeysAuthKey3rdl(d, i["auth_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["enc-key"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeysEncKey3rdl(d, i["enc_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["spi"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeysSpi3rdl(d, i["spi"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysAuthKey3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysEncKey3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysSpi3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkKeyRolloverInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkPeer3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkRetransmitInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkTransmitDelay3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf6AreaVirtualLink(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandRouterOspf6AreaVirtualLinkAuthentication3rdl(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("dead_interval"); ok || d.HasChange("dead_interval") {
		t, err := expandRouterOspf6AreaVirtualLinkDeadInterval3rdl(d, v, "dead_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dead-interval"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval"); ok || d.HasChange("hello_interval") {
		t, err := expandRouterOspf6AreaVirtualLinkHelloInterval3rdl(d, v, "hello_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_auth_alg"); ok || d.HasChange("ipsec_auth_alg") {
		t, err := expandRouterOspf6AreaVirtualLinkIpsecAuthAlg3rdl(d, v, "ipsec_auth_alg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-auth-alg"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_enc_alg"); ok || d.HasChange("ipsec_enc_alg") {
		t, err := expandRouterOspf6AreaVirtualLinkIpsecEncAlg3rdl(d, v, "ipsec_enc_alg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-enc-alg"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_keys"); ok || d.HasChange("ipsec_keys") {
		t, err := expandRouterOspf6AreaVirtualLinkIpsecKeys3rdl(d, v, "ipsec_keys")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-keys"] = t
		}
	}

	if v, ok := d.GetOk("key_rollover_interval"); ok || d.HasChange("key_rollover_interval") {
		t, err := expandRouterOspf6AreaVirtualLinkKeyRolloverInterval3rdl(d, v, "key_rollover_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-rollover-interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterOspf6AreaVirtualLinkName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok || d.HasChange("peer") {
		t, err := expandRouterOspf6AreaVirtualLinkPeer3rdl(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	if v, ok := d.GetOk("retransmit_interval"); ok || d.HasChange("retransmit_interval") {
		t, err := expandRouterOspf6AreaVirtualLinkRetransmitInterval3rdl(d, v, "retransmit_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retransmit-interval"] = t
		}
	}

	if v, ok := d.GetOk("transmit_delay"); ok || d.HasChange("transmit_delay") {
		t, err := expandRouterOspf6AreaVirtualLinkTransmitDelay3rdl(d, v, "transmit_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transmit-delay"] = t
		}
	}

	return &obj, nil
}
