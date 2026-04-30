// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaWebPortalBookmarkLlmSecureProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebPortalBookmarkLlmSecureProxyUpdate,
		Read:   resourceZtnaWebPortalBookmarkLlmSecureProxyRead,
		Update: resourceZtnaWebPortalBookmarkLlmSecureProxyUpdate,
		Delete: resourceZtnaWebPortalBookmarkLlmSecureProxyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"web_portal_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"all_llm_servers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"llm_servers": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceZtnaWebPortalBookmarkLlmSecureProxyUpdate(d *schema.ResourceData, m interface{}) error {
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
	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_portal_bookmark"] = web_portal_bookmark

	obj, err := getObjectZtnaWebPortalBookmarkLlmSecureProxy(d, false)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortalBookmarkLlmSecureProxy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateZtnaWebPortalBookmarkLlmSecureProxy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortalBookmarkLlmSecureProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ZtnaWebPortalBookmarkLlmSecureProxy")

	return resourceZtnaWebPortalBookmarkLlmSecureProxyRead(d, m)
}

func resourceZtnaWebPortalBookmarkLlmSecureProxyDelete(d *schema.ResourceData, m interface{}) error {
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
	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_portal_bookmark"] = web_portal_bookmark

	obj, err := getObjectZtnaWebPortalBookmarkLlmSecureProxy(d, true)

	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortalBookmarkLlmSecureProxy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateZtnaWebPortalBookmarkLlmSecureProxy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing ZtnaWebPortalBookmarkLlmSecureProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebPortalBookmarkLlmSecureProxyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
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
	if web_portal_bookmark == "" {
		web_portal_bookmark = importOptionChecking(m.(*FortiClient).Cfg, "web_portal_bookmark")
		if web_portal_bookmark == "" {
			return fmt.Errorf("Parameter web_portal_bookmark is missing")
		}
		if err = d.Set("web_portal_bookmark", web_portal_bookmark); err != nil {
			return fmt.Errorf("Error set params web_portal_bookmark: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_portal_bookmark"] = web_portal_bookmark

	o, err := c.ReadZtnaWebPortalBookmarkLlmSecureProxy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ZtnaWebPortalBookmarkLlmSecureProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebPortalBookmarkLlmSecureProxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebPortalBookmarkLlmSecureProxy resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebPortalBookmarkLlmSecureProxyAllLlmServers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkLlmSecureProxyLlmServers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectZtnaWebPortalBookmarkLlmSecureProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("all_llm_servers", flattenZtnaWebPortalBookmarkLlmSecureProxyAllLlmServers2edl(o["all-llm-servers"], d, "all_llm_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["all-llm-servers"], "ZtnaWebPortalBookmarkLlmSecureProxy-AllLlmServers"); ok {
			if err = d.Set("all_llm_servers", vv); err != nil {
				return fmt.Errorf("Error reading all_llm_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading all_llm_servers: %v", err)
		}
	}

	if err = d.Set("llm_servers", flattenZtnaWebPortalBookmarkLlmSecureProxyLlmServers2edl(o["llm-servers"], d, "llm_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["llm-servers"], "ZtnaWebPortalBookmarkLlmSecureProxy-LlmServers"); ok {
			if err = d.Set("llm_servers", vv); err != nil {
				return fmt.Errorf("Error reading llm_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading llm_servers: %v", err)
		}
	}

	return nil
}

func flattenZtnaWebPortalBookmarkLlmSecureProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaWebPortalBookmarkLlmSecureProxyAllLlmServers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkLlmSecureProxyLlmServers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectZtnaWebPortalBookmarkLlmSecureProxy(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("all_llm_servers"); ok || d.HasChange("all_llm_servers") {
		t, err := expandZtnaWebPortalBookmarkLlmSecureProxyAllLlmServers2edl(d, v, "all_llm_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["all-llm-servers"] = t
		}
	}

	if v, ok := d.GetOk("llm_servers"); ok || d.HasChange("llm_servers") {
		t, err := expandZtnaWebPortalBookmarkLlmSecureProxyLlmServers2edl(d, v, "llm_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["llm-servers"] = t
		}
	}

	return &obj, nil
}
