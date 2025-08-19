// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure network visibility settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNetworkVisibility() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNetworkVisibilityUpdate,
		Read:   resourceSystemNetworkVisibilityRead,
		Update: resourceSystemNetworkVisibilityUpdate,
		Delete: resourceSystemNetworkVisibilityDelete,

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
			"destination_hostname_visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination_visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hostname_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemNetworkVisibilityUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemNetworkVisibility(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNetworkVisibility resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNetworkVisibility(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemNetworkVisibility resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemNetworkVisibility")

	return resourceSystemNetworkVisibilityRead(d, m)
}

func resourceSystemNetworkVisibilityDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemNetworkVisibility(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNetworkVisibility resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNetworkVisibilityRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemNetworkVisibility(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemNetworkVisibility resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNetworkVisibility(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNetworkVisibility resource from API: %v", err)
	}
	return nil
}

func flattenSystemNetworkVisibilityDestinationHostnameVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNetworkVisibilityDestinationLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNetworkVisibilityDestinationVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNetworkVisibilityHostnameLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNetworkVisibilityHostnameTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNetworkVisibilitySourceLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNetworkVisibility(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("destination_hostname_visibility", flattenSystemNetworkVisibilityDestinationHostnameVisibility(o["destination-hostname-visibility"], d, "destination_hostname_visibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["destination-hostname-visibility"], "SystemNetworkVisibility-DestinationHostnameVisibility"); ok {
			if err = d.Set("destination_hostname_visibility", vv); err != nil {
				return fmt.Errorf("Error reading destination_hostname_visibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading destination_hostname_visibility: %v", err)
		}
	}

	if err = d.Set("destination_location", flattenSystemNetworkVisibilityDestinationLocation(o["destination-location"], d, "destination_location")); err != nil {
		if vv, ok := fortiAPIPatch(o["destination-location"], "SystemNetworkVisibility-DestinationLocation"); ok {
			if err = d.Set("destination_location", vv); err != nil {
				return fmt.Errorf("Error reading destination_location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading destination_location: %v", err)
		}
	}

	if err = d.Set("destination_visibility", flattenSystemNetworkVisibilityDestinationVisibility(o["destination-visibility"], d, "destination_visibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["destination-visibility"], "SystemNetworkVisibility-DestinationVisibility"); ok {
			if err = d.Set("destination_visibility", vv); err != nil {
				return fmt.Errorf("Error reading destination_visibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading destination_visibility: %v", err)
		}
	}

	if err = d.Set("hostname_limit", flattenSystemNetworkVisibilityHostnameLimit(o["hostname-limit"], d, "hostname_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname-limit"], "SystemNetworkVisibility-HostnameLimit"); ok {
			if err = d.Set("hostname_limit", vv); err != nil {
				return fmt.Errorf("Error reading hostname_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname_limit: %v", err)
		}
	}

	if err = d.Set("hostname_ttl", flattenSystemNetworkVisibilityHostnameTtl(o["hostname-ttl"], d, "hostname_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname-ttl"], "SystemNetworkVisibility-HostnameTtl"); ok {
			if err = d.Set("hostname_ttl", vv); err != nil {
				return fmt.Errorf("Error reading hostname_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname_ttl: %v", err)
		}
	}

	if err = d.Set("source_location", flattenSystemNetworkVisibilitySourceLocation(o["source-location"], d, "source_location")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-location"], "SystemNetworkVisibility-SourceLocation"); ok {
			if err = d.Set("source_location", vv); err != nil {
				return fmt.Errorf("Error reading source_location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_location: %v", err)
		}
	}

	return nil
}

func flattenSystemNetworkVisibilityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNetworkVisibilityDestinationHostnameVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNetworkVisibilityDestinationLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNetworkVisibilityDestinationVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNetworkVisibilityHostnameLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNetworkVisibilityHostnameTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNetworkVisibilitySourceLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNetworkVisibility(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("destination_hostname_visibility"); ok || d.HasChange("destination_hostname_visibility") {
		t, err := expandSystemNetworkVisibilityDestinationHostnameVisibility(d, v, "destination_hostname_visibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-hostname-visibility"] = t
		}
	}

	if v, ok := d.GetOk("destination_location"); ok || d.HasChange("destination_location") {
		t, err := expandSystemNetworkVisibilityDestinationLocation(d, v, "destination_location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-location"] = t
		}
	}

	if v, ok := d.GetOk("destination_visibility"); ok || d.HasChange("destination_visibility") {
		t, err := expandSystemNetworkVisibilityDestinationVisibility(d, v, "destination_visibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-visibility"] = t
		}
	}

	if v, ok := d.GetOk("hostname_limit"); ok || d.HasChange("hostname_limit") {
		t, err := expandSystemNetworkVisibilityHostnameLimit(d, v, "hostname_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname-limit"] = t
		}
	}

	if v, ok := d.GetOk("hostname_ttl"); ok || d.HasChange("hostname_ttl") {
		t, err := expandSystemNetworkVisibilityHostnameTtl(d, v, "hostname_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname-ttl"] = t
		}
	}

	if v, ok := d.GetOk("source_location"); ok || d.HasChange("source_location") {
		t, err := expandSystemNetworkVisibilitySourceLocation(d, v, "source_location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-location"] = t
		}
	}

	return &obj, nil
}
