package options

type Node struct {
	Network           string `mapstructure:"network" yaml:"network" json:"network"`
	Type              string `mapstructure:"type" yaml:"type" json:"type"`
	Id                int64  `mapstructure:"network_id" yaml:"network_id" json:"network_id"`
	Endpoint          string `mapstructure:"endpoint" yaml:"endpoint" json:"endpoint"`
	ConcurrentClients int    `mapstructure:"concurrent_clients" yaml:"concurrent_clients" json:"concurrent_clients"`
	Websocket         bool   `mapstructure:"supports_websocket" yaml:"supports_websocket" json:"supports_websocket"`
}

func (n *Node) GetNetwork() string {
	return n.Network
}

func (n *Node) GetType() string {
	return n.Type
}

func (n *Node) GetId() int64 {
	return n.Id
}

func (n *Node) GetEndpoint() string {
	return n.Endpoint
}

func (n *Node) GetConcurrentClients() int {
	return n.ConcurrentClients
}

func (n *Node) IsWebsocket() bool {
	return n.Websocket
}
