package deposits

import "messify/internal/domain/deposits/dto"

func ToDepositeResponse(deposit *Deposits) *dto.DepositsResponse{
	if deposit == nil {
		return nil
	}

	return  &dto.DepositsResponse{}
}

func ToDepositeResponses(deposites []*Deposits) []*dto.DepositsResponse{
	if len(deposites) == 0 {
		return nil
	}

	var depositeResponse []*dto.DepositsResponse

	for _, deposite := range deposites {
		depositeResponse = append(depositeResponse, ToDepositeResponse(deposite))
	}


	return depositeResponse
}