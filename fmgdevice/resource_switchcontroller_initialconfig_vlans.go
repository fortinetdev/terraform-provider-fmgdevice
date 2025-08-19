// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure initial template for auto-generated VLAN interfaces.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerInitialConfigVlans() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerInitialConfigVlansUpdate,
		Read:   resourceSwitchControllerInitialConfigVlansRead,
		Update: resourceSwitchControllerInitialConfigVlansUpdate,
		Delete: resourceSwitchControllerInitialConfigVlansDelete,

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
			"default_vlan": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"nac": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"nac_segment": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"quarantine": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rspan": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"video": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"voice": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerInitialConfigVlansUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerInitialConfigVlans(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerInitialConfigVlans resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerInitialConfigVlans(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerInitialConfigVlans resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerInitialConfigVlans")

	return resourceSwitchControllerInitialConfigVlansRead(d, m)
}

func resourceSwitchControllerInitialConfigVlansDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerInitialConfigVlans(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerInitialConfigVlans resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerInitialConfigVlansRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerInitialConfigVlans(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerInitialConfigVlans resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerInitialConfigVlans(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerInitialConfigVlans resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerInitialConfigVlansDefaultVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerInitialConfigVlansNac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerInitialConfigVlansNacSegment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerInitialConfigVlansQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerInitialConfigVlansRspan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerInitialConfigVlansVideo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerInitialConfigVlansVoice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerInitialConfigVlans(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("default_vlan", flattenSwitchControllerInitialConfigVlansDefaultVlan(o["default-vlan"], d, "default_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-vlan"], "SwitchControllerInitialConfigVlans-DefaultVlan"); ok {
			if err = d.Set("default_vlan", vv); err != nil {
				return fmt.Errorf("Error reading default_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_vlan: %v", err)
		}
	}

	if err = d.Set("nac", flattenSwitchControllerInitialConfigVlansNac(o["nac"], d, "nac")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac"], "SwitchControllerInitialConfigVlans-Nac"); ok {
			if err = d.Set("nac", vv); err != nil {
				return fmt.Errorf("Error reading nac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac: %v", err)
		}
	}

	if err = d.Set("nac_segment", flattenSwitchControllerInitialConfigVlansNacSegment(o["nac-segment"], d, "nac_segment")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-segment"], "SwitchControllerInitialConfigVlans-NacSegment"); ok {
			if err = d.Set("nac_segment", vv); err != nil {
				return fmt.Errorf("Error reading nac_segment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_segment: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenSwitchControllerInitialConfigVlansQuarantine(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "SwitchControllerInitialConfigVlans-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("rspan", flattenSwitchControllerInitialConfigVlansRspan(o["rspan"], d, "rspan")); err != nil {
		if vv, ok := fortiAPIPatch(o["rspan"], "SwitchControllerInitialConfigVlans-Rspan"); ok {
			if err = d.Set("rspan", vv); err != nil {
				return fmt.Errorf("Error reading rspan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rspan: %v", err)
		}
	}

	if err = d.Set("video", flattenSwitchControllerInitialConfigVlansVideo(o["video"], d, "video")); err != nil {
		if vv, ok := fortiAPIPatch(o["video"], "SwitchControllerInitialConfigVlans-Video"); ok {
			if err = d.Set("video", vv); err != nil {
				return fmt.Errorf("Error reading video: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading video: %v", err)
		}
	}

	if err = d.Set("voice", flattenSwitchControllerInitialConfigVlansVoice(o["voice"], d, "voice")); err != nil {
		if vv, ok := fortiAPIPatch(o["voice"], "SwitchControllerInitialConfigVlans-Voice"); ok {
			if err = d.Set("voice", vv); err != nil {
				return fmt.Errorf("Error reading voice: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voice: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerInitialConfigVlansFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerInitialConfigVlansDefaultVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerInitialConfigVlansNac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerInitialConfigVlansNacSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerInitialConfigVlansQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerInitialConfigVlansRspan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerInitialConfigVlansVideo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerInitialConfigVlansVoice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerInitialConfigVlans(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_vlan"); ok || d.HasChange("default_vlan") {
		t, err := expandSwitchControllerInitialConfigVlansDefaultVlan(d, v, "default_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-vlan"] = t
		}
	}

	if v, ok := d.GetOk("nac"); ok || d.HasChange("nac") {
		t, err := expandSwitchControllerInitialConfigVlansNac(d, v, "nac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac"] = t
		}
	}

	if v, ok := d.GetOk("nac_segment"); ok || d.HasChange("nac_segment") {
		t, err := expandSwitchControllerInitialConfigVlansNacSegment(d, v, "nac_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-segment"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandSwitchControllerInitialConfigVlansQuarantine(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("rspan"); ok || d.HasChange("rspan") {
		t, err := expandSwitchControllerInitialConfigVlansRspan(d, v, "rspan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rspan"] = t
		}
	}

	if v, ok := d.GetOk("video"); ok || d.HasChange("video") {
		t, err := expandSwitchControllerInitialConfigVlansVideo(d, v, "video")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["video"] = t
		}
	}

	if v, ok := d.GetOk("voice"); ok || d.HasChange("voice") {
		t, err := expandSwitchControllerInitialConfigVlansVoice(d, v, "voice")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voice"] = t
		}
	}

	return &obj, nil
}
