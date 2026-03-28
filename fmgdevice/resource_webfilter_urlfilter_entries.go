// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> URL filter entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterUrlfilterEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterUrlfilterEntriesCreate,
		Read:   resourceWebfilterUrlfilterEntriesRead,
		Update: resourceWebfilterUrlfilterEntriesUpdate,
		Delete: resourceWebfilterUrlfilterEntriesDelete,

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
			"urlfilter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antiphish_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_address_family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exempt": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"referrer_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_proxy_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebfilterUrlfilterEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	urlfilter := d.Get("urlfilter").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["urlfilter"] = urlfilter

	obj, err := getObjectWebfilterUrlfilterEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterUrlfilterEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebfilterUrlfilterEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebfilterUrlfilterEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebfilterUrlfilterEntries resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateWebfilterUrlfilterEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebfilterUrlfilterEntries resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceWebfilterUrlfilterEntriesRead(d, m)
			} else {
				return fmt.Errorf("Error creating WebfilterUrlfilterEntries resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebfilterUrlfilterEntriesRead(d, m)
}

func resourceWebfilterUrlfilterEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	urlfilter := d.Get("urlfilter").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["urlfilter"] = urlfilter

	obj, err := getObjectWebfilterUrlfilterEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterUrlfilterEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterUrlfilterEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterUrlfilterEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebfilterUrlfilterEntriesRead(d, m)
}

func resourceWebfilterUrlfilterEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	urlfilter := d.Get("urlfilter").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["urlfilter"] = urlfilter

	wsParams["adom"] = adomv

	err = c.DeleteWebfilterUrlfilterEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterUrlfilterEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterUrlfilterEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	urlfilter := d.Get("urlfilter").(string)
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
	if urlfilter == "" {
		urlfilter = importOptionChecking(m.(*FortiClient).Cfg, "urlfilter")
		if urlfilter == "" {
			return fmt.Errorf("Parameter urlfilter is missing")
		}
		if err = d.Set("urlfilter", urlfilter); err != nil {
			return fmt.Errorf("Error set params urlfilter: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["urlfilter"] = urlfilter

	o, err := c.ReadWebfilterUrlfilterEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterUrlfilterEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterUrlfilterEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterUrlfilterEntries resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterUrlfilterEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesAntiphishAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesDnsAddressFamily2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesExempt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterUrlfilterEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesReferrerHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesWebProxyProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWebfilterUrlfilterEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenWebfilterUrlfilterEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "WebfilterUrlfilterEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("antiphish_action", flattenWebfilterUrlfilterEntriesAntiphishAction2edl(o["antiphish-action"], d, "antiphish_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["antiphish-action"], "WebfilterUrlfilterEntries-AntiphishAction"); ok {
			if err = d.Set("antiphish_action", vv); err != nil {
				return fmt.Errorf("Error reading antiphish_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antiphish_action: %v", err)
		}
	}

	if err = d.Set("comment", flattenWebfilterUrlfilterEntriesComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WebfilterUrlfilterEntries-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("dns_address_family", flattenWebfilterUrlfilterEntriesDnsAddressFamily2edl(o["dns-address-family"], d, "dns_address_family")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-address-family"], "WebfilterUrlfilterEntries-DnsAddressFamily"); ok {
			if err = d.Set("dns_address_family", vv); err != nil {
				return fmt.Errorf("Error reading dns_address_family: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_address_family: %v", err)
		}
	}

	if err = d.Set("exempt", flattenWebfilterUrlfilterEntriesExempt2edl(o["exempt"], d, "exempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["exempt"], "WebfilterUrlfilterEntries-Exempt"); ok {
			if err = d.Set("exempt", vv); err != nil {
				return fmt.Errorf("Error reading exempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exempt: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWebfilterUrlfilterEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WebfilterUrlfilterEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("referrer_host", flattenWebfilterUrlfilterEntriesReferrerHost2edl(o["referrer-host"], d, "referrer_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["referrer-host"], "WebfilterUrlfilterEntries-ReferrerHost"); ok {
			if err = d.Set("referrer_host", vv); err != nil {
				return fmt.Errorf("Error reading referrer_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading referrer_host: %v", err)
		}
	}

	if err = d.Set("status", flattenWebfilterUrlfilterEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WebfilterUrlfilterEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenWebfilterUrlfilterEntriesType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "WebfilterUrlfilterEntries-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("url", flattenWebfilterUrlfilterEntriesUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "WebfilterUrlfilterEntries-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("web_proxy_profile", flattenWebfilterUrlfilterEntriesWebProxyProfile2edl(o["web-proxy-profile"], d, "web_proxy_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-proxy-profile"], "WebfilterUrlfilterEntries-WebProxyProfile"); ok {
			if err = d.Set("web_proxy_profile", vv); err != nil {
				return fmt.Errorf("Error reading web_proxy_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_proxy_profile: %v", err)
		}
	}

	return nil
}

func flattenWebfilterUrlfilterEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterUrlfilterEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesAntiphishAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesDnsAddressFamily2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesExempt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterUrlfilterEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesReferrerHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesWebProxyProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWebfilterUrlfilterEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandWebfilterUrlfilterEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("antiphish_action"); ok || d.HasChange("antiphish_action") {
		t, err := expandWebfilterUrlfilterEntriesAntiphishAction2edl(d, v, "antiphish_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antiphish-action"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWebfilterUrlfilterEntriesComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dns_address_family"); ok || d.HasChange("dns_address_family") {
		t, err := expandWebfilterUrlfilterEntriesDnsAddressFamily2edl(d, v, "dns_address_family")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-address-family"] = t
		}
	}

	if v, ok := d.GetOk("exempt"); ok || d.HasChange("exempt") {
		t, err := expandWebfilterUrlfilterEntriesExempt2edl(d, v, "exempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exempt"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWebfilterUrlfilterEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("referrer_host"); ok || d.HasChange("referrer_host") {
		t, err := expandWebfilterUrlfilterEntriesReferrerHost2edl(d, v, "referrer_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["referrer-host"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWebfilterUrlfilterEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandWebfilterUrlfilterEntriesType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandWebfilterUrlfilterEntriesUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("web_proxy_profile"); ok || d.HasChange("web_proxy_profile") {
		t, err := expandWebfilterUrlfilterEntriesWebProxyProfile2edl(d, v, "web_proxy_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-proxy-profile"] = t
		}
	}

	return &obj, nil
}
