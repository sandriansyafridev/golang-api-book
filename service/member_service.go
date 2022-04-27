package service

import (
	"github.com/sandriansyafridev/golang/api/book/model/dto"
	"github.com/sandriansyafridev/golang/api/book/model/entity"
	"github.com/sandriansyafridev/golang/api/book/model/formatter"
	"github.com/sandriansyafridev/golang/api/book/model/response"
	"github.com/sandriansyafridev/golang/api/book/repository"
)

type MemberService interface {
	FindAll() (membersResponse []response.MemberResponse, err error)
	FindByID(MemberID int) (memberResponse response.MemberResponse, err error)
	Delete(MemberID int) (err error)
	Create(request dto.MemberCreateDTO) (memberResponse response.MemberResponse, err error)
	Update(request dto.MemberUpdateDTO) (memberResponse response.MemberResponse, err error)
}

type MemberServiceImpl struct {
	repository.MemberRepository
}

func (memberService *MemberServiceImpl) FindAll() (membersResponse []response.MemberResponse, err error) {
	if members, err := memberService.MemberRepository.FindAll(); err != nil {
		return nil, err
	} else {
		return formatter.ToMembersResponse(members), nil
	}
}

func (memberService *MemberServiceImpl) FindByID(MemberID int) (memberResponse response.MemberResponse, err error) {
	if member, err := memberService.MemberRepository.FindByID(MemberID); err != nil {
		return memberResponse, err
	} else {
		return formatter.ToMemberResponse(member), nil
	}
}

func (memberService *MemberServiceImpl) Delete(MemberID int) (err error) {

	if member, err := memberService.MemberRepository.FindByID(MemberID); err != nil {
		return err
	} else {
		if err := memberService.MemberRepository.Delete(member); err != nil {
			return err
		} else {
			return nil
		}
	}

}

func (memberService *MemberServiceImpl) Create(request dto.MemberCreateDTO) (memberResponse response.MemberResponse, err error) {

	member := entity.Member{}
	member.Name = request.Name
	if memberCreated, err := memberService.MemberRepository.Create(member); err != nil {
		return memberResponse, err
	} else {
		return formatter.ToMemberResponse(memberCreated), nil
	}

}

func (memberService *MemberServiceImpl) Update(request dto.MemberUpdateDTO) (memberResponse response.MemberResponse, err error) {

	if member, err := memberService.MemberRepository.FindByID(request.ID); err != nil {
		return memberResponse, err
	} else {
		member.Name = request.Name
		if memberUpdated, err := memberService.MemberRepository.Update(member); err != nil {
			return memberResponse, err
		} else {
			return formatter.ToMemberResponse(memberUpdated), nil
		}
	}

}

func NewMemberServiceImpl(memberRepository repository.MemberRepository) *MemberServiceImpl {
	return &MemberServiceImpl{
		MemberRepository: memberRepository,
	}
}
