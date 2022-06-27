package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-bp-relations-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetBusinessPartnerRelationship(firstBusinessPartnerID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "BusinessPartnerRelationshipCollection":
			func() {
				c.BusinessPartnerRelationshipCollection(firstBusinessPartnerID)
				wg.Done()

			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) BusinessPartnerRelationshipCollection(firstBusinessPartnerID string) {
	data, err := c.callBusinessPartnerRelationshipSrvAPIRequirementBusinessPartnerRelationshipCollection("BusinessPartnerRelationshipCollection", firstBusinessPartnerID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBusinessPartnerRelationshipSrvAPIRequirementBusinessPartnerRelationshipCollection(api, firstBusinessPartnerID string) ([]sap_api_output_formatter.BusinessPartnerRelationshipCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithBusinessPartnerRelationshipCollection(req, firstBusinessPartnerID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToBusinessPartnerRelationshipCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithBusinessPartnerRelationshipCollection(req *http.Request, firstBusinessPartnerID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("FirstBusinessPartnerID eq '%s'", firstBusinessPartnerID))
	req.URL.RawQuery = params.Encode()
}
