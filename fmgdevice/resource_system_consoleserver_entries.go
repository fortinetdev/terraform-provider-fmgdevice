// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Entry used by console server.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemConsoleServerEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemConsoleServerEntriesCreate,
		Read:   resourceSystemConsoleServerEntriesRead,
		Update: resourceSystemConsoleServerEntriesUpdate,
		Delete: resourceSystemConsoleServerEntriesDelete,

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
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"slot_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemConsoleServerEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemConsoleServerEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemConsoleServerEntries resource while getting object: %v", err)
	}

	_, err = c.CreateSystemConsoleServerEntries(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemConsoleServerEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "slot_id")))

	return resourceSystemConsoleServerEntriesRead(d, m)
}

func resourceSystemConsoleServerEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemConsoleServerEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemConsoleServerEntries resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemConsoleServerEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemConsoleServerEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "slot_id")))

	return resourceSystemConsoleServerEntriesRead(d, m)
}

func resourceSystemConsoleServerEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemConsoleServerEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemConsoleServerEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemConsoleServerEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemConsoleServerEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemConsoleServerEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemConsoleServerEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemConsoleServerEntries resource from API: %v", err)
	}
	return nil
}

func flattenSystemConsoleServerEntriesPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemConsoleServerEntriesSlotId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemConsoleServerEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("port", flattenSystemConsoleServerEntriesPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemConsoleServerEntries-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("slot_id", flattenSystemConsoleServerEntriesSlotId2edl(o["slot-id"], d, "slot_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["slot-id"], "SystemConsoleServerEntries-SlotId"); ok {
			if err = d.Set("slot_id", vv); err != nil {
				return fmt.Errorf("Error reading slot_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slot_id: %v", err)
		}
	}

	return nil
}

func flattenSystemConsoleServerEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemConsoleServerEntriesPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleServerEntriesSlotId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemConsoleServerEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemConsoleServerEntriesPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("slot_id"); ok || d.HasChange("slot_id") {
		t, err := expandSystemConsoleServerEntriesSlotId2edl(d, v, "slot_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slot-id"] = t
		}
	}

	return &obj, nil
}
