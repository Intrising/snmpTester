package main

import (
	"./utils"
	"fmt"
	"strings"
	// "time"
)

const (
	check            = "\u2714"
	cross            = "\u274c"
	nothingInThisOID = "No Such Instance currently exists at this OID"
	portAmount       = 20

	typeGet     = 0
	typeWalk    = 1
	typeSet     = 2
	typeWalkSet = 4

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
	// deviceIP = "192.168.16.140 "

	snmpGetPrefix     = snmpGet + snmpUser + snmpSecurityLevel + snmpAuthentication + snmpAuthPassPhrase + snmpPrivateProtocol + snmpPrivatePassPhrase + deviceIP
	snmpWalkPrefix    = snmpWalk + snmpUser + snmpSecurityLevel + snmpAuthentication + snmpAuthPassPhrase + snmpPrivateProtocol + snmpPrivatePassPhrase + deviceIP
	snmpSetPrefix     = snmpSet + snmpUser + snmpSecurityLevel + snmpAuthentication + snmpAuthPassPhrase + snmpPrivateProtocol + snmpPrivatePassPhrase + deviceIP
	snmpSetTypeString = " string "
	snmpSetTypeInt    = " integer "
	snmpSetTypeIpaddr = " ipaddress "
	oidPrefix         = "1.3.6.1.4.1.37072.302.2.3."

	// no postfix means snmpget commnad
	// all command can be read so no postfix for read-only
	// -s   means read-write
	// -w-s means use cmd walk and read-write
	// @s   means type string
	// @i   means type integer
	// @ip  means type ipaddress
	// -ps  means ps
	// -m:  marks the known failed type:
	//                                  1 =>  need to set value to the device
	//                                  2 =>  Spec issue
	//                                  3 =>  snmp program issue

	// SYSTEM (1)
	systemName           = oidPrefix + "1.1.0@s-s"
	systemLocation       = oidPrefix + "1.2.0@s-s"
	systemContact        = oidPrefix + "1.3.0@s-s"
	systemDescr          = oidPrefix + "1.4.0@s-s"
	systemFwVersion      = oidPrefix + "1.5.0@s"
	systemMacaddress     = oidPrefix + "1.6.0@s"
	systemAutoLogoutTime = oidPrefix + "1.7.0@i-s"
	systemSerialNum      = oidPrefix + "1.8.0@i"

	// Setting(2)
	vlanPortCfgNum           = oidPrefix + "2.1.1.1.1@i-w"
	vlanMembers              = oidPrefix + "2.1.1.1.2@s-w"
	vlanTags                 = oidPrefix + "2.1.1.1.3@s-w-m:1"
	pvidCfgNum               = oidPrefix + "2.1.2.1.1@i-w"
	vlanPvid                 = oidPrefix + "2.1.2.1.2@i-w-s"
	vlanFrameType            = oidPrefix + "2.1.2.1.3@i-w-s"
	mvrCfgNum                = oidPrefix + "2.2.1.1.1@i-w-m:2" // Mvr (2.2)
	mvrCfgVid                = oidPrefix + "2.2.1.1.2@i-w-m:3"
	mvrIPAddr                = oidPrefix + "2.2.1.1.3@s-w"
	mvrMembers               = oidPrefix + "2.2.1.1.4@s-w"
	igmpEnableQuerier        = oidPrefix + "2.3.1.0@i-s" // Igmp (2.3)
	igmpQuerierVersion       = oidPrefix + "2.3.2.0@i-s"
	igmpEnableSnooping       = oidPrefix + "2.3.3.0@i-s"
	igmpEnableFloodWellKnown = oidPrefix + "2.3.4.0@i-s"
	igmpPortNum              = oidPrefix + "2.3.5.1.1@i-w" // IgmpRouterTable (2.3.5)
	igmpRouterStatus         = oidPrefix + "2.3.5.1.2@i-w-s"
	igmpFastLeaveStatus      = oidPrefix + "2.3.5.1.3@i-w-s"
	igmpVidNum               = oidPrefix + "2.3.6.1.1@i-w" // IgmpStatisticsTable (2.3.6)
	igmpStatusQuerier        = oidPrefix + "2.3.6.1.2@s-w"
	igmpQuerierTx            = oidPrefix + "2.3.6.1.3@i-w"
	igmpQuerierRx            = oidPrefix + "2.3.6.1.4@i-w"
	igmpV1Rx                 = oidPrefix + "2.3.6.1.5@i-w"
	igmpV2Rx                 = oidPrefix + "2.3.6.1.6@i-w"
	igmpV3Rx                 = oidPrefix + "2.3.6.1.7@i-w"
	igmpV2Leave              = oidPrefix + "2.3.6.1.8@i-w"
	igmpEntriesEntryIndex    = oidPrefix + "2.3.7.1.1@s-w-m:3" // IgmpEntriesTable (2.3.7)
	igmpEntriesEntryIPAddr   = oidPrefix + "2.3.7.1.2@s-w-m:3"
	igmpEntriesEntryVID      = oidPrefix + "2.3.7.1.3@i-w-m:3"
	igmpEntriesEntryMembers  = oidPrefix + "2.3.7.1.4@s-w-m:3"

	// Status (3)
	lldpPortNum     = oidPrefix + "3.1.1.1.1@i-w" // LLDPInfo (3.1)
	lldpInfoContent = oidPrefix + "3.1.1.1.2@s-w-m:1"

	// Warning (11)
	faultAlarmPowerCfgNum        = oidPrefix + "11.1.1.1.1@i-w-m:3" // FaultAlarm (11.1)
	faultAlarmPowerStatus        = oidPrefix + "11.1.1.1.2@i-w-s-m:3"
	faultAlarmPortCfgNum         = oidPrefix + "11.1.2.1.1@i-w-m:3"
	faultAlarmPortLinkStatus     = oidPrefix + "11.1.2.1.2@i-w-s-m:3" // ===============================> here
	eventDDMEnabled              = oidPrefix + "11.2.1.1.0@i-s-m:3"   // EventDDMEnabled (11.2.1)
	eventDDMTemperatureLower     = oidPrefix + "11.2.1.2.0@s-s-m:3"
	eventDDMTemperatureUpper     = oidPrefix + "11.2.1.3.0@s-s-m:3"
	eventDDMVoltageLower         = oidPrefix + "11.2.1.4.0@s-s-m:3"
	eventDDMVoltageUpper         = oidPrefix + "11.2.1.5.0@s-s-m:3"
	eventDDMTxBiasLower          = oidPrefix + "11.2.1.6.0@s-s-m:3"
	eventDDMTTxBiasUpper         = oidPrefix + "11.2.1.7.0@s-s-m:3"
	eventDDMTxPowerLower         = oidPrefix + "11.2.1.8.0@s-s-m:3"
	eventDDMTxPowerUpper         = oidPrefix + "11.2.1.9.0@s-s-m:3"
	eventDDMRxPowerLower         = oidPrefix + "11.2.1.10.0@s-s-m:3"
	eventDDMRxPowerUpper         = oidPrefix + "11.2.1.11.0@s-s-m:3"
	eventMonitorEnabled          = oidPrefix + "11.2.2.1.0@i-s-m:3" // EventMonitor (11.2.1)
	eventMonitorTemperatureLower = oidPrefix + "11.2.2.2.0@s-s-m:3"
	eventMonitorTemperatureUpper = oidPrefix + "11.2.2.3.0@s-s-m:3"
	eventMonitorVoltageLower     = oidPrefix + "11.2.2.4.0@s-s-m:3"
	eventMonitorVoltageUpper     = oidPrefix + "11.2.2.5.0@s-s-m:3"
	eventMonitorCurrentLower     = oidPrefix + "11.2.2.6.0@s-s-m:3"
	eventMonitorCurrentUpper     = oidPrefix + "11.2.2.7.0@s-s-m:3"
	eventMonitorPowerLower       = oidPrefix + "11.2.2.8.0@s-s-m:3"
	eventMonitorPowerUpper       = oidPrefix + "11.2.2.9.0@s-s-m:3"
	eventPOEAPortCfgNum          = oidPrefix + "11.2.3.1.1.1@i-m:3" // EventPOEA (11.2.3)
	eventPOEAPingEnabled         = oidPrefix + "11.2.3.1.1.2@i-s-m:3"
	eventPOEAPingIPAddr          = oidPrefix + "11.2.3.1.1.3@ip-s-m:3"
	eventPOEAPingInterval        = oidPrefix + "11.2.3.1.1.4@i-s-m:3"
	eventPOEAPingRetry           = oidPrefix + "11.2.3.1.1.5@i-s-m:3"
	eventPOEAPingReboot          = oidPrefix + "11.2.3.1.1.6@i-s-m:3"
	eventPOEAPingFailAction      = oidPrefix + "11.2.3.1.1.7@i-s-m:3"
	localLogEnable               = oidPrefix + "11.3.1.1.0@i-s-m:3" // ActionConfiguration (11.3)
	remoteSystemLogCfgNum        = oidPrefix + "11.3.2.1.1.1@i-m:3" // RemoteSystemLog (11.3.2)
	remoteSystemLogHost          = oidPrefix + "11.3.2.1.1.2@ip-m:3"
	remoteSystemLogTag           = oidPrefix + "11.3.2.1.1.3@s-m:3"
	remoteSystemLogFacility      = oidPrefix + "11.3.2.1.1.4@s-m:3"
	remoteSystemLogHostName      = oidPrefix + "11.3.2.1.1.5@s-m:3"
	emailEnable                  = oidPrefix + "11.3.3.1.0@i-s-m:3"   // email (11.3.3.1)
	emailServerUser              = oidPrefix + "11.3.3.2.1.0@s-s-m:3" // emailServer (11.3.3.2)
	emailServerPassword          = oidPrefix + "11.3.3.2.2.0@s-s-m:3"
	emailServerHost              = oidPrefix + "11.3.3.2.3.0@s-s-m:3"
	emailServerSSLEnable         = oidPrefix + "11.3.3.2.4.0@i-s-m:3"
	emailSender                  = oidPrefix + "11.3.3.3.0@s-s-m:3"
	emailSubject                 = oidPrefix + "11.3.3.4.0@-s-m:3"
	emailCloudEnable             = oidPrefix + "11.3.3.5.0@i-s-m:3"
	emailReceiverCfgNum          = oidPrefix + "11.3.3.6.1.1@i-m:3"
	emailReceiverHost            = oidPrefix + "11.3.3.6.1.2@s-m:3"
	smsEnable                    = oidPrefix + "11.3.4.1.0@i-s-m:3" // SMS (11.3.4)
	smsUser                      = oidPrefix + "11.3.4.2.0@s-s-m:3"
	smsPassword                  = oidPrefix + "11.3.4.3.0@s-s-m:3"
	smsSenderText                = oidPrefix + "11.3.4.4.0@s-s-m:3"
	smsReceiverCfgNum            = oidPrefix + "11.3.4.5.1.1@i-m:3"
	smsReceiverPhone             = oidPrefix + "11.3.4.5.1.2@s-m:3"
	snmpResponseLocale           = oidPrefix + "11.3.5.1.1.0@i-s-m:3" // Snmp (11.3.5)
	snmpCommunityCfgNum          = oidPrefix + "11.3.5.1.2.1.1@i-m:3"
	snmpCommunityCfgString       = oidPrefix + "11.3.5.1.2.1.2@s-m:3"
	snmpCommunityCfgReadOnly     = oidPrefix + "11.3.5.1.2.1.3@i-m:3"
	snmpTrapCfgNum               = oidPrefix + "11.3.5.2.1.1.1@i-m:3" // Trap (11.3.5.2)
	snmpTrapCfgCommunity         = oidPrefix + "11.3.5.2.1.1.2@s-m:3"
	snmpTrapCfgIPAddress         = oidPrefix + "11.3.5.2.1.1.3@ip-m:3"
	snmpTrapCfgVersion           = oidPrefix + "11.3.5.2.1.1.4@i-m:3"
	snmpV3UserCfgNum             = oidPrefix + "11.3.5.3.1.1.1@i-m:3" // V3User (11.3.5.3)
	snmpV3UserCfgName            = oidPrefix + "11.3.5.3.1.1.2@s-m:3"
	snmpV3UserCfgSecurityLevel   = oidPrefix + "11.3.5.3.1.1.3@i-m:3"
	snmpV3UserCfgAuthProtocal    = oidPrefix + "11.3.5.3.1.1.4@i-m:3"
	snmpV3UserCfgAuthPassword    = oidPrefix + "11.3.5.3.1.1.5@s-m:3"
	snmpV3UserCfgPrivProtocal    = oidPrefix + "11.3.5.3.1.1.6@s-m:3"
	snmpV3UserCfgPrivPassword    = oidPrefix + "11.3.5.3.1.1.7@s-m:3"
	doutCfgNum                   = oidPrefix + "11.3.6.1.1.1@i-m:3" // Dout (11.3.6)
	doutCfgEnable                = oidPrefix + "11.3.6.1.1.2@i-s-m:3"
	doutCfgAction                = oidPrefix + "11.3.6.1.1.3@i-s-m:3"
	deviceBootEvent              = oidPrefix + "11.4.1.1.0@i-s-m:3" // EventActionMap (11.4.1.1)
	authenticationFailureEvent   = oidPrefix + "11.4.1.2.0@i-s-m:3"
	authenticationSuccessEvent   = oidPrefix + "11.4.1.3.0@i-s-m:3"
	deviceDDMEvent               = oidPrefix + "11.4.1.4.0@i-s-m:3"
	devicePOEEvent               = oidPrefix + "11.4.1.5.0@i-s-m:3"
	devicePOEBEvent              = oidPrefix + "11.4.1.6.0@i-s-m:3"
	ringTopologyChangeEvent      = oidPrefix + "11.4.1.7.0@i-s-m:3"
	envMonitorEvent              = oidPrefix + "11.4.1.8.0@i-s-m:3"
	eventPortNumber              = oidPrefix + "11.4.2.1.1.1@i-m:3" // PortsEvent (11.4.2)
	eventPortEventLog            = oidPrefix + "11.4.2.1.1.2@i-s-m:3"
	eventPortEventsms            = oidPrefix + "11.4.2.1.1.3@i-s-m:3"
	eventPortEventSMTP           = oidPrefix + "11.4.2.1.1.4@i-s-m:3"
	eventPortEventsnmpTRAP       = oidPrefix + "11.4.2.1.1.5@i-s-m:3"
	eventPortEventdout1          = oidPrefix + "11.4.2.1.1.6@i-s-m:3"
	eventPortEventdout2          = oidPrefix + "11.4.2.1.1.7@i-s-m:3"
	eventPowerNumber             = oidPrefix + "11.4.3.1.1.1@i-m:3" // PowerEvent (11.4.3)
	eventPowerEventLog           = oidPrefix + "11.4.3.1.1.2@i-s-m:3"
	eventPowerEventsms           = oidPrefix + "11.4.3.1.1.3@i-s-m:3"
	eventPowerEventSMTP          = oidPrefix + "11.4.3.1.1.4@i-s-m:3"
	eventPowerEventsnmpTRAP      = oidPrefix + "11.4.3.1.1.5@i-s-m:3"
	eventPowerEventdout1         = oidPrefix + "11.4.3.1.1.6@i-s-m:3"
	eventPowerEventdout2         = oidPrefix + "11.4.3.1.1.7@i-s-m:3"
	eventDiNumber                = oidPrefix + "11.4.4.1.1.1@i-m:3" // DiEvent (11.4.4)
	eventDiEventLog              = oidPrefix + "11.4.4.1.1.2@i-s-m:3"
	eventDiEventsms              = oidPrefix + "11.4.4.1.1.3@i-s-m:3"
	eventDiEventSMTP             = oidPrefix + "11.4.4.1.1.4@i-s-m:3"
	eventDiEventsnmpTRAP         = oidPrefix + "11.4.4.1.1.5@i-s-m:3"
	eventDiEventdout1            = oidPrefix + "11.4.4.1.1.6@i-s-m:3"
	eventDiEventdout2            = oidPrefix + "11.4.4.1.1.7@i-s-m:3"

	// Monitoring (12)
	envVoltage                     = oidPrefix + "12.1.1.0@s" // ENVMonitor (12.1)
	envCurrent                     = oidPrefix + "12.1.2.0@s"
	envWalt                        = oidPrefix + "12.1.3.0@s"
	envTemperature                 = oidPrefix + "12.1.4.0@s"
	ddmPortNumber                  = oidPrefix + "12.2.1.1.1@i-m:3" // DDM (12.2)
	ddmTemperatureHighAlarm        = oidPrefix + "12.2.1.1.2@s-m:3"
	ddmTemperatureHighWarning      = oidPrefix + "12.2.1.1.3@s-m:3"
	ddmTemperatureCurrentValue     = oidPrefix + "12.2.1.1.4@s-m:3"
	ddmTemperatureLowWarning       = oidPrefix + "12.2.1.1.5@s-m:3"
	ddmTemperatureLowAlarm         = oidPrefix + "12.2.1.1.6@s-m:3"
	ddmVccHighAlarm                = oidPrefix + "12.2.1.1.7@s-m:3"
	ddmVccHighWarning              = oidPrefix + "12.2.1.1.8@s-m:3"
	ddmVccCurrentValue             = oidPrefix + "12.2.1.1.9@s-m:3"
	ddmVccLowWarning               = oidPrefix + "12.2.1.1.10@s-m:3"
	ddmVccLowAlarm                 = oidPrefix + "12.2.1.1.11@s-m:3"
	ddmBiasHighAlarm               = oidPrefix + "12.2.1.1.12@s-m:3"
	ddmBiasHighWarning             = oidPrefix + "12.2.1.1.13@s-m:3"
	ddmBiasCurrentValue            = oidPrefix + "12.2.1.1.14@s-m:3"
	ddmBiasLowWarning              = oidPrefix + "12.2.1.1.15@s-m:3"
	ddmBiasLowAlarm                = oidPrefix + "12.2.1.1.16@s-m:3"
	ddmTxPowerHighAlarm            = oidPrefix + "12.2.1.1.17@s-m:3"
	ddmTxPowerHighWarning          = oidPrefix + "12.2.1.1.18@s-m:3"
	ddmTxPowerCurrentValue         = oidPrefix + "12.2.1.1.19@s-m:3"
	ddmTxPowerLowWarning           = oidPrefix + "12.2.1.1.20@s-m:3"
	ddmTxPowerLowAlarm             = oidPrefix + "12.2.1.1.21@s-m:3"
	ddmRxPowerHighAlarm            = oidPrefix + "12.2.1.1.22@s-m:3"
	ddmRxPowerHighWarning          = oidPrefix + "12.2.1.1.23@s-m:3"
	ddmRxPowerCurrentValue         = oidPrefix + "12.2.1.1.24@s-m:3"
	ddmRxPowerLowWarning           = oidPrefix + "12.2.1.1.25@s-m:3"
	ddmRxPowerLowAlarm             = oidPrefix + "12.2.1.1.26@s-m:3"
	monitorPowerNumber             = oidPrefix + "12.3.1.1.1@i-m:3" // PowerMonitor (12.3)
	monitorPowerStatus             = oidPrefix + "12.3.1.1.2@i-m:3"
	monitorPoEPortCfgNum           = oidPrefix + "12.4.1.1.1@i-w-m:3" // POEMonitor (12.4)
	monitorPoEPortStatus           = oidPrefix + "12.4.1.1.2@s-w-m:3"
	monitorPoEPortClass            = oidPrefix + "12.4.1.1.3@s-w-m:3"
	monitorPoEPortPowerConsumption = oidPrefix + "12.4.1.1.4@s-w-m:3"
	monitorPoEPortCurrent          = oidPrefix + "12.4.1.1.5@s-w-m:3"
	monitorPoEPortVoltage          = oidPrefix + "12.4.1.1.6@s-w-m:3"
	monitorPoEPortTemperature      = oidPrefix + "12.4.1.1.7@s-w-m:3"
	cpuLoadingMonitor              = oidPrefix + "12.5.1.0@i" // CPULoadingMonitor (12.5)

	// SaveConfiguration (13)
	saveCfgMgtAction = oidPrefix + "13.1.0@s-s" // the value type to set is integer, but get will be string, set it string for now

	// FactoryDefault (14)
	factoryDefaultAction = oidPrefix + "14.1.0@s-s" // the value type to set is integer, but get will be string, set it string for now

	// SystemReboot (15)
	systemRebootAction = oidPrefix + "15.1.0@s-s" // the value type to set is integer, but get will be string, set it string for now

	// Maintenance (16)
	importConfiguration = oidPrefix + "16.1.0@s-s"
	upgrade             = oidPrefix + "16.2.0@s-s"
)

var oidMap map[string]*Task
var testValMap map[string]string
var taskEntry []*Task
var stats Stats

// Stats is the Statistics
type Stats struct {
	total       int
	pass        int
	failed      int
	marked      int
	unmarkedOID []string
}

func (s *Stats) init() {
	s.total = 0
	s.pass = 0
	s.failed = 0
	s.marked = 0
}

func (s *Stats) AddunmarkedOID(oid string) {
	s.unmarkedOID = append(s.unmarkedOID, oid)
}

func (s *Stats) AddPass() {
	s.pass++
	s.total++
}

func (s *Stats) AddFailed() {
	s.failed++
	s.total++
}

func (s *Stats) AddMarked() {
	s.marked++
}

func isGet(t *Task) bool {
	return t.taskType == typeGet || t.taskType == typeSet
}

func isWalk(t *Task) bool {
	return t.taskType == typeWalk || t.taskType == typeWalkSet
}

func taskTranslator(typeInt int) string {
	if typeInt == typeGet {
		return "get"
	} else if typeInt == typeWalk {
		return "walk"
	} else if typeInt == typeSet {
		return "set"
	} else if typeInt == typeWalkSet {
		return "walkSet"
	}
	return "sth wrong"
}

// Task is the task descirbed each snmp command and result
type Task struct {
	name     string
	taskType int //  0 = "get", 1 = "walk", 2 = "set" 	, 3 = "walkSet"
	oid      string

	getCmd  string
	walkCmd string
	setCmd  string

	valtype               string
	defaultVal            []string
	rawResult             string
	rawResultafterSet     string
	rawResultAfterDefault string
	testSuccess           string
	failedReason          string
	ps                    string
	failedtype            string
}

func parseTaskTypeFromCmd(oid string) int {
	if strings.Contains(oid, "-w-s") {
		return typeWalkSet
	} else if strings.Contains(oid, "-w") {
		return typeWalk
	} else if strings.Contains(oid, "-s") {
		return typeSet
	}
	return typeGet
}

func genTask(name, oid string) *Task {
	t := new(Task)
	t.Init(name, oid)
	return t
}

func parseValTypeFromCmd(oid string) string {
	if strings.Contains(oid, "@i") && !strings.Contains(oid, "@ip") {
		return strings.TrimSpace(snmpSetTypeInt)
	} else if strings.Contains(oid, "@s") {
		return strings.TrimSpace(snmpSetTypeString)
	} else if strings.Contains(oid, "@ip") {
		return strings.TrimSpace(snmpSetTypeIpaddr)
	}
	return "sthing wrong"
}

func parsePSFromCmd(oid string) string {
	if strings.Contains(oid, "-ps:") {
		return strings.Split(oid, "-ps:")[1]
	}
	return ""
}

func failedTypeTranlator(from string) string {
	switch from {
	case "1":
		return "Need to set value to the device, or connect to another device to get the value."
	case "2":
		return "Spec issue"
	case "3":
		return "snmp program issue"
	default:
		return ""
	}
}

func parseFailedType(oid string) string {
	if strings.Contains(oid, "-m:") {
		return strings.Split(oid, "-m:")[1]
	}
	return ""
}

func rmPostFix(oid string) string {
	return strings.Split(oid, "@")[0]
}

func probe(mainString, subString string) bool {
	return strings.Contains(mainString, subString)
}

func (t *Task) Failed(reason string) {
	t.testSuccess = cross
	t.failedReason = reason
	stats.AddFailed()
	if len(t.failedtype) > 0 {
		stats.AddMarked()
	} else {
		stats.AddunmarkedOID(t.oid)
	}
}

func (t *Task) Success() {
	t.testSuccess = check
	stats.AddPass()
}

func (t *Task) Init(taskName, oid string) {
	t.name = taskName
	t.taskType = parseTaskTypeFromCmd(oid)
	t.valtype = parseValTypeFromCmd(oid)
	t.failedtype = parseFailedType(oid)

	oid = rmPostFix(oid)
	t.oid = strings.Split(oid, oidPrefix)[1]
	if t.taskType == typeGet {
		t.getCmd = snmpGetPrefix + oid
	} else if t.taskType == typeWalk {
		t.walkCmd = snmpWalkPrefix + oid
	} else if t.taskType == typeSet {
		// We need to test all oid by get, and some oid has the read-write access, so we store the set cmd in task.setCmd
		//
		t.getCmd = snmpGetPrefix + oid
		t.setCmd = snmpSetPrefix + oid + t.valtype + testValMap[t.valtype]
	} else if t.taskType == typeWalkSet {
		t.walkCmd = snmpWalkPrefix + oid
		// @@Todo: handle the condition of walk and set
		// t.setCmd = snmpSetPrefix + oid + t.valtype + testValMap[t.valtype]
	}

}

func (t *Task) Exec() {
	fmt.Println("================================ type is ", taskTranslator(t.taskType))

	if isGet(t) {
		err, result := utils.ShellExec(t.getCmd)
		t.rawResult = result
		if err != nil {
			fmt.Println("exec error ", err)
		}
	} else if isWalk(t) {
		err, result := utils.ShellExec(t.walkCmd)
		if err != nil {
			fmt.Println("exec error ", err)
		}
		// work around to fix the issue that snpmwalk after snmpget will lead to the result that nothingInThisOID
		err, result = utils.ShellExec(t.walkCmd)
		if err != nil {
			fmt.Println("exec error ", err)
		}
		t.rawResult = result
	}

	t.handleFirstGet()
	t.printResult()

}

func (t *Task) handleFirstGet() {
	// fmt.Println("handleRawVal")
	// fmt.Println("Raw lines is =>", strings.Split(t.rawResult, "\n")[0])

	if probe(t.rawResult, nothingInThisOID) {
		t.Failed(nothingInThisOID)
		return
	}

	// Check the value type first
	if t.valtype == "string" && !probe(t.rawResult, "STRING:") {
		t.Failed("Expect type string, but probe other type")
		return
	} else if t.valtype == "integer" && !probe(t.rawResult, "INTEGER:") {
		t.Failed("Expect type integer, but probe other type")
		return
	} else if t.valtype == "ipaddress" && !probe(t.rawResult, "IPADDRESS:") {
		t.Failed("Expect type ipaddress, but probe other type")
		return
	}

	// Parse all the value from snmpwalk line by line
	if isWalk(t) {
		lines := strings.Split(t.rawResult, "\n")
		for _, line := range lines {
			// fmt.Println("line is ", line)
			if probe(line, "STRING: ") {
				val := strings.Split(line, "STRING: ")[1]
				val = strings.Replace(val, "\"", "", -1)
				t.defaultVal = append(t.defaultVal, val)
			} else if probe(line, "INTEGER:") {
				val := strings.Split(line, "INTEGER: ")[1]
				t.defaultVal = append(t.defaultVal, val)
			}
		}
		t.Success()

	} else if isGet(t) {
		if probe(t.rawResult, "STRING:") {
			val := strings.Split(t.rawResult, "STRING: ")[1]
			val = strings.Replace(val, "\"", "", -1)
			t.defaultVal = []string{val}
			t.Success()
		} else if probe(t.rawResult, "INTEGER:") {
			val := strings.Split(t.rawResult, "INTEGER: ")[1]
			t.defaultVal = []string{val}
			t.Success()
		}
	}
}

func (t *Task) printResult() {
	// fmt.Println("Raw Val is => ", t.rawResult)
	fmt.Println("Name:   ", t.name)
	fmt.Println("oid:    ", t.oid)
	if isGet(t) {
		fmt.Println("GetCmd:  ", t.getCmd)
	} else if isWalk(t) {
		fmt.Println("WalkCmd: ", t.walkCmd)
	}
	fmt.Println("GET "+t.name+" ===== >", t.defaultVal)
	fmt.Println("\n")
	fmt.Println("Test pass: ", t.testSuccess)
	if t.testSuccess == cross {
		if len(t.failedtype) > 0 {
			fmt.Println("Failed type:    ", failedTypeTranlator(t.failedtype))
		}
		fmt.Println("Failed reaseon: ", t.failedReason)
		fmt.Println("Raw Val is => ", t.rawResult)
	}
}

func init() {
	stats.init()

	testValMap = make(map[string]string)
	testValMap["string"] = "testWalter"
	testValMap["integer"] = "20"
	testValMap["ipaddress"] = "192.168.1.1"

	taskEntry = append(taskEntry, genTask("systemName", systemName))
	taskEntry = append(taskEntry, genTask("systemLocation", systemLocation))
	taskEntry = append(taskEntry, genTask("systemContact", systemContact))
	taskEntry = append(taskEntry, genTask("systemDescr", systemDescr))
	taskEntry = append(taskEntry, genTask("systemFwVersion", systemFwVersion))
	taskEntry = append(taskEntry, genTask("systemMacaddress", systemMacaddress))
	taskEntry = append(taskEntry, genTask("systemAutoLogoutTime", systemAutoLogoutTime))
	taskEntry = append(taskEntry, genTask("systemSerialNum", systemSerialNum))

	taskEntry = append(taskEntry, genTask("vlanPortCfgNum", vlanPortCfgNum))
	taskEntry = append(taskEntry, genTask("vlanMembers", vlanMembers))
	taskEntry = append(taskEntry, genTask("vlanTags", vlanTags))
	taskEntry = append(taskEntry, genTask("pvidCfgNum", pvidCfgNum))
	taskEntry = append(taskEntry, genTask("vlanPvid", vlanPvid))
	taskEntry = append(taskEntry, genTask("vlanFrameType", vlanFrameType))
	taskEntry = append(taskEntry, genTask("mvrCfgNum", mvrCfgNum))
	taskEntry = append(taskEntry, genTask("mvrCfgVid", mvrCfgVid))
	taskEntry = append(taskEntry, genTask("mvrIPAddr", mvrIPAddr))
	taskEntry = append(taskEntry, genTask("mvrMembers", mvrMembers))
	taskEntry = append(taskEntry, genTask("igmpEnableQuerier", igmpEnableQuerier))
	taskEntry = append(taskEntry, genTask("igmpQuerierVersion", igmpQuerierVersion))
	taskEntry = append(taskEntry, genTask("igmpEnableSnooping", igmpEnableSnooping))
	taskEntry = append(taskEntry, genTask("igmpEnableFloodWellKnown", igmpEnableFloodWellKnown))
	taskEntry = append(taskEntry, genTask("igmpPortNum", igmpPortNum))
	taskEntry = append(taskEntry, genTask("igmpRouterStatus", igmpRouterStatus))
	taskEntry = append(taskEntry, genTask("igmpFastLeaveStatus", igmpFastLeaveStatus))
	taskEntry = append(taskEntry, genTask("igmpVidNum", igmpVidNum))
	taskEntry = append(taskEntry, genTask("igmpStatusQuerier", igmpStatusQuerier))
	taskEntry = append(taskEntry, genTask("igmpQuerierTx", igmpQuerierTx))
	taskEntry = append(taskEntry, genTask("igmpQuerierRx", igmpQuerierRx))
	taskEntry = append(taskEntry, genTask("igmpV1Rx", igmpV1Rx))
	taskEntry = append(taskEntry, genTask("igmpV2Rx", igmpV2Rx))
	taskEntry = append(taskEntry, genTask("igmpV3Rx", igmpV3Rx))
	taskEntry = append(taskEntry, genTask("igmpV2Leave", igmpV2Leave))
	taskEntry = append(taskEntry, genTask("igmpEntriesEntryIndex", igmpEntriesEntryIndex))
	taskEntry = append(taskEntry, genTask("igmpEntriesEntryIPAddr", igmpEntriesEntryIPAddr))
	taskEntry = append(taskEntry, genTask("igmpEntriesEntryVID", igmpEntriesEntryVID))
	taskEntry = append(taskEntry, genTask("igmpEntriesEntryMembers", igmpEntriesEntryMembers))
	taskEntry = append(taskEntry, genTask("lldpPortNum", lldpPortNum))
	taskEntry = append(taskEntry, genTask("lldpInfoContent", lldpInfoContent))
	taskEntry = append(taskEntry, genTask("faultAlarmPowerCfgNum", faultAlarmPowerCfgNum))
	taskEntry = append(taskEntry, genTask("faultAlarmPowerStatus", faultAlarmPowerStatus))
	taskEntry = append(taskEntry, genTask("faultAlarmPortCfgNum", faultAlarmPortCfgNum))
	taskEntry = append(taskEntry, genTask("faultAlarmPortLinkStatus", faultAlarmPortLinkStatus))
	taskEntry = append(taskEntry, genTask("eventDDMEnabled", eventDDMEnabled))
	taskEntry = append(taskEntry, genTask("eventDDMTemperatureLower", eventDDMTemperatureLower))
	taskEntry = append(taskEntry, genTask("eventDDMTemperatureUpper", eventDDMTemperatureUpper))
	taskEntry = append(taskEntry, genTask("eventDDMVoltageLower", eventDDMVoltageLower))
	taskEntry = append(taskEntry, genTask("eventDDMVoltageUpper", eventDDMVoltageUpper))
	taskEntry = append(taskEntry, genTask("eventDDMTxBiasLower", eventDDMTxBiasLower))
	taskEntry = append(taskEntry, genTask("eventDDMTTxBiasUpper", eventDDMTTxBiasUpper))
	taskEntry = append(taskEntry, genTask("eventDDMTxPowerLower", eventDDMTxPowerLower))
	taskEntry = append(taskEntry, genTask("eventDDMTxPowerUpper", eventDDMTxPowerUpper))
	taskEntry = append(taskEntry, genTask("eventDDMRxPowerLower", eventDDMRxPowerLower))
	taskEntry = append(taskEntry, genTask("eventDDMRxPowerUpper", eventDDMRxPowerUpper))
	taskEntry = append(taskEntry, genTask("eventMonitorEnabled", eventMonitorEnabled))
	taskEntry = append(taskEntry, genTask("eventMonitorTemperatureLower", eventMonitorTemperatureLower))
	taskEntry = append(taskEntry, genTask("eventMonitorTemperatureUpper", eventMonitorTemperatureUpper))
	taskEntry = append(taskEntry, genTask("eventMonitorVoltageLower", eventMonitorVoltageLower))
	taskEntry = append(taskEntry, genTask("eventMonitorVoltageUpper", eventMonitorVoltageUpper))
	taskEntry = append(taskEntry, genTask("eventMonitorCurrentLower", eventMonitorCurrentLower))
	taskEntry = append(taskEntry, genTask("eventMonitorCurrentUpper", eventMonitorCurrentUpper))
	taskEntry = append(taskEntry, genTask("eventMonitorPowerLower", eventMonitorPowerLower))
	taskEntry = append(taskEntry, genTask("eventMonitorPowerUpper", eventMonitorPowerUpper))
	taskEntry = append(taskEntry, genTask("eventPOEAPortCfgNum", eventPOEAPortCfgNum))
	taskEntry = append(taskEntry, genTask("eventPOEAPingEnabled", eventPOEAPingEnabled))
	taskEntry = append(taskEntry, genTask("eventPOEAPingIPAddr", eventPOEAPingIPAddr))
	taskEntry = append(taskEntry, genTask("eventPOEAPingInterval", eventPOEAPingInterval))
	taskEntry = append(taskEntry, genTask("eventPOEAPingRetry", eventPOEAPingRetry))
	taskEntry = append(taskEntry, genTask("eventPOEAPingReboot", eventPOEAPingReboot))
	taskEntry = append(taskEntry, genTask("eventPOEAPingFailAction", eventPOEAPingFailAction))
	taskEntry = append(taskEntry, genTask("localLogEnable", localLogEnable))
	taskEntry = append(taskEntry, genTask("remoteSystemLogCfgNum", remoteSystemLogCfgNum))
	taskEntry = append(taskEntry, genTask("remoteSystemLogHost", remoteSystemLogHost))
	taskEntry = append(taskEntry, genTask("remoteSystemLogTag", remoteSystemLogTag))
	taskEntry = append(taskEntry, genTask("remoteSystemLogFacility", remoteSystemLogFacility))
	taskEntry = append(taskEntry, genTask("remoteSystemLogHostName", remoteSystemLogHostName))
	taskEntry = append(taskEntry, genTask("emailEnable", emailEnable))
	taskEntry = append(taskEntry, genTask("emailServerUser", emailServerUser))
	taskEntry = append(taskEntry, genTask("emailServerPassword", emailServerPassword))
	taskEntry = append(taskEntry, genTask("emailServerHost", emailServerHost))
	taskEntry = append(taskEntry, genTask("emailServerSSLEnable", emailServerSSLEnable))
	taskEntry = append(taskEntry, genTask("emailSender", emailSender))
	taskEntry = append(taskEntry, genTask("emailSubject", emailSubject))
	taskEntry = append(taskEntry, genTask("emailCloudEnable", emailCloudEnable))
	taskEntry = append(taskEntry, genTask("emailReceiverCfgNum", emailReceiverCfgNum))
	taskEntry = append(taskEntry, genTask("emailReceiverHost", emailReceiverHost))
	taskEntry = append(taskEntry, genTask("smsEnable", smsEnable))
	taskEntry = append(taskEntry, genTask("smsUser", smsUser))
	taskEntry = append(taskEntry, genTask("smsPassword", smsPassword))
	taskEntry = append(taskEntry, genTask("smsSenderText", smsSenderText))
	taskEntry = append(taskEntry, genTask("smsReceiverCfgNum", smsReceiverCfgNum))
	taskEntry = append(taskEntry, genTask("smsReceiverPhone", smsReceiverPhone))
	taskEntry = append(taskEntry, genTask("snmpResponseLocale", snmpResponseLocale))
	taskEntry = append(taskEntry, genTask("snmpCommunityCfgNum", snmpCommunityCfgNum))
	taskEntry = append(taskEntry, genTask("snmpCommunityCfgString", snmpCommunityCfgString))
	taskEntry = append(taskEntry, genTask("snmpCommunityCfgReadOnly", snmpCommunityCfgReadOnly))
	taskEntry = append(taskEntry, genTask("snmpTrapCfgNum", snmpTrapCfgNum))
	taskEntry = append(taskEntry, genTask("snmpTrapCfgCommunity", snmpTrapCfgCommunity))
	taskEntry = append(taskEntry, genTask("snmpTrapCfgIPAddress", snmpTrapCfgIPAddress))
	taskEntry = append(taskEntry, genTask("snmpTrapCfgVersion", snmpTrapCfgVersion))
	taskEntry = append(taskEntry, genTask("snmpV3UserCfgNum", snmpV3UserCfgNum))
	taskEntry = append(taskEntry, genTask("snmpV3UserCfgName", snmpV3UserCfgName))
	taskEntry = append(taskEntry, genTask("snmpV3UserCfgSecurityLevel", snmpV3UserCfgSecurityLevel))
	taskEntry = append(taskEntry, genTask("snmpV3UserCfgAuthProtocal", snmpV3UserCfgAuthProtocal))
	taskEntry = append(taskEntry, genTask("snmpV3UserCfgAuthPassword", snmpV3UserCfgAuthPassword))
	taskEntry = append(taskEntry, genTask("snmpV3UserCfgPrivProtocal", snmpV3UserCfgPrivProtocal))
	taskEntry = append(taskEntry, genTask("snmpV3UserCfgPrivPassword", snmpV3UserCfgPrivPassword))
	taskEntry = append(taskEntry, genTask("doutCfgNum", doutCfgNum))
	taskEntry = append(taskEntry, genTask("doutCfgEnable", doutCfgEnable))
	taskEntry = append(taskEntry, genTask("doutCfgAction", doutCfgAction))
	taskEntry = append(taskEntry, genTask("deviceBootEvent", deviceBootEvent))
	taskEntry = append(taskEntry, genTask("authenticationFailureEvent", authenticationFailureEvent))
	taskEntry = append(taskEntry, genTask("authenticationSuccessEvent", authenticationSuccessEvent))
	taskEntry = append(taskEntry, genTask("deviceDDMEvent", deviceDDMEvent))
	taskEntry = append(taskEntry, genTask("devicePOEEvent", devicePOEEvent))
	taskEntry = append(taskEntry, genTask("devicePOEBEvent", devicePOEBEvent))
	taskEntry = append(taskEntry, genTask("ringTopologyChangeEvent", ringTopologyChangeEvent))
	taskEntry = append(taskEntry, genTask("envMonitorEvent", envMonitorEvent))
	taskEntry = append(taskEntry, genTask("eventPortNumber", eventPortNumber))
	taskEntry = append(taskEntry, genTask("eventPortEventLog", eventPortEventLog))
	taskEntry = append(taskEntry, genTask("eventPortEventsms", eventPortEventsms))
	taskEntry = append(taskEntry, genTask("eventPortEventSMTP", eventPortEventSMTP))
	taskEntry = append(taskEntry, genTask("eventPortEventsnmpTRAP", eventPortEventsnmpTRAP))
	taskEntry = append(taskEntry, genTask("eventPortEventdout1", eventPortEventdout1))
	taskEntry = append(taskEntry, genTask("eventPortEventdout2", eventPortEventdout2))
	taskEntry = append(taskEntry, genTask("eventPowerNumber", eventPowerNumber))
	taskEntry = append(taskEntry, genTask("eventPowerEventLog", eventPowerEventLog))
	taskEntry = append(taskEntry, genTask("eventPowerEventsms", eventPowerEventsms))
	taskEntry = append(taskEntry, genTask("eventPowerEventSMTP", eventPowerEventSMTP))
	taskEntry = append(taskEntry, genTask("eventPowerEventsnmpTRAP", eventPowerEventsnmpTRAP))
	taskEntry = append(taskEntry, genTask("eventPowerEventdout1", eventPowerEventdout1))
	taskEntry = append(taskEntry, genTask("eventPowerEventdout2", eventPowerEventdout2))
	taskEntry = append(taskEntry, genTask("eventDiNumber", eventDiNumber))
	taskEntry = append(taskEntry, genTask("eventDiEventLog", eventDiEventLog))
	taskEntry = append(taskEntry, genTask("eventDiEventsms", eventDiEventsms))
	taskEntry = append(taskEntry, genTask("eventDiEventSMTP", eventDiEventSMTP))
	taskEntry = append(taskEntry, genTask("eventDiEventsnmpTRAP", eventDiEventsnmpTRAP))
	taskEntry = append(taskEntry, genTask("eventDiEventdout1", eventDiEventdout1))
	taskEntry = append(taskEntry, genTask("eventDiEventdout2", eventDiEventdout2))
	taskEntry = append(taskEntry, genTask("envVoltage", envVoltage))
	taskEntry = append(taskEntry, genTask("envCurrent", envCurrent))
	taskEntry = append(taskEntry, genTask("envWalt", envWalt))
	taskEntry = append(taskEntry, genTask("envTemperature", envTemperature))
	taskEntry = append(taskEntry, genTask("ddmPortNumber", ddmPortNumber))
	taskEntry = append(taskEntry, genTask("ddmTemperatureHighAlarm", ddmTemperatureHighAlarm))
	taskEntry = append(taskEntry, genTask("ddmTemperatureHighWarning", ddmTemperatureHighWarning))
	taskEntry = append(taskEntry, genTask("ddmTemperatureCurrentValue", ddmTemperatureCurrentValue))
	taskEntry = append(taskEntry, genTask("ddmTemperatureLowWarning", ddmTemperatureLowWarning))
	taskEntry = append(taskEntry, genTask("ddmTemperatureLowAlarm", ddmTemperatureLowAlarm))
	taskEntry = append(taskEntry, genTask("ddmVccHighAlarm", ddmVccHighAlarm))
	taskEntry = append(taskEntry, genTask("ddmVccHighWarning", ddmVccHighWarning))
	taskEntry = append(taskEntry, genTask("ddmVccCurrentValue", ddmVccCurrentValue))
	taskEntry = append(taskEntry, genTask("ddmVccLowWarning", ddmVccLowWarning))
	taskEntry = append(taskEntry, genTask("ddmVccLowAlarm", ddmVccLowAlarm))
	taskEntry = append(taskEntry, genTask("ddmBiasHighAlarm", ddmBiasHighAlarm))
	taskEntry = append(taskEntry, genTask("ddmBiasHighWarning", ddmBiasHighWarning))
	taskEntry = append(taskEntry, genTask("ddmBiasCurrentValue", ddmBiasCurrentValue))
	taskEntry = append(taskEntry, genTask("ddmBiasLowWarning", ddmBiasLowWarning))
	taskEntry = append(taskEntry, genTask("ddmBiasLowAlarm", ddmBiasLowAlarm))
	taskEntry = append(taskEntry, genTask("ddmTxPowerHighAlarm", ddmTxPowerHighAlarm))
	taskEntry = append(taskEntry, genTask("ddmTxPowerHighWarning", ddmTxPowerHighWarning))
	taskEntry = append(taskEntry, genTask("ddmTxPowerCurrentValue", ddmTxPowerCurrentValue))
	taskEntry = append(taskEntry, genTask("ddmTxPowerLowWarning", ddmTxPowerLowWarning))
	taskEntry = append(taskEntry, genTask("ddmTxPowerLowAlarm", ddmTxPowerLowAlarm))
	taskEntry = append(taskEntry, genTask("ddmRxPowerHighAlarm", ddmRxPowerHighAlarm))
	taskEntry = append(taskEntry, genTask("ddmRxPowerHighWarning", ddmRxPowerHighWarning))
	taskEntry = append(taskEntry, genTask("ddmRxPowerCurrentValue", ddmRxPowerCurrentValue))
	taskEntry = append(taskEntry, genTask("ddmRxPowerLowWarning", ddmRxPowerLowWarning))
	taskEntry = append(taskEntry, genTask("ddmRxPowerLowAlarm", ddmRxPowerLowAlarm))
	taskEntry = append(taskEntry, genTask("monitorPowerNumber", monitorPowerNumber))
	taskEntry = append(taskEntry, genTask("monitorPowerStatus", monitorPowerStatus))
	taskEntry = append(taskEntry, genTask("monitorPoEPortCfgNum", monitorPoEPortCfgNum))
	taskEntry = append(taskEntry, genTask("monitorPoEPortStatus", monitorPoEPortStatus))
	taskEntry = append(taskEntry, genTask("monitorPoEPortClass", monitorPoEPortClass))
	taskEntry = append(taskEntry, genTask("monitorPoEPortPowerConsumption", monitorPoEPortPowerConsumption))
	taskEntry = append(taskEntry, genTask("monitorPoEPortCurrent", monitorPoEPortCurrent))
	taskEntry = append(taskEntry, genTask("monitorPoEPortVoltage", monitorPoEPortVoltage))
	taskEntry = append(taskEntry, genTask("monitorPoEPortTemperature", monitorPoEPortTemperature))
	taskEntry = append(taskEntry, genTask("cpuLoadingMonitor", cpuLoadingMonitor))
	taskEntry = append(taskEntry, genTask("saveCfgMgtAction", saveCfgMgtAction))
	taskEntry = append(taskEntry, genTask("factoryDefaultAction", factoryDefaultAction))
	taskEntry = append(taskEntry, genTask("systemRebootAction", systemRebootAction))
	taskEntry = append(taskEntry, genTask("importConfiguration", importConfiguration))
	taskEntry = append(taskEntry, genTask("upgrade", upgrade))
	// here

}

func main() {
	defer func() {
		fmt.Println("test done")
	}()

	// fmt.Println("get command is", snmpGetPrefix+rmPostFix(systemContact))
	// _, b4 := utils.ShellExec(snmpGetPrefix + rmPostFix(systemContact))
	// fmt.Println("set command is", snmpSetPrefix+rmPostFix(systemContact))
	// cmd := "snmpwalk -v 3 -u walter -l authPriv -a MD5 -A 123456789 -x DES -X 123456789 192.168.15.10 1.3.6.1.4.1.37072.302.2.3.2.1.1.1.1"
	// _, setResult := utils.ShellExec(cmd)
	// fmt.Println("result is " + setResult)
	for _, val := range taskEntry {
		val.Exec()
	}
	fmt.Println("\n\n=================== Stats " + deviceIP + "=============================")
	fmt.Println("Pass:       ", stats.pass)
	fmt.Println("Failed:     ", stats.failed)
	fmt.Println("Fail marked:", stats.marked)
	if stats.marked != stats.failed {
		for _, val := range stats.unmarkedOID {
			fmt.Println(val)
		}
	}
	fmt.Println("Total:      ", stats.total)

}
