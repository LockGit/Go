/**
 * Created by Vim
 * User: lock
 * Date: 2018/12/3
 * Time: 14:59
 */
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func startPrintln(content string) {
	fmt.Println("\033[1;31m" + content + "\033[0m")
}

func hightRecord(content string) {
	fmt.Println("\033[1;32m" + content + "\033[0m")
}

type DnsInfo struct {
	ARecord   []string
	CName     string
	NsRecord  []string
	MxRecord  []string
	TxtRecord []string
	SrvRecord []string
}

func GetARecordInfo(domain string, dnsInfo *DnsInfo, wg *sync.WaitGroup) {
	ipRecords, _ := net.LookupIP(domain)
	for _, ip := range ipRecords {
		dnsInfo.ARecord = append(dnsInfo.ARecord, ip.String())
	}
	wg.Done()
}

func GetCnameInfo(domain string, dnsInfo *DnsInfo, wg *sync.WaitGroup) {
	cname, _ := net.LookupCNAME(domain)
	dnsInfo.CName = cname
	wg.Done()
}

func GetNsInfo(domain string, dnsInfo *DnsInfo, wg *sync.WaitGroup) {
	nameServer, _ := net.LookupNS(domain)
	for _, ns := range nameServer {
		dnsInfo.NsRecord = append(dnsInfo.NsRecord, ns.Host)
	}
	wg.Done()
}

func GetMxInfo(domain string, dnsInfo *DnsInfo, wg *sync.WaitGroup) {
	mxRecords, _ := net.LookupMX(domain)
	for _, mx := range mxRecords {
		mxInfo := mx.Host + fmt.Sprintf("%v", mx.Pref)
		dnsInfo.MxRecord = append(dnsInfo.MxRecord, mxInfo)
	}
	wg.Done()
}

func GetTxtInfo(domain string, dnsInfo *DnsInfo, wg *sync.WaitGroup) {
	txtRecords, _ := net.LookupTXT(domain)
	for _, txt := range txtRecords {
		dnsInfo.TxtRecord = append(dnsInfo.TxtRecord, txt)
	}
	wg.Done()
}

func GetSrvInfo(domain string, dnsInfo *DnsInfo, wg *sync.WaitGroup) {
	_, sRvs, err := net.LookupSRV("xmpp-server", "tcp", domain)
	if err == nil {
		for _, srv := range sRvs {
			s := fmt.Sprintf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
			srvInfo := strings.TrimSuffix(s, "\n")
			dnsInfo.SrvRecord = append(dnsInfo.SrvRecord, srvInfo)
		}
	}
	wg.Done()
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("please input the domain !")
		fmt.Println("example: fdns baidu.com")
		os.Exit(0)
	}

	var wg sync.WaitGroup
	dnsInfo := new(DnsInfo)
	domain := args[0]
	go GetARecordInfo(domain, dnsInfo, &wg)
	wg.Add(1)
	go GetCnameInfo(domain, dnsInfo, &wg)
	wg.Add(1)
	go GetNsInfo(domain, dnsInfo, &wg)
	wg.Add(1)
	go GetMxInfo(domain, dnsInfo, &wg)
	wg.Add(1)
	go GetTxtInfo(domain, dnsInfo, &wg)
	wg.Add(1)
	go GetSrvInfo(domain, dnsInfo, &wg)
	wg.Add(1)

	wg.Wait()
	startPrintln("Find A OR AAAA Record:")
	for _, ip := range dnsInfo.ARecord {
		hightRecord(ip)
	}
	startPrintln("----------")
	startPrintln("Cname:")
	hightRecord(dnsInfo.CName)
	startPrintln("----------")
	startPrintln("Ns Record:")
	for _, ns := range dnsInfo.NsRecord {
		hightRecord(ns)
	}
	startPrintln("----------")
	startPrintln("Mx Record:")
	for _, mx := range dnsInfo.MxRecord {
		hightRecord(mx)
	}
	startPrintln("----------")
	startPrintln("Txt Record:")
	for _, txt := range dnsInfo.TxtRecord {
		hightRecord(txt)
	}
	startPrintln("----------")
	startPrintln("Srv Record:")
	for _, srv := range dnsInfo.SrvRecord {
		hightRecord(srv)
	}
}
