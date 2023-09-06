package options

type Options struct {
	Nodes []*Node `mapstructure:"nodes" yaml:"nodes" json:"nodes"`
}

func (o *Options) GetNodes() []*Node {
	return o.Nodes
}

func (o *Options) GetNodeById(network string, networkId int64) *Node {
	for _, node := range o.Nodes {
		if node.Network == network && node.Id == networkId {
			return node
		}
	}

	return nil
}
