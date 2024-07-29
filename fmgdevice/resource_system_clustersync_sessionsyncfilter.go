// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device SystemClusterSyncSessionSyncFilter

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemClusterSyncSessionSyncFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemClusterSyncSessionSyncFilterUpdate,
		Read:   resourceSystemClusterSyncSessionSyncFilterRead,
		Update: resourceSystemClusterSyncSessionSyncFilterUpdate,
		Delete: resourceSystemClusterSyncSessionSyncFilterDelete,

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
			"custom_service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_port_range": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"src_port_range": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
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

func resourceSystemClusterSyncSessionSyncFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	cluster_sync := d.Get("cluster_sync").(string)
	paradict["device"] = device_name
	paradict["cluster_sync"] = cluster_sync

	obj, err := getObjectSystemClusterSyncSessionSyncFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemClusterSyncSessionSyncFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemClusterSyncSessionSyncFilter(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemClusterSyncSessionSyncFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemClusterSyncSessionSyncFilter")

	return resourceSystemClusterSyncSessionSyncFilterRead(d, m)
}

func resourceSystemClusterSyncSessionSyncFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	cluster_sync := d.Get("cluster_sync").(string)
	paradict["device"] = device_name
	paradict["cluster_sync"] = cluster_sync

	err = c.DeleteSystemClusterSyncSessionSyncFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemClusterSyncSessionSyncFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemClusterSyncSessionSyncFilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemClusterSyncSessionSyncFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemClusterSyncSessionSyncFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemClusterSyncSessionSyncFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemClusterSyncSessionSyncFilter resource from API: %v", err)
	}
	return nil
}

func flattenSystemClusterSyncSessionSyncFilterCustomService2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port_range"
		if _, ok := i["dst-port-range"]; ok {
			v := flattenSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange2edl(i["dst-port-range"], d, pre_append)
			tmp["dst_port_range"] = fortiAPISubPartPatch(v, "SystemClusterSyncSessionSyncFilter-CustomService-DstPortRange")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemClusterSyncSessionSyncFilterCustomServiceId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemClusterSyncSessionSyncFilter-CustomService-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := i["src-port-range"]; ok {
			v := flattenSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange2edl(i["src-port-range"], d, pre_append)
			tmp["src_port_range"] = fortiAPISubPartPatch(v, "SystemClusterSyncSessionSyncFilter-CustomService-SrcPortRange")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterDstaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemClusterSyncSessionSyncFilterDstaddr62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterDstintf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemClusterSyncSessionSyncFilterSrcaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemClusterSyncSessionSyncFilterSrcaddr62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterSrcintf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemClusterSyncSessionSyncFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("custom_service", flattenSystemClusterSyncSessionSyncFilterCustomService2edl(o["custom-service"], d, "custom_service")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-service"], "SystemClusterSyncSessionSyncFilter-CustomService"); ok {
				if err = d.Set("custom_service", vv); err != nil {
					return fmt.Errorf("Error reading custom_service: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_service"); ok {
			if err = d.Set("custom_service", flattenSystemClusterSyncSessionSyncFilterCustomService2edl(o["custom-service"], d, "custom_service")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-service"], "SystemClusterSyncSessionSyncFilter-CustomService"); ok {
					if err = d.Set("custom_service", vv); err != nil {
						return fmt.Errorf("Error reading custom_service: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_service: %v", err)
				}
			}
		}
	}

	if err = d.Set("dstaddr", flattenSystemClusterSyncSessionSyncFilterDstaddr2edl(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "SystemClusterSyncSessionSyncFilter-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenSystemClusterSyncSessionSyncFilterDstaddr62edl(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "SystemClusterSyncSessionSyncFilter-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenSystemClusterSyncSessionSyncFilterDstintf2edl(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "SystemClusterSyncSessionSyncFilter-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenSystemClusterSyncSessionSyncFilterSrcaddr2edl(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "SystemClusterSyncSessionSyncFilter-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenSystemClusterSyncSessionSyncFilterSrcaddr62edl(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "SystemClusterSyncSessionSyncFilter-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenSystemClusterSyncSessionSyncFilterSrcintf2edl(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "SystemClusterSyncSessionSyncFilter-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	return nil
}

func flattenSystemClusterSyncSessionSyncFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemClusterSyncSessionSyncFilterCustomService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port_range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-port-range"], _ = expandSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange2edl(d, i["dst_port_range"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemClusterSyncSessionSyncFilterCustomServiceId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-port-range"], _ = expandSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange2edl(d, i["src_port_range"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterDstaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemClusterSyncSessionSyncFilterDstaddr62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterDstintf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemClusterSyncSessionSyncFilterSrcaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemClusterSyncSessionSyncFilterSrcaddr62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterSrcintf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemClusterSyncSessionSyncFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("custom_service"); ok || d.HasChange("custom_service") {
		t, err := expandSystemClusterSyncSessionSyncFilterCustomService2edl(d, v, "custom_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-service"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandSystemClusterSyncSessionSyncFilterDstaddr2edl(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandSystemClusterSyncSessionSyncFilterDstaddr62edl(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandSystemClusterSyncSessionSyncFilterDstintf2edl(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandSystemClusterSyncSessionSyncFilterSrcaddr2edl(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandSystemClusterSyncSessionSyncFilterSrcaddr62edl(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandSystemClusterSyncSessionSyncFilterSrcintf2edl(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	return &obj, nil
}
