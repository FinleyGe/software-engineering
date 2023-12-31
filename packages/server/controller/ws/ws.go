package ws

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

type CallMessage struct {
	MessageType string `json:"messageType"`
	PatientID   uint64 `json:"patientID"`
	DoctorID    uint64 `json:"doctorID"`
}

var Connections sync.Map

func HandleWSConnection(c *gin.Context) {
	ws, err := websocket.Accept(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	id := c.GetUint64("id")
	Connections.Store(id, ws)

	defer Connections.Delete(id)
	defer ws.Close(websocket.StatusNormalClosure, "")

	<-c.Writer.CloseNotify()
}
