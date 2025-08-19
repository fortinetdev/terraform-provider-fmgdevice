// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Create SD-WAN duplication rule.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdwanDuplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdwanDuplicationCreate,
		Read:   resourceSystemSdwanDuplicationRead,
		Update: resourceSystemSdwanDuplicationUpdate,
		Delete: resourceSystemSdwanDuplicationDelete,

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
			"dstaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstintf": &schema.Schema{
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
			"packet_de_duplication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"packet_duplication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sla_match_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSdwanDuplicationCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSdwanDuplication(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanDuplication resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSdwanDuplication(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanDuplication resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSdwanDuplicationRead(d, m)
}

func resourceSystemSdwanDuplicationUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSdwanDuplication(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanDuplication resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSdwanDuplication(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanDuplication resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSdwanDuplicationRead(d, m)
}

func resourceSystemSdwanDuplicationDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSdwanDuplication(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdwanDuplication resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanDuplicationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdwanDuplication(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanDuplication resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdwanDuplication(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanDuplication resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdwanDuplicationDstaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationDstaddr62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationDstintf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationPacketDeDuplication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationPacketDuplication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationServiceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationSlaMatchService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationSrcaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationSrcaddr62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationSrcintf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemSdwanDuplication(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dstaddr", flattenSystemSdwanDuplicationDstaddr2edl(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "SystemSdwanDuplication-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenSystemSdwanDuplicationDstaddr62edl(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "SystemSdwanDuplication-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenSystemSdwanDuplicationDstintf2edl(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "SystemSdwanDuplication-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemSdwanDuplicationId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSdwanDuplication-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("packet_de_duplication", flattenSystemSdwanDuplicationPacketDeDuplication2edl(o["packet-de-duplication"], d, "packet_de_duplication")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-de-duplication"], "SystemSdwanDuplication-PacketDeDuplication"); ok {
			if err = d.Set("packet_de_duplication", vv); err != nil {
				return fmt.Errorf("Error reading packet_de_duplication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_de_duplication: %v", err)
		}
	}

	if err = d.Set("packet_duplication", flattenSystemSdwanDuplicationPacketDuplication2edl(o["packet-duplication"], d, "packet_duplication")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-duplication"], "SystemSdwanDuplication-PacketDuplication"); ok {
			if err = d.Set("packet_duplication", vv); err != nil {
				return fmt.Errorf("Error reading packet_duplication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_duplication: %v", err)
		}
	}

	if err = d.Set("service", flattenSystemSdwanDuplicationService2edl(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "SystemSdwanDuplication-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("service_id", flattenSystemSdwanDuplicationServiceId2edl(o["service-id"], d, "service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-id"], "SystemSdwanDuplication-ServiceId"); ok {
			if err = d.Set("service_id", vv); err != nil {
				return fmt.Errorf("Error reading service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_id: %v", err)
		}
	}

	if err = d.Set("sla_match_service", flattenSystemSdwanDuplicationSlaMatchService2edl(o["sla-match-service"], d, "sla_match_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-match-service"], "SystemSdwanDuplication-SlaMatchService"); ok {
			if err = d.Set("sla_match_service", vv); err != nil {
				return fmt.Errorf("Error reading sla_match_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_match_service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenSystemSdwanDuplicationSrcaddr2edl(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "SystemSdwanDuplication-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenSystemSdwanDuplicationSrcaddr62edl(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "SystemSdwanDuplication-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenSystemSdwanDuplicationSrcintf2edl(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "SystemSdwanDuplication-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	return nil
}

func flattenSystemSdwanDuplicationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdwanDuplicationDstaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationDstaddr62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationDstintf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationPacketDeDuplication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationPacketDuplication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationServiceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationSlaMatchService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationSrcaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationSrcaddr62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationSrcintf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemSdwanDuplication(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandSystemSdwanDuplicationDstaddr2edl(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandSystemSdwanDuplicationDstaddr62edl(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandSystemSdwanDuplicationDstintf2edl(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSdwanDuplicationId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("packet_de_duplication"); ok || d.HasChange("packet_de_duplication") {
		t, err := expandSystemSdwanDuplicationPacketDeDuplication2edl(d, v, "packet_de_duplication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-de-duplication"] = t
		}
	}

	if v, ok := d.GetOk("packet_duplication"); ok || d.HasChange("packet_duplication") {
		t, err := expandSystemSdwanDuplicationPacketDuplication2edl(d, v, "packet_duplication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-duplication"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandSystemSdwanDuplicationService2edl(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_id"); ok || d.HasChange("service_id") {
		t, err := expandSystemSdwanDuplicationServiceId2edl(d, v, "service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-id"] = t
		}
	}

	if v, ok := d.GetOk("sla_match_service"); ok || d.HasChange("sla_match_service") {
		t, err := expandSystemSdwanDuplicationSlaMatchService2edl(d, v, "sla_match_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-match-service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandSystemSdwanDuplicationSrcaddr2edl(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandSystemSdwanDuplicationSrcaddr62edl(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandSystemSdwanDuplicationSrcintf2edl(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	return &obj, nil
}
