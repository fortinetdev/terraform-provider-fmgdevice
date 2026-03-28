// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure VideoFilter profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVideofilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceVideofilterProfileCreate,
		Read:   resourceVideofilterProfileRead,
		Update: resourceVideofilterProfileUpdate,
		Delete: resourceVideofilterProfileDelete,

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dailymotion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"channel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"keyword": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
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
				},
			},
			"fortiguard_category": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"category_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vimeo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vimeo_restrict": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"youtube": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"youtube_channel_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"youtube_restrict": &schema.Schema{
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

func resourceVideofilterProfileCreate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectVideofilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating VideofilterProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVideofilterProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVideofilterProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VideofilterProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateVideofilterProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VideofilterProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceVideofilterProfileRead(d, m)
}

func resourceVideofilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectVideofilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating VideofilterProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVideofilterProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VideofilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVideofilterProfileRead(d, m)
}

func resourceVideofilterProfileDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteVideofilterProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VideofilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVideofilterProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVideofilterProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VideofilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVideofilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VideofilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenVideofilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileDailymotion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFilters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenVideofilterProfileFiltersAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "VideofilterProfile-Filters-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenVideofilterProfileFiltersCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "VideofilterProfile-Filters-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel"
		if _, ok := i["channel"]; ok {
			v := flattenVideofilterProfileFiltersChannel(i["channel"], d, pre_append)
			tmp["channel"] = fortiAPISubPartPatch(v, "VideofilterProfile-Filters-Channel")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenVideofilterProfileFiltersComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "VideofilterProfile-Filters-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenVideofilterProfileFiltersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "VideofilterProfile-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyword"
		if _, ok := i["keyword"]; ok {
			v := flattenVideofilterProfileFiltersKeyword(i["keyword"], d, pre_append)
			tmp["keyword"] = fortiAPISubPartPatch(v, "VideofilterProfile-Filters-Keyword")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenVideofilterProfileFiltersLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "VideofilterProfile-Filters-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenVideofilterProfileFiltersType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "VideofilterProfile-Filters-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVideofilterProfileFiltersAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFiltersCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFiltersChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFiltersComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFiltersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFiltersKeyword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVideofilterProfileFiltersLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFiltersType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFortiguardCategory(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenVideofilterProfileFortiguardCategoryFilters(i["filters"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenVideofilterProfileFortiguardCategoryFilters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenVideofilterProfileFortiguardCategoryFiltersAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "VideofilterProfileFortiguardCategory-Filters-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category_id"
		if _, ok := i["category-id"]; ok {
			v := flattenVideofilterProfileFortiguardCategoryFiltersCategoryId(i["category-id"], d, pre_append)
			tmp["category_id"] = fortiAPISubPartPatch(v, "VideofilterProfileFortiguardCategory-Filters-CategoryId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenVideofilterProfileFortiguardCategoryFiltersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "VideofilterProfileFortiguardCategory-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenVideofilterProfileFortiguardCategoryFiltersLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "VideofilterProfileFortiguardCategory-Filters-Log")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVideofilterProfileFortiguardCategoryFiltersAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFortiguardCategoryFiltersCategoryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFortiguardCategoryFiltersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFortiguardCategoryFiltersLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVideofilterProfileVimeo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileVimeoRestrict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileYoutube(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileYoutubeChannelFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVideofilterProfileYoutubeRestrict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVideofilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenVideofilterProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "VideofilterProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("dailymotion", flattenVideofilterProfileDailymotion(o["dailymotion"], d, "dailymotion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dailymotion"], "VideofilterProfile-Dailymotion"); ok {
			if err = d.Set("dailymotion", vv); err != nil {
				return fmt.Errorf("Error reading dailymotion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dailymotion: %v", err)
		}
	}

	if err = d.Set("default_action", flattenVideofilterProfileDefaultAction(o["default-action"], d, "default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-action"], "VideofilterProfile-DefaultAction"); ok {
			if err = d.Set("default_action", vv); err != nil {
				return fmt.Errorf("Error reading default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("filters", flattenVideofilterProfileFilters(o["filters"], d, "filters")); err != nil {
			if vv, ok := fortiAPIPatch(o["filters"], "VideofilterProfile-Filters"); ok {
				if err = d.Set("filters", vv); err != nil {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filters"); ok {
			if err = d.Set("filters", flattenVideofilterProfileFilters(o["filters"], d, "filters")); err != nil {
				if vv, ok := fortiAPIPatch(o["filters"], "VideofilterProfile-Filters"); ok {
					if err = d.Set("filters", vv); err != nil {
						return fmt.Errorf("Error reading filters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("fortiguard_category", flattenVideofilterProfileFortiguardCategory(o["fortiguard-category"], d, "fortiguard_category")); err != nil {
			if vv, ok := fortiAPIPatch(o["fortiguard-category"], "VideofilterProfile-FortiguardCategory"); ok {
				if err = d.Set("fortiguard_category", vv); err != nil {
					return fmt.Errorf("Error reading fortiguard_category: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fortiguard_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fortiguard_category"); ok {
			if err = d.Set("fortiguard_category", flattenVideofilterProfileFortiguardCategory(o["fortiguard-category"], d, "fortiguard_category")); err != nil {
				if vv, ok := fortiAPIPatch(o["fortiguard-category"], "VideofilterProfile-FortiguardCategory"); ok {
					if err = d.Set("fortiguard_category", vv); err != nil {
						return fmt.Errorf("Error reading fortiguard_category: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fortiguard_category: %v", err)
				}
			}
		}
	}

	if err = d.Set("log", flattenVideofilterProfileLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "VideofilterProfile-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenVideofilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VideofilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenVideofilterProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "VideofilterProfile-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("vimeo", flattenVideofilterProfileVimeo(o["vimeo"], d, "vimeo")); err != nil {
		if vv, ok := fortiAPIPatch(o["vimeo"], "VideofilterProfile-Vimeo"); ok {
			if err = d.Set("vimeo", vv); err != nil {
				return fmt.Errorf("Error reading vimeo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vimeo: %v", err)
		}
	}

	if err = d.Set("vimeo_restrict", flattenVideofilterProfileVimeoRestrict(o["vimeo-restrict"], d, "vimeo_restrict")); err != nil {
		if vv, ok := fortiAPIPatch(o["vimeo-restrict"], "VideofilterProfile-VimeoRestrict"); ok {
			if err = d.Set("vimeo_restrict", vv); err != nil {
				return fmt.Errorf("Error reading vimeo_restrict: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vimeo_restrict: %v", err)
		}
	}

	if err = d.Set("youtube", flattenVideofilterProfileYoutube(o["youtube"], d, "youtube")); err != nil {
		if vv, ok := fortiAPIPatch(o["youtube"], "VideofilterProfile-Youtube"); ok {
			if err = d.Set("youtube", vv); err != nil {
				return fmt.Errorf("Error reading youtube: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading youtube: %v", err)
		}
	}

	if err = d.Set("youtube_channel_filter", flattenVideofilterProfileYoutubeChannelFilter(o["youtube-channel-filter"], d, "youtube_channel_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["youtube-channel-filter"], "VideofilterProfile-YoutubeChannelFilter"); ok {
			if err = d.Set("youtube_channel_filter", vv); err != nil {
				return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
		}
	}

	if err = d.Set("youtube_restrict", flattenVideofilterProfileYoutubeRestrict(o["youtube-restrict"], d, "youtube_restrict")); err != nil {
		if vv, ok := fortiAPIPatch(o["youtube-restrict"], "VideofilterProfile-YoutubeRestrict"); ok {
			if err = d.Set("youtube_restrict", vv); err != nil {
				return fmt.Errorf("Error reading youtube_restrict: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading youtube_restrict: %v", err)
		}
	}

	return nil
}

func flattenVideofilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVideofilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileDailymotion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFilters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandVideofilterProfileFiltersAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandVideofilterProfileFiltersCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["channel"], _ = expandVideofilterProfileFiltersChannel(d, i["channel"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandVideofilterProfileFiltersComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandVideofilterProfileFiltersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyword"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keyword"], _ = expandVideofilterProfileFiltersKeyword(d, i["keyword"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandVideofilterProfileFiltersLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandVideofilterProfileFiltersType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVideofilterProfileFiltersAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersKeyword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVideofilterProfileFiltersLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFortiguardCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandVideofilterProfileFortiguardCategoryFilters(d, i["filters"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["filters"] = t
		}
	}

	return result, nil
}

func expandVideofilterProfileFortiguardCategoryFilters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandVideofilterProfileFortiguardCategoryFiltersAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category-id"], _ = expandVideofilterProfileFortiguardCategoryFiltersCategoryId(d, i["category_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandVideofilterProfileFortiguardCategoryFiltersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandVideofilterProfileFortiguardCategoryFiltersLog(d, i["log"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVideofilterProfileFortiguardCategoryFiltersAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFortiguardCategoryFiltersCategoryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFortiguardCategoryFiltersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFortiguardCategoryFiltersLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVideofilterProfileVimeo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileVimeoRestrict(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileYoutube(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileYoutubeChannelFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVideofilterProfileYoutubeRestrict(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVideofilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandVideofilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dailymotion"); ok || d.HasChange("dailymotion") {
		t, err := expandVideofilterProfileDailymotion(d, v, "dailymotion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dailymotion"] = t
		}
	}

	if v, ok := d.GetOk("default_action"); ok || d.HasChange("default_action") {
		t, err := expandVideofilterProfileDefaultAction(d, v, "default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-action"] = t
		}
	}

	if v, ok := d.GetOk("filters"); ok || d.HasChange("filters") {
		t, err := expandVideofilterProfileFilters(d, v, "filters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filters"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_category"); ok || d.HasChange("fortiguard_category") {
		t, err := expandVideofilterProfileFortiguardCategory(d, v, "fortiguard_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-category"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandVideofilterProfileLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVideofilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandVideofilterProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("vimeo"); ok || d.HasChange("vimeo") {
		t, err := expandVideofilterProfileVimeo(d, v, "vimeo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vimeo"] = t
		}
	}

	if v, ok := d.GetOk("vimeo_restrict"); ok || d.HasChange("vimeo_restrict") {
		t, err := expandVideofilterProfileVimeoRestrict(d, v, "vimeo_restrict")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vimeo-restrict"] = t
		}
	}

	if v, ok := d.GetOk("youtube"); ok || d.HasChange("youtube") {
		t, err := expandVideofilterProfileYoutube(d, v, "youtube")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube"] = t
		}
	}

	if v, ok := d.GetOk("youtube_channel_filter"); ok || d.HasChange("youtube_channel_filter") {
		t, err := expandVideofilterProfileYoutubeChannelFilter(d, v, "youtube_channel_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-channel-filter"] = t
		}
	}

	if v, ok := d.GetOk("youtube_restrict"); ok || d.HasChange("youtube_restrict") {
		t, err := expandVideofilterProfileYoutubeRestrict(d, v, "youtube_restrict")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-restrict"] = t
		}
	}

	return &obj, nil
}
