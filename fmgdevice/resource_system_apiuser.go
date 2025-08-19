// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure API users.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemApiUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemApiUserCreate,
		Read:   resourceSystemApiUserRead,
		Update: resourceSystemApiUserUpdate,
		Delete: resourceSystemApiUserDelete,

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
			"accprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"api_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cors_allow_origin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"peer_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trusthost": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ipv4_trusthost": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ipv6_trusthost": &schema.Schema{
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
			"vdom": &schema.Schema{
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

func resourceSystemApiUserCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemApiUser(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemApiUser resource while getting object: %v", err)
	}

	_, err = c.CreateSystemApiUser(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemApiUser resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemApiUserRead(d, m)
}

func resourceSystemApiUserUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemApiUser(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemApiUser resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemApiUser(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemApiUser resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemApiUserRead(d, m)
}

func resourceSystemApiUserDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemApiUser(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemApiUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemApiUserRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemApiUser(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemApiUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemApiUser(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemApiUser resource from API: %v", err)
	}
	return nil
}

func flattenSystemApiUserAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemApiUserComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserCorsAllowOrigin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserPeerAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserPeerGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemApiUserSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserTrusthost(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemApiUserTrusthostId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemApiUser-Trusthost-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_trusthost"
		if _, ok := i["ipv4-trusthost"]; ok {
			v := flattenSystemApiUserTrusthostIpv4Trusthost(i["ipv4-trusthost"], d, pre_append)
			tmp["ipv4_trusthost"] = fortiAPISubPartPatch(v, "SystemApiUser-Trusthost-Ipv4Trusthost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_trusthost"
		if _, ok := i["ipv6-trusthost"]; ok {
			v := flattenSystemApiUserTrusthostIpv6Trusthost(i["ipv6-trusthost"], d, pre_append)
			tmp["ipv6_trusthost"] = fortiAPISubPartPatch(v, "SystemApiUser-Trusthost-Ipv6Trusthost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSystemApiUserTrusthostType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SystemApiUser-Trusthost-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemApiUserTrusthostId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserTrusthostIpv4Trusthost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemApiUserTrusthostIpv6Trusthost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserTrusthostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemApiUser(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("accprofile", flattenSystemApiUserAccprofile(o["accprofile"], d, "accprofile")); err != nil {
		if vv, ok := fortiAPIPatch(o["accprofile"], "SystemApiUser-Accprofile"); ok {
			if err = d.Set("accprofile", vv); err != nil {
				return fmt.Errorf("Error reading accprofile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemApiUserComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "SystemApiUser-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("cors_allow_origin", flattenSystemApiUserCorsAllowOrigin(o["cors-allow-origin"], d, "cors_allow_origin")); err != nil {
		if vv, ok := fortiAPIPatch(o["cors-allow-origin"], "SystemApiUser-CorsAllowOrigin"); ok {
			if err = d.Set("cors_allow_origin", vv); err != nil {
				return fmt.Errorf("Error reading cors_allow_origin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cors_allow_origin: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemApiUserName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemApiUser-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("peer_auth", flattenSystemApiUserPeerAuth(o["peer-auth"], d, "peer_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-auth"], "SystemApiUser-PeerAuth"); ok {
			if err = d.Set("peer_auth", vv); err != nil {
				return fmt.Errorf("Error reading peer_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_auth: %v", err)
		}
	}

	if err = d.Set("peer_group", flattenSystemApiUserPeerGroup(o["peer-group"], d, "peer_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-group"], "SystemApiUser-PeerGroup"); ok {
			if err = d.Set("peer_group", vv); err != nil {
				return fmt.Errorf("Error reading peer_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_group: %v", err)
		}
	}

	if err = d.Set("schedule", flattenSystemApiUserSchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "SystemApiUser-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("trusthost", flattenSystemApiUserTrusthost(o["trusthost"], d, "trusthost")); err != nil {
			if vv, ok := fortiAPIPatch(o["trusthost"], "SystemApiUser-Trusthost"); ok {
				if err = d.Set("trusthost", vv); err != nil {
					return fmt.Errorf("Error reading trusthost: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading trusthost: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("trusthost"); ok {
			if err = d.Set("trusthost", flattenSystemApiUserTrusthost(o["trusthost"], d, "trusthost")); err != nil {
				if vv, ok := fortiAPIPatch(o["trusthost"], "SystemApiUser-Trusthost"); ok {
					if err = d.Set("trusthost", vv); err != nil {
						return fmt.Errorf("Error reading trusthost: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading trusthost: %v", err)
				}
			}
		}
	}

	if err = d.Set("vdom", flattenSystemApiUserVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemApiUser-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenSystemApiUserFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemApiUserAccprofile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemApiUserApiKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemApiUserComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserCorsAllowOrigin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserPeerAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserPeerGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemApiUserSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserTrusthost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemApiUserTrusthostId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_trusthost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv4-trusthost"], _ = expandSystemApiUserTrusthostIpv4Trusthost(d, i["ipv4_trusthost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_trusthost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-trusthost"], _ = expandSystemApiUserTrusthostIpv6Trusthost(d, i["ipv6_trusthost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSystemApiUserTrusthostType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemApiUserTrusthostId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserTrusthostIpv4Trusthost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemApiUserTrusthostIpv6Trusthost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserTrusthostType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemApiUser(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accprofile"); ok || d.HasChange("accprofile") {
		t, err := expandSystemApiUserAccprofile(d, v, "accprofile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	}

	if v, ok := d.GetOk("api_key"); ok || d.HasChange("api_key") {
		t, err := expandSystemApiUserApiKey(d, v, "api_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-key"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandSystemApiUserComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("cors_allow_origin"); ok || d.HasChange("cors_allow_origin") {
		t, err := expandSystemApiUserCorsAllowOrigin(d, v, "cors_allow_origin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cors-allow-origin"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemApiUserName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("peer_auth"); ok || d.HasChange("peer_auth") {
		t, err := expandSystemApiUserPeerAuth(d, v, "peer_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-auth"] = t
		}
	}

	if v, ok := d.GetOk("peer_group"); ok || d.HasChange("peer_group") {
		t, err := expandSystemApiUserPeerGroup(d, v, "peer_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-group"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandSystemApiUserSchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("trusthost"); ok || d.HasChange("trusthost") {
		t, err := expandSystemApiUserTrusthost(d, v, "trusthost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemApiUserVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
