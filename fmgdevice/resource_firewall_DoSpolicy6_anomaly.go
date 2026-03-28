// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Anomaly name.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallDosPolicy6Anomaly() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallDosPolicy6AnomalyCreate,
		Read:   resourceFirewallDosPolicy6AnomalyRead,
		Update: resourceFirewallDosPolicy6AnomalyUpdate,
		Delete: resourceFirewallDosPolicy6AnomalyDelete,

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
			"dos_policy6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"quarantine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quarantine_expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"quarantine_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"synproxy_tcp_mss": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"synproxy_tcp_sack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"synproxy_tcp_timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"synproxy_tcp_window": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"synproxy_tcp_windowscale": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"synproxy_tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"synproxy_ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"thresholddefault": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceFirewallDosPolicy6AnomalyCreate(d *schema.ResourceData, m interface{}) error {
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
	dos_policy6 := d.Get("dos_policy6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["dos_policy6"] = dos_policy6

	obj, err := getObjectFirewallDosPolicy6Anomaly(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallDosPolicy6Anomaly resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallDosPolicy6Anomaly(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallDosPolicy6Anomaly(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallDosPolicy6Anomaly resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallDosPolicy6Anomaly(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallDosPolicy6Anomaly resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallDosPolicy6AnomalyRead(d, m)
}

func resourceFirewallDosPolicy6AnomalyUpdate(d *schema.ResourceData, m interface{}) error {
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
	dos_policy6 := d.Get("dos_policy6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["dos_policy6"] = dos_policy6

	obj, err := getObjectFirewallDosPolicy6Anomaly(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDosPolicy6Anomaly resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallDosPolicy6Anomaly(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDosPolicy6Anomaly resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallDosPolicy6AnomalyRead(d, m)
}

func resourceFirewallDosPolicy6AnomalyDelete(d *schema.ResourceData, m interface{}) error {
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
	dos_policy6 := d.Get("dos_policy6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["dos_policy6"] = dos_policy6

	wsParams["adom"] = adomv

	err = c.DeleteFirewallDosPolicy6Anomaly(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallDosPolicy6Anomaly resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallDosPolicy6AnomalyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	dos_policy6 := d.Get("dos_policy6").(string)
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
	if dos_policy6 == "" {
		dos_policy6 = importOptionChecking(m.(*FortiClient).Cfg, "dos_policy6")
		if dos_policy6 == "" {
			return fmt.Errorf("Parameter dos_policy6 is missing")
		}
		if err = d.Set("dos_policy6", dos_policy6); err != nil {
			return fmt.Errorf("Error set params dos_policy6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["dos_policy6"] = dos_policy6

	o, err := c.ReadFirewallDosPolicy6Anomaly(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallDosPolicy6Anomaly resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallDosPolicy6Anomaly(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallDosPolicy6Anomaly resource from API: %v", err)
	}
	return nil
}

func flattenFirewallDosPolicy6AnomalyAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyQuarantine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyQuarantineExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyQuarantineLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTcpMss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTcpSack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTcpTimestamp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTcpWindow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTcpWindowscale2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyThresholdDefault2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallDosPolicy6Anomaly(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenFirewallDosPolicy6AnomalyAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "FirewallDosPolicy6Anomaly-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("log", flattenFirewallDosPolicy6AnomalyLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "FirewallDosPolicy6Anomaly-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallDosPolicy6AnomalyName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallDosPolicy6Anomaly-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenFirewallDosPolicy6AnomalyQuarantine2edl(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "FirewallDosPolicy6Anomaly-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("quarantine_expiry", flattenFirewallDosPolicy6AnomalyQuarantineExpiry2edl(o["quarantine-expiry"], d, "quarantine_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-expiry"], "FirewallDosPolicy6Anomaly-QuarantineExpiry"); ok {
			if err = d.Set("quarantine_expiry", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_expiry: %v", err)
		}
	}

	if err = d.Set("quarantine_log", flattenFirewallDosPolicy6AnomalyQuarantineLog2edl(o["quarantine-log"], d, "quarantine_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-log"], "FirewallDosPolicy6Anomaly-QuarantineLog"); ok {
			if err = d.Set("quarantine_log", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_log: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallDosPolicy6AnomalyStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallDosPolicy6Anomaly-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_mss", flattenFirewallDosPolicy6AnomalySynproxyTcpMss2edl(o["synproxy-tcp-mss"], d, "synproxy_tcp_mss")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-mss"], "FirewallDosPolicy6Anomaly-SynproxyTcpMss"); ok {
			if err = d.Set("synproxy_tcp_mss", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_mss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_mss: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_sack", flattenFirewallDosPolicy6AnomalySynproxyTcpSack2edl(o["synproxy-tcp-sack"], d, "synproxy_tcp_sack")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-sack"], "FirewallDosPolicy6Anomaly-SynproxyTcpSack"); ok {
			if err = d.Set("synproxy_tcp_sack", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_sack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_sack: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_timestamp", flattenFirewallDosPolicy6AnomalySynproxyTcpTimestamp2edl(o["synproxy-tcp-timestamp"], d, "synproxy_tcp_timestamp")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-timestamp"], "FirewallDosPolicy6Anomaly-SynproxyTcpTimestamp"); ok {
			if err = d.Set("synproxy_tcp_timestamp", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_timestamp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_timestamp: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_window", flattenFirewallDosPolicy6AnomalySynproxyTcpWindow2edl(o["synproxy-tcp-window"], d, "synproxy_tcp_window")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-window"], "FirewallDosPolicy6Anomaly-SynproxyTcpWindow"); ok {
			if err = d.Set("synproxy_tcp_window", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_window: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_window: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_windowscale", flattenFirewallDosPolicy6AnomalySynproxyTcpWindowscale2edl(o["synproxy-tcp-windowscale"], d, "synproxy_tcp_windowscale")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-windowscale"], "FirewallDosPolicy6Anomaly-SynproxyTcpWindowscale"); ok {
			if err = d.Set("synproxy_tcp_windowscale", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_windowscale: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_windowscale: %v", err)
		}
	}

	if err = d.Set("synproxy_tos", flattenFirewallDosPolicy6AnomalySynproxyTos2edl(o["synproxy-tos"], d, "synproxy_tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tos"], "FirewallDosPolicy6Anomaly-SynproxyTos"); ok {
			if err = d.Set("synproxy_tos", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tos: %v", err)
		}
	}

	if err = d.Set("synproxy_ttl", flattenFirewallDosPolicy6AnomalySynproxyTtl2edl(o["synproxy-ttl"], d, "synproxy_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-ttl"], "FirewallDosPolicy6Anomaly-SynproxyTtl"); ok {
			if err = d.Set("synproxy_ttl", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_ttl: %v", err)
		}
	}

	if err = d.Set("threshold", flattenFirewallDosPolicy6AnomalyThreshold2edl(o["threshold"], d, "threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold"], "FirewallDosPolicy6Anomaly-Threshold"); ok {
			if err = d.Set("threshold", vv); err != nil {
				return fmt.Errorf("Error reading threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold: %v", err)
		}
	}

	if err = d.Set("thresholddefault", flattenFirewallDosPolicy6AnomalyThresholdDefault2edl(o["threshold(default)"], d, "thresholddefault")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold(default)"], "FirewallDosPolicy6Anomaly-ThresholdDefault"); ok {
			if err = d.Set("thresholddefault", vv); err != nil {
				return fmt.Errorf("Error reading thresholddefault: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading thresholddefault: %v", err)
		}
	}

	return nil
}

func flattenFirewallDosPolicy6AnomalyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallDosPolicy6AnomalyAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyQuarantine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyQuarantineExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyQuarantineLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTcpMss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTcpSack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTcpTimestamp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTcpWindow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTcpWindowscale2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyThresholdDefault2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallDosPolicy6Anomaly(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandFirewallDosPolicy6AnomalyAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandFirewallDosPolicy6AnomalyLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallDosPolicy6AnomalyName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandFirewallDosPolicy6AnomalyQuarantine2edl(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_expiry"); ok || d.HasChange("quarantine_expiry") {
		t, err := expandFirewallDosPolicy6AnomalyQuarantineExpiry2edl(d, v, "quarantine_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-expiry"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_log"); ok || d.HasChange("quarantine_log") {
		t, err := expandFirewallDosPolicy6AnomalyQuarantineLog2edl(d, v, "quarantine_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-log"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallDosPolicy6AnomalyStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_mss"); ok || d.HasChange("synproxy_tcp_mss") {
		t, err := expandFirewallDosPolicy6AnomalySynproxyTcpMss2edl(d, v, "synproxy_tcp_mss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-mss"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_sack"); ok || d.HasChange("synproxy_tcp_sack") {
		t, err := expandFirewallDosPolicy6AnomalySynproxyTcpSack2edl(d, v, "synproxy_tcp_sack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-sack"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_timestamp"); ok || d.HasChange("synproxy_tcp_timestamp") {
		t, err := expandFirewallDosPolicy6AnomalySynproxyTcpTimestamp2edl(d, v, "synproxy_tcp_timestamp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-timestamp"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_window"); ok || d.HasChange("synproxy_tcp_window") {
		t, err := expandFirewallDosPolicy6AnomalySynproxyTcpWindow2edl(d, v, "synproxy_tcp_window")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-window"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_windowscale"); ok || d.HasChange("synproxy_tcp_windowscale") {
		t, err := expandFirewallDosPolicy6AnomalySynproxyTcpWindowscale2edl(d, v, "synproxy_tcp_windowscale")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-windowscale"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tos"); ok || d.HasChange("synproxy_tos") {
		t, err := expandFirewallDosPolicy6AnomalySynproxyTos2edl(d, v, "synproxy_tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tos"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_ttl"); ok || d.HasChange("synproxy_ttl") {
		t, err := expandFirewallDosPolicy6AnomalySynproxyTtl2edl(d, v, "synproxy_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-ttl"] = t
		}
	}

	if v, ok := d.GetOk("threshold"); ok || d.HasChange("threshold") {
		t, err := expandFirewallDosPolicy6AnomalyThreshold2edl(d, v, "threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold"] = t
		}
	}

	if v, ok := d.GetOk("thresholddefault"); ok || d.HasChange("thresholddefault") {
		t, err := expandFirewallDosPolicy6AnomalyThresholdDefault2edl(d, v, "thresholddefault")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold(default)"] = t
		}
	}

	return &obj, nil
}
