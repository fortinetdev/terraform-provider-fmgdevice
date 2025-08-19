// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Bluetooth Low Energy profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerBleProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerBleProfileCreate,
		Read:   resourceWirelessControllerBleProfileRead,
		Update: resourceWirelessControllerBleProfileUpdate,
		Delete: resourceWirelessControllerBleProfileDelete,

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
			"advertising": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"beacon_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ble_scanning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eddystone_instance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eddystone_namespace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eddystone_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eddystone_url_encode_hex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ibeacon_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"major_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minor_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"scan_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scan_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scan_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scan_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_window": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"txpower": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerBleProfileCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerBleProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerBleProfile resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerBleProfile(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerBleProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerBleProfileRead(d, m)
}

func resourceWirelessControllerBleProfileUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerBleProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerBleProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerBleProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerBleProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerBleProfileRead(d, m)
}

func resourceWirelessControllerBleProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerBleProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerBleProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerBleProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerBleProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerBleProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerBleProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerBleProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerBleProfileAdvertising(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerBleProfileBeaconInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileBleScanning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileEddystoneInstance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileEddystoneNamespace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileEddystoneUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileEddystoneUrlEncodeHex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileIbeaconUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileMajorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileMinorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileScanInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileScanPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileScanThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileScanTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileScanType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileScanWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileTxpower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerBleProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("advertising", flattenWirelessControllerBleProfileAdvertising(o["advertising"], d, "advertising")); err != nil {
		if vv, ok := fortiAPIPatch(o["advertising"], "WirelessControllerBleProfile-Advertising"); ok {
			if err = d.Set("advertising", vv); err != nil {
				return fmt.Errorf("Error reading advertising: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading advertising: %v", err)
		}
	}

	if err = d.Set("beacon_interval", flattenWirelessControllerBleProfileBeaconInterval(o["beacon-interval"], d, "beacon_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["beacon-interval"], "WirelessControllerBleProfile-BeaconInterval"); ok {
			if err = d.Set("beacon_interval", vv); err != nil {
				return fmt.Errorf("Error reading beacon_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading beacon_interval: %v", err)
		}
	}

	if err = d.Set("ble_scanning", flattenWirelessControllerBleProfileBleScanning(o["ble-scanning"], d, "ble_scanning")); err != nil {
		if vv, ok := fortiAPIPatch(o["ble-scanning"], "WirelessControllerBleProfile-BleScanning"); ok {
			if err = d.Set("ble_scanning", vv); err != nil {
				return fmt.Errorf("Error reading ble_scanning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ble_scanning: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerBleProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerBleProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("eddystone_instance", flattenWirelessControllerBleProfileEddystoneInstance(o["eddystone-instance"], d, "eddystone_instance")); err != nil {
		if vv, ok := fortiAPIPatch(o["eddystone-instance"], "WirelessControllerBleProfile-EddystoneInstance"); ok {
			if err = d.Set("eddystone_instance", vv); err != nil {
				return fmt.Errorf("Error reading eddystone_instance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eddystone_instance: %v", err)
		}
	}

	if err = d.Set("eddystone_namespace", flattenWirelessControllerBleProfileEddystoneNamespace(o["eddystone-namespace"], d, "eddystone_namespace")); err != nil {
		if vv, ok := fortiAPIPatch(o["eddystone-namespace"], "WirelessControllerBleProfile-EddystoneNamespace"); ok {
			if err = d.Set("eddystone_namespace", vv); err != nil {
				return fmt.Errorf("Error reading eddystone_namespace: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eddystone_namespace: %v", err)
		}
	}

	if err = d.Set("eddystone_url", flattenWirelessControllerBleProfileEddystoneUrl(o["eddystone-url"], d, "eddystone_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["eddystone-url"], "WirelessControllerBleProfile-EddystoneUrl"); ok {
			if err = d.Set("eddystone_url", vv); err != nil {
				return fmt.Errorf("Error reading eddystone_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eddystone_url: %v", err)
		}
	}

	if err = d.Set("eddystone_url_encode_hex", flattenWirelessControllerBleProfileEddystoneUrlEncodeHex(o["eddystone-url-encode-hex"], d, "eddystone_url_encode_hex")); err != nil {
		if vv, ok := fortiAPIPatch(o["eddystone-url-encode-hex"], "WirelessControllerBleProfile-EddystoneUrlEncodeHex"); ok {
			if err = d.Set("eddystone_url_encode_hex", vv); err != nil {
				return fmt.Errorf("Error reading eddystone_url_encode_hex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eddystone_url_encode_hex: %v", err)
		}
	}

	if err = d.Set("ibeacon_uuid", flattenWirelessControllerBleProfileIbeaconUuid(o["ibeacon-uuid"], d, "ibeacon_uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ibeacon-uuid"], "WirelessControllerBleProfile-IbeaconUuid"); ok {
			if err = d.Set("ibeacon_uuid", vv); err != nil {
				return fmt.Errorf("Error reading ibeacon_uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ibeacon_uuid: %v", err)
		}
	}

	if err = d.Set("major_id", flattenWirelessControllerBleProfileMajorId(o["major-id"], d, "major_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["major-id"], "WirelessControllerBleProfile-MajorId"); ok {
			if err = d.Set("major_id", vv); err != nil {
				return fmt.Errorf("Error reading major_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading major_id: %v", err)
		}
	}

	if err = d.Set("minor_id", flattenWirelessControllerBleProfileMinorId(o["minor-id"], d, "minor_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["minor-id"], "WirelessControllerBleProfile-MinorId"); ok {
			if err = d.Set("minor_id", vv); err != nil {
				return fmt.Errorf("Error reading minor_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minor_id: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerBleProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerBleProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("scan_interval", flattenWirelessControllerBleProfileScanInterval(o["scan-interval"], d, "scan_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-interval"], "WirelessControllerBleProfile-ScanInterval"); ok {
			if err = d.Set("scan_interval", vv); err != nil {
				return fmt.Errorf("Error reading scan_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_interval: %v", err)
		}
	}

	if err = d.Set("scan_period", flattenWirelessControllerBleProfileScanPeriod(o["scan-period"], d, "scan_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-period"], "WirelessControllerBleProfile-ScanPeriod"); ok {
			if err = d.Set("scan_period", vv); err != nil {
				return fmt.Errorf("Error reading scan_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_period: %v", err)
		}
	}

	if err = d.Set("scan_threshold", flattenWirelessControllerBleProfileScanThreshold(o["scan-threshold"], d, "scan_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-threshold"], "WirelessControllerBleProfile-ScanThreshold"); ok {
			if err = d.Set("scan_threshold", vv); err != nil {
				return fmt.Errorf("Error reading scan_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_threshold: %v", err)
		}
	}

	if err = d.Set("scan_time", flattenWirelessControllerBleProfileScanTime(o["scan-time"], d, "scan_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-time"], "WirelessControllerBleProfile-ScanTime"); ok {
			if err = d.Set("scan_time", vv); err != nil {
				return fmt.Errorf("Error reading scan_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_time: %v", err)
		}
	}

	if err = d.Set("scan_type", flattenWirelessControllerBleProfileScanType(o["scan-type"], d, "scan_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-type"], "WirelessControllerBleProfile-ScanType"); ok {
			if err = d.Set("scan_type", vv); err != nil {
				return fmt.Errorf("Error reading scan_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_type: %v", err)
		}
	}

	if err = d.Set("scan_window", flattenWirelessControllerBleProfileScanWindow(o["scan-window"], d, "scan_window")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-window"], "WirelessControllerBleProfile-ScanWindow"); ok {
			if err = d.Set("scan_window", vv); err != nil {
				return fmt.Errorf("Error reading scan_window: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_window: %v", err)
		}
	}

	if err = d.Set("txpower", flattenWirelessControllerBleProfileTxpower(o["txpower"], d, "txpower")); err != nil {
		if vv, ok := fortiAPIPatch(o["txpower"], "WirelessControllerBleProfile-Txpower"); ok {
			if err = d.Set("txpower", vv); err != nil {
				return fmt.Errorf("Error reading txpower: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading txpower: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerBleProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerBleProfileAdvertising(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerBleProfileBeaconInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileBleScanning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileEddystoneInstance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileEddystoneNamespace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileEddystoneUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileEddystoneUrlEncodeHex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileIbeaconUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileMajorId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileMinorId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileScanInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileScanPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileScanThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileScanTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileScanType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileScanWindow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileTxpower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerBleProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("advertising"); ok || d.HasChange("advertising") {
		t, err := expandWirelessControllerBleProfileAdvertising(d, v, "advertising")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertising"] = t
		}
	}

	if v, ok := d.GetOk("beacon_interval"); ok || d.HasChange("beacon_interval") {
		t, err := expandWirelessControllerBleProfileBeaconInterval(d, v, "beacon_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-interval"] = t
		}
	}

	if v, ok := d.GetOk("ble_scanning"); ok || d.HasChange("ble_scanning") {
		t, err := expandWirelessControllerBleProfileBleScanning(d, v, "ble_scanning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-scanning"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerBleProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_instance"); ok || d.HasChange("eddystone_instance") {
		t, err := expandWirelessControllerBleProfileEddystoneInstance(d, v, "eddystone_instance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-instance"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_namespace"); ok || d.HasChange("eddystone_namespace") {
		t, err := expandWirelessControllerBleProfileEddystoneNamespace(d, v, "eddystone_namespace")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-namespace"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_url"); ok || d.HasChange("eddystone_url") {
		t, err := expandWirelessControllerBleProfileEddystoneUrl(d, v, "eddystone_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-url"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_url_encode_hex"); ok || d.HasChange("eddystone_url_encode_hex") {
		t, err := expandWirelessControllerBleProfileEddystoneUrlEncodeHex(d, v, "eddystone_url_encode_hex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-url-encode-hex"] = t
		}
	}

	if v, ok := d.GetOk("ibeacon_uuid"); ok || d.HasChange("ibeacon_uuid") {
		t, err := expandWirelessControllerBleProfileIbeaconUuid(d, v, "ibeacon_uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibeacon-uuid"] = t
		}
	}

	if v, ok := d.GetOk("major_id"); ok || d.HasChange("major_id") {
		t, err := expandWirelessControllerBleProfileMajorId(d, v, "major_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["major-id"] = t
		}
	}

	if v, ok := d.GetOk("minor_id"); ok || d.HasChange("minor_id") {
		t, err := expandWirelessControllerBleProfileMinorId(d, v, "minor_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minor-id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerBleProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("scan_interval"); ok || d.HasChange("scan_interval") {
		t, err := expandWirelessControllerBleProfileScanInterval(d, v, "scan_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-interval"] = t
		}
	}

	if v, ok := d.GetOk("scan_period"); ok || d.HasChange("scan_period") {
		t, err := expandWirelessControllerBleProfileScanPeriod(d, v, "scan_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-period"] = t
		}
	}

	if v, ok := d.GetOk("scan_threshold"); ok || d.HasChange("scan_threshold") {
		t, err := expandWirelessControllerBleProfileScanThreshold(d, v, "scan_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-threshold"] = t
		}
	}

	if v, ok := d.GetOk("scan_time"); ok || d.HasChange("scan_time") {
		t, err := expandWirelessControllerBleProfileScanTime(d, v, "scan_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-time"] = t
		}
	}

	if v, ok := d.GetOk("scan_type"); ok || d.HasChange("scan_type") {
		t, err := expandWirelessControllerBleProfileScanType(d, v, "scan_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-type"] = t
		}
	}

	if v, ok := d.GetOk("scan_window"); ok || d.HasChange("scan_window") {
		t, err := expandWirelessControllerBleProfileScanWindow(d, v, "scan_window")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-window"] = t
		}
	}

	if v, ok := d.GetOk("txpower"); ok || d.HasChange("txpower") {
		t, err := expandWirelessControllerBleProfileTxpower(d, v, "txpower")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["txpower"] = t
		}
	}

	return &obj, nil
}
