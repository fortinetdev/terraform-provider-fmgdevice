// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> MSN Hotmail.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterProfileMsnHotmail() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterProfileMsnHotmailUpdate,
		Read:   resourceEmailfilterProfileMsnHotmailRead,
		Update: resourceEmailfilterProfileMsnHotmailUpdate,
		Delete: resourceEmailfilterProfileMsnHotmailDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceEmailfilterProfileMsnHotmailUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectEmailfilterProfileMsnHotmail(d)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterProfileMsnHotmail resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEmailfilterProfileMsnHotmail(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterProfileMsnHotmail resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("EmailfilterProfileMsnHotmail")

	return resourceEmailfilterProfileMsnHotmailRead(d, m)
}

func resourceEmailfilterProfileMsnHotmailDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteEmailfilterProfileMsnHotmail(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterProfileMsnHotmail resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterProfileMsnHotmailRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadEmailfilterProfileMsnHotmail(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EmailfilterProfileMsnHotmail resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterProfileMsnHotmail(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterProfileMsnHotmail resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterProfileMsnHotmailLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileMsnHotmailLogAll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEmailfilterProfileMsnHotmail(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("log", flattenEmailfilterProfileMsnHotmailLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "EmailfilterProfileMsnHotmail-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("log_all", flattenEmailfilterProfileMsnHotmailLogAll2edl(o["log-all"], d, "log_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-all"], "EmailfilterProfileMsnHotmail-LogAll"); ok {
			if err = d.Set("log_all", vv); err != nil {
				return fmt.Errorf("Error reading log_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_all: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterProfileMsnHotmailFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEmailfilterProfileMsnHotmailLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileMsnHotmailLogAll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterProfileMsnHotmail(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandEmailfilterProfileMsnHotmailLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("log_all"); ok || d.HasChange("log_all") {
		t, err := expandEmailfilterProfileMsnHotmailLogAll2edl(d, v, "log_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-all"] = t
		}
	}

	return &obj, nil
}
