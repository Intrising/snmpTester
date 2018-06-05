package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"encoding/hex"
	"encoding/binary"
	"time"
)

func ShellExec(format string, a ...interface{}) (error, string) {
	//fmt.Println("ShellExec : " , format ,  a)
	var outb, errb bytes.Buffer
	cmd := exec.Command("sh", "-c", fmt.Sprintf(format, a...))
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		err = fmt.Errorf(errb.String())
	}
	return err, outb.String()
}

func HasString(hay []string, needle string) bool {
	for _, s := range hay {
		if s == needle {
			return true
		}
	}
	return false
}

func ElementStringHasInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func ElementIntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func MatchPorts(a1 []int, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}

	sort.Ints(a1)
	sort.Ints(a2)

	for n, v := range a1 {
		if a2[n] != v {
			return false
		}
	}

	return true
}

func ORPorts(a1 []int, a2 []int) []int {
	tmp := make([]int, len(a1)+len(a2))
	copy(tmp, a1)
	index := len(a1)
	for _, v := range a2 {
		if ElementIntInSlice(v, a1) == false {
			tmp[index] = v
			index++
		}
	}
	result := make([]int, index)
	copy(result, tmp)
	sort.Ints(result)
	return result
}

func XORPorts(a1 []int, a2 []int) []int {
	tmp := make([]int, len(a1))
	index := 0
	for _, v := range a1 {
		if ElementIntInSlice(v, a2) == false {
			tmp[index] = v
			index++
		}
	}
	result := make([]int, index)
	copy(result, tmp)
	sort.Ints(result)
	return result
}

func ConvertArrayIntToString(intArray []int) []string {
	stringArray := make([]string, len(intArray))
	for n, v := range intArray {
		stringArray[n] = strconv.Itoa(v)
	}
	return stringArray
}

func PaddingMacAddr(macAddr string) string {
	//un used
	array := strings.Split(macAddr, ":")

	for n, v := range array {
		if len(v) == 1 {
			array[n] = "0" + v
		}
	}

	return strings.Join(array, ":")
}

func ConvertMacAddrStringToByteArray(mac string)([]byte){
	array , _ := hex.DecodeString(strings.Join(strings.Split(mac,":"),""))
	return array
}

func ConvertByteArrayToMacAddrString(array []byte)string{
	t := make([]string,0)
	for _,v := range array{
		t = append(t,hex.EncodeToString([]byte{v}))
	}
	return strings.Join(t,":")
}

func ConvertIpAddrStringByteArray(ip string)([]byte){
	buf := make([]byte,0,4)
	array := strings.Split(ip,".")
	for _,v:=range array{
		input, _ := strconv.Atoi(v)
		buf = append(buf,(byte)(input))
	}
	return buf
}

func ConvertByteArrayToIpAddrString(array []byte)string{
	t := make([]string,0)
	for _,v := range array{
		t = append(t,strconv.Itoa((int)(v)))
	}
	return strings.Join(t,".")
}

func VerifyCheckSum(verifyLength int, verifyValue []byte, isCDP bool)bool{
	//defer utils.MyDefer("")
	sum := (uint32)(0)
	if verifyLength%2 == 1{
		if isCDP && ((int)(verifyValue[verifyLength-1])>= 0x80){
			sum += 1
			sum += (uint32)(verifyValue[verifyLength-1]<<8)
		}else if isCDP{
			sum += (uint32)(verifyValue[verifyLength-1]<<8)
		}else{
			sum += (uint32)(verifyValue[verifyLength-1])
		}
		verifyLength = (verifyLength-1)/2
	}else{
		verifyLength = verifyLength/2
	}
	
	for verifyLength > 0{
		if isCDP{
			sum += (uint32)(binary.LittleEndian.Uint16(verifyValue[(verifyLength-1)*2:(verifyLength-1)*2+2]))
		}else{
			sum += (uint32)(binary.BigEndian.Uint16(verifyValue[(verifyLength-1)*2:(verifyLength-1)*2+2]))
		}
		verifyLength--
	}
	
	sum = ((sum >> 16) & 0xffff) + sum & 0xffff
	return (sum == 65535)
}

func GenCheckSum(verifyValue []byte , protocolType string)[]byte{
	//defer utils.MyDefer("")
	verifyLength:= len(verifyValue)
	sum := (uint32)(0)
	if verifyLength%2 == 1{
		if protocolType == "udp"{
			sum += 0xff<<8
		}else if (protocolType == "cdp") && ((int)(verifyValue[verifyLength-1])>= 0x80){
			sum -= 1<<8
			sum += (uint32)(verifyValue[verifyLength-1])
		}else{
			sum += (uint32)(verifyValue[verifyLength-1])
		}
		verifyLength = (verifyLength-1)/2
	}else{
		verifyLength = verifyLength/2
	}
	
	for verifyLength > 0{ 
		sum += (uint32)(binary.BigEndian.Uint16(verifyValue[(verifyLength-1)*2:(verifyLength-1)*2+2]))
		verifyLength--
	}
	
	sum =  ((sum >> 16) & 0xffff) + sum & 0xffff
	sum = 65535-sum
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, (uint16)(sum))
	
	return buf
}

func GenUdpCheckSum(ip []byte,serverip []byte,udpLength int,udpBuf []byte,bootpBuf []byte)[]byte{
	buf := make([]byte,udpLength + 12)
	buf = append(buf , []byte{0x00,0x11} ...)
	buf = append(buf , ip...)
	buf = append(buf , serverip...)
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, uint16(udpLength))
	buf = append(buf , bs...)
	buf = append(buf , udpBuf...)
	buf = append(buf , bootpBuf...)
	return GenCheckSum(buf,"udp")
}

func ConvertDotStringToInts(input string)[]int {
	//fmt.Println("ConvertDotStringToInts : input = " , len(input) , input )
	lists := strings.Split(input, ",")
	tmp := make([]int,len(lists))
	for _,v:=range lists {
		value, ok := strconv.Atoi(v)
		if ok == nil {
			tmp = append(tmp , value)
		}
	}
	return tmp
}

func ClearTimer(timer interface{}){
	//fmt.Println("ClearTimer timer = " , timer)
	switch f := timer.(type) {
		case *time.Timer:
			if timer != (*time.Timer)(nil){
				f.Stop()
			}
		case *time.Ticker:
			if timer != (*time.Ticker)(nil){
				f.Stop()
			}
		default:
			fmt.Println(" clearTimeTicker default")
	}
}

func MyDefer(info string){
	if err:=recover();err!=nil{
		fmt.Println(info + " MyDefer : ",err)
	}
}

func Diff(src , dst []int)[]int{
	//fmt.Println("Diff :  src = " , src , " dst = " ,dst)
	if len(src) == 0 {
		return dst
	}
	
	rtn := make([]int,0,len(src))
	for _,v:=range src {
		if ElementIntInSlice(v,dst) == false{
			rtn = append(rtn , v)
		}
	}
	
	return rtn 
}

func GetIndexInSlice(value int,array []int)int{
	for n,v:=range array {
		if v == value {
			return n
		}
	}
	
	return -1
}

func PrintMsg(msg ...interface{}){
	fmt.Println("PrintMsg : ", msg)
}