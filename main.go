package main

import (
	"fmt"
	"strings"

	"./utils"
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
	ospfRouterID                 = mib2Prefix + "14.1.1@s-s"
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
	ospfAreaID                              = mib2Prefix + "14.2.1.1@s"
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
	ospfStubAreaID     = mib2Prefix + "14.3.1.1@s"
	ospfStubTOS        = mib2Prefix + "14.3.1.2@i"
	ospfStubMetric     = mib2Prefix + "14.3.1.3@i"
	ospfStubStatus     = mib2Prefix + "14.3.1.4@i-s"
	ospfStubMetricType = mib2Prefix + "14.3.1.5@i-s"

	// ospfLsdbTable (14.4)
	// ospfLsdbEntry (not-accessible)(14.4.1)
	ospfLsdbAreaID        = mib2Prefix + "14.4.1.1@s"
	ospfLsdbType          = mib2Prefix + "14.4.1.2@i"
	ospfLsdbLsid          = mib2Prefix + "14.4.1.3@s"
	ospfLsdbRouterID      = mib2Prefix + "14.4.1.4@s"
	ospfLsdbSequence      = mib2Prefix + "14.4.1.5@i"
	ospfLsdbAge           = mib2Prefix + "14.4.1.6@i"
	ospfLsdbChecksum      = mib2Prefix + "14.4.1.7@i"
	ospfLsdbAdvertisement = mib2Prefix + "14.4.1.8@s"

	// ospfAreaRangeTable (14.5)
	// ospfAreaRangeEntry (not-accessible)(14.5.1)
	ospfAreaRangeAreaID = mib2Prefix + "14.5.1.1@s"
	ospfAreaRangeNet    = mib2Prefix + "14.5.1.2@s"
	ospfAreaRangeMask   = mib2Prefix + "14.5.1.3@s"
	ospfAreaRangeStatus = mib2Prefix + "14.5.1.4@i-s"
	ospfAreaRangeEffect = mib2Prefix + "14.5.1.6@i-s"

	// ospfHostTable (14.6)
	// ospfHostEntry (not-accessible)(14.6.1)
	ospfHostIPAddress = mib2Prefix + "14.6.1.1@s"
	ospfHostTOS       = mib2Prefix + "14.6.1.2@i"
	ospfHostMetric    = mib2Prefix + "14.6.1.3@i"
	ospfHostStatus    = mib2Prefix + "14.6.1.4@i-s"
	ospfHostAreaID    = mib2Prefix + "14.6.1.5@s"
	ospfHostCfgAreaID = mib2Prefix + "14.6.1.6@s-s"

	// ospfIfTable (14.7)
	// ospfIfEntry (not-accessible)(14.7.1)
	ospfIfIPAddress                = mib2Prefix + "14.7.1.1@s"
	ospfAddressLessIf              = mib2Prefix + "14.7.1.2@i"
	ospfIfAreaID                   = mib2Prefix + "14.7.1.3@s-s"
	ospfIfType                     = mib2Prefix + "14.7.1.4@i-s"
	ospfIfAdminStat                = mib2Prefix + "14.7.1.5@i-s"
	ospfIfRtrPriority              = mib2Prefix + "14.7.1.6@i-s"
	ospfIfTransitDelay             = mib2Prefix + "14.7.1.7@i-s"
	ospfIfRetransInterval          = mib2Prefix + "14.7.1.8@i-s"
	ospfIfHelloInterval            = mib2Prefix + "14.7.1.9@i-s"
	ospfIfRtrDeadInterval          = mib2Prefix + "14.7.1.10@i-s"
	ospfIfPollInterval             = mib2Prefix + "14.7.1.11@i-s"
	ospfIfState                    = mib2Prefix + "14.7.1.12@i"
	ospfIfDesignatedRouter         = mib2Prefix + "14.7.1.13@s"
	ospfIfBackupDesignatedRouter   = mib2Prefix + "14.7.1.14@s"
	ospfIfEvents                   = mib2Prefix + "14.7.1.15@i"
	ospfIfAuthKey                  = mib2Prefix + "14.7.1.16@s-s"
	ospfIfStatus                   = mib2Prefix + "14.7.1.17@i-s"
	ospfIfMulticastForwarding      = mib2Prefix + "14.7.1.18@i-s"
	ospfIfDemand                   = mib2Prefix + "14.7.1.19@i-s"
	ospfIfAuthType                 = mib2Prefix + "14.7.1.20@i-s"
	ospfIfLsaCount                 = mib2Prefix + "14.7.1.21@i"
	ospfIfLsaCksumSum              = mib2Prefix + "14.7.1.22@i"
	ospfIfDesignatedRouterID       = mib2Prefix + "14.7.1.23@s"
	ospfIfBackupDesignatedRouterID = mib2Prefix + "14.7.1.24@s"

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

	// RFC 4318 starts
	// taskEntry = append(taskEntry, genTask("dot1dStpVersion", dot1dStpVersion))
	// taskEntry = append(taskEntry, genTask("dot1dStpTxHoldCount", dot1dStpTxHoldCount))
	// taskEntry = append(taskEntry, genTask("dot1dStpPortProtocolMigration", dot1dStpPortProtocolMigration))
	// taskEntry = append(taskEntry, genTask("dot1dStpPortAdminEdgePort", dot1dStpPortAdminEdgePort))
	// taskEntry = append(taskEntry, genTask("dot1dStpPortOperEdgePort", dot1dStpPortOperEdgePort))
	// taskEntry = append(taskEntry, genTask("dot1dStpPortAdminPointToPoint", dot1dStpPortAdminPointToPoint))
	// taskEntry = append(taskEntry, genTask("dot1dStpPortOperPointToPoint", dot1dStpPortOperPointToPoint))
	taskEntry = append(taskEntry, genTask("dot1dStpPortAdminPathCost", dot1dStpPortAdminPathCost))

	// // no definition yet
	// taskEntry = append(taskEntry, genTask("rstpNotifications", rstpNotifications))
	// taskEntry = append(taskEntry, genTask("rstpObjects", rstpObjects))
	// // just groups
	// taskEntry = append(taskEntry, genTask("rstpBridgeGroup", rstpBridgeGroup))
	// taskEntry = append(taskEntry, genTask("rstpPortGroup", rstpPortGroup))
	// taskEntry = append(taskEntry, genTask("rstpCompliance", rstpCompliance))

	// RFC 4318 Ends

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
