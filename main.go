package main

import (
	"fmt"
	"math"
)

//create position type
type pairPositions struct {
	a int
	b int
}

// create node type
type Node struct {
	parent *Node 
	position pairPositions 
	g int
	h int
	f int
	wasSet bool
} 

func (n *Node) initTheNode(parentParm Node , positionParm pairPositions) {
n.parent=&parentParm
n.position=positionParm 
 }

//test if two nodes is equal
func (n Node) equalNodes(anotherNode Node) bool{

	return n.position == anotherNode.position
}

//Remove item from the slice
func removeItemFromSlice( slice [] Node  , index int) []Node {


return append( slice[:index] , slice[index+1:]... )
}

// function for reverse the slice
func reverseSlice(slice []pairPositions) []pairPositions{
mySlice := make([]pairPositions,len(slice))

	 copy(mySlice,slice)
	for i, j := 0, len(mySlice)-1; i < j; i, j = i+1, j-1 {
    mySlice[i], mySlice[j] = mySlice[j], mySlice[i]
	
}
return mySlice
}

// Function that will do everything
func astar(maze [][]int , start , end pairPositions) []pairPositions{
	//Create open and close list
var openList,closeList =  []	Node{} , [] Node {}
//Create start node
startNode:=Node{wasSet: true, position:start}
//Create end node
endNode:=Node{wasSet: true, position:end}
// Add the start and the end node
openList= append(openList, startNode )
closeList = append(closeList,endNode)

// Loop until you find the end
for len(openList) > 0 {
	var currentIndex int
	// Get the current node
	currentItem:=openList[0]
	for i,v:=range openList{
if v.f < currentItem.f{
	currentItem = v
	currentIndex = i
}

	}
	//Remove the current node from the open list and add it to the closed list
	openList= removeItemFromSlice(openList,currentIndex)
closeList = append(closeList, currentItem)

//Find the goal
if currentItem.equalNodes(endNode){
	path:=[]pairPositions{}
	theCurrentItemInPath:=currentItem
	for !theCurrentItemInPath.wasSet {
		path= append(path, theCurrentItemInPath.position)
theCurrentItemInPath= *theCurrentItemInPath.parent
	}
// Return the path after reverse it 
	return reverseSlice(path)
}
// Generate children
children:=[]Node{}
adjacentSquares:=[...]pairPositions{pairPositions{ 0, -1},pairPositions{  0, 1},pairPositions{ -1, 0}, pairPositions{ 1, 0}, pairPositions{ -1, -1}, pairPositions{ -1, 1},pairPositions{  1, -1},pairPositions{  1, 1}}
	

for _, newPosition := range adjacentSquares{
//Get node position
	nodePosition := pairPositions{currentItem.position.a + newPosition.a, currentItem.position.b + newPosition.b}
//Make sure within range
	if nodePosition.a > (len(maze) - 1) || nodePosition.a < 0 || nodePosition.b > (len(maze[len(maze)-1]) -1) || nodePosition.b < 0{
                continue
}

//Make sure walkable terrain
 if maze[nodePosition.a][nodePosition.b] != 0{
                continue
			}
			//Create new node
			    newNode := Node{parent: &currentItem,position:  nodePosition }
				//Append it
				children=append(children, newNode)
}

// Loop through children

for _,child:= range children{
	//Child is on the closed list
	for _, closedChild := range closeList {
		if child.equalNodes(closedChild){
			continue
		}
	}
//Create the f, g, and h values
child.g=currentItem.g+1

child.h = (int(math.Pow (float64(child.position.a - endNode.position.a) , 2)) + int(math.Pow (float64 (child.position.b - endNode.position.b) , 2)))
child.f=child.g+child.h
// Child is already in the open list
for _, openNode :=range openList{
                if child .equalNodes( openNode) && child.g > openNode.g{
                    continue
				}
}
//Add the child to the open list
openList=append(openList, child)
}
}

return nil

}


func main() {
//Simulation
  grid :=[][]int {{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
            {0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 1, 0, 1, 0, 0, 0},
            {0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 1, 0, 0, 1, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

	start:=pairPositions{1,3}
	goal:=pairPositions{5,2}
	
	path:=astar(grid,start,goal)
path=append([]pairPositions{start} , path...)
	fmt.Println(path)

} 