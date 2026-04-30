// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Custom Security Fabric permissions.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAccprofileSecfabgrpPermission() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAccprofileSecfabgrpPermissionUpdate,
		Read:   resourceSystemAccprofileSecfabgrpPermissionRead,
		Update: resourceSystemAccprofileSecfabgrpPermissionUpdate,
		Delete: resourceSystemAccprofileSecfabgrpPermissionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"accprofile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"csffoo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"csfsys": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mmsgtp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAccprofileSecfabgrpPermissionUpdate(d *schema.ResourceData, m interface{}) error {
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
	accprofile := d.Get("accprofile").(string)
	paradict["device"] = device_name
	paradict["accprofile"] = accprofile

	obj, err := getObjectSystemAccprofileSecfabgrpPermission(d, false)
	if err != nil {
		return fmt.Errorf("Error updating SystemAccprofileSecfabgrpPermission resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemAccprofileSecfabgrpPermission(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAccprofileSecfabgrpPermission resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemAccprofileSecfabgrpPermission")

	return resourceSystemAccprofileSecfabgrpPermissionRead(d, m)
}

func resourceSystemAccprofileSecfabgrpPermissionDelete(d *schema.ResourceData, m interface{}) error {
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
	accprofile := d.Get("accprofile").(string)
	paradict["device"] = device_name
	paradict["accprofile"] = accprofile

	obj, err := getObjectSystemAccprofileSecfabgrpPermission(d, true)

	if err != nil {
		return fmt.Errorf("Error updating SystemAccprofileSecfabgrpPermission resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemAccprofileSecfabgrpPermission(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing SystemAccprofileSecfabgrpPermission resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAccprofileSecfabgrpPermissionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	accprofile := d.Get("accprofile").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if accprofile == "" {
		accprofile = importOptionChecking(m.(*FortiClient).Cfg, "accprofile")
		if accprofile == "" {
			return fmt.Errorf("Parameter accprofile is missing")
		}
		if err = d.Set("accprofile", accprofile); err != nil {
			return fmt.Errorf("Error set params accprofile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["accprofile"] = accprofile

	o, err := c.ReadSystemAccprofileSecfabgrpPermission(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemAccprofileSecfabgrpPermission resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAccprofileSecfabgrpPermission(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAccprofileSecfabgrpPermission resource from API: %v", err)
	}
	return nil
}

func flattenSystemAccprofileSecfabgrpPermissionCsffoo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileSecfabgrpPermissionCsfsys2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileSecfabgrpPermissionMmsgtp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAccprofileSecfabgrpPermission(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("csffoo", flattenSystemAccprofileSecfabgrpPermissionCsffoo2edl(o["csffoo"], d, "csffoo")); err != nil {
		if vv, ok := fortiAPIPatch(o["csffoo"], "SystemAccprofileSecfabgrpPermission-Csffoo"); ok {
			if err = d.Set("csffoo", vv); err != nil {
				return fmt.Errorf("Error reading csffoo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading csffoo: %v", err)
		}
	}

	if err = d.Set("csfsys", flattenSystemAccprofileSecfabgrpPermissionCsfsys2edl(o["csfsys"], d, "csfsys")); err != nil {
		if vv, ok := fortiAPIPatch(o["csfsys"], "SystemAccprofileSecfabgrpPermission-Csfsys"); ok {
			if err = d.Set("csfsys", vv); err != nil {
				return fmt.Errorf("Error reading csfsys: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading csfsys: %v", err)
		}
	}

	if err = d.Set("mmsgtp", flattenSystemAccprofileSecfabgrpPermissionMmsgtp2edl(o["mmsgtp"], d, "mmsgtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["mmsgtp"], "SystemAccprofileSecfabgrpPermission-Mmsgtp"); ok {
			if err = d.Set("mmsgtp", vv); err != nil {
				return fmt.Errorf("Error reading mmsgtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mmsgtp: %v", err)
		}
	}

	return nil
}

func flattenSystemAccprofileSecfabgrpPermissionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAccprofileSecfabgrpPermissionCsffoo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSecfabgrpPermissionCsfsys2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSecfabgrpPermissionMmsgtp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAccprofileSecfabgrpPermission(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("csffoo"); ok || d.HasChange("csffoo") {
		t, err := expandSystemAccprofileSecfabgrpPermissionCsffoo2edl(d, v, "csffoo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csffoo"] = t
		}
	}

	if v, ok := d.GetOk("csfsys"); ok || d.HasChange("csfsys") {
		t, err := expandSystemAccprofileSecfabgrpPermissionCsfsys2edl(d, v, "csfsys")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csfsys"] = t
		}
	}

	if v, ok := d.GetOk("mmsgtp"); ok || d.HasChange("mmsgtp") {
		t, err := expandSystemAccprofileSecfabgrpPermissionMmsgtp2edl(d, v, "mmsgtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mmsgtp"] = t
		}
	}

	return &obj, nil
}
