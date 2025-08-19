// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Fabric connector configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCsfFabricConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCsfFabricConnectorCreate,
		Read:   resourceSystemCsfFabricConnectorRead,
		Update: resourceSystemCsfFabricConnectorUpdate,
		Delete: resourceSystemCsfFabricConnectorDelete,

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
			"accprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"configuration_write_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemCsfFabricConnectorCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemCsfFabricConnector(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCsfFabricConnector resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCsfFabricConnector(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemCsfFabricConnector resource: %v", err)
	}

	d.SetId(getStringKey(d, "serial"))

	return resourceSystemCsfFabricConnectorRead(d, m)
}

func resourceSystemCsfFabricConnectorUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemCsfFabricConnector(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsfFabricConnector resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCsfFabricConnector(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsfFabricConnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "serial"))

	return resourceSystemCsfFabricConnectorRead(d, m)
}

func resourceSystemCsfFabricConnectorDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemCsfFabricConnector(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCsfFabricConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCsfFabricConnectorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCsfFabricConnector(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsfFabricConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCsfFabricConnector(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsfFabricConnector resource from API: %v", err)
	}
	return nil
}

func flattenSystemCsfFabricConnectorAccprofile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCsfFabricConnectorConfigurationWriteAccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricConnectorSerial2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricConnectorVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemCsfFabricConnector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("accprofile", flattenSystemCsfFabricConnectorAccprofile2edl(o["accprofile"], d, "accprofile")); err != nil {
		if vv, ok := fortiAPIPatch(o["accprofile"], "SystemCsfFabricConnector-Accprofile"); ok {
			if err = d.Set("accprofile", vv); err != nil {
				return fmt.Errorf("Error reading accprofile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("configuration_write_access", flattenSystemCsfFabricConnectorConfigurationWriteAccess2edl(o["configuration-write-access"], d, "configuration_write_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["configuration-write-access"], "SystemCsfFabricConnector-ConfigurationWriteAccess"); ok {
			if err = d.Set("configuration_write_access", vv); err != nil {
				return fmt.Errorf("Error reading configuration_write_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading configuration_write_access: %v", err)
		}
	}

	if err = d.Set("serial", flattenSystemCsfFabricConnectorSerial2edl(o["serial"], d, "serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial"], "SystemCsfFabricConnector-Serial"); ok {
			if err = d.Set("serial", vv); err != nil {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemCsfFabricConnectorVdom2edl(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemCsfFabricConnector-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenSystemCsfFabricConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCsfFabricConnectorAccprofile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfFabricConnectorConfigurationWriteAccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricConnectorSerial2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricConnectorVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemCsfFabricConnector(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accprofile"); ok || d.HasChange("accprofile") {
		t, err := expandSystemCsfFabricConnectorAccprofile2edl(d, v, "accprofile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	}

	if v, ok := d.GetOk("configuration_write_access"); ok || d.HasChange("configuration_write_access") {
		t, err := expandSystemCsfFabricConnectorConfigurationWriteAccess2edl(d, v, "configuration_write_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["configuration-write-access"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		t, err := expandSystemCsfFabricConnectorSerial2edl(d, v, "serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemCsfFabricConnectorVdom2edl(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
