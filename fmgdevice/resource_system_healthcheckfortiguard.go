// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SD-WAN status checking or health checking. Identify a server predefine by FortiGuard and determine how SD-WAN verifies that FGT can communicate with it.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemHealthCheckFortiguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHealthCheckFortiguardCreate,
		Read:   resourceSystemHealthCheckFortiguardRead,
		Update: resourceSystemHealthCheckFortiguardUpdate,
		Delete: resourceSystemHealthCheckFortiguardDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"obsolete": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemHealthCheckFortiguardCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemHealthCheckFortiguard(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemHealthCheckFortiguard resource while getting object: %v", err)
	}

	_, err = c.CreateSystemHealthCheckFortiguard(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemHealthCheckFortiguard resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemHealthCheckFortiguardRead(d, m)
}

func resourceSystemHealthCheckFortiguardUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemHealthCheckFortiguard(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHealthCheckFortiguard resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHealthCheckFortiguard(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemHealthCheckFortiguard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemHealthCheckFortiguardRead(d, m)
}

func resourceSystemHealthCheckFortiguardDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemHealthCheckFortiguard(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHealthCheckFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHealthCheckFortiguardRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemHealthCheckFortiguard(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemHealthCheckFortiguard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHealthCheckFortiguard(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHealthCheckFortiguard resource from API: %v", err)
	}
	return nil
}

func flattenSystemHealthCheckFortiguardName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHealthCheckFortiguardObsolete(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHealthCheckFortiguardProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHealthCheckFortiguardServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHealthCheckFortiguard(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemHealthCheckFortiguardName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemHealthCheckFortiguard-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("obsolete", flattenSystemHealthCheckFortiguardObsolete(o["obsolete"], d, "obsolete")); err != nil {
		if vv, ok := fortiAPIPatch(o["obsolete"], "SystemHealthCheckFortiguard-Obsolete"); ok {
			if err = d.Set("obsolete", vv); err != nil {
				return fmt.Errorf("Error reading obsolete: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obsolete: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemHealthCheckFortiguardProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemHealthCheckFortiguard-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemHealthCheckFortiguardServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemHealthCheckFortiguard-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	return nil
}

func flattenSystemHealthCheckFortiguardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHealthCheckFortiguardName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckFortiguardObsolete(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckFortiguardProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckFortiguardServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHealthCheckFortiguard(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemHealthCheckFortiguardName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("obsolete"); ok || d.HasChange("obsolete") {
		t, err := expandSystemHealthCheckFortiguardObsolete(d, v, "obsolete")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obsolete"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemHealthCheckFortiguardProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemHealthCheckFortiguardServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	return &obj, nil
}
