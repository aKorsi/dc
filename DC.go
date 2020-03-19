package dc

type DependencyNameType string
type DependencyGroupNameType string

type DependencyContainer struct {
	builderContainer map[string]func() interface{}
	depContainer     map[string]interface{}
}

func NewDC() *DependencyContainer {
	return &DependencyContainer{
		builderContainer: map[string]func() interface{}{},
		depContainer:     map[string]interface{}{},
	}
}

func (dic *DependencyContainer) SetDependency(dependencyName DependencyNameType, depFunc func() interface{}) {
	dic.SetDependencyWithGroup("", dependencyName, depFunc)
}

func (dic *DependencyContainer) SetDependencyWithGroup(groupName DependencyGroupNameType, dependencyName DependencyNameType, depFunc func() interface{}) {
	dic.builderContainer[string(groupName)+string(dependencyName)] = depFunc
}

func (dic *DependencyContainer) GetSingletonDependency(dependencyName DependencyNameType) interface{} {
	return dic.GetSingletonDependencyWithGroup("", dependencyName)
}

func (dic *DependencyContainer) GetSingletonDependencyWithGroup(groupName DependencyGroupNameType, dependencyName DependencyNameType) interface{} {
	depData, _ := dic.depContainer[string(groupName)+string(dependencyName)]
	if depData == nil {
		builderFunc, _ := dic.builderContainer[string(groupName)+string(dependencyName)]
		depData = builderFunc()
		dic.depContainer[string(groupName)+string(dependencyName)] = depData
	}
	return depData
}

func (dic *DependencyContainer) GetDependency(dependencyName DependencyNameType) interface{} {
	return dic.GetDependencyWithGroup("", dependencyName)
}

func (dic *DependencyContainer) GetDependencyWithGroup(groupName DependencyGroupNameType, dependencyName DependencyNameType) interface{} {
	depFunc, _ := dic.builderContainer[string(groupName)+string(dependencyName)]
	return depFunc()
}

func (dic *DependencyContainer) DeleteDependency(dependencyName DependencyNameType) {
	dic.DeleteDependencyWithGroup("", dependencyName)
}

func (dic *DependencyContainer) DeleteDependencyWithGroup(groupName DependencyGroupNameType, dependencyName DependencyNameType) {
	delete(dic.depContainer, string(groupName)+string(dependencyName))
	delete(dic.builderContainer, string(groupName)+string(dependencyName))
}

func (dic *DependencyContainer) DeleteAll() {
	dic.builderContainer = map[string]func() interface{}{}
	dic.depContainer = map[string]interface{}{}
}
