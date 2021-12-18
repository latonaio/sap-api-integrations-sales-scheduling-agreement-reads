package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	PurchaseSchedulingAgreement struct {
		SchedulingAgreement       string      `json:"document_no"`
		Plant                     string      `json:"deliver_to"`
		TargetQuantity            string      `json:"quantity"`
		PickedQuantity            string      `json:"picked_quantity"`
		NetPriceAmount            string      `json:"price"`
	    Batch                     string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             string      `json:"quantity"`
		CompletedQuantity    string      `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 string      `json:"quantity"`
			CompletedQuantity        string      `json:"completed_quantity"`
			ErroredQuantity          string      `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity string      `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                         string `json:"api_schema"`
	Accepter                        []string `json:"accepter"`
	MaterialCode                      string `json:"material_code"`
	Supplier                          string `json:"plant/supplier"`
	Stock                             string `json:"stock"`
	PurchasingDocumentCategory        string `json:"document_type"`
	SSANumber                         string `json:"document_no"`
	ScheduleLineDeliveryDate          string `json:"planned_date"`
	ValidatedDate                     string `json:"validated_date"`
    Deleted                           bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey            string `json:"connection_key"`
	Result                   bool   `json:"result"`
	RedisKey                 string `json:"redis_key"`
	Filepath                 string `json:"filepath"`
	SalesSchedulingAgreement struct {
		SalesSchedulingAgreement       string `json:"SalesSchedulingAgreement"`
		CompanyCode                    string `json:"CompanyCode"`
		PurchasingDocumentCategory     string `json:"PurchasingDocumentCategory"`
		PurchasingDocumentType         string `json:"PurchasingDocumentType"`
		CreationDate                   string `json:"CreationDate"`
		Language                       string `json:"Language"`
		PurchasingOrganization         string `json:"PurchasingOrganization"`
		PurchasingGroup                string `json:"PurchasingGroup"`
		DocumentCurrency               string `json:"DocumentCurrency"`
		IncotermsClassification        string `json:"IncotermsClassification"`
		PaymentTerms                   string `json:"PaymentTerms"`
		NetPaymentDays                 string `json:"NetPaymentDays"`
		TargetAmount                   string `json:"TargetAmount"`
		ExchangeRate                   string `json:"ExchangeRate"`
		PurchasingDocumentOrderDate    string `json:"PurchasingDocumentOrderDate"`
		Supplier                       string `json:"Supplier"`
		SupplierAddressID              string `json:"SupplierAddressID"`
		ValidityStartDate              string `json:"ValidityStartDate"`
		ValidityEndDate                string `json:"ValidityEndDate"`
		PurchasingDocumentOrigin       string `json:"PurchasingDocumentOrigin"`
		PurchasingDocumentDeletionCode string `json:"PurchasingDocumentDeletionCode"`
		SupplierRespSalesPersonName    string `json:"SupplierRespSalesPersonName"`
		SupplierPhoneNumber            string `json:"SupplierPhoneNumber"`
		InvoicingParty                 string `json:"InvoicingParty"`
		SchedulingAgreementStatus      string `json:"SchedulingAgreementStatus"`
		HeaderPartner                  struct {
			SalesSchedulingAgreementItem string      `json:"SalesSchedulingAgreementItem"`
			PurchasingOrganization       string      `json:"PurchasingOrganization"`
			SupplierSubrange             string      `json:"SupplierSubrange"`
			Plant                        string      `json:"Plant"`
			PartnerFunction              string      `json:"PartnerFunction"`
			PartnerCounter               string      `json:"PartnerCounter"`
			Supplier                     string      `json:"Supplier"`
			DefaultPartner               interface{} `json:"DefaultPartner"`
		} `json:"HeaderPartner"`
		SalesSchedulingAgreementItem struct {
			SalesSchedulingAgreementItem   string      `json:"SalesSchedulingAgreementItem"`
			CompanyCode                    string      `json:"CompanyCode"`
			PurchasingDocumentCategory     string      `json:"PurchasingDocumentCategory"`
			PurchasingDocumentItemCategory string      `json:"PurchasingDocumentItemCategory"`
			PurchasingDocumentItemText     string      `json:"PurchasingDocumentItemText"`
			Material                       string      `json:"Material"`
			SupplierMaterialNumber         string      `json:"SupplierMaterialNumber"`
			MaterialGroup                  string      `json:"MaterialGroup"`
			Plant                          string      `json:"Plant"`
			ReferenceDeliveryAddressID     string      `json:"ReferenceDeliveryAddressID"`
			IncotermsClassification        string      `json:"IncotermsClassification"`
			OrderQuantityUnit              string      `json:"OrderQuantityUnit"`
			TargetQuantity                 string      `json:"TargetQuantity"`
			PurchaseRequisition            string      `json:"PurchaseRequisition"`
			PurchaseRequisitionItem        string      `json:"PurchaseRequisitionItem"`
			SchedAgrmtAgreedCumQty         string      `json:"SchedAgrmtAgreedCumQty"`
			SchedAgrmtCumQtyReconcileDate  string      `json:"SchedAgrmtCumQtyReconcileDate"`
			AccountAssignmentCategory      string      `json:"AccountAssignmentCategory"`
			NetPriceAmount                 string      `json:"NetPriceAmount"`
			NetPriceQuantity               string      `json:"NetPriceQuantity"`
			OrderPriceUnit                 string      `json:"OrderPriceUnit"`
			ProductType                    string      `json:"ProductType"`
			MaterialType                   string      `json:"MaterialType"`
			StorageLocation                string      `json:"StorageLocation"`
			DocumentCurrency               string      `json:"DocumentCurrency"`
			PurchasingInfoRecord           string      `json:"PurchasingInfoRecord"`
			PurchasingDocumentDeletionCode string      `json:"PurchasingDocumentDeletionCode"`
			UnderdelivTolrtdLmtRatioInPct  string      `json:"UnderdelivTolrtdLmtRatioInPct"`
			OverdelivTolrtdLmtRatioInPct   string      `json:"OverdelivTolrtdLmtRatioInPct"`
			UnlimitedOverdeliveryIsAllowed interface{} `json:"UnlimitedOverdeliveryIsAllowed"`
			StockType                      string      `json:"StockType"`
			TaxCode                        string      `json:"TaxCode"`
			TaxCountry                     string      `json:"TaxCountry"`
			GoodsReceiptIsExpected         interface{} `json:"GoodsReceiptIsExpected"`
			GoodsReceiptIsNonValuated      interface{} `json:"GoodsReceiptIsNonValuated"`
			InvoiceIsExpected              interface{} `json:"InvoiceIsExpected"`
			InvoiceIsGoodsReceiptBased     interface{} `json:"InvoiceIsGoodsReceiptBased"`
			EvaldRcptSettlmtIsAllowed      interface{} `json:"EvaldRcptSettlmtIsAllowed"`
			ItemScheduleLine               struct {
				ScheduleLine              string      `json:"ScheduleLine"`
				DelivDateCategory         string      `json:"DelivDateCategory"`
				ScheduleLineDeliveryDate  string      `json:"ScheduleLineDeliveryDate"`
				ScheduleLineDeliveryTime  string      `json:"ScheduleLineDeliveryTime"`
				OrderQuantityUnit         string      `json:"OrderQuantityUnit"`
				ScheduleLineOrderQuantity string      `json:"ScheduleLineOrderQuantity"`
				PurchaseRequisition       string      `json:"PurchaseRequisition"`
				PurchaseRequisitionItem   string      `json:"PurchaseRequisitionItem"`
				ScheduleLineIsFixed       interface{} `json:"ScheduleLineIsFixed"`
			} `json:"ItemScheduleLine"`
			ItemDeliveryAddress struct {
				DeliveryAddressID      string `json:"DeliveryAddressID"`
				AddressType            string `json:"AddressType"`
				StreetName             string `json:"StreetName"`
				PostalCode             string `json:"PostalCode"`
				CityName               string `json:"CityName"`
				MobileNumber           string `json:"MobileNumber"`
				Region                 string `json:"Region"`
				Country                string `json:"Country"`
				EmailAddress           string `json:"EmailAddress"`
				Plant                  string `json:"Plant"`
				CorrespondenceLanguage string `json:"CorrespondenceLanguage"`
				PhoneNumber            string `json:"PhoneNumber"`
				FaxNumber              string `json:"FaxNumber"`
			} `json:"ItemDeliveryAddress"`
		} `json:"SalesSchedulingAgreementItem"`
	} `json:"SalesSchedulingAgreement"`
	APISchema                string   `json:"api_schema"`
	Accepter                 []string `json:"accepter"`
	SSANumber                string   `json:"sales_scheduling_agreement"`
	Deleted                  bool     `json:"deleted"`
}
