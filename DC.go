package dc

type dependencyContainer struct {
	builderContainer map[string]func() interface{}
	depContainer     map[string]interface{}
}

type Container interface {
	SetDependency(dependencyName string, depFunc func() interface{})
	SetDependencyWithGroup(groupName string, dependencyName string, depFunc func() interface{})
	GetSingletonDependency(dependencyName string) interface{}
	GetSingletonDependencyWithGroup(groupName string, dependencyName string) interface{}
	GetDependency(dependencyName string) interface{}
	GetDependencyWithGroup(groupName string, dependencyName string) interface{}
	DeleteDependency(dependencyName string)
	DeleteDependencyWithGroup(groupName string, dependencyName string)
	DeleteAll()
}

func NewDC() Container {
	return &dependencyContainer{
		builderContainer: map[string]func() interface{}{},
		depContainer:     map[string]interface{}{},
	}
}

func (dic *dependencyContainer) SetDependency(dependencyName string, depFunc func() interface{}) {
	dic.SetDependencyWithGroup("", dependencyName, depFunc)
}

func (dic *dependencyContainer) SetDependencyWithGroup(groupName string, dependencyName string, depFunc func() interface{}) {
	dic.builderContainer[groupName+dependencyName] = depFunc
}

func (dic *dependencyContainer) GetSingletonDependency(dependencyName string) interface{} {
	return dic.GetSingletonDependencyWithGroup("", dependencyName)
}

func (dic *dependencyContainer) GetSingletonDependencyWithGroup(groupName string, dependencyName string) interface{} {
	depData, _ := dic.depContainer[groupName+dependencyName]
	if depData == nil {
		builderFunc, _ := dic.builderContainer[groupName+dependencyName]
		depData = builderFunc()
		dic.depContainer[groupName+dependencyName] = depData
	}
	return depData
}

func (dic *dependencyContainer) GetDependency(dependencyName string) interface{} {
	return dic.GetDependencyWithGroup("", dependencyName)
}

func (dic *dependencyContainer) GetDependencyWithGroup(groupName string, dependencyName string) interface{} {
	depFunc, _ := dic.builderContainer[groupName+dependencyName]
	return depFunc()
}

func (dic *dependencyContainer) DeleteDependency(dependencyName string) {
	dic.DeleteDependencyWithGroup("", dependencyName)
}

func (dic *dependencyContainer) DeleteDependencyWithGroup(groupName string, dependencyName string) {
	delete(dic.depContainer, groupName+dependencyName)
	delete(dic.builderContainer, groupName+dependencyName)
}

func (dic *dependencyContainer) DeleteAll() {
	dic.builderContainer = nil
	dic.depContainer = nil
}
