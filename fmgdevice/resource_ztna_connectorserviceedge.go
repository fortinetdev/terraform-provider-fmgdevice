// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device ZtnaConnectorServiceEdge

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaConnectorServiceEdge() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaConnectorServiceEdgeUpdate,
		Read:   resourceZtnaConnectorServiceEdgeRead,
		Update: resourceZtnaConnectorServiceEdgeUpdate,
		Delete: resourceZtnaConnectorServiceEdgeDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted_client_ca": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceZtnaConnectorServiceEdgeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaConnectorServiceEdge(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaConnectorServiceEdge resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateZtnaConnectorServiceEdge(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaConnectorServiceEdge resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ZtnaConnectorServiceEdge")

	return resourceZtnaConnectorServiceEdgeRead(d, m)
}

func resourceZtnaConnectorServiceEdgeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteZtnaConnectorServiceEdge(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaConnectorServiceEdge resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaConnectorServiceEdgeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaConnectorServiceEdge(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ZtnaConnectorServiceEdge resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaConnectorServiceEdge(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaConnectorServiceEdge resource from API: %v", err)
	}
	return nil
}

func flattenZtnaConnectorServiceEdgeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaConnectorServiceEdgePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaConnectorServiceEdgeServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaConnectorServiceEdgeSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaConnectorServiceEdgeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaConnectorServiceEdgeTrustedClientCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectZtnaConnectorServiceEdge(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("interface", flattenZtnaConnectorServiceEdgeInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ZtnaConnectorServiceEdge-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("port", flattenZtnaConnectorServiceEdgePort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ZtnaConnectorServiceEdge-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("server_cert", flattenZtnaConnectorServiceEdgeServerCert(o["server-cert"], d, "server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-cert"], "ZtnaConnectorServiceEdge-ServerCert"); ok {
			if err = d.Set("server_cert", vv); err != nil {
				return fmt.Errorf("Error reading server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_cert: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenZtnaConnectorServiceEdgeSslMaxVersion(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-version"], "ZtnaConnectorServiceEdge-SslMaxVersion"); ok {
			if err = d.Set("ssl_max_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("status", flattenZtnaConnectorServiceEdgeStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ZtnaConnectorServiceEdge-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trusted_client_ca", flattenZtnaConnectorServiceEdgeTrustedClientCa(o["trusted-client-ca"], d, "trusted_client_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusted-client-ca"], "ZtnaConnectorServiceEdge-TrustedClientCa"); ok {
			if err = d.Set("trusted_client_ca", vv); err != nil {
				return fmt.Errorf("Error reading trusted_client_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusted_client_ca: %v", err)
		}
	}

	return nil
}

func flattenZtnaConnectorServiceEdgeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaConnectorServiceEdgeInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaConnectorServiceEdgePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaConnectorServiceEdgeServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaConnectorServiceEdgeSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaConnectorServiceEdgeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaConnectorServiceEdgeTrustedClientCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectZtnaConnectorServiceEdge(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandZtnaConnectorServiceEdgeInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandZtnaConnectorServiceEdgePort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("server_cert"); ok || d.HasChange("server_cert") {
		t, err := expandZtnaConnectorServiceEdgeServerCert(d, v, "server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-cert"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok || d.HasChange("ssl_max_version") {
		t, err := expandZtnaConnectorServiceEdgeSslMaxVersion(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandZtnaConnectorServiceEdgeStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trusted_client_ca"); ok || d.HasChange("trusted_client_ca") {
		t, err := expandZtnaConnectorServiceEdgeTrustedClientCa(d, v, "trusted_client_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-client-ca"] = t
		}
	}

	return &obj, nil
}
