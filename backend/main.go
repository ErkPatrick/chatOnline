package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// permite o upgrade do protocolo HTTP para WebSocket.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Permitir conexões de qualquer origem
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Cliente struct {
	conn *websocket.Conn
	send chan string
}

// Hub gerencia todos os clientes e mensagens
type Hub struct {
	clientes    map[*Cliente]bool
	broadcast   chan string
	registrar   chan *Cliente
	desconectar chan *Cliente
}

var hub = Hub{
	clientes:    make(map[*Cliente]bool),
	broadcast:   make(chan string),
	registrar:   make(chan *Cliente),
	desconectar: make(chan *Cliente),
}

// Inicializa o Hub e gerencia as mensagens de clientes
func (h *Hub) run() {
	for {
		select {
		case cliente := <-h.registrar:
			h.clientes[cliente] = true
		case cliente := <-h.desconectar:
			if _, ok := h.clientes[cliente]; ok {
				delete(h.clientes, cliente)
				close(cliente.send)
			}
		case mensagem := <-h.broadcast:
			for cliente := range h.clientes {
				select {
				case cliente.send <- mensagem:
				default:
					close(cliente.send)
					delete(h.clientes, cliente)
				}
			}
		}
	}
}

// Lida com novas conexões WebSocket
func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Erro ao conectar WebSocket: %v", err)
		return
	}

	cliente := &Cliente{conn: conn, send: make(chan string, 256)}
	hub.registrar <- cliente

	go cliente.readMessages()
	go cliente.writeMessages()
}

// Lê mensagens recebidas do cliente
func (c *Cliente) readMessages() {
	defer func() {
		hub.desconectar <- c
		c.conn.Close()
	}()
	for {
		_, mensagem, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("Erro ao ler mensagem: %v", err)
			break
		}
		hub.broadcast <- string(mensagem)
	}
}

// Envia mensagens para o cliente
func (c *Cliente) writeMessages() {
	defer c.conn.Close()
	for mensagem := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, []byte(mensagem))
		if err != nil {
			log.Printf("Erro ao enviar mensagem: %v", err)
			break
		}
	}
}

func main() {

	go hub.run()

	// Define a rota para lidar com conexões WebSocket
	http.HandleFunc("/ws", handleConnections)

	// Inicia o servidor
	log.Println("Servidor iniciado na porta :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
