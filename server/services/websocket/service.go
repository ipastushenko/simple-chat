package websocket

import (
    "sync"
    "log"
    "github.com/gorilla/websocket"
)

type WebSocketService struct {
    clients map[*Client]bool
    users map[int]*Client
    register chan *Client
    unregister chan *Client
    broadcast chan *BroadcastInfo
}

var (
    service *WebSocketService
    once sync.Once
)

func NewWebSocketService() IWebSocketService {
    once.Do(func () {
        service = &WebSocketService{
            clients: make(map[*Client]bool),
            users: make(map[int]*Client),
            register: make(chan *Client),
            unregister: make(chan *Client),
            broadcast: make(chan *BroadcastInfo),
        }
        go service.run()
    })

    return service
}

func (service *WebSocketService) unregisterClient(client *Client) {
    if _, ok := service.users[client.UserId]; ok {
        delete(service.users, client.UserId)
    }
    if _, ok := service.clients[client]; ok {
        delete(service.clients, client)
    }
}

func (service *WebSocketService) registerClient(client *Client) {
    service.clients[client] = true
    service.users[client.UserId] = client
}

func (service *WebSocketService) broadcastClients(broadcastInfo *BroadcastInfo) {
    if len(broadcastInfo.userIds) > 0 {
        for userId := range broadcastInfo.userIds {
            client := service.users[userId]
            select {
            case client.Send <- broadcastInfo.message:
            default:
                service.unregisterClient(client)
            }
        }
    } else {
        for client := range service.clients {
            select {
            case client.Send <- broadcastInfo.message:
            default:
                service.unregisterClient(client)
            }
        }
    }
}

func (service *WebSocketService) run() {
    for {
        select {
        case client := <- service.register:
            service.registerClient(client)

        case client := <- service.unregister:
            service.unregisterClient(client)

        case broadcastInfo := <- service.broadcast:
            service.broadcastClients(broadcastInfo)
        }
    }
}

func (service *WebSocketService) InitConnection(
    connection *websocket.Conn,
    userData interface{},
) {
    data, ok := userData.(map[string]interface{});
    if !ok {
        connection.Close()
        return
    }
    userId, ok := data["user_id"].(float64)
    if !ok {
        connection.Close()
        return
    }

    log.Println("User ", userId, " has been connected")
    client := NewClient(service, connection, int(userId))
    service.register <- client
    go client.InitWriteHandler()
    go client.InitReadHandler()
}

func (service *WebSocketService) Broadcast(broadcastInfo *BroadcastInfo) {
    service.broadcast <- broadcastInfo
}
