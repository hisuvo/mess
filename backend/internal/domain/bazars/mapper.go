package bazars

import "messify/internal/domain/bazars/dto"

func ToBazarReponse(bazar *Bazar) *dto.ResponseBazar{
	if bazar == nil {
		return nil
	}

	return &dto.ResponseBazar{
		ID: bazar.ID,
		Total: bazar.TotalAmount,
	}
}