// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device SystemClusterSyncSessionSyncFilterCustomService

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemClusterSyncSessionSyncFilterCustomService() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemClusterSyncSessionSyncFilterCustomServiceCreate,
		Read:   resourceSystemClusterSyncSessionSyncFilterCustomServiceRead,
		Update: resourceSystemClusterSyncSessionSyncFilterCustomServiceUpdate,
		Delete: resourceSystemClusterSyncSessionSyncFilterCustomServiceDelete,

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
			"cluster_sync": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dst_port_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"src_port_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemClusterSyncSessionSyncFilterCustomServiceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	cluster_sync := d.Get("cluster_sync").(string)
	paradict["device"] = device_name
	paradict["cluster_sync"] = cluster_sync

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemClusterSyncSessionSyncFilterCustomService(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemClusterSyncSessionSyncFilterCustomService resource while getting object: %v", err)
	}

	_, err = c.CreateSystemClusterSyncSessionSyncFilterCustomService(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemClusterSyncSessionSyncFilterCustomService resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemClusterSyncSessionSyncFilterCustomServiceRead(d, m)
}

func resourceSystemClusterSyncSessionSyncFilterCustomServiceUpdate(d *schema.ResourceData, m interface{}) error {
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
	cluster_sync := d.Get("cluster_sync").(string)
	paradict["device"] = device_name
	paradict["cluster_sync"] = cluster_sync

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemClusterSyncSessionSyncFilterCustomService(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemClusterSyncSessionSyncFilterCustomService resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemClusterSyncSessionSyncFilterCustomService(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemClusterSyncSessionSyncFilterCustomService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemClusterSyncSessionSyncFilterCustomServiceRead(d, m)
}

func resourceSystemClusterSyncSessionSyncFilterCustomServiceDelete(d *schema.ResourceData, m interface{}) error {
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
	cluster_sync := d.Get("cluster_sync").(string)
	paradict["device"] = device_name
	paradict["cluster_sync"] = cluster_sync

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemClusterSyncSessionSyncFilterCustomService(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemClusterSyncSessionSyncFilterCustomService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemClusterSyncSessionSyncFilterCustomServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	cluster_sync := d.Get("cluster_sync").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if cluster_sync == "" {
		cluster_sync = importOptionChecking(m.(*FortiClient).Cfg, "cluster_sync")
		if cluster_sync == "" {
			return fmt.Errorf("Parameter cluster_sync is missing")
		}
		if err = d.Set("cluster_sync", cluster_sync); err != nil {
			return fmt.Errorf("Error set params cluster_sync: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["cluster_sync"] = cluster_sync

	o, err := c.ReadSystemClusterSyncSessionSyncFilterCustomService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemClusterSyncSessionSyncFilterCustomService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemClusterSyncSessionSyncFilterCustomService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemClusterSyncSessionSyncFilterCustomService resource from API: %v", err)
	}
	return nil
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemClusterSyncSessionSyncFilterCustomService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dst_port_range", flattenSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange3rdl(o["dst-port-range"], d, "dst_port_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-port-range"], "SystemClusterSyncSessionSyncFilterCustomService-DstPortRange"); ok {
			if err = d.Set("dst_port_range", vv); err != nil {
				return fmt.Errorf("Error reading dst_port_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_port_range: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemClusterSyncSessionSyncFilterCustomServiceId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemClusterSyncSessionSyncFilterCustomService-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("src_port_range", flattenSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange3rdl(o["src-port-range"], d, "src_port_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-port-range"], "SystemClusterSyncSessionSyncFilterCustomService-SrcPortRange"); ok {
			if err = d.Set("src_port_range", vv); err != nil {
				return fmt.Errorf("Error reading src_port_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_port_range: %v", err)
		}
	}

	return nil
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemClusterSyncSessionSyncFilterCustomService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dst_port_range"); ok || d.HasChange("dst_port_range") {
		t, err := expandSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange3rdl(d, v, "dst_port_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-port-range"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemClusterSyncSessionSyncFilterCustomServiceId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("src_port_range"); ok || d.HasChange("src_port_range") {
		t, err := expandSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange3rdl(d, v, "src_port_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-port-range"] = t
		}
	}

	return &obj, nil
}
