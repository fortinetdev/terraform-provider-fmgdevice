// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> POP3.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterProfilePop3() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterProfilePop3Update,
		Read:   resourceEmailfilterProfilePop3Read,
		Update: resourceEmailfilterProfilePop3Update,
		Delete: resourceEmailfilterProfilePop3Delete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"tag_msg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceEmailfilterProfilePop3Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEmailfilterProfilePop3(d)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterProfilePop3 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEmailfilterProfilePop3(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterProfilePop3 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("EmailfilterProfilePop3")

	return resourceEmailfilterProfilePop3Read(d, m)
}

func resourceEmailfilterProfilePop3Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteEmailfilterProfilePop3(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterProfilePop3 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterProfilePop3Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadEmailfilterProfilePop3(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EmailfilterProfilePop3 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterProfilePop3(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterProfilePop3 resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterProfilePop3Action2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3Log2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3LogAll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3TagMsg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3TagType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectEmailfilterProfilePop3(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenEmailfilterProfilePop3Action2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "EmailfilterProfilePop3-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("log", flattenEmailfilterProfilePop3Log2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "EmailfilterProfilePop3-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("log_all", flattenEmailfilterProfilePop3LogAll2edl(o["log-all"], d, "log_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-all"], "EmailfilterProfilePop3-LogAll"); ok {
			if err = d.Set("log_all", vv); err != nil {
				return fmt.Errorf("Error reading log_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_all: %v", err)
		}
	}

	if err = d.Set("tag_msg", flattenEmailfilterProfilePop3TagMsg2edl(o["tag-msg"], d, "tag_msg")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag-msg"], "EmailfilterProfilePop3-TagMsg"); ok {
			if err = d.Set("tag_msg", vv); err != nil {
				return fmt.Errorf("Error reading tag_msg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag_msg: %v", err)
		}
	}

	if err = d.Set("tag_type", flattenEmailfilterProfilePop3TagType2edl(o["tag-type"], d, "tag_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag-type"], "EmailfilterProfilePop3-TagType"); ok {
			if err = d.Set("tag_type", vv); err != nil {
				return fmt.Errorf("Error reading tag_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag_type: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterProfilePop3FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEmailfilterProfilePop3Action2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3Log2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3LogAll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3TagMsg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3TagType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectEmailfilterProfilePop3(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandEmailfilterProfilePop3Action2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandEmailfilterProfilePop3Log2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("log_all"); ok || d.HasChange("log_all") {
		t, err := expandEmailfilterProfilePop3LogAll2edl(d, v, "log_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-all"] = t
		}
	}

	if v, ok := d.GetOk("tag_msg"); ok || d.HasChange("tag_msg") {
		t, err := expandEmailfilterProfilePop3TagMsg2edl(d, v, "tag_msg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag-msg"] = t
		}
	}

	if v, ok := d.GetOk("tag_type"); ok || d.HasChange("tag_type") {
		t, err := expandEmailfilterProfilePop3TagType2edl(d, v, "tag_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag-type"] = t
		}
	}

	return &obj, nil
}
