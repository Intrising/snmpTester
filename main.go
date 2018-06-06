package main

import (
	"./utils"
	"fmt"
	"strings"
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
	snmpSetTypeString = " string "
	snmpSetTypeInt    = " integer "
	snmpSetTypeIpaddr = " ipaddress "
	oidPrefix         = "1.3.6.1.4.1.37072.302.2.3."

	// no postfix means snmpget commnad
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

var oidMap map[string]*Task
var testValMap map[string]string
var taskEntry []*Task
var stats Stats

// Stats is the Statistics
type Stats struct {
	total  int
	pass   int
	failed int
}

func (s *Stats) init() {
	s.total = 0
	s.pass = 0
	s.failed = 0
}

func (s *Stats) AddPass() {
	s.pass++
	s.total++
}

func (s *Stats) AddFailed() {
	s.failed++
	s.total++
}

// Task is the task descirbed each snmp command and result
type Task struct {
	name                  string
	taskType              string // "set", "get"
	getCmd                string
	setCmd                string
	valtype               string
	defaultVal            string
	rawResult             string
	rawResultafterSet     string
	rawResultAfterDefault string
	testSuccess           string
}

func parseTaskTypeFromCmd(oid string) string {
	if strings.Contains(oid, "-w") {
		return "set"
	} else {
		return "get"
	}
}

func parseValTypeFromCmd(oid string) string {
	if strings.Contains("@i", oid) && !strings.Contains("@ip", oid) {
		return snmpSetTypeInt
	} else if strings.Contains("@s", oid) {
		return snmpSetTypeString
	} else if strings.Contains("@ip", oid) {
		return snmpSetTypeIpaddr
	}
	return "sthing wrong"
}

func rmPostFix(oid string) string {
	return strings.Split(oid, "@")[0]
}

func (t *Task) Init(taskName, oid string) {
	t.name = taskName
	t.taskType = parseTaskTypeFromCmd(oid)
	t.valtype = parseValTypeFromCmd(oid)
	oid = rmPostFix(oid)
	t.getCmd = snmpGetPrefix + oid
	// We need to test all oid by get, and some oid has the read-write access, so we store the set cmd in task.setCmd
	//
	if t.taskType == "set" {
		t.setCmd = snmpSetPrefix + oid + t.valtype + testValMap[t.valtype]
	}
}

func (t *Task) Exec() {
	_, result := utils.ShellExec(t.getCmd)
	t.rawResult = result
	t.handleFirstGet()
	t.printResult()

}

func probe(mainString, subString string) bool {
	return strings.Contains(mainString, subString)
}

func (t *Task) handleFirstGet() {
	// fmt.Println("handleRawVal")
	// fmt.Println("Raw Val is => ", t.rawResult)
	if probe(t.rawResult, "No Such Instance currently exists at this OID") {
		t.testSuccess = cross
		stats.AddFailed()
	} else if probe(t.rawResult, "STRING:") {
		val := strings.Split(t.rawResult, "STRING: ")[1]
		val = strings.Replace(val, "\"", "", -1)
		t.defaultVal = val
		t.testSuccess = check
		stats.AddPass()
	} else if probe(t.rawResult, "INTEGER:") {
		val := strings.Split(t.rawResult, "INTEGER: ")[1]
		t.defaultVal = val
		t.testSuccess = check
		stats.AddPass()
	}
}

func (t *Task) printResult() {
	fmt.Println("\n")
	fmt.Println("Name: ", t.name)
	fmt.Println("GetCmd: ", t.getCmd)
	fmt.Println("GET "+t.name+" ===== >", t.defaultVal)
	fmt.Println("Test pass: ", t.testSuccess)
}

func genTask(name, oid string) *Task {
	t := new(Task)
	t.Init(name, oid)
	return t
}

func init() {
	stats.init()

	testValMap = make(map[string]string)
	testValMap["string"] = "testWalter"
	testValMap["integer"] = "20"
	testValMap["ipaddress"] = "192.168.1.1"

	// taskEntry = append(taskEntry, genTask(oid))
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
	taskEntry = append(taskEntry, genTask("mvrMemnters", mvrMemnters))
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

}

func main() {
	defer func() {
		fmt.Println("test done")
	}()

	// fmt.Println("get command is", snmpGetPrefix+rmPostFix(systemContact))
	// _, b4 := utils.ShellExec(snmpGetPrefix + rmPostFix(systemContact))
	// fmt.Println("set command is", snmpSetPrefix+rmPostFix(systemContact))
	// _, setResult := utils.ShellExec(snmpSetPrefix + rmPostFix(systemContact) + " string " + "helloWalter")
	// fmt.Println("result is " + b4)
	// fmt.Println("result is " + setResult)
	for _, val := range taskEntry {
		fmt.Println(val.name)
		val.Exec()
	}
	fmt.Println("================================================")
	fmt.Println("Pass:   ", stats.pass)
	fmt.Println("Failed: ", stats.failed)
	fmt.Println("Total:  ", stats.total)
}
