// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure speed test server list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSpeedTestServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSpeedTestServerCreate,
		Read:   resourceSystemSpeedTestServerRead,
		Update: resourceSystemSpeedTestServerUpdate,
		Delete: resourceSystemSpeedTestServerDelete,

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distance": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"latitude": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"longitude": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"user": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemSpeedTestServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemSpeedTestServer(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSpeedTestServer resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSpeedTestServer(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSpeedTestServer resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSpeedTestServerRead(d, m)
}

func resourceSystemSpeedTestServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemSpeedTestServer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestServer resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSpeedTestServer(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSpeedTestServerRead(d, m)
}

func resourceSystemSpeedTestServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteSystemSpeedTestServer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSpeedTestServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSpeedTestServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if device_vdom == "" {
		device_vdom = importOptionChecking(m.(*FortiClient).Cfg, "device_vdom")
		if device_vdom == "" {
			return fmt.Errorf("Parameter device_vdom is missing")
		}
		if err = d.Set("device_vdom", device_vdom); err != nil {
			return fmt.Errorf("Error set params device_vdom: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadSystemSpeedTestServer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpeedTestServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSpeedTestServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpeedTestServer resource from API: %v", err)
	}
	return nil
}

func flattenSystemSpeedTestServerHost(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			v := flattenSystemSpeedTestServerHostDistance(i["distance"], d, pre_append)
			tmp["distance"] = fortiAPISubPartPatch(v, "SystemSpeedTestServer-Host-Distance")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSpeedTestServerHostId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSpeedTestServer-Host-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemSpeedTestServerHostIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemSpeedTestServer-Host-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latitude"
		if _, ok := i["latitude"]; ok {
			v := flattenSystemSpeedTestServerHostLatitude(i["latitude"], d, pre_append)
			tmp["latitude"] = fortiAPISubPartPatch(v, "SystemSpeedTestServer-Host-Latitude")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "longitude"
		if _, ok := i["longitude"]; ok {
			v := flattenSystemSpeedTestServerHostLongitude(i["longitude"], d, pre_append)
			tmp["longitude"] = fortiAPISubPartPatch(v, "SystemSpeedTestServer-Host-Longitude")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenSystemSpeedTestServerHostPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "SystemSpeedTestServer-Host-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := i["user"]; ok {
			v := flattenSystemSpeedTestServerHostUser(i["user"], d, pre_append)
			tmp["user"] = fortiAPISubPartPatch(v, "SystemSpeedTestServer-Host-User")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSpeedTestServerHostDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestServerHostId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestServerHostIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestServerHostLatitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestServerHostLongitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestServerHostPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestServerHostUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSpeedTestServerTimestamp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSpeedTestServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("host", flattenSystemSpeedTestServerHost(o["host"], d, "host")); err != nil {
			if vv, ok := fortiAPIPatch(o["host"], "SystemSpeedTestServer-Host"); ok {
				if err = d.Set("host", vv); err != nil {
					return fmt.Errorf("Error reading host: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading host: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("host"); ok {
			if err = d.Set("host", flattenSystemSpeedTestServerHost(o["host"], d, "host")); err != nil {
				if vv, ok := fortiAPIPatch(o["host"], "SystemSpeedTestServer-Host"); ok {
					if err = d.Set("host", vv); err != nil {
						return fmt.Errorf("Error reading host: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading host: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSystemSpeedTestServerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSpeedTestServer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("timestamp", flattenSystemSpeedTestServerTimestamp(o["timestamp"], d, "timestamp")); err != nil {
		if vv, ok := fortiAPIPatch(o["timestamp"], "SystemSpeedTestServer-Timestamp"); ok {
			if err = d.Set("timestamp", vv); err != nil {
				return fmt.Errorf("Error reading timestamp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timestamp: %v", err)
		}
	}

	return nil
}

func flattenSystemSpeedTestServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSpeedTestServerHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distance"], _ = expandSystemSpeedTestServerHostDistance(d, i["distance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSpeedTestServerHostId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemSpeedTestServerHostIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latitude"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latitude"], _ = expandSystemSpeedTestServerHostLatitude(d, i["latitude"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "longitude"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["longitude"], _ = expandSystemSpeedTestServerHostLongitude(d, i["longitude"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandSystemSpeedTestServerHostPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandSystemSpeedTestServerHostPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user"], _ = expandSystemSpeedTestServerHostUser(d, i["user"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSpeedTestServerHostDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerHostId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerHostIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerHostLatitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerHostLongitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerHostPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSpeedTestServerHostPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerHostUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerTimestamp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSpeedTestServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandSystemSpeedTestServerHost(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSpeedTestServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("timestamp"); ok || d.HasChange("timestamp") {
		t, err := expandSystemSpeedTestServerTimestamp(d, v, "timestamp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timestamp"] = t
		}
	}

	return &obj, nil
}
