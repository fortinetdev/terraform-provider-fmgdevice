// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCentralManagementServerList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCentralManagementServerListCreate,
		Read:   resourceSystemCentralManagementServerListRead,
		Update: resourceSystemCentralManagementServerListUpdate,
		Delete: resourceSystemCentralManagementServerListDelete,

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
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"server_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_address6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemCentralManagementServerListCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemCentralManagementServerList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCentralManagementServerList resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCentralManagementServerList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemCentralManagementServerList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemCentralManagementServerListRead(d, m)
}

func resourceSystemCentralManagementServerListUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemCentralManagementServerList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagementServerList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCentralManagementServerList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagementServerList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemCentralManagementServerListRead(d, m)
}

func resourceSystemCentralManagementServerListDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemCentralManagementServerList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCentralManagementServerList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCentralManagementServerListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCentralManagementServerList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemCentralManagementServerList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCentralManagementServerList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCentralManagementServerList resource from API: %v", err)
	}
	return nil
}

func flattenSystemCentralManagementServerListAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListServerAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListServerAddress62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListServerType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemCentralManagementServerList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_type", flattenSystemCentralManagementServerListAddrType2edl(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "SystemCentralManagementServerList-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenSystemCentralManagementServerListFqdn2edl(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "SystemCentralManagementServerList-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemCentralManagementServerListId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemCentralManagementServerList-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("server_address", flattenSystemCentralManagementServerListServerAddress2edl(o["server-address"], d, "server_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-address"], "SystemCentralManagementServerList-ServerAddress"); ok {
			if err = d.Set("server_address", vv); err != nil {
				return fmt.Errorf("Error reading server_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_address: %v", err)
		}
	}

	if err = d.Set("server_address6", flattenSystemCentralManagementServerListServerAddress62edl(o["server-address6"], d, "server_address6")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-address6"], "SystemCentralManagementServerList-ServerAddress6"); ok {
			if err = d.Set("server_address6", vv); err != nil {
				return fmt.Errorf("Error reading server_address6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_address6: %v", err)
		}
	}

	if err = d.Set("server_type", flattenSystemCentralManagementServerListServerType2edl(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "SystemCentralManagementServerList-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	return nil
}

func flattenSystemCentralManagementServerListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCentralManagementServerListAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListServerAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListServerAddress62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListServerType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemCentralManagementServerList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandSystemCentralManagementServerListAddrType2edl(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok || d.HasChange("fqdn") {
		t, err := expandSystemCentralManagementServerListFqdn2edl(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemCentralManagementServerListId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("server_address"); ok || d.HasChange("server_address") {
		t, err := expandSystemCentralManagementServerListServerAddress2edl(d, v, "server_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-address"] = t
		}
	}

	if v, ok := d.GetOk("server_address6"); ok || d.HasChange("server_address6") {
		t, err := expandSystemCentralManagementServerListServerAddress62edl(d, v, "server_address6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-address6"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandSystemCentralManagementServerListServerType2edl(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	return &obj, nil
}
