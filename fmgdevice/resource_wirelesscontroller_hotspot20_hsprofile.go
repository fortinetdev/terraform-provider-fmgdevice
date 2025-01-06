// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure hotspot profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20HsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20HsProfileCreate,
		Read:   resourceWirelessControllerHotspot20HsProfileRead,
		Update: resourceWirelessControllerHotspot20HsProfileUpdate,
		Delete: resourceWirelessControllerHotspot20HsProfileDelete,

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
			"n3gpp_plmn": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"access_network_asra": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_esr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_internet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_uesa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"advice_of_charge": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"anqp_domain_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bss_transition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conn_cap": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"deauth_request_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dgaf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gas_comeback_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gas_fragmentation_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hessid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_addr_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"l2tif": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nai_realm": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"network_auth": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"oper_friendly_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"oper_icon": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"osu_provider": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"osu_provider_nai": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"osu_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pame_bi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_arp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_map": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"release": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"roaming_consortium": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"terms_and_conditions": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"venue_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"venue_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"venue_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"venue_url": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wan_metrics": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wba_charging_currency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wba_charging_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wba_data_clearing_provider": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wba_financial_clearing_provider": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wba_open_roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wnm_sleep_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20HsProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectWirelessControllerHotspot20HsProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20HsProfile resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20HsProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20HsProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20HsProfileRead(d, m)
}

func resourceWirelessControllerHotspot20HsProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectWirelessControllerHotspot20HsProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20HsProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20HsProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20HsProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20HsProfileRead(d, m)
}

func resourceWirelessControllerHotspot20HsProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteWirelessControllerHotspot20HsProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20HsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20HsProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerHotspot20HsProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20HsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20HsProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20HsProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20HsProfile3GppPlmn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileAccessNetworkAsra(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileAccessNetworkEsr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileAccessNetworkInternet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileAccessNetworkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileAccessNetworkUesa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileAdviceOfCharge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileAnqpDomainId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileBssTransition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileConnCap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileDeauthRequestTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileDgaf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileGasComebackDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileGasFragmentationLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileHessid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileIpAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileL2Tif(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileNaiRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileNetworkAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileOperFriendlyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileOperIcon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileOsuProvider(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileOsuProviderNai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileOsuSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfilePameBi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileProxyArp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileQosMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileRelease(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileRoamingConsortium(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileTermsAndConditions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileVenueGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileVenueName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileVenueType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileVenueUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileWanMetrics(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerHotspot20HsProfileWbaChargingCurrency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileWbaChargingRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileWbaDataClearingProvider(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileWbaFinancialClearingProvider(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileWbaOpenRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileWnmSleepMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20HsProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("n3gpp_plmn", flattenWirelessControllerHotspot20HsProfile3GppPlmn(o["3gpp-plmn"], d, "n3gpp_plmn")); err != nil {
		if vv, ok := fortiAPIPatch(o["3gpp-plmn"], "WirelessControllerHotspot20HsProfile-3GppPlmn"); ok {
			if err = d.Set("n3gpp_plmn", vv); err != nil {
				return fmt.Errorf("Error reading n3gpp_plmn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n3gpp_plmn: %v", err)
		}
	}

	if err = d.Set("access_network_asra", flattenWirelessControllerHotspot20HsProfileAccessNetworkAsra(o["access-network-asra"], d, "access_network_asra")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-network-asra"], "WirelessControllerHotspot20HsProfile-AccessNetworkAsra"); ok {
			if err = d.Set("access_network_asra", vv); err != nil {
				return fmt.Errorf("Error reading access_network_asra: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_network_asra: %v", err)
		}
	}

	if err = d.Set("access_network_esr", flattenWirelessControllerHotspot20HsProfileAccessNetworkEsr(o["access-network-esr"], d, "access_network_esr")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-network-esr"], "WirelessControllerHotspot20HsProfile-AccessNetworkEsr"); ok {
			if err = d.Set("access_network_esr", vv); err != nil {
				return fmt.Errorf("Error reading access_network_esr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_network_esr: %v", err)
		}
	}

	if err = d.Set("access_network_internet", flattenWirelessControllerHotspot20HsProfileAccessNetworkInternet(o["access-network-internet"], d, "access_network_internet")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-network-internet"], "WirelessControllerHotspot20HsProfile-AccessNetworkInternet"); ok {
			if err = d.Set("access_network_internet", vv); err != nil {
				return fmt.Errorf("Error reading access_network_internet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_network_internet: %v", err)
		}
	}

	if err = d.Set("access_network_type", flattenWirelessControllerHotspot20HsProfileAccessNetworkType(o["access-network-type"], d, "access_network_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-network-type"], "WirelessControllerHotspot20HsProfile-AccessNetworkType"); ok {
			if err = d.Set("access_network_type", vv); err != nil {
				return fmt.Errorf("Error reading access_network_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_network_type: %v", err)
		}
	}

	if err = d.Set("access_network_uesa", flattenWirelessControllerHotspot20HsProfileAccessNetworkUesa(o["access-network-uesa"], d, "access_network_uesa")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-network-uesa"], "WirelessControllerHotspot20HsProfile-AccessNetworkUesa"); ok {
			if err = d.Set("access_network_uesa", vv); err != nil {
				return fmt.Errorf("Error reading access_network_uesa: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_network_uesa: %v", err)
		}
	}

	if err = d.Set("advice_of_charge", flattenWirelessControllerHotspot20HsProfileAdviceOfCharge(o["advice-of-charge"], d, "advice_of_charge")); err != nil {
		if vv, ok := fortiAPIPatch(o["advice-of-charge"], "WirelessControllerHotspot20HsProfile-AdviceOfCharge"); ok {
			if err = d.Set("advice_of_charge", vv); err != nil {
				return fmt.Errorf("Error reading advice_of_charge: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading advice_of_charge: %v", err)
		}
	}

	if err = d.Set("anqp_domain_id", flattenWirelessControllerHotspot20HsProfileAnqpDomainId(o["anqp-domain-id"], d, "anqp_domain_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["anqp-domain-id"], "WirelessControllerHotspot20HsProfile-AnqpDomainId"); ok {
			if err = d.Set("anqp_domain_id", vv); err != nil {
				return fmt.Errorf("Error reading anqp_domain_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anqp_domain_id: %v", err)
		}
	}

	if err = d.Set("bss_transition", flattenWirelessControllerHotspot20HsProfileBssTransition(o["bss-transition"], d, "bss_transition")); err != nil {
		if vv, ok := fortiAPIPatch(o["bss-transition"], "WirelessControllerHotspot20HsProfile-BssTransition"); ok {
			if err = d.Set("bss_transition", vv); err != nil {
				return fmt.Errorf("Error reading bss_transition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bss_transition: %v", err)
		}
	}

	if err = d.Set("conn_cap", flattenWirelessControllerHotspot20HsProfileConnCap(o["conn-cap"], d, "conn_cap")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn-cap"], "WirelessControllerHotspot20HsProfile-ConnCap"); ok {
			if err = d.Set("conn_cap", vv); err != nil {
				return fmt.Errorf("Error reading conn_cap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_cap: %v", err)
		}
	}

	if err = d.Set("deauth_request_timeout", flattenWirelessControllerHotspot20HsProfileDeauthRequestTimeout(o["deauth-request-timeout"], d, "deauth_request_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["deauth-request-timeout"], "WirelessControllerHotspot20HsProfile-DeauthRequestTimeout"); ok {
			if err = d.Set("deauth_request_timeout", vv); err != nil {
				return fmt.Errorf("Error reading deauth_request_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deauth_request_timeout: %v", err)
		}
	}

	if err = d.Set("dgaf", flattenWirelessControllerHotspot20HsProfileDgaf(o["dgaf"], d, "dgaf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dgaf"], "WirelessControllerHotspot20HsProfile-Dgaf"); ok {
			if err = d.Set("dgaf", vv); err != nil {
				return fmt.Errorf("Error reading dgaf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dgaf: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenWirelessControllerHotspot20HsProfileDomainName(o["domain-name"], d, "domain_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-name"], "WirelessControllerHotspot20HsProfile-DomainName"); ok {
			if err = d.Set("domain_name", vv); err != nil {
				return fmt.Errorf("Error reading domain_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("gas_comeback_delay", flattenWirelessControllerHotspot20HsProfileGasComebackDelay(o["gas-comeback-delay"], d, "gas_comeback_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["gas-comeback-delay"], "WirelessControllerHotspot20HsProfile-GasComebackDelay"); ok {
			if err = d.Set("gas_comeback_delay", vv); err != nil {
				return fmt.Errorf("Error reading gas_comeback_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gas_comeback_delay: %v", err)
		}
	}

	if err = d.Set("gas_fragmentation_limit", flattenWirelessControllerHotspot20HsProfileGasFragmentationLimit(o["gas-fragmentation-limit"], d, "gas_fragmentation_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["gas-fragmentation-limit"], "WirelessControllerHotspot20HsProfile-GasFragmentationLimit"); ok {
			if err = d.Set("gas_fragmentation_limit", vv); err != nil {
				return fmt.Errorf("Error reading gas_fragmentation_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gas_fragmentation_limit: %v", err)
		}
	}

	if err = d.Set("hessid", flattenWirelessControllerHotspot20HsProfileHessid(o["hessid"], d, "hessid")); err != nil {
		if vv, ok := fortiAPIPatch(o["hessid"], "WirelessControllerHotspot20HsProfile-Hessid"); ok {
			if err = d.Set("hessid", vv); err != nil {
				return fmt.Errorf("Error reading hessid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hessid: %v", err)
		}
	}

	if err = d.Set("ip_addr_type", flattenWirelessControllerHotspot20HsProfileIpAddrType(o["ip-addr-type"], d, "ip_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-addr-type"], "WirelessControllerHotspot20HsProfile-IpAddrType"); ok {
			if err = d.Set("ip_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading ip_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_addr_type: %v", err)
		}
	}

	if err = d.Set("l2tif", flattenWirelessControllerHotspot20HsProfileL2Tif(o["l2tif"], d, "l2tif")); err != nil {
		if vv, ok := fortiAPIPatch(o["l2tif"], "WirelessControllerHotspot20HsProfile-L2Tif"); ok {
			if err = d.Set("l2tif", vv); err != nil {
				return fmt.Errorf("Error reading l2tif: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l2tif: %v", err)
		}
	}

	if err = d.Set("nai_realm", flattenWirelessControllerHotspot20HsProfileNaiRealm(o["nai-realm"], d, "nai_realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["nai-realm"], "WirelessControllerHotspot20HsProfile-NaiRealm"); ok {
			if err = d.Set("nai_realm", vv); err != nil {
				return fmt.Errorf("Error reading nai_realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nai_realm: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerHotspot20HsProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerHotspot20HsProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("network_auth", flattenWirelessControllerHotspot20HsProfileNetworkAuth(o["network-auth"], d, "network_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-auth"], "WirelessControllerHotspot20HsProfile-NetworkAuth"); ok {
			if err = d.Set("network_auth", vv); err != nil {
				return fmt.Errorf("Error reading network_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_auth: %v", err)
		}
	}

	if err = d.Set("oper_friendly_name", flattenWirelessControllerHotspot20HsProfileOperFriendlyName(o["oper-friendly-name"], d, "oper_friendly_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["oper-friendly-name"], "WirelessControllerHotspot20HsProfile-OperFriendlyName"); ok {
			if err = d.Set("oper_friendly_name", vv); err != nil {
				return fmt.Errorf("Error reading oper_friendly_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oper_friendly_name: %v", err)
		}
	}

	if err = d.Set("oper_icon", flattenWirelessControllerHotspot20HsProfileOperIcon(o["oper-icon"], d, "oper_icon")); err != nil {
		if vv, ok := fortiAPIPatch(o["oper-icon"], "WirelessControllerHotspot20HsProfile-OperIcon"); ok {
			if err = d.Set("oper_icon", vv); err != nil {
				return fmt.Errorf("Error reading oper_icon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oper_icon: %v", err)
		}
	}

	if err = d.Set("osu_provider", flattenWirelessControllerHotspot20HsProfileOsuProvider(o["osu-provider"], d, "osu_provider")); err != nil {
		if vv, ok := fortiAPIPatch(o["osu-provider"], "WirelessControllerHotspot20HsProfile-OsuProvider"); ok {
			if err = d.Set("osu_provider", vv); err != nil {
				return fmt.Errorf("Error reading osu_provider: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading osu_provider: %v", err)
		}
	}

	if err = d.Set("osu_provider_nai", flattenWirelessControllerHotspot20HsProfileOsuProviderNai(o["osu-provider-nai"], d, "osu_provider_nai")); err != nil {
		if vv, ok := fortiAPIPatch(o["osu-provider-nai"], "WirelessControllerHotspot20HsProfile-OsuProviderNai"); ok {
			if err = d.Set("osu_provider_nai", vv); err != nil {
				return fmt.Errorf("Error reading osu_provider_nai: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading osu_provider_nai: %v", err)
		}
	}

	if err = d.Set("osu_ssid", flattenWirelessControllerHotspot20HsProfileOsuSsid(o["osu-ssid"], d, "osu_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["osu-ssid"], "WirelessControllerHotspot20HsProfile-OsuSsid"); ok {
			if err = d.Set("osu_ssid", vv); err != nil {
				return fmt.Errorf("Error reading osu_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading osu_ssid: %v", err)
		}
	}

	if err = d.Set("pame_bi", flattenWirelessControllerHotspot20HsProfilePameBi(o["pame-bi"], d, "pame_bi")); err != nil {
		if vv, ok := fortiAPIPatch(o["pame-bi"], "WirelessControllerHotspot20HsProfile-PameBi"); ok {
			if err = d.Set("pame_bi", vv); err != nil {
				return fmt.Errorf("Error reading pame_bi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pame_bi: %v", err)
		}
	}

	if err = d.Set("proxy_arp", flattenWirelessControllerHotspot20HsProfileProxyArp(o["proxy-arp"], d, "proxy_arp")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-arp"], "WirelessControllerHotspot20HsProfile-ProxyArp"); ok {
			if err = d.Set("proxy_arp", vv); err != nil {
				return fmt.Errorf("Error reading proxy_arp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_arp: %v", err)
		}
	}

	if err = d.Set("qos_map", flattenWirelessControllerHotspot20HsProfileQosMap(o["qos-map"], d, "qos_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-map"], "WirelessControllerHotspot20HsProfile-QosMap"); ok {
			if err = d.Set("qos_map", vv); err != nil {
				return fmt.Errorf("Error reading qos_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_map: %v", err)
		}
	}

	if err = d.Set("release", flattenWirelessControllerHotspot20HsProfileRelease(o["release"], d, "release")); err != nil {
		if vv, ok := fortiAPIPatch(o["release"], "WirelessControllerHotspot20HsProfile-Release"); ok {
			if err = d.Set("release", vv); err != nil {
				return fmt.Errorf("Error reading release: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading release: %v", err)
		}
	}

	if err = d.Set("roaming_consortium", flattenWirelessControllerHotspot20HsProfileRoamingConsortium(o["roaming-consortium"], d, "roaming_consortium")); err != nil {
		if vv, ok := fortiAPIPatch(o["roaming-consortium"], "WirelessControllerHotspot20HsProfile-RoamingConsortium"); ok {
			if err = d.Set("roaming_consortium", vv); err != nil {
				return fmt.Errorf("Error reading roaming_consortium: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading roaming_consortium: %v", err)
		}
	}

	if err = d.Set("terms_and_conditions", flattenWirelessControllerHotspot20HsProfileTermsAndConditions(o["terms-and-conditions"], d, "terms_and_conditions")); err != nil {
		if vv, ok := fortiAPIPatch(o["terms-and-conditions"], "WirelessControllerHotspot20HsProfile-TermsAndConditions"); ok {
			if err = d.Set("terms_and_conditions", vv); err != nil {
				return fmt.Errorf("Error reading terms_and_conditions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading terms_and_conditions: %v", err)
		}
	}

	if err = d.Set("venue_group", flattenWirelessControllerHotspot20HsProfileVenueGroup(o["venue-group"], d, "venue_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["venue-group"], "WirelessControllerHotspot20HsProfile-VenueGroup"); ok {
			if err = d.Set("venue_group", vv); err != nil {
				return fmt.Errorf("Error reading venue_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading venue_group: %v", err)
		}
	}

	if err = d.Set("venue_name", flattenWirelessControllerHotspot20HsProfileVenueName(o["venue-name"], d, "venue_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["venue-name"], "WirelessControllerHotspot20HsProfile-VenueName"); ok {
			if err = d.Set("venue_name", vv); err != nil {
				return fmt.Errorf("Error reading venue_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading venue_name: %v", err)
		}
	}

	if err = d.Set("venue_type", flattenWirelessControllerHotspot20HsProfileVenueType(o["venue-type"], d, "venue_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["venue-type"], "WirelessControllerHotspot20HsProfile-VenueType"); ok {
			if err = d.Set("venue_type", vv); err != nil {
				return fmt.Errorf("Error reading venue_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading venue_type: %v", err)
		}
	}

	if err = d.Set("venue_url", flattenWirelessControllerHotspot20HsProfileVenueUrl(o["venue-url"], d, "venue_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["venue-url"], "WirelessControllerHotspot20HsProfile-VenueUrl"); ok {
			if err = d.Set("venue_url", vv); err != nil {
				return fmt.Errorf("Error reading venue_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading venue_url: %v", err)
		}
	}

	if err = d.Set("wan_metrics", flattenWirelessControllerHotspot20HsProfileWanMetrics(o["wan-metrics"], d, "wan_metrics")); err != nil {
		if vv, ok := fortiAPIPatch(o["wan-metrics"], "WirelessControllerHotspot20HsProfile-WanMetrics"); ok {
			if err = d.Set("wan_metrics", vv); err != nil {
				return fmt.Errorf("Error reading wan_metrics: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wan_metrics: %v", err)
		}
	}

	if err = d.Set("wba_charging_currency", flattenWirelessControllerHotspot20HsProfileWbaChargingCurrency(o["wba-charging-currency"], d, "wba_charging_currency")); err != nil {
		if vv, ok := fortiAPIPatch(o["wba-charging-currency"], "WirelessControllerHotspot20HsProfile-WbaChargingCurrency"); ok {
			if err = d.Set("wba_charging_currency", vv); err != nil {
				return fmt.Errorf("Error reading wba_charging_currency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wba_charging_currency: %v", err)
		}
	}

	if err = d.Set("wba_charging_rate", flattenWirelessControllerHotspot20HsProfileWbaChargingRate(o["wba-charging-rate"], d, "wba_charging_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["wba-charging-rate"], "WirelessControllerHotspot20HsProfile-WbaChargingRate"); ok {
			if err = d.Set("wba_charging_rate", vv); err != nil {
				return fmt.Errorf("Error reading wba_charging_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wba_charging_rate: %v", err)
		}
	}

	if err = d.Set("wba_data_clearing_provider", flattenWirelessControllerHotspot20HsProfileWbaDataClearingProvider(o["wba-data-clearing-provider"], d, "wba_data_clearing_provider")); err != nil {
		if vv, ok := fortiAPIPatch(o["wba-data-clearing-provider"], "WirelessControllerHotspot20HsProfile-WbaDataClearingProvider"); ok {
			if err = d.Set("wba_data_clearing_provider", vv); err != nil {
				return fmt.Errorf("Error reading wba_data_clearing_provider: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wba_data_clearing_provider: %v", err)
		}
	}

	if err = d.Set("wba_financial_clearing_provider", flattenWirelessControllerHotspot20HsProfileWbaFinancialClearingProvider(o["wba-financial-clearing-provider"], d, "wba_financial_clearing_provider")); err != nil {
		if vv, ok := fortiAPIPatch(o["wba-financial-clearing-provider"], "WirelessControllerHotspot20HsProfile-WbaFinancialClearingProvider"); ok {
			if err = d.Set("wba_financial_clearing_provider", vv); err != nil {
				return fmt.Errorf("Error reading wba_financial_clearing_provider: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wba_financial_clearing_provider: %v", err)
		}
	}

	if err = d.Set("wba_open_roaming", flattenWirelessControllerHotspot20HsProfileWbaOpenRoaming(o["wba-open-roaming"], d, "wba_open_roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["wba-open-roaming"], "WirelessControllerHotspot20HsProfile-WbaOpenRoaming"); ok {
			if err = d.Set("wba_open_roaming", vv); err != nil {
				return fmt.Errorf("Error reading wba_open_roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wba_open_roaming: %v", err)
		}
	}

	if err = d.Set("wnm_sleep_mode", flattenWirelessControllerHotspot20HsProfileWnmSleepMode(o["wnm-sleep-mode"], d, "wnm_sleep_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["wnm-sleep-mode"], "WirelessControllerHotspot20HsProfile-WnmSleepMode"); ok {
			if err = d.Set("wnm_sleep_mode", vv); err != nil {
				return fmt.Errorf("Error reading wnm_sleep_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wnm_sleep_mode: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20HsProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20HsProfile3GppPlmn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileAccessNetworkAsra(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileAccessNetworkEsr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileAccessNetworkInternet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileAccessNetworkType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileAccessNetworkUesa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileAdviceOfCharge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileAnqpDomainId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileBssTransition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileConnCap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileDeauthRequestTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileDgaf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileDomainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileGasComebackDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileGasFragmentationLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileHessid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileIpAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileL2Tif(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileNaiRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileNetworkAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileOperFriendlyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileOperIcon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileOsuProvider(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileOsuProviderNai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileOsuSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfilePameBi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileProxyArp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileQosMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileRelease(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileRoamingConsortium(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileTermsAndConditions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileVenueGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileVenueName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileVenueType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileVenueUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileWanMetrics(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerHotspot20HsProfileWbaChargingCurrency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileWbaChargingRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileWbaDataClearingProvider(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileWbaFinancialClearingProvider(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileWbaOpenRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileWnmSleepMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20HsProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n3gpp_plmn"); ok || d.HasChange("n3gpp_plmn") {
		t, err := expandWirelessControllerHotspot20HsProfile3GppPlmn(d, v, "n3gpp_plmn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["3gpp-plmn"] = t
		}
	}

	if v, ok := d.GetOk("access_network_asra"); ok || d.HasChange("access_network_asra") {
		t, err := expandWirelessControllerHotspot20HsProfileAccessNetworkAsra(d, v, "access_network_asra")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-asra"] = t
		}
	}

	if v, ok := d.GetOk("access_network_esr"); ok || d.HasChange("access_network_esr") {
		t, err := expandWirelessControllerHotspot20HsProfileAccessNetworkEsr(d, v, "access_network_esr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-esr"] = t
		}
	}

	if v, ok := d.GetOk("access_network_internet"); ok || d.HasChange("access_network_internet") {
		t, err := expandWirelessControllerHotspot20HsProfileAccessNetworkInternet(d, v, "access_network_internet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-internet"] = t
		}
	}

	if v, ok := d.GetOk("access_network_type"); ok || d.HasChange("access_network_type") {
		t, err := expandWirelessControllerHotspot20HsProfileAccessNetworkType(d, v, "access_network_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-type"] = t
		}
	}

	if v, ok := d.GetOk("access_network_uesa"); ok || d.HasChange("access_network_uesa") {
		t, err := expandWirelessControllerHotspot20HsProfileAccessNetworkUesa(d, v, "access_network_uesa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-uesa"] = t
		}
	}

	if v, ok := d.GetOk("advice_of_charge"); ok || d.HasChange("advice_of_charge") {
		t, err := expandWirelessControllerHotspot20HsProfileAdviceOfCharge(d, v, "advice_of_charge")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advice-of-charge"] = t
		}
	}

	if v, ok := d.GetOk("anqp_domain_id"); ok || d.HasChange("anqp_domain_id") {
		t, err := expandWirelessControllerHotspot20HsProfileAnqpDomainId(d, v, "anqp_domain_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anqp-domain-id"] = t
		}
	}

	if v, ok := d.GetOk("bss_transition"); ok || d.HasChange("bss_transition") {
		t, err := expandWirelessControllerHotspot20HsProfileBssTransition(d, v, "bss_transition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-transition"] = t
		}
	}

	if v, ok := d.GetOk("conn_cap"); ok || d.HasChange("conn_cap") {
		t, err := expandWirelessControllerHotspot20HsProfileConnCap(d, v, "conn_cap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-cap"] = t
		}
	}

	if v, ok := d.GetOk("deauth_request_timeout"); ok || d.HasChange("deauth_request_timeout") {
		t, err := expandWirelessControllerHotspot20HsProfileDeauthRequestTimeout(d, v, "deauth_request_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deauth-request-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dgaf"); ok || d.HasChange("dgaf") {
		t, err := expandWirelessControllerHotspot20HsProfileDgaf(d, v, "dgaf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dgaf"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok || d.HasChange("domain_name") {
		t, err := expandWirelessControllerHotspot20HsProfileDomainName(d, v, "domain_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("gas_comeback_delay"); ok || d.HasChange("gas_comeback_delay") {
		t, err := expandWirelessControllerHotspot20HsProfileGasComebackDelay(d, v, "gas_comeback_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-comeback-delay"] = t
		}
	}

	if v, ok := d.GetOk("gas_fragmentation_limit"); ok || d.HasChange("gas_fragmentation_limit") {
		t, err := expandWirelessControllerHotspot20HsProfileGasFragmentationLimit(d, v, "gas_fragmentation_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-fragmentation-limit"] = t
		}
	}

	if v, ok := d.GetOk("hessid"); ok || d.HasChange("hessid") {
		t, err := expandWirelessControllerHotspot20HsProfileHessid(d, v, "hessid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hessid"] = t
		}
	}

	if v, ok := d.GetOk("ip_addr_type"); ok || d.HasChange("ip_addr_type") {
		t, err := expandWirelessControllerHotspot20HsProfileIpAddrType(d, v, "ip_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("l2tif"); ok || d.HasChange("l2tif") {
		t, err := expandWirelessControllerHotspot20HsProfileL2Tif(d, v, "l2tif")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2tif"] = t
		}
	}

	if v, ok := d.GetOk("nai_realm"); ok || d.HasChange("nai_realm") {
		t, err := expandWirelessControllerHotspot20HsProfileNaiRealm(d, v, "nai_realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nai-realm"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerHotspot20HsProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("network_auth"); ok || d.HasChange("network_auth") {
		t, err := expandWirelessControllerHotspot20HsProfileNetworkAuth(d, v, "network_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-auth"] = t
		}
	}

	if v, ok := d.GetOk("oper_friendly_name"); ok || d.HasChange("oper_friendly_name") {
		t, err := expandWirelessControllerHotspot20HsProfileOperFriendlyName(d, v, "oper_friendly_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oper-friendly-name"] = t
		}
	}

	if v, ok := d.GetOk("oper_icon"); ok || d.HasChange("oper_icon") {
		t, err := expandWirelessControllerHotspot20HsProfileOperIcon(d, v, "oper_icon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oper-icon"] = t
		}
	}

	if v, ok := d.GetOk("osu_provider"); ok || d.HasChange("osu_provider") {
		t, err := expandWirelessControllerHotspot20HsProfileOsuProvider(d, v, "osu_provider")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osu-provider"] = t
		}
	}

	if v, ok := d.GetOk("osu_provider_nai"); ok || d.HasChange("osu_provider_nai") {
		t, err := expandWirelessControllerHotspot20HsProfileOsuProviderNai(d, v, "osu_provider_nai")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osu-provider-nai"] = t
		}
	}

	if v, ok := d.GetOk("osu_ssid"); ok || d.HasChange("osu_ssid") {
		t, err := expandWirelessControllerHotspot20HsProfileOsuSsid(d, v, "osu_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osu-ssid"] = t
		}
	}

	if v, ok := d.GetOk("pame_bi"); ok || d.HasChange("pame_bi") {
		t, err := expandWirelessControllerHotspot20HsProfilePameBi(d, v, "pame_bi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pame-bi"] = t
		}
	}

	if v, ok := d.GetOk("proxy_arp"); ok || d.HasChange("proxy_arp") {
		t, err := expandWirelessControllerHotspot20HsProfileProxyArp(d, v, "proxy_arp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-arp"] = t
		}
	}

	if v, ok := d.GetOk("qos_map"); ok || d.HasChange("qos_map") {
		t, err := expandWirelessControllerHotspot20HsProfileQosMap(d, v, "qos_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-map"] = t
		}
	}

	if v, ok := d.GetOk("release"); ok || d.HasChange("release") {
		t, err := expandWirelessControllerHotspot20HsProfileRelease(d, v, "release")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["release"] = t
		}
	}

	if v, ok := d.GetOk("roaming_consortium"); ok || d.HasChange("roaming_consortium") {
		t, err := expandWirelessControllerHotspot20HsProfileRoamingConsortium(d, v, "roaming_consortium")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roaming-consortium"] = t
		}
	}

	if v, ok := d.GetOk("terms_and_conditions"); ok || d.HasChange("terms_and_conditions") {
		t, err := expandWirelessControllerHotspot20HsProfileTermsAndConditions(d, v, "terms_and_conditions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["terms-and-conditions"] = t
		}
	}

	if v, ok := d.GetOk("venue_group"); ok || d.HasChange("venue_group") {
		t, err := expandWirelessControllerHotspot20HsProfileVenueGroup(d, v, "venue_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["venue-group"] = t
		}
	}

	if v, ok := d.GetOk("venue_name"); ok || d.HasChange("venue_name") {
		t, err := expandWirelessControllerHotspot20HsProfileVenueName(d, v, "venue_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["venue-name"] = t
		}
	}

	if v, ok := d.GetOk("venue_type"); ok || d.HasChange("venue_type") {
		t, err := expandWirelessControllerHotspot20HsProfileVenueType(d, v, "venue_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["venue-type"] = t
		}
	}

	if v, ok := d.GetOk("venue_url"); ok || d.HasChange("venue_url") {
		t, err := expandWirelessControllerHotspot20HsProfileVenueUrl(d, v, "venue_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["venue-url"] = t
		}
	}

	if v, ok := d.GetOk("wan_metrics"); ok || d.HasChange("wan_metrics") {
		t, err := expandWirelessControllerHotspot20HsProfileWanMetrics(d, v, "wan_metrics")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-metrics"] = t
		}
	}

	if v, ok := d.GetOk("wba_charging_currency"); ok || d.HasChange("wba_charging_currency") {
		t, err := expandWirelessControllerHotspot20HsProfileWbaChargingCurrency(d, v, "wba_charging_currency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wba-charging-currency"] = t
		}
	}

	if v, ok := d.GetOk("wba_charging_rate"); ok || d.HasChange("wba_charging_rate") {
		t, err := expandWirelessControllerHotspot20HsProfileWbaChargingRate(d, v, "wba_charging_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wba-charging-rate"] = t
		}
	}

	if v, ok := d.GetOk("wba_data_clearing_provider"); ok || d.HasChange("wba_data_clearing_provider") {
		t, err := expandWirelessControllerHotspot20HsProfileWbaDataClearingProvider(d, v, "wba_data_clearing_provider")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wba-data-clearing-provider"] = t
		}
	}

	if v, ok := d.GetOk("wba_financial_clearing_provider"); ok || d.HasChange("wba_financial_clearing_provider") {
		t, err := expandWirelessControllerHotspot20HsProfileWbaFinancialClearingProvider(d, v, "wba_financial_clearing_provider")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wba-financial-clearing-provider"] = t
		}
	}

	if v, ok := d.GetOk("wba_open_roaming"); ok || d.HasChange("wba_open_roaming") {
		t, err := expandWirelessControllerHotspot20HsProfileWbaOpenRoaming(d, v, "wba_open_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wba-open-roaming"] = t
		}
	}

	if v, ok := d.GetOk("wnm_sleep_mode"); ok || d.HasChange("wnm_sleep_mode") {
		t, err := expandWirelessControllerHotspot20HsProfileWnmSleepMode(d, v, "wnm_sleep_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wnm-sleep-mode"] = t
		}
	}

	return &obj, nil
}
