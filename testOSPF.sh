#!/bin/bash
deviceIP="192.168.16.165"
mibName="../private-mib-generator/output/Switch2/Lantech/nec/cat3/IPGS-6416XSFP-Layer3-v1.0.16.mib"
noAccessResult="Reason: noAccess"
noSuchObjectResult="No Such Object available on this agent at this OID"
noCreationResult="Reason: noCreation"
badValueReuslt="Reason: (badValue)"

setV3() {
  snmpset -v 3 -u $1 -l authPriv -a MD5 -A @Walter123 -x DES -X @Walter123 -m ${mibName} ${deviceIP} $2 $3 $4 > setResult.txt 2>&1  
  echo `cat setResult.txt`
}

getV3() {
  snmpget -v 3 -u $1 -l authPriv -a MD5 -A @Walter123 -x DES -X @Walter123 -m ${mibName} ${deviceIP} $2 > getResult.txt 2>&1
  echo `cat getResult.txt`
}

setV2c() {
  snmpset -v 2c -c private -m ${mibName} ${deviceIP} $1 $2 $3 > setResult.txt 2>&1  
  echo `cat setResult.txt`
}

getV2c() {
  snmpget -v 2c -c private -m ${mibName} ${deviceIP} $1 > getResult.txt 2>&1
  echo `cat getResult.txt`
}

walkV2c() {
  snmpwalk -v 2c -c private -m ${mibName} ${deviceIP} $1 > getResult.txt 2>&1
  echo `cat getResult.txt`
}

test_toggle_l3OSPFEnabled() {
  result=$(setV2c "l3OSPFEnable.0" i 1)
  assertContains ["Result should contain true(1)"] "$result" "true(1)"
  
  result=$(setV2c "l3OSPFEnable.0" i 2)
  assertContains ["Result should contain false(2)"] "$result" "false(2)"
  
  result=$(setV2c "l3OSPFEnable.0" i 1)
  assertContains ["Result should contain enabled"] "$result" "true(1)"
}

test_get_l3OSPFAddRouterTable_will_get_description() {
  result=$(getV2c "l3OSPFAddRouterTable.0")
  assertContains ["Result should contain This oid is used for adding l3OSPFRouterTable"] "$result" "This oid is used for adding l3OSPFRouterTable"
}

test_get_l3OSPFDeleteRouterTable_will_get_description() {
  result=$(getV2c "l3OSPFDeleteRouterTable.0")
  assertContains ["Result should contain This oid is used for deleting l3OSPFRouterTable"] "$result" "This oid is used for deleting l3OSPFRouterTable"
}

test_set_l3OSPFAddRouterTable() {
  result=$(setV2c "l3OSPFAddRouterTable.0" "s" "1.1.1.1/20/vlan-6")
  assertContains ["Result should contain 1.1.1.1/20/vlan-6"] "$result" "1.1.1.1/20/vlan-6"
  
  result=$(setV2c "l3OSPFAddRouterTable.0" "s" "1.1.1.2/20/vlan-6")
  assertContains ["Adding routerTable should get no-creationt "] "$result" "$noCreationResult"
}

test_get_l3OSPFRouterID() {
  result=$(getV2c "l3OSPFRouterID.1")
  assertContains ["Result should contain 1.1.1.1"] "$result" "1.1.1.1"
}

test_get_l3OSPFRouterAutoCostBandwith() {
  result=$(getV2c "l3OSPFAutoCostBandwidth.1")
  assertContains ["Result should contain 20"] "$result" "20"
}

test_get_l3OSPFRouterNetworks() {
  result=$(getV2c "l3OSPFNetworks.1")
  assertContains ["Result should contain vlan-6"] "$result" "vlan-6"
}

test_set_l3OSPFDeleteRouterTable_success() {
  result=$(setV2c "l3OSPFDeleteRouterTable.0" s 1.1.1.1)
  assertContains ["Result should contain 1.1.1.1"] "$result" "1.1.1.1"
}


test_set_l3OSPFAddInterfaceTable() {
  result=$(setV2c "l3OSPFAddInterfaceTable.0" "s" "vlan-20/0/1/0/10/40/2/5/0/key/keyID")
  assertContains ["Test setting non-exist network interface should get noCreation error"] "$result" "$noCreationResult"
  
  result=$(setV2c "l3OSPFAddInterfaceTable.0" "s" "1120/0/1/0/10/40/2/5/0/key/keyID")
  assertContains ["Test setting invalid interface format should get bad-value error"] "$result" "$badValueReuslt"

  result=$(setV2c "l3OSPFAddInterfaceTable.0" "s" "vlan-1/0/1/0/10/40/2/5/0/key/keyID")
  assertContains ["Result should contain vlan-1/0/1/0/10/40/2/5/0/key/keyID"] "$result" "vlan-1/0/1/0/10/40/2/5/0/key/keyID"
}

test_get_l3OSPFInterfaceTable() {
  result=$(getV2c "l3OSPFInterfaceNetworkInterface.1" )
  assertContains ["Result should contain vlan-1"] "$result" "vlan-1"

  result=$(getV2c "l3OSPFInterfaceArea.1" )
  assertContains ["Result should contain string 0"] "$result" "l3OSPFInterfaceArea.1 = STRING: 0"
  
  result=$(getV2c "l3OSPFInterfacePriority.1" )
  assertContains ["Result should contain integer 0"] "$result" "l3OSPFInterfacePriority.1 = INTEGER: 1"

  result=$(getV2c "l3OSPFInterfaceCost.1" )
  assertContains ["Result should contain integer 0"] "$result" "l3OSPFInterfaceCost.1 = INTEGER: 0"

  result=$(getV2c "l3OSPFInterfaceHelloInterval.1" )
  assertContains ["Result should contain integer 10"] "$result" "l3OSPFInterfaceHelloInterval.1 = INTEGER: 10"
  
  result=$(getV2c "l3OSPFInterfaceDeadInterval.1" )
  assertContains ["Result should contain integer 40"] "$result" "l3OSPFInterfaceDeadInterval.1 = INTEGER: 40"
  
  result=$(getV2c "l3OSPFInterfaceTransmitDelay.1" )
  assertContains ["Result should contain integer 2"] "$result" "l3OSPFInterfaceTransmitDelay.1 = INTEGER: 2"
 
  result=$(getV2c "l3OSPFInterfaceRetransmitInterval.1" )
  assertContains ["Result should contain integer 0"] "$result" "l3OSPFInterfaceRetransmitInterval.1 = INTEGER: 0"
  
  result=$(getV2c "l3OSPFInterfaceAuthenticationMode.1" )
  assertContains ["Result should contain none(0)"] "$result" "l3OSPFInterfaceAuthenticationMode.1 = INTEGER: none(0)"
  
  result=$(getV2c "l3OSPFInterfaceAuthenticationKey.1" )
  assertContains ["Result should contain key"] "$result" "l3OSPFInterfaceAuthenticationKey.1 = STRING: key"
  
  result=$(getV2c "l3OSPFInterfaceAuthenticationKeyID.1" )
  assertContains ["Result should contain keyID"] "$result" "l3OSPFInterfaceAuthenticationKeyID.1 = STRING: keyID"
}


test_set_l3OSPFDeleteInterfaceTable() {
  result=$(setV2c "l3OSPFDeleteInterfaceTable.0" "s" "vlan-1")
  assertContains ["Result should contain vlan-1"] "$result" "vlan-1"
}



# use following to execute all test fucntions
. shunit2