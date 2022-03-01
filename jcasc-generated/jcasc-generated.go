// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    jCasC, err := UnmarshalJCasC(bytes)
//    bytes, err = jCasC.Marshal()

package main

import "encoding/json"

func UnmarshalJCasC(data []byte) (JCasC, error) {
	var r JCasC
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *JCasC) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Jenkins Configuration as Code
type JCasC struct {
	ConfigurationAsCode            *ConfigurationBaseForTheConfigurationAsCodeClassifier            `json:"configuration-as-code,omitempty"`
	Credentials                    *ConfigurationBaseForTheCredentialsClassifier                    `json:"credentials,omitempty"`
	GlobalCredentialsConfiguration *ConfigurationBaseForTheGlobalCredentialsConfigurationClassifier `json:"globalCredentialsConfiguration,omitempty"`
	Jenkins                        *ConfigurationBaseForTheJenkinsClassifier                        `json:"jenkins,omitempty"`
	Jobs                           *ConfigurationBaseForTheJobsClassifier                           `json:"jobs,omitempty"`
	Security                       *ConfigurationBaseForTheSecurityClassifier                       `json:"security,omitempty"`
	Tool                           *ConfigurationBaseForTheToolClassifier                           `json:"tool,omitempty"`
	Unclassified                   *ConfigurationBaseForTheUnclassifiedClassifier                   `json:"unclassified,omitempty"`
}

type ConfigurationBaseForTheConfigurationAsCodeClassifier struct {
	Deprecated *Deprecated `json:"deprecated,omitempty"`
	Restricted *Restricted `json:"restricted,omitempty"`
	Unknown    *Deprecated `json:"unknown,omitempty"`
	Version    *Version    `json:"version,omitempty"`
}

type ConfigurationBaseForTheCredentialsClassifier struct {
	BasicSSHUserPrivateKey    *CredentialsBasicSSHUserPrivateKey    `json:"basicSSHUserPrivateKey,omitempty"`
	Certificate               *CredentialsCertificate               `json:"certificate,omitempty"`
	Credentials               *CredentialsCredentials               `json:"credentials,omitempty"`
	DirectEntry               *CredentialsDirectEntry               `json:"directEntry,omitempty"`
	Domain                    *CredentialsDomain                    `json:"domain,omitempty"`
	DomainCredentials         *CredentialsDomainCredentials         `json:"domainCredentials,omitempty"`
	DomainSpecification       *CredentialsDomainSpecification       `json:"domainSpecification,omitempty"`
	File                      *CredentialsFile                      `json:"file,omitempty"`
	Folder                    *Folder                               `json:"folder,omitempty"`
	FolderCredentialsProvider *FolderCredentialsProvider            `json:"folderCredentialsProvider,omitempty"`
	HostnamePortSpecification *CredentialsHostnamePortSpecification `json:"hostnamePortSpecification,omitempty"`
	HostnameSpecification     *CredentialsHostnameSpecification     `json:"hostnameSpecification,omitempty"`
	KeyStoreSource            *CredentialsKeyStoreSource            `json:"keyStoreSource,omitempty"`
	PathSpecification         *CredentialsPathSpecification         `json:"pathSpecification,omitempty"`
	PrivateKeySource          *CredentialsPrivateKeySource          `json:"privateKeySource,omitempty"`
	ProviderImpl              *ProviderImpl                         `json:"providerImpl,omitempty"`
	SchemeSpecification       *CredentialsSchemeSpecification       `json:"schemeSpecification,omitempty"`
	String                    *CredentialsString                    `json:"string,omitempty"`
	System                    *SystemClass                          `json:"system,omitempty"`
	Uploaded                  *CredentialsUploaded                  `json:"uploaded,omitempty"`
	User                      *CredentialsUser                      `json:"user,omitempty"`
	UserCredentialsProvider   *UserCredentialsProvider              `json:"userCredentialsProvider,omitempty"`
	UsernamePassword          *CredentialsUsernamePassword          `json:"usernamePassword,omitempty"`
}

type CredentialsBasicSSHUserPrivateKey struct {
	Description      *string                 `json:"description,omitempty"`
	ID               *string                 `json:"id,omitempty"`
	Passphrase       *string                 `json:"passphrase,omitempty"`
	PrivateKeySource *PurplePrivateKeySource `json:"privateKeySource,omitempty"`
	Scope            *Scope                  `json:"scope,omitempty"`
	Username         *string                 `json:"username,omitempty"`
	UsernameSecret   *bool                   `json:"usernameSecret,omitempty"` // By default, usernames are not masked in the build log.; You may choose to mask a credential's username if you consider it to be sensitive; information.; However, this can interfere with diagnostics.; For example, if the username is a common word it will cause unrelated occurrences of that; word to also be masked.; ; Regardless of this setting, the username will be displayed in the selection dropdown to; anyone permitted to reconfigure the credentials.
}

type PurplePrivateKeySource struct {
}

type CredentialsCertificate struct {
	Description    *string               `json:"description,omitempty"`
	ID             *string               `json:"id,omitempty"`
	KeyStoreSource *PurpleKeyStoreSource `json:"keyStoreSource,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; The source of the certificate.
	Password       *string               `json:"password,omitempty"`       // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; The password. The use of separate integrity and encryption passwords is not supported.
	Scope          *Scope                `json:"scope,omitempty"`
}

// <!--
// ~ The MIT License
// ~
// ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.
// ~
// ~ Permission is hereby granted, free of charge, to any person obtaining a copy
// ~ of this software and associated documentation files (the "Software"), to deal
// ~ in the Software without restriction, including without limitation the rights
// ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// ~ copies of the Software, and to permit persons to whom the Software is
// ~ furnished to do so, subject to the following conditions:
// ~
// ~ The above copyright notice and this permission notice shall be included in
// ~ all copies or substantial portions of the Software.
// ~
// ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// ~ THE SOFTWARE.
// -->
//
//
// The source of the certificate.
type PurpleKeyStoreSource struct {
}

type CredentialsCredentials struct {
	File                   interface{} `json:"file"`
	String                 interface{} `json:"string"`
	BasicSSHUserPrivateKey interface{} `json:"basicSSHUserPrivateKey"`
	Certificate            interface{} `json:"certificate"`
	UsernamePassword       interface{} `json:"usernamePassword"`
}

type CredentialsDirectEntry struct {
	PrivateKey *PurplePrivateKey `json:"privateKey,omitempty"`
}

type PurplePrivateKey struct {
}

type CredentialsDomain struct {
	Description    *string               `json:"description,omitempty"`
	Name           *string               `json:"name,omitempty"`
	Specifications []PurpleSpecification `json:"specifications,omitempty"`
}

type PurpleSpecification struct {
}

type CredentialsDomainCredentials struct {
	Credentials []PurpleCredential `json:"credentials,omitempty"`
	Domain      *PurpleDomain      `json:"domain,omitempty"`
}

type PurpleCredential struct {
}

type PurpleDomain struct {
}

type CredentialsDomainSpecification struct {
	SchemeSpecification       interface{} `json:"schemeSpecification"`
	HostnamePortSpecification interface{} `json:"hostnamePortSpecification"`
	PathSpecification         interface{} `json:"pathSpecification"`
	HostnameSpecification     interface{} `json:"hostnameSpecification"`
}

type CredentialsFile struct {
	Description *string            `json:"description,omitempty"`
	FileName    *string            `json:"fileName,omitempty"`
	ID          *string            `json:"id,omitempty"`
	Scope       *Scope             `json:"scope,omitempty"`
	SecretBytes *PurpleSecretBytes `json:"secretBytes,omitempty"`
}

type PurpleSecretBytes struct {
}

type Folder struct {
}

type FolderCredentialsProvider struct {
}

type CredentialsHostnamePortSpecification struct {
	Excludes *string `json:"excludes,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of excluded hostnames. (The * wildcard is permitted in; hostnames,; for example: jenkins-ci.org:*, *.jenkins-ci.org:80, jenkinsci.github.io:443.); ; The empty list implies no hostname:port is excluded. The excludes list is processed after; the includes; list.
	Includes *string `json:"includes,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of included hostnames. (The * wildcard is permitted in; hostnames,; for example: jenkins-ci.org:*, *.jenkins-ci.org:80, jenkinsci.github.io:443.); ; The empty list implies no hostname:port is excluded. The excludes list is processed after; the includes; list.
}

type CredentialsHostnameSpecification struct {
	Excludes *string `json:"excludes,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of excluded hostnames. (The * wildcard is permitted in; hostnames,; for example: jenkins-ci.org, *.jenkins-ci.org, jenkinsci.github.io.); ; The empty list implies no hostnames are excluded. The excludes list is processed after; the includes list.
	Includes *string `json:"includes,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of included hostnames. (The * wildcard is permitted in; hostnames,; for example: jenkins-ci.org, *.jenkins-ci.org, jenkinsci.github.io.); ; The empty list implies no hostnames are excluded. The excludes list is processed after; the includes list.
}

type CredentialsKeyStoreSource struct {
	Uploaded interface{} `json:"uploaded"`
}

type CredentialsPathSpecification struct {
	CaseSensitive *bool   `json:"caseSensitive,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2014, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; Select this option to require that the paths match taking character case into account.
	Excludes      *string `json:"excludes,omitempty"`      // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2014, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of excluded paths. (ANT style * and ** wildcards are permitted in; paths,; for example: /jenkins/github/*,/jenkins-ci/**/org,jenkinsci.github.io.); ; The empty list implies no paths are excluded. The excludes list is processed after the; includes list.
	Includes      *string `json:"includes,omitempty"`      // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2014, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of included paths. (ANT style * and ** wildcards are permitted in; paths,; for example: /jenkins/github/*,/jenkins-ci/**/org,jenkinsci.github.io.); ; The empty list implies no paths are excluded. The excludes list is processed after the; includes list.
}

type CredentialsPrivateKeySource struct {
	DirectEntry interface{} `json:"directEntry"`
}

type ProviderImpl struct {
}

type CredentialsSchemeSpecification struct {
	Schemes *string `json:"schemes,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of URI schemes (as defined in RFC-3986).; For example: http, https, ssh, sftp, imap.
}

type CredentialsString struct {
	Description *string       `json:"description,omitempty"`
	ID          *string       `json:"id,omitempty"`
	Scope       *Scope        `json:"scope,omitempty"`
	Secret      *PurpleSecret `json:"secret,omitempty"`
}

type PurpleSecret struct {
}

type SystemClass struct {
}

type CredentialsUploaded struct {
	UploadedKeystore *PurpleUploadedKeystore `json:"uploadedKeystore,omitempty"`
}

type PurpleUploadedKeystore struct {
}

type CredentialsUser struct {
}

type UserCredentialsProvider struct {
}

type CredentialsUsernamePassword struct {
	Description    *string `json:"description,omitempty"`
	ID             *string `json:"id,omitempty"`
	Password       *string `json:"password,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2012, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; The password.
	Scope          *Scope  `json:"scope,omitempty"`
	Username       *string `json:"username,omitempty"`       // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2012, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; The username.
	UsernameSecret *bool   `json:"usernameSecret,omitempty"` // Historically, not only passwords but usernames were masked in the build log.; Since these can interfere with diagnostics,; and cause unrelated occurrences of a common word to be masked,; you may choose to leave usernames unmasked if they are not sensitive.; Note that regardless of this setting, the username will be displayed to anyone permitted; to reconfigure the credentials.
}

type ConfigurationBaseForTheGlobalCredentialsConfigurationClassifier struct {
	Configuration                      *Configuration                      `json:"configuration,omitempty"`
	CredentialsProviderFilter          *CredentialsProviderFilter          `json:"credentialsProviderFilter,omitempty"`
	CredentialsProviderTypeRestriction *CredentialsProviderTypeRestriction `json:"credentialsProviderTypeRestriction,omitempty"`
	CredentialsTypeFilter              *CredentialsTypeFilter              `json:"credentialsTypeFilter,omitempty"`
	Excludes                           *Excludes                           `json:"excludes,omitempty"`
	Includes                           *Includes                           `json:"includes,omitempty"`
	None                               *GlobalCredentialsConfigurationNone `json:"none,omitempty"`
}

type Configuration struct {
	ProviderFilter *ProviderFilter `json:"providerFilter,omitempty"`
	Restrictions   []Restriction   `json:"restrictions,omitempty"`
	TypeFilter     *TypeFilter     `json:"typeFilter,omitempty"`
}

type ProviderFilter struct {
}

type Restriction struct {
}

type TypeFilter struct {
}

type CredentialsProviderFilter struct {
	Excludes interface{} `json:"excludes"`
	Includes interface{} `json:"includes"`
	None     interface{} `json:"none"`
}

type CredentialsProviderTypeRestriction struct {
	Excludes interface{} `json:"excludes"`
	Includes interface{} `json:"includes"`
}

type CredentialsTypeFilter struct {
	Excludes interface{} `json:"excludes"`
	Includes interface{} `json:"includes"`
	None     interface{} `json:"none"`
}

type Excludes struct {
	Provider *string `json:"provider,omitempty"`
	Type     *string `json:"type,omitempty"`
}

type Includes struct {
	Provider *string `json:"provider,omitempty"`
	Type     *string `json:"type,omitempty"`
}

type GlobalCredentialsConfigurationNone struct {
}

type ConfigurationBaseForTheJenkinsClassifier struct {
	Empty                                   *Jenkins                                 `json:",omitempty"`
	AgentProtocols                          *string                                  `json:"agentProtocols,omitempty"`
	All                                     *All                                     `json:"all,omitempty"`
	Always                                  *Always                                  `json:"always,omitempty"`
	APIToken                                *JenkinsAPIToken                         `json:"apiToken,omitempty"`
	AuthorizationMatrixNodeProperty         *AuthorizationMatrixNodeProperty         `json:"authorizationMatrixNodeProperty,omitempty"`
	AuthorizationStrategy                   *PurpleAuthorizationStrategy             `json:"authorizationStrategy,omitempty"`
	BasicSSHUserPrivateKey                  *JenkinsBasicSSHUserPrivateKey           `json:"basicSSHUserPrivateKey,omitempty"`
	BatchFile                               *JenkinsBatchFile                        `json:"batchFile,omitempty"`
	BranchStatusColumn                      *BranchStatusColumn                      `json:"branchStatusColumn,omitempty"`
	BuildButton                             *BuildButton                             `json:"buildButton,omitempty"`
	CAPTCHASupport                          interface{}                              `json:"captchaSupport"`
	Certificate                             *JenkinsCertificate                      `json:"certificate,omitempty"`
	Cloud                                   interface{}                              `json:"cloud"`
	Command                                 *JenkinsCommand                          `json:"command,omitempty"`
	ComputerLauncher                        *ComputerLauncher                        `json:"computerLauncher,omitempty"`
	Credentials                             *JenkinsCredentials                      `json:"credentials,omitempty"`
	CrumbIssuer                             *PurpleCrumbIssuer                       `json:"crumbIssuer,omitempty"`
	Demand                                  *Demand                                  `json:"demand,omitempty"`
	DescriptionColumn                       *DescriptionColumn                       `json:"descriptionColumn,omitempty"`
	DirectEntry                             *JenkinsDirectEntry                      `json:"directEntry,omitempty"`
	DisabledAdministrativeMonitors          *string                                  `json:"disabledAdministrativeMonitors,omitempty"`
	DisableRememberMe                       *bool                                    `json:"disableRememberMe,omitempty"`
	Domain                                  *JenkinsDomain                           `json:"domain,omitempty"`
	DomainCredentials                       *JenkinsDomainCredentials                `json:"domainCredentials,omitempty"`
	DomainSpecification                     *JenkinsDomainSpecification              `json:"domainSpecification,omitempty"`
	Entry                                   *Entry                                   `json:"entry,omitempty"`
	File                                    *JenkinsFile                             `json:"file,omitempty"`
	GitBranchSpecifierColumn                *GitBranchSpecifierColumn                `json:"gitBranchSpecifierColumn,omitempty"`
	HostnamePortSpecification               *JenkinsHostnamePortSpecification        `json:"hostnamePortSpecification,omitempty"`
	HostnameSpecification                   *JenkinsHostnameSpecification            `json:"hostnameSpecification,omitempty"`
	InheritanceStrategy                     *JenkinsInheritanceStrategy              `json:"inheritanceStrategy,omitempty"`
	Inheriting                              *Inheriting                              `json:"inheriting,omitempty"`
	InheritingGlobal                        *InheritingGlobal                        `json:"inheritingGlobal,omitempty"`
	ItemColumn                              *ItemColumn                              `json:"itemColumn,omitempty"`
	JDK                                     *JenkinsJDK                              `json:"jdk,omitempty"`
	JDKInstaller                            *JenkinsJDKInstaller                     `json:"jdkInstaller,omitempty"`
	JDKs                                    *JDKs                                    `json:"jDKs,omitempty"`
	Jenkins                                 *JenkinsClass                            `json:"jenkins,omitempty"`
	Jnlp                                    *Jnlp                                    `json:"jnlp,omitempty"`
	JobName                                 *JobName                                 `json:"jobName,omitempty"`
	KeyStoreSource                          *JenkinsKeyStoreSource                   `json:"keyStoreSource,omitempty"`
	KnownHostsFileKeyVerificationStrategy   *KnownHostsFileKeyVerificationStrategy   `json:"knownHostsFileKeyVerificationStrategy,omitempty"`
	LabelAtom                               *PurpleLabelAtom                         `json:"labelAtom,omitempty"`
	LabelAtomProperty                       interface{}                              `json:"labelAtomProperty"`
	LabelAtoms                              *LabelAtoms                              `json:"labelAtoms,omitempty"`
	LabelString                             *string                                  `json:"labelString,omitempty"`
	LastDuration                            *LastDuration                            `json:"lastDuration,omitempty"`
	LastFailure                             *LastFailure                             `json:"lastFailure,omitempty"`
	LastStable                              *LastStable                              `json:"lastStable,omitempty"`
	LastSuccess                             *LastSuccess                             `json:"lastSuccess,omitempty"`
	Legacy                                  *Legacy                                  `json:"legacy,omitempty"`
	List                                    *List                                    `json:"list,omitempty"`
	ListViewColumn                          *ListViewColumn                          `json:"listViewColumn,omitempty"`
	Local                                   *Local                                   `json:"local,omitempty"`
	Log                                     *FluffyLog                               `json:"log,omitempty"`
	LoggedInUsersCanDoAnything              *LoggedInUsersCanDoAnything              `json:"loggedInUsersCanDoAnything,omitempty"`
	LogRecorder                             *LogRecorder                             `json:"logRecorder,omitempty"`
	Mailer                                  *JenkinsMailer                           `json:"mailer,omitempty"`
	ManuallyProvidedKeyVerificationStrategy *ManuallyProvidedKeyVerificationStrategy `json:"manuallyProvidedKeyVerificationStrategy,omitempty"`
	ManuallyTrustedKeyVerificationStrategy  *ManuallyTrustedKeyVerificationStrategy  `json:"manuallyTrustedKeyVerificationStrategy,omitempty"`
	MarkupFormatter                         *FluffyMarkupFormatter                   `json:"markupFormatter,omitempty"`
	Maven                                   *JenkinsMaven                            `json:"maven,omitempty"`
	Mode                                    *Mode                                    `json:"mode,omitempty"`
	MyView                                  *JenkinsMyView                           `json:"myView,omitempty"`
	MyViewsTabBar                           *FluffyMyViewsTabBar                     `json:"myViewsTabBar,omitempty"`
	Node                                    *PurpleNode                              `json:"node,omitempty"`
	NodeName                                *string                                  `json:"nodeName,omitempty"`
	NodeProperty                            *FluffyNodeProperty                      `json:"nodeProperty,omitempty"`
	NonInheriting                           *NonInheriting                           `json:"nonInheriting,omitempty"`
	NonVerifyingKeyVerificationStrategy     *NonVerifyingKeyVerificationStrategy     `json:"nonVerifyingKeyVerificationStrategy,omitempty"`
	NoUsageStatistics                       *bool                                    `json:"noUsageStatistics,omitempty"`
	NumExecutors                            *int64                                   `json:"numExecutors,omitempty"`
	Pam                                     *Pam                                     `json:"pam,omitempty"`
	PathSpecification                       *JenkinsPathSpecification                `json:"pathSpecification,omitempty"`
	Pattern                                 *Pattern                                 `json:"pattern,omitempty"`
	Permanent                               *Permanent                               `json:"permanent,omitempty"`
	PlainText                               *PlainText                               `json:"plainText,omitempty"`
	PreferredProvider                       *PreferredProvider                       `json:"preferredProvider,omitempty"`
	PrivateKeySource                        *JenkinsPrivateKeySource                 `json:"privateKeySource,omitempty"`
	ProjectNamingStrategy                   *FluffyProjectNamingStrategy             `json:"projectNamingStrategy,omitempty"`
	Proxy                                   *FluffyProxy                             `json:"proxy,omitempty"`
	QuietPeriod                             *int64                                   `json:"quietPeriod,omitempty"`
	RawBuildsDir                            *string                                  `json:"rawBuildsDir,omitempty"`
	RawHTML                                 *RawHTML                                 `json:"rawHtml,omitempty"`
	RemotingSecurity                        *FluffyRemotingSecurity                  `json:"remotingSecurity,omitempty"`
	RetentionStrategy                       *JenkinsRetentionStrategy                `json:"retentionStrategy,omitempty"`
	Schedule                                *Schedule                                `json:"schedule,omitempty"`
	SchemeSpecification                     *JenkinsSchemeSpecification              `json:"schemeSpecification,omitempty"`
	SCMCheckoutRetryCount                   *int64                                   `json:"scmCheckoutRetryCount,omitempty"`
	SecurityRealm                           *FluffySecurityRealm                     `json:"securityRealm,omitempty"`
	SlaveAgentPort                          *int64                                   `json:"slaveAgentPort,omitempty"`
	SSH                                     *SSH                                     `json:"ssh,omitempty"`
	SSHHostKeyVerificationStrategy          *JenkinsSSHHostKeyVerificationStrategy   `json:"sshHostKeyVerificationStrategy,omitempty"`
	SSHPublicKey                            *SSHPublicKey                            `json:"sshPublicKey,omitempty"`
	Standard                                *JenkinsStandard                         `json:"standard,omitempty"`
	Status                                  *Status                                  `json:"status,omitempty"`
	StatusFilter                            *StatusFilter                            `json:"statusFilter,omitempty"`
	String                                  *JenkinsString                           `json:"string,omitempty"`
	SystemMessage                           *string                                  `json:"systemMessage,omitempty"`
	Target                                  *Target                                  `json:"target,omitempty"`
	Timezone                                *Timezone                                `json:"timezone,omitempty"`
	ToolInstaller                           *JenkinsToolInstaller                    `json:"toolInstaller,omitempty"`
	ToolLocation                            *ToolLocation                            `json:"toolLocation,omitempty"`
	ToolProperty                            *JenkinsToolProperty                     `json:"toolProperty,omitempty"`
	Unsecured                               *Unsecured                               `json:"unsecured,omitempty"`
	UpdateCenter                            *FluffyUpdateCenter                      `json:"updateCenter,omitempty"`
	UpdateSite                              *UpdateSite                              `json:"updateSite,omitempty"`
	Uploaded                                *JenkinsUploaded                         `json:"uploaded,omitempty"`
	UsernamePassword                        *JenkinsUsernamePassword                 `json:"usernamePassword,omitempty"`
	UserProperty                            *UserProperty                            `json:"userProperty,omitempty"`
	UserWithPassword                        *UserWithPassword                        `json:"userWithPassword,omitempty"`
	View                                    *PurpleView                              `json:"view,omitempty"`
	ViewJobFilter                           *ViewJobFilter                           `json:"viewJobFilter,omitempty"`
	ViewsTabBar                             *FluffyViewsTabBar                       `json:"viewsTabBar,omitempty"`
	Weather                                 *Weather                                 `json:"weather,omitempty"`
	Zip                                     *JenkinsZip                              `json:"zip,omitempty"`
}

type JenkinsAPIToken struct {
}

type All struct {
	Description *string       `json:"description,omitempty"`
	Name        *string       `json:"name,omitempty"`
	Properties  []AllProperty `json:"properties,omitempty"`
}

type AllProperty struct {
}

type Always struct {
}

type AuthorizationMatrixNodeProperty struct {
	InheritanceStrategy *AuthorizationMatrixNodePropertyInheritanceStrategy `json:"inheritanceStrategy,omitempty"`
	Permissions         *string                                             `json:"permissions,omitempty"`
}

type AuthorizationMatrixNodePropertyInheritanceStrategy struct {
}

type PurpleAuthorizationStrategy struct {
	ProjectMatrix              interface{} `json:"projectMatrix"`
	GlobalMatrix               interface{} `json:"globalMatrix"`
	Legacy                     interface{} `json:"legacy"`
	LoggedInUsersCanDoAnything interface{} `json:"loggedInUsersCanDoAnything"`
	Unsecured                  interface{} `json:"unsecured"`
}

type JenkinsBasicSSHUserPrivateKey struct {
	Description      *string                 `json:"description,omitempty"`
	ID               *string                 `json:"id,omitempty"`
	Passphrase       *string                 `json:"passphrase,omitempty"`
	PrivateKeySource *FluffyPrivateKeySource `json:"privateKeySource,omitempty"`
	Scope            *Scope                  `json:"scope,omitempty"`
	Username         *string                 `json:"username,omitempty"`
	UsernameSecret   *bool                   `json:"usernameSecret,omitempty"` // By default, usernames are not masked in the build log.; You may choose to mask a credential's username if you consider it to be sensitive; information.; However, this can interfere with diagnostics.; For example, if the username is a common word it will cause unrelated occurrences of that; word to also be masked.; ; Regardless of this setting, the username will be displayed in the selection dropdown to; anyone permitted to reconfigure the credentials.
}

type FluffyPrivateKeySource struct {
}

type JenkinsBatchFile struct {
	Command  *string `json:"command,omitempty"`
	Label    *string `json:"label,omitempty"`
	ToolHome *string `json:"toolHome,omitempty"`
}

type BranchStatusColumn struct {
}

type BuildButton struct {
}

type JenkinsCertificate struct {
	Description    *string               `json:"description,omitempty"`
	ID             *string               `json:"id,omitempty"`
	KeyStoreSource *FluffyKeyStoreSource `json:"keyStoreSource,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; The source of the certificate.
	Password       *string               `json:"password,omitempty"`       // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; The password. The use of separate integrity and encryption passwords is not supported.
	Scope          *Scope                `json:"scope,omitempty"`
}

// <!--
// ~ The MIT License
// ~
// ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.
// ~
// ~ Permission is hereby granted, free of charge, to any person obtaining a copy
// ~ of this software and associated documentation files (the "Software"), to deal
// ~ in the Software without restriction, including without limitation the rights
// ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// ~ copies of the Software, and to permit persons to whom the Software is
// ~ furnished to do so, subject to the following conditions:
// ~
// ~ The above copyright notice and this permission notice shall be included in
// ~ all copies or substantial portions of the Software.
// ~
// ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// ~ THE SOFTWARE.
// -->
//
//
// The source of the certificate.
type FluffyKeyStoreSource struct {
}

type JenkinsCommand struct {
	Command *string `json:"command,omitempty"` // Single command to launch an agent program, which controls the agent; computer and communicates with the controller. Jenkins assumes that; the executed program launches the agent.jar program on the correct; machine.; ; ; A copy of agent.jar can be downloaded from here.; ; ; ; In a simple case, this could be; something like ssh hostname java -jar ~/bin/agent.jar.; ; ; ; Note: the command can't rely on a shell to parse things, e.g. echo foo &gt; bar; baz.; If you need to do that, either use; sh -c or write the expression into a script and point to the script.; ; ; ; It is often a good idea to write a small shell script, like the following, on an; agent; so that you can control the location of Java and/or agent.jar, as well as set up any; environment variables specific to this node, such as PATH.; ; ; ; #!/bin/sh; exec java -jar ~/bin/agent.jar; ; ; ; You can use any command to run a process on the agent machine, such as RSH,; as long as stdin/stdout of the process on the controller will be connected to; those of java -jar ~/bin/agent.jar on the agent machine eventually.; ; ; ; In a larger deployment, it is also worth considering to load agent.jar from; a NFS-mounted common location, so that you don't have to update this file; on every agent machines every time you update Jenkins.; ; ; ; Setting this to ssh -v hostname may be useful for debugging connectivity; issue.
}

type ComputerLauncher struct {
	Jnlp    interface{} `json:"jnlp"`
	SSH     interface{} `json:"ssh"`
	Command interface{} `json:"command"`
}

type JenkinsCredentials struct {
	File                   interface{} `json:"file"`
	String                 interface{} `json:"string"`
	BasicSSHUserPrivateKey interface{} `json:"basicSSHUserPrivateKey"`
	Certificate            interface{} `json:"certificate"`
	UsernamePassword       interface{} `json:"usernamePassword"`
}

type PurpleCrumbIssuer struct {
	Standard interface{} `json:"standard"`
}

type Demand struct {
	IdleDelay     *int64 `json:"idleDelay,omitempty"`
	InDemandDelay *int64 `json:"inDemandDelay,omitempty"`
}

type DescriptionColumn struct {
}

type JenkinsDirectEntry struct {
	PrivateKey *FluffyPrivateKey `json:"privateKey,omitempty"`
}

type FluffyPrivateKey struct {
}

type JenkinsDomain struct {
	Description    *string               `json:"description,omitempty"`
	Name           *string               `json:"name,omitempty"`
	Specifications []FluffySpecification `json:"specifications,omitempty"`
}

type FluffySpecification struct {
}

type JenkinsDomainCredentials struct {
	Credentials []FluffyCredential `json:"credentials,omitempty"`
	Domain      *FluffyDomain      `json:"domain,omitempty"`
}

type FluffyCredential struct {
}

type FluffyDomain struct {
}

type JenkinsDomainSpecification struct {
	SchemeSpecification       interface{} `json:"schemeSpecification"`
	HostnamePortSpecification interface{} `json:"hostnamePortSpecification"`
	PathSpecification         interface{} `json:"pathSpecification"`
	HostnameSpecification     interface{} `json:"hostnameSpecification"`
}

type Jenkins struct {
	Disabled               *bool   `json:"disabled,omitempty"`               // Allows disabling Remoting Work Directory for the agent.; In such case the agent will be running in the legacy mode without logging enabled by; default.
	FailIfWorkDirIsMissing *bool   `json:"failIfWorkDirIsMissing,omitempty"` // If defined, Remoting will fail at startup if the target work directory is missing.; The option may be used to detect infrastructure issues like failed mount.
	InternalDir            *string `json:"internalDir,omitempty"`            // Defines a storage directory for the internal data.; This directory will be created within the Remoting working directory.
	WorkDirPath            *string `json:"workDirPath,omitempty"`            // If defined, a custom Remoting work directory will be used instead of the Agent Root; Directory.; This option has no environment variable resolution so far, it is recommended to use only; absolute paths.
}

type Entry struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type JenkinsFile struct {
	Description *string            `json:"description,omitempty"`
	FileName    *string            `json:"fileName,omitempty"`
	ID          *string            `json:"id,omitempty"`
	Scope       *Scope             `json:"scope,omitempty"`
	SecretBytes *FluffySecretBytes `json:"secretBytes,omitempty"`
}

type FluffySecretBytes struct {
}

type GitBranchSpecifierColumn struct {
}

type JenkinsHostnamePortSpecification struct {
	Excludes *string `json:"excludes,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of excluded hostnames. (The * wildcard is permitted in; hostnames,; for example: jenkins-ci.org:*, *.jenkins-ci.org:80, jenkinsci.github.io:443.); ; The empty list implies no hostname:port is excluded. The excludes list is processed after; the includes; list.
	Includes *string `json:"includes,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of included hostnames. (The * wildcard is permitted in; hostnames,; for example: jenkins-ci.org:*, *.jenkins-ci.org:80, jenkinsci.github.io:443.); ; The empty list implies no hostname:port is excluded. The excludes list is processed after; the includes; list.
}

type JenkinsHostnameSpecification struct {
	Excludes *string `json:"excludes,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of excluded hostnames. (The * wildcard is permitted in; hostnames,; for example: jenkins-ci.org, *.jenkins-ci.org, jenkinsci.github.io.); ; The empty list implies no hostnames are excluded. The excludes list is processed after; the includes list.
	Includes *string `json:"includes,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of included hostnames. (The * wildcard is permitted in; hostnames,; for example: jenkins-ci.org, *.jenkins-ci.org, jenkinsci.github.io.); ; The empty list implies no hostnames are excluded. The excludes list is processed after; the includes list.
}

type JenkinsInheritanceStrategy struct {
	InheritingGlobal interface{} `json:"inheritingGlobal"`
	NonInheriting    interface{} `json:"nonInheriting"`
	Inheriting       interface{} `json:"inheriting"`
}

type Inheriting struct {
}

type InheritingGlobal struct {
}

type ItemColumn struct {
}

type JenkinsJDK struct {
	Home       *string          `json:"home,omitempty"`
	Name       *string          `json:"name,omitempty"`
	Properties []PurpleProperty `json:"properties,omitempty"`
}

type PurpleProperty struct {
}

type JenkinsJDKInstaller struct {
	AcceptLicense *bool   `json:"acceptLicense,omitempty"`
	ID            *string `json:"id,omitempty"`
}

type JDKs struct {
}

type JenkinsClass struct {
	AgentProtocols                 *string                      `json:"agentProtocols,omitempty"`
	AuthorizationStrategy          *FluffyAuthorizationStrategy `json:"authorizationStrategy,omitempty"`
	Clouds                         []Cloud                      `json:"clouds,omitempty"`
	CrumbIssuer                    *FluffyCrumbIssuer           `json:"crumbIssuer,omitempty"`
	DisabledAdministrativeMonitors *string                      `json:"disabledAdministrativeMonitors,omitempty"`
	DisableRememberMe              *bool                        `json:"disableRememberMe,omitempty"`
	GlobalNodeProperties           []GlobalNodeProperty         `json:"globalNodeProperties,omitempty"`
	LabelAtoms                     []LabelAtomElement           `json:"labelAtoms,omitempty"`
	LabelString                    *string                      `json:"labelString,omitempty"`
	Log                            *PurpleLog                   `json:"log,omitempty"`
	MarkupFormatter                *PurpleMarkupFormatter       `json:"markupFormatter,omitempty"` // In such places as project description, user description, view description, and build; description,; Jenkins allows users to enter some free-form text that describes something.; ; This configuration determines how such free-form text is converted to HTML. By default,; Jenkins treats; the text as HTML and use it as-is unmodified (and this is default mainly because of the; backward compatibility.); ; ; While this is convenient and people often use it to load &lt;iframe>, &lt;script>. and so; on to; mash up data from other sources, this capability enables malicious users to mount; XSS attacks.; If the risk outweighs the benefit, install additional markup formatter plugins and use; them.
	Mode                           *Mode                        `json:"mode,omitempty"`
	MyViewsTabBar                  *PurpleMyViewsTabBar         `json:"myViewsTabBar,omitempty"`
	NodeName                       *string                      `json:"nodeName,omitempty"`
	NodeProperties                 []PurpleNodeProperty         `json:"nodeProperties,omitempty"`
	Nodes                          []NodeElement                `json:"nodes,omitempty"`
	NoUsageStatistics              *bool                        `json:"noUsageStatistics,omitempty"`
	NumExecutors                   *int64                       `json:"numExecutors,omitempty"`
	PrimaryView                    *PrimaryView                 `json:"primaryView,omitempty"`
	ProjectNamingStrategy          *PurpleProjectNamingStrategy `json:"projectNamingStrategy,omitempty"`
	Proxy                          *PurpleProxy                 `json:"proxy,omitempty"`
	QuietPeriod                    *int64                       `json:"quietPeriod,omitempty"`
	RemotingSecurity               *PurpleRemotingSecurity      `json:"remotingSecurity,omitempty"`
	SCMCheckoutRetryCount          *int64                       `json:"scmCheckoutRetryCount,omitempty"`
	SecurityRealm                  *PurpleSecurityRealm         `json:"securityRealm,omitempty"`
	SlaveAgentPort                 *int64                       `json:"slaveAgentPort,omitempty"`
	SystemMessage                  *string                      `json:"systemMessage,omitempty"`
	UpdateCenter                   *PurpleUpdateCenter          `json:"updateCenter,omitempty"`
	Views                          []ViewElement                `json:"views,omitempty"`
	ViewsTabBar                    *PurpleViewsTabBar           `json:"viewsTabBar,omitempty"`
}

type FluffyAuthorizationStrategy struct {
}

type Cloud struct {
}

type FluffyCrumbIssuer struct {
}

type GlobalNodeProperty struct {
}

type LabelAtomElement struct {
	Name       *string              `json:"name,omitempty"`
	Properties *LabelAtomProperties `json:"properties,omitempty"`
}

type LabelAtomProperties struct {
}

type PurpleLog struct {
}

// In such places as project description, user description, view description, and build
// description,
// Jenkins allows users to enter some free-form text that describes something.
//
// This configuration determines how such free-form text is converted to HTML. By default,
// Jenkins treats
// the text as HTML and use it as-is unmodified (and this is default mainly because of the
// backward compatibility.)
//
//
// While this is convenient and people often use it to load &lt;iframe>, &lt;script>. and so
// on to
// mash up data from other sources, this capability enables malicious users to mount
// XSS attacks.
// If the risk outweighs the benefit, install additional markup formatter plugins and use
// them.
type PurpleMarkupFormatter struct {
}

type PurpleMyViewsTabBar struct {
}

type PurpleNodeProperty struct {
}

type NodeElement struct {
}

type PrimaryView struct {
}

type PurpleProjectNamingStrategy struct {
}

type PurpleProxy struct {
}

type PurpleRemotingSecurity struct {
}

type PurpleSecurityRealm struct {
}

type PurpleUpdateCenter struct {
}

type ViewElement struct {
}

type PurpleViewsTabBar struct {
}

type Jnlp struct {
	Tunnel          *string          `json:"tunnel,omitempty"`
	Vmargs          *string          `json:"vmargs,omitempty"`    // If the agent JVM should be launched with additional VM arguments, such as "-Xmx256m",; specify those here. List of all the options are available; here.
	WebSocket       *bool            `json:"webSocket,omitempty"` // Use WebSocket to connect to the Jenkins master rather than the TCP port.; See JEP-222 for background.
	WorkDirSettings *WorkDirSettings `json:"workDirSettings,omitempty"`
}

type WorkDirSettings struct {
}

type JobName struct {
}

type JenkinsKeyStoreSource struct {
	Uploaded interface{} `json:"uploaded"`
}

type KnownHostsFileKeyVerificationStrategy struct {
}

type PurpleLabelAtom struct {
	Name       *string             `json:"name,omitempty"`
	Properties []LabelAtomProperty `json:"properties,omitempty"`
}

type LabelAtomProperty struct {
}

type LabelAtoms struct {
}

type LastDuration struct {
}

type LastFailure struct {
}

type LastStable struct {
}

type LastSuccess struct {
}

type Legacy struct {
	CAPTCHASupport *LegacyCAPTCHASupport `json:"captchaSupport,omitempty"`
}

type LegacyCAPTCHASupport struct {
}

type List struct {
	Columns      []Column       `json:"columns,omitempty"`
	Description  *string        `json:"description,omitempty"`
	IncludeRegex *string        `json:"includeRegex,omitempty"`
	JobFilters   []JobFilter    `json:"jobFilters,omitempty"`
	JobNames     *string        `json:"jobNames,omitempty"`
	Name         *string        `json:"name,omitempty"`
	Properties   []ListProperty `json:"properties,omitempty"`
	Recurse      *bool          `json:"recurse,omitempty"`
}

type Column struct {
}

type JobFilter struct {
}

type ListProperty struct {
}

type ListViewColumn struct {
	JobName                  interface{} `json:"jobName"`
	LastDuration             interface{} `json:"lastDuration"`
	BuildButton              interface{} `json:"buildButton"`
	GitBranchSpecifierColumn interface{} `json:"gitBranchSpecifierColumn"`
	LastStable               interface{} `json:"lastStable"`
	Weather                  interface{} `json:"weather"`
	BranchStatusColumn       interface{} `json:"branchStatusColumn"`
	ItemColumn               interface{} `json:"itemColumn"`
	LastSuccess              interface{} `json:"lastSuccess"`
	LastFailure              interface{} `json:"lastFailure"`
	DescriptionColumn        interface{} `json:"descriptionColumn"`
	Status                   interface{} `json:"status"`
}

type Local struct {
	AllowsSignup   *bool                `json:"allowsSignup,omitempty"` // This option allows users to create accounts by themselves, via the "sign up" link on the; top right shoulder of the page.; Make sure to not grant significant permissions to authenticated users, as anyone on the; network will be able to get them.; ; When this checkbox is unchecked, someone with the administrator role would have to create; accounts.; ; By default, Jenkins does not use captcha verification if the user creates an account by; themself.; If you'd like to enable captcha verification, install a captcha support plugin, e.g. the; Jenkins; JCaptcha Plugin.
	CAPTCHASupport *LocalCAPTCHASupport `json:"captchaSupport,omitempty"`
	EnableCAPTCHA  *bool                `json:"enableCaptcha,omitempty"`
	Users          []UserElement        `json:"users,omitempty"`
}

type LocalCAPTCHASupport struct {
}

type UserElement struct {
	Description *string         `json:"description,omitempty"`
	ID          *string         `json:"id,omitempty"`
	Name        *string         `json:"name,omitempty"`
	Password    *string         `json:"password,omitempty"`
	Properties  *UserProperties `json:"properties,omitempty"`
}

type UserProperties struct {
}

type FluffyLog struct {
}

type LogRecorder struct {
	Loggers []Logger `json:"loggers,omitempty"`
	Name    *string  `json:"name,omitempty"`
}

type Logger struct {
	Level *string `json:"level,omitempty"`
	Name  *string `json:"name,omitempty"`
}

type LoggedInUsersCanDoAnything struct {
	AllowAnonymousRead *bool `json:"allowAnonymousRead,omitempty"` // If checked, this will allow users who are not authenticated to access Jenkins in a; read-only mode.
}

type JenkinsMailer struct {
	EmailAddress *string `json:"emailAddress,omitempty"`
}

type ManuallyProvidedKeyVerificationStrategy struct {
	Key *string `json:"key,omitempty"` // The SSH key expected for this connection. This key should be in the form `algorithm; value` where algorithm is one of ssh-rsa or ssh-dss, and value is the Base 64 encoded; content of the key.
}

type ManuallyTrustedKeyVerificationStrategy struct {
	RequireInitialManualTrust *bool `json:"requireInitialManualTrust,omitempty"` // Require a user with Computer.CONFIGURE permission to authorise the key presented during; the first connection to this host before the connection will be allowed to be; established.; If this option is not enabled then the key presented on first connection for this host; will be automatically trusted and allowed for all subsequent connections without any; manual intervention.
}

type FluffyMarkupFormatter struct {
	RawHTML   interface{} `json:"rawHtml"`
	PlainText interface{} `json:"plainText"`
}

type JenkinsMaven struct {
	ID *string `json:"id,omitempty"`
}

type JenkinsMyView struct {
	PrimaryViewName *string `json:"primaryViewName,omitempty"`
}

type FluffyMyViewsTabBar struct {
	Standard interface{} `json:"standard"`
}

type PurpleNode struct {
	Permanent interface{} `json:"permanent"`
	Jenkins   interface{} `json:"jenkins"`
}

type FluffyNodeProperty struct {
	DisableDeferredWipeout interface{} `json:"disableDeferredWipeout"`
	ToolLocation           interface{} `json:"toolLocation"`
	EnvVars                interface{} `json:"envVars"`
	AuthorizationMatrix    interface{} `json:"authorizationMatrix"`
}

type NonInheriting struct {
}

type NonVerifyingKeyVerificationStrategy struct {
}

type Pam struct {
	CAPTCHASupport *PamCAPTCHASupport `json:"captchaSupport,omitempty"`
	ServiceName    *string            `json:"serviceName,omitempty"`
}

type PamCAPTCHASupport struct {
}

type JenkinsPathSpecification struct {
	CaseSensitive *bool   `json:"caseSensitive,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2014, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; Select this option to require that the paths match taking character case into account.
	Excludes      *string `json:"excludes,omitempty"`      // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2014, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of excluded paths. (ANT style * and ** wildcards are permitted in; paths,; for example: /jenkins/github/*,/jenkins-ci/**/org,jenkinsci.github.io.); ; The empty list implies no paths are excluded. The excludes list is processed after the; includes list.
	Includes      *string `json:"includes,omitempty"`      // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2014, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of included paths. (ANT style * and ** wildcards are permitted in; paths,; for example: /jenkins/github/*,/jenkins-ci/**/org,jenkinsci.github.io.); ; The empty list implies no paths are excluded. The excludes list is processed after the; includes list.
}

type Pattern struct {
	Description       *string `json:"description,omitempty"` // Provide a human-readable description to explain naming constraints.; This will be used as the error message when the job name does not match the pattern.
	ForceExistingJobs *bool   `json:"forceExistingJobs,omitempty"`
	NamePattern       *string `json:"namePattern,omitempty"`
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

type PermanentNodeProperty struct {
}

type PermanentRetentionStrategy struct {
}

type PlainText struct {
}

type PreferredProvider struct {
	ProviderID *string `json:"providerId,omitempty"` // Allows the user to select their preferred user interface when clicking links to Jenkins; from notifications (e.g. Email, Slack, or GitHub).
}

type JenkinsPrivateKeySource struct {
	DirectEntry interface{} `json:"directEntry"`
}

type FluffyProjectNamingStrategy struct {
	Standard interface{} `json:"standard"`
	Pattern  interface{} `json:"pattern"`
}

type FluffyProxy struct {
	Name           *string         `json:"name,omitempty"`        // If your Jenkins server sits behind a firewall and does not have the direct access to the; internet,; and if your server JVM is not configured appropriately; (See JDK networking properties for more details); to enable internet connection, you can specify the HTTP proxy server name in this field; to allow Jenkins; to install plugins on behalf of you. Note that Jenkins uses HTTPS to communicate with the; update center to download plugins.; ; ; Leaving this field empty; means Jenkins will try to connect to the internet directly.; ; ; If you are unsure about the value, check the browser proxy configuration.
	NoProxyHost    *string         `json:"noProxyHost,omitempty"` // Specify host name patterns that shouldn't go through the proxy, one host per line.; "*" is the wild card host name (such as "*.jenkins.io" or "www*.jenkins-ci.org")
	Port           *int64          `json:"port,omitempty"`        // This field works in conjunction with the proxy server field to specify the HTTP proxy; port.
	SecretPassword *SecretPassword `json:"secretPassword,omitempty"`
	TestURL        *string         `json:"testUrl,omitempty"`
	UserName       *string         `json:"userName,omitempty"` // This field works in conjunction with the proxy server field; to specify the username used to authenticate with the proxy.; ; ; If this proxy requires Microsofts NTLM; authentication scheme then the domain name can be encoded; within the username by prefixing the domain name followed; by a back-slash '\' before the username, e.g "ACME\John Doo".
}

type SecretPassword struct {
}

type RawHTML struct {
	DisableSyntaxHighlighting *bool `json:"disableSyntaxHighlighting,omitempty"`
}

type FluffyRemotingSecurity struct {
}

type JenkinsRetentionStrategy struct {
	Always   interface{} `json:"always"`
	Schedule interface{} `json:"schedule"`
	Demand   interface{} `json:"demand"`
}

type SSH struct {
	CredentialsID                  *string                            `json:"credentialsId,omitempty"`
	Host                           *string                            `json:"host,omitempty"`
	JavaPath                       *string                            `json:"javaPath,omitempty"`
	JVMOptions                     *string                            `json:"jvmOptions,omitempty"`
	LaunchTimeoutSeconds           *int64                             `json:"launchTimeoutSeconds,omitempty"`
	MaxNumRetries                  *int64                             `json:"maxNumRetries,omitempty"`
	Port                           *int64                             `json:"port,omitempty"`
	PrefixStartSlaveCmd            *string                            `json:"prefixStartSlaveCmd,omitempty"`
	RetryWaitTime                  *int64                             `json:"retryWaitTime,omitempty"`
	SSHHostKeyVerificationStrategy *SSHSSHHostKeyVerificationStrategy `json:"sshHostKeyVerificationStrategy,omitempty"`
	SuffixStartSlaveCmd            *string                            `json:"suffixStartSlaveCmd,omitempty"`
	TCPNoDelay                     *bool                              `json:"tcpNoDelay,omitempty"`
	WorkDir                        *string                            `json:"workDir,omitempty"`
}

type SSHSSHHostKeyVerificationStrategy struct {
}

type JenkinsSSHHostKeyVerificationStrategy struct {
	ManuallyTrustedKeyVerificationStrategy  interface{} `json:"manuallyTrustedKeyVerificationStrategy"`
	ManuallyProvidedKeyVerificationStrategy interface{} `json:"manuallyProvidedKeyVerificationStrategy"`
	NonVerifyingKeyVerificationStrategy     interface{} `json:"nonVerifyingKeyVerificationStrategy"`
	KnownHostsFileKeyVerificationStrategy   interface{} `json:"knownHostsFileKeyVerificationStrategy"`
}

type SSHPublicKey struct {
	AuthorizedKeys *string `json:"authorizedKeys,omitempty"` // List SSH public keys that are associated with the user account.; These keys can be used for example by Jenkins CLI.
}

type Schedule struct {
	KeepUpWhenActive *bool   `json:"keepUpWhenActive,omitempty"`
	StartTimeSpec    *string `json:"startTimeSpec,omitempty"`
	UpTimeMins       *int64  `json:"upTimeMins,omitempty"`
}

type JenkinsSchemeSpecification struct {
	Schemes *string `json:"schemes,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2013, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; A comma separated list of URI schemes (as defined in RFC-3986).; For example: http, https, ssh, sftp, imap.
}

type FluffySecurityRealm struct {
	Legacy interface{} `json:"legacy"`
	None   interface{} `json:"none"`
	Pam    interface{} `json:"pam"`
	Local  interface{} `json:"local"`
}

type JenkinsStandard struct {
}

type Status struct {
}

type StatusFilter struct {
	StatusFilter *bool `json:"statusFilter,omitempty"`
}

type JenkinsString struct {
	Description *string       `json:"description,omitempty"`
	ID          *string       `json:"id,omitempty"`
	Scope       *Scope        `json:"scope,omitempty"`
	Secret      *FluffySecret `json:"secret,omitempty"`
}

type FluffySecret struct {
}

type Target struct {
	Level *string `json:"level,omitempty"`
	Name  *string `json:"name,omitempty"`
}

type Timezone struct {
	TimeZoneName *string `json:"timeZoneName,omitempty"` // Specify user defined time zone for displaying time rather than system default.
}

type JenkinsToolInstaller struct {
	Zip          interface{} `json:"zip"`
	BatchFile    interface{} `json:"batchFile"`
	JDKInstaller interface{} `json:"jdkInstaller"`
	Maven        interface{} `json:"maven"`
	Command      interface{} `json:"command"`
}

type ToolLocation struct {
	Home *string `json:"home,omitempty"`
	Key  *string `json:"key,omitempty"`
}

type JenkinsToolProperty struct {
	InstallSource interface{} `json:"installSource"`
}

type Unsecured struct {
}

type FluffyUpdateCenter struct {
}

type UpdateSite struct {
	ID  *string `json:"id,omitempty"`
	URL *string `json:"url,omitempty"`
}

type JenkinsUploaded struct {
	UploadedKeystore *FluffyUploadedKeystore `json:"uploadedKeystore,omitempty"`
}

type FluffyUploadedKeystore struct {
}

type UserProperty struct {
	LastGrantedAuthorities  interface{} `json:"lastGrantedAuthorities"`
	SSHPublicKey            interface{} `json:"sshPublicKey"`
	UserSeed                interface{} `json:"userSeed"`
	Search                  interface{} `json:"search"`
	PaneStatus              interface{} `json:"paneStatus"`
	Password                interface{} `json:"password"`
	MyView                  interface{} `json:"myView"`
	APIToken                interface{} `json:"apiToken"`
	Timezone                interface{} `json:"timezone"`
	Mailer                  interface{} `json:"mailer"`
	UserCredentialsProperty interface{} `json:"userCredentialsProperty"`
	PreferredProvider       interface{} `json:"preferredProvider"`
}

type UserWithPassword struct {
	Description *string                    `json:"description,omitempty"`
	ID          *string                    `json:"id,omitempty"`
	Name        *string                    `json:"name,omitempty"`
	Password    *string                    `json:"password,omitempty"`
	Properties  []UserWithPasswordProperty `json:"properties,omitempty"`
}

type UserWithPasswordProperty struct {
}

type JenkinsUsernamePassword struct {
	Description    *string `json:"description,omitempty"`
	ID             *string `json:"id,omitempty"`
	Password       *string `json:"password,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2012, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; The password.
	Scope          *Scope  `json:"scope,omitempty"`
	Username       *string `json:"username,omitempty"`       // <!--; ~ The MIT License; ~; ~ Copyright (c) 2011-2012, CloudBees, Inc., Stephen Connolly.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; ; The username.
	UsernameSecret *bool   `json:"usernameSecret,omitempty"` // Historically, not only passwords but usernames were masked in the build log.; Since these can interfere with diagnostics,; and cause unrelated occurrences of a common word to be masked,; you may choose to leave usernames unmasked if they are not sensitive.; Note that regardless of this setting, the username will be displayed to anyone permitted; to reconfigure the credentials.
}

type PurpleView struct {
	Empty  interface{} `json:""`
	All    interface{} `json:"all"`
	Proxy  interface{} `json:"proxy"`
	MyView interface{} `json:"myView"`
	List   interface{} `json:"list"`
}

type ViewJobFilter struct {
	StatusFilter interface{} `json:"statusFilter"`
}

type FluffyViewsTabBar struct {
	Standard interface{} `json:"standard"`
}

type Weather struct {
}

type JenkinsZip struct {
	Label  *string `json:"label,omitempty"`
	Subdir *string `json:"subdir,omitempty"` // Optional subdirectory of the downloaded and unpacked archive to use as the tool's home; directory.
	URL    *string `json:"url,omitempty"`    // URL from which to download the tool in binary form.; Should be either a ZIP or a GZip-compressed TAR file.; The timestamp on the server will be compared to the local version (if any); so you can publish updates easily.; The URL must be accessible from the Jenkins controller but need not be accessible from; agents.
}

type ConfigurationBaseForTheJobsClassifier struct {
	Empty        *Jobs         `json:",omitempty"`
	ScriptSource *ScriptSource `json:"scriptSource,omitempty"`
}

type Jobs struct {
}

type ScriptSource struct {
	File   interface{} `json:"file"`
	Script interface{} `json:"script"`
	URL    interface{} `json:"url"`
}

type ConfigurationBaseForTheSecurityClassifier struct {
	APIToken                          *SecurityAPIToken                  `json:"apiToken,omitempty"`
	Crumb                             *Crumb                             `json:"crumb,omitempty"`
	GlobalJobDSLSecurityConfiguration *GlobalJobDSLSecurityConfiguration `json:"globalJobDslSecurityConfiguration,omitempty"`
	QueueItemAuthenticator            interface{}                        `json:"queueItemAuthenticator"`
	ScriptApproval                    *ScriptApproval                    `json:"scriptApproval,omitempty"`
	SSHD                              *SSHD                              `json:"sSHD,omitempty"`
	UpdateSiteWarningsConfiguration   *UpdateSiteWarningsConfiguration   `json:"updateSiteWarningsConfiguration,omitempty"`
}

type SecurityAPIToken struct {
	CreationOfLegacyTokenEnabled     *bool `json:"creationOfLegacyTokenEnabled,omitempty"`     // This option allows users to generate a legacy API token if they do not already have; one.; Because legacy tokens are deprecated, we recommend disabling it and having users instead; generate; new API tokens from the user configuration page.
	TokenGenerationOnCreationEnabled *bool `json:"tokenGenerationOnCreationEnabled,omitempty"` // This option causes a legacy API token to be generated automatically for newly created; users.; Because legacy tokens are deprecated, we recommend disabling it and having users instead; generate; new API tokens from the user configuration page as needed.
	UsageStatisticsEnabled           *bool `json:"usageStatisticsEnabled,omitempty"`           // If this option is enabled, then the date of the most recent use of each API token and the; total number of times; it has been used are stored in Jenkins.; This allows users to see if they have unused or outdated API tokens which should be; revoked.; ; ; This data is stored in your Jenkins instance and will not be used for any other purpose.
}

type Crumb struct {
}

type GlobalJobDSLSecurityConfiguration struct {
	UseScriptSecurity *bool `json:"useScriptSecurity,omitempty"`
}

type SSHD struct {
	Port *int64 `json:"port,omitempty"` // Jenkins can act as an SSH server to run a subset of Jenkins CLI commands.; Plugins may also use this service to expose additional commands. Specify this TCP/IP port; number.
}

type ScriptApproval struct {
}

type UpdateSiteWarningsConfiguration struct {
}

type ConfigurationBaseForTheToolClassifier struct {
	BatchFile              *ToolBatchFile              `json:"batchFile,omitempty"`
	Command                *ToolCommand                `json:"command,omitempty"`
	FilePath               *FilePath                   `json:"filePath,omitempty"`
	Git                    *ToolGit                    `json:"git,omitempty"`
	GlobalSettingsProvider *ToolGlobalSettingsProvider `json:"globalSettingsProvider,omitempty"`
	JDK                    *ToolJDK                    `json:"jdk,omitempty"`
	JDKInstaller           *ToolJDKInstaller           `json:"jdkInstaller,omitempty"`
	Jgit                   *Jgit                       `json:"jgit,omitempty"`
	Jgitapache             *Jgitapache                 `json:"jgitapache,omitempty"`
	Maven                  *ToolMaven                  `json:"maven,omitempty"`
	MavenGlobalConfig      *MavenGlobalConfig          `json:"mavenGlobalConfig,omitempty"`
	SettingsProvider       *ToolSettingsProvider       `json:"settingsProvider,omitempty"`
	Standard               *ToolStandard               `json:"standard,omitempty"`
	ToolInstallation       *ToolInstallation           `json:"toolInstallation,omitempty"`
	ToolInstaller          *ToolToolInstaller          `json:"toolInstaller,omitempty"`
	ToolProperty           *ToolToolProperty           `json:"toolProperty,omitempty"`
	Zip                    *ToolZip                    `json:"zip,omitempty"`
}

type ToolBatchFile struct {
	Command  *string `json:"command,omitempty"`
	Label    *string `json:"label,omitempty"`
	ToolHome *string `json:"toolHome,omitempty"`
}

type ToolCommand struct {
	Command  *string `json:"command,omitempty"`
	Label    *string `json:"label,omitempty"`
	ToolHome *string `json:"toolHome,omitempty"`
}

type FilePath struct {
	Path *string `json:"path,omitempty"` // Path to settings.xml file, relative to project workspace or absolute (variables are; supported).
}

type ToolGit struct {
	Home       *string       `json:"home,omitempty"` // Set path to git executable. This can be just "git" or complete path.
	Name       *string       `json:"name,omitempty"`
	Properties []GitProperty `json:"properties,omitempty"`
}

type GitProperty struct {
}

type ToolGlobalSettingsProvider struct {
	Standard interface{} `json:"standard"`
	FilePath interface{} `json:"filePath"`
}

type ToolJDK struct {
	Home       *string          `json:"home,omitempty"`
	Name       *string          `json:"name,omitempty"`
	Properties []FluffyProperty `json:"properties,omitempty"`
}

type FluffyProperty struct {
}

type ToolJDKInstaller struct {
	AcceptLicense *bool   `json:"acceptLicense,omitempty"`
	ID            *string `json:"id,omitempty"`
}

type Jgit struct {
}

type Jgitapache struct {
}

type ToolMaven struct {
	Home       *string         `json:"home,omitempty"`
	Name       *string         `json:"name,omitempty"`
	Properties []MavenProperty `json:"properties,omitempty"`
}

type MavenProperty struct {
}

type MavenGlobalConfig struct {
	GlobalSettingsProvider *MavenGlobalConfigGlobalSettingsProvider `json:"globalSettingsProvider,omitempty"`
	Installations          []Installation                           `json:"installations,omitempty"`
	SettingsProvider       *MavenGlobalConfigSettingsProvider       `json:"settingsProvider,omitempty"`
}

type MavenGlobalConfigGlobalSettingsProvider struct {
}

type Installation struct {
	Home       *string                 `json:"home,omitempty"`
	Name       *string                 `json:"name,omitempty"`
	Properties *InstallationProperties `json:"properties,omitempty"`
}

type InstallationProperties struct {
}

type MavenGlobalConfigSettingsProvider struct {
}

type ToolSettingsProvider struct {
	Standard interface{} `json:"standard"`
	FilePath interface{} `json:"filePath"`
}

type ToolStandard struct {
}

type ToolInstallation struct {
	JDK        interface{} `json:"jdk"`
	Git        interface{} `json:"git"`
	Maven      interface{} `json:"maven"`
	Jgit       interface{} `json:"jgit"`
	Jgitapache interface{} `json:"jgitapache"`
}

type ToolToolInstaller struct {
	Zip          interface{} `json:"zip"`
	BatchFile    interface{} `json:"batchFile"`
	JDKInstaller interface{} `json:"jdkInstaller"`
	Maven        interface{} `json:"maven"`
	Command      interface{} `json:"command"`
}

type ToolToolProperty struct {
	InstallSource interface{} `json:"installSource"`
}

type ToolZip struct {
	Label  *string `json:"label,omitempty"`
	Subdir *string `json:"subdir,omitempty"` // Optional subdirectory of the downloaded and unpacked archive to use as the tool's home; directory.
	URL    *string `json:"url,omitempty"`    // URL from which to download the tool in binary form.; Should be either a ZIP or a GZip-compressed TAR file.; The timestamp on the server will be compared to the local version (if any); so you can publish updates easily.; The URL must be accessible from the Jenkins controller but need not be accessible from; agents.
}

type ConfigurationBaseForTheUnclassifiedClassifier struct {
	Empty                               *Unclassified                        `json:",omitempty"`
	AdministrativeMonitorsConfiguration *AdministrativeMonitorsConfiguration `json:"administrativeMonitorsConfiguration,omitempty"`
	Ancestry                            *Ancestry                            `json:"ancestry,omitempty"`
	ArtifactManager                     *ArtifactManager                     `json:"artifactManager,omitempty"`
	ArtifactManagerFactory              interface{}                          `json:"artifactManagerFactory"`
	AssemblaWeb                         *AssemblaWeb                         `json:"assemblaWeb,omitempty"`
	AuthorInChangelog                   *AuthorInChangelog                   `json:"authorInChangelog,omitempty"`
	BitbucketServer                     *BitbucketServer                     `json:"bitbucketServer,omitempty"`
	BitbucketWeb                        *BitbucketWeb                        `json:"bitbucketWeb,omitempty"`
	BuildChooser                        *UnclassifiedBuildChooser            `json:"buildChooser,omitempty"`
	BuildChooserSetting                 *BuildChooserSetting                 `json:"buildChooserSetting,omitempty"`
	BuildDiscarder                      *BuildDiscarder                      `json:"buildDiscarder,omitempty"`
	BuildDiscarders                     *BuildDiscarders                     `json:"buildDiscarders,omitempty"`
	BuildSingleRevisionOnly             *BuildSingleRevisionOnly             `json:"buildSingleRevisionOnly,omitempty"`
	BuildStepOperation                  *BuildStepOperation                  `json:"buildStepOperation,omitempty"`
	BuiltInNode                         *BuiltInNode                         `json:"builtInNode,omitempty"`
	CasCGlobalConfig                    *CasCGlobalConfig                    `json:"casCGlobalConfig,omitempty"`
	CGit                                *CGit                                `json:"cGit,omitempty"`
	ChangelogToBranch                   *ChangelogToBranch                   `json:"changelogToBranch,omitempty"`
	CheckoutOption                      *CheckoutOption                      `json:"checkoutOption,omitempty"`
	CleanBeforeCheckout                 *CleanBeforeCheckout                 `json:"cleanBeforeCheckout,omitempty"`
	CleanCheckout                       *CleanCheckout                       `json:"cleanCheckout,omitempty"`
	CloneOption                         *CloneOption                         `json:"cloneOption,omitempty"`
	Default                             *Default                             `json:"default,omitempty"`
	DefaultFolderConfiguration          *DefaultFolderConfiguration          `json:"defaultFolderConfiguration,omitempty"`
	DefaultView                         *DefaultView                         `json:"defaultView,omitempty"`
	DisableRemotePoll                   *DisableRemotePoll                   `json:"disableRemotePoll,omitempty"`
	EnvVarsFilter                       *EnvVarsFilter                       `json:"envVarsFilter,omitempty"`
	File                                *UnclassifiedFile                    `json:"file,omitempty"`
	Fingerprints                        *Fingerprints                        `json:"fingerprints,omitempty"`
	FingerprintStorage                  *FingerprintStorage                  `json:"fingerprintStorage,omitempty"`
	FisheyeGit                          *FisheyeGit                          `json:"fisheyeGit,omitempty"`
	FolderHealthMetric                  *FolderHealthMetric                  `json:"folderHealthMetric,omitempty"`
	FromSCM                             *FromSCM                             `json:"fromScm,omitempty"`
	Git                                 *UnclassifiedGit                     `json:"git,omitempty"`
	GitBlit                             *GitBlit                             `json:"gitBlit,omitempty"`
	GitBranchDiscovery                  *GitBranchDiscovery                  `json:"gitBranchDiscovery,omitempty"`
	GithubWeb                           *GithubWeb                           `json:"githubWeb,omitempty"`
	Gitiles                             *Gitiles                             `json:"gitiles,omitempty"`
	GitLab                              *GitLab                              `json:"gitLab,omitempty"`
	GitLFSPull                          *GitLFSPull                          `json:"gitLFSPull,omitempty"`
	GitList                             *GitList                             `json:"gitList,omitempty"`
	GitoriousWeb                        *GitoriousWeb                        `json:"gitoriousWeb,omitempty"`
	GitRepositoryBrowser                *GitRepositoryBrowser                `json:"gitRepositoryBrowser,omitempty"`
	GitSCM                              *GitSCM                              `json:"gitSCM,omitempty"`
	GitSCMExtension                     *GitSCMExtension                     `json:"gitSCMExtension,omitempty"`
	GitTagDiscovery                     *GitTagDiscovery                     `json:"gitTagDiscovery,omitempty"`
	GitWeb                              *GitWeb                              `json:"gitWeb,omitempty"`
	GlobalBuildDiscarderStrategy        *GlobalBuildDiscarderStrategy        `json:"globalBuildDiscarderStrategy,omitempty"`
	GlobalDefaultFlowDurabilityLevel    *GlobalDefaultFlowDurabilityLevel    `json:"globalDefaultFlowDurabilityLevel,omitempty"`
	GlobalLibraries                     *GlobalLibraries                     `json:"globalLibraries,omitempty"`
	GogsGit                             *GogsGit                             `json:"gogsGit,omitempty"`
	HeadRegexFilter                     *HeadRegexFilter                     `json:"headRegexFilter,omitempty"`
	HeadWildcardFilter                  *HeadWildcardFilter                  `json:"headWildcardFilter,omitempty"`
	IgnoreNotifyCommit                  *IgnoreNotifyCommit                  `json:"ignoreNotifyCommit,omitempty"`
	Inverse                             *Inverse                             `json:"inverse,omitempty"`
	JobBuildDiscarder                   *JobBuildDiscarder                   `json:"jobBuildDiscarder,omitempty"`
	JunitTestResultStorage              *JunitTestResultStorage              `json:"junitTestResultStorage,omitempty"`
	KilnGit                             *KilnGit                             `json:"kilnGit,omitempty"`
	LegacySCM                           *LegacySCM                           `json:"legacySCM,omitempty"`
	LibraryRetriever                    *LibraryRetriever                    `json:"libraryRetriever,omitempty"`
	LocalBranch                         *LocalBranch                         `json:"localBranch,omitempty"`
	Location                            *Location                            `json:"location,omitempty"`
	LockableResourcesManager            *LockableResourcesManager            `json:"lockableResourcesManager,omitempty"`
	LogRotator                          *LogRotator                          `json:"logRotator,omitempty"`
	Mailer                              *UnclassifiedMailer                  `json:"mailer,omitempty"`
	MessageExclusion                    *MessageExclusion                    `json:"messageExclusion,omitempty"`
	ModernSCM                           *ModernSCM                           `json:"modernSCM,omitempty"`
	MyView                              *UnclassifiedMyView                  `json:"myView,omitempty"`
	NodeProperties                      *NodeProperties                      `json:"nodeProperties,omitempty"`
	None                                *UnclassifiedNone                    `json:"none,omitempty"`
	PathRestriction                     *PathRestriction                     `json:"pathRestriction,omitempty"`
	PerBuildTag                         *PerBuildTag                         `json:"perBuildTag,omitempty"`
	Phabricator                         *Phabricator                         `json:"phabricator,omitempty"`
	Plugin                              *Plugin                              `json:"plugin,omitempty"`
	PollSCM                             *PollSCM                             `json:"pollSCM,omitempty"`
	PreBuildMerge                       *PreBuildMerge                       `json:"preBuildMerge,omitempty"`
	PrimaryBranchHealthMetric           *PrimaryBranchHealthMetric           `json:"primaryBranchHealthMetric,omitempty"`
	ProjectNamingStrategy               *UnclassifiedProjectNamingStrategy   `json:"projectNamingStrategy,omitempty"`
	PruneStaleBranch                    *PruneStaleBranch                    `json:"pruneStaleBranch,omitempty"`
	PruneTags                           *PruneTags                           `json:"pruneTags,omitempty"`
	QuietPeriod                         *QuietPeriod                         `json:"quietPeriod,omitempty"`
	RedmineWeb                          *RedmineWeb                          `json:"redmineWeb,omitempty"`
	RelativeTargetDirectory             *RelativeTargetDirectory             `json:"relativeTargetDirectory,omitempty"`
	ResourceRoot                        *ResourceRoot                        `json:"resourceRoot,omitempty"`
	RhodeCode                           *RhodeCode                           `json:"rhodeCode,omitempty"`
	SCM                                 *UnclassifiedSCM                     `json:"scm,omitempty"`
	SCMName                             *SCMName                             `json:"scmName,omitempty"`
	SCMRetryCount                       *SCMRetryCount                       `json:"scmRetryCount,omitempty"`
	SCMSource                           *SCMSource                           `json:"sCMSource,omitempty"`
	SCMSourceTrait                      *SCMSourceTrait                      `json:"sCMSourceTrait,omitempty"`
	Shell                               *Shell                               `json:"shell,omitempty"`
	SimpleBuildDiscarder                *SimpleBuildDiscarder                `json:"simpleBuildDiscarder,omitempty"`
	Stash                               *Stash                               `json:"stash,omitempty"`
	SubmoduleConfig                     *SubmoduleConfig                     `json:"submoduleConfig,omitempty"`
	SubmoduleOption                     *SubmoduleOption                     `json:"submoduleOption,omitempty"`
	TFS2013Git                          *TFS2013Git                          `json:"tFS2013Git,omitempty"`
	Timestamper                         *Timestamper                         `json:"timestamper,omitempty"`
	UsageStatistics                     *UsageStatistics                     `json:"usageStatistics,omitempty"`
	UserExclusion                       *UserExclusion                       `json:"userExclusion,omitempty"`
	UserIdentity                        *UserIdentity                        `json:"userIdentity,omitempty"`
	ViewGitWeb                          *ViewGitWeb                          `json:"viewGitWeb,omitempty"`
	ViewsTabBar                         *UnclassifiedViewsTabBar             `json:"viewsTabBar,omitempty"`
	WipeWorkspace                       *WipeWorkspace                       `json:"wipeWorkspace,omitempty"`
	WorstChildHealthMetric              *WorstChildHealthMetric              `json:"worstChildHealthMetric,omitempty"`
}

type AdministrativeMonitorsConfiguration struct {
}

type Ancestry struct {
	AncestorCommitSha1 *string `json:"ancestorCommitSha1,omitempty"`
	MaximumAgeInDays   *int64  `json:"maximumAgeInDays,omitempty"`
}

type ArtifactManager struct {
}

type AssemblaWeb struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the root URL serving this repository (such as; https://www.assembla.com/code/PROJECT/git/).
}

type AuthorInChangelog struct {
}

type BitbucketServer struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the Bitbucket Server root URL for this repository (such as; https://bitbucket:7990/OWNER/REPO/).
}

type BitbucketWeb struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the root URL serving this repository (such as https://bitbucket.org/OWNER/REPO/).
}

type UnclassifiedBuildChooser struct {
	Inverse  interface{} `json:"inverse"`
	Default  interface{} `json:"default"`
	Ancestry interface{} `json:"ancestry"`
}

type BuildChooserSetting struct {
	BuildChooser *BuildChooserSettingBuildChooser `json:"buildChooser,omitempty"`
}

type BuildChooserSettingBuildChooser struct {
}

type BuildDiscarder struct {
	LogRotator interface{} `json:"logRotator"`
}

type BuildDiscarders struct {
}

type BuildSingleRevisionOnly struct {
}

type BuildStepOperation struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type BuiltInNode struct {
}

type CGit struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the root URL serving this repository (such as; http://cgit.example.com:port/group/REPO/).
}

type CasCGlobalConfig struct {
	ConfigurationPath *string `json:"configurationPath,omitempty"`
}

type ChangelogToBranch struct {
	Options *ChangelogToBranchOptions `json:"options,omitempty"`
}

type ChangelogToBranchOptions struct {
}

type CheckoutOption struct {
	Timeout *int64 `json:"timeout,omitempty"` // Specify a timeout (in minutes) for checkout.; This option overrides the default timeout of 10 minutes.; You can change the global git timeout via the property; org.jenkinsci.plugins.gitclient.Git.timeOut (see JENKINS-11286).; Note that property should be set on both controller and agent to have effect (see; JENKINS-22547).
}

type CleanBeforeCheckout struct {
	DeleteUntrackedNestedRepositories *bool `json:"deleteUntrackedNestedRepositories,omitempty"` // Deletes untracked submodules and any other subdirectories which contain .git directories.
}

type CleanCheckout struct {
	DeleteUntrackedNestedRepositories *bool `json:"deleteUntrackedNestedRepositories,omitempty"` // Deletes untracked submodules and any other subdirectories which contain .git directories.
}

type CloneOption struct {
	Depth        *int64  `json:"depth,omitempty"`        // Set shallow clone depth, so that git will only download recent history of the project,; saving time and disk space when you just want to access the latest commits of a; repository.
	HonorRefspec *bool   `json:"honorRefspec,omitempty"` // Perform initial clone using the refspec defined for the repository.; This can save time, data transfer and disk space when you only need; to access the references specified by the refspec.
	NoTags       *bool   `json:"noTags,omitempty"`       // Deselect this to perform a clone without tags, saving time and disk space when you just; want to access; what is specified by the refspec.
	Reference    *string `json:"reference,omitempty"`    // Specify a folder containing a repository that will be used by Git as a reference during; clone operations.; This option will be ignored if the folder is not available on the controller or agent; where the clone is being executed.
	Shallow      *bool   `json:"shallow,omitempty"`      // Perform shallow clone, so that git will not download the history of the project,; saving time and disk space when you just want to access the latest version of a; repository.
	Timeout      *int64  `json:"timeout,omitempty"`      // Specify a timeout (in minutes) for clone and fetch operations.; This option overrides the default timeout of 10 minutes.; You can change the global git timeout via the property; org.jenkinsci.plugins.gitclient.Git.timeOut (see JENKINS-11286).; Note that property should be set on both controller and agent to have effect (see; JENKINS-22547).
}

type Default struct {
}

type DefaultFolderConfiguration struct {
}

type DefaultView struct {
}

type DisableRemotePoll struct {
}

type Unclassified struct {
	Password *Password `json:"password,omitempty"`
	Username *string   `json:"username,omitempty"`
}

type Password struct {
}

type EnvVarsFilter struct {
}

type UnclassifiedFile struct {
}

type FingerprintStorage struct {
	File interface{} `json:"file"`
}

type Fingerprints struct {
	FingerprintCleanupDisabled *bool    `json:"fingerprintCleanupDisabled,omitempty"`
	Storage                    *Storage `json:"storage,omitempty"`
}

type Storage struct {
}

type FisheyeGit struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the URL of this repository in FishEye (such as; http://fisheye6.cenqua.com/browse/ant/).
}

type FolderHealthMetric struct {
	PrimaryBranchHealthMetric interface{} `json:"primaryBranchHealthMetric"`
	WorstChildHealthMetric    interface{} `json:"worstChildHealthMetric"`
}

type FromSCM struct {
	ID     *string     `json:"id,omitempty"`
	Name   *string     `json:"name,omitempty"` // The name of the SCM head/trunk/branch/tag that this source provides.
	SCM    *FromSCMSCM `json:"scm,omitempty"`
	Traits []Trait     `json:"traits,omitempty"`
}

type FromSCMSCM struct {
}

type Trait struct {
}

type UnclassifiedGit struct {
	Branches                          []Branch           `json:"branches,omitempty"` // List of branches to build.; Jenkins jobs are most effective when each job builds only a single branch.; When a single job builds multiple branches, the changelog comparisons between branches; often show no changes or incorrect changes.
	Browser                           *Browser           `json:"browser,omitempty"`  // Defines the repository browser that displays changes detected by the git plugin.
	BuildChooser                      *GitBuildChooser   `json:"buildChooser,omitempty"`
	DoGenerateSubmoduleConfigurations *bool              `json:"doGenerateSubmoduleConfigurations,omitempty"` // Removed facility that was intended to test combinations of git submodule versions.; Removed in git plugin 4.6.0.; Ignores the user provided value and always uses false as its value.
	Extensions                        []Extension        `json:"extensions,omitempty"`                        // Extensions add new behavior or modify existing plugin behavior for different uses.; Extensions help users more precisely tune plugin behavior to meet their needs.; ; ; Extensions include:; ; ; Clone extensions modify the git operations that retrieve remote changes into the agent; workspace.; The extensions can adjust the amount of history retrieved, how long the retrieval is; allowed to run, and other retrieval details.; ; ; Checkout extensions modify the git operations that place files in the workspace from the; git repository on the agent.; The extensions can adjust the maximum duration of the checkout operation, the use and; behavior of git submodules, the location of the workspace on the disc, and; more.; ; ; Changelog extensions adapt the source code difference calculations for different; cases.; ; ; Tagging extensions allow the plugin to apply tags in the current; workspace.; ; ; Build initiation extensions control the conditions that start a build.; They can ignore notifications of a change or force a deeper evaluation of the commits; when polling.; ; ; Merge extensions can optionally merge changes from other branches into the current branch; of the agent workspace.; They control the source branch for the merge and the options applied to the merge.
	GitTool                           *string            `json:"gitTool,omitempty"`                           // Absolute path to the git executable.; ; ; This is different from other Jenkins tool definitions.; Rather than providing the directory that contains the executable, you must provide the; complete path to the executable.; Setting '/usr/bin/git' would be correct, while setting '/usr/bin/' is not correct.
	SubmoduleCFG                      []SubmoduleCFG     `json:"submoduleCfg,omitempty"`                      // Removed facility that was intended to test combinations of git submodule versions.; Removed in git plugin 4.6.0.; Ignores the user provided value(s) and always uses empty values.
	UserRemoteConfigs                 []UserRemoteConfig `json:"userRemoteConfigs,omitempty"`                 // Specify the repository to track. This can be a URL or a local file path.; Note that for super-projects (repositories with submodules), only a local file; path or a complete URL is valid.  The following are examples of valid git URLs.; ; ssh://git@github.com/github/git.git; git@github.com:github/git.git (short notation for ssh protocol); ssh://user@other.host.com/~/repos/R.git (to access the repos/R.git; repository in the user's home directory); https://github.com/github/git.git; ; ; If the repository is a super-project, the; location from which to clone submodules is dependent on whether the repository; is bare or non-bare (i.e. has a working directory).; ; If the super-project is bare, the location of the submodules will be; taken from .gitmodules.; If the super-project is not bare, it is assumed that the; repository has each of its submodules cloned and checked out appropriately.; Thus, the submodules will be taken directly from a path like; ${SUPER_PROJECT_URL}/${SUBMODULE}, rather than relying on; information from .gitmodules.; ; ; For a local URL/path to a super-project,; git rev-parse --is-bare-repository; is used to detect whether the super-project is bare or not.; ; For a remote URL to a super-project, the ending of the URL determines whether; a bare or non-bare repository is assumed:; ; If the remote URL ends with /.git, a non-bare repository is; assumed.; If the remote URL does NOT end with /.git, a bare; repository is assumed.
}

type Branch struct {
	Name *string `json:"name,omitempty"`
}

// Defines the repository browser that displays changes detected by the git plugin.
type Browser struct {
}

type GitBuildChooser struct {
}

type Extension struct {
}

type SubmoduleCFG struct {
	Branches      *string `json:"branches,omitempty"` // List of branches to build.; Jenkins jobs are most effective when each job builds only a single branch.; When a single job builds multiple branches, the changelog comparisons between branches; often show no changes or incorrect changes.
	SubmoduleName *string `json:"submoduleName,omitempty"`
}

type UserRemoteConfig struct {
	CredentialsID *string `json:"credentialsId,omitempty"`
	Name          *string `json:"name,omitempty"`
	Refspec       *string `json:"refspec,omitempty"`
	URL           *string `json:"url,omitempty"`
}

type GitBlit struct {
	ProjectName *string `json:"projectName,omitempty"` // Specify the name of the project in GitBlit.
	RepoURL     *string `json:"repoUrl,omitempty"`     // Specify the root URL serving this repository.
}

type GitBranchDiscovery struct {
}

type GitLFSPull struct {
}

type GitLab struct {
	Version *string `json:"version,omitempty"` // Specify the major and minor version of GitLab you use (such as 9.1). If you; don't specify a version, a modern version of GitLab (>= 8.0) is assumed.
}

type GitList struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the root URL serving this repository (such as http://gitlistserver:port/REPO/).
}

type GitRepositoryBrowser struct {
	Gitiles                  interface{} `json:"gitiles"`
	CGit                     interface{} `json:"cGit"`
	GitWeb                   interface{} `json:"gitWeb"`
	GitList                  interface{} `json:"gitList"`
	GitBlitRepositoryBrowser interface{} `json:"gitBlitRepositoryBrowser"`
	GitoriousWeb             interface{} `json:"gitoriousWeb"`
	ViewGitWeb               interface{} `json:"viewGitWeb"`
	RedmineWeb               interface{} `json:"redmineWeb"`
	BitbucketWeb             interface{} `json:"bitbucketWeb"`
	Tfs2013                  interface{} `json:"tfs2013"`
	Phabricator              interface{} `json:"phabricator"`
	GogsGit                  interface{} `json:"gogsGit"`
	Fisheye                  interface{} `json:"fisheye"`
	BitbucketServer          interface{} `json:"bitbucketServer"`
	GitLab                   interface{} `json:"gitLab"`
	RhodeCode                interface{} `json:"rhodeCode"`
	GithubWeb                interface{} `json:"githubWeb"`
	AssemblaWeb              interface{} `json:"assemblaWeb"`
	KilnGit                  interface{} `json:"kilnGit"`
	Stash                    interface{} `json:"stash"`
}

type GitSCM struct {
	AddGitTagAction                  *bool   `json:"addGitTagAction,omitempty"`
	AllowSecondFetch                 *bool   `json:"allowSecondFetch,omitempty"`
	CreateAccountBasedOnEmail        *bool   `json:"createAccountBasedOnEmail,omitempty"`
	DisableGitToolChooser            *bool   `json:"disableGitToolChooser,omitempty"`
	GlobalConfigEmail                *string `json:"globalConfigEmail,omitempty"`
	GlobalConfigName                 *string `json:"globalConfigName,omitempty"`
	HideCredentials                  *bool   `json:"hideCredentials,omitempty"`
	ShowEntireCommitSummaryInChanges *bool   `json:"showEntireCommitSummaryInChanges,omitempty"`
	UseExistingAccountWithSameEmail  *bool   `json:"useExistingAccountWithSameEmail,omitempty"`
}

type GitSCMExtension struct {
	BuildChooserSetting     interface{} `json:"buildChooserSetting"`
	PreBuildMerge           interface{} `json:"preBuildMerge"`
	BuildSingleRevisionOnly interface{} `json:"buildSingleRevisionOnly"`
	GitLFSPull              interface{} `json:"gitLFSPull"`
	PathRestriction         interface{} `json:"pathRestriction"`
	SCMName                 interface{} `json:"scmName"`
	AuthorInChangelog       interface{} `json:"authorInChangelog"`
	CheckoutOption          interface{} `json:"checkoutOption"`
	UserIdentity            interface{} `json:"userIdentity"`
	DisableRemotePoll       interface{} `json:"disableRemotePoll"`
	IgnoreNotifyCommit      interface{} `json:"ignoreNotifyCommit"`
	WipeWorkspace           interface{} `json:"wipeWorkspace"`
	UserExclusion           interface{} `json:"userExclusion"`
	PruneStaleBranch        interface{} `json:"pruneStaleBranch"`
	SubmoduleOption         interface{} `json:"submoduleOption"`
	SparseCheckoutPaths     interface{} `json:"sparseCheckoutPaths"`
	PerBuildTag             interface{} `json:"perBuildTag"`
	CleanCheckout           interface{} `json:"cleanCheckout"`
	ChangelogToBranch       interface{} `json:"changelogToBranch"`
	CloneOption             interface{} `json:"cloneOption"`
	RelativeTargetDirectory interface{} `json:"relativeTargetDirectory"`
	CleanBeforeCheckout     interface{} `json:"cleanBeforeCheckout"`
	MessageExclusion        interface{} `json:"messageExclusion"`
	PruneTags               interface{} `json:"pruneTags"`
	LocalBranch             interface{} `json:"localBranch"`
}

type GitTagDiscovery struct {
}

type GitWeb struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the root URL serving this repository (such as; https://github.com/jenkinsci/jenkins.git).
}

type GithubWeb struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the HTTP URL for this repository's GitHub page (such as; https://github.com/jquery/jquery).
}

type Gitiles struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the root URL serving this repository (such as https://gwt.googlesource.com/gwt/).
}

type GitoriousWeb struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the root URL serving this repository (such as; https://gitorious.org/gitorious/mainline).
}

type GlobalBuildDiscarderStrategy struct {
	SimpleBuildDiscarder interface{} `json:"simpleBuildDiscarder"`
	JobBuildDiscarder    interface{} `json:"jobBuildDiscarder"`
}

type GlobalDefaultFlowDurabilityLevel struct {
}

type GlobalLibraries struct {
}

type GogsGit struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the root URL serving this repository (such as; http://gogs.example.com:port/username/some-repo-url.git).
}

type HeadRegexFilter struct {
	Regex *string `json:"regex,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2017, CloudBees, Inc.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; A Java regular expression to; restrict the names. Names that do not match the supplied regular expression will be; ignored.; NOTE: this filter will be applied to all branch like things, including change requests
}

type HeadWildcardFilter struct {
	Excludes *string `json:"excludes,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2017, CloudBees, Inc.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; Space-separated list of name patterns to ignore even if matched by the includes list.; For example: release alpha-* beta-*; NOTE: this filter will be applied to all branch like things, including change requests
	Includes *string `json:"includes,omitempty"` // <!--; ~ The MIT License; ~; ~ Copyright (c) 2017, CloudBees, Inc.; ~; ~ Permission is hereby granted, free of charge, to any person obtaining a copy; ~ of this software and associated documentation files (the "Software"), to deal; ~ in the Software without restriction, including without limitation the rights; ~ to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; ~ copies of the Software, and to permit persons to whom the Software is; ~ furnished to do so, subject to the following conditions:; ~; ~ The above copyright notice and this permission notice shall be included in; ~ all copies or substantial portions of the Software.; ~; ~ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; ~ IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; ~ FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; ~ AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; ~ LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; ~ OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; ~ THE SOFTWARE.; -->; ; Space-separated list of name patterns to consider.; You may use * as a wildcard; for example: master release*; NOTE: this filter will be applied to all branch like things, including change requests
}

type IgnoreNotifyCommit struct {
}

type Inverse struct {
}

type JobBuildDiscarder struct {
}

type JunitTestResultStorage struct {
	File interface{} `json:"file"`
}

type KilnGit struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the root URL serving this repository (such as; https://khanacademy.kilnhg.com/Code/Website/Group/webapp).
}

type LegacySCM struct {
	LibraryPath *string       `json:"libraryPath,omitempty"` // A relative path from the root of the SCM to the root of the library.; Leave this field blank if the root of the library is the root of the SCM.; Note that ".." is not permitted as a path component to avoid security issues.
	SCM         *LegacySCMSCM `json:"scm,omitempty"`
}

type LegacySCMSCM struct {
}

type LibraryRetriever struct {
	ModernSCM interface{} `json:"modernSCM"`
	LegacySCM interface{} `json:"legacySCM"`
}

type LocalBranch struct {
	LocalBranch *string `json:"localBranch,omitempty"`
}

type Location struct {
	AdminAddress *string `json:"adminAddress,omitempty"` // Notification e-mails from Jenkins to project owners will be sent; with this address in the from header. This can be just; "foo@acme.org" or it could be something like "Jenkins Daemon &lt;foo@acme.org>"
	URL          *string `json:"url,omitempty"`          // Optionally specify the HTTP address of the Jenkins installation, such; as http://yourhost.yourdomain/jenkins/. This value is used to; let Jenkins know how to refer to itself, ie. to display images or to; create links in emails.; ; ; This is necessary because Jenkins cannot reliably detect such a URL; from within itself.
}

type LockableResourcesManager struct {
}

type LogRotator struct {
	ArtifactDaysToKeepStr *string `json:"artifactDaysToKeepStr,omitempty"`
	ArtifactNumToKeepStr  *string `json:"artifactNumToKeepStr,omitempty"`
	DaysToKeepStr         *string `json:"daysToKeepStr,omitempty"`
	NumToKeepStr          *string `json:"numToKeepStr,omitempty"`
}

type UnclassifiedMailer struct {
	Authentication *Authentication `json:"authentication,omitempty"`
	Charset        *string         `json:"charset,omitempty"`
	DefaultSuffix  *string         `json:"defaultSuffix,omitempty"`
	ReplyToAddress *string         `json:"replyToAddress,omitempty"`
	SMTPHost       *string         `json:"smtpHost,omitempty"`
	SMTPPort       *string         `json:"smtpPort,omitempty"`
	UseSSL         *bool           `json:"useSsl,omitempty"`
	UseTLS         *bool           `json:"useTls,omitempty"`
}

type Authentication struct {
}

type MessageExclusion struct {
	ExcludedMessage *string `json:"excludedMessage,omitempty"` // If set, and Jenkins is set to poll for changes, Jenkins will ignore any revisions; committed with message matched to; Pattern when determining; if a build needs to be triggered. This can be used to exclude commits done by the build; itself from triggering another build,; assuming the build server commits the change with a distinct message.; Exclusion uses Pattern; matching; ; .*\[maven-release-plugin\].*; The example above illustrates that if only revisions with "[maven-release-plugin]"; message in first comment line; have been committed to the SCM a build will not occur.; ; You can create more complex patterns using embedded flag expressions.; (?s).*FOO.*; This example will search FOO message in all comment lines.
}

type ModernSCM struct {
	LibraryPath *string       `json:"libraryPath,omitempty"` // A relative path from the root of the SCM to the root of the library.; Leave this field blank if the root of the library is the root of the SCM.; Note that ".." is not permitted as a path component to avoid security issues.
	SCM         *ModernSCMSCM `json:"scm,omitempty"`
}

type ModernSCMSCM struct {
}

type UnclassifiedMyView struct {
}

type NodeProperties struct {
}

type UnclassifiedNone struct {
}

type PathRestriction struct {
	ExcludedRegions *string `json:"excludedRegions,omitempty"` // Each exclusion uses java regular expression pattern matching,; and must be separated by a new line.; ; ; myapp/src/main/web/.*\.html; myapp/src/main/web/.*\.jpeg; myapp/src/main/web/.*\.gif; ; The example above illustrates that if only html/jpeg/gif files have been committed to; the SCM a build will not occur.
	IncludedRegions *string `json:"includedRegions,omitempty"` // Each inclusion uses java regular expression pattern matching,; and must be separated by a new line.; An empty list implies that everything is included.; ; ; myapp/src/main/web/.*\.html; myapp/src/main/web/.*\.jpeg; myapp/src/main/web/.*\.gif; ; The example above illustrates that a build will only occur, if html/jpeg/gif files; have been committed to the SCM. Exclusions take precedence over inclusions, if there is; an overlap between included and excluded regions.
}

type PerBuildTag struct {
}

type Phabricator struct {
	Repo    *string `json:"repo,omitempty"`    // Specify the repository name in phabricator (such as the foo part of; phabricator.example.com/diffusion/foo/browse).
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the phabricator instance root URL (such as http://phabricator.example.com).
}

type Plugin struct {
}

type PollSCM struct {
	PollingThreadCount *int64 `json:"pollingThreadCount,omitempty"`
}

type PreBuildMerge struct {
	Options *PreBuildMergeOptions `json:"options,omitempty"`
}

type PreBuildMergeOptions struct {
}

type PrimaryBranchHealthMetric struct {
}

type UnclassifiedProjectNamingStrategy struct {
}

type PruneStaleBranch struct {
}

type PruneTags struct {
	PruneTags *bool `json:"pruneTags,omitempty"`
}

type QuietPeriod struct {
}

type RedmineWeb struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the root URL serving this repository (such as; http://SERVER/PATH/projects/PROJECT/repository).
}

type RelativeTargetDirectory struct {
	RelativeTargetDir *string `json:"relativeTargetDir,omitempty"` // Specify a local directory (relative to the workspace root); where the Git repository will be checked out. If left empty, the workspace root itself; will be used.; ; This extension should not be used in Jenkins Pipeline (either declarative or scripted).; Jenkins Pipeline already provides standard techniques for checkout to a subdirectory.; Use ws and; dir; in Jenkins Pipeline rather than this extension.
}

type ResourceRoot struct {
	URL *string `json:"url,omitempty"` // Jenkins serves many files that are potentially created by untrusted users, such as files; in project workspaces or archived artifacts.; When no resource root URL is defined, Jenkins will serve these files with the HTTP header; Content-Security-Policy ("CSP").; By default it is set to a value that disables many modern web features to prevent; cross-site scripting (XSS) and other attacks on Jenkins users accessing these files.; While the specific value for the CSP header is user configurable (and can even be; disabled), doing so is a trade-off between security and functionality.; ; ; If the resource root URL is defined, Jenkins will instead redirect requests for; user-created resource files to URLs starting with the URL configured here.; These URLs will not set the CSP header, allowing JavaScript and similar features to; work.; For this option to work as expected, the following constraints and considerations; apply:; ; ; The resource root URL must be a valid alternative choice for the Jenkins URL for requests; to be processed correctly.; The Jenkins URL must be set and it must be different from this resource root URL (in; fact, a different host name is required).; ; Once set, Jenkins will only serve resource URL requests via the resource root; URL.; All other requests will get HTTP 404 Not Found responses.; ; ; ; Once this URL has been set up correctly, Jenkins will redirect requests to workspaces,; archived artifacts, and similar collections of usually user-generated content to URLs; starting with the resource root URL.; Instead of a path like job/name_here/ws, resource URLs will contain a token encoding that; path, the user for which the URL was created, and when it was created.; These resource URLs access static files as if the user for which they were created would; access them:; If the user’s permission to access these files is removed, the corresponding resource; URLs will not work anymore either.; These URLs are accessible to anyone without authentication until they expire, so sharing; these URLs is akin to sharing the files directly.; ; Security considerations; Authentication; ; Resource URLs do not require authentication (users will not have a valid session for the; resource root URL).; Sharing a resource URL with another user, even one lacking Overall/Read permission for; Jenkins, will grant that user access to these files until the URLs expire.; ; Expiration; ; Resource URLs expire after 30 minutes by default.; Expired resource URLs will redirect users to their equivalent Jenkins URLs, so that the; user can reauthenticate, if necessary, and then be redirected back to a new resource URL; that will be valid for another 30 minutes.; This will generally be transparent to the user if they have a valid Jenkins session.; Otherwise, they will need to authenticate with Jenkins again.; However, when browsing pages with HTML frames, like Javadoc sites, the login form cannot; appear in a frame.; In these cases, users will need to reload the top-level frame to make the login form; appear.; ; ; To change how quickly resource URLs expire, set the system property; jenkins.security.ResourceDomainRootAction.validForMinutes to the desired value in; minutes.; Earlier expiration might make it harder to use these URLs, while later expiration; increases the likelihood of unauthorized users gaining access through URLs shared with; them by authorized users.; ; Authenticity; ; Resource URLs encode the URL, the user for which they were created, and their creation; timestamp.; Additionally, this string contains an HMAC to ensure the authenticity of the URL.; This prevents attackers from forging URLs that would grant them access to resource files; as if they were another user.
}

type RhodeCode struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the HTTP URL for this repository's RhodeCode page (such as; http://rhodecode.mydomain.com:5000/projects/PROJECT/repos/REPO/).
}

type UnclassifiedSCM struct {
	Git  interface{} `json:"git"`
	None interface{} `json:"none"`
}

type SCMName struct {
	Name *string `json:"name,omitempty"`
}

type SCMRetryCount struct {
}

type SCMSource struct {
	FromSCM interface{} `json:"fromScm"`
	Git     interface{} `json:"git"`
}

type SCMSourceTrait struct {
	CleanAfterCheckoutTrait       interface{} `json:"cleanAfterCheckoutTrait"`
	PruneStaleBranchTrait         interface{} `json:"pruneStaleBranchTrait"`
	GitLFSPullTrait               interface{} `json:"gitLFSPullTrait"`
	IgnoreOnPushNotificationTrait interface{} `json:"ignoreOnPushNotificationTrait"`
	GitBrowser                    interface{} `json:"gitBrowser"`
	RefSpecs                      interface{} `json:"refSpecs"`
	CloneOptionTrait              interface{} `json:"cloneOptionTrait"`
	SparseCheckoutPathsTrait      interface{} `json:"sparseCheckoutPathsTrait"`
	UserIdentityTrait             interface{} `json:"userIdentityTrait"`
	HeadWildcardFilter            interface{} `json:"headWildcardFilter"`
	SubmoduleOptionTrait          interface{} `json:"submoduleOptionTrait"`
	AuthorInChangelogTrait        interface{} `json:"authorInChangelogTrait"`
	GitTool                       interface{} `json:"gitTool"`
	CheckoutOptionTrait           interface{} `json:"checkoutOptionTrait"`
	WipeWorkspaceTrait            interface{} `json:"wipeWorkspaceTrait"`
	LocalBranchTrait              interface{} `json:"localBranchTrait"`
	GitBranchDiscovery            interface{} `json:"gitBranchDiscovery"`
	CleanBeforeCheckoutTrait      interface{} `json:"cleanBeforeCheckoutTrait"`
	PruneStaleTagTrait            interface{} `json:"pruneStaleTagTrait"`
	DiscoverOtherRefsTrait        interface{} `json:"discoverOtherRefsTrait"`
	GitTagDiscovery               interface{} `json:"gitTagDiscovery"`
	HeadRegexFilter               interface{} `json:"headRegexFilter"`
	RemoteName                    interface{} `json:"remoteName"`
}

type Shell struct {
	Shell *string `json:"shell,omitempty"`
}

type SimpleBuildDiscarder struct {
	Discarder *Discarder `json:"discarder,omitempty"`
}

type Discarder struct {
}

type Stash struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Specify the HTTP URL for this repository's Stash page (such as; http://stash.mydomain.com:7990/projects/PROJECT/repos/REPO/).
}

type SubmoduleConfig struct {
	Branches      *string `json:"branches,omitempty"`      // Removed in git plugin 4.6.0.
	SubmoduleName *string `json:"submoduleName,omitempty"` // Removed in git plugin 4.6.0.
}

type SubmoduleOption struct {
	Depth               *int64  `json:"depth,omitempty"`               // Set shallow clone depth, so that git will only download recent history of the project,; saving time and disk space when you just want to access the latest commits of a; repository.
	DisableSubmodules   *bool   `json:"disableSubmodules,omitempty"`   // By disabling support for submodules you can still keep using basic; git plugin functionality and just have Jenkins to ignore submodules; completely as if they didn't exist.
	ParentCredentials   *bool   `json:"parentCredentials,omitempty"`   // Use credentials from the default remote of the parent project.
	RecursiveSubmodules *bool   `json:"recursiveSubmodules,omitempty"` // Retrieve all submodules recursively; ; (uses '--recursive' option which requires git>=1.6.5)
	Reference           *string `json:"reference,omitempty"`           // Specify a folder containing a repository that will be used by Git as a reference during; clone operations.; This option will be ignored if the folder is not available on the controller or agent; where the clone is being executed.; To prepare a reference folder with multiple subprojects, create a bare git repository and; add all the remote urls then perform a fetch:; ; git init --bare; git remote add SubProject1 https://gitrepo.com/subproject1; git remote add SubProject2 https://gitrepo.com/subproject2; git fetch --all
	Shallow             *bool   `json:"shallow,omitempty"`             // Perform shallow clone, so that git will not download the history of the project,; saving time and disk space when you just want to access the latest version of a; repository.
	Threads             *int64  `json:"threads,omitempty"`             // Specify the number of threads that will be used to update submodules.; If unspecified, the command line git default thread count is used.
	Timeout             *int64  `json:"timeout,omitempty"`             // Specify a timeout (in minutes) for submodules operations.; This option overrides the default timeout of 10 minutes.; You can change the global git timeout via the property; org.jenkinsci.plugins.gitclient.Git.timeOut (see JENKINS-11286).; Note that property should be set on both controller and agent to have effect (see; JENKINS-22547).
	TrackingSubmodules  *bool   `json:"trackingSubmodules,omitempty"`  // Retrieve the tip of the configured branch in .gitmodules; ; (Uses '--remote' option which requires git>=1.8.2)
}

type TFS2013Git struct {
	RepoURL *string `json:"repoUrl,omitempty"` // Either the name of the remote whose URL should be used, or the URL of this; module in TFS (such as http://fisheye6.cenqua.com/tfs/PROJECT/_git/REPO/).; If empty (default), the URL of the "origin" repository is used.; If TFS is also used as the repository server, this can usually be left blank.
}

type Timestamper struct {
	AllPipelines      *bool   `json:"allPipelines,omitempty"`      // When checked, timestamps will be enabled for all Pipeline builds.; There is no need to use the timestamps {…} step in Scripted,; or the timestamps() option in Declarative.
	ElapsedTimeFormat *string `json:"elapsedTimeFormat,omitempty"` // <!--; The MIT License; ; Copyright (c) 2013 Steven G. Brown; ; Permission is hereby granted, free of charge, to any person obtaining a copy; of this software and associated documentation files (the "Software"), to deal; in the Software without restriction, including without limitation the rights; to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; copies of the Software, and to permit persons to whom the Software is; furnished to do so, subject to the following conditions:; ; The above copyright notice and this permission notice shall be included in; all copies or substantial portions of the Software.; ; THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; THE SOFTWARE.; -->; ; ; The elapsed time format defines how the timestamps will be rendered when the elapsed time; option has been selected.; The Commons Lang DurationFormatUtils pattern is used.; ; Default is: '&lt;b&gt;'HH:mm:ss.S'&lt;/b&gt; '.
	SystemTimeFormat  *string `json:"systemTimeFormat,omitempty"`  // <!--; The MIT License; ; Copyright (c) 2012 Frederik Fromm; Copyright (c) 2013 Steven G. Brown; ; Permission is hereby granted, free of charge, to any person obtaining a copy; of this software and associated documentation files (the "Software"), to deal; in the Software without restriction, including without limitation the rights; to use, copy, modify, merge, publish, distribute, sublicense, and/or sell; copies of the Software, and to permit persons to whom the Software is; furnished to do so, subject to the following conditions:; ; The above copyright notice and this permission notice shall be included in; all copies or substantial portions of the Software.; ; THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR; IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,; FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE; AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER; LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,; OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN; THE SOFTWARE.; -->; ; ; The system clock time format defines how the timestamps will be rendered.; The JDK SimpleDateFormat pattern is used.; ; Default is: '&lt;b&gt;'HH:mm:ss'&lt;/b&gt; '; For a more detailed format use: yyyy-MM-dd HH:mm:ss.SSS' | '.
}

type UsageStatistics struct {
}

type UserExclusion struct {
	ExcludedUsers *string `json:"excludedUsers,omitempty"` // If set, and Jenkins is set to poll for changes, Jenkins will ignore any revisions; committed by users in this list when determining if a build needs to be triggered. This; can be used to exclude commits done by the build itself from triggering another build,; assuming the build server commits the change with a distinct SCM user.; ; Using this behaviour will preclude the faster git ls-remote polling mechanism, forcing; polling to require a workspace thus sometimes triggering unwanted builds, as if you had; selected the Force polling using workspace extension as well.; Each exclusion uses exact string comparison and must be separated by a new line.; User names are only excluded if they exactly match one of the names in this list.; ; auto_build_user; The example above illustrates that if only revisions by "auto_build_user" have been; committed to the SCM a build will not occur.
}

type UserIdentity struct {
	Email *string `json:"email,omitempty"` // If given, "GIT_COMMITTER_EMAIL=[this]" and "GIT_AUTHOR_EMAIL=[this]" are set for builds.; This overrides whatever is in the global settings.
	Name  *string `json:"name,omitempty"`  // If given, "GIT_COMMITTER_NAME=[this]" and "GIT_AUTHOR_NAME=[this]" are set for builds.; This overrides whatever is in the global settings.
}

type ViewGitWeb struct {
	ProjectName *string `json:"projectName,omitempty"` // Specify the name of the project in ViewGit (e.g. scripts, scuttle etc. from; http://code.fealdia.org/viewgit/).
	RepoURL     *string `json:"repoUrl,omitempty"`     // Specify the root URL serving this repository (such as http://code.fealdia.org/viewgit/).
}

type UnclassifiedViewsTabBar struct {
}

type WipeWorkspace struct {
}

type WorstChildHealthMetric struct {
	Recursive *bool `json:"recursive,omitempty"` // Controls whether items within sub-folders will be considered as contributing to the; health of this folder.
}

type Deprecated string

const (
	DeprecatedReject Deprecated = "reject"
	DeprecatedWarn   Deprecated = "warn"
)

type Restricted string

const (
	Beta             Restricted = "beta"
	RestrictedReject Restricted = "reject"
	RestrictedWarn   Restricted = "warn"
)

type Version string

const (
	The1 Version = "1"
)

type Scope string

const (
	Global Scope = "GLOBAL"
	System Scope = "SYSTEM"
	User   Scope = "USER"
)

type Mode string

const (
	Exclusive Mode = "EXCLUSIVE"
	Normal    Mode = "NORMAL"
)
