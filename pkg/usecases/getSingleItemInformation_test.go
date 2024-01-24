package usecases_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	mock_usecases "warehouse-service/mocks/usecases"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"warehouse-service/pkg/entities"
)

const (
	getSingleItemInformationURL = "http://localhost:8080/warehouse-service/v1/items/%s"
)

var (
	mockWarehouseRepository *mock_usecases.MockWarehouseRepository
	ctxType                 = reflect.TypeOf((*context.Context)(nil)).Elem()
	mockCtrl                *gomock.Controller
	r                       *gin.Engine
	w                       *httptest.ResponseRecorder
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var _ = Describe("getSingleItemInformation", func() {
	var (
		barcode      string
		mockItem     *entities.ItemInformation
		returnError  error
		returnItem   *entities.ItemInformation
		requestTimes int
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockWarehouseRepository = mock_usecases.NewMockWarehouseRepository(mockCtrl)
		requestTimes = 1
		returnError = nil
		barcode = "123456789"
		returnItem = &entities.ItemInformation{
			BarcodePrefix: barcode,
			ItemName:      "test item",
			Description:   "test description",
			LocationName:  "test location",
			Quantity:      1,
		}
	})
	JustBeforeEach(func() {
		w = httptest.NewRecorder()
		requestBodyJSON, err := json.Marshal(returnItem)
		Expect(err).ToNot(HaveOccurred())

		req, err := http.NewRequest("GET", fmt.Sprintf(getSingleItemInformationURL, barcode), bytes.NewReader(requestBodyJSON))
		Expect(err).ToNot(HaveOccurred())
		mockWarehouseRepository.EXPECT().GetItemInformation(ctxType, barcode).Return(returnItem, returnError).Times(requestTimes)

		r.ServeHTTP(w, req)
	})
	Context("when the item exists", func() {
		It("should return the item information without error", func() {
			Expect(returnError).NotTo(HaveOccurred())
			Expect(returnItem).To(Equal(mockItem))
		})
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
})
