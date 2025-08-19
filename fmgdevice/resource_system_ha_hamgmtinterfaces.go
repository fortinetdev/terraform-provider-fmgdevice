// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Reserve interfaces to manage individual cluster units.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemHaHaMgmtInterfaces() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaHaMgmtInterfacesCreate,
		Read:   resourceSystemHaHaMgmtInterfacesRead,
		Update: resourceSystemHaHaMgmtInterfacesUpdate,
		Delete: resourceSystemHaHaMgmtInterfacesDelete,

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
			"dst": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemHaHaMgmtInterfacesCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemHaHaMgmtInterfaces(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemHaHaMgmtInterfaces resource while getting object: %v", err)
	}

	_, err = c.CreateSystemHaHaMgmtInterfaces(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemHaHaMgmtInterfaces resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemHaHaMgmtInterfacesRead(d, m)
}

func resourceSystemHaHaMgmtInterfacesUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemHaHaMgmtInterfaces(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaHaMgmtInterfaces resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaHaMgmtInterfaces(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaHaMgmtInterfaces resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemHaHaMgmtInterfacesRead(d, m)
}

func resourceSystemHaHaMgmtInterfacesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemHaHaMgmtInterfaces(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHaHaMgmtInterfaces resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaHaMgmtInterfacesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemHaHaMgmtInterfaces(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaHaMgmtInterfaces resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaHaMgmtInterfaces(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaHaMgmtInterfaces resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaHaMgmtInterfacesDst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaHaMgmtInterfacesGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfacesGateway62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfacesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfacesInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemHaHaMgmtInterfaces(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dst", flattenSystemHaHaMgmtInterfacesDst2edl(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "SystemHaHaMgmtInterfaces-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("gateway", flattenSystemHaHaMgmtInterfacesGateway2edl(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "SystemHaHaMgmtInterfaces-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("gateway6", flattenSystemHaHaMgmtInterfacesGateway62edl(o["gateway6"], d, "gateway6")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway6"], "SystemHaHaMgmtInterfaces-Gateway6"); ok {
			if err = d.Set("gateway6", vv); err != nil {
				return fmt.Errorf("Error reading gateway6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway6: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemHaHaMgmtInterfacesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemHaHaMgmtInterfaces-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemHaHaMgmtInterfacesInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemHaHaMgmtInterfaces-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func flattenSystemHaHaMgmtInterfacesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaHaMgmtInterfacesDst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemHaHaMgmtInterfacesGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesGateway62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemHaHaMgmtInterfaces(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandSystemHaHaMgmtInterfacesDst2edl(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandSystemHaHaMgmtInterfacesGateway2edl(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("gateway6"); ok || d.HasChange("gateway6") {
		t, err := expandSystemHaHaMgmtInterfacesGateway62edl(d, v, "gateway6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway6"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemHaHaMgmtInterfacesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemHaHaMgmtInterfacesInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	return &obj, nil
}
