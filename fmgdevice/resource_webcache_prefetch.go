// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device WebcachePrefetch

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebcachePrefetch() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebcachePrefetchCreate,
		Read:   resourceWebcachePrefetchRead,
		Update: resourceWebcachePrefetchUpdate,
		Delete: resourceWebcachePrefetchDelete,

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
			"crawl_depth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ignore_robots": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"repeat": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_agent": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebcachePrefetchCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebcachePrefetch(d)
	if err != nil {
		return fmt.Errorf("Error creating WebcachePrefetch resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebcachePrefetch(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebcachePrefetch(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebcachePrefetch resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebcachePrefetch(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebcachePrefetch resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebcachePrefetchRead(d, m)
}

func resourceWebcachePrefetchUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebcachePrefetch(d)
	if err != nil {
		return fmt.Errorf("Error updating WebcachePrefetch resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebcachePrefetch(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebcachePrefetch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebcachePrefetchRead(d, m)
}

func resourceWebcachePrefetchDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebcachePrefetch(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebcachePrefetch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebcachePrefetchRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebcachePrefetch(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebcachePrefetch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebcachePrefetch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebcachePrefetch resource from API: %v", err)
	}
	return nil
}

func flattenWebcachePrefetchCrawlDepth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcachePrefetchIgnoreRobots(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcachePrefetchInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcachePrefetchName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcachePrefetchRepeat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcachePrefetchStartDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcachePrefetchUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcachePrefetchUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcachePrefetchUserAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWebcachePrefetch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("crawl_depth", flattenWebcachePrefetchCrawlDepth(o["crawl-depth"], d, "crawl_depth")); err != nil {
		if vv, ok := fortiAPIPatch(o["crawl-depth"], "WebcachePrefetch-CrawlDepth"); ok {
			if err = d.Set("crawl_depth", vv); err != nil {
				return fmt.Errorf("Error reading crawl_depth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading crawl_depth: %v", err)
		}
	}

	if err = d.Set("ignore_robots", flattenWebcachePrefetchIgnoreRobots(o["ignore-robots"], d, "ignore_robots")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-robots"], "WebcachePrefetch-IgnoreRobots"); ok {
			if err = d.Set("ignore_robots", vv); err != nil {
				return fmt.Errorf("Error reading ignore_robots: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_robots: %v", err)
		}
	}

	if err = d.Set("interval", flattenWebcachePrefetchInterval(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "WebcachePrefetch-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("name", flattenWebcachePrefetchName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebcachePrefetch-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("repeat", flattenWebcachePrefetchRepeat(o["repeat"], d, "repeat")); err != nil {
		if vv, ok := fortiAPIPatch(o["repeat"], "WebcachePrefetch-Repeat"); ok {
			if err = d.Set("repeat", vv); err != nil {
				return fmt.Errorf("Error reading repeat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading repeat: %v", err)
		}
	}

	if err = d.Set("start_delay", flattenWebcachePrefetchStartDelay(o["start-delay"], d, "start_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-delay"], "WebcachePrefetch-StartDelay"); ok {
			if err = d.Set("start_delay", vv); err != nil {
				return fmt.Errorf("Error reading start_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_delay: %v", err)
		}
	}

	if err = d.Set("url", flattenWebcachePrefetchUrl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "WebcachePrefetch-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("user", flattenWebcachePrefetchUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "WebcachePrefetch-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("user_agent", flattenWebcachePrefetchUserAgent(o["user-agent"], d, "user_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-agent"], "WebcachePrefetch-UserAgent"); ok {
			if err = d.Set("user_agent", vv); err != nil {
				return fmt.Errorf("Error reading user_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent: %v", err)
		}
	}

	return nil
}

func flattenWebcachePrefetchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebcachePrefetchCrawlDepth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcachePrefetchIgnoreRobots(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcachePrefetchInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcachePrefetchName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcachePrefetchPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebcachePrefetchRepeat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcachePrefetchStartDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcachePrefetchUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcachePrefetchUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcachePrefetchUserAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWebcachePrefetch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("crawl_depth"); ok || d.HasChange("crawl_depth") {
		t, err := expandWebcachePrefetchCrawlDepth(d, v, "crawl_depth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crawl-depth"] = t
		}
	}

	if v, ok := d.GetOk("ignore_robots"); ok || d.HasChange("ignore_robots") {
		t, err := expandWebcachePrefetchIgnoreRobots(d, v, "ignore_robots")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-robots"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok || d.HasChange("interval") {
		t, err := expandWebcachePrefetchInterval(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebcachePrefetchName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandWebcachePrefetchPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("repeat"); ok || d.HasChange("repeat") {
		t, err := expandWebcachePrefetchRepeat(d, v, "repeat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["repeat"] = t
		}
	}

	if v, ok := d.GetOk("start_delay"); ok || d.HasChange("start_delay") {
		t, err := expandWebcachePrefetchStartDelay(d, v, "start_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-delay"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandWebcachePrefetchUrl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandWebcachePrefetchUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("user_agent"); ok || d.HasChange("user_agent") {
		t, err := expandWebcachePrefetchUserAgent(d, v, "user_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent"] = t
		}
	}

	return &obj, nil
}
