package usecases_test

import (
	"github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	mock_usecases "warehouse-service/mocks/usecases"
	"warehouse-service/pkg/drivers"
)

var _ = ginkgo.BeforeSuite(func() {
	mockWarehouseRepository := mock_usecases.NewMockWarehouseRepository(mockCtrl)

	r = drivers.SetupRouter(mockWarehouseRepository)
	// Listen on 8080
	go func() {
		defer GinkgoRecover()
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			Expect(err).ToNot(HaveOccurred())
		}
	}()
})
