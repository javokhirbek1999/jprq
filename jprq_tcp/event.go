package jprq_tcp

type TunnelStartedEvent struct {
	PublicServerPort  int `json:"public_server_port"`
	PrivateServerPort int `json:"private_server_port"`
}

type ConnectionReceivedEvent struct {
	PublicClientIP   string `json:"public_client_ip"`
	PublicClientPort int    `json:"public_client_port"`
}
