package main

import "fmt"

func main() {
	//demo of Graphic Object
	drawing := GraphicObject{"MyDrawing", "", nil}
	drawing.Children = append(drawing.Children, *NewCircle("Red"))
	drawing.Children = append(drawing.Children, *NewSquare("Yellow"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle("Blue"))
	group.Children = append(group.Children, *NewSquare("Blue"))
	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())

	//Neuron Example

	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1, layer2 := &NeuronLayer{}, &NeuronLayer{}

	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)

	fmt.Println(neuron1)
	fmt.Println(neuron2)

	neuron3 := &Neuron{}
	neuron3.ConnectTo(neuron1)
	neuron3.ConnectTo(neuron2)

	fmt.Println(neuron3)
}
