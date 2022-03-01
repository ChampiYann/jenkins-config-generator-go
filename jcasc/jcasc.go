// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    jCasC, err := UnmarshalJCasC(bytes)
//    bytes, err = jCasC.Marshal()

package jcasc

import (
	"github.com/ghodss/yaml"
)

// func UnmarshalJCasC(data []byte) (JCasC, error) {
// 	var r JCasC
// 	err := json.Unmarshal(data, &r)
// 	return r, err
// }

// func (r *JCasC) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(r)
// }

func (r *JCasC) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(r)
}

type JCasC struct {
	Jenkins *JenkinsClass `json:"jenkins,omitempty"`
	Jobs    []Job         `json:"jobs,omitempty"`
}

type Job struct {
	File   *string `json:"file,omitempty"`
	Script *string `json:"script,omitempty"`
	URL    *string `json:"url,omitempty"`
}

type JenkinsClass struct {
	// AgentProtocols                 *string                      `json:"agentProtocols,omitempty"`
	AuthorizationStrategy *AuthorizationStrategy `json:"authorizationStrategy,omitempty"`
	// Clouds                         []Cloud                      `json:"clouds,omitempty"`
	// CrumbIssuer                    *FluffyCrumbIssuer           `json:"crumbIssuer,omitempty"`
	// DisabledAdministrativeMonitors *string                      `json:"disabledAdministrativeMonitors,omitempty"`
	// DisableRememberMe              *bool                        `json:"disableRememberMe,omitempty"`
	// GlobalNodeProperties           []GlobalNodeProperty         `json:"globalNodeProperties,omitempty"`
	// LabelAtoms                     []LabelAtomElement           `json:"labelAtoms,omitempty"`
	// LabelString                    *string                      `json:"labelString,omitempty"`
	// Log                            *PurpleLog                   `json:"log,omitempty"`
	// MarkupFormatter                *PurpleMarkupFormatter       `json:"markupFormatter,omitempty"` // In such places as project description, user description, view description, and build; description,; Jenkins allows users to enter some free-form text that describes something.; ; This configuration determines how such free-form text is converted to HTML. By default,; Jenkins treats; the text as HTML and use it as-is unmodified (and this is default mainly because of the; backward compatibility.); ; ; While this is convenient and people often use it to load &lt;iframe>, &lt;script>. and so; on to; mash up data from other sources, this capability enables malicious users to mount; XSS attacks.; If the risk outweighs the benefit, install additional markup formatter plugins and use; them.
	// Mode                           *Mode                        `json:"mode,omitempty"`
	// MyViewsTabBar                  *PurpleMyViewsTabBar         `json:"myViewsTabBar,omitempty"`
	// NodeName                       *string                      `json:"nodeName,omitempty"`
	// NodeProperties                 []PurpleNodeProperty         `json:"nodeProperties,omitempty"`
	Nodes []Node `json:"nodes,omitempty"`
	// NoUsageStatistics              *bool                        `json:"noUsageStatistics,omitempty"`
	// NumExecutors                   *int64                       `json:"numExecutors,omitempty"`
	// PrimaryView                    *PrimaryView                 `json:"primaryView,omitempty"`
	// ProjectNamingStrategy          *PurpleProjectNamingStrategy `json:"projectNamingStrategy,omitempty"`
	// Proxy                          *PurpleProxy                 `json:"proxy,omitempty"`
	// QuietPeriod                    *int64                       `json:"quietPeriod,omitempty"`
	// RemotingSecurity               *PurpleRemotingSecurity      `json:"remotingSecurity,omitempty"`
	// SCMCheckoutRetryCount          *int64                       `json:"scmCheckoutRetryCount,omitempty"`
	SecurityRealm *SecurityRealm `json:"securityRealm,omitempty"`
	// SlaveAgentPort                 *int64                       `json:"slaveAgentPort,omitempty"`
	// SystemMessage                  *string                      `json:"systemMessage,omitempty"`
	// UpdateCenter                   *PurpleUpdateCenter          `json:"updateCenter,omitempty"`
	// Views                          []ViewElement                `json:"views,omitempty"`
	// ViewsTabBar                    *PurpleViewsTabBar           `json:"viewsTabBar,omitempty"`
}

type AuthorizationStrategy struct {
	ProjectMatrix              *AuthorizationMatrixNodeProperty `json:"projectMatrix,omitempty"`
	GlobalMatrix               interface{}                      `json:"globalMatrix,omitempty"`
	Legacy                     interface{}                      `json:"legacy,omitempty"`
	LoggedInUsersCanDoAnything interface{}                      `json:"loggedInUsersCanDoAnything,omitempty"`
	Unsecured                  interface{}                      `json:"unsecured,omitempty"`
}

type SecurityRealm struct {
	Legacy interface{} `json:"legacy,omitempty"`
	None   interface{} `json:"none,omitempty"`
	Pam    interface{} `json:"pam,omitempty"`
	Local  *Local      `json:"local,omitempty"`
}

type Local struct {
	AllowsSignup   *bool                `json:"allowsSignup,omitempty"` // This option allows users to create accounts by themselves, via the "sign up" link on the; top right shoulder of the page.; Make sure to not grant significant permissions to authenticated users, as anyone on the; network will be able to get them.; ; When this checkbox is unchecked, someone with the administrator role would have to create; accounts.; ; By default, Jenkins does not use captcha verification if the user creates an account by; themself.; If you'd like to enable captcha verification, install a captcha support plugin, e.g. the; Jenkins; JCaptcha Plugin.
	CAPTCHASupport *LocalCAPTCHASupport `json:"captchaSupport,omitempty"`
	EnableCAPTCHA  *bool                `json:"enableCaptcha,omitempty"`
	Users          []User               `json:"users,omitempty"`
}

type LocalCAPTCHASupport struct {
}

type User struct {
	Description *string        `json:"description,omitempty"`
	ID          *string        `json:"id,omitempty"`
	Name        *string        `json:"name,omitempty"`
	Password    *string        `json:"password,omitempty"`
	Properties  []UserProperty `json:"properties,omitempty"`
}

type UserProperty struct {
	LastGrantedAuthorities  interface{} `json:"lastGrantedAuthorities,omitempty"`
	SSHPublicKey            interface{} `json:"sshPublicKey,omitempty"`
	UserSeed                interface{} `json:"userSeed,omitempty"`
	Search                  interface{} `json:"search,omitempty"`
	PaneStatus              interface{} `json:"paneStatus,omitempty"`
	Password                interface{} `json:"password,omitempty"`
	MyView                  interface{} `json:"myView,omitempty"`
	APIToken                interface{} `json:"apiToken,omitempty"`
	Timezone                interface{} `json:"timezone,omitempty"`
	Mailer                  *Mailer     `json:"mailer,omitempty"`
	UserCredentialsProperty interface{} `json:"userCredentialsProperty,omitempty"`
	PreferredProvider       interface{} `json:"preferredProvider,omitempty"`
}

type Mailer struct {
	EmailAddress *string `json:"emailAddress,omitempty"`
}

type Node struct {
	Permanent *Permanent  `json:"permanent,omitempty"`
	Jenkins   interface{} `json:"jenkins,omitempty"`
}

type Permanent struct {
	LabelString       *string                     `json:"labelString,omitempty"`
	Launcher          *Launcher                   `json:"launcher,omitempty"`
	Mode              *Mode                       `json:"mode,omitempty"`
	Name              *string                     `json:"name,omitempty"`
	NodeDescription   *string                     `json:"nodeDescription,omitempty"`
	NodeName          *string                     `json:"nodeName,omitempty"`
	NodeProperties    []PermanentNodeProperty     `json:"nodeProperties,omitempty"`
	NumExecutors      *int64                      `json:"numExecutors,omitempty"`
	RemoteFS          *string                     `json:"remoteFS,omitempty"`
	RetentionStrategy *PermanentRetentionStrategy `json:"retentionStrategy,omitempty"`
}

type Launcher struct {
}

type PermanentRetentionStrategy struct {
}

type PermanentNodeProperty struct {
	DisableDeferredWipeout interface{}                      `json:"disableDeferredWipeout,omitempty"`
	ToolLocation           interface{}                      `json:"toolLocation,omitempty"`
	EnvVars                interface{}                      `json:"envVars,omitempty"`
	AuthorizationMatrix    *AuthorizationMatrixNodeProperty `json:"authorizationMatrix,omitempty"`
}

type AuthorizationMatrixNodeProperty struct {
	InheritanceStrategy *AuthorizationMatrixNodePropertyInheritanceStrategy `json:"inheritanceStrategy,omitempty"`
	Permissions         []string                                            `json:"permissions,omitempty"`
}

type AuthorizationMatrixNodePropertyInheritanceStrategy string

const (
	InheritingGlobal AuthorizationMatrixNodePropertyInheritanceStrategy = "inheritingGlobal"
	NonInheriting    AuthorizationMatrixNodePropertyInheritanceStrategy = "nonInheriting"
)

type Mode string

const (
	Exclusive Mode = "EXCLUSIVE"
	Normal    Mode = "NORMAL"
)
