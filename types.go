package main

type InterchangeHeader struct {
	AuthorizationInfoQualifier     string
	AuthorizationInformation       string
	SecurityInfoQualifier          string
	SecurityInfo                   string
	InterchangeSenderIDQualifier   string
	InterchangeSenderID            string
	InterchangeReceiverIDQualifier string
	InterchangeReceiverID          string
	InterchangeDate                string
	InterchangeTime                string
	InterchangeControlStandardsID  string
	InterchangeControlVersion      string
	InterchangeControlNumber       string
	AcknowledgmentRequested        string
	UsageIndicator                 string
	ComponentElementSeparator      string
}

type FunctionGroupHeader struct {
	FunctionalIDCode         string
	ApplicationSenderCode    string
	ApplicationReceiverCode  string
	Date                     string
	Time                     string
	GroupControlNumber       string
	ResponsibleAgencyCode    string
	VersionReleaseIndustryID string
}

type TransactionHeader struct {
	FunctionalIDCode         string
	ApplicationSenderCode    string
	ApplicationReceiverCode  string
	Date                     string
	Time                     string
	GroupControlNumber       string
	ResponsibleAgencyCode    string
	VersionReleaseIndustryID string
}

type IDValue struct {
	ID    string
	Value string
}

type Segment struct {
	ID       string
	Elements []IDValue
}

type Transaction struct {
	Header   TransactionHeader
	Segments []Segment
}

type FunctionGroup struct {
	Header       FunctionGroupHeader
	Transactions []Transaction
}

type Interchange struct {
	Header         InterchangeHeader
	FunctionGroups []FunctionGroup
}
