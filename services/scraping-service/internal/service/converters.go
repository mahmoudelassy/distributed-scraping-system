package service

import (
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"
)

func ToInternalQuery(dto QueryDTO) domain.Query {
	return domain.Query{
		Selector: dto.Selector,
		Label:    dto.Label,
		All:      dto.All,
	}
}

func ToInternalQueries(dtos []QueryDTO) []domain.Query {
	queries := make([]domain.Query, len(dtos))
	for i, dto := range dtos {
		queries[i] = ToInternalQuery(dto)
	}
	return queries
}

func ToResultDTO(result *domain.Result) ResultDTO {
	elements := make([]ElementDTO, len(result.Elements))
	for i, el := range result.Elements {
		elements[i] = ElementDTO{
			Text:       el.GetText(),
			Attributes: el.GetAttributes(),
		}
	}

	return ResultDTO{
		Selector: result.Query.Selector,
		Label:    result.Query.Label,
		Success:  result.Status == domain.QuerySuccess,
		Elements: elements,
	}
}

func ToResultDTOs(results []*domain.Result) []ResultDTO {
	dtos := make([]ResultDTO, len(results))
	for i, result := range results {
		dtos[i] = ToResultDTO(result)
	}
	return dtos
}

func ToErrorDTO(err error) *ErrorDTO {
	if err == nil {
		return nil
	}

	if scraperErr, ok := err.(*domain.ScraperError); ok {
		return &ErrorDTO{
			Type:    "scraper_error",
			Message: scraperErr.Error(),
			Stage:   scraperErr.Stage,
		}
	}

	return &ErrorDTO{
		Type:    "error",
		Message: err.Error(),
	}
}
