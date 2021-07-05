package common

// AuthorizedEmails represent emails to check whether they have access to certain dto
var AuthorizedEmails = []string{"apa-dev@healthcloud.co.ke", "apa-prod@healthcloud.co.ke"}

// AuthorizedPhones represent phonenumbers to check whether they have access to certain dto
var AuthorizedPhones = []string{"+254700000000"}

// Icon links for navactions
const (
	// StaticBase is the default path at which static assets are hosted
	StaticBase = "https://assets.healthcloud.co.ke"

	AgentNavActionURL    = StaticBase + "/actions/svg/agent_navaction.svg"
	ConsumerNavActionURL = StaticBase + "/actions/svg/consumer_navaction.svg"
	HelpNavActionURL     = StaticBase + "/actions/svg/help_navaction.svg"
	HomeNavActionURL     = StaticBase + "/actions/svg/home_navaction.svg"
	KYCNavActionURL      = StaticBase + "/actions/svg/kyc_navaction.svg"
	PartnerNavActionURL  = StaticBase + "/actions/svg/partner_navaction.svg"
	PatientNavActionURL  = StaticBase + "/actions/svg/patient_navaction.svg"
	RequestNavActionURL  = StaticBase + "/actions/svg/request_navaction.svg"
)

// On Tap Routes
const (
	HomeRoute                  = "/home"
	PatientRegistrationRoute   = "/addPatient"
	PatientIdentificationRoute = "/patients"
	GetHelpRouteRoute          = "/helpCenter"

	// Has KYC and Covers
	RequestsRoute = "/admin"

	AgentRegistrationRoute   = "/agentRegistration"
	AgentIdentificationRoute = "/agentIdentification"
)

// Navigation actions
const (
	HomeNavActionTitle       = "Home"
	HomeNavActionDescription = "Home Navigation action"

	HelpNavActionTitle       = "Help"
	HelpNavActionDescription = "Help Navigation action"

	PatientNavActionTitle            = "Patient"
	PatientNavActionDescription      = "Patient Navigation action"
	PatientRegistrationActionTitle   = "Patient Registration"
	PatientIdentificationActionTitle = "Patient Identification"

	RequestsNavActionTitle       = "Requests"
	RequestsNavActionDescription = "Requests Navigation action"

	AgentNavActionTitle            = "Agent"
	AgentNavActionDescription      = "Agent Navigation action"
	AgentRegistrationActionTitle   = "Agent Registration"
	AgentIdentificationActionTitle = "Agent Identification"

	AgentRegistrationTitle   = "Agent Registration"
	AgentIdentificationTitle = "Agent Identification"

	ConsumerNavActionTitle       = "Consumer"
	ConsumerNavActionDescription = "Consumer Navigation action"

	PartnerNavActionTitle       = "Partner"
	PartnerNavActionDescription = "Partner Navigation action"
)

// PubSub topic names
const (
	// CreateCustomerTopic is the TopicID for customer creation Topic
	CreateCustomerTopic = "customers.create"

	// CreateSupplierTopic is the TopicID for supplier creation Topic
	CreateSupplierTopic = "suppliers.create"

	// CreateCRMContact is the TopicID for CRM contact creation
	CreateCRMContact = "crm.contact.create"

	// UpdateCRMContact is the topicID for CRM contact updates
	UpdateCRMContact = "crm.contact.update"
)
