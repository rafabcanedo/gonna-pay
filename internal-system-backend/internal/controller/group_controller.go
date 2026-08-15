package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/configuration/validation"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/service"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/view/request"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/view/response"
)

type GroupController struct {
	service service.GroupService
}

func NewGroupController(service service.GroupService) *GroupController {
	return &GroupController{service: service}
}

func (gc *GroupController) CreateGroup(c *gin.Context) {
	ownerID := c.GetString("userID")

	var req request.CreateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := domains.NewGroupDomain(ownerID, req.Name, req.Category)

	created, err := gc.service.Create(domain, req.MemberIDs)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.NewGroupResponse(created))
}

func (gc *GroupController) FindAllGroups(c *gin.Context) {
	ownerID := c.GetString("userID")

	groups, err := gc.service.FindAll(ownerID)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	groupsResponse := response.NewGroupResponseList(groups)
	c.JSON(http.StatusOK, gin.H{
		"groups": groupsResponse,
		"total":  len(groupsResponse),
	})
}

func (gc *GroupController) FindGroupByID(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("userID")

	group, err := gc.service.FindByID(id, ownerID)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewGroupDetailResponse(group))
}

func (gc *GroupController) DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("userID")

	if err := gc.service.Delete(id, ownerID); err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "group deleted successfully"})
}

func (gc *GroupController) AddMember(c *gin.Context) {
	groupID := c.Param("id")
	ownerID := c.GetString("userID")

	var req request.AddMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	if err := gc.service.AddMember(groupID, req.ContactID, ownerID); err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "member added successfully"})
}

func (gc *GroupController) RemoveMember(c *gin.Context) {
	groupID := c.Param("id")
	contactID := c.Param("contactId")
	ownerID := c.GetString("userID")

	if err := gc.service.RemoveMember(groupID, contactID, ownerID); err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "member removed successfully"})
}
