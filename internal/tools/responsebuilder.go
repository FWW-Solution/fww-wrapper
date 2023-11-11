package tools

import "fww-wrapper/internal/data/dto"

func ResponseBuilder(data interface{}, meta interface{}) dto.BaseResponse {
	return dto.BaseResponse{
		Meta: meta,
		Data: data,
	}
}

func ResponseInternalServerError(err error) dto.BaseResponse {
	return dto.BaseResponse{
		Meta: dto.MetaResponse{
			StatusCode: "ERR500",
			IsSuccess:  false,
			Message:    err.Error(),
		},
		Data: nil,
	}

}
