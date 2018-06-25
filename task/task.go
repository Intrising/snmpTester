package task

import (
	"fmt"
	"strings"

	"../utils"
)

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
	failedtype            string
}

var taskEntry []*Task
var testValMap map[string]string

func (t *Task) failed(reason string) {
	t.testSuccess = cross
	t.failedReason = reason
	stats.AddFailed()
	if len(t.failedtype) > 0 {
		stats.AddMarked()
	} else {
		stats.AddunmarkedOID(t.oid)
	}
}

func (t *Task) success() {
	t.testSuccess = check
	stats.AddPass()
}

func (t *Task) init(taskName, oid string) {
	t.name = taskName
	t.taskType = parseTaskTypeFromCmd(oid)
	t.valtype = parseValTypeFromCmd(oid)
	t.failedtype = parseFailedType(oid)

	// parse oid
	oid = rmPostFix(oid)
	if strings.Contains(oid, privateMibOid) {
		// private privateMibOid
		t.oid = strings.Split(oid, privateMibOid)[1]
	} else if strings.Contains(oid, mib2Prefix) {
		// rfc4318 oid prefix
		t.oid = strings.Split(oid, mib2Prefix)[1]
	}

	// parse cmd
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

// Exec executes the task
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
		t.failed(nothingInThisOID)
		return
	}

	// Check the value type first
	if t.valtype == "string" && !probe(t.rawResult, "STRING:") {
		t.failed("Expect type string, but probe other type")
		return
	} else if t.valtype == "integer" && !probe(t.rawResult, "INTEGER:") && !probe(t.rawResult, "Counter32:") {
		t.failed("Expect type integer, but probe other type")
		return
	} else if t.valtype == "ipaddress" && !probe(t.rawResult, "IPADDRESS:") {
		t.failed("Expect type ipaddress, but probe other type")
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
			} else if probe(line, "Counter32:") {
				val := strings.Split(line, "Counter32: ")[1]
				t.defaultVal = append(t.defaultVal, val)
			}
		}
		t.success()

	} else if isGet(t) {
		if probe(t.rawResult, "STRING:") {
			val := strings.Split(t.rawResult, "STRING: ")[1]
			val = strings.Replace(val, "\"", "", -1)
			t.defaultVal = []string{val}
			t.success()
		} else if probe(t.rawResult, "INTEGER:") {
			val := strings.Split(t.rawResult, "INTEGER: ")[1]
			t.defaultVal = []string{val}
			t.success()
		} else if probe(t.rawResult, "Counter32:") {
			val := strings.Split(t.rawResult, "Counter32: ")[1]
			t.defaultVal = []string{val}
			t.success()
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

func probe(mainString, subString string) bool {
	return strings.Contains(mainString, subString)
}

func isGet(t *Task) bool {
	return t.taskType == typeGet || t.taskType == typeSet
}

func isWalk(t *Task) bool {
	return t.taskType == typeWalk || t.taskType == typeWalkSet
}

func isSet(t *Task) bool {
	return t.taskType == typeWalkSet || t.taskType == typeSet
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
	t.init(name, oid)
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
	case "4":
		return "Device not support"
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

// GetTaskEntry return the init taskEntry
func GetTaskEntry() []*Task {
	return taskEntry
}

func init() {
	testValMap = make(map[string]string)
	testValMap["string"] = "testWalter"
	testValMap["integer"] = "20"

	taskEntry = append(taskEntry, genTask("etherStatsIndex", etherStatsIndex))
	taskEntry = append(taskEntry, genTask("etherStatsDataSource", etherStatsDataSource))
	taskEntry = append(taskEntry, genTask("etherStatsDropEvents", etherStatsDropEvents))
	taskEntry = append(taskEntry, genTask("etherStatsOctets", etherStatsOctets))
	taskEntry = append(taskEntry, genTask("etherStatsPkts", etherStatsPkts))
	taskEntry = append(taskEntry, genTask("etherStatsBroadcastPkts", etherStatsBroadcastPkts))
	taskEntry = append(taskEntry, genTask("etherStatsMulticastPkts", etherStatsMulticastPkts))
	taskEntry = append(taskEntry, genTask("etherStatsCRCAlignErrors", etherStatsCRCAlignErrors))
	taskEntry = append(taskEntry, genTask("etherStatsUndersizePkts", etherStatsUndersizePkts))
	taskEntry = append(taskEntry, genTask("etherStatsOversizePkts", etherStatsOversizePkts))
	taskEntry = append(taskEntry, genTask("etherStatsFragments", etherStatsFragments))
	taskEntry = append(taskEntry, genTask("etherStatsJabbers", etherStatsJabbers))
	taskEntry = append(taskEntry, genTask("etherStatsCollisions", etherStatsCollisions))
	taskEntry = append(taskEntry, genTask("etherStatsPkts64Octets", etherStatsPkts64Octets))
	taskEntry = append(taskEntry, genTask("etherStatsPkts65to127Octets", etherStatsPkts65to127Octets))
	taskEntry = append(taskEntry, genTask("etherStatsPkts128to255Octets", etherStatsPkts128to255Octets))
	taskEntry = append(taskEntry, genTask("etherStatsPkts256to511Octets", etherStatsPkts256to511Octets))
	taskEntry = append(taskEntry, genTask("etherStatsPkts512to1023Octets", etherStatsPkts512to1023Octets))
	taskEntry = append(taskEntry, genTask("etherStatsPkts1024to1518Octets", etherStatsPkts1024to1518Octets))
	taskEntry = append(taskEntry, genTask("etherStatsOwner", etherStatsOwner))
	taskEntry = append(taskEntry, genTask("etherStatsStatus", etherStatsStatus))
	taskEntry = append(taskEntry, genTask("historyControlIndex", historyControlIndex))
	taskEntry = append(taskEntry, genTask("historyControlDataSource", historyControlDataSource))
	taskEntry = append(taskEntry, genTask("historyControlBucketsRequested", historyControlBucketsRequested))
	taskEntry = append(taskEntry, genTask("historyControlBucketsGranted", historyControlBucketsGranted))
	taskEntry = append(taskEntry, genTask("historyControlInterval", historyControlInterval))
	taskEntry = append(taskEntry, genTask("historyControlOwner", historyControlOwner))
	taskEntry = append(taskEntry, genTask("historyControlStatus", historyControlStatus))
	taskEntry = append(taskEntry, genTask("etherHistoryIndex", etherHistoryIndex))
	taskEntry = append(taskEntry, genTask("etherHistorySampleIndex", etherHistorySampleIndex))
	taskEntry = append(taskEntry, genTask("etherHistoryIntervalStart", etherHistoryIntervalStart))
	taskEntry = append(taskEntry, genTask("etherHistoryDropEvents", etherHistoryDropEvents))
	taskEntry = append(taskEntry, genTask("etherHistoryOctets", etherHistoryOctets))
	taskEntry = append(taskEntry, genTask("etherHistoryPkts", etherHistoryPkts))
	taskEntry = append(taskEntry, genTask("etherHistoryBroadcastPkts", etherHistoryBroadcastPkts))
	taskEntry = append(taskEntry, genTask("etherHistoryMulticastPkts", etherHistoryMulticastPkts))
	taskEntry = append(taskEntry, genTask("etherHistoryCRCAlignErrors", etherHistoryCRCAlignErrors))
	taskEntry = append(taskEntry, genTask("etherHistoryUndersizePkts", etherHistoryUndersizePkts))
	taskEntry = append(taskEntry, genTask("etherHistoryOversizePkts", etherHistoryOversizePkts))
	taskEntry = append(taskEntry, genTask("etherHistoryFragments", etherHistoryFragments))
	taskEntry = append(taskEntry, genTask("etherHistoryJabbers", etherHistoryJabbers))
	taskEntry = append(taskEntry, genTask("etherHistoryCollisions", etherHistoryCollisions))
	taskEntry = append(taskEntry, genTask("etherHistoryUtilization", etherHistoryUtilization))
	taskEntry = append(taskEntry, genTask("alarmIndex", alarmIndex))
	taskEntry = append(taskEntry, genTask("alarmInterval", alarmInterval))
	taskEntry = append(taskEntry, genTask("alarmVariable", alarmVariable))
	taskEntry = append(taskEntry, genTask("alarmSampleType", alarmSampleType))
	taskEntry = append(taskEntry, genTask("alarmValue", alarmValue))
	taskEntry = append(taskEntry, genTask("alarmStartupAlarm", alarmStartupAlarm))
	taskEntry = append(taskEntry, genTask("alarmRisingThreshold", alarmRisingThreshold))
	taskEntry = append(taskEntry, genTask("alarmFallingThreshold", alarmFallingThreshold))
	taskEntry = append(taskEntry, genTask("alarmRisingEventIndex", alarmRisingEventIndex))
	taskEntry = append(taskEntry, genTask("alarmFallingEventIndex", alarmFallingEventIndex))
	taskEntry = append(taskEntry, genTask("alarmOwner", alarmOwner))
	taskEntry = append(taskEntry, genTask("alarmStatus", alarmStatus))
	taskEntry = append(taskEntry, genTask("hostControlIndex", hostControlIndex))
	taskEntry = append(taskEntry, genTask("hostControlDataSource", hostControlDataSource))
	taskEntry = append(taskEntry, genTask("hostControlTableSize", hostControlTableSize))
	taskEntry = append(taskEntry, genTask("hostControlLastDeleteTime", hostControlLastDeleteTime))
	taskEntry = append(taskEntry, genTask("hostControlOwner", hostControlOwner))
	taskEntry = append(taskEntry, genTask("hostControlStatus", hostControlStatus))
	taskEntry = append(taskEntry, genTask("hostAddress", hostAddress))
	taskEntry = append(taskEntry, genTask("hostCreationOrder", hostCreationOrder))
	taskEntry = append(taskEntry, genTask("hostIndex", hostIndex))
	taskEntry = append(taskEntry, genTask("hostInPkts", hostInPkts))
	taskEntry = append(taskEntry, genTask("hostOutPkts", hostOutPkts))
	taskEntry = append(taskEntry, genTask("hostInOctets", hostInOctets))
	taskEntry = append(taskEntry, genTask("hostOutOctets", hostOutOctets))
	taskEntry = append(taskEntry, genTask("hostOutErrors", hostOutErrors))
	taskEntry = append(taskEntry, genTask("hostOutBroadcastPkts", hostOutBroadcastPkts))
	taskEntry = append(taskEntry, genTask("hostOutMulticastPkts", hostOutMulticastPkts))
	taskEntry = append(taskEntry, genTask("hostTimeAddress", hostTimeAddress))
	taskEntry = append(taskEntry, genTask("hostTimeCreationOrder", hostTimeCreationOrder))
	taskEntry = append(taskEntry, genTask("hostTimeIndex", hostTimeIndex))
	taskEntry = append(taskEntry, genTask("hostTimeInPkts", hostTimeInPkts))
	taskEntry = append(taskEntry, genTask("hostTimeOutPkts", hostTimeOutPkts))
	taskEntry = append(taskEntry, genTask("hostTimeInOctets", hostTimeInOctets))
	taskEntry = append(taskEntry, genTask("hostTimeOutOctets", hostTimeOutOctets))
	taskEntry = append(taskEntry, genTask("hostTimeOutErrors", hostTimeOutErrors))
	taskEntry = append(taskEntry, genTask("hostTimeOutBroadcastPkts", hostTimeOutBroadcastPkts))
	taskEntry = append(taskEntry, genTask("hostTimeOutMulticastPkts", hostTimeOutMulticastPkts))
	taskEntry = append(taskEntry, genTask("hostTopNControlIndex", hostTopNControlIndex))
	taskEntry = append(taskEntry, genTask("hostTopNHostIndex", hostTopNHostIndex))
	taskEntry = append(taskEntry, genTask("hostTopNRateBase", hostTopNRateBase))
	taskEntry = append(taskEntry, genTask("hostTopNTimeRemaining", hostTopNTimeRemaining))
	taskEntry = append(taskEntry, genTask("hostTopNDuration", hostTopNDuration))
	taskEntry = append(taskEntry, genTask("hostTopNRequestedSize", hostTopNRequestedSize))
	taskEntry = append(taskEntry, genTask("hostTopNGrantedSize", hostTopNGrantedSize))
	taskEntry = append(taskEntry, genTask("hostTopNStartTime", hostTopNStartTime))
	taskEntry = append(taskEntry, genTask("hostTopNOwner", hostTopNOwner))
	taskEntry = append(taskEntry, genTask("hostTopNStatus", hostTopNStatus))
	taskEntry = append(taskEntry, genTask("hostTopNReport", hostTopNReport))
	taskEntry = append(taskEntry, genTask("hostTopNIndex", hostTopNIndex))
	taskEntry = append(taskEntry, genTask("hostTopNAddress", hostTopNAddress))
	taskEntry = append(taskEntry, genTask("hostTopNRate", hostTopNRate))
	taskEntry = append(taskEntry, genTask("matrixControlIndex", matrixControlIndex))
	taskEntry = append(taskEntry, genTask("matrixControlDataSource", matrixControlDataSource))
	taskEntry = append(taskEntry, genTask("matrixControlTableSize", matrixControlTableSize))
	taskEntry = append(taskEntry, genTask("matrixControlLastDeleteTime", matrixControlLastDeleteTime))
	taskEntry = append(taskEntry, genTask("matrixControlOwner", matrixControlOwner))
	taskEntry = append(taskEntry, genTask("matrixControlStatus", matrixControlStatus))
	taskEntry = append(taskEntry, genTask("matrixSDSourceAddress", matrixSDSourceAddress))
	taskEntry = append(taskEntry, genTask("matrixSDDestAddress", matrixSDDestAddress))
	taskEntry = append(taskEntry, genTask("matrixSDIndex", matrixSDIndex))
	taskEntry = append(taskEntry, genTask("matrixSDPkts", matrixSDPkts))
	taskEntry = append(taskEntry, genTask("matrixSDOctets", matrixSDOctets))
	taskEntry = append(taskEntry, genTask("matrixSDErrors", matrixSDErrors))
	taskEntry = append(taskEntry, genTask("matrixDSSourceAddress", matrixDSSourceAddress))
	taskEntry = append(taskEntry, genTask("matrixDSDestAddress", matrixDSDestAddress))
	taskEntry = append(taskEntry, genTask("matrixDSIndex", matrixDSIndex))
	taskEntry = append(taskEntry, genTask("matrixDSPkts", matrixDSPkts))
	taskEntry = append(taskEntry, genTask("matrixDSOctets", matrixDSOctets))
	taskEntry = append(taskEntry, genTask("matrixDSErrors", matrixDSErrors))
	taskEntry = append(taskEntry, genTask("filterIndex", filterIndex))
	taskEntry = append(taskEntry, genTask("filterChannelIndex", filterChannelIndex))
	taskEntry = append(taskEntry, genTask("filterPktDataOffset", filterPktDataOffset))
	taskEntry = append(taskEntry, genTask("filterPktData", filterPktData))
	taskEntry = append(taskEntry, genTask("filterPktDataMask", filterPktDataMask))
	taskEntry = append(taskEntry, genTask("filterPktDataNotMask", filterPktDataNotMask))
	taskEntry = append(taskEntry, genTask("filterPktStatus", filterPktStatus))
	taskEntry = append(taskEntry, genTask("filterPktStatusMask", filterPktStatusMask))
	taskEntry = append(taskEntry, genTask("filterPktStatusNotMask", filterPktStatusNotMask))
	taskEntry = append(taskEntry, genTask("filterOwner", filterOwner))
	taskEntry = append(taskEntry, genTask("filterStatus", filterStatus))
	taskEntry = append(taskEntry, genTask("channelIndex", channelIndex))
	taskEntry = append(taskEntry, genTask("channelIfIndex", channelIfIndex))
	taskEntry = append(taskEntry, genTask("channelAcceptType", channelAcceptType))
	taskEntry = append(taskEntry, genTask("channelDataControl", channelDataControl))
	taskEntry = append(taskEntry, genTask("channelTurnOnEventIndex", channelTurnOnEventIndex))
	taskEntry = append(taskEntry, genTask("channelTurnOffEventIndex", channelTurnOffEventIndex))
	taskEntry = append(taskEntry, genTask("channelEventIndex", channelEventIndex))
	taskEntry = append(taskEntry, genTask("channelEventStatus", channelEventStatus))
	taskEntry = append(taskEntry, genTask("channelMatches", channelMatches))
	taskEntry = append(taskEntry, genTask("channelDescription", channelDescription))
	taskEntry = append(taskEntry, genTask("channelOwner", channelOwner))
	taskEntry = append(taskEntry, genTask("channelStatus", channelStatus))
	taskEntry = append(taskEntry, genTask("bufferControlIndex", bufferControlIndex))
	taskEntry = append(taskEntry, genTask("bufferControlChannelIndex", bufferControlChannelIndex))
	taskEntry = append(taskEntry, genTask("bufferControlFullStatus", bufferControlFullStatus))
	taskEntry = append(taskEntry, genTask("bufferControlFullAction", bufferControlFullAction))
	taskEntry = append(taskEntry, genTask("bufferControlCaptureSliceSize", bufferControlCaptureSliceSize))
	taskEntry = append(taskEntry, genTask("bufferControlDownloadSliceSize", bufferControlDownloadSliceSize))
	taskEntry = append(taskEntry, genTask("bufferControlDownloadOffset", bufferControlDownloadOffset))
	taskEntry = append(taskEntry, genTask("bufferControlMaxOctetsRequested", bufferControlMaxOctetsRequested))
	taskEntry = append(taskEntry, genTask("bufferControlMaxOctetsGranted", bufferControlMaxOctetsGranted))
	taskEntry = append(taskEntry, genTask("bufferControlCapturedPackets", bufferControlCapturedPackets))
	taskEntry = append(taskEntry, genTask("bufferControlTurnOnTime", bufferControlTurnOnTime))
	taskEntry = append(taskEntry, genTask("bufferControlOwner", bufferControlOwner))
	taskEntry = append(taskEntry, genTask("bufferControlStatus", bufferControlStatus))
	taskEntry = append(taskEntry, genTask("captureBufferControlIndex", captureBufferControlIndex))
	taskEntry = append(taskEntry, genTask("captureBufferIndex", captureBufferIndex))
	taskEntry = append(taskEntry, genTask("captureBufferPacketID", captureBufferPacketID))
	taskEntry = append(taskEntry, genTask("captureBufferPacketData", captureBufferPacketData))
	taskEntry = append(taskEntry, genTask("captureBufferPacketLength", captureBufferPacketLength))
	taskEntry = append(taskEntry, genTask("captureBufferPacketTime", captureBufferPacketTime))
	taskEntry = append(taskEntry, genTask("captureBufferPacketStatus", captureBufferPacketStatus))
	taskEntry = append(taskEntry, genTask("eventIndex", eventIndex))
	taskEntry = append(taskEntry, genTask("eventDescription", eventDescription))
	taskEntry = append(taskEntry, genTask("eventType", eventType))
	taskEntry = append(taskEntry, genTask("eventCommunity", eventCommunity))
	taskEntry = append(taskEntry, genTask("eventLastTimeSent", eventLastTimeSent))
	taskEntry = append(taskEntry, genTask("eventOwner", eventOwner))
	taskEntry = append(taskEntry, genTask("eventStatus", eventStatus))
	taskEntry = append(taskEntry, genTask("logEventIndex", logEventIndex))
	taskEntry = append(taskEntry, genTask("logIndex", logIndex))
	taskEntry = append(taskEntry, genTask("logTime", logTime))
	taskEntry = append(taskEntry, genTask("logDescription", logDescription))

	// taskEntry = append(taskEntry, genTask("mplsTunnelConfigured", mplsTunnelConfigured))
	// taskEntry = append(taskEntry, genTask("mplsTunnelActive", mplsTunnelActive))
	// taskEntry = append(taskEntry, genTask("mplsTunnelTEDistProto", mplsTunnelTEDistProto))
	// taskEntry = append(taskEntry, genTask("mplsTunnelMaxHops", mplsTunnelMaxHops))
	// taskEntry = append(taskEntry, genTask("mplsTunnelNotificationMaxRate", mplsTunnelNotificationMaxRate))

	// taskEntry = append(taskEntry, genTask("mplsTunnelIndexNext", mplsTunnelIndexNext))
	// taskEntry = append(taskEntry, genTask("mplsTunnelName", mplsTunnelName))
	// taskEntry = append(taskEntry, genTask("mplsTunnelDescr", mplsTunnelDescr))
	// taskEntry = append(taskEntry, genTask("mplsTunnelIsIf", mplsTunnelIsIf))
	// taskEntry = append(taskEntry, genTask("mplsTunnelIfIndex", mplsTunnelIfIndex))
	// taskEntry = append(taskEntry, genTask("mplsTunnelOwner", mplsTunnelOwner))
	// taskEntry = append(taskEntry, genTask("mplsTunnelRole", mplsTunnelRole))
	// taskEntry = append(taskEntry, genTask("mplsTunnelXCPointer", mplsTunnelXCPointer))
	// taskEntry = append(taskEntry, genTask("mplsTunnelSignallingProto", mplsTunnelSignallingProto))
	// taskEntry = append(taskEntry, genTask("mplsTunnelSetupPrio", mplsTunnelSetupPrio))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHoldingPrio", mplsTunnelHoldingPrio))
	// taskEntry = append(taskEntry, genTask("mplsTunnelSessionAttributes", mplsTunnelSessionAttributes))
	// taskEntry = append(taskEntry, genTask("mplsTunnelLocalProtectInUse", mplsTunnelLocalProtectInUse))
	// taskEntry = append(taskEntry, genTask("mplsTunnelResourcePointer", mplsTunnelResourcePointer))
	// taskEntry = append(taskEntry, genTask("mplsTunnelPrimaryInstance", mplsTunnelPrimaryInstance))
	// taskEntry = append(taskEntry, genTask("mplsTunnelInstancePriority", mplsTunnelInstancePriority))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopTableIndex", mplsTunnelHopTableIndex))
	// taskEntry = append(taskEntry, genTask("mplsTunnelPathInUse", mplsTunnelPathInUse))
	// taskEntry = append(taskEntry, genTask("mplsTunnelARHopTableIndex", mplsTunnelARHopTableIndex))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCHopTableIndex", mplsTunnelCHopTableIndex))
	// taskEntry = append(taskEntry, genTask("mplsTunnelIncludeAnyAffinity", mplsTunnelIncludeAnyAffinity))
	// taskEntry = append(taskEntry, genTask("mplsTunnelIncludeAllAffinity", mplsTunnelIncludeAllAffinity))
	// taskEntry = append(taskEntry, genTask("mplsTunnelExcludeAnyAffinity", mplsTunnelExcludeAnyAffinity))
	// taskEntry = append(taskEntry, genTask("mplsTunnelTotalUpTime", mplsTunnelTotalUpTime))
	// taskEntry = append(taskEntry, genTask("mplsTunnelInstanceUpTime", mplsTunnelInstanceUpTime))
	// taskEntry = append(taskEntry, genTask("mplsTunnelPrimaryUpTime", mplsTunnelPrimaryUpTime))
	// taskEntry = append(taskEntry, genTask("mplsTunnelPathChanges", mplsTunnelPathChanges))
	// taskEntry = append(taskEntry, genTask("mplsTunnelLastPathChange", mplsTunnelLastPathChange))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCreationTime", mplsTunnelCreationTime))
	// taskEntry = append(taskEntry, genTask("mplsTunnelStateTransitions", mplsTunnelStateTransitions))
	// taskEntry = append(taskEntry, genTask("mplsTunnelAdminStatus", mplsTunnelAdminStatus))
	// taskEntry = append(taskEntry, genTask("mplsTunnelOperStatus", mplsTunnelOperStatus))
	// taskEntry = append(taskEntry, genTask("mplsTunnelRowStatus", mplsTunnelRowStatus))
	// taskEntry = append(taskEntry, genTask("mplsTunnelStorageType", mplsTunnelStorageType))

	// taskEntry = append(taskEntry, genTask("mplsTunnelHopListIndexNext", mplsTunnelHopListIndexNext))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopAddrType", mplsTunnelHopAddrType))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopIPAddr", mplsTunnelHopIPAddr))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopIPPrefixLen", mplsTunnelHopIPPrefixLen))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopAsNumber", mplsTunnelHopAsNumber))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopAddrUnnum", mplsTunnelHopAddrUnnum))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopLspID", mplsTunnelHopLspID))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopType", mplsTunnelHopType))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopInclude", mplsTunnelHopInclude))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopPathOptionName", mplsTunnelHopPathOptionName))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopEntryPathComp", mplsTunnelHopEntryPathComp))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopRowStatus", mplsTunnelHopRowStatus))
	// taskEntry = append(taskEntry, genTask("mplsTunnelHopStorageType", mplsTunnelHopStorageType))
	// taskEntry = append(taskEntry, genTask("mplsTunnelResourceIndexNext", mplsTunnelResourceIndexNext))
	// taskEntry = append(taskEntry, genTask("mplsTunnelResourceMaxRate", mplsTunnelResourceMaxRate))
	// taskEntry = append(taskEntry, genTask("mplsTunnelResourceMeanRate", mplsTunnelResourceMeanRate))
	// taskEntry = append(taskEntry, genTask("mplsTunnelResourceMaxBurstSize", mplsTunnelResourceMaxBurstSize))
	// taskEntry = append(taskEntry, genTask("mplsTunnelResourceMeanBurstSize", mplsTunnelResourceMeanBurstSize))
	// taskEntry = append(taskEntry, genTask("mplsTunnelResourceExBurstSize", mplsTunnelResourceExBurstSize))
	// taskEntry = append(taskEntry, genTask("mplsTunnelResourceFrequency", mplsTunnelResourceFrequency))
	// taskEntry = append(taskEntry, genTask("mplsTunnelResourceWeight", mplsTunnelResourceWeight))
	// taskEntry = append(taskEntry, genTask("mplsTunnelResourceRowStatus", mplsTunnelResourceRowStatus))
	// taskEntry = append(taskEntry, genTask("mplsTunnelResourceStorageType", mplsTunnelResourceStorageType))
	// taskEntry = append(taskEntry, genTask("mplsTunnelARHopAddrType", mplsTunnelARHopAddrType))
	// taskEntry = append(taskEntry, genTask("mplsTunnelARHopIPAddr", mplsTunnelARHopIPAddr))
	// taskEntry = append(taskEntry, genTask("mplsTunnelARHopAddrUnnum", mplsTunnelARHopAddrUnnum))
	// taskEntry = append(taskEntry, genTask("mplsTunnelARHopLspID", mplsTunnelARHopLspID))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCHopAddrType", mplsTunnelCHopAddrType))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCHopIPAddr", mplsTunnelCHopIPAddr))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCHopIPPrefixLen", mplsTunnelCHopIPPrefixLen))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCHopAsNumber", mplsTunnelCHopAsNumber))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCHopAddrUnnum", mplsTunnelCHopAddrUnnum))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCHopLspID", mplsTunnelCHopLspID))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCHopType", mplsTunnelCHopType))
	// taskEntry = append(taskEntry, genTask("mplsTunnelPerfPackets", mplsTunnelPerfPackets))
	// taskEntry = append(taskEntry, genTask("mplsTunnelPerfHCPackets", mplsTunnelPerfHCPackets))
	// taskEntry = append(taskEntry, genTask("mplsTunnelPerfErrors", mplsTunnelPerfErrors))
	// taskEntry = append(taskEntry, genTask("mplsTunnelPerfBytes", mplsTunnelPerfBytes))
	// taskEntry = append(taskEntry, genTask("mplsTunnelPerfHCBytes", mplsTunnelPerfHCBytes))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCRLDPResMeanBurstSize", mplsTunnelCRLDPResMeanBurstSize))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCRLDPResExBurstSize", mplsTunnelCRLDPResExBurstSize))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCRLDPResFrequency", mplsTunnelCRLDPResFrequency))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCRLDPResWeight", mplsTunnelCRLDPResWeight))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCRLDPResFlags", mplsTunnelCRLDPResFlags))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCRLDPResRowStatus", mplsTunnelCRLDPResRowStatus))
	// taskEntry = append(taskEntry, genTask("mplsTunnelCRLDPResStorageType", mplsTunnelCRLDPResStorageType))
	// taskEntry = append(taskEntry, genTask("mplsTunnelNotificationEnable", mplsTunnelNotificationEnable))
	// // ************** rfc3812 **************

	// // ************** rfc3814 **************
	// taskEntry = append(taskEntry, genTask("mplsFTNIndexNext", mplsFTNIndexNext))
	// taskEntry = append(taskEntry, genTask("mplsFTNTableLastChanged", mplsFTNTableLastChanged))
	// taskEntry = append(taskEntry, genTask("mplsFTNRowStatus", mplsFTNRowStatus))
	// taskEntry = append(taskEntry, genTask("mplsFTNDescr", mplsFTNDescr))
	// taskEntry = append(taskEntry, genTask("mplsFTNMask", mplsFTNMask))
	// taskEntry = append(taskEntry, genTask("mplsFTNAddrType", mplsFTNAddrType))
	// taskEntry = append(taskEntry, genTask("mplsFTNSourceAddrMin", mplsFTNSourceAddrMin))
	// taskEntry = append(taskEntry, genTask("mplsFTNSourceAddrMax", mplsFTNSourceAddrMax))
	// taskEntry = append(taskEntry, genTask("mplsFTNDestAddrMin", mplsFTNDestAddrMin))
	// taskEntry = append(taskEntry, genTask("mplsFTNDestAddrMax", mplsFTNDestAddrMax))
	// taskEntry = append(taskEntry, genTask("mplsFTNSourcePortMin", mplsFTNSourcePortMin))
	// taskEntry = append(taskEntry, genTask("mplsFTNSourcePortMax", mplsFTNSourcePortMax))
	// taskEntry = append(taskEntry, genTask("mplsFTNDestPortMin", mplsFTNDestPortMin))
	// taskEntry = append(taskEntry, genTask("mplsFTNDestPortMax", mplsFTNDestPortMax))
	// taskEntry = append(taskEntry, genTask("mplsFTNProtocol", mplsFTNProtocol))
	// taskEntry = append(taskEntry, genTask("mplsFTNDscp", mplsFTNDscp))
	// taskEntry = append(taskEntry, genTask("mplsFTNActionType", mplsFTNActionType))
	// taskEntry = append(taskEntry, genTask("mplsFTNActionPointer", mplsFTNActionPointer))
	// taskEntry = append(taskEntry, genTask("mplsFTNStorageType", mplsFTNStorageType))

	// taskEntry = append(taskEntry, genTask("mplsFTNMapTableLastChanged", mplsFTNMapTableLastChanged))
	// taskEntry = append(taskEntry, genTask("mplsFTNMapRowStatus", mplsFTNMapRowStatus))
	// taskEntry = append(taskEntry, genTask("mplsFTNMapStorageType", mplsFTNMapStorageType))
	// taskEntry = append(taskEntry, genTask("mplsFTNPerfMatchedPackets", mplsFTNPerfMatchedPackets))
	// taskEntry = append(taskEntry, genTask("mplsFTNPerfMatchedOctets", mplsFTNPerfMatchedOctets))
	// taskEntry = append(taskEntry, genTask("mplsFTNPerfDiscontinuityTime", mplsFTNPerfDiscontinuityTime))

	// ******************** rfc 3812-14

	// rfc4750 starts
	// taskEntry = append(taskEntry, genTask("ospfRouterID", ospfRouterID))
	// taskEntry = append(taskEntry, genTask("ospfAdminStat", ospfAdminStat))
	// taskEntry = append(taskEntry, genTask("ospfVersionNumber", ospfVersionNumber))
	// taskEntry = append(taskEntry, genTask("ospfAreaBdrRtrStatus", ospfAreaBdrRtrStatus))
	// taskEntry = append(taskEntry, genTask("ospfASBdrRtrStatus", ospfASBdrRtrStatus))
	// taskEntry = append(taskEntry, genTask("ospfExternLsaCount", ospfExternLsaCount))
	// taskEntry = append(taskEntry, genTask("ospfExternLsaCksumSum", ospfExternLsaCksumSum))
	// taskEntry = append(taskEntry, genTask("ospfTOSSupport", ospfTOSSupport))
	// taskEntry = append(taskEntry, genTask("ospfOriginateNewLsas", ospfOriginateNewLsas))
	// taskEntry = append(taskEntry, genTask("ospfRxNewLsas", ospfRxNewLsas))
	// taskEntry = append(taskEntry, genTask("ospfExtLsdbLimit", ospfExtLsdbLimit))
	// taskEntry = append(taskEntry, genTask("ospfMulticastExtensions", ospfMulticastExtensions))
	// taskEntry = append(taskEntry, genTask("ospfExitOverflowInterval", ospfExitOverflowInterval))
	// taskEntry = append(taskEntry, genTask("ospfDemandExtensions", ospfDemandExtensions))
	// taskEntry = append(taskEntry, genTask("ospfRFC1583Compatibility", ospfRFC1583Compatibility))
	// taskEntry = append(taskEntry, genTask("ospfOpaqueLsaSupport", ospfOpaqueLsaSupport))
	// taskEntry = append(taskEntry, genTask("ospfReferenceBandwidth", ospfReferenceBandwidth))
	// taskEntry = append(taskEntry, genTask("ospfRestartSupport", ospfRestartSupport))
	// taskEntry = append(taskEntry, genTask("ospfRestartInterval", ospfRestartInterval))
	// taskEntry = append(taskEntry, genTask("ospfRestartStrictLsaChecking", ospfRestartStrictLsaChecking))
	// taskEntry = append(taskEntry, genTask("ospfRestartStatus", ospfRestartStatus))
	// taskEntry = append(taskEntry, genTask("ospfRestartAge", ospfRestartAge))
	// taskEntry = append(taskEntry, genTask("ospfRestartExitReason", ospfRestartExitReason))
	// taskEntry = append(taskEntry, genTask("ospfAsLsaCount", ospfAsLsaCount))
	// taskEntry = append(taskEntry, genTask("ospfAsLsaCksumSum", ospfAsLsaCksumSum))
	// taskEntry = append(taskEntry, genTask("ospfStubRouterSupport", ospfStubRouterSupport))
	// taskEntry = append(taskEntry, genTask("ospfStubRouterAdvertisement", ospfStubRouterAdvertisement))
	// taskEntry = append(taskEntry, genTask("ospfDiscontinuityTime", ospfDiscontinuityTime))
	// taskEntry = append(taskEntry, genTask("ospfAreaID", ospfAreaID))
	// taskEntry = append(taskEntry, genTask("ospfAuthType", ospfAuthType))
	// taskEntry = append(taskEntry, genTask("ospfImportAsExtern", ospfImportAsExtern))
	// taskEntry = append(taskEntry, genTask("ospfSpfRuns", ospfSpfRuns))
	// taskEntry = append(taskEntry, genTask("ospfAreaBdrRtrCount", ospfAreaBdrRtrCount))
	// taskEntry = append(taskEntry, genTask("ospfAsBdrRtrCount", ospfAsBdrRtrCount))
	// taskEntry = append(taskEntry, genTask("ospfAreaLsaCount", ospfAreaLsaCount))
	// taskEntry = append(taskEntry, genTask("ospfAreaLsaCksumSum", ospfAreaLsaCksumSum))
	// taskEntry = append(taskEntry, genTask("ospfAreaSummary", ospfAreaSummary))
	// taskEntry = append(taskEntry, genTask("ospfAreaStatus", ospfAreaStatus))
	// taskEntry = append(taskEntry, genTask("ospfAreaNssaTranslatorRole", ospfAreaNssaTranslatorRole))
	// taskEntry = append(taskEntry, genTask("ospfAreaNssaTranslatorState", ospfAreaNssaTranslatorState))
	// taskEntry = append(taskEntry, genTask("ospfAreaNssaTranslatorStabilityInterval", ospfAreaNssaTranslatorStabilityInterval))
	// taskEntry = append(taskEntry, genTask("ospfAreaNssaTranslatorEvents", ospfAreaNssaTranslatorEvents))
	// taskEntry = append(taskEntry, genTask("ospfStubAreaID", ospfStubAreaID))
	// taskEntry = append(taskEntry, genTask("ospfStubTOS", ospfStubTOS))
	// taskEntry = append(taskEntry, genTask("ospfStubMetric", ospfStubMetric))
	// taskEntry = append(taskEntry, genTask("ospfStubStatus", ospfStubStatus))
	// taskEntry = append(taskEntry, genTask("ospfStubMetricType", ospfStubMetricType))
	// taskEntry = append(taskEntry, genTask("ospfLsdbAreaID", ospfLsdbAreaID))
	// taskEntry = append(taskEntry, genTask("ospfLsdbType", ospfLsdbType))
	// taskEntry = append(taskEntry, genTask("ospfLsdbLsid", ospfLsdbLsid))
	// taskEntry = append(taskEntry, genTask("ospfLsdbRouterID", ospfLsdbRouterID))
	// taskEntry = append(taskEntry, genTask("ospfLsdbSequence", ospfLsdbSequence))
	// taskEntry = append(taskEntry, genTask("ospfLsdbAge", ospfLsdbAge))
	// taskEntry = append(taskEntry, genTask("ospfLsdbChecksum", ospfLsdbChecksum))
	// taskEntry = append(taskEntry, genTask("ospfLsdbAdvertisement", ospfLsdbAdvertisement))
	// taskEntry = append(taskEntry, genTask("ospfAreaRangeAreaID", ospfAreaRangeAreaID))
	// taskEntry = append(taskEntry, genTask("ospfAreaRangeNet", ospfAreaRangeNet))
	// taskEntry = append(taskEntry, genTask("ospfAreaRangeMask", ospfAreaRangeMask))
	// taskEntry = append(taskEntry, genTask("ospfAreaRangeStatus", ospfAreaRangeStatus))
	// taskEntry = append(taskEntry, genTask("ospfAreaRangeEffect", ospfAreaRangeEffect))
	// taskEntry = append(taskEntry, genTask("ospfHostIPAddress", ospfHostIPAddress))
	// taskEntry = append(taskEntry, genTask("ospfHostTOS", ospfHostTOS))
	// taskEntry = append(taskEntry, genTask("ospfHostMetric", ospfHostMetric))
	// taskEntry = append(taskEntry, genTask("ospfHostStatus", ospfHostStatus))
	// taskEntry = append(taskEntry, genTask("ospfHostCfgAreaID", ospfHostCfgAreaID))
	// taskEntry = append(taskEntry, genTask("ospfIfIPAddress", ospfIfIPAddress))
	// taskEntry = append(taskEntry, genTask("ospfAddressLessIf", ospfAddressLessIf))
	// taskEntry = append(taskEntry, genTask("ospfIfAreaID", ospfIfAreaID))
	// taskEntry = append(taskEntry, genTask("ospfIfType", ospfIfType))
	// taskEntry = append(taskEntry, genTask("ospfIfAdminStat", ospfIfAdminStat))
	// taskEntry = append(taskEntry, genTask("ospfIfRtrPriority", ospfIfRtrPriority))
	// taskEntry = append(taskEntry, genTask("ospfIfTransitDelay", ospfIfTransitDelay))
	// taskEntry = append(taskEntry, genTask("ospfIfRetransInterval", ospfIfRetransInterval))
	// taskEntry = append(taskEntry, genTask("ospfIfHelloInterval", ospfIfHelloInterval))
	// taskEntry = append(taskEntry, genTask("ospfIfRtrDeadInterval", ospfIfRtrDeadInterval))
	// taskEntry = append(taskEntry, genTask("ospfIfPollInterval", ospfIfPollInterval))
	// taskEntry = append(taskEntry, genTask("ospfIfState", ospfIfState))
	// taskEntry = append(taskEntry, genTask("ospfIfDesignatedRouter", ospfIfDesignatedRouter))
	// taskEntry = append(taskEntry, genTask("ospfIfBackupDesignatedRouter", ospfIfBackupDesignatedRouter))
	// taskEntry = append(taskEntry, genTask("ospfIfEvents", ospfIfEvents))
	// taskEntry = append(taskEntry, genTask("ospfIfAuthKey", ospfIfAuthKey))
	// taskEntry = append(taskEntry, genTask("ospfIfStatus", ospfIfStatus))
	// taskEntry = append(taskEntry, genTask("ospfIfMulticastForwarding", ospfIfMulticastForwarding))
	// taskEntry = append(taskEntry, genTask("ospfIfDemand", ospfIfDemand))
	// taskEntry = append(taskEntry, genTask("ospfIfAuthType", ospfIfAuthType))
	// taskEntry = append(taskEntry, genTask("ospfIfLsaCount", ospfIfLsaCount))
	// taskEntry = append(taskEntry, genTask("ospfIfLsaCksumSum", ospfIfLsaCksumSum))
	// taskEntry = append(taskEntry, genTask("ospfIfDesignatedRouterID", ospfIfDesignatedRouterID))
	// taskEntry = append(taskEntry, genTask("ospfIfBackupDesignatedRouterID", ospfIfBackupDesignatedRouterID))
	// taskEntry = append(taskEntry, genTask("ospfIfMetricIPAddress", ospfIfMetricIPAddress))
	// taskEntry = append(taskEntry, genTask("ospfIfMetricAddressLessIf", ospfIfMetricAddressLessIf))
	// taskEntry = append(taskEntry, genTask("ospfIfMetricTOS", ospfIfMetricTOS))
	// taskEntry = append(taskEntry, genTask("ospfIfMetricValue", ospfIfMetricValue))
	// taskEntry = append(taskEntry, genTask("ospfIfMetricStatus", ospfIfMetricStatus))
	// taskEntry = append(taskEntry, genTask("ospfVirtIfNeighbor", ospfVirtIfNeighbor))
	// taskEntry = append(taskEntry, genTask("ospfVirtIfTransitDelay", ospfVirtIfTransitDelay))
	// taskEntry = append(taskEntry, genTask("ospfVirtIfRetransInterval", ospfVirtIfRetransInterval))
	// taskEntry = append(taskEntry, genTask("ospfVirtIfHelloInterval", ospfVirtIfHelloInterval))
	// taskEntry = append(taskEntry, genTask("ospfVirtIfRtrDeadInterval", ospfVirtIfRtrDeadInterval))
	// taskEntry = append(taskEntry, genTask("ospfVirtIfState", ospfVirtIfState))
	// taskEntry = append(taskEntry, genTask("ospfVirtIfEvents", ospfVirtIfEvents))
	// taskEntry = append(taskEntry, genTask("ospfVirtIfAuthKey", ospfVirtIfAuthKey))
	// taskEntry = append(taskEntry, genTask("ospfVirtIfStatus", ospfVirtIfStatus))
	// taskEntry = append(taskEntry, genTask("ospfVirtIfAuthType", ospfVirtIfAuthType))
	// taskEntry = append(taskEntry, genTask("ospfVirtIfLsaCount", ospfVirtIfLsaCount))
	// taskEntry = append(taskEntry, genTask("ospfVirtIfLsaCksumSum", ospfVirtIfLsaCksumSum))
	// taskEntry = append(taskEntry, genTask("ospfNbrIPAddr", ospfNbrIPAddr))
	// taskEntry = append(taskEntry, genTask("ospfNbrAddressLessIndex", ospfNbrAddressLessIndex))
	// taskEntry = append(taskEntry, genTask("ospfNbrRtrID", ospfNbrRtrID))
	// taskEntry = append(taskEntry, genTask("ospfNbrOptions", ospfNbrOptions))
	// taskEntry = append(taskEntry, genTask("ospfNbrPriority", ospfNbrPriority))
	// taskEntry = append(taskEntry, genTask("ospfNbrState", ospfNbrState))
	// taskEntry = append(taskEntry, genTask("ospfNbrEvents", ospfNbrEvents))
	// taskEntry = append(taskEntry, genTask("ospfNbrLsRetransQLen", ospfNbrLsRetransQLen))
	// taskEntry = append(taskEntry, genTask("ospfNbmaNbrStatus", ospfNbmaNbrStatus))
	// taskEntry = append(taskEntry, genTask("ospfNbmaNbrPermanence", ospfNbmaNbrPermanence))
	// taskEntry = append(taskEntry, genTask("ospfNbrHelloSuppressed", ospfNbrHelloSuppressed))
	// taskEntry = append(taskEntry, genTask("ospfNbrRestartHelperStatus", ospfNbrRestartHelperStatus))
	// taskEntry = append(taskEntry, genTask("ospfNbrRestartHelperAge", ospfNbrRestartHelperAge))
	// taskEntry = append(taskEntry, genTask("ospfNbrRestartHelperExitReason", ospfNbrRestartHelperExitReason))
	// taskEntry = append(taskEntry, genTask("ospfVirtNbrArea", ospfVirtNbrArea))
	// taskEntry = append(taskEntry, genTask("ospfVirtNbrRtrID", ospfVirtNbrRtrID))
	// taskEntry = append(taskEntry, genTask("ospfVirtNbrIPAddr", ospfVirtNbrIPAddr))
	// taskEntry = append(taskEntry, genTask("ospfVirtNbrOptions", ospfVirtNbrOptions))
	// taskEntry = append(taskEntry, genTask("ospfVirtNbrState", ospfVirtNbrState))
	// taskEntry = append(taskEntry, genTask("ospfVirtNbrEvents", ospfVirtNbrEvents))
	// taskEntry = append(taskEntry, genTask("ospfVirtNbrLsRetransQLen", ospfVirtNbrLsRetransQLen))
	// taskEntry = append(taskEntry, genTask("ospfVirtNbrHelloSuppressed", ospfVirtNbrHelloSuppressed))
	// taskEntry = append(taskEntry, genTask("ospfVirtNbrRestartHelperStatus", ospfVirtNbrRestartHelperStatus))
	// taskEntry = append(taskEntry, genTask("ospfVirtNbrRestartHelperAge", ospfVirtNbrRestartHelperAge))
	// taskEntry = append(taskEntry, genTask("ospfVirtNbrRestartHelperExitReason", ospfVirtNbrRestartHelperExitReason))
	// taskEntry = append(taskEntry, genTask("ospfExtLsdbType", ospfExtLsdbType))
	// taskEntry = append(taskEntry, genTask("ospfExtLsdbLsid", ospfExtLsdbLsid))
	// taskEntry = append(taskEntry, genTask("ospfExtLsdbRouterID", ospfExtLsdbRouterID))
	// taskEntry = append(taskEntry, genTask("ospfExtLsdbSequence", ospfExtLsdbSequence))
	// taskEntry = append(taskEntry, genTask("ospfExtLsdbAge", ospfExtLsdbAge))
	// taskEntry = append(taskEntry, genTask("ospfExtLsdbChecksum", ospfExtLsdbChecksum))
	// taskEntry = append(taskEntry, genTask("ospfExtLsdbAdvertisement", ospfExtLsdbAdvertisement))
	// taskEntry = append(taskEntry, genTask("ospfAreaAggregateAreaID", ospfAreaAggregateAreaID))
	// taskEntry = append(taskEntry, genTask("ospfAreaAggregateLsdbType", ospfAreaAggregateLsdbType))
	// taskEntry = append(taskEntry, genTask("ospfAreaAggregateNet", ospfAreaAggregateNet))
	// taskEntry = append(taskEntry, genTask("ospfAreaAggregateMask", ospfAreaAggregateMask))
	// taskEntry = append(taskEntry, genTask("ospfAreaAggregateStatus", ospfAreaAggregateStatus))
	// taskEntry = append(taskEntry, genTask("ospfAreaAggregateEffect", ospfAreaAggregateEffect))
	// taskEntry = append(taskEntry, genTask("ospfAreaAggregateExtRouteTag", ospfAreaAggregateExtRouteTag))
	// taskEntry = append(taskEntry, genTask("ospfLocalLsdbSequence", ospfLocalLsdbSequence))
	// taskEntry = append(taskEntry, genTask("ospfLocalLsdbAge", ospfLocalLsdbAge))
	// taskEntry = append(taskEntry, genTask("ospfLocalLsdbChecksum", ospfLocalLsdbChecksum))
	// taskEntry = append(taskEntry, genTask("ospfLocalLsdbAdvertisement", ospfLocalLsdbAdvertisement))
	// taskEntry = append(taskEntry, genTask("ospfVirtLocalLsdbSequence", ospfVirtLocalLsdbSequence))
	// taskEntry = append(taskEntry, genTask("ospfVirtLocalLsdbAge", ospfVirtLocalLsdbAge))
	// taskEntry = append(taskEntry, genTask("ospfVirtLocalLsdbChecksum", ospfVirtLocalLsdbChecksum))
	// taskEntry = append(taskEntry, genTask("ospfVirtLocalLsdbAdvertisement", ospfVirtLocalLsdbAdvertisement))
	// taskEntry = append(taskEntry, genTask("ospfAsLsdbSequence", ospfAsLsdbSequence))
	// taskEntry = append(taskEntry, genTask("ospfAsLsdbAge", ospfAsLsdbAge))
	// taskEntry = append(taskEntry, genTask("ospfAsLsdbChecksum", ospfAsLsdbChecksum))
	// taskEntry = append(taskEntry, genTask("ospfAsLsdbAdvertisement", ospfAsLsdbAdvertisement))
	// taskEntry = append(taskEntry, genTask("ospfAreaLsaCountNumber", ospfAreaLsaCountNumber))
	// ************** rfc4750 ends **************

	// RFC 4318 starts
	// taskEntry = append(taskEntry, genTask("dot1dStpVersion", dot1dStpVersion))
	// taskEntry = append(taskEntry, genTask("dot1dStpTxHoldCount", dot1dStpTxHoldCount))
	// taskEntry = append(taskEntry, genTask("dot1dStpPortProtocolMigration", dot1dStpPortProtocolMigration))
	// taskEntry = append(taskEntry, genTask("dot1dStpPortAdminEdgePort", dot1dStpPortAdminEdgePort))
	// taskEntry = append(taskEntry, genTask("dot1dStpPortOperEdgePort", dot1dStpPortOperEdgePort))
	// taskEntry = append(taskEntry, genTask("dot1dStpPortAdminPointToPoint", dot1dStpPortAdminPointToPoint))
	// taskEntry = append(taskEntry, genTask("dot1dStpPortOperPointToPoint", dot1dStpPortOperPointToPoint))
	// taskEntry = append(taskEntry, genTask("dot1dStpPortAdminPathCost", dot1dStpPortAdminPathCost))

	// // no definition yet
	// taskEntry = append(taskEntry, genTask("rstpNotifications", rstpNotifications))
	// taskEntry = append(taskEntry, genTask("rstpObjects", rstpObjects))
	// RFC 4318 Ends

	// ************************* Private Mib start
	// taskEntry = append(taskEntry, genTask("systemName", systemName))
	// taskEntry = append(taskEntry, genTask("systemLocation", systemLocation))
	// taskEntry = append(taskEntry, genTask("systemContact", systemContact))
	// taskEntry = append(taskEntry, genTask("systemDescr", systemDescr))
	// taskEntry = append(taskEntry, genTask("systemFwVersion", systemFwVersion))
	// taskEntry = append(taskEntry, genTask("systemMacaddress", systemMacaddress))
	// taskEntry = append(taskEntry, genTask("systemAutoLogoutTime", systemAutoLogoutTime))
	// taskEntry = append(taskEntry, genTask("systemSerialNum", systemSerialNum))

	// taskEntry = append(taskEntry, genTask("vlanPortCfgNum", vlanPortCfgNum))
	// taskEntry = append(taskEntry, genTask("vlanMembers", vlanMembers))
	// taskEntry = append(taskEntry, genTask("vlanTags", vlanTags))
	// taskEntry = append(taskEntry, genTask("pvidCfgNum", pvidCfgNum))
	// taskEntry = append(taskEntry, genTask("vlanPvid", vlanPvid))
	// taskEntry = append(taskEntry, genTask("vlanFrameType", vlanFrameType))
	// taskEntry = append(taskEntry, genTask("mvrCfgNum", mvrCfgNum))
	// taskEntry = append(taskEntry, genTask("mvrCfgVid", mvrCfgVid))
	// taskEntry = append(taskEntry, genTask("mvrIPAddr", mvrIPAddr))
	// taskEntry = append(taskEntry, genTask("mvrMembers", mvrMembers))
	// taskEntry = append(taskEntry, genTask("igmpEnableQuerier", igmpEnableQuerier))
	// taskEntry = append(taskEntry, genTask("igmpQuerierVersion", igmpQuerierVersion))
	// taskEntry = append(taskEntry, genTask("igmpEnableSnooping", igmpEnableSnooping))
	// taskEntry = append(taskEntry, genTask("igmpEnableFloodWellKnown", igmpEnableFloodWellKnown))
	// taskEntry = append(taskEntry, genTask("igmpPortNum", igmpPortNum))
	// taskEntry = append(taskEntry, genTask("igmpRouterStatus", igmpRouterStatus))
	// taskEntry = append(taskEntry, genTask("igmpFastLeaveStatus", igmpFastLeaveStatus))
	// taskEntry = append(taskEntry, genTask("igmpVidNum", igmpVidNum))
	// taskEntry = append(taskEntry, genTask("igmpStatusQuerier", igmpStatusQuerier))
	// taskEntry = append(taskEntry, genTask("igmpQuerierTx", igmpQuerierTx))
	// taskEntry = append(taskEntry, genTask("igmpQuerierRx", igmpQuerierRx))
	// taskEntry = append(taskEntry, genTask("igmpV1Rx", igmpV1Rx))
	// taskEntry = append(taskEntry, genTask("igmpV2Rx", igmpV2Rx))
	// taskEntry = append(taskEntry, genTask("igmpV3Rx", igmpV3Rx))
	// taskEntry = append(taskEntry, genTask("igmpV2Leave", igmpV2Leave))
	// taskEntry = append(taskEntry, genTask("igmpEntriesEntryIndex", igmpEntriesEntryIndex))
	// taskEntry = append(taskEntry, genTask("igmpEntriesEntryIPAddr", igmpEntriesEntryIPAddr))
	// taskEntry = append(taskEntry, genTask("igmpEntriesEntryVID", igmpEntriesEntryVID))
	// taskEntry = append(taskEntry, genTask("igmpEntriesEntryMembers", igmpEntriesEntryMembers))
	// taskEntry = append(taskEntry, genTask("lldpPortNum", lldpPortNum))
	// taskEntry = append(taskEntry, genTask("lldpInfoContent", lldpInfoContent))
	// taskEntry = append(taskEntry, genTask("faultAlarmPowerCfgNum", faultAlarmPowerCfgNum))
	// taskEntry = append(taskEntry, genTask("faultAlarmPowerStatus", faultAlarmPowerStatus))
	// taskEntry = append(taskEntry, genTask("faultAlarmPortCfgNum", faultAlarmPortCfgNum))
	// taskEntry = append(taskEntry, genTask("faultAlarmPortLinkStatus", faultAlarmPortLinkStatus))
	// taskEntry = append(taskEntry, genTask("eventDDMEnabled", eventDDMEnabled))
	// taskEntry = append(taskEntry, genTask("eventDDMTemperatureLower", eventDDMTemperatureLower))
	// taskEntry = append(taskEntry, genTask("eventDDMTemperatureUpper", eventDDMTemperatureUpper))
	// taskEntry = append(taskEntry, genTask("eventDDMVoltageLower", eventDDMVoltageLower))
	// taskEntry = append(taskEntry, genTask("eventDDMVoltageUpper", eventDDMVoltageUpper))
	// taskEntry = append(taskEntry, genTask("eventDDMTxBiasLower", eventDDMTxBiasLower))
	// taskEntry = append(taskEntry, genTask("eventDDMTTxBiasUpper", eventDDMTTxBiasUpper))
	// taskEntry = append(taskEntry, genTask("eventDDMTxPowerLower", eventDDMTxPowerLower))
	// taskEntry = append(taskEntry, genTask("eventDDMTxPowerUpper", eventDDMTxPowerUpper))
	// taskEntry = append(taskEntry, genTask("eventDDMRxPowerLower", eventDDMRxPowerLower))
	// taskEntry = append(taskEntry, genTask("eventDDMRxPowerUpper", eventDDMRxPowerUpper))
	// taskEntry = append(taskEntry, genTask("eventMonitorEnabled", eventMonitorEnabled))
	// taskEntry = append(taskEntry, genTask("eventMonitorTemperatureLower", eventMonitorTemperatureLower))
	// taskEntry = append(taskEntry, genTask("eventMonitorTemperatureUpper", eventMonitorTemperatureUpper))
	// taskEntry = append(taskEntry, genTask("eventMonitorVoltageLower", eventMonitorVoltageLower))
	// taskEntry = append(taskEntry, genTask("eventMonitorVoltageUpper", eventMonitorVoltageUpper))
	// taskEntry = append(taskEntry, genTask("eventMonitorCurrentLower", eventMonitorCurrentLower))
	// taskEntry = append(taskEntry, genTask("eventMonitorCurrentUpper", eventMonitorCurrentUpper))
	// taskEntry = append(taskEntry, genTask("eventMonitorPowerLower", eventMonitorPowerLower))
	// taskEntry = append(taskEntry, genTask("eventMonitorPowerUpper", eventMonitorPowerUpper))
	// taskEntry = append(taskEntry, genTask("eventPOEAPortCfgNum", eventPOEAPortCfgNum))
	// taskEntry = append(taskEntry, genTask("eventPOEAPingEnabled", eventPOEAPingEnabled))
	// taskEntry = append(taskEntry, genTask("eventPOEAPingIPAddr", eventPOEAPingIPAddr))
	// taskEntry = append(taskEntry, genTask("eventPOEAPingInterval", eventPOEAPingInterval))
	// taskEntry = append(taskEntry, genTask("eventPOEAPingRetry", eventPOEAPingRetry))
	// taskEntry = append(taskEntry, genTask("eventPOEAPingReboot", eventPOEAPingReboot))
	// taskEntry = append(taskEntry, genTask("eventPOEAPingFailAction", eventPOEAPingFailAction))
	// taskEntry = append(taskEntry, genTask("localLogEnable", localLogEnable))
	// taskEntry = append(taskEntry, genTask("remoteSystemLogCfgNum", remoteSystemLogCfgNum))
	// taskEntry = append(taskEntry, genTask("remoteSystemLogHost", remoteSystemLogHost))
	// taskEntry = append(taskEntry, genTask("remoteSystemLogTag", remoteSystemLogTag))
	// taskEntry = append(taskEntry, genTask("remoteSystemLogFacility", remoteSystemLogFacility))
	// taskEntry = append(taskEntry, genTask("remoteSystemLogHostName", remoteSystemLogHostName))
	// taskEntry = append(taskEntry, genTask("emailEnable", emailEnable))
	// taskEntry = append(taskEntry, genTask("emailServerUser", emailServerUser))
	// taskEntry = append(taskEntry, genTask("emailServerPassword", emailServerPassword))
	// taskEntry = append(taskEntry, genTask("emailServerHost", emailServerHost))
	// taskEntry = append(taskEntry, genTask("emailServerSSLEnable", emailServerSSLEnable))
	// taskEntry = append(taskEntry, genTask("emailSender", emailSender))
	// taskEntry = append(taskEntry, genTask("emailSubject", emailSubject))
	// taskEntry = append(taskEntry, genTask("emailCloudEnable", emailCloudEnable))
	// taskEntry = append(taskEntry, genTask("emailReceiverCfgNum", emailReceiverCfgNum))
	// taskEntry = append(taskEntry, genTask("emailReceiverHost", emailReceiverHost))
	// taskEntry = append(taskEntry, genTask("smsEnable", smsEnable))
	// taskEntry = append(taskEntry, genTask("smsUser", smsUser))
	// taskEntry = append(taskEntry, genTask("smsPassword", smsPassword))
	// taskEntry = append(taskEntry, genTask("smsSenderText", smsSenderText))
	// taskEntry = append(taskEntry, genTask("smsReceiverCfgNum", smsReceiverCfgNum))
	// taskEntry = append(taskEntry, genTask("smsReceiverPhone", smsReceiverPhone))
	// taskEntry = append(taskEntry, genTask("snmpResponseLocale", snmpResponseLocale))
	// taskEntry = append(taskEntry, genTask("snmpCommunityCfgNum", snmpCommunityCfgNum))
	// taskEntry = append(taskEntry, genTask("snmpCommunityCfgString", snmpCommunityCfgString))
	// taskEntry = append(taskEntry, genTask("snmpCommunityCfgReadOnly", snmpCommunityCfgReadOnly))
	// taskEntry = append(taskEntry, genTask("snmpTrapCfgNum", snmpTrapCfgNum))
	// taskEntry = append(taskEntry, genTask("snmpTrapCfgCommunity", snmpTrapCfgCommunity))
	// taskEntry = append(taskEntry, genTask("snmpTrapCfgIPAddress", snmpTrapCfgIPAddress))
	// taskEntry = append(taskEntry, genTask("snmpTrapCfgVersion", snmpTrapCfgVersion))
	// taskEntry = append(taskEntry, genTask("snmpV3UserCfgNum", snmpV3UserCfgNum))
	// taskEntry = append(taskEntry, genTask("snmpV3UserCfgName", snmpV3UserCfgName))
	// taskEntry = append(taskEntry, genTask("snmpV3UserCfgSecurityLevel", snmpV3UserCfgSecurityLevel))
	// taskEntry = append(taskEntry, genTask("snmpV3UserCfgAuthProtocal", snmpV3UserCfgAuthProtocal))
	// taskEntry = append(taskEntry, genTask("snmpV3UserCfgAuthPassword", snmpV3UserCfgAuthPassword))
	// taskEntry = append(taskEntry, genTask("snmpV3UserCfgPrivProtocal", snmpV3UserCfgPrivProtocal))
	// taskEntry = append(taskEntry, genTask("snmpV3UserCfgPrivPassword", snmpV3UserCfgPrivPassword))
	// taskEntry = append(taskEntry, genTask("doutCfgNum", doutCfgNum))
	// taskEntry = append(taskEntry, genTask("doutCfgEnable", doutCfgEnable))
	// taskEntry = append(taskEntry, genTask("doutCfgAction", doutCfgAction))
	// taskEntry = append(taskEntry, genTask("deviceBootEvent", deviceBootEvent))
	// taskEntry = append(taskEntry, genTask("authenticationFailureEvent", authenticationFailureEvent))
	// taskEntry = append(taskEntry, genTask("authenticationSuccessEvent", authenticationSuccessEvent))
	// taskEntry = append(taskEntry, genTask("deviceDDMEvent", deviceDDMEvent))
	// taskEntry = append(taskEntry, genTask("devicePOEEvent", devicePOEEvent))
	// taskEntry = append(taskEntry, genTask("devicePOEBEvent", devicePOEBEvent))
	// taskEntry = append(taskEntry, genTask("ringTopologyChangeEvent", ringTopologyChangeEvent))
	// taskEntry = append(taskEntry, genTask("envMonitorEvent", envMonitorEvent))
	// taskEntry = append(taskEntry, genTask("eventPortNumber", eventPortNumber))
	// taskEntry = append(taskEntry, genTask("eventPortEventLog", eventPortEventLog))
	// taskEntry = append(taskEntry, genTask("eventPortEventsms", eventPortEventsms))
	// taskEntry = append(taskEntry, genTask("eventPortEventSMTP", eventPortEventSMTP))
	// taskEntry = append(taskEntry, genTask("eventPortEventsnmpTRAP", eventPortEventsnmpTRAP))
	// taskEntry = append(taskEntry, genTask("eventPortEventdout1", eventPortEventdout1))
	// taskEntry = append(taskEntry, genTask("eventPortEventdout2", eventPortEventdout2))
	// taskEntry = append(taskEntry, genTask("eventPowerNumber", eventPowerNumber))
	// taskEntry = append(taskEntry, genTask("eventPowerEventLog", eventPowerEventLog))
	// taskEntry = append(taskEntry, genTask("eventPowerEventsms", eventPowerEventsms))
	// taskEntry = append(taskEntry, genTask("eventPowerEventSMTP", eventPowerEventSMTP))
	// taskEntry = append(taskEntry, genTask("eventPowerEventsnmpTRAP", eventPowerEventsnmpTRAP))
	// taskEntry = append(taskEntry, genTask("eventPowerEventdout1", eventPowerEventdout1))
	// taskEntry = append(taskEntry, genTask("eventPowerEventdout2", eventPowerEventdout2))
	// taskEntry = append(taskEntry, genTask("eventDiNumber", eventDiNumber))
	// taskEntry = append(taskEntry, genTask("eventDiEventLog", eventDiEventLog))
	// taskEntry = append(taskEntry, genTask("eventDiEventsms", eventDiEventsms))
	// taskEntry = append(taskEntry, genTask("eventDiEventSMTP", eventDiEventSMTP))
	// taskEntry = append(taskEntry, genTask("eventDiEventsnmpTRAP", eventDiEventsnmpTRAP))
	// taskEntry = append(taskEntry, genTask("eventDiEventdout1", eventDiEventdout1))
	// taskEntry = append(taskEntry, genTask("eventDiEventdout2", eventDiEventdout2))
	// taskEntry = append(taskEntry, genTask("envVoltage", envVoltage))
	// taskEntry = append(taskEntry, genTask("envCurrent", envCurrent))
	// taskEntry = append(taskEntry, genTask("envWalt", envWalt))
	// taskEntry = append(taskEntry, genTask("envTemperature", envTemperature))
	// taskEntry = append(taskEntry, genTask("ddmPortNumber", ddmPortNumber))
	// taskEntry = append(taskEntry, genTask("ddmTemperatureHighAlarm", ddmTemperatureHighAlarm))
	// taskEntry = append(taskEntry, genTask("ddmTemperatureHighWarning", ddmTemperatureHighWarning))
	// taskEntry = append(taskEntry, genTask("ddmTemperatureCurrentValue", ddmTemperatureCurrentValue))
	// taskEntry = append(taskEntry, genTask("ddmTemperatureLowWarning", ddmTemperatureLowWarning))
	// taskEntry = append(taskEntry, genTask("ddmTemperatureLowAlarm", ddmTemperatureLowAlarm))
	// taskEntry = append(taskEntry, genTask("ddmVccHighAlarm", ddmVccHighAlarm))
	// taskEntry = append(taskEntry, genTask("ddmVccHighWarning", ddmVccHighWarning))
	// taskEntry = append(taskEntry, genTask("ddmVccCurrentValue", ddmVccCurrentValue))
	// taskEntry = append(taskEntry, genTask("ddmVccLowWarning", ddmVccLowWarning))
	// taskEntry = append(taskEntry, genTask("ddmVccLowAlarm", ddmVccLowAlarm))
	// taskEntry = append(taskEntry, genTask("ddmBiasHighAlarm", ddmBiasHighAlarm))
	// taskEntry = append(taskEntry, genTask("ddmBiasHighWarning", ddmBiasHighWarning))
	// taskEntry = append(taskEntry, genTask("ddmBiasCurrentValue", ddmBiasCurrentValue))
	// taskEntry = append(taskEntry, genTask("ddmBiasLowWarning", ddmBiasLowWarning))
	// taskEntry = append(taskEntry, genTask("ddmBiasLowAlarm", ddmBiasLowAlarm))
	// taskEntry = append(taskEntry, genTask("ddmTxPowerHighAlarm", ddmTxPowerHighAlarm))
	// taskEntry = append(taskEntry, genTask("ddmTxPowerHighWarning", ddmTxPowerHighWarning))
	// taskEntry = append(taskEntry, genTask("ddmTxPowerCurrentValue", ddmTxPowerCurrentValue))
	// taskEntry = append(taskEntry, genTask("ddmTxPowerLowWarning", ddmTxPowerLowWarning))
	// taskEntry = append(taskEntry, genTask("ddmTxPowerLowAlarm", ddmTxPowerLowAlarm))
	// taskEntry = append(taskEntry, genTask("ddmRxPowerHighAlarm", ddmRxPowerHighAlarm))
	// taskEntry = append(taskEntry, genTask("ddmRxPowerHighWarning", ddmRxPowerHighWarning))
	// taskEntry = append(taskEntry, genTask("ddmRxPowerCurrentValue", ddmRxPowerCurrentValue))
	// taskEntry = append(taskEntry, genTask("ddmRxPowerLowWarning", ddmRxPowerLowWarning))
	// taskEntry = append(taskEntry, genTask("ddmRxPowerLowAlarm", ddmRxPowerLowAlarm))
	// taskEntry = append(taskEntry, genTask("monitorPowerNumber", monitorPowerNumber))
	// taskEntry = append(taskEntry, genTask("monitorPowerStatus", monitorPowerStatus))
	// taskEntry = append(taskEntry, genTask("monitorPoEPortCfgNum", monitorPoEPortCfgNum))
	// taskEntry = append(taskEntry, genTask("monitorPoEPortStatus", monitorPoEPortStatus))
	// taskEntry = append(taskEntry, genTask("monitorPoEPortClass", monitorPoEPortClass))
	// taskEntry = append(taskEntry, genTask("monitorPoEPortPowerConsumption", monitorPoEPortPowerConsumption))
	// taskEntry = append(taskEntry, genTask("monitorPoEPortCurrent", monitorPoEPortCurrent))
	// taskEntry = append(taskEntry, genTask("monitorPoEPortVoltage", monitorPoEPortVoltage))
	// taskEntry = append(taskEntry, genTask("monitorPoEPortTemperature", monitorPoEPortTemperature))
	// taskEntry = append(taskEntry, genTask("cpuLoadingMonitor", cpuLoadingMonitor))
	// taskEntry = append(taskEntry, genTask("saveCfgMgtAction", saveCfgMgtAction))
	// taskEntry = append(taskEntry, genTask("factoryDefaultAction", factoryDefaultAction))
	// taskEntry = append(taskEntry, genTask("systemRebootAction", systemRebootAction))
	// taskEntry = append(taskEntry, genTask("importConfiguration", importConfiguration))
	// taskEntry = append(taskEntry, genTask("upgrade", upgrade))
	// ************************* Private Mib End
	// here
}
