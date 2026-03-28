// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device IcapLocalServerIcapService

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIcapLocalServerIcapService() *schema.Resource {
	return &schema.Resource{
		Create: resourceIcapLocalServerIcapServiceCreate,
		Read:   resourceIcapLocalServerIcapServiceRead,
		Update: resourceIcapLocalServerIcapServiceUpdate,
		Delete: resourceIcapLocalServerIcapServiceDelete,

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
			"local_server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dlp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dlp_sensor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"extension_headers": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"image_analyzer_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIcapLocalServerIcapServiceCreate(d *schema.ResourceData, m interface{}) error {
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
	local_server := d.Get("local_server").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["local_server"] = local_server

	obj, err := getObjectIcapLocalServerIcapService(d)
	if err != nil {
		return fmt.Errorf("Error creating IcapLocalServerIcapService resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("service_id")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIcapLocalServerIcapService(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIcapLocalServerIcapService(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IcapLocalServerIcapService resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateIcapLocalServerIcapService(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IcapLocalServerIcapService resource: %v", err)
		}

		if v != nil && v["service-id"] != nil {
			if vidn, ok := v["service-id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceIcapLocalServerIcapServiceRead(d, m)
			} else {
				return fmt.Errorf("Error creating IcapLocalServerIcapService resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "service_id")))

	return resourceIcapLocalServerIcapServiceRead(d, m)
}

func resourceIcapLocalServerIcapServiceUpdate(d *schema.ResourceData, m interface{}) error {
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
	local_server := d.Get("local_server").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["local_server"] = local_server

	obj, err := getObjectIcapLocalServerIcapService(d)
	if err != nil {
		return fmt.Errorf("Error updating IcapLocalServerIcapService resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIcapLocalServerIcapService(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IcapLocalServerIcapService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "service_id")))

	return resourceIcapLocalServerIcapServiceRead(d, m)
}

func resourceIcapLocalServerIcapServiceDelete(d *schema.ResourceData, m interface{}) error {
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
	local_server := d.Get("local_server").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["local_server"] = local_server

	wsParams["adom"] = adomv

	err = c.DeleteIcapLocalServerIcapService(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IcapLocalServerIcapService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIcapLocalServerIcapServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	local_server := d.Get("local_server").(string)
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
	if local_server == "" {
		local_server = importOptionChecking(m.(*FortiClient).Cfg, "local_server")
		if local_server == "" {
			return fmt.Errorf("Parameter local_server is missing")
		}
		if err = d.Set("local_server", local_server); err != nil {
			return fmt.Errorf("Error set params local_server: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["local_server"] = local_server

	o, err := c.ReadIcapLocalServerIcapService(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IcapLocalServerIcapService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIcapLocalServerIcapService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IcapLocalServerIcapService resource from API: %v", err)
	}
	return nil
}

func flattenIcapLocalServerIcapServiceAvProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIcapServiceDlpProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIcapServiceDlpSensor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIcapServiceExtensionHeaders2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIcapServiceImageAnalyzerProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIcapServiceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerIcapServiceProfileProtocolOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIcapServiceServiceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerIcapServiceWebfilterProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectIcapLocalServerIcapService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("av_profile", flattenIcapLocalServerIcapServiceAvProfile2edl(o["av-profile"], d, "av_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile"], "IcapLocalServerIcapService-AvProfile"); ok {
			if err = d.Set("av_profile", vv); err != nil {
				return fmt.Errorf("Error reading av_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenIcapLocalServerIcapServiceDlpProfile2edl(o["dlp-profile"], d, "dlp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-profile"], "IcapLocalServerIcapService-DlpProfile"); ok {
			if err = d.Set("dlp_profile", vv); err != nil {
				return fmt.Errorf("Error reading dlp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenIcapLocalServerIcapServiceDlpSensor2edl(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor"], "IcapLocalServerIcapService-DlpSensor"); ok {
			if err = d.Set("dlp_sensor", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("extension_headers", flattenIcapLocalServerIcapServiceExtensionHeaders2edl(o["extension-headers"], d, "extension_headers")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension-headers"], "IcapLocalServerIcapService-ExtensionHeaders"); ok {
			if err = d.Set("extension_headers", vv); err != nil {
				return fmt.Errorf("Error reading extension_headers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension_headers: %v", err)
		}
	}

	if err = d.Set("image_analyzer_profile", flattenIcapLocalServerIcapServiceImageAnalyzerProfile2edl(o["image-analyzer-profile"], d, "image_analyzer_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-analyzer-profile"], "IcapLocalServerIcapService-ImageAnalyzerProfile"); ok {
			if err = d.Set("image_analyzer_profile", vv); err != nil {
				return fmt.Errorf("Error reading image_analyzer_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_analyzer_profile: %v", err)
		}
	}

	if err = d.Set("name", flattenIcapLocalServerIcapServiceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "IcapLocalServerIcapService-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenIcapLocalServerIcapServiceProfileProtocolOptions2edl(o["profile-protocol-options"], d, "profile_protocol_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-protocol-options"], "IcapLocalServerIcapService-ProfileProtocolOptions"); ok {
			if err = d.Set("profile_protocol_options", vv); err != nil {
				return fmt.Errorf("Error reading profile_protocol_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("service_id", flattenIcapLocalServerIcapServiceServiceId2edl(o["service-id"], d, "service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-id"], "IcapLocalServerIcapService-ServiceId"); ok {
			if err = d.Set("service_id", vv); err != nil {
				return fmt.Errorf("Error reading service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_id: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenIcapLocalServerIcapServiceWebfilterProfile2edl(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "IcapLocalServerIcapService-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	return nil
}

func flattenIcapLocalServerIcapServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIcapLocalServerIcapServiceAvProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIcapServiceDlpProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIcapServiceDlpSensor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIcapServiceExtensionHeaders2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIcapServiceImageAnalyzerProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIcapServiceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerIcapServiceProfileProtocolOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIcapServiceServiceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerIcapServiceWebfilterProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectIcapLocalServerIcapService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("av_profile"); ok || d.HasChange("av_profile") {
		t, err := expandIcapLocalServerIcapServiceAvProfile2edl(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile"); ok || d.HasChange("dlp_profile") {
		t, err := expandIcapLocalServerIcapServiceDlpProfile2edl(d, v, "dlp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok || d.HasChange("dlp_sensor") {
		t, err := expandIcapLocalServerIcapServiceDlpSensor2edl(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("extension_headers"); ok || d.HasChange("extension_headers") {
		t, err := expandIcapLocalServerIcapServiceExtensionHeaders2edl(d, v, "extension_headers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-headers"] = t
		}
	}

	if v, ok := d.GetOk("image_analyzer_profile"); ok || d.HasChange("image_analyzer_profile") {
		t, err := expandIcapLocalServerIcapServiceImageAnalyzerProfile2edl(d, v, "image_analyzer_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-analyzer-profile"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandIcapLocalServerIcapServiceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok || d.HasChange("profile_protocol_options") {
		t, err := expandIcapLocalServerIcapServiceProfileProtocolOptions2edl(d, v, "profile_protocol_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("service_id"); ok || d.HasChange("service_id") {
		t, err := expandIcapLocalServerIcapServiceServiceId2edl(d, v, "service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-id"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandIcapLocalServerIcapServiceWebfilterProfile2edl(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	return &obj, nil
}
