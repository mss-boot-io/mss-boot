/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 9:32
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 9:32
 */

package mongodb

import "github.com/kamva/mgm/v3"

type Tabler interface {
	mgm.Model
	Make()
}
