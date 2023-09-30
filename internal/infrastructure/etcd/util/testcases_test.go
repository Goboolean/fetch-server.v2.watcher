package etcdutil_test

import etcdutil "github.com/Goboolean/fetch-system.master/internal/infrastructure/etcd/util"


type Worker struct {
	ID       string `etcd:"id"`
	Platform string `etcd:"platform"`
	Status   string `etcd:"status"`
}

func (w *Worker) Name() string {
	return "worker"
}


type Product struct {
	ID       string `etcd:"id"`
	Platform string `etcd:"platform"`
	Symbol   string `etcd:"symbol"`
	Worker 	 string `etcd:"worker"`
	Status   string `etcd:"status"`
}

func (p *Product) Name() string {
	return "product"
}


// Current version does not provide Nested struct, only flat struct is supported.
type Nested struct {
	ID string `etcd:"id"`
	Detail struct{
		Name string `etcd:"name"`
		Age int `etcd:"age"`
	} `etcd:"detail"`
}

func (n *Nested) Name() string {
	return "nested"
}



var cases []struct {
	name string
	str map[string]string
	model etcdutil.Model
	data etcdutil.Model
} = []struct{
	name string
	str map[string]string
	model etcdutil.Model
	data etcdutil.Model
}{
	{
		name: "Worker",
		str: map[string]string{
			"/worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d": "",
			"/worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/platform": "kis",
			"/worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/status": "active",
		},
		model: &Worker{},
		data: &Worker{
			ID: "9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d",
			Platform: "kis",
			Status: "active",
		},
	},
	{
		name: "Product",
		str: map[string]string{
			"/product/test.goboolean.kor": "",
			"/product/test.goboolean.kor/platform": "kis",
			"/product/test.goboolean.kor/symbol": "goboolean",
			"/product/test.goboolean.kor/worker": "9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d",
			"/product/test.goboolean.kor/status": "onsubscribe",
		},
		model: &Product{},
		data: &Product{
			ID: "test.goboolean.kor",
			Platform: "kis",
			Symbol: "goboolean",
			Worker: "9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d",
			Status: "onsubscribe",
		},
	},
	/*
	{
		name: "Nested Struct",
		str: map[string]string{
			"nested/mulmuri.dev": "",
			"nested/mulmuri.dev/detail": "",
			"nested/mulmuri.dev/detail/name": "goboolean",
			"nested/mulmuri.dev/detail/age": "1",
		},
		model: &Nested{},
		data: &Nested{
			ID: "mulmuri.dev",
			Detail: struct{
				Name string `etcd:"name"`
				Age int `etcd:"age"`
			} {
				Name: "goboolean",
				Age: 1,
			},
		},
	},
	*/
}