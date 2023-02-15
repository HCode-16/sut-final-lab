package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestNegativeEmployee(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("Check negative test", func(t *testing.T) {
		e := Employee{
			Name:       "hot",
			Email:      "hot@mail.com",
			EmployeeID: "LL12345678",
		}

		ok, err := govalidator.ValidateStruct(e)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("EmployeeID not true form"))
	})
}
