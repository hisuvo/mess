package tenants

import (
	"messify/internal/domain/tenants/dto"
)

func ToTenantResponse(tenant *Tenant) *dto.ResponseTenant {
	if tenant == nil {
		return nil
	}
	return &dto.ResponseTenant{
		ID:        tenant.ID,
		Name:      tenant.Name,
		Address:   tenant.Address,
		Phone:     tenant.Phone,
		Email:     tenant.Email,
		Status:    string(tenant.Status),
		CreatedAt: tenant.CreatedAt,
		UpdatedAt: tenant.UpdatedAt,
	}
}

func ToTenantsResponse(tenants []Tenant) []dto.ResponseTenant {
	if len(tenants) == 0 {
		return nil
	}
	var tenantsResponse []dto.ResponseTenant
	for _, tenant := range tenants {
		tenantsResponse = append(tenantsResponse, *ToTenantResponse(&tenant))
	}
	return tenantsResponse
}

func ToTenantResponseWithCount(tenants []Tenant, count int64) *dto.ResponseGetTenants {
	return &dto.ResponseGetTenants{
		Tenants: ToTenantsResponse(tenants),
		Total:   count,
	}
}
