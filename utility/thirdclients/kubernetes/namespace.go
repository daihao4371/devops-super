package kubernetes

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func (cli *Client) GetNamespaces() ([]string, error) {
	list, err := cli.CoreV1().Namespaces().List(cli.Ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var nsArr []string

	for _, nsItem := range list.Items {
		nsArr = append(nsArr, nsItem.Name)
	}

	return nsArr, nil
}

func (cli *Client) GetPersistentVolumeClaims(namespace string) ([]string, error) {
	list, err := cli.CoreV1().PersistentVolumeClaims(namespace).List(cli.Ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var arr []string

	for _, item := range list.Items {
		arr = append(arr, item.Name)
	}

	return arr, nil
}
