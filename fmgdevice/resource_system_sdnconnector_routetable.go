// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Azure route table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdnConnectorRouteTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdnConnectorRouteTableCreate,
		Read:   resourceSystemSdnConnectorRouteTableRead,
		Update: resourceSystemSdnConnectorRouteTableUpdate,
		Delete: resourceSystemSdnConnectorRouteTableDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"route": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"next_hop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemSdnConnectorRouteTableCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdnConnectorRouteTable(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdnConnectorRouteTable resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemSdnConnectorRouteTable(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemSdnConnectorRouteTable(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemSdnConnectorRouteTable resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemSdnConnectorRouteTable(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemSdnConnectorRouteTable resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSdnConnectorRouteTableRead(d, m)
}

func resourceSystemSdnConnectorRouteTableUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdnConnectorRouteTable(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnectorRouteTable resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSdnConnectorRouteTable(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnectorRouteTable resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSdnConnectorRouteTableRead(d, m)
}

func resourceSystemSdnConnectorRouteTableDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSdnConnectorRouteTable(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdnConnectorRouteTable resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdnConnectorRouteTableRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdnConnectorRouteTable(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemSdnConnectorRouteTable resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdnConnectorRouteTable(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnConnectorRouteTable resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdnConnectorRouteTableName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableResourceGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableRoute2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSdnConnectorRouteTableRouteName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSdnConnectorRouteTable-Route-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop"
		if _, ok := i["next-hop"]; ok {
			v := flattenSystemSdnConnectorRouteTableRouteNextHop2edl(i["next-hop"], d, pre_append)
			tmp["next_hop"] = fortiAPISubPartPatch(v, "SystemSdnConnectorRouteTable-Route-NextHop")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdnConnectorRouteTableRouteName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableRouteNextHop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableSubscriptionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdnConnectorRouteTable(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenSystemSdnConnectorRouteTableName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSdnConnectorRouteTable-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("resource_group", flattenSystemSdnConnectorRouteTableResourceGroup2edl(o["resource-group"], d, "resource_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["resource-group"], "SystemSdnConnectorRouteTable-ResourceGroup"); ok {
			if err = d.Set("resource_group", vv); err != nil {
				return fmt.Errorf("Error reading resource_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resource_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("route", flattenSystemSdnConnectorRouteTableRoute2edl(o["route"], d, "route")); err != nil {
			if vv, ok := fortiAPIPatch(o["route"], "SystemSdnConnectorRouteTable-Route"); ok {
				if err = d.Set("route", vv); err != nil {
					return fmt.Errorf("Error reading route: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading route: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("route"); ok {
			if err = d.Set("route", flattenSystemSdnConnectorRouteTableRoute2edl(o["route"], d, "route")); err != nil {
				if vv, ok := fortiAPIPatch(o["route"], "SystemSdnConnectorRouteTable-Route"); ok {
					if err = d.Set("route", vv); err != nil {
						return fmt.Errorf("Error reading route: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading route: %v", err)
				}
			}
		}
	}

	if err = d.Set("subscription_id", flattenSystemSdnConnectorRouteTableSubscriptionId2edl(o["subscription-id"], d, "subscription_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["subscription-id"], "SystemSdnConnectorRouteTable-SubscriptionId"); ok {
			if err = d.Set("subscription_id", vv); err != nil {
				return fmt.Errorf("Error reading subscription_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subscription_id: %v", err)
		}
	}

	return nil
}

func flattenSystemSdnConnectorRouteTableFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdnConnectorRouteTableName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableResourceGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSdnConnectorRouteTableRouteName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop"], _ = expandSystemSdnConnectorRouteTableRouteNextHop2edl(d, i["next_hop"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorRouteTableRouteName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableRouteNextHop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableSubscriptionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdnConnectorRouteTable(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSdnConnectorRouteTableName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("resource_group"); ok || d.HasChange("resource_group") {
		t, err := expandSystemSdnConnectorRouteTableResourceGroup2edl(d, v, "resource_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-group"] = t
		}
	}

	if v, ok := d.GetOk("route"); ok || d.HasChange("route") {
		t, err := expandSystemSdnConnectorRouteTableRoute2edl(d, v, "route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route"] = t
		}
	}

	if v, ok := d.GetOk("subscription_id"); ok || d.HasChange("subscription_id") {
		t, err := expandSystemSdnConnectorRouteTableSubscriptionId2edl(d, v, "subscription_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subscription-id"] = t
		}
	}

	return &obj, nil
}
