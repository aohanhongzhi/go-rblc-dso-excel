package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestCD02900021(t *testing.T) {

	var faqMap = make(map[string]string)

	faqMap["Does the operator have sufficient resources and knowledge to operate the IT system securely?"] = "finished IPS042"
	faqMap[`Are assets classified according their criticality and is the criticality documented in an asset management database?`] = `Yes, an authorization management concept is implemented based on the RBAC (Role-Based Access Control) model. The system features a comprehensive backend for personnel and permissions management, allowing for efficient granting, changing, revoking, and reviewing of access rights.`
	faqMap[`Is access to information and IT systems in accordance with the authentication procedures specified?`] = `Yes, we utilize SSL encryption for information transmission, a robust password design system, and comprehensive logging to provide strong authentication for accessing IT systems and applications with business records from public networks, such as the internet.`
	faqMap[`Does strong authentication provide access from public networks (e.g. internet) to IT systems and applications with business records?`] = `Yes, we utilize SSL encryption for information transmission, a robust password design system, and comprehensive logging to provide strong authentication for accessing IT systems and applications with business records from public networks, such as the internet.`
	faqMap[`Is authentication information protected during storage and transmission?`] = `Yes, authentication information is protected during transmission using SSL. Sensitive information is encrypted with RSA during transmission and is also stored using AES encryption.`
	faqMap[`Are cloud customers' cloud ressources and data strictly separated from other cloud customers in the cloud solution?`] = `Yes, we only use Azure cloud services.`
	faqMap[`Is the data stored and/or transmitted in encrypted form?`] = `Yes, authentication information is protected during transmission using SSL. Sensitive information is encrypted with RSA during transmission and is also stored using AES encryption.`
	faqMap[`Do the cryptographic algorithms and key lengths of the IT system comply with the C/IDS regulations?`] = `Yes, authentication information is protected during transmission using SSL. Sensitive information is encrypted with RSA during transmission and is also stored using AES encryption.`
	faqMap[`Is the functionality of the IT system limited to the legitimate purpose and are all built-in security functionalities activated?`] = `Yes, we have conducted a review with our legal colleagues, and the system's security settings, password strength, and firewall are all enabled.`
	faqMap[`Does a hardening concept exist for the webserver and are the corresponding requirements fulfilled?`] = `Yes, our server has undergone penetration testing, and all vulnerabilities have been addressed. The server is configured with a firewall, and we have minimized port exposure.`
	faqMap[`Does a current operating manual exist where security-relevant operating processes (System description, interfaces, responsibilities,  installed software, system update - patch management, authorization management - e.g. reference to AIM documentation, backup and recovery procedures, emergency procedures) are documented?`] = `Yes, we have standard operating documentation that covers these security-relevant processes.`
	faqMap[`Are development and test environments separated from productive environments?`] = `Yes, we have separate servers and domains for development and testing environments. The databases and Redis are also kept separate.`
	faqMap[`Do the data backup and recovery concepts guarantee the required recovery times and has this been tested?`] = `Yes, we have daily database backups, with the backed-up data stored both locally and in the cloud.`
	faqMap[`Are security-relevant events transmitted to a (central) server and logged?`] = `Yes, our system logs events and retains them for over six months in accordance with China's cybersecurity law.`
	faqMap[`Are access, write and change events logged on the IT systems?`] = `Yes, our system logs events and retains them for over six months in accordance with China's cybersecurity law.`
	faqMap[`Are the monitoring functions of the cloud solution used to detect security related events?`] = `Yes, we use monitoring functions to detect security-related events, such as identifying illegal IP addresses and adding them to the firewall for handling.`
	faqMap[`Are security patches/fixes and system updates applied according to Bosch regulations and deadlines? `] = `Yes, we apply security patches and system updates promptly according to emails from BD and DSO colleagues to address vulnerabilities and manage system patches.`
	faqMap[`Does the installed software originate from a trustworthy source?`] = `Yes, we install software through Bosch's software center or directly from official websites. Additionally, we do not install pirated software or software that does not comply with licensing agreements.`
	faqMap[`Do you know how and when to contact Bosch CERT and how to react to CERT advisories?`] = `Yes, I know how and when to contact Bosch CERT. The email for CERT is CERT@bosch.com.`
	faqMap[`Have the type, scope and repeat intervals of technical security checks (e.g. scans, penetration tests) been defined for the IT system and documented in the operating manual?`] = `Yes, our system undergoes penetration testing before going live, and DSO also regularly invites external professional teams to conduct security testing.`
	faqMap[`Is the IT system regularly checked by a security scan?`] = `Yes, our system undergoes penetration testing before going live, and DSO also regularly invites external professional teams to conduct security testing.`
	faqMap[`Is the cloud customer aware about its security responsibilities when using the cloud service (e.g., for secure configuration and operation of the cloud solution)?`] = `Yes, the cloud customer is made aware of its security responsibilities through prior security training, specifically ISP042. This training ensures that customers understand the importance of secure configuration and operation of the cloud solution. Additionally, all systems undergo penetration testing and acceptance procedures to further enhance security awareness and assurance.`
	faqMap[`Does the cloud customer know how the data can be removed/deleted from the cloud solution after the cloud service has ended?`] = `Yes, the cloud customer is informed about data removal procedures after the service has ended. Our system includes a user manual that outlines the methods for deleting data. Additionally, the user agreement specifies these procedures and provides contact information for the DSO department, including a phone number and email for any inquiries related to data deletion.`
	faqMap[`Is an authorization management concept implemented or are appropriate procedures provided for granting, changing, revoking and reviewing access rights?`] = `Yes, an authorization management concept is implemented based on the RBAC (Role-Based Access Control) model. The system features a comprehensive backend for personnel and permissions management, allowing for efficient granting, changing, revoking, and reviewing of access rights.`
	faqMap[`Is Bosch Single Sign-On used for the authentication of Bosch employees and employees of external service providers?`] = `No, our system is designed for general consumer users and does not involve Bosch internal users. Therefore, it does not integrate with Bosch Single Sign-On (SSO) for authentication of Bosch employees or employees of external service providers.`
	faqMap[`Does a password management concept exist and is it implemented?`] = `Yes, a password management concept exists and is implemented according to Bosch internal requirements. It mandates the use of uppercase and lowercase letters, numbers, and special characters. Additionally, passwords expire every three months, and the use of previous passwords is prohibited.`
	faqMap[`Have passwords for build in and standard accounts been changed and are they managed according to the regulations?`] = `Yes, a password management concept exists and is implemented according to Bosch internal requirements. It mandates the use of uppercase and lowercase letters, numbers, and special characters. Additionally, passwords expire every three months, and the use of previous passwords is prohibited.`

	OperatingManualV2, err := ListDir("/home/eric/Project/Go/go-rblc-dso-excel/2024-CD-02900-021", "xlsm")
	if err != nil {
		panic(err)
	}

	for _, op := range OperatingManualV2 {
		index := strings.Index(op, `.~`)
		if index > 0 {
			continue
		}
		println(op)
		f, err := excelize.OpenFile(op)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		sheetList := f.GetSheetList()
		for _, sheet := range sheetList {
			println(sheet)
		}
		sheet := "Checklist"
		// 遍历所有问题
		rows, err := f.GetRows(sheet)
		if err != nil {
			println(err)
			return
		}
		for index, row := range rows {
			if index < 10 {
				continue
			}
			fmt.Printf("%+v", row)
			number := strconv.Itoa(index)
			cell, err := f.GetCellValue(sheet, "F"+number)
			if err != nil {
				fmt.Println(err)
				continue
			}

			value, ok := faqMap[cell]
			if ok {
				f.SetCellStr(sheet, "I"+number, value)
			} else {
				fmt.Printf("\n%v.这个问题还没有答案【%v】 %v\n", index, cell, row)
			}
		}

		if false {
			// 获取工作表中指定单元格的值

			cell, err := f.GetCellValue(sheet, "F14")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(cell)

			f.SetCellStr(sheet, "I14", "finished IPS042")
			f.SetCellStr(sheet, "I15", "")
			f.SetCellStr(sheet, "I18", "Yes, the cloud customer is made aware of its security responsibilities through prior security training, specifically ISP042. This training ensures that customers understand the importance of secure configuration and operation of the cloud solution. Additionally, all systems undergo penetration testing and acceptance procedures to further enhance security awareness and assurance.")
			f.SetCellStr(sheet, "I34", "Yes, the cloud customer is informed about data removal procedures after the service has ended. Our system includes a user manual that outlines the methods for deleting data. Additionally, the user agreement specifies these procedures and provides contact information for the DSO department, including a phone number and email for any inquiries related to data deletion.")
			f.SetCellStr(sheet, "I40", "Yes, an authorization management concept is implemented based on the RBAC (Role-Based Access Control) model. The system features a comprehensive backend for personnel and permissions management, allowing for efficient granting, changing, revoking, and reviewing of access rights.")
			f.SetCellStr(sheet, "I52", "No, our system is designed for general consumer users and does not involve Bosch internal users. Therefore, it does not integrate with Bosch Single Sign-On (SSO) for authentication of Bosch employees or employees of external service providers.")
			f.SetCellStr(sheet, "I73", "Yes, a password management concept exists and is implemented according to Bosch internal requirements. It mandates the use of uppercase and lowercase letters, numbers, and special characters. Additionally, passwords expire every three months, and the use of previous passwords is prohibited.")
			f.SetCellStr(sheet, "I82", "Yes, a password management concept exists and is implemented according to Bosch internal requirements. It mandates the use of uppercase and lowercase letters, numbers, and special characters. Additionally, passwords expire every three months, and the use of previous passwords is prohibited.")
			f.SetCellStr(sheet, "I91", `Yes, we have established a comprehensive authorization and authentication system to ensure that data access aligns with the correct business logic.`)
			f.SetCellStr(sheet, "I95", `Yes, we utilize SSL encryption for information transmission, a robust password design system, and comprehensive logging to provide strong authentication for accessing IT systems and applications with business records from public networks, such as the internet.`)
			f.SetCellStr(sheet, "I99", `Yes, we utilize SSL encryption for information transmission, a robust password design system, and comprehensive logging to provide strong authentication for accessing IT systems and applications with business records from public networks, such as the internet.`)
			f.SetCellStr(sheet, "I101", `Yes, authentication information is protected during transmission using SSL. Sensitive information is encrypted with RSA during transmission and is also stored using AES encryption.`)
			f.SetCellStr(sheet, "I119", `Yes, we only use Azure cloud services.`)
			f.SetCellStr(sheet, "I124", `Yes, authentication information is protected during transmission using SSL. Sensitive information is encrypted with RSA during transmission and is also stored using AES encryption.`)
			f.SetCellStr(sheet, "I127", `Yes, authentication information is protected during transmission using SSL. Sensitive information is encrypted with RSA during transmission and is also stored using AES encryption.`)
			f.SetCellStr(sheet, "I153", `Yes, we have conducted a review with our legal colleagues, and the system's security settings, password strength, and firewall are all enabled.`)
			f.SetCellStr(sheet, "I154", `Yes, our server has undergone penetration testing, and all vulnerabilities have been addressed. The server is configured with a firewall, and we have minimized port exposure.`)
			f.SetCellStr(sheet, "I163", `Yes, we have standard operating documentation that covers these security-relevant processes.`)
			f.SetCellStr(sheet, "I173", `Yes, we have separate servers and domains for development and testing environments. The databases and Redis are also kept separate.`)
			f.SetCellStr(sheet, "I180", `Yes, we have daily database backups, with the backed-up data stored both locally and in the cloud.`)
			f.SetCellStr(sheet, "I184", `Yes, our system logs events and retains them for over six months in accordance with China's cybersecurity law.`)
			f.SetCellStr(sheet, "I194", `Yes, our system logs events and retains them for over six months in accordance with China's cybersecurity law.`)
			f.SetCellStr(sheet, "I196", `Yes, we use monitoring functions to detect security-related events, such as identifying illegal IP addresses and adding them to the firewall for handling.`)
			f.SetCellStr(sheet, "I200", `Yes, we apply security patches and system updates promptly according to emails from BD and DSO colleagues to address vulnerabilities and manage system patches.`)
			f.SetCellStr(sheet, "I207", `Yes, we install software through Bosch's software center or directly from official websites. Additionally, we do not install pirated software or software that does not comply with licensing agreements.`)
			f.SetCellStr(sheet, "I210", `Yes, I know how and when to contact Bosch CERT. The email for CERT is CERT@bosch.com.`)
			f.SetCellStr(sheet, "I239", `Yes, our system undergoes penetration testing before going live, and DSO also regularly invites external professional teams to conduct security testing.`)
			f.SetCellStr(sheet, "I241", `Yes, our system undergoes penetration testing before going live, and DSO also regularly invites external professional teams to conduct security testing.`)
		}

		f.Save()

	}
}
