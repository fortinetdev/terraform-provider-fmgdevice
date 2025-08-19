// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure online sign up (OSU) provider list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20H2QpOsuProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20H2QpOsuProviderCreate,
		Read:   resourceWirelessControllerHotspot20H2QpOsuProviderRead,
		Update: resourceWirelessControllerHotspot20H2QpOsuProviderUpdate,
		Delete: resourceWirelessControllerHotspot20H2QpOsuProviderDelete,

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
			"friendly_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"friendly_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lang": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"icon": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"osu_method": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"osu_nai": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_uri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_description": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lang": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"service_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerHotspot20H2QpOsuProviderCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20H2QpOsuProvider(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpOsuProvider resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20H2QpOsuProvider(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpOsuProvider resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20H2QpOsuProviderRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpOsuProviderUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20H2QpOsuProvider(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpOsuProvider resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20H2QpOsuProvider(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpOsuProvider resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20H2QpOsuProviderRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpOsuProviderDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerHotspot20H2QpOsuProvider(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20H2QpOsuProvider resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2QpOsuProviderRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerHotspot20H2QpOsuProvider(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpOsuProvider resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20H2QpOsuProvider(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpOsuProvider resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20H2QpOsuProviderFriendlyName(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "friendly_name"
		if _, ok := i["friendly-name"]; ok {
			v := flattenWirelessControllerHotspot20H2QpOsuProviderFriendlyNameFriendlyName(i["friendly-name"], d, pre_append)
			tmp["friendly_name"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20H2QpOsuProvider-FriendlyName-FriendlyName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenWirelessControllerHotspot20H2QpOsuProviderFriendlyNameIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20H2QpOsuProvider-FriendlyName-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := i["lang"]; ok {
			v := flattenWirelessControllerHotspot20H2QpOsuProviderFriendlyNameLang(i["lang"], d, pre_append)
			tmp["lang"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20H2QpOsuProvider-FriendlyName-Lang")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20H2QpOsuProviderFriendlyNameFriendlyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderFriendlyNameIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderFriendlyNameLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderIcon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20H2QpOsuProviderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderOsuMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20H2QpOsuProviderOsuNai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderServerUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescription(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := i["lang"]; ok {
			v := flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionLang(i["lang"], d, pre_append)
			tmp["lang"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20H2QpOsuProvider-ServiceDescription-Lang")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_description"
		if _, ok := i["service-description"]; ok {
			v := flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceDescription(i["service-description"], d, pre_append)
			tmp["service_description"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20H2QpOsuProvider-ServiceDescription-ServiceDescription")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := i["service-id"]; ok {
			v := flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceId(i["service-id"], d, pre_append)
			tmp["service_id"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20H2QpOsuProvider-ServiceDescription-ServiceId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20H2QpOsuProvider(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("friendly_name", flattenWirelessControllerHotspot20H2QpOsuProviderFriendlyName(o["friendly-name"], d, "friendly_name")); err != nil {
			if vv, ok := fortiAPIPatch(o["friendly-name"], "WirelessControllerHotspot20H2QpOsuProvider-FriendlyName"); ok {
				if err = d.Set("friendly_name", vv); err != nil {
					return fmt.Errorf("Error reading friendly_name: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading friendly_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("friendly_name"); ok {
			if err = d.Set("friendly_name", flattenWirelessControllerHotspot20H2QpOsuProviderFriendlyName(o["friendly-name"], d, "friendly_name")); err != nil {
				if vv, ok := fortiAPIPatch(o["friendly-name"], "WirelessControllerHotspot20H2QpOsuProvider-FriendlyName"); ok {
					if err = d.Set("friendly_name", vv); err != nil {
						return fmt.Errorf("Error reading friendly_name: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading friendly_name: %v", err)
				}
			}
		}
	}

	if err = d.Set("icon", flattenWirelessControllerHotspot20H2QpOsuProviderIcon(o["icon"], d, "icon")); err != nil {
		if vv, ok := fortiAPIPatch(o["icon"], "WirelessControllerHotspot20H2QpOsuProvider-Icon"); ok {
			if err = d.Set("icon", vv); err != nil {
				return fmt.Errorf("Error reading icon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icon: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerHotspot20H2QpOsuProviderName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerHotspot20H2QpOsuProvider-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("osu_method", flattenWirelessControllerHotspot20H2QpOsuProviderOsuMethod(o["osu-method"], d, "osu_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["osu-method"], "WirelessControllerHotspot20H2QpOsuProvider-OsuMethod"); ok {
			if err = d.Set("osu_method", vv); err != nil {
				return fmt.Errorf("Error reading osu_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading osu_method: %v", err)
		}
	}

	if err = d.Set("osu_nai", flattenWirelessControllerHotspot20H2QpOsuProviderOsuNai(o["osu-nai"], d, "osu_nai")); err != nil {
		if vv, ok := fortiAPIPatch(o["osu-nai"], "WirelessControllerHotspot20H2QpOsuProvider-OsuNai"); ok {
			if err = d.Set("osu_nai", vv); err != nil {
				return fmt.Errorf("Error reading osu_nai: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading osu_nai: %v", err)
		}
	}

	if err = d.Set("server_uri", flattenWirelessControllerHotspot20H2QpOsuProviderServerUri(o["server-uri"], d, "server_uri")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-uri"], "WirelessControllerHotspot20H2QpOsuProvider-ServerUri"); ok {
			if err = d.Set("server_uri", vv); err != nil {
				return fmt.Errorf("Error reading server_uri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_uri: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service_description", flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescription(o["service-description"], d, "service_description")); err != nil {
			if vv, ok := fortiAPIPatch(o["service-description"], "WirelessControllerHotspot20H2QpOsuProvider-ServiceDescription"); ok {
				if err = d.Set("service_description", vv); err != nil {
					return fmt.Errorf("Error reading service_description: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading service_description: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service_description"); ok {
			if err = d.Set("service_description", flattenWirelessControllerHotspot20H2QpOsuProviderServiceDescription(o["service-description"], d, "service_description")); err != nil {
				if vv, ok := fortiAPIPatch(o["service-description"], "WirelessControllerHotspot20H2QpOsuProvider-ServiceDescription"); ok {
					if err = d.Set("service_description", vv); err != nil {
						return fmt.Errorf("Error reading service_description: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading service_description: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20H2QpOsuProviderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20H2QpOsuProviderFriendlyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "friendly_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["friendly-name"], _ = expandWirelessControllerHotspot20H2QpOsuProviderFriendlyNameFriendlyName(d, i["friendly_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandWirelessControllerHotspot20H2QpOsuProviderFriendlyNameIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lang"], _ = expandWirelessControllerHotspot20H2QpOsuProviderFriendlyNameLang(d, i["lang"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderFriendlyNameFriendlyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderFriendlyNameIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderFriendlyNameLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderIcon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderOsuMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderOsuNai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderServerUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderServiceDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lang"], _ = expandWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionLang(d, i["lang"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-description"], _ = expandWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceDescription(d, i["service_description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-id"], _ = expandWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceId(d, i["service_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20H2QpOsuProvider(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("friendly_name"); ok || d.HasChange("friendly_name") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderFriendlyName(d, v, "friendly_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["friendly-name"] = t
		}
	}

	if v, ok := d.GetOk("icon"); ok || d.HasChange("icon") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderIcon(d, v, "icon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icon"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("osu_method"); ok || d.HasChange("osu_method") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderOsuMethod(d, v, "osu_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osu-method"] = t
		}
	}

	if v, ok := d.GetOk("osu_nai"); ok || d.HasChange("osu_nai") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderOsuNai(d, v, "osu_nai")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osu-nai"] = t
		}
	}

	if v, ok := d.GetOk("server_uri"); ok || d.HasChange("server_uri") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderServerUri(d, v, "server_uri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-uri"] = t
		}
	}

	if v, ok := d.GetOk("service_description"); ok || d.HasChange("service_description") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderServiceDescription(d, v, "service_description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-description"] = t
		}
	}

	return &obj, nil
}
