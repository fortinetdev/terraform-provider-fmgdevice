// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure web filter search engines.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterSearchEngine() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterSearchEngineCreate,
		Read:   resourceWebfilterSearchEngineRead,
		Update: resourceWebfilterSearchEngineUpdate,
		Delete: resourceWebfilterSearchEngineDelete,

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
			"charset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"safesearch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"safesearch_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWebfilterSearchEngineCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWebfilterSearchEngine(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterSearchEngine resource while getting object: %v", err)
	}

	_, err = c.CreateWebfilterSearchEngine(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterSearchEngine resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebfilterSearchEngineRead(d, m)
}

func resourceWebfilterSearchEngineUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWebfilterSearchEngine(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterSearchEngine resource while getting object: %v", err)
	}

	_, err = c.UpdateWebfilterSearchEngine(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterSearchEngine resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebfilterSearchEngineRead(d, m)
}

func resourceWebfilterSearchEngineDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebfilterSearchEngine(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterSearchEngine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterSearchEngineRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterSearchEngine(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterSearchEngine resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterSearchEngine(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterSearchEngine resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterSearchEngineCharset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterSearchEngineHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterSearchEngineName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterSearchEngineQuery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterSearchEngineSafesearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterSearchEngineSafesearchStr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterSearchEngineUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterSearchEngine(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("charset", flattenWebfilterSearchEngineCharset(o["charset"], d, "charset")); err != nil {
		if vv, ok := fortiAPIPatch(o["charset"], "WebfilterSearchEngine-Charset"); ok {
			if err = d.Set("charset", vv); err != nil {
				return fmt.Errorf("Error reading charset: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading charset: %v", err)
		}
	}

	if err = d.Set("hostname", flattenWebfilterSearchEngineHostname(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "WebfilterSearchEngine-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("name", flattenWebfilterSearchEngineName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebfilterSearchEngine-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("query", flattenWebfilterSearchEngineQuery(o["query"], d, "query")); err != nil {
		if vv, ok := fortiAPIPatch(o["query"], "WebfilterSearchEngine-Query"); ok {
			if err = d.Set("query", vv); err != nil {
				return fmt.Errorf("Error reading query: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query: %v", err)
		}
	}

	if err = d.Set("safesearch", flattenWebfilterSearchEngineSafesearch(o["safesearch"], d, "safesearch")); err != nil {
		if vv, ok := fortiAPIPatch(o["safesearch"], "WebfilterSearchEngine-Safesearch"); ok {
			if err = d.Set("safesearch", vv); err != nil {
				return fmt.Errorf("Error reading safesearch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading safesearch: %v", err)
		}
	}

	if err = d.Set("safesearch_str", flattenWebfilterSearchEngineSafesearchStr(o["safesearch-str"], d, "safesearch_str")); err != nil {
		if vv, ok := fortiAPIPatch(o["safesearch-str"], "WebfilterSearchEngine-SafesearchStr"); ok {
			if err = d.Set("safesearch_str", vv); err != nil {
				return fmt.Errorf("Error reading safesearch_str: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading safesearch_str: %v", err)
		}
	}

	if err = d.Set("url", flattenWebfilterSearchEngineUrl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "WebfilterSearchEngine-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func flattenWebfilterSearchEngineFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterSearchEngineCharset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterSearchEngineHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterSearchEngineName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterSearchEngineQuery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterSearchEngineSafesearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterSearchEngineSafesearchStr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterSearchEngineUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterSearchEngine(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("charset"); ok || d.HasChange("charset") {
		t, err := expandWebfilterSearchEngineCharset(d, v, "charset")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["charset"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandWebfilterSearchEngineHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebfilterSearchEngineName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("query"); ok || d.HasChange("query") {
		t, err := expandWebfilterSearchEngineQuery(d, v, "query")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query"] = t
		}
	}

	if v, ok := d.GetOk("safesearch"); ok || d.HasChange("safesearch") {
		t, err := expandWebfilterSearchEngineSafesearch(d, v, "safesearch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safesearch"] = t
		}
	}

	if v, ok := d.GetOk("safesearch_str"); ok || d.HasChange("safesearch_str") {
		t, err := expandWebfilterSearchEngineSafesearchStr(d, v, "safesearch_str")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safesearch-str"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandWebfilterSearchEngineUrl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
