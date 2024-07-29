// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IKE global attributes.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIke() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIkeUpdate,
		Read:   resourceSystemIkeRead,
		Update: resourceSystemIkeUpdate,
		Delete: resourceSystemIkeDelete,

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
			"dh_group_1": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_14": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_15": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_16": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_17": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_18": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_19": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_2": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_20": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_21": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_27": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_28": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_29": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_30": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_31": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_32": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_group_5": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dh_keypair_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh_keypair_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dh_keypair_throttle": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh_multiprocess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh_worker_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"embryonic_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemIkeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemIke(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemIke resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIke(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemIke resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemIke")

	return resourceSystemIkeRead(d, m)
}

func resourceSystemIkeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemIke(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIke resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIkeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemIke(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemIke resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIke(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemIke resource from API: %v", err)
	}
	return nil
}

func flattenSystemIkeDhGroup1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup1Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup1KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup1KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup1Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup1Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup1KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup1KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup1Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup14(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup14Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup14KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup14KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup14Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup14Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup14KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup14KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup14Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup15(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup15Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup15KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup15KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup15Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup15Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup15KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup15KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup15Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup16(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup16Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup16KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup16KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup16Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup16Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup16KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup16KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup16Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup17(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup17Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup17KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup17KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup17Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup17Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup17KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup17KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup17Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup18(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup18Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup18KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup18KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup18Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup18Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup18KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup18KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup18Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup19(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup19Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup19KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup19KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup19Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup19Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup19KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup19KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup19Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup2Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup2KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup2KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup2Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup2Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup2KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup2KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup2Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup20(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup20Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup20KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup20KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup20Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup20Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup20KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup20KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup20Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup21(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup21Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup21KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup21KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup21Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup21Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup21KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup21KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup21Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup27(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup27Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup27KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup27KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup27Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup27Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup27KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup27KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup27Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup28(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup28Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup28KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup28KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup28Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup28Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup28KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup28KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup28Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup29(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup29Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup29KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup29KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup29Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup29Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup29KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup29KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup29Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup30(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup30Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup30KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup30KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup30Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup30Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup30KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup30KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup30Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup31(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup31Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup31KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup31KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup31Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup31Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup31KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup31KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup31Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup32(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup32Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup32KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup32KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup32Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup32Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup32KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup32KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup32Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup5(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenSystemIkeDhGroup5Id(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup5KeypairCache(i["keypair-cache"], d, pre_append)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup5KeypairCount(i["keypair-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup5Mode(i["mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup5Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup5KeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup5KeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhGroup5Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhKeypairCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhKeypairCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhKeypairThrottle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhMultiprocess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeDhWorkerCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIkeEmbryonicLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemIke(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("dh_group_1", flattenSystemIkeDhGroup1(o["dh-group-1"], d, "dh_group_1")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-1"], "SystemIke-DhGroup1"); ok {
				if err = d.Set("dh_group_1", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_1: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_1"); ok {
			if err = d.Set("dh_group_1", flattenSystemIkeDhGroup1(o["dh-group-1"], d, "dh_group_1")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-1"], "SystemIke-DhGroup1"); ok {
					if err = d.Set("dh_group_1", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_1: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_14", flattenSystemIkeDhGroup14(o["dh-group-14"], d, "dh_group_14")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-14"], "SystemIke-DhGroup14"); ok {
				if err = d.Set("dh_group_14", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_14: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_14: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_14"); ok {
			if err = d.Set("dh_group_14", flattenSystemIkeDhGroup14(o["dh-group-14"], d, "dh_group_14")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-14"], "SystemIke-DhGroup14"); ok {
					if err = d.Set("dh_group_14", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_14: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_14: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_15", flattenSystemIkeDhGroup15(o["dh-group-15"], d, "dh_group_15")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-15"], "SystemIke-DhGroup15"); ok {
				if err = d.Set("dh_group_15", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_15: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_15: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_15"); ok {
			if err = d.Set("dh_group_15", flattenSystemIkeDhGroup15(o["dh-group-15"], d, "dh_group_15")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-15"], "SystemIke-DhGroup15"); ok {
					if err = d.Set("dh_group_15", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_15: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_15: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_16", flattenSystemIkeDhGroup16(o["dh-group-16"], d, "dh_group_16")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-16"], "SystemIke-DhGroup16"); ok {
				if err = d.Set("dh_group_16", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_16: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_16: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_16"); ok {
			if err = d.Set("dh_group_16", flattenSystemIkeDhGroup16(o["dh-group-16"], d, "dh_group_16")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-16"], "SystemIke-DhGroup16"); ok {
					if err = d.Set("dh_group_16", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_16: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_16: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_17", flattenSystemIkeDhGroup17(o["dh-group-17"], d, "dh_group_17")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-17"], "SystemIke-DhGroup17"); ok {
				if err = d.Set("dh_group_17", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_17: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_17: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_17"); ok {
			if err = d.Set("dh_group_17", flattenSystemIkeDhGroup17(o["dh-group-17"], d, "dh_group_17")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-17"], "SystemIke-DhGroup17"); ok {
					if err = d.Set("dh_group_17", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_17: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_17: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_18", flattenSystemIkeDhGroup18(o["dh-group-18"], d, "dh_group_18")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-18"], "SystemIke-DhGroup18"); ok {
				if err = d.Set("dh_group_18", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_18: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_18: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_18"); ok {
			if err = d.Set("dh_group_18", flattenSystemIkeDhGroup18(o["dh-group-18"], d, "dh_group_18")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-18"], "SystemIke-DhGroup18"); ok {
					if err = d.Set("dh_group_18", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_18: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_18: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_19", flattenSystemIkeDhGroup19(o["dh-group-19"], d, "dh_group_19")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-19"], "SystemIke-DhGroup19"); ok {
				if err = d.Set("dh_group_19", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_19: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_19: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_19"); ok {
			if err = d.Set("dh_group_19", flattenSystemIkeDhGroup19(o["dh-group-19"], d, "dh_group_19")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-19"], "SystemIke-DhGroup19"); ok {
					if err = d.Set("dh_group_19", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_19: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_19: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_2", flattenSystemIkeDhGroup2(o["dh-group-2"], d, "dh_group_2")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-2"], "SystemIke-DhGroup2"); ok {
				if err = d.Set("dh_group_2", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_2: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_2"); ok {
			if err = d.Set("dh_group_2", flattenSystemIkeDhGroup2(o["dh-group-2"], d, "dh_group_2")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-2"], "SystemIke-DhGroup2"); ok {
					if err = d.Set("dh_group_2", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_2: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_2: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_20", flattenSystemIkeDhGroup20(o["dh-group-20"], d, "dh_group_20")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-20"], "SystemIke-DhGroup20"); ok {
				if err = d.Set("dh_group_20", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_20: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_20: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_20"); ok {
			if err = d.Set("dh_group_20", flattenSystemIkeDhGroup20(o["dh-group-20"], d, "dh_group_20")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-20"], "SystemIke-DhGroup20"); ok {
					if err = d.Set("dh_group_20", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_20: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_20: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_21", flattenSystemIkeDhGroup21(o["dh-group-21"], d, "dh_group_21")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-21"], "SystemIke-DhGroup21"); ok {
				if err = d.Set("dh_group_21", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_21: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_21: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_21"); ok {
			if err = d.Set("dh_group_21", flattenSystemIkeDhGroup21(o["dh-group-21"], d, "dh_group_21")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-21"], "SystemIke-DhGroup21"); ok {
					if err = d.Set("dh_group_21", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_21: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_21: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_27", flattenSystemIkeDhGroup27(o["dh-group-27"], d, "dh_group_27")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-27"], "SystemIke-DhGroup27"); ok {
				if err = d.Set("dh_group_27", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_27: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_27: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_27"); ok {
			if err = d.Set("dh_group_27", flattenSystemIkeDhGroup27(o["dh-group-27"], d, "dh_group_27")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-27"], "SystemIke-DhGroup27"); ok {
					if err = d.Set("dh_group_27", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_27: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_27: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_28", flattenSystemIkeDhGroup28(o["dh-group-28"], d, "dh_group_28")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-28"], "SystemIke-DhGroup28"); ok {
				if err = d.Set("dh_group_28", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_28: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_28: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_28"); ok {
			if err = d.Set("dh_group_28", flattenSystemIkeDhGroup28(o["dh-group-28"], d, "dh_group_28")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-28"], "SystemIke-DhGroup28"); ok {
					if err = d.Set("dh_group_28", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_28: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_28: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_29", flattenSystemIkeDhGroup29(o["dh-group-29"], d, "dh_group_29")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-29"], "SystemIke-DhGroup29"); ok {
				if err = d.Set("dh_group_29", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_29: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_29: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_29"); ok {
			if err = d.Set("dh_group_29", flattenSystemIkeDhGroup29(o["dh-group-29"], d, "dh_group_29")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-29"], "SystemIke-DhGroup29"); ok {
					if err = d.Set("dh_group_29", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_29: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_29: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_30", flattenSystemIkeDhGroup30(o["dh-group-30"], d, "dh_group_30")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-30"], "SystemIke-DhGroup30"); ok {
				if err = d.Set("dh_group_30", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_30: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_30: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_30"); ok {
			if err = d.Set("dh_group_30", flattenSystemIkeDhGroup30(o["dh-group-30"], d, "dh_group_30")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-30"], "SystemIke-DhGroup30"); ok {
					if err = d.Set("dh_group_30", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_30: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_30: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_31", flattenSystemIkeDhGroup31(o["dh-group-31"], d, "dh_group_31")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-31"], "SystemIke-DhGroup31"); ok {
				if err = d.Set("dh_group_31", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_31: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_31: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_31"); ok {
			if err = d.Set("dh_group_31", flattenSystemIkeDhGroup31(o["dh-group-31"], d, "dh_group_31")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-31"], "SystemIke-DhGroup31"); ok {
					if err = d.Set("dh_group_31", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_31: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_31: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_32", flattenSystemIkeDhGroup32(o["dh-group-32"], d, "dh_group_32")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-32"], "SystemIke-DhGroup32"); ok {
				if err = d.Set("dh_group_32", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_32: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_32: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_32"); ok {
			if err = d.Set("dh_group_32", flattenSystemIkeDhGroup32(o["dh-group-32"], d, "dh_group_32")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-32"], "SystemIke-DhGroup32"); ok {
					if err = d.Set("dh_group_32", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_32: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_32: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dh_group_5", flattenSystemIkeDhGroup5(o["dh-group-5"], d, "dh_group_5")); err != nil {
			if vv, ok := fortiAPIPatch(o["dh-group-5"], "SystemIke-DhGroup5"); ok {
				if err = d.Set("dh_group_5", vv); err != nil {
					return fmt.Errorf("Error reading dh_group_5: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dh_group_5: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_5"); ok {
			if err = d.Set("dh_group_5", flattenSystemIkeDhGroup5(o["dh-group-5"], d, "dh_group_5")); err != nil {
				if vv, ok := fortiAPIPatch(o["dh-group-5"], "SystemIke-DhGroup5"); ok {
					if err = d.Set("dh_group_5", vv); err != nil {
						return fmt.Errorf("Error reading dh_group_5: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dh_group_5: %v", err)
				}
			}
		}
	}

	if err = d.Set("dh_keypair_cache", flattenSystemIkeDhKeypairCache(o["dh-keypair-cache"], d, "dh_keypair_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["dh-keypair-cache"], "SystemIke-DhKeypairCache"); ok {
			if err = d.Set("dh_keypair_cache", vv); err != nil {
				return fmt.Errorf("Error reading dh_keypair_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dh_keypair_cache: %v", err)
		}
	}

	if err = d.Set("dh_keypair_count", flattenSystemIkeDhKeypairCount(o["dh-keypair-count"], d, "dh_keypair_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["dh-keypair-count"], "SystemIke-DhKeypairCount"); ok {
			if err = d.Set("dh_keypair_count", vv); err != nil {
				return fmt.Errorf("Error reading dh_keypair_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dh_keypair_count: %v", err)
		}
	}

	if err = d.Set("dh_keypair_throttle", flattenSystemIkeDhKeypairThrottle(o["dh-keypair-throttle"], d, "dh_keypair_throttle")); err != nil {
		if vv, ok := fortiAPIPatch(o["dh-keypair-throttle"], "SystemIke-DhKeypairThrottle"); ok {
			if err = d.Set("dh_keypair_throttle", vv); err != nil {
				return fmt.Errorf("Error reading dh_keypair_throttle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dh_keypair_throttle: %v", err)
		}
	}

	if err = d.Set("dh_mode", flattenSystemIkeDhMode(o["dh-mode"], d, "dh_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["dh-mode"], "SystemIke-DhMode"); ok {
			if err = d.Set("dh_mode", vv); err != nil {
				return fmt.Errorf("Error reading dh_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dh_mode: %v", err)
		}
	}

	if err = d.Set("dh_multiprocess", flattenSystemIkeDhMultiprocess(o["dh-multiprocess"], d, "dh_multiprocess")); err != nil {
		if vv, ok := fortiAPIPatch(o["dh-multiprocess"], "SystemIke-DhMultiprocess"); ok {
			if err = d.Set("dh_multiprocess", vv); err != nil {
				return fmt.Errorf("Error reading dh_multiprocess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dh_multiprocess: %v", err)
		}
	}

	if err = d.Set("dh_worker_count", flattenSystemIkeDhWorkerCount(o["dh-worker-count"], d, "dh_worker_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["dh-worker-count"], "SystemIke-DhWorkerCount"); ok {
			if err = d.Set("dh_worker_count", vv); err != nil {
				return fmt.Errorf("Error reading dh_worker_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dh_worker_count: %v", err)
		}
	}

	if err = d.Set("embryonic_limit", flattenSystemIkeEmbryonicLimit(o["embryonic-limit"], d, "embryonic_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["embryonic-limit"], "SystemIke-EmbryonicLimit"); ok {
			if err = d.Set("embryonic_limit", vv); err != nil {
				return fmt.Errorf("Error reading embryonic_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading embryonic_limit: %v", err)
		}
	}

	return nil
}

func flattenSystemIkeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemIkeDhGroup1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup1Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup1KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup1KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup1Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup1Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup1KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup1KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup1Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup14(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup14Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup14KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup14KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup14Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup14Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup14KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup14KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup14Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup15(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup15Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup15KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup15KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup15Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup15Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup15KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup15KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup15Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup16(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup16Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup16KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup16KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup16Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup16Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup16KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup16KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup16Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup17(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup17Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup17KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup17KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup17Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup17Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup17KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup17KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup17Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup18(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup18Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup18KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup18KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup18Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup18Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup18KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup18KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup18Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup19(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup19Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup19KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup19KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup19Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup19Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup19KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup19KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup19Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup2Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup2KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup2KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup2Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup2Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup2KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup2KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup2Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup20(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup20Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup20KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup20KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup20Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup20Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup20KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup20KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup20Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup21(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup21Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup21KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup21KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup21Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup21Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup21KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup21KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup21Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup27(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup27Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup27KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup27KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup27Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup27Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup27KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup27KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup27Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup28(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup28Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup28KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup28KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup28Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup28Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup28KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup28KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup28Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup29(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup29Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup29KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup29KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup29Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup29Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup29KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup29KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup29Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup30(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup30Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup30KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup30KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup30Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup30Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup30KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup30KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup30Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup31(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup31Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup31KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup31KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup31Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup31Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup31KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup31KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup31Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup32(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup32Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup32KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup32KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup32Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup32Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup32KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup32KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup32Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandSystemIkeDhGroup5Id(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-cache"], _ = expandSystemIkeDhGroup5KeypairCache(d, i["keypair_cache"], pre_append)
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keypair-count"], _ = expandSystemIkeDhGroup5KeypairCount(d, i["keypair_count"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandSystemIkeDhGroup5Mode(d, i["mode"], pre_append)
	}

	return result, nil
}

func expandSystemIkeDhGroup5Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup5KeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup5KeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup5Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhKeypairCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhKeypairCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhKeypairThrottle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhMultiprocess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhWorkerCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeEmbryonicLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIke(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dh_group_1"); ok || d.HasChange("dh_group_1") {
		t, err := expandSystemIkeDhGroup1(d, v, "dh_group_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-1"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_14"); ok || d.HasChange("dh_group_14") {
		t, err := expandSystemIkeDhGroup14(d, v, "dh_group_14")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-14"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_15"); ok || d.HasChange("dh_group_15") {
		t, err := expandSystemIkeDhGroup15(d, v, "dh_group_15")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-15"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_16"); ok || d.HasChange("dh_group_16") {
		t, err := expandSystemIkeDhGroup16(d, v, "dh_group_16")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-16"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_17"); ok || d.HasChange("dh_group_17") {
		t, err := expandSystemIkeDhGroup17(d, v, "dh_group_17")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-17"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_18"); ok || d.HasChange("dh_group_18") {
		t, err := expandSystemIkeDhGroup18(d, v, "dh_group_18")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-18"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_19"); ok || d.HasChange("dh_group_19") {
		t, err := expandSystemIkeDhGroup19(d, v, "dh_group_19")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-19"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_2"); ok || d.HasChange("dh_group_2") {
		t, err := expandSystemIkeDhGroup2(d, v, "dh_group_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-2"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_20"); ok || d.HasChange("dh_group_20") {
		t, err := expandSystemIkeDhGroup20(d, v, "dh_group_20")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-20"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_21"); ok || d.HasChange("dh_group_21") {
		t, err := expandSystemIkeDhGroup21(d, v, "dh_group_21")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-21"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_27"); ok || d.HasChange("dh_group_27") {
		t, err := expandSystemIkeDhGroup27(d, v, "dh_group_27")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-27"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_28"); ok || d.HasChange("dh_group_28") {
		t, err := expandSystemIkeDhGroup28(d, v, "dh_group_28")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-28"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_29"); ok || d.HasChange("dh_group_29") {
		t, err := expandSystemIkeDhGroup29(d, v, "dh_group_29")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-29"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_30"); ok || d.HasChange("dh_group_30") {
		t, err := expandSystemIkeDhGroup30(d, v, "dh_group_30")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-30"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_31"); ok || d.HasChange("dh_group_31") {
		t, err := expandSystemIkeDhGroup31(d, v, "dh_group_31")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-31"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_32"); ok || d.HasChange("dh_group_32") {
		t, err := expandSystemIkeDhGroup32(d, v, "dh_group_32")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-32"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_5"); ok || d.HasChange("dh_group_5") {
		t, err := expandSystemIkeDhGroup5(d, v, "dh_group_5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-5"] = t
		}
	}

	if v, ok := d.GetOk("dh_keypair_cache"); ok || d.HasChange("dh_keypair_cache") {
		t, err := expandSystemIkeDhKeypairCache(d, v, "dh_keypair_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-keypair-cache"] = t
		}
	}

	if v, ok := d.GetOk("dh_keypair_count"); ok || d.HasChange("dh_keypair_count") {
		t, err := expandSystemIkeDhKeypairCount(d, v, "dh_keypair_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-keypair-count"] = t
		}
	}

	if v, ok := d.GetOk("dh_keypair_throttle"); ok || d.HasChange("dh_keypair_throttle") {
		t, err := expandSystemIkeDhKeypairThrottle(d, v, "dh_keypair_throttle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-keypair-throttle"] = t
		}
	}

	if v, ok := d.GetOk("dh_mode"); ok || d.HasChange("dh_mode") {
		t, err := expandSystemIkeDhMode(d, v, "dh_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-mode"] = t
		}
	}

	if v, ok := d.GetOk("dh_multiprocess"); ok || d.HasChange("dh_multiprocess") {
		t, err := expandSystemIkeDhMultiprocess(d, v, "dh_multiprocess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-multiprocess"] = t
		}
	}

	if v, ok := d.GetOk("dh_worker_count"); ok || d.HasChange("dh_worker_count") {
		t, err := expandSystemIkeDhWorkerCount(d, v, "dh_worker_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-worker-count"] = t
		}
	}

	if v, ok := d.GetOk("embryonic_limit"); ok || d.HasChange("embryonic_limit") {
		t, err := expandSystemIkeEmbryonicLimit(d, v, "embryonic_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["embryonic-limit"] = t
		}
	}

	return &obj, nil
}
