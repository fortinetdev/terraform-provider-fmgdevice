// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure AWS external account list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdnConnectorExternalAccountList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdnConnectorExternalAccountListCreate,
		Read:   resourceSystemSdnConnectorExternalAccountListRead,
		Update: resourceSystemSdnConnectorExternalAccountListUpdate,
		Delete: resourceSystemSdnConnectorExternalAccountListDelete,

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
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"region_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"role_arn": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemSdnConnectorExternalAccountListCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdnConnectorExternalAccountList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdnConnectorExternalAccountList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("role_arn")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemSdnConnectorExternalAccountList(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemSdnConnectorExternalAccountList(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemSdnConnectorExternalAccountList resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemSdnConnectorExternalAccountList(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemSdnConnectorExternalAccountList resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "role_arn"))

	return resourceSystemSdnConnectorExternalAccountListRead(d, m)
}

func resourceSystemSdnConnectorExternalAccountListUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdnConnectorExternalAccountList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnectorExternalAccountList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSdnConnectorExternalAccountList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnectorExternalAccountList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "role_arn"))

	return resourceSystemSdnConnectorExternalAccountListRead(d, m)
}

func resourceSystemSdnConnectorExternalAccountListDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSdnConnectorExternalAccountList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdnConnectorExternalAccountList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdnConnectorExternalAccountListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdnConnectorExternalAccountList(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemSdnConnectorExternalAccountList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdnConnectorExternalAccountList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnConnectorExternalAccountList resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdnConnectorExternalAccountListExternalId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorExternalAccountListRegionList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnConnectorExternalAccountListRoleArn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdnConnectorExternalAccountList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("external_id", flattenSystemSdnConnectorExternalAccountListExternalId2edl(o["external-id"], d, "external_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-id"], "SystemSdnConnectorExternalAccountList-ExternalId"); ok {
			if err = d.Set("external_id", vv); err != nil {
				return fmt.Errorf("Error reading external_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_id: %v", err)
		}
	}

	if err = d.Set("region_list", flattenSystemSdnConnectorExternalAccountListRegionList2edl(o["region-list"], d, "region_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["region-list"], "SystemSdnConnectorExternalAccountList-RegionList"); ok {
			if err = d.Set("region_list", vv); err != nil {
				return fmt.Errorf("Error reading region_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region_list: %v", err)
		}
	}

	if err = d.Set("role_arn", flattenSystemSdnConnectorExternalAccountListRoleArn2edl(o["role-arn"], d, "role_arn")); err != nil {
		if vv, ok := fortiAPIPatch(o["role-arn"], "SystemSdnConnectorExternalAccountList-RoleArn"); ok {
			if err = d.Set("role_arn", vv); err != nil {
				return fmt.Errorf("Error reading role_arn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role_arn: %v", err)
		}
	}

	return nil
}

func flattenSystemSdnConnectorExternalAccountListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdnConnectorExternalAccountListExternalId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorExternalAccountListRegionList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorExternalAccountListRoleArn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdnConnectorExternalAccountList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("external_id"); ok || d.HasChange("external_id") {
		t, err := expandSystemSdnConnectorExternalAccountListExternalId2edl(d, v, "external_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-id"] = t
		}
	}

	if v, ok := d.GetOk("region_list"); ok || d.HasChange("region_list") {
		t, err := expandSystemSdnConnectorExternalAccountListRegionList2edl(d, v, "region_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region-list"] = t
		}
	}

	if v, ok := d.GetOk("role_arn"); ok || d.HasChange("role_arn") {
		t, err := expandSystemSdnConnectorExternalAccountListRoleArn2edl(d, v, "role_arn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role-arn"] = t
		}
	}

	return &obj, nil
}
