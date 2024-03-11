package rpx_handling

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// ResponseTransformerMiddleware es un middleware que transforma la respuesta antes de enviarla al cliente
func ResponseTransformerMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Llamar al siguiente middleware en la cadena
		if err := ctx.Next(); err != nil {
			return err
		}

		// Transformar la respuesta
		body := ctx.Response().Body()

		// Modificar el cuerpo de la respuesta según sea necesario
		modifiedBody := transformResponseBody(ctx, body)

		// Asignar el cuerpo modificado a la respuesta
		ctx.Response().SetBody(modifiedBody)

		// Continuar con el siguiente middleware
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
