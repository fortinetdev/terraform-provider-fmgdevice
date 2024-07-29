// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemStandaloneClusterClusterPeerSessionSyncFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemStandaloneClusterClusterPeerSessionSyncFilterUpdate,
		Read:   resourceSystemStandaloneClusterClusterPeerSessionSyncFilterRead,
		Update: resourceSystemStandaloneClusterClusterPeerSessionSyncFilterUpdate,
		Delete: resourceSystemStandaloneClusterClusterPeerSessionSyncFilterDelete,

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
			"cluster_peer": &schema.Schema{
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
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceSystemStandaloneClusterClusterPeerSessionSyncFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	cluster_peer := d.Get("cluster_peer").(string)
	paradict["device"] = device_name
	paradict["cluster_peer"] = cluster_peer

	obj, err := getObjectSystemStandaloneClusterClusterPeerSessionSyncFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneClusterClusterPeerSessionSyncFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemStandaloneClusterClusterPeerSessionSyncFilter(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneClusterClusterPeerSessionSyncFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemStandaloneClusterClusterPeerSessionSyncFilter")

	return resourceSystemStandaloneClusterClusterPeerSessionSyncFilterRead(d, m)
}

func resourceSystemStandaloneClusterClusterPeerSessionSyncFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	cluster_peer := d.Get("cluster_peer").(string)
	paradict["device"] = device_name
	paradict["cluster_peer"] = cluster_peer

	err = c.DeleteSystemStandaloneClusterClusterPeerSessionSyncFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemStandaloneClusterClusterPeerSessionSyncFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStandaloneClusterClusterPeerSessionSyncFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	cluster_peer := d.Get("cluster_peer").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if cluster_peer == "" {
		cluster_peer = importOptionChecking(m.(*FortiClient).Cfg, "cluster_peer")
		if cluster_peer == "" {
			return fmt.Errorf("Parameter cluster_peer is missing")
		}
		if err = d.Set("cluster_peer", cluster_peer); err != nil {
			return fmt.Errorf("Error set params cluster_peer: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["cluster_peer"] = cluster_peer

	o, err := c.ReadSystemStandaloneClusterClusterPeerSessionSyncFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneClusterClusterPeerSessionSyncFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemStandaloneClusterClusterPeerSessionSyncFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneClusterClusterPeerSessionSyncFilter resource from API: %v", err)
	}
	return nil
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange3rdl(i["dst-port-range"], d, pre_append)
			tmp["dst_port_range"] = fortiAPISubPartPatch(v, "SystemStandaloneClusterClusterPeerSessionSyncFilter-CustomService-DstPortRange")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemStandaloneClusterClusterPeerSessionSyncFilter-CustomService-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := i["src-port-range"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange3rdl(i["src-port-range"], d, pre_append)
			tmp["src_port_range"] = fortiAPISubPartPatch(v, "SystemStandaloneClusterClusterPeerSessionSyncFilter-CustomService-SrcPortRange")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemStandaloneClusterClusterPeerSessionSyncFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("custom_service", flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService3rdl(o["custom-service"], d, "custom_service")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-service"], "SystemStandaloneClusterClusterPeerSessionSyncFilter-CustomService"); ok {
				if err = d.Set("custom_service", vv); err != nil {
					return fmt.Errorf("Error reading custom_service: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_service"); ok {
			if err = d.Set("custom_service", flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService3rdl(o["custom-service"], d, "custom_service")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-service"], "SystemStandaloneClusterClusterPeerSessionSyncFilter-CustomService"); ok {
					if err = d.Set("custom_service", vv); err != nil {
						return fmt.Errorf("Error reading custom_service: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_service: %v", err)
				}
			}
		}
	}

	if err = d.Set("dstaddr", flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr3rdl(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "SystemStandaloneClusterClusterPeerSessionSyncFilter-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr63rdl(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "SystemStandaloneClusterClusterPeerSessionSyncFilter-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf3rdl(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "SystemStandaloneClusterClusterPeerSessionSyncFilter-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr3rdl(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "SystemStandaloneClusterClusterPeerSessionSyncFilter-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr63rdl(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "SystemStandaloneClusterClusterPeerSessionSyncFilter-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf3rdl(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "SystemStandaloneClusterClusterPeerSessionSyncFilter-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	return nil
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["dst-port-range"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange3rdl(d, i["dst_port_range"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-port-range"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange3rdl(d, i["src_port_range"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemStandaloneClusterClusterPeerSessionSyncFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("custom_service"); ok || d.HasChange("custom_service") {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService3rdl(d, v, "custom_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-service"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr3rdl(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr63rdl(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf3rdl(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr3rdl(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr63rdl(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf3rdl(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	return &obj, nil
}
