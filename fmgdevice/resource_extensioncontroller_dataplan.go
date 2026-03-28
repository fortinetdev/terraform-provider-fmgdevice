// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender dataplan configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerDataplan() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerDataplanCreate,
		Read:   resourceExtensionControllerDataplanRead,
		Update: resourceExtensionControllerDataplanUpdate,
		Delete: resourceExtensionControllerDataplanDelete,

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
			"apn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"billing_date": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"carrier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"iccid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"modem_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monthly_fee": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"overage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"pdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preferred_subnet": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"private_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signal_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"signal_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"slot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceExtensionControllerDataplanCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerDataplan(d)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerDataplan resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadExtensionControllerDataplan(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateExtensionControllerDataplan(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ExtensionControllerDataplan resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateExtensionControllerDataplan(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ExtensionControllerDataplan resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerDataplanRead(d, m)
}

func resourceExtensionControllerDataplanUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerDataplan(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerDataplan resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateExtensionControllerDataplan(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerDataplan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerDataplanRead(d, m)
}

func resourceExtensionControllerDataplanDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteExtensionControllerDataplan(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerDataplan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerDataplanRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerDataplan(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ExtensionControllerDataplan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerDataplan(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerDataplan resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerDataplanApn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanBillingDate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanIccid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanModemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanMonthlyFee(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanOverage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanPdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanPreferredSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanPrivateNetwork(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanSignalPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanSignalThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanSlot(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerDataplanUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerDataplan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("apn", flattenExtensionControllerDataplanApn(o["apn"], d, "apn")); err != nil {
		if vv, ok := fortiAPIPatch(o["apn"], "ExtensionControllerDataplan-Apn"); ok {
			if err = d.Set("apn", vv); err != nil {
				return fmt.Errorf("Error reading apn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apn: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenExtensionControllerDataplanAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "ExtensionControllerDataplan-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("billing_date", flattenExtensionControllerDataplanBillingDate(o["billing-date"], d, "billing_date")); err != nil {
		if vv, ok := fortiAPIPatch(o["billing-date"], "ExtensionControllerDataplan-BillingDate"); ok {
			if err = d.Set("billing_date", vv); err != nil {
				return fmt.Errorf("Error reading billing_date: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading billing_date: %v", err)
		}
	}

	if err = d.Set("capacity", flattenExtensionControllerDataplanCapacity(o["capacity"], d, "capacity")); err != nil {
		if vv, ok := fortiAPIPatch(o["capacity"], "ExtensionControllerDataplan-Capacity"); ok {
			if err = d.Set("capacity", vv); err != nil {
				return fmt.Errorf("Error reading capacity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capacity: %v", err)
		}
	}

	if err = d.Set("carrier", flattenExtensionControllerDataplanCarrier(o["carrier"], d, "carrier")); err != nil {
		if vv, ok := fortiAPIPatch(o["carrier"], "ExtensionControllerDataplan-Carrier"); ok {
			if err = d.Set("carrier", vv); err != nil {
				return fmt.Errorf("Error reading carrier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading carrier: %v", err)
		}
	}

	if err = d.Set("iccid", flattenExtensionControllerDataplanIccid(o["iccid"], d, "iccid")); err != nil {
		if vv, ok := fortiAPIPatch(o["iccid"], "ExtensionControllerDataplan-Iccid"); ok {
			if err = d.Set("iccid", vv); err != nil {
				return fmt.Errorf("Error reading iccid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iccid: %v", err)
		}
	}

	if err = d.Set("modem_id", flattenExtensionControllerDataplanModemId(o["modem-id"], d, "modem_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem-id"], "ExtensionControllerDataplan-ModemId"); ok {
			if err = d.Set("modem_id", vv); err != nil {
				return fmt.Errorf("Error reading modem_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem_id: %v", err)
		}
	}

	if err = d.Set("monthly_fee", flattenExtensionControllerDataplanMonthlyFee(o["monthly-fee"], d, "monthly_fee")); err != nil {
		if vv, ok := fortiAPIPatch(o["monthly-fee"], "ExtensionControllerDataplan-MonthlyFee"); ok {
			if err = d.Set("monthly_fee", vv); err != nil {
				return fmt.Errorf("Error reading monthly_fee: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monthly_fee: %v", err)
		}
	}

	if err = d.Set("name", flattenExtensionControllerDataplanName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ExtensionControllerDataplan-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("overage", flattenExtensionControllerDataplanOverage(o["overage"], d, "overage")); err != nil {
		if vv, ok := fortiAPIPatch(o["overage"], "ExtensionControllerDataplan-Overage"); ok {
			if err = d.Set("overage", vv); err != nil {
				return fmt.Errorf("Error reading overage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overage: %v", err)
		}
	}

	if err = d.Set("pdn", flattenExtensionControllerDataplanPdn(o["pdn"], d, "pdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdn"], "ExtensionControllerDataplan-Pdn"); ok {
			if err = d.Set("pdn", vv); err != nil {
				return fmt.Errorf("Error reading pdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdn: %v", err)
		}
	}

	if err = d.Set("preferred_subnet", flattenExtensionControllerDataplanPreferredSubnet(o["preferred-subnet"], d, "preferred_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-subnet"], "ExtensionControllerDataplan-PreferredSubnet"); ok {
			if err = d.Set("preferred_subnet", vv); err != nil {
				return fmt.Errorf("Error reading preferred_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_subnet: %v", err)
		}
	}

	if err = d.Set("private_network", flattenExtensionControllerDataplanPrivateNetwork(o["private-network"], d, "private_network")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-network"], "ExtensionControllerDataplan-PrivateNetwork"); ok {
			if err = d.Set("private_network", vv); err != nil {
				return fmt.Errorf("Error reading private_network: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_network: %v", err)
		}
	}

	if err = d.Set("signal_period", flattenExtensionControllerDataplanSignalPeriod(o["signal-period"], d, "signal_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["signal-period"], "ExtensionControllerDataplan-SignalPeriod"); ok {
			if err = d.Set("signal_period", vv); err != nil {
				return fmt.Errorf("Error reading signal_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signal_period: %v", err)
		}
	}

	if err = d.Set("signal_threshold", flattenExtensionControllerDataplanSignalThreshold(o["signal-threshold"], d, "signal_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["signal-threshold"], "ExtensionControllerDataplan-SignalThreshold"); ok {
			if err = d.Set("signal_threshold", vv); err != nil {
				return fmt.Errorf("Error reading signal_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signal_threshold: %v", err)
		}
	}

	if err = d.Set("slot", flattenExtensionControllerDataplanSlot(o["slot"], d, "slot")); err != nil {
		if vv, ok := fortiAPIPatch(o["slot"], "ExtensionControllerDataplan-Slot"); ok {
			if err = d.Set("slot", vv); err != nil {
				return fmt.Errorf("Error reading slot: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slot: %v", err)
		}
	}

	if err = d.Set("type", flattenExtensionControllerDataplanType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ExtensionControllerDataplan-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("username", flattenExtensionControllerDataplanUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "ExtensionControllerDataplan-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerDataplanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerDataplanApn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanBillingDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanIccid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanModemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanMonthlyFee(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanOverage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerDataplanPdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanPreferredSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanPrivateNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanSignalPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanSignalThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanSlot(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerDataplanUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerDataplan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("apn"); ok || d.HasChange("apn") {
		t, err := expandExtensionControllerDataplanApn(d, v, "apn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apn"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandExtensionControllerDataplanAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("billing_date"); ok || d.HasChange("billing_date") {
		t, err := expandExtensionControllerDataplanBillingDate(d, v, "billing_date")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["billing-date"] = t
		}
	}

	if v, ok := d.GetOk("capacity"); ok || d.HasChange("capacity") {
		t, err := expandExtensionControllerDataplanCapacity(d, v, "capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capacity"] = t
		}
	}

	if v, ok := d.GetOk("carrier"); ok || d.HasChange("carrier") {
		t, err := expandExtensionControllerDataplanCarrier(d, v, "carrier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier"] = t
		}
	}

	if v, ok := d.GetOk("iccid"); ok || d.HasChange("iccid") {
		t, err := expandExtensionControllerDataplanIccid(d, v, "iccid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iccid"] = t
		}
	}

	if v, ok := d.GetOk("modem_id"); ok || d.HasChange("modem_id") {
		t, err := expandExtensionControllerDataplanModemId(d, v, "modem_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-id"] = t
		}
	}

	if v, ok := d.GetOk("monthly_fee"); ok || d.HasChange("monthly_fee") {
		t, err := expandExtensionControllerDataplanMonthlyFee(d, v, "monthly_fee")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monthly-fee"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandExtensionControllerDataplanName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("overage"); ok || d.HasChange("overage") {
		t, err := expandExtensionControllerDataplanOverage(d, v, "overage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandExtensionControllerDataplanPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("pdn"); ok || d.HasChange("pdn") {
		t, err := expandExtensionControllerDataplanPdn(d, v, "pdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdn"] = t
		}
	}

	if v, ok := d.GetOk("preferred_subnet"); ok || d.HasChange("preferred_subnet") {
		t, err := expandExtensionControllerDataplanPreferredSubnet(d, v, "preferred_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-subnet"] = t
		}
	}

	if v, ok := d.GetOk("private_network"); ok || d.HasChange("private_network") {
		t, err := expandExtensionControllerDataplanPrivateNetwork(d, v, "private_network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-network"] = t
		}
	}

	if v, ok := d.GetOk("signal_period"); ok || d.HasChange("signal_period") {
		t, err := expandExtensionControllerDataplanSignalPeriod(d, v, "signal_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-period"] = t
		}
	}

	if v, ok := d.GetOk("signal_threshold"); ok || d.HasChange("signal_threshold") {
		t, err := expandExtensionControllerDataplanSignalThreshold(d, v, "signal_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-threshold"] = t
		}
	}

	if v, ok := d.GetOk("slot"); ok || d.HasChange("slot") {
		t, err := expandExtensionControllerDataplanSlot(d, v, "slot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slot"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandExtensionControllerDataplanType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandExtensionControllerDataplanUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
