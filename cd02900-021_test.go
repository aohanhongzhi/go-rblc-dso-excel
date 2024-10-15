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
	faqMap[`Has the operator disclosed for which security classes the IT system can be used?`] = `Certainly, all our systems are involved in DSPiE, and they have been assessed for security levels.`
	faqMap[`Is there a central asset management database where the IT systems and the respective operators are documented?`] = `bosch leanix`
	faqMap[`Does the authorization management concept include the handling of privileged or administrative accounts?`] = `We have created a document for access control management.`
	faqMap[`Are the C/IDS requirements for secure authentication fulfilled?`] = `Yes, our Azure server management system is integrated with Bosch's Single Sign-On (SSO). Additionally, all servers are configured for key-based authentication to ensure security`
	faqMap[`Are the used keys (e.g., pre-shared keys) and certificates changed on a regular basis?`] = `We change our keys every six months.`
	faqMap[`Is there a change management process for managing changes to the IT system?`] = `Yes, we manage changes through email communication or handover processes`
	faqMap[`Is the cloud customer aware of the critical operation procedures (e.g. deletion of virtual servers) and are these functions documented (e.g. in an operational concept)?`] = `Yes, we have documented the purpose of these applications using notes in the Azure management portal.`
	faqMap[`Does the solution for authentication the Bosch customer's context fulfill the requirements of C/IDS?`] = `Yes, we have designed the password system strictly according to Bosch's requirements, which includes a combination of over 12 characters with uppercase and lowercase letters, numbers, and special characters. We use a high-strength encryption algorithm, AES-256.`
	faqMap[`Is the interactive access to the IT system disabled or terminated in case of user absence?`] = `Yes, we disable login and access for users who have left or are no longer using the system.`
	faqMap[`Are there adequate data backup and recovery concepts for the IT system?`] = `We have daily Microsoft backups, and in addition, we use automated scripts to back up all data, which is then stored on Bosch cloud storage.`
	faqMap[`Have the security requirements been agreed with the external service providers?`] = `NDA`
	faqMap[`Are security requirements defined in external supplier agreements (e.g. contracts) and agreed by the supplier?`] = `NDA`
	faqMap[`Was it checked whether a penetration test is necessary and, if so, was it carried out?`] = `NDA`
	faqMap[`Was it checked whether a penetration test is necessary and, if so, was it carried out?`] = `The PO will send an email to the colleague responsible for security in the department to conduct a relevant assessment. If a penetration test is deemed necessary, it will be carried out.`
	faqMap[`Is the protection against malware (virus protection) according to C/IDS specifications ensured or is the system isolated via network separation measures (e.g. ITM specifications)?`] = `Yes, we have implemented according to the C/IDS specifications. Firstly, we have disabled ping to mitigate ping attacks. Additionally, we have enhanced logging and auditing of visitor IP addresses. Since our business operations are primarily within China, we have blocked all requests from foreign IPs, effectively preventing DDoS attacks.`
	faqMap[`Is a central Identity Management system (IdM) provided by CI used for the compliant management of access rights?`] = `This is not an internal Bosch system; it is intended for external users. Although it is not integrated with IdM, we do have relevant documentation that adheres to the requirements of Bosch's DSO department to record user permissions.`

	faqMap[`Is there a central management database where the interfaces of the IT system are documented?`] = `Yes, we have interface development documentation for all our projects, which includes all the interfaces.`
	faqMap[`Is the test data being used protected (according to its security class) and deleted from the development or test systems at the end of the tests?`] = `Yes, we will delete all test data at the end of the tests.`
	faqMap[``] = ``

	OperatingManualV2, err := ListDir("/home/eric/Project/Go/go-rblc-dso-excel/2024-Operating_Manual/1", "xlsm")
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
