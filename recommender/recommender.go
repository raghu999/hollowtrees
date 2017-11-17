package recommender

import (
	"github.com/banzaicloud/hollowtrees/conf"
	"github.com/sirupsen/logrus"
)

type AZRecommendation map[string][]InstanceTypeInfo

type InstanceTypeInfo struct {
	InstanceTypeName   string
	CurrentPrice       string
	AvgPriceFor24Hours float32
	OnDemandPrice      string
	CostScore          float32
	StabilityScore     float32
}

var log *logrus.Logger

func RecommendSpotInstanceTypes(region string, az string, baseInstanceType string) (AZRecommendation, error) {
	log = conf.Logger()
	log.Info(region, az, baseInstanceType)

	// validate region and base instance type
	// get instance types based on base instance type from pricing api (based on cpus, mem, etc..)
	// compute cost/ondemand/stabilityscore/avgprice/currentprice in seleced AZ

	var azRecommendations = AZRecommendation{
		"eu-west-1a": []InstanceTypeInfo{
			InstanceTypeInfo{
				InstanceTypeName:   "c5.xlarge",
				CurrentPrice:       "0.2",
				AvgPriceFor24Hours: 0.1,
				OnDemandPrice:      "0.22",
				CostScore:          0.3,
				StabilityScore:     0.5,
			},
			InstanceTypeInfo{
				InstanceTypeName:   "c5.2xlarge",
				CurrentPrice:       "0.065",
				AvgPriceFor24Hours: 0.07,
				OnDemandPrice:      "0.25",
				CostScore:          0.95,
				StabilityScore:     0.99,
			},
		},

		"eu-west-1b": []InstanceTypeInfo{
			InstanceTypeInfo{
				InstanceTypeName:   "c5.xlarge",
				CurrentPrice:       "0.1",
				AvgPriceFor24Hours: 0.08,
				OnDemandPrice:      "0.22",
				CostScore:          0.6,
				StabilityScore:     0.8,
			},
			InstanceTypeInfo{
				InstanceTypeName:   "c5.2xlarge",
				CurrentPrice:       "0.06",
				AvgPriceFor24Hours: 0.065,
				OnDemandPrice:      "0.25",
				CostScore:          0.99,
				StabilityScore:     0.99,
			},
		},
	}

	return azRecommendations, nil
}