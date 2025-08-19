// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Trusthost.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemApiUserTrusthost() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemApiUserTrusthostCreate,
		Read:   resourceSystemApiUserTrusthostRead,
		Update: resourceSystemApiUserTrusthostUpdate,
		Delete: resourceSystemApiUserTrusthostDelete,

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
			"api_user": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceSystemApiUserTrusthostCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	api_user := d.Get("api_user").(string)
	paradict["device"] = device_name
	paradict["api_user"] = api_user

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemApiUserTrusthost(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemApiUserTrusthost resource while getting object: %v", err)
	}

	_, err = c.CreateSystemApiUserTrusthost(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemApiUserTrusthost resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemApiUserTrusthostRead(d, m)
}

func resourceSystemApiUserTrusthostUpdate(d *schema.ResourceData, m interface{}) error {
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
	api_user := d.Get("api_user").(string)
	paradict["device"] = device_name
	paradict["api_user"] = api_user

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemApiUserTrusthost(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemApiUserTrusthost resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemApiUserTrusthost(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemApiUserTrusthost resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemApiUserTrusthostRead(d, m)
}

func resourceSystemApiUserTrusthostDelete(d *schema.ResourceData, m interface{}) error {
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
	api_user := d.Get("api_user").(string)
	paradict["device"] = device_name
	paradict["api_user"] = api_user

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemApiUserTrusthost(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemApiUserTrusthost resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemApiUserTrusthostRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	api_user := d.Get("api_user").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if api_user == "" {
		api_user = importOptionChecking(m.(*FortiClient).Cfg, "api_user")
		if api_user == "" {
			return fmt.Errorf("Parameter api_user is missing")
		}
		if err = d.Set("api_user", api_user); err != nil {
			return fmt.Errorf("Error set params api_user: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["api_user"] = api_user

	o, err := c.ReadSystemApiUserTrusthost(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemApiUserTrusthost resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemApiUserTrusthost(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemApiUserTrusthost resource from API: %v", err)
	}
	return nil
}

func flattenSystemApiUserTrusthostId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserTrusthostIpv4Trusthost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemApiUserTrusthostIpv6Trusthost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemApiUserTrusthostType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemApiUserTrusthost(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemApiUserTrusthostId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemApiUserTrusthost-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ipv4_trusthost", flattenSystemApiUserTrusthostIpv4Trusthost2edl(o["ipv4-trusthost"], d, "ipv4_trusthost")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-trusthost"], "SystemApiUserTrusthost-Ipv4Trusthost"); ok {
			if err = d.Set("ipv4_trusthost", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_trusthost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_trusthost: %v", err)
		}
	}

	if err = d.Set("ipv6_trusthost", flattenSystemApiUserTrusthostIpv6Trusthost2edl(o["ipv6-trusthost"], d, "ipv6_trusthost")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-trusthost"], "SystemApiUserTrusthost-Ipv6Trusthost"); ok {
			if err = d.Set("ipv6_trusthost", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_trusthost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_trusthost: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemApiUserTrusthostType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemApiUserTrusthost-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSystemApiUserTrusthostFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemApiUserTrusthostId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserTrusthostIpv4Trusthost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemApiUserTrusthostIpv6Trusthost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemApiUserTrusthostType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemApiUserTrusthost(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemApiUserTrusthostId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_trusthost"); ok || d.HasChange("ipv4_trusthost") {
		t, err := expandSystemApiUserTrusthostIpv4Trusthost2edl(d, v, "ipv4_trusthost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-trusthost"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_trusthost"); ok || d.HasChange("ipv6_trusthost") {
		t, err := expandSystemApiUserTrusthostIpv6Trusthost2edl(d, v, "ipv6_trusthost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-trusthost"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemApiUserTrusthostType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
