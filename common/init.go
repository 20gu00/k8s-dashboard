package common

//初始化clientset promethues istio 配置文件
func Init() error {
	var err error

	//初始化kubeclientset
	if err = kubeClient(); err != nil {
		return err
	}

	return nil

}
