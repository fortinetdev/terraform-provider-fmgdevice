// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device SystemHaLic

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemHaLic() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaLicCreate,
		Read:   resourceSystemHaLicRead,
		Update: resourceSystemHaLicUpdate,
		Delete: resourceSystemHaLicDelete,

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
			"last_sync_seat": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"last_sync_seat_dlp": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"last_sync_seat_fcas": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"last_sync_seat_fnbi": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"last_sync_time": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"serial_num": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemHaLicCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemHaLic(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemHaLic resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("serial_num")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemHaLic(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemHaLic(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemHaLic resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemHaLic(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemHaLic resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "serial_num"))

	return resourceSystemHaLicRead(d, m)
}

func resourceSystemHaLicUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemHaLic(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaLic resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemHaLic(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaLic resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "serial_num"))

	return resourceSystemHaLicRead(d, m)
}

func resourceSystemHaLicDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteSystemHaLic(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHaLic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaLicRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemHaLic(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemHaLic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaLic(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaLic resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaLicSerialNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHaLic(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("serial_num", flattenSystemHaLicSerialNum(o["serial-num"], d, "serial_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial-num"], "SystemHaLic-SerialNum"); ok {
			if err = d.Set("serial_num", vv); err != nil {
				return fmt.Errorf("Error reading serial_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial_num: %v", err)
		}
	}

	return nil
}

func flattenSystemHaLicFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaLicLastSyncSeat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaLicLastSyncSeatDlp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaLicLastSyncSeatFcas(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaLicLastSyncSeatFnbi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaLicLastSyncTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaLicSerialNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHaLic(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("last_sync_seat"); ok || d.HasChange("last_sync_seat") {
		t, err := expandSystemHaLicLastSyncSeat(d, v, "last_sync_seat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-sync-seat"] = t
		}
	}

	if v, ok := d.GetOk("last_sync_seat_dlp"); ok || d.HasChange("last_sync_seat_dlp") {
		t, err := expandSystemHaLicLastSyncSeatDlp(d, v, "last_sync_seat_dlp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-sync-seat-dlp"] = t
		}
	}

	if v, ok := d.GetOk("last_sync_seat_fcas"); ok || d.HasChange("last_sync_seat_fcas") {
		t, err := expandSystemHaLicLastSyncSeatFcas(d, v, "last_sync_seat_fcas")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-sync-seat-fcas"] = t
		}
	}

	if v, ok := d.GetOk("last_sync_seat_fnbi"); ok || d.HasChange("last_sync_seat_fnbi") {
		t, err := expandSystemHaLicLastSyncSeatFnbi(d, v, "last_sync_seat_fnbi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-sync-seat-fnbi"] = t
		}
	}

	if v, ok := d.GetOk("last_sync_time"); ok || d.HasChange("last_sync_time") {
		t, err := expandSystemHaLicLastSyncTime(d, v, "last_sync_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-sync-time"] = t
		}
	}

	if v, ok := d.GetOk("serial_num"); ok || d.HasChange("serial_num") {
		t, err := expandSystemHaLicSerialNum(d, v, "serial_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-num"] = t
		}
	}

	return &obj, nil
}
