package rpx_handling

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func ResponseTransformerMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := ctx.Next(); err != nil {
			return err
		}

		body := ctx.Response().Body()

		modifiedBody := transformResponseBody(ctx, body)
		ctx.Response().SetBody(modifiedBody)

		return nil
	}
}

func transformResponseBody(ctx *fiber.Ctx, body []byte) []byte {
	statusCode := ctx.Response().StatusCode()

	success := statusCode == fiber.StatusOK

	var response interface{}
	if success {
		response = struct {
			Success bool   `json:"success"`
			Content []byte `json:"content"`
		}{
			Success: success,
			Content: body,
		}
	} else {
		response = struct {
			Success bool   `json:"success"`
			Error   []byte `json:"error"`
			Content []byte `json:"content"`
		}{
			Success: success,
			Error:   body,
		}
	}

	modifiedBody, _ := json.Marshal(response)
	return modifiedBody
}
