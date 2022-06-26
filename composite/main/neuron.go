package main

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range n.Neurons {
		result = append(result, &n.Neurons[i])
	}
	return result
}

type NeuronInterface interface {
	Iter() []*Neuron
}

func Connect(In NeuronInterface, Out NeuronInterface) {
	for _, l := range In.Iter() {
		for _, r := range Out.Iter() {
			l.ConnectTo(r)
		}
	}
}
