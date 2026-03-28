---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_antivirus_profile"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure AntiVirus profiles.
---

# fmgdevice_antivirus_profile
<i>This object will be purged after policy copy and install.</i> Configure AntiVirus profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `cifs`: `fmgdevice_antivirus_profile_cifs`
>- `content_disarm`: `fmgdevice_antivirus_profile_contentdisarm`
>- `ftp`: `fmgdevice_antivirus_profile_ftp`
>- `http`: `fmgdevice_antivirus_profile_http`
>- `imap`: `fmgdevice_antivirus_profile_imap`
>- `mapi`: `fmgdevice_antivirus_profile_mapi`
>- `nac_quar`: `fmgdevice_antivirus_profile_nacquar`
>- `nntp`: `fmgdevice_antivirus_profile_nntp`
>- `outbreak_prevention`: `fmgdevice_antivirus_profile_outbreakprevention`
>- `pop3`: `fmgdevice_antivirus_profile_pop3`
>- `smtp`: `fmgdevice_antivirus_profile_smtp`
>- `ssh`: `fmgdevice_antivirus_profile_ssh`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `analytics_bl_filetype` - Only submit files matching this DLP file-pattern to FortiSandbox.
* `analytics_accept_filetype` - Only submit files matching this DLP file-pattern to FortiSandbox (post-transfer scan only).
* `analytics_db` - Enable/disable using the FortiSandbox signature database to supplement the AV signature databases. Valid values: `disable`, `enable`.

* `analytics_ignore_filetype` - Do not submit files matching this DLP file-pattern to FortiSandbox (post-transfer scan only).
* `av_virus_log` - Enable/disable AntiVirus logging. Valid values: `disable`, `enable`.

* `cifs` - Cifs. The structure of `cifs` block is documented below.
* `comment` - Comment.
* `content_disarm` - Content-Disarm. The structure of `content_disarm` block is documented below.
* `ems_threat_feed` - Enable/disable use of EMS threat feed when performing AntiVirus scan. Analyzes files including the content of archives. Valid values: `disable`, `enable`.

* `extended_log` - Enable/disable extended logging for antivirus. Valid values: `disable`, `enable`.

* `external_blocklist` - One or more external malware block lists.
* `external_blocklist_archive_scan` - External-Blocklist-Archive-Scan. Valid values: `disable`, `enable`.

* `external_blocklist_enable_all` - Enable/disable all external blocklists. Valid values: `disable`, `enable`.

* `feature_set` - Flow/proxy feature set. Valid values: `proxy`, `flow`.

* `fortiai_error_action` - Action to take if FortiAI encounters an error. Valid values: `block`, `log-only`, `ignore`.

* `fortiai_timeout_action` - Fortiai-Timeout-Action. Valid values: `block`, `log-only`, `ignore`.

* `fortindr_error_action` - Action to take if FortiNDR encounters an error. Valid values: `log-only`, `block`, `ignore`.

* `fortindr_timeout_action` - Action to take if FortiNDR encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.

* `fortisandbox_error_action` - Action to take if FortiSandbox inline scan encounters an error. Valid values: `log-only`, `block`, `ignore`.

* `fortisandbox_max_upload` - Maximum size of files that can be uploaded to FortiSandbox in Mbytes.
* `fortisandbox_mode` - FortiSandbox scan modes. Valid values: `inline`, `analytics-suspicious`, `analytics-everything`.

* `fortisandbox_timeout_action` - Action to take if FortiSandbox inline scan encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.

* `ftp` - Ftp. The structure of `ftp` block is documented below.
* `http` - Http. The structure of `http` block is documented below.
* `imap` - Imap. The structure of `imap` block is documented below.
* `mapi` - Mapi. The structure of `mapi` block is documented below.
* `mobile_malware_db` - Enable/disable using the mobile malware signature database. Valid values: `disable`, `enable`.

* `nac_quar` - Nac-Quar. The structure of `nac_quar` block is documented below.
* `name` - Profile name.
* `nntp` - Nntp. The structure of `nntp` block is documented below.
* `outbreak_prevention` - Outbreak-Prevention. The structure of `outbreak_prevention` block is documented below.
* `outbreak_prevention_archive_scan` - Enable/disable outbreak-prevention archive scanning. Valid values: `disable`, `enable`.

* `pop3` - Pop3. The structure of `pop3` block is documented below.
* `replacemsg_group` - Replacement message group customized for this profile.
* `scan_mode` - Configure scan mode (default or legacy). Valid values: `legacy`, `default`.

* `smtp` - Smtp. The structure of `smtp` block is documented below.
* `ssh` - Ssh. The structure of `ssh` block is documented below.
* `analytics_max_upload` - Analytics-Max-Upload.
* `analytics_wl_filetype` - Do not submit files matching this DLP file-pattern to FortiSandbox.
* `ftgd_analytics` - Ftgd-Analytics. Valid values: `disable`, `suspicious`, `everything`.

* `av_block_log` - Av-Block-Log. Valid values: `disable`, `enable`.


The `cifs` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `options` - Enable/disable CIFS AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `quarantine`, `avmonitor`.

* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.

* `fortisandbox` - Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.

* `malware_stream` - Enable 0-day malware-stream scanning. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `content_disarm` block supports:

* `analytics_suspicious` - Enable/disable using CDR as a secondary method for determining suspicous files for analytics. Valid values: `disable`, `enable`.

* `cover_page` - Enable/disable inserting a cover page into the disarmed document. Valid values: `disable`, `enable`.

* `detect_only` - Enable/disable only detect disarmable files, do not alter content. Valid values: `disable`, `enable`.

* `error_action` - Action to be taken if CDR engine encounters an unrecoverable error. Valid values: `block`, `log-only`, `ignore`.

* `office_action` - Enable/disable stripping of PowerPoint action events in Microsoft Office documents. Valid values: `disable`, `enable`.

* `office_dde` - Enable/disable stripping of Dynamic Data Exchange events in Microsoft Office documents. Valid values: `disable`, `enable`.

* `office_embed` - Enable/disable stripping of embedded objects in Microsoft Office documents. Valid values: `disable`, `enable`.

* `office_hylink` - Enable/disable stripping of hyperlinks in Microsoft Office documents. Valid values: `disable`, `enable`.

* `office_linked` - Enable/disable stripping of linked objects in Microsoft Office documents. Valid values: `disable`, `enable`.

* `office_macro` - Enable/disable stripping of macros in Microsoft Office documents. Valid values: `disable`, `enable`.

* `original_file_destination` - Destination to send original file if active content is removed. Valid values: `fortisandbox`, `quarantine`, `discard`.

* `pdf_act_form` - Enable/disable stripping of PDF document actions that submit data to other targets. Valid values: `disable`, `enable`.

* `pdf_act_gotor` - Enable/disable stripping of PDF document actions that access other PDF documents. Valid values: `disable`, `enable`.

* `pdf_act_java` - Enable/disable stripping of PDF document actions that execute JavaScript code. Valid values: `disable`, `enable`.

* `pdf_act_launch` - Enable/disable stripping of PDF document actions that launch other applications. Valid values: `disable`, `enable`.

* `pdf_act_movie` - Enable/disable stripping of PDF document actions that play a movie. Valid values: `disable`, `enable`.

* `pdf_act_sound` - Enable/disable stripping of PDF document actions that play a sound. Valid values: `disable`, `enable`.

* `pdf_embedfile` - Enable/disable stripping of embedded files in PDF documents. Valid values: `disable`, `enable`.

* `pdf_hyperlink` - Enable/disable stripping of hyperlinks from PDF documents. Valid values: `disable`, `enable`.

* `pdf_javacode` - Enable/disable stripping of JavaScript code in PDF documents. Valid values: `disable`, `enable`.


The `ftp` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `options` - Enable/disable FTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `quarantine`, `avmonitor`.

* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.

* `fortisandbox` - Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.

* `malware_stream` - Enable 0-day malware-stream scanning. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `http` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan. Valid values: `disable`, `enable`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `options` - Enable/disable HTTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `quarantine`, `avmonitor`.

* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.

* `fortisandbox` - Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.

* `malware_stream` - Enable 0-day malware-stream scanning. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.

* `unknown_content_encoding` - Configure the action the FortiGate unit will take on unknown content-encoding. Valid values: `block`, `inspect`, `bypass`.


The `imap` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan. Valid values: `disable`, `enable`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.

* `options` - Enable/disable IMAP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `quarantine`, `avmonitor`.

* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.

* `fortisandbox` - Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.

* `malware_stream` - Enable 0-day malware-stream scanning. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `mapi` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.

* `options` - Enable/disable MAPI AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `quarantine`, `avmonitor`.

* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.

* `fortisandbox` - Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.

* `malware_stream` - Enable 0-day malware-stream scanning. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `nac_quar` block supports:

* `expiry` - Duration of quarantine.
* `infected` - Enable/Disable quarantining infected hosts to the banned user list. Valid values: `none`, `quar-src-ip`.

* `log` - Enable/disable AntiVirus quarantine logging. Valid values: `disable`, `enable`.


The `nntp` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `options` - Enable/disable NNTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `quarantine`, `avmonitor`.

* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.

* `fortisandbox` - Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.

* `malware_stream` - Enable 0-day malware-stream scanning. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `outbreak_prevention` block supports:

* `external_blocklist` - Enable/disable external malware blocklist. Valid values: `disable`, `enable`.

* `ftgd_service` - Enable/disable FortiGuard Virus outbreak prevention service. Valid values: `disable`, `enable`.


The `pop3` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan. Valid values: `disable`, `enable`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.

* `options` - Enable/disable POP3 AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `quarantine`, `avmonitor`.

* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.

* `fortisandbox` - Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.

* `malware_stream` - Enable 0-day malware-stream scanning. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `smtp` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan. Valid values: `disable`, `enable`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.

* `options` - Enable/disable SMTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `quarantine`, `avmonitor`.

* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.

* `fortisandbox` - Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.

* `malware_stream` - Enable 0-day malware-stream scanning. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.


The `ssh` block supports:

* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `multipart`, `nested`, `mailbomb`, `unhandled`, `partiallycorrupted`, `timeout`.

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `monitor`, `block`.

* `emulator` - Enable/disable the virus emulator. Valid values: `disable`, `enable`.

* `options` - Enable/disable SFTP and SCP AntiVirus scanning, monitoring, and quarantine. Valid values: `avmonitor`, `quarantine`, `scan`.

* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `monitor`, `block`.

* `fortindr` - Enable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.

* `fortisandbox` - Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.

* `malware_stream` - Enable 0-day malware-stream scanning. Analyzes files including the content of archives. Valid values: `disable`, `monitor`, `block`.

* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable`, `block`, `monitor`.

* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Antivirus Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_antivirus_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

