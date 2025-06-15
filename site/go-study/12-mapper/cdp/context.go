package cdp

type Context struct {
	Channel  string    `json:"channel"`
	Messages []Message `json:"messages"`
	Identity Identity  `json:"identity"`
	Sender   Sender    `json:"sender"`
	Tenant   Tenant    `json:"tenant"`
}

type Arg struct {
	ServiceCode   string `json:"serviceCode"`
	BundleName    string `json:"bundleName"`
	BundleVersion string `json:"bundleVersion"`
}

type Message struct {
	Args          Arg      `json:"args"`
	MessageCode   string   `json:"messageCode"`
	MessageType   string   `json:"messageType"`
	Priority      string   `json:"priority"`
	Severity      string   `json:"severity"`
	Message       string   `json:"message"`
	ObjectKey     string   `json:"objectKey"`
	ObjectPointer string   `json:"objectPointer"`
	ValueNames    []string `json:"valueNames"`
}

type Identity struct {
	TransactionId     string `json:"transactionId"`
	ConversationId    string `json:"conversationId"`
	MessageId         string `json:"messageId"`
	MessageInstanceId string `json:"messageInstanceId"`
}

type Sender struct {
	DeputyId string `json:"deputyId"`
}

type Tenant struct {
	TenantId    string `json:"tenantId"`
	TenantOrgId string `json:"tenantOrgId"`
	ConsumerId  string `json:"consumerId"`
}
