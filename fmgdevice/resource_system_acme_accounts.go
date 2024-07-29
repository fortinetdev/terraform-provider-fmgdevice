// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ACME accounts list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAcmeAccounts() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAcmeAccountsCreate,
		Read:   resourceSystemAcmeAccountsRead,
		Update: resourceSystemAcmeAccountsUpdate,
		Delete: resourceSystemAcmeAccountsDelete,

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
			"ca_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"privatekey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
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

func resourceSystemAcmeAccountsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAcmeAccounts(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAcmeAccounts resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAcmeAccounts(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAcmeAccounts resource: %v", err)
	}

	d.SetId(getStringKey(d, "fosid"))

	return resourceSystemAcmeAccountsRead(d, m)
}

func resourceSystemAcmeAccountsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAcmeAccounts(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAcmeAccounts resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAcmeAccounts(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAcmeAccounts resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "fosid"))

	return resourceSystemAcmeAccountsRead(d, m)
}

func resourceSystemAcmeAccountsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemAcmeAccounts(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAcmeAccounts resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAcmeAccountsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAcmeAccounts(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAcmeAccounts resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAcmeAccounts(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAcmeAccounts resource from API: %v", err)
	}
	return nil
}

func flattenSystemAcmeAccountsCaUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeAccountsEmail2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeAccountsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeAccountsPrivatekey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeAccountsStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAcmeAccountsUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAcmeAccounts(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ca_url", flattenSystemAcmeAccountsCaUrl2edl(o["ca_url"], d, "ca_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca_url"], "SystemAcmeAccounts-CaUrl"); ok {
			if err = d.Set("ca_url", vv); err != nil {
				return fmt.Errorf("Error reading ca_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca_url: %v", err)
		}
	}

	if err = d.Set("email", flattenSystemAcmeAccountsEmail2edl(o["email"], d, "email")); err != nil {
		if vv, ok := fortiAPIPatch(o["email"], "SystemAcmeAccounts-Email"); ok {
			if err = d.Set("email", vv); err != nil {
				return fmt.Errorf("Error reading email: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemAcmeAccountsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemAcmeAccounts-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("privatekey", flattenSystemAcmeAccountsPrivatekey2edl(o["privatekey"], d, "privatekey")); err != nil {
		if vv, ok := fortiAPIPatch(o["privatekey"], "SystemAcmeAccounts-Privatekey"); ok {
			if err = d.Set("privatekey", vv); err != nil {
				return fmt.Errorf("Error reading privatekey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading privatekey: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemAcmeAccountsStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemAcmeAccounts-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("url", flattenSystemAcmeAccountsUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "SystemAcmeAccounts-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func flattenSystemAcmeAccountsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAcmeAccountsCaUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeAccountsEmail2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeAccountsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeAccountsPrivatekey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeAccountsStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAcmeAccountsUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAcmeAccounts(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ca_url"); ok || d.HasChange("ca_url") {
		t, err := expandSystemAcmeAccountsCaUrl2edl(d, v, "ca_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca_url"] = t
		}
	}

	if v, ok := d.GetOk("email"); ok || d.HasChange("email") {
		t, err := expandSystemAcmeAccountsEmail2edl(d, v, "email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemAcmeAccountsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("privatekey"); ok || d.HasChange("privatekey") {
		t, err := expandSystemAcmeAccountsPrivatekey2edl(d, v, "privatekey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["privatekey"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemAcmeAccountsStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandSystemAcmeAccountsUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
