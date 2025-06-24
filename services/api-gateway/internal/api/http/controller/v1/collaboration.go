package v1

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/S1riyS/go-whiteboard/api-gateway/internal/api"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/request"
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/response"
	"github.com/S1riyS/go-whiteboard/api-gateway/pkg/logger/slogext"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins for now (not recommended for production)
	// TODO: use CORS from config
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// CollaborationClient interface remains the same
type CollaborationClient interface {
	Draw(ctx context.Context, whiteboardID string, req *request.CollaborationDrawPayload) (string, error)
}

// CollaborationWSController remains the same
type CollaborationWSController struct {
	logger *slog.Logger
	client CollaborationClient
	hub    *WSHub
}

func NewCollaborationWSController(logger *slog.Logger, client CollaborationClient, hub *WSHub) *CollaborationWSController {
	return &CollaborationWSController{
		logger: logger,
		client: client,
		hub:    hub,
	}
}

func (c *CollaborationWSController) HandleRequest(ctx *gin.Context) {
	const mark = "CollaborationWSController.HandleRequest"
	logger := c.logger.With(slog.String("mark", mark))

	// Retrieve UUID
	id, err := uuid.Parse(ctx.Param("whiteboardID"))
	if err != nil {
		logger.Warn("Invalid UUID format", slog.String("uuid", ctx.Param("uuid")), slogext.Err(err))
		api.NewBadRequestError("Invalid UUID format").WriteToContext(ctx)
		return
	}

	// Upgrade HTTP connection to WebSocket
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.Error("Failed to upgrade connection to WebSocket", slogext.Err(err))
		return
	}
	defer ws.Close()

	// TODO: tweak if needed
	// Configure connection settings
	// ws.SetReadLimit(1024)                                // Max message size in bytes
	// ws.SetReadDeadline(time.Now().Add(60 * time.Second)) // Timeout for read
	// ws.SetPongHandler(func(string) error {
	// 	ws.SetReadDeadline(time.Now().Add(60 * time.Second))
	// 	return nil
	// })

	// Add client to hub
	c.hub.Join(id.String(), ws)
	defer c.hub.Leave(id.String(), ws)

	// WebSocket event loop
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				logger.Error("Unexpected close error", slogext.Err(err))
			} else {
				logger.Debug("WebSocket closed", slogext.Err(err))
			}
			break
		}

		switch messageType {
		case websocket.CloseMessage:
			logger.Debug("Received close message")
			return // Exit the loop on close message

		case websocket.PingMessage:
			logger.Debug("Received ping message")
			// Respond with pong (handled automatically by SetPingHandler)
			if err := ws.WriteControl(websocket.PongMessage, message, time.Now().Add(10*time.Second)); err != nil {
				logger.Error("Failed to send pong", slogext.Err(err))
				return
			}

		case websocket.PongMessage:
			logger.Debug("Received pong message")
			// Update read deadline (handled by SetPongHandler)

		case websocket.TextMessage:
			logger.Debug("Received text message", slog.String("message", string(message)))

			// Process message
			c.handleCollaborationOperation(ctx, ws, id.String(), message)

		case websocket.BinaryMessage:
			logger.Warn("Received binary message (not supported)")
			errorMsg := map[string]string{"error": "binary messages not supported"}
			if err := ws.WriteJSON(errorMsg); err != nil {
				logger.Error("Failed to send error message", slogext.Err(err))
				return
			}

		default:
			logger.Warn("Received unknown message type", slog.Int("type", messageType))
		}
	}
}

func (c *CollaborationWSController) handleCollaborationOperation(ctx *gin.Context, ws *websocket.Conn, whiteboardID string, payload json.RawMessage) {
	var message request.CollaborationMessage
	if err := json.Unmarshal(payload, &message); err != nil {
		c.logger.Warn("Failed to parse JSON message", slogext.Err(err))
		return
	}

	c.logger.Debug("Parsed message", slog.String("message", string(payload)))

	// Tmp constants
	notImplementedErr := api.NewBadRequestError("not implemented")

	switch message.Type {
	case request.CollaborationTypeJoin:
		c.logger.Warn("CollaborationTypeJoin is not implemented yet", slogext.Err(notImplementedErr))
		notImplementedErr.WriteToWebsocket(ws)

	case request.CollaborationTypeDraw:
		var drawPayload request.CollaborationDrawPayload
		if err := json.Unmarshal(message.Payload, &drawPayload); err != nil {
			api.NewBadRequestError("failed to parse draw payload").WriteToWebsocket(ws)
			c.logger.Warn("Failed to parse draw payload", slogext.Err(err))
			return
		}
		c.handleDraw(ctx, ws, whiteboardID, &drawPayload)

	case request.CollaborationTypeDelete:
		c.logger.Warn("CollaborationTypeDelete is not implemented yet", slogext.Err(notImplementedErr))
		notImplementedErr.WriteToWebsocket(ws)

	default:
		c.logger.Warn("Unknown collaboration type", slog.String("type", string(message.Type)))
		api.NewBadRequestError("unknown collaboration type").WriteToWebsocket(ws)
	}
}

func (c *CollaborationWSController) handleDraw(ctx *gin.Context, ws *websocket.Conn, whiteboardID string, payload *request.CollaborationDrawPayload) {
	const mark = "CollaborationWSController.handleDraw"
	logger := c.logger.With(slog.String("mark", mark))

	// * NOTE: No need to recieve drawing from gRPC.
	// * By design, drawings are not modified in microservices.
	// * Hence we can use already parsed drawing.
	id, err := c.client.Draw(ctx.Request.Context(), whiteboardID, payload)
	if err != nil {
		logger.Error("Failed to draw", slogext.Err(err))
		api.NewInternalError().WriteToWebsocket(ws)
		return
	}

	resp := response.DrawResponse{
		ID:   id,
		Type: payload.ElementType,
		Data: payload.Data,
	}

	// Broadcast new drawing
	c.hub.Broadcast(whiteboardID, resp)
}
