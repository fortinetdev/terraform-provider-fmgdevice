// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DHCP IP range configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDhcpTemplateIpRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcpTemplateIpRangeCreate,
		Read:   resourceSystemDhcpTemplateIpRangeRead,
		Update: resourceSystemDhcpTemplateIpRangeUpdate,
		Delete: resourceSystemDhcpTemplateIpRangeDelete,

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
			"template": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"lease_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"oui_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oui_string": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"reserve": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uci_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uci_string": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vci_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vci_string": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemDhcpTemplateIpRangeCreate(d *schema.ResourceData, m interface{}) error {
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
	template := d.Get("template").(string)
	paradict["device"] = device_name
	paradict["template"] = template

	obj, err := getObjectSystemDhcpTemplateIpRange(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcpTemplateIpRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemDhcpTemplateIpRange(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemDhcpTemplateIpRange(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemDhcpTemplateIpRange resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemDhcpTemplateIpRange(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemDhcpTemplateIpRange resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpTemplateIpRangeRead(d, m)
}

func resourceSystemDhcpTemplateIpRangeUpdate(d *schema.ResourceData, m interface{}) error {
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
	template := d.Get("template").(string)
	paradict["device"] = device_name
	paradict["template"] = template

	obj, err := getObjectSystemDhcpTemplateIpRange(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpTemplateIpRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemDhcpTemplateIpRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpTemplateIpRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpTemplateIpRangeRead(d, m)
}

func resourceSystemDhcpTemplateIpRangeDelete(d *schema.ResourceData, m interface{}) error {
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
	template := d.Get("template").(string)
	paradict["device"] = device_name
	paradict["template"] = template

	wsParams["adom"] = adomv

	err = c.DeleteSystemDhcpTemplateIpRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcpTemplateIpRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcpTemplateIpRangeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	template := d.Get("template").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if template == "" {
		template = importOptionChecking(m.(*FortiClient).Cfg, "template")
		if template == "" {
			return fmt.Errorf("Parameter template is missing")
		}
		if err = d.Set("template", template); err != nil {
			return fmt.Errorf("Error set params template: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["template"] = template

	o, err := c.ReadSystemDhcpTemplateIpRange(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemDhcpTemplateIpRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcpTemplateIpRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcpTemplateIpRange resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcpTemplateIpRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeIpCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeLeaseTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeOuiMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeOuiString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateIpRangeReserve2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeUciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeUciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateIpRangeVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateIpRangeVendor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDhcpTemplateIpRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemDhcpTemplateIpRangeId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemDhcpTemplateIpRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_count", flattenSystemDhcpTemplateIpRangeIpCount2edl(o["ip-count"], d, "ip_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-count"], "SystemDhcpTemplateIpRange-IpCount"); ok {
			if err = d.Set("ip_count", vv); err != nil {
				return fmt.Errorf("Error reading ip_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_count: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenSystemDhcpTemplateIpRangeLeaseTime2edl(o["lease-time"], d, "lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["lease-time"], "SystemDhcpTemplateIpRange-LeaseTime"); ok {
			if err = d.Set("lease_time", vv); err != nil {
				return fmt.Errorf("Error reading lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("oui_match", flattenSystemDhcpTemplateIpRangeOuiMatch2edl(o["oui-match"], d, "oui_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["oui-match"], "SystemDhcpTemplateIpRange-OuiMatch"); ok {
			if err = d.Set("oui_match", vv); err != nil {
				return fmt.Errorf("Error reading oui_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oui_match: %v", err)
		}
	}

	if err = d.Set("oui_string", flattenSystemDhcpTemplateIpRangeOuiString2edl(o["oui-string"], d, "oui_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["oui-string"], "SystemDhcpTemplateIpRange-OuiString"); ok {
			if err = d.Set("oui_string", vv); err != nil {
				return fmt.Errorf("Error reading oui_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oui_string: %v", err)
		}
	}

	if err = d.Set("reserve", flattenSystemDhcpTemplateIpRangeReserve2edl(o["reserve"], d, "reserve")); err != nil {
		if vv, ok := fortiAPIPatch(o["reserve"], "SystemDhcpTemplateIpRange-Reserve"); ok {
			if err = d.Set("reserve", vv); err != nil {
				return fmt.Errorf("Error reading reserve: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reserve: %v", err)
		}
	}

	if err = d.Set("uci_match", flattenSystemDhcpTemplateIpRangeUciMatch2edl(o["uci-match"], d, "uci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["uci-match"], "SystemDhcpTemplateIpRange-UciMatch"); ok {
			if err = d.Set("uci_match", vv); err != nil {
				return fmt.Errorf("Error reading uci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uci_match: %v", err)
		}
	}

	if err = d.Set("uci_string", flattenSystemDhcpTemplateIpRangeUciString2edl(o["uci-string"], d, "uci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["uci-string"], "SystemDhcpTemplateIpRange-UciString"); ok {
			if err = d.Set("uci_string", vv); err != nil {
				return fmt.Errorf("Error reading uci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uci_string: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenSystemDhcpTemplateIpRangeVciMatch2edl(o["vci-match"], d, "vci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-match"], "SystemDhcpTemplateIpRange-VciMatch"); ok {
			if err = d.Set("vci_match", vv); err != nil {
				return fmt.Errorf("Error reading vci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

	if err = d.Set("vci_string", flattenSystemDhcpTemplateIpRangeVciString2edl(o["vci-string"], d, "vci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-string"], "SystemDhcpTemplateIpRange-VciString"); ok {
			if err = d.Set("vci_string", vv); err != nil {
				return fmt.Errorf("Error reading vci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_string: %v", err)
		}
	}

	if err = d.Set("vendor", flattenSystemDhcpTemplateIpRangeVendor2edl(o["vendor"], d, "vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["vendor"], "SystemDhcpTemplateIpRange-Vendor"); ok {
			if err = d.Set("vendor", vv); err != nil {
				return fmt.Errorf("Error reading vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	return nil
}

func flattenSystemDhcpTemplateIpRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDhcpTemplateIpRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeIpCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeLeaseTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeOuiMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeOuiString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateIpRangeReserve2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeUciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeUciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateIpRangeVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateIpRangeVendor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDhcpTemplateIpRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemDhcpTemplateIpRangeId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_count"); ok || d.HasChange("ip_count") {
		t, err := expandSystemDhcpTemplateIpRangeIpCount2edl(d, v, "ip_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-count"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok || d.HasChange("lease_time") {
		t, err := expandSystemDhcpTemplateIpRangeLeaseTime2edl(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("oui_match"); ok || d.HasChange("oui_match") {
		t, err := expandSystemDhcpTemplateIpRangeOuiMatch2edl(d, v, "oui_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oui-match"] = t
		}
	}

	if v, ok := d.GetOk("oui_string"); ok || d.HasChange("oui_string") {
		t, err := expandSystemDhcpTemplateIpRangeOuiString2edl(d, v, "oui_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oui-string"] = t
		}
	}

	if v, ok := d.GetOk("reserve"); ok || d.HasChange("reserve") {
		t, err := expandSystemDhcpTemplateIpRangeReserve2edl(d, v, "reserve")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reserve"] = t
		}
	}

	if v, ok := d.GetOk("uci_match"); ok || d.HasChange("uci_match") {
		t, err := expandSystemDhcpTemplateIpRangeUciMatch2edl(d, v, "uci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uci-match"] = t
		}
	}

	if v, ok := d.GetOk("uci_string"); ok || d.HasChange("uci_string") {
		t, err := expandSystemDhcpTemplateIpRangeUciString2edl(d, v, "uci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uci-string"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok || d.HasChange("vci_match") {
		t, err := expandSystemDhcpTemplateIpRangeVciMatch2edl(d, v, "vci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok || d.HasChange("vci_string") {
		t, err := expandSystemDhcpTemplateIpRangeVciString2edl(d, v, "vci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok || d.HasChange("vendor") {
		t, err := expandSystemDhcpTemplateIpRangeVendor2edl(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	return &obj, nil
}
