package formatter

import (
	"github.com/sandriansyafridev/golang/api/book/model/entity"
	"github.com/sandriansyafridev/golang/api/book/model/response"
)

func ToMemberResponse(member entity.Member) (memberResponse response.MemberResponse) {
	memberResponse.ID = member.ID
	memberResponse.Name = member.Name
	memberResponse.CreatedAt = member.CreatedAt
	memberResponse.UpdatedAt = member.UpdatedAt
	return
}

func ToMembersResponse(members []entity.Member) (membersResponse []response.MemberResponse) {

	for _, member := range members {
		membersResponse = append(membersResponse, ToMemberResponse(member))
	}

	return

}
