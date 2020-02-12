package dc

type DependencyNameType string
type DependencyGroupNameType string

type DependencyContainer struct {
	container map[string]interface{}
}

func NewDC() *DependencyContainer {
	return &DependencyContainer{
		container: map[string]interface{}{},
	}
}

func (dic *DependencyContainer) SetDependency(dependencyName DependencyNameType, dependency interface{}) {
	dic.SetDependencyWithGroup("", dependencyName, dependency)
}

func (dic *DependencyContainer) SetDependencyWithGroup(groupName DependencyGroupNameType, dependencyName DependencyNameType, dependency interface{}) {
	dic.container[string(groupName)+string(dependencyName)] = dependency
}

func (dic *DependencyContainer) GetDependency(dependencyName DependencyNameType) interface{} {
	return dic.GetDependencyWithGroup("", dependencyName)
}

func (dic *DependencyContainer) GetDependencyWithGroup(groupName DependencyGroupNameType, dependencyName DependencyNameType) interface{} {
	dep, _ := dic.container[string(groupName)+string(dependencyName)]
	return dep
}

func (dic *DependencyContainer) DeleteDependency(dependencyName DependencyNameType) {
	dic.DeleteDependencyWithGroup("", dependencyName)
}

func (dic *DependencyContainer) DeleteDependencyWithGroup(groupName DependencyGroupNameType, dependencyName DependencyNameType) {
	delete(dic.container, string(groupName)+string(dependencyName))
}

func (dic *DependencyContainer) DeleteAll() {
	dic.container = map[string]interface{}{}
}
