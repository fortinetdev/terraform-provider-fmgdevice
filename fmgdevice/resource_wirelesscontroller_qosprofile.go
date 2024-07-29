// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WiFi quality of service (QoS) profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerQosProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerQosProfileCreate,
		Read:   resourceWirelessControllerQosProfileRead,
		Update: resourceWirelessControllerQosProfileUpdate,
		Delete: resourceWirelessControllerQosProfileDelete,

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
			"bandwidth_admission_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bandwidth_capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"burst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"call_admission_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"call_capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"downlink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"downlink_sta": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dscp_wmm_be": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"dscp_wmm_bk": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"dscp_wmm_mapping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp_wmm_vi": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"dscp_wmm_vo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"uplink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uplink_sta": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wmm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wmm_be_dscp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wmm_bk_dscp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wmm_dscp_marking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wmm_uapsd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wmm_vi_dscp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wmm_vo_dscp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerQosProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerQosProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerQosProfile resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerQosProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerQosProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerQosProfileRead(d, m)
}

func resourceWirelessControllerQosProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerQosProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerQosProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerQosProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerQosProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerQosProfileRead(d, m)
}

func resourceWirelessControllerQosProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerQosProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerQosProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerQosProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerQosProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerQosProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerQosProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerQosProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerQosProfileBandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileBandwidthCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileBurst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileCallAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileCallCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDownlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDownlinkSta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDscpWmmBe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWirelessControllerQosProfileDscpWmmBk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWirelessControllerQosProfileDscpWmmMapping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDscpWmmVi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWirelessControllerQosProfileDscpWmmVo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWirelessControllerQosProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileUplink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileUplinkSta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmBeDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmBkDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmDscpMarking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmUapsd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmViDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmVoDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerQosProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bandwidth_admission_control", flattenWirelessControllerQosProfileBandwidthAdmissionControl(o["bandwidth-admission-control"], d, "bandwidth_admission_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-admission-control"], "WirelessControllerQosProfile-BandwidthAdmissionControl"); ok {
			if err = d.Set("bandwidth_admission_control", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_admission_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_admission_control: %v", err)
		}
	}

	if err = d.Set("bandwidth_capacity", flattenWirelessControllerQosProfileBandwidthCapacity(o["bandwidth-capacity"], d, "bandwidth_capacity")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-capacity"], "WirelessControllerQosProfile-BandwidthCapacity"); ok {
			if err = d.Set("bandwidth_capacity", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_capacity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_capacity: %v", err)
		}
	}

	if err = d.Set("burst", flattenWirelessControllerQosProfileBurst(o["burst"], d, "burst")); err != nil {
		if vv, ok := fortiAPIPatch(o["burst"], "WirelessControllerQosProfile-Burst"); ok {
			if err = d.Set("burst", vv); err != nil {
				return fmt.Errorf("Error reading burst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading burst: %v", err)
		}
	}

	if err = d.Set("call_admission_control", flattenWirelessControllerQosProfileCallAdmissionControl(o["call-admission-control"], d, "call_admission_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-admission-control"], "WirelessControllerQosProfile-CallAdmissionControl"); ok {
			if err = d.Set("call_admission_control", vv); err != nil {
				return fmt.Errorf("Error reading call_admission_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_admission_control: %v", err)
		}
	}

	if err = d.Set("call_capacity", flattenWirelessControllerQosProfileCallCapacity(o["call-capacity"], d, "call_capacity")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-capacity"], "WirelessControllerQosProfile-CallCapacity"); ok {
			if err = d.Set("call_capacity", vv); err != nil {
				return fmt.Errorf("Error reading call_capacity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_capacity: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerQosProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerQosProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("downlink", flattenWirelessControllerQosProfileDownlink(o["downlink"], d, "downlink")); err != nil {
		if vv, ok := fortiAPIPatch(o["downlink"], "WirelessControllerQosProfile-Downlink"); ok {
			if err = d.Set("downlink", vv); err != nil {
				return fmt.Errorf("Error reading downlink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading downlink: %v", err)
		}
	}

	if err = d.Set("downlink_sta", flattenWirelessControllerQosProfileDownlinkSta(o["downlink-sta"], d, "downlink_sta")); err != nil {
		if vv, ok := fortiAPIPatch(o["downlink-sta"], "WirelessControllerQosProfile-DownlinkSta"); ok {
			if err = d.Set("downlink_sta", vv); err != nil {
				return fmt.Errorf("Error reading downlink_sta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading downlink_sta: %v", err)
		}
	}

	if err = d.Set("dscp_wmm_be", flattenWirelessControllerQosProfileDscpWmmBe(o["dscp-wmm-be"], d, "dscp_wmm_be")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-wmm-be"], "WirelessControllerQosProfile-DscpWmmBe"); ok {
			if err = d.Set("dscp_wmm_be", vv); err != nil {
				return fmt.Errorf("Error reading dscp_wmm_be: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_wmm_be: %v", err)
		}
	}

	if err = d.Set("dscp_wmm_bk", flattenWirelessControllerQosProfileDscpWmmBk(o["dscp-wmm-bk"], d, "dscp_wmm_bk")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-wmm-bk"], "WirelessControllerQosProfile-DscpWmmBk"); ok {
			if err = d.Set("dscp_wmm_bk", vv); err != nil {
				return fmt.Errorf("Error reading dscp_wmm_bk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_wmm_bk: %v", err)
		}
	}

	if err = d.Set("dscp_wmm_mapping", flattenWirelessControllerQosProfileDscpWmmMapping(o["dscp-wmm-mapping"], d, "dscp_wmm_mapping")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-wmm-mapping"], "WirelessControllerQosProfile-DscpWmmMapping"); ok {
			if err = d.Set("dscp_wmm_mapping", vv); err != nil {
				return fmt.Errorf("Error reading dscp_wmm_mapping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_wmm_mapping: %v", err)
		}
	}

	if err = d.Set("dscp_wmm_vi", flattenWirelessControllerQosProfileDscpWmmVi(o["dscp-wmm-vi"], d, "dscp_wmm_vi")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-wmm-vi"], "WirelessControllerQosProfile-DscpWmmVi"); ok {
			if err = d.Set("dscp_wmm_vi", vv); err != nil {
				return fmt.Errorf("Error reading dscp_wmm_vi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_wmm_vi: %v", err)
		}
	}

	if err = d.Set("dscp_wmm_vo", flattenWirelessControllerQosProfileDscpWmmVo(o["dscp-wmm-vo"], d, "dscp_wmm_vo")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-wmm-vo"], "WirelessControllerQosProfile-DscpWmmVo"); ok {
			if err = d.Set("dscp_wmm_vo", vv); err != nil {
				return fmt.Errorf("Error reading dscp_wmm_vo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_wmm_vo: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerQosProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerQosProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uplink", flattenWirelessControllerQosProfileUplink(o["uplink"], d, "uplink")); err != nil {
		if vv, ok := fortiAPIPatch(o["uplink"], "WirelessControllerQosProfile-Uplink"); ok {
			if err = d.Set("uplink", vv); err != nil {
				return fmt.Errorf("Error reading uplink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uplink: %v", err)
		}
	}

	if err = d.Set("uplink_sta", flattenWirelessControllerQosProfileUplinkSta(o["uplink-sta"], d, "uplink_sta")); err != nil {
		if vv, ok := fortiAPIPatch(o["uplink-sta"], "WirelessControllerQosProfile-UplinkSta"); ok {
			if err = d.Set("uplink_sta", vv); err != nil {
				return fmt.Errorf("Error reading uplink_sta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uplink_sta: %v", err)
		}
	}

	if err = d.Set("wmm", flattenWirelessControllerQosProfileWmm(o["wmm"], d, "wmm")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm"], "WirelessControllerQosProfile-Wmm"); ok {
			if err = d.Set("wmm", vv); err != nil {
				return fmt.Errorf("Error reading wmm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm: %v", err)
		}
	}

	if err = d.Set("wmm_be_dscp", flattenWirelessControllerQosProfileWmmBeDscp(o["wmm-be-dscp"], d, "wmm_be_dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm-be-dscp"], "WirelessControllerQosProfile-WmmBeDscp"); ok {
			if err = d.Set("wmm_be_dscp", vv); err != nil {
				return fmt.Errorf("Error reading wmm_be_dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm_be_dscp: %v", err)
		}
	}

	if err = d.Set("wmm_bk_dscp", flattenWirelessControllerQosProfileWmmBkDscp(o["wmm-bk-dscp"], d, "wmm_bk_dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm-bk-dscp"], "WirelessControllerQosProfile-WmmBkDscp"); ok {
			if err = d.Set("wmm_bk_dscp", vv); err != nil {
				return fmt.Errorf("Error reading wmm_bk_dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm_bk_dscp: %v", err)
		}
	}

	if err = d.Set("wmm_dscp_marking", flattenWirelessControllerQosProfileWmmDscpMarking(o["wmm-dscp-marking"], d, "wmm_dscp_marking")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm-dscp-marking"], "WirelessControllerQosProfile-WmmDscpMarking"); ok {
			if err = d.Set("wmm_dscp_marking", vv); err != nil {
				return fmt.Errorf("Error reading wmm_dscp_marking: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm_dscp_marking: %v", err)
		}
	}

	if err = d.Set("wmm_uapsd", flattenWirelessControllerQosProfileWmmUapsd(o["wmm-uapsd"], d, "wmm_uapsd")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm-uapsd"], "WirelessControllerQosProfile-WmmUapsd"); ok {
			if err = d.Set("wmm_uapsd", vv); err != nil {
				return fmt.Errorf("Error reading wmm_uapsd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm_uapsd: %v", err)
		}
	}

	if err = d.Set("wmm_vi_dscp", flattenWirelessControllerQosProfileWmmViDscp(o["wmm-vi-dscp"], d, "wmm_vi_dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm-vi-dscp"], "WirelessControllerQosProfile-WmmViDscp"); ok {
			if err = d.Set("wmm_vi_dscp", vv); err != nil {
				return fmt.Errorf("Error reading wmm_vi_dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm_vi_dscp: %v", err)
		}
	}

	if err = d.Set("wmm_vo_dscp", flattenWirelessControllerQosProfileWmmVoDscp(o["wmm-vo-dscp"], d, "wmm_vo_dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm-vo-dscp"], "WirelessControllerQosProfile-WmmVoDscp"); ok {
			if err = d.Set("wmm_vo_dscp", vv); err != nil {
				return fmt.Errorf("Error reading wmm_vo_dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm_vo_dscp: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerQosProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerQosProfileBandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileBandwidthCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileBurst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileCallAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileCallCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDownlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDownlinkSta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDscpWmmBe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerQosProfileDscpWmmBk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerQosProfileDscpWmmMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDscpWmmVi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerQosProfileDscpWmmVo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerQosProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileUplink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileUplinkSta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmBeDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmBkDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmDscpMarking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmUapsd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmViDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmVoDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerQosProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bandwidth_admission_control"); ok || d.HasChange("bandwidth_admission_control") {
		t, err := expandWirelessControllerQosProfileBandwidthAdmissionControl(d, v, "bandwidth_admission_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-admission-control"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_capacity"); ok || d.HasChange("bandwidth_capacity") {
		t, err := expandWirelessControllerQosProfileBandwidthCapacity(d, v, "bandwidth_capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-capacity"] = t
		}
	}

	if v, ok := d.GetOk("burst"); ok || d.HasChange("burst") {
		t, err := expandWirelessControllerQosProfileBurst(d, v, "burst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["burst"] = t
		}
	}

	if v, ok := d.GetOk("call_admission_control"); ok || d.HasChange("call_admission_control") {
		t, err := expandWirelessControllerQosProfileCallAdmissionControl(d, v, "call_admission_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-admission-control"] = t
		}
	}

	if v, ok := d.GetOk("call_capacity"); ok || d.HasChange("call_capacity") {
		t, err := expandWirelessControllerQosProfileCallCapacity(d, v, "call_capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-capacity"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerQosProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("downlink"); ok || d.HasChange("downlink") {
		t, err := expandWirelessControllerQosProfileDownlink(d, v, "downlink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlink"] = t
		}
	}

	if v, ok := d.GetOk("downlink_sta"); ok || d.HasChange("downlink_sta") {
		t, err := expandWirelessControllerQosProfileDownlinkSta(d, v, "downlink_sta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlink-sta"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_be"); ok || d.HasChange("dscp_wmm_be") {
		t, err := expandWirelessControllerQosProfileDscpWmmBe(d, v, "dscp_wmm_be")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-be"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_bk"); ok || d.HasChange("dscp_wmm_bk") {
		t, err := expandWirelessControllerQosProfileDscpWmmBk(d, v, "dscp_wmm_bk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-bk"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_mapping"); ok || d.HasChange("dscp_wmm_mapping") {
		t, err := expandWirelessControllerQosProfileDscpWmmMapping(d, v, "dscp_wmm_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-mapping"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_vi"); ok || d.HasChange("dscp_wmm_vi") {
		t, err := expandWirelessControllerQosProfileDscpWmmVi(d, v, "dscp_wmm_vi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-vi"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_vo"); ok || d.HasChange("dscp_wmm_vo") {
		t, err := expandWirelessControllerQosProfileDscpWmmVo(d, v, "dscp_wmm_vo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-vo"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerQosProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uplink"); ok || d.HasChange("uplink") {
		t, err := expandWirelessControllerQosProfileUplink(d, v, "uplink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uplink"] = t
		}
	}

	if v, ok := d.GetOk("uplink_sta"); ok || d.HasChange("uplink_sta") {
		t, err := expandWirelessControllerQosProfileUplinkSta(d, v, "uplink_sta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uplink-sta"] = t
		}
	}

	if v, ok := d.GetOk("wmm"); ok || d.HasChange("wmm") {
		t, err := expandWirelessControllerQosProfileWmm(d, v, "wmm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm"] = t
		}
	}

	if v, ok := d.GetOk("wmm_be_dscp"); ok || d.HasChange("wmm_be_dscp") {
		t, err := expandWirelessControllerQosProfileWmmBeDscp(d, v, "wmm_be_dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-be-dscp"] = t
		}
	}

	if v, ok := d.GetOk("wmm_bk_dscp"); ok || d.HasChange("wmm_bk_dscp") {
		t, err := expandWirelessControllerQosProfileWmmBkDscp(d, v, "wmm_bk_dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-bk-dscp"] = t
		}
	}

	if v, ok := d.GetOk("wmm_dscp_marking"); ok || d.HasChange("wmm_dscp_marking") {
		t, err := expandWirelessControllerQosProfileWmmDscpMarking(d, v, "wmm_dscp_marking")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-dscp-marking"] = t
		}
	}

	if v, ok := d.GetOk("wmm_uapsd"); ok || d.HasChange("wmm_uapsd") {
		t, err := expandWirelessControllerQosProfileWmmUapsd(d, v, "wmm_uapsd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-uapsd"] = t
		}
	}

	if v, ok := d.GetOk("wmm_vi_dscp"); ok || d.HasChange("wmm_vi_dscp") {
		t, err := expandWirelessControllerQosProfileWmmViDscp(d, v, "wmm_vi_dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-vi-dscp"] = t
		}
	}

	if v, ok := d.GetOk("wmm_vo_dscp"); ok || d.HasChange("wmm_vo_dscp") {
		t, err := expandWirelessControllerQosProfileWmmVoDscp(d, v, "wmm_vo_dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-vo-dscp"] = t
		}
	}

	return &obj, nil
}
