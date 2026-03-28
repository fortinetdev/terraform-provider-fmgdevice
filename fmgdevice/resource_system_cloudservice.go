// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure system cloud service.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCloudService() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCloudServiceCreate,
		Read:   resourceSystemCloudServiceRead,
		Update: resourceSystemCloudServiceUpdate,
		Delete: resourceSystemCloudServiceDelete,

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
			"gck_access_token_lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gck_keyid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gck_private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gck_service_account": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"traffic_vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemCloudServiceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemCloudService(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCloudService resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemCloudService(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemCloudService(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemCloudService resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemCloudService(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemCloudService resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemCloudServiceRead(d, m)
}

func resourceSystemCloudServiceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemCloudService(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCloudService resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemCloudService(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemCloudService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemCloudServiceRead(d, m)
}

func resourceSystemCloudServiceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemCloudService(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCloudService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCloudServiceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCloudService(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemCloudService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCloudService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCloudService resource from API: %v", err)
	}
	return nil
}

func flattenSystemCloudServiceGckAccessTokenLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCloudServiceGckKeyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCloudServiceGckPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCloudServiceGckServiceAccount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCloudServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCloudServiceTrafficVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCloudServiceVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemCloudService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("gck_access_token_lifetime", flattenSystemCloudServiceGckAccessTokenLifetime(o["gck-access-token-lifetime"], d, "gck_access_token_lifetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["gck-access-token-lifetime"], "SystemCloudService-GckAccessTokenLifetime"); ok {
			if err = d.Set("gck_access_token_lifetime", vv); err != nil {
				return fmt.Errorf("Error reading gck_access_token_lifetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gck_access_token_lifetime: %v", err)
		}
	}

	if err = d.Set("gck_keyid", flattenSystemCloudServiceGckKeyid(o["gck-keyid"], d, "gck_keyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["gck-keyid"], "SystemCloudService-GckKeyid"); ok {
			if err = d.Set("gck_keyid", vv); err != nil {
				return fmt.Errorf("Error reading gck_keyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gck_keyid: %v", err)
		}
	}

	if err = d.Set("gck_private_key", flattenSystemCloudServiceGckPrivateKey(o["gck-private-key"], d, "gck_private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["gck-private-key"], "SystemCloudService-GckPrivateKey"); ok {
			if err = d.Set("gck_private_key", vv); err != nil {
				return fmt.Errorf("Error reading gck_private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gck_private_key: %v", err)
		}
	}

	if err = d.Set("gck_service_account", flattenSystemCloudServiceGckServiceAccount(o["gck-service-account"], d, "gck_service_account")); err != nil {
		if vv, ok := fortiAPIPatch(o["gck-service-account"], "SystemCloudService-GckServiceAccount"); ok {
			if err = d.Set("gck_service_account", vv); err != nil {
				return fmt.Errorf("Error reading gck_service_account: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gck_service_account: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemCloudServiceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemCloudService-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("traffic_vdom", flattenSystemCloudServiceTrafficVdom(o["traffic-vdom"], d, "traffic_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-vdom"], "SystemCloudService-TrafficVdom"); ok {
			if err = d.Set("traffic_vdom", vv); err != nil {
				return fmt.Errorf("Error reading traffic_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_vdom: %v", err)
		}
	}

	if err = d.Set("vendor", flattenSystemCloudServiceVendor(o["vendor"], d, "vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["vendor"], "SystemCloudService-Vendor"); ok {
			if err = d.Set("vendor", vv); err != nil {
				return fmt.Errorf("Error reading vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	return nil
}

func flattenSystemCloudServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCloudServiceGckAccessTokenLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCloudServiceGckKeyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCloudServiceGckPrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCloudServiceGckServiceAccount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCloudServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCloudServiceTrafficVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCloudServiceVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCloudService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("gck_access_token_lifetime"); ok || d.HasChange("gck_access_token_lifetime") {
		t, err := expandSystemCloudServiceGckAccessTokenLifetime(d, v, "gck_access_token_lifetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gck-access-token-lifetime"] = t
		}
	}

	if v, ok := d.GetOk("gck_keyid"); ok || d.HasChange("gck_keyid") {
		t, err := expandSystemCloudServiceGckKeyid(d, v, "gck_keyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gck-keyid"] = t
		}
	}

	if v, ok := d.GetOk("gck_private_key"); ok || d.HasChange("gck_private_key") {
		t, err := expandSystemCloudServiceGckPrivateKey(d, v, "gck_private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gck-private-key"] = t
		}
	}

	if v, ok := d.GetOk("gck_service_account"); ok || d.HasChange("gck_service_account") {
		t, err := expandSystemCloudServiceGckServiceAccount(d, v, "gck_service_account")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gck-service-account"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemCloudServiceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("traffic_vdom"); ok || d.HasChange("traffic_vdom") {
		t, err := expandSystemCloudServiceTrafficVdom(d, v, "traffic_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-vdom"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok || d.HasChange("vendor") {
		t, err := expandSystemCloudServiceVendor(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	return &obj, nil
}
