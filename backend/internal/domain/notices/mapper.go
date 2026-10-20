package notices

import "messify/internal/domain/notices/dto"

func ToNoticeResponse(notice *Notice) *dto.ResponseNotice{
	if notice == nil {
		return nil
	}

	return &dto.ResponseNotice{
		ID: notice.ID,
		Title: notice.Title,
		Description: notice.Description,
		Status: string(notice.Status),
	}
}

func ToNoticeResponses(notices []*Notice) []*dto.ResponseNotice{
	var results []*dto.ResponseNotice

	if len(notices) <= 0 {
		return results
	}


	for _, notice := range notices {
		results = append(results, ToNoticeResponse(notice))
	}

	return results
}