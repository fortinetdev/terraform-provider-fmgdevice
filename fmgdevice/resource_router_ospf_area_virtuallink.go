// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSPF virtual link configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspfAreaVirtualLink() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspfAreaVirtualLinkCreate,
		Read:   resourceRouterOspfAreaVirtualLinkRead,
		Update: resourceRouterOspfAreaVirtualLinkUpdate,
		Delete: resourceRouterOspfAreaVirtualLinkDelete,

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
			"authentication_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
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
			"keychain": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"md5_keychain": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"md5_keys": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"key_string": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
					},
				},
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

func resourceRouterOspfAreaVirtualLinkCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspfAreaVirtualLink(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfAreaVirtualLink resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspfAreaVirtualLink(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfAreaVirtualLink resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterOspfAreaVirtualLinkRead(d, m)
}

func resourceRouterOspfAreaVirtualLinkUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspfAreaVirtualLink(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfAreaVirtualLink resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspfAreaVirtualLink(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfAreaVirtualLink resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterOspfAreaVirtualLinkRead(d, m)
}

func resourceRouterOspfAreaVirtualLinkDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterOspfAreaVirtualLink(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspfAreaVirtualLink resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfAreaVirtualLinkRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspfAreaVirtualLink(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfAreaVirtualLink resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspfAreaVirtualLink(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfAreaVirtualLink resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspfAreaVirtualLinkAuthentication3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkDeadInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkHelloInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkKeychain3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaVirtualLinkMd5Keychain3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaVirtualLinkMd5Keys3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspfAreaVirtualLinkMd5KeysId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspfAreaVirtualLink-Md5Keys-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfAreaVirtualLinkMd5KeysId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkPeer3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkRetransmitInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkTransmitDelay3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspfAreaVirtualLink(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("authentication", flattenRouterOspfAreaVirtualLinkAuthentication3rdl(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "RouterOspfAreaVirtualLink-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("dead_interval", flattenRouterOspfAreaVirtualLinkDeadInterval3rdl(o["dead-interval"], d, "dead_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["dead-interval"], "RouterOspfAreaVirtualLink-DeadInterval"); ok {
			if err = d.Set("dead_interval", vv); err != nil {
				return fmt.Errorf("Error reading dead_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dead_interval: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenRouterOspfAreaVirtualLinkHelloInterval3rdl(o["hello-interval"], d, "hello_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-interval"], "RouterOspfAreaVirtualLink-HelloInterval"); ok {
			if err = d.Set("hello_interval", vv); err != nil {
				return fmt.Errorf("Error reading hello_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_interval: %v", err)
		}
	}

	if err = d.Set("keychain", flattenRouterOspfAreaVirtualLinkKeychain3rdl(o["keychain"], d, "keychain")); err != nil {
		if vv, ok := fortiAPIPatch(o["keychain"], "RouterOspfAreaVirtualLink-Keychain"); ok {
			if err = d.Set("keychain", vv); err != nil {
				return fmt.Errorf("Error reading keychain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keychain: %v", err)
		}
	}

	if err = d.Set("md5_keychain", flattenRouterOspfAreaVirtualLinkMd5Keychain3rdl(o["md5-keychain"], d, "md5_keychain")); err != nil {
		if vv, ok := fortiAPIPatch(o["md5-keychain"], "RouterOspfAreaVirtualLink-Md5Keychain"); ok {
			if err = d.Set("md5_keychain", vv); err != nil {
				return fmt.Errorf("Error reading md5_keychain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading md5_keychain: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("md5_keys", flattenRouterOspfAreaVirtualLinkMd5Keys3rdl(o["md5-keys"], d, "md5_keys")); err != nil {
			if vv, ok := fortiAPIPatch(o["md5-keys"], "RouterOspfAreaVirtualLink-Md5Keys"); ok {
				if err = d.Set("md5_keys", vv); err != nil {
					return fmt.Errorf("Error reading md5_keys: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading md5_keys: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("md5_keys"); ok {
			if err = d.Set("md5_keys", flattenRouterOspfAreaVirtualLinkMd5Keys3rdl(o["md5-keys"], d, "md5_keys")); err != nil {
				if vv, ok := fortiAPIPatch(o["md5-keys"], "RouterOspfAreaVirtualLink-Md5Keys"); ok {
					if err = d.Set("md5_keys", vv); err != nil {
						return fmt.Errorf("Error reading md5_keys: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading md5_keys: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenRouterOspfAreaVirtualLinkName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterOspfAreaVirtualLink-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("peer", flattenRouterOspfAreaVirtualLinkPeer3rdl(o["peer"], d, "peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer"], "RouterOspfAreaVirtualLink-Peer"); ok {
			if err = d.Set("peer", vv); err != nil {
				return fmt.Errorf("Error reading peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer: %v", err)
		}
	}

	if err = d.Set("retransmit_interval", flattenRouterOspfAreaVirtualLinkRetransmitInterval3rdl(o["retransmit-interval"], d, "retransmit_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["retransmit-interval"], "RouterOspfAreaVirtualLink-RetransmitInterval"); ok {
			if err = d.Set("retransmit_interval", vv); err != nil {
				return fmt.Errorf("Error reading retransmit_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retransmit_interval: %v", err)
		}
	}

	if err = d.Set("transmit_delay", flattenRouterOspfAreaVirtualLinkTransmitDelay3rdl(o["transmit-delay"], d, "transmit_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["transmit-delay"], "RouterOspfAreaVirtualLink-TransmitDelay"); ok {
			if err = d.Set("transmit_delay", vv); err != nil {
				return fmt.Errorf("Error reading transmit_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transmit_delay: %v", err)
		}
	}

	return nil
}

func flattenRouterOspfAreaVirtualLinkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspfAreaVirtualLinkAuthentication3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkAuthenticationKey3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaVirtualLinkDeadInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkHelloInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkKeychain3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaVirtualLinkMd5Keychain3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaVirtualLinkMd5Keys3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfAreaVirtualLinkMd5KeysId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-string"], _ = expandRouterOspfAreaVirtualLinkMd5KeysKeyString3rdl(d, i["key_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaVirtualLinkMd5KeysId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkMd5KeysKeyString3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaVirtualLinkName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkPeer3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkRetransmitInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkTransmitDelay3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspfAreaVirtualLink(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandRouterOspfAreaVirtualLinkAuthentication3rdl(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("authentication_key"); ok || d.HasChange("authentication_key") {
		t, err := expandRouterOspfAreaVirtualLinkAuthenticationKey3rdl(d, v, "authentication_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication-key"] = t
		}
	}

	if v, ok := d.GetOk("dead_interval"); ok || d.HasChange("dead_interval") {
		t, err := expandRouterOspfAreaVirtualLinkDeadInterval3rdl(d, v, "dead_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dead-interval"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval"); ok || d.HasChange("hello_interval") {
		t, err := expandRouterOspfAreaVirtualLinkHelloInterval3rdl(d, v, "hello_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval"] = t
		}
	}

	if v, ok := d.GetOk("keychain"); ok || d.HasChange("keychain") {
		t, err := expandRouterOspfAreaVirtualLinkKeychain3rdl(d, v, "keychain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keychain"] = t
		}
	}

	if v, ok := d.GetOk("md5_keychain"); ok || d.HasChange("md5_keychain") {
		t, err := expandRouterOspfAreaVirtualLinkMd5Keychain3rdl(d, v, "md5_keychain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["md5-keychain"] = t
		}
	}

	if v, ok := d.GetOk("md5_keys"); ok || d.HasChange("md5_keys") {
		t, err := expandRouterOspfAreaVirtualLinkMd5Keys3rdl(d, v, "md5_keys")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["md5-keys"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterOspfAreaVirtualLinkName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok || d.HasChange("peer") {
		t, err := expandRouterOspfAreaVirtualLinkPeer3rdl(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	if v, ok := d.GetOk("retransmit_interval"); ok || d.HasChange("retransmit_interval") {
		t, err := expandRouterOspfAreaVirtualLinkRetransmitInterval3rdl(d, v, "retransmit_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retransmit-interval"] = t
		}
	}

	if v, ok := d.GetOk("transmit_delay"); ok || d.HasChange("transmit_delay") {
		t, err := expandRouterOspfAreaVirtualLinkTransmitDelay3rdl(d, v, "transmit_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transmit-delay"] = t
		}
	}

	return &obj, nil
}
