# 树型结构
## 使用代码演示
```Go
import (
	"encoding/json"
	"github.com/go-kb/tree"
  "fmt"
)

type Model struct {
	Id       int      `json:"id"`
	Pid      int      `json:"pid"`
	Name     string   `json:"name"`
	Children []*Model `json:"children,omitempty"`
}

func (d *Model) TreeId() int {
	return d.Id
}

func (d *Model) TreeParentId() int {
	return d.Pid
}


func (d *Model) TreeChildren(node interface{}) {
	d.Children = append(d.Children, node.(*Model))
}

 
  model := make([]Model, 0)
	model = append(model,
		Model{
			Id:   1,
			Pid:  0,
			Name: "顶级1",
		},
		Model{
			Id:   2,
			Pid:  0,
			Name: "顶级2",
		},
		Model{
			Id:   3,
			Pid:  0,
			Name: "顶级3",
		},
		Model{
			Id:   4,
			Pid:  2,
			Name: "一级2",
		},
	)
  
```
### 方法1
```Go
func main() {
	nodeArray := make([]tree.TreeNode, len(model))
	for i := 0; i < len(model); i++ {
		nodeArray[i] = &model[i]
	}
	result := tree.Node(nodeArray)
  fmt.Println(result)
  }

```
### 方法2
```GO
func main() {
result := tree.Tree(model, func(node interface{}) tree.TreeNode {
		model := node.(Model)
		return &model
	})
  fmt.Println(result)
  }
```
