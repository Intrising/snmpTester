package task

const (
	check            = "\u2714"
	cross            = "\u274c"
	nothingInThisOID = "No Such Instance currently exists at this OID"
	portAmount       = 20

	typeGet     = 0
	typeWalk    = 1
	typeSet     = 2
	typeWalkSet = 3

	snmpGet  = "snmpget -v 3 -t 10 "
	snmpWalk = "snmpwalk -v 3 -t 10 "
	snmpSet  = "snmpset -v 3 "

	snmpUser              = "-u walter "
	snmpSecurityLevel     = "-l authPriv "
	snmpAuthentication    = "-a MD5 "
	snmpAuthPassPhrase    = "-A 123456789 "
	snmpPrivateProtocol   = "-x DES "
	snmpPrivatePassPhrase = "-X 123456789 "

	deviceIP = "192.168.17.100 "
	// deviceIP = "192.168.15.10 "
	// deviceIP = "192.168.16.140 "

	snmpGetPrefix     = snmpGet + snmpUser + snmpSecurityLevel + snmpAuthentication + snmpAuthPassPhrase + snmpPrivateProtocol + snmpPrivatePassPhrase + deviceIP
	snmpWalkPrefix    = snmpWalk + snmpUser + snmpSecurityLevel + snmpAuthentication + snmpAuthPassPhrase + snmpPrivateProtocol + snmpPrivatePassPhrase + deviceIP
	snmpSetPrefix     = snmpSet + snmpUser + snmpSecurityLevel + snmpAuthentication + snmpAuthPassPhrase + snmpPrivateProtocol + snmpPrivatePassPhrase + deviceIP
	snmpSetTypeString = " string "
	snmpSetTypeInt    = " integer "
	snmpSetTypeIpaddr = " ipaddress "
	privateMibOid     = "1.3.6.1.4.1.37072.302.3.1."
	oldPrivateMibOid  = "1.3.6.1.4.1.37072.302.2.3."
	mib2Prefix        = "1.3.6.1.2.1."
	rfc4318           = ""

	// no postfix means snmpget commnad
	// all command can be read so no postfix for read-only
	// -s   means read-write
	// -w-s means use cmd walk and read-write
	//
	// ATTENTION: The type here is not the exact type in the MIB, but the type from snmpMIB program
	// @s   means type string
	// @i   means type integer
	// @c   means type Counter32
	// ATTENTION: The type here is not the exact type in the MIB, but the type from snmpMIB program
	//
	// -ps  means ps
	// -m:  marks the known failed type:
	//                                  1 =>  need to set value to the device
	//                                  2 =>  Spec issue
	//                                  3 =>  snmp program issue
	//                                  4 =>  Device not support

	// ************** rfc2819 Begin **************
	// etherStatsTable (16.1.1)
	// etherStatsEntry (16.1.1.1)
	etherStatsIndex                = mib2Prefix + "16.1.1.1.1@i-w"
	etherStatsDataSource           = mib2Prefix + "16.1.1.1.2@s-w-s-m:4"
	etherStatsDropEvents           = mib2Prefix + "16.1.1.1.3@i-w"
	etherStatsOctets               = mib2Prefix + "16.1.1.1.4@i-w"
	etherStatsPkts                 = mib2Prefix + "16.1.1.1.5@i-w"
	etherStatsBroadcastPkts        = mib2Prefix + "16.1.1.1.6@i-w"
	etherStatsMulticastPkts        = mib2Prefix + "16.1.1.1.7@i-w"
	etherStatsCRCAlignErrors       = mib2Prefix + "16.1.1.1.8@i-w"
	etherStatsUndersizePkts        = mib2Prefix + "16.1.1.1.9@i-w"
	etherStatsOversizePkts         = mib2Prefix + "16.1.1.1.10@i-w"
	etherStatsFragments            = mib2Prefix + "16.1.1.1.11@i-w"
	etherStatsJabbers              = mib2Prefix + "16.1.1.1.12@i-w"
	etherStatsCollisions           = mib2Prefix + "16.1.1.1.13@i-w"
	etherStatsPkts64Octets         = mib2Prefix + "16.1.1.1.14@i-w"
	etherStatsPkts65to127Octets    = mib2Prefix + "16.1.1.1.15@i-w"
	etherStatsPkts128to255Octets   = mib2Prefix + "16.1.1.1.16@i-w"
	etherStatsPkts256to511Octets   = mib2Prefix + "16.1.1.1.17@i-w"
	etherStatsPkts512to1023Octets  = mib2Prefix + "16.1.1.1.18@i-w"
	etherStatsPkts1024to1518Octets = mib2Prefix + "16.1.1.1.19@i-w"
	etherStatsOwner                = mib2Prefix + "16.1.1.1.20@s-w-s"
	etherStatsStatus               = mib2Prefix + "16.1.1.1.21@i-w-s"

	// historyControlTable (16.2.1)
	// historyControlEntry (16.2.1.1)
	historyControlIndex            = mib2Prefix + "16.2.1.1.1@i-w"
	historyControlDataSource       = mib2Prefix + "16.2.1.1.2@s-w-s-m:4"
	historyControlBucketsRequested = mib2Prefix + "16.2.1.1.3@i-w-s"
	historyControlBucketsGranted   = mib2Prefix + "16.2.1.1.4@i-w"
	historyControlInterval         = mib2Prefix + "16.2.1.1.5@i-w-s"
	historyControlOwner            = mib2Prefix + "16.2.1.1.6@s-w-s"
	historyControlStatus           = mib2Prefix + "16.2.1.1.7@i-w-s"

	// etherHistoryTable (16.2.2)
	// etherHistoryEntry (16.2.2.1)
	etherHistoryIndex          = mib2Prefix + "16.2.2.1.1@i-w"
	etherHistorySampleIndex    = mib2Prefix + "16.2.2.1.2@i-w"
	etherHistoryIntervalStart  = mib2Prefix + "16.2.2.1.3@i-w"
	etherHistoryDropEvents     = mib2Prefix + "16.2.2.1.4@i-w"
	etherHistoryOctets         = mib2Prefix + "16.2.2.1.5@i-w"
	etherHistoryPkts           = mib2Prefix + "16.2.2.1.6@i-w"
	etherHistoryBroadcastPkts  = mib2Prefix + "16.2.2.1.7@i-w"
	etherHistoryMulticastPkts  = mib2Prefix + "16.2.2.1.8@i-w"
	etherHistoryCRCAlignErrors = mib2Prefix + "16.2.2.1.9@i-w"
	etherHistoryUndersizePkts  = mib2Prefix + "16.2.2.1.10@i-w"
	etherHistoryOversizePkts   = mib2Prefix + "16.2.2.1.11@i-w"
	etherHistoryFragments      = mib2Prefix + "16.2.2.1.12@i-w"
	etherHistoryJabbers        = mib2Prefix + "16.2.2.1.13@i-w"
	etherHistoryCollisions     = mib2Prefix + "16.2.2.1.14@i-w"
	etherHistoryUtilization    = mib2Prefix + "16.2.2.1.15@i-w"

	// alarmTable (16.3.1)
	// alarmEntry (16.3.1.1)
	alarmIndex             = mib2Prefix + "16.3.1.1.1@i-w-s"
	alarmInterval          = mib2Prefix + "16.3.1.1.2@i-w-s"
	alarmVariable          = mib2Prefix + "16.3.1.1.3@s-w-s-m:4"
	alarmSampleType        = mib2Prefix + "16.3.1.1.4@i-w-s"
	alarmValue             = mib2Prefix + "16.3.1.1.5@i-w"
	alarmStartupAlarm      = mib2Prefix + "16.3.1.1.6@i-w-s"
	alarmRisingThreshold   = mib2Prefix + "16.3.1.1.7@i-w-s"
	alarmFallingThreshold  = mib2Prefix + "16.3.1.1.8@i-w-s"
	alarmRisingEventIndex  = mib2Prefix + "16.3.1.1.9@i-w-s"
	alarmFallingEventIndex = mib2Prefix + "16.3.1.1.10@i-w-s"
	alarmOwner             = mib2Prefix + "16.3.1.1.11@s-w-s-m:4"
	alarmStatus            = mib2Prefix + "16.3.1.1.12@i-w-s"

	// hostControlTable (16.4.1)
	// hostControlEntry (16.4.1.1)
	hostControlIndex          = mib2Prefix + "16.4.1.1.1@i-w"
	hostControlDataSource     = mib2Prefix + "16.4.1.1.2@s-w-s-m:4"
	hostControlTableSize      = mib2Prefix + "16.4.1.1.3@i-w"
	hostControlLastDeleteTime = mib2Prefix + "16.4.1.1.4@i-w"
	hostControlOwner          = mib2Prefix + "16.4.1.1.5@s-w-s-m:4"
	hostControlStatus         = mib2Prefix + "16.4.1.1.6@i-w-s"

	// hostTable (16.4.2)
	// hostEntry (16.4.2.1)
	hostAddress          = mib2Prefix + "16.4.2.1.1@s-w-m:4"
	hostCreationOrder    = mib2Prefix + "16.4.2.1.2@i-w"
	hostIndex            = mib2Prefix + "16.4.2.1.3@i-w"
	hostInPkts           = mib2Prefix + "16.4.2.1.4@i-w"
	hostOutPkts          = mib2Prefix + "16.4.2.1.5@i-w"
	hostInOctets         = mib2Prefix + "16.4.2.1.6@i-w"
	hostOutOctets        = mib2Prefix + "16.4.2.1.7@i-w"
	hostOutErrors        = mib2Prefix + "16.4.2.1.8@i-w"
	hostOutBroadcastPkts = mib2Prefix + "16.4.2.1.9@i-w"
	hostOutMulticastPkts = mib2Prefix + "16.4.2.1.10@i-w"

	// hostTimeTable (16.4.3)
	// hostTimeEntry (16.4.3.1)
	hostTimeAddress          = mib2Prefix + "16.4.3.1.1@s-w-m:4"
	hostTimeCreationOrder    = mib2Prefix + "16.4.3.1.2@i-w"
	hostTimeIndex            = mib2Prefix + "16.4.3.1.3@i-w"
	hostTimeInPkts           = mib2Prefix + "16.4.3.1.4@i-w"
	hostTimeOutPkts          = mib2Prefix + "16.4.3.1.5@i-w"
	hostTimeInOctets         = mib2Prefix + "16.4.3.1.6@i-w"
	hostTimeOutOctets        = mib2Prefix + "16.4.3.1.7@i-w"
	hostTimeOutErrors        = mib2Prefix + "16.4.3.1.8@i-w"
	hostTimeOutBroadcastPkts = mib2Prefix + "16.4.3.1.9@i-w"
	hostTimeOutMulticastPkts = mib2Prefix + "16.4.3.1.10@i-w"

	// hostTopNControlTable (16.5.1)
	// hostTopNControlEntry (16.5.1.1)
	hostTopNControlIndex  = mib2Prefix + "16.5.1.1.1@i-w"
	hostTopNHostIndex     = mib2Prefix + "16.5.1.1.2@i-w-s"
	hostTopNRateBase      = mib2Prefix + "16.5.1.1.3@i-w-s"
	hostTopNTimeRemaining = mib2Prefix + "16.5.1.1.4@i-w-s"
	hostTopNDuration      = mib2Prefix + "16.5.1.1.5@i-w"
	hostTopNRequestedSize = mib2Prefix + "16.5.1.1.6@i-w"
	hostTopNGrantedSize   = mib2Prefix + "16.5.1.1.7@i-w"
	hostTopNStartTime     = mib2Prefix + "16.5.1.1.8@i-w"
	hostTopNOwner         = mib2Prefix + "16.5.1.1.9@s-w-s-m:4"
	hostTopNStatus        = mib2Prefix + "16.5.1.1.10@i-w-s"

	// hostTopNTable (16.5.2)
	// hostTopNEntry (16.5.2.1)
	hostTopNReport  = mib2Prefix + "16.5.2.1.1@i-w"
	hostTopNIndex   = mib2Prefix + "16.5.2.1.2@i-w"
	hostTopNAddress = mib2Prefix + "16.5.2.1.3@s-w-m:4"
	hostTopNRate    = mib2Prefix + "16.5.2.1.4@i-w"

	// matrixControlTable (16.6.1)
	// matrixControlEntry (16.6.1.1)
	matrixControlIndex          = mib2Prefix + "16.6.1.1.1@i-w"
	matrixControlDataSource     = mib2Prefix + "16.6.1.1.2@s-w-m:4"
	matrixControlTableSize      = mib2Prefix + "16.6.1.1.3@i-w"
	matrixControlLastDeleteTime = mib2Prefix + "16.6.1.1.4@i-w"
	matrixControlOwner          = mib2Prefix + "16.6.1.1.5@s-w-m:4"
	matrixControlStatus         = mib2Prefix + "16.6.1.1.6@i-w"

	// matrixSDTable (16.6.2)
	// matrixSDEntry (16.6.2.1)
	matrixSDSourceAddress = mib2Prefix + "16.6.2.1.1@s-w-m:4"
	matrixSDDestAddress   = mib2Prefix + "16.6.2.1.2@s-w-m:4"
	matrixSDIndex         = mib2Prefix + "16.6.2.1.3@i-w"
	matrixSDPkts          = mib2Prefix + "16.6.2.1.4@i-w"
	matrixSDOctets        = mib2Prefix + "16.6.2.1.5@i-w"
	matrixSDErrors        = mib2Prefix + "16.6.2.1.6@i-w"

	// matrixDSTable (16.6.3)
	// matrixDSEntry (16.6.3.1)
	matrixDSSourceAddress = mib2Prefix + "16.6.3.1.1@s-w-m:4"
	matrixDSDestAddress   = mib2Prefix + "16.6.3.1.2@s-w-m:4"
	matrixDSIndex         = mib2Prefix + "16.6.3.1.3@i-w"
	matrixDSPkts          = mib2Prefix + "16.6.3.1.4@i-w"
	matrixDSOctets        = mib2Prefix + "16.6.3.1.5@i-w"
	matrixDSErrors        = mib2Prefix + "16.6.3.1.6@i-w"

	// filterTable (16.7.1)
	// filterEntry (16.7.1.1)
	filterIndex            = mib2Prefix + "16.7.1.1.1@i-w"
	filterChannelIndex     = mib2Prefix + "16.7.1.1.2@i-w-s"
	filterPktDataOffset    = mib2Prefix + "16.7.1.1.3@i-w-s"
	filterPktData          = mib2Prefix + "16.7.1.1.4@s-w-s-m:4"
	filterPktDataMask      = mib2Prefix + "16.7.1.1.5@s-w-s-m:4"
	filterPktDataNotMask   = mib2Prefix + "16.7.1.1.6@s-w-s-m:4"
	filterPktStatus        = mib2Prefix + "16.7.1.1.7@i-w-s"
	filterPktStatusMask    = mib2Prefix + "16.7.1.1.8@i-w-s"
	filterPktStatusNotMask = mib2Prefix + "16.7.1.1.9@i-w-s"
	filterOwner            = mib2Prefix + "16.7.1.1.10@s-w-s-m:4"
	filterStatus           = mib2Prefix + "16.7.1.1.11@i-w-s"

	// channelTable (16.7.2)
	// channelEntry (16.7.2.1)
	channelIndex             = mib2Prefix + "16.7.2.1.1@i-w"
	channelIfIndex           = mib2Prefix + "16.7.2.1.2@i-w-s"
	channelAcceptType        = mib2Prefix + "16.7.2.1.3@i-w-s"
	channelDataControl       = mib2Prefix + "16.7.2.1.4@i-w-s"
	channelTurnOnEventIndex  = mib2Prefix + "16.7.2.1.5@i-w-s"
	channelTurnOffEventIndex = mib2Prefix + "16.7.2.1.6@i-w-s"
	channelEventIndex        = mib2Prefix + "16.7.2.1.7@i-w-s"
	channelEventStatus       = mib2Prefix + "16.7.2.1.8@i-w-s"
	channelMatches           = mib2Prefix + "16.7.2.1.9@i-w"
	channelDescription       = mib2Prefix + "16.7.2.1.10@s-w-s-m:4"
	channelOwner             = mib2Prefix + "16.7.2.1.11@s-w-s-m:4"
	channelStatus            = mib2Prefix + "16.7.2.1.12@i-w-s"

	// bufferControlTable (16.8.1)
	// bufferControlEntry (16.8.1.1)
	bufferControlIndex              = mib2Prefix + "16.8.1.1.1@i-w"
	bufferControlChannelIndex       = mib2Prefix + "16.8.1.1.2@i-w-s"
	bufferControlFullStatus         = mib2Prefix + "16.8.1.1.3@i-w"
	bufferControlFullAction         = mib2Prefix + "16.8.1.1.4@i-w-s"
	bufferControlCaptureSliceSize   = mib2Prefix + "16.8.1.1.5@i-w-s"
	bufferControlDownloadSliceSize  = mib2Prefix + "16.8.1.1.6@i-w-s"
	bufferControlDownloadOffset     = mib2Prefix + "16.8.1.1.7@i-w-s"
	bufferControlMaxOctetsRequested = mib2Prefix + "16.8.1.1.8@i-w-s"
	bufferControlMaxOctetsGranted   = mib2Prefix + "16.8.1.1.9@i-w"
	bufferControlCapturedPackets    = mib2Prefix + "16.8.1.1.10@i-w"
	bufferControlTurnOnTime         = mib2Prefix + "16.8.1.1.11@i-w"
	bufferControlOwner              = mib2Prefix + "16.8.1.1.12@s-w-s-m:4"
	bufferControlStatus             = mib2Prefix + "16.8.1.1.13@i-w-s"

	// captureBufferTable (16.8.2)
	// captureBufferEntry (16.8.2.1)
	captureBufferControlIndex = mib2Prefix + "16.8.2.1.1@i-w"
	captureBufferIndex        = mib2Prefix + "16.8.2.1.2@i-w"
	captureBufferPacketID     = mib2Prefix + "16.8.2.1.3@i-w"
	captureBufferPacketData   = mib2Prefix + "16.8.2.1.4@s-w-m:4"
	captureBufferPacketLength = mib2Prefix + "16.8.2.1.5@i-w"
	captureBufferPacketTime   = mib2Prefix + "16.8.2.1.6@i-w"
	captureBufferPacketStatus = mib2Prefix + "16.8.2.1.7@i-w"

	// eventTable (16.9.1)
	// eventEntry (16.9.1.1)
	eventIndex        = mib2Prefix + "16.9.1.1.1@i-w"
	eventDescription  = mib2Prefix + "16.9.1.1.2@s-w-s-m:4"
	eventType         = mib2Prefix + "16.9.1.1.3@i-w-s"
	eventCommunity    = mib2Prefix + "16.9.1.1.4@s-w-s-m:4"
	eventLastTimeSent = mib2Prefix + "16.9.1.1.5@i-w"
	eventOwner        = mib2Prefix + "16.9.1.1.6@s-w-s-m:4"
	eventStatus       = mib2Prefix + "16.9.1.1.7@i-w-s"

	// logTable (16.9.2)
	// logEntry (16.9.2.1)
	logEventIndex  = mib2Prefix + "16.9.2.1.1@i-w"
	logIndex       = mib2Prefix + "16.9.2.1.2@i-w"
	logTime        = mib2Prefix + "16.9.2.1.3@i-w"
	logDescription = mib2Prefix + "16.9.2.1.4@s-w-m:4"

	// ************** rfc2819 End **************

	// ************** rfc4750 **************
	// ospfGeneralGroup (14.1)
	ospfRouterID                 = mib2Prefix + "14.1.1@s-s-m:4"
	ospfAdminStat                = mib2Prefix + "14.1.2@i-s"
	ospfVersionNumber            = mib2Prefix + "14.1.3@i"
	ospfAreaBdrRtrStatus         = mib2Prefix + "14.1.4@i"
	ospfASBdrRtrStatus           = mib2Prefix + "14.1.5@i-s"
	ospfExternLsaCount           = mib2Prefix + "14.1.6@i"
	ospfExternLsaCksumSum        = mib2Prefix + "14.1.7@i"
	ospfTOSSupport               = mib2Prefix + "14.1.8@i-s"
	ospfOriginateNewLsas         = mib2Prefix + "14.1.9@i"
	ospfRxNewLsas                = mib2Prefix + "14.1.10@i"
	ospfExtLsdbLimit             = mib2Prefix + "14.1.11@i-s"
	ospfMulticastExtensions      = mib2Prefix + "14.1.12@i-s"
	ospfExitOverflowInterval     = mib2Prefix + "14.1.13@i-s"
	ospfDemandExtensions         = mib2Prefix + "14.1.14@i-s"
	ospfRFC1583Compatibility     = mib2Prefix + "14.1.15@i-s"
	ospfOpaqueLsaSupport         = mib2Prefix + "14.1.16@i"
	ospfReferenceBandwidth       = mib2Prefix + "14.1.17@i-s"
	ospfRestartSupport           = mib2Prefix + "14.1.18@i-s"
	ospfRestartInterval          = mib2Prefix + "14.1.19@i-s"
	ospfRestartStrictLsaChecking = mib2Prefix + "14.1.20@i-s"
	ospfRestartStatus            = mib2Prefix + "14.1.21@i"
	ospfRestartAge               = mib2Prefix + "14.1.22@i"
	ospfRestartExitReason        = mib2Prefix + "14.1.23@i"
	ospfAsLsaCount               = mib2Prefix + "14.1.24@i"
	ospfAsLsaCksumSum            = mib2Prefix + "14.1.25@i"
	ospfStubRouterSupport        = mib2Prefix + "14.1.26@i"
	ospfStubRouterAdvertisement  = mib2Prefix + "14.1.27@i-s"
	ospfDiscontinuityTime        = mib2Prefix + "14.1.28@i"

	// ospfAreaTable (14.2)
	// ospfAreaEntry (not-accessible)(14.2.1)
	ospfAreaID                              = mib2Prefix + "14.2.1.1@s-m:4"
	ospfAuthType                            = mib2Prefix + "14.2.1.2@i-s"
	ospfImportAsExtern                      = mib2Prefix + "14.2.1.3@i-s"
	ospfSpfRuns                             = mib2Prefix + "14.2.1.4@i"
	ospfAreaBdrRtrCount                     = mib2Prefix + "14.2.1.5@i"
	ospfAsBdrRtrCount                       = mib2Prefix + "14.2.1.6@i"
	ospfAreaLsaCount                        = mib2Prefix + "14.2.1.7@i"
	ospfAreaLsaCksumSum                     = mib2Prefix + "14.2.1.8@i"
	ospfAreaSummary                         = mib2Prefix + "14.2.1.9@i-s"
	ospfAreaStatus                          = mib2Prefix + "14.2.1.10@i-s"
	ospfAreaNssaTranslatorRole              = mib2Prefix + "14.2.1.11@i-s"
	ospfAreaNssaTranslatorState             = mib2Prefix + "14.2.1.12@i"
	ospfAreaNssaTranslatorStabilityInterval = mib2Prefix + "14.2.1.13@i-s"
	ospfAreaNssaTranslatorEvents            = mib2Prefix + "14.2.1.14@i"

	// ospfStubAreaTable (14.3)
	// ospfStubAreaEntry (not-accessible)(14.3.1)
	ospfStubAreaID     = mib2Prefix + "14.3.1.1@s-m:4"
	ospfStubTOS        = mib2Prefix + "14.3.1.2@i"
	ospfStubMetric     = mib2Prefix + "14.3.1.3@i"
	ospfStubStatus     = mib2Prefix + "14.3.1.4@i-s"
	ospfStubMetricType = mib2Prefix + "14.3.1.5@i-s"

	// ospfLsdbTable (14.4)
	// ospfLsdbEntry (not-accessible)(14.4.1)
	ospfLsdbAreaID        = mib2Prefix + "14.4.1.1@s-m:4"
	ospfLsdbType          = mib2Prefix + "14.4.1.2@i"
	ospfLsdbLsid          = mib2Prefix + "14.4.1.3@s-m:4"
	ospfLsdbRouterID      = mib2Prefix + "14.4.1.4@s-m:4"
	ospfLsdbSequence      = mib2Prefix + "14.4.1.5@i"
	ospfLsdbAge           = mib2Prefix + "14.4.1.6@i"
	ospfLsdbChecksum      = mib2Prefix + "14.4.1.7@i"
	ospfLsdbAdvertisement = mib2Prefix + "14.4.1.8@s-m:4"

	// ospfAreaRangeTable (14.5)
	// ospfAreaRangeEntry (not-accessible)(14.5.1)
	ospfAreaRangeAreaID = mib2Prefix + "14.5.1.1@s-m:4"
	ospfAreaRangeNet    = mib2Prefix + "14.5.1.2@s-m:4"
	ospfAreaRangeMask   = mib2Prefix + "14.5.1.3@s-m:4"
	ospfAreaRangeStatus = mib2Prefix + "14.5.1.4@i-s"
	ospfAreaRangeEffect = mib2Prefix + "14.5.1.5@i-s"

	// ospfHostTable (14.6)
	// ospfHostEntry (not-accessible)(14.6.1)
	ospfHostIPAddress = mib2Prefix + "14.6.1.1@s-m:4"
	ospfHostTOS       = mib2Prefix + "14.6.1.2@i"
	ospfHostMetric    = mib2Prefix + "14.6.1.3@i"
	ospfHostStatus    = mib2Prefix + "14.6.1.4@i-s"
	// ospfHostAreaID    = mib2Prefix + "14.6.1.5@s"  => Deprecated
	ospfHostCfgAreaID = mib2Prefix + "14.6.1.6@s-s-m:4"

	// ospfIfTable (14.7)
	// ospfIfEntry (not-accessible)(14.7.1)
	ospfIfIPAddress                = mib2Prefix + "14.7.1.1@s-m:4"
	ospfAddressLessIf              = mib2Prefix + "14.7.1.2@i"
	ospfIfAreaID                   = mib2Prefix + "14.7.1.3@s-s-m:4"
	ospfIfType                     = mib2Prefix + "14.7.1.4@i-s"
	ospfIfAdminStat                = mib2Prefix + "14.7.1.5@i-s"
	ospfIfRtrPriority              = mib2Prefix + "14.7.1.6@i-s"
	ospfIfTransitDelay             = mib2Prefix + "14.7.1.7@i-s"
	ospfIfRetransInterval          = mib2Prefix + "14.7.1.8@i-s"
	ospfIfHelloInterval            = mib2Prefix + "14.7.1.9@i-s"
	ospfIfRtrDeadInterval          = mib2Prefix + "14.7.1.10@i-s"
	ospfIfPollInterval             = mib2Prefix + "14.7.1.11@i-s"
	ospfIfState                    = mib2Prefix + "14.7.1.12@i"
	ospfIfDesignatedRouter         = mib2Prefix + "14.7.1.13@s-m:4"
	ospfIfBackupDesignatedRouter   = mib2Prefix + "14.7.1.14@s-m:4"
	ospfIfEvents                   = mib2Prefix + "14.7.1.15@i"
	ospfIfAuthKey                  = mib2Prefix + "14.7.1.16@s-s"
	ospfIfStatus                   = mib2Prefix + "14.7.1.17@i-s"
	ospfIfMulticastForwarding      = mib2Prefix + "14.7.1.18@i-s"
	ospfIfDemand                   = mib2Prefix + "14.7.1.19@i-s"
	ospfIfAuthType                 = mib2Prefix + "14.7.1.20@i-s"
	ospfIfLsaCount                 = mib2Prefix + "14.7.1.21@i"
	ospfIfLsaCksumSum              = mib2Prefix + "14.7.1.22@i"
	ospfIfDesignatedRouterID       = mib2Prefix + "14.7.1.23@s-m:4"
	ospfIfBackupDesignatedRouterID = mib2Prefix + "14.7.1.24@s-m:4"

	// ospfIfMetricTable (14.8)
	// ospfIfMetricEntry (not-accessible)(14.8.1)
	ospfIfMetricIPAddress     = mib2Prefix + "14.8.1.1@s"
	ospfIfMetricAddressLessIf = mib2Prefix + "14.8.1.2@i"
	ospfIfMetricTOS           = mib2Prefix + "14.8.1.3@i"
	ospfIfMetricValue         = mib2Prefix + "14.8.1.4@i-s"
	ospfIfMetricStatus        = mib2Prefix + "14.8.1.5@i-s"

	// ospfVirtIfTable (14.9)
	// ospfVirtIfEntry (not-accessible)(14.9.1)
	ospfVirtIfAreaID          = mib2Prefix + "14.9.1.1@s"
	ospfVirtIfNeighbor        = mib2Prefix + "14.9.1.2@s"
	ospfVirtIfTransitDelay    = mib2Prefix + "14.9.1.3@i-s"
	ospfVirtIfRetransInterval = mib2Prefix + "14.9.1.4@i-s"
	ospfVirtIfHelloInterval   = mib2Prefix + "14.9.1.5@i-s"
	ospfVirtIfRtrDeadInterval = mib2Prefix + "14.9.1.6@i-s"
	ospfVirtIfState           = mib2Prefix + "14.9.1.7@i-s"
	ospfVirtIfEvents          = mib2Prefix + "14.9.1.8@i"
	ospfVirtIfAuthKey         = mib2Prefix + "14.9.1.9@s-s"
	ospfVirtIfStatus          = mib2Prefix + "14.9.1.10@i-s"
	ospfVirtIfAuthType        = mib2Prefix + "14.9.1.11@i-s"
	ospfVirtIfLsaCount        = mib2Prefix + "14.9.1.12@i"
	ospfVirtIfLsaCksumSum     = mib2Prefix + "14.9.1.13@i"

	// ospfNbrTable (14.10)
	// ospfNbrEntry (not-accessible)(14.10.1)
	ospfNbrIPAddr                  = mib2Prefix + "14.9.10.1@s"
	ospfNbrAddressLessIndex        = mib2Prefix + "14.9.10.2@i"
	ospfNbrRtrID                   = mib2Prefix + "14.9.10.3@s"
	ospfNbrOptions                 = mib2Prefix + "14.9.10.4@i"
	ospfNbrPriority                = mib2Prefix + "14.9.10.5@i-s"
	ospfNbrState                   = mib2Prefix + "14.9.10.6@i"
	ospfNbrEvents                  = mib2Prefix + "14.9.10.7@i"
	ospfNbrLsRetransQLen           = mib2Prefix + "14.9.10.8@i"
	ospfNbmaNbrStatus              = mib2Prefix + "14.9.10.9@i-s"
	ospfNbmaNbrPermanence          = mib2Prefix + "14.9.10.10@i"
	ospfNbrHelloSuppressed         = mib2Prefix + "14.9.10.11@i"
	ospfNbrRestartHelperStatus     = mib2Prefix + "14.9.10.12@i"
	ospfNbrRestartHelperAge        = mib2Prefix + "14.9.10.13@i"
	ospfNbrRestartHelperExitReason = mib2Prefix + "14.9.10.14@i"

	// ospfVirtNbrTable (14.11)
	// ospfVirtNbrEntry (not-accessible)(14.11.1)
	ospfVirtNbrArea                    = mib2Prefix + "14.11.1.1@s"
	ospfVirtNbrRtrID                   = mib2Prefix + "14.11.1.2@s"
	ospfVirtNbrIPAddr                  = mib2Prefix + "14.11.1.3@s"
	ospfVirtNbrOptions                 = mib2Prefix + "14.11.1.4@i"
	ospfVirtNbrState                   = mib2Prefix + "14.11.1.5@i"
	ospfVirtNbrEvents                  = mib2Prefix + "14.11.1.6@i"
	ospfVirtNbrLsRetransQLen           = mib2Prefix + "14.11.1.7@i"
	ospfVirtNbrHelloSuppressed         = mib2Prefix + "14.11.1.8@i"
	ospfVirtNbrRestartHelperStatus     = mib2Prefix + "14.11.1.9@i"
	ospfVirtNbrRestartHelperAge        = mib2Prefix + "14.11.1.10@i"
	ospfVirtNbrRestartHelperExitReason = mib2Prefix + "14.11.1.11@i"

	// ospfExtLsdbTable (14.12)
	// ospfExtLsdbEntry (not-accessible)(14.12.1)
	ospfExtLsdbType          = mib2Prefix + "14.12.1.1@i"
	ospfExtLsdbLsid          = mib2Prefix + "14.12.1.2@i"
	ospfExtLsdbRouterID      = mib2Prefix + "14.12.1.3@s"
	ospfExtLsdbSequence      = mib2Prefix + "14.12.1.4@i"
	ospfExtLsdbAge           = mib2Prefix + "14.12.1.5@i"
	ospfExtLsdbChecksum      = mib2Prefix + "14.12.1.6@i"
	ospfExtLsdbAdvertisement = mib2Prefix + "14.12.1.7@s"

	// ospfAreaAggregateTable (14.14)
	// ospfAreaAggregateEntry (not-accessible)(14.14.1)
	ospfAreaAggregateAreaID      = mib2Prefix + "14.14.1.1@s"
	ospfAreaAggregateLsdbType    = mib2Prefix + "14.14.1.2@i"
	ospfAreaAggregateNet         = mib2Prefix + "14.14.1.3@s"
	ospfAreaAggregateMask        = mib2Prefix + "14.14.1.4@s"
	ospfAreaAggregateStatus      = mib2Prefix + "14.14.1.5@i-s"
	ospfAreaAggregateEffect      = mib2Prefix + "14.14.1.6@i-s"
	ospfAreaAggregateExtRouteTag = mib2Prefix + "14.14.1.7@i-s"

	// ospfLocalLsdbTable (14.17)
	// ospfLocalLsdbEntry (not-accessible)(14.17.1)
	ospfLocalLsdbSequence      = mib2Prefix + "14.17.1.6@i"
	ospfLocalLsdbAge           = mib2Prefix + "14.17.1.7@i"
	ospfLocalLsdbChecksum      = mib2Prefix + "14.17.1.8@i"
	ospfLocalLsdbAdvertisement = mib2Prefix + "14.17.1.9@s"

	// ospfVirtLocalLsdbTable (14.18)
	// ospfVirtLocalLsdbEntry (not-accessible)(14.18.1)

	ospfVirtLocalLsdbSequence      = mib2Prefix + "14.18.1.6@i"
	ospfVirtLocalLsdbAge           = mib2Prefix + "14.18.1.7@i"
	ospfVirtLocalLsdbChecksum      = mib2Prefix + "14.18.1.8@i"
	ospfVirtLocalLsdbAdvertisement = mib2Prefix + "14.18.1.9@s"

	// ospfAsLsdbTable (14.19)
	// ospfAsLsdbEntry (not-accessible)(14.19.1)

	ospfAsLsdbSequence      = mib2Prefix + "14.19.1.4@i"
	ospfAsLsdbAge           = mib2Prefix + "14.19.1.5@i"
	ospfAsLsdbChecksum      = mib2Prefix + "14.19.1.6@i"
	ospfAsLsdbAdvertisement = mib2Prefix + "14.19.1.7@s"

	// ospfAreaLsaCountTable (14.20)
	// ospfAreaLsaCountEntry (not-accessible)(14.20.1)
	ospfAreaLsaCountNumber = mib2Prefix + "14.20.1.3@i"

	// ************** rfc4750 **************

	// ************** rfc4318 **************
	// RFC 4318 starts
	dot1dStpVersion     = mib2Prefix + "17.2.16@i-s"
	dot1dStpTxHoldCount = mib2Prefix + "17.2.17@i-s"
	// dot1dStpExtPortTable (17.2.19)
	// dot1dStpExtPortEntry (not-accessible)(17.2.19.1)
	dot1dStpPortProtocolMigration = mib2Prefix + "17.2.19.1.1@i-w-s"
	dot1dStpPortAdminEdgePort     = mib2Prefix + "17.2.19.1.2@i-w-s"
	dot1dStpPortOperEdgePort      = mib2Prefix + "17.2.19.1.3@i-w"
	dot1dStpPortAdminPointToPoint = mib2Prefix + "17.2.19.1.4@i-w-s"
	dot1dStpPortOperPointToPoint  = mib2Prefix + "17.2.19.1.5@i-w"
	dot1dStpPortAdminPathCost     = mib2Prefix + "17.2.19.1.6@i-w-s"

	// Following 2 don't have the data type in rfc4318 yet
	rstpNotifications = mib2Prefix + "134.0@i"
	rstpObjects       = mib2Prefix + "134.1@i"

	// Following 2 are the OBJECT-GROUP type, need details
	// rstpGroups (134.2.1)
	rstpBridgeGroup = mib2Prefix + "134.2.1.1@i"
	rstpPortGroup   = mib2Prefix + "134.2.1.1@i"

	// rstpCompliances (134.2.2)
	rstpCompliance = mib2Prefix + "134.2.2.1@i"
	// ************** rfc4318 **************

	// ************** rfc3877 alarm mib **************

	alarmModelLastChanged = mib2Prefix + "118.1.1.1@i"
	// alarmModelTable (118.1.1.2)
	// alarmModelEntry (118.1.1.2.1)
	// alarmModelIndex           = mib2Prefix + "118.1.1.2.1.1@i-w" (not-accessible)
	// alarmModelState           = mib2Prefix + "118.1.1.2.1.2@i-w" (not-accessible)
	alarmModelNotificationID  = mib2Prefix + "118.1.1.2.1.3@s-w-s"
	alarmModelVarbindIndex    = mib2Prefix + "118.1.1.2.1.4@i-w-s"
	alarmModelVarbindValue    = mib2Prefix + "118.1.1.2.1.5@i-w-s"
	alarmModelDescription     = mib2Prefix + "118.1.1.2.1.6@s-w-s"
	alarmModelSpecificPointer = mib2Prefix + "118.1.1.2.1.7@s-w-s"
	alarmModelVarbindSubtree  = mib2Prefix + "118.1.1.2.1.8@s-w-s"
	alarmModelResourcePrefix  = mib2Prefix + "118.1.1.2.1.9@s-w-s"
	alarmModelRowStatus       = mib2Prefix + "118.1.1.2.1.10@i-w-s"

	alarmActiveLastChanged = mib2Prefix + "118.1.2.1@i"
	// alarmActiveTable (118.1.2.2)
	// alarmActiveEntry (118.1.2.2.1)
	// alarmListName                = mib2Prefix + "118.1.2.2.1.1@s" (not-accessible)
	// alarmActiveDateAndTime       = mib2Prefix + "118.1.2.2.1.2@i" (not-accessible)
	// alarmActiveIndex             = mib2Prefix + "118.1.2.2.1.3@i" (not-accessible)
	alarmActiveEngineID          = mib2Prefix + "118.1.2.2.1.4@s-w"
	alarmActiveEngineAddressType = mib2Prefix + "118.1.2.2.1.5@i-w"
	alarmActiveEngineAddress     = mib2Prefix + "118.1.2.2.1.6@s-w"
	alarmActiveContextName       = mib2Prefix + "118.1.2.2.1.7@s-w"
	alarmActiveVariables         = mib2Prefix + "118.1.2.2.1.8@i-w"
	alarmActiveNotificationID    = mib2Prefix + "118.1.2.2.1.9@s-w"
	alarmActiveResourceID        = mib2Prefix + "118.1.2.2.1.10@s-w"
	alarmActiveDescription       = mib2Prefix + "118.1.2.2.1.11@s-w"
	alarmActiveLogPointer        = mib2Prefix + "118.1.2.2.1.12@s-w"
	alarmActiveModelPointer      = mib2Prefix + "118.1.2.2.1.13@s-w"
	alarmActiveSpecificPointer   = mib2Prefix + "118.1.2.2.1.14@s-w"

	// alarmActiveVariableTable (118.1.2.3)
	// alarmActiveVariableEntry (118.1.2.3.1)
	// alarmActiveVariableIndex          = mib2Prefix + "118.1.2.3.1.1@i" (not-accessible)
	alarmActiveVariableID             = mib2Prefix + "118.1.2.3.1.2@s-w"
	alarmActiveVariableValueType      = mib2Prefix + "118.1.2.3.1.3@i-w"
	alarmActiveVariableCounter32Val   = mib2Prefix + "118.1.2.3.1.4@i-w"
	alarmActiveVariableUnsigned32Val  = mib2Prefix + "118.1.2.3.1.5@i-w"
	alarmActiveVariableTimeTicksVal   = mib2Prefix + "118.1.2.3.1.6@i-w"
	alarmActiveVariableInteger32Val   = mib2Prefix + "118.1.2.3.1.7@i-w"
	alarmActiveVariableOctetStringVal = mib2Prefix + "118.1.2.3.1.8@s-w"
	alarmActiveVariableIPAddressVal   = mib2Prefix + "118.1.2.3.1.9@s-w"
	alarmActiveVariableOidVal         = mib2Prefix + "118.1.2.3.1.10@s-w"
	alarmActiveVariableCounter64Val   = mib2Prefix + "118.1.2.3.1.11@i-w"
	alarmActiveVariableOpaqueVal      = mib2Prefix + "118.1.2.3.1.12@i-w"

	// alarmActiveStatsTable (118.1.2.4)
	// alarmActiveStatsEntry (118.1.2.4.1)
	alarmActiveStatsActiveCurrent = mib2Prefix + "118.1.2.4.1.1@i-w"
	alarmActiveStatsActives       = mib2Prefix + "118.1.2.4.1.2@i-w"
	alarmActiveStatsLastRaise     = mib2Prefix + "118.1.2.4.1.3@i-w"
	alarmActiveStatsLastClear     = mib2Prefix + "118.1.2.4.1.4@i-w"

	alarmClearMaximum = mib2Prefix + "118.1.3.1@i-s"
	// alarmClearTable (118.1.3.2)
	// alarmClearEntry (118.1.3.2.1)
	// alarmClearIndex             = mib2Prefix + "118.1.3.2.1.1@i" (not-accessible)
	// alarmClearDateAndTime       = mib2Prefix + "118.1.3.2.1.2@i" (not-accessible)
	alarmClearEngineID          = mib2Prefix + "118.1.3.2.1.3@s-w"
	alarmClearEngineAddressType = mib2Prefix + "118.1.3.2.1.4@s-w"
	alarmClearEngineAddress     = mib2Prefix + "118.1.3.2.1.5@s-w"
	alarmClearContextName       = mib2Prefix + "118.1.3.2.1.6@s-w"
	alarmClearNotificationID    = mib2Prefix + "118.1.3.2.1.7@s-w"
	alarmClearResourceID        = mib2Prefix + "118.1.3.2.1.8@s-w"
	alarmClearLogIndex          = mib2Prefix + "118.1.3.2.1.9@i-w"
	alarmClearModelPointer      = mib2Prefix + "118.1.3.2.1.10@s-w"

	// ituAlarmTable (121.1.1.1)
	// ituAlarmEntry (121.1.1.1.1)
	// ituAlarmPerceivedSeverity = mib2Prefix + "121.1.1.1.1.1@i" (not-accessible)
	ituAlarmEventType      = mib2Prefix + "121.1.1.1.1.2@i-s-w"
	ituAlarmProbableCause  = mib2Prefix + "121.1.1.1.1.3@i-s-w"
	ituAlarmAdditionalText = mib2Prefix + "121.1.1.1.1.4@s-s-w"
	ituAlarmGenericModel   = mib2Prefix + "121.1.1.1.1.5@s-s-w"

	// ituAlarmActiveTable (121.1.2.1)
	// ituAlarmActiveEntry (121.1.2.1.1)
	ituAlarmActiveTrendIndication = mib2Prefix + "121.1.2.1.1.1@i-w"
	ituAlarmActiveDetector        = mib2Prefix + "121.1.2.1.1.2@s-w"
	ituAlarmActiveServiceProvider = mib2Prefix + "121.1.2.1.1.3@s-w"
	ituAlarmActiveServiceUser     = mib2Prefix + "121.1.2.1.1.4@s-w"

	// ituAlarmActiveStatsTable (121.1.2.2)
	// ituAlarmActiveStatsEntry (121.1.2.2.1)
	ituAlarmActiveStatsIndeterminateCurrent = mib2Prefix + "121.1.2.2.1.1@i-w"
	ituAlarmActiveStatsCriticalCurrent      = mib2Prefix + "121.1.2.2.1.2@i-w"
	ituAlarmActiveStatsMajorCurrent         = mib2Prefix + "121.1.2.2.1.3@i-w"
	ituAlarmActiveStatsMinorCurrent         = mib2Prefix + "121.1.2.2.1.4@i-w"
	ituAlarmActiveStatsWarningCurrent       = mib2Prefix + "121.1.2.2.1.5@i-w"
	ituAlarmActiveStatsIndeterminates       = mib2Prefix + "121.1.2.2.1.6@i-w"
	ituAlarmActiveStatsCriticals            = mib2Prefix + "121.1.2.2.1.7@i-w"
	ituAlarmActiveStatsMajors               = mib2Prefix + "121.1.2.2.1.8@i-w"
	ituAlarmActiveStatsMinors               = mib2Prefix + "121.1.2.2.1.9@i-w"
	ituAlarmActiveStatsWarnings             = mib2Prefix + "121.1.2.2.1.10@i-w"

	// ************** rfc3877 alarm mib **************

	// ************** rfc3812 **************
	mplsTunnelConfigured          = mib2Prefix + "10.166.3.1.1@i"
	mplsTunnelActive              = mib2Prefix + "10.166.3.1.2@i"
	mplsTunnelTEDistProto         = mib2Prefix + "10.166.3.1.3@i"
	mplsTunnelMaxHops             = mib2Prefix + "10.166.3.1.4@i"
	mplsTunnelNotificationMaxRate = mib2Prefix + "10.166.3.1.5@i-s"

	mplsTunnelIndexNext = mib2Prefix + "10.166.3.2.1@i"
	// mplsTunnelTable (10.166.3.2.2)
	// mplsTunnelEntry (not-accessible)(10.166.3.2.2.1)
	mplsTunnelName               = mib2Prefix + "10.166.3.2.2.1.5@s-s-m:4"
	mplsTunnelDescr              = mib2Prefix + "10.166.3.2.2.1.6@s-s-m:4"
	mplsTunnelIsIf               = mib2Prefix + "10.166.3.2.2.1.7@i-s"
	mplsTunnelIfIndex            = mib2Prefix + "10.166.3.2.2.1.8@i-s"
	mplsTunnelOwner              = mib2Prefix + "10.166.3.2.2.1.9@i"
	mplsTunnelRole               = mib2Prefix + "10.166.3.2.2.1.10@i-s"
	mplsTunnelXCPointer          = mib2Prefix + "10.166.3.2.2.1.11@s-s-m:4"
	mplsTunnelSignallingProto    = mib2Prefix + "10.166.3.2.2.1.12@i-s"
	mplsTunnelSetupPrio          = mib2Prefix + "10.166.3.2.2.1.13@i-s"
	mplsTunnelHoldingPrio        = mib2Prefix + "10.166.3.2.2.1.14@i-s"
	mplsTunnelSessionAttributes  = mib2Prefix + "10.166.3.2.2.1.15@i-s"
	mplsTunnelLocalProtectInUse  = mib2Prefix + "10.166.3.2.2.1.16@i-s"
	mplsTunnelResourcePointer    = mib2Prefix + "10.166.3.2.2.1.17@s-s-m:4"
	mplsTunnelPrimaryInstance    = mib2Prefix + "10.166.3.2.2.1.18@i"
	mplsTunnelInstancePriority   = mib2Prefix + "10.166.3.2.2.1.19@i-s"
	mplsTunnelHopTableIndex      = mib2Prefix + "10.166.3.2.2.1.20@i-s"
	mplsTunnelPathInUse          = mib2Prefix + "10.166.3.2.2.1.21@i-s"
	mplsTunnelARHopTableIndex    = mib2Prefix + "10.166.3.2.2.1.22@i"
	mplsTunnelCHopTableIndex     = mib2Prefix + "10.166.3.2.2.1.23@i"
	mplsTunnelIncludeAnyAffinity = mib2Prefix + "10.166.3.2.2.1.24@i-s"
	mplsTunnelIncludeAllAffinity = mib2Prefix + "10.166.3.2.2.1.25@i-s"
	mplsTunnelExcludeAnyAffinity = mib2Prefix + "10.166.3.2.2.1.26@i-s"
	mplsTunnelTotalUpTime        = mib2Prefix + "10.166.3.2.2.1.27@i"
	mplsTunnelInstanceUpTime     = mib2Prefix + "10.166.3.2.2.1.28@i"
	mplsTunnelPrimaryUpTime      = mib2Prefix + "10.166.3.2.2.1.29@i"
	mplsTunnelPathChanges        = mib2Prefix + "10.166.3.2.2.1.30@i"
	mplsTunnelLastPathChange     = mib2Prefix + "10.166.3.2.2.1.31@i"
	mplsTunnelCreationTime       = mib2Prefix + "10.166.3.2.2.1.32@i"
	mplsTunnelStateTransitions   = mib2Prefix + "10.166.3.2.2.1.33@i"
	mplsTunnelAdminStatus        = mib2Prefix + "10.166.3.2.2.1.34@i-s"
	mplsTunnelOperStatus         = mib2Prefix + "10.166.3.2.2.1.35@i"
	mplsTunnelRowStatus          = mib2Prefix + "10.166.3.2.2.1.36@i-s"
	mplsTunnelStorageType        = mib2Prefix + "10.166.3.2.2.1.37@i-s"

	mplsTunnelHopListIndexNext = mib2Prefix + "10.166.3.2.3@i"
	// mplsTunnelHopTable (10.166.3.2.4)
	// mplsTunnelHopEntry (not-accessible)(10.166.3.2.4.1)
	mplsTunnelHopAddrType       = mib2Prefix + "10.166.3.2.4.1.4@i-s"
	mplsTunnelHopIPAddr         = mib2Prefix + "10.166.3.2.4.1.5@s-s-m:4"
	mplsTunnelHopIPPrefixLen    = mib2Prefix + "10.166.3.2.4.1.6@i-s"
	mplsTunnelHopAsNumber       = mib2Prefix + "10.166.3.2.4.1.7@s-s-m:4"
	mplsTunnelHopAddrUnnum      = mib2Prefix + "10.166.3.2.4.1.8@s-s-m:4"
	mplsTunnelHopLspID          = mib2Prefix + "10.166.3.2.4.1.9@s-s-m:4"
	mplsTunnelHopType           = mib2Prefix + "10.166.3.2.4.1.10@i-s"
	mplsTunnelHopInclude        = mib2Prefix + "10.166.3.2.4.1.11@i-s"
	mplsTunnelHopPathOptionName = mib2Prefix + "10.166.3.2.4.1.12@s-s-m:4"
	mplsTunnelHopEntryPathComp  = mib2Prefix + "10.166.3.2.4.1.13@i-s-m:4"
	mplsTunnelHopRowStatus      = mib2Prefix + "10.166.3.2.4.1.14@i-s"
	mplsTunnelHopStorageType    = mib2Prefix + "10.166.3.2.4.1.15@i-s"

	mplsTunnelResourceIndexNext = mib2Prefix + "10.166.3.2.5@i"

	// mplsTunnelResourceTable (not-accessible)(10.166.3.2.6.0)
	// mplsTunnelResourceEntry (not-accessible)(10.166.3.2.6.1)
	mplsTunnelResourceMaxRate       = mib2Prefix + "10.166.3.2.6.1.2@i-s"
	mplsTunnelResourceMeanRate      = mib2Prefix + "10.166.3.2.6.1.3@i-s"
	mplsTunnelResourceMaxBurstSize  = mib2Prefix + "10.166.3.2.6.1.4@i-s"
	mplsTunnelResourceMeanBurstSize = mib2Prefix + "10.166.3.2.6.1.5@i-s"
	mplsTunnelResourceExBurstSize   = mib2Prefix + "10.166.3.2.6.1.6@i-s"
	mplsTunnelResourceFrequency     = mib2Prefix + "10.166.3.2.6.1.7@i-s"
	mplsTunnelResourceWeight        = mib2Prefix + "10.166.3.2.6.1.8@i-s"
	mplsTunnelResourceRowStatus     = mib2Prefix + "10.166.3.2.6.1.9@i-s"
	mplsTunnelResourceStorageType   = mib2Prefix + "10.166.3.2.6.1.10@i-s"

	// mplsTunnelARHopTable (not-accessible)(10.166.3.2.7)
	// mplsTunnelARHopEntry (not-accessible)(10.166.3.2.7.1)
	mplsTunnelARHopAddrType  = mib2Prefix + "10.166.3.2.7.1.3@i"
	mplsTunnelARHopIPAddr    = mib2Prefix + "10.166.3.2.7.1.4@s-m:4"
	mplsTunnelARHopAddrUnnum = mib2Prefix + "10.166.3.2.7.1.5@s-m:4"
	mplsTunnelARHopLspID     = mib2Prefix + "10.166.3.2.7.1.6@s-m:4"

	// mplsTunnelCHopTable (not-accessible)(10.166.3.2.8)
	// mplsTunnelCHopEntry (not-accessible)(10.166.3.2.8.1)
	mplsTunnelCHopAddrType    = mib2Prefix + "10.166.3.2.8.1.3@i"
	mplsTunnelCHopIPAddr      = mib2Prefix + "10.166.3.2.8.1.4@s-m:4"
	mplsTunnelCHopIPPrefixLen = mib2Prefix + "10.166.3.2.8.1.5@i"
	mplsTunnelCHopAsNumber    = mib2Prefix + "10.166.3.2.8.1.6@s-m:4"
	mplsTunnelCHopAddrUnnum   = mib2Prefix + "10.166.3.2.8.1.7@s-m:4"
	mplsTunnelCHopLspID       = mib2Prefix + "10.166.3.2.8.1.8@s-m:4"
	mplsTunnelCHopType        = mib2Prefix + "10.166.3.2.8.1.9@i"

	// mplsTunnelPerfTable (not-accessible)(10.166.3.2.9.)
	// mplsTunnelPerfEntry (not-accessible)(10.166.3.2.9.1)
	mplsTunnelPerfPackets   = mib2Prefix + "10.166.3.2.9.1.1@i"
	mplsTunnelPerfHCPackets = mib2Prefix + "10.166.3.2.9.1.2@i"
	mplsTunnelPerfErrors    = mib2Prefix + "10.166.3.2.9.1.3@i"
	mplsTunnelPerfBytes     = mib2Prefix + "10.166.3.2.9.1.4@i"
	mplsTunnelPerfHCBytes   = mib2Prefix + "10.166.3.2.9.1.5@i"

	// mplsTunnelCRLDPResTable (not-accessible)(10.166.3.2.10)
	// mplsTunnelCRLDPResEntry (not-accessible)(10.166.3.2.10.1)
	mplsTunnelCRLDPResMeanBurstSize = mib2Prefix + "10.166.3.2.10.1.1@i-s"
	mplsTunnelCRLDPResExBurstSize   = mib2Prefix + "10.166.3.2.10.1.2@i-s"
	mplsTunnelCRLDPResFrequency     = mib2Prefix + "10.166.3.2.10.1.3@i-s"
	mplsTunnelCRLDPResWeight        = mib2Prefix + "10.166.3.2.10.1.4@i-s"
	mplsTunnelCRLDPResFlags         = mib2Prefix + "10.166.3.2.10.1.5@i-s"
	mplsTunnelCRLDPResRowStatus     = mib2Prefix + "10.166.3.2.10.1.6@i-s"
	mplsTunnelCRLDPResStorageType   = mib2Prefix + "10.166.3.2.10.1.7@i-s"

	mplsTunnelNotificationEnable = mib2Prefix + "10.166.3.2.11@i-s"
	// ************** rfc3812 **************

	// ************** rfc3814 **************
	mplsFTNIndexNext        = mib2Prefix + "10.166.8.1.1@i"
	mplsFTNTableLastChanged = mib2Prefix + "10.166.8.1.2@i"

	// mplsFTNTable  (not-accessible)(10.166.8.1.3)
	// mplsFTNEntry (not-accessible)(10.166.8.1.3.1)
	mplsFTNRowStatus     = mib2Prefix + "10.166.8.1.3.1.2@i-s"
	mplsFTNDescr         = mib2Prefix + "10.166.8.1.3.1.3@s-s-m:4"
	mplsFTNMask          = mib2Prefix + "10.166.8.1.3.1.4@i-s"
	mplsFTNAddrType      = mib2Prefix + "10.166.8.1.3.1.5@i-s"
	mplsFTNSourceAddrMin = mib2Prefix + "10.166.8.1.3.1.6@s-s-m:4"
	mplsFTNSourceAddrMax = mib2Prefix + "10.166.8.1.3.1.7@s-s-m:4"
	mplsFTNDestAddrMin   = mib2Prefix + "10.166.8.1.3.1.8@s-s-m:4"
	mplsFTNDestAddrMax   = mib2Prefix + "10.166.8.1.3.1.9@s-s-m:4"
	mplsFTNSourcePortMin = mib2Prefix + "10.166.8.1.3.1.10@i-s"
	mplsFTNSourcePortMax = mib2Prefix + "10.166.8.1.3.1.11@i-s"
	mplsFTNDestPortMin   = mib2Prefix + "10.166.8.1.3.1.12@i-s"
	mplsFTNDestPortMax   = mib2Prefix + "10.166.8.1.3.1.13@i-s"
	mplsFTNProtocol      = mib2Prefix + "10.166.8.1.3.1.14@i-s"
	mplsFTNDscp          = mib2Prefix + "10.166.8.1.3.1.15@i-s"
	mplsFTNActionType    = mib2Prefix + "10.166.8.1.3.1.16@i-s"
	mplsFTNActionPointer = mib2Prefix + "10.166.8.1.3.1.17@s-s-m:4"
	mplsFTNStorageType   = mib2Prefix + "10.166.8.1.3.1.18@i-s"

	mplsFTNMapTableLastChanged = mib2Prefix + "10.166.8.1.4@i"

	// mplsFTNMapTable  (not-accessible)(10.166.8.1.5)
	// mplsFTNMapEntry (not-accessible)(10.166.8.1.5.1)
	mplsFTNMapRowStatus   = mib2Prefix + "10.166.8.1.5.1.4@i-s"
	mplsFTNMapStorageType = mib2Prefix + "10.166.8.1.5.1.5@i-s"

	// mplsFTNPerfTable (not-accessible)(10.166.8.1.6)
	// MplsFTNPerfEntry (not-accessible)(10.166.8.1.6.1)
	mplsFTNPerfMatchedPackets    = mib2Prefix + "10.166.8.1.6.1.3@i"
	mplsFTNPerfMatchedOctets     = mib2Prefix + "10.166.8.1.6.1.4@i"
	mplsFTNPerfDiscontinuityTime = mib2Prefix + "10.166.8.1.6.1.5@i"

	// ************** rfc3814 **************

	// ************** Private MIB *********************
	// SYSTEM (1)

	// IPGS-6416XSFP-MIB systemInfo                                    node         1.3.6.1.4.1.37072.302.3.1.1
	// IPGS-6416XSFP-MIB configuration                                 node         1.3.6.1.4.1.37072.302.3.1.1.1
	systemName     = privateMibOid + "1.1.1@s"
	systemDescr    = privateMibOid + "1.1.2@s"
	systemLocation = privateMibOid + "1.1.3@s"
	systemContact  = privateMibOid + "1.1.4@s"
	// IPGS-6416XSFP-MIB information                                   node         1.3.6.1.4.1.37072.302.3.1.1.2
	deviceTimeInfo  = privateMibOid + "1.2.1@s"
	uptime          = privateMibOid + "1.2.2@i"
	softwareVersion = privateMibOid + "1.2.3@s"
	macAddressInfo  = privateMibOid + "1.2.4@s"
	hardwareModel   = privateMibOid + "1.2.5@s"
	hardwareDescr   = privateMibOid + "1.2.6@s"
	serialNum       = privateMibOid + "1.2.7@s"
	// IPGS-6416XSFP-MIB ipAddressCfg                                  node         1.3.6.1.4.1.37072.302.3.1.1.3
	dhcpClientEnable = privateMibOid + "1.3.1@s"
	devIPAddress     = privateMibOid + "1.3.2@s"
	networkMask      = privateMibOid + "1.3.3@s"
	gateway          = privateMibOid + "1.3.4@s"
	dns              = privateMibOid + "1.3.5@s"
	currentIPAddress = privateMibOid + "1.3.6@s"
	currentNetmask   = privateMibOid + "1.3.7@s"
	currentGateway   = privateMibOid + "1.3.8@s"
	currentDNS       = privateMibOid + "1.3.9@s"
	// IPGS-6416XSFP-MIB time                                          node         1.3.6.1.4.1.37072.302.3.1.1.4
	timeMethod       = privateMibOid + "1.4.1@i"
	deviceManualTime = privateMibOid + "1.4.2@s"
	deviceTimeUTC    = privateMibOid + "1.4.3@i"
	timezone         = privateMibOid + "1.4.4@i"
	ntpServer        = privateMibOid + "1.4.5@s"
	// IPGS-6416XSFP-MIB monitor                                       node         1.3.6.1.4.1.37072.302.3.1.1.5
	deviceVoltage     = privateMibOid + "1.5.1@i"
	deviceCurrent     = privateMibOid + "1.5.2@i"
	devicePower       = privateMibOid + "1.5.3@i"
	deviceTemperature = privateMibOid + "1.5.4@i"
	deviceCPUUsage    = privateMibOid + "1.5.5@i"
	deviceRAMUsage    = privateMibOid + "1.5.6@i"
	// IPGS-6416XSFP-MIB ports                                         node         1.3.6.1.4.1.37072.302.3.1.2
	// IPGS-6416XSFP-MIB portBasic                                     node         1.3.6.1.4.1.37072.302.3.1.2.1
	// IPGS-6416XSFP-MIB portCfgTable                                  table        1.3.6.1.4.1.37072.302.3.1.2.1.1
	// IPGS-6416XSFP-MIB portCfgEntry                                  row          1.3.6.1.4.1.37072.302.3.1.2.1.1.1
	portIndex              = privateMibOid + "2.1.1.1.1@i"
	portType               = privateMibOid + "2.1.1.1.2@s"
	portDescr              = privateMibOid + "2.1.1.1.3@s"
	portEnabled            = privateMibOid + "2.1.1.1.4@i"
	portFlowControlEnabled = privateMibOid + "2.1.1.1.5@i"
	portSpeed              = privateMibOid + "2.1.1.1.6@i"
	// IPGS-6416XSFP-MIB portStatusTable                               table        1.3.6.1.4.1.37072.302.3.1.2.1.2
	// IPGS-6416XSFP-MIB portStatusTableEntry                          row          1.3.6.1.4.1.37072.302.3.1.2.1.2.1
	portStatusIndex               = privateMibOid + "2.1.2.1.1@i"
	portLinkStatus                = privateMibOid + "2.1.2.1.2@i"
	portCurrentEnabled            = privateMibOid + "2.1.2.1.3@i"
	portCurrentFlowControlEnabled = privateMibOid + "2.1.2.1.4@i"
	portCurrentSpeed              = privateMibOid + "2.1.2.1.5@i"
	// IPGS-6416XSFP-MIB portRMONTable                                 node         1.3.6.1.4.1.37072.302.3.1.2.1.3
	// IPGS-6416XSFP-MIB portTrafficTable                              table        1.3.6.1.4.1.37072.302.3.1.2.1.3.1
	// IPGS-6416XSFP-MIB portTrafficTableEntry                         row          1.3.6.1.4.1.37072.302.3.1.2.1.3.1.1
	portTrafficIndex = privateMibOid + "2.1.3.1.1.1@i"
	portTransmitting = privateMibOid + "2.1.3.1.1.2@i"
	portReceiving    = privateMibOid + "2.1.3.1.1.3@i"
	// IPGS-6416XSFP-MIB portErrorStatisticsTable                      table        1.3.6.1.4.1.37072.302.3.1.2.1.3.2
	// IPGS-6416XSFP-MIB portErrorStatisticsTableEntry                 row          1.3.6.1.4.1.37072.302.3.1.2.1.3.2.1
	portErrorStatisticsIndex = privateMibOid + "2.1.3.2.1.1@i"
	txMACErrorPacket         = privateMibOid + "2.1.3.2.1.2@i"
	rxMACErrorPacket         = privateMibOid + "2.1.3.2.1.3@i"
	collision                = privateMibOid + "2.1.3.2.1.4@i"
	lateCollision            = privateMibOid + "2.1.3.2.1.5@i"
	excessiveCollision       = privateMibOid + "2.1.3.2.1.6@i"
	badCRCPacket             = privateMibOid + "2.1.3.2.1.7@i"
	jabberPacket             = privateMibOid + "2.1.3.2.1.8@i"
	oversizePacket           = privateMibOid + "2.1.3.2.1.9@i"
	undersizePacket          = privateMibOid + "2.1.3.2.1.10@i"
	fragments                = privateMibOid + "2.1.3.2.1.11@i"
	packetSentMultipleTimes  = privateMibOid + "2.1.3.2.1.12@i"
	deferredSentFrames       = privateMibOid + "2.1.3.2.1.13@i"
	unknownControlFrames     = privateMibOid + "2.1.3.2.1.14@i"
	insufficientDropPacket   = privateMibOid + "2.1.3.2.1.15@i"
	// IPGS-6416XSFP-MIB portFlowControlStatisticsTable                table        1.3.6.1.4.1.37072.302.3.1.2.1.3.3
	// IPGS-6416XSFP-MIB portFlowControlStatisticsTableEntry           row          1.3.6.1.4.1.37072.302.3.1.2.1.3.3.1
	portFlowControlStatisticsIndex = privateMibOid + "2.1.3.3.1.1@i"
	txFlowControlFrames            = privateMibOid + "2.1.3.3.1.2@i"
	rxGoodFlowControlFrames        = privateMibOid + "2.1.3.3.1.3@i"
	rxBadFlowControlFrames         = privateMibOid + "2.1.3.3.1.4@i"
	// IPGS-6416XSFP-MIB portCastTypeStatisticsTable                   table        1.3.6.1.4.1.37072.302.3.1.2.1.3.4
	// IPGS-6416XSFP-MIB portCastTypeStatisticsTableEntry              row          1.3.6.1.4.1.37072.302.3.1.2.1.3.4.1
	portCastTypeStatisticsIndex = privateMibOid + "2.1.3.4.1.1@i"
	txGoodUnicastFrames         = privateMibOid + "2.1.3.4.1.2@i"
	rxGoodUnicastFrames         = privateMibOid + "2.1.3.4.1.3@i"
	txBroadCastFrames           = privateMibOid + "2.1.3.4.1.4@i"
	rxBroadCastFrames           = privateMibOid + "2.1.3.4.1.5@i"
	txMulticastFrames           = privateMibOid + "2.1.3.4.1.6@i"
	rxMulticastFrames           = privateMibOid + "2.1.3.4.1.7@i"
	// IPGS-6416XSFP-MIB portPacketSizeStatisticsTable                 table        1.3.6.1.4.1.37072.302.3.1.2.1.3.5
	// IPGS-6416XSFP-MIB portPacketSizeStatisticsTableEntry            row          1.3.6.1.4.1.37072.302.3.1.2.1.3.5.1
	portPacketSizeStatisticsIndex = privateMibOid + "2.1.3.5.1.1@i"
	bytes64Packet                 = privateMibOid + "2.1.3.5.1.2@i"
	bytes65to127Packet            = privateMibOid + "2.1.3.5.1.3@i"
	bytes128to255Packet           = privateMibOid + "2.1.3.5.1.4@i"
	bytes256to511Packet           = privateMibOid + "2.1.3.5.1.5@i"
	bytes512to1023Packet          = privateMibOid + "2.1.3.5.1.6@i"
	bytes1024toMAXPacket          = privateMibOid + "2.1.3.5.1.7@i"
	// IPGS-6416XSFP-MIB portBasicStatisticsTable                      table        1.3.6.1.4.1.37072.302.3.1.2.1.3.6
	// IPGS-6416XSFP-MIB portBasicStatisticsTableEntry                 row          1.3.6.1.4.1.37072.302.3.1.2.1.3.6.1
	portBasicStatisticsIndex = privateMibOid + "2.1.3.6.1.1@i"
	txGoodFrames             = privateMibOid + "2.1.3.6.1.2@i"
	rxGoodFrames             = privateMibOid + "2.1.3.6.1.3@i"
	rxBadFrames              = privateMibOid + "2.1.3.6.1.4@i"
	txGoodPackets            = privateMibOid + "2.1.3.6.1.5@i"
	rxGoodPackets            = privateMibOid + "2.1.3.6.1.6@i"
	rxBadPackets             = privateMibOid + "2.1.3.6.1.7@i"
	clearAllPortStatistics   = privateMibOid + "2.1.3.7@i"
	// IPGS-6416XSFP-MIB portMirroring                                 node         1.3.6.1.4.1.37072.302.3.1.2.2
	ingressCfg = privateMibOid + "2.2.1@s"
	egressCfg  = privateMibOid + "2.2.2@s"
	// IPGS-6416XSFP-MIB rateLimiting                                  node         1.3.6.1.4.1.37072.302.3.1.2.3
	// IPGS-6416XSFP-MIB rateLimitTable                                table        1.3.6.1.4.1.37072.302.3.1.2.3.1
	// IPGS-6416XSFP-MIB rateLimitTableEntry                           row          1.3.6.1.4.1.37072.302.3.1.2.3.1.1
	portRateLimitIndex    = privateMibOid + "2.3.1.1.1@i"
	ingressLimitFrameType = privateMibOid + "2.3.1.1.2@i"
	ingressBandwidthLimit = privateMibOid + "2.3.1.1.3@i"
	egressBandwidthLimit  = privateMibOid + "2.3.1.1.4@i"
	// IPGS-6416XSFP-MIB aggregation                                   node         1.3.6.1.4.1.37072.302.3.1.2.4
	// IPGS-6416XSFP-MIB aggregationCfgTable                           table        1.3.6.1.4.1.37072.302.3.1.2.4.1
	// IPGS-6416XSFP-MIB aggregationCfgTableEntry                      row          1.3.6.1.4.1.37072.302.3.1.2.4.1.1
	aggregationCfgGroupIndex = privateMibOid + "2.4.1.1.1@i"
	trunkingType             = privateMibOid + "2.4.1.1.2@i"
	trunkingMembers          = privateMibOid + "2.4.1.1.3@s"
	// IPGS-6416XSFP-MIB aggregationStatusTable                        table        1.3.6.1.4.1.37072.302.3.1.2.4.2
	// IPGS-6416XSFP-MIB aggregationStatusTableEntry                   row          1.3.6.1.4.1.37072.302.3.1.2.4.2.1
	aggregationStatusGroupIndex = privateMibOid + "2.4.2.1.1@i"
	currentTrunkingType         = privateMibOid + "2.4.2.1.2@s"
	currentTrunkingMember       = privateMibOid + "2.4.2.1.3@s"
	// IPGS-6416XSFP-MIB poe                                           node         1.3.6.1.4.1.37072.302.3.1.3
	// IPGS-6416XSFP-MIB poeSys                                        node         1.3.6.1.4.1.37072.302.3.1.3.1
	poeMaxPower          = privateMibOid + "3.1.1@i"
	poeLegacyModeEnabled = privateMibOid + "3.1.2@i"
	// IPGS-6416XSFP-MIB poePorts                                      node         1.3.6.1.4.1.37072.302.3.1.3.2
	// IPGS-6416XSFP-MIB poeCfgTable                                   table        1.3.6.1.4.1.37072.302.3.1.3.2.1
	// IPGS-6416XSFP-MIB poeCfgTableEntry                              row          1.3.6.1.4.1.37072.302.3.1.3.2.1.1
	// poePortIndex                = privateMibOid + "3.2.1.1.1"
	poeEnabled                  = privateMibOid + "3.2.1.1.2@i"
	poePowerLimit               = privateMibOid + "3.2.1.1.3@i"
	poeSchedulingEnabled        = privateMibOid + "3.2.1.1.4@i"
	poeAliveDetectEnabled       = privateMibOid + "3.2.1.1.5@i"
	poeAliveDetectPingIP        = privateMibOid + "3.2.1.1.6@s"
	poeAliveDetectInterval      = privateMibOid + "3.2.1.1.7@i"
	poeAliveDetectRetryCount    = privateMibOid + "3.2.1.1.8@i"
	poeAliveDetectFailureAction = privateMibOid + "3.2.1.1.9@i"
	// IPGS-6416XSFP-MIB poeSchedule                                   node         1.3.6.1.4.1.37072.302.3.1.3.3
	poeSchdulingSun = privateMibOid + "3.3.1@s"
	poeSchdulingMon = privateMibOid + "3.3.2@s"
	poeSchdulingTue = privateMibOid + "3.3.3@s"
	poeSchdulingWed = privateMibOid + "3.3.4@s"
	poeSchdulingThu = privateMibOid + "3.3.5@s"
	poeSchdulingFri = privateMibOid + "3.3.6@s"
	poeSchdulingSat = privateMibOid + "3.3.7@s"
	// IPGS-6416XSFP-MIB poeStatus                                     node         1.3.6.1.4.1.37072.302.3.1.3.4
	poeFirmwareVersion  = privateMibOid + "3.4.1@s"
	poePowerConsumption = privateMibOid + "3.4.2@i"
	poePowerMainVoltage = privateMibOid + "3.4.3@i"
	poePowerMainCurrent = privateMibOid + "3.4.4@i"
	// IPGS-6416XSFP-MIB poePortstatusTable                            table        1.3.6.1.4.1.37072.302.3.1.3.4.5
	// IPGS-6416XSFP-MIB poePortStatusTableEntry                       row          1.3.6.1.4.1.37072.302.3.1.3.4.5.1
	poePortEnabled         = privateMibOid + "3.4.5.1.2@i"
	poePortLinked          = privateMibOid + "3.4.5.1.3@i"
	poePortState           = privateMibOid + "3.4.5.1.4@s"
	poePortTemperature     = privateMibOid + "3.4.5.1.5@i"
	poePortCurrent         = privateMibOid + "3.4.5.1.6@i"
	poePortPower           = privateMibOid + "3.4.5.1.7@i"
	poePortDeterminedClass = privateMibOid + "3.4.5.1.8@s"
	// IPGS-6416XSFP-MIB poeAliveDetectStatusTable                     table        1.3.6.1.4.1.37072.302.3.1.3.4.6
	// IPGS-6416XSFP-MIB poeAliveDetectStatusTableEntry                row          1.3.6.1.4.1.37072.302.3.1.3.4.6.1
	poeAliveDetectStatusEnabled    = privateMibOid + "3.4.6.1.2@i"
	poeAliveDetectStatusPingIP     = privateMibOid + "3.4.6.1.3@s"
	poeAliveDetectStatusPingResult = privateMibOid + "3.4.6.1.4@s"
	// IPGS-6416XSFP-MIB vlan                                          node         1.3.6.1.4.1.37072.302.3.1.4
	vlanMode = privateMibOid + "4.1@i"
	// IPGS-6416XSFP-MIB vlanDot1q                                     node         1.3.6.1.4.1.37072.302.3.1.4.2
	// IPGS-6416XSFP-MIB vlanDot1qGroupsTableGroup                     node         1.3.6.1.4.1.37072.302.3.1.4.2.1
	vlanDot1qGroupsAddCfg    = privateMibOid + "4.2.1.1@s"
	vlanDot1qGroupsDeleteCfg = privateMibOid + "4.2.1.2@s"
	vlanDot1qGroupsCfgRow    = privateMibOid + "4.2.1.3@s"
	// IPGS-6416XSFP-MIB vlanDot1qGroupsTable                          table        1.3.6.1.4.1.37072.302.3.1.4.2.1.4
	// IPGS-6416XSFP-MIB vlanDot1qGroupsTableEntry                     row          1.3.6.1.4.1.37072.302.3.1.4.2.1.4.1
	vlanDot1qGroupsName           = privateMibOid + "4.2.1.4.1.2@s"
	vlanDot1qGroupsPortMembers    = privateMibOid + "4.2.1.4.1.3@s"
	vlanDot1qGroupsPortMemberTags = privateMibOid + "4.2.1.4.1.4@s"
	// IPGS-6416XSFP-MIB vlanDot1qPVIDTable                            table        1.3.6.1.4.1.37072.302.3.1.4.2.2
	// IPGS-6416XSFP-MIB vlanDot1qPVIDTableEntry                       row          1.3.6.1.4.1.37072.302.3.1.4.2.2.1
	vlanDot1qPVID                 = privateMibOid + "4.2.2.1.2@i"
	vlanDot1qAcceptableFrameType  = privateMibOid + "4.2.2.1.3@i"
	vlanDot1qManagementPermission = privateMibOid + "4.2.3@i"
	vlanDot1qGVRPEnabled          = privateMibOid + "4.2.4@i"
	// IPGS-6416XSFP-MIB gvrpStatusTable                               table        1.3.6.1.4.1.37072.302.3.1.4.2.5
	// IPGS-6416XSFP-MIB gvrpStatusTableEntry                          row          1.3.6.1.4.1.37072.302.3.1.4.2.5.1
	gvrpCurrentVlanID             = privateMibOid + "4.2.5.1.1@i"
	gvrpCurrentMembers            = privateMibOid + "4.2.5.1.2@s"
	gvrpCurrentMembersCount       = privateMibOid + "4.2.5.1.3@i"
	gvrpCurrentTaggedMembers      = privateMibOid + "4.2.5.1.4@s"
	gvrpCurrentTaggedMembersCount = privateMibOid + "4.2.5.1.5@i"
	gvrpCurrentStatus             = privateMibOid + "4.2.5.1.6@s"
	// IPGS-6416XSFP-MIB vlanPortBased                                 node         1.3.6.1.4.1.37072.302.3.1.4.3
	// IPGS-6416XSFP-MIB vlanPortBasedGroupsTableGroup                 node         1.3.6.1.4.1.37072.302.3.1.4.3.1
	vlanPortBasedGroupsAddCfg    = privateMibOid + "4.3.1.1@s"
	vlanPortBasedGroupsDeleteCfg = privateMibOid + "4.3.1.2@s"
	vlanPortBasedGroupsCfgRow    = privateMibOid + "4.3.1.3@s"
	// IPGS-6416XSFP-MIB vlanPortBasedGroupsTable                      table        1.3.6.1.4.1.37072.302.3.1.4.3.1.4
	// IPGS-6416XSFP-MIB vlanPortBasedGroupsTableEntry                 row          1.3.6.1.4.1.37072.302.3.1.4.3.1.4.1
	vlanPortBasedGroupsID          = privateMibOid + "4.3.1.4.1.1@i"
	vlanPortBasedGroupsPortMembers = privateMibOid + "4.3.1.4.1.2@s"
	// IPGS-6416XSFP-MIB qos                                           node         1.3.6.1.4.1.37072.302.3.1.5
	// IPGS-6416XSFP-MIB qosPolicy                                     node         1.3.6.1.4.1.37072.302.3.1.5.1
	qosUseQueuingSchema = privateMibOid + "5.1.1@i"
	qosPriorityType     = privateMibOid + "5.1.2@i"
	// IPGS-6416XSFP-MIB qosWeightedFairTrafficRatioTable              table        1.3.6.1.4.1.37072.302.3.1.5.1.3
	// IPGS-6416XSFP-MIB qosWeightedFairTrafficRatioTableEntry         row          1.3.6.1.4.1.37072.302.3.1.5.1.3.1
	qosTrafficIndex = privateMibOid + "5.1.3.1.1@i"
	qosTrafficRatio = privateMibOid + "5.1.3.1.2@i"
	// IPGS-6416XSFP-MIB qosCosPriorityNumberTable                     table        1.3.6.1.4.1.37072.302.3.1.5.1.4
	// IPGS-6416XSFP-MIB qosCosPriorityNumberTableEntry                row          1.3.6.1.4.1.37072.302.3.1.5.1.4.1
	qosCosPriorityIndex = privateMibOid + "5.1.4.1.1@i"
	qosCosPriority      = privateMibOid + "5.1.4.1.2@i"
	// IPGS-6416XSFP-MIB qosDscpPriorityNumberTable                    table        1.3.6.1.4.1.37072.302.3.1.5.1.5
	// IPGS-6416XSFP-MIB qosDscpPriorityNumberTableEntry               row          1.3.6.1.4.1.37072.302.3.1.5.1.5.1
	qosDscpPriorityIndex = privateMibOid + "5.1.5.1.1@i"
	qosDscpPriority      = privateMibOid + "5.1.5.1.2@i"
	// IPGS-6416XSFP-MIB multicast                                     node         1.3.6.1.4.1.37072.302.3.1.6
	// IPGS-6416XSFP-MIB gmrp                                          node         1.3.6.1.4.1.37072.302.3.1.6.1
	gmrpEnabled = privateMibOid + "6.1.1@i"
	// IPGS-6416XSFP-MIB gmrpStaticCfgTableGroup                       node         1.3.6.1.4.1.37072.302.3.1.6.1.2
	gmrpAddStaticCfg    = privateMibOid + "6.1.2.1@s"
	gmrpDeleteStaticCfg = privateMibOid + "6.1.2.2@s"
	gmrpStaticCfgRow    = privateMibOid + "6.1.2.3@s"
	// IPGS-6416XSFP-MIB gmrpStaticCfgTable                            table        1.3.6.1.4.1.37072.302.3.1.6.1.2.4
	// IPGS-6416XSFP-MIB gmrpStaticCfgTableEntry                       row          1.3.6.1.4.1.37072.302.3.1.6.1.2.4.1
	gmrpPortNumbers = privateMibOid + "6.1.2.4.1.3@s"
	// IPGS-6416XSFP-MIB gmrpStatusTable                               table        1.3.6.1.4.1.37072.302.3.1.6.1.3
	// IPGS-6416XSFP-MIB gmrpStatusTableEntry                          row          1.3.6.1.4.1.37072.302.3.1.6.1.3.1
	gmrpCurrentStaticPortNumbers  = privateMibOid + "6.1.3.1.2@s"
	gmrpCurrentDynamicPortNumbers = privateMibOid + "6.1.3.1.3@s"
	// IPGS-6416XSFP-MIB igmpSnooping                                  node         1.3.6.1.4.1.37072.302.3.1.6.2
	igmpQuerierEnabled                        = privateMibOid + "6.2.1@s-s"
	igmpQuerierVersion                        = privateMibOid + "6.2.2@s-s"
	igmpSnoopingEnabled                       = privateMibOid + "6.2.3@s-s"
	igmpUnregisterFloodingEnabled             = privateMibOid + "6.2.4@s-s"
	igmpFloodWellKnownMulticastTrafficEnabled = privateMibOid + "6.2.5@s-s"
	// IPGS-6416XSFP-MIB igmpPortsCfgTable                             table        1.3.6.1.4.1.37072.302.3.1.6.2.6
	// IPGS-6416XSFP-MIB igmpPortCfgTableEntry                         row          1.3.6.1.4.1.37072.302.3.1.6.2.6.1
	igmpPortIndex         = privateMibOid + "6.2.6.1.1@i"
	igmpRouterPortEnabled = privateMibOid + "6.2.6.1.2@i"
	igmpFastLeaveEnabled  = privateMibOid + "6.2.6.1.3@i"
	// IPGS-6416XSFP-MIB igmpStaticEntryGroup                          node         1.3.6.1.4.1.37072.302.3.1.6.2.7
	igmpStaticEntryAdd    = privateMibOid + "6.2.7.1@s"
	igmpStaticEntryDelete = privateMibOid + "6.2.7.2@s"
	igmpStaticEntryRow    = privateMibOid + "6.2.7.3@s"
	// IPGS-6416XSFP-MIB igmpStaticEntryTable                          table        1.3.6.1.4.1.37072.302.3.1.6.2.7.4
	// IPGS-6416XSFP-MIB igmpStaticEntryTableEntry                     row          1.3.6.1.4.1.37072.302.3.1.6.2.7.4.1
	igmpStaticVlanID      = privateMibOid + "6.2.7.4.1.1@i"
	igmpStaticAddress     = privateMibOid + "6.2.7.4.1.2@s"
	igmpStaticPortMembers = privateMibOid + "6.2.7.4.1.3@s"
	// IPGS-6416XSFP-MIB igmpRouterPortStatusTable                     table        1.3.6.1.4.1.37072.302.3.1.6.2.8
	// IPGS-6416XSFP-MIB igmpRouterPortStatusTableEntry                row          1.3.6.1.4.1.37072.302.3.1.6.2.8.1
	igmpRouterPortStatusPortNumber = privateMibOid + "6.2.8.1.1@i"
	igmpRouterPortRole             = privateMibOid + "6.2.8.1.2@s"
	// IPGS-6416XSFP-MIB igmpGroupStatusTable                          table        1.3.6.1.4.1.37072.302.3.1.6.2.9
	// IPGS-6416XSFP-MIB igmpGroupStatusTableEntry                     row          1.3.6.1.4.1.37072.302.3.1.6.2.9.1
	igmpGroupMulticastAddress   = privateMibOid + "6.2.9.1.1@s"
	igmpGroupVlanID             = privateMibOid + "6.2.9.1.2@i"
	igmpGroupMembershipInterval = privateMibOid + "6.2.9.1.3@i"
	// IPGS-6416XSFP-MIB igmpStatisticsTable                           table        1.3.6.1.4.1.37072.302.3.1.6.2.10
	// IPGS-6416XSFP-MIB igmpStatisticsTableEntry                      row          1.3.6.1.4.1.37072.302.3.1.6.2.10.1
	igmpStatisticsVlanID                 = privateMibOid + "6.2.10.1.1@i"
	igmpStatisticsQuerierStatus          = privateMibOid + "6.2.10.1.2@s"
	igmpStatisticsQuerierTransmitted     = privateMibOid + "6.2.10.1.3@i"
	igmpStatisticsQuerierReceived        = privateMibOid + "6.2.10.1.4@i"
	igmpStatisticsV1ReportsReceived      = privateMibOid + "6.2.10.1.5@i"
	igmpStatisticsV2LeaveReceived        = privateMibOid + "6.2.10.1.6@i"
	igmpStatisticsV2ReportsReceivedCount = privateMibOid + "6.2.10.1.7@i"
	igmpStatisticsV3ReportsReceivedCount = privateMibOid + "6.2.10.1.8@i"
	// IPGS-6416XSFP-MIB discovery                                     node         1.3.6.1.4.1.37072.302.3.1.7
	// IPGS-6416XSFP-MIB lldp                                          node         1.3.6.1.4.1.37072.302.3.1.7.1
	lldpEnabled    = privateMibOid + "7.1.1@i"
	lldpTxInterval = privateMibOid + "7.1.2@i"
	lldpLiveTime   = privateMibOid + "7.1.3@i"
	// IPGS-6416XSFP-MIB lldpModeTable                                 table        1.3.6.1.4.1.37072.302.3.1.7.1.4
	// IPGS-6416XSFP-MIB lldpModeTableEntry                            row          1.3.6.1.4.1.37072.302.3.1.7.1.4.1
	lldpPortIndex = privateMibOid + "7.1.4.1.1@i"
	lldpModeValue = privateMibOid + "7.1.4.1.2@i"
	// IPGS-6416XSFP-MIB lldpNeighborInfoTable                         table        1.3.6.1.4.1.37072.302.3.1.7.1.5
	// IPGS-6416XSFP-MIB lldpNeighborInfoTableEntry                    row          1.3.6.1.4.1.37072.302.3.1.7.1.5.1
	lldpNeighborLocalPort         = privateMibOid + "7.1.5.1.1@i"
	lldpNeighborChassisID         = privateMibOid + "7.1.5.1.2@s"
	lldpNeighborRemotePort        = privateMibOid + "7.1.5.1.3@s"
	lldpNeighborPortDescr         = privateMibOid + "7.1.5.1.4@s"
	lldpNeighborSystemName        = privateMibOid + "7.1.5.1.5@s"
	lldpNeighborSystemCapability  = privateMibOid + "7.1.5.1.6@s"
	lldpNeighborManagementAddress = privateMibOid + "7.1.5.1.7@s"
	// IPGS-6416XSFP-MIB lldpStatisticTable                            table        1.3.6.1.4.1.37072.302.3.1.7.1.6
	// IPGS-6416XSFP-MIB lldpStatisticTableEntry                       row          1.3.6.1.4.1.37072.302.3.1.7.1.6.1
	lldpStatisticEntryIndex            = privateMibOid + "7.1.6.1.1@i"
	lldpNeighborsAgedOut               = privateMibOid + "7.1.6.1.2@i"
	lldpNeighborsAdd                   = privateMibOid + "7.1.6.1.3@i"
	lldpNeighborsDelete                = privateMibOid + "7.1.6.1.4@i"
	lldpNeighborsFramesDiscarded       = privateMibOid + "7.1.6.1.5@i"
	lldpNeighborsFramesReceivedInError = privateMibOid + "7.1.6.1.6@i"
	lldpNeighborsFramesIn              = privateMibOid + "7.1.6.1.7@i"
	lldpNeighborsFramesOut             = privateMibOid + "7.1.6.1.8@i"
	lldpNeighborsTLVsDiscarded         = privateMibOid + "7.1.6.1.9@i"
	lldpNeighborsTLVsUnrecongnized     = privateMibOid + "7.1.6.1.10@i"
	// IPGS-6416XSFP-MIB cdp                                           node         1.3.6.1.4.1.37072.302.3.1.7.2
	cdpEnabled      = privateMibOid + "7.2.1@i"
	cdpUpdateTime   = privateMibOid + "7.2.2@i"
	cdpHoldTime     = privateMibOid + "7.2.3@i"
	cdpTxPackets    = privateMibOid + "7.2.4@i"
	cdpRxPackets    = privateMibOid + "7.2.5@i"
	cdpClearPackets = privateMibOid + "7.2.6@i"
	// IPGS-6416XSFP-MIB cdpStatisticTable                             table        1.3.6.1.4.1.37072.302.3.1.7.2.7
	// IPGS-6416XSFP-MIB cdpStatisticTableEntry                        row          1.3.6.1.4.1.37072.302.3.1.7.2.7.1
	cdpLocalPort             = privateMibOid + "7.2.7.1.1@i"
	cdpVersion               = privateMibOid + "7.2.7.1.2@s"
	cdpDeviceID              = privateMibOid + "7.2.7.1.3@s"
	cdpRemotePortID          = privateMibOid + "7.2.7.1.4@s"
	cdpRemotePlatform        = privateMibOid + "7.2.7.1.5@s"
	cdpRemoteSoftwareVersion = privateMibOid + "7.2.7.1.6@s"
	cdpAgeOutTTL             = privateMibOid + "7.2.7.1.7@i"
	cdpRemoteAddress         = privateMibOid + "7.2.7.1.8@i"
	// IPGS-6416XSFP-MIB topology                                      node         1.3.6.1.4.1.37072.302.3.1.7.3
	// IPGS-6416XSFP-MIB topologyNodeTable                             table        1.3.6.1.4.1.37072.302.3.1.7.3.1
	// IPGS-6416XSFP-MIB topologyNodeTableEntry                        row          1.3.6.1.4.1.37072.302.3.1.7.3.1.1
	topoNodeMacAddress = privateMibOid + "7.3.1.1.1@s"
	topoNodeIPAddress  = privateMibOid + "7.3.1.1.2@s"
	// IPGS-6416XSFP-MIB topologyLinkTable                             table        1.3.6.1.4.1.37072.302.3.1.7.3.2
	// IPGS-6416XSFP-MIB topologyLinkTableEntry                        row          1.3.6.1.4.1.37072.302.3.1.7.3.2.1
	topoLinkIndex    = privateMibOid + "7.3.2.1.1@i"
	topoFromPortNum  = privateMibOid + "7.3.2.1.2@i"
	topoToMacAddress = privateMibOid + "7.3.2.1.3@s"
	topoToPortNum    = privateMibOid + "7.3.2.1.4@i"
	// IPGS-6416XSFP-MIB topologyRingTable                             table        1.3.6.1.4.1.37072.302.3.1.7.3.3
	// IPGS-6416XSFP-MIB topologyRingTableEntry                        row          1.3.6.1.4.1.37072.302.3.1.7.3.3.1
	topoRingID          = privateMibOid + "7.3.3.1.1@i"
	topoRingType        = privateMibOid + "7.3.3.1.2@s"
	topoRingState       = privateMibOid + "7.3.3.1.3@s"
	topoRingRole        = privateMibOid + "7.3.3.1.4@s"
	topoRingPort0       = privateMibOid + "7.3.3.1.5@i"
	topoRingPort1       = privateMibOid + "7.3.3.1.6@i"
	topoIsBlockingPort0 = privateMibOid + "7.3.3.1.7@i"
	topoIsBlockingPort1 = privateMibOid + "7.3.3.1.8@i"
	// IPGS-6416XSFP-MIB dhcp                                          node         1.3.6.1.4.1.37072.302.3.1.8
	// IPGS-6416XSFP-MIB dhcpServer                                    node         1.3.6.1.4.1.37072.302.3.1.8.1
	dhcpServerEnabled      = privateMibOid + "8.1.1@i"
	dhcpServerIPRangeFirst = privateMibOid + "8.1.2@s"
	dhcpServerIPRangeLast  = privateMibOid + "8.1.3@s"
	dhcpServerNetmask      = privateMibOid + "8.1.4@s"
	dhcpServerGateway      = privateMibOid + "8.1.5@s"
	dhcpServerDNS          = privateMibOid + "8.1.6@s"
	dhcpServerLeaseTime    = privateMibOid + "8.1.7@i"
	// IPGS-6416XSFP-MIB dhcpServerMacBasedTableGroup                  node         1.3.6.1.4.1.37072.302.3.1.8.1.8
	dhcpServerMacBasedAddCfg    = privateMibOid + "8.1.8.1@s"
	dhcpServerMacBasedDeleteCfg = privateMibOid + "8.1.8.2@s"
	dhcpServerMacBasedCfgRow    = privateMibOid + "8.1.8.3@s"
	// IPGS-6416XSFP-MIB dhcpServerMacBasedTable                       table        1.3.6.1.4.1.37072.302.3.1.8.1.8.4
	// IPGS-6416XSFP-MIB dhcpServerMacBasedTableEntry                  row          1.3.6.1.4.1.37072.302.3.1.8.1.8.4.1
	dhcpServerMacBasedID         = privateMibOid + "8.1.8.4.1.1@s"
	dhcpServerMacBasedIP         = privateMibOid + "8.1.8.4.1.2@s"
	dhcpServerOption66ServerName = privateMibOid + "8.1.10@s"
	// IPGS-6416XSFP-MIB dhcpServerOption82TableGroup                  node         1.3.6.1.4.1.37072.302.3.1.8.1.11
	dhcpServerOption82AddCfg    = privateMibOid + "8.1.11.1@s"
	dhcpServerOption82DeleteCfg = privateMibOid + "8.1.11.2@s"
	dhcpServerOption82CfgRow    = privateMibOid + "8.1.11.3@s"
	// IPGS-6416XSFP-MIB dhcpServerOption82Table                       table        1.3.6.1.4.1.37072.302.3.1.8.1.11.4
	// IPGS-6416XSFP-MIB dhcpServerOption82TableEntry                  row          1.3.6.1.4.1.37072.302.3.1.8.1.11.4.1
	dhcpServerOption82RemoteID  = privateMibOid + "8.1.11.4.1.1@s"
	dhcpServerOption82CircuitID = privateMibOid + "8.1.11.4.1.2@s"
	dhcpServerOption82IPFirst   = privateMibOid + "8.1.11.4.1.3@s"
	dhcpServerOption82IPLast    = privateMibOid + "8.1.11.4.1.4@s"
	dhcpServerOption82Netmask   = privateMibOid + "8.1.11.4.1.5@s"
	dhcpServerOption82Gateway   = privateMibOid + "8.1.11.4.1.6@s"
	dhcpServerOption82DNS       = privateMibOid + "8.1.11.4.1.7@s"
	dhcpServerOPtion82LeaseTime = privateMibOid + "8.1.11.4.1.8@i"
	dhcpServerOptionPXEEnabled  = privateMibOid + "8.1.12@i"
	// IPGS-6416XSFP-MIB dhcpServerOptionPXECfgTableGroup              node         1.3.6.1.4.1.37072.302.3.1.8.1.13
	dhcpServerOptionPXEAddCfg     = privateMibOid + "8.1.13.1@s"
	dhcpServerOptionPXEDeleteCfg  = privateMibOid + "8.1.13.2@s"
	dhcpServerOptionPXECfgRow     = privateMibOid + "8.1.13.3@s"
	dhcpServerOptionAddPXEInfo    = privateMibOid + "8.1.14@s"
	dhcpServerOptionDeletePXEInfo = privateMibOid + "8.1.15@s"
	// IPGS-6416XSFP-MIB dhcpServerOptionPXECfg1InfoTable              table        1.3.6.1.4.1.37072.302.3.1.8.1.16
	// IPGS-6416XSFP-MIB dhcpServerOptionPXECfg1InfoTableEntry         row          1.3.6.1.4.1.37072.302.3.1.8.1.16.1
	dhcpServerOptionPXECfg1InfoIndex              = privateMibOid + "8.1.16.1.1@i"
	dhcpServerOptionPXECfg1InfoTFTPServerName     = privateMibOid + "8.1.16.1.2@s"
	dhcpServerOptionPXECfg1InfoBootFileName       = privateMibOid + "8.1.16.1.3@s"
	dhcpServerOptionPXECfg1InfoSystemArchitecture = privateMibOid + "8.1.16.1.4@i"
	dhcpServerOptionPXECfg1InfoTFTPServerIP       = privateMibOid + "8.1.16.1.5@i"
	// IPGS-6416XSFP-MIB dhcpServerOptionPXECfg2InfoTable              table        1.3.6.1.4.1.37072.302.3.1.8.1.17
	// IPGS-6416XSFP-MIB dhcpServerOptionPXECfg2InfoTableEntry         row          1.3.6.1.4.1.37072.302.3.1.8.1.17.1
	dhcpServerOptionPXECfg2InfoIndex              = privateMibOid + "8.1.17.1.1@i"
	dhcpServerOptionPXECfg2InfoTFTPServerName     = privateMibOid + "8.1.17.1.2@s"
	dhcpServerOptionPXECfg2InfoBootFileName       = privateMibOid + "8.1.17.1.3@s"
	dhcpServerOptionPXECfg2InfoSystemArchitecture = privateMibOid + "8.1.17.1.4@i"
	dhcpServerOptionPXECfg2InfoTFTPServerIP       = privateMibOid + "8.1.17.1.5@i"
	// IPGS-6416XSFP-MIB dhcpServerOptionPXECfg3InfoTable              table        1.3.6.1.4.1.37072.302.3.1.8.1.18
	// IPGS-6416XSFP-MIB dhcpServerOptionPXECfg3InfoTableEntry         row          1.3.6.1.4.1.37072.302.3.1.8.1.18.1
	dhcpServerOptionPXECfg3InfoIndex              = privateMibOid + "8.1.18.1.1@i"
	dhcpServerOptionPXECfg3InfoTFTPServerName     = privateMibOid + "8.1.18.1.2@s"
	dhcpServerOptionPXECfg3InfoBootFileName       = privateMibOid + "8.1.18.1.3@s"
	dhcpServerOptionPXECfg3InfoSystemArchitecture = privateMibOid + "8.1.18.1.4@i"
	dhcpServerOptionPXECfg3InfoTFTPServerIP       = privateMibOid + "8.1.18.1.5@i"
	// IPGS-6416XSFP-MIB dhcpServerOptionPXECfg4InfoTable              table        1.3.6.1.4.1.37072.302.3.1.8.1.19
	// IPGS-6416XSFP-MIB dhcpServerOptionPXECfg4InfoTableEntry         row          1.3.6.1.4.1.37072.302.3.1.8.1.19.1
	dhcpServerOptionPXECfg4InfoIndex              = privateMibOid + "8.1.19.1.1@i"
	dhcpServerOptionPXECfg4InfoTFTPServerName     = privateMibOid + "8.1.19.1.2@s"
	dhcpServerOptionPXECfg4InfoBootFileName       = privateMibOid + "8.1.19.1.3@s"
	dhcpServerOptionPXECfg4InfoSystemArchitecture = privateMibOid + "8.1.19.1.4@i"
	dhcpServerOptionPXECfg4InfoTFTPServerIP       = privateMibOid + "8.1.19.1.5@i"
	// IPGS-6416XSFP-MIB dhcpServerStatusTable                         table        1.3.6.1.4.1.37072.302.3.1.8.1.20
	// IPGS-6416XSFP-MIB dhcpServerStatusTableEntry                    row          1.3.6.1.4.1.37072.302.3.1.8.1.20.1
	dhcpCurrentMacAddr             = privateMibOid + "8.1.20.1.1@s"
	dhcpCurrentIPAddr              = privateMibOid + "8.1.20.1.2@s"
	dhcpCurrentName                = privateMibOid + "8.1.20.1.3@s"
	dhcpCurrentAvailableLeasedTime = privateMibOid + "8.1.20.1.4@i"
	// IPGS-6416XSFP-MIB dhcpRelay                                     node         1.3.6.1.4.1.37072.302.3.1.8.2
	dhcpRelayCfgEnabled  = privateMibOid + "8.2.1@i"
	dhcpRelayCfgRemoteID = privateMibOid + "8.2.2@s"
	// IPGS-6416XSFP-MIB dhcpRelayCfgTable                             table        1.3.6.1.4.1.37072.302.3.1.8.2.3
	// IPGS-6416XSFP-MIB dhcpRelayServerCfgTableEntry                  row          1.3.6.1.4.1.37072.302.3.1.8.2.3.1
	dhcpRelayServerVlanID              = privateMibOid + "8.2.3.1.1@i"
	dhcpRelayServerIP                  = privateMibOid + "8.2.3.1.2@s"
	dhcpRelayStatisticsFromClient      = privateMibOid + "8.2.4@i"
	dhcpRelayStatisticsFromServer      = privateMibOid + "8.2.5@i"
	dhcpRelayStatisticsTxToClient      = privateMibOid + "8.2.6@i"
	dhcpRelayStatisticsTxToClientError = privateMibOid + "8.2.7@i"
	dhcpRelayStatisticsTxToServer      = privateMibOid + "8.2.8@i"
	dhcpRelayStatisticsTxToServerError = privateMibOid + "8.2.9@i"
	dhcpRelayStatisticsUnknown         = privateMibOid + "8.2.10@i"
	dhcpRelayServerPingStatus          = privateMibOid + "8.2.11@i"
	// IPGS-6416XSFP-MIB dhcpRelayCircuitIDTable                       table        1.3.6.1.4.1.37072.302.3.1.8.2.12
	// IPGS-6416XSFP-MIB dhcpRelayCircuitIDTableEntry                  row          1.3.6.1.4.1.37072.302.3.1.8.2.12.1
	dhcpRelayCircuitPortNo = privateMibOid + "8.2.12.1.1@i"
	dhcpRelayCircuitID     = privateMibOid + "8.2.12.1.2@s"
	// IPGS-6416XSFP-MIB dhcpSnooping                                  node         1.3.6.1.4.1.37072.302.3.1.8.3
	dhcpSnoopingEnabled = privateMibOid + "8.3.1@i"
	// IPGS-6416XSFP-MIB dhcpSnoopingCfgTable                          table        1.3.6.1.4.1.37072.302.3.1.8.3.2
	// IPGS-6416XSFP-MIB dhcpSnoopingCfgTableEntry                     row          1.3.6.1.4.1.37072.302.3.1.8.3.2.1
	dhcpSnoopingPortNo = privateMibOid + "8.3.2.1.1@i"
	dhcpSnoopingMode   = privateMibOid + "8.3.2.1.2@i"
	// IPGS-6416XSFP-MIB dhcpSnoopingInfoTable                         table        1.3.6.1.4.1.37072.302.3.1.8.3.3
	// IPGS-6416XSFP-MIB dhcpSnoopingInfoTableEntry                    row          1.3.6.1.4.1.37072.302.3.1.8.3.3.1
	dhcpSnoopingInfoPortNo      = privateMibOid + "8.3.3.1.1@i"
	dhcpSnoopingInfoMacAddr     = privateMibOid + "8.3.3.1.2@s"
	dhcpSnoopingInfoIpAddr      = privateMibOid + "8.3.3.1.3@s"
	dhcpSnoopingInfoLeaseTime   = privateMibOid + "8.3.3.1.4@i"
	dhcpSnoopingInfoBindingType = privateMibOid + "8.3.3.1.5@s"
	dhcpSnoopingInfoVlanID      = privateMibOid + "8.3.3.1.6@i"
	// IPGS-6416XSFP-MIB dhcpSnoopingRXStatisticTable                  table        1.3.6.1.4.1.37072.302.3.1.8.3.4
	// IPGS-6416XSFP-MIB dhcpSnoopingRXStatisticTableEntry             row          1.3.6.1.4.1.37072.302.3.1.8.3.4.1
	dhcpSnoopingRXPort            = privateMibOid + "8.3.4.1.1@i"
	dhcpSnoopingRXDiscover        = privateMibOid + "8.3.4.1.2@i"
	dhcpSnoopingRXOffer           = privateMibOid + "8.3.4.1.3@i"
	dhcpSnoopingRXRequest         = privateMibOid + "8.3.4.1.4@i"
	dhcpSnoopingRXDecline         = privateMibOid + "8.3.4.1.5@i"
	dhcpSnoopingRXAck             = privateMibOid + "8.3.4.1.6@i"
	dhcpSnoopingRXNak             = privateMibOid + "8.3.4.1.7@i"
	dhcpSnoopingRXRelease         = privateMibOid + "8.3.4.1.8@i"
	dhcpSnoopingRXInform          = privateMibOid + "8.3.4.1.9@i"
	dhcpSnoopingRXLeaseQuery      = privateMibOid + "8.3.4.1.10@i"
	dhcpSnoopingRXLeaseUnassigned = privateMibOid + "8.3.4.1.11@i"
	dhcpSnoopingRXLeaseUnKnown    = privateMibOid + "8.3.4.1.12@i"
	dhcpSnoopingRXLeaseActive     = privateMibOid + "8.3.4.1.13@i"
	// IPGS-6416XSFP-MIB dhcpSnoopingTXStatisticTable                  table        1.3.6.1.4.1.37072.302.3.1.8.3.5
	// IPGS-6416XSFP-MIB dhcpSnoopingTXStatisticTableEntry             row          1.3.6.1.4.1.37072.302.3.1.8.3.5.1
	dhcpSnoopingTXPort            = privateMibOid + "8.3.5.1.1@i"
	dhcpSnoopingTXDiscover        = privateMibOid + "8.3.5.1.2@i"
	dhcpSnoopingTXOffer           = privateMibOid + "8.3.5.1.3@i"
	dhcpSnoopingTXRequest         = privateMibOid + "8.3.5.1.4@i"
	dhcpSnoopingTXDecline         = privateMibOid + "8.3.5.1.5@i"
	dhcpSnoopingTXAck             = privateMibOid + "8.3.5.1.6@i"
	dhcpSnoopingTXNak             = privateMibOid + "8.3.5.1.7@i"
	dhcpSnoopingTXRelease         = privateMibOid + "8.3.5.1.8@i"
	dhcpSnoopingTXInform          = privateMibOid + "8.3.5.1.9@i"
	dhcpSnoopingTXLeaseQuery      = privateMibOid + "8.3.5.1.10@i"
	dhcpSnoopingTXLeaseUnassigned = privateMibOid + "8.3.5.1.11@i"
	dhcpSnoopingTXLeaseUnKnown    = privateMibOid + "8.3.5.1.12@i"
	dhcpSnoopingTXLeaseActive     = privateMibOid + "8.3.5.1.13@i"
	// IPGS-6416XSFP-MIB dhcpServerPortBasedTable                      table        1.3.6.1.4.1.37072.302.3.1.8.9
	// IPGS-6416XSFP-MIB dhcpServerPortBasedTableEntry                 row          1.3.6.1.4.1.37072.302.3.1.8.9.1
	dhcpServerPortBasedIndex      = privateMibOid + "8.9.1.1@i"
	dhcpServerPortBasedIP         = privateMibOid + "8.9.1.2@s"
	dhcpServerPortBasedNotOfferIP = privateMibOid + "8.9.1.3@i"
	// IPGS-6416XSFP-MIB redundant                                     node         1.3.6.1.4.1.37072.302.3.1.9
	// IPGS-6416XSFP-MIB stp                                           node         1.3.6.1.4.1.37072.302.3.1.9.1
	stpMode           = privateMibOid + "9.1.1@i-s"
	stpName           = privateMibOid + "9.1.2@s-s"
	stpRevision       = privateMibOid + "9.1.3@i"
	stpForwardDelay   = privateMibOid + "9.1.4@i"
	stpMaxAge         = privateMibOid + "9.1.5@i"
	stpMaxHops        = privateMibOid + "9.1.6@i"
	stpBridgePriority = privateMibOid + "9.1.7@i"
	// IPGS-6416XSFP-MIB stpCISTCfgTable                               table        1.3.6.1.4.1.37072.302.3.1.9.1.8
	// IPGS-6416XSFP-MIB stpCISTCfgTableEntry                          row          1.3.6.1.4.1.37072.302.3.1.9.1.8.1
	stpCISTPortIndex = privateMibOid + "9.1.8.1.1@i"
	stpCISTEnabled   = privateMibOid + "9.1.8.1.2@i"
	stpCISTPathCost  = privateMibOid + "9.1.8.1.3@i"
	stpCISTPriority  = privateMibOid + "9.1.8.1.4@i"
	stpCISTEdgeMode  = privateMibOid + "9.1.8.1.5@i"
	stpCISTP2PMode   = privateMibOid + "9.1.8.1.6@s"
	// IPGS-6416XSFP-MIB stpMSTIInstanceTable                          table        1.3.6.1.4.1.37072.302.3.1.9.1.9
	// IPGS-6416XSFP-MIB stpMSTIInstanceTableEntry                     row          1.3.6.1.4.1.37072.302.3.1.9.1.9.1
	stpMSTIIndex    = privateMibOid + "9.1.9.1.1@i"
	stpMSTIEnabled  = privateMibOid + "9.1.9.1.2@i"
	stpMSTIVlans    = privateMibOid + "9.1.9.1.3@s"
	stpMSTIPriority = privateMibOid + "9.1.9.1.4@i"
	// IPGS-6416XSFP-MIB stpMSTIInstancePortDetailTable                table        1.3.6.1.4.1.37072.302.3.1.9.1.10
	// IPGS-6416XSFP-MIB stpMSTIInstancePortDetailTableEntry           row          1.3.6.1.4.1.37072.302.3.1.9.1.10.1
	stpMSTIPortDetailIndex = privateMibOid + "9.1.10.1.1@i"
	stpMSTIPortNo          = privateMibOid + "9.1.10.1.2@i"
	stpMSTIPortPathCost    = privateMibOid + "9.1.10.1.3@i"
	stpMSTIPortPriority    = privateMibOid + "9.1.10.1.4@i"
	// IPGS-6416XSFP-MIB stpBridgeStatusTable                          table        1.3.6.1.4.1.37072.302.3.1.9.1.11
	// IPGS-6416XSFP-MIB stpBridgeStatusTableEntry                     row          1.3.6.1.4.1.37072.302.3.1.9.1.11.1
	stpBridgeNo       = privateMibOid + "9.1.11.1.1@s"
	stpBridgeID       = privateMibOid + "9.1.11.1.2@s"
	stpBridgeRootID   = privateMibOid + "9.1.11.1.3@s"
	stpBridgeRootPort = privateMibOid + "9.1.11.1.4@i"
	stpBridgeRootCost = privateMibOid + "9.1.11.1.5@i"
	// IPGS-6416XSFP-MIB stpPortStatusTable                            table        1.3.6.1.4.1.37072.302.3.1.9.1.12
	// IPGS-6416XSFP-MIB stpPortStatusTableEntry                       row          1.3.6.1.4.1.37072.302.3.1.9.1.12.1
	stpPortStatusPortIndex = privateMibOid + "9.1.12.1.1@i"
	stpPortStatusCIST0     = privateMibOid + "9.1.12.1.2@s"
	stpPortStatusMSTI1     = privateMibOid + "9.1.12.1.3@s"
	stpPortStatusMSTI2     = privateMibOid + "9.1.12.1.4@s"
	stpPortStatusMSTI3     = privateMibOid + "9.1.12.1.5@s"
	stpPortStatusMSTI4     = privateMibOid + "9.1.12.1.6@s"
	stpPortStatusMSTI5     = privateMibOid + "9.1.12.1.7@s"
	stpPortStatusMSTI6     = privateMibOid + "9.1.12.1.8@s"
	stpPortStatusMSTI7     = privateMibOid + "9.1.12.1.9@s"
	stpPortStatusMSTI8     = privateMibOid + "9.1.12.1.10@s"
	// IPGS-6416XSFP-MIB loopProtection                                node         1.3.6.1.4.1.37072.302.3.1.9.2
	loopProtectionEnable        = privateMibOid + "9.2.1@i"
	loopProtectionEnableOnPorts = privateMibOid + "9.2.2@s"
	loopProtectionInterval      = privateMibOid + "9.2.3@i"
	loopProtectionShutDown      = privateMibOid + "9.2.4@i"
	// IPGS-6416XSFP-MIB loopProtectionStatusTable                     table        1.3.6.1.4.1.37072.302.3.1.9.2.5
	// IPGS-6416XSFP-MIB loopProtectionStatusTableEntry                row          1.3.6.1.4.1.37072.302.3.1.9.2.5.1
	loopProtectionPortIndex    = privateMibOid + "9.2.5.1.1@i"
	loopProtectionIsLooping    = privateMibOid + "9.2.5.1.2@i"
	loopProtectionLoopCounts   = privateMibOid + "9.2.5.1.3@i"
	loopProtectionLastLoopTime = privateMibOid + "9.2.5.1.4@s"
	// IPGS-6416XSFP-MIB dot8032ERPS                                   node         1.3.6.1.4.1.37072.302.3.1.9.3
	dot8032RingMode = privateMibOid + "9.3.1@i"
	// IPGS-6416XSFP-MIB dot8032CfgTableGroup                          node         1.3.6.1.4.1.37072.302.3.1.9.3.2
	dot8032AddCfg    = privateMibOid + "9.3.2.1@s"
	dot8032DeleteCfg = privateMibOid + "9.3.2.2@s"
	dot8032CfgRow    = privateMibOid + "9.3.2.3@s"
	// IPGS-6416XSFP-MIB dot8032CfgTable                               table        1.3.6.1.4.1.37072.302.3.1.9.3.2.4
	// IPGS-6416XSFP-MIB dot8032CfgTableEntry                          row          1.3.6.1.4.1.37072.302.3.1.9.3.2.4.1
	dot8032RingID                = privateMibOid + "9.3.2.4.1.1@i"
	dot8032RingEnabled           = privateMibOid + "9.3.2.4.1.2@i"
	dot8032RingRole              = privateMibOid + "9.3.2.4.1.3@i"
	dot8032RingType              = privateMibOid + "9.3.2.4.1.4@i"
	dot8032RingPort0             = privateMibOid + "9.3.2.4.1.5@i"
	dot8032RingPort1             = privateMibOid + "9.3.2.4.1.6@i"
	dot8032NodeFailureProtection = privateMibOid + "9.3.2.4.1.7@i"
	dot8032DetectMiswiring       = privateMibOid + "9.3.2.4.1.8@i"
	// IPGS-6416XSFP-MIB dot8032StatusTable                            table        1.3.6.1.4.1.37072.302.3.1.9.3.3
	// IPGS-6416XSFP-MIB dot8032StatusTableEntry                       row          1.3.6.1.4.1.37072.302.3.1.9.3.3.1
	dot8032CurrentRingID          = privateMibOid + "9.3.3.1.1@i"
	dot8032CurrentRingEnabled     = privateMibOid + "9.3.3.1.2@i"
	dot8032CurrentRingState       = privateMibOid + "9.3.3.1.3@s"
	dot8032CurrentRingRole        = privateMibOid + "9.3.3.1.4@s"
	dot8032CurrentRingType        = privateMibOid + "9.3.3.1.5@s"
	dot8032CurrentRingPort0Status = privateMibOid + "9.3.3.1.6@s"
	dot8032CurrentRingPort1Status = privateMibOid + "9.3.3.1.7@s"
	// IPGS-6416XSFP-MIB dualHoming                                    node         1.3.6.1.4.1.37072.302.3.1.9.4
	// IPGS-6416XSFP-MIB dualHomingTableGroup                          node         1.3.6.1.4.1.37072.302.3.1.9.4.1
	dualHomingAddCfg    = privateMibOid + "9.4.1.1@s"
	dualHomingDeleteCfg = privateMibOid + "9.4.1.2@s"
	dualHomingCfgRow    = privateMibOid + "9.4.1.3@s"
	// IPGS-6416XSFP-MIB dualHomingTable                               table        1.3.6.1.4.1.37072.302.3.1.9.4.1.4
	// IPGS-6416XSFP-MIB dualHomingTableEntry                          row          1.3.6.1.4.1.37072.302.3.1.9.4.1.4.1
	dualHomingRingID  = privateMibOid + "9.4.1.4.1.1@i"
	dualHomingEnabled = privateMibOid + "9.4.1.4.1.2@i"
	dualHomingRole    = privateMibOid + "9.4.1.4.1.3@s"
	dualHomingPort    = privateMibOid + "9.4.1.4.1.4@i"
	// IPGS-6416XSFP-MIB dualHomingStatusTable                         table        1.3.6.1.4.1.37072.302.3.1.9.4.2
	// IPGS-6416XSFP-MIB dualHomingStatusTableEntry                    row          1.3.6.1.4.1.37072.302.3.1.9.4.2.1
	dualHomingCurrentRingID  = privateMibOid + "9.4.2.1.1@i"
	dualHomingCurrentEnabled = privateMibOid + "9.4.2.1.2@i"
	dualHomingCurrentStatus  = privateMibOid + "9.4.2.1.3@s"
	dualHomingCurrentPort    = privateMibOid + "9.4.2.1.4@i"
	dualHomingCurrentRole    = privateMibOid + "9.4.2.1.5@s"
	// IPGS-6416XSFP-MIB mrp                                           node         1.3.6.1.4.1.37072.302.3.1.9.5
	// IPGS-6416XSFP-MIB mrpCfgTableGroup                              node         1.3.6.1.4.1.37072.302.3.1.9.5.1
	mrpAddCfg    = privateMibOid + "9.5.1.1@s"
	mrpDeleteCfg = privateMibOid + "9.5.1.2@s"
	mrpCfgRow    = privateMibOid + "9.5.1.3@s"
	// IPGS-6416XSFP-MIB mrpCfgTable                                   table        1.3.6.1.4.1.37072.302.3.1.9.5.1.4
	// IPGS-6416XSFP-MIB mrpCfgTableEntry                              row          1.3.6.1.4.1.37072.302.3.1.9.5.1.4.1
	mrpIndex         = privateMibOid + "9.5.1.4.1.1@i"
	mrpRole          = privateMibOid + "9.5.1.4.1.2@i"
	mrpDomainID      = privateMibOid + "9.5.1.4.1.3@s"
	mrpVlanID        = privateMibOid + "9.5.1.4.1.4@i"
	mrpPrimaryPort   = privateMibOid + "9.5.1.4.1.5@i"
	mrpSecondaryPort = privateMibOid + "9.5.1.4.1.6@i"
	// IPGS-6416XSFP-MIB mrpStatusTable                                table        1.3.6.1.4.1.37072.302.3.1.9.5.2
	// IPGS-6416XSFP-MIB mrpStatusTableEntry                           row          1.3.6.1.4.1.37072.302.3.1.9.5.2.1
	mrpStatusIndex         = privateMibOid + "9.5.2.1.1@i"
	mrpStatusRole          = privateMibOid + "9.5.2.1.2@i"
	mrpStatusDomainID      = privateMibOid + "9.5.2.1.3@s"
	mrpRoleState           = privateMibOid + "9.5.2.1.4@s"
	mrpRingState           = privateMibOid + "9.5.2.1.5@s"
	mrpPrimaryPortState    = privateMibOid + "9.5.2.1.6@s"
	mrpSecondaryPortState  = privateMibOid + "9.5.2.1.7@s"
	mrpRingTransitionCount = privateMibOid + "9.5.2.1.8@s"
	// IPGS-6416XSFP-MIB deviceSecurity                                node         1.3.6.1.4.1.37072.302.3.1.10
	// IPGS-6416XSFP-MIB macAddressTable                               node         1.3.6.1.4.1.37072.302.3.1.10.1
	// IPGS-6416XSFP-MIB staticMacAddressTableGroup                    node         1.3.6.1.4.1.37072.302.3.1.10.1.1
	staticMacAddressAddCfg    = privateMibOid + "10.1.1.1@s"
	staticMacAddressDeleteCfg = privateMibOid + "10.1.1.2@s"
	staticMacAddressCfgRow    = privateMibOid + "10.1.1.3@s"
	// IPGS-6416XSFP-MIB staticMacAddressTable                         table        1.3.6.1.4.1.37072.302.3.1.10.1.1.4
	// IPGS-6416XSFP-MIB staticMacAddressTableEntry                    row          1.3.6.1.4.1.37072.302.3.1.10.1.1.4.1
	staticMacAddressIndex  = privateMibOid + "10.1.1.4.1.1@i"
	staticMacAddress       = privateMibOid + "10.1.1.4.1.2@s"
	staticMacAddressVlanID = privateMibOid + "10.1.1.4.1.3@i"
	staticMacAddressPort   = privateMibOid + "10.1.1.4.1.4@i"
	// IPGS-6416XSFP-MIB macFilterTableGroup                           node         1.3.6.1.4.1.37072.302.3.1.10.1.2
	macFilterAddCfg    = privateMibOid + "10.1.2.1@s"
	macFilterDeleteCfg = privateMibOid + "10.1.2.2@s"
	macFilterCfgRow    = privateMibOid + "10.1.2.3@s"
	// IPGS-6416XSFP-MIB macFilterTable                                table        1.3.6.1.4.1.37072.302.3.1.10.1.2.4
	// IPGS-6416XSFP-MIB macFilterTableEntry                           row          1.3.6.1.4.1.37072.302.3.1.10.1.2.4.1
	macFilterEntryIndex = privateMibOid + "10.1.2.4.1.1@i"
	macFilterMacAddress = privateMibOid + "10.1.2.4.1.2@s"
	macFilterVlanID     = privateMibOid + "10.1.2.4.1.3@i"
	// IPGS-6416XSFP-MIB allMACAddressTable                            table        1.3.6.1.4.1.37072.302.3.1.10.1.3
	// IPGS-6416XSFP-MIB allMACAddressTableEntry                       row          1.3.6.1.4.1.37072.302.3.1.10.1.3.1
	allMACAddressVlanID = privateMibOid + "10.1.3.1.1@i"
	allMACAddressType   = privateMibOid + "10.1.3.1.2@s"
	allMACAddressValue  = privateMibOid + "10.1.3.1.3@s"
	allMACAddressPort   = privateMibOid + "10.1.3.1.4@i"
	// IPGS-6416XSFP-MIB accessControl                                 node         1.3.6.1.4.1.37072.302.3.1.10.2
	// IPGS-6416XSFP-MIB accessControlTableGroup                       node         1.3.6.1.4.1.37072.302.3.1.10.2.1
	accessControlAddCfg    = privateMibOid + "10.2.1.1@s"
	accessControlDeleteCfg = privateMibOid + "10.2.1.2@s"
	accessControlCfgRow    = privateMibOid + "10.2.1.3@s"
	// IPGS-6416XSFP-MIB accessControlTable                            table        1.3.6.1.4.1.37072.302.3.1.10.2.1.4
	// IPGS-6416XSFP-MIB accessControlTableEntry                       row          1.3.6.1.4.1.37072.302.3.1.10.2.1.4.1
	accessControlIndex              = privateMibOid + "10.2.1.4.1.1@i"
	accessControlDirection          = privateMibOid + "10.2.1.4.1.2@i"
	accessControlLookupRule         = privateMibOid + "10.2.1.4.1.3@i"
	accessControlIPProtocol         = privateMibOid + "10.2.1.4.1.4@i"
	accessControlSourceAddress      = privateMibOid + "10.2.1.4.1.5@s"
	accessControlSourceMask         = privateMibOid + "10.2.1.4.1.6@s"
	accessControlSourcePort         = privateMibOid + "10.2.1.4.1.7@i"
	accessControlDestinationAddress = privateMibOid + "10.2.1.4.1.8@s"
	accessControlDestinationMask    = privateMibOid + "10.2.1.4.1.9@s"
	accessControlDestinationPort    = privateMibOid + "10.2.1.4.1.10@i"
	accessControlPorts              = privateMibOid + "10.2.1.4.1.11@s"
	accessControlAction             = privateMibOid + "10.2.1.4.1.12@i"
	// IPGS-6416XSFP-MIB remoteAdmin                                   node         1.3.6.1.4.1.37072.302.3.1.10.3
	// IPGS-6416XSFP-MIB aaaConfiguration                              node         1.3.6.1.4.1.37072.302.3.1.10.4
	// IPGS-6416XSFP-MIB eventLog                                      node         1.3.6.1.4.1.37072.302.3.1.11
	// IPGS-6416XSFP-MIB event                                         node         1.3.6.1.4.1.37072.302.3.1.11.1
	// IPGS-6416XSFP-MIB eventDDMTableGroup                            node         1.3.6.1.4.1.37072.302.3.1.11.1.1
	eventDDMTableUpdateCfg = privateMibOid + "11.1.1.1@s"
	eventDDMTableRowCfg    = privateMibOid + "11.1.1.2@s"
	// IPGS-6416XSFP-MIB eventDDMTable                                 table        1.3.6.1.4.1.37072.302.3.1.11.1.1.4
	// IPGS-6416XSFP-MIB eventDDMTableEntry                            row          1.3.6.1.4.1.37072.302.3.1.11.1.1.4.1
	eventDDMPortNo           = privateMibOid + "11.1.1.4.1.1@i"
	eventDDMVoltageLower     = privateMibOid + "11.1.1.4.1.2@i"
	eventDDMVoltageUpper     = privateMibOid + "11.1.1.4.1.3@i"
	eventDDMRxPowerLower     = privateMibOid + "11.1.1.4.1.4@i"
	eventDDMRxPowerUpper     = privateMibOid + "11.1.1.4.1.5@i"
	eventDDMTxPowerLower     = privateMibOid + "11.1.1.4.1.6@i"
	eventDDMTxPowerUpper     = privateMibOid + "11.1.1.4.1.7@i"
	eventDDMTxBiasLower      = privateMibOid + "11.1.1.4.1.8@i"
	eventDDMTxBiasUpper      = privateMibOid + "11.1.1.4.1.9@i"
	eventDDMTemperatureLower = privateMibOid + "11.1.1.4.1.10@i"
	eventDDMTemperatureUpper = privateMibOid + "11.1.1.4.1.11@i"
	// IPGS-6416XSFP-MIB eventMonitor                                  node         1.3.6.1.4.1.37072.302.3.1.11.1.2
	eventMonitorVoltageLower     = privateMibOid + "11.1.2.1@i"
	eventMonitorVoltageUpper     = privateMibOid + "11.1.2.2@i"
	eventMonitorPowerLower       = privateMibOid + "11.1.2.3@i"
	eventMonitorPowerUpper       = privateMibOid + "11.1.2.4@i"
	eventMonitorCurrentLower     = privateMibOid + "11.1.2.5@i"
	eventMonitorCurrentUpper     = privateMibOid + "11.1.2.6@i"
	eventMonitorTemperatureLower = privateMibOid + "11.1.2.7@i"
	eventMonitorTemparatureUpper = privateMibOid + "11.1.2.8@i"
	// IPGS-6416XSFP-MIB eventDigitalInputTable                        table        1.3.6.1.4.1.37072.302.3.1.11.1.3
	// IPGS-6416XSFP-MIB eventDigitalInputTableEntry                   row          1.3.6.1.4.1.37072.302.3.1.11.1.3.1
	eventDigitalInputIndex     = privateMibOid + "11.1.3.1.1@i"
	eventDigitalInputCondition = privateMibOid + "11.1.3.1.2@s"
	eventDigitalInputDescr     = privateMibOid + "11.1.3.1.3@s"
	// IPGS-6416XSFP-MIB eventActions                                  node         1.3.6.1.4.1.37072.302.3.1.11.2
	// IPGS-6416XSFP-MIB remotesystemLog                               node         1.3.6.1.4.1.37072.302.3.1.11.2.1
	// IPGS-6416XSFP-MIB remoteSystemLogGroup                          node         1.3.6.1.4.1.37072.302.3.1.11.2.1.1
	remoteSystemLogAddCfg    = privateMibOid + "11.2.1.1.1@s"
	remoteSystemLogDeleteCfg = privateMibOid + "11.2.1.1.2@s"
	remoteSystemLogCfgRow    = privateMibOid + "11.2.1.1.3@s"
	// IPGS-6416XSFP-MIB remoteSystemLogCfgTable                       table        1.3.6.1.4.1.37072.302.3.1.11.2.1.1.4
	// IPGS-6416XSFP-MIB remoteSystemLogCfgTableEntry                  row          1.3.6.1.4.1.37072.302.3.1.11.2.1.1.4.1
	remoteSystemLogCfgNum   = privateMibOid + "11.2.1.1.4.1.1@i"
	remoteSystemLogHost     = privateMibOid + "11.2.1.1.4.1.2@s"
	remoteSystemLogTag      = privateMibOid + "11.2.1.1.4.1.3@s"
	remoteSystemLogFacility = privateMibOid + "11.2.1.1.4.1.4@s"
	// IPGS-6416XSFP-MIB email                                         node         1.3.6.1.4.1.37072.302.3.1.11.2.2
	emailCfg = privateMibOid + "11.2.2.1@s"
	// IPGS-6416XSFP-MIB emailReceiverGroup                            node         1.3.6.1.4.1.37072.302.3.1.11.2.2.2
	emailReceiverAdd    = privateMibOid + "11.2.2.2.1@s"
	emailReceiverDelete = privateMibOid + "11.2.2.2.2@s"
	// IPGS-6416XSFP-MIB emailReceiverCfgTable                         table        1.3.6.1.4.1.37072.302.3.1.11.2.2.2.3
	// IPGS-6416XSFP-MIB emailReceiverCfgTableEntry                    row          1.3.6.1.4.1.37072.302.3.1.11.2.2.2.3.1
	emailReceiverCfgNum = privateMibOid + "11.2.2.2.3.1.1@i"
	emailReceiver       = privateMibOid + "11.2.2.2.3.1.2@s"
	// IPGS-6416XSFP-MIB sms                                           node         1.3.6.1.4.1.37072.302.3.1.11.2.3
	smsUsername = privateMibOid + "11.2.3.1@s-s"
	smsPassword = privateMibOid + "11.2.3.2@s-s"
	// IPGS-6416XSFP-MIB smsPhoneNumberGroup                           node         1.3.6.1.4.1.37072.302.3.1.11.2.3.3
	smsPhoneNumberAdd    = privateMibOid + "11.2.3.3.1@s-s"
	smsPhoneNumberDelete = privateMibOid + "11.2.3.3.2@s-s"
	// IPGS-6416XSFP-MIB smsPhoneNumberCfgTable                        table        1.3.6.1.4.1.37072.302.3.1.11.2.3.3.3
	// IPGS-6416XSFP-MIB smsPhoneNumberCfgTableEntry                   row          1.3.6.1.4.1.37072.302.3.1.11.2.3.3.3.1
	smsPhoneNumberCfgIndex = privateMibOid + "11.2.3.3.3.1.1@i"
	smsPhoneNumber         = privateMibOid + "11.2.3.3.3.1.2@s"
	// IPGS-6416XSFP-MIB dout                                          node         1.3.6.1.4.1.37072.302.3.1.11.2.4
	// IPGS-6416XSFP-MIB doutTable                                     table        1.3.6.1.4.1.37072.302.3.1.11.2.4.1
	// IPGS-6416XSFP-MIB doutTableEntry                                row          1.3.6.1.4.1.37072.302.3.1.11.2.4.1.1
	doutCfgNum = privateMibOid + "11.2.4.1.1.1@i"
	doutAction = privateMibOid + "11.2.4.1.1.2@s-s"
	doutStatus = privateMibOid + "11.2.4.1.1.3@s-s"
	testRelay  = privateMibOid + "11.2.5@i-s"
	// IPGS-6416XSFP-MIB eventActionMap                                node         1.3.6.1.4.1.37072.302.3.1.11.3
	// IPGS-6416XSFP-MIB systemEvent                                   node         1.3.6.1.4.1.37072.302.3.1.11.3.1
	authenticationSuccessEvent = privateMibOid + "11.3.1.1@s-s"
	authenticationFailedEvent  = privateMibOid + "11.3.1.2@s-s"
	deviceBootEvent            = privateMibOid + "11.3.1.3@s-s"
	ringTopologyChangeEvent    = privateMibOid + "11.3.1.4@s-s"
	// IPGS-6416XSFP-MIB poeEvent                                      node         1.3.6.1.4.1.37072.302.3.1.11.3.2
	poeDetectionFailedEvent = privateMibOid + "11.3.2.1@s-s"
	// IPGS-6416XSFP-MIB monitorEvent                                  node         1.3.6.1.4.1.37072.302.3.1.11.3.3
	hardWareMonitorEvent = privateMibOid + "11.3.3.1@s-s"
	// IPGS-6416XSFP-MIB ddmEventTable                                 table        1.3.6.1.4.1.37072.302.3.1.11.3.3.2
	// IPGS-6416XSFP-MIB ddmEventTableEntry                            row          1.3.6.1.4.1.37072.302.3.1.11.3.3.2.1
	ddmPortIndex  = privateMibOid + "11.3.3.2.1.1@i"
	ddmEventValue = privateMibOid + "11.3.3.2.1.2@s-s"
	// IPGS-6416XSFP-MIB diEventTable                                  table        1.3.6.1.4.1.37072.302.3.1.11.3.4
	// IPGS-6416XSFP-MIB diEventTableEntry                             row          1.3.6.1.4.1.37072.302.3.1.11.3.4.1
	diIndex      = privateMibOid + "11.3.4.1.1@i"
	diEventValue = privateMibOid + "11.3.4.1.2@s-s"
	// IPGS-6416XSFP-MIB powerEventTable                               table        1.3.6.1.4.1.37072.302.3.1.11.3.5
	// IPGS-6416XSFP-MIB powerEventTableEntry                          row          1.3.6.1.4.1.37072.302.3.1.11.3.5.1
	powerIndex    = privateMibOid + "11.3.5.1.1@i-s"
	powerOnEvent  = privateMibOid + "11.3.5.1.2@s-s"
	powerOffEvent = privateMibOid + "11.3.5.1.3@s-s"
	// IPGS-6416XSFP-MIB linkChangeEvent                               node         1.3.6.1.4.1.37072.302.3.1.11.3.6
	// IPGS-6416XSFP-MIB linkChangeEventTableEntry                     row          1.3.6.1.4.1.37072.302.3.1.11.3.6.1
	linkChangePortIndex = privateMibOid + "11.3.6.1.1@i"
	linkUpEvent         = privateMibOid + "11.3.6.1.2@s-s"
	linkDownEvent       = privateMibOid + "11.3.6.1.3@s-s"
	// IPGS-6416XSFP-MIB logs                                          node         1.3.6.1.4.1.37072.302.3.1.11.4
	enableLogs = privateMibOid + "11.4.1@i-s"
	// IPGS-6416XSFP-MIB logsTable                                     table        1.3.6.1.4.1.37072.302.3.1.11.4.2
	// IPGS-6416XSFP-MIB logsTableEntry                                row          1.3.6.1.4.1.37072.302.3.1.11.4.2.1
	logContent = privateMibOid + "11.4.2.1.2@s"
	// IPGS-6416XSFP-MIB diagnostic                                    node         1.3.6.1.4.1.37072.302.3.1.12
	// IPGS-6416XSFP-MIB arpTable                                      table        1.3.6.1.4.1.37072.302.3.1.12.1
	// IPGS-6416XSFP-MIB arpTableEntry                                 row          1.3.6.1.4.1.37072.302.3.1.12.1.1
	arpTableMacAddress = privateMibOid + "12.1.1.1@s"
	arpTableIPAddr     = privateMibOid + "12.1.1.2@s"
	// IPGS-6416XSFP-MIB ddmStatusTable                                table        1.3.6.1.4.1.37072.302.3.1.12.2
	// IPGS-6416XSFP-MIB ddmStatusTableEntry                           row          1.3.6.1.4.1.37072.302.3.1.12.2.1
	ddmStatusPortNum     = privateMibOid + "12.2.1.1@i"
	ddmStatusPortType    = privateMibOid + "12.2.1.2@s"
	ddmStatusLinked      = privateMibOid + "12.2.1.3@i"
	ddmStatusPlugIn      = privateMibOid + "12.2.1.4@i"
	ddmStatusSupportDDM  = privateMibOid + "12.2.1.5@i"
	ddmStatusTemperature = privateMibOid + "12.2.1.6@i"
	ddmStatusVoltage     = privateMibOid + "12.2.1.7@i"
	ddmStatusTxBias      = privateMibOid + "12.2.1.8@i"
	ddmStatusTxPower     = privateMibOid + "12.2.1.9@i"
	ddmStatusRxPopwer    = privateMibOid + "12.2.1.10@i"
	// IPGS-6416XSFP-MIB snmpCfg                                       node         1.3.6.1.4.1.37072.302.3.1.13
	// IPGS-6416XSFP-MIB maintenance                                   node         1.3.6.1.4.1.37072.302.3.1.14
	saveConfiguration = privateMibOid + "14.1@i-s"
	resetDefault      = privateMibOid + "14.2@i-s"
	upgradePath       = privateMibOid + "14.3@s-s"
	upgradeImage      = privateMibOid + "14.4@i"
	reboot            = privateMibOid + "14.5@i"

// IPGS-6416XSFP-MIB traps                                         node         1.3.6.1.4.1.37072.302.3.1.15

// ************** Private MIB *********************
)
