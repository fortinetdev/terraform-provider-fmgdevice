// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> DNS translation settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDnsfilterProfileDnsTranslation() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsfilterProfileDnsTranslationCreate,
		Read:   resourceDnsfilterProfileDnsTranslationRead,
		Update: resourceDnsfilterProfileDnsTranslationUpdate,
		Delete: resourceDnsfilterProfileDnsTranslationDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDnsfilterProfileDnsTranslationCreate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectDnsfilterProfileDnsTranslation(d)
	if err != nil {
		return fmt.Errorf("Error creating DnsfilterProfileDnsTranslation resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDnsfilterProfileDnsTranslation(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDnsfilterProfileDnsTranslation(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DnsfilterProfileDnsTranslation resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateDnsfilterProfileDnsTranslation(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DnsfilterProfileDnsTranslation resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceDnsfilterProfileDnsTranslationRead(d, m)
			} else {
				return fmt.Errorf("Error creating DnsfilterProfileDnsTranslation resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDnsfilterProfileDnsTranslationRead(d, m)
}

func resourceDnsfilterProfileDnsTranslationUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectDnsfilterProfileDnsTranslation(d)
	if err != nil {
		return fmt.Errorf("Error updating DnsfilterProfileDnsTranslation resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDnsfilterProfileDnsTranslation(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DnsfilterProfileDnsTranslation resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDnsfilterProfileDnsTranslationRead(d, m)
}

func resourceDnsfilterProfileDnsTranslationDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteDnsfilterProfileDnsTranslation(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DnsfilterProfileDnsTranslation resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDnsfilterProfileDnsTranslationRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadDnsfilterProfileDnsTranslation(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DnsfilterProfileDnsTranslation resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDnsfilterProfileDnsTranslation(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DnsfilterProfileDnsTranslation resource from API: %v", err)
	}
	return nil
}

func flattenDnsfilterProfileDnsTranslationAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationDst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationDst62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationNetmask2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationSrc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationSrc62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDnsfilterProfileDnsTranslation(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_type", flattenDnsfilterProfileDnsTranslationAddrType2edl(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "DnsfilterProfileDnsTranslation-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("dst", flattenDnsfilterProfileDnsTranslationDst2edl(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "DnsfilterProfileDnsTranslation-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("dst6", flattenDnsfilterProfileDnsTranslationDst62edl(o["dst6"], d, "dst6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst6"], "DnsfilterProfileDnsTranslation-Dst6"); ok {
			if err = d.Set("dst6", vv); err != nil {
				return fmt.Errorf("Error reading dst6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst6: %v", err)
		}
	}

	if err = d.Set("fosid", flattenDnsfilterProfileDnsTranslationId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "DnsfilterProfileDnsTranslation-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("netmask", flattenDnsfilterProfileDnsTranslationNetmask2edl(o["netmask"], d, "netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["netmask"], "DnsfilterProfileDnsTranslation-Netmask"); ok {
			if err = d.Set("netmask", vv); err != nil {
				return fmt.Errorf("Error reading netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netmask: %v", err)
		}
	}

	if err = d.Set("prefix", flattenDnsfilterProfileDnsTranslationPrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "DnsfilterProfileDnsTranslation-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("src", flattenDnsfilterProfileDnsTranslationSrc2edl(o["src"], d, "src")); err != nil {
		if vv, ok := fortiAPIPatch(o["src"], "DnsfilterProfileDnsTranslation-Src"); ok {
			if err = d.Set("src", vv); err != nil {
				return fmt.Errorf("Error reading src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("src6", flattenDnsfilterProfileDnsTranslationSrc62edl(o["src6"], d, "src6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src6"], "DnsfilterProfileDnsTranslation-Src6"); ok {
			if err = d.Set("src6", vv); err != nil {
				return fmt.Errorf("Error reading src6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src6: %v", err)
		}
	}

	if err = d.Set("status", flattenDnsfilterProfileDnsTranslationStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "DnsfilterProfileDnsTranslation-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenDnsfilterProfileDnsTranslationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDnsfilterProfileDnsTranslationAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationDst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationDst62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationNetmask2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationSrc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationSrc62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDnsfilterProfileDnsTranslation(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandDnsfilterProfileDnsTranslationAddrType2edl(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandDnsfilterProfileDnsTranslationDst2edl(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("dst6"); ok || d.HasChange("dst6") {
		t, err := expandDnsfilterProfileDnsTranslationDst62edl(d, v, "dst6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst6"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandDnsfilterProfileDnsTranslationId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("netmask"); ok || d.HasChange("netmask") {
		t, err := expandDnsfilterProfileDnsTranslationNetmask2edl(d, v, "netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netmask"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandDnsfilterProfileDnsTranslationPrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandDnsfilterProfileDnsTranslationSrc2edl(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("src6"); ok || d.HasChange("src6") {
		t, err := expandDnsfilterProfileDnsTranslationSrc62edl(d, v, "src6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src6"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandDnsfilterProfileDnsTranslationStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
