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

func resourceNsxtServiceChainServiceIndex() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtServiceChainServiceIndexCreate,
		Read:   resourceNsxtServiceChainServiceIndexRead,
		Update: resourceNsxtServiceChainServiceIndexUpdate,
		Delete: resourceNsxtServiceChainServiceIndexDelete,

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
			"service_chain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reverse_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vd": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceNsxtServiceChainServiceIndexCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["service_chain"] = service_chain

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectNsxtServiceChainServiceIndex(d)
	if err != nil {
		return fmt.Errorf("Error creating NsxtServiceChainServiceIndex resource while getting object: %v", err)
	}

	_, err = c.CreateNsxtServiceChainServiceIndex(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating NsxtServiceChainServiceIndex resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceNsxtServiceChainServiceIndexRead(d, m)
}

func resourceNsxtServiceChainServiceIndexUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["service_chain"] = service_chain

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectNsxtServiceChainServiceIndex(d)
	if err != nil {
		return fmt.Errorf("Error updating NsxtServiceChainServiceIndex resource while getting object: %v", err)
	}

	_, err = c.UpdateNsxtServiceChainServiceIndex(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating NsxtServiceChainServiceIndex resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceNsxtServiceChainServiceIndexRead(d, m)
}

func resourceNsxtServiceChainServiceIndexDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["service_chain"] = service_chain

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteNsxtServiceChainServiceIndex(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting NsxtServiceChainServiceIndex resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceNsxtServiceChainServiceIndexRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	o, err := c.ReadNsxtServiceChainServiceIndex(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading NsxtServiceChainServiceIndex resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectNsxtServiceChainServiceIndex(d, o)
	if err != nil {
		return fmt.Errorf("Error reading NsxtServiceChainServiceIndex resource from API: %v", err)
	}
	return nil
}

func flattenNsxtServiceChainServiceIndexId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenNsxtServiceChainServiceIndexName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenNsxtServiceChainServiceIndexReverseIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenNsxtServiceChainServiceIndexVd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectNsxtServiceChainServiceIndex(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenNsxtServiceChainServiceIndexId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "NsxtServiceChainServiceIndex-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenNsxtServiceChainServiceIndexName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "NsxtServiceChainServiceIndex-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("reverse_index", flattenNsxtServiceChainServiceIndexReverseIndex2edl(o["reverse-index"], d, "reverse_index")); err != nil {
		if vv, ok := fortiAPIPatch(o["reverse-index"], "NsxtServiceChainServiceIndex-ReverseIndex"); ok {
			if err = d.Set("reverse_index", vv); err != nil {
				return fmt.Errorf("Error reading reverse_index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reverse_index: %v", err)
		}
	}

	if err = d.Set("vd", flattenNsxtServiceChainServiceIndexVd2edl(o["vd"], d, "vd")); err != nil {
		if vv, ok := fortiAPIPatch(o["vd"], "NsxtServiceChainServiceIndex-Vd"); ok {
			if err = d.Set("vd", vv); err != nil {
				return fmt.Errorf("Error reading vd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vd: %v", err)
		}
	}

	return nil
}

func flattenNsxtServiceChainServiceIndexFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandNsxtServiceChainServiceIndexId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainServiceIndexName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainServiceIndexReverseIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainServiceIndexVd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectNsxtServiceChainServiceIndex(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandNsxtServiceChainServiceIndexId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandNsxtServiceChainServiceIndexName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("reverse_index"); ok || d.HasChange("reverse_index") {
		t, err := expandNsxtServiceChainServiceIndexReverseIndex2edl(d, v, "reverse_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reverse-index"] = t
		}
	}

	if v, ok := d.GetOk("vd"); ok || d.HasChange("vd") {
		t, err := expandNsxtServiceChainServiceIndexVd2edl(d, v, "vd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vd"] = t
		}
	}

	return &obj, nil
}
