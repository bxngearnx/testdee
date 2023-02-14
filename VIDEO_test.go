package testdee

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type VIDEO struct {
	gorm.Model

	Name string `valid:"required~Name cannot be blank"`
	Url  string `gorm:"UniqueIndex" valid:"url"`
}

func TestNameVideo(t *testing.T) {
	g := NewGomegaWithT(t)

	u := VIDEO{
		Name: "", //ผิด
		Url:  "https://www.youtube.com/",
	}

	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name cannot be blank"))
}
func TestUrlVideo(t *testing.T) {
	g := NewGomegaWithT(t)

	u := VIDEO{
		Name: "ABC",
		Url:  "youtube", //ผิด
	}

	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Url: youtube does not validate as url"))
}
func TestValidVideo(t *testing.T) {
	g := NewGomegaWithT(t)

	u := VIDEO{
		Name: "ABC",
		Url:  "https://www.youtube.com/",
	}

	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
	fmt.Println("LOL")
}
