package tree

import (
	"reflect"
)

/**
 * TreeNode
 * @Description: 创建 TreeNode
 * @date 2021-09-16 22:41:59
 * @author 七秒记忆 <274397981@qq.com>
 */
type TreeNode interface {
	TreeId() int
	TreeParentId() int
	TreeChildren(interface{})
}

/**
 * Node
 * @Description: 生成节点菜单
 * @param data
 * @return []TreeNode
 * @date 2021-09-16 22:42:10
 * @author 七秒记忆 <274397981@qq.com>
 */
func Node(data []TreeNode) []TreeNode {
	maxLen := len(data)
	var rootNode []TreeNode = nil
	///<找出根节点,根节点的特点，没有父节点
	for i := 0; i < maxLen; i++ {
		///< 统计每个节点的父节点出现的次数，父节点出现0次就是根节点
		count := 0
		for j := 0; j < maxLen; j++ {
			///< 如果有节点的ID == i的parentID 那么j就是父节点
			if data[j].TreeId() == data[i].TreeParentId() {
				count++
				data[j].TreeChildren(data[i])
			}
		}
		if count == 0 {
			rootNode = append(rootNode, data[i])
		}
	}
	return rootNode
}

/**
 * Tree
 * @Description: 生成树节点菜单
 * @param dest
 * @param fu
 * @return []TreeNode
 * @date 2021-09-16 22:42:19
 * @author 七秒记忆 <274397981@qq.com>
 */
func Tree(dest interface{}, fu func(node interface{}) TreeNode) []TreeNode {
	refValue := reflect.ValueOf(dest)
	if refValue.Kind() == reflect.Ptr {
		refValue = refValue.Elem()
	}
	nodeArray := make([]TreeNode, refValue.Len())
	for i := 0; i < refValue.Len(); i++ {
		nodeArray[i] = fu(refValue.Index(i).Interface())
	}
	result := Node(nodeArray)
	return result
}
