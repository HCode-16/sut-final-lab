package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"required~Name can not Blank"`
	Email      string
	EmployeeID string `valid:"matches(^[MJS]\\d{8}$)~EmployeeID not true form"`
}

func TestEmployeeValidation(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("Check possitive test", func(t *testing.T) {
		e := Employee{
			Name:       "nat",
			Email:      "hot@mail.com",
			EmployeeID: "M12345678",
		}

		ok, err := govalidator.ValidateStruct(e)
		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())
	})
}
