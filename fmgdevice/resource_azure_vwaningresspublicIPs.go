// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Display Azure vWAN SLB ingress public IPs.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAzureVwanIngressPublicIps() *schema.Resource {
	return &schema.Resource{
		Create: resourceAzureVwanIngressPublicIpsCreate,
		Read:   resourceAzureVwanIngressPublicIpsRead,
		Update: resourceAzureVwanIngressPublicIpsUpdate,
		Delete: resourceAzureVwanIngressPublicIpsDelete,

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
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceAzureVwanIngressPublicIpsCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectAzureVwanIngressPublicIps(d)
	if err != nil {
		return fmt.Errorf("Error creating AzureVwanIngressPublicIps resource while getting object: %v", err)
	}

	_, err = c.CreateAzureVwanIngressPublicIps(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating AzureVwanIngressPublicIps resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceAzureVwanIngressPublicIpsRead(d, m)
}

func resourceAzureVwanIngressPublicIpsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectAzureVwanIngressPublicIps(d)
	if err != nil {
		return fmt.Errorf("Error updating AzureVwanIngressPublicIps resource while getting object: %v", err)
	}

	_, err = c.UpdateAzureVwanIngressPublicIps(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AzureVwanIngressPublicIps resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceAzureVwanIngressPublicIpsRead(d, m)
}

func resourceAzureVwanIngressPublicIpsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteAzureVwanIngressPublicIps(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AzureVwanIngressPublicIps resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAzureVwanIngressPublicIpsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAzureVwanIngressPublicIps(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading AzureVwanIngressPublicIps resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAzureVwanIngressPublicIps(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AzureVwanIngressPublicIps resource from API: %v", err)
	}
	return nil
}

func flattenAzureVwanIngressPublicIpsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAzureVwanIngressPublicIpsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAzureVwanIngressPublicIps(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip", flattenAzureVwanIngressPublicIpsIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "AzureVwanIngressPublicIps-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("name", flattenAzureVwanIngressPublicIpsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "AzureVwanIngressPublicIps-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenAzureVwanIngressPublicIpsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAzureVwanIngressPublicIpsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAzureVwanIngressPublicIpsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAzureVwanIngressPublicIps(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandAzureVwanIngressPublicIpsIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandAzureVwanIngressPublicIpsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
