// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Nodes which will be included in the upgrade.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFederatedUpgradeNodeList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFederatedUpgradeNodeListCreate,
		Read:   resourceSystemFederatedUpgradeNodeListRead,
		Update: resourceSystemFederatedUpgradeNodeListUpdate,
		Delete: resourceSystemFederatedUpgradeNodeListDelete,

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
			"coordinating_fortigate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maximum_minutes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"setup_time": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"timing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upgrade_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemFederatedUpgradeNodeListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemFederatedUpgradeNodeList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemFederatedUpgradeNodeList resource while getting object: %v", err)
	}

	_, err = c.CreateSystemFederatedUpgradeNodeList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemFederatedUpgradeNodeList resource: %v", err)
	}

	d.SetId(getStringKey(d, "serial"))

	return resourceSystemFederatedUpgradeNodeListRead(d, m)
}

func resourceSystemFederatedUpgradeNodeListUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemFederatedUpgradeNodeList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFederatedUpgradeNodeList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFederatedUpgradeNodeList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemFederatedUpgradeNodeList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "serial"))

	return resourceSystemFederatedUpgradeNodeListRead(d, m)
}

func resourceSystemFederatedUpgradeNodeListDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemFederatedUpgradeNodeList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFederatedUpgradeNodeList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFederatedUpgradeNodeListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemFederatedUpgradeNodeList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemFederatedUpgradeNodeList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFederatedUpgradeNodeList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFederatedUpgradeNodeList resource from API: %v", err)
	}
	return nil
}

func flattenSystemFederatedUpgradeNodeListCoordinatingFortigate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListDeviceType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListMaximumMinutes2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListSerial2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListSetupTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFederatedUpgradeNodeListTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFederatedUpgradeNodeListTiming2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListUpgradePath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFederatedUpgradeNodeList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("coordinating_fortigate", flattenSystemFederatedUpgradeNodeListCoordinatingFortigate2edl(o["coordinating-fortigate"], d, "coordinating_fortigate")); err != nil {
		if vv, ok := fortiAPIPatch(o["coordinating-fortigate"], "SystemFederatedUpgradeNodeList-CoordinatingFortigate"); ok {
			if err = d.Set("coordinating_fortigate", vv); err != nil {
				return fmt.Errorf("Error reading coordinating_fortigate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading coordinating_fortigate: %v", err)
		}
	}

	if err = d.Set("device_type", flattenSystemFederatedUpgradeNodeListDeviceType2edl(o["device-type"], d, "device_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-type"], "SystemFederatedUpgradeNodeList-DeviceType"); ok {
			if err = d.Set("device_type", vv); err != nil {
				return fmt.Errorf("Error reading device_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_type: %v", err)
		}
	}

	if err = d.Set("maximum_minutes", flattenSystemFederatedUpgradeNodeListMaximumMinutes2edl(o["maximum-minutes"], d, "maximum_minutes")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-minutes"], "SystemFederatedUpgradeNodeList-MaximumMinutes"); ok {
			if err = d.Set("maximum_minutes", vv); err != nil {
				return fmt.Errorf("Error reading maximum_minutes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_minutes: %v", err)
		}
	}

	if err = d.Set("serial", flattenSystemFederatedUpgradeNodeListSerial2edl(o["serial"], d, "serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial"], "SystemFederatedUpgradeNodeList-Serial"); ok {
			if err = d.Set("serial", vv); err != nil {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	if err = d.Set("setup_time", flattenSystemFederatedUpgradeNodeListSetupTime2edl(o["setup-time"], d, "setup_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["setup-time"], "SystemFederatedUpgradeNodeList-SetupTime"); ok {
			if err = d.Set("setup_time", vv); err != nil {
				return fmt.Errorf("Error reading setup_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading setup_time: %v", err)
		}
	}

	if err = d.Set("time", flattenSystemFederatedUpgradeNodeListTime2edl(o["time"], d, "time")); err != nil {
		if vv, ok := fortiAPIPatch(o["time"], "SystemFederatedUpgradeNodeList-Time"); ok {
			if err = d.Set("time", vv); err != nil {
				return fmt.Errorf("Error reading time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}

	if err = d.Set("timing", flattenSystemFederatedUpgradeNodeListTiming2edl(o["timing"], d, "timing")); err != nil {
		if vv, ok := fortiAPIPatch(o["timing"], "SystemFederatedUpgradeNodeList-Timing"); ok {
			if err = d.Set("timing", vv); err != nil {
				return fmt.Errorf("Error reading timing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timing: %v", err)
		}
	}

	if err = d.Set("upgrade_path", flattenSystemFederatedUpgradeNodeListUpgradePath2edl(o["upgrade-path"], d, "upgrade_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["upgrade-path"], "SystemFederatedUpgradeNodeList-UpgradePath"); ok {
			if err = d.Set("upgrade_path", vv); err != nil {
				return fmt.Errorf("Error reading upgrade_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upgrade_path: %v", err)
		}
	}

	return nil
}

func flattenSystemFederatedUpgradeNodeListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFederatedUpgradeNodeListCoordinatingFortigate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListDeviceType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListMaximumMinutes2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListSerial2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListSetupTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFederatedUpgradeNodeListTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFederatedUpgradeNodeListTiming2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListUpgradePath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFederatedUpgradeNodeList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("coordinating_fortigate"); ok || d.HasChange("coordinating_fortigate") {
		t, err := expandSystemFederatedUpgradeNodeListCoordinatingFortigate2edl(d, v, "coordinating_fortigate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["coordinating-fortigate"] = t
		}
	}

	if v, ok := d.GetOk("device_type"); ok || d.HasChange("device_type") {
		t, err := expandSystemFederatedUpgradeNodeListDeviceType2edl(d, v, "device_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-type"] = t
		}
	}

	if v, ok := d.GetOk("maximum_minutes"); ok || d.HasChange("maximum_minutes") {
		t, err := expandSystemFederatedUpgradeNodeListMaximumMinutes2edl(d, v, "maximum_minutes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-minutes"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		t, err := expandSystemFederatedUpgradeNodeListSerial2edl(d, v, "serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	if v, ok := d.GetOk("setup_time"); ok || d.HasChange("setup_time") {
		t, err := expandSystemFederatedUpgradeNodeListSetupTime2edl(d, v, "setup_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["setup-time"] = t
		}
	}

	if v, ok := d.GetOk("time"); ok || d.HasChange("time") {
		t, err := expandSystemFederatedUpgradeNodeListTime2edl(d, v, "time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time"] = t
		}
	}

	if v, ok := d.GetOk("timing"); ok || d.HasChange("timing") {
		t, err := expandSystemFederatedUpgradeNodeListTiming2edl(d, v, "timing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timing"] = t
		}
	}

	if v, ok := d.GetOk("upgrade_path"); ok || d.HasChange("upgrade_path") {
		t, err := expandSystemFederatedUpgradeNodeListUpgradePath2edl(d, v, "upgrade_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upgrade-path"] = t
		}
	}

	return &obj, nil
}
