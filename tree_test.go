package tree_test

import (
	"encoding/json"
	"github.com/go-kb/tree"
	"testing"
)

/**
 * Model
 * @Description:
 * @date 2021-09-16 22:43:42
 * @author 七秒记忆 <274397981@qq.com>
 */
type Model struct {
	Id       int      `json:"id"`
	Pid      int      `json:"pid"`
	Name     string   `json:"name"`
	Children []*Model `json:"children,omitempty"`
}

/**
 * TreeId
 * @Description: ID
 * @receiver d
 * @return int
 * @date 2021-09-16 22:43:33
 * @author 七秒记忆 <274397981@qq.com>
 */
func (d *Model) TreeId() int {
	return d.Id
}

/**
 * TreeParentId
 * @Description: 父级ID
 * @receiver d
 * @return int
 * @date 2021-09-16 22:43:10
 * @author 七秒记忆 <274397981@qq.com>
 */
func (d *Model) TreeParentId() int {
	return d.Pid
}

/**
 * TreeChildren
 * @Description: 子级
 * @receiver d
 * @param node
 * @date 2021-09-16 22:42:57
 * @author 七秒记忆 <274397981@qq.com>
 */
func (d *Model) TreeChildren(node interface{}) {
	d.Children = append(d.Children, node.(*Model))
}

/**
 * Test
 * @Description:
 * @param t
 * @date 2021-09-16 22:42:36
 * @author 七秒记忆 <274397981@qq.com>
 */
func Test(t *testing.T) {
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
	/**转换树结构**/
	nodeArray := make([]tree.TreeNode, len(model))
	for i := 0; i < len(model); i++ {
		nodeArray[i] = &model[i]
	}
	result := tree.Node(nodeArray)
	data := `[{"id":1,"pid":0,"name":"顶级1"},{"id":2,"pid":0,"name":"顶级2","children":[{"id":4,"pid":2,"name":"一级2"}]},{"id":3,"pid":0,"name":"顶级3"}]`
	a, _ := json.Marshal(result)
	if string(a) != data {
		t.Fatal("test fail")
	}
}

/**
 * Test2
 * @Description:
 * @param t
 * @date 2021-09-16 22:42:32
 * @author 七秒记忆 <274397981@qq.com>
 */
func Test2(t *testing.T) {
	m := make([]Model, 0)
	m = append(m,
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
	result := tree.Tree(m, func(node interface{}) tree.TreeNode {
		model := node.(Model)
		return &model
	})
	res, _ := json.Marshal(result)
	data := `[{"id":1,"pid":0,"name":"顶级1"},{"id":2,"pid":0,"name":"顶级2","children":[{"id":4,"pid":2,"name":"一级2"}]},{"id":3,"pid":0,"name":"顶级3"}]`
	if string(res) != data {
		t.Fatal("test fail")
	}

}
