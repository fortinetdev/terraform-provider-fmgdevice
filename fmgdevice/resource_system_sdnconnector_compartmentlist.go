// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure OCI compartment list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdnConnectorCompartmentList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdnConnectorCompartmentListCreate,
		Read:   resourceSystemSdnConnectorCompartmentListRead,
		Update: resourceSystemSdnConnectorCompartmentListUpdate,
		Delete: resourceSystemSdnConnectorCompartmentListDelete,

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
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemSdnConnectorCompartmentListCreate(d *schema.ResourceData, m interface{}) error {
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
	sdn_connector := d.Get("sdn_connector").(string)
	paradict["device"] = device_name
	paradict["sdn_connector"] = sdn_connector

	obj, err := getObjectSystemSdnConnectorCompartmentList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdnConnectorCompartmentList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("compartment_id")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemSdnConnectorCompartmentList(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemSdnConnectorCompartmentList(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemSdnConnectorCompartmentList resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemSdnConnectorCompartmentList(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemSdnConnectorCompartmentList resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "compartment_id"))

	return resourceSystemSdnConnectorCompartmentListRead(d, m)
}

func resourceSystemSdnConnectorCompartmentListUpdate(d *schema.ResourceData, m interface{}) error {
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
	sdn_connector := d.Get("sdn_connector").(string)
	paradict["device"] = device_name
	paradict["sdn_connector"] = sdn_connector

	obj, err := getObjectSystemSdnConnectorCompartmentList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnectorCompartmentList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSdnConnectorCompartmentList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnectorCompartmentList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "compartment_id"))

	return resourceSystemSdnConnectorCompartmentListRead(d, m)
}

func resourceSystemSdnConnectorCompartmentListDelete(d *schema.ResourceData, m interface{}) error {
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
	sdn_connector := d.Get("sdn_connector").(string)
	paradict["device"] = device_name
	paradict["sdn_connector"] = sdn_connector

	wsParams["adom"] = adomv

	err = c.DeleteSystemSdnConnectorCompartmentList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdnConnectorCompartmentList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdnConnectorCompartmentListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	sdn_connector := d.Get("sdn_connector").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if sdn_connector == "" {
		sdn_connector = importOptionChecking(m.(*FortiClient).Cfg, "sdn_connector")
		if sdn_connector == "" {
			return fmt.Errorf("Parameter sdn_connector is missing")
		}
		if err = d.Set("sdn_connector", sdn_connector); err != nil {
			return fmt.Errorf("Error set params sdn_connector: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["sdn_connector"] = sdn_connector

	o, err := c.ReadSystemSdnConnectorCompartmentList(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemSdnConnectorCompartmentList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdnConnectorCompartmentList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnConnectorCompartmentList resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdnConnectorCompartmentListCompartmentId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdnConnectorCompartmentList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("compartment_id", flattenSystemSdnConnectorCompartmentListCompartmentId2edl(o["compartment-id"], d, "compartment_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["compartment-id"], "SystemSdnConnectorCompartmentList-CompartmentId"); ok {
			if err = d.Set("compartment_id", vv); err != nil {
				return fmt.Errorf("Error reading compartment_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading compartment_id: %v", err)
		}
	}

	return nil
}

func flattenSystemSdnConnectorCompartmentListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdnConnectorCompartmentListCompartmentId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdnConnectorCompartmentList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("compartment_id"); ok || d.HasChange("compartment_id") {
		t, err := expandSystemSdnConnectorCompartmentListCompartmentId2edl(d, v, "compartment_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compartment-id"] = t
		}
	}

	return &obj, nil
}
