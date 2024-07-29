// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure EVPN instance.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemEvpn() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemEvpnCreate,
		Read:   resourceSystemEvpnRead,
		Update: resourceSystemEvpnUpdate,
		Delete: resourceSystemEvpnDelete,

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
			"arp_suppression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"export_rt": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"import_rt": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_local_learning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemEvpnCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectSystemEvpn(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemEvpn resource while getting object: %v", err)
	}

	_, err = c.CreateSystemEvpn(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemEvpn resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemEvpnRead(d, m)
}

func resourceSystemEvpnUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectSystemEvpn(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemEvpn resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemEvpn(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemEvpn resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemEvpnRead(d, m)
}

func resourceSystemEvpnDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteSystemEvpn(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemEvpn resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemEvpnRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemEvpn(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemEvpn resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemEvpn(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemEvpn resource from API: %v", err)
	}
	return nil
}

func flattenSystemEvpnArpSuppression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemEvpnExportRt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemEvpnId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemEvpnImportRt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemEvpnIpLocalLearning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemEvpnRd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemEvpn(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("arp_suppression", flattenSystemEvpnArpSuppression(o["arp-suppression"], d, "arp_suppression")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-suppression"], "SystemEvpn-ArpSuppression"); ok {
			if err = d.Set("arp_suppression", vv); err != nil {
				return fmt.Errorf("Error reading arp_suppression: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_suppression: %v", err)
		}
	}

	if err = d.Set("export_rt", flattenSystemEvpnExportRt(o["export-rt"], d, "export_rt")); err != nil {
		if vv, ok := fortiAPIPatch(o["export-rt"], "SystemEvpn-ExportRt"); ok {
			if err = d.Set("export_rt", vv); err != nil {
				return fmt.Errorf("Error reading export_rt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading export_rt: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemEvpnId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemEvpn-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("import_rt", flattenSystemEvpnImportRt(o["import-rt"], d, "import_rt")); err != nil {
		if vv, ok := fortiAPIPatch(o["import-rt"], "SystemEvpn-ImportRt"); ok {
			if err = d.Set("import_rt", vv); err != nil {
				return fmt.Errorf("Error reading import_rt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading import_rt: %v", err)
		}
	}

	if err = d.Set("ip_local_learning", flattenSystemEvpnIpLocalLearning(o["ip-local-learning"], d, "ip_local_learning")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-local-learning"], "SystemEvpn-IpLocalLearning"); ok {
			if err = d.Set("ip_local_learning", vv); err != nil {
				return fmt.Errorf("Error reading ip_local_learning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_local_learning: %v", err)
		}
	}

	if err = d.Set("rd", flattenSystemEvpnRd(o["rd"], d, "rd")); err != nil {
		if vv, ok := fortiAPIPatch(o["rd"], "SystemEvpn-Rd"); ok {
			if err = d.Set("rd", vv); err != nil {
				return fmt.Errorf("Error reading rd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rd: %v", err)
		}
	}

	return nil
}

func flattenSystemEvpnFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemEvpnArpSuppression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnExportRt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemEvpnId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnImportRt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemEvpnIpLocalLearning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnRd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemEvpn(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("arp_suppression"); ok || d.HasChange("arp_suppression") {
		t, err := expandSystemEvpnArpSuppression(d, v, "arp_suppression")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-suppression"] = t
		}
	}

	if v, ok := d.GetOk("export_rt"); ok || d.HasChange("export_rt") {
		t, err := expandSystemEvpnExportRt(d, v, "export_rt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["export-rt"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemEvpnId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("import_rt"); ok || d.HasChange("import_rt") {
		t, err := expandSystemEvpnImportRt(d, v, "import_rt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["import-rt"] = t
		}
	}

	if v, ok := d.GetOk("ip_local_learning"); ok || d.HasChange("ip_local_learning") {
		t, err := expandSystemEvpnIpLocalLearning(d, v, "ip_local_learning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-local-learning"] = t
		}
	}

	if v, ok := d.GetOk("rd"); ok || d.HasChange("rd") {
		t, err := expandSystemEvpnRd(d, v, "rd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rd"] = t
		}
	}

	return &obj, nil
}
