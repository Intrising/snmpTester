package main

import (
	"./utils"
	"fmt"
)

const (
	check            = "\u2714"
	cross            = "\u274c"
	nothingInThisOID = "No Such Instance currently exists at this OID"

	snmpGet  = "snmpget -v 3 "
	snmpWalk = "snmpwalk -v 3 "
	snmpSet  = "snmpset -v 3 "

	snmpUser              = "-u walter "
	snmpSecurityLevel     = "-l authPriv "
	snmpAuthentication    = "-a MD5 "
	snmpAuthPassPhrase    = "-A 123456789 "
	snmpPrivateProtocol   = "-x DES "
	snmpPrivatePassPhrase = "-X 123456789 "

	deviceIP = "192.168.15.10 "

	snmpGetPrefix     = snmpGet + snmpUser + snmpSecurityLevel + snmpAuthentication + snmpAuthPassPhrase + snmpPrivateProtocol + snmpPrivatePassPhrase + deviceIP
	snmpWalkPrefix    = snmpWalk + snmpUser + snmpSecurityLevel + snmpAuthentication + snmpAuthPassPhrase + snmpPrivateProtocol + snmpPrivatePassPhrase + deviceIP
	snmpSetPrefix     = snmpSet + snmpUser + snmpSecurityLevel + snmpAuthentication + snmpAuthPassPhrase + snmpPrivateProtocol + snmpPrivatePassPhrase + deviceIP
	snmpSetTypeString = "string"
	snmpSetTypeInt    = "integer"
	snmpSetTypeIpaddr = "ipaddress"
	oidPrefix         = "1.3.6.1.4.1.37072.302.2.3."

	// no postfix means snmpwalk commnad
	// all command can be read so no postfix for read-only
	// -w  means read-write
	// @s  means type string
	// @i  means type integer
	// @ip means type ipaddress

	// SYSTEM (1)
	systemName           = oidPrefix + "1.1.0@s-w"
	systemLocation       = oidPrefix + "1.2.0@s-w"
	systemContact        = oidPrefix + "1.3.0@s-w"
	systemDescr          = oidPrefix + "1.4.0@s-w"
	systemFwVersion      = oidPrefix + "1.5.0@s"
	systemMacaddress     = oidPrefix + "1.6.0@s"
	systemAutoLogoutTime = oidPrefix + "1.7.0@i-w"
	systemSerialNum      = oidPrefix + "1.8.0@i"

	// Setting(2)
	vlanWalk                 = oidPrefix + "2" // VLan (2.1)
	vlanPortCfgNum           = oidPrefix + "2.1.1.1.1@i"
	vlanMembers              = oidPrefix + "2.1.1.1.2@s"
	vlanTags                 = oidPrefix + "2.1.1.1.3@s"
	pvidCfgNum               = oidPrefix + "2.1.2.1.1@i"
	vlanPvid                 = oidPrefix + "2.1.2.1.2@i-w"
	vlanFrameType            = oidPrefix + "2.1.2.1.3@i-w"
	mvrCfgNum                = oidPrefix + "2.2.1.1.1@i" // Mvr (2.2)
	mvrCfgVid                = oidPrefix + "2.2.1.1.2@i"
	mvrIPAddr                = oidPrefix + "2.2.1.1.3@ip"
	mvrMemnters              = oidPrefix + "2.2.1.1.4@s"
	igmpEnableQuerier        = oidPrefix + "2.3.1.0@s-w" // Igmp (2.3)
	igmpQuerierVersion       = oidPrefix + "2.3.2.0@s-w"
	igmpEnableSnooping       = oidPrefix + "2.3.3.0@s-w"
	igmpEnableFloodWellKnown = oidPrefix + "2.3.4.0@i-w"
	igmpPortNum              = oidPrefix + "2.3.5.1.1@i" // IgmpRouterTable (2.3.5)
	igmpRouterStatus         = oidPrefix + "2.3.5.1.2@i-w"
	igmpFastLeaveStatus      = oidPrefix + "2.3.5.1.3@i-w"
	igmpVidNum               = oidPrefix + "2.3.6.1.1@i" // IgmpStatisticsTable (2.3.6)
	igmpStatusQuerier        = oidPrefix + "2.3.6.1.2@s"
	igmpQuerierTx            = oidPrefix + "2.3.6.1.3@i"
	igmpQuerierRx            = oidPrefix + "2.3.6.1.4@i"
	igmpV1Rx                 = oidPrefix + "2.3.6.1.5@i"
	igmpV2Rx                 = oidPrefix + "2.3.6.1.6@i"
	igmpV3Rx                 = oidPrefix + "2.3.6.1.7@i"
	igmpV2Leave              = oidPrefix + "2.3.6.1.8@i"
	igmpEntriesEntryIndex    = oidPrefix + "2.3.7.1.1@i" // IgmpEntriesTable (2.3.7)
	igmpEntriesEntryIPAddr   = oidPrefix + "2.3.7.1.2@ip"
	igmpEntriesEntryVID      = oidPrefix + "2.3.7.1.3@i"
	igmpEntriesEntryMembers  = oidPrefix + "2.3.7.1.4@s"

	// Status (3)
	lldpPortNum     = oidPrefix + "3.1.1.1.1@i" // LLDPInfo (3.1)
	lldpInfoContent = oidPrefix + "3.1.1.1.2@i"

	// Warning (11)
	faultAlarmPowerCfgNum        = oidPrefix + "11.1.1.1.1@i" // FaultAlarm (11.1)
	faultAlarmPowerStatus        = oidPrefix + "11.1.1.1.2@i-w"
	faultAlarmPortCfgNum         = oidPrefix + "11.1.2.1.1@i"
	faultAlarmPortLinkStatus     = oidPrefix + "11.1.2.1.2@i-w"
	eventDDMEnabled              = oidPrefix + "11.2.1.1.0@i-w" // EventDDMEnabled (11.2.1)
	eventDDMTemperatureLower     = oidPrefix + "11.2.1.2.0@s-w"
	eventDDMTemperatureUpper     = oidPrefix + "11.2.1.3.0@s-w"
	eventDDMVoltageLower         = oidPrefix + "11.2.1.4.0@s-w"
	eventDDMVoltageUpper         = oidPrefix + "11.2.1.5.0@s-w"
	eventDDMTxBiasLower          = oidPrefix + "11.2.1.6.0@s-w"
	eventDDMTTxBiasUpper         = oidPrefix + "11.2.1.7.0@s-w"
	eventDDMTxPowerLower         = oidPrefix + "11.2.1.8.0@s-w"
	eventDDMTxPowerUpper         = oidPrefix + "11.2.1.9.0@s-w"
	eventDDMRxPowerLower         = oidPrefix + "11.2.1.10.0@s-w"
	eventDDMRxPowerUpper         = oidPrefix + "11.2.1.11.0@s-w"
	eventMonitorEnabled          = oidPrefix + "11.2.2.1.0@i-w" // EventMonitor (11.2.1)
	eventMonitorTemperatureLower = oidPrefix + "11.2.2.2.0@s-w"
	eventMonitorTemperatureUpper = oidPrefix + "11.2.2.3.0@s-w"
	eventMonitorVoltageLower     = oidPrefix + "11.2.2.4.0@s-w"
	eventMonitorVoltageUpper     = oidPrefix + "11.2.2.5.0@s-w"
	eventMonitorCurrentLower     = oidPrefix + "11.2.2.6.0@s-w"
	eventMonitorCurrentUpper     = oidPrefix + "11.2.2.7.0@s-w"
	eventMonitorPowerLower       = oidPrefix + "11.2.2.8.0@s-w"
	eventMonitorPowerUpper       = oidPrefix + "11.2.2.9.0@s-w"
	eventPOEAPortCfgNum          = oidPrefix + "11.2.3.1.1.1@i" // EventPOEA (11.2.3)
	eventPOEAPingEnabled         = oidPrefix + "11.2.3.1.1.2@i-w"
	eventPOEAPingIPAddr          = oidPrefix + "11.2.3.1.1.3@ip-w"
	eventPOEAPingInterval        = oidPrefix + "11.2.3.1.1.4@i-w"
	eventPOEAPingRetry           = oidPrefix + "11.2.3.1.1.5@i-w"
	eventPOEAPingReboot          = oidPrefix + "11.2.3.1.1.6@i-w"
	eventPOEAPingFailAction      = oidPrefix + "11.2.3.1.1.7@i-w"
	localLogEnable               = oidPrefix + "11.3.1.1.0@i-w" // ActionConfiguration (11.3)
	remoteSystemLogCfgNum        = oidPrefix + "11.3.2.1.1.1@i" // RemoteSystemLog (11.3.2)
	remoteSystemLogHost          = oidPrefix + "11.3.2.1.1.2@ip"
	remoteSystemLogTag           = oidPrefix + "11.3.2.1.1.3@s"
	remoteSystemLogFacility      = oidPrefix + "11.3.2.1.1.4@s"
	remoteSystemLogHostName      = oidPrefix + "11.3.2.1.1.5@s"
	emailEnable                  = oidPrefix + "11.3.3.1.0@i-w"   // email (11.3.3.1)
	emailServerUser              = oidPrefix + "11.3.3.2.1.0@s-w" // emailServer (11.3.3.2)
	emailServerPassword          = oidPrefix + "11.3.3.2.2.0@s-w"
	emailServerHost              = oidPrefix + "11.3.3.2.3.0@s-w"
	emailServerSSLEnable         = oidPrefix + "11.3.3.2.4.0@i-w"
	emailSender                  = oidPrefix + "11.3.3.3.0@s-w"
	emailSubject                 = oidPrefix + "11.3.3.4.0@-w"
	emailCloudEnable             = oidPrefix + "11.3.3.5.0@i-w"
	emailReceiverCfgNum          = oidPrefix + "11.3.3.6.1.1@i"
	emailReceiverHost            = oidPrefix + "11.3.3.6.1.2@s"
	smsEnable                    = oidPrefix + "11.3.4.1.0@i-w" // SMS (11.3.4)
	smsUser                      = oidPrefix + "11.3.4.2.0@s-w"
	smsPassword                  = oidPrefix + "11.3.4.3.0@s-w"
	smsSenderText                = oidPrefix + "11.3.4.4.0@s-w"
	smsReceiverCfgNum            = oidPrefix + "11.3.4.5.1.1@i"
	smsReceiverPhone             = oidPrefix + "11.3.4.5.1.2@s"
	snmpResponseLocale           = oidPrefix + "11.3.5.1.1.0@i-w" // Snmp (11.3.5)
	snmpCommunityCfgNum          = oidPrefix + "11.3.5.1.2.1.1@i"
	snmpCommunityCfgString       = oidPrefix + "11.3.5.1.2.1.2@s"
	snmpCommunityCfgReadOnly     = oidPrefix + "11.3.5.1.2.1.3@i"
	snmpTrapCfgNum               = oidPrefix + "11.3.5.2.1.1.1@i" // Trap (11.3.5.2)
	snmpTrapCfgCommunity         = oidPrefix + "11.3.5.2.1.1.2@s"
	snmpTrapCfgIPAddress         = oidPrefix + "11.3.5.2.1.1.3@ip"
	snmpTrapCfgVersion           = oidPrefix + "11.3.5.2.1.1.4@i"
	snmpV3UserCfgNum             = oidPrefix + "11.3.5.3.1.1.1@i" // V3User (11.3.5.3)
	snmpV3UserCfgName            = oidPrefix + "11.3.5.3.1.1.2@s"
	snmpV3UserCfgSecurityLevel   = oidPrefix + "11.3.5.3.1.1.3@i"
	snmpV3UserCfgAuthProtocal    = oidPrefix + "11.3.5.3.1.1.4@i"
	snmpV3UserCfgAuthPassword    = oidPrefix + "11.3.5.3.1.1.5@s"
	snmpV3UserCfgPrivProtocal    = oidPrefix + "11.3.5.3.1.1.6@s"
	snmpV3UserCfgPrivPassword    = oidPrefix + "11.3.5.3.1.1.7@s"
	doutCfgNum                   = oidPrefix + "11.3.6.1.1.1@i" // Dout (11.3.6)
	doutCfgEnable                = oidPrefix + "11.3.6.1.1.2@i-w"
	doutCfgAction                = oidPrefix + "11.3.6.1.1.3@i-w"
	deviceBootEvent              = oidPrefix + "11.4.1.1.0@i-w" // EventActionMap (11.4.1.1)
	authenticationFailureEvent   = oidPrefix + "11.4.1.2.0@i-w"
	authenticationSuccessEvent   = oidPrefix + "11.4.1.3.0@i-w"
	deviceDDMEvent               = oidPrefix + "11.4.1.4.0@i-w"
	devicePOEEvent               = oidPrefix + "11.4.1.5.0@i-w"
	devicePOEBEvent              = oidPrefix + "11.4.1.6.0@i-w"
	ringTopologyChangeEvent      = oidPrefix + "11.4.1.7.0@i-w"
	envMonitorEvent              = oidPrefix + "11.4.1.8.0@i-w"
	eventPortNumber              = oidPrefix + "11.4.2.1.1.1@i" // PortsEvent (11.4.2)
	eventPortEventLog            = oidPrefix + "11.4.2.1.1.2@i-w"
	eventPortEventsms            = oidPrefix + "11.4.2.1.1.3@i-w"
	eventPortEventSMTP           = oidPrefix + "11.4.2.1.1.4@i-w"
	eventPortEventsnmpTRAP       = oidPrefix + "11.4.2.1.1.5@i-w"
	eventPortEventdout1          = oidPrefix + "11.4.2.1.1.6@i-w"
	eventPortEventdout2          = oidPrefix + "11.4.2.1.1.7@i-w"
	eventPowerNumber             = oidPrefix + "11.4.3.1.1.1@i" // PowerEvent (11.4.3)
	eventPowerEventLog           = oidPrefix + "11.4.3.1.1.2@i-w"
	eventPowerEventsms           = oidPrefix + "11.4.3.1.1.3@i-w"
	eventPowerEventSMTP          = oidPrefix + "11.4.3.1.1.4@i-w"
	eventPowerEventsnmpTRAP      = oidPrefix + "11.4.3.1.1.5@i-w"
	eventPowerEventdout1         = oidPrefix + "11.4.3.1.1.6@i-w"
	eventPowerEventdout2         = oidPrefix + "11.4.3.1.1.7@i-w"
	eventDiNumber                = oidPrefix + "11.4.4.1.1.1@i" // DiEvent (11.4.4)
	eventDiEventLog              = oidPrefix + "11.4.4.1.1.2@i-w"
	eventDiEventsms              = oidPrefix + "11.4.4.1.1.3@i-w"
	eventDiEventSMTP             = oidPrefix + "11.4.4.1.1.4@i-w"
	eventDiEventsnmpTRAP         = oidPrefix + "11.4.4.1.1.5@i-w"
	eventDiEventdout1            = oidPrefix + "11.4.4.1.1.6@i-w"
	eventDiEventdout2            = oidPrefix + "11.4.4.1.1.7@i-w"

	// Monitoring (12)
	envVoltage                     = oidPrefix + "12.1.1.0@s" // ENVMonitor (12.1)
	envCurrent                     = oidPrefix + "12.1.2.0@s"
	envWalt                        = oidPrefix + "12.1.3.0@s"
	envTemperature                 = oidPrefix + "12.1.4.0@s"
	ddmPortNumber                  = oidPrefix + "12.2.1.1.1@i" // DDM (12.2)
	ddmTemperatureHighAlarm        = oidPrefix + "12.2.1.1.2@s"
	ddmTemperatureHighWarning      = oidPrefix + "12.2.1.1.3@s"
	ddmTemperatureCurrentValue     = oidPrefix + "12.2.1.1.4@s"
	ddmTemperatureLowWarning       = oidPrefix + "12.2.1.1.5@s"
	ddmTemperatureLowAlarm         = oidPrefix + "12.2.1.1.6@s"
	ddmVccHighAlarm                = oidPrefix + "12.2.1.1.7@s"
	ddmVccHighWarning              = oidPrefix + "12.2.1.1.8@s"
	ddmVccCurrentValue             = oidPrefix + "12.2.1.1.9@s"
	ddmVccLowWarning               = oidPrefix + "12.2.1.1.10@s"
	ddmVccLowAlarm                 = oidPrefix + "12.2.1.1.11@"
	ddmBiasHighAlarm               = oidPrefix + "12.2.1.1.12@s"
	ddmBiasHighWarning             = oidPrefix + "12.2.1.1.13@s"
	ddmBiasCurrentValue            = oidPrefix + "12.2.1.1.14@s"
	ddmBiasLowWarning              = oidPrefix + "12.2.1.1.15@s"
	ddmBiasLowAlarm                = oidPrefix + "12.2.1.1.16@s"
	ddmTxPowerHighAlarm            = oidPrefix + "12.2.1.1.17@s"
	ddmTxPowerHighWarning          = oidPrefix + "12.2.1.1.18@s"
	ddmTxPowerCurrentValue         = oidPrefix + "12.2.1.1.19@s"
	ddmTxPowerLowWarning           = oidPrefix + "12.2.1.1.20@s"
	ddmTxPowerLowAlarm             = oidPrefix + "12.2.1.1.21@s"
	ddmRxPowerHighAlarm            = oidPrefix + "12.2.1.1.22@s"
	ddmRxPowerHighWarning          = oidPrefix + "12.2.1.1.23@s"
	ddmRxPowerCurrentValue         = oidPrefix + "12.2.1.1.24@s"
	ddmRxPowerLowWarning           = oidPrefix + "12.2.1.1.25@s"
	ddmRxPowerLowAlarm             = oidPrefix + "12.2.1.1.26@s"
	monitorPowerNumber             = oidPrefix + "12.3.1.1.1@i" // PowerMonitor (12.3)
	monitorPowerStatus             = oidPrefix + "12.3.1.1.2@i"
	monitorPoEPortCfgNum           = oidPrefix + "12.4.1.1.1@i" // POEMonitor (12.4)
	monitorPoEPortStatus           = oidPrefix + "12.4.1.1.2@s"
	monitorPoEPortClass            = oidPrefix + "12.4.1.1.3@s"
	monitorPoEPortPowerConsumption = oidPrefix + "12.4.1.1.4@s"
	monitorPoEPortCurrent          = oidPrefix + "12.4.1.1.5@s"
	monitorPoEPortVoltage          = oidPrefix + "12.4.1.1.6@s"
	monitorPoEPortTemperature      = oidPrefix + "12.4.1.1.7@s"
	cpuLoadingMonitor              = oidPrefix + "12.5.1.0@i" // CPULoadingMonitor (12.5)

	// SaveConfiguration (13)
	saveCfgMgtAction = oidPrefix + "13.1.0@i-w"

	// FactoryDefault (14)
	factoryDefaultAction = oidPrefix + "14.1.0@i-w"

	// SystemReboot (15)
	systemRebootAction = oidPrefix + "15.1.0@i-w"

	// Maintenance (16)
	importConfiguration = oidPrefix + "16.1.0@s-w"
	upgrade             = oidPrefix + "16.2.0@s-w"
)

// TestTask is the task descirbed each snmp command and result
type TestTask struct {
	taskType          string // "set", "get", "walk"
	cmd               string
	valtype           string
	defaultVal        interface{}
	rawResult         string
	rawResultafterSet string
	testSuccess       string
}

func (t *TestTask) Init(taskType, cmd string) *TestTask {
	testTask := new(TestTask)
	testTask.taskType = taskType
	testTask.cmd = cmd
	return t
}

func (t *TestTask) Exec() {
	// if t.taskType == "get" {
	// 	_, result := utils.ShellExec(snmpGetPrefix + t.cmd)
	// 	t.result = result
	// } else if t.taskType == "set" {
	// 	_, b4 := utils.ShellExec(snmpGetPrefix + t.cmd)
	// } else if t.taskType == "walk" {
	//
	// }
}

func main() {
	defer func() {
		fmt.Println("test done")
	}()

	fmt.Println("command is", snmpSetPrefix+systemContact)

	_, b4 := utils.ShellExec(snmpGetPrefix + systemContact)
	_, setResult := utils.ShellExec(snmpSetPrefix + systemContact)
	fmt.Println("result is " + b4)
	fmt.Println("result is " + setResult)
}
