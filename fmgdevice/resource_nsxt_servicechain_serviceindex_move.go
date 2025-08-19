// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure service index.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNsxtServiceChainServiceIndexMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtServiceChainServiceIndexMoveUpdate,
		Read:   resourceNsxtServiceChainServiceIndexMoveRead,
		Update: resourceNsxtServiceChainServiceIndexMoveUpdate,
		Delete: resourceNsxtServiceChainServiceIndexMoveDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"state_pos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"service_chain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service_index": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNsxtServiceChainServiceIndexMoveUpdate(d *schema.ResourceData, m interface{}) error {
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
	service_chain := d.Get("service_chain").(string)
	service_index := d.Get("service_index").(string)
	paradict["device"] = device_name
	paradict["service_chain"] = service_chain
	paradict["service_index"] = service_index

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	target := d.Get("target").(string)
	obj, err := getObjectNsxtServiceChainServiceIndexMove(d)
	if err != nil {
		return fmt.Errorf("Error updating NsxtServiceChainServiceIndexMove resource while getting object: %v", err)
	}

	_, err = c.UpdateNsxtServiceChainServiceIndexMove(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating NsxtServiceChainServiceIndexMove resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("NsxtServiceChainServiceIndexMove" + "_" + device_name + "_" + service_chain + "_" + service_index + "_" + target)

	return resourceNsxtServiceChainServiceIndexMoveRead(d, m)
}

func resourceNsxtServiceChainServiceIndexMoveDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceNsxtServiceChainServiceIndexMoveRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	sid := d.Get("service_index").(string)
	did := d.Get("target").(string)
	action := d.Get("option").(string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	service_chain := d.Get("service_chain").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if service_chain == "" {
		service_chain = importOptionChecking(m.(*FortiClient).Cfg, "service_chain")
		if service_chain == "" {
			return fmt.Errorf("Parameter service_chain is missing")
		}
		if err = d.Set("service_chain", service_chain); err != nil {
			return fmt.Errorf("Error set params service_chain: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["service_chain"] = service_chain

	o, err := c.ReadNsxtServiceChainServiceIndexMove(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading NsxtServiceChainServiceIndexMove resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if o != nil {
		i := 1
		now_sid := -1
		now_did := -1

		for _, z := range o {
			if v, ok := z.(map[string]interface{}); ok {
				if _, ok := v["id"]; !ok {
					return fmt.Errorf("Error reading NsxtServiceChainServiceIndexMove resource: id doesn't exist.")
				}

				vid := fmt.Sprintf("%v", v["id"])

				if vid == sid {
					now_sid = i
				}

				if vid == did {
					now_did = i
				}
			} else {
				return fmt.Errorf("Error reading NsxtServiceChainServiceIndexMove resource: no valid map string.")
			}

			i += 1
		}

		state_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_pos = "id(" + sid + ") and target(" + did + ") were deleted"
			} else if now_sid == -1 {
				state_pos = "id(" + sid + ") was deleted"
			} else if now_did == -1 {
				state_pos = "target(" + did + ") was deleted"
			}
		} else {
			bconsistent := true
			if action == "before" {
				if now_sid != now_did-1 {
					bconsistent = false
				}
			}

			if action == "after" {
				if now_sid != now_did+1 {
					bconsistent = false
				}
			}

			if bconsistent == false {
				relative_pos := now_sid - now_did

				if relative_pos > 0 {
					state_pos = "id(" + sid + ") is " + strconv.Itoa(relative_pos) + " behind target(" + did + ")"
				} else {
					state_pos = "id(" + sid + ") is " + strconv.Itoa(-relative_pos) + " ahead of target(" + did + ")"
				}
			}
		}

		d.Set("state_pos", state_pos)
	}

	return nil
}

func flattenNsxtServiceChainServiceIndexMoveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandNsxtServiceChainServiceIndexMoveTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainServiceIndexMoveOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectNsxtServiceChainServiceIndexMove(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandNsxtServiceChainServiceIndexMoveTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandNsxtServiceChainServiceIndexMoveOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
