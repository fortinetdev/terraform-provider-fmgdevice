// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Fabric device configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCsfFabricDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCsfFabricDeviceCreate,
		Read:   resourceSystemCsfFabricDeviceRead,
		Update: resourceSystemCsfFabricDeviceUpdate,
		Delete: resourceSystemCsfFabricDeviceDelete,

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
			"access_token": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"device_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemCsfFabricDeviceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemCsfFabricDevice(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCsfFabricDevice resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCsfFabricDevice(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemCsfFabricDevice resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemCsfFabricDeviceRead(d, m)
}

func resourceSystemCsfFabricDeviceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemCsfFabricDevice(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsfFabricDevice resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCsfFabricDevice(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsfFabricDevice resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemCsfFabricDeviceRead(d, m)
}

func resourceSystemCsfFabricDeviceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemCsfFabricDevice(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCsfFabricDevice resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCsfFabricDeviceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCsfFabricDevice(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsfFabricDevice resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCsfFabricDevice(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsfFabricDevice resource from API: %v", err)
	}
	return nil
}

func flattenSystemCsfFabricDeviceDeviceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceHttpsPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemCsfFabricDevice(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("device_ip", flattenSystemCsfFabricDeviceDeviceIp2edl(o["device-ip"], d, "device_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-ip"], "SystemCsfFabricDevice-DeviceIp"); ok {
			if err = d.Set("device_ip", vv); err != nil {
				return fmt.Errorf("Error reading device_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_ip: %v", err)
		}
	}

	if err = d.Set("https_port", flattenSystemCsfFabricDeviceHttpsPort2edl(o["https-port"], d, "https_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-port"], "SystemCsfFabricDevice-HttpsPort"); ok {
			if err = d.Set("https_port", vv); err != nil {
				return fmt.Errorf("Error reading https_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_port: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemCsfFabricDeviceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemCsfFabricDevice-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemCsfFabricDeviceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCsfFabricDeviceAccessToken2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfFabricDeviceDeviceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDeviceHttpsPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDeviceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCsfFabricDevice(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_token"); ok || d.HasChange("access_token") {
		t, err := expandSystemCsfFabricDeviceAccessToken2edl(d, v, "access_token")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-token"] = t
		}
	}

	if v, ok := d.GetOk("device_ip"); ok || d.HasChange("device_ip") {
		t, err := expandSystemCsfFabricDeviceDeviceIp2edl(d, v, "device_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-ip"] = t
		}
	}

	if v, ok := d.GetOk("https_port"); ok || d.HasChange("https_port") {
		t, err := expandSystemCsfFabricDeviceHttpsPort2edl(d, v, "https_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-port"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemCsfFabricDeviceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
