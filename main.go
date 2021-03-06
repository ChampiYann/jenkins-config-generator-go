package main

import (
	"fmt"
	"jenkins-config-generator-go/jcasc"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Input
	projectName := "Project 4 Awesome"
	adminUserMail := "hello-world.admin@example.org"
	numberOfNodes := 2

	// Format input
	projectNameDashed := strings.Replace(projectName, " ", "-", -1)
	// projectNameCaps := strings.ToUpper(strings.Replace(projectName, " ", "_", -1))

	// Create configuration object
	config := new(jcasc.JCasC)

	// Add customer folder
	folderName := "folder('" + projectName + "')"
	job := jcasc.Job{Script: &folderName}
	config.Jobs = append(config.Jobs, job)

	// Create admin user
	adminPassword := ""
	adminUser := jcasc.User{Name: &adminUserMail, ID: &adminUserMail, Password: &adminPassword}
	mailerProperty := jcasc.UserProperty{Mailer: &jcasc.Mailer{EmailAddress: &adminUserMail}}
	adminUser.Properties = append(adminUser.Properties, mailerProperty)

	config.Jenkins = &jcasc.JenkinsClass{SecurityRealm: &jcasc.SecurityRealm{Local: &jcasc.Local{}}}
	config.Jenkins.SecurityRealm.Local.Users = append(config.Jenkins.SecurityRealm.Local.Users, adminUser)

	globalPermissions := []string{
		"USER:Overall/Read:" + adminUserMail,
	}
	config.Jenkins.AuthorizationStrategy = &jcasc.AuthorizationStrategy{ProjectMatrix: &jcasc.AuthorizationMatrixNodeProperty{Permissions: globalPermissions}}

	// Create SA
	saUserName := projectNameDashed + "-sa-build-1"
	saPassword := ""
	saUser := jcasc.User{Name: &saUserName, ID: &saUserName, Password: &saPassword}
	config.Jenkins.SecurityRealm.Local.Users = append(config.Jenkins.SecurityRealm.Local.Users, saUser)

	// Add Node(s)
	for i := 0; i < numberOfNodes; i++ {
		nodeName := projectNameDashed + "-" + strconv.Itoa(i+1)
		node := jcasc.Node{Permanent: &jcasc.Permanent{Name: &nodeName}}
		nodeAuthorizationMatrixInheritanceStrategy := jcasc.InheritingGlobal
		nodePermissions := []string{
			"USER:Agent/Build:" + saUserName,
			"USER:Agent/Connect:" + adminUserMail,
			"USER:Agent/Build:" + adminUserMail,
			"USER:Agent/Configure:" + adminUserMail,
			"USER:Agent/Disconnect:" + adminUserMail,
		}
		nodeProperty := jcasc.PermanentNodeProperty{AuthorizationMatrix: &jcasc.AuthorizationMatrixNodeProperty{InheritanceStrategy: &nodeAuthorizationMatrixInheritanceStrategy, Permissions: nodePermissions}}
		node.Permanent.NodeProperties = append(node.Permanent.NodeProperties, nodeProperty)

		config.Jenkins.Nodes = append(config.Jenkins.Nodes, node)
	}

	bytes, err := config.MarshalYAML()
	if err == nil {
		fmt.Println(string(bytes))
	}

	err = os.WriteFile(projectNameDashed+".yaml", bytes, 0644)
	check(err)
}
