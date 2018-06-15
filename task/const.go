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

	snmpGet  = "snmpget -v 3 "
	snmpWalk = "snmpwalk -v 3 "
	snmpSet  = "snmpset -v 3 "

	snmpUser              = "-u walter "
	snmpSecurityLevel     = "-l authPriv "
	snmpAuthentication    = "-a MD5 "
	snmpAuthPassPhrase    = "-A 123456789 "
	snmpPrivateProtocol   = "-x DES "
	snmpPrivatePassPhrase = "-X 123456789 "

	deviceIP = "192.168.1.171 "
	// deviceIP = "192.168.15.10 "
	// deviceIP = "192.168.16.140 "

	snmpGetPrefix     = snmpGet + snmpUser + snmpSecurityLevel + snmpAuthentication + snmpAuthPassPhrase + snmpPrivateProtocol + snmpPrivatePassPhrase + deviceIP
	snmpWalkPrefix    = snmpWalk + snmpUser + snmpSecurityLevel + snmpAuthentication + snmpAuthPassPhrase + snmpPrivateProtocol + snmpPrivatePassPhrase + deviceIP
	snmpSetPrefix     = snmpSet + snmpUser + snmpSecurityLevel + snmpAuthentication + snmpAuthPassPhrase + snmpPrivateProtocol + snmpPrivatePassPhrase + deviceIP
	snmpSetTypeString = " string "
	snmpSetTypeInt    = " integer "
	snmpSetTypeIpaddr = " ipaddress "
	privateMibOid     = "1.3.6.1.4.1.37072.302.2.3."
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
	// ATTENTION: The type here is not the exact type in the MIB, but the type from snmpMIB program
	//
	// -ps  means ps
	// -m:  marks the known failed type:
	//                                  1 =>  need to set value to the device
	//                                  2 =>  Spec issue
	//                                  3 =>  snmp program issue

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
	dot1dStpPortProtocolMigration = mib2Prefix + "17.2.19.1.1@i-s-w"
	dot1dStpPortAdminEdgePort     = mib2Prefix + "17.2.19.1.2@i-s-w"
	dot1dStpPortOperEdgePort      = mib2Prefix + "17.2.19.1.3@i-w"
	dot1dStpPortAdminPointToPoint = mib2Prefix + "17.2.19.1.4@i-s-w"
	dot1dStpPortOperPointToPoint  = mib2Prefix + "17.2.19.1.5@i-w"
	dot1dStpPortAdminPathCost     = mib2Prefix + "17.2.19.1.6@i-s-w"

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

	// ************** Private MIB *********************
	// SYSTEM (1)
	systemName           = privateMibOid + "1.1.0@s-s"
	systemLocation       = privateMibOid + "1.2.0@s-s"
	systemContact        = privateMibOid + "1.3.0@s-s"
	systemDescr          = privateMibOid + "1.4.0@s-s"
	systemFwVersion      = privateMibOid + "1.5.0@s"
	systemMacaddress     = privateMibOid + "1.6.0@s"
	systemAutoLogoutTime = privateMibOid + "1.7.0@i-s"
	systemSerialNum      = privateMibOid + "1.8.0@i"

	// Setting(2)
	vlanPortCfgNum           = privateMibOid + "2.1.1.1.1@i-w"
	vlanMembers              = privateMibOid + "2.1.1.1.2@s-w"
	vlanTags                 = privateMibOid + "2.1.1.1.3@s-w-m:1"
	pvidCfgNum               = privateMibOid + "2.1.2.1.1@i-w"
	vlanPvid                 = privateMibOid + "2.1.2.1.2@i-w-s"
	vlanFrameType            = privateMibOid + "2.1.2.1.3@i-w-s"
	mvrCfgNum                = privateMibOid + "2.2.1.1.1@i-w-m:2" // Mvr (2.2)
	mvrCfgVid                = privateMibOid + "2.2.1.1.2@i-w-"
	mvrIPAddr                = privateMibOid + "2.2.1.1.3@s-w"
	mvrMembers               = privateMibOid + "2.2.1.1.4@s-w"
	igmpEnableQuerier        = privateMibOid + "2.3.1.0@i-s" // Igmp (2.3)
	igmpQuerierVersion       = privateMibOid + "2.3.2.0@i-s"
	igmpEnableSnooping       = privateMibOid + "2.3.3.0@i-s"
	igmpEnableFloodWellKnown = privateMibOid + "2.3.4.0@i-s"
	igmpPortNum              = privateMibOid + "2.3.5.1.1@i-w" // IgmpRouterTable (2.3.5)
	igmpRouterStatus         = privateMibOid + "2.3.5.1.2@i-w-s"
	igmpFastLeaveStatus      = privateMibOid + "2.3.5.1.3@i-w-s"
	igmpVidNum               = privateMibOid + "2.3.6.1.1@i-w" // IgmpStatisticsTable (2.3.6)
	igmpStatusQuerier        = privateMibOid + "2.3.6.1.2@s-w"
	igmpQuerierTx            = privateMibOid + "2.3.6.1.3@i-w"
	igmpQuerierRx            = privateMibOid + "2.3.6.1.4@i-w"
	igmpV1Rx                 = privateMibOid + "2.3.6.1.5@i-w"
	igmpV2Rx                 = privateMibOid + "2.3.6.1.6@i-w"
	igmpV3Rx                 = privateMibOid + "2.3.6.1.7@i-w"
	igmpV2Leave              = privateMibOid + "2.3.6.1.8@i-w"
	igmpEntriesEntryIndex    = privateMibOid + "2.3.7.1.1@s-w" // IgmpEntriesTable (2.3.7)
	igmpEntriesEntryIPAddr   = privateMibOid + "2.3.7.1.2@s-w"
	igmpEntriesEntryVID      = privateMibOid + "2.3.7.1.3@i-w"
	igmpEntriesEntryMembers  = privateMibOid + "2.3.7.1.4@i-w"

	// Status (3)
	lldpPortNum     = privateMibOid + "3.1.1.1.1@i-w" // LLDPInfo (3.1)
	lldpInfoContent = privateMibOid + "3.1.1.1.2@s-w-m:1"

	// Warning (11)
	faultAlarmPowerCfgNum        = privateMibOid + "11.1.1.1.1@i-w-m:3" // FaultAlarm (11.1)
	faultAlarmPowerStatus        = privateMibOid + "11.1.1.1.2@i-w-s-m:3"
	faultAlarmPortCfgNum         = privateMibOid + "11.1.2.1.1@i-w-m:3"
	faultAlarmPortLinkStatus     = privateMibOid + "11.1.2.1.2@i-w-s-m:3" // ===============================> here
	eventDDMEnabled              = privateMibOid + "11.2.1.1.0@i-s-m:3"   // EventDDMEnabled (11.2.1)
	eventDDMTemperatureLower     = privateMibOid + "11.2.1.2.0@s-s-m:3"
	eventDDMTemperatureUpper     = privateMibOid + "11.2.1.3.0@s-s-m:3"
	eventDDMVoltageLower         = privateMibOid + "11.2.1.4.0@s-s-m:3"
	eventDDMVoltageUpper         = privateMibOid + "11.2.1.5.0@s-s-m:3"
	eventDDMTxBiasLower          = privateMibOid + "11.2.1.6.0@s-s-m:3"
	eventDDMTTxBiasUpper         = privateMibOid + "11.2.1.7.0@s-s-m:3"
	eventDDMTxPowerLower         = privateMibOid + "11.2.1.8.0@s-s-m:3"
	eventDDMTxPowerUpper         = privateMibOid + "11.2.1.9.0@s-s-m:3"
	eventDDMRxPowerLower         = privateMibOid + "11.2.1.10.0@s-s-m:3"
	eventDDMRxPowerUpper         = privateMibOid + "11.2.1.11.0@s-s-m:3"
	eventMonitorEnabled          = privateMibOid + "11.2.2.1.0@i-s-m:3" // EventMonitor (11.2.1)
	eventMonitorTemperatureLower = privateMibOid + "11.2.2.2.0@s-s-m:3"
	eventMonitorTemperatureUpper = privateMibOid + "11.2.2.3.0@s-s-m:3"
	eventMonitorVoltageLower     = privateMibOid + "11.2.2.4.0@s-s-m:3"
	eventMonitorVoltageUpper     = privateMibOid + "11.2.2.5.0@s-s-m:3"
	eventMonitorCurrentLower     = privateMibOid + "11.2.2.6.0@s-s-m:3"
	eventMonitorCurrentUpper     = privateMibOid + "11.2.2.7.0@s-s-m:3"
	eventMonitorPowerLower       = privateMibOid + "11.2.2.8.0@s-s-m:3"
	eventMonitorPowerUpper       = privateMibOid + "11.2.2.9.0@s-s-m:3"
	eventPOEAPortCfgNum          = privateMibOid + "11.2.3.1.1.1@i-w" // EventPOEA (11.2.3)
	eventPOEAPingEnabled         = privateMibOid + "11.2.3.1.1.2@i-w-s"
	eventPOEAPingIPAddr          = privateMibOid + "11.2.3.1.1.3@s-s-m:1"
	eventPOEAPingInterval        = privateMibOid + "11.2.3.1.1.4@i-s-m:3"
	eventPOEAPingRetry           = privateMibOid + "11.2.3.1.1.5@i-s-m:3"
	eventPOEAPingReboot          = privateMibOid + "11.2.3.1.1.6@i-s-m:3"
	eventPOEAPingFailAction      = privateMibOid + "11.2.3.1.1.7@i-s-m:3"
	localLogEnable               = privateMibOid + "11.3.1.1.0@i-s-m:3" // ActionConfiguration (11.3)
	remoteSystemLogCfgNum        = privateMibOid + "11.3.2.1.1.1@i-m:3" // RemoteSystemLog (11.3.2)
	remoteSystemLogHost          = privateMibOid + "11.3.2.1.1.2@ip-m:3"
	remoteSystemLogTag           = privateMibOid + "11.3.2.1.1.3@s-m:3"
	remoteSystemLogFacility      = privateMibOid + "11.3.2.1.1.4@s-m:3"
	remoteSystemLogHostName      = privateMibOid + "11.3.2.1.1.5@s-m:3"
	// emailEnable                  = privateMibOid + "11.3.3.1.0@i-s-m:3"   // email (11.3.3.1) email is always enable, so no need to check
	emailServerUser      = privateMibOid + "11.3.3.2.1.0@s-s-m:3" // emailServer (11.3.3.2)
	emailServerPassword  = privateMibOid + "11.3.3.2.2.0@s-s-m:3"
	emailServerHost      = privateMibOid + "11.3.3.2.3.0@s-s-m:3"
	emailServerSSLEnable = privateMibOid + "11.3.3.2.4.0@i-s-m:3"
	emailSender          = privateMibOid + "11.3.3.3.0@s-s-m:3"
	emailSubject         = privateMibOid + "11.3.3.4.0@-s-m:3"
	emailCloudEnable     = privateMibOid + "11.3.3.5.0@i-s-m:3"
	emailReceiverCfgNum  = privateMibOid + "11.3.3.6.1.1@i-m:3"
	emailReceiverHost    = privateMibOid + "11.3.3.6.1.2@s-m:3"
	// smsEnable                    = privateMibOid + "11.3.4.1.0@i-s-m:3" // SMS (11.3.4)
	smsUser     = privateMibOid + "11.3.4.2.0@s-s-m:3"
	smsPassword = privateMibOid + "11.3.4.3.0@s-s-m:3"
	// smsSenderText              = privateMibOid + "11.3.4.4.0@s-s-m:3"   // no more senderText
	smsReceiverCfgNum          = privateMibOid + "11.3.4.5.1.1@i-m:3"
	smsReceiverPhone           = privateMibOid + "11.3.4.5.1.2@s-m:3"
	snmpResponseLocale         = privateMibOid + "11.3.5.1.1.0@i-s-m:3" // Snmp (11.3.5)
	snmpCommunityCfgNum        = privateMibOid + "11.3.5.1.2.1.1@i-m:3"
	snmpCommunityCfgString     = privateMibOid + "11.3.5.1.2.1.2@s-m:3"
	snmpCommunityCfgReadOnly   = privateMibOid + "11.3.5.1.2.1.3@i-m:3"
	snmpTrapCfgNum             = privateMibOid + "11.3.5.2.1.1.1@i-m:3" // Trap (11.3.5.2)
	snmpTrapCfgCommunity       = privateMibOid + "11.3.5.2.1.1.2@s-m:3"
	snmpTrapCfgIPAddress       = privateMibOid + "11.3.5.2.1.1.3@ip-m:3"
	snmpTrapCfgVersion         = privateMibOid + "11.3.5.2.1.1.4@i-m:3"
	snmpV3UserCfgNum           = privateMibOid + "11.3.5.3.1.1.1@i-m:3" // V3User (11.3.5.3)
	snmpV3UserCfgName          = privateMibOid + "11.3.5.3.1.1.2@s-m:3"
	snmpV3UserCfgSecurityLevel = privateMibOid + "11.3.5.3.1.1.3@i-m:3"
	snmpV3UserCfgAuthProtocal  = privateMibOid + "11.3.5.3.1.1.4@i-m:3"
	snmpV3UserCfgAuthPassword  = privateMibOid + "11.3.5.3.1.1.5@s-m:3"
	snmpV3UserCfgPrivProtocal  = privateMibOid + "11.3.5.3.1.1.6@s-m:3"
	snmpV3UserCfgPrivPassword  = privateMibOid + "11.3.5.3.1.1.7@s-m:3"
	doutCfgNum                 = privateMibOid + "11.3.6.1.1.1@i-m:3" // Dout (11.3.6)
	doutCfgEnable              = privateMibOid + "11.3.6.1.1.2@i-s-m:3"
	doutCfgAction              = privateMibOid + "11.3.6.1.1.3@i-s-m:3"
	deviceBootEvent            = privateMibOid + "11.4.1.1.0@i-s-m:3" // EventActionMap (11.4.1.1)
	authenticationFailureEvent = privateMibOid + "11.4.1.2.0@i-s-m:3"
	authenticationSuccessEvent = privateMibOid + "11.4.1.3.0@i-s-m:3"
	deviceDDMEvent             = privateMibOid + "11.4.1.4.0@i-s-m:3"
	devicePOEEvent             = privateMibOid + "11.4.1.5.0@i-s-m:3"
	devicePOEBEvent            = privateMibOid + "11.4.1.6.0@i-s-m:3"
	ringTopologyChangeEvent    = privateMibOid + "11.4.1.7.0@i-s-m:3"
	envMonitorEvent            = privateMibOid + "11.4.1.8.0@i-s-m:3"
	eventPortNumber            = privateMibOid + "11.4.2.1.1.1@i-m:3" // PortsEvent (11.4.2)
	eventPortEventLog          = privateMibOid + "11.4.2.1.1.2@i-s-m:3"
	eventPortEventsms          = privateMibOid + "11.4.2.1.1.3@i-s-m:3"
	eventPortEventSMTP         = privateMibOid + "11.4.2.1.1.4@i-s-m:3"
	eventPortEventsnmpTRAP     = privateMibOid + "11.4.2.1.1.5@i-s-m:3"
	eventPortEventdout1        = privateMibOid + "11.4.2.1.1.6@i-s-m:3"
	eventPortEventdout2        = privateMibOid + "11.4.2.1.1.7@i-s-m:3"
	eventPowerNumber           = privateMibOid + "11.4.3.1.1.1@i-m:3" // PowerEvent (11.4.3)
	eventPowerEventLog         = privateMibOid + "11.4.3.1.1.2@i-s-m:3"
	eventPowerEventsms         = privateMibOid + "11.4.3.1.1.3@i-s-m:3"
	eventPowerEventSMTP        = privateMibOid + "11.4.3.1.1.4@i-s-m:3"
	eventPowerEventsnmpTRAP    = privateMibOid + "11.4.3.1.1.5@i-s-m:3"
	eventPowerEventdout1       = privateMibOid + "11.4.3.1.1.6@i-s-m:3"
	eventPowerEventdout2       = privateMibOid + "11.4.3.1.1.7@i-s-m:3"
	eventDiNumber              = privateMibOid + "11.4.4.1.1.1@i-m:3" // DiEvent (11.4.4)
	eventDiEventLog            = privateMibOid + "11.4.4.1.1.2@i-s-m:3"
	eventDiEventsms            = privateMibOid + "11.4.4.1.1.3@i-s-m:3"
	eventDiEventSMTP           = privateMibOid + "11.4.4.1.1.4@i-s-m:3"
	eventDiEventsnmpTRAP       = privateMibOid + "11.4.4.1.1.5@i-s-m:3"
	eventDiEventdout1          = privateMibOid + "11.4.4.1.1.6@i-s-m:3"
	eventDiEventdout2          = privateMibOid + "11.4.4.1.1.7@i-s-m:3"

	// Monitoring (12)
	envVoltage                     = privateMibOid + "12.1.1.0@s" // ENVMonitor (12.1)
	envCurrent                     = privateMibOid + "12.1.2.0@s"
	envWalt                        = privateMibOid + "12.1.3.0@s"
	envTemperature                 = privateMibOid + "12.1.4.0@s"
	ddmPortNumber                  = privateMibOid + "12.2.1.1.1@i-m:3" // DDM (12.2)
	ddmTemperatureHighAlarm        = privateMibOid + "12.2.1.1.2@s-m:3"
	ddmTemperatureHighWarning      = privateMibOid + "12.2.1.1.3@s-m:3"
	ddmTemperatureCurrentValue     = privateMibOid + "12.2.1.1.4@s-m:3"
	ddmTemperatureLowWarning       = privateMibOid + "12.2.1.1.5@s-m:3"
	ddmTemperatureLowAlarm         = privateMibOid + "12.2.1.1.6@s-m:3"
	ddmVccHighAlarm                = privateMibOid + "12.2.1.1.7@s-m:3"
	ddmVccHighWarning              = privateMibOid + "12.2.1.1.8@s-m:3"
	ddmVccCurrentValue             = privateMibOid + "12.2.1.1.9@s-m:3"
	ddmVccLowWarning               = privateMibOid + "12.2.1.1.10@s-m:3"
	ddmVccLowAlarm                 = privateMibOid + "12.2.1.1.11@s-m:3"
	ddmBiasHighAlarm               = privateMibOid + "12.2.1.1.12@s-m:3"
	ddmBiasHighWarning             = privateMibOid + "12.2.1.1.13@s-m:3"
	ddmBiasCurrentValue            = privateMibOid + "12.2.1.1.14@s-m:3"
	ddmBiasLowWarning              = privateMibOid + "12.2.1.1.15@s-m:3"
	ddmBiasLowAlarm                = privateMibOid + "12.2.1.1.16@s-m:3"
	ddmTxPowerHighAlarm            = privateMibOid + "12.2.1.1.17@s-m:3"
	ddmTxPowerHighWarning          = privateMibOid + "12.2.1.1.18@s-m:3"
	ddmTxPowerCurrentValue         = privateMibOid + "12.2.1.1.19@s-m:3"
	ddmTxPowerLowWarning           = privateMibOid + "12.2.1.1.20@s-m:3"
	ddmTxPowerLowAlarm             = privateMibOid + "12.2.1.1.21@s-m:3"
	ddmRxPowerHighAlarm            = privateMibOid + "12.2.1.1.22@s-m:3"
	ddmRxPowerHighWarning          = privateMibOid + "12.2.1.1.23@s-m:3"
	ddmRxPowerCurrentValue         = privateMibOid + "12.2.1.1.24@s-m:3"
	ddmRxPowerLowWarning           = privateMibOid + "12.2.1.1.25@s-m:3"
	ddmRxPowerLowAlarm             = privateMibOid + "12.2.1.1.26@s-m:3"
	monitorPowerNumber             = privateMibOid + "12.3.1.1.1@i-m:3" // PowerMonitor (12.3)
	monitorPowerStatus             = privateMibOid + "12.3.1.1.2@i-m:3"
	monitorPoEPortCfgNum           = privateMibOid + "12.4.1.1.1@i-w-m:3" // POEMonitor (12.4)
	monitorPoEPortStatus           = privateMibOid + "12.4.1.1.2@s-w-m:3"
	monitorPoEPortClass            = privateMibOid + "12.4.1.1.3@s-w-m:3"
	monitorPoEPortPowerConsumption = privateMibOid + "12.4.1.1.4@s-w-m:3"
	monitorPoEPortCurrent          = privateMibOid + "12.4.1.1.5@s-w-m:3"
	monitorPoEPortVoltage          = privateMibOid + "12.4.1.1.6@s-w-m:3"
	monitorPoEPortTemperature      = privateMibOid + "12.4.1.1.7@s-w-m:3"
	cpuLoadingMonitor              = privateMibOid + "12.5.1.0@i" // CPULoadingMonitor (12.5)

	// SaveConfiguration (13)
	saveCfgMgtAction = privateMibOid + "13.1.0@s-s" // the value type to set is integer, but get will be string, set it string for now

	// FactoryDefault (14)
	factoryDefaultAction = privateMibOid + "14.1.0@s-s" // the value type to set is integer, but get will be string, set it string for now

	// SystemReboot (15)
	systemRebootAction = privateMibOid + "15.1.0@s-s" // the value type to set is integer, but get will be string, set it string for now

	// Maintenance (16)
	importConfiguration = privateMibOid + "16.1.0@s-s"
	upgrade             = privateMibOid + "16.2.0@s-s"

	// ************** Private MIB *********************
)
