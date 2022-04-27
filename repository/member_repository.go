package repository

import (
	"github.com/sandriansyafridev/golang/api/book/model/entity"
	"gorm.io/gorm"
)

type MemberRepository interface {
	FindAll() (members []entity.Member, err error)
	FindByID(MemberID int) (member entity.Member, err error)
	Delete(member entity.Member) (err error)
	Create(member entity.Member) (entity.Member, error)
	Update(member entity.Member) (entity.Member, error)
}

type MemberRepositoryImpl struct {
	DB *gorm.DB
}

func (memberRepository *MemberRepositoryImpl) FindAll() (members []entity.Member, err error) {

	if err := memberRepository.DB.Find(&members).Error; err != nil {
		return nil, err
	} else {
		return members, nil
	}

}
func (memberRepository *MemberRepositoryImpl) FindByID(MemberID int) (member entity.Member, err error) {

	if err := memberRepository.DB.First(&member, MemberID).Error; err != nil {
		return member, err
	} else {
		return member, nil
	}
}
func (memberRepository *MemberRepositoryImpl) Delete(member entity.Member) (err error) {

	if err := memberRepository.DB.Delete(&member).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (memberRepository *MemberRepositoryImpl) Create(member entity.Member) (entity.Member, error) {
	if err := memberRepository.DB.Create(&member).Error; err != nil {
		return member, err
	} else {
		return member, nil
	}
}

func (memberRepository *MemberRepositoryImpl) Update(member entity.Member) (entity.Member, error) {
	if err := memberRepository.DB.Save(&member).Error; err != nil {
		return member, err
	} else {
		return member, nil
	}
}

func NewMemberRepositoryImpl(db *gorm.DB) *MemberRepositoryImpl {
	return &MemberRepositoryImpl{
		DB: db,
	}
}
