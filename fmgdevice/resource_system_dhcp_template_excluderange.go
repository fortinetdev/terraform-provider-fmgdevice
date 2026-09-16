// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Exclude one or more ranges of IP addresses from being assigned to clients.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDhcpTemplateExcludeRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcpTemplateExcludeRangeCreate,
		Read:   resourceSystemDhcpTemplateExcludeRangeRead,
		Update: resourceSystemDhcpTemplateExcludeRangeUpdate,
		Delete: resourceSystemDhcpTemplateExcludeRangeDelete,

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
			"start_ip_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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

func resourceSystemDhcpTemplateExcludeRangeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcpTemplateExcludeRange(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcpTemplateExcludeRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemDhcpTemplateExcludeRange(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemDhcpTemplateExcludeRange(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemDhcpTemplateExcludeRange resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemDhcpTemplateExcludeRange(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemDhcpTemplateExcludeRange resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpTemplateExcludeRangeRead(d, m)
}

func resourceSystemDhcpTemplateExcludeRangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcpTemplateExcludeRange(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpTemplateExcludeRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemDhcpTemplateExcludeRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpTemplateExcludeRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcpTemplateExcludeRangeRead(d, m)
}

func resourceSystemDhcpTemplateExcludeRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemDhcpTemplateExcludeRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcpTemplateExcludeRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcpTemplateExcludeRangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDhcpTemplateExcludeRange(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemDhcpTemplateExcludeRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcpTemplateExcludeRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcpTemplateExcludeRange resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcpTemplateExcludeRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeIpCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeLeaseTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeOuiMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeOuiString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateExcludeRangeStartIpIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeUciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeUciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateExcludeRangeVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcpTemplateExcludeRangeVendor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDhcpTemplateExcludeRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemDhcpTemplateExcludeRangeId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemDhcpTemplateExcludeRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_count", flattenSystemDhcpTemplateExcludeRangeIpCount2edl(o["ip-count"], d, "ip_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-count"], "SystemDhcpTemplateExcludeRange-IpCount"); ok {
			if err = d.Set("ip_count", vv); err != nil {
				return fmt.Errorf("Error reading ip_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_count: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenSystemDhcpTemplateExcludeRangeLeaseTime2edl(o["lease-time"], d, "lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["lease-time"], "SystemDhcpTemplateExcludeRange-LeaseTime"); ok {
			if err = d.Set("lease_time", vv); err != nil {
				return fmt.Errorf("Error reading lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("oui_match", flattenSystemDhcpTemplateExcludeRangeOuiMatch2edl(o["oui-match"], d, "oui_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["oui-match"], "SystemDhcpTemplateExcludeRange-OuiMatch"); ok {
			if err = d.Set("oui_match", vv); err != nil {
				return fmt.Errorf("Error reading oui_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oui_match: %v", err)
		}
	}

	if err = d.Set("oui_string", flattenSystemDhcpTemplateExcludeRangeOuiString2edl(o["oui-string"], d, "oui_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["oui-string"], "SystemDhcpTemplateExcludeRange-OuiString"); ok {
			if err = d.Set("oui_string", vv); err != nil {
				return fmt.Errorf("Error reading oui_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oui_string: %v", err)
		}
	}

	if err = d.Set("start_ip_index", flattenSystemDhcpTemplateExcludeRangeStartIpIndex2edl(o["start-ip-index"], d, "start_ip_index")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip-index"], "SystemDhcpTemplateExcludeRange-StartIpIndex"); ok {
			if err = d.Set("start_ip_index", vv); err != nil {
				return fmt.Errorf("Error reading start_ip_index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip_index: %v", err)
		}
	}

	if err = d.Set("uci_match", flattenSystemDhcpTemplateExcludeRangeUciMatch2edl(o["uci-match"], d, "uci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["uci-match"], "SystemDhcpTemplateExcludeRange-UciMatch"); ok {
			if err = d.Set("uci_match", vv); err != nil {
				return fmt.Errorf("Error reading uci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uci_match: %v", err)
		}
	}

	if err = d.Set("uci_string", flattenSystemDhcpTemplateExcludeRangeUciString2edl(o["uci-string"], d, "uci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["uci-string"], "SystemDhcpTemplateExcludeRange-UciString"); ok {
			if err = d.Set("uci_string", vv); err != nil {
				return fmt.Errorf("Error reading uci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uci_string: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenSystemDhcpTemplateExcludeRangeVciMatch2edl(o["vci-match"], d, "vci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-match"], "SystemDhcpTemplateExcludeRange-VciMatch"); ok {
			if err = d.Set("vci_match", vv); err != nil {
				return fmt.Errorf("Error reading vci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

	if err = d.Set("vci_string", flattenSystemDhcpTemplateExcludeRangeVciString2edl(o["vci-string"], d, "vci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-string"], "SystemDhcpTemplateExcludeRange-VciString"); ok {
			if err = d.Set("vci_string", vv); err != nil {
				return fmt.Errorf("Error reading vci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_string: %v", err)
		}
	}

	if err = d.Set("vendor", flattenSystemDhcpTemplateExcludeRangeVendor2edl(o["vendor"], d, "vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["vendor"], "SystemDhcpTemplateExcludeRange-Vendor"); ok {
			if err = d.Set("vendor", vv); err != nil {
				return fmt.Errorf("Error reading vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	return nil
}

func flattenSystemDhcpTemplateExcludeRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDhcpTemplateExcludeRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeIpCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeLeaseTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeOuiMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeOuiString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateExcludeRangeStartIpIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeUciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeUciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateExcludeRangeVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcpTemplateExcludeRangeVendor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDhcpTemplateExcludeRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemDhcpTemplateExcludeRangeId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_count"); ok || d.HasChange("ip_count") {
		t, err := expandSystemDhcpTemplateExcludeRangeIpCount2edl(d, v, "ip_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-count"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok || d.HasChange("lease_time") {
		t, err := expandSystemDhcpTemplateExcludeRangeLeaseTime2edl(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("oui_match"); ok || d.HasChange("oui_match") {
		t, err := expandSystemDhcpTemplateExcludeRangeOuiMatch2edl(d, v, "oui_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oui-match"] = t
		}
	}

	if v, ok := d.GetOk("oui_string"); ok || d.HasChange("oui_string") {
		t, err := expandSystemDhcpTemplateExcludeRangeOuiString2edl(d, v, "oui_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oui-string"] = t
		}
	}

	if v, ok := d.GetOk("start_ip_index"); ok || d.HasChange("start_ip_index") {
		t, err := expandSystemDhcpTemplateExcludeRangeStartIpIndex2edl(d, v, "start_ip_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip-index"] = t
		}
	}

	if v, ok := d.GetOk("uci_match"); ok || d.HasChange("uci_match") {
		t, err := expandSystemDhcpTemplateExcludeRangeUciMatch2edl(d, v, "uci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uci-match"] = t
		}
	}

	if v, ok := d.GetOk("uci_string"); ok || d.HasChange("uci_string") {
		t, err := expandSystemDhcpTemplateExcludeRangeUciString2edl(d, v, "uci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uci-string"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok || d.HasChange("vci_match") {
		t, err := expandSystemDhcpTemplateExcludeRangeVciMatch2edl(d, v, "vci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok || d.HasChange("vci_string") {
		t, err := expandSystemDhcpTemplateExcludeRangeVciString2edl(d, v, "vci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok || d.HasChange("vendor") {
		t, err := expandSystemDhcpTemplateExcludeRangeVendor2edl(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	return &obj, nil
}
