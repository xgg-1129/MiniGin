package MiniGin

import (
	"strings"
)

type node struct {
	Pattern string
	Value string
	Dim bool
	children []*node
}

func (n *node) matchChild(part string)*node{
	for _,child := range n.children{
		if part == child.Value || child.Dim{
			return child
		}
	}
	return nil
}
func (n *node) matchChildren(part string)[]*node{
	res := make([]*node,0)
	for _,child := range n.children{
		if part==child.Value||child.Dim{
			res=append(res,child)
		}
	}
	return res
}
func ParsePattern(pattern string)[]string{
	split:=strings.Split(pattern,"/")
	parts:= make([]string,0)
	for _,item := range split{
		if item!=""{
			parts=append(parts,item)
			if item[0] == '*'{
				break
			}
		}
	}
	return parts
}
func (n *node) Insert(pattern string,parts []string,height int)  {
	if height == len(parts){
		n.Pattern=pattern
		return
	}
	part:=parts[height]
	child := n.matchChild(part)
	if child==nil{
		child = &node{
			Pattern:  "",
			Value:    part,
			Dim:      part[0] == ':' || part[0] == '*',
			children: nil,
		}
		n.children=append(n.children,child)
	}
	child.Insert(pattern,parts,height+1)
}
func (n *node) Search(parts []string,height int)*node{
	if len(parts)==height || strings.HasPrefix(n.Value, "*"){
		if n.Pattern==""{
			return nil
		}
		return n
	}
	children:=n.matchChildren(parts[height])
	for _,child := range children{
		res:=child.Search(parts,height+1)
		if res!=nil{
			return res
		}
	}
	return nil
}