// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device SystemAdminTrusthosts

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdminTrusthosts() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminTrusthostsCreate,
		Read:   resourceSystemAdminTrusthostsRead,
		Update: resourceSystemAdminTrusthostsUpdate,
		Delete: resourceSystemAdminTrusthostsDelete,

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
			"admin": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ipv4": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAdminTrusthostsCreate(d *schema.ResourceData, m interface{}) error {
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
	admin := d.Get("admin").(string)
	paradict["device"] = device_name
	paradict["admin"] = admin

	obj, err := getObjectSystemAdminTrusthosts(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminTrusthosts resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemAdminTrusthosts(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemAdminTrusthosts(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemAdminTrusthosts resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemAdminTrusthosts(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemAdminTrusthosts resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAdminTrusthostsRead(d, m)
}

func resourceSystemAdminTrusthostsUpdate(d *schema.ResourceData, m interface{}) error {
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
	admin := d.Get("admin").(string)
	paradict["device"] = device_name
	paradict["admin"] = admin

	obj, err := getObjectSystemAdminTrusthosts(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminTrusthosts resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemAdminTrusthosts(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminTrusthosts resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAdminTrusthostsRead(d, m)
}

func resourceSystemAdminTrusthostsDelete(d *schema.ResourceData, m interface{}) error {
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
	admin := d.Get("admin").(string)
	paradict["device"] = device_name
	paradict["admin"] = admin

	wsParams["adom"] = adomv

	err = c.DeleteSystemAdminTrusthosts(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminTrusthosts resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminTrusthostsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	admin := d.Get("admin").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if admin == "" {
		admin = importOptionChecking(m.(*FortiClient).Cfg, "admin")
		if admin == "" {
			return fmt.Errorf("Parameter admin is missing")
		}
		if err = d.Set("admin", admin); err != nil {
			return fmt.Errorf("Error set params admin: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["admin"] = admin

	o, err := c.ReadSystemAdminTrusthosts(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemAdminTrusthosts resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminTrusthosts(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminTrusthosts resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminTrusthostsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTrusthostsIpv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminTrusthostsIpv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTrusthostsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminTrusthosts(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemAdminTrusthostsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemAdminTrusthosts-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ipv4", flattenSystemAdminTrusthostsIpv42edl(o["ipv4"], d, "ipv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4"], "SystemAdminTrusthosts-Ipv4"); ok {
			if err = d.Set("ipv4", vv); err != nil {
				return fmt.Errorf("Error reading ipv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenSystemAdminTrusthostsIpv62edl(o["ipv6"], d, "ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6"], "SystemAdminTrusthosts-Ipv6"); ok {
			if err = d.Set("ipv6", vv); err != nil {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemAdminTrusthostsType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemAdminTrusthosts-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminTrusthostsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminTrusthostsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthostsIpv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminTrusthostsIpv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthostsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminTrusthosts(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemAdminTrusthostsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ipv4"); ok || d.HasChange("ipv4") {
		t, err := expandSystemAdminTrusthostsIpv42edl(d, v, "ipv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok || d.HasChange("ipv6") {
		t, err := expandSystemAdminTrusthostsIpv62edl(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemAdminTrusthostsType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
